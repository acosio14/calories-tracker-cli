package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type GoalManager struct{}

func (g *GoalManager) EditGoal(u UserManager) error {
	user, err := u.SelectUser()
	if err != nil {
		return err
	}

	var newGoal string
	fmt.Println("What is your new goal? (lose/maintain/gain)")
	fmt.Scan(&newGoal)

	var currentWeight float64
	fmt.Println("What is your current weight?(lbs)")
	fmt.Scan(&currentWeight)

	var goalWeight float64
	fmt.Println("What is your goal weight?(lbs)")
	fmt.Scan(&goalWeight)

	var goalTimeline float64
	fmt.Println("How many weeks to reach goal?")
	fmt.Scan(&goalTimeline)

	goalRate := (goalWeight - currentWeight) / goalTimeline

	user.Goal.Type = newGoal
	user.Goal.Weight = goalWeight
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

func (g *GoalManager) EditDailyCalories(u UserManager, calories float64) error {
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
