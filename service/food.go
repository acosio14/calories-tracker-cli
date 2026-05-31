package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type FoodTracker struct{}

func (f *FoodTracker) AddFoodItem(
	u UserManager,
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
	//Date
	today := time.Now()
	hour := today.Hour()

	if hour < 11 {
		meal = "breakfast"
	} else if hour < 18 {
		meal = "lunch"
	} else {
		meal = "dinner"
	}
	totalCalories := calories * (quantity / servingSize)

	foodEntry := domain.FoodItem{
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

func (f *FoodTracker) ViewRemainingCalories() error {
	return nil
}
