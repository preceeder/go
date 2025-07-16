package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// WritePidToFile 写入当前pid到文件中
func WritePidToFile(fileNames ...string) (pid string, err error) {
	pidi := os.Getpid()
	pid = strconv.Itoa(pidi)
	var file *os.File
	var fileName string = "socket.pid"
	if len(fileNames) > 0 {
		fileName = fileNames[0]
	}
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		fmt.Println("WritePidToFile error 1", err.Error())
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Print("WritePidToFile close fs error", err.Error())
		}
	}(file)

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte(pid))
	if err != nil {
		fmt.Println("WritePidToFile error 2", err.Error())

		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("WritePidToFile error 3", err.Error())
		return
	}
	return
}

// CleanPidFile 清空pid文件
func CleanPidFile(fileNames ...string) (data []byte, err error) {
	var file *os.File
	var fileName string = "socket.pid"
	if len(fileNames) > 0 {
		fileName = fileNames[0]
	}
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
	if err != nil {
		fmt.Println("CleanPidFile error 1", err.Error())
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Print("CleanPidFile close fs error", err.Error())
		}
	}(file)

	// 读出内容返回
	data, err = io.ReadAll(file)
	if err != nil {
		return
	}

	err = file.Truncate(0)

	//// 写空值
	//writer := bufio.NewWriter(file)
	//_, err = writer.Write([]byte(""))
	//if err != nil {
	//	fmt.Println("CleanPidFile error 2", err.Error())
	//	return
	//}
	//
	//err = writer.Flush()
	//if err != nil {
	//	fmt.Println("CleanPidFile error 3", err.Error())
	//	return
	//}
	return
}
