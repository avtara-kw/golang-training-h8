package helpers

func IsProductiveAge(gender string, age int) bool {
	if gender == "pria" {
		return age > 17 && age < 60
	} else if gender == "wanita" {
		return age > 19 && age < 60
	} else {
		return false
	}
}
