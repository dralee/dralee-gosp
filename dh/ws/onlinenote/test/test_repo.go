/*
测试仓储
2025.4.23 by dralee
*/
package main

import (
	"draleeonlinenote/note"
	"fmt"
	"log"
)

func test() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/notedb?charset=utf8mb4&parseTime=True&loc=Local"
	repo, err := note.NewDefaultNoteRepository(dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = repo.Save(&note.Note{Name: "test", Content: "test", CreatorId: 1})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("success")
}

func main() {
	test()
}
