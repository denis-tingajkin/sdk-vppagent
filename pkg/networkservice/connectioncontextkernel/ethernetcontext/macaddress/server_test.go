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

package macaddress

import (
	"context"
	"testing"

	"github.com/ligato/vpp-agent/api/models/linux"
	"github.com/networkservicemesh/sdk/pkg/networkservice/core/next"
	"github.com/stretchr/testify/assert"

	"github.com/networkservicemesh/sdk-vppagent/pkg/networkservice/vppagent"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/kernel"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connectioncontext"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice"
)

func TestServerBasic(t *testing.T) {
	ctx := vppagent.WithConfig(context.Background())
	config := vppagent.Config(ctx)
	config.LinuxConfig = &linux.ConfigData{
		Interfaces: []*linux.Interface{
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
				EthernetContext: &connectioncontext.EthernetContext{
					DstMac: "0a-1b-3c-4d-5e-6f",
				},
			},
		},
	}
	server := next.NewNetworkServiceServer(NewServer())
	cc, err := server.Request(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.NotEqual(t, nil, cc)
	for _, iface := range config.LinuxConfig.Interfaces {
		if iface.Name == "DST-"+request.Connection.Id {
			assert.Equal(t, iface.PhysAddress, request.Connection.Context.EthernetContext.DstMac)
			return
		}
	}
	assert.FailNow(t, "interface DST-1 not found in vpp-config")
}
