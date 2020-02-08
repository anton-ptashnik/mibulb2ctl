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
	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "Set power state",
	Run: func(cmd *cobra.Command, args []string) {
		bulb.GetPower()
	},
}

var powerOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Power off the bulb",
	Run: func(cmd *cobra.Command, args []string) {
		bulb.SetPower(false)
	},
}
var powerOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Power on the bulb",
	Run: func(cmd *cobra.Command, args []string) {
		bulb.SetPower(true)
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)
	powerCmd.AddCommand(powerOnCmd)
	powerCmd.AddCommand(powerOffCmd)
}
