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

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Calculate the all of time of tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := todoist.NewClient()
		if err != nil {
			fmt.Println(err)
			return
		}

		projects, err := client.GetProjects()
		if err != nil {
			fmt.Println(err)
			return
		}
		var routineIDs []int
		routines := []string{
			"Daily Routine",
			"Weekly Routine",
			"Monthly Routine",
			"3Month Routine",
		}
		for _, project := range projects {
			for _, routine := range routines {
				if project.Name == routine {
					routineIDs = append(routineIDs, project.ID)
				}
			}
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
		sumRoutine := 0
		for _, task := range tasks {
			isRoutine := false
			for _, routine := range routineIDs {
				if int64(routine) == task.ProjectID {
					isRoutine = true
				}
			}
			if isRoutine {
				for _, label := range task.Labels {
					for i, id := range labelIDs {
						if label == id {
							sumRoutine += time[i]
						}
					}
				}
				continue
			}
			for _, label := range task.Labels {
				for i, id := range labelIDs {
					if label == id {
						sum += time[i]
					}
				}
			}
		}
		fmt.Printf("Sum    : %2dh %2dm\nRoutine: %2dh %2dm\n",
			sum/60, sum%60, sumRoutine/60, sumRoutine%60)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
