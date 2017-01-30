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
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	spin "github.com/tj/go-spin"
)

var name, alias, desc, archive string
var f *os.File
var fi os.FileInfo
var bar *pb.ProgressBar
var err error

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var appsURL string
		var resp *http.Response
		// var wg sync.WaitGroup
		//
		// wg.Add(2)

		viper.SetConfigName("app")
		configPath := getPath()
		viper.AddConfigPath(*configPath)

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Config file not found...")
		} else {
			appsURL = viper.GetString("urls.apps")
		}

		if f, err = os.Open(archive); err != nil {
			log.Fatal(err)
		}

		if fi, err = f.Stat(); err != nil {
			log.Fatal(err)
		}

		bar = pb.New64(fi.Size()).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
		bar.Start()

		r, w := io.Pipe()
		mpw := multipart.NewWriter(w)

		go func() {
			var part io.Writer
			defer w.Close()
			defer f.Close()

			if err = mpw.WriteField("name", name); err != nil {
				log.Fatal(err)
			}
			if err = mpw.WriteField("alias", alias); err != nil {
				log.Fatal(err)
			}
			if err = mpw.WriteField("desc", desc); err != nil {
				log.Fatal(err)
			}
			//part is assigned the mpw aka the body that contains name/alias/desc & file
			if part, err = mpw.CreateFormFile("file", fi.Name()); err != nil {
				log.Fatal(err)
			}
			//part is a now a writer that duplicates its writes to both part & bar
			part = io.MultiWriter(part, bar)
			//the file f is now written to both part & bar as part of the copy
			if _, err := io.Copy(part, f); err != nil {
				log.Fatal(err)
			}
			if err := mpw.Close(); err != nil {
				log.Fatal(err)
			}
			//wg.Done()
			bar.FinishPrint("Upload Complete")

		}()

		go func() {
			// wg.Wait()
			req, err := http.NewRequest("POST", appsURL+"/"+alias, r)
			req.Header.Set("Content-Type", mpw.FormDataContentType())
			if err != nil {
				log.Fatal(err)
			}
			resp, err = http.DefaultClient.Do(req)

			if err != nil {
				log.Fatal(err)
			}
		}()

		s := spin.New()
		for resp == nil {
			if bar.Get() == bar.Total {
				fmt.Printf("\r  \033[36mProcessing\033[m %s ", s.Next())
				time.Sleep(100 * time.Millisecond)
			}
		}

		defer resp.Body.Close()
		ret, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(ret))
	},
}

func init() {
	pushCmd.Flags().StringVarP(&name, "name", "n", "", "Application name")
	pushCmd.Flags().StringVarP(&alias, "alias", "a", "", "Application alias")
	pushCmd.Flags().StringVarP(&desc, "desc", "d", "", "Application description")
	pushCmd.Flags().StringVar(&archive, "archive", "", "Path to application archive.zip")
	appCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
