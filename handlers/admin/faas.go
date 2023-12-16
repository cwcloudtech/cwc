package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func HandleGetFunctions(pretty *bool) {
	client, _ := admin.NewClient()
	functions, err := client.GetAllFunctions()
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(functions)
	} else {
		if *pretty {
			displayFunctionsAsTable(*functions)
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
}

func HandleGetFunctionOwner(id *string, pretty *bool) {
	client, _ := admin.NewClient()
	owner, err := client.GetFunctionOwnerById(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if *pretty {
		fmt.Printf("  ➤ ID: %s\n", strconv.Itoa(owner.Id))
		fmt.Printf("  ➤ Username: %s\n", owner.Username)
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
	client, _ := admin.NewClient()
	invocations, err := client.GetAllInvocations()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(invocations)
	} else if *pretty {
		displayInvocationsAsTable(*invocations)
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
	client, _ := admin.NewClient()
	invoker, err := client.GetInvocationInvokerById(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if *pretty {
		fmt.Printf("  ➤ ID: %s\n", strconv.Itoa(invoker.Id))
		fmt.Printf("  ➤ Username: %s\n", invoker.Username)
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
	client, _ := admin.NewClient()
	triggers, err := client.GetAllTriggers()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(triggers)
	} else {
		if *pretty {
			displayTriggersAsTable(*triggers)
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
}

func HandleGetTriggerOwner(id *string, pretty *bool) {
	client, _ := admin.NewClient()
	owner, err := client.GetTriggerOwnerById(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if *pretty {
		fmt.Printf("  ➤ ID: %s\n", strconv.Itoa(owner.Id))
		fmt.Printf("  ➤ Username: %s\n", owner.Username)
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
