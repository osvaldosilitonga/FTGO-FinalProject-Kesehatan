package utils

import (
	"strings"

	"github.com/labstack/echo/v4"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcError(c echo.Context, err error) error {
	if e, ok := status.FromError(err); ok {
		errMsg := e.Message()
		index := strings.Index(errMsg, "desc")
		msg := strings.Join(strings.Fields(errMsg[index+1:]), " ")

		switch e.Code() {
		case codes.InvalidArgument:
			return ErrorMessage(c, &ApiBadRequest, msg)
		case codes.NotFound:
			return ErrorMessage(c, &ApiNotFound, msg)
		default:
			return ErrorMessage(c, &ApiInternalServer, "internal server error")
		}
	}
	return nil
}
