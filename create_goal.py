from user import UserProfile
from dataclasses import dataclass

@dataclass
class HealthPlan:
    """
    A class for calculating recommended health metrics for desired goal.
    """

    goal: str
    diet_type: str
    current_goal_calories: int
    current_weight_delta: float
    desired_weight_delta: float


    def calculate_recommended_calories(
            self,
            my_user: UserProfile, 
            tracking: bool = False
    ) -> int:
        """
        Function to calculate caloires for user based on goal.

        Will be used to recalculate the maintenance calories of the user depending on how much
        their weight has changed. With new maintance calories, it find a more accurate weekly calorie
        goal to reach the desired weight change.

        Args:
            my_user: UserProfile Class that contains all the person's current body metrics.
            tracking: Flag that tells wether user is tracking their weight.
        
        Returns
            recommended_calories: New calories to reach the goal weight change.
    
        """
        
        if tracking:
            maintenance_calories = self.current_goal_calories - (500 * self.current_weight_delta)
        else:
            maintenance_calories = my_user.calculate_maintance_calories()
        
        return maintenance_calories + ( self.desired_weight_delta * 500 ) #recommended calories
    
    def calculate_macronutrients(
            self,
            my_user: UserProfile
    ) -> int:
        """ Function to calculate user's macros. """
        
        macro_calories = { 
            'Protien': 4,
            'Carbohydrates': 4,
            'Fats': 9
        }
        
        # Protein needs, grams per kg of body weight
        protein_intake_need = {
            'lose': 2.4,
            'gain': 2.2,
            'maintain': 2.0
        }

        protein_grams = (my_user.weight_lbs * 0.453) * protein_intake_need(self.goal)
        
        protein_calories = protein_grams * macro_calories.get('Protein')

        if self.diet_type == 'balanced':
            fats_calories = 0.20 * self.current_goal_calories # 20% of current cal are for fats 
            fats_grams = fats_calories / macro_calories.get('Fats')

            carb_calories = self.current_goal_calories - (protein_calories + fats_calories)
            carbohydrates_grams = carb_calories / macro_calories.get('Carbohydrates')

        elif self.diet_type == 'low-carb':
            carbohydrates_grams = 130
            carb_calories = carbohydrates_grams * macro_calories.get('Carbohydrates')

            fats_calories = self.current_goal_calories - (protein_calories + carb_calories)
            fats_grams = fats_calories / macro_calories.get('Fats')

        
        return protein_grams, carbohydrates_grams, fats_grams


def main():
    ...

if __name__ == "__main__":
    main()
