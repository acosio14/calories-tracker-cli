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
    current_weight_delta: float | int
    desired_weight_delta: float | int


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
        
        ''' Protein Needs g per kg of body weigth:
        Sedentary Adult, 0.8 g/kg, Moderate 1.2 - 1.6, Active, 1.6 - 2.0, Very Active 2.0 - 2.2
        Active Adult, 1.2 - 2.0 g/kg
        Endurance athletes, 1.2 - 1.6 g/kg
        Strength/Power athletes, 1.6 - 2.2 g/kg
        Fat loss (cutting), 2.0 - 2.4 g/kg
        Muslce gain (bulking), 1.6 - 2.2 g/kg
        Older Adults, 1.2 - 2.0 g/kg 
        '''



def main():
    ...

if __name__ == "__main__":
    main()