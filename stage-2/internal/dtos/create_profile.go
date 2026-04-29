package dtos

import "github.com/google/uuid"

type CreateProfileRequest struct {
	Name string `json:"name" binding:"required"`
}

type ProfileResponse struct {
	Id              	 uuid.UUID  	 `json:"id"`
	Name               string 			 `json:"name"`
	Gender             string 			 `json:"gender"`
	GenderProbability  float64			 `json:"gender_probability"`
	Age                int    			 `json:"age"`
	AgeGroup           string 			 `json:"age_group"`
	CountryId          string 			 `json:"country_id"`
	CountryName        string 			 `json:"country_name"`
	CountryProbability float64			 `json:"country_probability"`
	CreatedAt          string 			 `json:"created_at"`
}
