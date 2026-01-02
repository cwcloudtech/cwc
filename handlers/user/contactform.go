package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleGetForms(forms *[]client.ContactForm, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayFormsAsTable(*forms)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(forms)
	} else {
		var formsDisplay []client.ContactForm
		for i, form := range *forms {
			formsDisplay = append(formsDisplay, client.ContactForm{
				Id:         form.Id,
				Name:       form.Name,
				Hash:       form.Hash,
				MailFrom:   form.MailFrom,
				MailTo:     form.MailTo,
				Updated_at: form.Updated_at,
			})
			formsDisplay[i].Id = form.Id
		}
		utils.PrintMultiRow(client.ContactForm{}, formsDisplay)
	}
}

func HandleGetForm(form *client.ContactForm, pretty *bool) {
	var formDisplay client.ContactForm
	formDisplay.Id = form.Id
	formDisplay.Name = form.Name
	formDisplay.Hash = form.Hash
	formDisplay.MailFrom = form.MailFrom
	formDisplay.MailTo = form.MailTo
	formDisplay.Updated_at = form.Updated_at

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Found contact form", formDisplay)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(form)
	} else {
		utils.PrintRow(formDisplay)
	}
}

func PrepareAddForm(form *client.ContactForm) (client.ContactForm, error) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_form, err := c.AddForm(*form)
	utils.ExitIfError(err)
	return *created_form, err
}

func HandleAddForm(createdForm *client.ContactForm, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Form successfully created", *createdForm)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(createdForm)
	} else {
		utils.PrintRow(*createdForm)
	}
}

func HandleUpdateForm(formId *string, updatedForm *client.ContactForm) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	form, err := c.GetFormById(*formId)
	utils.ExitIfError(err)

	if utils.IsNotBlank(updatedForm.MailFrom) {
		form.MailFrom = updatedForm.MailFrom
	}

	if utils.IsNotBlank(updatedForm.MailTo) {
		form.MailTo = updatedForm.MailTo
	}

	if utils.IsNotBlank(updatedForm.LogoUrl) {
		form.LogoUrl = updatedForm.LogoUrl
	}

	if utils.IsNotBlank(updatedForm.CopyrightName) {
		form.CopyrightName = updatedForm.CopyrightName
	}

	if utils.IsNotBlank(updatedForm.Name) {
		form.Name = updatedForm.Name
	} else {
		form.Name = utils.ShortName(form.Name, form.Hash)
	}

	_, updateError := c.UpdateFormById(*formId, *form)
	utils.ExitIfError(updateError)

	fmt.Println("Contact form successfully updated")
}

func HandleDeleteForm(formId *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteFormById(*formId)
	utils.ExitIfError(err)

	fmt.Println("Form successfully deleted")
}

func displayFormsAsTable(forms []client.ContactForm) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Hash", "MailFrom", "MailTo", "Updated_at"})

	if len(forms) == 0 {
		table.Append([]string{"No contact form available", "404", "404", "404", "404", "404"})
	} else {
		for _, form := range forms {
			table.Append([]string{
				form.Id,
				form.Name,
				form.Hash,
				form.MailFrom,
				form.MailTo,
				form.Updated_at,
			})
		}
		table.Render()
	}
}
