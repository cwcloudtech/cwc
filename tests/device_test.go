package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleAddDevice(t *testing.T){
	createdDevice := &client.Device{
		Id:            "test-id",
		Username:      "test-user",
		Typeobject_id: "test-typeobject-id",
		Active:        true,
	}
	
	pretyy := true
	user.HandleAddDevice(createdDevice, &pretyy)
}

func TestHandleGetDevices(t *testing.T) {
    mockDevices := []client.Device{
        {
            Id:            "device-1",
            Username:      "user-1",
            Typeobject_id: "typeobject-1",
            Active:        true,
        },
        {
            Id:            "device-2",
            Username:      "user-2",
            Typeobject_id: "typeobject-2",
            Active:        false,
        },
		{
            Id:            "device-3",
            Username:      "user-3",
            Typeobject_id: "typeobject-3",
            Active:        false,
        },
		{
            Id:            "device-4",
            Username:      "user-4",
            Typeobject_id: "typeobject-4",
            Active:        false,
        },
    }

    var pretty bool = true
    user.HandleGetDevices(&mockDevices, &pretty)
}

func TestHandleGetDevice(t *testing.T) {
    mockDevice := &client.Device{
        Id:            "device-1",
        Username:      "user-1",
        Typeobject_id: "typeobject-1",
        Active:        true,
    }

    pretty := false
    user.HandleGetDevice(mockDevice, &pretty)
}
