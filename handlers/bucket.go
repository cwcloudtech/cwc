package handlers

import (
	"cwc/client"
	"fmt"
	"os"
)

func HandleDeleteBucket(id *string) {
	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.DeleteBucket(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Bucket %v successfully deleted\n", *id)
}

func HandleUpdateBucket(id *string) {

	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.UpdateBucket(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Bucket %v successfully updated\n", *id)

}

func HandleGetBucket() {

	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

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
