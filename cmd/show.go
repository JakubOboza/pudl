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
	"strconv"

	"github.com/JakubOboza/pudl/pudl"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show me in browser the news i want to see",
	Long: `show will try to open browser directly to link you want to see
	eg. "pudl show 1" will show the 1 cached url. When you are using list or cache you get a list of news items.
	You can use their ids in show command`,
	Run: func(cmd *cobra.Command, args []string) {

		enableLogs, _ := rootCmd.Flags().GetBool("logs")
		if enableLogs {
			pudl.EnableLogs()
		}

		index, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println(err)
			return
		}

		err = pudl.Show(index)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
