package buildyaml

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type BuildYaml struct {
	Name  string `yaml:"name"`
	Kind  string `yaml:"kind"`
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Name    string   `yaml:"name"`
	Command []string `yaml:"command"`
	Timeout string   `yaml:"timeout"`
}

type BuildYamlSvc struct {
	YamlFile string
}

func NewBuildYamlSvc() *BuildYamlSvc {
	return &BuildYamlSvc{}
}

func (b *BuildYamlSvc) SetYamlFile(yamlFileName string) *BuildYamlSvc {
	b.YamlFile = yamlFileName
	return b
}

func (b *BuildYamlSvc) Parse() (*BuildYaml, error) {
	if b.YamlFile == "" {
		defaultYamlFile, err := b.getDefaultYamlFile()
		if err != nil {
			return nil, err
		}
		b.YamlFile = defaultYamlFile
	}

	yamlContent, err := ioutil.ReadFile(b.YamlFile)
	if err != nil {
		return nil, err
	}

	buildYaml := BuildYaml{}
	err = yaml.Unmarshal(yamlContent, &buildYaml)
	if err != nil {
		return nil, err
	}

	return &buildYaml, nil
}

func (b *BuildYamlSvc) getDefaultYamlFile() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	yamlFile := path + "/build.yaml"
	return yamlFile, nil
}
