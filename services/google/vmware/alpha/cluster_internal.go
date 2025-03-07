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

func (r *Cluster) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "nodeTypeConfigs"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.PrivateCloud, "PrivateCloud"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterNodeTypeConfigs) validate() error {
	if err := dcl.Required(r, "nodeCount"); err != nil {
		return err
	}
	return nil
}
func (r *Cluster) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://vmwareengine.googleapis.com/v1/", params)
}

func (r *Cluster) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"privateCloud": dcl.ValueOrEmptyString(nr.PrivateCloud),
		"name":         dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{privateCloud}}/clusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Cluster) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"privateCloud": dcl.ValueOrEmptyString(nr.PrivateCloud),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{privateCloud}}/clusters", nr.basePath(), userBasePath, params), nil

}

func (r *Cluster) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"privateCloud": dcl.ValueOrEmptyString(nr.PrivateCloud),
		"name":         dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{privateCloud}}/clusters?clusterId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Cluster) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"privateCloud": dcl.ValueOrEmptyString(nr.PrivateCloud),
		"name":         dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{privateCloud}}/clusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

// clusterApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type clusterApiOperation interface {
	do(context.Context, *Cluster, *Client) error
}

// newUpdateClusterUpdateClusterRequest creates a request for an
// Cluster resource's UpdateCluster update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateClusterRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandClusterNodeTypeConfigsMap(c, f.NodeTypeConfigs, res); err != nil {
		return nil, fmt.Errorf("error expanding NodeTypeConfigs into nodeTypeConfigs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["nodeTypeConfigs"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s", *f.Project, *f.Location, *f.PrivateCloud, *f.Name)

	return req, nil
}

// marshalUpdateClusterUpdateClusterRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateClusterRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateClusterUpdateClusterOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateCluster")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateClusterRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateClusterRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listClusterRaw(ctx context.Context, r *Cluster, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ClusterMaxPage {
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

