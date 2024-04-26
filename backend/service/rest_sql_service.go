package service

import (
	"context"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"io"
	"net/http"
	"strings"
)

type RestSqlService struct {
	ctx context.Context
}

func NewRestSqlService() *RestSqlService {
	return &RestSqlService{}
}

func (a *RestSqlService) Startup(ctx context.Context) {
	a.ctx = ctx
}

type ResultSqlResult struct {
	//返回数据
	Result string
	//错误信息
	Msg string
}

func (a *RestSqlService) post(sql string) (string, string) {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body), ""
}

/*// Greet returns a greeting for the given name
func (a *RestSqlService) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *RestSqlService) getConn(config ConnectionConfig) (*sql.DB, error) {
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

func (a *RestSqlService) ListDatabases(config ConnectionConfig) []string {
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

func (a *RestSqlService) ListSuperTable(config ConnectionConfig, databaseName string, superTableSearch string) []string {
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

func (a *RestSqlService) ListChildTable(config ConnectionConfig, databaseName string, superTable string, childTableSearch string) []string {
	querySql := "SELECT DISTINCT TBNAME FROM `" + databaseName + "`.`" + superTable + "` ORDER BY TBNAME ASC;"
	if childTableSearch != "" {
		querySql = "SELECT DISTINCT TBNAME FROM `" + databaseName + "`.`" + superTable + "`" + "where TBNAME LIKE \"%" + childTableSearch + "%\" ORDER BY TBNAME ASC;"
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
		//log.Println(childName)
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

func (a *RestSqlService) PageData1(config ConnectionConfig, databaseName string, table string, query Query) PageData {
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

func (a *RestSqlService) SqlQuery(config ConnectionConfig, sql string) PageData {

	var p PageData
	dbConn, err := a.getConn(config)
	if err != nil {
		return p
	}
	rows, err1 := dbConn.Query(sql)
	if err1 != nil {
		log.Println("failed to select from data1, err:", err)
		return p
	}
	defer rows.Close()
	rowList := make([]map[string]interface{}, 2000)
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
		if rowNum > 2000 {
			break
		}
	}
	rowList1 := make([]map[string]interface{}, rowNum)
	for i := 0; i < rowNum; i++ {
		rowList1[i] = rowList[i]
	}
	p.Data = rowList1
	fmt.Println(rowNum)
	return p
}

func (a *RestSqlService) getPageSql(query Query, databaseName string, table string, queryType int) string {
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
			sqlLimit = "LIMIT " + fmt.Sprintf("%d", query.Size) + " OFFSET " + fmt.Sprintf("%d", (query.Current-1)*query.Size)
		} else {
			sqlLimit = " LIMIT 50 OFFSET " + fmt.Sprintf("%d", (query.Current-1)*query.Size)
		}
	}

	sql1 := "SELECT " + sqlType + " FROM `" + databaseName + "`.`" + table + "` " + sqlWhere + sqlOrder + sqlLimit
	fmt.Println(sql1)
	return sql1
}

type TableField struct {
	Field  string `json:"field" yaml:"field"`
	Type   string `json:"type" yaml:"type"`
	Length int    `json:"length" yaml:"length"`
	Note   string `json:"note" yaml:"note"`
}

func (a *RestSqlService) DescTable(config ConnectionConfig, databaseName string, superTable string) []TableField {
	querySql := "DESC `" + databaseName + "`.`" + superTable + "`;"
	dbConn, err := a.getConn(config)
	if err != nil {
		return []TableField{}
	}

	rows, err := dbConn.Query(querySql)
	if err != nil {
		log.Println("failed to desc table, err:", err)
		return []TableField{}
	}
	defer rows.Close()
	var slice []TableField
	for rows.Next() {
		var tableField TableField

		err := rows.Scan(&tableField.Field, &tableField.Type, &tableField.Length, &tableField.Note)
		if err != nil {
			log.Println("scan error:\n", err)
			return []TableField{}
		}
		slice = append(slice, tableField)
	}
	log.Println(slice)
	return slice
}

func (a *RestSqlService) Check(strType string) interface{} {
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
*/
