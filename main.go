package main

import (
	"fmt"
	"os"

	"mcui/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcui",
	Short: "Memcached UI in Go",
}

func main() {
	rootCmd.AddCommand(cmd.ServeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
