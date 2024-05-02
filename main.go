package testmek

import (
	"Test_Mek/constant"
	"Test_Mek/domain/model"
	"Test_Mek/interfaces/http"
	"Test_Mek/usecase"

	"Test_Mek/infrastructure/persistance"

	"github.com/gofiber/fiber/v2"
)

func main() {

	repos, err := persistance.NewRepositories()
	if err != nil {

	}
	constant.Repos = repos
	// start init usecase
	employeeUsecase := usecase.NewEmployeeUsecase(repos.Db, repos.EmployeeRepo)

	// end init usecase

	// start init handler
	employeeHandler := http.NewEmployeeHandler(employeeUsecase)

	// init server with config
	cfgFiber := fiber.Config{
		BodyLimit: 5 * 1024 * 1024, // 100 mb limit upload
	}
	app := fiber.New(cfgFiber)
	uploademployee := app.Group("employee")
	// add health check to bulk-upload-service
	uploademployee.Get("healthcheck", func(c *fiber.Ctx) error {

		return model.NewBaseResponse(c, new(model.BaseResp).OK("OK"))
	})

	v1 := uploademployee.Group("v1")
	{
		v1.Put("list", employeeHandler.EditEmployeeHandler)
		v1.Post("add", employeeHandler.AddEmployeeHandler)
		v1.Put("edit", employeeHandler.EditEmployeeHandler)
		v1.Get("view", employeeHandler.ViewEmployeeHandler)
		v1.Get("delete", employeeHandler.DeleteEmployeeHandler)

	}

	// run server

	app.Listen(":" + "8080")
}
