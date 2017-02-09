// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// pushdirCmd represents the pushdir command
var pushdirCmd = &cobra.Command{
	Use:   "pushdir",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var dir string

		if dir, err = os.Getwd(); err != nil {
			log.Fatal(err)
		}

		file := findExt(dir, "sln")
		filename := strings.TrimSuffix(file[0], "sln")
		archivename := filename + "zip"
		archiveOutput := path.Join(dir, archivename)
		command := exec.Command("acs", "NewPackage", "-Sln", file[0], "-O", archiveOutput)

		cmdOutput := &bytes.Buffer{}
		command.Stdout = cmdOutput
		err := command.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(cmdOutput.Bytes()))
	},
}

func findExt(dir string, ext string) []string {

	var files []string
	filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}

func init() {
	appCmd.AddCommand(pushdirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushdirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
