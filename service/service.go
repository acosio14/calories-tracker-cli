package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

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

	var userGoal string
	fmt.Println("What is your goal? (lose/maintain/gain)")
	fmt.Scan(&userGoal)

	var initWeight int
	fmt.Println("What is your current weight?")
	fmt.Scan(&initWeight)

	var goalWeight int
	fmt.Println("What is your goal weight?")
	fmt.Scan(&goalWeight)

	var goalTimeline int
	fmt.Println("How many weeks to reach goal?")
	fmt.Scan(&goalTimeline)

	goalRate := (goalWeight - initWeight) / goalTimeline

	goal := domain.Goal{
		Type:   userGoal,
		Weight: goalWeight,
		Rate:   goalRate,
	}
	today := time.Now().Format("January 2, 2006")

	weight := make([]domain.Weight, 1)
	weight[0] = domain.Weight{
		Date:  today,
		Value: initWeight,
	}

	foodItem := make([]domain.FoodItem, 1)

	user := domain.User{
		Name:        name,
		Age:         age,
		Gender:      gender,
		Height:      height,
		Goal:        goal,
		Weight:      weight,
		FoodJournal: foodItem,
	}

	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", user.Name)
	err = os.WriteFile(filename, userData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserManager) SelectUser(name string) error {
	//unmarshals json file pertaining to user. How does in keep it? Cache?
	//Ability to write name or list names and select?

	jsonFile := fmt.Sprintf("%s.json", name)
	jsonContent, err := os.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	var user domain.User
	err = json.Unmarshal(jsonContent, &user)
	if err != nil {
		return err
	}

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
