package configuration

import (
	"fmt"
	"strings"
)

func (nwa *networkAddress) GetHostName() string {
	return nwa.Ip
}

func (nwa *networkAddress) GetPort() int {
	return nwa.Port
}

type projectInfo struct {
	projectInfo    *multiProject
	nestedProjects map[string]*nestedProject
	projectId      int
	versionId      int

	objectsMap                 map[int]*ObjectInfo
	deviceMap                  map[int]*Device
	objectParametersMap        map[int]*ObjectParameter
	objectParametersMappingMap map[ParameterMappingKey]*ObjectParameterMapping
	stationMap                 map[int]*stationInfo
	spanMap                    map[int]*spanInfo
	objectTypeMap              map[int]*objectTypeInfo
	objectAttributesMappingMap map[ParameterMappingKey]*ObjectAttributeMapping
	objectAttributesMap        map[int]*ObjectAttribute
	objectAttributeValuesMap   map[attributeValueKey]string
	objectsToHosts             map[int]int
	uniStatesMap               []*UniStateMappingInfo
	uniPlaces                  []*UniPlaceInfo
	nwaStateToEventMap         map[int]int
}

func (project *projectInfo) GetId() int {
	return project.projectId
}

func (project *projectInfo) GetVersionId() int {
	return project.versionId
}

func (project *projectInfo) GetDeviceMap() map[int]*Device {
	return project.getDeviceMap()
}

func (project *projectInfo) GetObjects() map[int]*ObjectInfo {
	return project.getObjects()
}

func (project *projectInfo) GetObjectHost(objectId int) int {
	return project.getObjectsToHostsMap()[objectId]
}

func (project *projectInfo) GetObjectInfo(id int) *ObjectInfo {
	return project.getObjects()[id]
}

func (project *projectInfo) GetDeviceInfo(deviceId int) *Device {
	return project.GetDeviceMap()[deviceId]
}

func (project *projectInfo) GetAttributeValue(attributeId int, objectId int) *string {
	s, ok := project.getObjectAttributeValuesMap()[attributeValueKey{attributeId: attributeId, objectId: objectId}]

	if ok {
		return &s
	}

	return nil
}

func (project *projectInfo) GetAttributeInfo(id int) ObjectAttributeInformation {
	result := project.getObjectAttributesMap()[id]

	if result == nil {
		return nil
	}

	return result
}

func (project *projectInfo) GetObjectParameterInfo(id int) *ObjectParameter {
	return project.getObjectParametersMap()[id]
}

func (project *projectInfo) GetObjectParametersMappingsMap() map[ParameterMappingKey]*ObjectParameterMapping {

	if project.objectParametersMappingMap != nil {
		return project.objectParametersMappingMap
	}

	result := make(map[ParameterMappingKey]*ObjectParameterMapping)

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.getObjectParametersMappingsMap()

		for key, pmInfo := range nestedMap {
			result[key] = pmInfo
		}
	}

	project.objectParametersMappingMap = result
	return result
}

func (project *projectInfo) GetObjectAttributeMappingsMap() map[ParameterMappingKey]*ObjectAttributeMapping {

	if project.objectAttributesMappingMap != nil {
		return project.objectAttributesMappingMap
	}

	result := make(map[ParameterMappingKey]*ObjectAttributeMapping)

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.getObjectAttributesMappingsMap()

		for key, pmInfo := range nestedMap {
			result[key] = pmInfo
		}
	}

	project.objectAttributesMappingMap = result
	return result
}

func (project *projectInfo) GetUniStates() []*UniStateMappingInfo {

	// TODO: Optimize it
	if project.uniStatesMap != nil {
		return project.uniStatesMap
	}

	var result []*UniStateMappingInfo

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.UniStates

		for index := range nestedMap {
			result = append(result, &nestedMap[index])
		}
	}

	project.uniStatesMap = result

	return result
}

func (project *projectInfo) GetUniPlaces() []*UniPlaceInfo {
	if project.uniPlaces != nil {
		return project.uniPlaces
	}

	var result []*UniPlaceInfo

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.UniPlaces

		for index := range nestedMap {
			result = append(result, &nestedMap[index])
		}
	}

	project.uniPlaces = result
	return result
}

