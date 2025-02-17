package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleGetAdapters(t *testing.T) {
	mockAdapters := &client.AiAdaptersResponse{
		Adapters: []string{"adapter_cwc_1", "adapter_cwc_2", "adapter_cwc_3"},
	}
	pretty := true
	user.HandleGetAiAdapters(mockAdapters, &pretty)
}
