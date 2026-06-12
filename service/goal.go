package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/acosio14/calories-tracker-cli/cli"
)

type GoalManager struct{}

func (g *GoalManager) EditGoal(u cli.UserManager) error {
	user, err := u.SelectUser()
	if err != nil {
		return err
	}

	fmt.Println("What is your new goal? (lose/maintain/gain)")
	newGoal, _ := GetInput(os.Stdin)

	fmt.Println("What is your current weight?(lbs)")
	currentWeight, _ := GetInput(os.Stdin)

	fmt.Println("What is your goal weight?(lbs)")
	goalWeight, _ := GetInput(os.Stdin)

	fmt.Println("How many weeks to reach goal?")
	goalTimeline, _ := GetInput(os.Stdin)

	goalRate := (goalWeight.(float64) - currentWeight.(float64)) / goalTimeline.(float64)

	user.Goal.Type = newGoal.(string)
	user.Goal.Weight = goalWeight.(float64)
	user.Goal.Rate = goalRate

	userData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", user.Name)
	outputPath := filepath.Join("output", filename)
	err = os.WriteFile(outputPath, userData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (g *GoalManager) EditDailyCalories(u cli.UserManager, calories float64) error {
	user, err := u.SelectUser()
	if err != nil {
		return err
	}
	user.Goal.DailyCalories = calories

	userData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", user.Name)
	outputPath := filepath.Join("output", filename)
	err = os.WriteFile(outputPath, userData, 0644)
	if err != nil {
		return err
	}
	return nil
}
