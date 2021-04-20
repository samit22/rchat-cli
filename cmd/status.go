/*
Copyright © 2021 Samit Ghimire <samitghimire@gamil.com

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
	"strings"

	"github.com/samit22/rchat/rchatwrapr"
	"github.com/spf13/cobra"
)

var status string
var validStatus = []string{"online", "away", "busy", "offline"}

type UpdateStatusIface interface {
	SetUserStatus(message, status string) (err error)
}

// statusCmd represents the setting status command
var updStatusCmd = &cobra.Command{
	Use:   "update-status [message]",
	Short: "Update status will change the status of the user",
	Long:  `User status message can be updated with this command, along with the status message the status can also be Updated. Updating status is optional using the flag 'online', 'away', 'busy', 'offline'`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := rchatwrapr.Client{}
		updateUserStatus(args, client)
	},
}

func updateUserStatus(args []string, updater UpdateStatusIface) (err error) {
	var valid bool
	if status != "" {
		for _, st := range validStatus {
			if st == status {
				valid = true
				break
			}
		}
	} else {
		valid = true
	}
	if !valid {
		err = fmt.Errorf("invalid status %s", status)
		fmt.Printf("Invalid status, it can be one of %v \n", strings.Join(validStatus, ", "))
		return
	}
	message := strings.Join(args, " ")

	err = updater.SetUserStatus(message, status)
	if err != nil {
		fmt.Printf("Failed to update status err %v \n", err)
	} else {
		fmt.Println("Successfully updated user status message")
	}
	return
}

func init() {
	updStatusCmd.Flags().StringVarP(&status, "status", "s", "", "Pass the status as flag which can be one of online, away, busy or offline")
}
