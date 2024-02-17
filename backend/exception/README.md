### Exception

Why need exception?

1. decide the access token expired

Used Method:
```go
err := Biz_Call()
if err == TokenExpired {

}
```

String comparison may cause the errors

```go
access token expired %f minutes

hasPrefix("access token expired")
```

Usually designed as an exception code (Error Code):
```go
// err.ErrorCode == xxxx
if exception.IsTokenExpired(err) {

}
```

How to Design the Exception?  

Error Scenes
```go
func XXX() error
```

Error in go is an interface
```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

Error implementation provided in fmt package
```go
type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

type wrapErrors struct {
	msg  string
	errs []error
}

func (e *wrapErrors) Error() string {
	return e.msg
}

func (e *wrapErrors) Unwrap() []error {
	return e.errs
}
```

How to customize the exception?

```go
func NewAPIException(code int, msg string) *APIException {
	return &APIException{
		code: code,
		msg:  msg,
	}
}

// APIException error customize
type APIException struct {
	code int
	msg  string
}

func (e *APIException) Error() string {
	return e.msg
}

func (e *APIException) Code() int {
	return e.code
}
```

Define the exceptions

1. customize TokenExpired 5000
```go
// this module to define the customize exception error
// token expired %f minitues
// ErrXXXXX to define the customize error, Convenient to quick search in package
var (
	ErrAccessTokenExpired  = exception.NewAPIException(5000, "access token expired")
	ErrRefreshTokenExpired = exception.NewAPIException(5001, "refresh token expired")
)
```

2. use customize exception
```go
if aDelta > 0 {
    return ErrAccessTokenExpired.WithMessagef("access token expired %f minutes", aDelta)
}
```

#### Determine whether exceptions are equal

Determine exceptions based on Code
```go
if e, ok := err.(*exception.APIException); ok {
    t.Log(e.String())
    // Determine whether the exception is a TokenExpired exception
    if e.Code == token.ErrAccessTokenExpired.Code {
        t.Log(e.String())
    }
}
```

```go
// exception.IsException(err, token.ErrAccessTokenExpired)
// Give a method for exceptions judgment
func IsException(err error, e *APIException) bool {
	if targe, ok := err.(*APIException); ok {
		return targe.Code == e.Code
	}

	return false
}
```