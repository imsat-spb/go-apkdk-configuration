package configuration

import "encoding/xml"

type multiProject struct {
	XMLName xml.Name `xml:"Project"`
	// Хосты. В составном файле проекта обязательно должен быть один хост
	Hosts              []hostInfo `xml:"Hosts>Host"`
	SupportUniProtocol bool       `xml:"SupportUniProtocol,attr"`
	SendObjectStates   bool       `xml:"SendObjectStates,attr"`
}

type sensorRange struct {
	XMLName xml.Name `xml:"Range"`
	IdFrom  int      `xml:"IdFrom,attr"`
	IdTo    int      `xml:"IdTo,attr"`
}

type sensorRanges struct {
	XMLName   xml.Name      `xml:"Ranges"`
	Ranges    []sensorRange `xml:"Range"`
	IsInclude bool          `xml:"IsInclude,attr"`
}

type Device struct {
	XMLName       xml.Name      `xml:"Device"`
	Ranges        *sensorRanges `xml:"Ranges,omitempty"`
	Id            int           `xml:"ID,attr"`
	BitsPerSensor int           `xml:"BitsOnSensor,attr"`
	SensorCount   int           `xml:"SensorCount,attr"`
}

type NestedHost struct {
	XMLName xml.Name `xml:"Host"`
	Devices []Device `xml:"DataHub>Device"`
	Id      int      `xml:"Number,attr"`
}

type ObjectInfo struct {
	XMLName      xml.Name `xml:"Object"`
	Id           int      `xml:"Id,attr"`
	Name         string   `xml:"Name,attr"`
	TypeId       int      `xml:"TypeId,attr"`
	StationId    int      `xml:"StationId,attr"`
	SpanId       *int     `xml:"SpanId,attr"`
	MainObjectId *int     `xml:"MainObjectId,attr"`
}

type objectTypeInfo struct {
	XMLName xml.Name `xml:"ObjectType"`
	Id      int      `xml:"Id,attr"`
	Name    string   `xml:"Name,attr"`
}

type stationInfo struct {
	XMLName xml.Name `xml:"Station"`
	Id      int      `xml:"Id,attr"`
	Name    string   `xml:"Name,attr"`
}

type spanInfo struct {
	XMLName xml.Name `xml:"Span"`
	Id      int      `xml:"Id,attr"`
	Name    string   `xml:"Name,attr"`
}

type ObjectParameter struct {
	XMLName       xml.Name `xml:"ObjectParameter"`
	Id            int      `xml:"Id,attr"`
	Name          string   `xml:"Name,attr"`
	ShortName     string   `xml:"ShortName,attr"`
	UnitOfMeasure string   `xml:"UnitOfMeasure,attr"`
}

type ObjectAttribute struct {
	XMLName       xml.Name `xml:"ObjectAttributeType"`
	Id            int      `xml:"Id,attr"`
	Name          string   `xml:"Name,attr"`
	UnitOfMeasure string   `xml:"UnitOfMeasure,attr"`
}

type ObjectAttributeValue struct {
	XMLName  xml.Name `xml:"ObjectAttribute"`
	Id       int      `xml:"AttributeTypeId,attr"`
	ObjectId int      `xml:"ObjectId,attr"`
	Value    string   `xml:"Value,attr"`
}

type networkAddress struct {
	XMLName xml.Name `xml:"Address"`
	Ip      string   `xml:"Ip,attr"`
	Port    int      `xml:"Port,attr"`
}

type hostInfo struct {
	XMLName xml.Name `xml:"Host"`
	// Точка подключения может быть не задана. В этом случае АРМ не подключится
	ConnectionPoint            *networkAddress `xml:"ExternalConnectionPoint>Address,omitempty"`
	UniProtocolConnectionPoint *networkAddress `xml:"UniProtocolExternalConnectionPoint>Address,omitempty"`
	// Адрес сервера сообщений для посылки данных, если не указан, то сервер сообщений не будет использоваться
	MessageServerConnectionPoint *networkAddress `xml:"MessageServerConnectionPoint>Address,omitempty"`
	// Адрес сервера сообщений для посылки данных, если не указан, то сервер сообщений не будет использоваться
	ArchiveServerConnectionPoint *networkAddress `xml:"ArchiveServerConnectionPoint>Address,omitempty"`
	// Точки приема датаграмм. По крайней мере одна должна быть задана
	InboundPoints []networkAddress `xml:"Inbound>Address"`
	// Адрес веб сервера для диагностики, если не указан, то веб сервер
	WebServerConnectionPoint *networkAddress `xml:"WebServerConnectionPoint>Address,omitempty"`
	Id                       int             `xml:"Number,attr"`
}

type ObjectParameterMapping struct {
	XMLName  xml.Name `xml:"ObjectParameterMapping"`
	Id       int      `xml:"ParameterId,attr"`
	ObjectId int      `xml:"ObjectId,attr"`
	DeviceId int      `xml:"DeviceId,attr"`
	SensorId int      `xml:"SensorId,attr"`
}

