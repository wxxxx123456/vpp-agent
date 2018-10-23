// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package descriptor

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/ligato/cn-infra/logging"
	scheduler "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	ifdescriptor "github.com/ligato/vpp-agent/plugins/vppv2/ifplugin/descriptor"
	"github.com/ligato/vpp-agent/plugins/vppv2/l3plugin/descriptor/adapter"
	"github.com/ligato/vpp-agent/plugins/vppv2/l3plugin/vppcalls"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/interfaces"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/l3"
)

const (
	// StaticRouteDescriptorName is the name of the descriptor for static routes.
	StaticRouteDescriptorName = "vpp-static-route"

	// dependency labels
	routeOutInterfaceDep = "interface-exists"
)

// RouteDescriptor teaches KVScheduler how to configure Linux routes.
type RouteDescriptor struct {
	log       logging.Logger
	l3Handler vppcalls.RouteVppAPI
	scheduler scheduler.KVScheduler
}

// NewRouteDescriptor creates a new instance of the Route descriptor.
func NewRouteDescriptor(scheduler scheduler.KVScheduler,
	l3Handler vppcalls.RouteVppAPI, log logging.PluginLogger) *RouteDescriptor {

	return &RouteDescriptor{
		scheduler: scheduler,
		l3Handler: l3Handler,
		log:       log.NewLogger("route-descriptor"),
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter) with
// the KVScheduler.
func (d *RouteDescriptor) GetDescriptor() *adapter.StaticRouteDescriptor {
	return &adapter.StaticRouteDescriptor{
		Name: StaticRouteDescriptorName,
		KeySelector: func(key string) bool {
			return strings.HasPrefix(key, l3.VrfPrefix)
		},
		ValueTypeName:      proto.MessageName(&l3.StaticRoute{}),
		ValueComparator:    d.EquivalentRoutes,
		NBKeyPrefix:        l3.VrfPrefix,
		Add:                d.Add,
		Delete:             d.Delete,
		Modify:             d.Modify,
		IsRetriableFailure: d.IsRetriableFailure,
		Dependencies:       d.Dependencies,
		DerivedValues:      d.DerivedValues,
		Dump:               d.Dump,
		DumpDependencies:   []string{ifdescriptor.InterfaceDescriptorName},
	}
}

// EquivalentRoutes is case-insensitive comparison function for l3.LinuxStaticRoute.
func (d *RouteDescriptor) EquivalentRoutes(key string, oldRoute, newRoute *l3.StaticRoute) bool {
	return proto.Equal(oldRoute, newRoute)
}

// IsRetriableFailure returns <false> for errors related to invalid configuration.
func (d *RouteDescriptor) IsRetriableFailure(err error) bool {
	return false // nothing retriable
}

// Add adds Linux route.
func (d *RouteDescriptor) Add(key string, route *l3.StaticRoute) (metadata interface{}, err error) {
	err = d.l3Handler.VppAddRoute(route, route.GetOutgoingInterface())
	if err != nil {
		return nil, errors.Errorf("failed to add VPP route: %v", err)
	}

	return nil, nil //fmt.Errorf("not implemented")
}

// Delete removes Linux route.
func (d *RouteDescriptor) Delete(key string, route *l3.StaticRoute, metadata interface{}) error {
	return fmt.Errorf("not implemented")
}

// Modify is able to change route scope, metric and GW address.
func (d *RouteDescriptor) Modify(key string, oldRoute, newRoute *l3.StaticRoute, oldMetadata interface{}) (newMetadata interface{}, err error) {
	return nil, fmt.Errorf("not implemented")
}

// Dependencies lists dependencies for a Linux route.
func (d *RouteDescriptor) Dependencies(key string, route *l3.StaticRoute) []scheduler.Dependency {
	var dependencies []scheduler.Dependency
	// the outgoing interface must exist and be UP
	if route.OutgoingInterface != "" {
		dependencies = append(dependencies, scheduler.Dependency{
			Label: routeOutInterfaceDep,
			Key:   interfaces.InterfaceKey(route.OutgoingInterface),
		})
	}
	// GW must be routable
	/*gwAddr := net.ParseIP(getGwAddr(route))
	if gwAddr != nil && !gwAddr.IsUnspecified() {
		dependencies = append(dependencies, scheduler.Dependency{
			Label: routeGwReachabilityDep,
			AnyOf: func(key string) bool {
				dstAddr, ifName, err := l3.ParseStaticLinkLocalRouteKey(key)
				if err == nil && ifName == route.OutgoingInterface && dstAddr.Contains(gwAddr) {
					// GW address is neighbour as told by another link-local route
					return true
				}
				ifName, addr, err := ifmodel.ParseInterfaceAddressKey(key)
				if err == nil && ifName == route.OutgoingInterface && addr.Contains(gwAddr) {
					// GW address is inside the local network of the outgoing interface
					// as given by the assigned IP address
					return true
				}
				return false
			},
		})
	}*/
	return dependencies
}

// DerivedValues derives empty value under StaticLinkLocalRouteKey if route is link-local.
// It is used in dependencies for network reachability of a route gateway (see above).
func (d *RouteDescriptor) DerivedValues(key string, route *l3.StaticRoute) (derValues []scheduler.KeyValuePair) {
	/*if route.Scope == l3.LinuxStaticRoute_LINK {
		derValues = append(derValues, scheduler.KeyValuePair{
			Key:   l3.StaticLinkLocalRouteKey(route.DstNetwork, route.OutgoingInterface),
			Value: &prototypes.Empty{},
		})
	}*/
	return derValues
}

// Dump returns all routes associated with interfaces managed by this agent.
func (d *RouteDescriptor) Dump(correlate []adapter.StaticRouteKVWithMetadata) (
	dump []adapter.StaticRouteKVWithMetadata, err error) {

	// Retrieve VPP route configuration
	staticRoutes, err := d.l3Handler.DumpStaticRoutes()
	if err != nil {
		return nil, errors.Errorf("failed to dump VPP routes: %v", err)
	}

	for _, staticRoute := range staticRoutes {
		dump = append(dump, adapter.StaticRouteKVWithMetadata{
			Key:    l3.RouteKey(staticRoute.Route.VrfId, staticRoute.Route.DstNetwork, staticRoute.Route.NextHopAddr),
			Value:  staticRoute.Route,
			Origin: scheduler.FromSB,
		})
	}

	d.log.Debugf("Dumped %d Static Routes: %v", len(dump), dump)
	return dump, nil
}

// getGwAddr returns the GW address chosen in the given route, handling the cases
// when it is left undefined.
/*func getGwAddr(route *l3.StaticRoute) string {
	if route.GwAddr == "" {
		if ipv6, _ := addrs.IsIPv6(route.DstNetwork); ipv6 {
			return ipv6AddrAny
		}
		return ipv4AddrAny
	}
	return route.GwAddr
}
*/
