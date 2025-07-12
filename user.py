# User Profile

class User:

    def __init__(self, name, age, height, weight):
        self.name = name
        self.age = age
        self.height = height
        self.weight = weight

    def calculate_bmi_imperial(self):
        bmi = ( self.weight / pow( self.height, 2) ) * 703
         
        # Need to compare bmi to recommended numbers, show risk

        return bmi
    
    def calculate_body_fat_percentage():
        pass

    