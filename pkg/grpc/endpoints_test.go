/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//nolint:nosnakecase
package grpc

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/dapr/dapr/pkg/config"
	runtimev1pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
)

func TestEndpointCompleteness(t *testing.T) {
	// Get the list of endpoints in the runtime
	runtimeEndpoints := []string{}
	prefix := "/" + runtimev1pb.Dapr_ServiceDesc.ServiceName + "/"
	for _, m := range runtimev1pb.Dapr_ServiceDesc.Methods {
		runtimeEndpoints = append(runtimeEndpoints, prefix+m.MethodName)
	}
	for _, m := range runtimev1pb.Dapr_ServiceDesc.Streams {
		runtimeEndpoints = append(runtimeEndpoints, prefix+m.StreamName)
	}
	sort.Strings(runtimeEndpoints)

	// Get the list of endpoints in this package (regardless of group)
	packageEndpoints := []string{}
	for _, g := range endpoints {
		packageEndpoints = append(packageEndpoints, g...)
	}
	sort.Strings(packageEndpoints)

	assert.Equal(t, runtimeEndpoints, packageEndpoints, "the list of endpoints defined in this package does not match the endpoints defined in the %s gRPC service", runtimev1pb.Dapr_ServiceDesc.ServiceName)
}

func TestSetAPIEndpointsMiddlewareUnary(t *testing.T) {
	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	}

	t.Run("state.v1 endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "state",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["state.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "state.v1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("state.v1alpha1 endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "state",
				Version:  "v1alpha1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["state.v1alpha1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "state.v1alpha1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("publish endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "publish",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["publish.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "publish.v1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("actors endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "actors",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["actors.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "actors.v1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("bindings endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "bindings",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["bindings.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "bindings.v1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("secrets endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "secrets",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["secrets.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "secrets.v1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("metadata endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "metadata",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["metadata.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "metadata.v1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("shutdown endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "shutdown",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["shutdown.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "shutdown.v1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("invoke endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "invoke",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, e := range endpoints["invoke.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for k, v := range endpoints {
			if k != "invoke.v1" {
				for _, e := range v {
					_, err := f(nil, nil, &grpc.UnaryServerInfo{
						FullMethod: e,
					}, h)
					assert.Error(t, err)
				}
			}
		}
	})

	t.Run("no rules, all endpoints are allowed", func(t *testing.T) {
		f := setAPIEndpointsMiddlewareUnary(nil)

		for _, e := range endpoints["invoke.v1"] {
			_, err := f(nil, nil, &grpc.UnaryServerInfo{
				FullMethod: e,
			}, h)
			assert.NoError(t, err)
		}

		for _, v := range endpoints {
			for _, e := range v {
				_, err := f(nil, nil, &grpc.UnaryServerInfo{
					FullMethod: e,
				}, h)
				assert.NoError(t, err)
			}
		}
	})

	t.Run("protocol mismatch, rule not applied", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "state",
				Version:  "v1",
				Protocol: "http",
			},
		}

		f := setAPIEndpointsMiddlewareUnary(a)

		for _, v := range endpoints {
			for _, e := range v {
				_, err := f(nil, nil, &grpc.UnaryServerInfo{
					FullMethod: e,
				}, h)
				assert.NoError(t, err)
			}
		}
	})
}
