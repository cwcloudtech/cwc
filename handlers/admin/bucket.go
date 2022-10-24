package admin

import (
	"cwc/admin"
	"fmt"
	"os"
)

func HandleAddBucket(user_email *string, name *string, reg_type *string) {
	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	created_bucket, err := client.AdminAddBucket(*user_email, *name, *reg_type)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\tstatus\taccess_key\tsecret_key\tendpoint\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_bucket.Id, created_bucket.CreatedAt, created_bucket.Name, created_bucket.Status, created_bucket.AccessKey, created_bucket.SecretKey, created_bucket.Endpoint)
}

func HandleDeleteBucket(id *string) {
	client, err := admin.NewClient()
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

	client, err := admin.NewClient()
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

func HandleGetBuckets() {

	client, err := admin.NewClient()
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


func HandleGetBucket(id *string) {

	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	bucket, err := client.GetBucket(*id)

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID\tcreated_at\tname\tstatus\taccess_key\tsecret_key\tendpoint\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", bucket.Id, bucket.CreatedAt, bucket.Name, bucket.Status, bucket.AccessKey, bucket.SecretKey, bucket.Endpoint)


	return
}
