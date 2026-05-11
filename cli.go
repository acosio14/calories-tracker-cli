package main

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "calorie-tracker",
		Short: "Tracks Calories",
	}
	rootCmd.AddCommand(User())
	rootCmd.AddCommand(Add())
	rootCmd.AddCommand(Edit())
	rootCmd.AddCommand(View())

	return rootCmd
}

func User() *cobra.Command {
	//create
	//select
	userCmd := &cobra.Command{
		Use:   "User",
		Short: "Creater or Select user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	userCmd.Flags().AddFlag()
	return userCmd
}

func Add() *cobra.Command {
	//Add: Goal, Calories, Weight, or Food item
	return nil
}

func Edit() *cobra.Command {
	//Edit: Goal or Food Item
	return nil
}

func View() *cobra.Command {
	//View: Remaining calories in the day
	// Weight Progress (Plot)
	return nil
}
