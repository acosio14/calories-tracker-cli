package service

import (
	"fmt"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
	"github.com/acosio14/calories-tracker-cli/storage"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type WeightTracker struct{}

func AddWeight(weight float64, date int) error {
	user, err := storage.LoadUser()
	if err != nil {
		return err
	}

	// Need to find the highest index(ID) then increase it by one
	var count int
	for i := range user.WeightTracker {
		count = i
	}
	count++

	weight_entry := domain.Weight{
		ID:    count,
		Date:  time.Now(),
		Value: weight,
	}
	user.WeightTracker = append(user.WeightTracker, weight_entry)

	err = storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil

}

func DeleteWeightEntry(weightEntryID int) error {
	user, err := storage.LoadUser()
	if err != nil {
		return err
	}
	// Delete index through slicing
	user.WeightTracker = append(user.WeightTracker[:weightEntryID], user.WeightTracker[weightEntryID+1])

	err = storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func DisplayWeightProgress() error {
	user, err := storage.LoadUser()
	if err != nil {
		return err
	}

	pts := make(plotter.XYs, len(user.WeightTracker))
	for i := range user.WeightTracker {
		pts[i].X = float64(user.WeightTracker[i].Date.Unix())
		pts[i].Y = float64(user.WeightTracker[i].Value)
	}

	p := plot.New()
	p.Title.Text = "Weight"
	p.Y.Label.Text = "(LBS)"
	p.X.Tick.Marker = plot.TimeTicks{
		Format: "2006-01-02",
	}

	line, scatter, err := plotter.NewLinePoints(pts)
	if err != nil {
		return fmt.Errorf("failed to create line points: %v", err)
	}
	p.Add(line, scatter)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "weight.png"); err != nil {
		return fmt.Errorf("failed to save plot as png")
	}
	return nil
}