func (project *projectInfo) GetConnectionPoints() ConnectionPointsInformation {
	return project
}

func (project *projectInfo) GetWebServerConnectionPoint() NetworkAddressInformation {
	host := project.getUniServerHost()

	if host == nil || host.WebServerConnectionPoint == nil {
		return nil
	}

	return host.WebServerConnectionPoint
}

func (project *projectInfo) GetMessageServerConnectionPoint() NetworkAddressInformation {
	host := project.getUniServerHost()

	if host == nil || host.MessageServerConnectionPoint == nil {
		return nil
	}

	return host.MessageServerConnectionPoint
}

func (project *projectInfo) GetExternalSystemConnectionPoint() NetworkAddressInformation {
	host := project.getUniServerHost()

	if host == nil || host.ConnectionPoint == nil {
		return nil
	}

	return host.ConnectionPoint
}

func (project *projectInfo) GetArchiveServerConnectionPoint() NetworkAddressInformation {
	host := project.getUniServerHost()

	if host == nil || host.ArchiveServerConnectionPoint == nil {
		return nil
	}

	return host.ArchiveServerConnectionPoint
}

func (project *projectInfo) GetUniProtocolConnectionPoint() NetworkAddressInformation {
	host := project.getUniServerHost()

	if host == nil || host.UniProtocolConnectionPoint == nil {
		return nil
	}

	return host.UniProtocolConnectionPoint
}

func (project *projectInfo) GetUdpConnectionPoint() NetworkAddressInformation {
	host := project.getUniServerHost()

	if host == nil {
		return nil
	}

	if len(host.InboundPoints) != 1 {
		return nil
	}

	return &host.InboundPoints[0]
}

func (project *projectInfo) GetNwaStatesToEvents() map[int]int {
	if project.nwaStateToEventMap != nil {
		return project.nwaStateToEventMap
	}

	result := make(map[int]int)

	for _, innerProject := range project.nestedProjects {
		nestedNwaStatesMap := innerProject.getNwaStateToEventsMap()

		for stateId, eventId := range nestedNwaStatesMap {
			result[stateId] = eventId
		}
	}

	project.nwaStateToEventMap = result

	return result
}

func (project *projectInfo) getStationsMap() map[int]*stationInfo {

	if project.stationMap != nil {
		return project.stationMap
	}

	result := make(map[int]*stationInfo)

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.getStationsMap()

		for key, sInfo := range nestedMap {
			result[key] = sInfo
		}
	}

	project.stationMap = result
	return result
}

func (project *projectInfo) getSpansMap() map[int]*spanInfo {

	if project.spanMap != nil {
		return project.spanMap
	}

	result := make(map[int]*spanInfo)

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.getSpansMap()

		for key, sInfo := range nestedMap {
			result[key] = sInfo
		}
	}

	project.spanMap = result
	return result
}

func (project *projectInfo) getObjectTypesMap() map[int]*objectTypeInfo {

	if project.objectTypeMap != nil {
		return project.objectTypeMap
	}

	result := make(map[int]*objectTypeInfo)

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.getObjectTypesMap()

		for key, sInfo := range nestedMap {
			result[key] = sInfo
		}
	}

	project.objectTypeMap = result
	return result
}

func (project *projectInfo) getStationName(id int) string {
	s := project.getStationsMap()[id]

	if s == nil {
		return ""
	}

	return s.Name
}

func (project *projectInfo) getSpanName(id int) string {
	s := project.getSpansMap()[id]

	if s == nil {
		return ""
	}

	return s.Name
}

func (project *projectInfo) getObjectTypeName(id int) string {
	t := project.getObjectTypesMap()[id]

	if t == nil {
		return ""
	}

	return t.Name
}

func (project *projectInfo) getObjectsToHostsMap() map[int]int {
	if project.objectsToHosts != nil {
		return project.objectsToHosts
	}

	result := make(map[int]int)

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.getObjectsToHostsMap()

		for objId, hostId := range nestedMap {
			result[objId] = hostId
		}
	}
	project.objectsToHosts = result
	return result
}

