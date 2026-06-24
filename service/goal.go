package service

import (
	"fmt"
	"os"
)

type GoalManager struct {
	Storage Storage
}

func (g *GoalManager) EditGoal() error {
	user, err := g.Storage.LoadUser()
	if err != nil {
		return err
	}

	fmt.Println("What is your new goal? (lose/maintain/gain)")
	newGoal, _ := readString(os.Stdin)

	fmt.Println("What is your current weight?(lbs)")
	currentWeight, _ := readFloat(os.Stdin)

	fmt.Println("What is your goal weight?(lbs)")
	goalWeight, _ := readFloat(os.Stdin)

	fmt.Println("How many weeks to reach goal?")
	goalTimeline, _ := readFloat(os.Stdin)

	goalRate := (goalWeight - currentWeight) / goalTimeline

	user.Goal.Type = newGoal
	user.Goal.Weight = goalWeight
	user.Goal.Rate = goalRate

	err = g.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %w", err)
	}

	return nil
}

func (g *GoalManager) EditDailyCalories(calories float64) error {
	user, err := g.Storage.LoadUser()
	if err != nil {
		return err
	}
	user.Goal.DailyCalories = calories

	err = g.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %w", err)
	}

	return nil
}
