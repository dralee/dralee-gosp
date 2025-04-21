/*
主键分析
2025.4.18 by dralee
*/
package main

import (
	"fmt"
	"keyanalysis/utils"
	"strings"
	"sync"
)

const (
	KeyReg       = `\((\d+),`
	FileIndexReg = `FAT_(\d+)_`
)

type KeyAnalysis struct {
	srcPath    string
	keys       []string
	existsKeys map[string][]string
	mutex      sync.Mutex
	channel    chan [2]string
	quit       chan bool
}

// 实例化对象
func NewKeyAnalysis(srcPath string) *KeyAnalysis {
	return &KeyAnalysis{
		srcPath:    srcPath,
		keys:       make([]string, 0),
		existsKeys: map[string][]string{},
		mutex:      sync.Mutex{},
		channel:    make(chan [2]string, 10),
		quit:       make(chan bool),
	}
}

// // 忽略文件
// func (k *KeyAnalysis) isIgnoreFile(srcPath string) bool {
// 	fileName := utils.FileName(srcPath)
// 	fileExt := utils.FileExt(srcPath)
// 	if !strings.HasPrefix(fileName, "FAT_") {
// 		return true
// 	}

// 	if strings.Contains(fileName, "xxljob") {
// 		return true
// 	}

// 	if fileExt != ".sql" {
// 		return true
// 	}

// 	num := utils.FindGroupNum(FileIndexReg, fileName, 1)
// 	return num < 3
// }

// 扫描
func (k *KeyAnalysis) Scan() bool {
	if k.srcPath == "" {
		return false
	}

	files := utils.ListFiles(k.srcPath)
	var wg sync.WaitGroup

	go k.ListenKeys()

	for _, file := range files {
		if file.FileType != utils.File {
			continue
		}
		if IsIgnoreFile(file.FilePath) {
			Info("ignore file: %s\n", file.FilePath)
			continue
		}
		wg.Add(1)
		go k.scanFile(file.FilePath, &wg)
	}

	go func() {
		wg.Wait()
		k.quit <- true
	}()

	Info("waiting for keys")
	<-k.quit
	Info("scan all done")

	return true
}

// 保存
func (k *KeyAnalysis) Save(fileName string) bool {
	keyContent := strings.Join(k.keys, ",")
	content := fmt.Sprintf("keys(%d):\n", len(k.keys))
	content += keyContent
	content += "\n\n===============================================\n"
	for fileName, keys := range k.existsKeys {
		content += fmt.Sprintf("%s: %s\n", fileName, strings.Join(keys, ","))
	}
	err := utils.WriteString(fileName, content)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = utils.WriteString(SrcKeyFile, keyContent)
	return err == nil
}

// 扫描单个文件
func (k *KeyAnalysis) scanFile(fileName string, wg *sync.WaitGroup) bool {
	defer wg.Done()

	if !utils.FileExists(fileName) {
		return false
	}

	Info("scanning \"%s\"...\n", fileName)
	content := utils.ReadString(fileName)
	keys := utils.FindAllGroup(KeyReg, content, 1)
	//k.addKeys(fileName, keys)
	for _, key := range keys {
		k.channel <- [2]string{fileName, key}
	}
	Info("scan \"%s\" done\n", fileName)
	return true
}

// 监听主键
func (k *KeyAnalysis) ListenKeys() {
	for {
		select {
		case data, ok := <-k.channel:
			if !ok {
				Info("channel closed")
				return
			}
			fileName := data[0]
			key := data[1]
			if utils.Contains(k.keys, key) {
				k.existsKeys[fileName] = append(k.existsKeys[fileName], key)
				continue
			}
			k.keys = append(k.keys, key)
		case <-k.quit:
			Info("quit")
			close(k.channel)
			return
		}
	}
}

func (k *KeyAnalysis) addKeys(fileName string, keys []string) bool {
	if keys == nil {
		return false
	}
	defer k.mutex.Unlock()
	k.mutex.Lock()
	for _, key := range keys {
		if utils.Contains(k.keys, key) {
			k.existsKeys[fileName] = append(k.existsKeys[fileName], key)
			continue
		}
		k.keys = append(k.keys, key)
	}
	return true
}
