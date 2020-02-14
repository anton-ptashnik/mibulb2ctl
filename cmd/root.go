/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/iAutomator/mibulb2"

	"github.com/spf13/viper"
)

var cfgFile string
var cfgFilePath = path.Join(os.Getenv("userprofile"), ".bulbctl.env")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bulbctl",
	Short: "CLI for controlling MiBulb 2",
	Long: "Top level command usage displays state of the corresponding property." +
		"Property is set by providing the value corresponding to the property in question",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() == discoverCmd.Name() {
			return
		}

		viper.SetConfigFile(cfgFilePath)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("No bulbs were setup to be managed. Use '%s' to command to setup\n", discoverCmd.Name())
			os.Exit(1)
		}

		bulbSummary := mibulb2.BulbSummary{}
		viper.Unmarshal(&bulbSummary)
		bulb = mibulb2.Bulb{bulbSummary}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var bulb mibulb2.Bulb
