/*
分析脚本主键
2025.4.18 by dralee
*/
package main

import (
	"flag"
	"fmt"
	"keyanalysis/utils"
)

func test() {
	//fmt.Println("args:", os.Args)
	// args 1

	name := flag.String("name", "", "name") // go run key-analysis.go --name=lee --age=22
	age := flag.Int("age", 0, "age")
	source := flag.String("source", "", "source")
	fmt.Println("===============")
	flag.Parse()
	fmt.Println("name:", *name, "age:", *age)

	FileUtils := utils.FileInfo{}
	r := utils.FileExists(*name)
	files := utils.ListFiles(*source)
	fmt.Println(FileUtils, r)
	fmt.Println(files)

	content := utils.ReadString("/home/dralee/f/gitlab/z9yun-group/z9yun-database-script/Base/message/FAT/v0.1/v0.1.0/FAT_03_message_v0_1_0_table_data_init.sql")
	//keys := utils.MatchAll(`\((\d+),`, content)
	keys := utils.FindAllGroup(`\((\d+),`, content, 1)
	fmt.Println("keys:", keys)

}

func main() {
	source := flag.String("source", "", "source")
	flag.Parse()

	fmt.Println("source:", *source)
	if !utils.FileExists(*source) {
		fmt.Printf("source %s not exists\n", *source)
		return
	}

	ka := NewKeyAnalysis(*source)
	ka.Scan()
	ka.Save("result.log")
	fmt.Println("keys:", ka.keys)
	fmt.Println("existsKeys:", ka.existsKeys)

}
