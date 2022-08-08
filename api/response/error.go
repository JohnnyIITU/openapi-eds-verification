package response

import "openapi-eds-verification/log"

type ErrorResponse struct {
	Message string
}

func NewErrorResponse(message string) *ErrorResponse {
	log.GetLogger().Warningf("Error[%d] : %s", message)

	return &ErrorResponse{
		Message: message,
	}
}
