package configuration

import (
	"encoding/xml"
)

type TestProjectData struct {
	Devices           map[int]*Device
	Objects           map[int]*ObjectInfo
	ObjectsToHosts    map[int]int
	NwaStatesToEvents map[int]int
	Parameters        map[int]*ObjectParameter
	Attributes        map[int]*ObjectAttribute
	AttributeValues   []ObjectAttributeValue
	ParameterMappings map[ParameterMappingKey]*ObjectParameterMapping
	AttributeMappings map[ParameterMappingKey]*ObjectAttributeMapping
	Hosts             []NestedHost
	UniPlaces         []UniPlaceInfo
	UniStates         []UniStateMappingInfo
}

/*type ProjectInformation interface {
	GetUniStates() []*UniStateMappingInfo
	GetUniPlaces() []*UniPlaceInfo
}*/

func (td *TestProjectData) GetId() int {
	return 1
}

func (td *TestProjectData) GetVersionId() int {
	return 1
}

func (td *TestProjectData) GetObjects() map[int]*ObjectInfo {
	return td.Objects
}

func (td *TestProjectData) GetObjectInfo(id int) *ObjectInfo {
	o, ok := td.Objects[id]

	if ok {
		return o
	} else {
		return nil
	}
}

func (td *TestProjectData) GetObjectHost(objectId int) int {
	return td.ObjectsToHosts[objectId]
}

func (td *TestProjectData) GetDeviceMap() map[int]*Device {
	return td.Devices
}

func (td *TestProjectData) GetDeviceInfo(deviceId int) *Device {
	d, ok := td.Devices[deviceId]

	if ok {
		return d
	} else {
		return nil
	}
}

func (td *TestProjectData) GetObjectParameterInfo(id int) *ObjectParameter {
	p, ok := td.Parameters[id]

	if ok {
		return p
	} else {
		return nil
	}
}

func (td *TestProjectData) GetAttributeInfo(id int) ObjectAttributeInformation {
	a, ok := td.Attributes[id]

	if ok {
		return a
	} else {
		return nil
	}
}

func (td *TestProjectData) GetAttributeValue(attributeId int, objectId int) *string {
	for _, v := range td.AttributeValues {
		if v.Id == objectId && v.Id == attributeId {
			return &v.Value
		}
	}
	return nil
}

func (td *TestProjectData) GetObjectParametersMappingsMap() map[ParameterMappingKey]*ObjectParameterMapping {
	return td.ParameterMappings
}

func (td *TestProjectData) GetObjectAttributeMappingsMap() map[ParameterMappingKey]*ObjectAttributeMapping {
	return td.AttributeMappings
}

func (td *TestProjectData) GetConnectionPoints() ConnectionPointsInformation {
	return nil
}

func (td *TestProjectData) GetNwaStatesToEvents() map[int]int {
	return td.NwaStatesToEvents
}

func (td *TestProjectData) GetUniStates() []*UniStateMappingInfo {
	var result []*UniStateMappingInfo

	for index := range td.UniStates {
		result = append(result, &td.UniStates[index])
	}

	return result
}

func (td *TestProjectData) GetUniPlaces() []*UniPlaceInfo {
	var result []*UniPlaceInfo

	for index := range td.UniPlaces {
		result = append(result, &td.UniPlaces[index])
	}

	return result
}

func CreateTestProjectInfoFromXml(projectXml string) *TestProjectData {
	var project nestedProject
	xmlData := []byte(projectXml)
	if err := xml.Unmarshal(xmlData, &project); err != nil {
		panic("не удалось загрузить файл проекта")
	}

	return &TestProjectData{
		Devices:           project.getDeviceMap(),
		Objects:           project.getObjectsMap(),
		ObjectsToHosts:    project.getObjectsToHostsMap(),
		NwaStatesToEvents: project.getNwaStateToEventsMap(),
		Hosts:             project.Hosts,
		Attributes:        project.getObjectAttributesMap(),
		AttributeValues:   project.AttributeValues,
		AttributeMappings: project.getObjectAttributesMappingsMap(),
		Parameters:        project.getObjectParametersMap(),
		ParameterMappings: project.getObjectParametersMappingsMap(),
		UniPlaces:         project.UniPlaces,
		UniStates:         project.UniStates,
	}
}
