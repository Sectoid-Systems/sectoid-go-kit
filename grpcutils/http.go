package grpcutils

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCErrorToHTTPStatus converts a gRPC error code into the corresponding HTTP status code.
// If there's an error converting, the status will be 500.
//
// See: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go#L16
func GRPCErrorToHTTPStatus(err error) (int, error) {
	if err == nil {
		return http.StatusOK, nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("non-gRPC error: %v", err)
	}

	switch st.Code() {
	case codes.OK:
		return http.StatusOK, nil
	case codes.Canceled:
		return http.StatusRequestTimeout, nil
	case codes.Unknown:
		return http.StatusInternalServerError, nil
	case codes.InvalidArgument:
		return http.StatusBadRequest, nil
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout, nil
	case codes.NotFound:
		return http.StatusNotFound, nil
	case codes.AlreadyExists:
		return http.StatusConflict, nil
	case codes.PermissionDenied:
		return http.StatusForbidden, nil
	case codes.Unauthenticated:
		return http.StatusUnauthorized, nil
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests, nil
	case codes.FailedPrecondition:
		// This deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP status
		return http.StatusBadRequest, nil
	case codes.Aborted:
		return http.StatusConflict, nil
	case codes.OutOfRange:
		return http.StatusBadRequest, nil
	case codes.Unimplemented:
		return http.StatusNotImplemented, nil
	case codes.Internal:
		return http.StatusInternalServerError, nil
	case codes.Unavailable:
		return http.StatusServiceUnavailable, nil
	case codes.DataLoss:
		return http.StatusInternalServerError, nil
	default:
		return http.StatusInternalServerError, fmt.Errorf("unknown gRPC error code: %v", st.Code())
	}
}
