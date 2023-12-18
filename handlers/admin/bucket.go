package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
)

func HandleAddBucket(user_email *string, name *string, reg_type *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_bucket, err := c.AdminAddBucket(*user_email, *name, *reg_type)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(created_bucket)
	} else {
		utils.PrintRow(*created_bucket)
	}
}

func HandleDeleteBucket(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteBucket(*id)
	utils.ExitIfError(err)

	fmt.Printf("Bucket %v successfully deleted\n", *id)
}

func HandleUpdateBucket(id *string, email *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateBucket(*id, *email)
	utils.ExitIfError(err)

	fmt.Printf("Bucket %v successfully updated\n", *id)
}

func HandleGetBuckets() {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	buckets, err := c.GetAllBuckets()
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(buckets)
	} else {
		utils.PrintMultiRow(admin.Bucket{}, *buckets)
	}
}

func HandleGetBucket(id *string, pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	bucket, err := c.GetBucket(*id)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Bucket's informations", *bucket)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(bucket)
	} else {
		utils.PrintRow(*bucket)
	}
}
