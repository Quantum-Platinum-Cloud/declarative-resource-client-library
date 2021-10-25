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

func (r *Service) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "template"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.BinaryAuthorization) {
		if err := r.BinaryAuthorization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Template) {
		if err := r.Template.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TerminalCondition) {
		if err := r.TerminalCondition.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceBinaryAuthorization) validate() error {
	return nil
}
func (r *ServiceTemplate) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Scaling) {
		if err := r.Scaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.VPCAccess) {
		if err := r.VPCAccess.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateScaling) validate() error {
	return nil
}
func (r *ServiceTemplateVPCAccess) validate() error {
	return nil
}
func (r *ServiceTemplateContainers) validate() error {
	if err := dcl.Required(r, "image"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Resources) {
		if err := r.Resources.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateContainersEnv) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Value", "ValueSource"}, r.Value, r.ValueSource); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ValueSource) {
		if err := r.ValueSource.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateContainersEnvValueSource) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SecretKeyRef) {
		if err := r.SecretKeyRef.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateContainersEnvValueSourceSecretKeyRef) validate() error {
	if err := dcl.Required(r, "secret"); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTemplateContainersResources) validate() error {
	return nil
}
func (r *ServiceTemplateContainersPorts) validate() error {
	return nil
}
func (r *ServiceTemplateContainersVolumeMounts) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "mountPath"); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTemplateVolumes) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Secret", "CloudSqlInstance"}, r.Secret, r.CloudSqlInstance); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Secret) {
		if err := r.Secret.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudSqlInstance) {
		if err := r.CloudSqlInstance.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateVolumesSecret) validate() error {
	if err := dcl.Required(r, "secret"); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTemplateVolumesSecretItems) validate() error {
	if err := dcl.Required(r, "path"); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTemplateVolumesCloudSqlInstance) validate() error {
	return nil
}
func (r *ServiceTraffic) validate() error {
	return nil
}
func (r *ServiceGooglecloudrunopv2Condition) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Reason", "InternalReason", "DomainMappingReason", "RevisionReason", "JobReason"}, r.Reason, r.InternalReason, r.DomainMappingReason, r.RevisionReason, r.JobReason); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTrafficStatuses) validate() error {
	return nil
}
func (r *Service) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://run.googleapis.com/v2/", params)
}

func (r *Service) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Service) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services", nr.basePath(), userBasePath, params), nil

}

func (r *Service) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services?serviceId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Service) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Service) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}", nr.basePath(), userBasePath, fields)
}

func (r *Service) SetPolicyVerb() string {
	return "POST"
}

func (r *Service) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}", nr.basePath(), userBasePath, fields)
}

func (r *Service) IAMPolicyVersion() int {
	return 3
}

// serviceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type serviceApiOperation interface {
	do(context.Context, *Service, *Client) error
}

