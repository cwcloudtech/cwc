package user

import (
	"cwc/admin"
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleDeleteBucket(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteBucket(*id)
	utils.ExitIfError(err)

	fmt.Printf("Bucket %v successfully deleted\n", *id)
}

func HandleUpdateBucket(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateBucket(*id)
	utils.ExitIfError(err)

	fmt.Printf("Bucket %v successfully updated\n", *id)
}

func HandleGetBuckets(pretty *bool) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	buckets, err := c.GetAllBuckets()
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		displayBucketsAsTable(*buckets)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(buckets)
	} else {
		utils.PrintMultiRow(admin.Bucket{}, *buckets)
	}
}

func HandleGetBucket(id *string, pretty *bool) {
	c, err := client.NewClient()
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

func displayBucketsAsTable(buckets []client.Bucket) {
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
