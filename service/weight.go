package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type WeightTracker struct{}

func (w *WeightTracker) AddWeight(weight float32) error {
	// Appends a weight entry to the user's weight list
	user, err := SelectUser()
	if err != nil {
		return err
	}

	weight_entry := domain.Weight{
		Date:  time.Now().Format("January 2, 2006"),
		Value: weight,
	}
	user.Weight = append(user.Weight, weight_entry)

	userData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s.json", user.Name)
	outputPath := filepath.Join("output", filename)
	err = os.WriteFile(outputPath, userData, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (w *WeightTracker) DisplayWeightProgress() error {
	return nil
}
