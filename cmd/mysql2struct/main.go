package main

import (
	"github.com/oddminng/mysql2gostruct/internal/mysql2struct"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var mysqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql 转换和处理",
	Long:  "sql 转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mysql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 转换",
	Long:  "sql 转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &mysql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := mysql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err:%v", err)
		}

		template := mysql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	mysqlCmd.AddCommand(mysql2structCmd)
	mysql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	mysql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	mysql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的地址HOST")
	mysql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的字符编码")
	mysql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	mysql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	mysql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名")
	rootCmd.AddCommand(mysqlCmd)
}

func main() {
	err := Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
