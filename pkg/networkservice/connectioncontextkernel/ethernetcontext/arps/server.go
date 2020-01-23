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

// Package arps provides networkservice chain elements for setting ethernet context specific arps
package arps

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ligato/vpp-agent/api/models/linux"
	"github.com/networkservicemesh/sdk/pkg/networkservice/core/next"

	"github.com/networkservicemesh/sdk-vppagent/pkg/networkservice/vppagent"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/kernel"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice"
)

type setKernelArpServer struct {
}

func (s *setKernelArpServer) Request(ctx context.Context, request *networkservice.NetworkServiceRequest) (*connection.Connection, error) {
	if mechanism := kernel.ToMechanism(request.GetConnection().GetMechanism()); mechanism != nil {
		config := vppagent.Config(ctx)
		if !request.Connection.GetContext().IsEthernetContextEmtpy() {
			config.LinuxConfig.ArpEntries = append(config.LinuxConfig.ArpEntries, &linux.ARPEntry{
				IpAddress: strings.Split(request.GetConnection().GetContext().IpContext.DstIpAddr, "/")[0],
				Interface: "SRC-" + request.Connection.Id,
				HwAddress: request.Connection.GetContext().EthernetContext.DstMac,
			})
		}
	}
	return next.Server(ctx).Request(ctx, request)
}

func (s *setKernelArpServer) Close(ctx context.Context, conn *connection.Connection) (*empty.Empty, error) {
	if mechanism := kernel.ToMechanism(conn.GetMechanism()); mechanism != nil {
		config := vppagent.Config(ctx)
		if !conn.GetContext().IsEthernetContextEmtpy() {
			config.LinuxConfig.ArpEntries = append(config.LinuxConfig.ArpEntries, &linux.ARPEntry{
				IpAddress: strings.Split(conn.GetContext().IpContext.DstIpAddr, "/")[0],
				Interface: "SRC-" + conn.Id,
				HwAddress: conn.GetContext().EthernetContext.DstMac,
			})
		}
	}
	return next.Server(ctx).Close(ctx, conn)
}

// NewServer creates a NetworkServiceServer chain element to set the ARP for SRC interface
func NewServer() networkservice.NetworkServiceServer {
	return &setKernelArpServer{}
}
