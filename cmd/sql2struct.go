package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/taoruicheng/tour/internal/sql2struct"
)

var mysql2stuct = &cobra.Command{
	Use:   "mysql2stuct",
	Short: "读取mysql的表结构，转化为go的struct",
	Long:  "该子命令支持将mysql的表结构转化为go的struct，使用如下：./tour mysql2stuct --ip 10.151.3.169 --port 3308 -u root -p root --dbName nb_smart_building --tableName menu",
	Run: func(cmd *cobra.Command, args []string) {
		mysqlDbInfo := &sql2struct.DBInfo{
			DBType:   "mysql",
			Host:     mysqlIp + ":" + strconv.Itoa(mysqlPort),
			UserName: mysqlUserName,
			Password: mysqlPassword,
			Charset:  mysqlCharset,
		}

		dbModel := sql2struct.NewDBmodel(mysqlDbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("数据库连接异常:%v", err)
			return
		}

		menuColums, err := dbModel.GetColumns(dbName, tableName)

		if err != nil {
			log.Fatalf("获取列表异常:%v", err)
			return
		}

		structTemplate := sql2struct.NewStructTemplate()
		structColumn := structTemplate.AssemblyColumns(menuColums)
		s, err2 := structTemplate.Generate(tableName, structColumn)
		if err2 != nil {
			log.Fatalf("创建模板异常:%v", err)
			return
		}
		log.Println(s)

	},
}
var mysqlIp string
var mysqlPort int
var mysqlUserName string
var mysqlPassword string
var mysqlCharset string
var dbName string
var tableName string

func init() {
	mysql2stuct.Flags().StringVar(&mysqlIp, "ip", "127.0.0.1", "IP地址:10.151.3.169,默认127.0.0.1")
	mysql2stuct.Flags().IntVar(&mysqlPort, "port", 3306, "端口")
	mysql2stuct.Flags().StringVarP(&mysqlUserName, "username", "u", "root", "mysql登陆的用户，默认root")
	mysql2stuct.Flags().StringVarP(&mysqlPassword, "password", "p", "", "mysql登陆的密码")
	mysql2stuct.Flags().StringVarP(&mysqlCharset, "Charset", "c", "utf8mb3", "Charset，默认utf8mb3")

	mysql2stuct.Flags().StringVar(&dbName, "dbName", "", "数据库名字")
	mysql2stuct.MarkFlagRequired("dbName")
	mysql2stuct.Flags().StringVar(&tableName, "tableName", "", "表名")
	mysql2stuct.MarkFlagRequired("tableName")
}
