// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type SlotVirtualNetworkSwiftConnectionId struct {
	SubscriptionId string
	ResourceGroup  string
	SiteName       string
	SlotName       string
	ConfigName     string
}

func NewSlotVirtualNetworkSwiftConnectionID(subscriptionId, resourceGroup, siteName, slotName, configName string) SlotVirtualNetworkSwiftConnectionId {
	return SlotVirtualNetworkSwiftConnectionId{
		SubscriptionId: subscriptionId,
		ResourceGroup:  resourceGroup,
		SiteName:       siteName,
		SlotName:       slotName,
		ConfigName:     configName,
	}
}

func (id SlotVirtualNetworkSwiftConnectionId) String() string {
	segments := []string{
		fmt.Sprintf("Config Name %q", id.ConfigName),
		fmt.Sprintf("Slot Name %q", id.SlotName),
		fmt.Sprintf("Site Name %q", id.SiteName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Slot Virtual Network Swift Connection", segmentsStr)
}

func (id SlotVirtualNetworkSwiftConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/config/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.SiteName, id.SlotName, id.ConfigName)
}

// SlotVirtualNetworkSwiftConnectionID parses a SlotVirtualNetworkSwiftConnection ID into an SlotVirtualNetworkSwiftConnectionId struct
func SlotVirtualNetworkSwiftConnectionID(input string) (*SlotVirtualNetworkSwiftConnectionId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("parsing %q as an SlotVirtualNetworkSwiftConnection ID: %+v", input, err)
	}

	resourceId := SlotVirtualNetworkSwiftConnectionId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, errors.New("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, errors.New("ID was missing the 'resourceGroups' element")
	}

	if resourceId.SiteName, err = id.PopSegment("sites"); err != nil {
		return nil, err
	}
	if resourceId.SlotName, err = id.PopSegment("slots"); err != nil {
		return nil, err
	}
	if resourceId.ConfigName, err = id.PopSegment("config"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
