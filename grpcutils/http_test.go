package grpcutils

import (
	"fmt"
	"net/http"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TestGRPCErrorToHTTPStatus tests the GRPCErrorToHTTPStatus function.
func TestGRPCErrorToHTTPStatus(t *testing.T) {
	tests := []struct {
		name         string
		input        error
		expectedCode int
		expectedErr  error
	}{
		{"Nil error", nil, http.StatusOK, nil},
		{"gRPC OK", status.Error(codes.OK, "ok"), http.StatusOK, nil},
		{"gRPC Canceled", status.Error(codes.Canceled, "canceled"), http.StatusRequestTimeout, nil},
		{"gRPC Unknown", status.Error(codes.Unknown, "unknown"), http.StatusInternalServerError, nil},
		{"gRPC InvalidArgument", status.Error(codes.InvalidArgument, "invalid argument"), http.StatusBadRequest, nil},
		{"gRPC DeadlineExceeded", status.Error(codes.DeadlineExceeded, "deadline exceeded"), http.StatusGatewayTimeout, nil},
		{"gRPC NotFound", status.Error(codes.NotFound, "not found"), http.StatusNotFound, nil},
		{"gRPC AlreadyExists", status.Error(codes.AlreadyExists, "already exists"), http.StatusConflict, nil},
		{"gRPC PermissionDenied", status.Error(codes.PermissionDenied, "permission denied"), http.StatusForbidden, nil},
		{"gRPC Unauthenticated", status.Error(codes.Unauthenticated, "unauthenticated"), http.StatusUnauthorized, nil},
		{"gRPC ResourceExhausted", status.Error(codes.ResourceExhausted, "resource exhausted"), http.StatusTooManyRequests, nil},
		{"gRPC FailedPrecondition", status.Error(codes.FailedPrecondition, "failed precondition"), http.StatusBadRequest, nil},
		{"gRPC Aborted", status.Error(codes.Aborted, "aborted"), http.StatusConflict, nil},
		{"gRPC OutOfRange", status.Error(codes.OutOfRange, "out of range"), http.StatusBadRequest, nil},
		{"gRPC Unimplemented", status.Error(codes.Unimplemented, "unimplemented"), http.StatusNotImplemented, nil},
		{"gRPC Internal", status.Error(codes.Internal, "internal"), http.StatusInternalServerError, nil},
		{"gRPC Unavailable", status.Error(codes.Unavailable, "unavailable"), http.StatusServiceUnavailable, nil},
		{"gRPC DataLoss", status.Error(codes.DataLoss, "data loss"), http.StatusInternalServerError, nil},
		{"Unknown gRPC code", status.Error(codes.Code(999), "unknown code"), http.StatusInternalServerError, fmt.Errorf("unknown gRPC error code: %v", codes.Code(999))},
		{"Non-gRPC error", fmt.Errorf("some error"), http.StatusInternalServerError, fmt.Errorf("non-gRPC error: some error")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, err := GRPCErrorToHTTPStatus(tt.input)
			if statusCode != tt.expectedCode {
				t.Errorf("GRPCErrorToHTTPStatus(%v) = %v, expected %v", tt.input, statusCode, tt.expectedCode)
			}

			if tt.expectedErr == nil && err != nil || tt.expectedErr != nil && err == nil || tt.expectedErr != nil && err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("GRPCErrorToHTTPStatus(%v) error = %v, expected %v", tt.input, err, tt.expectedErr)
			}
		})
	}
}
