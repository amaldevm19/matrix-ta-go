package router

import (
	"github.com/amaldevm19/go_matrix_tna/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupBranchRoutes(api fiber.Router) {

	branchApiRoutes := api.Group("/branch")
	branchApiRoutes.Post("/", handler.AddNewBranch)
	branchApiRoutes.Put("/", handler.UpdateBranch)
	branchApiRoutes.Delete("/:BranchCode", handler.DeleteBranch)

}
