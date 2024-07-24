
# grpcutils Package

The `grpcutils` package provides utility functions for working with gRPC errors and converting them into corresponding HTTP status codes.

## Functions

### GRPCErrorToHTTPStatus

The `GRPCErrorToHTTPStatus` function converts a gRPC error code into the corresponding HTTP status code. If there's an error converting, the status will be 500.

```go
func GRPCErrorToHTTPStatus(err error) (int, error)
```

#### Parameters

- **err error**: The gRPC error to convert.

#### Returns

- **int**: The corresponding HTTP status code.
- **error**: An error if the conversion fails.

#### Mapping of gRPC codes to HTTP status codes

- **codes.OK**: `http.StatusOK` (200)
- **codes.Canceled**: `http.StatusRequestTimeout` (408)
- **codes.Unknown**: `http.StatusInternalServerError` (500)
- **codes.InvalidArgument**: `http.StatusBadRequest` (400)
- **codes.DeadlineExceeded**: `http.StatusGatewayTimeout` (504)
- **codes.NotFound**: `http.StatusNotFound` (404)
- **codes.AlreadyExists**: `http.StatusConflict` (409)
- **codes.PermissionDenied**: `http.StatusForbidden` (403)
- **codes.Unauthenticated**: `http.StatusUnauthorized` (401)
- **codes.ResourceExhausted**: `http.StatusTooManyRequests` (429)
- **codes.FailedPrecondition**: `http.StatusBadRequest` (400)
- **codes.Aborted**: `http.StatusConflict` (409)
- **codes.OutOfRange**: `http.StatusBadRequest` (400)
- **codes.Unimplemented**: `http.StatusNotImplemented` (501)
- **codes.Internal**: `http.StatusInternalServerError` (500)
- **codes.Unavailable**: `http.StatusServiceUnavailable` (503)
- **codes.DataLoss**: `http.StatusInternalServerError` (500)

#### Example

```go
grpcErr := status.Error(codes.NotFound, "resource not found")
httpStatus, err := grpcutils.GRPCErrorToHTTPStatus(grpcErr)
if err != nil {
    log.Fatalf("Error converting gRPC error: %v", err)
}
fmt.Printf("HTTP Status: %d", httpStatus)
```

### References

For more details, see the [gRPC Gateway Errors documentation](https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go#L16).
