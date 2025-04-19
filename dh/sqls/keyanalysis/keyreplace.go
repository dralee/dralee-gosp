/*
主键替换
2025.4.19 by dralee
*/
package main

import (
	"keyanalysis/utils"
	"strings"
)

type KeyReplace struct {
	srcSqlPath string
	srcKeys    []string
	newKeys    []string
	msg        map[string]string
}

/*
实例化
srcKeyFileName 源主键文件,
newKeyFileName 新主键文件,
srcSqlPath 源sql文件
*/
func NewKeyReplace(srcKeyFileName, newKeyFileName string, srcSqlPath string) *KeyReplace {
	kp := KeyReplace{
		srcSqlPath: srcSqlPath,
		srcKeys:    strings.Split(utils.ReadString(srcKeyFileName), ","),
		newKeys:    strings.Split(utils.ReadString(newKeyFileName), ","),
		msg:        make(map[string]string),
	}
	if !kp.checkKeys() {
		panic("src key and new key num not equal")
	}
	return &kp
}

// 执行
func (kp *KeyReplace) Run() {
	files := utils.ListFiles(kp.srcSqlPath)
	for _, file := range files {
		if IsIgnoreFile(file.FilePath) {
			continue
		}
		if !kp.replace(file.FilePath) {
			kp.msg[file.FilePath] = "replace failed"
		}
	}
}

func (kp *KeyReplace) isIgnoreKey(key string) bool {
	nKey := utils.ToUInt64(key)
	// 小于13位的数值直接忽略
	return len(key) > 13 && nKey > 0
}

func (kp *KeyReplace) formatKey(content string) string {
	for i := 1; i < len(kp.srcKeys); i++ {
		if kp.isIgnoreKey(kp.srcKeys[i]) {
			continue
		}
		content = strings.ReplaceAll(content, kp.srcKeys[i], kp.newKeys[i])
	}
	return content
}

func (kp *KeyReplace) replace(fileName string) bool {
	content := utils.ReadString(fileName)
	content = kp.formatKey(content)
	err := utils.WriteString(fileName, content)
	return err == nil
}

func (kp *KeyReplace) checkKeys() bool {
	return len(kp.srcKeys) != len(kp.newKeys)
}
