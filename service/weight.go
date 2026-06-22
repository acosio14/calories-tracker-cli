package service

import (
	"cmp"
	"fmt"
	"slices"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type WeightTracker struct {
	Storage Storage
}

func (w *WeightTracker) AddWeight(weight float64, date time.Time) error {
	user, err := w.Storage.LoadUser()
	if err != nil {
		return err
	}

	var maxID int
	if len(user.WeightTracker) == 0 {
		return fmt.Errorf("weightTracker slice is empty")
	} else {
		maxID = slices.MaxFunc(user.WeightTracker,
			func(a, b domain.Weight) int {
				return cmp.Compare(a.ID, b.ID)
			},
		).ID
	}

	weight_entry := domain.Weight{
		ID:    maxID + 1,
		Date:  date,
		Value: weight,
	}
	user.WeightTracker = append(user.WeightTracker, weight_entry)

	err = w.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil

}

func (w *WeightTracker) DeleteWeightEntry(weightInputID int) error {
	user, err := w.Storage.LoadUser()
	if err != nil {
		return err
	}
	// Delete index through slicing
	var weightIndex int
	for i, weight := range user.WeightTracker {
		if weightInputID == weight.ID {
			weightIndex = i
		}
	}
	user.WeightTracker = slices.Delete(user.WeightTracker, weightIndex, weightIndex+1)

	err = w.Storage.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (w *WeightTracker) DisplayWeightProgress() error {
	user, err := w.Storage.LoadUser()
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
