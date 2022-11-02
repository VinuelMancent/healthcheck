package cmd

import (
	"Healthcheck/model"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
	"time"
)

func ConnectToDocker() (error, context.Context, *client.Client) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err, nil, nil
	}
	return nil, ctx, cli
}

func StartContainer(ctx context.Context, client *client.Client, containerID string) error {
	if err := client.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}

func GetAllContainers(ctx context.Context, client *client.Client) (error, []types.Container) {
	containers, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err, nil
	}
	return nil, containers
}

func GetAllUnhealthyContainers(ctx context.Context, client *client.Client) (error, []types.Container) {
	err, allContainers := GetAllContainers(ctx, client)
	unhealthyContainers := make([]types.Container, 0)
	if err != nil {
		return err, nil
	}
	for _, element := range allContainers {
		if strings.Contains(element.Status, "(unhealthy)") {
			unhealthyContainers = append(unhealthyContainers, element)
		}
	}
	return nil, unhealthyContainers
}

func GetContainersByName(ctx context.Context, client *client.Client, containerNames []string) (error, []types.Container) {
	pathToComposeYaml := "C:\\Users\\Vincent\\Documents\\Uni\\Thesis\\Thesis\\DockerComposeDaprPoc\\WithDaprWithoutDurable - LB"
	err, containers := GetAllContainers(ctx, client)
	if err != nil {
		return err, nil
	}
	containersToReturn := make([]types.Container, 0)
	//iterate over each container to check its names
	for _, container := range containers {
		//iterate over each name of the container
		for _, name := range container.Names {
			//iterate over each name of the containernames to compare with current container current name
			for _, nameToCompare := range containerNames {
				//full container name = /+pathToComposeYaml+name+index (alles kleingeschrieben)
				fullContainerName := model.ContainerName{
					Name:      nameToCompare,
					Directory: pathToComposeYaml,
					Index:     1,
				}
				toCompare := fullContainerName.String()
				if name == toCompare {
					containersToReturn = append(containersToReturn, container)
				}
			}
		}
	}
	return nil, containersToReturn
}

func RestartContainers(ctx context.Context, client *client.Client, containers []types.Container) error {
	duration, err := time.ParseDuration("250ms")
	if err != nil {
		return err
	}
	for _, element := range containers {
		client.ContainerRestart(ctx, element.ID, &duration)
	}
	return nil
}
