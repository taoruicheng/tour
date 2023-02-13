package sql2struct

import (
	"encoding/json"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	mysqlDbInfo := &DBInfo{
		DBType:   "mysql",
		Host:     "10.151.3.169:3308",
		UserName: "root",
		Password: "root",
		Charset:  "utf8mb3",
	}
	dbModel := NewDBmodel(mysqlDbInfo)
	err := dbModel.Connect()
	if err != nil {
		t.Fatalf("数据库连接异常:%v", err)
	}

	t.Log("test")
	menuColums, err := dbModel.GetColumns("nb_smart_building", "menu")

	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(menuColums)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(b))
}
