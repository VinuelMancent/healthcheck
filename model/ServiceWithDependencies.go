package model

type Service struct {
	Name string `json: "name"`
}
type ServiceWithDependencies struct {
	Name              string    `json: "name"`
	DependentServices []Service `json: "dependentServices"`
}
