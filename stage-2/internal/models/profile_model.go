package models

type Profile struct {
	BaseModel

	Name 							 string 				`json:"name" gorm:"uniqueIndex:idx_profile_name"`
	Gender             string  				`json:"gender"`
	Age                int     				`json:"age"`
	AgeGroup           string  				`json:"age_group"`
	CountryId          string  				`json:"country_id"`
	CountryName        string  				`json:"country_name"`
	GenderProbability  float64 				`json:"gender_probability" gorm:"not null;check:gender_probability >= 0 AND gender_probability <= 1"`
	CountryProbability float64 				`json:"country_probability" gorm:"not null;check:country_probability >= 0 AND country_probability <= 1"`
}
