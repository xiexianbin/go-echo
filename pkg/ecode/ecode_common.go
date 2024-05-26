// https://github.com/golang/go/blob/go1.22.3/src/net/http/Code.go
package ecode

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

var codeCommonMessages = map[Code]string{
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
