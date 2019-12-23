package errors

import "github.com/plexmediamanager/service/errors"

const (
    ServiceID       errors.Service      =   3
)

var (
    NetworkAPIResponseError = errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeNetwork,
            ErrorNumber:    1,
        },
        Message:    "TMDB API returned error: Code (%d): %s",
    }
)
