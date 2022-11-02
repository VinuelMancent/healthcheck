package model

import (
	"os"
	"strconv"
	"strings"
)

type ContainerName struct {
	Directory string `json:"directory" yaml:"directory"`
	Name      string `json:"name" yaml:"name"`
	Index     int    `json:"index" yaml:"index"`
}

func (containerName *ContainerName) String() string {
	builder := strings.Builder{}
	builder.WriteString("/")
	builder.WriteString(strings.ReplaceAll(strings.ToLower(containerName.getLastDirectory(containerName.Directory)), " ", ""))
	builder.WriteString("-")
	builder.WriteString(containerName.Name)
	builder.WriteString("-")
	builder.WriteString(strconv.Itoa(containerName.Index))
	return builder.String()
}

func (containerName *ContainerName) getLastDirectory(fullDirectory string) string {
	//parts := filepath.SplitList(fullDirectory)
	parts := strings.Split(fullDirectory, string(os.PathSeparator))
	return parts[len(parts)-1]
}
