package domain

import "time"

type Goal struct {
	Type          string  `json:"type"`
	Weight        float64 `json:"weight"`
	Rate          float64 `json:"rate"`
	DailyCalories float64 `json:"daily_calories"`
}

type Weight struct {
	ID    int       `json:"id"`
	Date  time.Time `json:"date"`
	Value float64   `json:"weight_value"`
}

type FoodItem struct {
	ID                 int       `json:"id"`
	Date               time.Time `json:"date"` //Future: Find way to make this a struct of calendar dates not time included
	Meal               string    `json:"meal"`
	Name               string    `json:"name"`
	ServingSize        float64   `json:"servings"`
	CaloriesPerServing float64   `json:"calories_per_serving"`
	Quantity           float64   `json:"quantity"`
	TotalCalories      float64   `json:"total_calories"`
}

type FoodItemInput struct {
	Date        *time.Time
	Meal        *string
	Name        *string
	ServingSize *float64
	Calories    *float64
	Quantity    *float64
}

type User struct {
	Name          string     `json:"name"`
	Age           int        `json:"age"`
	Gender        string     `json:"gender"`
	Height        int        `json:"height"`
	Goal          Goal       `json:"goal"`
	WeightTracker []Weight   `json:"weight"`
	FoodJournal   []FoodItem `json:"food_items"`
}
