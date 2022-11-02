package cmd

import (
	"Healthcheck/model"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadFileToServiceWithDependencies(filepath string) (model.ServiceWithDependencies, error) {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return model.ServiceWithDependencies{}, err
	}
	Service := model.ServiceWithDependencies{}
	err = yaml.Unmarshal(dat, &Service)
	if err != nil {
		return model.ServiceWithDependencies{}, err
	}
	return Service, nil
}
