package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleListInstancesTypes(t *testing.T){
    mockInstancesTypes := &client.InstancesTypes{
        Types: []string{"type-1","type-2","type-3"},
    }
    pretty := true
    user.HandleListInstancesTypes(mockInstancesTypes, &pretty)
}

func TestHandleGetInstances(t *testing.T){
    mockInstances := []client.Instance{
        {
            Id:  1,
            Name: "instance-1",
        },
        {
            Id:  2,
            Name: "instance-2",
        },
        {
            Id:  3,
            Name: "instance-3",
        },
    }
    pretty := true
    user.HandleGetInstances(&mockInstances, &pretty)
}

func TestHandleGetInstance(t *testing.T){
    mockInstance := &client.Instance{
        Id:  1,
        Name: "instance-1",
    }
    pretty := true
    user.HandleGetInstance(mockInstance, &pretty)
}
