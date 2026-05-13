package main

import (
	"github.com/spf13/cobra"
)

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
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Select user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	userCmd.AddCommand(CreateUser())
	return userCmd
}

func CreateUser() *cobra.Command {
	createUserCmd := &cobra.Command{
		Use:   "create",
		Short: "Create User.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return createUserCmd

}

func Add() *cobra.Command {
	//Add: Goal, Calories, Weight, or Food item
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add Goal, FoodItem or weight.",
		Long:  "Add Goal, FoodItem or weight to User's tracking history.",
	}
	addCmd.AddCommand(AddGoal())
	addCmd.AddCommand(AddFoodItem())
	addCmd.AddCommand(AddWeigth())
	return addCmd
}

func AddGoal() *cobra.Command {
	addGoalCmd := &cobra.Command{
		Use:   "goal",
		Short: "Add weight goal.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return addGoalCmd
}

func AddFoodItem() *cobra.Command {
	var meal string
	addFoodCmd := &cobra.Command{
		Use:   "food",
		Short: "Add Food Item.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	addFoodCmd.Flags().StringVar(&meal, "meal", "breakfast", "Meal of the day(breakfast, Lunch, Dinner)")
	return addFoodCmd
}

func AddWeigth() *cobra.Command {
	addWeigthCmd := &cobra.Command{
		Use:   "weight",
		Short: "Add weight",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return addWeigthCmd
}

func Edit() *cobra.Command {
	//Edit: Goal or Food Item
	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit Goal or Food Item.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	editCmd.AddCommand(EditGoal())
	editCmd.AddCommand(EditFoodItem())
	return editCmd
}

func EditGoal() *cobra.Command {
	editGoalCmd := &cobra.Command{
		Use:   "goal",
		Short: "Edit Goal.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return editGoalCmd
}

func EditFoodItem() *cobra.Command {
	editFoodItem := &cobra.Command{
		Use:   "food",
		Short: "Edit food item.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return editFoodItem
}

func View() *cobra.Command {
	//View: Remaining calories in the day
	// Weight Progress (Plot)
	viewCmd := &cobra.Command{
		Use:   "add",
		Short: "Add Goal, FoodItem or weight.",
		Long:  "Add Goal, FoodItem or weight to User's tracking history.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	viewCmd.AddCommand(ViewLeftoverCalories())
	viewCmd.AddCommand(ViewPlot())
	return viewCmd
}

func ViewLeftoverCalories() *cobra.Command {
	viewCaloriesCmd := &cobra.Command{
		Use:   "calories",
		Short: "Remaining calories for the day/week.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	//default: today. Monday, Tuesday, etc.
	//default: current. Last (week)
	viewCaloriesCmd.Flags().String("day", "today", "Show remaining calories the day.")
	viewCaloriesCmd.Flags().String("week", "current", "Show remaining calories for current week.")
	return viewCaloriesCmd
}

func ViewPlot() *cobra.Command {
	viewPlotCmd := &cobra.Command{
		Use:   "plot",
		Short: "Show plot",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	// week. month, year
	viewPlotCmd.Flags().String("weigth", "month", "Plot weight for time period.")
	return viewPlotCmd
}
