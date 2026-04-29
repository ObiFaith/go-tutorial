package utils

func AgeGroup(age int) string {
	switch {
		case age < 13:
			return "child"
		case age < 20:
			return "teenager"
		case age < 60:
			return "adult"
		default:
			return "senior"
	}
}