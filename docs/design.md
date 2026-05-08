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
- Display: functions to display metrics

Tracker(Core):
- User
- Calories
- Weight
- Macros

CLI:
- Add
- Edit
- View
- Set

Store:
- json_file.go

Display:
- Plot weight per week, day or year