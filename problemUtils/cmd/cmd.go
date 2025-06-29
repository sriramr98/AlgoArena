package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sriramr98/dsa_problem_utils/problems"
)

func AddProblemCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [PROBLEM_LIST_META_PATH] [PROBLEMS_PATH]",
		Short: "Add a list new of new problems",
		Long:  "Add a list of new problems based on the provided metadata",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Please provide the path to the problem list metadata file and the problems path.")
				return
			}
			problemListMetaPath := args[0]
			problemsPath := args[1]

			if problemListMetaPath == "" {
				fmt.Println("Problem list metadata path cannot be empty.")
				fmt.Println("Problems list metadata should have the format specified in problem_creation_schema.json")
				return
			}

			if problemsPath == "" {
				fmt.Println("Problems path cannot be empty.")
				return
			}

			err := problems.AddProblems(problemListMetaPath, problemsPath)
			if err != nil {
				fmt.Printf("Error adding problems: %v\n", err)
			} else {
				fmt.Println("Problems added successfully.")
			}
		},
	}

	return cmd
}

func EditProblemCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit an existing problem",
		Long:  "Edit an existing problem in the problem set.",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for editing a problem
			println("Editing an existing problem...")
		},
	}

	return cmd
}

func RemoveProblemCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a problem",
		Long:  "Remove a problem from the problem set.",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for removing a problem
			println("Removing a problem...")
		},
	}

	return cmd
}

func ValidateProblemCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate [PROBLEMS_PATH]",
		Short: "Validate a problem",
		Long:  "Validate the structure and content of a problem.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide the path to the problem file.")
				return
			}
			problemPath := args[0]
			problems.ValidateProblem(problemPath)
		},
	}

	return cmd
}
