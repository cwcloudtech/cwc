package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleGetDeployment(t *testing.T) {
    mockDeployment := client.DeploymentByIdResponse{
        Name:        "test-deployment-name",
        Namespace:   "test-namespace",
        Pods:       []client.Pod{},
        Containers:  []client.Container{},
        Project:     client.DeploymentProject{},
        Environment: client.DeploymentEnvironment{},
    }

    pretty := false
    user.HandleGetDeployment(&mockDeployment, &pretty)
}

func TestHandleGetDeployments(t *testing.T) {
    mockDeployment := []client.Deployment{
        {
            Id:          "test-deployment-id",
            Name:        "test-deployment-name",
            Description: "test-deployment-description",
            Hash:        "test-deployment-hash",
            Cluster_id:  1,
            Project_id:  2,
            Env_id:      3,
            User_id:     4,
            Created_at:  "test-created-at",
            Namespace:   "test-namespace",
        },
    }

    pretty := false
    user.HandleGetDeployments(&mockDeployment, &pretty)
}
