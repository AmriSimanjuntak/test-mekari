package model

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewBaseResponse(c *fiber.Ctx, res *BaseResp) error {
	return c.Status(http.StatusOK).JSON(res)
}

func ResponseErr(c *fiber.Ctx, status int, msg string) error {
	return c.Status(http.StatusUnprocessableEntity).JSON(
		fiber.Map{
			"responseCode": fmt.Sprintf("%d", status),
			"responseDesc": msg,
			"data":         nil,
		})
}

func (br *BaseResp) OK(data interface{}) *BaseResp {
	return &BaseResp{ResponseCode: "00", ResponseDesc: "success", Data: data}
}

func (br *BaseResp) Err(err error) *BaseResp {
	br.ResponseDesc = err.Error()
	br.ResponseCode = fmt.Sprintf("%d", http.StatusUnprocessableEntity)
	return br
}
