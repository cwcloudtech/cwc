package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestAdminHandleGetInstances(t *testing.T){
	mockInstances := []adminEntity.Instance{
		{
			Id:  1,
			Name: "instance-1",
		},
		{
			Id:  2,
			Name: "instance-2",
		},
		{
			Id:  3,
			Name: "instance-3",
		},
	}
	pretty := true
	admin.HandleGetInstances(&mockInstances, &pretty)
}

func TestAdminHandleGetInstance(t *testing.T){
	mockInstance := &adminEntity.Instance{
		Id:  1,
		Name: "instance-1",
	}
	pretty := true
	admin.HandleGetInstance(mockInstance, &pretty)
}

func TestHandleAddInstance(t *testing.T){
	mockInstance := &adminEntity.Instance{
		Id:  1,
		Name: "instance-1",
		Zone: "zone-1",
		Region: "region-1",
		Instance_type: "instance_type-1",
		Root_dns_zone: "root_dns_zone-1",
		Environment: "environment-1",
		Project: 1,
		Email: "email-1",
		Project_name: "project_name-1",
		Project_url: "project_url-1",
	}
	user_email := "email-1"
	name := "instance-1"
	project_id := 1
	project_name := "project_name-1"
	project_url := "project_url-1"
	env := "environment-1"
	instance_type := "instance_type-1"
	zone := "zone-1"
	dns_zone := "root_dns_zone-1"
	
	admin.HandleAddInstance(mockInstance, &user_email, &name, &project_id, &project_name, &project_url, &env, &instance_type, &zone, &dns_zone)
}
