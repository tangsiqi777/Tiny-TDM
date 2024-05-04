package service

import (
	"context"
	"encoding/base64"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"io"
	"net/http"
	"strings"
	"time"
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

type RestSqlResult struct {
	//返回数据
	Result string
	//错误信息
	Msg string
}

func (a *RestSqlService) DoPost(config ConnectionConfig, sql string) RestSqlResult {
	restSqlResult := RestSqlResult{}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", "http://"+config.Addr+":"+fmt.Sprintf("%d", config.Port)+"/rest/sql", strings.NewReader(sql))
	if err != nil {
		fmt.Println("连接参数错误")
		restSqlResult.Result = ""
		restSqlResult.Msg = "连接参数错误:" + err.Error()
		return restSqlResult
	}
	userAndPwd := "Basic " + base64.StdEncoding.EncodeToString([]byte((config.Username + ":" + config.Password)))
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Authorization", userAndPwd)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("连接错误")
		restSqlResult.Result = ""
		restSqlResult.Msg = "连接错误:" + err.Error()
		return restSqlResult
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取数据错误")
		restSqlResult.Result = ""
		restSqlResult.Msg = "读取数据错误:" + err.Error()
		return restSqlResult
	}
	restSqlResult.Result = string(body)
	restSqlResult.Msg = ""
	return restSqlResult
}

func (a *RestSqlService) ListDatabases(config ConnectionConfig) RestSqlResult {
	sql := "show databases;"
	return a.DoPost(config, sql)
}

func (a *RestSqlService) ListSuperTable(config ConnectionConfig, databaseName string, superTableSearch string) RestSqlResult {
	querySql := "show `" + databaseName + "`.stables;"
	if superTableSearch != "" {
		querySql = "show `" + databaseName + "`.stables LIKE \"%" + superTableSearch + "%\";"
	}
	fmt.Print(querySql)
	return a.DoPost(config, querySql)
}

func (a *RestSqlService) ListChildTable(config ConnectionConfig, databaseName string, superTable string, childTableSearch string) RestSqlResult {
	querySql := "SELECT DISTINCT TBNAME FROM `" + databaseName + "`.`" + superTable + "` ORDER BY TBNAME ASC;"
	if childTableSearch != "" {
		querySql = "SELECT DISTINCT TBNAME FROM `" + databaseName + "`.`" + superTable + "`" + "where TBNAME LIKE \"%" + childTableSearch + "%\" ORDER BY TBNAME ASC;"
	}
	return a.DoPost(config, querySql)
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

func (a *RestSqlService) PageData(config ConnectionConfig, databaseName string, table string, query Query) RestSqlResult {
	dataSql := a.getPageSql(query, databaseName, table, 0)
	return a.DoPost(config, dataSql)

}

func (a *RestSqlService) CountData(config ConnectionConfig, databaseName string, table string, query Query) RestSqlResult {
	countSql := a.getPageSql(query, databaseName, table, 1)
	return a.DoPost(config, countSql)
}

func (a *RestSqlService) DescTable(config ConnectionConfig, databaseName string, superTable string) RestSqlResult {
	querySql := "DESC `" + databaseName + "`.`" + superTable + "`;"
	return a.DoPost(config, querySql)
}

func (a *RestSqlService) SqlQuery(config ConnectionConfig, sql string) RestSqlResult {
	return a.DoPost(config, sql)
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
		sqlType = " COUNT(*) AS Total"
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
