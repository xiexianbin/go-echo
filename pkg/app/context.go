// Copyright 2024 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/xiexianbin/go-echo-demo/pkg/code"
)

type CustomeContext struct {
	echo.Context

	Errors []error

	// Response
	xresp Response
}

// CatchError temporary cache error
func (c *CustomeContext) CatchError(err error) {
	if err != nil {
		c.Errors = append(c.Errors, err)
	}
}

// HasErr just is ready has error
func (c *CustomeContext) HasError() bool {
	return len(c.Errors) > 0
}

// Must use after HasErr() == true, return error
func (c *CustomeContext) Must(codes ...code.Code) error {
	if !c.HasError() {
		panic("no errors found")
	}
	_code := code.CodeBadRequest
	if len(codes) > 0 {
		_code = codes[0]
	}
	msgs := make([]string, len(c.Errors))
	for _, err := range c.Errors {
		msgs = append(msgs, err.Error())
	}
	return c.Context.JSON(_code.HTTPCode(), Response{
		Code:    _code,
		Message: fmt.Sprintf("%s: %s", _code.Message(), msgs),
	})
}

type Page struct {
	Page    int            `json:"page"`
	Size    int            `json:"size"`
	Orderby map[string]any `json:"orderby"`
}

// ParsePage get page info
func (c *CustomeContext) ParsePage() *Page {
	return nil
}

func (c *CustomeContext) WithCode(code code.Code) *CustomeContext {
	c.xresp.Code = code
	return c
}

func (c *CustomeContext) WithCodeOK() *CustomeContext {
	c.xresp.Code = code.CodeOK
	return c
}

func (c *CustomeContext) WithCodeCreated() *CustomeContext {
	c.xresp.Code = code.CodeCreated
	return c
}

func (c *CustomeContext) WithCodeAccepted() *CustomeContext {
	c.xresp.Code = code.CodeCreated
	return c
}
func (c *CustomeContext) WithCodeNoContent() *CustomeContext {
	c.xresp.Code = code.CodeNoContent
	return c
}

func (c *CustomeContext) WithCodeBadRequest() *CustomeContext {
	c.xresp.Code = code.CodeBadRequest
	return c
}

func (c *CustomeContext) WithCodeUnauthorized() *CustomeContext {
	c.xresp.Code = code.CodeUnauthorized
	return c
}

func (c *CustomeContext) WithCodeForbidden() *CustomeContext {
	c.xresp.Code = code.CodeForbidden
	return c
}

func (c *CustomeContext) WithCodeNotFound() *CustomeContext {
	c.xresp.Code = code.CodeNotFound
	return c
}

func (c *CustomeContext) WithMessage(msg string) *CustomeContext {
	c.xresp.Message = fmt.Sprintf(
		"%s: %s",
		c.xresp.Code.Message(),
		msg)
	return c
}

func (c *CustomeContext) WithMessagef(format string, args ...any) *CustomeContext {
	c.xresp.Message = fmt.Sprintf(
		"%s: %s",
		c.xresp.Code.Message(),
		fmt.Sprintf(format, args...))
	return c
}

func (c *CustomeContext) WithErrorMessagef(err error, format string, a ...any) *CustomeContext {
	c.xresp.Message = fmt.Sprintf(
		"%s: %s, err: %s",
		c.xresp.Code.Message(),
		fmt.Sprintf(format, a...),
		err)
	return c
}

func (c *CustomeContext) WithData(data ...any) *CustomeContext {
	if len(data) == 1 {
		c.xresp.Data = data[0]
	} else if len(data) > 1 {
		c.xresp.Data = data
	}
	return c
}

// RJSON resturn JSON to client
func (c *CustomeContext) RJSON() error {
	if c.xresp.Reason == "" {
		c.xresp.Reason = strings.Replace(c.xresp.Code.String(), "Code", "", 1)
	}
	return c.Context.JSON(c.xresp.Code.HTTPCode(), c.xresp)
}
