package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleGetEnvironments(t *testing.T) {
    mockEnvironments := []client.Environment{
        {
            Id:         1,
            Name:       "environment-1",
            Path:       "path-1",
            Description: "description-1",
        },
        {
            Id:         2,
            Name:       "environment-2",
            Path:       "path-2",
            Description: "description-2",
        },
        {
            Id:         3,
            Name:       "environment-3",
            Path:       "path-3",
            Description: "description-3",
        },
    }
    pretty := true
    user.HandleGetEnvironments(&mockEnvironments , &pretty)
}

func TestHandleGetEnvironment(t *testing.T) {
    mockEnvironment := &client.Environment{
        Id:         1,
        Name:       "environment-1",
        Path:       "path-1",
        Description: "description-1",
    }
    pretty := true
    user.HandleGetEnvironment(mockEnvironment , &pretty)
}
