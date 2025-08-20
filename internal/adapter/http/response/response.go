package response

import (
	"errors"
	"go-log-saas/internal/adapter/http/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal error")
	// ErrDataNotFound is an error for when requested data is not found
	ErrDataNotFound = errors.New("data not found")
	// ErrNoUpdatedData is an error for when no data is provided to update
	ErrNoUpdatedData = errors.New("no data to update")
	// ErrConflictingData is an error for when data conflicts with existing data
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")
	// ErrForbidden is an error for when the user is forbidden to access the resource
	ErrForbidden = errors.New("user is forbidden to access the resource")
)

// newResponse is a helper function to create a response body
func newResponse(id string, status string) dto.IngestResponse {
	return dto.IngestResponse{
		ID:     id,
		Status: status,
	}
}

// errorStatusMap is a map of defined error messages and their corresponding http status codes
var errorStatusMap = map[error]int{
	ErrInternal:        http.StatusInternalServerError,
	ErrDataNotFound:    http.StatusNotFound,
	ErrConflictingData: http.StatusConflict,
	ErrForbidden:       http.StatusForbidden,
	ErrNoUpdatedData:   http.StatusBadRequest,
}

// validationError sends an error response for some specific request validation error
func ValidationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func HandleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

// handleAbort sends an error response and aborts the request with the specified status code and error message
func HandleAbort(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.AbortWithStatusJSON(statusCode, errRsp)
}

func HandleValidationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

// parseError parses error messages from the error object and returns a slice of error messages
func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

// errorResponse represents an error response body format
type errorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

// newErrorResponse is a helper function to create an error response body
func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

// handleSuccess sends a success response with the specified status code and optional data
func HandleSuccess(ctx *gin.Context, id, status string) {
	rsp := newResponse(id, status)
	ctx.JSON(http.StatusOK, rsp)
}
