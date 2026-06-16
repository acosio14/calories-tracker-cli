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

	err = u.SaveUser(user, "output")
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

	err = u.SaveUser(user, "output")
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}
