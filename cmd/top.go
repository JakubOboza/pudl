/*
Copyright © 2022 Jakub OBoza <jakub.oboza@gmail.com>

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

	"github.com/JakubOboza/pudl/pudl"
	"github.com/spf13/cobra"
)

// topCmd represents the top command
var topCmd = &cobra.Command{
	Use:   "top",
	Short: "top shows only few most recent news items from pudelek.pl",
	Long:  `Use top if you want to see only few recent items. By default it is 3`,
	Run: func(cmd *cobra.Command, args []string) {

		enableLogs, _ := rootCmd.Flags().GetBool("logs")
		if enableLogs {
			pudl.EnableLogs()
		}

		err := pudl.Top()
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}
