//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package l3plugin

import (
	"go.ligato.io/cn-infra/v2/health/statuscheck"
	"go.ligato.io/cn-infra/v2/logging"

	"go.ligato.io/vpp-agent/v3/plugins/govppmux"
	"go.ligato.io/vpp-agent/v3/plugins/kvscheduler"
	"go.ligato.io/vpp-agent/v3/plugins/netalloc"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin"
)

// DefaultPlugin is a default instance of IfPlugin.
var DefaultPlugin = *NewPlugin()

// NewPlugin creates a new Plugin with the provides Options
func NewPlugin(opts ...Option) *L3Plugin {
	p := &L3Plugin{}

	p.PluginName = "vpp-l3plugin"
	p.StatusCheck = &statuscheck.DefaultPlugin
	p.KVScheduler = &kvscheduler.DefaultPlugin
	p.VPP = &govppmux.DefaultPlugin
	p.IfPlugin = &ifplugin.DefaultPlugin
	p.AddrAlloc = &netalloc.DefaultPlugin

	for _, o := range opts {
		o(p)
	}

	if p.Log == nil {
		p.Log = logging.ForPlugin(p.String())
	}

	return p
}

// Option is a function that can be used in NewPlugin to customize Plugin.
type Option func(*L3Plugin)

// UseDeps returns Option that can inject custom dependencies.
func UseDeps(f func(*Deps)) Option {
	return func(p *L3Plugin) {
		f(&p.Deps)
	}
}
