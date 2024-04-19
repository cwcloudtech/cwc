package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
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
