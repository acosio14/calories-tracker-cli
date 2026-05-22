package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type UserManager struct {
	statePath string
}

func (u *UserManager) CreateUser(name string) error {
	var age int
	fmt.Println("What is your age?:")
	fmt.Scan(&age)

	var gender string
	fmt.Println("What is your gender?") //add M/F, truncate to lower
	fmt.Scan(&gender)

	var height int
	fmt.Println("What is your height?") //add in inches
	fmt.Scan(&height)

	var userGoal string
	fmt.Println("What is your goal? (lose/maintain/gain)") //truncate to lower case, and if misspelled
	fmt.Scan(&userGoal)

	var initWeight float32
	fmt.Println("What is your current weight?") //add lbs
	fmt.Scan(&initWeight)

	var goalWeight float32
	fmt.Println("What is your goal weight?")
	fmt.Scan(&goalWeight)

	var goalTimeline float32
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

	userData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		return err
	}
	outputFolder := "output"
	err = os.MkdirAll(outputFolder, 0644)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", user.Name)
	outputPath := filepath.Join(outputFolder, filename)
	err = os.WriteFile(outputPath, userData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserManager) SelectUser(name string) (*domain.User, error) {

	jsonFile := fmt.Sprintf("%s.json", name)
	jsonContent, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("read user %q: %w", name, err)
	}

	var user domain.User
	err = json.Unmarshal(jsonContent, &user)
	if err != nil {
		return nil, fmt.Errorf("parse user %q: %w", name, err)
	}

	return &user, nil
}
