package service

import (
	"fmt"
	"time"

	"github.com/acosio14/calories-tracker-cli/cli"
	"github.com/acosio14/calories-tracker-cli/domain"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type WeightTracker struct{}

func (w *WeightTracker) AddWeight(u cli.UserManager, weight float64, date int) error {
	user, err := u.LoadUser()
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

	err = u.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil

}

func (w *WeightTracker) DeleteWeightEntry(u cli.UserManager, weightEntryID int) error {
	user, err := u.LoadUser()
	if err != nil {
		return err
	}
	// Delete index through slicing
	user.WeightTracker = append(user.WeightTracker[:weightEntryID], user.WeightTracker[weightEntryID+1])

	err = u.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user, %v", err)
	}

	return nil
}

func (w *WeightTracker) DisplayWeightProgress() error {
	u := &UserManager{}
	user, err := u.LoadUser()
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
