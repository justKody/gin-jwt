package helpers

import (
	"time"
)

// later
func GenerateAllTokens(email string, firstName string, lastName string, userType string, userId string) (signedToken string, signedRefreshToken string, err error) {
	claims := &model.SignedDetails{
		Email: email,
		First_name: firstName,
		Last_name: lastName,
		User_type: userType,
		User_id: userId,
	}
}