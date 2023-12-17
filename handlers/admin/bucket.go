package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleAddBucket(user_email *string, name *string, reg_type *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	created_bucket, err := c.AdminAddBucket(*user_email, *name, *reg_type)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(created_bucket)
	} else {
		utils.PrintRow(*created_bucket)
	}
}

func HandleDeleteBucket(id *string) {
	c, err := admin.NewClient()
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

func HandleUpdateBucket(id *string, email *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.UpdateBucket(*id, *email)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Bucket %v successfully updated\n", *id)
}

func HandleGetBuckets() {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	buckets, err := c.GetAllBuckets()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(buckets)
	} else {
		utils.PrintMultiRow(admin.Bucket{}, *buckets)
	}
}

func HandleGetBucket(id *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	bucket, err := c.GetBucket(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(bucket)
	} else {
		utils.PrintRow(*bucket)
	}
}
