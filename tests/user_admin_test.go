package test

import (
	adminEntity "cwc/admin"
	"cwc/handlers/admin"
	"testing"
)

func TestHandleGetUsers(t *testing.T) {
	mockUsers := &adminEntity.ResponseUsers{
		Result: []adminEntity.User{
			{
				Id:                 1,
				Email:              "cwc_test@comwork.io",
				RegistrationNumber: "registration_number-1",
				Address:            "address-1",
				CompanyName:        "comwork_io",
				ContactInfo:        "012345698",
				IsAdmin:            true,
				Confirmed:          true,
			},
			{
				Id:                 2,
				Email:              "cwc_test@comwork.io",
				RegistrationNumber: "registration_number-2",
				Address:            "address-2",
				CompanyName:        "comwork_io",
				ContactInfo:        "012345698",
				IsAdmin:            false,
				Confirmed:          false,
			},
			{
				Id:                 3,
				Email:              "cwc_test@comwork.io",
				RegistrationNumber: "registration_number-3",
				Address:            "address-3",
				CompanyName:        "comwork_io",
				ContactInfo:        "012345698",
				IsAdmin:            false,
				Confirmed:          false,
			},
			{
				Id:                 4,
				Email:              "email-4",
				RegistrationNumber: "registration_number-4",
				Address:            "address-4",
				CompanyName:        "company_name-4",
				ContactInfo:        "contact_info-4",
				IsAdmin:            false,
				Confirmed:          false,
			},
		},
	}
	pretty := true
	admin.HandleGetUsers(mockUsers, &pretty)
}

func TestGetUser(t *testing.T) {
	mockUser := &adminEntity.ResponseUser{
		Result: adminEntity.User{
			Id:                 1,
			Email:              "cwc_test@comwork.io",
			RegistrationNumber: "registration_number-1",
			Address:            "address-1",
			CompanyName:        "comwork_io",
			ContactInfo:        "012345698",
			IsAdmin:            true,
			Confirmed:          true,
		},
	}
	pretty := true
	admin.HandleGetUser(mockUser, &pretty)
}
