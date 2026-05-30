package domain

import "time"

type Goal struct {
	Type   string  `json:"type"`
	Weight float64 `json:"weight"`
	Rate   float64 `json:"rate"`
}

type Weight struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"weight_value"`
}

type FoodItem struct {
	Date               time.Time `json:"date"`
	Meal               string    `json:"meal"`
	Name               string    `json:"name"`
	ServingSize        int       `json:"servings"`
	CaloriesPerServing int       `json:"calories_per_serving"`
	Quantity           int       `json:"quantity"`
	TotalCalories      int       `json:"total_calories"`
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
