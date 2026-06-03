package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/acosio14/calories-tracker-cli/domain"
	"github.com/spf13/cobra"
)

type UserManager interface {
	CreateUser(name string) error
	SelectUser() (*domain.User, error)
}

type GoalManager interface {
	EditGoal(u UserManager) error
	EditDailyCalories(u UserManager, calories float64) error
}

type FoodTracker interface {
	AddFoodItem(
		u UserManager,
		foodItem string,
		calories int,
		servingSize int,
		quantity int,
		meal string,
	) error
	EditFoodItem() error
	DeleteFoodItem(u UserManager, foodItemID int)
	ViewRemainingCalories() error
}

type WeightTracker interface {
	AddWeight(u UserManager, weight float64) error
	DeleteWeightEntry(u UserManager, weightEntryID int)
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
	rootCmd.AddCommand(c.CreateUserCmd())
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
			foodItem := args[0]
			meal, _ := cmd.Flags().GetString("meal")
			calories, _ := cmd.Flags().GetInt("calories")
			servingSize, _ := cmd.Flags().GetInt("serving-size")
			quantity, _ := cmd.Flags().GetInt("quantity")
			// add food oatmeal --calories 160 --serving-size 35g --quantity 43g --meal breakfast
			return c.Food.AddFoodItem(c.User, foodItem, calories, servingSize, quantity, meal)
		},
	}
	addFoodCmd.Flags().String("meal", "", "Meal of the day(breakfast, Lunch, Dinner)")
	if err := addFoodCmd.MarkFlagRequired("meal"); err != nil {
		panic(err)
	}
	addFoodCmd.Flags().Int("calories", 0, "Calories per serving for food item.")
	if err := addFoodCmd.MarkFlagRequired("calories"); err != nil {
		panic(err)
	}
	addFoodCmd.Flags().Int("serving-size", 0, "Serving size of food item.")
	if err := addFoodCmd.MarkFlagRequired("serving-size"); err != nil {
		panic(err)
	}
	addFoodCmd.Flags().Int("quantity", 0, "Number of servings.")
	if err := addFoodCmd.MarkFlagRequired("quantity"); err != nil {
		panic(err)
	}

	return addFoodCmd
}

func (c *CLI) AddWeigthCmd() *cobra.Command {
	addWeigthCmd := &cobra.Command{
		Use:   "weight",
		Short: "Add weight",
		RunE: func(cmd *cobra.Command, args []string) error {
			weight_arg, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				return fmt.Errorf("invalid weight value: %w", err)
			}
			return c.Weight.AddWeight(c.User, weight_arg)
			// Need optional date flag, say I measured yesterday and wrote it down but didn't add it
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
	editCmd.AddCommand(c.EditDailyCaloriesCmd())
	editCmd.AddCommand(c.EditFoodItemCmd())
	return editCmd
}

func (c *CLI) EditGoalCmd() *cobra.Command {
	editGoalCmd := &cobra.Command{
		Use:   "goal",
		Short: "Edit Goal.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Goal.EditGoal(c.User)
		},
	}
	return editGoalCmd
}

func (c *CLI) EditDailyCaloriesCmd() *cobra.Command {
	editGoalCmd := &cobra.Command{
		Use:   "daily-calories",
		Short: "Update daily calories limit.",
		RunE: func(cmd *cobra.Command, args []string) error {
			newCalories, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				return fmt.Errorf("invalid weight value: %w", err)
			}
			return c.Goal.EditDailyCalories(c.User, newCalories)
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

func (c *CLI) DeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete food item or weight entry",
	}
	deleteCmd.AddCommand(c.DeleteFoodItem())
	deleteCmd.AddCommand(c.DeleteWeightEntry())

	return deleteCmd
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
