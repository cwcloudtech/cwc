package user

import (
	"cwc/admin"
	"cwc/client"
	"cwc/utils"
	"fmt"
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

func HandleGetBuckets() {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	buckets, err := c.GetAllBuckets()
	utils.ExitIfError(err)

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(buckets)
	} else {
		utils.PrintMultiRow(admin.Bucket{}, *buckets)
	}
}

func HandleGetBucket(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	bucket, err := c.GetBucket(*id)
	utils.ExitIfError(err)

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(bucket)
	} else {
		utils.PrintRow(*bucket)
	}
}
