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
	"io/ioutil"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/iAutomator/mibulb2"

	"github.com/spf13/cobra"
)

var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover bulbs for future communication",
	Run: func(cmd *cobra.Command, args []string) {
		si := make(chan bool)
		foundBulb := mibulb2.Search(si)

		foundBulbIDs := make(map[int]bool)
		var foundBulbs []mibulb2.BulbSummary
		endTime := time.Now().Add(time.Second * 9)
		fmt.Println("Discovering...")
		for endTime.After(time.Now()) {
			select {
			case newlyFoundOne := <-foundBulb:
				if _, present := foundBulbIDs[newlyFoundOne.Id]; !present {
					foundBulbIDs[newlyFoundOne.Id] = true
					foundBulbs = append(foundBulbs, newlyFoundOne)
				}
			default:
				time.Sleep(time.Second * 3)
			}
		}
		si <- true
		if len(foundBulbs) == 0 {
			fmt.Println("No bulbs found")
			return
		}

		targetBulb := targetBulbPrompt(foundBulbs)
		err := setupToBeManaged(targetBulb)
		if err == nil {
			fmt.Println("The selected bulb is setup and ready to be managed")
		} else {
			fmt.Println("Error ocurred creating a config file:", err)
		}
	},
}

func targetBulbPrompt(options []mibulb2.BulbSummary) mibulb2.BulbSummary {
	var targetDeviceIndx int
	var optionsSurvey []string
	for _, v := range options {
		optionsSurvey = append(optionsSurvey, fmt.Sprintf("%+v", v))
	}

	prompt := &survey.Select{
		Message: "Select the bulb you'd like to control:",
		Options: optionsSurvey,
	}
	survey.AskOne(prompt, &targetDeviceIndx)
	return options[targetDeviceIndx]
}

func setupToBeManaged(bs mibulb2.BulbSummary) error {
	bsInDotenv := fmt.Sprintf("Id=%v\nIp=%v\nModel=%v", bs.Id, bs.Ip, bs.Model)
	return ioutil.WriteFile(cfgFilePath, []byte(bsInDotenv), 0)
}

func init() {
	rootCmd.AddCommand(discoverCmd)
}
