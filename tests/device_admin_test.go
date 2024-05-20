package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestAdminHandleGetDevices(t *testing.T) {
    mockdevices := []adminEntity.Device{
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
    admin.HandleGetDevices(&mockdevices, &pretty)
}
