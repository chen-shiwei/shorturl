package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use: "shorturl",
	}
)

func Execute() {

	rootCmd.AddCommand(apiCmd)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("run fail:%s\n", err.Error())
		os.Exit(0)
	}
}
