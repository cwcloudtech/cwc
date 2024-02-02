package test

import (
    "cwc/client"
    "cwc/handlers/user"
    "testing"
)

func TestHandleGetModels(t *testing.T) {
    mockModels := &client.ModelsResponse{
        Models: []string{"model_cwc_1", "model_cwc_2", "model_cwc_3"},
    }
    pretty := true
    user.HandleGetModels(mockModels, &pretty)
}
