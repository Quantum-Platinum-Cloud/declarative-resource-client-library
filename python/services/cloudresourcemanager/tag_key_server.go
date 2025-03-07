// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	cloudresourcemanagerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudresourcemanager/cloudresourcemanager_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager"
)

// TagKeyServer implements the gRPC interface for TagKey.
type TagKeyServer struct{}

// ProtoToTagKeyPurposeEnum converts a TagKeyPurposeEnum enum from its proto representation.
func ProtoToCloudresourcemanagerTagKeyPurposeEnum(e cloudresourcemanagerpb.CloudresourcemanagerTagKeyPurposeEnum) *cloudresourcemanager.TagKeyPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudresourcemanagerpb.CloudresourcemanagerTagKeyPurposeEnum_name[int32(e)]; ok {
		e := cloudresourcemanager.TagKeyPurposeEnum(n[len("CloudresourcemanagerTagKeyPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToTagKey converts a TagKey resource from its proto representation.
func ProtoToTagKey(p *cloudresourcemanagerpb.CloudresourcemanagerTagKey) *cloudresourcemanager.TagKey {
	obj := &cloudresourcemanager.TagKey{
		Name:           dcl.StringOrNil(p.GetName()),
		Parent:         dcl.StringOrNil(p.GetParent()),
		ShortName:      dcl.StringOrNil(p.GetShortName()),
		NamespacedName: dcl.StringOrNil(p.GetNamespacedName()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		Purpose:        ProtoToCloudresourcemanagerTagKeyPurposeEnum(p.GetPurpose()),
	}
	return obj
}

// TagKeyPurposeEnumToProto converts a TagKeyPurposeEnum enum to its proto representation.
func CloudresourcemanagerTagKeyPurposeEnumToProto(e *cloudresourcemanager.TagKeyPurposeEnum) cloudresourcemanagerpb.CloudresourcemanagerTagKeyPurposeEnum {
	if e == nil {
		return cloudresourcemanagerpb.CloudresourcemanagerTagKeyPurposeEnum(0)
	}
	if v, ok := cloudresourcemanagerpb.CloudresourcemanagerTagKeyPurposeEnum_value["TagKeyPurposeEnum"+string(*e)]; ok {
		return cloudresourcemanagerpb.CloudresourcemanagerTagKeyPurposeEnum(v)
	}
	return cloudresourcemanagerpb.CloudresourcemanagerTagKeyPurposeEnum(0)
}

// TagKeyToProto converts a TagKey resource to its proto representation.
func TagKeyToProto(resource *cloudresourcemanager.TagKey) *cloudresourcemanagerpb.CloudresourcemanagerTagKey {
	p := &cloudresourcemanagerpb.CloudresourcemanagerTagKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetShortName(dcl.ValueOrEmptyString(resource.ShortName))
	p.SetNamespacedName(dcl.ValueOrEmptyString(resource.NamespacedName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetPurpose(CloudresourcemanagerTagKeyPurposeEnumToProto(resource.Purpose))
	mPurposeData := make(map[string]string, len(resource.PurposeData))
	for k, r := range resource.PurposeData {
		mPurposeData[k] = r
	}
	p.SetPurposeData(mPurposeData)

	return p
}

// applyTagKey handles the gRPC request by passing it to the underlying TagKey Apply() method.
func (s *TagKeyServer) applyTagKey(ctx context.Context, c *cloudresourcemanager.Client, request *cloudresourcemanagerpb.ApplyCloudresourcemanagerTagKeyRequest) (*cloudresourcemanagerpb.CloudresourcemanagerTagKey, error) {
	p := ProtoToTagKey(request.GetResource())
	res, err := c.ApplyTagKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TagKeyToProto(res)
	return r, nil
}

// applyCloudresourcemanagerTagKey handles the gRPC request by passing it to the underlying TagKey Apply() method.
func (s *TagKeyServer) ApplyCloudresourcemanagerTagKey(ctx context.Context, request *cloudresourcemanagerpb.ApplyCloudresourcemanagerTagKeyRequest) (*cloudresourcemanagerpb.CloudresourcemanagerTagKey, error) {
	cl, err := createConfigTagKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTagKey(ctx, cl, request)
}

// DeleteTagKey handles the gRPC request by passing it to the underlying TagKey Delete() method.
func (s *TagKeyServer) DeleteCloudresourcemanagerTagKey(ctx context.Context, request *cloudresourcemanagerpb.DeleteCloudresourcemanagerTagKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTagKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTagKey(ctx, ProtoToTagKey(request.GetResource()))

}

// ListCloudresourcemanagerTagKey is a no-op method because TagKey has no list method.
func (s *TagKeyServer) ListCloudresourcemanagerTagKey(_ context.Context, _ *cloudresourcemanagerpb.ListCloudresourcemanagerTagKeyRequest) (*cloudresourcemanagerpb.ListCloudresourcemanagerTagKeyResponse, error) {
	return nil, nil
}

func createConfigTagKey(ctx context.Context, service_account_file string) (*cloudresourcemanager.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudresourcemanager.NewClient(conf), nil
}
