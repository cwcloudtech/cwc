package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	// "github.com/spf13/cobra"
)

func AddObjectTypeInInteractiveMode (objectType *client.ObjectType) {
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

func PrepareAddObjectType(objectType *client.ObjectType, interactive *bool) (*client.ObjectType, error) {
	if *interactive {
		AddObjectTypeInInteractiveMode(objectType)
	}

	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_objectType, err := c.CreateObjectType(*objectType)
	utils.ExitIfError(err)

	return created_objectType, err
}

func HandleAddObjectType(createdObjectType *client.ObjectType, pretty *bool) {
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
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteObjectTypeById(*objectTypeId)
	utils.ExitIfError(err)

	fmt.Println("Object type successfully deleted")
}

func HandleGetObjectTypes(objectTypes *[]client.ObjectType, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayObjectTypesAsTable(*objectTypes)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(objectTypes)
	} else {
		var objectTypesDisplay []client.ObjectTypesDisplay
		for i, objectType := range *objectTypes {
			objectTypesDisplay = append(objectTypesDisplay, client.ObjectTypesDisplay{
				Id: objectType.Id,
				Name: objectType.Content.Name,
				Public: objectType.Content.Public,
				DecodingFunction: objectType.Content.DecodingFunction,
			})
			objectTypesDisplay[i].Id = objectType.Id
		}

		utils.PrintMultiRow(client.ObjectTypesDisplay{}, objectTypesDisplay)
	}
}

func displayObjectTypesAsTable(objectTypes []client.ObjectType) {
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

func HandleGetObjectType(objectType *client.ObjectType, pretty *bool) {
	var objectTypeDisplay client.ObjectTypesDisplay
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