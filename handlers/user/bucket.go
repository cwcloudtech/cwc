package user

import (
	"cwc/admin"
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleDeleteBucket(id *string) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.DeleteBucket(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Bucket %v successfully deleted\n", *id)
}

func HandleUpdateBucket(id *string) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.UpdateBucket(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Bucket %v successfully updated\n", *id)
}

func HandleGetBuckets() {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	buckets, err := c.GetAllBuckets()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(buckets)
	} else {
		utils.PrintMultiRow(admin.Bucket{}, *buckets)
	}
}

func HandleGetBucket(id *string) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	bucket, err := c.GetBucket(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(bucket)
	} else {
		utils.PrintRow(*bucket)
	}
}
