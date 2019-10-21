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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register -u [Username] -p [Password] -e [Email] -t[PhoneNumber]",
	Short: "a new user start his/her journey here",
	Long:  `To regist, the user must input a username, password, email and phone number.:`,

	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phoneNumber")
		entity.GetExsitingUsers()
		if entity.QueryUserByName(username) {
			fmt.Println("sorry, the name has been used, please try another one.")
			return
		}
		entity.AddUser(entity.User{Username: username, Password: password, Email: email, PhoneNumber: phone})
		entity.SaveUsers()

		fmt.Println("User has successfully registered:")
		fmt.Println("userName :		" + username)
		fmt.Println("password :		" + password)
		fmt.Println("email    :		" + email)
		fmt.Println("phone    :		" + phone)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("username", "u", "", "new user's username")
	registerCmd.Flags().StringP("password", "p", "", "new user's password")
	registerCmd.Flags().StringP("email", "e", "", "new user's email")
	registerCmd.Flags().StringP("phoneNumber", "t", "", "phone new user's phone")
}
