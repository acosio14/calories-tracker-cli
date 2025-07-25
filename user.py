# User Profile
class UserProfile:

    def __init__(
        self, name: str, age: int, gender: str, height_ft: float, weight_lbs: float,
    ) -> None:
        
        self.name = name
        self.age = age
        self.gender = gender
        self.height_ft = height_ft
        self.weight_lbs = weight_lbs


    def calculate_waist_to_hip_ratio( 
        self,
        waist_inches: float,
        hip_inches: float,
    ) -> float:
        
        return waist_inches / hip_inches
    

    def calculate_bmi_imperial(self) -> float:
        # Need to compare bmi to recommended numbers, show risk
        weight = self.weight_lbs
        height_squared = pow( self.height_ft * 12, 2)

        return round(( weight / height_squared ) * 703, 1)
         

    def calculate_body_fat_percentage(self, waist_hip_ratio: float) -> float:
        # simplified Navy method

        if self.gender == "male":
            body_fat_percentage = 76.76 * waist_hip_ratio - 49.82

        if self.gender == "female":
            body_fat_percentage = 70.41 * waist_hip_ratio - 58.25

        return round(body_fat_percentage,2)
    

    def calculate_maintenance_calories(self, activity_level) -> int:
        
        weight_kg = self.weight_lbs * 0.453       #lbs to kgs
        height_cm = self.height_ft * 30.48     #(convert ft to cm (To-Do: Convert 5'9" to inches first)

        # Basal Matabolic Rate (Harris-Benedict Equation)
        if self.gender == "male":
            bmr = 66.5 + (13.75 * weight_kg) + (5.003 * height_cm) - (6.75 * self.age)
        if self.gender == "female":
            bmr = 655.1 + (9.563 * weight_kg) + (1.850 * height_cm) - (4.676 * self.age)
        
        activity_factor = {
            'sedantary': 1.2,           # little to no exercise
            'lightly active': 1.375,    # light exercise/sports 1-3 days/week
            'moderate active': 1.55,    # moderate exercise/sports 3-5 days/week
            'very active': 1.725,       # hard exercise/sports 6-7 days a week
            'extra active': 1.9,        # very hard exercise/sports & physical job
        }

        activity_factor_value = activity_factor.get(activity_level.lower())

        if activity_factor_value is not None:
            maintenance_calories = activity_factor_value * bmr
        else:
            raise ValueError(f"{activity_level} is not a correct activity level.")

        return round(maintenance_calories)


    def calculate_macronutrients(self, diet_type: str) -> int:
        """ Function to calculate user's macros. """
        
        macro_calories = { 
            'Protein': 4.0,
            'Carbohydrates': 4.0,
            'Fats': 9.0
        }
        
        # Protein needs, grams per kg of body weight
        protein_intake_need = {
            'lose': 2.4,
            'gain': 2.2,
            'maintain': 2.0
        }
        protein_grams = round(
            (float(self.weight_lbs) * 0.453) * protein_intake_need.get(self.goal)
        )

        protein_calories = protein_grams * macro_calories.get('Protein')

        if diet_type == 'balanced':
            fats_calories = 0.20 * self.current_goal_calories # 20% of current cal are for fats 
            fats_grams = round(fats_calories / macro_calories.get('Fats'))

            carb_calories = self.current_goal_calories - (protein_calories + fats_calories)
            carbohydrates_grams = round(carb_calories / macro_calories.get('Carbohydrates'))

        elif diet_type == 'low-carb':
            carbohydrates_grams = 130.0
            carb_calories = carbohydrates_grams * macro_calories.get('Carbohydrates')

            fats_calories = self.current_goal_calories - (protein_calories + carb_calories)
            fats_grams = round(fats_calories / macro_calories.get('Fats'))

        return protein_grams, carbohydrates_grams, fats_grams
    

    def calculate_initial_calories(self, goal: str, desired_weight_delta: float):
        """ Function that calculates the initial recommended calories based on goal."""

        init_maintenance_calories = self.calculate_maintenance_calories()
        
        if goal == 'lose':
            recommend_calories = init_maintenance_calories - 500 * abs(desired_weight_delta)
        elif goal == 'gain':
            recommend_calories = init_maintenance_calories + 500 * abs(desired_weight_delta)
        else:
            recommend_calories = init_maintenance_calories

        return recommend_calories


    def calculate_recommended_calories(
        self, goal: str, 
        desired_weight_delta: float,
        current_calories: float,
        current_weight_delta: float,
        tracking: bool = False
    ) -> int:
        """
        Function to calculate caloires for user based on goal.

        Will be used to recalculate the maintenance calories of the user depending on how much
        their weight has changed. With new maintenance calories, it find a more accurate weekly calorie
        goal to reach the desired weight change.

        Args:
            my_user: UserProfile Class that contains all the person's current body metrics.
            tracking: Flag that tells wether user is tracking their weight.
        
        Returns
            recommended_calories: New calories to reach the goal weight change.
    
        """
        
        if tracking:
            maintenance_calories = current_calories - (500 * current_weight_delta)
        else:
            self.calculate_initial_calories(goal, desired_weight_delta)
        
        return maintenance_calories + ( self.desired_weight_delta * 500 ) #recommended calories
    


