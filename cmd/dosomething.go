// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/viper"
)

// dosomethingCmd represents the dosomething command
var dosomethingCmd = &cobra.Command{
	Use:   "dosomething",
	Short: "It does something, anything really.",
	Long:  `dosomething subcommand prints the sentence `,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("dosomething called")
	//},
	Run: doSomethingMain,
}

func init() {
	rootCmd.AddCommand(dosomethingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dosomethingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dosomethingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doSomethingMain(cmd *cobra.Command, args []string) {
	value := viper.GetString("number")
	fmt.Printf("Viper test: %s\n", value)
	fmt.Println("dosomething called")
	value := viper.GetString("number")
	fmt.Printf("Viper test: %s\n", value)
	fmt.Printf("Printing Config Tree: %+v\n", cfgTree)
	fmt.Println("")

	for i := 0; i < cfgTree.DoSomething.Repeat; i++ {
		fmt.Println(cfgTree.Number*i, cfgTree.DoSomething.Sentence)
	}
}
