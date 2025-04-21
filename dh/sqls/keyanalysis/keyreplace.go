/*
主键替换
2025.4.19 by dralee
*/
package main

import (
	"fmt"
	"keyanalysis/utils"
	"strings"
	"sync"

	"github.com/cloudflare/ahocorasick"
)

type KeyReplace struct {
	srcSqlPath string
	srcKeys    []string
	newKeys    []string
	msg        map[string]string
	//repChan    chan string
	quit       chan bool
	trie       *utils.Trie
	replaceMap map[string]string
}

/*
实例化
srcKeyFileName 源主键文件,
newKeyFileName 新主键文件,
srcSqlPath 源sql文件
*/
func NewKeyReplace(srcKeyFileName, newKeyFileName string, srcSqlPath string) *KeyReplace {
	if srcKeyFileName == "" {
		panic("src key file name is empty")
	}
	if newKeyFileName == "" {
		panic("new key file name is empty")
	}
	if !utils.FileExists(srcKeyFileName) {
		panic(fmt.Sprintf("src key file \"%s\" not exists", srcKeyFileName))
	}
	if !utils.FileExists(newKeyFileName) {
		panic(fmt.Sprintf("new key file \"%s\" not exists", newKeyFileName))
	}
	if srcSqlPath == "" {
		panic("src sql path is empty")
	}

	kp := KeyReplace{
		srcSqlPath: srcSqlPath,
		srcKeys:    strings.Split(utils.ReadString(srcKeyFileName), ","),
		newKeys:    strings.Split(utils.ReadString(newKeyFileName), ","),
		msg:        make(map[string]string),
		//repChan:    make(chan string, 10),
		quit: make(chan bool),
		trie: utils.NewTrie(16),
	}
	if !kp.checkKeys() {
		panic("src key and new key num not equal")
	}

	kp.initTrie()

	return &kp
}

// 执行
func (kp *KeyReplace) Run() error {
	if !kp.checkKeys() {
		return fmt.Errorf("src key and new key num not equal")
	}

	if !utils.FileExists(kp.srcSqlPath) {
		return fmt.Errorf("src sql path \"%s\" not exists", kp.srcSqlPath)
	}

	//go kp.ListenKeys()

	wg := sync.WaitGroup{}
	files := utils.ListFiles(kp.srcSqlPath)
	var count int
	for _, file := range files {
		if IsIgnoreFile(file.FilePath) {
			continue
		}
		//go kp.replace(file.FilePath)
		/*if !kp.replace(file.FilePath) {
			kp.msg[file.FilePath] = "replace failed"
		}*/
		go kp.doReplace(file.FilePath, &wg)
		wg.Add(1)
		count++
		//break // 测试运行一次
	}

	go func() {
		wg.Wait()
		kp.quit <- true
	}()

	Info("waiting for replacing keys")
	<-kp.quit
	Info("finish.")

	if count == 0 {
		return fmt.Errorf("there not any file for replacing keys")
	}

	return nil
}

func (kp *KeyReplace) initTrie() {
	replaceMap := map[string]string{}
	for i, k := range kp.srcKeys {
		replaceMap[k] = kp.newKeys[i]
	}
	// 关键词入Trie
	for k := range replaceMap {
		kp.trie.Insert(k)
	}
	kp.replaceMap = replaceMap

	utils.WriteMapString("./data/temp.log", kp.replaceMap)
}

func (kp *KeyReplace) doReplace(fileName string, wg *sync.WaitGroup) bool {
	defer wg.Done()
	if !kp.replace(fileName) {
		kp.msg[fileName] = "replace failed"
	}
	return true
}

/*
func (kp *KeyReplace) ListenKeys() {
	for {
		select {
		case fileName, ok := <-kp.repChan:
			if !ok {
				Info("channel closed")
				return
			}
			if !kp.replace(fileName) {
				kp.msg[fileName] = "replace failed"
			}
		case <-kp.quit:
			close(kp.repChan)
			Info("end of replacing keys")
			return
		}
	}
}*/

func (kp *KeyReplace) isIgnoreKey(key string) bool {
	nKey := utils.ToUInt64(key)
	// 小于13位的数值直接忽略
	return len(key) > 13 && nKey > 0
}

func (kp *KeyReplace) formatKeyCustomTrie(content string) string {
	newContent := kp.trie.Replace(content, kp.replaceMap)
	return newContent
}

func (kp *KeyReplace) formatKeyTrie(content string) string {
	replaceMap := map[string]string{}
	var keywords []string
	for i := 1; i < len(kp.srcKeys); i++ {
		if kp.isIgnoreKey(kp.srcKeys[i]) {
			continue
		}
		replaceMap[kp.srcKeys[i]] = kp.newKeys[i]
		keywords = append(keywords, kp.srcKeys[i])
	}

	// 关键词列表
	for k := range replaceMap {
		keywords = append(keywords, k)
	}

	// 构建 AC 自动机
	matcher := ahocorasick.NewStringMatcher(keywords)

	// 先把文本分成字符数组，便于替换
	runes := []rune(content)
	positions := matcher.Match([]byte(content))

	// 标记每个匹配词的开始位置
	replaceAt := make([][2]int, 0)
	for _, idx := range positions {
		word := keywords[idx]
		loc := strings.Index(content, word) // 注意，这里不精确匹配重复内容，适合一次性替换的场景
		if loc >= 0 {
			replaceAt = append(replaceAt, [2]int{loc, len(word)})
			content = strings.Replace(content, word, strings.Repeat("$", len(word)), 1) // 避免重复匹配
		}
	}

	// 重建字符串
	newContent := string(runes)
	for _, r := range replaceAt {
		oldWord := string(runes[r[0] : r[0]+r[1]])
		newContent = strings.Replace(newContent, oldWord, replaceMap[oldWord], 1)
	}

	return newContent
}

// 循环方式，非常低效
func (kp *KeyReplace) formatKey(content string) string {
	Info("formatting keys...\n==>%s", content)
	Info("keys: %d, %d", len(kp.srcKeys), len(kp.newKeys))
	for i := 1; i < len(kp.srcKeys); i++ {
		if kp.isIgnoreKey(kp.srcKeys[i]) {
			continue
		}
		content = strings.ReplaceAll(content, kp.srcKeys[i], kp.newKeys[i])
	}
	Info("keys formatted\n==>%s", content)
	return content
}

func (kp *KeyReplace) replace(fileName string) bool {
	Info("replacing \"%s\"...\n", fileName)
	content := utils.ReadString(fileName)
	//content = kp.formatKey(content)
	//content = kp.formatKeyTrie(content)
	content = kp.formatKeyCustomTrie(content)
	err := utils.WriteString(fileName, content)
	Info("replace \"%s\" done\n", fileName)
	return err == nil
}

func (kp *KeyReplace) checkKeys() bool {
	return len(kp.srcKeys) != len(kp.newKeys)
}
