package pk

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// PwdAndCopy is a function to copy current directory to clipboard
var copyDir = &cobra.Command{
	Use:     "copydir",
	Aliases: []string{"cpdir"},
	Short:   "Copy current directory to clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the current working directory
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
		}

		com := exec.Command("pwd")
		com.Dir = currentDir

		// Run the command and get the output
		output, err := com.Output()
		if err != nil {
			fmt.Println("Error running command:", err)
		}

		// Copy the output to clipboard
		err = clipboard.WriteAll(string(output))
		if err != nil {
			fmt.Println("Error copying to clipboard:", err)
		}
		fmt.Println("Current directory copied to clipboard!")
	},
}

func init() {
	rootCmd.AddCommand(copyDir)
}
