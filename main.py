from user import UserProfile
from create_goal import HealthPlan

James = UserProfile("James", 24, "male", 6, 180, 36, 40, 'Lightly Active')

waist_to_hip_ratio = James.calculate_waist_to_hip_ratio()

print(f"{James.name} is {James.age} years old {James.gender}.\n")
print("His body metrics are:")
print(f"WHR (waist-to-hip ratio): {waist_to_hip_ratio}")
print(f"BMI (body-mass-index): {James.calculate_bmi_imperial()}")
print(f"Body fat percentage: {James.calculate_body_fat_percentage(waist_to_hip_ratio)}%")

print(f"\nEstimated maintance calories: {James.calculate_maintance_calories()}")

weight_delta = -1
goal_calories = James.calculate_maintance_calories() + (weight_delta * 500)

james_goal = HealthPlan('lose','balanced', goal_calories, weight_delta, weight_delta)

print(f"Recommended calories: {james_goal.calculate_recommended_calories(James)}")
protein, carbs, fats = james_goal.calculate_macronutrients(James)
print(
f"""
Macros 
  Protein: {protein}g
  Carbs: {carbs}g
  Fats: {fats}g
"""
)