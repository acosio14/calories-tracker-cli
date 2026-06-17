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

func GetString(r io.Reader) (string, error) {
	var returnValue string
	_, err := fmt.Fscan(r, &returnValue)
	if err != nil {
		return "", err
	}

	return returnValue, nil
}

func GetInt(r io.Reader) (int, error) {
	var returnValue int
	_, err := fmt.Fscan(r, &returnValue)
	if err != nil {
		return 0, err
	}

	return returnValue, nil
}

func GetFloat64(r io.Reader) (float64, error) {
	var returnValue float64
	_, err := fmt.Fscan(r, &returnValue)
	if err != nil {
		return 0.0, err
	}

	return returnValue, nil
}

func (u *UserManager) CreateUser(name string) error {

	fmt.Print("Enter Birthday (MM/DD/YYYY):")
	birthday, _ := GetInt(os.Stdin)
	// Need to parse b-day into proper digit in order to calculate age
	age := time.Now().Day() - birthday

	fmt.Print("Gender (M/F): ")
	gender, _ := GetString(os.Stdin)

	fmt.Print("Height (inches): ")
	height, _ := GetInt(os.Stdin)

	fmt.Print("Goal (lose/maintain/gain): ")
	userGoal, _ := GetString(os.Stdin)

	fmt.Print("Current weight (lbs): ")
	currentWeight, _ := GetFloat64(os.Stdin)

	fmt.Print("Goal weight (lbs): ")
	goalWeight, _ := GetFloat64(os.Stdin)

	fmt.Print("Duration (weeks): ")
	goalTimeline, _ := GetFloat64(os.Stdin)

	fmt.Print("Daily Calories intake: ")
	dailyCalories, _ := GetFloat64(os.Stdin)

	goalRate := (goalWeight - currentWeight) / goalTimeline

	goal := domain.Goal{
		Type:          userGoal,
		Weight:        goalWeight,
		Rate:          goalRate,
		DailyCalories: dailyCalories,
	}
	today := time.Now()

	weight := make([]domain.Weight, 1)
	weight[0] = domain.Weight{
		ID:    0,
		Date:  today,
		Value: currentWeight,
	}

	foodItem := make([]domain.FoodItem, 1)

	user := &domain.User{
		Name:          name,
		Age:           age,
		Gender:        gender,
		Height:        height,
		Goal:          goal,
		WeightTracker: weight,
		FoodJournal:   foodItem,
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home dir %v", err)
	}
	outputFolder := filepath.Join(homeDir, "Projects/calories-tracker-cli/output")
	err = os.MkdirAll(outputFolder, 0755)
	if err != nil {
		return fmt.Errorf("error creating output folder %v", err)
	}
	err = u.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (u *UserManager) LoadUser() (*domain.User, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error finding home dir %v", err)
	}
	outputFolder := filepath.Join(homeDir, "Projects/calories-tracker-cli/output")

	entries, err := os.ReadDir(outputFolder)
	if err != nil {
		return nil, fmt.Errorf("Error with reading output directory %v", err)
	}

	var user domain.User
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

func (u *UserManager) SaveUser(user *domain.User) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home dir %v", err)
	}
	outputFolder := filepath.Join(homeDir, "Projects/calories-tracker-cli/output")

	userData, err := json.MarshalIndent(user, "", "	")
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
