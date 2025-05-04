package configuration

import "fmt"

const (
	ErrorMsgNoSingleHost         = "в файле проекта сервера должна быть конфигурация для одного хоста"
	ErrorMsgNoInboundEndpoint    = "в файле проекта сервера должен быть указан адрес для приема данных от ЦП"
	ErrorMsgIncorrectProjectName = "неправильное имя файла проекта"
)

func GetIncorrectProjectNameErrorMessage(projectFileName string) string {
	return fmt.Sprintf("%s %s", ErrorMsgIncorrectProjectName, projectFileName)
}
