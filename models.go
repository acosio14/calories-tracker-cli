package main

type Goal struct {
	Type   string `json:"type"`
	Weight int    `json:"weight"`
	Rate   int    `json:"rate"`
}

type WeightProgress struct {
	Date  int `json:"date"`
	Value int `json:"weight_value"`
}

type FoodItem struct {
	Date               int    `json:"date"`
	Meal               string `json:"meal"`
	Name               string `json:"nanme"`
	Servings           int    `json:"servings"`
	CaloriesPerServing int    `json:"calories_per_serving"`
	TotalCalories      int    `json:"total_calories"`
}

type TrackerUser struct {
	Name   string         `json:"name"`
	Age    int            `json:"age"`
	Gender string         `json:"gender"`
	Height int            `json:"height"`
	Goal   Goal           `json:"goal"`
	Weight WeightProgress `json:"weight"`
	Food   FoodItem       `json:"food"`
}
