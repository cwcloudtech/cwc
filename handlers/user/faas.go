package user

import (
	"bufio"
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func HandleGetLanguages(languages *client.LanguagesResponse, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPrettyArray("Available languages", languages.Languages)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(languages)
	} else {
		utils.PrintArray(languages.Languages)
	}
}

func HandleGetTriggerKinds(triggerKinds *client.TriggerKindsResponse, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPrettyArray("Available trigger kinds", triggerKinds.TriggerKinds)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(triggerKinds)
	} else {
		utils.PrintArray(triggerKinds.TriggerKinds)
	}
}

func HandleGetFunctions(functions *[]client.Function, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayFunctionsAsTable(*functions)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(functions)
	} else {
		var functionsDisplay []client.FunctionDisplay
		for i, function := range *functions {
			functionsDisplay = append(functionsDisplay, client.FunctionDisplay{
				Id:         function.Id,
				Owner_id:   function.Owner_id,
				Is_public:  function.Is_public,
				Name:       function.Content.Name,
				Language:   function.Content.Language,
				Created_at: function.Created_at,
				Updated_at: function.Updated_at,
			})
			functionsDisplay[i].Id = function.Id
		}

		utils.PrintMultiRow(client.FunctionDisplay{}, functionsDisplay)
	}
}

func HandleGetFunction(function *client.Function, pretty *bool) {
	var functionDisplay client.FunctionDisplay
	functionDisplay.Id = function.Id
	functionDisplay.Owner_id = function.Owner_id
	functionDisplay.Is_public = function.Is_public
	functionDisplay.Name = function.Content.Name
	functionDisplay.Language = function.Content.Language
	functionDisplay.Created_at = function.Created_at
	functionDisplay.Updated_at = function.Updated_at

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Found function", functionDisplay)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(function)
	} else {
		utils.PrintRow(functionDisplay)
	}
}

func displayFunctionsAsTable(functions []client.Function) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Language", "Public", "Created At", "Updated At"})

	if len(functions) == 0 {
		// If there are no functions available, display a message in a single cell.
		table.Append([]string{"No functions available", "404", "404", "404", "404", "404"})
	} else {
		for _, function := range functions {
			table.Append([]string{
				function.Id,
				function.Content.Name,
				function.Content.Language,
				fmt.Sprintf("%t", function.Is_public),
				function.Created_at,
				function.Updated_at,
			})
		}
	}

	table.Render() // Render the table
}

func HandleDeleteFunction(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteFunctionById(*id)
	utils.ExitIfError(err)

	fmt.Printf("Function successfully deleted\n")
}

