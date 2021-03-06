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
	"text/template"
	"os"

	"github.com/spf13/cobra"
	"github.com/satori/go.uuid"
)

type UUID struct {
	Uuid uuid.UUID
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: CreateInitFile,
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func CreateInitFile(*cobra.Command, []string) {
	gopath := os.Getenv("GOPATH")
	current, err  := os.Getwd()

	if err != nil {
		panic(err)
	}

	//template に変数を適用
	tpl := template.Must(template.ParseFiles(gopath + "/src/github.com/kazuminn/xhyve-gomana/template/sample.tpl"))

	//create uuid
	member := UUID{uuid.NewV4()} 

	//setting output file
	file, err := os.Create(current + "/Config")
	defer file.Close()

	if err != nil {
		panic(err)
        }
	

	//output file
	err = tpl.Execute(file, member)

	if err != nil {
		panic(err)
    	}
}
