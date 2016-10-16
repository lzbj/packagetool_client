// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"../model"
)

var deploymentfile string


// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy the package based on deployment.yaml file",
	Long:  `deploy the package based on deployment.yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
		fmt.Println("Print: " + deploymentfile)
		var deployment model.Deployment
		data, err:=ioutil.ReadFile(deploymentfile)
		if err!=nil{
			fmt.Errorf("error happened %v",err)
		}
		err=model.Deployer.Unmarshal(data,&deployment)
		if err !=nil{
			fmt.Errorf("error happened %v",err)
		}
		content, err:=model.Deployer.Marshal(&deployment)
                if err!=nil{
			fmt.Errorf("error happened %v",err)
		}

                fmt.Printf("file content is:\n\n %s\n", string(content))

	},
}

func init() {
	deployCmd.Flags().StringVarP(&deploymentfile,"file", "f", "deployment.yaml", "The deployment yaml file to parse, " +
		"default is deployment.yaml in current folder")
	RootCmd.AddCommand(deployCmd)
}