func AddFunctionInInteractiveMode(function *client.Function) {
	// Prompt for Regexp
	fmt.Print("Enter Regexp (or press Enter for empty): ")
	fmt.Scanln(&function.Content.Regexp)

	// Prompt for Args array
	fmt.Println("Enter Args (one per line, press Enter for each entry; leave an empty line to finish):")
	for {
		var arg string
		_, err := fmt.Scanln(&arg)
		if nil != err {
			break 
		}
		function.Content.Args = append(function.Content.Args, arg)
	}
	if len(function.Content.Args) == 0 {
		function.Content.Args = []string{}
	}

	fmt.Print("Do you want to make the function public? [Y/N]: ")
	fmt.Scanln(&function.Is_public)

	fmt.Print("Do you want to add environment variables? [Y/N]: ")
	var addEnvVars string
	fmt.Scanln(&addEnvVars)
	if addEnvVars == "y" || addEnvVars == "Y" {
		fmt.Println("Enter environment variables (one per line, press Enter for each entry; leave an empty line to finish):")
		function.Content.Env = make(map[string]string) // Initialize the map
		for {
			fmt.Print("  ➤ Key: ")
			var envVarKey string
			_, err := fmt.Scanln(&envVarKey)
			if nil != err {
				break 
			}
			fmt.Print("  ➤ Value: ")
			var envVarValue string
			_, err = fmt.Scanln(&envVarValue)
			if nil != err {
				break 
			}
			function.Content.Env[envVarKey] = envVarValue
			fmt.Print("--------------------\n")
		}
	}
	if len(function.Content.Env) == 0 {
		function.Content.Env = map[string]string{}
	}

	fmt.Print("Do you want to add callbacks? [Y/N]: ")
	var addCallbacks string
	fmt.Scanln(&addCallbacks)
	if addCallbacks == "y" || addCallbacks == "Y" {
		fmt.Println("Enter callbacks (press Enter for each callback; leave an empty line to finish):")
		function.Content.Callbacks = []client.CallbacksContent{}
		for {
			var callback client.CallbacksContent
			fmt.Print("  ➤ Type (Current available types are http, websocket, and mqtt): ")
			_, err := fmt.Scanln(&callback.Type)
			if nil != err {
				break 
			}

			if callback.Type != "http" && callback.Type != "websocket" && callback.Type != "mqtt" {
				fmt.Println("Invalid callback type. Available types are http, websocket, and mqtt")
			}

			fmt.Print("  ➤ Endpoint: ")
			_, err = fmt.Scanln(&callback.Endpoint)
			if nil != err {
				break 
			}

			if callback.Type == "http" {
				fmt.Print("  ➤ Token: ")
				_, err = fmt.Scanln(&callback.Token)
				if nil != err {
					break 
				}
			} else {
				fmt.Print("  ➤ Client ID: ")
				_, err = fmt.Scanln(&callback.Client_id)
				if nil != err {
					break 
				}

				fmt.Print("  ➤ User Data: ")
				_, err = fmt.Scanln(&callback.User_data)
				if nil != err {
					break 
				}

				fmt.Print("  ➤ Username: ")
				_, err = fmt.Scanln(&callback.Username)
				if nil != err {
					break 
				}

				fmt.Print("  ➤ Password: ")
				_, err = fmt.Scanln(&callback.Password)
				if nil != err {
					break 
				}

				fmt.Print("  ➤ Port: ")
				_, err = fmt.Scanln(&callback.Port)
				if nil != err {
					break 
				}

				fmt.Print("  ➤ Subscription: ")
				_, err = fmt.Scanln(&callback.Subscription)
				if nil != err {
					break 
				}

				fmt.Print("  ➤ QoS: ")
				_, err = fmt.Scanln(&callback.Qos)
				if nil != err {
					break 
				}

				fmt.Print("  ➤ Topic: ")
				_, err = fmt.Scanln(&callback.Topic)
				if nil != err {
					break 
				}
			}
			function.Content.Callbacks = append(function.Content.Callbacks, callback)
			fmt.Print("--------------------\n")
		}
	}

	c, err := client.NewClient()
	utils.ExitIfError(err)

	// assign the code template after choosing the language
	code_template, err := c.GetFunctionCodeTemplate(function.Content.Args, function.Content.Language)
	utils.ExitIfError(err)

	fmt.Print("Do you want to add code? [Y/N]: ")
	var addCode string
	fmt.Scanln(&addCode)

	if addCode == "y" || addCode == "Y" {
		editorCommand := utils.GetSystemEditor()

		// Create a temporary file with a specific name and path
		tempFileName := "temp-code-editor.txt"
		tempFile, err := os.Create(tempFileName)
		utils.ExitIfErrorWithMsg("Error creating temporary file", err)

		defer tempFile.Close()
		defer os.Remove(tempFileName)

		// Write the code_template to the temporary file
		_, err = tempFile.WriteString(*code_template)
		utils.ExitIfErrorWithMsg("Error writing code_template to the temporary file", err)

		// Prompt the user to write code in the editor
		fmt.Printf("Please write your code in the text editor that opens. Save and close the editor when done.\n")

		cmd := exec.Command(editorCommand, tempFileName)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		utils.ExitIfErrorWithMsg("Error opening the text editor", err)

		// Read the code from the temporary file
		codeBytes, err := os.ReadFile(tempFileName)
		utils.ExitIfErrorWithMsg("Error reading code from the text editor", err)

		function.Content.Code = string(codeBytes)
	}

	fmt.Printf("code: %s\n", function.Content.Code)
}

func PrepareAddFunction(function *client.Function, interactive *bool) (*client.Function, error) {
	language_response, err := client.GetLanguages()
	utils.ExitIfError(err)

	isLanguageAllowed := false
	for _, allowedLang := range language_response.Languages {
		if function.Content.Language == allowedLang {
			isLanguageAllowed = true
			break
		}
	}

	utils.ExitIfNeeded(fmt.Sprintf("Invalid language. Allowed languages are: %s", strings.Join(language_response.Languages, ", ")), !isLanguageAllowed)

	if *interactive {
		AddFunctionInInteractiveMode(function)
	}

	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_function, err := c.AddFunction(*function)
	utils.ExitIfError(err)
	return created_function, err
}

func HandleAddFunction(createdFunction *client.Function, pretty *bool) {
	if createdFunction == nil {
		// Handle the nil case to prevent panic
		fmt.Println("Error: createdFunction is nil")
		return
	}
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Function successfully created", *createdFunction)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(*createdFunction)
	} else {
		utils.PrintRow(*createdFunction)
	}
}

