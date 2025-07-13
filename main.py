from user import User



James = User("James", 24, "male", 6, 180, 36, 40, 'Lightly Active')

waist_to_hip_ratio = James.calculate_waist_to_hip_ratio()

print(f"{James.name} is {James.age} years old.\n")
print("His body metrics are:")
print(f"WHR (waist-to-hip ratio): {waist_to_hip_ratio}")
print(f"BMI (body-mass-index): {James.calculate_bmi_imperial()}")
print(f"Body fat percentage: {James.calculate_body_fat_percentage(waist_to_hip_ratio)}%")

print(f"\nEstimated maintance calories: {James.calculate_maintance_calories()}")

