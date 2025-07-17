from user import User

def create_goal_plan(goal: str, tracking: bool = True) -> int:
    # Currently only able to recalibrate to weight of 1

    delta = 500
    current_weight_delta = 1
    current_calories = 2500

    weight_delta = { 'lose': -1, 'maintain': 0, 'gain': 1 }

    maintenance = User.calculate_maintance_calories()
    
    if tracking:
        maintenance = current_calories - (500 * current_weight_delta)
    
    calories = maintenance + ( weight_delta.get(goal) * delta )
    
    return calories

def main():
    ...

if __name__ == "__main__":
    main()