func HandleUpdateFunction(id *string, updated_function *client.Function, interactive *bool) {
	language_response, err := client.GetLanguages()
	utils.ExitIfError(err)

	c, err := client.NewClient()
	utils.ExitIfError(err)

	function, err := c.GetFunctionById(*id)
	utils.ExitIfError(err)

	isLanguageAllowed := false

	if *interactive {
		// prompt to update language
		fmt.Printf("Current language: %s\n", function.Content.Language)
		fmt.Printf("Available languages are: %s\n", strings.Join(language_response.Languages, ", "))
		fmt.Printf("Enter new language (or press Enter to keep the current one): ")
		fmt.Scanln(&function.Content.Language)

		for _, allowedLang := range language_response.Languages {
			if function.Content.Language == allowedLang {
				isLanguageAllowed = true
				break
			}
		}

		utils.ExitIfNeeded(fmt.Sprintf("Allowed languages are: %s", strings.Join(language_response.Languages, ", ")), !isLanguageAllowed)

		// Prompt for Regexp
		fmt.Printf("Current regexp: %s\n", function.Content.Regexp)
		fmt.Print("Enter new regexp (or press Enter to keep the current one): ")
		fmt.Scanln(&function.Content.Regexp)

		// Prompt for function name
		fmt.Printf("Current name: %s\n", function.Content.Name)
		fmt.Print("Enter new name (or press Enter to keep the current one): ")
		fmt.Scanln(&function.Content.Name)

		// Ask if the function should be public
		if function.Is_public {
			fmt.Print("Current function status: Public\n")
		} else {
			fmt.Print("Current function status: Private\n")
		}

		fmt.Print("Do you want to make the change the function status? [Y/N]: ")

		var answer string
		fmt.Scanln(&answer)
		if answer == "y" || answer == "Y" {
			function.Is_public = !function.Is_public
		}

		// Prompt for new Args array
		utils.PrintPrettyArray("Current args", function.Content.Args)
		fmt.Println("Enter new Args (one per line, press Enter for each entry; leave an empty line to finish):")
		for {
			var arg string
			_, err := fmt.Scanln(&arg)
			if nil != err {
				break 
			}
			function.Content.Args = append(function.Content.Args, arg)
		}

		// Prompt for new code
		fmt.Print("Do you want to update code? [Y/N]: ")
		var updateCode string
		fmt.Scanln(&updateCode)

		if updateCode == "y" || updateCode == "Y" {
			editorCommand := utils.GetSystemEditor()

			// Create a temporary file with a specific name and path
			tempFileName := "temp-code-editor-update.txt"
			tempFile, err := os.Create(tempFileName)
			utils.ExitIfErrorWithMsg("Error creating temporary file", err)

			defer os.Remove(tempFileName)

			// Write the current code to the temporary file
			_, err = tempFile.WriteString(function.Content.Code)
			utils.ExitIfErrorWithMsg("Error writing current code to the temporary file", err)

			// Prompt the user to edit the code in the editor
			fmt.Printf("Please update your code in the text editor that opens. Save and close the editor when done.\n")

			cmd := exec.Command(editorCommand, tempFileName)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err = cmd.Run()
			utils.ExitIfErrorWithMsg("Error opening the text editor", err)

			// Read the updated code from the temporary file
			updatedCodeBytes, err := os.ReadFile(tempFileName)
			utils.ExitIfErrorWithMsg("Error reading updated code from the text editor", err)

			// Update the function's code with the edited code
			function.Content.Code = string(updatedCodeBytes)
		}
	} else {
		// If interactive mode is not enabled, update only the fields that are not empty
		if utils.IsNotBlank(updated_function.Content.Language) {
			function.Content.Language = updated_function.Content.Language
		}

		if utils.IsNotBlank(updated_function.Content.Regexp) {
			function.Content.Regexp = updated_function.Content.Regexp
		}

		if utils.IsNotBlank(updated_function.Content.Name) {
			function.Content.Name = updated_function.Content.Name
		}

		if updated_function.Is_public {
			function.Is_public = updated_function.Is_public
		}

		if len(updated_function.Content.Args) > 0 {
			function.Content.Args = updated_function.Content.Args
		}

		if utils.IsNotBlank(updated_function.Content.Code) {
			function.Content.Code = updated_function.Content.Code
		}
	}

	_, err = c.UpdateFunction(*function)
	utils.ExitIfError(err)

	fmt.Println("Function successfully updated")
}

