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

	entity "github.com/homework/agenda/entity"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login -u [Username] -p [Password]",
	Short: "user can log in agenda here.",
	Long:  `user can log in agenda here by inputing their username and password.:`,

	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		entity.GetExsitingUsers()
		if !entity.QueryUserByNameAndPassword(username, password) {
			fmt.Println("user does not exist, please check your username and password.")
			return
		}
		fmt.Println("Hi~ " + username + ", welcome to agenda!")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "new user's username")
	loginCmd.Flags().StringP("password", "p", "", "new user's password")
}
