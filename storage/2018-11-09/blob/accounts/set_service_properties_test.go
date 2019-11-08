package accounts

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/storage/mgmt/storage"
	"github.com/tombuildsstuff/giovanni/testhelpers"
)

func TestContainerLifecycle(t *testing.T) {
	client, err := testhelpers.Build()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.TODO()

	resourceGroup := fmt.Sprintf("acctestrg-%d", testhelpers.RandomInt())
	accountName := fmt.Sprintf("acctestsa%s", testhelpers.RandomString())

	_, err = client.BuildTestResources(ctx, resourceGroup, accountName, storage.StorageV2)
	if err != nil {
		t.Fatal(err)
	}
	defer client.DestroyTestResources(ctx, resourceGroup, accountName)

	accountsClient := NewWithEnvironment(client.Environment)
	accountsClient.Client = client.PrepareWithStorageResourceManagerAuth(accountsClient.Client)

	input := StorageServiceProperties{}
	_, err = accountsClient.SetServiceProperties(ctx, accountName, input)
	if err != nil {
		t.Fatal(fmt.Errorf("error setting properties: %s", err))
	}

	var index = "index.html"
	var enabled = true
	var errorDocument = "404.html"

	input = StorageServiceProperties{
		StaticWebsite: &StaticWebsite{
			Enabled: &enabled,
			IndexDocument: &index,
			ErrorDocument404Path: &errorDocument,
		},
	}

	_, err = accountsClient.SetServiceProperties(ctx, accountName, input)
	if err != nil {
		t.Fatal(fmt.Errorf("error setting properties: %s", err))
	}
}
