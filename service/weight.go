package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type WeightTracker struct{}

func (w *WeightTracker) AddWeight(u UserManager, weight float64) error {
	user, err := u.SelectUser()
	if err != nil {
		return err
	}

	weight_entry := domain.Weight{
		Date:  time.Now(),
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
	u := &UserManager{}
	user, err := u.SelectUser()
	if err != nil {
		return err
	}

	pts := make(plotter.XYs, len(user.Weight))
	for i := range user.Weight {
		pts[i].X = float64(user.Weight[i].Date.Unix())
		pts[i].Y = float64(user.Weight[i].Value)
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
