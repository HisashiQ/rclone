package trafficmanager

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// EndpointMonitorStatus enumerates the values for endpoint monitor status.
type EndpointMonitorStatus string

const (
	// CheckingEndpoint specifies the checking endpoint state for endpoint monitor status.
	CheckingEndpoint EndpointMonitorStatus = "CheckingEndpoint"
	// Degraded specifies the degraded state for endpoint monitor status.
	Degraded EndpointMonitorStatus = "Degraded"
	// Disabled specifies the disabled state for endpoint monitor status.
	Disabled EndpointMonitorStatus = "Disabled"
	// Inactive specifies the inactive state for endpoint monitor status.
	Inactive EndpointMonitorStatus = "Inactive"
	// Online specifies the online state for endpoint monitor status.
	Online EndpointMonitorStatus = "Online"
	// Stopped specifies the stopped state for endpoint monitor status.
	Stopped EndpointMonitorStatus = "Stopped"
)

// EndpointStatus enumerates the values for endpoint status.
type EndpointStatus string

const (
	// EndpointStatusDisabled specifies the endpoint status disabled state for endpoint status.
	EndpointStatusDisabled EndpointStatus = "Disabled"
	// EndpointStatusEnabled specifies the endpoint status enabled state for endpoint status.
	EndpointStatusEnabled EndpointStatus = "Enabled"
)

// MonitorProtocol enumerates the values for monitor protocol.
type MonitorProtocol string

const (
	// HTTP specifies the http state for monitor protocol.
	HTTP MonitorProtocol = "HTTP"
	// HTTPS specifies the https state for monitor protocol.
	HTTPS MonitorProtocol = "HTTPS"
	// TCP specifies the tcp state for monitor protocol.
	TCP MonitorProtocol = "TCP"
)

// ProfileMonitorStatus enumerates the values for profile monitor status.
type ProfileMonitorStatus string

const (
	// ProfileMonitorStatusCheckingEndpoints specifies the profile monitor status checking endpoints state for profile
	// monitor status.
	ProfileMonitorStatusCheckingEndpoints ProfileMonitorStatus = "CheckingEndpoints"
	// ProfileMonitorStatusDegraded specifies the profile monitor status degraded state for profile monitor status.
	ProfileMonitorStatusDegraded ProfileMonitorStatus = "Degraded"
	// ProfileMonitorStatusDisabled specifies the profile monitor status disabled state for profile monitor status.
	ProfileMonitorStatusDisabled ProfileMonitorStatus = "Disabled"
	// ProfileMonitorStatusInactive specifies the profile monitor status inactive state for profile monitor status.
	ProfileMonitorStatusInactive ProfileMonitorStatus = "Inactive"
	// ProfileMonitorStatusOnline specifies the profile monitor status online state for profile monitor status.
	ProfileMonitorStatusOnline ProfileMonitorStatus = "Online"
)

// ProfileStatus enumerates the values for profile status.
type ProfileStatus string

const (
	// ProfileStatusDisabled specifies the profile status disabled state for profile status.
	ProfileStatusDisabled ProfileStatus = "Disabled"
	// ProfileStatusEnabled specifies the profile status enabled state for profile status.
	ProfileStatusEnabled ProfileStatus = "Enabled"
)

// TrafficRoutingMethod enumerates the values for traffic routing method.
type TrafficRoutingMethod string

const (
	// Geographic specifies the geographic state for traffic routing method.
	Geographic TrafficRoutingMethod = "Geographic"
	// Performance specifies the performance state for traffic routing method.
	Performance TrafficRoutingMethod = "Performance"
	// Priority specifies the priority state for traffic routing method.
	Priority TrafficRoutingMethod = "Priority"
	// Weighted specifies the weighted state for traffic routing method.
	Weighted TrafficRoutingMethod = "Weighted"
)

