//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/CheckSkuAvailability.json
func ExampleManagementClient_CheckSKUAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().CheckSKUAvailability(ctx, "westus", armcognitiveservices.CheckSKUAvailabilityParameter{
		Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		Kind: to.Ptr("Face"),
		SKUs: []*string{
			to.Ptr("S0")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SKUAvailabilityListResult = armcognitiveservices.SKUAvailabilityListResult{
	// 	Value: []*armcognitiveservices.SKUAvailability{
	// 		{
	// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 			Kind: to.Ptr("Face"),
	// 			SKUAvailable: to.Ptr(true),
	// 			SKUName: to.Ptr("S0"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/CheckDomainAvailability.json
func ExampleManagementClient_CheckDomainAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().CheckDomainAvailability(ctx, armcognitiveservices.CheckDomainAvailabilityParameter{
		Type:          to.Ptr("Microsoft.CognitiveServices/accounts"),
		SubdomainName: to.Ptr("contosodemoapp1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DomainAvailability = armcognitiveservices.DomainAvailability{
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 	IsSubdomainAvailable: to.Ptr(false),
	// 	Reason: to.Ptr("Sub domain name 'contosodemoapp1' is not valid"),
	// 	SubdomainName: to.Ptr("contosodemoapp1"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/CalculateModelCapacity.json
func ExampleManagementClient_CalculateModelCapacity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().CalculateModelCapacity(ctx, armcognitiveservices.CalculateModelCapacityParameter{
		Model: &armcognitiveservices.DeploymentModel{
			Name:    to.Ptr("gpt-4"),
			Format:  to.Ptr("OpenAI"),
			Version: to.Ptr("0613"),
		},
		SKUName: to.Ptr("ProvisionedManaged"),
		Workloads: []*armcognitiveservices.ModelCapacityCalculatorWorkload{
			{
				RequestParameters: &armcognitiveservices.ModelCapacityCalculatorWorkloadRequestParam{
					AvgGeneratedTokens: to.Ptr[int64](50),
					AvgPromptTokens:    to.Ptr[int64](30),
				},
				RequestPerMinute: to.Ptr[int64](10),
			},
			{
				RequestParameters: &armcognitiveservices.ModelCapacityCalculatorWorkloadRequestParam{
					AvgGeneratedTokens: to.Ptr[int64](20),
					AvgPromptTokens:    to.Ptr[int64](60),
				},
				RequestPerMinute: to.Ptr[int64](20),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CalculateModelCapacityResult = armcognitiveservices.CalculateModelCapacityResult{
	// 	EstimatedCapacity: &armcognitiveservices.CalculateModelCapacityResultEstimatedCapacity{
	// 		DeployableValue: to.Ptr[int32](400),
	// 		Value: to.Ptr[int32](346),
	// 	},
	// 	Model: &armcognitiveservices.DeploymentModel{
	// 		Name: to.Ptr("gpt-4"),
	// 		Format: to.Ptr("OpenAI"),
	// 		Version: to.Ptr("0613"),
	// 	},
	// 	SKUName: to.Ptr("ProvisionedManaged"),
	// }
}
