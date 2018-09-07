// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"bufio"
	"log"
	"os"

	"github.com/shoebillk/sbs/client"
	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put a file to sbs server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("put called")

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatal(err)
		}
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err)
		}

		b, err := client.NewBlobServiceClient(host, port)
		if err != nil {
			log.Fatal(err)
		}

		c, err := client.NewClient(b)
		if err != nil {
			log.Fatal(err)
		}

		file := args[0]
		log.Printf("file : %s\n", file)
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		reader := bufio.NewReader(f)

		ID := file
		pushStatus, err := c.Push(ID, reader)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%#v", pushStatus)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	putCmd.Flags().StringP("host", "s", DefaultHost, "Host string of server")
	putCmd.Flags().IntP("port", "p", DefaultPort, "Port number of server")
}
