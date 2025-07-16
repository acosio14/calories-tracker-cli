from user import User

def create_goal_plan(goal: str, tracking: bool = True) -> int:
    maintance = User.calculate_maintance_calories()
    
    if not tracking:
        calories_delta = 500
    else:
        ...
        # linear regression of table, data, stored values.

    goals = { 'gain': 1, 'lose': -1, 'maintain': 0}
    
    calories = maintance + ( goals.get(goal) * calories_delta )
    
    return calories