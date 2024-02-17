package exception

import "net/http"

var (
	ErrBadRequest     = NewAPIException(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).WithHttpCode(http.StatusBadRequest)
	ErrUnauthorized   = NewAPIException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)).WithHttpCode(http.StatusUnauthorized)
	ErrPermissionDeny = NewAPIException(http.StatusForbidden, http.StatusText(http.StatusForbidden)).WithHttpCode(http.StatusForbidden)
)
