package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateProject(t *testing.T) {

	var testProject1 = multiProject{
		Hosts: []hostInfo{{Id: 1}, {Id: 2}},
	}
	var testProject2 = multiProject{
		Hosts: []hostInfo{{Id: 1}},
	}

	var testProject3 = multiProject{
		Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
			{Ip: "127.0.0.1", Port: 10000},
		}}},
	}
	tests := []struct {
		name         string
		project      *multiProject
		errorMessage string
	}{
		{"NotSingleHost", &testProject1, ErrorMsgNoSingleHost},
		{"NoInboundPoint", &testProject2, ErrorMsgNoInboundEndpoint},
		{name: "Ok", project: &testProject3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.project.validate()

			if err != nil {
				assert.Equal(t, test.errorMessage, err.Error())
			} else {
				assert.Empty(t, test.errorMessage)
			}
		})
	}
}

func TestGetUniServerPoint(t *testing.T) {

	var testProject1 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}}},
			},
		},
	}
	var testProject3 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}}, UniProtocolConnectionPoint: &networkAddress{Ip: "127.0.0.1", Port: 20000},
			}}}}

	tests := []struct {
		name    string
		project *projectInfo
		found   bool
		ip      string
		port    int
	}{
		{name: "No uni server", project: &testProject1, found: false},
		{name: "Ok", project: &testProject3, found: true, ip: "127.0.0.1", port: 20000},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			extPoint := test.project.GetUniProtocolConnectionPoint()

			if !test.found {
				assert.Nil(t, extPoint)
			} else {
				assert.Equal(t, test.ip, extPoint.GetHostName())
				assert.Equal(t, test.port, extPoint.GetPort())
			}
		})
	}
}

func TestUdpConnectionPoint(t *testing.T) {

	var testProject1 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1}},
		},
	}
	var testProject3 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}},
			}}}}

	tests := []struct {
		name    string
		project *projectInfo
		found   bool
		ip      string
		port    int
	}{
		{name: "No udp point", project: &testProject1, found: false},
		{name: "Ok", project: &testProject3, found: true, ip: "127.0.0.1", port: 10000},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			extPoint := test.project.GetUdpConnectionPoint()

			if !test.found {
				assert.Nil(t, extPoint)
			} else {
				assert.Equal(t, test.ip, extPoint.GetHostName())
				assert.Equal(t, test.port, extPoint.GetPort())
			}
		})
	}
}

func TestGetExternalPoint(t *testing.T) {

	var testProject1 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}}},
			},
		},
	}
	var testProject3 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}}, ConnectionPoint: &networkAddress{Ip: "127.0.0.1", Port: 8000},
			}}}}

	tests := []struct {
		name    string
		project *projectInfo
		found   bool
		ip      string
		port    int
	}{
		{name: "No External point", project: &testProject1, found: false},
		{name: "Ok", project: &testProject3, found: true, ip: "127.0.0.1", port: 8000},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			extPoint := test.project.GetExternalSystemConnectionPoint()

			if !test.found {
				assert.Nil(t, extPoint)
			} else {
				assert.Equal(t, test.ip, extPoint.GetHostName())
				assert.Equal(t, test.port, extPoint.GetPort())
			}
		})
	}
}

func TestGetMessageServerPoint(t *testing.T) {

	var testProject1 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}}},
			},
		},
	}
	var testProject3 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}}, MessageServerConnectionPoint: &networkAddress{Ip: "127.0.0.1", Port: 9000},
			}}}}

	tests := []struct {
		name    string
		project *projectInfo
		found   bool
		ip      string
		port    int
	}{
		{name: "No message server", project: &testProject1, found: false},
		{name: "Ok", project: &testProject3, found: true, ip: "127.0.0.1", port: 9000},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			extPoint := test.project.GetMessageServerConnectionPoint()

			if !test.found {
				assert.Nil(t, extPoint)
			} else {
				assert.Equal(t, test.ip, extPoint.GetHostName())
				assert.Equal(t, test.port, extPoint.GetPort())
			}
		})
	}
}

