from user import UserProfile
from dataclasses import dataclass

@dataclass
class HealthPlan:
    goal: str = 'lose' # lose, gain, maintain weight
    diet_type: str = 'balanced'
    current_goal_calories: int = 1500 # Calories to lose, gain (subtract by 1 lb = 500 cal)
    current_weight_delta: float | int = - 1 # Default of lose 1 lb/week
    desired_weight_delta: float | int = - 1


    def calculate_recommended_calories(
            self,
            my_user: UserProfile, 
            tracking: bool = False
    ) -> int:
        
        if tracking:
            maintenance_calories = self.current_goal_calories - (500 * self.current_weight_delta)
        else:
            maintenance_calories = my_user.calculate_maintance_calories()
        
        return maintenance_calories + ( self.desired_weight_delta * 500 ) #recommended calories

def main():
    ...

if __name__ == "__main__":
    main()