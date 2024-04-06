package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"log"
	"time"
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

func (a *App) ListDatabases() []string {
	var taosUri = "root:taosdata@http(192.168.56.19:6041)/"
	taos, err := sql.Open("taosRestful", taosUri)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return nil
	} else {
		fmt.Println("success")
	}

	rows, err := taos.Query("show databases;")
	if err != nil {
		log.Fatalln("failed to select from table, err:", err)
	}
	defer rows.Close()

	var slice []string

	for rows.Next() {
		var (
			name string
		)
		err := rows.Scan(&name)
		if err != nil {
			log.Fatalln("scan error:\n", err)
			return nil
		}
		slice = append(slice, name)
		log.Println(name)
	}
	return slice
}

func (a *App) ListSuperTable(databaseName string) []string {
	var taosUri = "root:taosdata@http(192.168.56.19:6041)/"
	taos, err := sql.Open("taosRestful", taosUri)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return nil
	} else {
		fmt.Println("success")
	}

	rows, err := taos.Query("show `" + databaseName + "`.stables;")
	if err != nil {
		log.Fatalln("failed to select from table, err:", err)
	}
	defer rows.Close()

	var slice []string

	for rows.Next() {
		var (
			stable_name string
		)
		err := rows.Scan(&stable_name)
		if err != nil {
			log.Fatalln("scan error:\n", err)
			return nil
		}
		slice = append(slice, stable_name)
		log.Println(stable_name)
	}
	return slice
}

func (a *App) ListChildTable(databaseName string, superTable string) []string {
	var taosUri = "root:taosdata@http(192.168.56.19:6041)/"
	taos, err := sql.Open("taosRestful", taosUri)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return nil
	} else {
		fmt.Println("success")
	}

	rows, err := taos.Query("SELECT DISTINCT TBNAME FROM `" + databaseName + "`.`" + superTable + "`;")
	if err != nil {
		log.Fatalln("failed to select from table, err:", err)
	}
	defer rows.Close()

	var slice []string

	for rows.Next() {
		var (
			tbname string
		)
		err := rows.Scan(&tbname)
		if err != nil {
			log.Fatalln("scan error:\n", err)
			return nil
		}
		slice = append(slice, tbname)
		log.Println(tbname)
	}
	return slice
}

func (a *App) PageData(databaseName string, table string) [][]interface{} {
	var taosUri = "root:taosdata@http(192.168.56.19:6041)/"
	taos, err := sql.Open("taosRestful", taosUri)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return nil
	} else {
		fmt.Println("success")
	}

	rows, err := taos.Query("SELECT * FROM `power`.`meters` LIMIT 100;")
	if err != nil {
		log.Fatalln("failed to select from table, err:", err)
	}
	defer rows.Close()

	vals := make([][]interface{}, 100)
	var rowNum = 0
	for rows.Next() {
		cols, _ := rows.Columns()
		row := make([]interface{}, len(cols))
		vals[rowNum] = row
		for i, _ := range cols {
			t, _ := rows.ColumnTypes()
			row[i] = a.Check(t[i].ScanType().Name())
		}
		err := rows.Scan(row...)
		if err != nil {
			log.Fatalln("\nscan error:", err)
		}
		fmt.Print(row[1])
		rowNum++
	}
	return vals
}

type PageData struct {
	Data       []map[string]interface{}
	HeaderList []string
}

func (a *App) PageData1(databaseName string, table string) PageData {
	var taosUri = "root:taosdata@http(192.168.56.19:6041)/"
	taos, err := sql.Open("taosRestful", taosUri)
	var p PageData
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return p
	} else {
		fmt.Println("success")
	}

	rows, err := taos.Query("SELECT * FROM `" + databaseName + "`.`" + table + "` LIMIT 100;")
	if err != nil {
		log.Fatalln("failed to select from table, err:", err)
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
			log.Fatalln("\nscan error:", err)
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
