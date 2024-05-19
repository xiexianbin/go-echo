// https://github.com/golang/go/blob/go1.22.3/src/net/http/Code.go
package e

import (
	"net/http"
	"strings"
)

// https://pkg.go.dev/golang.org/x/tools/cmd/stringer

//go:generate stringer -type=Code
type Code int

const (
	// common err code biz id is 100000
	CodeOK Code = 100200 + iota
	CodeCreated
	CodeAccepted
	CodeNonAuthoritativeInfo
	CodeNoContent

	CodeBadRequest Code = 100400 + iota
	CodeUnauthorized
	CodePaymentRequired
	CodeForbidden
	CodeNotFound // 100404
	CodeMethodNotAllowed
	CodeNotAcceptable
	CodeProxyAuthRequired
	CodeRequestTimeout
	CodeConflict
	CodeUnKnown
)

var codeMessageMap = map[Code]string{
	CodeOK: "OK",

	CodeCreated:              "Created",
	CodeAccepted:             "Accepted",
	CodeNonAuthoritativeInfo: "Non-Authoritative Information",
	CodeNoContent:            "No Content",

	CodeBadRequest:        "Bad Request",
	CodeUnauthorized:      "Unauthorized",
	CodePaymentRequired:   "Payment Required",
	CodeForbidden:         "Forbidden",
	CodeNotFound:          "Not Found",
	CodeMethodNotAllowed:  "Method Not Allowed",
	CodeNotAcceptable:     "Not Acceptable",
	CodeProxyAuthRequired: "Proxy Authentication Required",
	CodeRequestTimeout:    "Request Timeout",
	CodeConflict:          "Conflict",

	CodeUnKnown: "UnKnown",
}

type icode interface {
	HTTPCode() int
	Message() string
}

func (c Code) HTTPCode() int {
	if c == CodeOK || strings.Contains(c.String(), "OK") { // 200
		return http.StatusOK
	} else if strings.Contains(c.String(), "Created") { // 201
		return http.StatusCreated
	} else if strings.Contains(c.String(), "Accepted") { // 202
		return http.StatusAccepted
	} else if strings.Contains(c.String(), "BadRequest") { // 400
		return http.StatusBadRequest
	} else if strings.Contains(c.String(), "Unauthorized") { // 401
		return http.StatusUnauthorized
	} else if strings.Contains(c.String(), "Forbidden") { // 403
		return http.StatusForbidden
	} else if strings.Contains(c.String(), "NotFound") { // 404
		return http.StatusNotFound
	} else if strings.Contains(c.String(), "Conflict") { // 409
		return http.StatusConflict
	}

	return http.StatusBadRequest
}

func (c Code) Message() string {
	if v, ok := codeMessageMap[c]; ok {
		return v
	}
	return codeMessageMap[CodeUnKnown]
}
