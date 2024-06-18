package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestHandleGetClusters(t *testing.T) {
    mockClusters := []adminEntity.Cluster{
        {
            Id:               1,
            KubeconfigFileId: 1,
            Name:        "name",
            Platform: "platform",
            Version: "version",
            Created_at: "created_at",
        },
    }

    pretty := true
    admin.HandleGetClusters(&mockClusters, &pretty)
}
