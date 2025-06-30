package main

import (
	"fmt"
	"github.com/sriramr98/dsa_problem_utils/cmd"
	"os"

	"github.com/spf13/cobra"
)

func getProblemCmd() *cobra.Command {
	problemCmd := &cobra.Command{
		Use:   "problem",
		Short: "Manage problems",
		Long:  "A command to manage problems, including adding, editing, removing and validating problems.",
	}

	problemCmd.AddCommand(cmd.AddProblemCmd())
	problemCmd.AddCommand(cmd.EditProblemCmd())
	problemCmd.AddCommand(cmd.RemoveProblemCmd())
	problemCmd.AddCommand(cmd.ValidateProblemCmd())

	return problemCmd
}

func main() {
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the Problem Utils CLI!")
			fmt.Println("Use the --help flag to see available commands.")
		},
	}

	rootCmd.AddCommand(getProblemCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
