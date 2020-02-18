//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/16 14:28 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package main

import (
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strings"
)

type ExcelData struct {
	FieldName     string //第一列 字段名称
	FieldComment  string // 第二列 字段描述
	FieldDataType string //第三列 字段数据类型
	IsPrimary     bool   //第四列 是否为主键
	Describe      string //第五列 字段描述附加
	IsNull        string //第六列 是否为空
	DefaultValue  string //第七列 默认值
}

func main() {
	var xlsFile,comment string
	flag.StringVar(&xlsFile, "f", "", "excel的名称")
	flag.StringVar(&comment, "c", "", "数据表名")
	flag.Parse()

	//sql语句开头第一句
	var sqlStr string
	sqlStr += fmt.Sprintf("CREATE TABLE `%s` (\n", xlsFile)
	//excel的路径
	xlsPath := fmt.Sprintf("./xls/%s.xlsx", xlsFile)
	//要生成的sql文件
	fileName := fmt.Sprintf("./sql/%s.sql",xlsFile)
	dstFile,err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()

	//打开并读取excel文件
	f, err := excelize.OpenFile(xlsPath)
	if err != nil {
		panic("excle文件读取失败")
		return
	}

	rows := f.GetRows("Sheet1")
	var datas []ExcelData
	var primaryKey string

	for i, row := range rows {
		if i < 1 {
			continue
		}
		var data ExcelData
		for j, colCell := range row {
			if i==1 && j==0{
				primaryKey= colCell  //主键
			}
			//第一列
			if j == 0 {
				data.FieldName = colCell
			}
			//第二列
			if j == 1 {
				data.FieldComment = colCell
			}

			//第三列
			if j == 2 {
				data.FieldDataType = colCell
			}

			//第四列
			if j == 3 {
				if colCell == "Y" {
					data.IsPrimary = true
				} else {
					data.IsPrimary = false
				}

			}

			//第五列
			if j == 4 {
				data.Describe = colCell
			}

			//第六列
			if j == 5 {
				if colCell == "N" {
					data.IsNull = "NOT NULL"
				} else {
					data.IsNull = ""
				}
			}

			//第七列
			if j == 6 {
				if colCell != "" {
					data.DefaultValue = strings.Join([]string{"DEFAULT '",colCell,"'"},"")
				} else {
					data.DefaultValue = ""
				}
			}

		}
		datas = append(datas, data)
	}

	for _, v := range datas {
		if v.IsPrimary {
			sqlStr += strings.Join([]string{"`", v.FieldName, "`", " ", v.FieldDataType, " unsigned ", v.IsNull, " AUTO_INCREMENT COMMENT '", v.FieldComment, "',\n"}, "")
		} else {
			sqlStr += strings.Join([]string{"`", v.FieldName, "`", " ", v.FieldDataType, " ", v.IsNull," ",v.DefaultValue, " COMMENT '", v.FieldComment, "',\n"}, "")
		}
	}

	sqlStr+=strings.Join([]string{"PRIMARY KEY (`",primaryKey,"`)\n"},"")
	sqlStr+=") ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='"+comment+"';"
	fmt.Println(sqlStr)
	dstFile.WriteString(sqlStr+"\n")


}
