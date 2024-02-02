package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestAdminGetEnvironments(t *testing.T) {
	mockEnvironments := []adminEntity.Environment {
		{
			Id:            1,
			Name:          "environment-1",
			Path:          "path-1",
			Roles:         "roles-1",
			IsPrivate:     false,
			Description:   "description-1",
			EnvironmentTemplate: "template-1",
			DocTemplate:   "doc_template-1",
			SubDomains:    "subdomains-1",
		},
		{
			Id:            2,
			Name:          "environment-2",
			Path:          "path-2",
			Roles:         "roles-2",
			IsPrivate:     false,
			Description:   "description-2",
			EnvironmentTemplate: "template-2",
			DocTemplate:   "doc_template-2",
			SubDomains:    "subdomains-2",
		},
		{
			Id:            3,
			Name:          "environment-3",
			Path:          "path-3",
			Roles:         "roles-3",
			IsPrivate:     false,
			Description:   "description-3",
			EnvironmentTemplate: "template-3",
			DocTemplate:   "doc_template-3",
			SubDomains:    "subdomains-3",
		},
	}
	pretty := true
	admin.HandleGetEnvironments(&mockEnvironments, &pretty)
}
