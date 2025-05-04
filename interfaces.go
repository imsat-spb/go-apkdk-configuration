package configuration

type NetworkAddressInformation interface {
	GetHostName() string
	GetPort() int
}

type ConnectionPointsInformation interface {
	GetMessageServerConnectionPoint() NetworkAddressInformation
	GetExternalSystemConnectionPoint() NetworkAddressInformation
	GetArchiveServerConnectionPoint() NetworkAddressInformation
	GetUniProtocolConnectionPoint() NetworkAddressInformation
	GetUdpConnectionPoint() NetworkAddressInformation
	GetWebServerConnectionPoint() NetworkAddressInformation
}

type ObjectAttributeInformation interface {
	GetUnitOfMeasure() string
	GetName() string
	GetId() int
}

type ProjectInformation interface {
	GetId() int
	GetVersionId() int
	GetDeviceMap() map[int]*Device
	GetObjects() map[int]*ObjectInfo
	GetObjectHost(objectId int) int
	GetObjectInfo(id int) *ObjectInfo
	GetDeviceInfo(deviceId int) *Device
	GetAttributeInfo(id int) ObjectAttributeInformation
	GetAttributeValue(attributeId int, objectId int) *string
	GetObjectParameterInfo(id int) *ObjectParameter
	GetObjectParametersMappingsMap() map[ParameterMappingKey]*ObjectParameterMapping
	GetObjectAttributeMappingsMap() map[ParameterMappingKey]*ObjectAttributeMapping
	GetUniStates() []*UniStateMappingInfo
	GetUniPlaces() []*UniPlaceInfo
	GetConnectionPoints() ConnectionPointsInformation
	GetNwaStatesToEvents() map[int]int
}
