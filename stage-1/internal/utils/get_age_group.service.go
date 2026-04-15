package utils

func GetAgeGroup(age int) string {
	if age < 13 {
		return "child"
	} else if age >= 13 && age < 20 {
		return "teenager"
	} else if age >= 20 && age < 60 {
		return "adult"
	} else {
		return "senior"
	}
}