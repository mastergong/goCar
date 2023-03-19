package database

import (
	"carapi/models"
	"encoding/json"

	_ "github.com/denisenkom/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"fmt"
)

// Database variables

var dbName = ""
var userDb = ""
var pwdDb = ""
var serverDb = ""

var DB *gorm.DB

func Connect() {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;", serverDb, userDb, pwdDb)

	connection, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.PSNCar{})
	connection.AutoMigrate(&models.PSNHisUsedCar{})

	GetUsedId()

}

func Disconnect() {

	if DB != nil {
		defer DB.Distinct()
	}
}

func GetUsedId() int {

	_sql := ""

	getJSON(_sql)

	return 0
}

func getJSON(sqlString string) (string, error) {
	rows, err := DB.Raw(sqlString).Rows()
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
