package main

import (
	"Healthcheck/cmd"
)

func main() {

	cmd.Execute()
	err, context, client := cmd.ConnectToDocker()
	if err != nil {
		return
	}
	//err, unhealthyContainers := cmd.GetAllUnhealthyContainers(context, client)
	//if err != nil {
	//return
	//}
	//cmd.RestartContainers(context, client, unhealthyContainers)
	//for _, element := range unhealthyContainers {
	//	fmt.Println(element.ID)
	//}
	names := make([]string, 1)
	//names[0] = "test"
	//names[1] = "redis"
	names[0] = "fineCollection"
	err, containers := cmd.GetContainersByName(context, client, names)
	if err != nil {
		return
	}
	cmd.RestartContainers(context, client, containers)
	//service, err := cmd.ReadFileToServiceWithDependencies("C:\\Users\\Vincent\\Documents\\Go\\src\\Healthcheck\\test.yaml")
	//fmt.Println(service)
	/*
		err, context, client := cmd.ConnectToDocker()
		if err != nil {
			return
		}
		cmd.ListenForEvents(context, client)*/
}
