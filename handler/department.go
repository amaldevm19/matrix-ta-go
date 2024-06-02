package handler

import (
	"fmt"

	"github.com/amaldevm19/go_matrix_tna/database"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type Department struct {
	DepartmentId    string  `json:"department_id"`
	DepartmentName  string  `json:"department_name"`
	DepartmentCode  string  `json:"department_code"`
	TnaDepartmentId *string `json:"tna_department_id"`
}

func GetAllDepartments(c *fiber.Ctx) error {
	type Departments struct {
		Departments []Department `json:"departments"`
	}
	db := database.COSEC_DB
	rows, err := db.Query("SELECT DepartmentId, DepartmentName, DepartmentCode, TnaDepartmentId FROM Mx_DepartmentMst order by DepartmentId")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := Departments{}

	for rows.Next() {
		department := Department{}
		if err := rows.Scan(&department.DepartmentId, &department.DepartmentName, &department.DepartmentCode, &department.TnaDepartmentId); err != nil {
			return err
		}
		result.Departments = append(result.Departments, department)
	}
	return c.JSON(result)
}

func GetSingleDepartment(c *fiber.Ctx) error {
	db := database.COSEC_DB
	DepartmentId := c.Params("DepartmentId")
	department := Department{}
	row := db.QueryRow("SELECT * FROM Mx_DepartmentMst WHERE DepartmentId = ?", DepartmentId)
	if err := row.Scan(&department.DepartmentId, &department.DepartmentName, &department.DepartmentCode, &department.TnaDepartmentId); err != nil {
		new_err := &Response{Status: "failed", Error: err.Error()}
		return c.Status(200).JSON(new_err)
	}
	return c.JSON(department)
}

func AddNewDepartment(c *fiber.Ctx) error {
	db := database.COSEC_DB
	new_department := new(Department)

	if err := c.BodyParser(&new_department); err != nil {
		new_err := &Response{Status: "failed", Error: err.Error()}
		return c.Status(200).JSON(new_err)
	}

	_, err := db.Exec("INSERT INTO Mx_DepartmentMst (DepartmentName, DepartmentCode) VALUES ($1,$2)", new_department.DepartmentName, new_department.DepartmentCode)
	if err != nil {
		new_err := &Response{Status: "failed", Error: err.Error()}
		return c.Status(200).JSON(new_err)
	}

	new_succ := &Response{Status: "ok", Error: ""}
	return c.JSON(new_succ)
}

func UpdateDepartment(c *fiber.Ctx) error {
	db := database.COSEC_DB
	new_department := new(Department)
	DepartmentCode := c.Query("DepartmentCode")
	if err := c.BodyParser(&new_department); err != nil {
		new_err := &Response{Status: "failed", Error: err.Error()}
		return c.Status(200).JSON(new_err)
	}

	update := squirrel.Update("Mx_DepartmentMst").Where(squirrel.Eq{"DepartmentCode": DepartmentCode})

	if new_department.DepartmentName != "" {
		update = update.Set("DepartmentName", new_department.DepartmentName)
	}

	update = update.Set("TnaDepartmentId", new_department.TnaDepartmentId)

	if new_department.DepartmentCode != "" {
		update = update.Set("DepartmentCode", new_department.DepartmentCode)
	}
	sql, args, err := update.ToSql()
	if err != nil {
		new_err := &Response{Status: "failed", Error: err.Error()}
		return c.Status(200).JSON(new_err)
	}
	result, err := db.Exec(sql, args...)
	if err != nil {
		new_err := &Response{Status: "failed", Error: err.Error()}
		return c.Status(200).JSON(new_err)
	}

	if row_count, _ := result.RowsAffected(); row_count == 0 {
		new_err := &Response{Status: "failed", Error: fmt.Sprintf("Not found DepartmentCode : %s", DepartmentCode)}
		return c.Status(200).JSON(new_err)
	}

	new_succ := &Response{Status: "ok", Error: ""}
	return c.JSON(new_succ)
}

func DeleteDepartment(c *fiber.Ctx) error {
	db := database.COSEC_DB
	DepartmentCode := c.Params("DepartmentCode")

	res, _ := db.Exec("DELETE FROM Mx_DepartmentMst WHERE DepartmentCode = $1", DepartmentCode)

	if row_count, _ := res.RowsAffected(); row_count == 0 {
		new_err := &Response{Status: "failed", Error: fmt.Sprintf("Not found DepartmentCode : %s", DepartmentCode)}
		return c.Status(200).JSON(new_err)
	}
	new_succ := &Response{Status: "ok", Error: ""}
	return c.JSON(new_succ)
}
