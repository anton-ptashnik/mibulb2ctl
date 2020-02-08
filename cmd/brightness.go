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
	"strconv"

	"github.com/spf13/cobra"
)

var brightnessCmd = &cobra.Command{
	Use:   "brightness",
	Short: "Control brightness property. Use value in percent (1-100) to set state",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			v, _ := strconv.Atoi(args[0])
			bulb.SetBrightness(v)
		} else {
			bulb.GetBrightness()
		}
	},
}

func init() {
	rootCmd.AddCommand(brightnessCmd)

	brightnessCmd.Args = cobra.MaximumNArgs(1)
}