func (project *projectInfo) getObjects() map[int]*ObjectInfo {
	if project.objectsMap != nil {
		return project.objectsMap
	}

	result := make(map[int]*ObjectInfo)

	for _, innerProject := range project.nestedProjects {
		nestedMap := innerProject.getObjectsMap()

		for _, objInfo := range nestedMap {
			result[objInfo.Id] = objInfo
		}
	}
	project.objectsMap = result
	return result
}

func (project *projectInfo) getUniServerHost() *hostInfo {
	if project.projectInfo == nil {
		return nil
	}

	if len(project.projectInfo.Hosts) != 1 {
		return nil
	}

	return &project.projectInfo.Hosts[0]
}

func (project *projectInfo) getDeviceMap() map[int]*Device {
	if project.deviceMap != nil {
		return project.deviceMap
	}

	result := make(map[int]*Device)

	for _, innerProject := range project.nestedProjects {
		nestedDevMap := innerProject.getDeviceMap()

		for _, devInfo := range nestedDevMap {
			result[devInfo.Id] = devInfo
		}
	}

	project.deviceMap = result
	return result
}

func (project *projectInfo) getObjectParametersMap() map[int]*ObjectParameter {
	if project.objectParametersMap != nil {
		return project.objectParametersMap
	}

	result := make(map[int]*ObjectParameter)

	for _, innerProject := range project.nestedProjects {
		nestedParamMap := innerProject.getObjectParametersMap()

		for _, paramInfo := range nestedParamMap {
			result[paramInfo.Id] = paramInfo
		}
	}

	project.objectParametersMap = result

	return result
}

func (project *projectInfo) getObjectAttributeValuesMap() map[attributeValueKey]string {
	if project.objectAttributeValuesMap != nil {
		return project.objectAttributeValuesMap
	}

	result := make(map[attributeValueKey]string)

	for _, innerProject := range project.nestedProjects {
		nestedParamMap := innerProject.getObjectAttributeValuesMap()

		for k, v := range nestedParamMap {
			result[k] = v
		}
	}

	project.objectAttributeValuesMap = result

	return result
}

func (project *projectInfo) getObjectAttributesMap() map[int]*ObjectAttribute {
	if project.objectAttributesMap != nil {
		return project.objectAttributesMap
	}

	result := make(map[int]*ObjectAttribute)

	for _, innerProject := range project.nestedProjects {
		nestedParamMap := innerProject.getObjectAttributesMap()

		for _, attrInfo := range nestedParamMap {
			result[attrInfo.Id] = attrInfo
		}
	}

	project.objectAttributesMap = result

	return result
}

func (project *multiProject) validate() error {
	if project.Hosts == nil || len(project.Hosts) != 1 {
		return fmt.Errorf(ErrorMsgNoSingleHost)
	}

	if project.Hosts[0].InboundPoints == nil || len(project.Hosts[0].InboundPoints) != 1 {
		return fmt.Errorf(ErrorMsgNoInboundEndpoint)
	}

	// TODO: other checks here
	return nil
}

func (parameter *ObjectParameter) GetParameterDisplayName() string {
	if parameter.ShortName != "" {
		return parameter.ShortName
	}
	return parameter.Name
}

func (parameter *ObjectParameter) GetUnitOfMeasureDisplayName() string {
	if parameter.UnitOfMeasure == "" {
		return ""
	}

	parts := strings.Split(parameter.UnitOfMeasure, ",")

	if len(parts) < 2 {
		return ""
	}

	return strings.Trim(parts[len(parts)-1], " ")
}

type ParameterMappingKey struct {
	objectId  int
	measureId int
}

type UniStateMappingKey struct {
}

func (key *ParameterMappingKey) GetObjectId() int {
	return key.objectId
}

func (key *ParameterMappingKey) GetMeasureId() int {
	return key.measureId
}

func NewParameterMappingKey(objectId int, measureId int) ParameterMappingKey {
	return ParameterMappingKey{objectId: objectId, measureId: measureId}
}
