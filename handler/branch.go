package handler

import (
	"fmt"

	"github.com/amaldevm19/go_matrix_tna/database"
	"github.com/amaldevm19/go_matrix_tna/helpers"

	"github.com/gofiber/fiber/v2"
)

type Branch struct {
	BranchName string `json:"branch_name"`
	BranchCode string `json:"branch_code"`
}

func AddNewBranch(c *fiber.Ctx) error {
	new_branch := new(Branch)

	if err := c.BodyParser(&new_branch); err != nil {
		newErr := &Response{Status: "failed", Error: "Invalid request"}
		return c.Status(200).JSON(newErr)
	}

	// Validate that branch_name and branch_code are not empty
	if new_branch.BranchName == "" || new_branch.BranchCode == "" {
		newErr := &Response{Status: "failed", Error: "Values cannot be empty"}
		return c.Status(200).JSON(newErr)
	}

	status, response_text, resp_err := helpers.InsertItem(fmt.Sprintf("branch?action=set;code=%s;name=%s", new_branch.BranchCode, new_branch.BranchName))
	if !status {
		if resp_err != nil {
			new_err := &Response{Status: "failed", Error: resp_err.Error()}
			return c.Status(200).JSON(new_err)
		}
		new_err := &Response{Status: "failed", Error: response_text}
		return c.Status(200).JSON(new_err)
	}
	db := database.TNA_PROXY_DB
	_, err := db.Exec("INSERT INTO Px_BranchMst (BranchName, BranchCode) VALUES ($1,$2)", new_branch.BranchName, new_branch.BranchCode)
	if err != nil {
		new_err := &Response{Status: "failed", Error: "Failed to store in T&A ID in local server"}
		return c.Status(200).JSON(new_err)
	}

	success_response := &Response{Status: "ok", Error: ""}
	return c.JSON(success_response)
}

func UpdateBranch(c *fiber.Ctx) error {
	new_branch := new(Branch)
	// BranchCode := c.Query("BranchCode")
	if err := c.BodyParser(&new_branch); err != nil {
		new_err := &Response{Status: "failed", Error: err.Error()}
		return c.Status(200).JSON(new_err)
	}

	new_succ := &Response{Status: "ok", Error: ""}
	return c.JSON(new_succ)
}

func DeleteBranch(c *fiber.Ctx) error {
	db := database.COSEC_DB
	BranchCode := c.Params("BranchCode")

	res, _ := db.Exec("DELETE FROM Mx_BranchMst WHERE BranchCode = $1", BranchCode)

	if row_count, _ := res.RowsAffected(); row_count == 0 {
		new_err := &Response{Status: "failed", Error: fmt.Sprintf("Not found BranchCode : %s", BranchCode)}
		return c.Status(200).JSON(new_err)
	}
	new_succ := &Response{Status: "ok", Error: ""}
	return c.JSON(new_succ)
}
