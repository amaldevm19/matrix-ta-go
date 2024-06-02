package router

import (
	"github.com/amaldevm19/go_matrix_tna/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupDepartmentRoutes(api fiber.Router) {

	departmentApiRoutes := api.Group("/department")
	departmentApiRoutes.Get("/", handler.GetAllDepartments)
	departmentApiRoutes.Get("/:DepartmentId", handler.GetSingleDepartment)
	departmentApiRoutes.Post("/", handler.AddNewDepartment)
	departmentApiRoutes.Put("/", handler.UpdateDepartment)
	departmentApiRoutes.Delete("/:DepartmentCode", handler.DeleteDepartment)

}
