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
	Total      int64
}

type Query struct {
	//0 降序 1升序
	TimeOrder int    `json:"timeOrder" yaml:"timeOrder"`
	TimeStart string `json:"timeStart" yaml:"timeStart"`
	TimeEnd   string `json:"timeEnd" yaml:"timeEnd"`
	PrimaryId string `json:"primaryId" yaml:"primaryId"`
	Size      int    `json:"size" yaml:"size"`
	Current   int    `json:"current" yaml:"current"`
}

func (a *App) PageData1(config service.ConnectionConfig, databaseName string, table string, query Query) PageData {
	queryStr, _ := json.Marshal(query)
	fmt.Println("go query：%s", string(queryStr))
	fmt.Println("go query：%s", a.getPageSql(query, databaseName, table, 0))
	fmt.Println("go query：%s", a.getPageSql(query, databaseName, table, 1))
	var p PageData
	dbConn, err := a.getConn(config)
	if err != nil {
		return p
	}
	rows, err1 := dbConn.Query(a.getPageSql(query, databaseName, table, 0))
	total, err2 := dbConn.Query(a.getPageSql(query, databaseName, table, 1))
	if err1 != nil || err2 != nil {
		log.Println("failed to select from data1, err:", err)
		return p
	}
	defer rows.Close()
	defer total.Close()
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
		rowNum++
	}
	rowList1 := make([]map[string]interface{}, rowNum)
	for i := 0; i < rowNum; i++ {
		rowList1[i] = rowList[i]
	}
	p.Data = rowList1
	fmt.Println(rowNum)

	var totalNum int64 = 0
	total.Next()
	err3 := total.Scan(&totalNum)
	if err3 != nil {
		return p
	}

	p.Total = totalNum
	fmt.Println(totalNum)

	return p
}

func (a *App) getPageSql(query Query, databaseName string, table string, queryType int) string {
	sqlWhere := ""
	sqlOrder := ""
	sqlLimit := ""
	sqlType := " * "

	if query.TimeStart != "" {
		sqlWhere = "WHERE " + query.PrimaryId + " >=  " + "\"" + query.TimeStart + "\""
		if query.TimeEnd != "" {
			sqlWhere = sqlWhere + " AND " + query.PrimaryId + " <= " + "\"" + query.TimeEnd + "\""
		}
	} else {
		if query.TimeEnd != "" {
			sqlWhere = " AND " + query.PrimaryId + " <= " + "\"" + query.TimeEnd + "\""
		}
	}

	if queryType == 1 {
		sqlType = " COUNT(*) "
	} else {
		if query.TimeOrder == 0 {
			sqlOrder = " ORDER BY " + query.PrimaryId + " DESC "
		} else {
			sqlOrder = " ORDER BY " + query.PrimaryId + " ASC "
		}

		if query.Size != 0 {
			sqlLimit = "LIMIT " + fmt.Sprintf("%d", query.Size) + "OFFSET " + fmt.Sprintf("%d", (query.Current-1)*query.Size)
		} else {
			sqlLimit = " LIMIT 50 OFFSET " + fmt.Sprintf("%d", (query.Current-1)*query.Size)
		}
	}

	sql1 := "SELECT " + sqlType + " FROM `" + databaseName + "`.`" + table + "` " + sqlWhere + sqlOrder + sqlLimit
	return sql1
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