func TestGetArchiveServerPoint(t *testing.T) {

	var testProject1 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}}},
			},
		},
	}
	var testProject3 = projectInfo{
		projectInfo: &multiProject{
			Hosts: []hostInfo{{Id: 1, InboundPoints: []networkAddress{
				{Ip: "127.0.0.1", Port: 10000}}, ArchiveServerConnectionPoint: &networkAddress{Ip: "127.0.0.1", Port: 9200},
			}}}}

	tests := []struct {
		name    string
		project *projectInfo
		found   bool
		ip      string
		port    int
	}{
		{name: "No archive server", project: &testProject1, found: false},
		{name: "Ok", project: &testProject3, found: true, ip: "127.0.0.1", port: 9200},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			extPoint := test.project.GetArchiveServerConnectionPoint()

			if !test.found {
				assert.Nil(t, extPoint)
			} else {
				assert.Equal(t, test.ip, extPoint.GetHostName())
				assert.Equal(t, test.port, extPoint.GetPort())
			}
		})
	}
}

func TestMeasureInfo(t *testing.T) {

	var testMeasure1 = ObjectParameter{
		Id:            1,
		Name:          "name1",
		UnitOfMeasure: "Вольты, В",
	}

	var testMeasure2 = ObjectParameter{
		Id:            1,
		Name:          "name2",
		ShortName:     "SN",
		UnitOfMeasure: "Безразмерная",
	}

	var testMeasure3 = ObjectParameter{
		Id:   1,
		Name: "some name",
	}

	tests := []struct {
		name                 string
		measure              *ObjectParameter
		displayName          string
		displayUnitOfMeasure string
	}{
		{"NoShortName", &testMeasure1, testMeasure1.Name, "В"},
		{"WithShortName", &testMeasure2, testMeasure2.ShortName, ""},
		{"NoUnitOfMeasure", &testMeasure3, testMeasure3.Name, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			displayName := test.measure.GetParameterDisplayName()
			displayUnitOfMeasure := test.measure.GetUnitOfMeasureDisplayName()

			assert.Equal(t, test.displayName, displayName)
			assert.Equal(t, test.displayUnitOfMeasure, displayUnitOfMeasure)
		})
	}
}

func TestGetSpanName(t *testing.T) {

	var testProject1 = projectInfo{nestedProjects: make(map[string]*nestedProject)}

	var testProject2 = projectInfo{nestedProjects: map[string]*nestedProject{"project_1_1.zip": {
		Spans: []spanInfo{{Id: 1, Name: "test"}, {Id: 2, Name: "test_1"}}},
	}}

	var testProject3 = projectInfo{nestedProjects: map[string]*nestedProject{"project_1_1.zip": {
		Spans: []spanInfo{{Id: 3, Name: "test"}, {Id: 2, Name: "test_1"}}},
	}}

	tests := []struct {
		name    string
		project *projectInfo
		result  string
		id      int
	}{
		{"NoSpans", &testProject1, "", 1},
		{"SpanFound", &testProject2, "test", 1},
		{"SpanNotFound", &testProject3, "", 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			name := test.project.getSpanName(test.id)

			assert.Equal(t, test.result, name)
		})
	}
}

func TestGetStationName(t *testing.T) {

	var testProject1 = projectInfo{}
	var testProject2 = projectInfo{nestedProjects: map[string]*nestedProject{
		"project_1_1.zip": {Stations: []stationInfo{{Id: 1, Name: "test"}, {Id: 2, Name: "test_1"}}},
	}}
	var testProject3 = projectInfo{nestedProjects: map[string]*nestedProject{
		"project_1_1.zip": {Stations: []stationInfo{{Id: 3, Name: "test"}, {Id: 2, Name: "test_1"}}},
	}}

	tests := []struct {
		name    string
		project *projectInfo
		result  string
		id      int
	}{
		{"NoStations", &testProject1, "", 1},
		{"StationFound", &testProject2, "test", 1},
		{"StationNotFound", &testProject3, "", 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			name := test.project.getStationName(test.id)

			assert.Equal(t, test.result, name)
		})
	}
}

