/*
文件目录工具
2025.4.18 by dralee
*/
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileType int // 文件类型
const (
	Unknow FileType = iota
	File
	Dir
)

func (f FileType) String() string {
	switch f {
	case File:
		return "file"
	case Dir:
		return "dir"
	default:
		return "unknow"
	}
}

type FileInfo struct {
	FileType FileType
	FilePath string
}

func FileExists(path string) bool {
	//fmt.Println("FileExists:", path)
	_, err := os.Stat(path)
	return err == nil
}

func FileName(path string) string {
	return filepath.Base(path)
}

func FileDir(path string) string {
	return filepath.Dir(path)
}

func FileExt(path string) string {
	return filepath.Ext(path)
}

func FileWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(FileExt(path))])
}

func ListFiles(path string) []FileInfo {
	files := []FileInfo{}
	filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			files = append(files, FileInfo{Dir, path})
		} else {
			files = append(files, FileInfo{File, path})
		}
		return nil
	})
	return files
}

func Read(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return b
}

func ReadString(path string) string {
	return string(Read(path))
}

func Write(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func WriteString(path string, data string) error {
	return Write(path, []byte(data))
}

// 将map写入文件
func WriteMapString(fileName string, strMap map[string]string) {
	builder := strings.Builder{}
	for k, v := range strMap {
		builder.WriteString(fmt.Sprintf("%s:%s\n", k, v))
	}
	WriteString(fileName, builder.String())
}
