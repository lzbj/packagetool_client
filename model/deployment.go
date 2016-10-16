package model

import (
	"gopkg.in/yaml.v2"
	"log"
)

var Deployer *DeploymentManager

func init() {
	Deployer = &DeploymentManager{}
}

type ParseDeploymentYaml interface {
	unmarshal(input []byte, deploy *Deployment) error
	marshal(deployment *Deployment) (output []byte, err error)
}

type DeploymentManager struct {
	deployments []*Deployment
	lastID      uint32
}

type Action struct {
	Name     string                   `yaml:"name"`
	Version  string                   `yaml:"version"`
	Function string                   `yaml:"function"`
	Runtime  string                   `yaml:"runtime"`
	Input    []map[string]interface{} `yaml:"inputs"`
	Output   []map[string]interface{} `yaml:"outputs"`
}

type Deployment struct {
	Packagename string   `yaml:"packagename"`
	Version     string   `yaml:"version"`
	License     string   `yaml:"license"`
	Actions     []Action `yaml:"actions"`
}

func (dm *DeploymentManager) Unmarshal(input []byte, deploy *Deployment) error {
	err := yaml.Unmarshal(input, deploy)
	if err != nil {
		log.Fatalf("error happened during unmarshal :%v", err)
		return err
	}
	return nil
}

func (dm *DeploymentManager) Marshal(deployment *Deployment) (output []byte, err error) {
	data, err := yaml.Marshal(deployment)
	if err != nil {
		log.Fatalf("err happened during marshal :%v", err)
		return nil, err
	}
	return data, nil
}
