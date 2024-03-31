package utils

type ErrorResponse struct {
	ErrorCode string
	ErrorString string
	Message     string
}

type SuccessResponse struct {
	Message string
}

func NewErrorResponse(errorCode string, errorString string, message string) ErrorResponse {
	return ErrorResponse{
		ErrorCode: errorCode, ErrorString: errorString, Message: message,
	}
}

func NewSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{
		Message: message,
	}
}