type ObjectAttributeMapping struct {
	XMLName  xml.Name `xml:"ObjectAttributeMapping"`
	Id       int      `xml:"AttributeTypeId,attr"`
	ObjectId int      `xml:"ObjectId,attr"`
	DeviceId int      `xml:"DeviceId,attr"`
	SensorId int      `xml:"SensorId,attr"`
}

type UniParameterInfo struct {
	XMLName        xml.Name `xml:"UniParameter"`
	Id             int      `xml:"ControlParamId,attr"`
	UniParamTypeId int      `xml:"UniParamTypeId,attr"`
	ParameterId    int      `xml:"MeasureTypeId,attr"`
	MinAttrId      *int     `xml:"NormMinAttributeId,attr"`
	MaxAttrId      *int     `xml:"NormMaxAttributeId,attr"`
	ExtAttrId      *int     `xml:"NormExtendedAttributeId,attr"`
}

type uniDiagStateInfo struct {
	XMLName        xml.Name `xml:"UniDiagState"`
	UniDiagStateId int      `xml:"UniDiagStateId,attr"`
	FailureId      int      `xml:"FailureId,attr"`
}

type UniObjectInfo struct {
	XMLName       xml.Name           `xml:"UniObject"`
	UniParameters []UniParameterInfo `xml:"UniParameters>UniParameter"`
	UniDiagStates []uniDiagStateInfo `xml:"UniDiagStates>UniDiagState"`
	ObjectId      int                `xml:"ObjectId,attr"`
	UniTypeId     int                `xml:"UniTypeId,attr"`
	HasUniStates  bool               `xml:"HasUniStates,attr"`
}

type UniPlaceInfo struct {
	XMLName    xml.Name        `xml:"UniPlace"`
	UniObjects []UniObjectInfo `xml:"UniObjects>UniObject"`
	Id         int             `xml:"UniPlaceId,attr"`
}

type UniStateMappingInfo struct {
	XMLName       xml.Name `xml:"UniState"`
	UniTypeId     int      `xml:"UniTypeId,attr"`
	UniStateId    int      `xml:"UniStateId,attr"`
	ObjectStateId int      `xml:"ObjectStateId,attr"`
}

type objectIdInfo struct {
	XMLName  xml.Name `xml:"Object"`
	ObjectId int      `xml:"Id,attr"`
}

type objectsOnHostInfo struct {
	XMLName   xml.Name       `xml:"Host"`
	HostId    int            `xml:"HostId,attr"`
	ObjectIds []objectIdInfo `xml:"Object"`
}

type nwaState struct {
	XMLName xml.Name `xml:"FlowState"`
	Id      int      `xml:"StateId,attr"`
	EventId *int     `xml:"EventId,attr"`
}

type normalWorkflowAlgorithm struct {
	XMLName    xml.Name   `xml:"ObjectNormalFlow"`
	FlowStates []nwaState `xml:"FlowState"`
}

type nestedProject struct {
	XMLName xml.Name `xml:"Project"`
	// Объекты контроля
	Objects []ObjectInfo `xml:"Objects>Object"`
	// Измерения
	Parameters      []ObjectParameter      `xml:"ObjectParameterTypes>ObjectParameter"`
	Attributes      []ObjectAttribute      `xml:"ObjectAttributeTypes>ObjectAttributeType"`
	AttributeValues []ObjectAttributeValue `xml:"ObjectAttributes>ObjectAttribute"`
	// Привязки измерений к устройствам
	ParameterMappings []ObjectParameterMapping `xml:"ObjectParameterMappings>ObjectParameterMapping"`
	AttributeMappings []ObjectAttributeMapping `xml:"ObjectAttributeMappings>ObjectAttributeMapping"`
	UniPlaces         []UniPlaceInfo           `xml:"UniProtocol>UniPlaces>UniPlace"`
	UniStates         []UniStateMappingInfo    `xml:"UniProtocol>UniStates>UniState"`
	// Хосты
	Hosts []NestedHost `xml:"Hosts>Host"`
	// Алгоритмы нормальной работы
	NormalWorkFlows []normalWorkflowAlgorithm `xml:"ObjectNormalFlows>ObjectNormalFlow"`
	// Справочная информация - станции, перегоны, типы объектов
	Stations      []stationInfo       `xml:"Stations>Station"`
	Spans         []spanInfo          `xml:"Spans>Span"`
	ObjectTypes   []objectTypeInfo    `xml:"ObjectTypes>ObjectType"`
	ObjectsOnHost []objectsOnHostInfo `xml:"ObjectsToHosts>Host"`

	objectsMap                 map[int]*ObjectInfo
	objectParametersMap        map[int]*ObjectParameter
	objectAttributesMap        map[int]*ObjectAttribute
	objectParametersMappingMap map[ParameterMappingKey]*ObjectParameterMapping
	objectAttributesMappingMap map[ParameterMappingKey]*ObjectAttributeMapping
	objectAttributeValuesMap   map[attributeValueKey]string
	stationMap                 map[int]*stationInfo
	objectTypeMap              map[int]*objectTypeInfo
	spanMap                    map[int]*spanInfo
	deviceMap                  map[int]*Device
	objectsToHosts             map[int]int
	nwaStateToEventMap         map[int]int
}
