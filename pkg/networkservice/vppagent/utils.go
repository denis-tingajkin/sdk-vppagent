// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
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

package vppagent

import (
	"fmt"

	"github.com/ligato/vpp-agent/api/configurator"
	"github.com/ligato/vpp-agent/api/models/linux"
)

// SourceLinuxInterface returns source linux interface if exist or nil
func SourceLinuxInterface(config *configurator.Config, id string) *linux.Interface {
	srcName := fmt.Sprintf("SRC-%v", id)
	return findLinuxInterfaceByName(config, srcName)
}

// DestinationLinuxInterface returns destination linux interface if exist or nil
func DestinationLinuxInterface(config *configurator.Config, id string) *linux.Interface {
	dstName := fmt.Sprintf("DST-%v", id)
	return findLinuxInterfaceByName(config, dstName)
}

func findLinuxInterfaceByName(config *configurator.Config, name string) *linux.Interface {
	if config == nil {
		return nil
	}
	if config.LinuxConfig == nil {
		return nil
	}
	if config.LinuxConfig.Interfaces == nil {
		return nil
	}
	for _, iface := range config.LinuxConfig.Interfaces {
		if iface.Name == name {
			return iface
		}
	}
	return nil
}
