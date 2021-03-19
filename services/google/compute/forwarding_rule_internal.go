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
package compute

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *ForwardingRule) validate() error {

	if err := dcl.Required(r, "project"); err != nil {
		return err
	}
	return nil
}
func (r *ForwardingRuleMetadataFilter) validate() error {
	return nil
}
func (r *ForwardingRuleMetadataFilterFilterLabel) validate() error {
	return nil
}

func forwardingRuleGetURL(userBasePath string, r *ForwardingRule) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/forwardingRules/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/forwardingRules/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func forwardingRuleListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/forwardingRules", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/forwardingRules", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func forwardingRuleCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/forwardingRules", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/forwardingRules", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func forwardingRuleDeleteURL(userBasePath string, r *ForwardingRule) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/forwardingRules/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/forwardingRules/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// forwardingRuleApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type forwardingRuleApiOperation interface {
	do(context.Context, *ForwardingRule, *Client) error
}

func (c *Client) listForwardingRuleRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := forwardingRuleListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ForwardingRuleMaxPage {
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

type listForwardingRuleOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listForwardingRule(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*ForwardingRule, string, error) {
	b, err := c.listForwardingRuleRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listForwardingRuleOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ForwardingRule
	for _, v := range m.Items {
		res := flattenForwardingRule(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllForwardingRule(ctx context.Context, f func(*ForwardingRule) bool, resources []*ForwardingRule) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteForwardingRule(ctx, res)
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

type deleteForwardingRuleOperation struct{}

func (op *deleteForwardingRuleOperation) do(ctx context.Context, r *ForwardingRule, c *Client) error {

	_, err := c.GetForwardingRule(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ForwardingRule not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetForwardingRule checking for existence. error: %v", err)
		return err
	}

	u, err := forwardingRuleDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		return err
	}
	_, err = c.GetForwardingRule(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createForwardingRuleOperation struct {
	response map[string]interface{}
}

func (op *createForwardingRuleOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createForwardingRuleOperation) do(ctx context.Context, r *ForwardingRule, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := forwardingRuleCreateURL(c.Config.BasePath, project, location)

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
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetForwardingRule(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getForwardingRuleRaw(ctx context.Context, r *ForwardingRule) ([]byte, error) {

	u, err := forwardingRuleGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) forwardingRuleDiffsForRawDesired(ctx context.Context, rawDesired *ForwardingRule, opts ...dcl.ApplyOption) (initial, desired *ForwardingRule, diffs []forwardingRuleDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ForwardingRule
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ForwardingRule); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ForwardingRule, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetForwardingRule(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ForwardingRule resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ForwardingRule resource: %v", err)
		}
		c.Config.Logger.Info("Found that ForwardingRule resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeForwardingRuleDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for ForwardingRule: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ForwardingRule: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeForwardingRuleInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ForwardingRule: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeForwardingRuleDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ForwardingRule: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffForwardingRule(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeForwardingRuleInitialState(rawInitial, rawDesired *ForwardingRule) (*ForwardingRule, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeForwardingRuleDesiredState(rawDesired, rawInitial *ForwardingRule, opts ...dcl.ApplyOption) (*ForwardingRule, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.AllPorts) {
		rawDesired.AllPorts = rawInitial.AllPorts
	}
	if dcl.IsZeroValue(rawDesired.AllowGlobalAccess) {
		rawDesired.AllowGlobalAccess = rawInitial.AllowGlobalAccess
	}
	if dcl.StringCanonicalize(rawDesired.BackendService, rawInitial.BackendService) {
		rawDesired.BackendService = rawInitial.BackendService
	}
	if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawInitial.CreationTimestamp) {
		rawDesired.CreationTimestamp = rawInitial.CreationTimestamp
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.IPAddress, rawInitial.IPAddress) {
		rawDesired.IPAddress = rawInitial.IPAddress
	}
	if dcl.IsZeroValue(rawDesired.IPProtocol) {
		rawDesired.IPProtocol = rawInitial.IPProtocol
	}
	if dcl.IsZeroValue(rawDesired.IPVersion) {
		rawDesired.IPVersion = rawInitial.IPVersion
	}
	if dcl.IsZeroValue(rawDesired.IsMirroringCollector) {
		rawDesired.IsMirroringCollector = rawInitial.IsMirroringCollector
	}
	if dcl.IsZeroValue(rawDesired.LoadBalancingScheme) {
		rawDesired.LoadBalancingScheme = rawInitial.LoadBalancingScheme
	}
	if dcl.IsZeroValue(rawDesired.MetadataFilter) {
		rawDesired.MetadataFilter = rawInitial.MetadataFilter
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.IsZeroValue(rawDesired.NetworkTier) {
		rawDesired.NetworkTier = rawInitial.NetworkTier
	}
	if dcl.StringCanonicalize(rawDesired.PortRange, rawInitial.PortRange) {
		rawDesired.PortRange = rawInitial.PortRange
	}
	if dcl.IsZeroValue(rawDesired.Ports) {
		rawDesired.Ports = rawInitial.Ports
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.StringCanonicalize(rawDesired.ServiceLabel, rawInitial.ServiceLabel) {
		rawDesired.ServiceLabel = rawInitial.ServiceLabel
	}
	if dcl.StringCanonicalize(rawDesired.ServiceName, rawInitial.ServiceName) {
		rawDesired.ServiceName = rawInitial.ServiceName
	}
	if dcl.StringCanonicalize(rawDesired.Subnetwork, rawInitial.Subnetwork) {
		rawDesired.Subnetwork = rawInitial.Subnetwork
	}
	if dcl.StringCanonicalize(rawDesired.Target, rawInitial.Target) {
		rawDesired.Target = rawInitial.Target
	}
	if dcl.StringCanonicalize(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeForwardingRuleNewState(c *Client, rawNew, rawDesired *ForwardingRule) (*ForwardingRule, error) {

	if dcl.IsEmptyValueIndirect(rawNew.AllPorts) && dcl.IsEmptyValueIndirect(rawDesired.AllPorts) {
		rawNew.AllPorts = rawDesired.AllPorts
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AllowGlobalAccess) && dcl.IsEmptyValueIndirect(rawDesired.AllowGlobalAccess) {
		rawNew.AllowGlobalAccess = rawDesired.AllowGlobalAccess
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BackendService) && dcl.IsEmptyValueIndirect(rawDesired.BackendService) {
		rawNew.BackendService = rawDesired.BackendService
	} else {
		if dcl.StringCanonicalize(rawDesired.BackendService, rawNew.BackendService) {
			rawNew.BackendService = rawDesired.BackendService
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
		if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawNew.CreationTimestamp) {
			rawNew.CreationTimestamp = rawDesired.CreationTimestamp
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAddress) && dcl.IsEmptyValueIndirect(rawDesired.IPAddress) {
		rawNew.IPAddress = rawDesired.IPAddress
	} else {
		if dcl.StringCanonicalize(rawDesired.IPAddress, rawNew.IPAddress) {
			rawNew.IPAddress = rawDesired.IPAddress
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPProtocol) && dcl.IsEmptyValueIndirect(rawDesired.IPProtocol) {
		rawNew.IPProtocol = rawDesired.IPProtocol
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPVersion) && dcl.IsEmptyValueIndirect(rawDesired.IPVersion) {
		rawNew.IPVersion = rawDesired.IPVersion
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IsMirroringCollector) && dcl.IsEmptyValueIndirect(rawDesired.IsMirroringCollector) {
		rawNew.IsMirroringCollector = rawDesired.IsMirroringCollector
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LoadBalancingScheme) && dcl.IsEmptyValueIndirect(rawDesired.LoadBalancingScheme) {
		rawNew.LoadBalancingScheme = rawDesired.LoadBalancingScheme
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MetadataFilter) && dcl.IsEmptyValueIndirect(rawDesired.MetadataFilter) {
		rawNew.MetadataFilter = rawDesired.MetadataFilter
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.StringCanonicalize(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkTier) && dcl.IsEmptyValueIndirect(rawDesired.NetworkTier) {
		rawNew.NetworkTier = rawDesired.NetworkTier
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PortRange) && dcl.IsEmptyValueIndirect(rawDesired.PortRange) {
		rawNew.PortRange = rawDesired.PortRange
	} else {
		if dcl.StringCanonicalize(rawDesired.PortRange, rawNew.PortRange) {
			rawNew.PortRange = rawDesired.PortRange
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Ports) && dcl.IsEmptyValueIndirect(rawDesired.Ports) {
		rawNew.Ports = rawDesired.Ports
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceLabel) && dcl.IsEmptyValueIndirect(rawDesired.ServiceLabel) {
		rawNew.ServiceLabel = rawDesired.ServiceLabel
	} else {
		if dcl.StringCanonicalize(rawDesired.ServiceLabel, rawNew.ServiceLabel) {
			rawNew.ServiceLabel = rawDesired.ServiceLabel
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceName) && dcl.IsEmptyValueIndirect(rawDesired.ServiceName) {
		rawNew.ServiceName = rawDesired.ServiceName
	} else {
		if dcl.StringCanonicalize(rawDesired.ServiceName, rawNew.ServiceName) {
			rawNew.ServiceName = rawDesired.ServiceName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Subnetwork) && dcl.IsEmptyValueIndirect(rawDesired.Subnetwork) {
		rawNew.Subnetwork = rawDesired.Subnetwork
	} else {
		if dcl.StringCanonicalize(rawDesired.Subnetwork, rawNew.Subnetwork) {
			rawNew.Subnetwork = rawDesired.Subnetwork
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Target) && dcl.IsEmptyValueIndirect(rawDesired.Target) {
		rawNew.Target = rawDesired.Target
	} else {
		if dcl.StringCanonicalize(rawDesired.Target, rawNew.Target) {
			rawNew.Target = rawDesired.Target
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.StringCanonicalize(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeForwardingRuleMetadataFilter(des, initial *ForwardingRuleMetadataFilter, opts ...dcl.ApplyOption) *ForwardingRuleMetadataFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.FilterMatchCriteria) {
		des.FilterMatchCriteria = initial.FilterMatchCriteria
	}
	if dcl.IsZeroValue(des.FilterLabel) {
		des.FilterLabel = initial.FilterLabel
	}

	return des
}

func canonicalizeNewForwardingRuleMetadataFilter(c *Client, des, nw *ForwardingRuleMetadataFilter) *ForwardingRuleMetadataFilter {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewForwardingRuleMetadataFilterSet(c *Client, des, nw []ForwardingRuleMetadataFilter) []ForwardingRuleMetadataFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []ForwardingRuleMetadataFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareForwardingRuleMetadataFilter(c, &d, &n) {
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

func canonicalizeForwardingRuleMetadataFilterFilterLabel(des, initial *ForwardingRuleMetadataFilterFilterLabel, opts ...dcl.ApplyOption) *ForwardingRuleMetadataFilterFilterLabel {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewForwardingRuleMetadataFilterFilterLabel(c *Client, des, nw *ForwardingRuleMetadataFilterFilterLabel) *ForwardingRuleMetadataFilterFilterLabel {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) || dcl.IsZeroValue(des.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewForwardingRuleMetadataFilterFilterLabelSet(c *Client, des, nw []ForwardingRuleMetadataFilterFilterLabel) []ForwardingRuleMetadataFilterFilterLabel {
	if des == nil {
		return nw
	}
	var reorderedNew []ForwardingRuleMetadataFilterFilterLabel
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareForwardingRuleMetadataFilterFilterLabel(c, &d, &n) {
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

type forwardingRuleDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         forwardingRuleApiOperation
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffForwardingRule(c *Client, desired, actual *ForwardingRule, opts ...dcl.ApplyOption) ([]forwardingRuleDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []forwardingRuleDiff
	if !reflect.DeepEqual(desired.AllPorts, actual.AllPorts) {
		c.Config.Logger.Infof("Detected diff in AllPorts.\nDESIRED: %v\nACTUAL: %v", desired.AllPorts, actual.AllPorts)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "AllPorts",
		})
	}
	if !reflect.DeepEqual(desired.AllowGlobalAccess, actual.AllowGlobalAccess) {
		c.Config.Logger.Infof("Detected diff in AllowGlobalAccess.\nDESIRED: %v\nACTUAL: %v", desired.AllowGlobalAccess, actual.AllowGlobalAccess)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "AllowGlobalAccess",
		})
	}
	if !dcl.IsZeroValue(desired.BackendService) && !dcl.StringCanonicalize(desired.BackendService, actual.BackendService) {
		c.Config.Logger.Infof("Detected diff in BackendService.\nDESIRED: %v\nACTUAL: %v", desired.BackendService, actual.BackendService)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "BackendService",
		})
	}
	if !dcl.IsZeroValue(desired.CreationTimestamp) && !dcl.StringCanonicalize(desired.CreationTimestamp, actual.CreationTimestamp) {
		c.Config.Logger.Infof("Detected diff in CreationTimestamp.\nDESIRED: %v\nACTUAL: %v", desired.CreationTimestamp, actual.CreationTimestamp)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "CreationTimestamp",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.IPAddress) && !dcl.StringCanonicalize(desired.IPAddress, actual.IPAddress) {
		c.Config.Logger.Infof("Detected diff in IPAddress.\nDESIRED: %v\nACTUAL: %v", desired.IPAddress, actual.IPAddress)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "IPAddress",
		})
	}
	if !reflect.DeepEqual(desired.IPProtocol, actual.IPProtocol) {
		c.Config.Logger.Infof("Detected diff in IPProtocol.\nDESIRED: %v\nACTUAL: %v", desired.IPProtocol, actual.IPProtocol)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "IPProtocol",
		})
	}
	if !reflect.DeepEqual(desired.IPVersion, actual.IPVersion) {
		c.Config.Logger.Infof("Detected diff in IPVersion.\nDESIRED: %v\nACTUAL: %v", desired.IPVersion, actual.IPVersion)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "IPVersion",
		})
	}
	if !reflect.DeepEqual(desired.IsMirroringCollector, actual.IsMirroringCollector) {
		c.Config.Logger.Infof("Detected diff in IsMirroringCollector.\nDESIRED: %v\nACTUAL: %v", desired.IsMirroringCollector, actual.IsMirroringCollector)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "IsMirroringCollector",
		})
	}
	if !reflect.DeepEqual(desired.LoadBalancingScheme, actual.LoadBalancingScheme) {
		c.Config.Logger.Infof("Detected diff in LoadBalancingScheme.\nDESIRED: %v\nACTUAL: %v", desired.LoadBalancingScheme, actual.LoadBalancingScheme)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "LoadBalancingScheme",
		})
	}
	if compareForwardingRuleMetadataFilterSlice(c, desired.MetadataFilter, actual.MetadataFilter) {
		c.Config.Logger.Infof("Detected diff in MetadataFilter.\nDESIRED: %v\nACTUAL: %v", desired.MetadataFilter, actual.MetadataFilter)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "MetadataFilter",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Network) && !dcl.StringCanonicalize(desired.Network, actual.Network) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %v\nACTUAL: %v", desired.Network, actual.Network)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Network",
		})
	}
	if !reflect.DeepEqual(desired.NetworkTier, actual.NetworkTier) {
		c.Config.Logger.Infof("Detected diff in NetworkTier.\nDESIRED: %v\nACTUAL: %v", desired.NetworkTier, actual.NetworkTier)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "NetworkTier",
		})
	}
	if !dcl.IsZeroValue(desired.PortRange) && !dcl.StringCanonicalize(desired.PortRange, actual.PortRange) {
		c.Config.Logger.Infof("Detected diff in PortRange.\nDESIRED: %v\nACTUAL: %v", desired.PortRange, actual.PortRange)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "PortRange",
		})
	}
	if !dcl.StringSliceEquals(desired.Ports, actual.Ports) {
		c.Config.Logger.Infof("Detected diff in Ports.\nDESIRED: %v\nACTUAL: %v", desired.Ports, actual.Ports)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Ports",
		})
	}
	if !dcl.IsZeroValue(desired.Region) && !dcl.StringCanonicalize(desired.Region, actual.Region) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %v\nACTUAL: %v", desired.Region, actual.Region)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Region",
		})
	}
	if !dcl.IsZeroValue(desired.SelfLink) && !dcl.StringCanonicalize(desired.SelfLink, actual.SelfLink) {
		c.Config.Logger.Infof("Detected diff in SelfLink.\nDESIRED: %v\nACTUAL: %v", desired.SelfLink, actual.SelfLink)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "SelfLink",
		})
	}
	if !dcl.IsZeroValue(desired.ServiceLabel) && !dcl.StringCanonicalize(desired.ServiceLabel, actual.ServiceLabel) {
		c.Config.Logger.Infof("Detected diff in ServiceLabel.\nDESIRED: %v\nACTUAL: %v", desired.ServiceLabel, actual.ServiceLabel)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "ServiceLabel",
		})
	}
	if !dcl.IsZeroValue(desired.ServiceName) && !dcl.StringCanonicalize(desired.ServiceName, actual.ServiceName) {
		c.Config.Logger.Infof("Detected diff in ServiceName.\nDESIRED: %v\nACTUAL: %v", desired.ServiceName, actual.ServiceName)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "ServiceName",
		})
	}
	if !dcl.IsZeroValue(desired.Subnetwork) && !dcl.StringCanonicalize(desired.Subnetwork, actual.Subnetwork) {
		c.Config.Logger.Infof("Detected diff in Subnetwork.\nDESIRED: %v\nACTUAL: %v", desired.Subnetwork, actual.Subnetwork)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Subnetwork",
		})
	}
	if !dcl.IsZeroValue(desired.Target) && !dcl.StringCanonicalize(desired.Target, actual.Target) {
		c.Config.Logger.Infof("Detected diff in Target.\nDESIRED: %v\nACTUAL: %v", desired.Target, actual.Target)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Target",
		})
	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.StringCanonicalize(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %v\nACTUAL: %v", desired.Project, actual.Project)
		diffs = append(diffs, forwardingRuleDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
		})
	}
	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []forwardingRuleDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareForwardingRuleMetadataFilter(c *Client, desired, actual *ForwardingRuleMetadataFilter) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.FilterMatchCriteria == nil && desired.FilterMatchCriteria != nil && !dcl.IsEmptyValueIndirect(desired.FilterMatchCriteria) {
		c.Config.Logger.Infof("desired FilterMatchCriteria %s - but actually nil", dcl.SprintResource(desired.FilterMatchCriteria))
		return true
	}
	if !reflect.DeepEqual(desired.FilterMatchCriteria, actual.FilterMatchCriteria) && !dcl.IsZeroValue(desired.FilterMatchCriteria) {
		c.Config.Logger.Infof("Diff in FilterMatchCriteria. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FilterMatchCriteria), dcl.SprintResource(actual.FilterMatchCriteria))
		return true
	}
	if actual.FilterLabel == nil && desired.FilterLabel != nil && !dcl.IsEmptyValueIndirect(desired.FilterLabel) {
		c.Config.Logger.Infof("desired FilterLabel %s - but actually nil", dcl.SprintResource(desired.FilterLabel))
		return true
	}
	if compareForwardingRuleMetadataFilterFilterLabelSlice(c, desired.FilterLabel, actual.FilterLabel) && !dcl.IsZeroValue(desired.FilterLabel) {
		c.Config.Logger.Infof("Diff in FilterLabel. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FilterLabel), dcl.SprintResource(actual.FilterLabel))
		return true
	}
	return false
}

func compareForwardingRuleMetadataFilterSlice(c *Client, desired, actual []ForwardingRuleMetadataFilter) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleMetadataFilter, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareForwardingRuleMetadataFilter(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ForwardingRuleMetadataFilter, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareForwardingRuleMetadataFilterMap(c *Client, desired, actual map[string]ForwardingRuleMetadataFilter) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleMetadataFilter, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ForwardingRuleMetadataFilter, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareForwardingRuleMetadataFilter(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ForwardingRuleMetadataFilter, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareForwardingRuleMetadataFilterFilterLabel(c *Client, desired, actual *ForwardingRuleMetadataFilterFilterLabel) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareForwardingRuleMetadataFilterFilterLabelSlice(c *Client, desired, actual []ForwardingRuleMetadataFilterFilterLabel) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleMetadataFilterFilterLabel, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareForwardingRuleMetadataFilterFilterLabel(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ForwardingRuleMetadataFilterFilterLabel, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareForwardingRuleMetadataFilterFilterLabelMap(c *Client, desired, actual map[string]ForwardingRuleMetadataFilterFilterLabel) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleMetadataFilterFilterLabel, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ForwardingRuleMetadataFilterFilterLabel, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareForwardingRuleMetadataFilterFilterLabel(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ForwardingRuleMetadataFilterFilterLabel, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareForwardingRuleIPProtocolEnumSlice(c *Client, desired, actual []ForwardingRuleIPProtocolEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleIPProtocolEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareForwardingRuleIPProtocolEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ForwardingRuleIPProtocolEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareForwardingRuleIPProtocolEnum(c *Client, desired, actual *ForwardingRuleIPProtocolEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareForwardingRuleIPVersionEnumSlice(c *Client, desired, actual []ForwardingRuleIPVersionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleIPVersionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareForwardingRuleIPVersionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ForwardingRuleIPVersionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareForwardingRuleIPVersionEnum(c *Client, desired, actual *ForwardingRuleIPVersionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareForwardingRuleLoadBalancingSchemeEnumSlice(c *Client, desired, actual []ForwardingRuleLoadBalancingSchemeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleLoadBalancingSchemeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareForwardingRuleLoadBalancingSchemeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ForwardingRuleLoadBalancingSchemeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareForwardingRuleLoadBalancingSchemeEnum(c *Client, desired, actual *ForwardingRuleLoadBalancingSchemeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareForwardingRuleMetadataFilterFilterMatchCriteriaEnumSlice(c *Client, desired, actual []ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleMetadataFilterFilterMatchCriteriaEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareForwardingRuleMetadataFilterFilterMatchCriteriaEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ForwardingRuleMetadataFilterFilterMatchCriteriaEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareForwardingRuleMetadataFilterFilterMatchCriteriaEnum(c *Client, desired, actual *ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareForwardingRuleNetworkTierEnumSlice(c *Client, desired, actual []ForwardingRuleNetworkTierEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ForwardingRuleNetworkTierEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareForwardingRuleNetworkTierEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ForwardingRuleNetworkTierEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareForwardingRuleNetworkTierEnum(c *Client, desired, actual *ForwardingRuleNetworkTierEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *ForwardingRule) urlNormalized() *ForwardingRule {
	normalized := deepcopy.Copy(*r).(ForwardingRule)
	normalized.BackendService = dcl.SelfLinkToName(r.BackendService)
	normalized.CreationTimestamp = dcl.SelfLinkToName(r.CreationTimestamp)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.IPAddress = dcl.SelfLinkToName(r.IPAddress)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.PortRange = dcl.SelfLinkToName(r.PortRange)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.ServiceLabel = dcl.SelfLinkToName(r.ServiceLabel)
	normalized.ServiceName = dcl.SelfLinkToName(r.ServiceName)
	normalized.Subnetwork = dcl.SelfLinkToName(r.Subnetwork)
	normalized.Target = dcl.SelfLinkToName(r.Target)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *ForwardingRule) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *ForwardingRule) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *ForwardingRule) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *ForwardingRule) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ForwardingRule resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ForwardingRule) marshal(c *Client) ([]byte, error) {
	m, err := expandForwardingRule(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ForwardingRule: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalForwardingRule decodes JSON responses into the ForwardingRule resource schema.
func unmarshalForwardingRule(b []byte, c *Client) (*ForwardingRule, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapForwardingRule(m, c)
}

func unmarshalMapForwardingRule(m map[string]interface{}, c *Client) (*ForwardingRule, error) {

	return flattenForwardingRule(c, m), nil
}

// expandForwardingRule expands ForwardingRule into a JSON request object.
func expandForwardingRule(c *Client, f *ForwardingRule) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.AllPorts; !dcl.IsEmptyValueIndirect(v) {
		m["allPorts"] = v
	}
	if v := f.AllowGlobalAccess; !dcl.IsEmptyValueIndirect(v) {
		m["allowGlobalAccess"] = v
	}
	if v := f.BackendService; !dcl.IsEmptyValueIndirect(v) {
		m["backendService"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["IPAddress"] = v
	}
	if v := f.IPProtocol; !dcl.IsEmptyValueIndirect(v) {
		m["IPProtocol"] = v
	}
	if v := f.IPVersion; !dcl.IsEmptyValueIndirect(v) {
		m["ipVersion"] = v
	}
	if v := f.IsMirroringCollector; !dcl.IsEmptyValueIndirect(v) {
		m["isMirroringCollector"] = v
	}
	if v := f.LoadBalancingScheme; !dcl.IsEmptyValueIndirect(v) {
		m["loadBalancingScheme"] = v
	}
	if v, err := expandForwardingRuleMetadataFilterSlice(c, f.MetadataFilter); err != nil {
		return nil, fmt.Errorf("error expanding MetadataFilter into metadataFilters: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadataFilters"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.NetworkTier; !dcl.IsEmptyValueIndirect(v) {
		m["networkTier"] = v
	}
	if v := f.PortRange; !dcl.IsEmptyValueIndirect(v) {
		m["portRange"] = v
	}
	if v := f.Ports; !dcl.IsEmptyValueIndirect(v) {
		m["ports"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.ServiceLabel; !dcl.IsEmptyValueIndirect(v) {
		m["serviceLabel"] = v
	}
	if v := f.ServiceName; !dcl.IsEmptyValueIndirect(v) {
		m["serviceName"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}
	if v := f.Target; !dcl.IsEmptyValueIndirect(v) {
		m["target"] = v
	}
	if v := f.Project; !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenForwardingRule flattens ForwardingRule from a JSON request object into the
// ForwardingRule type.
func flattenForwardingRule(c *Client, i interface{}) *ForwardingRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &ForwardingRule{}
	r.AllPorts = dcl.FlattenBool(m["allPorts"])
	r.AllowGlobalAccess = dcl.FlattenBool(m["allowGlobalAccess"])
	r.BackendService = dcl.FlattenString(m["backendService"])
	r.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	r.Description = dcl.FlattenString(m["description"])
	r.IPAddress = dcl.FlattenString(m["IPAddress"])
	r.IPProtocol = flattenForwardingRuleIPProtocolEnum(m["IPProtocol"])
	r.IPVersion = flattenForwardingRuleIPVersionEnum(m["ipVersion"])
	r.IsMirroringCollector = dcl.FlattenBool(m["isMirroringCollector"])
	r.LoadBalancingScheme = flattenForwardingRuleLoadBalancingSchemeEnum(m["loadBalancingScheme"])
	r.MetadataFilter = flattenForwardingRuleMetadataFilterSlice(c, m["metadataFilters"])
	r.Name = dcl.FlattenString(m["name"])
	r.Network = dcl.FlattenString(m["network"])
	r.NetworkTier = flattenForwardingRuleNetworkTierEnum(m["networkTier"])
	r.PortRange = dcl.FlattenString(m["portRange"])
	r.Ports = dcl.FlattenStringSlice(m["ports"])
	r.Region = dcl.FlattenString(m["region"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.ServiceLabel = dcl.FlattenString(m["serviceLabel"])
	r.ServiceName = dcl.FlattenString(m["serviceName"])
	r.Subnetwork = dcl.FlattenString(m["subnetwork"])
	r.Target = dcl.FlattenString(m["target"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandForwardingRuleMetadataFilterMap expands the contents of ForwardingRuleMetadataFilter into a JSON
// request object.
func expandForwardingRuleMetadataFilterMap(c *Client, f map[string]ForwardingRuleMetadataFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandForwardingRuleMetadataFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandForwardingRuleMetadataFilterSlice expands the contents of ForwardingRuleMetadataFilter into a JSON
// request object.
func expandForwardingRuleMetadataFilterSlice(c *Client, f []ForwardingRuleMetadataFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandForwardingRuleMetadataFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenForwardingRuleMetadataFilterMap flattens the contents of ForwardingRuleMetadataFilter from a JSON
// response object.
func flattenForwardingRuleMetadataFilterMap(c *Client, i interface{}) map[string]ForwardingRuleMetadataFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ForwardingRuleMetadataFilter{}
	}

	if len(a) == 0 {
		return map[string]ForwardingRuleMetadataFilter{}
	}

	items := make(map[string]ForwardingRuleMetadataFilter)
	for k, item := range a {
		items[k] = *flattenForwardingRuleMetadataFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenForwardingRuleMetadataFilterSlice flattens the contents of ForwardingRuleMetadataFilter from a JSON
// response object.
func flattenForwardingRuleMetadataFilterSlice(c *Client, i interface{}) []ForwardingRuleMetadataFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []ForwardingRuleMetadataFilter{}
	}

	if len(a) == 0 {
		return []ForwardingRuleMetadataFilter{}
	}

	items := make([]ForwardingRuleMetadataFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenForwardingRuleMetadataFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandForwardingRuleMetadataFilter expands an instance of ForwardingRuleMetadataFilter into a JSON
// request object.
func expandForwardingRuleMetadataFilter(c *Client, f *ForwardingRuleMetadataFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FilterMatchCriteria; !dcl.IsEmptyValueIndirect(v) {
		m["filterMatchCriteria"] = v
	}
	if v, err := expandForwardingRuleMetadataFilterFilterLabelSlice(c, f.FilterLabel); err != nil {
		return nil, fmt.Errorf("error expanding FilterLabel into filterLabels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["filterLabels"] = v
	}

	return m, nil
}

// flattenForwardingRuleMetadataFilter flattens an instance of ForwardingRuleMetadataFilter from a JSON
// response object.
func flattenForwardingRuleMetadataFilter(c *Client, i interface{}) *ForwardingRuleMetadataFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ForwardingRuleMetadataFilter{}
	r.FilterMatchCriteria = flattenForwardingRuleMetadataFilterFilterMatchCriteriaEnum(m["filterMatchCriteria"])
	r.FilterLabel = flattenForwardingRuleMetadataFilterFilterLabelSlice(c, m["filterLabels"])

	return r
}

// expandForwardingRuleMetadataFilterFilterLabelMap expands the contents of ForwardingRuleMetadataFilterFilterLabel into a JSON
// request object.
func expandForwardingRuleMetadataFilterFilterLabelMap(c *Client, f map[string]ForwardingRuleMetadataFilterFilterLabel) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandForwardingRuleMetadataFilterFilterLabel(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandForwardingRuleMetadataFilterFilterLabelSlice expands the contents of ForwardingRuleMetadataFilterFilterLabel into a JSON
// request object.
func expandForwardingRuleMetadataFilterFilterLabelSlice(c *Client, f []ForwardingRuleMetadataFilterFilterLabel) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandForwardingRuleMetadataFilterFilterLabel(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenForwardingRuleMetadataFilterFilterLabelMap flattens the contents of ForwardingRuleMetadataFilterFilterLabel from a JSON
// response object.
func flattenForwardingRuleMetadataFilterFilterLabelMap(c *Client, i interface{}) map[string]ForwardingRuleMetadataFilterFilterLabel {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ForwardingRuleMetadataFilterFilterLabel{}
	}

	if len(a) == 0 {
		return map[string]ForwardingRuleMetadataFilterFilterLabel{}
	}

	items := make(map[string]ForwardingRuleMetadataFilterFilterLabel)
	for k, item := range a {
		items[k] = *flattenForwardingRuleMetadataFilterFilterLabel(c, item.(map[string]interface{}))
	}

	return items
}

// flattenForwardingRuleMetadataFilterFilterLabelSlice flattens the contents of ForwardingRuleMetadataFilterFilterLabel from a JSON
// response object.
func flattenForwardingRuleMetadataFilterFilterLabelSlice(c *Client, i interface{}) []ForwardingRuleMetadataFilterFilterLabel {
	a, ok := i.([]interface{})
	if !ok {
		return []ForwardingRuleMetadataFilterFilterLabel{}
	}

	if len(a) == 0 {
		return []ForwardingRuleMetadataFilterFilterLabel{}
	}

	items := make([]ForwardingRuleMetadataFilterFilterLabel, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenForwardingRuleMetadataFilterFilterLabel(c, item.(map[string]interface{})))
	}

	return items
}

// expandForwardingRuleMetadataFilterFilterLabel expands an instance of ForwardingRuleMetadataFilterFilterLabel into a JSON
// request object.
func expandForwardingRuleMetadataFilterFilterLabel(c *Client, f *ForwardingRuleMetadataFilterFilterLabel) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenForwardingRuleMetadataFilterFilterLabel flattens an instance of ForwardingRuleMetadataFilterFilterLabel from a JSON
// response object.
func flattenForwardingRuleMetadataFilterFilterLabel(c *Client, i interface{}) *ForwardingRuleMetadataFilterFilterLabel {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ForwardingRuleMetadataFilterFilterLabel{}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// flattenForwardingRuleIPProtocolEnumSlice flattens the contents of ForwardingRuleIPProtocolEnum from a JSON
// response object.
func flattenForwardingRuleIPProtocolEnumSlice(c *Client, i interface{}) []ForwardingRuleIPProtocolEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ForwardingRuleIPProtocolEnum{}
	}

	if len(a) == 0 {
		return []ForwardingRuleIPProtocolEnum{}
	}

	items := make([]ForwardingRuleIPProtocolEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenForwardingRuleIPProtocolEnum(item.(interface{})))
	}

	return items
}

// flattenForwardingRuleIPProtocolEnum asserts that an interface is a string, and returns a
// pointer to a *ForwardingRuleIPProtocolEnum with the same value as that string.
func flattenForwardingRuleIPProtocolEnum(i interface{}) *ForwardingRuleIPProtocolEnum {
	s, ok := i.(string)
	if !ok {
		return ForwardingRuleIPProtocolEnumRef("")
	}

	return ForwardingRuleIPProtocolEnumRef(s)
}

// flattenForwardingRuleIPVersionEnumSlice flattens the contents of ForwardingRuleIPVersionEnum from a JSON
// response object.
func flattenForwardingRuleIPVersionEnumSlice(c *Client, i interface{}) []ForwardingRuleIPVersionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ForwardingRuleIPVersionEnum{}
	}

	if len(a) == 0 {
		return []ForwardingRuleIPVersionEnum{}
	}

	items := make([]ForwardingRuleIPVersionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenForwardingRuleIPVersionEnum(item.(interface{})))
	}

	return items
}

// flattenForwardingRuleIPVersionEnum asserts that an interface is a string, and returns a
// pointer to a *ForwardingRuleIPVersionEnum with the same value as that string.
func flattenForwardingRuleIPVersionEnum(i interface{}) *ForwardingRuleIPVersionEnum {
	s, ok := i.(string)
	if !ok {
		return ForwardingRuleIPVersionEnumRef("")
	}

	return ForwardingRuleIPVersionEnumRef(s)
}

// flattenForwardingRuleLoadBalancingSchemeEnumSlice flattens the contents of ForwardingRuleLoadBalancingSchemeEnum from a JSON
// response object.
func flattenForwardingRuleLoadBalancingSchemeEnumSlice(c *Client, i interface{}) []ForwardingRuleLoadBalancingSchemeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ForwardingRuleLoadBalancingSchemeEnum{}
	}

	if len(a) == 0 {
		return []ForwardingRuleLoadBalancingSchemeEnum{}
	}

	items := make([]ForwardingRuleLoadBalancingSchemeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenForwardingRuleLoadBalancingSchemeEnum(item.(interface{})))
	}

	return items
}

// flattenForwardingRuleLoadBalancingSchemeEnum asserts that an interface is a string, and returns a
// pointer to a *ForwardingRuleLoadBalancingSchemeEnum with the same value as that string.
func flattenForwardingRuleLoadBalancingSchemeEnum(i interface{}) *ForwardingRuleLoadBalancingSchemeEnum {
	s, ok := i.(string)
	if !ok {
		return ForwardingRuleLoadBalancingSchemeEnumRef("")
	}

	return ForwardingRuleLoadBalancingSchemeEnumRef(s)
}

// flattenForwardingRuleMetadataFilterFilterMatchCriteriaEnumSlice flattens the contents of ForwardingRuleMetadataFilterFilterMatchCriteriaEnum from a JSON
// response object.
func flattenForwardingRuleMetadataFilterFilterMatchCriteriaEnumSlice(c *Client, i interface{}) []ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ForwardingRuleMetadataFilterFilterMatchCriteriaEnum{}
	}

	if len(a) == 0 {
		return []ForwardingRuleMetadataFilterFilterMatchCriteriaEnum{}
	}

	items := make([]ForwardingRuleMetadataFilterFilterMatchCriteriaEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenForwardingRuleMetadataFilterFilterMatchCriteriaEnum(item.(interface{})))
	}

	return items
}

// flattenForwardingRuleMetadataFilterFilterMatchCriteriaEnum asserts that an interface is a string, and returns a
// pointer to a *ForwardingRuleMetadataFilterFilterMatchCriteriaEnum with the same value as that string.
func flattenForwardingRuleMetadataFilterFilterMatchCriteriaEnum(i interface{}) *ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	s, ok := i.(string)
	if !ok {
		return ForwardingRuleMetadataFilterFilterMatchCriteriaEnumRef("")
	}

	return ForwardingRuleMetadataFilterFilterMatchCriteriaEnumRef(s)
}

// flattenForwardingRuleNetworkTierEnumSlice flattens the contents of ForwardingRuleNetworkTierEnum from a JSON
// response object.
func flattenForwardingRuleNetworkTierEnumSlice(c *Client, i interface{}) []ForwardingRuleNetworkTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ForwardingRuleNetworkTierEnum{}
	}

	if len(a) == 0 {
		return []ForwardingRuleNetworkTierEnum{}
	}

	items := make([]ForwardingRuleNetworkTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenForwardingRuleNetworkTierEnum(item.(interface{})))
	}

	return items
}

// flattenForwardingRuleNetworkTierEnum asserts that an interface is a string, and returns a
// pointer to a *ForwardingRuleNetworkTierEnum with the same value as that string.
func flattenForwardingRuleNetworkTierEnum(i interface{}) *ForwardingRuleNetworkTierEnum {
	s, ok := i.(string)
	if !ok {
		return ForwardingRuleNetworkTierEnumRef("")
	}

	return ForwardingRuleNetworkTierEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ForwardingRule) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalForwardingRule(b, c)
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
