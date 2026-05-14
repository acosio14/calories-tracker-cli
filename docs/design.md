# Design 

## Commands
- Create User
- Select User
- Set Goal: maintain, lose, gain
- Add meal/food per day
    - FoodItem -> quantity? servings?
- Modify/Update food entry
- Add weight
- Display progress: weight (weekly or monthly?)
- View remaining calories in the day (or week?)
- Update Calories (daily/weekly) to match goal (weight change rate)

## Funcitonality
- Core: Has main business logic (domain type and rules)
- CLI: commands to core
- Store: storage of data json_file or sqlite (db for meals, calories, weight)
    - Storage and Retirval of file (parser?)
    - Able to repeate values (i.e yesterday had 2 eggs, today the same, just repate instead of retyping)
        - Create meal -> reusable?
- Display: functions to display metrics

Tracker(Core):
- User
- Calories
- Weight
- Macros

CLI:
- User
    - (Name) Select User
    - (--create) Create User
- Add
    - Goal
    - Food Item
    - Weight
- Edit
    - Goal(Gain, Lose, Maintain)
    - Food Item
- View
    - Remaining Calories for day/week
    - Plot Weight Progress

Store:
- json_file.go

Display:
- Plot weight per week, day or year