package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type UserManager struct {
	statePath string
}

func (u *UserManager) CreateUser(name string) error {
	var age int
	fmt.Println("What is your age?:")
	fmt.Scan(&age)

	var gender string
	fmt.Println("What is your gender?")
	fmt.Scan(&gender)

	var height int
	fmt.Println("What is your height?")
	fmt.Scan(&height)

	var userGoal string
	fmt.Println("What is your goal? (lose/maintain/gain)")
	fmt.Scan(&userGoal)

	var initWeight int
	fmt.Println("What is your current weight?")
	fmt.Scan(&initWeight)

	var goalWeight int
	fmt.Println("What is your goal weight?")
	fmt.Scan(&goalWeight)

	var goalTimeline int
	fmt.Println("How many weeks to reach goal?")
	fmt.Scan(&goalTimeline)

	goalRate := (goalWeight - initWeight) / goalTimeline

	goal := domain.Goal{
		Type:   userGoal,
		Weight: goalWeight,
		Rate:   goalRate,
	}
	today := time.Now().Format("January 2, 2006")

	weight := make([]domain.Weight, 1)
	weight[0] = domain.Weight{
		Date:  today,
		Value: initWeight,
	}

	foodItem := make([]domain.FoodItem, 1)

	user := domain.User{
		Name:        name,
		Age:         age,
		Gender:      gender,
		Height:      height,
		Goal:        goal,
		Weight:      weight,
		FoodJournal: foodItem,
	}

	userData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", user.Name)
	err = os.WriteFile(filename, userData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserManager) SelectUser(name string) (*domain.User, error) {

	jsonFile := fmt.Sprintf("%s.json", name)
	jsonContent, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("read user %q: %w", name, err)
	}

	var user domain.User
	err = json.Unmarshal(jsonContent, &user)
	if err != nil {
		return nil, fmt.Errorf("parse user %q: %w", name, err)
	}

	return &user, nil
}

func (u *UserManager) SetCurrentUser(name string) error {
	if _, err := u.SelectUser(name); err != nil {
		return err
	}
	return os.WriteFile(u.statePath, []byte(name), 0600)
}

func (u *UserManager) CurrentUser() (*domain.User, error) {
	name, err := os.ReadFile(u.statePath)
	if err != nil {
		return nil, fmt.Errorf("no user selected: %w", err)
	}
	return u.SelectUser(strings.TrimSpace(string(name)))
}
