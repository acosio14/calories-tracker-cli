package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/acosio14/calories-tracker-cli/cli"
	"github.com/acosio14/calories-tracker-cli/domain"
)

type FoodTracker struct{}

func (f *FoodTracker) AddFoodItem(
	u cli.UserManager,
	foodItem string,
	calories int,
	servingSize int,
	quantity int,
	meal string,
) error {
	user, err := u.SelectUser()
	if err != nil {
		return err
	}

	today := time.Now()
	hour := today.Hour()
	if meal == "" {
		switch {
		case hour < 11:
			meal = "breakfast"
		case hour < 18:
			meal = "lunch"
		default:
			meal = "dinner"
		}
	}
	totalCalories := calories * (quantity / servingSize)

	// Need to find the highest index(ID) then increase it by one
	var count int
	for i := range user.FoodJournal {
		count = i
	}
	count++

	foodEntry := domain.FoodItem{
		ID:                 count,
		Date:               today,
		Meal:               meal,
		Name:               foodItem,
		ServingSize:        servingSize,
		CaloriesPerServing: calories,
		Quantity:           quantity,
		TotalCalories:      totalCalories,
	}
	user.FoodJournal = append(user.FoodJournal, foodEntry)

	userData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s.json", user.Name)
	outputPath := filepath.Join("output", filename)
	err = os.WriteFile(outputPath, userData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *FoodTracker) EditFoodItem() error {
	return nil
}

func (f *FoodTracker) DeleteFoodItem(u cli.UserManager, foodItemID int) error {
	return nil
}

func (f *FoodTracker) ViewRemainingCalories() error {
	return nil
}
