package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleAddData(t *testing.T){
	created_data := &client.Data{
		Id:            "test-id",
		Device_id: "test-typeobject-id",
        Content:   "content",
	}
    
	user.HandleAddData(created_data)
}
