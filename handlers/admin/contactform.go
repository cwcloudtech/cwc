package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func HandleGetForms(forms *[]admin.ContactForm, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayFormsAsTable(*forms)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(forms)
	} else {
		var formsDisplay []admin.ContactForm
		for i, form := range *forms {
			formsDisplay = append(formsDisplay, admin.ContactForm{
				Id:         form.Id,
				Name:       form.Name,
				Hash:       form.Hash,
				MailFrom:   form.MailFrom,
				MailTo:     form.MailTo,
				UserId:     form.UserId,
				Updated_at: form.Updated_at,
			})
			formsDisplay[i].Id = form.Id
		}
		utils.PrintMultiRow(admin.ContactForm{}, formsDisplay)
	}
}

func HandleGetForm(form *admin.ContactForm, pretty *bool) {
	var formDisplay admin.ContactForm
	formDisplay.Id = form.Id
	formDisplay.Hash = form.Hash
	formDisplay.MailFrom = form.MailFrom
	formDisplay.MailTo = form.MailTo
	formDisplay.UserId = form.UserId
	formDisplay.Updated_at = form.Updated_at

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Found contact form", formDisplay)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(form)
	} else {
		utils.PrintRow(formDisplay)
	}
}

func PrepareAddForm(form *admin.ContactForm) (admin.ContactForm, error) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_form, err := c.AddForm(*form)
	utils.ExitIfError(err)
	return *created_form, err
}

func HandleAddForm(createdForm *admin.ContactForm, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Contact form successfully created", *createdForm)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(createdForm)
	} else {
		utils.PrintRow(*createdForm)
	}
}

func HandleUpdateForm(formId *string, updatedForm *admin.ContactForm) {
	c, err := admin.NewClient()
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
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteFormById(*formId)
	utils.ExitIfError(err)

	fmt.Println("Contact form successfully deleted")
}

func displayFormsAsTable(forms []admin.ContactForm) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Hash", "MailFrom", "MailTo", "UserId", "Updated_at"})

	if len(forms) == 0 {
		table.Append([]string{"No contact forms available", "404", "404", "404", "404", "404", "404"})
	} else {
		for _, form := range forms {
			table.Append([]string{
				form.Id,
				form.Name,
				form.Hash,
				form.MailFrom,
				form.MailTo,
				strconv.Itoa(form.UserId),
				form.Updated_at,
			})
		}
		table.Render()
	}
}
