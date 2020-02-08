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

var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Set timer. Use value in range 1-60 (mins) to set timer",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			bulb.GetTimer()
		} else {
			timeVal, _ := strconv.Atoi(args[0])
			bulb.SetTimer(timeVal)
		}
	},
}
var discardTimerCmd = &cobra.Command{
	Use:   "discard",
	Short: "Discard active timer",
	Run: func(cmd *cobra.Command, args []string) {
		bulb.DiscardTimer()
	},
}

func init() {
	rootCmd.AddCommand(timerCmd)
	timerCmd.AddCommand(discardTimerCmd)
}
