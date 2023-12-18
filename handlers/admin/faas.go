package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleGetFunctions(pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	functions, err := c.GetAllFunctions()
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		displayFunctionsAsTable(*functions)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(functions)
	} else {
		var functionsDisplay []admin.FunctionDisplay
		for i, function := range *functions {
			functionsDisplay = append(functionsDisplay, admin.FunctionDisplay{
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

		utils.PrintMultiRow(admin.FunctionDisplay{}, functionsDisplay)
	}
}

func HandleGetFunctionOwner(id *string, pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	owner, err := c.GetFunctionOwnerById(*id)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Owner found", *owner)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(owner)
	} else {
		utils.PrintRow(*owner)
	}
}

func displayFunctionsAsTable(functions []admin.Function) {
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

func HandleGetInvocations(pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	invocations, err := c.GetAllInvocations()
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		displayInvocationsAsTable(*invocations)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(invocations)
	} else {
		var invocationsDisplay []admin.InvocationDisplay
		for i, invocation := range *invocations {
			invocationsDisplay = append(invocationsDisplay, admin.InvocationDisplay{
				Id:          invocation.Id,
				Invoker_id:  invocation.Invoker_id,
				Function_id: invocation.Content.Function_id,
				State:       invocation.Content.State,
				Created_at:  invocation.Created_at,
				Updated_at:  invocation.Updated_at,
			})
			invocationsDisplay[i].Id = invocation.Id
		}

		utils.PrintMultiRow(admin.InvocationDisplay{}, invocationsDisplay)
	}
}

func HandleGetInvocationInvoker(id *string, pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	invoker, err := c.GetInvocationInvokerById(*id)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Invoker found", *invoker)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(invoker)
	} else {
		utils.PrintRow(*invoker)
	}
}

func displayInvocationsAsTable(invocations []admin.Invocation) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Function ID", "State", "Created At", "Updated At"})
	if len(invocations) == 0 {
		// If there are no invocations available, display a message in a single cell.
		table.Append([]string{"No invocations available", "404", "404", "404", "404"})
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

func HandleGetTriggers(pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	triggers, err := c.GetAllTriggers()
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		displayTriggersAsTable(*triggers)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(triggers)
	} else {
		var triggersDisplay []admin.TriggerDisplay
		for i, trigger := range *triggers {
			triggersDisplay = append(triggersDisplay, admin.TriggerDisplay{
				Id:          trigger.Id,
				Kind:        trigger.Kind,
				Owner_id:    trigger.Owner_id,
				Name:        trigger.Content.Name,
				Cron_expr:   trigger.Content.Cron_expr,
				Function_id: trigger.Content.Function_id,
				Created_at:  trigger.Created_at,
				Updated_at:  trigger.Updated_at,
			})
			triggersDisplay[i].Id = trigger.Id
		}

		utils.PrintMultiRow(admin.TriggerDisplay{}, triggersDisplay)
	}
}

func HandleGetTriggerOwner(id *string, pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	owner, err := c.GetTriggerOwnerById(*id)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Owner found", *owner)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(owner)
	} else {
		utils.PrintRow(*owner)
	}
}

func displayTriggersAsTable(triggers []admin.Trigger) {
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
