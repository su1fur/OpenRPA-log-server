package api

import (
	"net/http"

	"openRPA-log-server/app/models"

	"strings"

	"io"

	"github.com/gin-gonic/gin"
	"openRPA-log-server/app/lib/tool"
)

type Api struct {
	*gin.Context
	Request  *http.Request
	Response Response
	Args     Args
}

type Response struct {
	Status      int
	ContentType string
	writer      io.Writer
}

type Args struct {
	Message   string
	ErrorCode int
}

func (c *Api) SetMessage(s string) *Api {
	c.Args.Message = s
	return c
}

func (c *Api) GetMessage() string {
	message := c.Args.Message
	return message
}

func (c *Api) SetErrorCode(i int) *Api {
	c.Args.ErrorCode = i
	return c
}

func (c *Api) GetErrorCode() int {
	code := c.Args.ErrorCode
	return code
}

func (c *Api) GetUint64IDFromParam() (id uint64, err error) {
	id, err = tool.ParseUint64FromStr(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *Api) GetStrIDFromParam() (id string, err error) {
	id = c.Param("id")
	if id == "" {
		return "", err
	}
	return id, nil
}

// Bad Request Error を返すやつ
func (c *Api) HandleBadRequestError() *Api {
	c.Response.Status = http.StatusBadRequest
	return c.RenderJSON(nil)
}

// 作成しようとしたリソースが既にある場合
func (c *Api) HandleConflictError() *Api {
	c.Response.Status = http.StatusConflict
	c.SetMessage("作成しようとしたリソースがすでに存在しています")
	return c.RenderJSON(nil)
}

// 指定できないIDを指定したとき
func (c *Api) UnSpecifiedId() *Api {
	c.SetMessage("idを指定することはできません。")
	return c
}

// Not Found Error を返すやつ
func (c *Api) HandleNotFoundError() *Api {
	c.Response.Status = http.StatusNotFound
	return c.RenderJSON(nil)
}

// Internal Server Error を返すやつ
func (c *Api) HandleInternalServerError() *Api {
	c.Response.Status = http.StatusInternalServerError
	return c.RenderJSON(nil)
}

// 403 Forbidden を返す
func (c *Api) HandleForbiddenError() *Api {
	c.Response.Status = http.StatusForbidden
	return c.RenderJSON(nil)
}

func (c *Api) DeleteFailure() *Api {
	c.SetMessage("Record Delete Failure")
	return c
}

// Precondition Failed を返すやつ(存在しないリソースを更新した場合等)
func (c *Api) HandlePreconditionFailed() *Api {
	c.Response.Status = http.StatusPreconditionFailed
	return c.RenderJSON(nil)
}

// バージョンチェック
func (c *Api) CheckVersion() *Api {
	version := c.Request.Header.Get("App-Version")
	if !strings.Contains(version, "0.0.1") {
		c.Response.Status = http.StatusMethodNotAllowed
		return c.RenderJSON(nil)
	}
	return nil
}

type ResultJSON struct {
	Data      interface{} `json:"data"`
	ErrorCode int         `json:"error_code"`
	Message   string      `json:"message"`
}

func (c *Api) Success() *Api {
	return c.SetMessage("Record Success").RenderJSON(nil)
}

func (c *Api) RenderJSON(data interface{}) *Api {
	if c.Response.Status == 0 {
		c.Response.Status = http.StatusOK
	}

	if err := models.DB.Error; err != nil {
		return c.SetMessage(err.Error()).HandleInternalServerError()
	}

	resultJSON := ResultJSON{
		Data:      data,
		ErrorCode: c.GetErrorCode(),
		Message:   c.GetMessage(),
	}

	c.Status(c.Response.Status)
	c.IndentedJSON(c.Response.Status, resultJSON)
	return nil
}
