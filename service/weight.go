package service

type WeightTracker struct{}

func (w *WeightTracker) AddWeight() error {
	user, err := SelectUser()
	if err != nil {
		return err
	}

	//Now add weight to user domain

}

func (w *WeightTracker) DisplayWeightProgress() error {
	return nil
}
