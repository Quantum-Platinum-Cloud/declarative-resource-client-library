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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/clouddeploy/beta/clouddeploy_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/beta"
)

// DeliveryPipelineServer implements the gRPC interface for DeliveryPipeline.
type DeliveryPipelineServer struct{}

// ProtoToDeliveryPipelineSerialPipeline converts a DeliveryPipelineSerialPipeline object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipeline(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipeline) *beta.DeliveryPipelineSerialPipeline {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipeline{}
	for _, r := range p.GetStages() {
		obj.Stages = append(obj.Stages, *ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStages(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStages converts a DeliveryPipelineSerialPipelineStages object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStages(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStages) *beta.DeliveryPipelineSerialPipelineStages {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStages{
		TargetId: dcl.StringOrNil(p.GetTargetId()),
		Strategy: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy(p.GetStrategy()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategy converts a DeliveryPipelineSerialPipelineStagesStrategy object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy) *beta.DeliveryPipelineSerialPipelineStagesStrategy {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategy{
		Standard: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard(p.GetStandard()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandard converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard) *beta.DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyStandard{
		Verify: dcl.Bool(p.GetVerify()),
	}
	return obj
}

// ProtoToDeliveryPipelineCondition converts a DeliveryPipelineCondition object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineCondition(p *betapb.ClouddeployBetaDeliveryPipelineCondition) *beta.DeliveryPipelineCondition {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineCondition{
		PipelineReadyCondition:  ProtoToClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition(p.GetPipelineReadyCondition()),
		TargetsPresentCondition: ProtoToClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition(p.GetTargetsPresentCondition()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionPipelineReadyCondition converts a DeliveryPipelineConditionPipelineReadyCondition object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition(p *betapb.ClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition) *beta.DeliveryPipelineConditionPipelineReadyCondition {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineConditionPipelineReadyCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsPresentCondition converts a DeliveryPipelineConditionTargetsPresentCondition object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition(p *betapb.ClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition) *beta.DeliveryPipelineConditionTargetsPresentCondition {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineConditionTargetsPresentCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetMissingTargets() {
		obj.MissingTargets = append(obj.MissingTargets, r)
	}
	return obj
}

// ProtoToDeliveryPipeline converts a DeliveryPipeline resource from its proto representation.
func ProtoToDeliveryPipeline(p *betapb.ClouddeployBetaDeliveryPipeline) *beta.DeliveryPipeline {
	obj := &beta.DeliveryPipeline{
		Name:           dcl.StringOrNil(p.GetName()),
		Uid:            dcl.StringOrNil(p.GetUid()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		SerialPipeline: ProtoToClouddeployBetaDeliveryPipelineSerialPipeline(p.GetSerialPipeline()),
		Condition:      ProtoToClouddeployBetaDeliveryPipelineCondition(p.GetCondition()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
		Suspended:      dcl.Bool(p.GetSuspended()),
	}
	return obj
}

// DeliveryPipelineSerialPipelineToProto converts a DeliveryPipelineSerialPipeline object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineToProto(o *beta.DeliveryPipelineSerialPipeline) *betapb.ClouddeployBetaDeliveryPipelineSerialPipeline {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipeline{}
	sStages := make([]*betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStages, len(o.Stages))
	for i, r := range o.Stages {
		sStages[i] = ClouddeployBetaDeliveryPipelineSerialPipelineStagesToProto(&r)
	}
	p.SetStages(sStages)
	return p
}

// DeliveryPipelineSerialPipelineStagesToProto converts a DeliveryPipelineSerialPipelineStages object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesToProto(o *beta.DeliveryPipelineSerialPipelineStages) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStages {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStages{}
	p.SetTargetId(dcl.ValueOrEmptyString(o.TargetId))
	p.SetStrategy(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyToProto(o.Strategy))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyToProto converts a DeliveryPipelineSerialPipelineStagesStrategy object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategy) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy{}
	p.SetStandard(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o.Standard))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyStandard) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard{}
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	return p
}

// DeliveryPipelineConditionToProto converts a DeliveryPipelineCondition object to its proto representation.
func ClouddeployBetaDeliveryPipelineConditionToProto(o *beta.DeliveryPipelineCondition) *betapb.ClouddeployBetaDeliveryPipelineCondition {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineCondition{}
	p.SetPipelineReadyCondition(ClouddeployBetaDeliveryPipelineConditionPipelineReadyConditionToProto(o.PipelineReadyCondition))
	p.SetTargetsPresentCondition(ClouddeployBetaDeliveryPipelineConditionTargetsPresentConditionToProto(o.TargetsPresentCondition))
	return p
}

// DeliveryPipelineConditionPipelineReadyConditionToProto converts a DeliveryPipelineConditionPipelineReadyCondition object to its proto representation.
func ClouddeployBetaDeliveryPipelineConditionPipelineReadyConditionToProto(o *beta.DeliveryPipelineConditionPipelineReadyCondition) *betapb.ClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// DeliveryPipelineConditionTargetsPresentConditionToProto converts a DeliveryPipelineConditionTargetsPresentCondition object to its proto representation.
func ClouddeployBetaDeliveryPipelineConditionTargetsPresentConditionToProto(o *beta.DeliveryPipelineConditionTargetsPresentCondition) *betapb.ClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	sMissingTargets := make([]string, len(o.MissingTargets))
	for i, r := range o.MissingTargets {
		sMissingTargets[i] = r
	}
	p.SetMissingTargets(sMissingTargets)
	return p
}

// DeliveryPipelineToProto converts a DeliveryPipeline resource to its proto representation.
func DeliveryPipelineToProto(resource *beta.DeliveryPipeline) *betapb.ClouddeployBetaDeliveryPipeline {
	p := &betapb.ClouddeployBetaDeliveryPipeline{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetSerialPipeline(ClouddeployBetaDeliveryPipelineSerialPipelineToProto(resource.SerialPipeline))
	p.SetCondition(ClouddeployBetaDeliveryPipelineConditionToProto(resource.Condition))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSuspended(dcl.ValueOrEmptyBool(resource.Suspended))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) applyDeliveryPipeline(ctx context.Context, c *beta.Client, request *betapb.ApplyClouddeployBetaDeliveryPipelineRequest) (*betapb.ClouddeployBetaDeliveryPipeline, error) {
	p := ProtoToDeliveryPipeline(request.GetResource())
	res, err := c.ApplyDeliveryPipeline(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DeliveryPipelineToProto(res)
	return r, nil
}

// applyClouddeployBetaDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) ApplyClouddeployBetaDeliveryPipeline(ctx context.Context, request *betapb.ApplyClouddeployBetaDeliveryPipelineRequest) (*betapb.ClouddeployBetaDeliveryPipeline, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDeliveryPipeline(ctx, cl, request)
}

// DeleteDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Delete() method.
func (s *DeliveryPipelineServer) DeleteClouddeployBetaDeliveryPipeline(ctx context.Context, request *betapb.DeleteClouddeployBetaDeliveryPipelineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDeliveryPipeline(ctx, ProtoToDeliveryPipeline(request.GetResource()))

}

// ListClouddeployBetaDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipelineList() method.
func (s *DeliveryPipelineServer) ListClouddeployBetaDeliveryPipeline(ctx context.Context, request *betapb.ListClouddeployBetaDeliveryPipelineRequest) (*betapb.ListClouddeployBetaDeliveryPipelineResponse, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDeliveryPipeline(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ClouddeployBetaDeliveryPipeline
	for _, r := range resources.Items {
		rp := DeliveryPipelineToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListClouddeployBetaDeliveryPipelineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDeliveryPipeline(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
