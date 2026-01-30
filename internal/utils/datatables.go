package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Row map[string]any

type DataTable struct {
	db             *gorm.DB
	baseQuery      string
	columns        []string
	addColumns     map[string]func(Row) any
	editColumns    map[string]func(any, Row) any
	additionalData map[string]any
}

func New(db *gorm.DB, query string, columns []string) *DataTable {
	return &DataTable{
		db:             db,
		baseQuery:      query,
		columns:        columns,
		addColumns:     map[string]func(Row) any{},
		editColumns:    map[string]func(any, Row) any{},
		additionalData: map[string]any{},
	}
}

func (dt *DataTable) With(data map[string]any) *DataTable {
	for k, v := range data {
		dt.additionalData[k] = v
	}
	return dt
}

func (dt *DataTable) AddColumn(name string, fn func(Row) any) *DataTable {
	dt.addColumns[name] = fn
	return dt
}

func (dt *DataTable) EditColumn(name string, fn func(any, Row) any) *DataTable {
	dt.editColumns[name] = fn
	return dt
}

func (dt *DataTable) Make(c *gin.Context) {
	draw, _ := strconv.Atoi(c.DefaultQuery("draw", "1"))
	start, _ := strconv.Atoi(c.DefaultQuery("start", "0"))
	length, _ := strconv.Atoi(c.DefaultQuery("length", "10"))
	search := c.Query("search[value]")

	orderColIdx, _ := strconv.Atoi(c.DefaultQuery("order[0][column]", "0"))
	orderDir := strings.ToUpper(c.DefaultQuery("order[0][dir]", "ASC"))
	orderColumn := dt.columns[orderColIdx]

	var total int64
	countSQL := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS t", dt.baseQuery)
	dt.db.Raw(countSQL).Scan(&total)

	query := dt.baseQuery
	args := []any{}

	if search != "" {
		parts := []string{}
		for _, col := range dt.columns {
			parts = append(parts, fmt.Sprintf("%s LIKE ?", col))
			args = append(args, "%"+search+"%")
		}
		query += " WHERE " + strings.Join(parts, " OR ")
	}

	var filtered int64
	filteredSQL := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS f", query)
	dt.db.Raw(filteredSQL, args...).Scan(&filtered)

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT ? OFFSET ?", orderColumn, orderDir)
	args = append(args, length, start)

	rows := []Row{}
	dt.db.Raw(query, args...).Scan(&rows)

	for _, row := range rows {
		for col, fn := range dt.editColumns {
			if val, ok := row[col]; ok {
				row[col] = fn(val, row)
			}
		}
		for col, fn := range dt.addColumns {
			row[col] = fn(row)
		}
	}

	resp := gin.H{
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": filtered,
		"data":            rows,
	}

	for k, v := range dt.additionalData {
		resp[k] = v
	}

	c.JSON(http.StatusOK, resp)
}

// Cara Menggunakan DataTable di Controller Gin-Gonic
/*
func MaterialTable(c *gin.Context) {
	query := `
		SELECT id, mold_number, model_name, lamp_name, type
		FROM materials
	`

	datatables.
		New(database.DB, query, []string{
			"mold_number",
			"model_name",
			"lamp_name",
			"type",
		}).
		AddColumn("action", func(row datatables.Row) any {
			return fmt.Sprintf(`
				<button class="btn btn-sm btn-danger" onclick="delete(%v)">Delete</button>
			`, row["id"])
		}).
		EditColumn("type", func(val any, row datatables.Row) any {
			return strings.ToUpper(fmt.Sprint(val))
		}).
		Make(c)
}
*/
