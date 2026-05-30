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

	// I can make --meal automatic based of time and if user wishes he can use
	// serving and quantity have to be same units to make math correct else error
	// the optional flag to add it in a diff meal
	// PLACEHOLDER: need to decide correct times for breakfast, lunch, dinner
	if today > midnight && today < midday {
		meal = "breakfast"
	} else if today > midday && today < dinner {
		meal = "lunch"
	} else if today > dinner {
		meal = "dinner"
	} else {
		fmt.Println("meal out of range")
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
