package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType string

	Host     string
	Username string
	Password string

	Charset string
}

type TableColumn struct {
	ColumnName string
	DataType   string
	IsNullable string

	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

const InformationSchema = "information_schema"

var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error

	dbConnStr := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.Username,
		m.DBInfo.Password,
		m.DBInfo.Host,
		InformationSchema,
		m.DBInfo.Charset,
	)

	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dbConnStr)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := strings.Join([]string{
		"SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, COLUMN_TYPE, IS_NULLABLE, COLUMN_COMMENT ",
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? ",
	}, "")

	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("no data in table")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn

		err = rows.Scan(
			&column.ColumnName,
			&column.DataType,
			&column.ColumnKey,
			&column.ColumnType,
			&column.IsNullable,
			&column.ColumnComment,
		)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
