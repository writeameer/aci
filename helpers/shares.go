package helpers

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-07-01/storage"
)

func CreateStorage() {
	// Authenticate with Azure
	authorizer, sid := AzureAuth()

	// Setup ARM Client

	client := storage.NewAccountsClient(sid)
	client.Authorizer = authorizer

	location := "Australia East"
	sku := storage.Sku{
		Name: storage.StandardLRS,
		Tier: storage.Standard,
	}
	kind := storage.StorageV2
	accountCreateFuture, err := client.Create(ctx, "aci-example", "storagetestbee", storage.AccountCreateParameters{
		Location: &location,
		Kind:     kind,
		Sku:      &sku,
	})

	PrintError(err)

	err = accountCreateFuture.WaitForCompletion(ctx, client.BaseClient.Client)
	PrintError(err)

	account, err := accountCreateFuture.Result(client)

	log.Println(account.Name)
}