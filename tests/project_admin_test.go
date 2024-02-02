package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestAddProject(t *testing.T) {
	userEmail := "unit_test_cwc@comwork.io"
	projectName := "project-cwc"
	host := "host-cwc"
	token := "token-cwc"
	gitUsername := "gituser_cwc_test"
	namespace := "unit_test_cwc"
	mockProject := &adminEntity.Project{
		Name: projectName,
	}
	admin.HandleAddProject(mockProject, &userEmail, &projectName, &host, &token, &gitUsername, &namespace)
}
