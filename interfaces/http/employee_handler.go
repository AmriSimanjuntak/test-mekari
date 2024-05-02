package http

import (
	"Test_Mek/domain/model"
	"Test_Mek/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type employeeHandler struct {
	employeeUc usecase.IEmployeeUsecase
}

func NewEmployeeHandler(employeeUc usecase.IEmployeeUsecase) *employeeHandler {
	return &employeeHandler{employeeUc: employeeUc}
}
func (h *employeeHandler) ListEmployeeHandler(c *fiber.Ctx) error {
	req := model.Employee{}
	err := c.BodyParser(&req)
	if err != nil {
		rc, _ := strconv.Atoi("02")
		return model.ResponseErr(c, rc, "Invalid Req Data")
	}
	return model.NewBaseResponse(c, h.employeeUc.ListEmployeeUseCase(&req))
}

func (h *employeeHandler) AddEmployeeHandler(c *fiber.Ctx) error {
	req := model.Employee{}
	err := c.BodyParser(&req)
	if err != nil {
		rc, _ := strconv.Atoi("02")
		return model.ResponseErr(c, rc, "Invalid Req Data")
	}
	return model.NewBaseResponse(c, h.employeeUc.AddEmployeeUseCase(&req))
}

func (h *employeeHandler) EditEmployeeHandler(c *fiber.Ctx) error {
	req := model.Employee{}
	err := c.BodyParser(&req)
	if err != nil {
		rc, _ := strconv.Atoi("02")
		return model.ResponseErr(c, rc, "Invalid Req Data")
	}
	return model.NewBaseResponse(c, h.employeeUc.EditEmployeeUseCase(&req))
}

func (h *employeeHandler) DeleteEmployeeHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		rc, _ := strconv.Atoi("02")
		return model.ResponseErr(c, rc, "Invalid Req Data")
	}

	return model.NewBaseResponse(c, h.employeeUc.DeleteEmployeeUseCase(id))
}

func (h *employeeHandler) ViewEmployeeHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		rc, _ := strconv.Atoi("02")
		return model.ResponseErr(c, rc, "Invalid Req Data")
	}

	return model.NewBaseResponse(c, h.employeeUc.ViewEmployeeUseCase(id))
}
