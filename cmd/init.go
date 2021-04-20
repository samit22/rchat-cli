/*
Copyright © 2021 Samit Ghimire samitghimire@gmail.com>
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
	"net/url"

	"github.com/samit22/rchat/rchatwrapr"
	"github.com/spf13/cobra"
)

var initRchatCliCmd = &cobra.Command{
	Use:   "init",
	Short: "Store server and  login information for rocket chat",
	Long: `Rocket chat requires a domain name.
	Also to make the API request we need to authenticate.
	Please provide the username or email and the password.
	The username and password is encoded and saved in you local machine, you
	can remove it anytime using 'rm ~/.rchat-config'
`,
	Run: func(cmd *cobra.Command, args []string) {

		initRchat(args)
	},
}

func init() {
	rootCmd.AddCommand(initRchatCliCmd)
}

func initRchat(args []string) (err error) {
	fmt.Print("Enter Rocket Chat Domain URL: ")
	var domain, user, password string
	fmt.Scanln(&domain)
	domain, err = validateDomain(domain)
	if err != nil {
		fmt.Println("Invalid domain address please use a valid one should have http or https")
		return
	}
	fmt.Print("Enter Username or email: ")
	fmt.Scanln(&user)
	err = validateEmptyString(user)
	if err != nil {
		fmt.Println("User can not be empty use either username or email")
		return
	}
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)
	err = validateEmptyString(password)
	if err != nil {
		fmt.Println("Password can not be empty")
		return
	}

	fmt.Printf("validating inputs, making request to %s \n", domain)
	c := rchatwrapr.Client{
		DomainURL: domain,
		User:      user,
		Password:  password,
	}
	c.ValidateRequest()
	return
}

func validateDomain(domain string) (parsedDomain string, err error) {
	if domain == "" {
		err = fmt.Errorf("domain can not be empty")
		return
	}
	u, err := url.Parse(domain)
	if err != nil {
		err = fmt.Errorf("invalid domain address, %v", err)
		return
	}
	if u.Host == "" || u.Scheme == "" {
		err = fmt.Errorf("invalid domain address please enter a valid one with the scheme")
		return
	}
	parsedDomain = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	return
}

func validateEmptyString(inp string) (err error) {
	if inp == "" {
		err = fmt.Errorf("empty input value")
	}
	return
}
