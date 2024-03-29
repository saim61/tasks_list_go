package utils

type ErrorResponse struct {
	ErrorString string
	Message     string
}

type SuccessResponse struct {
	Message string
}

func NewErrorResponse(errorString string, message string) ErrorResponse {
	return ErrorResponse{
		ErrorString: errorString, Message: message,
	}
}

func NewSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{
		Message: message,
	}
}
