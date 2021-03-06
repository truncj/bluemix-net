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
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string

// bindCmd represents the bind command
var bindCmd = &cobra.Command{
	Use:   "bind",
	Short: "Unbind an application with a Bluemix service (wip)",
	Run: func(cmd *cobra.Command, args []string) {
		var bindURL string

		viper.SetConfigName("app")
		configPath := getPath()
		viper.AddConfigPath(*configPath)

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Config file not found...")
		} else {
			bindURL = viper.GetString("urls.bind")
		}

		req, err := http.NewRequest("POST", bindURL+"/"+alias+"/"+version, nil)
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
	appCmd.AddCommand(bindCmd)
	bindCmd.Flags().StringVarP(&alias, "alias", "a", "", "Application alias")
	bindCmd.Flags().StringVarP(&version, "version", "v", "", "Version alias")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bindCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bindCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
