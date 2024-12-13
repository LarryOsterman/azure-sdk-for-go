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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListNetworkSecurityPerimeterConfigurations.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().NewListPager("resourceGroupName", "accountName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.NetworkSecurityPerimeterConfigurationList = armcognitiveservices.NetworkSecurityPerimeterConfigurationList{
		// 	Value: []*armcognitiveservices.NetworkSecurityPerimeterConfiguration{
		// 		{
		// 			Name: to.Ptr("networkSecurityPerimeterConfigurationName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/networkSecurityPerimeterConfigurations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/networkSecurityPerimeterConfigurations"),
		// 			Properties: &armcognitiveservices.NetworkSecurityPerimeterConfigurationProperties{
		// 				NetworkSecurityPerimeter: &armcognitiveservices.NetworkSecurityPerimeter{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Network/networkSecurityPerimeters/perimeter"),
		// 					Location: to.Ptr("East US"),
		// 					PerimeterGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 				Profile: &armcognitiveservices.NetworkSecurityPerimeterProfileInfo{
		// 					Name: to.Ptr("profileName"),
		// 					AccessRules: []*armcognitiveservices.NetworkSecurityPerimeterAccessRule{
		// 						{
		// 							Name: to.Ptr("ruleName"),
		// 							Properties: &armcognitiveservices.NetworkSecurityPerimeterAccessRuleProperties{
		// 								AddressPrefixes: []*string{
		// 									to.Ptr("148.0.0.0/8"),
		// 									to.Ptr("152.4.6.0/24")},
		// 									Direction: to.Ptr(armcognitiveservices.NspAccessRuleDirectionInbound),
		// 								},
		// 						}},
		// 						AccessRulesVersion: to.Ptr[int64](1),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ResourceAssociation: &armcognitiveservices.NetworkSecurityPerimeterConfigurationAssociationInfo{
		// 						Name: to.Ptr("associationName"),
		// 						AccessMode: to.Ptr("Enforced"),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/GetNetworkSecurityPerimeterConfigurations.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().Get(ctx, "resourceGroupName", "accountName", "NSPConfigurationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSecurityPerimeterConfiguration = armcognitiveservices.NetworkSecurityPerimeterConfiguration{
	// 	Name: to.Ptr("networkSecurityPerimeterConfigurationName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/networkSecurityPerimeterConfigurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/networkSecurityPerimeterConfigurations"),
	// 	Properties: &armcognitiveservices.NetworkSecurityPerimeterConfigurationProperties{
	// 		NetworkSecurityPerimeter: &armcognitiveservices.NetworkSecurityPerimeter{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Network/networkSecurityPerimeters/perimeter"),
	// 			Location: to.Ptr("East US"),
	// 			PerimeterGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Profile: &armcognitiveservices.NetworkSecurityPerimeterProfileInfo{
	// 			Name: to.Ptr("profileName"),
	// 			AccessRules: []*armcognitiveservices.NetworkSecurityPerimeterAccessRule{
	// 				{
	// 					Name: to.Ptr("ruleName"),
	// 					Properties: &armcognitiveservices.NetworkSecurityPerimeterAccessRuleProperties{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("148.0.0.0/8"),
	// 							to.Ptr("152.4.6.0/24")},
	// 							Direction: to.Ptr(armcognitiveservices.NspAccessRuleDirectionInbound),
	// 						},
	// 				}},
	// 				AccessRulesVersion: to.Ptr[int64](1),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ResourceAssociation: &armcognitiveservices.NetworkSecurityPerimeterConfigurationAssociationInfo{
	// 				Name: to.Ptr("associationName"),
	// 				AccessMode: to.Ptr("Enforced"),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ReconcileNetworkSecurityPerimeterConfigurations.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_BeginReconcile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().BeginReconcile(ctx, "resourceGroupName", "accountName", "NSPConfigurationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSecurityPerimeterConfiguration = armcognitiveservices.NetworkSecurityPerimeterConfiguration{
	// 	Name: to.Ptr("networkSecurityPerimeterConfigurationName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/networkSecurityPerimeterConfigurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/networkSecurityPerimeterConfigurations"),
	// 	Properties: &armcognitiveservices.NetworkSecurityPerimeterConfigurationProperties{
	// 		NetworkSecurityPerimeter: &armcognitiveservices.NetworkSecurityPerimeter{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Network/networkSecurityPerimeters/perimeter"),
	// 			Location: to.Ptr("East US"),
	// 			PerimeterGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Profile: &armcognitiveservices.NetworkSecurityPerimeterProfileInfo{
	// 			Name: to.Ptr("profileName"),
	// 			AccessRules: []*armcognitiveservices.NetworkSecurityPerimeterAccessRule{
	// 				{
	// 					Name: to.Ptr("ruleName"),
	// 					Properties: &armcognitiveservices.NetworkSecurityPerimeterAccessRuleProperties{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("148.0.0.0/8"),
	// 							to.Ptr("152.4.6.0/24")},
	// 							Direction: to.Ptr(armcognitiveservices.NspAccessRuleDirectionInbound),
	// 						},
	// 				}},
	// 				AccessRulesVersion: to.Ptr[int64](1),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ResourceAssociation: &armcognitiveservices.NetworkSecurityPerimeterConfigurationAssociationInfo{
	// 				Name: to.Ptr("associationName"),
	// 				AccessMode: to.Ptr("Enforced"),
	// 			},
	// 		},
	// 	}
}
