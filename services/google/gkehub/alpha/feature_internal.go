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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Feature) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ResourceState) {
		if err := r.ResourceState.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Spec) {
		if err := r.Spec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureResourceState) validate() error {
	return nil
}
func (r *FeatureSpec) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Multiclusteringress) {
		if err := r.Multiclusteringress.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Cloudauditlogging) {
		if err := r.Cloudauditlogging.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureSpecMulticlusteringress) validate() error {
	if err := dcl.Required(r, "configMembership"); err != nil {
		return err
	}
	return nil
}
func (r *FeatureSpecCloudauditlogging) validate() error {
	return nil
}
func (r *FeatureState) validate() error {
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Servicemesh) {
		if err := r.Servicemesh.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureStateState) validate() error {
	return nil
}
func (r *FeatureStateServicemesh) validate() error {
	return nil
}
func (r *FeatureStateServicemeshAnalysisMessages) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MessageBase) {
		if err := r.MessageBase.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureStateServicemeshAnalysisMessagesMessageBase) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Type) {
		if err := r.Type.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureStateServicemeshAnalysisMessagesMessageBaseType) validate() error {
	return nil
}
func (r *Feature) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://gkehub.googleapis.com/v1alpha/", params)
}

// featureApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type featureApiOperation interface {
	do(context.Context, *Feature, *Client) error
}

// newUpdateFeatureUpdateFeatureRequest creates a request for an
// Feature resource's UpdateFeature update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFeatureUpdateFeatureRequest(ctx context.Context, f *Feature, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandFeatureSpec(c, f.Spec, res); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["spec"] = v
	}
	return req, nil
}

