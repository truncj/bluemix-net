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

var component, count string

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Scale the components of an application to specific quantity",
	Run: func(cmd *cobra.Command, args []string) {
		var scaleURL string

		viper.SetConfigName("app")
		configPath := getPath()
		viper.AddConfigPath(*configPath)

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Config file not found...")
		} else {
			scaleURL = viper.GetString("urls.scale")
		}

		req, err := http.NewRequest("POST", scaleURL+"/"+alias+"/"+version+"/"+component+"/"+count, nil)
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
	appCmd.AddCommand(scaleCmd)
	scaleCmd.Flags().StringVarP(&alias, "alias", "a", "", "Application alias")
	scaleCmd.Flags().StringVarP(&version, "version", "v", "", "Version alias")
	scaleCmd.Flags().StringVarP(&component, "component", "c", "", "Component alias")
	scaleCmd.Flags().StringVarP(&count, "quantity", "q", "", "Requested component quantity")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scaleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scaleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
