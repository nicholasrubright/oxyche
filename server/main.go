package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type ProxyCache struct {
	httpCode      int
	contentLength int64
	contentType   string
	body          []byte
}

var (
	cache map[string]ProxyCache
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

func setupRouter(origin string) *gin.Engine {
	r := gin.Default()

	r.Any("/*any", func(c *gin.Context) {

		path := c.Param("any")
		method := c.Request.Method
		if strings.HasPrefix(path, "/cache") && method == http.MethodDelete {
			clear(cache)
			fmt.Println("Cache CLEARED")
			c.Status(http.StatusOK)
			return
		}

		origin_url, err := url.JoinPath(origin, c.Request.URL.Path)

		if err != nil {
			panic(err)
		}

		val, ok := cache[origin_url]
		if ok {
			fmt.Println("Cache HIT")
			cache_reader := bytes.NewReader(val.body)
			c.DataFromReader(val.httpCode, val.contentLength, val.contentType, cache_reader, nil)
			return
		}

		response, err := http.Get(origin_url)
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		bodyBytes, err := io.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		reader := bytes.NewReader(bodyBytes)
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		cache[origin_url] = ProxyCache{
			httpCode:      http.StatusOK,
			contentLength: contentLength,
			contentType:   contentType,
			body:          bodyBytes,
		}

		fmt.Println("Cache MISS")

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)

	})

	return r
}

func main() {
	cfg := getConfig()

	cache = make(map[string]ProxyCache)
	r := setupRouter(cfg.Server.Origin)
	r.Run(":" + cfg.Server.Port)
}
