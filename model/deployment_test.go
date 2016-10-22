package model

import (
	"io/ioutil"
	"testing"
)

var deployment_yaml = "../testdata/deployment.yaml"

func TestParseDeploymentYAML(t *testing.T) {
	data, err := ioutil.ReadFile(deployment_yaml)
	if err!=nil{
		panic(err)
	}
	var deployment DeploymentYAML
	err=Deployer.Unmarshal(data, &deployment); if err!= nil {
		panic(err)
	}
	var expectedPackagename = "helloworld"
	if deployment.Package.Packagename != expectedPackagename {
		t.Error("get package name failed")
	}

	var expectedDesc = "input person address"
	if deployment.Package.Actions[0].Input[1]["description"] != expectedDesc {
		t.Error("get input failed")
	}

}
