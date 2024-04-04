package service

import (
	"bufio"
	"os"

	"github.com/mouday/cron-admin/src/utils"
)

func AppendLog(taskId string, taskLogId string, text string) {
	utils.MakeDir("./logs/" + taskId)

	logFile := "./logs/" + taskId + "/" + taskLogId + ".log"

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(text + "\n")
	//Flush将缓存的文件真正写入到文件中
	write.Flush()

}

func ReadLog(taskId string, taskLogId string) string {
	logFile := "./logs/" + taskId + "/" + taskLogId + ".log"

	content, err := os.ReadFile(logFile)
	if err != nil {
		panic(err)
	}

	return string(content)
}
