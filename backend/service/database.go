package service

import (
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"log"
)

func ListDatabases() []string {
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

func ListSuperTable() []string {
	var taosUri = "root:taosdata@http(192.168.56.19:6041)/"
	taos, err := sql.Open("taosRestful", taosUri)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return nil
	} else {
		fmt.Println("success")
	}

	rows, err := taos.Query("show stables;")
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
