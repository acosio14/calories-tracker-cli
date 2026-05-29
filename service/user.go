package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type UserManager struct{}

func (u *UserManager) CreateUser(name string) error {
	var age int
	fmt.Println("What is your age?:") //change to use birthday and calculate age
	fmt.Scan(&age)

	var gender string
	fmt.Println("What is your gender?(M/F)")
	fmt.Scan(&gender)

	var height int
	fmt.Println("What is your height?(inches)")
	fmt.Scan(&height)

	var userGoal string
	fmt.Println("What is your goal? (lose/maintain/gain)")
	fmt.Scan(&userGoal)

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

	goal := domain.Goal{
		Type:   userGoal,
		Weight: goalWeight,
		Rate:   goalRate,
	}
	today := time.Now()

	weight := make([]domain.Weight, 1)
	weight[0] = domain.Weight{
		Date:  today,
		Value: currentWeight,
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

func (u *UserManager) SelectUser() (*domain.User, error) {
	var user domain.User

	entries, err := os.ReadDir("../output")
	if err != nil {
		return nil, fmt.Errorf("Error with reading output directory %v", err)
	}

	if len(entries) < 1 {
		return nil, fmt.Errorf("User not created")
	} else {
		jsonFile := string(entries[0].Name())
		jsonContent, err := os.ReadFile(jsonFile)
		if err != nil {
			return nil, fmt.Errorf("read user %q: %w", jsonFile, err)
		}

		err = json.Unmarshal(jsonContent, &user)
		if err != nil {
			return nil, fmt.Errorf("parse user %q: %w", jsonFile, err)
		}
	}

	return &user, nil
}
