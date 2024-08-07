package restserver

import (
	"fmt"
	"log"

	"github.com/Beelzebub0/go-crud-boilerplate/src/business/entity"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

const (
	BAD_REQUEST           = "Bad Request"
	INTERNAL_SERVER_ERROR = "Internal Server Error"
	UNAUTHORIZED          = "Unauthorized"
)

var (
	errMsg = map[int]string{
		400: BAD_REQUEST,
		401: UNAUTHORIZED,
		500: INTERNAL_SERVER_ERROR,
	}
)

func (e *rest) HttpSuccess(c *gin.Context, statusCode int, resp interface{}, p *entity.Pagination) {
	var msg string
	switch statusCode {
	case 200:
		msg = "Success"
	case 201:
		msg = "Created"
	default:
		msg = "Success"
	}

	r := entity.HttpResponse{
		Code:       statusCode,
		Data:       resp,
		Pagination: p,
		Message:    msg,
	}

	c.JSON(statusCode, r)
}

func (e *rest) HttpError(c *gin.Context, statusCode int, err error) {
	var msg string

	msg, ok := errMsg[statusCode]
	if !ok {
		msg = "Unknown Error"
	}

	r := entity.HttpResponse{
		Code:    statusCode,
		Message: msg,
	}

	if err != nil {
		var (
			stacktrace      string
			laststacktrace  string
			stacktracecount int64
		)

		// Check if error implement stacktrace
		if err, ok := err.(stackTracer); ok {
			for _, f := range err.StackTrace() {
				stacktrace = stacktrace + fmt.Sprintf("%+s:%d\n", f, f)
				if stacktracecount == 0 {
					laststacktrace = stacktrace
					stacktracecount = stacktracecount + 1
				}
			}
		}

		// Get error message
		sysmsg := fmt.Sprint(err)

		// Print error on log
		log.Println(sysmsg)
		if stacktrace != "" {
			log.Println(stacktrace)
		}

		if e.conf.Gin.Mode == "debug" {
			// If on debug mode
			// geenerate error metadata for http response
			r.Metadata = &entity.HttpMetadata{
				Error: &entity.HttpError{
					SystemMessage:  sysmsg,
					Stacktrace:     laststacktrace,
					FullStacktrace: stacktrace,
				},
			}
		}

	}

	c.AbortWithStatusJSON(statusCode, r)
}

func (e *rest) SetCookie(c *gin.Context, value string, deleteCookie bool) {
	var maxAge int

	if deleteCookie {
		maxAge = -1
	} else {
		maxAge = e.conf.SessionCookie.MaxAge
	}

	c.SetCookie(e.conf.SessionCookie.Name, value, maxAge, "/", e.conf.SessionCookie.Domain, e.conf.SessionCookie.Secure, e.conf.SessionCookie.HttpOnly)
}
