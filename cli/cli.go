package cli

import (
	"fmt"
	"os"

	"github.com/acosio14/calories-tracker-cli/domain"
	"github.com/spf13/cobra"
)

type UserManager interface {
	CreateUser(name string) error
	SelectUser(name string) (*domain.User, error)
}

type GoalManager interface {
	EditCaloriesGoal() error
}

type FoodTracker interface {
	AddFoodItem() error
	EditFoodItem() error
	ViewRemainingCalories() error
}

type WeightTracker interface {
	AddWeight() error
	DisplayWeightProgress() error
}

type CLI struct {
	User   UserManager
	Goal   GoalManager
	Food   FoodTracker
	Weight WeightTracker
}

func NewCLI(User UserManager, Goal GoalManager, Food FoodTracker, Weight WeightTracker) *CLI {
	return &CLI{
		User:   User,
		Goal:   Goal,
		Food:   Food,
		Weight: Weight,
	}
}

func (c *CLI) NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "calorie-tracker",
		Short: "Tracks Calories",
	}
	rootCmd.AddCommand(c.SelectUserCmd())
	rootCmd.AddCommand(c.AddCmd())
	rootCmd.AddCommand(c.EditCmd())
	rootCmd.AddCommand(c.ViewCmd())

	return rootCmd
}

func (c *CLI) CreateUserCmd() *cobra.Command {
	createUserCmd := &cobra.Command{
		Use:   "create",
		Short: "Create User.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entries, err := os.ReadDir("./output")
			if err != nil {
				fmt.Printf("Error reading directory %v", err)
				return err
			}
			if len(entries) == 0 {
				return c.User.CreateUser(args[0])
			} else {
				fmt.Printf("A user file already exist: %s\n", entries[0].Name())
				return err
			}
		},
	}
	return createUserCmd

}

func (c *CLI) SelectUserCmd() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Select user.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := c.User.SelectUser(args[0])
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Selected user: %s\n", user.Name)
			return nil
		},
	}
	userCmd.AddCommand(c.CreateUserCmd())
	return userCmd
}

func (c *CLI) AddCmd() *cobra.Command {
	//Add: Goal, Calories, Weight, or Food item
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add Goal, FoodItem or weight.",
		Long:  "Add Goal, FoodItem or weight to User's tracking history.",
	}
	addCmd.AddCommand(c.AddFoodItemCmd())
	addCmd.AddCommand(c.AddWeigthCmd())
	return addCmd
}

func (c *CLI) AddFoodItemCmd() *cobra.Command {
	addFoodCmd := &cobra.Command{
		Use:   "food",
		Short: "Add Food Item.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Food.AddFoodItem()
		},
	}
	addFoodCmd.Flags().String("meal", "", "Meal of the day(breakfast, Lunch, Dinner)")
	if err := addFoodCmd.MarkFlagRequired("meal"); err != nil {
		panic(err)
	}
	return addFoodCmd
}

func (c *CLI) AddWeigthCmd() *cobra.Command {
	addWeigthCmd := &cobra.Command{
		Use:   "weight",
		Short: "Add weight",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Weight.AddWeight()
		},
	}
	return addWeigthCmd
}

func (c *CLI) EditCmd() *cobra.Command {
	//Edit: Goal or Food Item
	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit Goal or Food Item.",
	}
	editCmd.AddCommand(c.EditGoalCmd())
	editCmd.AddCommand(c.EditFoodItemCmd())
	return editCmd
}

func (c *CLI) EditGoalCmd() *cobra.Command {
	editGoalCmd := &cobra.Command{
		Use:   "goal",
		Short: "Edit Goal.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Goal.EditCaloriesGoal()
		},
	}
	return editGoalCmd
}

func (c *CLI) EditFoodItemCmd() *cobra.Command {
	editFoodItem := &cobra.Command{
		Use:   "food",
		Short: "Edit food item.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Food.EditFoodItem()
		},
	}
	return editFoodItem
}

func (c *CLI) ViewCmd() *cobra.Command {
	//View: Remaining calories in the day
	// Weight Progress (Plot)
	viewCmd := &cobra.Command{
		Use:   "view",
		Short: "View remaining calories or weight progress.",
	}
	viewCmd.AddCommand(c.ViewLeftoverCaloriesCmd())
	viewCmd.AddCommand(c.ViewPlotCmd())
	return viewCmd
}

func (c *CLI) ViewLeftoverCaloriesCmd() *cobra.Command {
	viewCaloriesCmd := &cobra.Command{
		Use:   "calories",
		Short: "Remaining calories for the day/week.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Food.ViewRemainingCalories()
		},
	}
	//default: today. Monday, Tuesday, etc.
	//default: current. Last (week)
	viewCaloriesCmd.Flags().String("day", "today", "Show remaining calories the day.")
	viewCaloriesCmd.Flags().String("week", "current", "Show remaining calories for current week.")
	return viewCaloriesCmd
}

func (c *CLI) ViewPlotCmd() *cobra.Command {
	viewPlotCmd := &cobra.Command{
		Use:   "plot",
		Short: "Show plot",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Weight.DisplayWeightProgress()
		},
	}
	// week. month, year
	viewPlotCmd.Flags().String("weigth", "month", "Plot weight for time period.")
	return viewPlotCmd
}
