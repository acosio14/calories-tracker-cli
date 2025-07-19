from user import UserProfile
from dataclasses import dataclass

@dataclass
class HealthPlan:
    goal: str
    diet_type: str
    current_calories: int
    current_weight_delta: float | int
    desired_weight_delta: float | int


    def calculate_recommended_calories(
            self,
            my_user: UserProfile,
            goal: str, 
            tracking: bool = False
    ) -> int:
        
        weight_delta = abs(self.desired_weight_delta)

        # weight delta need to be negative and positive.
        # Need to verify calculations
        # Might need to change dict to tuple. Don't think I need all 3
        calories_delta = 500 * self.desired_weight_delta
        weight_delta = {'lose': self.desired_weight_delta, 
                        'maintain': 0, 
                        'gain': self.desired_weight_delta,
                        }

        
        maintenance = my_user.calculate_maintance_calories()
        
        if tracking:
            maintenance = self.current_calories - (500 * self.current_weight_delta)
        
        return maintenance + ( weight_delta.get(goal) * calories_delta )

def main():
    ...

if __name__ == "__main__":
    main()