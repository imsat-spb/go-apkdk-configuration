package configuration

func (project *nestedProject) getObjectParametersMappingsMap() map[ParameterMappingKey]*ObjectParameterMapping {

	if project.objectParametersMappingMap != nil {
		return project.objectParametersMappingMap
	}

	var measureMappingsMap = make(map[ParameterMappingKey]*ObjectParameterMapping)

	for i := range project.ParameterMappings {
		pm := &project.ParameterMappings[i]

		measureMappingsMap[NewParameterMappingKey(pm.ObjectId, pm.Id)] = pm
	}

	project.objectParametersMappingMap = measureMappingsMap
	return measureMappingsMap
}

func (project *nestedProject) getObjectAttributesMappingsMap() map[ParameterMappingKey]*ObjectAttributeMapping {

	if project.objectAttributesMappingMap != nil {
		return project.objectAttributesMappingMap
	}

	result := make(map[ParameterMappingKey]*ObjectAttributeMapping)

	for i := range project.AttributeMappings {
		pm := &project.AttributeMappings[i]

		result[NewParameterMappingKey(pm.ObjectId, pm.Id)] = pm
	}

	project.objectAttributesMappingMap = result
	return result
}

func (project *nestedProject) getDeviceMap() map[int]*Device {
	if project.deviceMap != nil {
		return project.deviceMap
	}

	devMap := make(map[int]*Device)

	for i := range project.Hosts {
		h := &project.Hosts[i]

		for j := range h.Devices {
			d := &h.Devices[j]
			devMap[d.Id] = d
		}
	}
	project.deviceMap = devMap
	return devMap
}

func (project *nestedProject) getStationsMap() map[int]*stationInfo {

	if project.stationMap != nil {
		return project.stationMap
	}

	var stationsMap = make(map[int]*stationInfo)

	for i := range project.Stations {
		s := &project.Stations[i]

		stationsMap[s.Id] = s
	}

	project.stationMap = stationsMap
	return stationsMap
}

func (project *nestedProject) getSpansMap() map[int]*spanInfo {

	if project.spanMap != nil {
		return project.spanMap
	}

	var spansMap = make(map[int]*spanInfo)

	for i := range project.Spans {
		s := &project.Spans[i]

		spansMap[s.Id] = s
	}

	project.spanMap = spansMap
	return spansMap
}

func (project *nestedProject) getObjectTypesMap() map[int]*objectTypeInfo {

	if project.objectTypeMap != nil {
		return project.objectTypeMap
	}

	var typesMap = make(map[int]*objectTypeInfo)

	for i := range project.ObjectTypes {
		t := &project.ObjectTypes[i]

		typesMap[t.Id] = t
	}

	project.objectTypeMap = typesMap
	return typesMap
}

func (project *nestedProject) getObjectsToHostsMap() map[int]int {
	if project.objectsToHosts != nil {
		return project.objectsToHosts
	}

	var objInfoMap = make(map[int]int, len(project.Objects))
	for _, objectIds := range project.ObjectsOnHost {
		for _, id := range objectIds.ObjectIds {
			objInfoMap[id.ObjectId] = objectIds.HostId
		}
	}

	project.objectsToHosts = objInfoMap

	return objInfoMap
}

func (project *nestedProject) getObjectsMap() map[int]*ObjectInfo {
	if project.objectsMap != nil {
		return project.objectsMap
	}

	var objInfoMap = make(map[int]*ObjectInfo, len(project.Objects))
	for i := range project.Objects {
		p := &project.Objects[i]
		objInfoMap[p.Id] = p
	}

	project.objectsMap = objInfoMap

	return objInfoMap
}

func (project *nestedProject) getObjectParametersMap() map[int]*ObjectParameter {
	if project.objectParametersMap != nil {
		return project.objectParametersMap
	}

	var measureInfoMap = make(map[int]*ObjectParameter, len(project.Parameters))
	for i := range project.Parameters {
		p := &project.Parameters[i]
		measureInfoMap[p.Id] = p
	}

	project.objectParametersMap = measureInfoMap

	return measureInfoMap
}

func (project *nestedProject) getObjectAttributeValuesMap() map[attributeValueKey]string {
	if project.objectAttributeValuesMap != nil {
		return project.objectAttributeValuesMap
	}

	result := make(map[attributeValueKey]string, len(project.AttributeValues))
	for _, v := range project.AttributeValues {
		result[attributeValueKey{v.Id, v.ObjectId}] = v.Value
	}

	project.objectAttributeValuesMap = result

	return result
}

func (project *nestedProject) getObjectAttributesMap() map[int]*ObjectAttribute {
	if project.objectAttributesMap != nil {
		return project.objectAttributesMap
	}

	result := make(map[int]*ObjectAttribute, len(project.Attributes))
	for i := range project.Attributes {
		p := &project.Attributes[i]
		result[p.Id] = p
	}

	project.objectAttributesMap = result

	return result
}

func (project *nestedProject) getNwaStateToEventsMap() map[int]int {
	if project.nwaStateToEventMap != nil {
		return project.nwaStateToEventMap
	}

	result := make(map[int]int)

	for _, nwa := range project.NormalWorkFlows {
		for _, nwaState := range nwa.FlowStates {
			if nwaState.EventId == nil {
				continue
			}
			result[nwaState.Id] = *nwaState.EventId
		}
	}

	project.nwaStateToEventMap = result

	return result
}
