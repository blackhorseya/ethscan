package er

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddErrorHandlingMiddleware global handle *gin.Context error middleware
func AddErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()
			c.Errors = c.Errors[:0]

			switch err.Err.(type) {
			case *Error:
				appError := err.Err.(*Error)
				c.AbortWithStatusJSON(appError.Status, appError)
				break
			default:
				errUnknown := New(http.StatusInternalServerError, 50099, "Internal server error", "Unknown error")
				c.AbortWithStatusJSON(errUnknown.Status, errUnknown)
				break
			}
		}()

		c.Next()
	}
}
