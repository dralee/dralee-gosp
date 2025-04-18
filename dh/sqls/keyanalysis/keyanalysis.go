package main

import (
	"fmt"
	"keyanalysis/utils"
	"strings"
	"sync"
)

const (
	KeyReg = `\((\d+),`
)

type KeyAnalysis struct {
	srcPath    string
	keys       []string
	existsKeys map[string][]string
	mutex      sync.Mutex
	channel    chan [2]string
	quit       chan bool
}

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
		wg.Add(1)
		go k.scanFile(file.FilePath, &wg)
	}

	go func() {
		wg.Wait()
		k.quit <- true
	}()

	fmt.Println("waiting for keys")
	<-k.quit
	fmt.Println("scan all done")

	return true
}

func (k *KeyAnalysis) Save(fileName string) bool {
	content := strings.Join(k.keys, ",")
	content += "===============================================\n"
	for fileName, keys := range k.existsKeys {
		content += fmt.Sprintf("%s: %s\n", fileName, strings.Join(keys, ","))
	}
	err := utils.WriteString(fileName, content)
	return err == nil
}

func (k *KeyAnalysis) scanFile(fileName string, wg *sync.WaitGroup) bool {
	defer wg.Done()

	if !utils.FileExists(fileName) {
		return false
	}

	content := utils.ReadString(fileName)
	keys := utils.FindAllGroup(KeyReg, content, 1)
	//k.addKeys(fileName, keys)
	for _, key := range keys {
		k.channel <- [2]string{fileName, key}
	}
	fmt.Printf("scan \"%s\" done\n", fileName)
	return true
}

func (k *KeyAnalysis) ListenKeys() {
	for {
		select {
		case data, ok := <-k.channel:
			if !ok {
				fmt.Println("channel closed")
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
			fmt.Println("quit")
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
