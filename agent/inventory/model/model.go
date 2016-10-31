// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package inventory contains routines that periodically updates basic instance inventory to Inventory service
package model

//TODO: add all inventory types here
//TODO: if all attributes of inventory types become strong typed then we can directly refer to aws-sdk rather
//than defining everything here

const (
	// AWSInstanceInformation is inventory type of instance information
	AWSInstanceInformation = "AWS:InstanceInformation"
	// InventoryPluginName is name of inventory plugin
	InventoryPluginName = "Inventory"
	// Enabled represents constant string used to enable various components of inventory plugin
	Enabled = "Enabled"
	// ErrorThreshold represents error threshold for inventory plugin
	ErrorThreshold = 10
	// InventoryPolicyDocName represents name of inventory policy doc
	InventoryPolicyDocName = "policy.json"
	// SizeLimitKBPerInventoryType represents size limit in KB for 1 inventory data type
	SizeLimitKBPerInventoryType = 200
	// TotalSizeLimitKB represents size limit in KB for 1 PutInventory API call
	TotalSizeLimitKB = 1024
)

// Item encapsulates an inventory item
type Item struct {
	Name string
	//content depends on inventory type - hence set as interface{} here.
	//e.g: for application - it will contain []ApplicationData,
	//for instanceInformation - it will contain []InstanceInformation.
	Content       interface{}
	ContentHash   string
	SchemaVersion string
	CaptureTime   string
}

// InstanceInformation captures all attributes present in AWS:InstanceInformation inventory type
type InstanceInformation struct {
	AgentStatus     string
	AgentVersion    string
	ComputerName    string
	IPAddress       string
	InstanceID      string
	PlatformName    string
	PlatformType    string
	PlatformVersion string
}

// ApplicationData captures all attributes present in AWS:Application inventory type
type ApplicationData struct {
	Name            string
	Publisher       string
	Version         string
	InstalledTime   string
	ApplicationType string
	Architecture    string
	URL             string
}

// NetworkData captures all attributes present in AWS:Network inventory type
type NetworkData struct {
	Name       string
	SubnetMask string `json:",omitempty"`
	Gateway    string `json:",omitempty"`
	DHCPServer string `json:",omitempty"`
	DNSServer  string `json:",omitempty"`
	MacAddress string
	IPV4       string
	IPV6       string
}

// WindowsUpdateData captures all attributes present in AWS:WindowsUpdate inventory type
type WindowsUpdateData struct {
	HotFixID      string
	Description   string
	InstalledTime string
	InstalledBy   string
}

// Config captures all various properties (including optional) that can be supplied to a gatherer.
// NOTE: Not all properties will be applicable to all gatherers.
// E.g: Applications gatherer uses Collection, Files use Filters, Custom uses Collection & Location.
type Config struct {
	Collection string
	Filters    []string
	Location   string
}

// Policy defines how an inventory policy document looks like
// TODO: this struct might change depending on the type of data associate plugin provides to inventory plugin
// For e.g: this will incorporate association & runId after integrating with associate plugin.
type Policy struct {
	InventoryPolicy map[string]Config
}

// CustomInventoryItem represents the schema of custom inventory item
type CustomInventoryItem struct {
	TypeName      string
	SchemaVersion string
	Content       interface{}
}