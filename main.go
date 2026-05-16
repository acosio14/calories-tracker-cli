package main

import (
	"os"

	"github.com/acosio14/calories-tracker-cli/cli"
	"github.com/acosio14/calories-tracker-cli/service"
)

func main() {
	userSvc := &service.UserManager{}
	goalSvc := &service.GoalManager{}
	foodSvc := &service.FoodTracker{}
	weightSvc := &service.WeightTracker{}

	cli := cli.NewCLI(userSvc, goalSvc, foodSvc, weightSvc)
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
