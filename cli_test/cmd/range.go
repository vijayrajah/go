// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rangeCmd represents the range command
var (
	rangeCmd = &cobra.Command{
		Use:   "range",
		Short: "prints a range",
		Long:  `Prints rang from 0 to the number specified`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("range called.. The args are ", args, " the end it ", endVar)
			printRange(endVar)
		},
	}
	endVar int
)

func init() {

	rangeCmd.Flags().IntVarP(&endVar, "end", "e", 10, "Specifiy end to print the numbers.. from 0...END")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rangeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rangeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printRange(end int) {

	for i := 0; i <= end; i++ {
		fmt.Println(i)
	}

}
