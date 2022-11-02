package model

type Service struct {
	Name string `json:"name" yaml:"name"`
}
type ServiceWithDependencies struct {
	Name              string    `json:"name" yaml:"name"`
	DependentServices []Service `json:"dependentServices" yaml:"dependentServices"`
}
