package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidProjectName(t *testing.T) {
	tests := []struct {
		name        string
		projectName string
		isValid     bool
	}{
		{"NotValidNoVersion", "project_10.zip", false},
		{"NotValidNoZipExtension", "project_10_2.tar", false},
		{"NotProjectPrefix", "proj_10_2.zip", false},
		{"Valid", "project_10_2.zip", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isValid := isValidProjectName(test.projectName)
			assert.Equal(t, test.isValid, isValid)
		})
	}
}

func TestGetProjectAndVersionIds(t *testing.T) {
	tests := []struct {
		name        string
		projectName string
		id          int
		versionId   int
		isValid     bool
	}{
		{"NotValidNoVersion", "project_10.zip", 0, 0, false},
		{"NotValidNoZipExtension", "project_10_2.tar", 0, 0, false},
		{"NotProjectPrefix", "proj_10_2.zip", 0, 0, false},
		{"NotProjectId", "project_xx_2.zip", 0, 0, false},
		{"ProjectIdNegative", "project_-10_2.zip", 0, 0, false},
		{"NotVersionId", "project_10_xx.zip", 0, 0, false},
		{"VersionIdNegative", "project_10_-2.zip", 0, 0, false},
		{"Valid1", "project_10_2.zip", 10, 2, true},
		{"Valid2", "project_17_1025.zip", 17, 1025, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, ver, err := getProjectAndVersionIds(test.projectName)

			if !test.isValid {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, test.id, id)
				assert.Equal(t, test.versionId, ver)
			}
		})
	}
}
