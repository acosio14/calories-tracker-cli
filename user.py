# User Profile
class UserProfile:

    def __init__(
        self,
        name: str,
        age: int,
        gender: str,
        height_ft: float,
        weight_lbs: float,
    ) -> None:
        
        self.name = name
        self.age = age
        self.gender = gender
        self.height_ft = height_ft
        self.weight_lbs = weight_lbs

    activity_level: str
    goal: str = None
    diet_type: str             
    current_goal_calories: int
    current_weight_delta: int
    desired_weight_delta: int

    def calculate_waist_to_hip_ratio( 
        self,
        waist_inches: float,
        hip_inches: float,
    ) -> float:
        
        waist_to_hip_ratio = waist_inches / hip_inches

        return waist_to_hip_ratio
    
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
    
    def calculate_maintance_calories(self) -> int:
        
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

        activity_factor_value = activity_factor.get(self.activity_level.lower())

        if activity_factor_value is not None:
            maintance_calories = activity_factor_value * bmr
        else:
            raise ValueError(f"{self.activity_level} is not a correct activity level.")

        return round(maintance_calories)

    def calculate_recommended_calories( self, tracking: bool = False) -> int:
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
            maintenance_calories = self.calculate_maintance_calories()
        
        return maintenance_calories + ( self.desired_weight_delta * 500 ) #recommended calories
    
    def calculate_macronutrients(self) -> int:
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

        if self.diet_type == 'balanced':
            fats_calories = 0.20 * self.current_goal_calories # 20% of current cal are for fats 
            fats_grams = round(fats_calories / macro_calories.get('Fats'))

            carb_calories = self.current_goal_calories - (protein_calories + fats_calories)
            carbohydrates_grams = round(carb_calories / macro_calories.get('Carbohydrates'))

        elif self.diet_type == 'low-carb':
            carbohydrates_grams = 130.0
            carb_calories = carbohydrates_grams * macro_calories.get('Carbohydrates')

            fats_calories = self.current_goal_calories - (protein_calories + carb_calories)
            fats_grams = round(fats_calories / macro_calories.get('Fats'))

        return protein_grams, carbohydrates_grams, fats_grams

