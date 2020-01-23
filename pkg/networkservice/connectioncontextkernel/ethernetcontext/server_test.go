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

package ethernetcontext

import (
	"context"
	"testing"

	"github.com/ligato/vpp-agent/api/configurator"
	"github.com/ligato/vpp-agent/api/models/linux"
	"github.com/networkservicemesh/sdk/pkg/networkservice/core/next"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/networkservicemesh/sdk-vppagent/pkg/networkservice/vppagent"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/kernel"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connectioncontext"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice"
)

type testDumpConfiguratorClient struct {
}

func (t *testDumpConfiguratorClient) Get(ctx context.Context, in *configurator.GetRequest, opts ...grpc.CallOption) (*configurator.GetResponse, error) {
	panic("implement me")
}

func (t *testDumpConfiguratorClient) Update(ctx context.Context, in *configurator.UpdateRequest, opts ...grpc.CallOption) (*configurator.UpdateResponse, error) {
	panic("implement me")
}

func (t *testDumpConfiguratorClient) Delete(ctx context.Context, in *configurator.DeleteRequest, opts ...grpc.CallOption) (*configurator.DeleteResponse, error) {
	panic("implement me")
}

func (t *testDumpConfiguratorClient) Dump(ctx context.Context, in *configurator.DumpRequest, opts ...grpc.CallOption) (*configurator.DumpResponse, error) {
	return &configurator.DumpResponse{
		Dump: &configurator.Config{
			LinuxConfig: &linux.ConfigData{
				Interfaces: []*linux.Interface{
					{
						Name:        "DST-1",
						PhysAddress: "0a-1b-3c-4d-5e-6f",
					},
				},
			},
		},
	}, nil
}

func (t *testDumpConfiguratorClient) Notify(ctx context.Context, in *configurator.NotificationRequest, opts ...grpc.CallOption) (configurator.Configurator_NotifyClient, error) {
	panic("implement me")
}

func TestServerBasic(t *testing.T) {
	ctx := vppagent.WithConfig(context.Background())
	config := vppagent.Config(ctx)
	config.LinuxConfig = &linux.ConfigData{
		Interfaces: []*linux.Interface{
			{
				Name: "SRC-1",
			},
			{
				Name: "DST-1",
			},
		},
	}
	request := &networkservice.NetworkServiceRequest{
		Connection: &connection.Connection{
			Id: "1",
			Mechanism: &connection.Mechanism{
				Type: kernel.MECHANISM,
			},
			Context: &connectioncontext.ConnectionContext{
				IpContext: &connectioncontext.IPContext{
					DstIpAddr: "172.16.1.2",
				},
			},
		},
	}

	server := next.NewNetworkServiceServer(&setEthernetContextServer{
		client: &testDumpConfiguratorClient{},
	})
	cc, err := server.Request(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.NotNil(t, request.Connection.Context.EthernetContext)
	assert.Equal(t, request.Connection.Context.EthernetContext.DstMac, "0a-1b-3c-4d-5e-6f")
}
