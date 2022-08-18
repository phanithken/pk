package pk

import (
	"fmt"
	"github.com/spf13/cobra"
)

var todoistCmd = &cobra.Command{
	Use: "todoist",
	Aliases: []string{""},
	Short: "A list of command to interact with todoist",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Oup! Empty todoist command")
	},
}

func init() {
	rootCmd.AddCommand(todoistCmd)
}