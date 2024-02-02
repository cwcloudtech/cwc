package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestAdminHandleGetRegistry(t *testing.T) {
	mockRegistry := &adminEntity.Registry{
		Id:        1,
		Name:      "registry-cwc",
		Status:    "active",
		CreatedAt: "2024-01-01T00:00:00Z",
		AccessKey: "accessKeycwc",
		Endpoint:  "https://registry.example.com",
		SecretKey: "secretKeycwc",
		Region:    "void",
		Type:      "type-cwc",
		Email:     "cwc_test@comwork.io",
	}
	pretty := true
	admin.HandleGetRegistry(mockRegistry, &pretty)
}

func TestGetRegistries(t *testing.T){
	mockRegistries := []adminEntity.Registry {
		{
			Id:        1,
			Name:      "registry-1",
			Status:    "active",
			CreatedAt: "2024-01-01T00:00:00Z",
			AccessKey: "accessKey123",
			Endpoint:  "https://registry.example.com",
			SecretKey: "secretKey123",
			Region:    "void",
			Type:      "type-cwc",
			Email:     "cwc_test@comwork.io",
		},
		{
			Id:        2,
			Name:      "registry-2",
			Status:    "active",
			CreatedAt: "2024-01-01T00:00:00Z",
			AccessKey: "accessKey123",
			Endpoint:  "https://registry.example.com",
			SecretKey: "secretKey123",
			Region:    "void",
			Type:      "type-cwc",
			Email:     "cwc_test@comwork.io",
		},
		{
			Id:        3,
			Name:      "registry-3",
			Status:    "active",
			CreatedAt: "2021-01-01T00:00:00Z",
			AccessKey: "accessKey123",
			Endpoint:  "https://registry.example.com",
			SecretKey: "secretKey123",
			Region:    "void",
			Type:      "type-cwc",
			Email:     "cwc_test@comwork.io",
		},
	}
	pretty := true
	admin.HandleGetRegistries(&mockRegistries, &pretty)
}
