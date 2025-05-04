package configuration

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const projectZipNamePattern = "^project_\\d+_\\d+.zip$"

func isValidProjectName(fileName string) bool {
	return regexp.MustCompile(projectZipNamePattern).MatchString(fileName)
}

func getProjectAndVersionIds(fileName string) (int, int, error) {
	if !isValidProjectName(fileName) {
		return -1, -1, errors.New(GetIncorrectProjectNameErrorMessage(fileName))
	}
	s := strings.Split(fileName, "_")
	id, err := strconv.Atoi(s[1])

	if err != nil {
		return -1, -1, fmt.Errorf("ошибка при получении идентификатора проекта %s", fileName)
	}

	versionId, err := strconv.Atoi(strings.Split(s[2], ".")[0])

	if err != nil {
		return -1, -1, fmt.Errorf("ошибка при получении версии проекта %s", fileName)
	}

	return id, versionId, nil
}
