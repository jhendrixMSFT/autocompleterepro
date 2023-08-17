package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5/fake"
)

func main() {
	_ = fake.VirtualMachinesServer{
		Get: func(ctx context.Context, resourceGroupName, vmName string, options *armcompute.VirtualMachinesClientGetOptions) (resp fake.Responder[armcompute.VirtualMachinesClientGetResponse], errResp fake.ErrorResponder) {
		},
	}
}
