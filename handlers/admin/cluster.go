package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleGetClusters(clusters *[]admin.Cluster, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayClustersAsTable(*clusters)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(clusters)
	} else {
		utils.PrintMultiRow(admin.Cluster{}, clusters)
	}
}

func displayClustersAsTable(clusters []admin.Cluster) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Platform", "Version", "Created At"})

	if len(clusters) == 0 {
		fmt.Println("No clusters found")
	} else {
		for _, cluster := range clusters {
			table.Append([]string{
				fmt.Sprintf("%d", cluster.Id),
				cluster.Name,
				cluster.Platform,
				cluster.Version,
				cluster.Created_at,
			})
		}
	}

	table.Render()
}

func HandleDeleteCluster(clusterId *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteCluster(*clusterId)
	utils.ExitIfError(err)

	fmt.Printf("Cluster with id %v successfully deleted\n", *clusterId)
}
