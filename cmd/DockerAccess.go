package cmd

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
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
