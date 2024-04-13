package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"log"
	"time"
	"tinytdm/backend/service"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) getConn(config service.ConnectionConfig) (*sql.DB, error) {
	configStr, _ := json.Marshal(config)
	fmt.Println("go conn：%s", string(configStr))
	var url = config.Username + ":" + config.Password + "@http(" + config.Addr + ":" + fmt.Sprintf("%d", config.Port) + ")/"
	dbConn, err := sql.Open("taosRestful", url)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return nil, err
	} else {
		fmt.Println("success conn")
		return dbConn, nil
	}
}

func (a *App) getConn1() (*sql.DB, error) {
	var url = "root:taosdata@http(192.168.56.19:6041)/"
	dbConn, err := sql.Open("taosRestful", url)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return nil, err
	} else {
		fmt.Println("success conn")
		return dbConn, nil
	}
}

func (a *App) ListDatabases(config service.ConnectionConfig) []string {
	dbConn, err := a.getConn(config)
	if err != nil {
		fmt.Println("empty db")
		return []string{}
	}
	rows, err := dbConn.Query("show databases;")
	if err != nil {
		log.Println("failed to select database from table, err:", err)
		return []string{}
	}
	defer rows.Close()

	var slice []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Println("scan error:\n", err)
			return []string{}
		}
		slice = append(slice, name)
		log.Println(name)
	}
	return slice
}

func (a *App) ListSuperTable(config service.ConnectionConfig, databaseName string, superTableSearch string) []string {
	querySql := "show `" + databaseName + "`.stables;"
	if superTableSearch != "" {
		querySql = "show `" + databaseName + "`.stables LIKE \"%" + superTableSearch + "%\";"
	}
	fmt.Print(querySql)
	dbConn, err := a.getConn(config)
	if err != nil {
		return []string{}
	}
	rows, err := dbConn.Query(querySql)
	if err != nil {
		log.Println("failed to select from super table, err:", err)
		return []string{}
	}
	defer rows.Close()
	var slice []string
	for rows.Next() {
		var stableName string
		err := rows.Scan(&stableName)
		if err != nil {
			log.Println("scan error:\n", err)
			return []string{}
		}
		slice = append(slice, stableName)
		log.Println(stableName)
	}
	return slice
}

func (a *App) ListChildTable(config service.ConnectionConfig, databaseName string, superTable string, childTableSearch string) []string {
	querySql := "SELECT DISTINCT TBNAME FROM `" + databaseName + "`.`" + superTable + "`;"
	if childTableSearch != "" {
		querySql = "SELECT DISTINCT TBNAME FROM `" + databaseName + "`.`" + superTable + "`" + "where TBNAME LIKE \"%" + childTableSearch + "%\";"
	}
	dbConn, err := a.getConn(config)
	if err != nil {
		return []string{}
	}

	rows, err := dbConn.Query(querySql)
	if err != nil {
		log.Println("failed to select from table, err:", err)
		return []string{}
	}
	defer rows.Close()
	var slice []string
	for rows.Next() {
		var childName string

		err := rows.Scan(&childName)
		if err != nil {
			log.Println("scan error:\n", err)
			return []string{}
		}
		slice = append(slice, childName)
		log.Println(childName)
	}
	return slice
}

type PageData struct {
	Data       []map[string]interface{}
	HeaderList []string
}

func (a *App) PageData1(config service.ConnectionConfig, databaseName string, table string) PageData {
	var p PageData
	dbConn, err := a.getConn(config)
	if err != nil {
		return p
	}
	rows, err := dbConn.Query("SELECT * FROM `" + databaseName + "`.`" + table + "` LIMIT 100;")
	if err != nil {
		log.Println("failed to select from data1, err:", err)
		return p
	}
	defer rows.Close()
	rowList := make([]map[string]interface{}, 100)
	var rowNum = 0
	for rows.Next() {
		cols, _ := rows.Columns()
		row := make([]interface{}, len(cols))
		for i, _ := range cols {
			t, _ := rows.ColumnTypes()
			row[i] = a.Check(t[i].ScanType().Name())

		}
		err := rows.Scan(row...)

		rowMap := make(map[string]interface{})
		headerList := make([]string, len(cols))
		for i := 0; i < len(cols); i++ {
			t, _ := rows.ColumnTypes()
			columnName := t[i].Name()
			headerList[i] = columnName
			rowMap[columnName] = row[i]
		}
		if rowNum == 0 {
			p.HeaderList = headerList
		}
		rowList[rowNum] = rowMap
		if err != nil {
			log.Println("\nscan error:", err)
		}
		fmt.Print(row[1])
		rowNum++
	}
	p.Data = rowList
	return p
}

func (a *App) Check(strType string) interface{} {
	switch strType {
	case "NullString":
		return new(string)
	case "NullInt64":
		return new(int64)
	case "NullInt32":
		return new(int32)
	case "NullByte":
		return new(byte)
	case "NullFloat32":
		return new(float32)
	case "NullFloat64":
		return new(float64)
	case "NullBool":
		return new(bool)
	case "NullTime":
		return new(time.Time)
	}
	return nil
}
