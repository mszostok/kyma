package errorhandler

import (
	"encoding/json"

	"net/http"

	"github.com/mszostok/kyma/components/connector-service/internal/httpconsts"
	"github.com/mszostok/kyma/components/connector-service/internal/httperrors"
)

type ErrorHandler struct {
	Message string
	Code    int
}

func NewErrorHandler(code int, message string) *ErrorHandler {
	return &ErrorHandler{Message: message, Code: code}
}

func (eh *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseBody := httperrors.ErrorResponse{Code: eh.Code, Error: eh.Message}

	w.Header().Set(httpconsts.HeaderContentType, httpconsts.ContentTypeApplicationJson)
	w.WriteHeader(eh.Code)
	json.NewEncoder(w).Encode(responseBody)
}
