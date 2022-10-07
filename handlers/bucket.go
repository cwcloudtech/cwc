package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleDeleteBucket(deleteCmd *flag.FlagSet, id *string) {

	deleteCmd.Parse(os.Args[3:])
	if *id == "" {
		fmt.Println("id is required to delete your bucket")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	err := client.DeleteBucket(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Bucket %v successfully deleted\n", *id)
}

func HandleUpdateBucket(updateCmd *flag.FlagSet, id *string) {
	updateCmd.Parse(os.Args[3:])
	if *id == "" {
		fmt.Println("id are required")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	err := client.UpdateBucket(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Bucket %v successfully updated\n", *id)

}

func HandleGetBucket(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[3:])
	if !*all && *id == "" {
		fmt.Println("id is required or specify --all to get all buckets.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	if *all {

		buckets, err := client.GetAllBuckets()

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("ID\tcreated_at\tname\tstatus\taccess_key\tsecret_key\tendpoint\n")
		for _, bucket := range *buckets {
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", bucket.Id, bucket.CreatedAt, bucket.Name, bucket.Status, bucket.AccessKey, bucket.SecretKey, bucket.Endpoint)

		}

		return
	}

	if *id != "" {
		bucket, err := client.GetBucket(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID\tcreated_at\tname\tstatus\taccess_key\tsecret_key\tendpoint\n")
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", bucket.Id, bucket.CreatedAt, bucket.Name, bucket.Status, bucket.AccessKey, bucket.SecretKey, bucket.Endpoint)

		return
	}
}
