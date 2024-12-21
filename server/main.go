package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Origin string `yaml:"origin"`
		Port   string `yaml:"port"`
	}
}

func getConfig() Config {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println("Can't open config: ", err)
	}

	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Problem decoding config: ", err)
	}

	return cfg
}

func fetchRequest(url string) interface{} {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var body interface{}
	err = json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		panic(err)
	}

	return body
}

func setupRouter(origin string) *gin.Engine {
	r := gin.Default()

	r.Any("/*any", func(c *gin.Context) {

		result := fetchRequest(origin + c.Request.URL.Path)
		c.JSON(200, result)
	})

	return r
}

func main() {
	cfg := getConfig()
	r := setupRouter(cfg.Server.Origin)
	r.Run(":" + cfg.Server.Port)
}
