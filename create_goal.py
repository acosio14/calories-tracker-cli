from user import User

def create_goal_plan(goal: str, tracking: bool = True) -> int:
    
    delta = 500
    weight_delta = 1
    current_maintenance = 2500

    maintenance = User.calculate_maintance_calories()
    
    if tracking:
        maintenance = current_maintenance - (500 * weight_delta)

    goals = { 'gain': 1, 'lose': -1, 'maintain': 0}
    
    calories = maintenance + ( goals.get(goal) * delta )
    
    return calories

def main():
    ...

if __name__ == "__main__":
    main()