package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleAddObjectType(t *testing.T){
	mockObjectType := &client.ObjectType{
		Id:      "1",
		Content: client.ObjectTypeContent{
			Name:             "MockObjectName",
			Public:           false,
			DecodingFunction: "mockDecodingFunction",
			Triggers:         []string{"trigger1", "trigger2"},
		},
	}

	pretyy := true
	user.HandleAddObjectType(mockObjectType, &pretyy)
}

func TestHandleGetObjectTypes(t *testing.T) {
	mockObjectTypes := []client.ObjectType{
		{
			Id:      "1",
			Content: client.ObjectTypeContent{
				Name: "MockObjectName1", 
				Public: false, 
				DecodingFunction: "mockDecodingFunction1", 
				Triggers: []string{"trigger1", "trigger2"}},
		},
		{
			Id:      "2",
			Content: client.ObjectTypeContent{
				Name: "MockObjectName2", 
				Public: false, 
				DecodingFunction: "mockDecodingFunction2", 
				Triggers: []string{"trigger1", "trigger2"}},
		},
		{
			Id:      "3",
			Content: client.ObjectTypeContent{
				Name: "MockObjectName3", 
				Public: false,
				DecodingFunction: "mockDecodingFunction3", 
				Triggers: []string{"trigger1", "trigger2"}},
		},
	}

	var pretty bool = true
	user.HandleGetObjectTypes(&mockObjectTypes, &pretty)
}

func TestHandleGetObjectType(t *testing.T) {
    mockObjectType := &client.ObjectType{
        Id: "1",
        Content: client.ObjectTypeContent{
            Name:       "Test Object Type",
            Public:     true,
            DecodingFunction: "decodeFunc_for_test",
        },
    }

    pretty := false
    user.HandleGetObjectType(mockObjectType, &pretty)
}
