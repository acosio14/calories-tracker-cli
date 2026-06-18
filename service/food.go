package service

import (
	"fmt"
	"time"

	"github.com/acosio14/calories-tracker-cli/cli"
	"github.com/acosio14/calories-tracker-cli/domain"
)

type FoodTracker struct{}

func (f *FoodTracker) AddFoodItem(
	u cli.UserManager,
	foodItem string,
	calories float64,
	servingSize float64,
	quantity float64,
	meal string,
) error {
	user, err := u.LoadUser()
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

	err = u.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (f *FoodTracker) EditFoodItem(u cli.UserManager, foodItemID int, flag cli.FoodItemInput) error {
	// flags = date, meal, name, servingSize, calories, quantity
	user, err := u.LoadUser()
	if err != nil {
		return err
	}

	var flagIndex []int
	for i, flag := range flags {
		if flag != "" && flag != 0 {
			flagIndex = append(flagIndex, i)
		}
	}

	// find out if foodItemID is 1 based or 0 based
	for _, flagEnum := range flagIndex {
		switch flagEnum {
		case 0:
			user.FoodJournal[foodItemID-1].Date = *flag.Date
		case 1:
			//meal
			user.FoodJournal[foodItemID-1].Meal = *flag.Meal
		case 2:
			//name
			user.FoodJournal[foodItemID-1].Name = *flag.Name
		case 3:
			//serving-size
			user.FoodJournal[foodItemID-1].ServingSize = *flag.ServingSize
		case 4:
			//calories
			user.FoodJournal[foodItemID-1].TotalCalories = *flag.Calories
		case 5:
			//quantity
			user.FoodJournal[foodItemID-1].Quantity = *flag.Quantity
		default:
			fmt.Println("Error")
		}
	}

	err = u.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (f *FoodTracker) DeleteFoodItem(u cli.UserManager, foodItemID int) error {
	user, err := u.LoadUser()
	if err != nil {
		return err
	}
	// Delete index through slicing
	user.FoodJournal = append(user.FoodJournal[:foodItemID], user.FoodJournal[foodItemID+1])

	err = u.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (f *FoodTracker) ViewRemainingCalories(u cli.UserManager, date time.Time) error {
	user, err := u.LoadUser()
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
