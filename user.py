# User Profile
from dataclasses import dataclass

@dataclass
class UserProfile:
    name: str
    age: int
    gender: str
    height_ft: int
    weight_lbs: int
    waist_in: int
    hip_in: int
    activity_level: str


    def calculate_waist_to_hip_ratio(self):
        waist_to_hip_ratio = self.waist_in / self.hip_in

        return waist_to_hip_ratio
    
    def calculate_bmi_imperial(self):
        height_in = self.height_ft * 12

        bmi = ( self.weight_lbs / pow( height_in, 2) ) * 703
         
        # Need to compare bmi to recommended numbers, show risk

        return round(bmi,1)
    
    def calculate_body_fat_percentage(self, waist_hip_ratio):
        # simplified Navy method

        if self.gender == "male":
            body_fat_percentage = 76.76 * waist_hip_ratio - 49.82

        if self.gender == "female":
            body_fat_percentage = 70.41 * waist_hip_ratio - 58.25

        return round(body_fat_percentage,2)
    
    def calculate_maintance_calories(self):
        
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
