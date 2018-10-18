package httperrors

import (
	"net/http"

	"github.com/mszostok/kyma/components/gateway/internal/apperrors"
)

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func errorCodeToHttpStatus(code int) int {
	switch code {
	case apperrors.CodeInternal:
		return http.StatusInternalServerError
	case apperrors.CodeNotFound:
		return http.StatusNotFound
	case apperrors.CodeAlreadyExists:
		return http.StatusConflict
	case apperrors.CodeWrongInput:
		return http.StatusBadRequest
	case apperrors.CodeUpstreamServerCallFailed:
		return http.StatusBadGateway
	default:
		return http.StatusInternalServerError
	}
}

func AppErrorToResponse(appError apperrors.AppError) (status int, body ErrorResponse) {
	httpCode := errorCodeToHttpStatus(appError.Code())
	return httpCode, ErrorResponse{httpCode, appError.Error()}
}
