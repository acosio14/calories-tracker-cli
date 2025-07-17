from user import User

def create_goal_plan(my_user: User, goal: str, tracking: bool = False) -> int:
    # Currently only able to recalibrate to weight of 1

    current_weight_delta = 1
    current_calories = 2500
    
    desired_weight_delta = 1
    weight_delta = abs(desired_weight_delta)

    # weight delta need to be negative and positive.
    # Need to verify calculations
    # Might need to change dict to tuple. Don't think I need all 3
    calories_delta = 500 * desired_weight_delta
    weight_delta = {'lose': desired_weight_delta, 
                    'maintain': 0, 
                    'gain': desired_weight_delta,
                    }

    
    maintenance = my_user.calculate_maintance_calories()
    
    if tracking:
        maintenance = current_calories - (500 * current_weight_delta)
    
    calories = maintenance + ( weight_delta.get(goal) * calories_delta )
    
    return calories

def main():
    ...

if __name__ == "__main__":
    main()