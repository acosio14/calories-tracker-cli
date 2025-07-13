# User Profile
from dataclasses import dataclass

@dataclass
class User:
    name: str
    age: int
    gender: str
    height_ft: int
    weight_lbs: int
    waist_in: int
    hip_in: int


    def calculate_waist_to_hip_ratio(self):
        waist_to_hip_ratio = self.waist_in / self.hip_in

        return waist_to_hip_ratio
    
    def calculate_bmi_imperial(self):
        bmi = ( self.weight_lbs / pow( self.height_ft, 2) ) * 703
         
        # Need to compare bmi to recommended numbers, show risk

        return bmi
    
    def calculate_body_fat_percentage(self, waist_hip_ratio):
        # simplified Navy method

        if self.gender is "male":
            body_fat_percentage = 76.76 * waist_hip_ratio - 49.82

        if self.gender is "female":
            body_fat_percentage = 70.41 * waist_hip_ratio - 58.25

        return body_fat_percentage
    
    def calculate_maintance_calories(self):
        # Mifflin-St Jeor Equation
        
        weight_kg = self.weight_lbs * 0.453       #lbs to kgs
        height_cm = self.height_ft * 0.393701     #inches to cm (To-Do: Convert 5'9" to inches first)

        # Basal Matabolic Rate
        if self.gender is "male":
            bmr = (10 * weight_kg) + (6.25 * height_cm) - (5 * self.age) + 5
        if self.gender is "female":
            bmr = (10 * weight_kg) + (6.25 * height_cm) - (5 * self.age) - 161
        
        activity_factor = {
            'sedantary': 1.2,           #little to no exercise
            'lightly active': 1.375,    # light exercise/sports 1-3 days/week
            'moderate active': 1.55,    # moderate exercise/sports 3-5 days/week
            'Very Active': 1.725,       # hard exercise/sports 6-7 days a week
            'Extra Active': 1.9,        # very hard exercise/sports & physical job
        }

        maintance_calories = activity_factor * bmr

        return maintance_calories
