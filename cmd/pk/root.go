package pk

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use: "pk",
	Version: version,
	Short: "pk - a self productivity tool made with ❤️",
	Long: "pk - a self productivity tool made with ❤️",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oupp! There is an error '%s'", err)
		os.Exit(1)
	}
}