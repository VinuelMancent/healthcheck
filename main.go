package main

import (
	"Healthcheck/cmd"
	"fmt"
)

func main() {

	cmd.Execute()
	err, context, client := cmd.ConnectToDocker()
	if err != nil {
		return
	}
	err, unhealthyContainers := cmd.GetAllUnhealthyContainers(context, client)
	if err != nil {
		return
	}
	cmd.RestartContainers(context, client, unhealthyContainers)
	for _, element := range unhealthyContainers {
		fmt.Println(element.ID)
	}
	/*
		err, context, client := cmd.ConnectToDocker()
		if err != nil {
			return
		}
		cmd.ListenForEvents(context, client)*/
}
