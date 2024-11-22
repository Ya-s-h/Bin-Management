package validation

import (
	connection "assignment/renie/db"
	model "assignment/renie/models"
)

type ErrorResponse struct {
	Error            bool
	FailedField      string
	FailedFieldValue string
}

func DuplicateEmailAddress(data *model.User) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	db := connection.ConnectToDb()
	var user model.User
	res := db.First(&user, "email = ?", data.Email)
	if res.Error == nil {
		var elem ErrorResponse
		elem.FailedField = "email"
		elem.FailedFieldValue = data.Email
		// elem.Tag = res.Tag()           // Export struct tag
		// elem.Value = res.Value()       // Export field value
		elem.Error = true
		validationErrors = append(validationErrors, elem)
	}
	return validationErrors
}
