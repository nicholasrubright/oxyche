package internal

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"gopkg.in/yaml.v3"
)

const (
	containerName = "oxyche-server"
	imageName     = "oxyche-server:latest"
	hostPort      = "8000"
	containerPort = "8080"
)

func BuildImage() error {
	serverDir := filepath.Join(".", "server")

	buildCmd := exec.Command("docker", "build", "-t", imageName, serverDir)

	fmt.Printf("Building docker image %s...\n", imageName)

	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)
	}

	fmt.Printf("Successfully built Docker image %s\n", imageName)

	return nil

}

// func checkContainerExist() error {
// 	ctx := context.Background()

// 	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

// 	if err != nil {
// 		return fmt.Errorf("Unable to connect to docker client")
// 	}

// 	_, err = client.ContainerAPIClient.ContainerInspect(ctx, containerName)

// 	if err != nil {
// 		return fmt.Errorf("container %s already exists. Stop and remove it first", containerName)
// 	}

// 	return nil
// }

func CreateAndStartContainer() error {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	containerConfig := &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			nat.Port(containerPort + "/tcp"): struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(containerPort + "/tcp"): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: hostPort,
				},
			},
		},
	}

	cont, err := cli.ContainerCreate(
		ctx, containerConfig,
		hostConfig, nil, nil, containerName)

	if err != nil {
		panic(err)
	}

	cli.ContainerStart(ctx, cont.ID, container.StartOptions{})
	fmt.Printf("Container %s is started", cont.ID)

	return nil

}

func StopContainer() error {

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	timeout := 10
	if err := cli.ContainerStop(ctx, containerName, container.StopOptions{Timeout: &timeout}); err != nil {
		return fmt.Errorf("failed to stop container: %v", err)
	}

	if err := cli.ContainerRemove(ctx, containerName, container.RemoveOptions{}); err != nil {
		return fmt.Errorf("failed to remove container: %v", err)
	}

	fmt.Printf("Container %s stopped and removed successfully\n", containerName)
	return nil

}

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port   int    `yaml:"port"`
	Origin string `yaml:"origin"`
}

func UpdateConfig(origin string, port int) error {

	config := Config{
		Server: Server{
			Port:   port,
			Origin: origin,
		},
	}

	data, err := yaml.Marshal(&config)

	if err != nil {
		panic(err)
	}

	config_yaml := filepath.Join(".", "server", "config.yml")

	err = os.WriteFile(config_yaml, data, 0644)
	if err != nil {
		panic(err)
	}

	return nil

}
