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

// Package getmac provides networkservice chain elements for applying ethernet context destination mac
package getmac

import (
	"context"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ligato/vpp-agent/api/configurator"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connectioncontext"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice"
	"github.com/networkservicemesh/sdk/pkg/networkservice/core/next"
	"github.com/sirupsen/logrus"

	"github.com/networkservicemesh/sdk-vppagent/pkg/networkservice/vppagent"
)

// NewServer creates a NetworkServiceServer chain element to set the EthernetContext for Kernel connection request
func NewServer(сс *grpc.ClientConn) networkservice.NetworkServiceServer {
	return &getMacServer{
		client: configurator.NewConfiguratorClient(сс),
	}
}

type getMacServer struct {
	client configurator.ConfiguratorClient
}

func (s *getMacServer) Request(ctx context.Context, request *networkservice.NetworkServiceRequest) (*connection.Connection, error) {
	conn, err := next.Server(ctx).Request(ctx, request)
	if err == nil && request.GetConnection().GetContext().IsEthernetContextEmtpy() {
		request.GetConnection().GetContext().EthernetContext = s.createEthernetContext(conn)
	}
	return conn, err
}

func (s *getMacServer) Close(ctx context.Context, conn *connection.Connection) (*empty.Empty, error) {
	return next.Server(ctx).Close(ctx, conn)
}

func (s *getMacServer) createEthernetContext(conn *connection.Connection) *connectioncontext.EthernetContext {
	dump, err := s.client.Dump(context.Background(), &configurator.DumpRequest{})
	if err != nil {
		logrus.Errorf("An error during ConfiguratorClient.Dump, err: %v", err.Error())
	}
	iface := vppagent.DestinationLinuxInterface(dump.Dump, conn.Id)
	return &connectioncontext.EthernetContext{
		DstMac: iface.PhysAddress,
	}
}
