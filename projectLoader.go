package configuration

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func getProjectEntry(file *zip.File) (*multiProject, error) {
	xmlFile, err := file.Open()

	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать содержимое файла проекта сервера в архиве")
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать содержимое файла проекта сервера в архиве")
	}

	var project multiProject
	if err := xml.Unmarshal(byteValue, &project); err != nil {
		return nil, fmt.Errorf("не удалось загрузить файл проекта сервера из архива")
	}
	return &project, nil
}

func getNestedProjectEntry(file *zip.File) (*nestedProject, error) {
	xmlFile, err := file.Open()

	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать содержимое вложенного файла проекта %s", file.Name)
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать содержимое вложенного файла проекта %s", file.Name)
	}

	var project nestedProject
	if err := xml.Unmarshal(byteValue, &project); err != nil {
		return nil, fmt.Errorf("не удалось загрузить файл вложенного проекта %s", file.Name)
	}
	return &project, nil
}

func getNestedZippedProjectEntry(file *zip.File) (*nestedProject, error) {
	zipXmlFile, err := file.Open()

	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать содержимое вложенного файла проекта %s", file.Name)
	}

	defer zipXmlFile.Close()

	byteValue, err := ioutil.ReadAll(zipXmlFile)

	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать содержимое вложенного файла проекта %s", file.Name)
	}

	subProjectId, _, err := getProjectAndVersionIds(file.Name)

	if err != nil {
		return nil, err
	}

	projectEntryName := fmt.Sprintf("project_%d.prj", subProjectId)

	reader, err := zip.NewReader(bytes.NewReader(byteValue), int64(len(byteValue)))

	for _, f := range reader.File {
		if f.Name == projectEntryName {
			return getNestedProjectEntry(f)
		}
	}

	return nil, fmt.Errorf("не найден файл проекта в архиве %s", file.Name)
}

func loadServerProjectInfo(reader *zip.ReadCloser, projectId int) (*multiProject, map[string]*nestedProject, error) {
	var projectEntryName = fmt.Sprintf("project_%d.prj", projectId)

	nestedProjects := make(map[string]*nestedProject)

	var project *multiProject
	var err error
	for _, f := range reader.File {
		if f.Name == projectEntryName {
			project, err = getProjectEntry(f)
			if err != nil {
				return nil, nestedProjects, err
			}
		} else {
			// Проверяем на допустимое имя вложенного файла проекта
			if isValidProjectName(f.Name) {
				if nestedProjectEntry, err := getNestedZippedProjectEntry(f); err != nil {
					return project, nestedProjects, err
				} else {
					nestedProjects[f.Name] = nestedProjectEntry
				}
			}
		}
	}

	if project == nil {
		return nil, nestedProjects, fmt.Errorf("не найден файл проекта сервера в архиве")
	}
	return project, nestedProjects, nil

}

func FindMaxProjectVersion(projectFolder string, projectId int) (int, error) {
	files, err := ioutil.ReadDir(projectFolder)

	if err != nil {
		return -1, err
	}

	var currentVersion = -1

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		var name = f.Name()

		if !isValidProjectName(name) {
			continue
		}

		s := strings.Split(name, "_")

		id, _ := strconv.Atoi(s[1])

		if id != projectId {
			continue
		}

		s = strings.Split(s[2], ".")

		verId, _ := strconv.Atoi(s[0])

		if verId > currentVersion {
			currentVersion = verId
		}
	}

	return currentVersion, nil
}

func LoadServerProjectInfo(projectFilePath string) (ProjectInformation, error) {
	_, fileName := filepath.Split(projectFilePath)

	projectId, versionId, err := getProjectAndVersionIds(fileName)

	if err != nil {
		return nil, err
	}

	reader, err := zip.OpenReader(projectFilePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файла проекта %s", projectFilePath)
	}

	defer reader.Close()

	var project *multiProject
	var nestedProjects map[string]*nestedProject

	project, nestedProjects, err = loadServerProjectInfo(reader, projectId)

	if err != nil {
		return nil, err
	}

	result := &projectInfo{
		projectInfo:    project,
		nestedProjects: nestedProjects,
		projectId:      projectId,
		versionId:      versionId,
	}
	if err := project.validate(); err != nil {
		return result, err
	}

	return result, nil
}
