package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	wbFile := "../../../CMDB/CMDB信息总览.xlsx"
	wbData := readExcel(wbFile, "cmdb")
	generateWBSQL(wbData)

	vbFile := "../../../CMDB/VBCMDB.xlsx"
	vbData := readExcel(vbFile, "Sheet1")
	generateVBSQL(vbData)
}

func readExcel(filename, sheet string) [][]string {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func escape(s string) string {
	s = strings.ReplaceAll(s, "'", "''")
	s = strings.ReplaceAll(s, "\\", "\\\\")
	return s
}

func generateWBSQL(rows [][]string) {
	f, _ := os.Create("../../sql/insert_wb_cmdb.sql")
	defer f.Close()

	f.WriteString("-- WB CMDB数据导入\n")
	f.WriteString("INSERT INTO `wb_cmdb_info` (`system_name`, `environment`, `ip_address`, `port`, `status`, `owner`, `remark`) VALUES\n")

	var values []string
	for i, row := range rows {
		if i == 0 || len(row) < 2 {
			continue
		}

		systemName := ""
		environment := ""
		ipAddress := ""
		port := ""
		status := ""
		owner := ""
		remark := ""

		if len(row) > 0 {
			systemName = escape(strings.TrimSpace(row[0]))
		}
		if len(row) > 1 {
			environment = escape(strings.TrimSpace(row[1]))
		}
		if len(row) > 2 {
			ipAddress = escape(strings.TrimSpace(row[2]))
		}
		if len(row) > 3 {
			port = escape(strings.TrimSpace(row[3]))
		}
		if len(row) > 4 {
			status = escape(strings.TrimSpace(row[4]))
		}
		if len(row) > 5 {
			owner = escape(strings.TrimSpace(row[5]))
		}
		if len(row) > 6 {
			remark = escape(strings.TrimSpace(row[6]))
		}

		if systemName == "" {
			continue
		}

		values = append(values, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s')",
			systemName, environment, ipAddress, port, status, owner, remark))
	}

	f.WriteString(strings.Join(values, ",\n"))
	f.WriteString(";\n")
	fmt.Println("WB CMDB SQL生成完成: sql/insert_wb_cmdb.sql")
}

func generateVBSQL(rows [][]string) {
	f, _ := os.Create("../../sql/insert_vb_cmdb.sql")
	defer f.Close()

	f.WriteString("-- VB CMDB数据导入\n")
	f.WriteString("INSERT INTO `vb_cmdb_info` (`system_name`, `environment`, `ip_address`, `port`, `status`, `owner`, `remark`) VALUES\n")

	var values []string
	for i, row := range rows {
		if i == 0 || len(row) < 2 {
			continue
		}

		systemName := ""
		environment := ""
		ipAddress := ""
		port := ""
		status := ""
		owner := ""
		remark := ""

		if len(row) > 0 {
			systemName = escape(strings.TrimSpace(row[0]))
		}
		if len(row) > 1 {
			environment = escape(strings.TrimSpace(row[1]))
		}
		if len(row) > 2 {
			ipAddress = escape(strings.TrimSpace(row[2]))
		}
		if len(row) > 3 {
			port = escape(strings.TrimSpace(row[3]))
		}
		if len(row) > 4 {
			status = escape(strings.TrimSpace(row[4]))
		}
		if len(row) > 5 {
			owner = escape(strings.TrimSpace(row[5]))
		}
		if len(row) > 6 {
			remark = escape(strings.TrimSpace(row[6]))
		}

		if systemName == "" {
			continue
		}

		values = append(values, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s')",
			systemName, environment, ipAddress, port, status, owner, remark))
	}

	f.WriteString(strings.Join(values, ",\n"))
	f.WriteString(";\n")
	fmt.Println("VB CMDB SQL生成完成: sql/insert_vb_cmdb.sql")
}
