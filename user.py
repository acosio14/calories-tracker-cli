# User Profile

class User:

    def __init__(self, name, age, gender, height, weight):
        self.name = name
        self.age = age
        self.gender = gender
        self.height = height
        self.weight = weight

    def calculate_bmi_imperial(self):
        bmi = ( self.weight / pow( self.height, 2) ) * 703
         
        # Need to compare bmi to recommended numbers, show risk

        return bmi
    
    def calculate_body_fat_percentage(self, waist_hip_ratio):
        # simplified Navy method

        if self.gender is "male":
            body_fat_percentage = 76.76 * waist_hip_ratio - 49.82

        if self.gender is "female":
            body_fat_percentage = 70.41 * waist_hip_ratio - 58.25

        return body_fat_percentage