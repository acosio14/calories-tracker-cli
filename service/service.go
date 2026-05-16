package service

import (
	"encoding/json"
	"fmt"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type UserManager struct{}

func (u *UserManager) CreateUser(name string) error {
	var age int
	fmt.Println("What is your age?:")
	fmt.Scan(&age)

	var gender string
	fmt.Println("What is your gender?")
	fmt.Scan(&gender)

	var height int
	fmt.Println("What is your height?")
	fmt.Scan(&height)

	var goal int
	fmt.Println("What is your goal? (lose/maintain/gain)")
	fmt.Scan(&goal)

	var weight int
	fmt.Println("What is your current weight?")
	fmt.Scan(&weight)
	// Need to get date, and add date and weight to weight struct

	//Validate if User does not exist

	user := domain.User{
		Name:   name,
		Age:    age,
		Gender: gender,
		Height: height,
		//Goal:   goal,
		//Weight: weight,
	}
	json.Marshal(user)

	return nil
}

func (u *UserManager) SelectUser() error {
	return nil
}

type GoalManager struct{}

func (g *GoalManager) AddCaloriesGoal() error {
	return nil
}

func (g *GoalManager) EditCaloriesGoal() error {
	return nil
}

type FoodTracker struct{}

func (f *FoodTracker) AddFoodItem() error {
	return nil
}

func (f *FoodTracker) EditFoodItem() error {
	return nil
}

func (f *FoodTracker) ViewRemainingCalories() error {
	return nil
}

type WeightTracker struct{}

func (w *WeightTracker) AddWeight() error {
	return nil
}

func (w *WeightTracker) DisplayWeightProgress() error {
	return nil
}