func HandleGetInvocations(invocations *[]client.Invocation, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayInvocationsAsTable(*invocations)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(invocations)
	} else {
		var invocationsDisplay []client.InvocationDisplay
		for i, invocation := range *invocations {
			invocationsDisplay = append(invocationsDisplay, client.InvocationDisplay{
				Id:          invocation.Id,
				Invoker_id:  invocation.Invoker_id,
				Function_id: invocation.Content.Function_id,
				State:       invocation.Content.State,
				Created_at:  invocation.Created_at,
				Updated_at:  invocation.Updated_at,
			})
			invocationsDisplay[i].Id = invocation.Id
		}

		utils.PrintMultiRow(client.InvocationDisplay{}, invocationsDisplay)
	}
}

func HandleGetInvocation(invocation *client.Invocation, pretty *bool) {
	var invocationDisplay client.InvocationDisplay
	invocationDisplay.Id = invocation.Id
	invocationDisplay.Invoker_id = invocation.Invoker_id
	invocationDisplay.Function_id = invocation.Content.Function_id
	invocationDisplay.State = invocation.Content.State
	invocationDisplay.Created_at = invocation.Created_at
	invocationDisplay.Updated_at = invocation.Updated_at

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Found invocation", invocationDisplay)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(invocation)
	} else {
		utils.PrintRow(invocationDisplay)
	}
}

func displayInvocationsAsTable(invocations []client.Invocation) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Function ID", "State", "Created At", "Updated At"})
	if len(invocations) == 0 {
		// If there are no invocations available, display a message in a single cell.
		table.Append([]string{"No invocations available", "404", "404", "404", "404", "404"})
	} else {
		for _, invocation := range invocations {
			table.Append([]string{
				invocation.Id,
				invocation.Content.Function_id,
				invocation.Content.State,
				invocation.Created_at,
				invocation.Updated_at,
			})
		}
	}
	table.Render() // Render the table
}

func PrepareAddInvocation(content *client.InvocationAddContent, argument_values *[]string, interactive *bool, synchronous *bool) (*client.Invocation, error) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	var id = &content.Function_id
	function, _ := c.GetFunctionById(*id)
	args := function.Content.Args

	if *interactive {
		// Prompt values of the existing arguments
		if len(args) > 0 {
			fmt.Println("Enter Args (one per line, press Enter for each entry; leave an empty line to finish):")
			for _, arg := range args {
				fmt.Printf("  ➤ %s: ", arg)
				var value string
				fmt.Scanln(&value)
				content.Args = append(content.Args, client.Argument{Key: arg, Value: value})
			}
		}
	} else {
		utils.ExitIfNeeded(fmt.Sprintf("Invalid number of arguments. Expected %d arguments, got %d", len(args), len(*argument_values)), len(*argument_values) != len(args))

		if len(args) > 0 {
			for i, arg := range args {
				content.Args = append(content.Args, client.Argument{Key: arg, Value: (*argument_values)[i]})
			}
		}
	}

	created_invocation, err := c.AddInvocation(*content, *synchronous)
	utils.ExitIfError(err)

	return created_invocation, err
}

func HandleAddInvocation(created_invocation *client.Invocation, pretty *bool) {
	var invocationDisplay client.InvocationDisplay
	invocationDisplay.Id = created_invocation.Id
	invocationDisplay.State = created_invocation.Content.State
	invocationDisplay.Created_at = created_invocation.Created_at
	invocationDisplay.Updated_at = created_invocation.Updated_at
	invocationDisplay.Function_id = created_invocation.Content.Function_id
	invocationDisplay.Result = created_invocation.Content.Result

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Invocation successfully created", invocationDisplay)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(*created_invocation)
	} else {
		utils.PrintRow(invocationDisplay)
	}
}

func HandleDeleteInvocation(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteInvocationById(*id)
	utils.ExitIfError(err)

	fmt.Println("Invocation successfully deleted")
}

func HandleTruncateInvocations() {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.TruncateInvocations()
	utils.ExitIfError(err)

	fmt.Println("Invocations successfully truncated")
}

func HandleGetTriggers(triggers *[]client.Trigger, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayTriggersAsTable(*triggers)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(triggers)
	} else {
		var triggersDisplay []client.TriggerDisplay
		for i, trigger := range *triggers {
			triggersDisplay = append(triggersDisplay, client.TriggerDisplay{
				Id:          trigger.Id,
				Function_id: trigger.Content.Function_id,
				Kind:        trigger.Kind,
				Name:        trigger.Content.Name,
				Cron_expr:   trigger.Content.Cron_expr,
				Created_at:  trigger.Created_at,
				Updated_at:  trigger.Updated_at,
			})
			triggersDisplay[i].Id = trigger.Id
		}
		utils.PrintMultiRow(client.TriggerDisplay{}, triggersDisplay)
	}
}