// marshalUpdateFeatureUpdateFeatureRequest converts the update into
// the final JSON request body.
func marshalUpdateFeatureUpdateFeatureRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFeatureUpdateFeatureOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listFeatureRaw(ctx context.Context, r *Feature, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FeatureMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listFeatureOperation struct {
	Resources []map[string]interface{} `json:"resources"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listFeature(ctx context.Context, r *Feature, pageToken string, pageSize int32) ([]*Feature, string, error) {
	b, err := c.listFeatureRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFeatureOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Feature
	for _, v := range m.Resources {
		res, err := unmarshalMapFeature(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFeature(ctx context.Context, f func(*Feature) bool, resources []*Feature) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFeature(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteFeatureOperation struct{}

func (op *deleteFeatureOperation) do(ctx context.Context, r *Feature, c *Client) error {
	r, err := c.GetFeature(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Feature not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetFeature checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetFeature(ctx, r)
		if dcl.IsNotFound(err) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFeatureOperation struct {
	response map[string]interface{}
}

func (op *createFeatureOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFeatureOperation) do(ctx context.Context, r *Feature, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetFeature(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFeatureRaw(ctx context.Context, r *Feature) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) featureDiffsForRawDesired(ctx context.Context, rawDesired *Feature, opts ...dcl.ApplyOption) (initial, desired *Feature, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Feature
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Feature); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Feature, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFeature(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Feature resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Feature resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Feature resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFeatureDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Feature: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Feature: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractFeatureFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFeatureInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Feature: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFeatureDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Feature: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFeature(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFeatureInitialState(rawInitial, rawDesired *Feature) (*Feature, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFeatureDesiredState(rawDesired, rawInitial *Feature, opts ...dcl.ApplyOption) (*Feature, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ResourceState = canonicalizeFeatureResourceState(rawDesired.ResourceState, nil, opts...)
		rawDesired.Spec = canonicalizeFeatureSpec(rawDesired.Spec, nil, opts...)
		rawDesired.State = canonicalizeFeatureState(rawDesired.State, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Feature{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.Spec = canonicalizeFeatureSpec(rawDesired.Spec, rawInitial.Spec, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeFeatureNewState(c *Client, rawNew, rawDesired *Feature) (*Feature, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceState) && dcl.IsEmptyValueIndirect(rawDesired.ResourceState) {
		rawNew.ResourceState = rawDesired.ResourceState
	} else {
		rawNew.ResourceState = canonicalizeNewFeatureResourceState(c, rawDesired.ResourceState, rawNew.ResourceState)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Spec) && dcl.IsEmptyValueIndirect(rawDesired.Spec) {
		rawNew.Spec = rawDesired.Spec
	} else {
		rawNew.Spec = canonicalizeNewFeatureSpec(c, rawDesired.Spec, rawNew.Spec)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
		rawNew.State = canonicalizeNewFeatureState(c, rawDesired.State, rawNew.State)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeFeatureResourceState(des, initial *FeatureResourceState, opts ...dcl.ApplyOption) *FeatureResourceState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureResourceState{}

	return cDes
}

func canonicalizeFeatureResourceStateSlice(des, initial []FeatureResourceState, opts ...dcl.ApplyOption) []FeatureResourceState {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureResourceState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureResourceState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureResourceState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureResourceState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureResourceState(c *Client, des, nw *FeatureResourceState) *FeatureResourceState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureResourceState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.HasResources, nw.HasResources) {
		nw.HasResources = des.HasResources
	}

	return nw
}

func canonicalizeNewFeatureResourceStateSet(c *Client, des, nw []FeatureResourceState) []FeatureResourceState {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureResourceState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureResourceStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureResourceStateSlice(c *Client, des, nw []FeatureResourceState) []FeatureResourceState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureResourceState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureResourceState(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpec(des, initial *FeatureSpec, opts ...dcl.ApplyOption) *FeatureSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpec{}

	cDes.Multiclusteringress = canonicalizeFeatureSpecMulticlusteringress(des.Multiclusteringress, initial.Multiclusteringress, opts...)
	cDes.Cloudauditlogging = canonicalizeFeatureSpecCloudauditlogging(des.Cloudauditlogging, initial.Cloudauditlogging, opts...)

	return cDes
}

func canonicalizeFeatureSpecSlice(des, initial []FeatureSpec, opts ...dcl.ApplyOption) []FeatureSpec {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpec, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpec(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpec, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpec(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpec(c *Client, des, nw *FeatureSpec) *FeatureSpec {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpec while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Multiclusteringress = canonicalizeNewFeatureSpecMulticlusteringress(c, des.Multiclusteringress, nw.Multiclusteringress)
	nw.Cloudauditlogging = canonicalizeNewFeatureSpecCloudauditlogging(c, des.Cloudauditlogging, nw.Cloudauditlogging)

	return nw
}

func canonicalizeNewFeatureSpecSet(c *Client, des, nw []FeatureSpec) []FeatureSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureSpecSlice(c *Client, des, nw []FeatureSpec) []FeatureSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpec(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecMulticlusteringress(des, initial *FeatureSpecMulticlusteringress, opts ...dcl.ApplyOption) *FeatureSpecMulticlusteringress {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpecMulticlusteringress{}

	if dcl.IsZeroValue(des.ConfigMembership) || (dcl.IsEmptyValueIndirect(des.ConfigMembership) && dcl.IsEmptyValueIndirect(initial.ConfigMembership)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ConfigMembership = initial.ConfigMembership
	} else {
		cDes.ConfigMembership = des.ConfigMembership
	}

	return cDes
}

func canonicalizeFeatureSpecMulticlusteringressSlice(des, initial []FeatureSpecMulticlusteringress, opts ...dcl.ApplyOption) []FeatureSpecMulticlusteringress {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpecMulticlusteringress, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpecMulticlusteringress(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpecMulticlusteringress, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpecMulticlusteringress(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpecMulticlusteringress(c *Client, des, nw *FeatureSpecMulticlusteringress) *FeatureSpecMulticlusteringress {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpecMulticlusteringress while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFeatureSpecMulticlusteringressSet(c *Client, des, nw []FeatureSpecMulticlusteringress) []FeatureSpecMulticlusteringress {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpecMulticlusteringress
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureSpecMulticlusteringressNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureSpecMulticlusteringressSlice(c *Client, des, nw []FeatureSpecMulticlusteringress) []FeatureSpecMulticlusteringress {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpecMulticlusteringress
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecMulticlusteringress(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecCloudauditlogging(des, initial *FeatureSpecCloudauditlogging, opts ...dcl.ApplyOption) *FeatureSpecCloudauditlogging {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpecCloudauditlogging{}

	if dcl.StringArrayCanonicalize(des.AllowlistedServiceAccounts, initial.AllowlistedServiceAccounts) {
		cDes.AllowlistedServiceAccounts = initial.AllowlistedServiceAccounts
	} else {
		cDes.AllowlistedServiceAccounts = des.AllowlistedServiceAccounts
	}

	return cDes
}

func canonicalizeFeatureSpecCloudauditloggingSlice(des, initial []FeatureSpecCloudauditlogging, opts ...dcl.ApplyOption) []FeatureSpecCloudauditlogging {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpecCloudauditlogging, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpecCloudauditlogging(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpecCloudauditlogging, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpecCloudauditlogging(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpecCloudauditlogging(c *Client, des, nw *FeatureSpecCloudauditlogging) *FeatureSpecCloudauditlogging {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpecCloudauditlogging while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.AllowlistedServiceAccounts, nw.AllowlistedServiceAccounts) {
		nw.AllowlistedServiceAccounts = des.AllowlistedServiceAccounts
	}

	return nw
}

func canonicalizeNewFeatureSpecCloudauditloggingSet(c *Client, des, nw []FeatureSpecCloudauditlogging) []FeatureSpecCloudauditlogging {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpecCloudauditlogging
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureSpecCloudauditloggingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureSpecCloudauditloggingSlice(c *Client, des, nw []FeatureSpecCloudauditlogging) []FeatureSpecCloudauditlogging {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpecCloudauditlogging
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecCloudauditlogging(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureState(des, initial *FeatureState, opts ...dcl.ApplyOption) *FeatureState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureState{}

	return cDes
}

func canonicalizeFeatureStateSlice(des, initial []FeatureState, opts ...dcl.ApplyOption) []FeatureState {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureState(c *Client, des, nw *FeatureState) *FeatureState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.State = canonicalizeNewFeatureStateState(c, des.State, nw.State)
	nw.Servicemesh = canonicalizeNewFeatureStateServicemesh(c, des.Servicemesh, nw.Servicemesh)

	return nw
}

func canonicalizeNewFeatureStateSet(c *Client, des, nw []FeatureState) []FeatureState {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureStateSlice(c *Client, des, nw []FeatureState) []FeatureState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureState(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureStateState(des, initial *FeatureStateState, opts ...dcl.ApplyOption) *FeatureStateState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureStateState{}

	return cDes
}

func canonicalizeFeatureStateStateSlice(des, initial []FeatureStateState, opts ...dcl.ApplyOption) []FeatureStateState {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureStateState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureStateState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureStateState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureStateState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureStateState(c *Client, des, nw *FeatureStateState) *FeatureStateState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureStateState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringCanonicalize(des.UpdateTime, nw.UpdateTime) {
		nw.UpdateTime = des.UpdateTime
	}

	return nw
}

func canonicalizeNewFeatureStateStateSet(c *Client, des, nw []FeatureStateState) []FeatureStateState {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureStateState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureStateStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureStateStateSlice(c *Client, des, nw []FeatureStateState) []FeatureStateState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureStateState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateState(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureStateServicemesh(des, initial *FeatureStateServicemesh, opts ...dcl.ApplyOption) *FeatureStateServicemesh {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureStateServicemesh{}

	return cDes
}

func canonicalizeFeatureStateServicemeshSlice(des, initial []FeatureStateServicemesh, opts ...dcl.ApplyOption) []FeatureStateServicemesh {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureStateServicemesh, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureStateServicemesh(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureStateServicemesh, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureStateServicemesh(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureStateServicemesh(c *Client, des, nw *FeatureStateServicemesh) *FeatureStateServicemesh {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureStateServicemesh while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.AnalysisMessages = canonicalizeNewFeatureStateServicemeshAnalysisMessagesSlice(c, des.AnalysisMessages, nw.AnalysisMessages)

	return nw
}

func canonicalizeNewFeatureStateServicemeshSet(c *Client, des, nw []FeatureStateServicemesh) []FeatureStateServicemesh {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureStateServicemesh
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureStateServicemeshNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureStateServicemeshSlice(c *Client, des, nw []FeatureStateServicemesh) []FeatureStateServicemesh {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureStateServicemesh
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateServicemesh(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureStateServicemeshAnalysisMessages(des, initial *FeatureStateServicemeshAnalysisMessages, opts ...dcl.ApplyOption) *FeatureStateServicemeshAnalysisMessages {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureStateServicemeshAnalysisMessages{}

	return cDes
}

func canonicalizeFeatureStateServicemeshAnalysisMessagesSlice(des, initial []FeatureStateServicemeshAnalysisMessages, opts ...dcl.ApplyOption) []FeatureStateServicemeshAnalysisMessages {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureStateServicemeshAnalysisMessages, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureStateServicemeshAnalysisMessages(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureStateServicemeshAnalysisMessages, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureStateServicemeshAnalysisMessages(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureStateServicemeshAnalysisMessages(c *Client, des, nw *FeatureStateServicemeshAnalysisMessages) *FeatureStateServicemeshAnalysisMessages {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureStateServicemeshAnalysisMessages while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MessageBase = canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBase(c, des.MessageBase, nw.MessageBase)
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringArrayCanonicalize(des.ResourcePaths, nw.ResourcePaths) {
		nw.ResourcePaths = des.ResourcePaths
	}

	return nw
}

func canonicalizeNewFeatureStateServicemeshAnalysisMessagesSet(c *Client, des, nw []FeatureStateServicemeshAnalysisMessages) []FeatureStateServicemeshAnalysisMessages {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureStateServicemeshAnalysisMessages
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureStateServicemeshAnalysisMessagesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureStateServicemeshAnalysisMessagesSlice(c *Client, des, nw []FeatureStateServicemeshAnalysisMessages) []FeatureStateServicemeshAnalysisMessages {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureStateServicemeshAnalysisMessages
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateServicemeshAnalysisMessages(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureStateServicemeshAnalysisMessagesMessageBase(des, initial *FeatureStateServicemeshAnalysisMessagesMessageBase, opts ...dcl.ApplyOption) *FeatureStateServicemeshAnalysisMessagesMessageBase {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureStateServicemeshAnalysisMessagesMessageBase{}

	return cDes
}

func canonicalizeFeatureStateServicemeshAnalysisMessagesMessageBaseSlice(des, initial []FeatureStateServicemeshAnalysisMessagesMessageBase, opts ...dcl.ApplyOption) []FeatureStateServicemeshAnalysisMessagesMessageBase {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureStateServicemeshAnalysisMessagesMessageBase, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureStateServicemeshAnalysisMessagesMessageBase(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureStateServicemeshAnalysisMessagesMessageBase, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureStateServicemeshAnalysisMessagesMessageBase(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBase(c *Client, des, nw *FeatureStateServicemeshAnalysisMessagesMessageBase) *FeatureStateServicemeshAnalysisMessagesMessageBase {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureStateServicemeshAnalysisMessagesMessageBase while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Type = canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBaseType(c, des.Type, nw.Type)
	if dcl.StringCanonicalize(des.DocumentationUrl, nw.DocumentationUrl) {
		nw.DocumentationUrl = des.DocumentationUrl
	}

	return nw
}

func canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBaseSet(c *Client, des, nw []FeatureStateServicemeshAnalysisMessagesMessageBase) []FeatureStateServicemeshAnalysisMessagesMessageBase {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureStateServicemeshAnalysisMessagesMessageBase
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureStateServicemeshAnalysisMessagesMessageBaseNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBaseSlice(c *Client, des, nw []FeatureStateServicemeshAnalysisMessagesMessageBase) []FeatureStateServicemeshAnalysisMessagesMessageBase {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureStateServicemeshAnalysisMessagesMessageBase
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBase(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureStateServicemeshAnalysisMessagesMessageBaseType(des, initial *FeatureStateServicemeshAnalysisMessagesMessageBaseType, opts ...dcl.ApplyOption) *FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureStateServicemeshAnalysisMessagesMessageBaseType{}

	return cDes
}

func canonicalizeFeatureStateServicemeshAnalysisMessagesMessageBaseTypeSlice(des, initial []FeatureStateServicemeshAnalysisMessagesMessageBaseType, opts ...dcl.ApplyOption) []FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureStateServicemeshAnalysisMessagesMessageBaseType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureStateServicemeshAnalysisMessagesMessageBaseType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureStateServicemeshAnalysisMessagesMessageBaseType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureStateServicemeshAnalysisMessagesMessageBaseType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBaseType(c *Client, des, nw *FeatureStateServicemeshAnalysisMessagesMessageBaseType) *FeatureStateServicemeshAnalysisMessagesMessageBaseType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureStateServicemeshAnalysisMessagesMessageBaseType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.DisplayName, nw.DisplayName) {
		nw.DisplayName = des.DisplayName
	}
	if dcl.StringCanonicalize(des.Code, nw.Code) {
		nw.Code = des.Code
	}

	return nw
}

func canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBaseTypeSet(c *Client, des, nw []FeatureStateServicemeshAnalysisMessagesMessageBaseType) []FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureStateServicemeshAnalysisMessagesMessageBaseType
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFeatureStateServicemeshAnalysisMessagesMessageBaseTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBaseTypeSlice(c *Client, des, nw []FeatureStateServicemeshAnalysisMessagesMessageBaseType) []FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureStateServicemeshAnalysisMessagesMessageBaseType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateServicemeshAnalysisMessagesMessageBaseType(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffFeature(c *Client, desired, actual *Feature, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceState, actual.ResourceState, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureResourceStateNewStyle, EmptyObject: EmptyFeatureResourceState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Spec, actual.Spec, dcl.DiffInfo{ObjectFunction: compareFeatureSpecNewStyle, EmptyObject: EmptyFeatureSpec, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Spec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureStateNewStyle, EmptyObject: EmptyFeatureState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareFeatureResourceStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureResourceState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureResourceState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureResourceState or *FeatureResourceState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureResourceState)
	if !ok {
		actualNotPointer, ok := a.(FeatureResourceState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureResourceState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HasResources, actual.HasResources, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HasResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpec)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpec or *FeatureSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpec)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Multiclusteringress, actual.Multiclusteringress, dcl.DiffInfo{ObjectFunction: compareFeatureSpecMulticlusteringressNewStyle, EmptyObject: EmptyFeatureSpecMulticlusteringress, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Multiclusteringress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Cloudauditlogging, actual.Cloudauditlogging, dcl.DiffInfo{ObjectFunction: compareFeatureSpecCloudauditloggingNewStyle, EmptyObject: EmptyFeatureSpecCloudauditlogging, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Cloudauditlogging")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecMulticlusteringressNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpecMulticlusteringress)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpecMulticlusteringress)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecMulticlusteringress or *FeatureSpecMulticlusteringress", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpecMulticlusteringress)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpecMulticlusteringress)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecMulticlusteringress", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigMembership, actual.ConfigMembership, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("ConfigMembership")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecCloudauditloggingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpecCloudauditlogging)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpecCloudauditlogging)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecCloudauditlogging or *FeatureSpecCloudauditlogging", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpecCloudauditlogging)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpecCloudauditlogging)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecCloudauditlogging", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowlistedServiceAccounts, actual.AllowlistedServiceAccounts, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("AllowlistedServiceAccounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureState or *FeatureState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureState)
	if !ok {
		actualNotPointer, ok := a.(FeatureState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureStateStateNewStyle, EmptyObject: EmptyFeatureStateState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Servicemesh, actual.Servicemesh, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureStateServicemeshNewStyle, EmptyObject: EmptyFeatureStateServicemesh, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Servicemesh")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureStateState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureStateState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateState or *FeatureStateState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureStateState)
	if !ok {
		actualNotPointer, ok := a.(FeatureStateState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateServicemeshNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureStateServicemesh)
	if !ok {
		desiredNotPointer, ok := d.(FeatureStateServicemesh)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateServicemesh or *FeatureStateServicemesh", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureStateServicemesh)
	if !ok {
		actualNotPointer, ok := a.(FeatureStateServicemesh)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateServicemesh", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AnalysisMessages, actual.AnalysisMessages, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureStateServicemeshAnalysisMessagesNewStyle, EmptyObject: EmptyFeatureStateServicemeshAnalysisMessages, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AnalysisMessages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateServicemeshAnalysisMessagesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureStateServicemeshAnalysisMessages)
	if !ok {
		desiredNotPointer, ok := d.(FeatureStateServicemeshAnalysisMessages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateServicemeshAnalysisMessages or *FeatureStateServicemeshAnalysisMessages", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureStateServicemeshAnalysisMessages)
	if !ok {
		actualNotPointer, ok := a.(FeatureStateServicemeshAnalysisMessages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateServicemeshAnalysisMessages", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MessageBase, actual.MessageBase, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureStateServicemeshAnalysisMessagesMessageBaseNewStyle, EmptyObject: EmptyFeatureStateServicemeshAnalysisMessagesMessageBase, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MessageBase")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourcePaths, actual.ResourcePaths, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourcePaths")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateServicemeshAnalysisMessagesMessageBaseNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureStateServicemeshAnalysisMessagesMessageBase)
	if !ok {
		desiredNotPointer, ok := d.(FeatureStateServicemeshAnalysisMessagesMessageBase)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateServicemeshAnalysisMessagesMessageBase or *FeatureStateServicemeshAnalysisMessagesMessageBase", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureStateServicemeshAnalysisMessagesMessageBase)
	if !ok {
		actualNotPointer, ok := a.(FeatureStateServicemeshAnalysisMessagesMessageBase)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateServicemeshAnalysisMessagesMessageBase", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureStateServicemeshAnalysisMessagesMessageBaseTypeNewStyle, EmptyObject: EmptyFeatureStateServicemeshAnalysisMessagesMessageBaseType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Level, actual.Level, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Level")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DocumentationUrl, actual.DocumentationUrl, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DocumentationUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateServicemeshAnalysisMessagesMessageBaseTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureStateServicemeshAnalysisMessagesMessageBaseType)
	if !ok {
		desiredNotPointer, ok := d.(FeatureStateServicemeshAnalysisMessagesMessageBaseType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateServicemeshAnalysisMessagesMessageBaseType or *FeatureStateServicemeshAnalysisMessagesMessageBaseType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureStateServicemeshAnalysisMessagesMessageBaseType)
	if !ok {
		actualNotPointer, ok := a.(FeatureStateServicemeshAnalysisMessagesMessageBaseType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateServicemeshAnalysisMessagesMessageBaseType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Feature) urlNormalized() *Feature {
	normalized := dcl.Copy(*r).(Feature)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Feature) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateFeature" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Feature resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Feature) marshal(c *Client) ([]byte, error) {
	m, err := expandFeature(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Feature: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFeature decodes JSON responses into the Feature resource schema.
func unmarshalFeature(b []byte, c *Client, res *Feature) (*Feature, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFeature(m, c, res)
}

func unmarshalMapFeature(m map[string]interface{}, c *Client, res *Feature) (*Feature, error) {

	flattened := flattenFeature(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandFeature expands Feature into a JSON request object.
func expandFeature(c *Client, f *Feature) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/features/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandFeatureSpec(c, f.Spec, res); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["spec"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenFeature flattens Feature from a JSON request object into the
// Feature type.
func flattenFeature(c *Client, i interface{}, res *Feature) *Feature {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Feature{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.ResourceState = flattenFeatureResourceState(c, m["resourceState"], res)
	resultRes.Spec = flattenFeatureSpec(c, m["spec"], res)
	resultRes.State = flattenFeatureState(c, m["state"], res)
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.DeleteTime = dcl.FlattenString(m["deleteTime"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandFeatureResourceStateMap expands the contents of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceStateMap(c *Client, f map[string]FeatureResourceState, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureResourceState(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureResourceStateSlice expands the contents of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceStateSlice(c *Client, f []FeatureResourceState, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureResourceState(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureResourceStateMap flattens the contents of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceStateMap(c *Client, i interface{}, res *Feature) map[string]FeatureResourceState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureResourceState{}
	}

	if len(a) == 0 {
		return map[string]FeatureResourceState{}
	}

	items := make(map[string]FeatureResourceState)
	for k, item := range a {
		items[k] = *flattenFeatureResourceState(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureResourceStateSlice flattens the contents of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceStateSlice(c *Client, i interface{}, res *Feature) []FeatureResourceState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureResourceState{}
	}

	if len(a) == 0 {
		return []FeatureResourceState{}
	}

	items := make([]FeatureResourceState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureResourceState(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureResourceState expands an instance of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceState(c *Client, f *FeatureResourceState, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureResourceState flattens an instance of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceState(c *Client, i interface{}, res *Feature) *FeatureResourceState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureResourceState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureResourceState
	}
	r.State = flattenFeatureResourceStateStateEnum(m["state"])
	r.HasResources = dcl.FlattenBool(m["hasResources"])

	return r
}

// expandFeatureSpecMap expands the contents of FeatureSpec into a JSON
// request object.
func expandFeatureSpecMap(c *Client, f map[string]FeatureSpec, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpec(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecSlice expands the contents of FeatureSpec into a JSON
// request object.
func expandFeatureSpecSlice(c *Client, f []FeatureSpec, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpec(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecMap flattens the contents of FeatureSpec from a JSON
// response object.
func flattenFeatureSpecMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpec{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpec{}
	}

	items := make(map[string]FeatureSpec)
	for k, item := range a {
		items[k] = *flattenFeatureSpec(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecSlice flattens the contents of FeatureSpec from a JSON
// response object.
func flattenFeatureSpecSlice(c *Client, i interface{}, res *Feature) []FeatureSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpec{}
	}

	if len(a) == 0 {
		return []FeatureSpec{}
	}

	items := make([]FeatureSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpec(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpec expands an instance of FeatureSpec into a JSON
// request object.
func expandFeatureSpec(c *Client, f *FeatureSpec, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureSpecMulticlusteringress(c, f.Multiclusteringress, res); err != nil {
		return nil, fmt.Errorf("error expanding Multiclusteringress into multiclusteringress: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["multiclusteringress"] = v
	}
	if v, err := expandFeatureSpecCloudauditlogging(c, f.Cloudauditlogging, res); err != nil {
		return nil, fmt.Errorf("error expanding Cloudauditlogging into cloudauditlogging: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudauditlogging"] = v
	}

	return m, nil
}

// flattenFeatureSpec flattens an instance of FeatureSpec from a JSON
// response object.
func flattenFeatureSpec(c *Client, i interface{}, res *Feature) *FeatureSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpec
	}
	r.Multiclusteringress = flattenFeatureSpecMulticlusteringress(c, m["multiclusteringress"], res)
	r.Cloudauditlogging = flattenFeatureSpecCloudauditlogging(c, m["cloudauditlogging"], res)

	return r
}

// expandFeatureSpecMulticlusteringressMap expands the contents of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringressMap(c *Client, f map[string]FeatureSpecMulticlusteringress, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecMulticlusteringress(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecMulticlusteringressSlice expands the contents of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringressSlice(c *Client, f []FeatureSpecMulticlusteringress, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecMulticlusteringress(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecMulticlusteringressMap flattens the contents of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringressMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecMulticlusteringress {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecMulticlusteringress{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecMulticlusteringress{}
	}

	items := make(map[string]FeatureSpecMulticlusteringress)
	for k, item := range a {
		items[k] = *flattenFeatureSpecMulticlusteringress(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecMulticlusteringressSlice flattens the contents of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringressSlice(c *Client, i interface{}, res *Feature) []FeatureSpecMulticlusteringress {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecMulticlusteringress{}
	}

	if len(a) == 0 {
		return []FeatureSpecMulticlusteringress{}
	}

	items := make([]FeatureSpecMulticlusteringress, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecMulticlusteringress(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpecMulticlusteringress expands an instance of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringress(c *Client, f *FeatureSpecMulticlusteringress, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ConfigMembership; !dcl.IsEmptyValueIndirect(v) {
		m["configMembership"] = v
	}

	return m, nil
}

// flattenFeatureSpecMulticlusteringress flattens an instance of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringress(c *Client, i interface{}, res *Feature) *FeatureSpecMulticlusteringress {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecMulticlusteringress{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpecMulticlusteringress
	}
	r.ConfigMembership = dcl.FlattenString(m["configMembership"])

	return r
}

// expandFeatureSpecCloudauditloggingMap expands the contents of FeatureSpecCloudauditlogging into a JSON
// request object.
func expandFeatureSpecCloudauditloggingMap(c *Client, f map[string]FeatureSpecCloudauditlogging, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecCloudauditlogging(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecCloudauditloggingSlice expands the contents of FeatureSpecCloudauditlogging into a JSON
// request object.
func expandFeatureSpecCloudauditloggingSlice(c *Client, f []FeatureSpecCloudauditlogging, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecCloudauditlogging(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecCloudauditloggingMap flattens the contents of FeatureSpecCloudauditlogging from a JSON
// response object.
func flattenFeatureSpecCloudauditloggingMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecCloudauditlogging {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecCloudauditlogging{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecCloudauditlogging{}
	}

	items := make(map[string]FeatureSpecCloudauditlogging)
	for k, item := range a {
		items[k] = *flattenFeatureSpecCloudauditlogging(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecCloudauditloggingSlice flattens the contents of FeatureSpecCloudauditlogging from a JSON
// response object.
func flattenFeatureSpecCloudauditloggingSlice(c *Client, i interface{}, res *Feature) []FeatureSpecCloudauditlogging {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecCloudauditlogging{}
	}

	if len(a) == 0 {
		return []FeatureSpecCloudauditlogging{}
	}

	items := make([]FeatureSpecCloudauditlogging, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecCloudauditlogging(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpecCloudauditlogging expands an instance of FeatureSpecCloudauditlogging into a JSON
// request object.
func expandFeatureSpecCloudauditlogging(c *Client, f *FeatureSpecCloudauditlogging, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowlistedServiceAccounts; v != nil {
		m["allowlistedServiceAccounts"] = v
	}

	return m, nil
}

// flattenFeatureSpecCloudauditlogging flattens an instance of FeatureSpecCloudauditlogging from a JSON
// response object.
func flattenFeatureSpecCloudauditlogging(c *Client, i interface{}, res *Feature) *FeatureSpecCloudauditlogging {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecCloudauditlogging{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpecCloudauditlogging
	}
	r.AllowlistedServiceAccounts = dcl.FlattenStringSlice(m["allowlistedServiceAccounts"])

	return r
}

// expandFeatureStateMap expands the contents of FeatureState into a JSON
// request object.
func expandFeatureStateMap(c *Client, f map[string]FeatureState, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureState(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateSlice expands the contents of FeatureState into a JSON
// request object.
func expandFeatureStateSlice(c *Client, f []FeatureState, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureState(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateMap flattens the contents of FeatureState from a JSON
// response object.
func flattenFeatureStateMap(c *Client, i interface{}, res *Feature) map[string]FeatureState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureState{}
	}

	if len(a) == 0 {
		return map[string]FeatureState{}
	}

	items := make(map[string]FeatureState)
	for k, item := range a {
		items[k] = *flattenFeatureState(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureStateSlice flattens the contents of FeatureState from a JSON
// response object.
func flattenFeatureStateSlice(c *Client, i interface{}, res *Feature) []FeatureState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureState{}
	}

	if len(a) == 0 {
		return []FeatureState{}
	}

	items := make([]FeatureState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureState(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureState expands an instance of FeatureState into a JSON
// request object.
func expandFeatureState(c *Client, f *FeatureState, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureState flattens an instance of FeatureState from a JSON
// response object.
func flattenFeatureState(c *Client, i interface{}, res *Feature) *FeatureState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureState
	}
	r.State = flattenFeatureStateState(c, m["state"], res)
	r.Servicemesh = flattenFeatureStateServicemesh(c, m["servicemesh"], res)

	return r
}

// expandFeatureStateStateMap expands the contents of FeatureStateState into a JSON
// request object.
func expandFeatureStateStateMap(c *Client, f map[string]FeatureStateState, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureStateState(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateStateSlice expands the contents of FeatureStateState into a JSON
// request object.
func expandFeatureStateStateSlice(c *Client, f []FeatureStateState, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureStateState(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateStateMap flattens the contents of FeatureStateState from a JSON
// response object.
func flattenFeatureStateStateMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateState{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateState{}
	}

	items := make(map[string]FeatureStateState)
	for k, item := range a {
		items[k] = *flattenFeatureStateState(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureStateStateSlice flattens the contents of FeatureStateState from a JSON
// response object.
func flattenFeatureStateStateSlice(c *Client, i interface{}, res *Feature) []FeatureStateState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateState{}
	}

	if len(a) == 0 {
		return []FeatureStateState{}
	}

	items := make([]FeatureStateState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateState(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureStateState expands an instance of FeatureStateState into a JSON
// request object.
func expandFeatureStateState(c *Client, f *FeatureStateState, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureStateState flattens an instance of FeatureStateState from a JSON
// response object.
func flattenFeatureStateState(c *Client, i interface{}, res *Feature) *FeatureStateState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureStateState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureStateState
	}
	r.Code = flattenFeatureStateStateCodeEnum(m["code"])
	r.Description = dcl.FlattenString(m["description"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// expandFeatureStateServicemeshMap expands the contents of FeatureStateServicemesh into a JSON
// request object.
func expandFeatureStateServicemeshMap(c *Client, f map[string]FeatureStateServicemesh, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureStateServicemesh(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateServicemeshSlice expands the contents of FeatureStateServicemesh into a JSON
// request object.
func expandFeatureStateServicemeshSlice(c *Client, f []FeatureStateServicemesh, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureStateServicemesh(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateServicemeshMap flattens the contents of FeatureStateServicemesh from a JSON
// response object.
func flattenFeatureStateServicemeshMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateServicemesh {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateServicemesh{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateServicemesh{}
	}

	items := make(map[string]FeatureStateServicemesh)
	for k, item := range a {
		items[k] = *flattenFeatureStateServicemesh(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureStateServicemeshSlice flattens the contents of FeatureStateServicemesh from a JSON
// response object.
func flattenFeatureStateServicemeshSlice(c *Client, i interface{}, res *Feature) []FeatureStateServicemesh {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateServicemesh{}
	}

	if len(a) == 0 {
		return []FeatureStateServicemesh{}
	}

	items := make([]FeatureStateServicemesh, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateServicemesh(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureStateServicemesh expands an instance of FeatureStateServicemesh into a JSON
// request object.
func expandFeatureStateServicemesh(c *Client, f *FeatureStateServicemesh, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureStateServicemesh flattens an instance of FeatureStateServicemesh from a JSON
// response object.
func flattenFeatureStateServicemesh(c *Client, i interface{}, res *Feature) *FeatureStateServicemesh {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureStateServicemesh{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureStateServicemesh
	}
	r.AnalysisMessages = flattenFeatureStateServicemeshAnalysisMessagesSlice(c, m["analysisMessages"], res)

	return r
}

// expandFeatureStateServicemeshAnalysisMessagesMap expands the contents of FeatureStateServicemeshAnalysisMessages into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessagesMap(c *Client, f map[string]FeatureStateServicemeshAnalysisMessages, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureStateServicemeshAnalysisMessages(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateServicemeshAnalysisMessagesSlice expands the contents of FeatureStateServicemeshAnalysisMessages into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessagesSlice(c *Client, f []FeatureStateServicemeshAnalysisMessages, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureStateServicemeshAnalysisMessages(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateServicemeshAnalysisMessagesMap flattens the contents of FeatureStateServicemeshAnalysisMessages from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateServicemeshAnalysisMessages {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateServicemeshAnalysisMessages{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateServicemeshAnalysisMessages{}
	}

	items := make(map[string]FeatureStateServicemeshAnalysisMessages)
	for k, item := range a {
		items[k] = *flattenFeatureStateServicemeshAnalysisMessages(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureStateServicemeshAnalysisMessagesSlice flattens the contents of FeatureStateServicemeshAnalysisMessages from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesSlice(c *Client, i interface{}, res *Feature) []FeatureStateServicemeshAnalysisMessages {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateServicemeshAnalysisMessages{}
	}

	if len(a) == 0 {
		return []FeatureStateServicemeshAnalysisMessages{}
	}

	items := make([]FeatureStateServicemeshAnalysisMessages, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateServicemeshAnalysisMessages(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureStateServicemeshAnalysisMessages expands an instance of FeatureStateServicemeshAnalysisMessages into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessages(c *Client, f *FeatureStateServicemeshAnalysisMessages, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureStateServicemeshAnalysisMessages flattens an instance of FeatureStateServicemeshAnalysisMessages from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessages(c *Client, i interface{}, res *Feature) *FeatureStateServicemeshAnalysisMessages {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureStateServicemeshAnalysisMessages{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureStateServicemeshAnalysisMessages
	}
	r.MessageBase = flattenFeatureStateServicemeshAnalysisMessagesMessageBase(c, m["messageBase"], res)
	r.Description = dcl.FlattenString(m["description"])
	r.ResourcePaths = dcl.FlattenStringSlice(m["resourcePaths"])
	r.Args = dcl.FlattenKeyValuePairs(m["args"])

	return r
}

// expandFeatureStateServicemeshAnalysisMessagesMessageBaseMap expands the contents of FeatureStateServicemeshAnalysisMessagesMessageBase into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessagesMessageBaseMap(c *Client, f map[string]FeatureStateServicemeshAnalysisMessagesMessageBase, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureStateServicemeshAnalysisMessagesMessageBase(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateServicemeshAnalysisMessagesMessageBaseSlice expands the contents of FeatureStateServicemeshAnalysisMessagesMessageBase into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessagesMessageBaseSlice(c *Client, f []FeatureStateServicemeshAnalysisMessagesMessageBase, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureStateServicemeshAnalysisMessagesMessageBase(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBaseMap flattens the contents of FeatureStateServicemeshAnalysisMessagesMessageBase from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBaseMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateServicemeshAnalysisMessagesMessageBase {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateServicemeshAnalysisMessagesMessageBase{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateServicemeshAnalysisMessagesMessageBase{}
	}

	items := make(map[string]FeatureStateServicemeshAnalysisMessagesMessageBase)
	for k, item := range a {
		items[k] = *flattenFeatureStateServicemeshAnalysisMessagesMessageBase(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBaseSlice flattens the contents of FeatureStateServicemeshAnalysisMessagesMessageBase from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBaseSlice(c *Client, i interface{}, res *Feature) []FeatureStateServicemeshAnalysisMessagesMessageBase {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateServicemeshAnalysisMessagesMessageBase{}
	}

	if len(a) == 0 {
		return []FeatureStateServicemeshAnalysisMessagesMessageBase{}
	}

	items := make([]FeatureStateServicemeshAnalysisMessagesMessageBase, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateServicemeshAnalysisMessagesMessageBase(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureStateServicemeshAnalysisMessagesMessageBase expands an instance of FeatureStateServicemeshAnalysisMessagesMessageBase into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessagesMessageBase(c *Client, f *FeatureStateServicemeshAnalysisMessagesMessageBase, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBase flattens an instance of FeatureStateServicemeshAnalysisMessagesMessageBase from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBase(c *Client, i interface{}, res *Feature) *FeatureStateServicemeshAnalysisMessagesMessageBase {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureStateServicemeshAnalysisMessagesMessageBase{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureStateServicemeshAnalysisMessagesMessageBase
	}
	r.Type = flattenFeatureStateServicemeshAnalysisMessagesMessageBaseType(c, m["type"], res)
	r.Level = flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(m["level"])
	r.DocumentationUrl = dcl.FlattenString(m["documentationUrl"])

	return r
}

// expandFeatureStateServicemeshAnalysisMessagesMessageBaseTypeMap expands the contents of FeatureStateServicemeshAnalysisMessagesMessageBaseType into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessagesMessageBaseTypeMap(c *Client, f map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseType, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureStateServicemeshAnalysisMessagesMessageBaseType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateServicemeshAnalysisMessagesMessageBaseTypeSlice expands the contents of FeatureStateServicemeshAnalysisMessagesMessageBaseType into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessagesMessageBaseTypeSlice(c *Client, f []FeatureStateServicemeshAnalysisMessagesMessageBaseType, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureStateServicemeshAnalysisMessagesMessageBaseType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBaseTypeMap flattens the contents of FeatureStateServicemeshAnalysisMessagesMessageBaseType from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBaseTypeMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseType{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseType{}
	}

	items := make(map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseType)
	for k, item := range a {
		items[k] = *flattenFeatureStateServicemeshAnalysisMessagesMessageBaseType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBaseTypeSlice flattens the contents of FeatureStateServicemeshAnalysisMessagesMessageBaseType from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBaseTypeSlice(c *Client, i interface{}, res *Feature) []FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateServicemeshAnalysisMessagesMessageBaseType{}
	}

	if len(a) == 0 {
		return []FeatureStateServicemeshAnalysisMessagesMessageBaseType{}
	}

	items := make([]FeatureStateServicemeshAnalysisMessagesMessageBaseType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateServicemeshAnalysisMessagesMessageBaseType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureStateServicemeshAnalysisMessagesMessageBaseType expands an instance of FeatureStateServicemeshAnalysisMessagesMessageBaseType into a JSON
// request object.
func expandFeatureStateServicemeshAnalysisMessagesMessageBaseType(c *Client, f *FeatureStateServicemeshAnalysisMessagesMessageBaseType, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBaseType flattens an instance of FeatureStateServicemeshAnalysisMessagesMessageBaseType from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBaseType(c *Client, i interface{}, res *Feature) *FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureStateServicemeshAnalysisMessagesMessageBaseType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureStateServicemeshAnalysisMessagesMessageBaseType
	}
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Code = dcl.FlattenString(m["code"])

	return r
}

// flattenFeatureResourceStateStateEnumMap flattens the contents of FeatureResourceStateStateEnum from a JSON
// response object.
func flattenFeatureResourceStateStateEnumMap(c *Client, i interface{}, res *Feature) map[string]FeatureResourceStateStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureResourceStateStateEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureResourceStateStateEnum{}
	}

	items := make(map[string]FeatureResourceStateStateEnum)
	for k, item := range a {
		items[k] = *flattenFeatureResourceStateStateEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureResourceStateStateEnumSlice flattens the contents of FeatureResourceStateStateEnum from a JSON
// response object.
func flattenFeatureResourceStateStateEnumSlice(c *Client, i interface{}, res *Feature) []FeatureResourceStateStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureResourceStateStateEnum{}
	}

	if len(a) == 0 {
		return []FeatureResourceStateStateEnum{}
	}

	items := make([]FeatureResourceStateStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureResourceStateStateEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureResourceStateStateEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureResourceStateStateEnum with the same value as that string.
func flattenFeatureResourceStateStateEnum(i interface{}) *FeatureResourceStateStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureResourceStateStateEnumRef(s)
}

// flattenFeatureStateStateCodeEnumMap flattens the contents of FeatureStateStateCodeEnum from a JSON
// response object.
func flattenFeatureStateStateCodeEnumMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateStateCodeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateStateCodeEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateStateCodeEnum{}
	}

	items := make(map[string]FeatureStateStateCodeEnum)
	for k, item := range a {
		items[k] = *flattenFeatureStateStateCodeEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureStateStateCodeEnumSlice flattens the contents of FeatureStateStateCodeEnum from a JSON
// response object.
func flattenFeatureStateStateCodeEnumSlice(c *Client, i interface{}, res *Feature) []FeatureStateStateCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateStateCodeEnum{}
	}

	if len(a) == 0 {
		return []FeatureStateStateCodeEnum{}
	}

	items := make([]FeatureStateStateCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateStateCodeEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureStateStateCodeEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureStateStateCodeEnum with the same value as that string.
func flattenFeatureStateStateCodeEnum(i interface{}) *FeatureStateStateCodeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureStateStateCodeEnumRef(s)
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumMap flattens the contents of FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum{}
	}

	items := make(map[string]FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum)
	for k, item := range a {
		items[k] = *flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumSlice flattens the contents of FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum from a JSON
// response object.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumSlice(c *Client, i interface{}, res *Feature) []FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum{}
	}

	if len(a) == 0 {
		return []FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum{}
	}

	items := make([]FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum with the same value as that string.
func flattenFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(i interface{}) *FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Feature) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFeature(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type featureDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         featureApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToFeatureDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]featureDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []featureDiff
	// For each operation name, create a featureDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := featureDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToFeatureApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToFeatureApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (featureApiOperation, error) {
	switch opName {

	case "updateFeatureUpdateFeatureOperation":
		return &updateFeatureUpdateFeatureOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractFeatureFields(r *Feature) error {
	vResourceState := r.ResourceState
	if vResourceState == nil {
		// note: explicitly not the empty object.
		vResourceState = &FeatureResourceState{}
	}
	if err := extractFeatureResourceStateFields(r, vResourceState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResourceState) {
		r.ResourceState = vResourceState
	}
	vSpec := r.Spec
	if vSpec == nil {
		// note: explicitly not the empty object.
		vSpec = &FeatureSpec{}
	}
	if err := extractFeatureSpecFields(r, vSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSpec) {
		r.Spec = vSpec
	}
	vState := r.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &FeatureState{}
	}
	if err := extractFeatureStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		r.State = vState
	}
	return nil
}
func extractFeatureResourceStateFields(r *Feature, o *FeatureResourceState) error {
	return nil
}
func extractFeatureSpecFields(r *Feature, o *FeatureSpec) error {
	vMulticlusteringress := o.Multiclusteringress
	if vMulticlusteringress == nil {
		// note: explicitly not the empty object.
		vMulticlusteringress = &FeatureSpecMulticlusteringress{}
	}
	if err := extractFeatureSpecMulticlusteringressFields(r, vMulticlusteringress); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMulticlusteringress) {
		o.Multiclusteringress = vMulticlusteringress
	}
	vCloudauditlogging := o.Cloudauditlogging
	if vCloudauditlogging == nil {
		// note: explicitly not the empty object.
		vCloudauditlogging = &FeatureSpecCloudauditlogging{}
	}
	if err := extractFeatureSpecCloudauditloggingFields(r, vCloudauditlogging); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudauditlogging) {
		o.Cloudauditlogging = vCloudauditlogging
	}
	return nil
}
func extractFeatureSpecMulticlusteringressFields(r *Feature, o *FeatureSpecMulticlusteringress) error {
	return nil
}
func extractFeatureSpecCloudauditloggingFields(r *Feature, o *FeatureSpecCloudauditlogging) error {
	return nil
}
func extractFeatureStateFields(r *Feature, o *FeatureState) error {
	vState := o.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &FeatureStateState{}
	}
	if err := extractFeatureStateStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		o.State = vState
	}
	vServicemesh := o.Servicemesh
	if vServicemesh == nil {
		// note: explicitly not the empty object.
		vServicemesh = &FeatureStateServicemesh{}
	}
	if err := extractFeatureStateServicemeshFields(r, vServicemesh); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vServicemesh) {
		o.Servicemesh = vServicemesh
	}
	return nil
}
func extractFeatureStateStateFields(r *Feature, o *FeatureStateState) error {
	return nil
}
func extractFeatureStateServicemeshFields(r *Feature, o *FeatureStateServicemesh) error {
	return nil
}
func extractFeatureStateServicemeshAnalysisMessagesFields(r *Feature, o *FeatureStateServicemeshAnalysisMessages) error {
	vMessageBase := o.MessageBase
	if vMessageBase == nil {
		// note: explicitly not the empty object.
		vMessageBase = &FeatureStateServicemeshAnalysisMessagesMessageBase{}
	}
	if err := extractFeatureStateServicemeshAnalysisMessagesMessageBaseFields(r, vMessageBase); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMessageBase) {
		o.MessageBase = vMessageBase
	}
	return nil
}
func extractFeatureStateServicemeshAnalysisMessagesMessageBaseFields(r *Feature, o *FeatureStateServicemeshAnalysisMessagesMessageBase) error {
	vType := o.Type
	if vType == nil {
		// note: explicitly not the empty object.
		vType = &FeatureStateServicemeshAnalysisMessagesMessageBaseType{}
	}
	if err := extractFeatureStateServicemeshAnalysisMessagesMessageBaseTypeFields(r, vType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vType) {
		o.Type = vType
	}
	return nil
}
func extractFeatureStateServicemeshAnalysisMessagesMessageBaseTypeFields(r *Feature, o *FeatureStateServicemeshAnalysisMessagesMessageBaseType) error {
	return nil
}

func postReadExtractFeatureFields(r *Feature) error {
	vResourceState := r.ResourceState
	if vResourceState == nil {
		// note: explicitly not the empty object.
		vResourceState = &FeatureResourceState{}
	}
	if err := postReadExtractFeatureResourceStateFields(r, vResourceState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResourceState) {
		r.ResourceState = vResourceState
	}
	vSpec := r.Spec
	if vSpec == nil {
		// note: explicitly not the empty object.
		vSpec = &FeatureSpec{}
	}
	if err := postReadExtractFeatureSpecFields(r, vSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSpec) {
		r.Spec = vSpec
	}
	vState := r.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &FeatureState{}
	}
	if err := postReadExtractFeatureStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		r.State = vState
	}
	return nil
}
func postReadExtractFeatureResourceStateFields(r *Feature, o *FeatureResourceState) error {
	return nil
}
func postReadExtractFeatureSpecFields(r *Feature, o *FeatureSpec) error {
	vMulticlusteringress := o.Multiclusteringress
	if vMulticlusteringress == nil {
		// note: explicitly not the empty object.
		vMulticlusteringress = &FeatureSpecMulticlusteringress{}
	}
	if err := extractFeatureSpecMulticlusteringressFields(r, vMulticlusteringress); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMulticlusteringress) {
		o.Multiclusteringress = vMulticlusteringress
	}
	vCloudauditlogging := o.Cloudauditlogging
	if vCloudauditlogging == nil {
		// note: explicitly not the empty object.
		vCloudauditlogging = &FeatureSpecCloudauditlogging{}
	}
	if err := extractFeatureSpecCloudauditloggingFields(r, vCloudauditlogging); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudauditlogging) {
		o.Cloudauditlogging = vCloudauditlogging
	}
	return nil
}
func postReadExtractFeatureSpecMulticlusteringressFields(r *Feature, o *FeatureSpecMulticlusteringress) error {
	return nil
}
func postReadExtractFeatureSpecCloudauditloggingFields(r *Feature, o *FeatureSpecCloudauditlogging) error {
	return nil
}
func postReadExtractFeatureStateFields(r *Feature, o *FeatureState) error {
	vState := o.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &FeatureStateState{}
	}
	if err := extractFeatureStateStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		o.State = vState
	}
	vServicemesh := o.Servicemesh
	if vServicemesh == nil {
		// note: explicitly not the empty object.
		vServicemesh = &FeatureStateServicemesh{}
	}
	if err := extractFeatureStateServicemeshFields(r, vServicemesh); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vServicemesh) {
		o.Servicemesh = vServicemesh
	}
	return nil
}
func postReadExtractFeatureStateStateFields(r *Feature, o *FeatureStateState) error {
	return nil
}
func postReadExtractFeatureStateServicemeshFields(r *Feature, o *FeatureStateServicemesh) error {
	return nil
}
func postReadExtractFeatureStateServicemeshAnalysisMessagesFields(r *Feature, o *FeatureStateServicemeshAnalysisMessages) error {
	vMessageBase := o.MessageBase
	if vMessageBase == nil {
		// note: explicitly not the empty object.
		vMessageBase = &FeatureStateServicemeshAnalysisMessagesMessageBase{}
	}
	if err := extractFeatureStateServicemeshAnalysisMessagesMessageBaseFields(r, vMessageBase); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMessageBase) {
		o.MessageBase = vMessageBase
	}
	return nil
}
func postReadExtractFeatureStateServicemeshAnalysisMessagesMessageBaseFields(r *Feature, o *FeatureStateServicemeshAnalysisMessagesMessageBase) error {
	vType := o.Type
	if vType == nil {
		// note: explicitly not the empty object.
		vType = &FeatureStateServicemeshAnalysisMessagesMessageBaseType{}
	}
	if err := extractFeatureStateServicemeshAnalysisMessagesMessageBaseTypeFields(r, vType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vType) {
		o.Type = vType
	}
	return nil
}
func postReadExtractFeatureStateServicemeshAnalysisMessagesMessageBaseTypeFields(r *Feature, o *FeatureStateServicemeshAnalysisMessagesMessageBaseType) error {
	return nil
}