// CheckTrafficManagerRelativeDNSNameAvailabilityParameters is parameters supplied to check Traffic Manager name
// operation.
type CheckTrafficManagerRelativeDNSNameAvailabilityParameters struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// CloudError is an error returned by the Azure Resource Manager
type CloudError struct {
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody is the content of an error returned by the Azure Resource Manager
type CloudErrorBody struct {
	Code    *string           `json:"code,omitempty"`
	Message *string           `json:"message,omitempty"`
	Target  *string           `json:"target,omitempty"`
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// DeleteOperationResult is the result of the request or operation.
type DeleteOperationResult struct {
	autorest.Response `json:"-"`
	OperationResult   *bool `json:"boolean,omitempty"`
}

// DNSConfig is class containing DNS settings in a Traffic Manager profile.
type DNSConfig struct {
	RelativeName *string `json:"relativeName,omitempty"`
	Fqdn         *string `json:"fqdn,omitempty"`
	TTL          *int64  `json:"ttl,omitempty"`
}

// Endpoint is class representing a Traffic Manager endpoint.
type Endpoint struct {
	autorest.Response   `json:"-"`
	ID                  *string `json:"id,omitempty"`
	Name                *string `json:"name,omitempty"`
	Type                *string `json:"type,omitempty"`
	*EndpointProperties `json:"properties,omitempty"`
}

// EndpointProperties is class representing a Traffic Manager endpoint properties.
type EndpointProperties struct {
	TargetResourceID      *string               `json:"targetResourceId,omitempty"`
	Target                *string               `json:"target,omitempty"`
	EndpointStatus        EndpointStatus        `json:"endpointStatus,omitempty"`
	Weight                *int64                `json:"weight,omitempty"`
	Priority              *int64                `json:"priority,omitempty"`
	EndpointLocation      *string               `json:"endpointLocation,omitempty"`
	EndpointMonitorStatus EndpointMonitorStatus `json:"endpointMonitorStatus,omitempty"`
	MinChildEndpoints     *int64                `json:"minChildEndpoints,omitempty"`
	GeoMapping            *[]string             `json:"geoMapping,omitempty"`
}

// GeographicHierarchy is class representing the Geographic hierarchy used with the Geographic traffic routing method.
type GeographicHierarchy struct {
	autorest.Response              `json:"-"`
	ID                             *string `json:"id,omitempty"`
	Name                           *string `json:"name,omitempty"`
	Type                           *string `json:"type,omitempty"`
	*GeographicHierarchyProperties `json:"properties,omitempty"`
}

// GeographicHierarchyProperties is class representing the properties of the Geographic hierarchy used with the
// Geographic traffic routing method.
type GeographicHierarchyProperties struct {
	GeographicHierarchy *Region `json:"geographicHierarchy,omitempty"`
}

// HeatMapEndpoint is class which is a sparse representation of a Traffic Manager endpoint.
type HeatMapEndpoint struct {
	ResourceID *string `json:"resourceId,omitempty"`
	EndpointID *int32  `json:"endpointId,omitempty"`
}

// HeatMapModel is class representing a Traffic Manager HeatMap.
type HeatMapModel struct {
	autorest.Response  `json:"-"`
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	Type               *string `json:"type,omitempty"`
	*HeatMapProperties `json:"properties,omitempty"`
}

// HeatMapProperties is class representing a Traffic Manager HeatMap properties.
type HeatMapProperties struct {
	StartTime    *date.Time         `json:"startTime,omitempty"`
	EndTime      *date.Time         `json:"endTime,omitempty"`
	Endpoints    *[]HeatMapEndpoint `json:"endpoints,omitempty"`
	TrafficFlows *[]TrafficFlow     `json:"trafficFlows,omitempty"`
}

// MonitorConfig is class containing endpoint monitoring settings in a Traffic Manager profile.
type MonitorConfig struct {
	ProfileMonitorStatus      ProfileMonitorStatus `json:"profileMonitorStatus,omitempty"`
	Protocol                  MonitorProtocol      `json:"protocol,omitempty"`
	Port                      *int64               `json:"port,omitempty"`
	Path                      *string              `json:"path,omitempty"`
	IntervalInSeconds         *int64               `json:"intervalInSeconds,omitempty"`
	TimeoutInSeconds          *int64               `json:"timeoutInSeconds,omitempty"`
	ToleratedNumberOfFailures *int64               `json:"toleratedNumberOfFailures,omitempty"`
}

// NameAvailability is class representing a Traffic Manager Name Availability response.
type NameAvailability struct {
	autorest.Response `json:"-"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
	NameAvailable     *bool   `json:"nameAvailable,omitempty"`
	Reason            *string `json:"reason,omitempty"`
	Message           *string `json:"message,omitempty"`
}

// Profile is class representing a Traffic Manager profile.
type Profile struct {
	autorest.Response  `json:"-"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Type               *string             `json:"type,omitempty"`
	Tags               *map[string]*string `json:"tags,omitempty"`
	Location           *string             `json:"location,omitempty"`
	*ProfileProperties `json:"properties,omitempty"`
}

// ProfileListResult is the list Traffic Manager profiles operation response.
type ProfileListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Profile `json:"value,omitempty"`
}

// ProfileProperties is class representing the Traffic Manager profile properties.
type ProfileProperties struct {
	ProfileStatus        ProfileStatus        `json:"profileStatus,omitempty"`
	TrafficRoutingMethod TrafficRoutingMethod `json:"trafficRoutingMethod,omitempty"`
	DNSConfig            *DNSConfig           `json:"dnsConfig,omitempty"`
	MonitorConfig        *MonitorConfig       `json:"monitorConfig,omitempty"`
	Endpoints            *[]Endpoint          `json:"endpoints,omitempty"`
}

// ProxyResource is the resource model definition for a ARM proxy resource. It will have everything other than required
// location and tags
type ProxyResource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// QueryExperience is class representing a Traffic Manager HeatMap query experience properties.
type QueryExperience struct {
	EndpointID *int32   `json:"endpointId,omitempty"`
	QueryCount *int32   `json:"queryCount,omitempty"`
	Latency    *float64 `json:"latency,omitempty"`
}

// Region is class representing a region in the Geographic hierarchy used with the Geographic traffic routing method.
type Region struct {
	Code    *string   `json:"code,omitempty"`
	Name    *string   `json:"name,omitempty"`
	Regions *[]Region `json:"regions,omitempty"`
}

// Resource is the core properties of ARM resources
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// TrackedResource is the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
	Location *string             `json:"location,omitempty"`
}

// TrafficFlow is class representing a Traffic Manager HeatMap traffic flow properties.
type TrafficFlow struct {
	SourceIP         *string            `json:"sourceIp,omitempty"`
	Latitude         *float64           `json:"latitude,omitempty"`
	Longitude        *float64           `json:"longitude,omitempty"`
	QueryExperiences *[]QueryExperience `json:"queryExperiences,omitempty"`
}