func HandleGetTrigger(trigger *client.Trigger, pretty *bool) {
	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(trigger)
	} else {
		var triggerDisplay client.TriggerDisplay
		triggerDisplay.Id = trigger.Id
		triggerDisplay.Function_id = trigger.Content.Function_id
		triggerDisplay.Kind = trigger.Kind
		triggerDisplay.Name = trigger.Content.Name
		triggerDisplay.Cron_expr = trigger.Content.Cron_expr
		triggerDisplay.Created_at = trigger.Created_at
		triggerDisplay.Updated_at = trigger.Updated_at

		if config.IsPrettyFormatExpected(pretty) {
			utils.PrintPretty("Found trigger", triggerDisplay)
		} else if config.GetDefaultFormat() == "json" {
			utils.PrintJson(trigger)
		} else {
			utils.PrintRow(triggerDisplay)
		}
	}
}

func displayTriggersAsTable(triggers []client.Trigger) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Function ID", "Kind", "Name", "Cron Expr", "Created At", "Updated At"})

	if len(triggers) == 0 {
		// If there are no functions available, display a message in a single cell.
		table.Append([]string{"No triggers available", "404", "404", "404", "404"})
	} else {
		for _, trigger := range triggers {
			table.Append([]string{
				trigger.Id,
				trigger.Content.Function_id,
				trigger.Kind,
				trigger.Content.Name,
				trigger.Content.Cron_expr,
				trigger.Created_at,
				trigger.Updated_at,
			})
		}
	}

	table.Render()
}

func HandleAddTrigger(trigger *client.Trigger, argument_values *[]string, interactive *bool, pretty *bool) {
	triggerKinds, _ := client.GetKinds()
	isTriggerKindAllowed := false
	var id = &trigger.Content.Function_id
	args, _ := client.GetFunctionByIdArgs(*id)

	if *interactive {
		// Create a scanner to read user input
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt for trigger kind
		fmt.Printf("Enter one of these Trigger Kinds:\n")
		for _, available_triggerKind := range triggerKinds.TriggerKinds {
			fmt.Printf("  ➤ %v\n", available_triggerKind)
		}

		fmt.Printf("Trigger kind: ")
		scanner.Scan()
		trigger.Kind = scanner.Text()
		for _, available_triggerKind := range triggerKinds.TriggerKinds {
			if trigger.Kind == available_triggerKind {
				isTriggerKindAllowed = true
				break
			}
		}

		utils.ExitIfNeeded(fmt.Sprintf("Invalid trigger kind. Allowed trigger kinds are: %s", strings.Join(triggerKinds.TriggerKinds, ", ")), !isTriggerKindAllowed)

		// Prompt for trigger name
		fmt.Printf("Enter Trigger name: ")
		scanner.Scan()
		trigger.Content.Name = scanner.Text()

		// Prompt for trigger cron expression
		fmt.Printf("Enter Trigger cron expression: ")
		scanner.Scan()
		trigger.Content.Cron_expr = scanner.Text()

		// Prompt values of the existing arguments
		if len(args) > 0 {
			fmt.Println("Enter Args (one per line, press Enter for each entry; leave an empty line to finish):")
			for _, arg := range args {
				fmt.Printf("  ➤ %s: ", arg)
				scanner.Scan()
				value := scanner.Text()
				trigger.Content.Args = append(trigger.Content.Args, client.Argument{Key: arg, Value: value})
			}
		}
	} else {
		utils.ExitIfNeeded(fmt.Sprintf("Invalid number of arguments. Expected %d arguments, got %d\n", len(args), len(*argument_values)), len(*argument_values) != len(args))

		if len(*argument_values) > 0 {
			for i, arg := range args {
				trigger.Content.Args = append(trigger.Content.Args, client.Argument{Key: arg, Value: (*argument_values)[i]})
			}
		}
	}

	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_trigger, err := c.AddTrigger(*trigger)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Trigger successfully created", *created_trigger)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(created_trigger)
	} else {
		utils.PrintRow(*created_trigger)
	}
}

func HandleDeleteTrigger(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteTriggerById(*id)
	utils.ExitIfError(err)

	fmt.Println("Trigger successfully deleted")
}

func HandleTruncateTriggers() {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.TruncateTriggers()
	utils.ExitIfError(err)

	fmt.Println("Triggers successfully truncated")
}
