package configuration

import (
	"encoding/xml"
)

type TestProjectData struct {
	Objects           []ObjectInfo
	Parameters        []ObjectParameter
	Attributes        []ObjectAttribute
	AttributeValues   []ObjectAttributeValue
	ParameterMappings []ObjectParameterMapping
	AttributeMappings []ObjectAttributeMapping
	Hosts             []NestedHost
	UniPlaces         []UniPlaceInfo
	UniStates         []UniStateMappingInfo
}

func CreateTestProjectInfoFromXml(projectXml string) *TestProjectData {
	var project nestedProject
	xmlData := []byte(projectXml)
	if err := xml.Unmarshal(xmlData, &project); err != nil {
		panic("не удалось загрузить файл проекта")
	}

	return &TestProjectData{
		Objects:           project.Objects,
		Hosts:             project.Hosts,
		Attributes:        project.Attributes,
		AttributeValues:   project.AttributeValues,
		AttributeMappings: project.AttributeMappings,
		Parameters:        project.Parameters,
		ParameterMappings: project.ParameterMappings,
		UniPlaces:         project.UniPlaces,
		UniStates:         project.UniStates,
	}
}
