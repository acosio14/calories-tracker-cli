package main

import (
	"os"

	"github.com/acosio14/calories-tracker-cli/cli"
	"github.com/acosio14/calories-tracker-cli/service"
	"github.com/acosio14/calories-tracker-cli/storage"
)

func main() {
	JSONStorage := &storage.JSONStorage{}
	foodSvc := &service.FoodTracker{
		Storage: JSONStorage,
	}
	goalSvc := &service.GoalManager{
		Storage: JSONStorage,
	}
	userSvc := &service.UserManager{
		Storage: JSONStorage,
	}
	weightSvc := &service.WeightTracker{
		Storage: JSONStorage,
	}
	cli := cli.NewCLI(userSvc, foodSvc, weightSvc, goalSvc)
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
