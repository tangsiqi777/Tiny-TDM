package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"log"
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

func (a *App) pageData(databaseName string, superTable string) []string {
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

	cols, err := rows.Columns() // Remember to check err afterwards
	vals := make([]interface{}, len(cols))
	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}
	for rows.Next() {
		err = rows.Scan(vals...)
	}
	return nil
}