func TestGetDeviceInfo(t *testing.T) {

	var testProject1 = projectInfo{}
	var testProject2 = projectInfo{nestedProjects: map[string]*nestedProject{
		"project_1_1.zip": {
			Stations: []stationInfo{{Id: 3, Name: "test"}, {Id: 2, Name: "test_1"}},
			Hosts:    []NestedHost{{Id: 1000, Devices: []Device{{Id: 1000, BitsPerSensor: 32, SensorCount: 100}}}},
		}}}
	var testProject3 = projectInfo{nestedProjects: map[string]*nestedProject{
		"project_1_1.zip": {
			Stations: []stationInfo{{Id: 3, Name: "test"}, {Id: 2, Name: "test_1"}},
			Hosts: []NestedHost{
				{Id: 1000, Devices: []Device{{Id: 1000, BitsPerSensor: 32, SensorCount: 100}}},
				{Id: 2000, Devices: []Device{{Id: 1001, BitsPerSensor: 2, SensorCount: 512}}}},
		}}}

	tests := []struct {
		name          string
		project       *projectInfo
		found         bool
		id            int
		sensorCount   int
		bitsPerSensor int
	}{
		{name: "NotFound", project: &testProject1, found: false},
		{name: "Single", project: &testProject2, found: true, id: 1000, sensorCount: 100, bitsPerSensor: 32},
		{name: "Multiple", project: &testProject3, found: true, id: 1001, sensorCount: 512, bitsPerSensor: 2},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			devInfo := test.project.GetDeviceInfo(test.id)

			if !test.found {
				assert.Nil(t, devInfo)
			} else {
				assert.Equal(t, test.id, devInfo.Id)
				assert.Equal(t, test.sensorCount, devInfo.SensorCount)
			}
			// BitPerSensor, ranges
		})
	}
}

func TestGetObjectTypeName(t *testing.T) {

	var testProject1 = projectInfo{}
	var testProject2 = projectInfo{nestedProjects: map[string]*nestedProject{
		"project_1_1.zip": {ObjectTypes: []objectTypeInfo{{Id: 1, Name: "test"}, {Id: 2, Name: "test_1"}}},
	}}
	var testProject3 = projectInfo{nestedProjects: map[string]*nestedProject{
		"project_1_1.zip": {ObjectTypes: []objectTypeInfo{{Id: 3, Name: "test"}, {Id: 2, Name: "test_1"}}},
	}}

	tests := []struct {
		name    string
		project *projectInfo
		result  string
		id      int
	}{
		{"NoTypes", &testProject1, "", 1},
		{"TypeFound", &testProject2, "test", 1},
		{"TypeNotFound", &testProject3, "", 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			name := test.project.getObjectTypeName(test.id)

			assert.Equal(t, test.result, name)
		})
	}
}

func TestGetObjectInfo(t *testing.T) {

	var testProject1 = projectInfo{}

	testObject2_1 := ObjectInfo{Id: 1, Name: "test", TypeId: 1, StationId: 100}
	testObject2_2 := ObjectInfo{Id: 2, Name: "test_1", TypeId: 1, StationId: 200}

	var testProject2 = projectInfo{
		nestedProjects: map[string]*nestedProject{"project_1_1.zip": {
			Objects: []ObjectInfo{testObject2_1, testObject2_2},
		}},
	}
	testObject3_1 := ObjectInfo{Id: 3, Name: "test", TypeId: 1, StationId: 100}
	testObject3_2 := ObjectInfo{Id: 2, Name: "test_1", TypeId: 2, StationId: 200}
	var testProject3 = projectInfo{
		nestedProjects: map[string]*nestedProject{"project_1_1.zip": {
			Objects: []ObjectInfo{testObject3_1, testObject3_2},
		}},
	}

	tests := []struct {
		name    string
		project *projectInfo
		result  *ObjectInfo
		id      int
	}{
		{"NoObjects", &testProject1, nil, 1},
		{"ObjectFound", &testProject2, &testObject2_1, 1},
		{"ObjectNotFound", &testProject3, nil, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			obj := test.project.GetObjectInfo(test.id)

			assert.Equal(t, test.result, obj)
		})
	}
}

