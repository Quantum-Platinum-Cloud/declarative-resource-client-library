// Copyright 2021 Google LLC. All Rights Reserved.
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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudresourcemanager/alpha/cloudresourcemanager_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/alpha"
)

// Server implements the gRPC interface for Folder.
type FolderServer struct{}

// ProtoToFolderStateEnum converts a FolderStateEnum enum from its proto representation.
func ProtoToCloudresourcemanagerAlphaFolderStateEnum(e alphapb.CloudresourcemanagerAlphaFolderStateEnum) *alpha.FolderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudresourcemanagerAlphaFolderStateEnum_name[int32(e)]; ok {
		e := alpha.FolderStateEnum(n[len("CloudresourcemanagerAlphaFolderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFolder converts a Folder resource from its proto representation.
func ProtoToFolder(p *alphapb.CloudresourcemanagerAlphaFolder) *alpha.Folder {
	obj := &alpha.Folder{
		Name:        dcl.StringOrNil(p.Name),
		Parent:      dcl.StringOrNil(p.Parent),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		State:       ProtoToCloudresourcemanagerAlphaFolderStateEnum(p.GetState()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:  dcl.StringOrNil(p.GetDeleteTime()),
		Etag:        dcl.StringOrNil(p.Etag),
	}
	return obj
}

// FolderStateEnumToProto converts a FolderStateEnum enum to its proto representation.
func CloudresourcemanagerAlphaFolderStateEnumToProto(e *alpha.FolderStateEnum) alphapb.CloudresourcemanagerAlphaFolderStateEnum {
	if e == nil {
		return alphapb.CloudresourcemanagerAlphaFolderStateEnum(0)
	}
	if v, ok := alphapb.CloudresourcemanagerAlphaFolderStateEnum_value["FolderStateEnum"+string(*e)]; ok {
		return alphapb.CloudresourcemanagerAlphaFolderStateEnum(v)
	}
	return alphapb.CloudresourcemanagerAlphaFolderStateEnum(0)
}

// FolderToProto converts a Folder resource to its proto representation.
func FolderToProto(resource *alpha.Folder) *alphapb.CloudresourcemanagerAlphaFolder {
	p := &alphapb.CloudresourcemanagerAlphaFolder{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Parent:      dcl.ValueOrEmptyString(resource.Parent),
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		State:       CloudresourcemanagerAlphaFolderStateEnumToProto(resource.State),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:  dcl.ValueOrEmptyString(resource.DeleteTime),
		Etag:        dcl.ValueOrEmptyString(resource.Etag),
	}

	return p
}

// ApplyFolder handles the gRPC request by passing it to the underlying Folder Apply() method.
func (s *FolderServer) applyFolder(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudresourcemanagerAlphaFolderRequest) (*alphapb.CloudresourcemanagerAlphaFolder, error) {
	p := ProtoToFolder(request.GetResource())
	res, err := c.ApplyFolder(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FolderToProto(res)
	return r, nil
}

// ApplyFolder handles the gRPC request by passing it to the underlying Folder Apply() method.
func (s *FolderServer) ApplyCloudresourcemanagerAlphaFolder(ctx context.Context, request *alphapb.ApplyCloudresourcemanagerAlphaFolderRequest) (*alphapb.CloudresourcemanagerAlphaFolder, error) {
	cl, err := createConfigFolder(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFolder(ctx, cl, request)
}

// DeleteFolder handles the gRPC request by passing it to the underlying Folder Delete() method.
func (s *FolderServer) DeleteCloudresourcemanagerAlphaFolder(ctx context.Context, request *alphapb.DeleteCloudresourcemanagerAlphaFolderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFolder(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFolder(ctx, ProtoToFolder(request.GetResource()))

}

// ListCloudresourcemanagerAlphaFolder handles the gRPC request by passing it to the underlying FolderList() method.
func (s *FolderServer) ListCloudresourcemanagerAlphaFolder(ctx context.Context, request *alphapb.ListCloudresourcemanagerAlphaFolderRequest) (*alphapb.ListCloudresourcemanagerAlphaFolderResponse, error) {
	cl, err := createConfigFolder(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFolder(ctx, request.Parent)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudresourcemanagerAlphaFolder
	for _, r := range resources.Items {
		rp := FolderToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListCloudresourcemanagerAlphaFolderResponse{Items: protos}, nil
}

func createConfigFolder(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
