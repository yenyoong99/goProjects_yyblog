package exception

import "net/http"

var (
	ErrBadRequest     = NewAPIException(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).WithHttpCode(http.StatusBadRequest).WithMessage("Bad Request.")
	ErrUnauthorized   = NewAPIException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)).WithHttpCode(http.StatusUnauthorized).WithMessage("Please Login.")
	ErrPermissionDeny = NewAPIException(http.StatusForbidden, http.StatusText(http.StatusForbidden)).WithHttpCode(http.StatusForbidden).WithMessage("Access Denied.")
)
