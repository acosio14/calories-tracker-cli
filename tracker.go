package main

type Tracker struct {
}

func (t *Tracker) CreateUser() error {
	// probably add init weight
	return nil
}

func (t *Tracker) SelectUser() error {
	return nil
}

func (t *Tracker) SetCaloriesGoal() error {
	// manually set calories
	return nil
}

func (t *Tracker) UpdateCaloriesGoal() error {
	// manually update calories
	return nil
}

func (t *Tracker) AddWeight() error {
	return nil
}

func (t *Tracker) AddFoodItem() error {
	return nil
}

func (t *Tracker) ModifyFoodItem() error {
	return nil
}

func (t *Tracker) ViewRemaindingCalories() error {
	return nil
}

func (t *Tracker) DisplayWeightProgress() error {
	return nil
}