// newUpdateServiceUpdateServiceRequest creates a request for an
// Service resource's UpdateService update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateServiceUpdateServiceRequest(ctx context.Context, f *Service, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/locations/%s/services/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		req["annotations"] = v
	}
	if v := f.Client; !dcl.IsEmptyValueIndirect(v) {
		req["client"] = v
	}
	if v := f.ClientVersion; !dcl.IsEmptyValueIndirect(v) {
		req["clientVersion"] = v
	}
	if v := f.Ingress; !dcl.IsEmptyValueIndirect(v) {
		req["ingress"] = v
	}
	if v := f.LaunchStage; !dcl.IsEmptyValueIndirect(v) {
		req["launchStage"] = v
	}
	if v, err := expandServiceBinaryAuthorization(c, f.BinaryAuthorization); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into binaryAuthorization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["binaryAuthorization"] = v
	}
	if v, err := expandServiceTemplate(c, f.Template); err != nil {
		return nil, fmt.Errorf("error expanding Template into template: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["template"] = v
	}
	if v, err := expandServiceTrafficSlice(c, f.Traffic); err != nil {
		return nil, fmt.Errorf("error expanding Traffic into traffic: %w", err)
	} else if v != nil {
		req["traffic"] = v
	}
	if v, err := expandServiceGooglecloudrunopv2Condition(c, f.TerminalCondition); err != nil {
		return nil, fmt.Errorf("error expanding TerminalCondition into terminalCondition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["terminalCondition"] = v
	}
	b, err := c.getServiceRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	return req, nil
}

// marshalUpdateServiceUpdateServiceRequest converts the update into
// the final JSON request body.
func marshalUpdateServiceUpdateServiceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateServiceUpdateServiceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServiceUpdateServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	_, err := c.GetService(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateService")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateServiceUpdateServiceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateServiceUpdateServiceRequest(c, req)
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

func (c *Client) listServiceRaw(ctx context.Context, r *Service, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ServiceMaxPage {
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

type listServiceOperation struct {
	Services []map[string]interface{} `json:"services"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listService(ctx context.Context, r *Service, pageToken string, pageSize int32) ([]*Service, string, error) {
	b, err := c.listServiceRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServiceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Service
	for _, v := range m.Services {
		res, err := unmarshalMapService(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllService(ctx context.Context, f func(*Service) bool, resources []*Service) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteService(ctx, res)
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

type deleteServiceOperation struct{}

func (op *deleteServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	r, err := c.GetService(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Service not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetService checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Service: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createServiceOperation struct {
	response map[string]interface{}
}

func (op *createServiceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
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

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string in %v, was %T", op.response, op.response["name"])
	}
	r.Name = &name

	if _, err := c.GetService(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getServiceRaw(ctx context.Context, r *Service) ([]byte, error) {

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

func (c *Client) serviceDiffsForRawDesired(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (initial, desired *Service, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Service
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Service); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Service, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeServiceDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetService(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Service resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Service resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Service resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Service: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Service: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServiceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Service: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Service: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffService(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServiceInitialState(rawInitial, rawDesired *Service) (*Service, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServiceDesiredState(rawDesired, rawInitial *Service, opts ...dcl.ApplyOption) (*Service, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.BinaryAuthorization = canonicalizeServiceBinaryAuthorization(rawDesired.BinaryAuthorization, nil, opts...)
		rawDesired.Template = canonicalizeServiceTemplate(rawDesired.Template, nil, opts...)
		rawDesired.TerminalCondition = canonicalizeServiceGooglecloudrunopv2Condition(rawDesired.TerminalCondition, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Service{}
	if dcl.IsZeroValue(rawDesired.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Annotations) {
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
	}
	if dcl.StringCanonicalize(rawDesired.Client, rawInitial.Client) {
		canonicalDesired.Client = rawInitial.Client
	} else {
		canonicalDesired.Client = rawDesired.Client
	}
	if dcl.StringCanonicalize(rawDesired.ClientVersion, rawInitial.ClientVersion) {
		canonicalDesired.ClientVersion = rawInitial.ClientVersion
	} else {
		canonicalDesired.ClientVersion = rawDesired.ClientVersion
	}
	if dcl.IsZeroValue(rawDesired.Ingress) {
		canonicalDesired.Ingress = rawInitial.Ingress
	} else {
		canonicalDesired.Ingress = rawDesired.Ingress
	}
	if dcl.IsZeroValue(rawDesired.LaunchStage) {
		canonicalDesired.LaunchStage = rawInitial.LaunchStage
	} else {
		canonicalDesired.LaunchStage = rawDesired.LaunchStage
	}
	canonicalDesired.BinaryAuthorization = canonicalizeServiceBinaryAuthorization(rawDesired.BinaryAuthorization, rawInitial.BinaryAuthorization, opts...)
	canonicalDesired.Template = canonicalizeServiceTemplate(rawDesired.Template, rawInitial.Template, opts...)
	canonicalDesired.Traffic = canonicalizeServiceTrafficSlice(rawDesired.Traffic, rawInitial.Traffic, opts...)
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

func canonicalizeServiceNewState(c *Client, rawNew, rawDesired *Service) (*Service, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Uid) && dcl.IsNotReturnedByServer(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Generation) && dcl.IsNotReturnedByServer(rawDesired.Generation) {
		rawNew.Generation = rawDesired.Generation
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Labels) && dcl.IsNotReturnedByServer(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Annotations) && dcl.IsNotReturnedByServer(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.CreateTime) && dcl.IsNotReturnedByServer(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.DeleteTime) && dcl.IsNotReturnedByServer(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.ExpireTime) && dcl.IsNotReturnedByServer(rawDesired.ExpireTime) {
		rawNew.ExpireTime = rawDesired.ExpireTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Creator) && dcl.IsNotReturnedByServer(rawDesired.Creator) {
		rawNew.Creator = rawDesired.Creator
	} else {
		if dcl.StringCanonicalize(rawDesired.Creator, rawNew.Creator) {
			rawNew.Creator = rawDesired.Creator
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.LastModifier) && dcl.IsNotReturnedByServer(rawDesired.LastModifier) {
		rawNew.LastModifier = rawDesired.LastModifier
	} else {
		if dcl.StringCanonicalize(rawDesired.LastModifier, rawNew.LastModifier) {
			rawNew.LastModifier = rawDesired.LastModifier
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Client) && dcl.IsNotReturnedByServer(rawDesired.Client) {
		rawNew.Client = rawDesired.Client
	} else {
		if dcl.StringCanonicalize(rawDesired.Client, rawNew.Client) {
			rawNew.Client = rawDesired.Client
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.ClientVersion) && dcl.IsNotReturnedByServer(rawDesired.ClientVersion) {
		rawNew.ClientVersion = rawDesired.ClientVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.ClientVersion, rawNew.ClientVersion) {
			rawNew.ClientVersion = rawDesired.ClientVersion
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Ingress) && dcl.IsNotReturnedByServer(rawDesired.Ingress) {
		rawNew.Ingress = rawDesired.Ingress
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.LaunchStage) && dcl.IsNotReturnedByServer(rawDesired.LaunchStage) {
		rawNew.LaunchStage = rawDesired.LaunchStage
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.BinaryAuthorization) && dcl.IsNotReturnedByServer(rawDesired.BinaryAuthorization) {
		rawNew.BinaryAuthorization = rawDesired.BinaryAuthorization
	} else {
		rawNew.BinaryAuthorization = canonicalizeNewServiceBinaryAuthorization(c, rawDesired.BinaryAuthorization, rawNew.BinaryAuthorization)
	}

	if dcl.IsNotReturnedByServer(rawNew.Template) && dcl.IsNotReturnedByServer(rawDesired.Template) {
		rawNew.Template = rawDesired.Template
	} else {
		rawNew.Template = canonicalizeNewServiceTemplate(c, rawDesired.Template, rawNew.Template)
	}

	if dcl.IsNotReturnedByServer(rawNew.Traffic) && dcl.IsNotReturnedByServer(rawDesired.Traffic) {
		rawNew.Traffic = rawDesired.Traffic
	} else {
		rawNew.Traffic = canonicalizeNewServiceTrafficSlice(c, rawDesired.Traffic, rawNew.Traffic)
	}

	if dcl.IsNotReturnedByServer(rawNew.ObservedGeneration) && dcl.IsNotReturnedByServer(rawDesired.ObservedGeneration) {
		rawNew.ObservedGeneration = rawDesired.ObservedGeneration
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.TerminalCondition) && dcl.IsNotReturnedByServer(rawDesired.TerminalCondition) {
		rawNew.TerminalCondition = rawDesired.TerminalCondition
	} else {
		rawNew.TerminalCondition = canonicalizeNewServiceGooglecloudrunopv2Condition(c, rawDesired.TerminalCondition, rawNew.TerminalCondition)
	}

	if dcl.IsNotReturnedByServer(rawNew.Conditions) && dcl.IsNotReturnedByServer(rawDesired.Conditions) {
		rawNew.Conditions = rawDesired.Conditions
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.LatestReadyRevision) && dcl.IsNotReturnedByServer(rawDesired.LatestReadyRevision) {
		rawNew.LatestReadyRevision = rawDesired.LatestReadyRevision
	} else {
		if dcl.NameToSelfLink(rawDesired.LatestReadyRevision, rawNew.LatestReadyRevision) {
			rawNew.LatestReadyRevision = rawDesired.LatestReadyRevision
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.LatestCreatedRevision) && dcl.IsNotReturnedByServer(rawDesired.LatestCreatedRevision) {
		rawNew.LatestCreatedRevision = rawDesired.LatestCreatedRevision
	} else {
		if dcl.NameToSelfLink(rawDesired.LatestCreatedRevision, rawNew.LatestCreatedRevision) {
			rawNew.LatestCreatedRevision = rawDesired.LatestCreatedRevision
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.TrafficStatuses) && dcl.IsNotReturnedByServer(rawDesired.TrafficStatuses) {
		rawNew.TrafficStatuses = rawDesired.TrafficStatuses
	} else {
		rawNew.TrafficStatuses = canonicalizeNewServiceTrafficStatusesSlice(c, rawDesired.TrafficStatuses, rawNew.TrafficStatuses)
	}

	if dcl.IsNotReturnedByServer(rawNew.Uri) && dcl.IsNotReturnedByServer(rawDesired.Uri) {
		rawNew.Uri = rawDesired.Uri
	} else {
		if dcl.StringCanonicalize(rawDesired.Uri, rawNew.Uri) {
			rawNew.Uri = rawDesired.Uri
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Reconciling) && dcl.IsNotReturnedByServer(rawDesired.Reconciling) {
		rawNew.Reconciling = rawDesired.Reconciling
	} else {
		if dcl.BoolCanonicalize(rawDesired.Reconciling, rawNew.Reconciling) {
			rawNew.Reconciling = rawDesired.Reconciling
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Etag) && dcl.IsNotReturnedByServer(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeServiceBinaryAuthorization(des, initial *ServiceBinaryAuthorization, opts ...dcl.ApplyOption) *ServiceBinaryAuthorization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceBinaryAuthorization{}

	if dcl.StringCanonicalize(des.Policy, initial.Policy) || dcl.IsZeroValue(des.Policy) {
		cDes.Policy = initial.Policy
	} else {
		cDes.Policy = des.Policy
	}
	if dcl.StringCanonicalize(des.BreakglassJustification, initial.BreakglassJustification) || dcl.IsZeroValue(des.BreakglassJustification) {
		cDes.BreakglassJustification = initial.BreakglassJustification
	} else {
		cDes.BreakglassJustification = des.BreakglassJustification
	}

	return cDes
}

func canonicalizeServiceBinaryAuthorizationSlice(des, initial []ServiceBinaryAuthorization, opts ...dcl.ApplyOption) []ServiceBinaryAuthorization {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceBinaryAuthorization, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceBinaryAuthorization(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceBinaryAuthorization, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceBinaryAuthorization(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceBinaryAuthorization(c *Client, des, nw *ServiceBinaryAuthorization) *ServiceBinaryAuthorization {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceBinaryAuthorization while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Policy, nw.Policy) {
		nw.Policy = des.Policy
	}
	if dcl.StringCanonicalize(des.BreakglassJustification, nw.BreakglassJustification) {
		nw.BreakglassJustification = des.BreakglassJustification
	}

	return nw
}

func canonicalizeNewServiceBinaryAuthorizationSet(c *Client, des, nw []ServiceBinaryAuthorization) []ServiceBinaryAuthorization {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceBinaryAuthorization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceBinaryAuthorizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceBinaryAuthorizationSlice(c *Client, des, nw []ServiceBinaryAuthorization) []ServiceBinaryAuthorization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceBinaryAuthorization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceBinaryAuthorization(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplate(des, initial *ServiceTemplate, opts ...dcl.ApplyOption) *ServiceTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplate{}

	if dcl.NameToSelfLink(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	} else {
		cDes.Labels = des.Labels
	}
	if dcl.IsZeroValue(des.Annotations) {
		des.Annotations = initial.Annotations
	} else {
		cDes.Annotations = des.Annotations
	}
	cDes.Scaling = canonicalizeServiceTemplateScaling(des.Scaling, initial.Scaling, opts...)
	cDes.VPCAccess = canonicalizeServiceTemplateVPCAccess(des.VPCAccess, initial.VPCAccess, opts...)
	if dcl.IsZeroValue(des.ContainerConcurrency) {
		des.ContainerConcurrency = initial.ContainerConcurrency
	} else {
		cDes.ContainerConcurrency = des.ContainerConcurrency
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		cDes.Timeout = initial.Timeout
	} else {
		cDes.Timeout = des.Timeout
	}
	if dcl.StringCanonicalize(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		cDes.ServiceAccount = initial.ServiceAccount
	} else {
		cDes.ServiceAccount = des.ServiceAccount
	}
	cDes.Containers = canonicalizeServiceTemplateContainersSlice(des.Containers, initial.Containers, opts...)
	cDes.Volumes = canonicalizeServiceTemplateVolumesSlice(des.Volumes, initial.Volumes, opts...)
	if dcl.BoolCanonicalize(des.Confidential, initial.Confidential) || dcl.IsZeroValue(des.Confidential) {
		cDes.Confidential = initial.Confidential
	} else {
		cDes.Confidential = des.Confidential
	}
	if dcl.IsZeroValue(des.ExecutionEnvironment) {
		des.ExecutionEnvironment = initial.ExecutionEnvironment
	} else {
		cDes.ExecutionEnvironment = des.ExecutionEnvironment
	}

	return cDes
}

func canonicalizeServiceTemplateSlice(des, initial []ServiceTemplate, opts ...dcl.ApplyOption) []ServiceTemplate {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplate(c *Client, des, nw *ServiceTemplate) *ServiceTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	nw.Scaling = canonicalizeNewServiceTemplateScaling(c, des.Scaling, nw.Scaling)
	nw.VPCAccess = canonicalizeNewServiceTemplateVPCAccess(c, des.VPCAccess, nw.VPCAccess)
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	if dcl.StringCanonicalize(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	nw.Containers = canonicalizeNewServiceTemplateContainersSlice(c, des.Containers, nw.Containers)
	nw.Volumes = canonicalizeNewServiceTemplateVolumesSlice(c, des.Volumes, nw.Volumes)
	if dcl.BoolCanonicalize(des.Confidential, nw.Confidential) {
		nw.Confidential = des.Confidential
	}

	return nw
}

func canonicalizeNewServiceTemplateSet(c *Client, des, nw []ServiceTemplate) []ServiceTemplate {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplate
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateSlice(c *Client, des, nw []ServiceTemplate) []ServiceTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateScaling(des, initial *ServiceTemplateScaling, opts ...dcl.ApplyOption) *ServiceTemplateScaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateScaling{}

	if dcl.IsZeroValue(des.MinInstanceCount) {
		des.MinInstanceCount = initial.MinInstanceCount
	} else {
		cDes.MinInstanceCount = des.MinInstanceCount
	}
	if dcl.IsZeroValue(des.MaxInstanceCount) {
		des.MaxInstanceCount = initial.MaxInstanceCount
	} else {
		cDes.MaxInstanceCount = des.MaxInstanceCount
	}

	return cDes
}

func canonicalizeServiceTemplateScalingSlice(des, initial []ServiceTemplateScaling, opts ...dcl.ApplyOption) []ServiceTemplateScaling {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateScaling, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateScaling(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateScaling, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateScaling(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateScaling(c *Client, des, nw *ServiceTemplateScaling) *ServiceTemplateScaling {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateScaling while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewServiceTemplateScalingSet(c *Client, des, nw []ServiceTemplateScaling) []ServiceTemplateScaling {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateScaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateScalingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateScalingSlice(c *Client, des, nw []ServiceTemplateScaling) []ServiceTemplateScaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateScaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateScaling(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVPCAccess(des, initial *ServiceTemplateVPCAccess, opts ...dcl.ApplyOption) *ServiceTemplateVPCAccess {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVPCAccess{}

	if dcl.NameToSelfLink(des.Connector, initial.Connector) || dcl.IsZeroValue(des.Connector) {
		cDes.Connector = initial.Connector
	} else {
		cDes.Connector = des.Connector
	}
	if dcl.IsZeroValue(des.Egress) {
		des.Egress = initial.Egress
	} else {
		cDes.Egress = des.Egress
	}

	return cDes
}

func canonicalizeServiceTemplateVPCAccessSlice(des, initial []ServiceTemplateVPCAccess, opts ...dcl.ApplyOption) []ServiceTemplateVPCAccess {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVPCAccess, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVPCAccess(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVPCAccess, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVPCAccess(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVPCAccess(c *Client, des, nw *ServiceTemplateVPCAccess) *ServiceTemplateVPCAccess {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVPCAccess while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.Connector, nw.Connector) {
		nw.Connector = des.Connector
	}

	return nw
}

func canonicalizeNewServiceTemplateVPCAccessSet(c *Client, des, nw []ServiceTemplateVPCAccess) []ServiceTemplateVPCAccess {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateVPCAccess
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateVPCAccessNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateVPCAccessSlice(c *Client, des, nw []ServiceTemplateVPCAccess) []ServiceTemplateVPCAccess {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVPCAccess
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVPCAccess(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainers(des, initial *ServiceTemplateContainers, opts ...dcl.ApplyOption) *ServiceTemplateContainers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainers{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Image, initial.Image) || dcl.IsZeroValue(des.Image) {
		cDes.Image = initial.Image
	} else {
		cDes.Image = des.Image
	}
	if dcl.StringArrayCanonicalize(des.Command, initial.Command) || dcl.IsZeroValue(des.Command) {
		cDes.Command = initial.Command
	} else {
		cDes.Command = des.Command
	}
	if dcl.StringArrayCanonicalize(des.Args, initial.Args) || dcl.IsZeroValue(des.Args) {
		cDes.Args = initial.Args
	} else {
		cDes.Args = des.Args
	}
	cDes.Env = canonicalizeServiceTemplateContainersEnvSlice(des.Env, initial.Env, opts...)
	cDes.Resources = canonicalizeServiceTemplateContainersResources(des.Resources, initial.Resources, opts...)
	cDes.Ports = canonicalizeServiceTemplateContainersPortsSlice(des.Ports, initial.Ports, opts...)
	cDes.VolumeMounts = canonicalizeServiceTemplateContainersVolumeMountsSlice(des.VolumeMounts, initial.VolumeMounts, opts...)

	return cDes
}

func canonicalizeServiceTemplateContainersSlice(des, initial []ServiceTemplateContainers, opts ...dcl.ApplyOption) []ServiceTemplateContainers {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainers, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainers(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainers, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainers(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainers(c *Client, des, nw *ServiceTemplateContainers) *ServiceTemplateContainers {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainers while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Image, nw.Image) {
		nw.Image = des.Image
	}
	if dcl.StringArrayCanonicalize(des.Command, nw.Command) {
		nw.Command = des.Command
	}
	if dcl.StringArrayCanonicalize(des.Args, nw.Args) {
		nw.Args = des.Args
	}
	nw.Env = canonicalizeNewServiceTemplateContainersEnvSlice(c, des.Env, nw.Env)
	nw.Resources = canonicalizeNewServiceTemplateContainersResources(c, des.Resources, nw.Resources)
	nw.Ports = canonicalizeNewServiceTemplateContainersPortsSlice(c, des.Ports, nw.Ports)
	nw.VolumeMounts = canonicalizeNewServiceTemplateContainersVolumeMountsSlice(c, des.VolumeMounts, nw.VolumeMounts)

	return nw
}

func canonicalizeNewServiceTemplateContainersSet(c *Client, des, nw []ServiceTemplateContainers) []ServiceTemplateContainers {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateContainers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateContainersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateContainersSlice(c *Client, des, nw []ServiceTemplateContainers) []ServiceTemplateContainers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainers(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersEnv(des, initial *ServiceTemplateContainersEnv, opts ...dcl.ApplyOption) *ServiceTemplateContainersEnv {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Value != nil || (initial != nil && initial.Value != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ValueSource) {
			des.Value = nil
			if initial != nil {
				initial.Value = nil
			}
		}
	}

	if des.ValueSource != nil || (initial != nil && initial.ValueSource != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Value) {
			des.ValueSource = nil
			if initial != nil {
				initial.ValueSource = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersEnv{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}
	cDes.ValueSource = canonicalizeServiceTemplateContainersEnvValueSource(des.ValueSource, initial.ValueSource, opts...)

	return cDes
}

func canonicalizeServiceTemplateContainersEnvSlice(des, initial []ServiceTemplateContainersEnv, opts ...dcl.ApplyOption) []ServiceTemplateContainersEnv {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersEnv, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersEnv(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersEnv, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersEnv(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersEnv(c *Client, des, nw *ServiceTemplateContainersEnv) *ServiceTemplateContainersEnv {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersEnv while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}
	nw.ValueSource = canonicalizeNewServiceTemplateContainersEnvValueSource(c, des.ValueSource, nw.ValueSource)

	return nw
}

func canonicalizeNewServiceTemplateContainersEnvSet(c *Client, des, nw []ServiceTemplateContainersEnv) []ServiceTemplateContainersEnv {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateContainersEnv
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateContainersEnvNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateContainersEnvSlice(c *Client, des, nw []ServiceTemplateContainersEnv) []ServiceTemplateContainersEnv {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersEnv
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersEnv(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersEnvValueSource(des, initial *ServiceTemplateContainersEnvValueSource, opts ...dcl.ApplyOption) *ServiceTemplateContainersEnvValueSource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersEnvValueSource{}

	cDes.SecretKeyRef = canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRef(des.SecretKeyRef, initial.SecretKeyRef, opts...)

	return cDes
}

func canonicalizeServiceTemplateContainersEnvValueSourceSlice(des, initial []ServiceTemplateContainersEnvValueSource, opts ...dcl.ApplyOption) []ServiceTemplateContainersEnvValueSource {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersEnvValueSource, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersEnvValueSource(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersEnvValueSource, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersEnvValueSource(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersEnvValueSource(c *Client, des, nw *ServiceTemplateContainersEnvValueSource) *ServiceTemplateContainersEnvValueSource {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersEnvValueSource while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.SecretKeyRef = canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRef(c, des.SecretKeyRef, nw.SecretKeyRef)

	return nw
}

func canonicalizeNewServiceTemplateContainersEnvValueSourceSet(c *Client, des, nw []ServiceTemplateContainersEnvValueSource) []ServiceTemplateContainersEnvValueSource {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateContainersEnvValueSource
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateContainersEnvValueSourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateContainersEnvValueSourceSlice(c *Client, des, nw []ServiceTemplateContainersEnvValueSource) []ServiceTemplateContainersEnvValueSource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersEnvValueSource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersEnvValueSource(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRef(des, initial *ServiceTemplateContainersEnvValueSourceSecretKeyRef, opts ...dcl.ApplyOption) *ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersEnvValueSourceSecretKeyRef{}

	if dcl.NameToSelfLink(des.Secret, initial.Secret) || dcl.IsZeroValue(des.Secret) {
		cDes.Secret = initial.Secret
	} else {
		cDes.Secret = des.Secret
	}
	if dcl.NameToSelfLink(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRefSlice(des, initial []ServiceTemplateContainersEnvValueSourceSecretKeyRef, opts ...dcl.ApplyOption) []ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersEnvValueSourceSecretKeyRef, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRef(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersEnvValueSourceSecretKeyRef, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRef(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRef(c *Client, des, nw *ServiceTemplateContainersEnvValueSourceSecretKeyRef) *ServiceTemplateContainersEnvValueSourceSecretKeyRef {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersEnvValueSourceSecretKeyRef while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.Secret, nw.Secret) {
		nw.Secret = des.Secret
	}
	if dcl.NameToSelfLink(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRefSet(c *Client, des, nw []ServiceTemplateContainersEnvValueSourceSecretKeyRef) []ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateContainersEnvValueSourceSecretKeyRef
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateContainersEnvValueSourceSecretKeyRefNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, des, nw []ServiceTemplateContainersEnvValueSourceSecretKeyRef) []ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersEnvValueSourceSecretKeyRef
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRef(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersResources(des, initial *ServiceTemplateContainersResources, opts ...dcl.ApplyOption) *ServiceTemplateContainersResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersResources{}

	if dcl.IsZeroValue(des.Limits) {
		des.Limits = initial.Limits
	} else {
		cDes.Limits = des.Limits
	}
	if dcl.BoolCanonicalize(des.CpuIdle, initial.CpuIdle) || dcl.IsZeroValue(des.CpuIdle) {
		cDes.CpuIdle = initial.CpuIdle
	} else {
		cDes.CpuIdle = des.CpuIdle
	}

	return cDes
}

func canonicalizeServiceTemplateContainersResourcesSlice(des, initial []ServiceTemplateContainersResources, opts ...dcl.ApplyOption) []ServiceTemplateContainersResources {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersResources(c *Client, des, nw *ServiceTemplateContainersResources) *ServiceTemplateContainersResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.CpuIdle, nw.CpuIdle) {
		nw.CpuIdle = des.CpuIdle
	}

	return nw
}

func canonicalizeNewServiceTemplateContainersResourcesSet(c *Client, des, nw []ServiceTemplateContainersResources) []ServiceTemplateContainersResources {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateContainersResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateContainersResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateContainersResourcesSlice(c *Client, des, nw []ServiceTemplateContainersResources) []ServiceTemplateContainersResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersResources(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersPorts(des, initial *ServiceTemplateContainersPorts, opts ...dcl.ApplyOption) *ServiceTemplateContainersPorts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersPorts{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.ContainerPort) {
		des.ContainerPort = initial.ContainerPort
	} else {
		cDes.ContainerPort = des.ContainerPort
	}

	return cDes
}

func canonicalizeServiceTemplateContainersPortsSlice(des, initial []ServiceTemplateContainersPorts, opts ...dcl.ApplyOption) []ServiceTemplateContainersPorts {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersPorts, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersPorts(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersPorts, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersPorts(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersPorts(c *Client, des, nw *ServiceTemplateContainersPorts) *ServiceTemplateContainersPorts {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersPorts while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceTemplateContainersPortsSet(c *Client, des, nw []ServiceTemplateContainersPorts) []ServiceTemplateContainersPorts {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateContainersPorts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateContainersPortsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateContainersPortsSlice(c *Client, des, nw []ServiceTemplateContainersPorts) []ServiceTemplateContainersPorts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersPorts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersPorts(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersVolumeMounts(des, initial *ServiceTemplateContainersVolumeMounts, opts ...dcl.ApplyOption) *ServiceTemplateContainersVolumeMounts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersVolumeMounts{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.MountPath, initial.MountPath) || dcl.IsZeroValue(des.MountPath) {
		cDes.MountPath = initial.MountPath
	} else {
		cDes.MountPath = des.MountPath
	}

	return cDes
}

func canonicalizeServiceTemplateContainersVolumeMountsSlice(des, initial []ServiceTemplateContainersVolumeMounts, opts ...dcl.ApplyOption) []ServiceTemplateContainersVolumeMounts {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersVolumeMounts, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersVolumeMounts(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersVolumeMounts, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersVolumeMounts(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersVolumeMounts(c *Client, des, nw *ServiceTemplateContainersVolumeMounts) *ServiceTemplateContainersVolumeMounts {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersVolumeMounts while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.MountPath, nw.MountPath) {
		nw.MountPath = des.MountPath
	}

	return nw
}

func canonicalizeNewServiceTemplateContainersVolumeMountsSet(c *Client, des, nw []ServiceTemplateContainersVolumeMounts) []ServiceTemplateContainersVolumeMounts {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateContainersVolumeMounts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateContainersVolumeMountsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateContainersVolumeMountsSlice(c *Client, des, nw []ServiceTemplateContainersVolumeMounts) []ServiceTemplateContainersVolumeMounts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersVolumeMounts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersVolumeMounts(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVolumes(des, initial *ServiceTemplateVolumes, opts ...dcl.ApplyOption) *ServiceTemplateVolumes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Secret != nil || (initial != nil && initial.Secret != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudSqlInstance) {
			des.Secret = nil
			if initial != nil {
				initial.Secret = nil
			}
		}
	}

	if des.CloudSqlInstance != nil || (initial != nil && initial.CloudSqlInstance != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Secret) {
			des.CloudSqlInstance = nil
			if initial != nil {
				initial.CloudSqlInstance = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVolumes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.Secret = canonicalizeServiceTemplateVolumesSecret(des.Secret, initial.Secret, opts...)
	cDes.CloudSqlInstance = canonicalizeServiceTemplateVolumesCloudSqlInstance(des.CloudSqlInstance, initial.CloudSqlInstance, opts...)

	return cDes
}

func canonicalizeServiceTemplateVolumesSlice(des, initial []ServiceTemplateVolumes, opts ...dcl.ApplyOption) []ServiceTemplateVolumes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVolumes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVolumes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVolumes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVolumes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVolumes(c *Client, des, nw *ServiceTemplateVolumes) *ServiceTemplateVolumes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVolumes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Secret = canonicalizeNewServiceTemplateVolumesSecret(c, des.Secret, nw.Secret)
	nw.CloudSqlInstance = canonicalizeNewServiceTemplateVolumesCloudSqlInstance(c, des.CloudSqlInstance, nw.CloudSqlInstance)

	return nw
}

func canonicalizeNewServiceTemplateVolumesSet(c *Client, des, nw []ServiceTemplateVolumes) []ServiceTemplateVolumes {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateVolumes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateVolumesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateVolumesSlice(c *Client, des, nw []ServiceTemplateVolumes) []ServiceTemplateVolumes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVolumes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVolumes(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVolumesSecret(des, initial *ServiceTemplateVolumesSecret, opts ...dcl.ApplyOption) *ServiceTemplateVolumesSecret {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVolumesSecret{}

	if dcl.StringCanonicalize(des.Secret, initial.Secret) || dcl.IsZeroValue(des.Secret) {
		cDes.Secret = initial.Secret
	} else {
		cDes.Secret = des.Secret
	}
	cDes.Items = canonicalizeServiceTemplateVolumesSecretItemsSlice(des.Items, initial.Items, opts...)
	if dcl.IsZeroValue(des.DefaultMode) {
		des.DefaultMode = initial.DefaultMode
	} else {
		cDes.DefaultMode = des.DefaultMode
	}

	return cDes
}

func canonicalizeServiceTemplateVolumesSecretSlice(des, initial []ServiceTemplateVolumesSecret, opts ...dcl.ApplyOption) []ServiceTemplateVolumesSecret {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVolumesSecret, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVolumesSecret(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVolumesSecret, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVolumesSecret(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVolumesSecret(c *Client, des, nw *ServiceTemplateVolumesSecret) *ServiceTemplateVolumesSecret {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVolumesSecret while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Secret, nw.Secret) {
		nw.Secret = des.Secret
	}
	nw.Items = canonicalizeNewServiceTemplateVolumesSecretItemsSlice(c, des.Items, nw.Items)

	return nw
}

func canonicalizeNewServiceTemplateVolumesSecretSet(c *Client, des, nw []ServiceTemplateVolumesSecret) []ServiceTemplateVolumesSecret {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateVolumesSecret
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateVolumesSecretNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateVolumesSecretSlice(c *Client, des, nw []ServiceTemplateVolumesSecret) []ServiceTemplateVolumesSecret {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVolumesSecret
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVolumesSecret(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVolumesSecretItems(des, initial *ServiceTemplateVolumesSecretItems, opts ...dcl.ApplyOption) *ServiceTemplateVolumesSecretItems {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVolumesSecretItems{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}
	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}

	return cDes
}

func canonicalizeServiceTemplateVolumesSecretItemsSlice(des, initial []ServiceTemplateVolumesSecretItems, opts ...dcl.ApplyOption) []ServiceTemplateVolumesSecretItems {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVolumesSecretItems, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVolumesSecretItems(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVolumesSecretItems, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVolumesSecretItems(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVolumesSecretItems(c *Client, des, nw *ServiceTemplateVolumesSecretItems) *ServiceTemplateVolumesSecretItems {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVolumesSecretItems while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewServiceTemplateVolumesSecretItemsSet(c *Client, des, nw []ServiceTemplateVolumesSecretItems) []ServiceTemplateVolumesSecretItems {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateVolumesSecretItems
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateVolumesSecretItemsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateVolumesSecretItemsSlice(c *Client, des, nw []ServiceTemplateVolumesSecretItems) []ServiceTemplateVolumesSecretItems {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVolumesSecretItems
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVolumesSecretItems(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVolumesCloudSqlInstance(des, initial *ServiceTemplateVolumesCloudSqlInstance, opts ...dcl.ApplyOption) *ServiceTemplateVolumesCloudSqlInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVolumesCloudSqlInstance{}

	if dcl.StringArrayCanonicalize(des.Connections, initial.Connections) || dcl.IsZeroValue(des.Connections) {
		cDes.Connections = initial.Connections
	} else {
		cDes.Connections = des.Connections
	}

	return cDes
}

func canonicalizeServiceTemplateVolumesCloudSqlInstanceSlice(des, initial []ServiceTemplateVolumesCloudSqlInstance, opts ...dcl.ApplyOption) []ServiceTemplateVolumesCloudSqlInstance {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVolumesCloudSqlInstance, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVolumesCloudSqlInstance(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVolumesCloudSqlInstance, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVolumesCloudSqlInstance(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVolumesCloudSqlInstance(c *Client, des, nw *ServiceTemplateVolumesCloudSqlInstance) *ServiceTemplateVolumesCloudSqlInstance {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVolumesCloudSqlInstance while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Connections, nw.Connections) {
		nw.Connections = des.Connections
	}

	return nw
}

func canonicalizeNewServiceTemplateVolumesCloudSqlInstanceSet(c *Client, des, nw []ServiceTemplateVolumesCloudSqlInstance) []ServiceTemplateVolumesCloudSqlInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTemplateVolumesCloudSqlInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTemplateVolumesCloudSqlInstanceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTemplateVolumesCloudSqlInstanceSlice(c *Client, des, nw []ServiceTemplateVolumesCloudSqlInstance) []ServiceTemplateVolumesCloudSqlInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVolumesCloudSqlInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVolumesCloudSqlInstance(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTraffic(des, initial *ServiceTraffic, opts ...dcl.ApplyOption) *ServiceTraffic {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTraffic{}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.NameToSelfLink(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	} else {
		cDes.Percent = des.Percent
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		cDes.Tag = initial.Tag
	} else {
		cDes.Tag = des.Tag
	}

	return cDes
}

func canonicalizeServiceTrafficSlice(des, initial []ServiceTraffic, opts ...dcl.ApplyOption) []ServiceTraffic {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTraffic, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTraffic(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTraffic, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTraffic(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTraffic(c *Client, des, nw *ServiceTraffic) *ServiceTraffic {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTraffic while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) {
		nw.Tag = des.Tag
	}

	return nw
}

func canonicalizeNewServiceTrafficSet(c *Client, des, nw []ServiceTraffic) []ServiceTraffic {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTraffic
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTrafficNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTrafficSlice(c *Client, des, nw []ServiceTraffic) []ServiceTraffic {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTraffic
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTraffic(c, &d, &n))
	}

	return items
}

func canonicalizeServiceGooglecloudrunopv2Condition(des, initial *ServiceGooglecloudrunopv2Condition, opts ...dcl.ApplyOption) *ServiceGooglecloudrunopv2Condition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Reason != nil || (initial != nil && initial.Reason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.InternalReason, des.DomainMappingReason, des.RevisionReason, des.JobReason) {
			des.Reason = nil
			if initial != nil {
				initial.Reason = nil
			}
		}
	}

	if des.InternalReason != nil || (initial != nil && initial.InternalReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.DomainMappingReason, des.RevisionReason, des.JobReason) {
			des.InternalReason = nil
			if initial != nil {
				initial.InternalReason = nil
			}
		}
	}

	if des.DomainMappingReason != nil || (initial != nil && initial.DomainMappingReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.InternalReason, des.RevisionReason, des.JobReason) {
			des.DomainMappingReason = nil
			if initial != nil {
				initial.DomainMappingReason = nil
			}
		}
	}

	if des.RevisionReason != nil || (initial != nil && initial.RevisionReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.InternalReason, des.DomainMappingReason, des.JobReason) {
			des.RevisionReason = nil
			if initial != nil {
				initial.RevisionReason = nil
			}
		}
	}

	if des.JobReason != nil || (initial != nil && initial.JobReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.InternalReason, des.DomainMappingReason, des.RevisionReason) {
			des.JobReason = nil
			if initial != nil {
				initial.JobReason = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceGooglecloudrunopv2Condition{}

	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	} else {
		cDes.State = des.State
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		cDes.Message = initial.Message
	} else {
		cDes.Message = des.Message
	}
	if dcl.IsZeroValue(des.LastTransitionTime) {
		des.LastTransitionTime = initial.LastTransitionTime
	} else {
		cDes.LastTransitionTime = des.LastTransitionTime
	}
	if dcl.IsZeroValue(des.Severity) {
		des.Severity = initial.Severity
	} else {
		cDes.Severity = des.Severity
	}
	if dcl.IsZeroValue(des.Reason) {
		des.Reason = initial.Reason
	} else {
		cDes.Reason = des.Reason
	}
	if dcl.IsZeroValue(des.InternalReason) {
		des.InternalReason = initial.InternalReason
	} else {
		cDes.InternalReason = des.InternalReason
	}
	if dcl.IsZeroValue(des.DomainMappingReason) {
		des.DomainMappingReason = initial.DomainMappingReason
	} else {
		cDes.DomainMappingReason = des.DomainMappingReason
	}
	if dcl.IsZeroValue(des.RevisionReason) {
		des.RevisionReason = initial.RevisionReason
	} else {
		cDes.RevisionReason = des.RevisionReason
	}
	if dcl.IsZeroValue(des.JobReason) {
		des.JobReason = initial.JobReason
	} else {
		cDes.JobReason = des.JobReason
	}

	return cDes
}

func canonicalizeServiceGooglecloudrunopv2ConditionSlice(des, initial []ServiceGooglecloudrunopv2Condition, opts ...dcl.ApplyOption) []ServiceGooglecloudrunopv2Condition {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceGooglecloudrunopv2Condition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceGooglecloudrunopv2Condition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceGooglecloudrunopv2Condition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceGooglecloudrunopv2Condition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceGooglecloudrunopv2Condition(c *Client, des, nw *ServiceGooglecloudrunopv2Condition) *ServiceGooglecloudrunopv2Condition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceGooglecloudrunopv2Condition while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}

	return nw
}

func canonicalizeNewServiceGooglecloudrunopv2ConditionSet(c *Client, des, nw []ServiceGooglecloudrunopv2Condition) []ServiceGooglecloudrunopv2Condition {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceGooglecloudrunopv2Condition
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceGooglecloudrunopv2ConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceGooglecloudrunopv2ConditionSlice(c *Client, des, nw []ServiceGooglecloudrunopv2Condition) []ServiceGooglecloudrunopv2Condition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceGooglecloudrunopv2Condition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceGooglecloudrunopv2Condition(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTrafficStatuses(des, initial *ServiceTrafficStatuses, opts ...dcl.ApplyOption) *ServiceTrafficStatuses {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTrafficStatuses{}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.NameToSelfLink(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	} else {
		cDes.Percent = des.Percent
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		cDes.Tag = initial.Tag
	} else {
		cDes.Tag = des.Tag
	}
	if dcl.StringCanonicalize(des.Uri, initial.Uri) || dcl.IsZeroValue(des.Uri) {
		cDes.Uri = initial.Uri
	} else {
		cDes.Uri = des.Uri
	}

	return cDes
}

func canonicalizeServiceTrafficStatusesSlice(des, initial []ServiceTrafficStatuses, opts ...dcl.ApplyOption) []ServiceTrafficStatuses {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTrafficStatuses, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTrafficStatuses(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTrafficStatuses, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTrafficStatuses(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTrafficStatuses(c *Client, des, nw *ServiceTrafficStatuses) *ServiceTrafficStatuses {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTrafficStatuses while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) {
		nw.Tag = des.Tag
	}
	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}

	return nw
}

func canonicalizeNewServiceTrafficStatusesSet(c *Client, des, nw []ServiceTrafficStatuses) []ServiceTrafficStatuses {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceTrafficStatuses
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceTrafficStatusesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceTrafficStatusesSlice(c *Client, des, nw []ServiceTrafficStatuses) []ServiceTrafficStatuses {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTrafficStatuses
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTrafficStatuses(c, &d, &n))
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
func diffService(c *Client, desired, actual *Service, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpireTime, actual.ExpireTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpireTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Creator, actual.Creator, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Creator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifier, actual.LastModifier, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Client, actual.Client, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Client")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientVersion, actual.ClientVersion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ClientVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ingress, actual.Ingress, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Ingress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LaunchStage, actual.LaunchStage, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("LaunchStage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BinaryAuthorization, actual.BinaryAuthorization, dcl.Info{ObjectFunction: compareServiceBinaryAuthorizationNewStyle, EmptyObject: EmptyServiceBinaryAuthorization, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("BinaryAuthorization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Template, actual.Template, dcl.Info{ObjectFunction: compareServiceTemplateNewStyle, EmptyObject: EmptyServiceTemplate, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Template")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Traffic, actual.Traffic, dcl.Info{ObjectFunction: compareServiceTrafficNewStyle, EmptyObject: EmptyServiceTraffic, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Traffic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ObservedGeneration, actual.ObservedGeneration, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObservedGeneration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TerminalCondition, actual.TerminalCondition, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceGooglecloudrunopv2ConditionNewStyle, EmptyObject: EmptyServiceGooglecloudrunopv2Condition, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TerminalCondition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Conditions, actual.Conditions, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceGooglecloudrunopv2ConditionNewStyle, EmptyObject: EmptyServiceGooglecloudrunopv2Condition, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Conditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestReadyRevision, actual.LatestReadyRevision, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LatestReadyRevision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestCreatedRevision, actual.LatestCreatedRevision, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LatestCreatedRevision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrafficStatuses, actual.TrafficStatuses, dcl.Info{OutputOnly: true, ObjectFunction: compareServiceTrafficStatusesNewStyle, EmptyObject: EmptyServiceTrafficStatuses, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TrafficStatuses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reconciling, actual.Reconciling, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Reconciling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareServiceBinaryAuthorizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceBinaryAuthorization)
	if !ok {
		desiredNotPointer, ok := d.(ServiceBinaryAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceBinaryAuthorization or *ServiceBinaryAuthorization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceBinaryAuthorization)
	if !ok {
		actualNotPointer, ok := a.(ServiceBinaryAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceBinaryAuthorization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Policy, actual.Policy, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Policy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BreakglassJustification, actual.BreakglassJustification, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("BreakglassJustification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplate)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplate or *ServiceTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplate)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scaling, actual.Scaling, dcl.Info{ObjectFunction: compareServiceTemplateScalingNewStyle, EmptyObject: EmptyServiceTemplateScaling, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Scaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VPCAccess, actual.VPCAccess, dcl.Info{ObjectFunction: compareServiceTemplateVPCAccessNewStyle, EmptyObject: EmptyServiceTemplateVPCAccess, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("VpcAccess")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerConcurrency, actual.ContainerConcurrency, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ContainerConcurrency")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Containers, actual.Containers, dcl.Info{ObjectFunction: compareServiceTemplateContainersNewStyle, EmptyObject: EmptyServiceTemplateContainers, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Containers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Volumes, actual.Volumes, dcl.Info{ObjectFunction: compareServiceTemplateVolumesNewStyle, EmptyObject: EmptyServiceTemplateVolumes, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Volumes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Confidential, actual.Confidential, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Confidential")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionEnvironment, actual.ExecutionEnvironment, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ExecutionEnvironment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateScalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateScaling)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateScaling or *ServiceTemplateScaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateScaling)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateScaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinInstanceCount, actual.MinInstanceCount, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MinInstanceCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxInstanceCount, actual.MaxInstanceCount, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MaxInstanceCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVPCAccessNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVPCAccess)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVPCAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVPCAccess or *ServiceTemplateVPCAccess", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVPCAccess)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVPCAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVPCAccess", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Connector, actual.Connector, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Connector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Egress, actual.Egress, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Egress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainers)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainers or *ServiceTemplateContainers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainers)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Image")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Env, actual.Env, dcl.Info{ObjectFunction: compareServiceTemplateContainersEnvNewStyle, EmptyObject: EmptyServiceTemplateContainersEnv, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Env")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.Info{ObjectFunction: compareServiceTemplateContainersResourcesNewStyle, EmptyObject: EmptyServiceTemplateContainersResources, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.Info{ObjectFunction: compareServiceTemplateContainersPortsNewStyle, EmptyObject: EmptyServiceTemplateContainersPorts, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VolumeMounts, actual.VolumeMounts, dcl.Info{ObjectFunction: compareServiceTemplateContainersVolumeMountsNewStyle, EmptyObject: EmptyServiceTemplateContainersVolumeMounts, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("VolumeMounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersEnvNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersEnv)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnv or *ServiceTemplateContainersEnv", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersEnv)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnv", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueSource, actual.ValueSource, dcl.Info{ObjectFunction: compareServiceTemplateContainersEnvValueSourceNewStyle, EmptyObject: EmptyServiceTemplateContainersEnvValueSource, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ValueSource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersEnvValueSourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersEnvValueSource)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersEnvValueSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnvValueSource or *ServiceTemplateContainersEnvValueSource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersEnvValueSource)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersEnvValueSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnvValueSource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SecretKeyRef, actual.SecretKeyRef, dcl.Info{ObjectFunction: compareServiceTemplateContainersEnvValueSourceSecretKeyRefNewStyle, EmptyObject: EmptyServiceTemplateContainersEnvValueSourceSecretKeyRef, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("SecretKeyRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersEnvValueSourceSecretKeyRefNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersEnvValueSourceSecretKeyRef)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersEnvValueSourceSecretKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnvValueSourceSecretKeyRef or *ServiceTemplateContainersEnvValueSourceSecretKeyRef", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersEnvValueSourceSecretKeyRef)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersEnvValueSourceSecretKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnvValueSourceSecretKeyRef", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersResources)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersResources or *ServiceTemplateContainersResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersResources)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Limits, actual.Limits, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Limits")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuIdle, actual.CpuIdle, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("CpuIdle")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersPortsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersPorts)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersPorts or *ServiceTemplateContainersPorts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersPorts)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersPorts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerPort, actual.ContainerPort, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ContainerPort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersVolumeMountsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersVolumeMounts)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersVolumeMounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersVolumeMounts or *ServiceTemplateContainersVolumeMounts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersVolumeMounts)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersVolumeMounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersVolumeMounts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MountPath, actual.MountPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MountPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVolumesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVolumes)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumes or *ServiceTemplateVolumes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVolumes)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.Info{ObjectFunction: compareServiceTemplateVolumesSecretNewStyle, EmptyObject: EmptyServiceTemplateVolumesSecret, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudSqlInstance, actual.CloudSqlInstance, dcl.Info{ObjectFunction: compareServiceTemplateVolumesCloudSqlInstanceNewStyle, EmptyObject: EmptyServiceTemplateVolumesCloudSqlInstance, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("CloudSqlInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVolumesSecretNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVolumesSecret)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVolumesSecret)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesSecret or *ServiceTemplateVolumesSecret", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVolumesSecret)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVolumesSecret)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesSecret", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Items, actual.Items, dcl.Info{ObjectFunction: compareServiceTemplateVolumesSecretItemsNewStyle, EmptyObject: EmptyServiceTemplateVolumesSecretItems, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Items")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultMode, actual.DefaultMode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("DefaultMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVolumesSecretItemsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVolumesSecretItems)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVolumesSecretItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesSecretItems or *ServiceTemplateVolumesSecretItems", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVolumesSecretItems)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVolumesSecretItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesSecretItems", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVolumesCloudSqlInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVolumesCloudSqlInstance)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVolumesCloudSqlInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesCloudSqlInstance or *ServiceTemplateVolumesCloudSqlInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVolumesCloudSqlInstance)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVolumesCloudSqlInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesCloudSqlInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Connections, actual.Connections, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Connections")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTrafficNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTraffic)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTraffic)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTraffic or *ServiceTraffic", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTraffic)
	if !ok {
		actualNotPointer, ok := a.(ServiceTraffic)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTraffic", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceGooglecloudrunopv2ConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceGooglecloudrunopv2Condition)
	if !ok {
		desiredNotPointer, ok := d.(ServiceGooglecloudrunopv2Condition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceGooglecloudrunopv2Condition or *ServiceGooglecloudrunopv2Condition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceGooglecloudrunopv2Condition)
	if !ok {
		actualNotPointer, ok := a.(ServiceGooglecloudrunopv2Condition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceGooglecloudrunopv2Condition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastTransitionTime, actual.LastTransitionTime, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("LastTransitionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Severity, actual.Severity, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Severity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reason, actual.Reason, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Reason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InternalReason, actual.InternalReason, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("InternalReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DomainMappingReason, actual.DomainMappingReason, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("DomainMappingReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionReason, actual.RevisionReason, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("RevisionReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.JobReason, actual.JobReason, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("JobReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTrafficStatusesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTrafficStatuses)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTrafficStatuses)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTrafficStatuses or *ServiceTrafficStatuses", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTrafficStatuses)
	if !ok {
		actualNotPointer, ok := a.(ServiceTrafficStatuses)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTrafficStatuses", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
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
func (r *Service) urlNormalized() *Service {
	normalized := dcl.Copy(*r).(Service)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Creator = dcl.SelfLinkToName(r.Creator)
	normalized.LastModifier = dcl.SelfLinkToName(r.LastModifier)
	normalized.Client = dcl.SelfLinkToName(r.Client)
	normalized.ClientVersion = dcl.SelfLinkToName(r.ClientVersion)
	normalized.LatestReadyRevision = dcl.SelfLinkToName(r.LatestReadyRevision)
	normalized.LatestCreatedRevision = dcl.SelfLinkToName(r.LatestCreatedRevision)
	normalized.Uri = dcl.SelfLinkToName(r.Uri)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Service) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateService" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Service resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Service) marshal(c *Client) ([]byte, error) {
	m, err := expandService(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Service: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalService decodes JSON responses into the Service resource schema.
func unmarshalService(b []byte, c *Client) (*Service, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapService(m, c)
}

func unmarshalMapService(m map[string]interface{}, c *Client) (*Service, error) {

	flattened := flattenService(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandService expands Service into a JSON request object.
func expandService(c *Client, f *Service) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/services/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Annotations; dcl.ValueShouldBeSent(v) {
		m["annotations"] = v
	}
	if v := f.Client; dcl.ValueShouldBeSent(v) {
		m["client"] = v
	}
	if v := f.ClientVersion; dcl.ValueShouldBeSent(v) {
		m["clientVersion"] = v
	}
	if v := f.Ingress; dcl.ValueShouldBeSent(v) {
		m["ingress"] = v
	}
	if v := f.LaunchStage; dcl.ValueShouldBeSent(v) {
		m["launchStage"] = v
	}
	if v, err := expandServiceBinaryAuthorization(c, f.BinaryAuthorization); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into binaryAuthorization: %w", err)
	} else if v != nil {
		m["binaryAuthorization"] = v
	}
	if v, err := expandServiceTemplate(c, f.Template); err != nil {
		return nil, fmt.Errorf("error expanding Template into template: %w", err)
	} else if v != nil {
		m["template"] = v
	}
	if v, err := expandServiceTrafficSlice(c, f.Traffic); err != nil {
		return nil, fmt.Errorf("error expanding Traffic into traffic: %w", err)
	} else {
		m["traffic"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenService flattens Service from a JSON request object into the
// Service type.
func flattenService(c *Client, i interface{}) *Service {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Service{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.Description = dcl.FlattenString(m["description"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.Generation = dcl.FlattenInteger(m["generation"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.DeleteTime = dcl.FlattenString(m["deleteTime"])
	res.ExpireTime = dcl.FlattenString(m["expireTime"])
	res.Creator = dcl.FlattenString(m["creator"])
	res.LastModifier = dcl.FlattenString(m["lastModifier"])
	res.Client = dcl.FlattenString(m["client"])
	res.ClientVersion = dcl.FlattenString(m["clientVersion"])
	res.Ingress = flattenServiceIngressEnum(m["ingress"])
	res.LaunchStage = flattenServiceLaunchStageEnum(m["launchStage"])
	res.BinaryAuthorization = flattenServiceBinaryAuthorization(c, m["binaryAuthorization"])
	res.Template = flattenServiceTemplate(c, m["template"])
	res.Traffic = flattenServiceTrafficSlice(c, m["traffic"])
	res.ObservedGeneration = dcl.FlattenInteger(m["observedGeneration"])
	res.TerminalCondition = flattenServiceGooglecloudrunopv2Condition(c, m["terminalCondition"])
	res.Conditions = flattenServiceGooglecloudrunopv2ConditionSlice(c, m["conditions"])
	res.LatestReadyRevision = dcl.FlattenString(m["latestReadyRevision"])
	res.LatestCreatedRevision = dcl.FlattenString(m["latestCreatedRevision"])
	res.TrafficStatuses = flattenServiceTrafficStatusesSlice(c, m["trafficStatuses"])
	res.Uri = dcl.FlattenString(m["uri"])
	res.Reconciling = dcl.FlattenBool(m["reconciling"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandServiceBinaryAuthorizationMap expands the contents of ServiceBinaryAuthorization into a JSON
// request object.
func expandServiceBinaryAuthorizationMap(c *Client, f map[string]ServiceBinaryAuthorization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceBinaryAuthorization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceBinaryAuthorizationSlice expands the contents of ServiceBinaryAuthorization into a JSON
// request object.
func expandServiceBinaryAuthorizationSlice(c *Client, f []ServiceBinaryAuthorization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceBinaryAuthorization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceBinaryAuthorizationMap flattens the contents of ServiceBinaryAuthorization from a JSON
// response object.
func flattenServiceBinaryAuthorizationMap(c *Client, i interface{}) map[string]ServiceBinaryAuthorization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceBinaryAuthorization{}
	}

	if len(a) == 0 {
		return map[string]ServiceBinaryAuthorization{}
	}

	items := make(map[string]ServiceBinaryAuthorization)
	for k, item := range a {
		items[k] = *flattenServiceBinaryAuthorization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceBinaryAuthorizationSlice flattens the contents of ServiceBinaryAuthorization from a JSON
// response object.
func flattenServiceBinaryAuthorizationSlice(c *Client, i interface{}) []ServiceBinaryAuthorization {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceBinaryAuthorization{}
	}

	if len(a) == 0 {
		return []ServiceBinaryAuthorization{}
	}

	items := make([]ServiceBinaryAuthorization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceBinaryAuthorization(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceBinaryAuthorization expands an instance of ServiceBinaryAuthorization into a JSON
// request object.
func expandServiceBinaryAuthorization(c *Client, f *ServiceBinaryAuthorization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Policy; !dcl.IsEmptyValueIndirect(v) {
		m["policy"] = v
	}
	if v := f.BreakglassJustification; !dcl.IsEmptyValueIndirect(v) {
		m["breakglassJustification"] = v
	}

	return m, nil
}

// flattenServiceBinaryAuthorization flattens an instance of ServiceBinaryAuthorization from a JSON
// response object.
func flattenServiceBinaryAuthorization(c *Client, i interface{}) *ServiceBinaryAuthorization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceBinaryAuthorization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceBinaryAuthorization
	}
	r.Policy = dcl.FlattenString(m["policy"])
	r.BreakglassJustification = dcl.FlattenString(m["breakglassJustification"])

	return r
}

// expandServiceTemplateMap expands the contents of ServiceTemplate into a JSON
// request object.
func expandServiceTemplateMap(c *Client, f map[string]ServiceTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplate(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateSlice expands the contents of ServiceTemplate into a JSON
// request object.
func expandServiceTemplateSlice(c *Client, f []ServiceTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplate(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateMap flattens the contents of ServiceTemplate from a JSON
// response object.
func flattenServiceTemplateMap(c *Client, i interface{}) map[string]ServiceTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplate{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplate{}
	}

	items := make(map[string]ServiceTemplate)
	for k, item := range a {
		items[k] = *flattenServiceTemplate(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateSlice flattens the contents of ServiceTemplate from a JSON
// response object.
func flattenServiceTemplateSlice(c *Client, i interface{}) []ServiceTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplate{}
	}

	if len(a) == 0 {
		return []ServiceTemplate{}
	}

	items := make([]ServiceTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplate(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplate expands an instance of ServiceTemplate into a JSON
// request object.
func expandServiceTemplate(c *Client, f *ServiceTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v, err := expandServiceTemplateScaling(c, f.Scaling); err != nil {
		return nil, fmt.Errorf("error expanding Scaling into scaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scaling"] = v
	}
	if v, err := expandServiceTemplateVPCAccess(c, f.VPCAccess); err != nil {
		return nil, fmt.Errorf("error expanding VPCAccess into vpcAccess: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["vpcAccess"] = v
	}
	if v := f.ContainerConcurrency; !dcl.IsEmptyValueIndirect(v) {
		m["containerConcurrency"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v, err := expandServiceTemplateContainersSlice(c, f.Containers); err != nil {
		return nil, fmt.Errorf("error expanding Containers into containers: %w", err)
	} else if v != nil {
		m["containers"] = v
	}
	if v, err := expandServiceTemplateVolumesSlice(c, f.Volumes); err != nil {
		return nil, fmt.Errorf("error expanding Volumes into volumes: %w", err)
	} else if v != nil {
		m["volumes"] = v
	}
	if v := f.Confidential; !dcl.IsEmptyValueIndirect(v) {
		m["confidential"] = v
	}
	if v := f.ExecutionEnvironment; !dcl.IsEmptyValueIndirect(v) {
		m["executionEnvironment"] = v
	}

	return m, nil
}

// flattenServiceTemplate flattens an instance of ServiceTemplate from a JSON
// response object.
func flattenServiceTemplate(c *Client, i interface{}) *ServiceTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplate
	}
	r.Revision = dcl.FlattenString(m["revision"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	r.Scaling = flattenServiceTemplateScaling(c, m["scaling"])
	r.VPCAccess = flattenServiceTemplateVPCAccess(c, m["vpcAccess"])
	r.ContainerConcurrency = dcl.FlattenInteger(m["containerConcurrency"])
	r.Timeout = dcl.FlattenString(m["timeout"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.Containers = flattenServiceTemplateContainersSlice(c, m["containers"])
	r.Volumes = flattenServiceTemplateVolumesSlice(c, m["volumes"])
	r.Confidential = dcl.FlattenBool(m["confidential"])
	r.ExecutionEnvironment = flattenServiceTemplateExecutionEnvironmentEnum(m["executionEnvironment"])

	return r
}

// expandServiceTemplateScalingMap expands the contents of ServiceTemplateScaling into a JSON
// request object.
func expandServiceTemplateScalingMap(c *Client, f map[string]ServiceTemplateScaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateScaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateScalingSlice expands the contents of ServiceTemplateScaling into a JSON
// request object.
func expandServiceTemplateScalingSlice(c *Client, f []ServiceTemplateScaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateScaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateScalingMap flattens the contents of ServiceTemplateScaling from a JSON
// response object.
func flattenServiceTemplateScalingMap(c *Client, i interface{}) map[string]ServiceTemplateScaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateScaling{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateScaling{}
	}

	items := make(map[string]ServiceTemplateScaling)
	for k, item := range a {
		items[k] = *flattenServiceTemplateScaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateScalingSlice flattens the contents of ServiceTemplateScaling from a JSON
// response object.
func flattenServiceTemplateScalingSlice(c *Client, i interface{}) []ServiceTemplateScaling {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateScaling{}
	}

	if len(a) == 0 {
		return []ServiceTemplateScaling{}
	}

	items := make([]ServiceTemplateScaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateScaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateScaling expands an instance of ServiceTemplateScaling into a JSON
// request object.
func expandServiceTemplateScaling(c *Client, f *ServiceTemplateScaling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinInstanceCount; !dcl.IsEmptyValueIndirect(v) {
		m["minInstanceCount"] = v
	}
	if v := f.MaxInstanceCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstanceCount"] = v
	}

	return m, nil
}

// flattenServiceTemplateScaling flattens an instance of ServiceTemplateScaling from a JSON
// response object.
func flattenServiceTemplateScaling(c *Client, i interface{}) *ServiceTemplateScaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateScaling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateScaling
	}
	r.MinInstanceCount = dcl.FlattenInteger(m["minInstanceCount"])
	r.MaxInstanceCount = dcl.FlattenInteger(m["maxInstanceCount"])

	return r
}

// expandServiceTemplateVPCAccessMap expands the contents of ServiceTemplateVPCAccess into a JSON
// request object.
func expandServiceTemplateVPCAccessMap(c *Client, f map[string]ServiceTemplateVPCAccess) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVPCAccess(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVPCAccessSlice expands the contents of ServiceTemplateVPCAccess into a JSON
// request object.
func expandServiceTemplateVPCAccessSlice(c *Client, f []ServiceTemplateVPCAccess) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVPCAccess(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVPCAccessMap flattens the contents of ServiceTemplateVPCAccess from a JSON
// response object.
func flattenServiceTemplateVPCAccessMap(c *Client, i interface{}) map[string]ServiceTemplateVPCAccess {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVPCAccess{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVPCAccess{}
	}

	items := make(map[string]ServiceTemplateVPCAccess)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVPCAccess(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateVPCAccessSlice flattens the contents of ServiceTemplateVPCAccess from a JSON
// response object.
func flattenServiceTemplateVPCAccessSlice(c *Client, i interface{}) []ServiceTemplateVPCAccess {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVPCAccess{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVPCAccess{}
	}

	items := make([]ServiceTemplateVPCAccess, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVPCAccess(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateVPCAccess expands an instance of ServiceTemplateVPCAccess into a JSON
// request object.
func expandServiceTemplateVPCAccess(c *Client, f *ServiceTemplateVPCAccess) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Connector; !dcl.IsEmptyValueIndirect(v) {
		m["connector"] = v
	}
	if v := f.Egress; !dcl.IsEmptyValueIndirect(v) {
		m["egress"] = v
	}

	return m, nil
}

// flattenServiceTemplateVPCAccess flattens an instance of ServiceTemplateVPCAccess from a JSON
// response object.
func flattenServiceTemplateVPCAccess(c *Client, i interface{}) *ServiceTemplateVPCAccess {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVPCAccess{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVPCAccess
	}
	r.Connector = dcl.FlattenString(m["connector"])
	r.Egress = flattenServiceTemplateVPCAccessEgressEnum(m["egress"])

	return r
}

// expandServiceTemplateContainersMap expands the contents of ServiceTemplateContainers into a JSON
// request object.
func expandServiceTemplateContainersMap(c *Client, f map[string]ServiceTemplateContainers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersSlice expands the contents of ServiceTemplateContainers into a JSON
// request object.
func expandServiceTemplateContainersSlice(c *Client, f []ServiceTemplateContainers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersMap flattens the contents of ServiceTemplateContainers from a JSON
// response object.
func flattenServiceTemplateContainersMap(c *Client, i interface{}) map[string]ServiceTemplateContainers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainers{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainers{}
	}

	items := make(map[string]ServiceTemplateContainers)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateContainersSlice flattens the contents of ServiceTemplateContainers from a JSON
// response object.
func flattenServiceTemplateContainersSlice(c *Client, i interface{}) []ServiceTemplateContainers {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainers{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainers{}
	}

	items := make([]ServiceTemplateContainers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainers(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateContainers expands an instance of ServiceTemplateContainers into a JSON
// request object.
func expandServiceTemplateContainers(c *Client, f *ServiceTemplateContainers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Image; !dcl.IsEmptyValueIndirect(v) {
		m["image"] = v
	}
	if v := f.Command; v != nil {
		m["command"] = v
	}
	if v := f.Args; v != nil {
		m["args"] = v
	}
	if v, err := expandServiceTemplateContainersEnvSlice(c, f.Env); err != nil {
		return nil, fmt.Errorf("error expanding Env into env: %w", err)
	} else if v != nil {
		m["env"] = v
	}
	if v, err := expandServiceTemplateContainersResources(c, f.Resources); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resources"] = v
	}
	if v, err := expandServiceTemplateContainersPortsSlice(c, f.Ports); err != nil {
		return nil, fmt.Errorf("error expanding Ports into ports: %w", err)
	} else if v != nil {
		m["ports"] = v
	}
	if v, err := expandServiceTemplateContainersVolumeMountsSlice(c, f.VolumeMounts); err != nil {
		return nil, fmt.Errorf("error expanding VolumeMounts into volumeMounts: %w", err)
	} else if v != nil {
		m["volumeMounts"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainers flattens an instance of ServiceTemplateContainers from a JSON
// response object.
func flattenServiceTemplateContainers(c *Client, i interface{}) *ServiceTemplateContainers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainers
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Image = dcl.FlattenString(m["image"])
	r.Command = dcl.FlattenStringSlice(m["command"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.Env = flattenServiceTemplateContainersEnvSlice(c, m["env"])
	r.Resources = flattenServiceTemplateContainersResources(c, m["resources"])
	r.Ports = flattenServiceTemplateContainersPortsSlice(c, m["ports"])
	r.VolumeMounts = flattenServiceTemplateContainersVolumeMountsSlice(c, m["volumeMounts"])

	return r
}

// expandServiceTemplateContainersEnvMap expands the contents of ServiceTemplateContainersEnv into a JSON
// request object.
func expandServiceTemplateContainersEnvMap(c *Client, f map[string]ServiceTemplateContainersEnv) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersEnv(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersEnvSlice expands the contents of ServiceTemplateContainersEnv into a JSON
// request object.
func expandServiceTemplateContainersEnvSlice(c *Client, f []ServiceTemplateContainersEnv) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersEnv(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersEnvMap flattens the contents of ServiceTemplateContainersEnv from a JSON
// response object.
func flattenServiceTemplateContainersEnvMap(c *Client, i interface{}) map[string]ServiceTemplateContainersEnv {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersEnv{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersEnv{}
	}

	items := make(map[string]ServiceTemplateContainersEnv)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersEnv(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateContainersEnvSlice flattens the contents of ServiceTemplateContainersEnv from a JSON
// response object.
func flattenServiceTemplateContainersEnvSlice(c *Client, i interface{}) []ServiceTemplateContainersEnv {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersEnv{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersEnv{}
	}

	items := make([]ServiceTemplateContainersEnv, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersEnv(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateContainersEnv expands an instance of ServiceTemplateContainersEnv into a JSON
// request object.
func expandServiceTemplateContainersEnv(c *Client, f *ServiceTemplateContainersEnv) (map[string]interface{}, error) {
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
	if v, err := expandServiceTemplateContainersEnvValueSource(c, f.ValueSource); err != nil {
		return nil, fmt.Errorf("error expanding ValueSource into valueSource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["valueSource"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersEnv flattens an instance of ServiceTemplateContainersEnv from a JSON
// response object.
func flattenServiceTemplateContainersEnv(c *Client, i interface{}) *ServiceTemplateContainersEnv {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersEnv{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersEnv
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])
	r.ValueSource = flattenServiceTemplateContainersEnvValueSource(c, m["valueSource"])

	return r
}

// expandServiceTemplateContainersEnvValueSourceMap expands the contents of ServiceTemplateContainersEnvValueSource into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceMap(c *Client, f map[string]ServiceTemplateContainersEnvValueSource) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersEnvValueSource(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersEnvValueSourceSlice expands the contents of ServiceTemplateContainersEnvValueSource into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceSlice(c *Client, f []ServiceTemplateContainersEnvValueSource) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersEnvValueSource(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersEnvValueSourceMap flattens the contents of ServiceTemplateContainersEnvValueSource from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceMap(c *Client, i interface{}) map[string]ServiceTemplateContainersEnvValueSource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersEnvValueSource{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersEnvValueSource{}
	}

	items := make(map[string]ServiceTemplateContainersEnvValueSource)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersEnvValueSource(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateContainersEnvValueSourceSlice flattens the contents of ServiceTemplateContainersEnvValueSource from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceSlice(c *Client, i interface{}) []ServiceTemplateContainersEnvValueSource {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersEnvValueSource{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersEnvValueSource{}
	}

	items := make([]ServiceTemplateContainersEnvValueSource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersEnvValueSource(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateContainersEnvValueSource expands an instance of ServiceTemplateContainersEnvValueSource into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSource(c *Client, f *ServiceTemplateContainersEnvValueSource) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandServiceTemplateContainersEnvValueSourceSecretKeyRef(c, f.SecretKeyRef); err != nil {
		return nil, fmt.Errorf("error expanding SecretKeyRef into secretKeyRef: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secretKeyRef"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersEnvValueSource flattens an instance of ServiceTemplateContainersEnvValueSource from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSource(c *Client, i interface{}) *ServiceTemplateContainersEnvValueSource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersEnvValueSource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersEnvValueSource
	}
	r.SecretKeyRef = flattenServiceTemplateContainersEnvValueSourceSecretKeyRef(c, m["secretKeyRef"])

	return r
}

// expandServiceTemplateContainersEnvValueSourceSecretKeyRefMap expands the contents of ServiceTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceSecretKeyRefMap(c *Client, f map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersEnvValueSourceSecretKeyRef(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersEnvValueSourceSecretKeyRefSlice expands the contents of ServiceTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, f []ServiceTemplateContainersEnvValueSourceSecretKeyRef) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersEnvValueSourceSecretKeyRef(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersEnvValueSourceSecretKeyRefMap flattens the contents of ServiceTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceSecretKeyRefMap(c *Client, i interface{}) map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	items := make(map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersEnvValueSourceSecretKeyRef(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateContainersEnvValueSourceSecretKeyRefSlice flattens the contents of ServiceTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, i interface{}) []ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	items := make([]ServiceTemplateContainersEnvValueSourceSecretKeyRef, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersEnvValueSourceSecretKeyRef(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateContainersEnvValueSourceSecretKeyRef expands an instance of ServiceTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceSecretKeyRef(c *Client, f *ServiceTemplateContainersEnvValueSourceSecretKeyRef) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Secret; !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersEnvValueSourceSecretKeyRef flattens an instance of ServiceTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceSecretKeyRef(c *Client, i interface{}) *ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersEnvValueSourceSecretKeyRef{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersEnvValueSourceSecretKeyRef
	}
	r.Secret = dcl.FlattenString(m["secret"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandServiceTemplateContainersResourcesMap expands the contents of ServiceTemplateContainersResources into a JSON
// request object.
func expandServiceTemplateContainersResourcesMap(c *Client, f map[string]ServiceTemplateContainersResources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersResources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersResourcesSlice expands the contents of ServiceTemplateContainersResources into a JSON
// request object.
func expandServiceTemplateContainersResourcesSlice(c *Client, f []ServiceTemplateContainersResources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersResources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersResourcesMap flattens the contents of ServiceTemplateContainersResources from a JSON
// response object.
func flattenServiceTemplateContainersResourcesMap(c *Client, i interface{}) map[string]ServiceTemplateContainersResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersResources{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersResources{}
	}

	items := make(map[string]ServiceTemplateContainersResources)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersResources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateContainersResourcesSlice flattens the contents of ServiceTemplateContainersResources from a JSON
// response object.
func flattenServiceTemplateContainersResourcesSlice(c *Client, i interface{}) []ServiceTemplateContainersResources {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersResources{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersResources{}
	}

	items := make([]ServiceTemplateContainersResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersResources(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateContainersResources expands an instance of ServiceTemplateContainersResources into a JSON
// request object.
func expandServiceTemplateContainersResources(c *Client, f *ServiceTemplateContainersResources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Limits; !dcl.IsEmptyValueIndirect(v) {
		m["limits"] = v
	}
	if v := f.CpuIdle; !dcl.IsEmptyValueIndirect(v) {
		m["cpuIdle"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersResources flattens an instance of ServiceTemplateContainersResources from a JSON
// response object.
func flattenServiceTemplateContainersResources(c *Client, i interface{}) *ServiceTemplateContainersResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersResources
	}
	r.Limits = dcl.FlattenKeyValuePairs(m["limits"])
	r.CpuIdle = dcl.FlattenBool(m["cpuIdle"])

	return r
}

// expandServiceTemplateContainersPortsMap expands the contents of ServiceTemplateContainersPorts into a JSON
// request object.
func expandServiceTemplateContainersPortsMap(c *Client, f map[string]ServiceTemplateContainersPorts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersPorts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersPortsSlice expands the contents of ServiceTemplateContainersPorts into a JSON
// request object.
func expandServiceTemplateContainersPortsSlice(c *Client, f []ServiceTemplateContainersPorts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersPorts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersPortsMap flattens the contents of ServiceTemplateContainersPorts from a JSON
// response object.
func flattenServiceTemplateContainersPortsMap(c *Client, i interface{}) map[string]ServiceTemplateContainersPorts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersPorts{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersPorts{}
	}

	items := make(map[string]ServiceTemplateContainersPorts)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersPorts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateContainersPortsSlice flattens the contents of ServiceTemplateContainersPorts from a JSON
// response object.
func flattenServiceTemplateContainersPortsSlice(c *Client, i interface{}) []ServiceTemplateContainersPorts {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersPorts{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersPorts{}
	}

	items := make([]ServiceTemplateContainersPorts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersPorts(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateContainersPorts expands an instance of ServiceTemplateContainersPorts into a JSON
// request object.
func expandServiceTemplateContainersPorts(c *Client, f *ServiceTemplateContainersPorts) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ContainerPort; !dcl.IsEmptyValueIndirect(v) {
		m["containerPort"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersPorts flattens an instance of ServiceTemplateContainersPorts from a JSON
// response object.
func flattenServiceTemplateContainersPorts(c *Client, i interface{}) *ServiceTemplateContainersPorts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersPorts{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersPorts
	}
	r.Name = dcl.FlattenString(m["name"])
	r.ContainerPort = dcl.FlattenInteger(m["containerPort"])

	return r
}

// expandServiceTemplateContainersVolumeMountsMap expands the contents of ServiceTemplateContainersVolumeMounts into a JSON
// request object.
func expandServiceTemplateContainersVolumeMountsMap(c *Client, f map[string]ServiceTemplateContainersVolumeMounts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersVolumeMounts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersVolumeMountsSlice expands the contents of ServiceTemplateContainersVolumeMounts into a JSON
// request object.
func expandServiceTemplateContainersVolumeMountsSlice(c *Client, f []ServiceTemplateContainersVolumeMounts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersVolumeMounts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersVolumeMountsMap flattens the contents of ServiceTemplateContainersVolumeMounts from a JSON
// response object.
func flattenServiceTemplateContainersVolumeMountsMap(c *Client, i interface{}) map[string]ServiceTemplateContainersVolumeMounts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersVolumeMounts{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersVolumeMounts{}
	}

	items := make(map[string]ServiceTemplateContainersVolumeMounts)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersVolumeMounts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateContainersVolumeMountsSlice flattens the contents of ServiceTemplateContainersVolumeMounts from a JSON
// response object.
func flattenServiceTemplateContainersVolumeMountsSlice(c *Client, i interface{}) []ServiceTemplateContainersVolumeMounts {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersVolumeMounts{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersVolumeMounts{}
	}

	items := make([]ServiceTemplateContainersVolumeMounts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersVolumeMounts(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateContainersVolumeMounts expands an instance of ServiceTemplateContainersVolumeMounts into a JSON
// request object.
func expandServiceTemplateContainersVolumeMounts(c *Client, f *ServiceTemplateContainersVolumeMounts) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.MountPath; !dcl.IsEmptyValueIndirect(v) {
		m["mountPath"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersVolumeMounts flattens an instance of ServiceTemplateContainersVolumeMounts from a JSON
// response object.
func flattenServiceTemplateContainersVolumeMounts(c *Client, i interface{}) *ServiceTemplateContainersVolumeMounts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersVolumeMounts{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersVolumeMounts
	}
	r.Name = dcl.FlattenString(m["name"])
	r.MountPath = dcl.FlattenString(m["mountPath"])

	return r
}

// expandServiceTemplateVolumesMap expands the contents of ServiceTemplateVolumes into a JSON
// request object.
func expandServiceTemplateVolumesMap(c *Client, f map[string]ServiceTemplateVolumes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVolumes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVolumesSlice expands the contents of ServiceTemplateVolumes into a JSON
// request object.
func expandServiceTemplateVolumesSlice(c *Client, f []ServiceTemplateVolumes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVolumes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVolumesMap flattens the contents of ServiceTemplateVolumes from a JSON
// response object.
func flattenServiceTemplateVolumesMap(c *Client, i interface{}) map[string]ServiceTemplateVolumes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVolumes{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVolumes{}
	}

	items := make(map[string]ServiceTemplateVolumes)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVolumes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateVolumesSlice flattens the contents of ServiceTemplateVolumes from a JSON
// response object.
func flattenServiceTemplateVolumesSlice(c *Client, i interface{}) []ServiceTemplateVolumes {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVolumes{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVolumes{}
	}

	items := make([]ServiceTemplateVolumes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVolumes(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateVolumes expands an instance of ServiceTemplateVolumes into a JSON
// request object.
func expandServiceTemplateVolumes(c *Client, f *ServiceTemplateVolumes) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandServiceTemplateVolumesSecret(c, f.Secret); err != nil {
		return nil, fmt.Errorf("error expanding Secret into secret: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v, err := expandServiceTemplateVolumesCloudSqlInstance(c, f.CloudSqlInstance); err != nil {
		return nil, fmt.Errorf("error expanding CloudSqlInstance into cloudSqlInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudSqlInstance"] = v
	}

	return m, nil
}

// flattenServiceTemplateVolumes flattens an instance of ServiceTemplateVolumes from a JSON
// response object.
func flattenServiceTemplateVolumes(c *Client, i interface{}) *ServiceTemplateVolumes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVolumes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVolumes
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Secret = flattenServiceTemplateVolumesSecret(c, m["secret"])
	r.CloudSqlInstance = flattenServiceTemplateVolumesCloudSqlInstance(c, m["cloudSqlInstance"])

	return r
}

// expandServiceTemplateVolumesSecretMap expands the contents of ServiceTemplateVolumesSecret into a JSON
// request object.
func expandServiceTemplateVolumesSecretMap(c *Client, f map[string]ServiceTemplateVolumesSecret) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVolumesSecret(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVolumesSecretSlice expands the contents of ServiceTemplateVolumesSecret into a JSON
// request object.
func expandServiceTemplateVolumesSecretSlice(c *Client, f []ServiceTemplateVolumesSecret) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVolumesSecret(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVolumesSecretMap flattens the contents of ServiceTemplateVolumesSecret from a JSON
// response object.
func flattenServiceTemplateVolumesSecretMap(c *Client, i interface{}) map[string]ServiceTemplateVolumesSecret {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVolumesSecret{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVolumesSecret{}
	}

	items := make(map[string]ServiceTemplateVolumesSecret)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVolumesSecret(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateVolumesSecretSlice flattens the contents of ServiceTemplateVolumesSecret from a JSON
// response object.
func flattenServiceTemplateVolumesSecretSlice(c *Client, i interface{}) []ServiceTemplateVolumesSecret {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVolumesSecret{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVolumesSecret{}
	}

	items := make([]ServiceTemplateVolumesSecret, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVolumesSecret(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateVolumesSecret expands an instance of ServiceTemplateVolumesSecret into a JSON
// request object.
func expandServiceTemplateVolumesSecret(c *Client, f *ServiceTemplateVolumesSecret) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Secret; !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v, err := expandServiceTemplateVolumesSecretItemsSlice(c, f.Items); err != nil {
		return nil, fmt.Errorf("error expanding Items into items: %w", err)
	} else if v != nil {
		m["items"] = v
	}
	if v := f.DefaultMode; !dcl.IsEmptyValueIndirect(v) {
		m["defaultMode"] = v
	}

	return m, nil
}

// flattenServiceTemplateVolumesSecret flattens an instance of ServiceTemplateVolumesSecret from a JSON
// response object.
func flattenServiceTemplateVolumesSecret(c *Client, i interface{}) *ServiceTemplateVolumesSecret {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVolumesSecret{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVolumesSecret
	}
	r.Secret = dcl.FlattenString(m["secret"])
	r.Items = flattenServiceTemplateVolumesSecretItemsSlice(c, m["items"])
	r.DefaultMode = dcl.FlattenInteger(m["defaultMode"])

	return r
}

// expandServiceTemplateVolumesSecretItemsMap expands the contents of ServiceTemplateVolumesSecretItems into a JSON
// request object.
func expandServiceTemplateVolumesSecretItemsMap(c *Client, f map[string]ServiceTemplateVolumesSecretItems) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVolumesSecretItems(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVolumesSecretItemsSlice expands the contents of ServiceTemplateVolumesSecretItems into a JSON
// request object.
func expandServiceTemplateVolumesSecretItemsSlice(c *Client, f []ServiceTemplateVolumesSecretItems) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVolumesSecretItems(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVolumesSecretItemsMap flattens the contents of ServiceTemplateVolumesSecretItems from a JSON
// response object.
func flattenServiceTemplateVolumesSecretItemsMap(c *Client, i interface{}) map[string]ServiceTemplateVolumesSecretItems {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVolumesSecretItems{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVolumesSecretItems{}
	}

	items := make(map[string]ServiceTemplateVolumesSecretItems)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVolumesSecretItems(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateVolumesSecretItemsSlice flattens the contents of ServiceTemplateVolumesSecretItems from a JSON
// response object.
func flattenServiceTemplateVolumesSecretItemsSlice(c *Client, i interface{}) []ServiceTemplateVolumesSecretItems {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVolumesSecretItems{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVolumesSecretItems{}
	}

	items := make([]ServiceTemplateVolumesSecretItems, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVolumesSecretItems(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateVolumesSecretItems expands an instance of ServiceTemplateVolumesSecretItems into a JSON
// request object.
func expandServiceTemplateVolumesSecretItems(c *Client, f *ServiceTemplateVolumesSecretItems) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenServiceTemplateVolumesSecretItems flattens an instance of ServiceTemplateVolumesSecretItems from a JSON
// response object.
func flattenServiceTemplateVolumesSecretItems(c *Client, i interface{}) *ServiceTemplateVolumesSecretItems {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVolumesSecretItems{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVolumesSecretItems
	}
	r.Path = dcl.FlattenString(m["path"])
	r.Version = dcl.FlattenString(m["version"])
	r.Mode = dcl.FlattenInteger(m["mode"])

	return r
}

// expandServiceTemplateVolumesCloudSqlInstanceMap expands the contents of ServiceTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandServiceTemplateVolumesCloudSqlInstanceMap(c *Client, f map[string]ServiceTemplateVolumesCloudSqlInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVolumesCloudSqlInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVolumesCloudSqlInstanceSlice expands the contents of ServiceTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandServiceTemplateVolumesCloudSqlInstanceSlice(c *Client, f []ServiceTemplateVolumesCloudSqlInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVolumesCloudSqlInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVolumesCloudSqlInstanceMap flattens the contents of ServiceTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenServiceTemplateVolumesCloudSqlInstanceMap(c *Client, i interface{}) map[string]ServiceTemplateVolumesCloudSqlInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVolumesCloudSqlInstance{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVolumesCloudSqlInstance{}
	}

	items := make(map[string]ServiceTemplateVolumesCloudSqlInstance)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVolumesCloudSqlInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTemplateVolumesCloudSqlInstanceSlice flattens the contents of ServiceTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenServiceTemplateVolumesCloudSqlInstanceSlice(c *Client, i interface{}) []ServiceTemplateVolumesCloudSqlInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVolumesCloudSqlInstance{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVolumesCloudSqlInstance{}
	}

	items := make([]ServiceTemplateVolumesCloudSqlInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVolumesCloudSqlInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTemplateVolumesCloudSqlInstance expands an instance of ServiceTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandServiceTemplateVolumesCloudSqlInstance(c *Client, f *ServiceTemplateVolumesCloudSqlInstance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Connections; v != nil {
		m["connections"] = v
	}

	return m, nil
}

// flattenServiceTemplateVolumesCloudSqlInstance flattens an instance of ServiceTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenServiceTemplateVolumesCloudSqlInstance(c *Client, i interface{}) *ServiceTemplateVolumesCloudSqlInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVolumesCloudSqlInstance{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVolumesCloudSqlInstance
	}
	r.Connections = dcl.FlattenStringSlice(m["connections"])

	return r
}

// expandServiceTrafficMap expands the contents of ServiceTraffic into a JSON
// request object.
func expandServiceTrafficMap(c *Client, f map[string]ServiceTraffic) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTraffic(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTrafficSlice expands the contents of ServiceTraffic into a JSON
// request object.
func expandServiceTrafficSlice(c *Client, f []ServiceTraffic) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTraffic(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTrafficMap flattens the contents of ServiceTraffic from a JSON
// response object.
func flattenServiceTrafficMap(c *Client, i interface{}) map[string]ServiceTraffic {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTraffic{}
	}

	if len(a) == 0 {
		return map[string]ServiceTraffic{}
	}

	items := make(map[string]ServiceTraffic)
	for k, item := range a {
		items[k] = *flattenServiceTraffic(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTrafficSlice flattens the contents of ServiceTraffic from a JSON
// response object.
func flattenServiceTrafficSlice(c *Client, i interface{}) []ServiceTraffic {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTraffic{}
	}

	if len(a) == 0 {
		return []ServiceTraffic{}
	}

	items := make([]ServiceTraffic, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTraffic(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTraffic expands an instance of ServiceTraffic into a JSON
// request object.
func expandServiceTraffic(c *Client, f *ServiceTraffic) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tag"] = v
	}

	return m, nil
}

// flattenServiceTraffic flattens an instance of ServiceTraffic from a JSON
// response object.
func flattenServiceTraffic(c *Client, i interface{}) *ServiceTraffic {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTraffic{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTraffic
	}
	r.Type = flattenServiceTrafficTypeEnum(m["type"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Tag = dcl.FlattenString(m["tag"])

	return r
}

// expandServiceGooglecloudrunopv2ConditionMap expands the contents of ServiceGooglecloudrunopv2Condition into a JSON
// request object.
func expandServiceGooglecloudrunopv2ConditionMap(c *Client, f map[string]ServiceGooglecloudrunopv2Condition) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceGooglecloudrunopv2Condition(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceGooglecloudrunopv2ConditionSlice expands the contents of ServiceGooglecloudrunopv2Condition into a JSON
// request object.
func expandServiceGooglecloudrunopv2ConditionSlice(c *Client, f []ServiceGooglecloudrunopv2Condition) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceGooglecloudrunopv2Condition(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceGooglecloudrunopv2ConditionMap flattens the contents of ServiceGooglecloudrunopv2Condition from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionMap(c *Client, i interface{}) map[string]ServiceGooglecloudrunopv2Condition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGooglecloudrunopv2Condition{}
	}

	if len(a) == 0 {
		return map[string]ServiceGooglecloudrunopv2Condition{}
	}

	items := make(map[string]ServiceGooglecloudrunopv2Condition)
	for k, item := range a {
		items[k] = *flattenServiceGooglecloudrunopv2Condition(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionSlice flattens the contents of ServiceGooglecloudrunopv2Condition from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionSlice(c *Client, i interface{}) []ServiceGooglecloudrunopv2Condition {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGooglecloudrunopv2Condition{}
	}

	if len(a) == 0 {
		return []ServiceGooglecloudrunopv2Condition{}
	}

	items := make([]ServiceGooglecloudrunopv2Condition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGooglecloudrunopv2Condition(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceGooglecloudrunopv2Condition expands an instance of ServiceGooglecloudrunopv2Condition into a JSON
// request object.
func expandServiceGooglecloudrunopv2Condition(c *Client, f *ServiceGooglecloudrunopv2Condition) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v := f.LastTransitionTime; !dcl.IsEmptyValueIndirect(v) {
		m["lastTransitionTime"] = v
	}
	if v := f.Severity; !dcl.IsEmptyValueIndirect(v) {
		m["severity"] = v
	}
	if v := f.Reason; !dcl.IsEmptyValueIndirect(v) {
		m["reason"] = v
	}
	if v := f.InternalReason; !dcl.IsEmptyValueIndirect(v) {
		m["internalReason"] = v
	}
	if v := f.DomainMappingReason; !dcl.IsEmptyValueIndirect(v) {
		m["domainMappingReason"] = v
	}
	if v := f.RevisionReason; !dcl.IsEmptyValueIndirect(v) {
		m["revisionReason"] = v
	}
	if v := f.JobReason; !dcl.IsEmptyValueIndirect(v) {
		m["jobReason"] = v
	}

	return m, nil
}

// flattenServiceGooglecloudrunopv2Condition flattens an instance of ServiceGooglecloudrunopv2Condition from a JSON
// response object.
func flattenServiceGooglecloudrunopv2Condition(c *Client, i interface{}) *ServiceGooglecloudrunopv2Condition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceGooglecloudrunopv2Condition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceGooglecloudrunopv2Condition
	}
	r.Type = dcl.FlattenString(m["type"])
	r.State = flattenServiceGooglecloudrunopv2ConditionStateEnum(m["state"])
	r.Message = dcl.FlattenString(m["message"])
	r.LastTransitionTime = dcl.FlattenString(m["lastTransitionTime"])
	r.Severity = flattenServiceGooglecloudrunopv2ConditionSeverityEnum(m["severity"])
	r.Reason = flattenServiceGooglecloudrunopv2ConditionReasonEnum(m["reason"])
	r.InternalReason = flattenServiceGooglecloudrunopv2ConditionInternalReasonEnum(m["internalReason"])
	r.DomainMappingReason = flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum(m["domainMappingReason"])
	r.RevisionReason = flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnum(m["revisionReason"])
	r.JobReason = flattenServiceGooglecloudrunopv2ConditionJobReasonEnum(m["jobReason"])

	return r
}

// expandServiceTrafficStatusesMap expands the contents of ServiceTrafficStatuses into a JSON
// request object.
func expandServiceTrafficStatusesMap(c *Client, f map[string]ServiceTrafficStatuses) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTrafficStatuses(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTrafficStatusesSlice expands the contents of ServiceTrafficStatuses into a JSON
// request object.
func expandServiceTrafficStatusesSlice(c *Client, f []ServiceTrafficStatuses) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTrafficStatuses(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTrafficStatusesMap flattens the contents of ServiceTrafficStatuses from a JSON
// response object.
func flattenServiceTrafficStatusesMap(c *Client, i interface{}) map[string]ServiceTrafficStatuses {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTrafficStatuses{}
	}

	if len(a) == 0 {
		return map[string]ServiceTrafficStatuses{}
	}

	items := make(map[string]ServiceTrafficStatuses)
	for k, item := range a {
		items[k] = *flattenServiceTrafficStatuses(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceTrafficStatusesSlice flattens the contents of ServiceTrafficStatuses from a JSON
// response object.
func flattenServiceTrafficStatusesSlice(c *Client, i interface{}) []ServiceTrafficStatuses {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTrafficStatuses{}
	}

	if len(a) == 0 {
		return []ServiceTrafficStatuses{}
	}

	items := make([]ServiceTrafficStatuses, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTrafficStatuses(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceTrafficStatuses expands an instance of ServiceTrafficStatuses into a JSON
// request object.
func expandServiceTrafficStatuses(c *Client, f *ServiceTrafficStatuses) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tag"] = v
	}
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}

	return m, nil
}

// flattenServiceTrafficStatuses flattens an instance of ServiceTrafficStatuses from a JSON
// response object.
func flattenServiceTrafficStatuses(c *Client, i interface{}) *ServiceTrafficStatuses {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTrafficStatuses{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTrafficStatuses
	}
	r.Type = flattenServiceTrafficStatusesTypeEnum(m["type"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Tag = dcl.FlattenString(m["tag"])
	r.Uri = dcl.FlattenString(m["uri"])

	return r
}

// flattenServiceIngressEnumMap flattens the contents of ServiceIngressEnum from a JSON
// response object.
func flattenServiceIngressEnumMap(c *Client, i interface{}) map[string]ServiceIngressEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceIngressEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceIngressEnum{}
	}

	items := make(map[string]ServiceIngressEnum)
	for k, item := range a {
		items[k] = *flattenServiceIngressEnum(item.(interface{}))
	}

	return items
}

// flattenServiceIngressEnumSlice flattens the contents of ServiceIngressEnum from a JSON
// response object.
func flattenServiceIngressEnumSlice(c *Client, i interface{}) []ServiceIngressEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceIngressEnum{}
	}

	if len(a) == 0 {
		return []ServiceIngressEnum{}
	}

	items := make([]ServiceIngressEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceIngressEnum(item.(interface{})))
	}

	return items
}

// flattenServiceIngressEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceIngressEnum with the same value as that string.
func flattenServiceIngressEnum(i interface{}) *ServiceIngressEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceIngressEnumRef("")
	}

	return ServiceIngressEnumRef(s)
}

// flattenServiceLaunchStageEnumMap flattens the contents of ServiceLaunchStageEnum from a JSON
// response object.
func flattenServiceLaunchStageEnumMap(c *Client, i interface{}) map[string]ServiceLaunchStageEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceLaunchStageEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceLaunchStageEnum{}
	}

	items := make(map[string]ServiceLaunchStageEnum)
	for k, item := range a {
		items[k] = *flattenServiceLaunchStageEnum(item.(interface{}))
	}

	return items
}

// flattenServiceLaunchStageEnumSlice flattens the contents of ServiceLaunchStageEnum from a JSON
// response object.
func flattenServiceLaunchStageEnumSlice(c *Client, i interface{}) []ServiceLaunchStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceLaunchStageEnum{}
	}

	if len(a) == 0 {
		return []ServiceLaunchStageEnum{}
	}

	items := make([]ServiceLaunchStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceLaunchStageEnum(item.(interface{})))
	}

	return items
}

// flattenServiceLaunchStageEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceLaunchStageEnum with the same value as that string.
func flattenServiceLaunchStageEnum(i interface{}) *ServiceLaunchStageEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceLaunchStageEnumRef("")
	}

	return ServiceLaunchStageEnumRef(s)
}

// flattenServiceTemplateVPCAccessEgressEnumMap flattens the contents of ServiceTemplateVPCAccessEgressEnum from a JSON
// response object.
func flattenServiceTemplateVPCAccessEgressEnumMap(c *Client, i interface{}) map[string]ServiceTemplateVPCAccessEgressEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVPCAccessEgressEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVPCAccessEgressEnum{}
	}

	items := make(map[string]ServiceTemplateVPCAccessEgressEnum)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVPCAccessEgressEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTemplateVPCAccessEgressEnumSlice flattens the contents of ServiceTemplateVPCAccessEgressEnum from a JSON
// response object.
func flattenServiceTemplateVPCAccessEgressEnumSlice(c *Client, i interface{}) []ServiceTemplateVPCAccessEgressEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVPCAccessEgressEnum{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVPCAccessEgressEnum{}
	}

	items := make([]ServiceTemplateVPCAccessEgressEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVPCAccessEgressEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTemplateVPCAccessEgressEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTemplateVPCAccessEgressEnum with the same value as that string.
func flattenServiceTemplateVPCAccessEgressEnum(i interface{}) *ServiceTemplateVPCAccessEgressEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceTemplateVPCAccessEgressEnumRef("")
	}

	return ServiceTemplateVPCAccessEgressEnumRef(s)
}

// flattenServiceTemplateExecutionEnvironmentEnumMap flattens the contents of ServiceTemplateExecutionEnvironmentEnum from a JSON
// response object.
func flattenServiceTemplateExecutionEnvironmentEnumMap(c *Client, i interface{}) map[string]ServiceTemplateExecutionEnvironmentEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateExecutionEnvironmentEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateExecutionEnvironmentEnum{}
	}

	items := make(map[string]ServiceTemplateExecutionEnvironmentEnum)
	for k, item := range a {
		items[k] = *flattenServiceTemplateExecutionEnvironmentEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTemplateExecutionEnvironmentEnumSlice flattens the contents of ServiceTemplateExecutionEnvironmentEnum from a JSON
// response object.
func flattenServiceTemplateExecutionEnvironmentEnumSlice(c *Client, i interface{}) []ServiceTemplateExecutionEnvironmentEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateExecutionEnvironmentEnum{}
	}

	if len(a) == 0 {
		return []ServiceTemplateExecutionEnvironmentEnum{}
	}

	items := make([]ServiceTemplateExecutionEnvironmentEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateExecutionEnvironmentEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTemplateExecutionEnvironmentEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTemplateExecutionEnvironmentEnum with the same value as that string.
func flattenServiceTemplateExecutionEnvironmentEnum(i interface{}) *ServiceTemplateExecutionEnvironmentEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceTemplateExecutionEnvironmentEnumRef("")
	}

	return ServiceTemplateExecutionEnvironmentEnumRef(s)
}

// flattenServiceTrafficTypeEnumMap flattens the contents of ServiceTrafficTypeEnum from a JSON
// response object.
func flattenServiceTrafficTypeEnumMap(c *Client, i interface{}) map[string]ServiceTrafficTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTrafficTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTrafficTypeEnum{}
	}

	items := make(map[string]ServiceTrafficTypeEnum)
	for k, item := range a {
		items[k] = *flattenServiceTrafficTypeEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTrafficTypeEnumSlice flattens the contents of ServiceTrafficTypeEnum from a JSON
// response object.
func flattenServiceTrafficTypeEnumSlice(c *Client, i interface{}) []ServiceTrafficTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTrafficTypeEnum{}
	}

	if len(a) == 0 {
		return []ServiceTrafficTypeEnum{}
	}

	items := make([]ServiceTrafficTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTrafficTypeEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTrafficTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTrafficTypeEnum with the same value as that string.
func flattenServiceTrafficTypeEnum(i interface{}) *ServiceTrafficTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceTrafficTypeEnumRef("")
	}

	return ServiceTrafficTypeEnumRef(s)
}

// flattenServiceGooglecloudrunopv2ConditionStateEnumMap flattens the contents of ServiceGooglecloudrunopv2ConditionStateEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionStateEnumMap(c *Client, i interface{}) map[string]ServiceGooglecloudrunopv2ConditionStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGooglecloudrunopv2ConditionStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceGooglecloudrunopv2ConditionStateEnum{}
	}

	items := make(map[string]ServiceGooglecloudrunopv2ConditionStateEnum)
	for k, item := range a {
		items[k] = *flattenServiceGooglecloudrunopv2ConditionStateEnum(item.(interface{}))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionStateEnumSlice flattens the contents of ServiceGooglecloudrunopv2ConditionStateEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionStateEnumSlice(c *Client, i interface{}) []ServiceGooglecloudrunopv2ConditionStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGooglecloudrunopv2ConditionStateEnum{}
	}

	if len(a) == 0 {
		return []ServiceGooglecloudrunopv2ConditionStateEnum{}
	}

	items := make([]ServiceGooglecloudrunopv2ConditionStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGooglecloudrunopv2ConditionStateEnum(item.(interface{})))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionStateEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceGooglecloudrunopv2ConditionStateEnum with the same value as that string.
func flattenServiceGooglecloudrunopv2ConditionStateEnum(i interface{}) *ServiceGooglecloudrunopv2ConditionStateEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceGooglecloudrunopv2ConditionStateEnumRef("")
	}

	return ServiceGooglecloudrunopv2ConditionStateEnumRef(s)
}

// flattenServiceGooglecloudrunopv2ConditionSeverityEnumMap flattens the contents of ServiceGooglecloudrunopv2ConditionSeverityEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionSeverityEnumMap(c *Client, i interface{}) map[string]ServiceGooglecloudrunopv2ConditionSeverityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGooglecloudrunopv2ConditionSeverityEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceGooglecloudrunopv2ConditionSeverityEnum{}
	}

	items := make(map[string]ServiceGooglecloudrunopv2ConditionSeverityEnum)
	for k, item := range a {
		items[k] = *flattenServiceGooglecloudrunopv2ConditionSeverityEnum(item.(interface{}))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionSeverityEnumSlice flattens the contents of ServiceGooglecloudrunopv2ConditionSeverityEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionSeverityEnumSlice(c *Client, i interface{}) []ServiceGooglecloudrunopv2ConditionSeverityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGooglecloudrunopv2ConditionSeverityEnum{}
	}

	if len(a) == 0 {
		return []ServiceGooglecloudrunopv2ConditionSeverityEnum{}
	}

	items := make([]ServiceGooglecloudrunopv2ConditionSeverityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGooglecloudrunopv2ConditionSeverityEnum(item.(interface{})))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionSeverityEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceGooglecloudrunopv2ConditionSeverityEnum with the same value as that string.
func flattenServiceGooglecloudrunopv2ConditionSeverityEnum(i interface{}) *ServiceGooglecloudrunopv2ConditionSeverityEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceGooglecloudrunopv2ConditionSeverityEnumRef("")
	}

	return ServiceGooglecloudrunopv2ConditionSeverityEnumRef(s)
}

// flattenServiceGooglecloudrunopv2ConditionReasonEnumMap flattens the contents of ServiceGooglecloudrunopv2ConditionReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionReasonEnumMap(c *Client, i interface{}) map[string]ServiceGooglecloudrunopv2ConditionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGooglecloudrunopv2ConditionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceGooglecloudrunopv2ConditionReasonEnum{}
	}

	items := make(map[string]ServiceGooglecloudrunopv2ConditionReasonEnum)
	for k, item := range a {
		items[k] = *flattenServiceGooglecloudrunopv2ConditionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionReasonEnumSlice flattens the contents of ServiceGooglecloudrunopv2ConditionReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionReasonEnumSlice(c *Client, i interface{}) []ServiceGooglecloudrunopv2ConditionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGooglecloudrunopv2ConditionReasonEnum{}
	}

	if len(a) == 0 {
		return []ServiceGooglecloudrunopv2ConditionReasonEnum{}
	}

	items := make([]ServiceGooglecloudrunopv2ConditionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGooglecloudrunopv2ConditionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceGooglecloudrunopv2ConditionReasonEnum with the same value as that string.
func flattenServiceGooglecloudrunopv2ConditionReasonEnum(i interface{}) *ServiceGooglecloudrunopv2ConditionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceGooglecloudrunopv2ConditionReasonEnumRef("")
	}

	return ServiceGooglecloudrunopv2ConditionReasonEnumRef(s)
}

// flattenServiceGooglecloudrunopv2ConditionInternalReasonEnumMap flattens the contents of ServiceGooglecloudrunopv2ConditionInternalReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionInternalReasonEnumMap(c *Client, i interface{}) map[string]ServiceGooglecloudrunopv2ConditionInternalReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGooglecloudrunopv2ConditionInternalReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceGooglecloudrunopv2ConditionInternalReasonEnum{}
	}

	items := make(map[string]ServiceGooglecloudrunopv2ConditionInternalReasonEnum)
	for k, item := range a {
		items[k] = *flattenServiceGooglecloudrunopv2ConditionInternalReasonEnum(item.(interface{}))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionInternalReasonEnumSlice flattens the contents of ServiceGooglecloudrunopv2ConditionInternalReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionInternalReasonEnumSlice(c *Client, i interface{}) []ServiceGooglecloudrunopv2ConditionInternalReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGooglecloudrunopv2ConditionInternalReasonEnum{}
	}

	if len(a) == 0 {
		return []ServiceGooglecloudrunopv2ConditionInternalReasonEnum{}
	}

	items := make([]ServiceGooglecloudrunopv2ConditionInternalReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGooglecloudrunopv2ConditionInternalReasonEnum(item.(interface{})))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionInternalReasonEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceGooglecloudrunopv2ConditionInternalReasonEnum with the same value as that string.
func flattenServiceGooglecloudrunopv2ConditionInternalReasonEnum(i interface{}) *ServiceGooglecloudrunopv2ConditionInternalReasonEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceGooglecloudrunopv2ConditionInternalReasonEnumRef("")
	}

	return ServiceGooglecloudrunopv2ConditionInternalReasonEnumRef(s)
}

// flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnumMap flattens the contents of ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnumMap(c *Client, i interface{}) map[string]ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum{}
	}

	items := make(map[string]ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum)
	for k, item := range a {
		items[k] = *flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum(item.(interface{}))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnumSlice flattens the contents of ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnumSlice(c *Client, i interface{}) []ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum{}
	}

	if len(a) == 0 {
		return []ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum{}
	}

	items := make([]ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum(item.(interface{})))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum with the same value as that string.
func flattenServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum(i interface{}) *ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnumRef("")
	}

	return ServiceGooglecloudrunopv2ConditionDomainMappingReasonEnumRef(s)
}

// flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnumMap flattens the contents of ServiceGooglecloudrunopv2ConditionRevisionReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnumMap(c *Client, i interface{}) map[string]ServiceGooglecloudrunopv2ConditionRevisionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGooglecloudrunopv2ConditionRevisionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceGooglecloudrunopv2ConditionRevisionReasonEnum{}
	}

	items := make(map[string]ServiceGooglecloudrunopv2ConditionRevisionReasonEnum)
	for k, item := range a {
		items[k] = *flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnumSlice flattens the contents of ServiceGooglecloudrunopv2ConditionRevisionReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnumSlice(c *Client, i interface{}) []ServiceGooglecloudrunopv2ConditionRevisionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGooglecloudrunopv2ConditionRevisionReasonEnum{}
	}

	if len(a) == 0 {
		return []ServiceGooglecloudrunopv2ConditionRevisionReasonEnum{}
	}

	items := make([]ServiceGooglecloudrunopv2ConditionRevisionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceGooglecloudrunopv2ConditionRevisionReasonEnum with the same value as that string.
func flattenServiceGooglecloudrunopv2ConditionRevisionReasonEnum(i interface{}) *ServiceGooglecloudrunopv2ConditionRevisionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceGooglecloudrunopv2ConditionRevisionReasonEnumRef("")
	}

	return ServiceGooglecloudrunopv2ConditionRevisionReasonEnumRef(s)
}

// flattenServiceGooglecloudrunopv2ConditionJobReasonEnumMap flattens the contents of ServiceGooglecloudrunopv2ConditionJobReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionJobReasonEnumMap(c *Client, i interface{}) map[string]ServiceGooglecloudrunopv2ConditionJobReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceGooglecloudrunopv2ConditionJobReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceGooglecloudrunopv2ConditionJobReasonEnum{}
	}

	items := make(map[string]ServiceGooglecloudrunopv2ConditionJobReasonEnum)
	for k, item := range a {
		items[k] = *flattenServiceGooglecloudrunopv2ConditionJobReasonEnum(item.(interface{}))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionJobReasonEnumSlice flattens the contents of ServiceGooglecloudrunopv2ConditionJobReasonEnum from a JSON
// response object.
func flattenServiceGooglecloudrunopv2ConditionJobReasonEnumSlice(c *Client, i interface{}) []ServiceGooglecloudrunopv2ConditionJobReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceGooglecloudrunopv2ConditionJobReasonEnum{}
	}

	if len(a) == 0 {
		return []ServiceGooglecloudrunopv2ConditionJobReasonEnum{}
	}

	items := make([]ServiceGooglecloudrunopv2ConditionJobReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceGooglecloudrunopv2ConditionJobReasonEnum(item.(interface{})))
	}

	return items
}

// flattenServiceGooglecloudrunopv2ConditionJobReasonEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceGooglecloudrunopv2ConditionJobReasonEnum with the same value as that string.
func flattenServiceGooglecloudrunopv2ConditionJobReasonEnum(i interface{}) *ServiceGooglecloudrunopv2ConditionJobReasonEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceGooglecloudrunopv2ConditionJobReasonEnumRef("")
	}

	return ServiceGooglecloudrunopv2ConditionJobReasonEnumRef(s)
}

// flattenServiceTrafficStatusesTypeEnumMap flattens the contents of ServiceTrafficStatusesTypeEnum from a JSON
// response object.
func flattenServiceTrafficStatusesTypeEnumMap(c *Client, i interface{}) map[string]ServiceTrafficStatusesTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTrafficStatusesTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTrafficStatusesTypeEnum{}
	}

	items := make(map[string]ServiceTrafficStatusesTypeEnum)
	for k, item := range a {
		items[k] = *flattenServiceTrafficStatusesTypeEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTrafficStatusesTypeEnumSlice flattens the contents of ServiceTrafficStatusesTypeEnum from a JSON
// response object.
func flattenServiceTrafficStatusesTypeEnumSlice(c *Client, i interface{}) []ServiceTrafficStatusesTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTrafficStatusesTypeEnum{}
	}

	if len(a) == 0 {
		return []ServiceTrafficStatusesTypeEnum{}
	}

	items := make([]ServiceTrafficStatusesTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTrafficStatusesTypeEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTrafficStatusesTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTrafficStatusesTypeEnum with the same value as that string.
func flattenServiceTrafficStatusesTypeEnum(i interface{}) *ServiceTrafficStatusesTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceTrafficStatusesTypeEnumRef("")
	}

	return ServiceTrafficStatusesTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Service) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalService(b, c)
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

type serviceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serviceApiOperation
}

func convertFieldDiffsToServiceDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]serviceDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []serviceDiff
	// For each operation name, create a serviceDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := serviceDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToServiceApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToServiceApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (serviceApiOperation, error) {
	switch opName {

	case "updateServiceUpdateServiceOperation":
		return &updateServiceUpdateServiceOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractServiceFields(r *Service) error {
	vBinaryAuthorization := r.BinaryAuthorization
	if vBinaryAuthorization == nil {
		// note: explicitly not the empty object.
		vBinaryAuthorization = &ServiceBinaryAuthorization{}
	}
	if err := extractServiceBinaryAuthorizationFields(r, vBinaryAuthorization); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vBinaryAuthorization) {
		r.BinaryAuthorization = vBinaryAuthorization
	}
	vTemplate := r.Template
	if vTemplate == nil {
		// note: explicitly not the empty object.
		vTemplate = &ServiceTemplate{}
	}
	if err := extractServiceTemplateFields(r, vTemplate); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTemplate) {
		r.Template = vTemplate
	}
	// *ServiceGooglecloudrunopv2Condition is a reused type - that's not compatible with function extractors.

	return nil
}
func extractServiceBinaryAuthorizationFields(r *Service, o *ServiceBinaryAuthorization) error {
	return nil
}
func extractServiceTemplateFields(r *Service, o *ServiceTemplate) error {
	vScaling := o.Scaling
	if vScaling == nil {
		// note: explicitly not the empty object.
		vScaling = &ServiceTemplateScaling{}
	}
	if err := extractServiceTemplateScalingFields(r, vScaling); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vScaling) {
		o.Scaling = vScaling
	}
	vVPCAccess := o.VPCAccess
	if vVPCAccess == nil {
		// note: explicitly not the empty object.
		vVPCAccess = &ServiceTemplateVPCAccess{}
	}
	if err := extractServiceTemplateVPCAccessFields(r, vVPCAccess); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vVPCAccess) {
		o.VPCAccess = vVPCAccess
	}
	return nil
}
func extractServiceTemplateScalingFields(r *Service, o *ServiceTemplateScaling) error {
	return nil
}
func extractServiceTemplateVPCAccessFields(r *Service, o *ServiceTemplateVPCAccess) error {
	return nil
}
func extractServiceTemplateContainersFields(r *Service, o *ServiceTemplateContainers) error {
	vResources := o.Resources
	if vResources == nil {
		// note: explicitly not the empty object.
		vResources = &ServiceTemplateContainersResources{}
	}
	if err := extractServiceTemplateContainersResourcesFields(r, vResources); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vResources) {
		o.Resources = vResources
	}
	return nil
}
func extractServiceTemplateContainersEnvFields(r *Service, o *ServiceTemplateContainersEnv) error {
	vValueSource := o.ValueSource
	if vValueSource == nil {
		// note: explicitly not the empty object.
		vValueSource = &ServiceTemplateContainersEnvValueSource{}
	}
	if err := extractServiceTemplateContainersEnvValueSourceFields(r, vValueSource); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vValueSource) {
		o.ValueSource = vValueSource
	}
	return nil
}
func extractServiceTemplateContainersEnvValueSourceFields(r *Service, o *ServiceTemplateContainersEnvValueSource) error {
	vSecretKeyRef := o.SecretKeyRef
	if vSecretKeyRef == nil {
		// note: explicitly not the empty object.
		vSecretKeyRef = &ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}
	if err := extractServiceTemplateContainersEnvValueSourceSecretKeyRefFields(r, vSecretKeyRef); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecretKeyRef) {
		o.SecretKeyRef = vSecretKeyRef
	}
	return nil
}
func extractServiceTemplateContainersEnvValueSourceSecretKeyRefFields(r *Service, o *ServiceTemplateContainersEnvValueSourceSecretKeyRef) error {
	return nil
}
func extractServiceTemplateContainersResourcesFields(r *Service, o *ServiceTemplateContainersResources) error {
	return nil
}
func extractServiceTemplateContainersPortsFields(r *Service, o *ServiceTemplateContainersPorts) error {
	return nil
}
func extractServiceTemplateContainersVolumeMountsFields(r *Service, o *ServiceTemplateContainersVolumeMounts) error {
	return nil
}
func extractServiceTemplateVolumesFields(r *Service, o *ServiceTemplateVolumes) error {
	vSecret := o.Secret
	if vSecret == nil {
		// note: explicitly not the empty object.
		vSecret = &ServiceTemplateVolumesSecret{}
	}
	if err := extractServiceTemplateVolumesSecretFields(r, vSecret); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecret) {
		o.Secret = vSecret
	}
	vCloudSqlInstance := o.CloudSqlInstance
	if vCloudSqlInstance == nil {
		// note: explicitly not the empty object.
		vCloudSqlInstance = &ServiceTemplateVolumesCloudSqlInstance{}
	}
	if err := extractServiceTemplateVolumesCloudSqlInstanceFields(r, vCloudSqlInstance); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCloudSqlInstance) {
		o.CloudSqlInstance = vCloudSqlInstance
	}
	return nil
}
func extractServiceTemplateVolumesSecretFields(r *Service, o *ServiceTemplateVolumesSecret) error {
	return nil
}
func extractServiceTemplateVolumesSecretItemsFields(r *Service, o *ServiceTemplateVolumesSecretItems) error {
	return nil
}
func extractServiceTemplateVolumesCloudSqlInstanceFields(r *Service, o *ServiceTemplateVolumesCloudSqlInstance) error {
	return nil
}
func extractServiceTrafficFields(r *Service, o *ServiceTraffic) error {
	return nil
}
func extractServiceGooglecloudrunopv2ConditionFields(r *Service, o *ServiceGooglecloudrunopv2Condition) error {
	return nil
}
func extractServiceTrafficStatusesFields(r *Service, o *ServiceTrafficStatuses) error {
	return nil
}

func postReadExtractServiceFields(r *Service) error {
	vBinaryAuthorization := r.BinaryAuthorization
	if vBinaryAuthorization == nil {
		// note: explicitly not the empty object.
		vBinaryAuthorization = &ServiceBinaryAuthorization{}
	}
	if err := postReadExtractServiceBinaryAuthorizationFields(r, vBinaryAuthorization); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vBinaryAuthorization) {
		r.BinaryAuthorization = vBinaryAuthorization
	}
	vTemplate := r.Template
	if vTemplate == nil {
		// note: explicitly not the empty object.
		vTemplate = &ServiceTemplate{}
	}
	if err := postReadExtractServiceTemplateFields(r, vTemplate); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTemplate) {
		r.Template = vTemplate
	}
	// *ServiceGooglecloudrunopv2Condition is a reused type - that's not compatible with function extractors.

	return nil
}
func postReadExtractServiceBinaryAuthorizationFields(r *Service, o *ServiceBinaryAuthorization) error {
	return nil
}
func postReadExtractServiceTemplateFields(r *Service, o *ServiceTemplate) error {
	vScaling := o.Scaling
	if vScaling == nil {
		// note: explicitly not the empty object.
		vScaling = &ServiceTemplateScaling{}
	}
	if err := extractServiceTemplateScalingFields(r, vScaling); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vScaling) {
		o.Scaling = vScaling
	}
	vVPCAccess := o.VPCAccess
	if vVPCAccess == nil {
		// note: explicitly not the empty object.
		vVPCAccess = &ServiceTemplateVPCAccess{}
	}
	if err := extractServiceTemplateVPCAccessFields(r, vVPCAccess); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vVPCAccess) {
		o.VPCAccess = vVPCAccess
	}
	return nil
}
func postReadExtractServiceTemplateScalingFields(r *Service, o *ServiceTemplateScaling) error {
	return nil
}
func postReadExtractServiceTemplateVPCAccessFields(r *Service, o *ServiceTemplateVPCAccess) error {
	return nil
}
func postReadExtractServiceTemplateContainersFields(r *Service, o *ServiceTemplateContainers) error {
	vResources := o.Resources
	if vResources == nil {
		// note: explicitly not the empty object.
		vResources = &ServiceTemplateContainersResources{}
	}
	if err := extractServiceTemplateContainersResourcesFields(r, vResources); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vResources) {
		o.Resources = vResources
	}
	return nil
}
func postReadExtractServiceTemplateContainersEnvFields(r *Service, o *ServiceTemplateContainersEnv) error {
	vValueSource := o.ValueSource
	if vValueSource == nil {
		// note: explicitly not the empty object.
		vValueSource = &ServiceTemplateContainersEnvValueSource{}
	}
	if err := extractServiceTemplateContainersEnvValueSourceFields(r, vValueSource); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vValueSource) {
		o.ValueSource = vValueSource
	}
	return nil
}
func postReadExtractServiceTemplateContainersEnvValueSourceFields(r *Service, o *ServiceTemplateContainersEnvValueSource) error {
	vSecretKeyRef := o.SecretKeyRef
	if vSecretKeyRef == nil {
		// note: explicitly not the empty object.
		vSecretKeyRef = &ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}
	if err := extractServiceTemplateContainersEnvValueSourceSecretKeyRefFields(r, vSecretKeyRef); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecretKeyRef) {
		o.SecretKeyRef = vSecretKeyRef
	}
	return nil
}
func postReadExtractServiceTemplateContainersEnvValueSourceSecretKeyRefFields(r *Service, o *ServiceTemplateContainersEnvValueSourceSecretKeyRef) error {
	return nil
}
func postReadExtractServiceTemplateContainersResourcesFields(r *Service, o *ServiceTemplateContainersResources) error {
	return nil
}
func postReadExtractServiceTemplateContainersPortsFields(r *Service, o *ServiceTemplateContainersPorts) error {
	return nil
}
func postReadExtractServiceTemplateContainersVolumeMountsFields(r *Service, o *ServiceTemplateContainersVolumeMounts) error {
	return nil
}
func postReadExtractServiceTemplateVolumesFields(r *Service, o *ServiceTemplateVolumes) error {
	vSecret := o.Secret
	if vSecret == nil {
		// note: explicitly not the empty object.
		vSecret = &ServiceTemplateVolumesSecret{}
	}
	if err := extractServiceTemplateVolumesSecretFields(r, vSecret); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecret) {
		o.Secret = vSecret
	}
	vCloudSqlInstance := o.CloudSqlInstance
	if vCloudSqlInstance == nil {
		// note: explicitly not the empty object.
		vCloudSqlInstance = &ServiceTemplateVolumesCloudSqlInstance{}
	}
	if err := extractServiceTemplateVolumesCloudSqlInstanceFields(r, vCloudSqlInstance); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCloudSqlInstance) {
		o.CloudSqlInstance = vCloudSqlInstance
	}
	return nil
}
func postReadExtractServiceTemplateVolumesSecretFields(r *Service, o *ServiceTemplateVolumesSecret) error {
	return nil
}
func postReadExtractServiceTemplateVolumesSecretItemsFields(r *Service, o *ServiceTemplateVolumesSecretItems) error {
	return nil
}
func postReadExtractServiceTemplateVolumesCloudSqlInstanceFields(r *Service, o *ServiceTemplateVolumesCloudSqlInstance) error {
	return nil
}
func postReadExtractServiceTrafficFields(r *Service, o *ServiceTraffic) error {
	return nil
}
func postReadExtractServiceGooglecloudrunopv2ConditionFields(r *Service, o *ServiceGooglecloudrunopv2Condition) error {
	return nil
}
func postReadExtractServiceTrafficStatusesFields(r *Service, o *ServiceTrafficStatuses) error {
	return nil
}
