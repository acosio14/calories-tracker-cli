package main

import (
	"os"

	"github.com/acosio14/calories-tracker-cli/cli"
	service "github.com/acosio14/calories-tracker-cli/services"
)

func main() {
	userSvc := service.UserService{}
	goalSvc := service.GoalService{}
	foodSvc := service.FoodService{}
	weightSvc := service.WeightService{}

	cli := cli.NewCLI(userSvc, goalSvc, foodSvc, weightSvc)
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

// The svc above should implmement the interfaces in cli
// UserService would be struct type UserServiceStruct{}
// And it would have methods -> CreateUser() and SelectUser()
// to complete the interface
