package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

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

// "root:123456@tcp(127.0.0.1:3306)/test"
const c_information_schema_connect_sql string = "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=true&loc=Local&autocommit=true&timeout=5s&readTimeout=100s&writeTimeout=100s"

func NewDBmodel(dbInfo *DBInfo) *DBModel {
	return &DBModel{DBInfo: dbInfo}
}
func (dbModel *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf(c_information_schema_connect_sql, dbModel.DBInfo.UserName, dbModel.DBInfo.Password, dbModel.DBInfo.Host, dbModel.DBInfo.Charset)
	log.Printf("数据库连接语句:%s \n", dsn)
	dbModel.DBEngine, err = sql.Open(dbModel.DBInfo.DBType, dsn)
	if err != nil {
		log.Fatalln("链接数据库报错 ", err)
		return err
	}
	err = dbModel.DBEngine.Ping()
	if err != nil {
		log.Fatalln("ping数据库报错", err)
		return err
	}
	log.Println("数据库连接成功")
	dbModel.DBEngine.SetConnMaxLifetime(100 * time.Second)
	dbModel.DBEngine.SetMaxIdleConns(1)
	dbModel.DBEngine.SetMaxOpenConns(1)
	return nil
}
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
