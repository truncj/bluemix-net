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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an application",
	Run: func(cmd *cobra.Command, args []string) {
		var appsURL string
		var dir string

		viper.SetConfigName("app")
		configPath := getPath()
		viper.AddConfigPath(*configPath)

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Config file not found...")
		} else {
			appsURL = viper.GetString("urls.apps")
		}

		if dir, err = os.Getwd(); err != nil {
			log.Fatal(err)
		}

		if alias == "" {

			file := *findExt(dir, "sln")
			filename := strings.TrimSuffix(file[0], ".sln")
			val := confirm("Are you sure you want to delete "+filename+" from Bluemix?", 3)
			if !val {
				log.Fatal(val)
			}
			//Alias must be abc/0-9
			rp := regexp.MustCompile("[^0-9a-zA-Z]+")
			name = rp.ReplaceAllString(filename, "")
			alias = name

			//Alias must be 20 characters or less
			if len(filename) > 20 {
				alias = alias[:20]
			}
		}

		req, err := http.NewRequest("DELETE", appsURL+"/"+alias, nil)
		resp, _ := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Printf(string(body))
	},
}

func init() {
	appCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&alias, "alias", "a", "", "Application alias")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
