package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func AddObjectTypeInInteractiveMode (objectType *admin.ObjectType) {
	// Prompt the id of the owner
	fmt.Print("Enter the id of the owner: ")
	fmt.Scanln(&objectType.User_id)

	// Prompt for the name of the object type
	fmt.Print("Enter the name of the object type (or press Enter to skip): ")
	fmt.Scanln(&objectType.Content.Name)

	// Prompt for the public status of the object type
	fmt.Print("Is the object type public? (Y/N): ")
	fmt.Scanln(&objectType.Content.Public)

	// Prompt for the decoding function of the object type
	fmt.Print("Enter the decoding function of the object type: ")
	fmt.Scanln(&objectType.Content.DecodingFunction)
	if objectType.Content.DecodingFunction == "" {
		fmt.Println("The decoding function is required")
		fmt.Print("--------------------")
		fmt.Print("Enter the decoding function of the object type")
		fmt.Scanln(&objectType.Content.DecodingFunction)
	}

	// Prompt to ask if the user want to add triggers
	fmt.Print("Do you want to add triggers? (Y/N): ")
	var addTriggers string
	fmt.Scanln(&addTriggers)
	if addTriggers == "y" || addTriggers == "Y" {
		// Prompt for the triggers of the object type
		fmt.Println("Enter the trigger id (one per line, press Enter for each entry; leave an empty line to finish): ")
		for {
			var trigger string
			fmt.Print("  ➤ Trigger id: ")
			fmt.Scanln(&trigger)
			if trigger == "" {
				break
			}
			objectType.Content.Triggers = append(objectType.Content.Triggers, trigger)
		}
		if len(objectType.Content.Triggers) == 0 {
			objectType.Content.Triggers = []string{}
		}
	}
}

func PrepareAddObjectType(objectType *admin.ObjectType, interactive *bool) (*admin.ObjectType, error) {
	if *interactive {
		AddObjectTypeInInteractiveMode(objectType)
	}

	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_objectType, err := c.CreateObjectType(*objectType)
	utils.ExitIfError(err)

	return created_objectType, err
}

func HandleAddObjectType(createdObjectType *admin.ObjectType, pretty *bool) {
	if createdObjectType == nil {
		fmt.Println("Object type not created")
		return
	}
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Object type successfully created", *createdObjectType)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(createdObjectType)
	} else {
		utils.PrintRow(createdObjectType)
	}
}

func HandleDeleteObjectType(objectTypeId *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteObjectTypeById(*objectTypeId)
	utils.ExitIfError(err)

	fmt.Println("Object type successfully deleted")
}

func HandleGetObjectTypes(objectTypes *[]admin.ObjectType, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayObjectTypesAsTable(*objectTypes)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(objectTypes)
	} else {
		var objectTypesDisplay []admin.ObjectTypesDisplay
		for i, objectType := range *objectTypes {
			objectTypesDisplay = append(objectTypesDisplay, admin.ObjectTypesDisplay{
				Id: objectType.Id,
				Name: objectType.Content.Name,
				Public: objectType.Content.Public,
				DecodingFunction: objectType.Content.DecodingFunction,
			})
			objectTypesDisplay[i].Id = objectType.Id
		}

		utils.PrintMultiRow(admin.ObjectTypesDisplay{}, objectTypesDisplay)
	}
}

func displayObjectTypesAsTable(objectTypes []admin.ObjectType) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Public", "Decoding Function"})
	
	if len(objectTypes) == 0 {
		table.Append([]string{"No object types available", "404", "404", "404"})
	} else {
		for _, objectType := range objectTypes {
			table.Append([]string{
				objectType.Id,
				objectType.Content.Name,
				fmt.Sprintf("%t", objectType.Content.Public),
				objectType.Content.DecodingFunction,
			})
		}
	}
	table.Render()
}

func HandleGetObjectType(objectType *admin.ObjectType, pretty *bool) {
	var objectTypeDisplay admin.ObjectTypesDisplay
	objectTypeDisplay.Id = objectType.Id
	objectTypeDisplay.Name = objectType.Content.Name
	objectTypeDisplay.Public = objectType.Content.Public
	objectTypeDisplay.DecodingFunction = objectType.Content.DecodingFunction

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Object type details", objectTypeDisplay)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(objectType)
	} else {
		utils.PrintRow(objectTypeDisplay)
	}
}

