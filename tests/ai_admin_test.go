package test 

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestAdminHandleGetFunctions(t *testing.T) {
	mockFunctions := []adminEntity.Function{
		{
			Id:        "1",
			Owner_id:  1,
			Is_public: true,
			Content: adminEntity.FunctionContent{
				Name:     "function-1",
				Language: "language-1",
			},
			Created_at: "22/01/2024",
			Updated_at: "24/01/2024",
		},
	}
	pretty := true
	admin.HandleGetFunctions(&mockFunctions, &pretty)
}

func TestHandleGetFunctionOwner(t *testing.T) {
	mockFunctionOwner := adminEntity.FunctionOwner{
		Id:       1,
		Username: "unit_test_cwc_user_1",
	}
	pretty := true
	admin.HandleGetFunctionOwner(&mockFunctionOwner, &pretty)	
}

func TestHandleGetAllInvocations(t *testing.T) {
	invocations := []adminEntity.Invocation{
		{
			Id:         "1",
			Invoker_id: 1,
			Content: adminEntity.InvocationContent{
				Function_id: "1",
				State:       "running",
			},
			Created_at: "2024-01-01T00:00:00Z",
			Updated_at: "2024-01-01T00:00:00Z",
		},
		{
			Id:         "2",
			Invoker_id: 2,
			Content: adminEntity.InvocationContent{
				Function_id: "2",
				State:       "completed",
			},
			Created_at: "2024-01-02T00:00:00Z",
			Updated_at: "2024-01-02T00:00:00Z",
		},
		{
			Id:         "3",
			Invoker_id: 3,
			Content: adminEntity.InvocationContent{
				Function_id: "3",
				State:       "failed",
			},
			Created_at: "2024-01-03T00:00:00Z",
			Updated_at: "2024-01-03T00:00:00Z",
		},
	}
	pretty := true
	admin.HandleGetInvocations(&invocations, &pretty)
}

func TestGetInvocationInvoker(t *testing.T) {
    mockinvocation := adminEntity.InvocationInvoker{
		Id:         1,
		Username: "admin",
    }
    pretty := true
    admin.HandleGetInvocationInvoker(&mockinvocation, &pretty)
}

func TestAdminHandleGetTriggers(t *testing.T){
	mockTrigger := []adminEntity.Trigger{
		{
			Id: "1",
			Kind: "kind-1",
			Owner_id: 1,
			Created_at: "22/01/2024",
			Updated_at: "24/01/2024",
		},
		{
			Id: "2",
			Kind: "kind-2",
			Owner_id: 2,
			Created_at: "22/01/2024",
			Updated_at: "24/01/2024",
		},
		{
			Id: "3",
			Kind: "kind-3",
			Owner_id: 3,
			Created_at: "22/01/2024",
			Updated_at: "24/01/2024",
		},
	}
	pretty := true
	admin.HandleGetTriggers(&mockTrigger,&pretty)
}

func TestGetTriggerOwner(t *testing.T){
	mockTriggerOwner := adminEntity.TriggerOwner{
		Id: 1,
		Username: "cwc_user",
	}
	pretty := true
	admin.HandleGetTriggerOwner(&mockTriggerOwner, &pretty)
}
