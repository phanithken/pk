package pk

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use: "pk",
	Version: version,
	Short: "productivity tool made with ❤️",
	Long: "productivity tool made with ❤️",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: show guide messages
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Oupp! There is an error '%s'", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}