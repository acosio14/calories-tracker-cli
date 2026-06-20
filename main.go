package main

import (
	"os"

	"github.com/acosio14/calories-tracker-cli/cli"
	"github.com/acosio14/calories-tracker-cli/service"
	"github.com/acosio14/calories-tracker-cli/storage"
)

func main() {
	JSONStorage := &storage.JSONStorage{}
	food_svc := &service.FoodTracker{
		Storage: JSONStorage,
	}
	goal_svc := &service.GoalManager{
		Storage: JSONStorage,
	}
	user_svc := &service.UserManager{
		Storage: JSONStorage,
	}
	weight_svc := &service.WeightTracker{
		Storage: JSONStorage,
	}
	cli := cli.NewCLI(user_svc, food_svc, weight_svc, goal_svc)
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
