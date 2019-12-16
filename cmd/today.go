/*
Copyright © 2019 orion0616 earth.nobu.light@gmail.com

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

	"github.com/orion0616/sealion/todoist"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Total time of today's tasks",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := todoist.NewClient()
		if err != nil {
			fmt.Println(err)
			return
		}
		labels := []string{
			"5分",
			"15分",
			"30分",
			"1時間",
		}
		time := []int{
			5, 15, 30, 60,
		}
		labelIDs, err := client.CreateLabelIDs(labels)
		if err != nil {
			fmt.Println(err)
			return
		}
		tasks, err := client.GetAllTasks()
		if err != nil {
			fmt.Println(err)
			return
		}
		sum := 0
		for _, task := range tasks {
			// TODO: filtering
			for _, label := range task.Labels {
				for i, id := range labelIDs {
					if label == id {
						sum += time[i]
					}
				}
			}
		}
		fmt.Printf("Sum: %2dh %2dm\n", sum/60, sum%60)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
