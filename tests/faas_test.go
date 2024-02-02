package test

import (
    "cwc/handlers/user"
    "cwc/client"
    "testing"
)

func TestHandleGetTriggerKinds(t *testing.T) {
    mockTriggerKinds := client.TriggerKindsResponse{
        TriggerKinds: []string{"kind-1", "kind-2", "kind-3"},
    }
    pretty := true
    user.HandleGetTriggerKinds(&mockTriggerKinds, &pretty)
}

func TestHandleGetTriggers(t *testing.T) {
    mockTriggers := []client.Trigger{
        {
            Id: "test",
            Content: client.TriggerContent{
                Function_id: "147",
                Name:        "cwc_test",
                Cron_expr:   "test_cron",
            },
        },
    }
    pretty := true
    user.HandleGetTriggers(&mockTriggers, &pretty)
}

func TestHandleGetTrigger(t *testing.T) {
    mockTrigger := client.Trigger{
        Id: "test",
        Content: client.TriggerContent{
            Function_id: "879",
            Name:        "cwc_test",
            Cron_expr:   "test_cron",
            Args:        []client.Argument{},
        },
    }
    pretty := true
    user.HandleGetTrigger(&mockTrigger, &pretty)
}

func TestHandleGetLanguages(t *testing.T) {
    mockLanguage := client.LanguagesResponse{
        Languages: []string{
        "pyhton",
        "go",
        "code"},
    }
    pretty := true
    user.HandleGetLanguages(&mockLanguage, &pretty)
}

func TestHandleGetFunctions(t *testing.T) {
	mockFunctions := []client.Function{
		{
			Id:        "1",
			Owner_id:  1,
			Is_public: true,
			Content:   client.FunctionContent{
				Name:      "function-1",
				Language:  "language-1",
			},
			Created_at: "created_at-1",
			Updated_at: "updated_at-1",
		},
		{
			Id:        "2",
			Owner_id:  2,
			Is_public: false,
			Content:   client.FunctionContent{
				Name:      "function-2",
				Language:  "language-2",
			},
			Created_at: "created_at-2",
			Updated_at: "updated_at-2",
		},
		{
			Id:        "3",
			Owner_id:  3,
			Is_public: true,
			Content:   client.FunctionContent{
				Name:      "function-3",
				Language:  "language-3",
			},
			Created_at: "created_at-3",
			Updated_at: "updated_at-3",
		},
	}
	
	pretty := true
	user.HandleGetFunctions(&mockFunctions, &pretty)
}

func TestHandleGetFunction(t *testing.T) {
	mockFunction := &client.Function{
		Id:        "1",
		Owner_id:  1,
		Is_public: true,
		Content:   client.FunctionContent{
			Name:      "function-1",
			Language:  "language-1",
		},
		Created_at: "created_at-1",
		Updated_at: "updated_at-1",
	}
	pretty := true
	user.HandleGetFunction(mockFunction, &pretty)
}

func TestHandleAddFunction(t *testing.T){
	mockFunction := &client.Function{
		Id:        "1",
		Owner_id:  1,
		Is_public: true,
		Content:   client.FunctionContent{
			Name:      "function-1",
			Language:  "python",
		},
		Created_at: "created_at-1",
		Updated_at: "updated_at-1",
	}
	pretty := true
	user.HandleAddFunction(mockFunction, &pretty)  
}

func TestHandleGetInvocations(t *testing.T) {
	mockInvocations := []client.Invocation{
		{
			Id:         "1",
			Invoker_id: 2,
			Content: client.InvocationContent{
				Function_id: "3",
				State:       "4",
			},
			Created_at: "5",
			Updated_at: "6",
		},
	}
	pretty := true
	user.HandleGetInvocations(&mockInvocations, &pretty)
}

func TestHandleGetInvocation(t *testing.T) {
	mockInvocation := &client.Invocation{
		Id:         "test-invocation-id",
		Invoker_id: 123,
		Content: client.InvocationContent{
			Function_id: "test-function-id",
			State:       "SUCCESS",
		},
		Created_at: "2024-01-01T00:00:00Z",
		Updated_at: "2024-01-02T00:00:00Z",
	}
	pretty := true
	user.HandleGetInvocation(mockInvocation, &pretty)
}

func TestHandleAddInvocation(t *testing.T) {
	mockInvocation := &client.Invocation{
		Id:         "test-invocation-id",
		Invoker_id: 123,
		Content: client.InvocationContent{
			Function_id: "test-function-id",
			State:       "SUCCESS",
		},
		Created_at: "2024-01-01T00:00:00Z",
		Updated_at: "2024-01-02T00:00:00Z",
	}
	pretty := true
	user.HandleAddInvocation(mockInvocation, &pretty)
}
