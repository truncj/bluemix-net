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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// unbindCmd represents the unbind command
var unbindCmd = &cobra.Command{
	Use:   "unbind",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var unbindURL string

		viper.SetConfigName("app")
		viper.AddConfigPath("./cmd/config")

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Config file not found...")
		} else {
			unbindURL = viper.GetString("urls.unbind")
		}

		req, err := http.NewRequest("DELETE", unbindURL+"/"+alias+"/"+version, nil)
		resp, _ := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		fmt.Printf(string(body))
	},
}

func init() {
	appCmd.AddCommand(unbindCmd)
	unbindCmd.Flags().StringVarP(&alias, "alias", "a", "", "Application alias")
	unbindCmd.Flags().StringVarP(&version, "version", "v", "", "Version alias")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unbindCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unbindCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
