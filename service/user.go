package service

import (
	"fmt"
	"io"
	"os"
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

func (u *UserManager) CreateUser(name string) error {

	fmt.Print("Enter Birthday (MM/DD/YYYY):")
	birthdayString, _ := readString(os.Stdin)

	birthday, err := time.Parse("01/02/2006", birthdayString)
	if err != nil {
		return err
	}
	// Need to parse b-day into proper digit in order to calculate age
	age := time.Since(birthday).Hours() / 24 / 365

	fmt.Print("Gender (M/F): ")
	gender, _ := readString(os.Stdin)

	fmt.Print("Height (inches): ")
	height, _ := readInt(os.Stdin)

	fmt.Print("Goal (lose/maintain/gain): ")
	userGoal, _ := readString(os.Stdin)

	fmt.Print("Current weight (lbs): ")
	currentWeight, _ := readFloat(os.Stdin)

	fmt.Print("Goal weight (lbs): ")
	goalWeight, _ := readFloat(os.Stdin)

	fmt.Print("Duration (weeks): ")
	goalTimeline, _ := readFloat(os.Stdin)

	fmt.Print("Daily Calories intake: ")
	dailyCalories, _ := readFloat(os.Stdin)

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

	storage.OutputFolder(storage.Options{CreateFolder: true})
	err = u.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}
