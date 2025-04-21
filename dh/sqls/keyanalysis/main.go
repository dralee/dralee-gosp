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

func testTrie() {
	trie := utils.NewTrie(10)

	// 要替换的关键词表
	replaceMap := map[string]string{
		"Go":      "Golang",
		"OpenAI":  "Open AI",
		"ChatGPT": "Chat GPT",
		"Open":    "Close",
	}

	// 把关键词塞进Trie
	for k := range replaceMap {
		trie.Insert(k)
	}

	text := "I love Go and ChatGPT developed by OpenAI!I love Go Open and ChatGPT developed by OpenAI!I love Go and ChatGPT developed by OpenAI!"
	newText := trie.Replace(text, replaceMap)

	fmt.Println(newText)
}

func scan(srcSqlPath string) {
	ka := NewKeyAnalysis(srcSqlPath)
	ka.Scan()
	ka.Save(fmt.Sprintf("%s/%s", OutputPath, "result.log"))
	Info("keys: %v", ka.keys)
	Info("existsKeys: %v", ka.existsKeys)
}

func replace(srcSqlPath, srcKeyFile, newKeyFile string) {
	kp := NewKeyReplace(srcKeyFile, newKeyFile, srcSqlPath)
	err := kp.Run()
	if err != nil {
		Error(err)
	} else {
		Info("replace success")
	}
}

func main() {
	//testTrie()
	//return
	source := flag.String("source", "", "the sqls path for analysis or replace.")
	optType := flag.String("optType", "", "operate type to analysis: Scan or Replace")
	srcKeyFile := flag.String("srcKeyFile", "", "src key file for replace")
	newKeyFile := flag.String("newKeyFile", "", "new key file for replace")
	flag.Parse()

	initLogger()
	keyOperateType, err := ParseKeyOperateType(*optType)
	if err != nil {
		Error(err)
		return
	}

	Info("source: %s", *source)
	if !utils.FileExists(*source) {
		Errorf("source %s not exists\n", *source)
		return
	}

	if keyOperateType == Scan {
		scan(*source)
	} else {
		replace(*source, *srcKeyFile, *newKeyFile)
	}
}
