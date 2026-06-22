package service

import (
	"cmp"
	"fmt"
	"slices"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type FoodTracker struct {
	Storage StorageInterface
}

func (f *FoodTracker) AddFoodItem(
	foodItem string,
	calories float64,
	servingSize float64,
	quantity float64,
	meal string,
) error {
	user, err := f.Storage.LoadUser()
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

	var maxID int
	if len(user.FoodJournal) == 0 {
		maxID = -1
	} else {
		maxID = slices.MaxFunc(user.FoodJournal,
			func(a, b domain.FoodItem) int {
				return cmp.Compare(a.ID, b.ID)
			},
		).ID
	}

	foodEntry := domain.FoodItem{
		ID:                 maxID + 1,
		Date:               today,
		Meal:               meal,
		Name:               foodItem,
		ServingSize:        servingSize,
		CaloriesPerServing: calories,
		Quantity:           quantity,
		TotalCalories:      totalCalories,
	}
	user.FoodJournal = append(user.FoodJournal, foodEntry)

	err = f.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (f *FoodTracker) EditFoodItem(foodInputID int, flag domain.FoodItemInput) error {
	// flags = date, meal, name, servingSize, calories, quantity
	user, err := f.Storage.LoadUser()
	if err != nil {
		return err
	}

	var item *domain.FoodItem
	for i, food := range user.FoodJournal {
		if foodInputID == food.ID {
			item = &user.FoodJournal[i]
		}
	}

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

	if item.ServingSize <= 0 {
		return fmt.Errorf("serving size less than or equals to zero")
	}
	item.TotalCalories = item.CaloriesPerServing * (item.Quantity / item.ServingSize) // 40g/30g * (100 cal/g)

	err = f.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (f *FoodTracker) DeleteFoodItem(foodItemID int) error {
	user, err := f.Storage.LoadUser()
	if err != nil {
		return err
	}
	// Delete index
	var foodIndex int
	for i, food := range user.FoodJournal {
		if foodItemID == food.ID {
			foodIndex = i
		}
	}
	user.FoodJournal = slices.Delete(user.FoodJournal, foodIndex, foodIndex+1)

	err = f.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (f *FoodTracker) ViewRemainingCalories(date time.Time) error {
	user, err := f.Storage.LoadUser()
	if err != nil {
		return err
	}
	var dailyCalories float64 = 0
	year2, month2, day2 := date.Date()
	for _, foodItem := range user.FoodJournal {
		year, month, day := foodItem.Date.Date()
		if year == year2 && month == month2 && day == day2 {
			fmt.Printf("%s | %.2f | %.2f\n", foodItem.Name, foodItem.Quantity, foodItem.TotalCalories)
			dailyCalories += foodItem.TotalCalories
		}
	}

	remainingCalories := user.Goal.DailyCalories - dailyCalories
	fmt.Printf("Remaining Calories for %v: %.2f\n", date, remainingCalories)

	return nil
}
