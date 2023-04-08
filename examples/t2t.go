package main

import (
	"fmt"
	"github.com/tiantianlikeu/converter"
)

func main() {
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly:       false,
		JsonTagToHump:     true, // 驼峰处理
		JsonTagFirstLower: true, // json  tag 驼峰首字母小写
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		//SeperatFile: false,
	})
	err := t2t.
		SavePath("./model.go").
		Dsn("root:root@tcp(localhost:3306)/database?charset=utf8mb4").
		TagKey("gorm").
		EnableJsonTag(true).
		EnableFormTag(true).
		EnableXmlTag(true).
		Table("tablename").
		Run()
	fmt.Println(err)
}
