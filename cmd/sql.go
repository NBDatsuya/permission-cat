package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"permission-cat/config"
	"permission-cat/internal/sql2struct"
)

var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql transforming",
	Long:  `sql transforming`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2StructCmd = &cobra.Command{
	Use:   "struct",
	Short: "generate go struct from database table",
	Long:  `generate go struct from database table`,
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   config.Conf.Sql2Struct.Type,
			Host:     config.Conf.Sql2Struct.Host,
			Username: config.Conf.Sql2Struct.Username,
			Password: config.Conf.Sql2Struct.Password,
			Charset:  config.Conf.Sql2Struct.Charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
			return
		}

		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
			return
		}

		templateStruct := sql2struct.NewStructTemplate()
		templateColumns := templateStruct.AssemblyColumns(columns)

		err = templateStruct.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("templateStruct.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2StructCmd)
	sql2StructCmd.Flags().StringVarP(&dbName, "db", "b", "", "db name")
	sql2StructCmd.Flags().StringVarP(&tableName, "table", "t", "", "table name")
}
