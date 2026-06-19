package service

import (
	"fmt"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
	"github.com/acosio14/calories-tracker-cli/storage"
)

func AddFoodItem(
	foodItem string,
	calories float64,
	servingSize float64,
	quantity float64,
	meal string,
) error {
	user, err := storage.LoadUser()
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

	err = storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func EditFoodItem(foodItemID int, flag domain.FoodItemInput) error {
	// flags = date, meal, name, servingSize, calories, quantity
	user, err := storage.LoadUser()
	if err != nil {
		return err
	}
	item := &user.FoodJournal[foodItemID]

	if flag.Date != nil {
		item.Date = *flag.Date
	}
	if flag.Meal != nil {
		item.Meal = *flag.Meal
	}
	if flag.Name != nil {
		item.Name = *flag.Name
	}
	if flag.ServingSize != nil {
		item.ServingSize = *flag.ServingSize
	}
	if flag.Calories != nil {
		item.CaloriesPerServing = *flag.Calories
	}
	if flag.Quantity != nil {
		item.Quantity = *flag.Quantity
	}

	err = storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func DeleteFoodItem(foodItemID int) error {
	user, err := storage.LoadUser()
	if err != nil {
		return err
	}
	// Delete index through slicing
	user.FoodJournal = append(user.FoodJournal[:foodItemID], user.FoodJournal[foodItemID+1])

	err = storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func ViewRemainingCalories(date time.Time) error {
	user, err := storage.LoadUser()
	if err != nil {
		return err
	}

	var dailyCalories float64 = 0
	for _, foodItem := range user.FoodJournal {
		if foodItem.Date.Equal(date) {
			fmt.Printf("%s | %.2f | %.2f\n", foodItem.Name, foodItem.Quantity, foodItem.TotalCalories)
			dailyCalories += foodItem.TotalCalories
		}
	}

	remainingCalories := user.Goal.DailyCalories - dailyCalories
	fmt.Printf("Remaining Calories for %v: %.2f\n", date, remainingCalories)

	return nil
}
