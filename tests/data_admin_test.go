package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestAdminHandleGetNumericData(t *testing.T) {
    numericData := []adminEntity.NumericData{
        {
            Id:         "id1",
            Data_id:    "data_id1",
            Device_id:  "device_id1",
            Value:      123.456,
            Created_at: "created_at1",
        },
        {
            Id:         "id2",
            Data_id:    "data_id2",
            Device_id:  "device_id2",
            Value:      789.012,
            Created_at: "created_at2",
        },
        {
            Id:         "id3",
            Data_id:    "data_id3",
            Device_id:  "device_id",
            Value:      789.014,
            Created_at: "created_at3",
        },
    }

    pretty := false
    admin.HandleGetNumericData(&numericData, &pretty)
}

func TestAdminHandleGetStringData(t *testing.T) {
    stringData := []adminEntity.StringData{
        {
            Id:         "id1",
            Data_id:    "data_id1",
            Device_id:  "device_id1",
            Value:      "some_value",
            Created_at: "created_at1",
        },
        {
            Id:         "id2",
            Data_id:    "data_id2",
            Device_id:  "device_id2",
            Value:      "another_value",
            Created_at: "created_at2",
        },
        {
            Id:         "id3",
            Data_id:    "data_id3",
            Device_id:  "device_id3",
            Value:      "another_value",
            Created_at: "created_at3",
        },
    }

    pretty := false
    admin.HandleGetStringData(&stringData, &pretty)
}
