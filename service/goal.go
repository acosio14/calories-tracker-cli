package service

import (
	"fmt"
	"os"

	"github.com/acosio14/calories-tracker-cli/cli"
)

type GoalManager struct{}

func (g *GoalManager) EditGoal(u cli.UserManager) error {
	user, err := u.LoadUser()
	if err != nil {
		return err
	}

	fmt.Println("What is your new goal? (lose/maintain/gain)")
	newGoal, _ := GetString(os.Stdin)

	fmt.Println("What is your current weight?(lbs)")
	currentWeight, _ := GetFloat64(os.Stdin)

	fmt.Println("What is your goal weight?(lbs)")
	goalWeight, _ := GetFloat64(os.Stdin)

	fmt.Println("How many weeks to reach goal?")
	goalTimeline, _ := GetFloat64(os.Stdin)

	goalRate := (goalWeight - currentWeight) / goalTimeline

	user.Goal.Type = newGoal
	user.Goal.Weight = goalWeight
	user.Goal.Rate = goalRate

	err = u.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (g *GoalManager) EditDailyCalories(u cli.UserManager, calories float64) error {
	user, err := u.LoadUser()
	if err != nil {
		return err
	}
	user.Goal.DailyCalories = calories

	err = u.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}
