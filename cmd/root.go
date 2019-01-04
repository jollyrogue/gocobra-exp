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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigTreeMain struct {
	Servers []string
	Secure  bool
	Domain  string
}

type ConfigTreeDoSomething struct {
	Repeat   int
	Sentence string
}

type ConfigTree struct {
	Number      int
	Main        ConfigTreeMain
	DoSomething ConfigTreeDoSomething
}

var cfgFile string
var cfgTree ConfigTree

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocobra-exp",
	Short: "Experimenting with Cobra.",
	Long: `Experimenting with building Go applications with Cobra.
This is for me to work out problems in an simple, controlled environment.

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
	Run: mainConfigPrint,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.gocobra-exp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		err := viper.Unmarshal(&cfgTree)
		if err != nil {
			fmt.Printf("Unable to decode config into struct, %v\n", err)
		}
	}

}

func mainConfigPrint(cmd *cobra.Command, args []string) {
	fmt.Println("Printing config tree.")
	fmt.Printf("Full Config Tree:\n%+v\n", cfgTree)
	fmt.Println("")
	fmt.Printf("root.Number: %d\n", cfgTree.Number)
	fmt.Printf("root.Main.Servers: %s\n", cfgTree.Main.Servers)
	for i, v := range cfgTree.Main.Servers {
		fmt.Printf("\t%d: %s\n", i, v)
	}
	fmt.Printf("root.Main.Secure: %t\n", cfgTree.Main.Secure)
	fmt.Printf("root.Main.Domain: %s\n", cfgTree.Main.Domain)
	fmt.Printf("root.DoSomething.Repeat: %d\n", cfgTree.DoSomething.Repeat)
	fmt.Printf("root.DoSomething.Sentence: %s\n", cfgTree.DoSomething.Sentence)
}
