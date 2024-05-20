package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestAdminHandleAddObjectType(t *testing.T){
	mockcreated_objectType := &adminEntity.ObjectType{
		Id:      "1",
		Content: adminEntity.ObjectTypeContent{
			Name:             "MockObjectName",
			Public:           false,
			DecodingFunction: "mockDecodingFunction",
			Triggers:         []string{"trigger1", "trigger2"},
		},
	}

	pretyy := true
	admin.HandleAddObjectType(mockcreated_objectType, &pretyy)
}

func TestAdminHandleGetObjectTypes(t *testing.T) {
	mockObjectTypes := []adminEntity.ObjectType{
		{
			Id:      "1",
			Content: adminEntity.ObjectTypeContent{
				Name: "MockObjectName1", 
				Public: false, 
				DecodingFunction: "mockDecodingFunction1", 
				Triggers: []string{"trigger1", "trigger2"}},
		},
		{
			Id:      "2",
			Content: adminEntity.ObjectTypeContent{
				Name: "MockObjectName2", 
				Public: false, 
				DecodingFunction: "mockDecodingFunction2", 
				Triggers: []string{"trigger1", "trigger2"}},
		},
		{
			Id:      "3",
			Content: adminEntity.ObjectTypeContent{
				Name: "MockObjectName3", 
				Public: false, 
				DecodingFunction: "mockDecodingFunction3", 
				Triggers: []string{"trigger1", "trigger2"}},
		},
	}

	var pretty bool = true
	admin.HandleGetObjectTypes(&mockObjectTypes, &pretty)
}

func TestAdminHandleGetObjectType(t *testing.T) {
    mockObjectType := &adminEntity.ObjectType{
        Id: "1",
        Content: adminEntity.ObjectTypeContent{
            Name:       "Test Object Type",
            Public:     true,
            DecodingFunction: "decodeFunc_for_test",
        },
    }

    pretty := false
    admin.HandleGetObjectType(mockObjectType, &pretty)
}
