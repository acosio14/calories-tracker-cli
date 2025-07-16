from user import User

def create_goal_plan(goal: str) -> int:
    maintance = User.calculate_maintance_calories()

    if goal == "gain":
        calories = maintance + 500
    if goal == "lose":
        calories = maintance - 500
    if goal == "maintain":
        calories = maintance
    
    return calories