package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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

func HandleTransferBucketOwnership(id *string, email *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateBucket(*id, *email)
	utils.ExitIfError(err)

	fmt.Printf("Bucket with id %v successfully transferred to this email owner: %v\n", *id, *email)
}

func HandleRenewBucketCredentials(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateBucket(*id)
	utils.ExitIfError(err)

	fmt.Printf("Bucket with id %v successfully renewed\n", *id)
}

func HandleGetBuckets(buckets *[]admin.Bucket, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayBucketsAsTable(*buckets)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(buckets)
	} else {
		utils.PrintMultiRow(admin.Bucket{}, *buckets)
	}
}

func HandleGetBucket(bucket *admin.Bucket, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Bucket's informations", *bucket)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(bucket)
	} else {
		utils.PrintRow(*bucket)
	}
}

func displayBucketsAsTable(buckets []admin.Bucket) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Type", "Endpoint", "Region", "Created at"})

	if len(buckets) == 0 {
		fmt.Println("No buckets found")
	} else {
		for _, bucket := range buckets {
			table.Append([]string{
				fmt.Sprintf("%d", bucket.Id),
				bucket.Name,
				bucket.Type,
				bucket.Endpoint,
				bucket.Region,
				bucket.CreatedAt,
			})
		}
	}

	table.Render()
}
