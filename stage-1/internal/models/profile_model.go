package models

import "gorm.io/gorm"

type Profile struct {
	BaseModel

	Name 							 string 				`json:"name" gorm:"uniqueIndex:idx_profile_name"`
	Gender             string  				`json:"gender"`
	SampleSize         int     				`json:"sample_size"`
	Age                int     				`json:"age"`
	AgeGroup           string  				`json:"age_group"`
	CountryID          string  				`json:"country_id"`
	GenderProbability  float64 				`gorm:"not null;check:gender_probability >= 0 AND gender_probability <= 1"`
	CountryProbability float64 				`gorm:"not null;check:country_probability >= 0 AND country_probability <= 1"`
	DeletedAt 				 gorm.DeletedAt `gorm:"index"`
}
