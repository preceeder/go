package base

import (
	"bufio"
	"io"
	"log/slog"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("file open fail", "file_name", filePath, "error", err.Error())
		return nil, err
	}
	defer file.Close()
	fd, err := io.ReadAll(file)
	if err != nil {
		slog.Error("read to fd fail", "error", err.Error())
		return nil, err
	}
	return fd, nil
}

// WriteFile 清空文件后在写
func WriteFile(filePath string) (*os.File, *bufio.Writer) {
	// os.FileMode(0777) -rwxrwxrwx
	//os.FileMode(0666) -rw-rw-rw-
	//os.FileMode(0644) -rw-r--r--
	// 写完后最好 用 	writer.Flush() 保存到磁盘

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
	if err != nil {
		slog.Error("打开文件失败", "error", err.Error())
	}
	file.Truncate(0)
	//defer file.Close()

	writer := bufio.NewWriter(file)
	return file, writer
}