func UpdateObjectTypeInInteractiveMode(objectType *admin.ObjectType) {
	// Prompt the id of the owner
	fmt.Println("Current owner id: ", objectType.User_id)
	fmt.Print("Enter the id of the owner: ")
	fmt.Scanln(&objectType.User_id)

	// Prompt for the name of the object type
	fmt.Println("Current name: ", objectType.Content.Name)
	fmt.Print("Enter the name of the object type (or press Enter to skip): ")
	fmt.Scanln(&objectType.Content.Name)

	// Prompt for the public status of the object type
	fmt.Println("Current public status: ", objectType.Content.Public)
	fmt.Print("Is the object type public? (Y/N): ")
	fmt.Scanln(&objectType.Content.Public)

	// Prompt for the decoding function of the object type
	fmt.Println("Current decoding function: ", objectType.Content.DecodingFunction)
	fmt.Print("Enter the decoding function of the object type (or press Enter to skip): ")
	fmt.Scanln(&objectType.Content.DecodingFunction)
	if objectType.Content.DecodingFunction == "" {
		fmt.Println("The decoding function is required")
		fmt.Print("--------------------")
		fmt.Print("Enter the decoding function of the object type")
		fmt.Scanln(&objectType.Content.DecodingFunction)
	}

	// Prompt to ask if the user want to add triggers
	fmt.Println("Current triggers Ids: ", objectType.Content.Triggers)
	fmt.Print("Do you want to recreate triggers? (Y/N): ")
	var addTriggers string
	fmt.Scanln(&addTriggers)
	if addTriggers == "y" || addTriggers == "Y" {
		// Prompt for the triggers of the object type
		fmt.Println("Enter the trigger id (one per line, press Enter for each entry; leave an empty line to finish): ")
		for {
			var trigger string
			fmt.Print("  ➤ Trigger id: ")
			fmt.Scanln(&trigger)
			if trigger == "" {
				break
			}
			objectType.Content.Triggers = append(objectType.Content.Triggers, trigger)
		}
		if len(objectType.Content.Triggers) == 0 {
			objectType.Content.Triggers = []string{}
		}
	}
}

func HandleUpdateObjectType(id *string, updated_objectType *admin.ObjectType, interactive *bool) {

	c, err := admin.NewClient()
	utils.ExitIfError(err)

	objectType, err := c.GetObjectTypeById(*id)
	utils.ExitIfError(err)

	if *interactive {
		UpdateObjectTypeInInteractiveMode(objectType)
	} else {
		if updated_objectType.User_id != 0 {
			objectType.User_id = updated_objectType.User_id
		}

		if utils.IsNotBlank(updated_objectType.Content.Name) {
			objectType.Content.Name = updated_objectType.Content.Name
		}

		if utils.IsNotBlank(updated_objectType.Content.DecodingFunction) {
			objectType.Content.DecodingFunction = updated_objectType.Content.DecodingFunction
		}

		if len(updated_objectType.Content.Triggers) > 0 {
			objectType.Content.Triggers = updated_objectType.Content.Triggers
		}

		if updated_objectType.Content.Public {
			objectType.Content.Public = updated_objectType.Content.Public
		}
	}

	_, err = c.UpdateObjectType(*objectType)
	utils.ExitIfError(err)

	fmt.Println("Object type successfully updated")
}

func displayDevicesAsTable(devices []admin.Device) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Object Type ID", "Username", "Active"})

	if len(devices) == 0 {
		table.Append([]string{"No devices available", "404", "404", "404"})
	} else {
		for _, device := range devices {
			table.Append([]string{
				device.Id,
				device.Typeobject_id,
				device.Username,
				fmt.Sprintf("%t", device.Active),
			})
		}
	}
	table.Render()
}

func HandleGetDevices(devices *[]admin.Device, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayDevicesAsTable(*devices)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(devices)
	} else {
		var devicesDisplay []admin.DeviceDisplay
		for i, device := range *devices {
			devicesDisplay = append(devicesDisplay, admin.DeviceDisplay(device))
			devicesDisplay[i].Id = device.Id
		}
		utils.PrintMultiRow(admin.DeviceDisplay{}, devicesDisplay)
	}
}

func HandleDeleteDevice(deviceId *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteDeviceById(*deviceId)
	utils.ExitIfError(err)

	fmt.Println("Device successfully deleted")
}

func displayNumericDataAsTable(numericData []admin.NumericData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Device ID", "Value", "Created At"})

	if len(numericData) == 0 {
		table.Append([]string{"No numeric data available", "404", "404", "404"})
	} else {
		for _, data := range numericData {
			table.Append([]string{
				data.Id,
				data.Device_id,
				fmt.Sprintf("%f", data.Value),
				data.Created_at,
			})
		}
	}
	table.Render()
}

func HandleGetNumericData(numericData *[]admin.NumericData, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayNumericDataAsTable(*numericData)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(numericData)
	} else {
		var numericDataDisplay []admin.NumericData
		for i, data := range *numericData {
			numericDataDisplay = append(numericDataDisplay, admin.NumericData(data))
			numericDataDisplay[i].Id = data.Id
		}
		utils.PrintMultiRow(admin.NumericData{}, numericDataDisplay)
	}
}

func displayStringDataAsTable(stringData []admin.StringData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Device ID", "Value", "Created At"})

	if len(stringData) == 0 {
		table.Append([]string{"No string data available", "404", "404", "404"})
	} else {
		for _, data := range stringData {
			table.Append([]string{
				data.Id,
				data.Device_id,
				data.Value,
				data.Created_at,
			})
		}
	}
	table.Render()
}

func HandleGetStringData(stringData *[]admin.StringData, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayStringDataAsTable(*stringData)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(stringData)
	} else {
		var stringDataDisplay []admin.StringData
		for i, data := range *stringData {
			stringDataDisplay = append(stringDataDisplay, admin.StringData(data))
			stringDataDisplay[i].Id = data.Id
		}
		utils.PrintMultiRow(admin.StringData{}, stringDataDisplay)
	}
}
