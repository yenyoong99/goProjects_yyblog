package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/exception"
)

// Success API, return data
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// Failed API return err, API Exception
func Failed(c *gin.Context, err error) {
	// Construct exception data
	var resp *exception.APIException
	if e, ok := err.(*exception.APIException); ok {
		resp = e
	} else {
		resp = exception.NewAPIException(
			500,
			http.StatusText(http.StatusInternalServerError),
		).WithMessage(err.Error()).WithHttpCode(500)
	}

	// return exception
	c.JSON(resp.HttpCode, resp)

	// abort
	c.Abort()
}
