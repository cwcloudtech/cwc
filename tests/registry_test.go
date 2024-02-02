package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleGetRegistries(t *testing.T) {
    mockRegistries := []client.Registry{
        {
            Id:        1,
            Name:      "registry-1",
            Status:    "status-1",
            CreatedAt: "created_at-1",
            AccessKey: "access_key-1",
            Endpoint:  "endpoint-1",
            SecretKey: "secret_key-1",
            Region:    "region-1",
            Type:      "type-1",
        },
        {
            Id:        2,
            Name:      "registry-2",
            Status:    "status-2",
            CreatedAt: "created_at-2",
            AccessKey: "access_key-2",
            Endpoint:  "endpoint-2",
            SecretKey: "secret_key-2",
            Region:    "region-2",
            Type:      "type-2",
        },
        {
            Id:        3,
            Name:      "registry-3",
            Status:    "status-3",
            CreatedAt: "created_at-3",
            AccessKey: "access_key-3",
            Endpoint:  "endpoint-3",
            SecretKey: "secret_key-3",
            Region:    "region-3",
            Type:      "type-3",
        },
    }
    pretty := true
    user.HandleGetRegistries(&mockRegistries, &pretty)
}

func TestHandleGetRegistry(t *testing.T) {
    mockRegistry := &client.Registry{
        Id:      1,
        Name:    "registry-1",
        Status:  "status-1",
        CreatedAt: "created_at-1",
        AccessKey: "access_key-1",
        Endpoint: "endpoint-1",
        SecretKey: "secret_key-1",
        Region:  "region-1",
        Type:    "type-1",
    }
    pretty := true
    user.HandleGetRegistry(mockRegistry, &pretty)
}
