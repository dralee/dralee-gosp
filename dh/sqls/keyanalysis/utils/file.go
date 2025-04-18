/*
文件目录工具
2025.4.18 by dralee
*/
package utils

import (
	"os"
	"path/filepath"
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
