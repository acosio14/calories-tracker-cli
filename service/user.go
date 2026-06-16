package service

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type UserManager struct{}

func GetInput(r io.Reader) (any, error) {
	var returnValue any
	_, err := fmt.Fscan(r, &returnValue)
	if err != nil {
		return "", err
	}

	return returnValue, nil
}

func (u *UserManager) CreateUser(name string) error {

	fmt.Print("Enter Birthday (MM/DD/YYYY):")
	birthday, _ := GetInput(os.Stdin)
	// Need to parse b-day into proper digit in order to calculate age
	age := time.Now().Day() - birthday.(int)

	fmt.Print("Gender (M/F): ")
	gender, _ := GetInput(os.Stdin)

	fmt.Print("Height (inches): ")
	height, _ := GetInput(os.Stdin)

	fmt.Print("Goal (lose/maintain/gain): ")
	userGoal, _ := GetInput(os.Stdin)

	fmt.Print("Current weight (lbs): ")
	currentWeight, _ := GetInput(os.Stdin)

	fmt.Print("Goal weight (lbs): ")
	goalWeight, _ := GetInput(os.Stdin)

	fmt.Print("Duration (weeks): ")
	goalTimeline, _ := GetInput(os.Stdin)

	fmt.Print("Daily Calories intake: ")
	dailyCalories, _ := GetInput(os.Stdin)

	goalRate := (goalWeight.(float64) - currentWeight.(float64)) / goalTimeline.(float64)

	goal := domain.Goal{
		Type:          userGoal.(string),
		Weight:        goalWeight.(float64),
		Rate:          goalRate,
		DailyCalories: dailyCalories.(float64),
	}
	today := time.Now()

	weight := make([]domain.Weight, 1)
	weight[0] = domain.Weight{
		ID:    0,
		Date:  today,
		Value: currentWeight.(float64),
	}

	foodItem := make([]domain.FoodItem, 1)

	user := domain.User{
		Name:          name,
		Age:           age,
		Gender:        gender.(string),
		Height:        height.(int),
		Goal:          goal,
		WeightTracker: weight,
		FoodJournal:   foodItem,
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

func (u *UserManager) LoadUser() (*domain.User, error) {
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

func (u *UserManager) SaveUser(User *domain.User) error {
	return nil
}
