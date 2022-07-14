package pk

import (
	"fmt"
	"github.com/phanithken/pk-cli/pkg/pk"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use: "reverse",
	Aliases: []string{""},
	Short: "Reverses a string",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := pk.Reverse(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}