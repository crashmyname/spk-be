package services

import (
	"bytes"
	"errors"
	"spk-be/internal/database"
	"spk-be/internal/dto"
	"spk-be/internal/models"
	"spk-be/internal/repositories"
	"spk-be/internal/utils"
	"strconv"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func GetAllUser() ([]models.User, error) {
	return repositories.GetAllUser()
}

func GetUserPaginate(limit int, offset int, search string) ([]models.User, int64, error) {
	return repositories.GetPaginateUser(limit, offset, search)
}

func GetUserByID(id int) (*models.User, error) {
	return repositories.GetUserByID(id)
}

func GetUserName() ([]dto.UserNameDTO, error) {
	return repositories.GetUserName()
}

func CreateUser(name, username, email, password, section, role string) (*models.User, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		UUID:     uuid.NewString(),
		Name:     name,
		Username: username,
		Password: hashed,
		Email:    email,
		Section:  section,
		Role:     role,
	}

	err = repositories.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(id int, data map[string]interface{}) (*models.User, error) {
	return repositories.UpdateUser(id, data)
}

func DeleteUser(id int) (*models.User, error) {
	return repositories.DeleteUser(id)
}

func IsUsernameExsist(username string) bool {
	var user models.User
	err := database.DB.Where("username= ?", username).First(&user).Error

	if err != nil {
		return !errors.Is(err, gorm.ErrRecordNotFound)
	}

	return true
}

func ExportExcelUser() (*excelize.File, error) {
	f := excelize.NewFile()
	sheet := "Users"

	f.SetSheetName("Sheet1", sheet)
	index, _ := f.GetSheetIndex(sheet)
	f.SetActiveSheet(index)

	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4F46E5"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	f.SetCellStyle(sheet, "A1", "F1", headerStyle)

	dataStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical: "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	headers := []string{
		"NO",
		"Username",
		"Name",
		"Email",
		"Section",
		"Role",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	users, _ := repositories.ExportUser()
	for i, u := range users {
		row := i + 2
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), i+1)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), u.Username)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), u.Name)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), u.Email)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), u.Section)
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), u.Role)
	}

	lastRow := len(users) + 1
	if lastRow >= 2 {
		f.SetCellStyle(sheet, "A2", "F"+strconv.Itoa(lastRow), dataStyle)
	}

	f.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	return f, nil
}

func ExportCSVUser() ([]models.User, error) {
	return repositories.ExportUser()
}

func ExportPDFUser() ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "User Report")
	pdf.Ln(12)

	headers := []string{"No", "Username", "Name", "Email", "Section", "Role"}
	widths := []float64{10, 35, 40, 50, 60, 30}

	pdf.SetFont("Arial", "B", 10)
	for i, h := range headers {
		pdf.CellFormat(widths[i], 7, h, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	users, _ := repositories.ExportUser()

	for i, u := range users {
		pdf.CellFormat(10, 7, strconv.Itoa(i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 7, u.Username, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 7, u.Name, "1", 0, "", false, 0, "")
		pdf.CellFormat(50, 7, u.Email, "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 7, u.Section, "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 7, u.Role, "1", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)

	return buf.Bytes(), err
}
