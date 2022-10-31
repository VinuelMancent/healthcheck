package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ListenForEvents(ctx context.Context, client *client.Client) error {
	eventMessages, err := client.Events(ctx, types.EventsOptions{})
	if err != nil {
		return nil
	}
	fmt.Println(eventMessages)
	return nil
}