type listClusterOperation struct {
	Clusters []map[string]interface{} `json:"clusters"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listCluster(ctx context.Context, r *Cluster, pageToken string, pageSize int32) ([]*Cluster, string, error) {
	b, err := c.listClusterRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listClusterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Cluster
	for _, v := range m.Clusters {
		res, err := unmarshalMapCluster(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		res.PrivateCloud = r.PrivateCloud
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllCluster(ctx context.Context, f func(*Cluster) bool, resources []*Cluster) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteCluster(ctx, res)
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

type deleteClusterOperation struct{}

func (op *deleteClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	r, err := c.GetCluster(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Cluster not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetCluster checking for existence. error: %v", err)
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
		_, err := c.GetCluster(ctx, r)
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
type createClusterOperation struct {
	response map[string]interface{}
}

func (op *createClusterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
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

	if _, err := c.GetCluster(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getClusterRaw(ctx context.Context, r *Cluster) ([]byte, error) {

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

func (c *Client) clusterDiffsForRawDesired(ctx context.Context, rawDesired *Cluster, opts ...dcl.ApplyOption) (initial, desired *Cluster, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Cluster
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Cluster); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Cluster, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCluster(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Cluster resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Cluster resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Cluster resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Cluster: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Cluster: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractClusterFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeClusterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Cluster: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Cluster: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffCluster(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeClusterInitialState(rawInitial, rawDesired *Cluster) (*Cluster, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeClusterDesiredState(rawDesired, rawInitial *Cluster, opts ...dcl.ApplyOption) (*Cluster, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &Cluster{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.NodeTypeConfigs) || (dcl.IsEmptyValueIndirect(rawDesired.NodeTypeConfigs) && dcl.IsEmptyValueIndirect(rawInitial.NodeTypeConfigs)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.NodeTypeConfigs = rawInitial.NodeTypeConfigs
	} else {
		canonicalDesired.NodeTypeConfigs = rawDesired.NodeTypeConfigs
	}
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
	if dcl.NameToSelfLink(rawDesired.PrivateCloud, rawInitial.PrivateCloud) {
		canonicalDesired.PrivateCloud = rawInitial.PrivateCloud
	} else {
		canonicalDesired.PrivateCloud = rawDesired.PrivateCloud
	}

	return canonicalDesired, nil
}

func canonicalizeClusterNewState(c *Client, rawNew, rawDesired *Cluster) (*Cluster, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Management) && dcl.IsEmptyValueIndirect(rawDesired.Management) {
		rawNew.Management = rawDesired.Management
	} else {
		if dcl.BoolCanonicalize(rawDesired.Management, rawNew.Management) {
			rawNew.Management = rawDesired.Management
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NodeTypeConfigs) && dcl.IsEmptyValueIndirect(rawDesired.NodeTypeConfigs) {
		rawNew.NodeTypeConfigs = rawDesired.NodeTypeConfigs
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.PrivateCloud = rawDesired.PrivateCloud

	return rawNew, nil
}

func canonicalizeClusterNodeTypeConfigs(des, initial *ClusterNodeTypeConfigs, opts ...dcl.ApplyOption) *ClusterNodeTypeConfigs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterNodeTypeConfigs{}

	if dcl.IsZeroValue(des.NodeCount) || (dcl.IsEmptyValueIndirect(des.NodeCount) && dcl.IsEmptyValueIndirect(initial.NodeCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NodeCount = initial.NodeCount
	} else {
		cDes.NodeCount = des.NodeCount
	}
	if dcl.IsZeroValue(des.CustomCoreCount) || (dcl.IsEmptyValueIndirect(des.CustomCoreCount) && dcl.IsEmptyValueIndirect(initial.CustomCoreCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CustomCoreCount = initial.CustomCoreCount
	} else {
		cDes.CustomCoreCount = des.CustomCoreCount
	}

	return cDes
}

func canonicalizeClusterNodeTypeConfigsSlice(des, initial []ClusterNodeTypeConfigs, opts ...dcl.ApplyOption) []ClusterNodeTypeConfigs {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterNodeTypeConfigs, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterNodeTypeConfigs(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterNodeTypeConfigs, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterNodeTypeConfigs(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterNodeTypeConfigs(c *Client, des, nw *ClusterNodeTypeConfigs) *ClusterNodeTypeConfigs {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterNodeTypeConfigs while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewClusterNodeTypeConfigsSet(c *Client, des, nw []ClusterNodeTypeConfigs) []ClusterNodeTypeConfigs {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeTypeConfigs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterNodeTypeConfigsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterNodeTypeConfigsSlice(c *Client, des, nw []ClusterNodeTypeConfigs) []ClusterNodeTypeConfigs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterNodeTypeConfigs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeTypeConfigs(c, &d, &n))
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
func diffCluster(c *Client, desired, actual *Cluster, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Management, actual.Management, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Management")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeTypeConfigs, actual.NodeTypeConfigs, dcl.DiffInfo{ObjectFunction: compareClusterNodeTypeConfigsNewStyle, EmptyObject: EmptyClusterNodeTypeConfigs, OperationSelector: dcl.TriggersOperation("updateClusterUpdateClusterOperation")}, fn.AddNest("NodeTypeConfigs")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.PrivateCloud, actual.PrivateCloud, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrivateCloud")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareClusterNodeTypeConfigsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeTypeConfigs)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeTypeConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeTypeConfigs or *ClusterNodeTypeConfigs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeTypeConfigs)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeTypeConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeTypeConfigs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NodeCount, actual.NodeCount, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateClusterUpdateClusterOperation")}, fn.AddNest("NodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomCoreCount, actual.CustomCoreCount, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateClusterUpdateClusterOperation")}, fn.AddNest("CustomCoreCount")); len(ds) != 0 || err != nil {
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
func (r *Cluster) urlNormalized() *Cluster {
	normalized := dcl.Copy(*r).(Cluster)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.PrivateCloud = dcl.SelfLinkToName(r.PrivateCloud)
	return &normalized
}

func (r *Cluster) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateCluster" {
		fields := map[string]interface{}{
			"project":      dcl.ValueOrEmptyString(nr.Project),
			"location":     dcl.ValueOrEmptyString(nr.Location),
			"privateCloud": dcl.ValueOrEmptyString(nr.PrivateCloud),
			"name":         dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{privateCloud}}/clusters/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Cluster resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Cluster) marshal(c *Client) ([]byte, error) {
	m, err := expandCluster(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Cluster: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalCluster decodes JSON responses into the Cluster resource schema.
func unmarshalCluster(b []byte, c *Client, res *Cluster) (*Cluster, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapCluster(m, c, res)
}

func unmarshalMapCluster(m map[string]interface{}, c *Client, res *Cluster) (*Cluster, error) {

	flattened := flattenCluster(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandCluster expands Cluster into a JSON request object.
func expandCluster(c *Client, f *Cluster) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandClusterNodeTypeConfigsMap(c, f.NodeTypeConfigs, res); err != nil {
		return nil, fmt.Errorf("error expanding NodeTypeConfigs into nodeTypeConfigs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["nodeTypeConfigs"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding PrivateCloud into privateCloud: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["privateCloud"] = v
	}

	return m, nil
}

// flattenCluster flattens Cluster from a JSON request object into the
// Cluster type.
func flattenCluster(c *Client, i interface{}, res *Cluster) *Cluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Cluster{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.State = flattenClusterStateEnum(m["state"])
	resultRes.Management = dcl.FlattenBool(m["management"])
	resultRes.Uid = dcl.FlattenString(m["uid"])
	resultRes.NodeTypeConfigs = flattenClusterNodeTypeConfigsMap(c, m["nodeTypeConfigs"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.PrivateCloud = dcl.FlattenString(m["privateCloud"])

	return resultRes
}

// expandClusterNodeTypeConfigsMap expands the contents of ClusterNodeTypeConfigs into a JSON
// request object.
func expandClusterNodeTypeConfigsMap(c *Client, f map[string]ClusterNodeTypeConfigs, res *Cluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeTypeConfigs(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeTypeConfigsSlice expands the contents of ClusterNodeTypeConfigs into a JSON
// request object.
func expandClusterNodeTypeConfigsSlice(c *Client, f []ClusterNodeTypeConfigs, res *Cluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeTypeConfigs(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeTypeConfigsMap flattens the contents of ClusterNodeTypeConfigs from a JSON
// response object.
func flattenClusterNodeTypeConfigsMap(c *Client, i interface{}, res *Cluster) map[string]ClusterNodeTypeConfigs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeTypeConfigs{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeTypeConfigs{}
	}

	items := make(map[string]ClusterNodeTypeConfigs)
	for k, item := range a {
		items[k] = *flattenClusterNodeTypeConfigs(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenClusterNodeTypeConfigsSlice flattens the contents of ClusterNodeTypeConfigs from a JSON
// response object.
func flattenClusterNodeTypeConfigsSlice(c *Client, i interface{}, res *Cluster) []ClusterNodeTypeConfigs {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeTypeConfigs{}
	}

	if len(a) == 0 {
		return []ClusterNodeTypeConfigs{}
	}

	items := make([]ClusterNodeTypeConfigs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeTypeConfigs(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandClusterNodeTypeConfigs expands an instance of ClusterNodeTypeConfigs into a JSON
// request object.
func expandClusterNodeTypeConfigs(c *Client, f *ClusterNodeTypeConfigs, res *Cluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["nodeCount"] = v
	}
	if v := f.CustomCoreCount; !dcl.IsEmptyValueIndirect(v) {
		m["customCoreCount"] = v
	}

	return m, nil
}

// flattenClusterNodeTypeConfigs flattens an instance of ClusterNodeTypeConfigs from a JSON
// response object.
func flattenClusterNodeTypeConfigs(c *Client, i interface{}, res *Cluster) *ClusterNodeTypeConfigs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeTypeConfigs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterNodeTypeConfigs
	}
	r.NodeCount = dcl.FlattenInteger(m["nodeCount"])
	r.CustomCoreCount = dcl.FlattenInteger(m["customCoreCount"])

	return r
}

// flattenClusterStateEnumMap flattens the contents of ClusterStateEnum from a JSON
// response object.
func flattenClusterStateEnumMap(c *Client, i interface{}, res *Cluster) map[string]ClusterStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterStateEnum{}
	}

	items := make(map[string]ClusterStateEnum)
	for k, item := range a {
		items[k] = *flattenClusterStateEnum(item.(interface{}))
	}

	return items
}

// flattenClusterStateEnumSlice flattens the contents of ClusterStateEnum from a JSON
// response object.
func flattenClusterStateEnumSlice(c *Client, i interface{}, res *Cluster) []ClusterStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterStateEnum{}
	}

	if len(a) == 0 {
		return []ClusterStateEnum{}
	}

	items := make([]ClusterStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterStateEnum(item.(interface{})))
	}

	return items
}

// flattenClusterStateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStateEnum with the same value as that string.
func flattenClusterStateEnum(i interface{}) *ClusterStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ClusterStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Cluster) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalCluster(b, c, r)
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
		if nr.PrivateCloud == nil && ncr.PrivateCloud == nil {
			c.Config.Logger.Info("Both PrivateCloud fields null - considering equal.")
		} else if nr.PrivateCloud == nil || ncr.PrivateCloud == nil {
			c.Config.Logger.Info("Only one PrivateCloud field is null - considering unequal.")
			return false
		} else if *nr.PrivateCloud != *ncr.PrivateCloud {
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

type clusterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         clusterApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToClusterDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]clusterDiff, error) {
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
	var diffs []clusterDiff
	// For each operation name, create a clusterDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := clusterDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToClusterApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToClusterApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (clusterApiOperation, error) {
	switch opName {

	case "updateClusterUpdateClusterOperation":
		return &updateClusterUpdateClusterOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractClusterFields(r *Cluster) error {
	return nil
}
func extractClusterNodeTypeConfigsFields(r *Cluster, o *ClusterNodeTypeConfigs) error {
	return nil
}

func postReadExtractClusterFields(r *Cluster) error {
	return nil
}
func postReadExtractClusterNodeTypeConfigsFields(r *Cluster, o *ClusterNodeTypeConfigs) error {
	return nil
}
