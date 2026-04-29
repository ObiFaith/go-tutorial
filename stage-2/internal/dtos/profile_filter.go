package dtos

type ProfileFilter struct {
	MinAge 								string
	MaxAge 								string
	Gender    					  string
	AgeGroup  					  string
	CountryId 					  string
	MinGenderProbability  string
	MinCountryProbability string
}
