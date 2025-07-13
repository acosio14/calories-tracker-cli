from user import User



James = User("James", 24, "male", 6, 180, 36, 40)

waist_to_hip_ratio = James.calculate_waist_to_hip_ratio()

print(f"{James.name} is {James.age} years old")
print("His body metrics are:")
print(f"WHR (waist-to-hip ratio): {waist_to_hip_ratio}")
print(f"BMI (body-mass-index): {James.calculate_bmi_imperial()}")
print(f"Body fat %: {James.calculate_body_fat_percentage(waist_to_hip_ratio)}")

print(f"Estimated maintance calories: {James.calculate_maintance_calories()}")