func TestGetObjectParameterInfo(t *testing.T) {

	var testProject1 = projectInfo{nestedProjects: make(map[string]*nestedProject)}

	var np2 = nestedProject{
		Parameters: []ObjectParameter{
			{Id: 1, Name: "test", UnitOfMeasure: "Амперы,А"},
			{Id: 2, Name: "test_1", UnitOfMeasure: "Вольты,В"}},
	}

	testProject2 := projectInfo{nestedProjects: map[string]*nestedProject{"project_1_1.zip": &np2}}
	var np3 = nestedProject{
		Parameters: []ObjectParameter{
			{Id: 3, Name: "test", UnitOfMeasure: "Амперы,А"},
			{Id: 2, Name: "test_1", UnitOfMeasure: "Вольты,В"}},
	}

	testProject3 := projectInfo{nestedProjects: map[string]*nestedProject{"project_1_1.zip": &np3}}

	tests := []struct {
		name    string
		project *projectInfo
		result  *ObjectParameter
		id      int
	}{
		{"NoMeasure", &testProject1, nil, 1},
		{"MeasureFound", &testProject2, &np2.Parameters[0], 1},
		{"MeasureNotFound", &testProject3, nil, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			obj := test.project.GetObjectParameterInfo(test.id)

			assert.Equal(t, test.result, obj)
		})
	}
}

func TestGetHostObjects(t *testing.T) {

	var np1 = nestedProject{
		ObjectsOnHost: []objectsOnHostInfo{
			{HostId: 100, ObjectIds: []objectIdInfo{{ObjectId: 1000}, {ObjectId: 2000}}},
			{HostId: 200, ObjectIds: []objectIdInfo{{ObjectId: 1001}, {ObjectId: 1002}}},
		},
	}

	var np2 = nestedProject{
		ObjectsOnHost: []objectsOnHostInfo{
			{HostId: 700, ObjectIds: []objectIdInfo{{ObjectId: 7000}, {ObjectId: 7001}}},
			{HostId: 800, ObjectIds: []objectIdInfo{{ObjectId: 8000}, {ObjectId: 8001}}},
		},
	}

	testProject := projectInfo{nestedProjects: map[string]*nestedProject{
		"project_1_1.zip": &np1,
		"project_2_1.zip": &np2}}

	tests := []struct {
		name     string
		hostId   int
		objectId int
	}{
		{"test1", 100, 1000},
		{"test2", 200, 1002},
		{"test3", 0, 10000},
		{"test4", 700, 7001},
		{"test5", 800, 8000},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hostId := testProject.GetObjectHost(test.objectId)

			assert.Equal(t, test.hostId, hostId)
		})
	}
}

func TestGetAttribute(t *testing.T) {

	var np1 = nestedProject{
		Attributes: []ObjectAttribute{
			{Id: 200, Name: "attr1", UnitOfMeasure: "V"},
			{Id: 500, Name: "attr2", UnitOfMeasure: "A"},
		},
	}

	var np2 = nestedProject{
		Attributes: []ObjectAttribute{
			{Id: 300, Name: "attr3", UnitOfMeasure: "V"},
			{Id: 500, Name: "attr2", UnitOfMeasure: "A"},
		},
	}

	testProject := projectInfo{nestedProjects: map[string]*nestedProject{
		"project_1_1.zip": &np1,
		"project_2_1.zip": &np2}}

	tests := []struct {
		name          string
		attrId        int
		attrName      string
		unitOfMeasure string
		result        bool
	}{
		{"test1", 200, "attr1", "V", true},
		{"test2", 500, "attr2", "A", true},
		{"test3", 300, "attr3", "V", true},
		{name: "test4", attrId: 700, result: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			aInfo := testProject.GetAttributeInfo(test.attrId)
			if !test.result {
				assert.Nil(t, aInfo)
			} else {
				assert.Equal(t, test.unitOfMeasure, aInfo.GetUnitOfMeasure())
				assert.Equal(t, test.attrName, aInfo.GetName())
				assert.Equal(t, test.attrId, aInfo.GetId())
			}
		})
	}
}
