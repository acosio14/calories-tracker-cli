package service

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
	"github.com/acosio14/calories-tracker-cli/storage"
)

func readString(r io.Reader) (string, error) {
	var returnValue string
	_, err := fmt.Fscan(r, &returnValue)
	if err != nil {
		return "", err
	}

	return returnValue, nil
}

func readInt(r io.Reader) (int, error) {
	var returnValue int
	_, err := fmt.Fscan(r, &returnValue)
	if err != nil {
		return 0, err
	}

	return returnValue, nil
}

func readFloat(r io.Reader) (float64, error) {
	var returnValue float64
	_, err := fmt.Fscan(r, &returnValue)
	if err != nil {
		return 0.0, err
	}

	return returnValue, nil
}

type UserManager struct {
	Storage Storage
}

func userExists(userFile string) error {
	outputFolder, err := storage.OutputFolderPath()
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", userFile)
	outputFile := filepath.Join(*outputFolder, filename)
	entries, err := os.ReadDir(outputFile)
	if err != nil {
		return fmt.Errorf("Error reading directory %w", err)
	}
	if len(entries) != 0 {
		fmt.Printf("A user file already exist: %s\n", entries[0].Name())
		return err
	}

	return nil
}

func (u *UserManager) CreateUser(name string) error {
	err := userExists(name) // This prevents having to go through the entire questions and then finding out there is already one.
	if err != nil {
		return err
	}

	fmt.Print("Enter Birthday (MM/DD/YYYY):")
	birthdayString, err := readString(os.Stdin)
	if err != nil {
		return fmt.Errorf("error with input %w", err)
	}

	birthday, err := time.Parse("01/02/2006", birthdayString)
	if err != nil {
		return err
	}
	// Need to parse b-day into proper digit in order to calculate age
	age := time.Since(birthday).Hours() / 24 / 365

	fmt.Print("Gender (M/F): ")
	gender, err := readString(os.Stdin)
	if err != nil {
		return fmt.Errorf("error with input %w", err)
	}

	fmt.Print("Height (inches): ")
	height, err := readInt(os.Stdin)
	if err != nil {
		return fmt.Errorf("error with input %w", err)
	}

	fmt.Print("Goal (lose/maintain/gain): ")
	userGoal, err := readString(os.Stdin)
	if err != nil {
		return fmt.Errorf("error with input %w", err)
	}

	fmt.Print("Current weight (lbs): ")
	currentWeight, err := readFloat(os.Stdin)
	if err != nil {
		return fmt.Errorf("error with input %w", err)
	}

	fmt.Print("Goal weight (lbs): ")
	goalWeight, err := readFloat(os.Stdin)
	if err != nil {
		return fmt.Errorf("error with input %w", err)
	}

	fmt.Print("Duration (weeks): ")
	goalTimeline, err := readFloat(os.Stdin)
	if err != nil {
		return fmt.Errorf("error with input %w", err)
	}

	fmt.Print("Daily Calories intake: ")
	dailyCalories, err := readFloat(os.Stdin)
	if err != nil {
		return fmt.Errorf("error with input %w", err)
	}

	goalRate := (goalWeight - currentWeight) / goalTimeline

	goal := domain.Goal{
		Type:          userGoal,
		Weight:        goalWeight,
		Rate:          goalRate,
		DailyCalories: dailyCalories,
	}
	today := time.Now()

	weight := make([]domain.Weight, 0)
	weight = append(
		weight,
		domain.Weight{
			ID:    0,
			Date:  today,
			Value: currentWeight,
		},
	)

	foodItem := make([]domain.FoodItem, 0)

	user := &domain.User{
		Name:          name,
		Age:           int(age),
		Gender:        gender,
		Height:        height,
		Goal:          goal,
		WeightTracker: weight,
		FoodJournal:   foodItem,
	}

	err = u.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %w", err)
	}

	return nil
}
