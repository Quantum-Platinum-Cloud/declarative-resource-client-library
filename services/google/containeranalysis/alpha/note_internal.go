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
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Note) validate() error {

	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"Image", "Package", "Deployment", "Discovery", "Attestation"}, r.Image, r.Package, r.Deployment, r.Discovery, r.Attestation); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Image) {
		if err := r.Image.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Package) {
		if err := r.Package.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Discovery) {
		if err := r.Discovery.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Deployment) {
		if err := r.Deployment.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Attestation) {
		if err := r.Attestation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteRelatedUrl) validate() error {
	return nil
}
func (r *NoteImage) validate() error {
	if err := dcl.Required(r, "resourceUrl"); err != nil {
		return err
	}
	if err := dcl.Required(r, "fingerprint"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Fingerprint) {
		if err := r.Fingerprint.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteImageFingerprint) validate() error {
	if err := dcl.Required(r, "v1Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "v2Blob"); err != nil {
		return err
	}
	return nil
}
func (r *NotePackage) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	return nil
}
func (r *NotePackageDistribution) validate() error {
	if err := dcl.Required(r, "cpeUri"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.LatestVersion) {
		if err := r.LatestVersion.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NotePackageDistributionLatestVersion) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteDiscovery) validate() error {
	if err := dcl.Required(r, "analysisKind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteDeployment) validate() error {
	if err := dcl.Required(r, "resourceUri"); err != nil {
		return err
	}
	return nil
}
func (r *NoteAttestation) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Hint) {
		if err := r.Hint.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteAttestationHint) validate() error {
	if err := dcl.Required(r, "humanReadableName"); err != nil {
		return err
	}
	return nil
}
func (r *Note) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://containeranalysis.googleapis.com/v1alpha1/", params)
}

func (r *Note) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/notes/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Note) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/notes", nr.basePath(), userBasePath, params), nil

}

func (r *Note) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/notes?noteId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Note) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/notes/{{name}}", nr.basePath(), userBasePath, params), nil
}

// noteApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type noteApiOperation interface {
	do(context.Context, *Note, *Client) error
}

// newUpdateNoteUpdateNoteRequest creates a request for an
// Note resource's UpdateNote update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNoteUpdateNoteRequest(ctx context.Context, f *Note, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.ShortDescription; !dcl.IsEmptyValueIndirect(v) {
		req["shortDescription"] = v
	}
	if v := f.LongDescription; !dcl.IsEmptyValueIndirect(v) {
		req["longDescription"] = v
	}
	if v, err := expandNoteRelatedUrlSlice(c, f.RelatedUrl); err != nil {
		return nil, fmt.Errorf("error expanding RelatedUrl into relatedUrl: %w", err)
	} else if v != nil {
		req["relatedUrl"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		req["expirationTime"] = v
	}
	if v, err := expandNoteImage(c, f.Image); err != nil {
		return nil, fmt.Errorf("error expanding Image into image: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["image"] = v
	}
	if v, err := expandNotePackage(c, f.Package); err != nil {
		return nil, fmt.Errorf("error expanding Package into package: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["package"] = v
	}
	if v, err := expandNoteDiscovery(c, f.Discovery); err != nil {
		return nil, fmt.Errorf("error expanding Discovery into discovery: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["discovery"] = v
	}
	if v, err := expandNoteDeployment(c, f.Deployment); err != nil {
		return nil, fmt.Errorf("error expanding Deployment into deployment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["deployment"] = v
	}
	if v, err := expandNoteAttestation(c, f.Attestation); err != nil {
		return nil, fmt.Errorf("error expanding Attestation into attestation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["attestation"] = v
	}
	return req, nil
}

// marshalUpdateNoteUpdateNoteRequest converts the update into
// the final JSON request body.
func marshalUpdateNoteUpdateNoteRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNoteUpdateNoteOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNoteUpdateNoteOperation) do(ctx context.Context, r *Note, c *Client) error {
	_, err := c.GetNote(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateNote")
	if err != nil {
		return err
	}

	req, err := newUpdateNoteUpdateNoteRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateNoteUpdateNoteRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listNoteRaw(ctx context.Context, r *Note, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != NoteMaxPage {
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

type listNoteOperation struct {
	Notes []map[string]interface{} `json:"notes"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listNote(ctx context.Context, r *Note, pageToken string, pageSize int32) ([]*Note, string, error) {
	b, err := c.listNoteRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listNoteOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Note
	for _, v := range m.Notes {
		res, err := unmarshalMapNote(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNote(ctx context.Context, f func(*Note) bool, resources []*Note) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNote(ctx, res)
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

type deleteNoteOperation struct{}

func (op *deleteNoteOperation) do(ctx context.Context, r *Note, c *Client) error {
	r, err := c.GetNote(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Note not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetNote checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Note: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetNote(ctx, r)
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createNoteOperation struct {
	response map[string]interface{}
}

func (op *createNoteOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNoteOperation) do(ctx context.Context, r *Note, c *Client) error {
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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetNote(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNoteRaw(ctx context.Context, r *Note) ([]byte, error) {

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

func (c *Client) noteDiffsForRawDesired(ctx context.Context, rawDesired *Note, opts ...dcl.ApplyOption) (initial, desired *Note, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Note
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Note); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Note, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNote(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Note resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Note resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Note resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNoteDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Note: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Note: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNoteInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Note: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNoteDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Note: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNote(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeNoteInitialState(rawInitial, rawDesired *Note) (*Note, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.Image) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Image = EmptyNoteImage
		}
	}

	if !dcl.IsZeroValue(rawInitial.Package) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Image, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Package = EmptyNotePackage
		}
	}

	if !dcl.IsZeroValue(rawInitial.Deployment) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Image, rawInitial.Package, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Deployment = EmptyNoteDeployment
		}
	}

	if !dcl.IsZeroValue(rawInitial.Discovery) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Attestation) {
			rawInitial.Discovery = EmptyNoteDiscovery
		}
	}

	if !dcl.IsZeroValue(rawInitial.Attestation) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery) {
			rawInitial.Attestation = EmptyNoteAttestation
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeNoteDesiredState(rawDesired, rawInitial *Note, opts ...dcl.ApplyOption) (*Note, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Image = canonicalizeNoteImage(rawDesired.Image, nil, opts...)
		rawDesired.Package = canonicalizeNotePackage(rawDesired.Package, nil, opts...)
		rawDesired.Discovery = canonicalizeNoteDiscovery(rawDesired.Discovery, nil, opts...)
		rawDesired.Deployment = canonicalizeNoteDeployment(rawDesired.Deployment, nil, opts...)
		rawDesired.Attestation = canonicalizeNoteAttestation(rawDesired.Attestation, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.Image != nil || rawInitial.Image != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			rawDesired.Image = nil
			rawInitial.Image = nil
		}
	}

	if rawDesired.Package != nil || rawInitial.Package != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Image, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			rawDesired.Package = nil
			rawInitial.Package = nil
		}
	}

	if rawDesired.Deployment != nil || rawInitial.Deployment != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Image, rawDesired.Package, rawDesired.Discovery, rawDesired.Attestation) {
			rawDesired.Deployment = nil
			rawInitial.Deployment = nil
		}
	}

	if rawDesired.Discovery != nil || rawInitial.Discovery != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Attestation) {
			rawDesired.Discovery = nil
			rawInitial.Discovery = nil
		}
	}

	if rawDesired.Attestation != nil || rawInitial.Attestation != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery) {
			rawDesired.Attestation = nil
			rawInitial.Attestation = nil
		}
	}

	canonicalDesired := &Note{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.ShortDescription, rawInitial.ShortDescription) {
		canonicalDesired.ShortDescription = rawInitial.ShortDescription
	} else {
		canonicalDesired.ShortDescription = rawDesired.ShortDescription
	}
	if dcl.StringCanonicalize(rawDesired.LongDescription, rawInitial.LongDescription) {
		canonicalDesired.LongDescription = rawInitial.LongDescription
	} else {
		canonicalDesired.LongDescription = rawDesired.LongDescription
	}
	canonicalDesired.RelatedUrl = canonicalizeNoteRelatedUrlSlice(rawDesired.RelatedUrl, rawInitial.RelatedUrl, opts...)
	if dcl.IsZeroValue(rawDesired.ExpirationTime) {
		canonicalDesired.ExpirationTime = rawInitial.ExpirationTime
	} else {
		canonicalDesired.ExpirationTime = rawDesired.ExpirationTime
	}
	canonicalDesired.Image = canonicalizeNoteImage(rawDesired.Image, rawInitial.Image, opts...)
	canonicalDesired.Package = canonicalizeNotePackage(rawDesired.Package, rawInitial.Package, opts...)
	canonicalDesired.Discovery = canonicalizeNoteDiscovery(rawDesired.Discovery, rawInitial.Discovery, opts...)
	canonicalDesired.Deployment = canonicalizeNoteDeployment(rawDesired.Deployment, rawInitial.Deployment, opts...)
	canonicalDesired.Attestation = canonicalizeNoteAttestation(rawDesired.Attestation, rawInitial.Attestation, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeNoteNewState(c *Client, rawNew, rawDesired *Note) (*Note, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.ShortDescription) && dcl.IsNotReturnedByServer(rawDesired.ShortDescription) {
		rawNew.ShortDescription = rawDesired.ShortDescription
	} else {
		if dcl.StringCanonicalize(rawDesired.ShortDescription, rawNew.ShortDescription) {
			rawNew.ShortDescription = rawDesired.ShortDescription
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.LongDescription) && dcl.IsNotReturnedByServer(rawDesired.LongDescription) {
		rawNew.LongDescription = rawDesired.LongDescription
	} else {
		if dcl.StringCanonicalize(rawDesired.LongDescription, rawNew.LongDescription) {
			rawNew.LongDescription = rawDesired.LongDescription
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.RelatedUrl) && dcl.IsNotReturnedByServer(rawDesired.RelatedUrl) {
		rawNew.RelatedUrl = rawDesired.RelatedUrl
	} else {
		rawNew.RelatedUrl = canonicalizeNewNoteRelatedUrlSlice(c, rawDesired.RelatedUrl, rawNew.RelatedUrl)
	}

	if dcl.IsNotReturnedByServer(rawNew.ExpirationTime) && dcl.IsNotReturnedByServer(rawDesired.ExpirationTime) {
		rawNew.ExpirationTime = rawDesired.ExpirationTime
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

	if dcl.IsNotReturnedByServer(rawNew.Image) && dcl.IsNotReturnedByServer(rawDesired.Image) {
		rawNew.Image = rawDesired.Image
	} else {
		rawNew.Image = canonicalizeNewNoteImage(c, rawDesired.Image, rawNew.Image)
	}

	if dcl.IsNotReturnedByServer(rawNew.Package) && dcl.IsNotReturnedByServer(rawDesired.Package) {
		rawNew.Package = rawDesired.Package
	} else {
		rawNew.Package = canonicalizeNewNotePackage(c, rawDesired.Package, rawNew.Package)
	}

	if dcl.IsNotReturnedByServer(rawNew.Discovery) && dcl.IsNotReturnedByServer(rawDesired.Discovery) {
		rawNew.Discovery = rawDesired.Discovery
	} else {
		rawNew.Discovery = canonicalizeNewNoteDiscovery(c, rawDesired.Discovery, rawNew.Discovery)
	}

	if dcl.IsNotReturnedByServer(rawNew.Deployment) && dcl.IsNotReturnedByServer(rawDesired.Deployment) {
		rawNew.Deployment = rawDesired.Deployment
	} else {
		rawNew.Deployment = canonicalizeNewNoteDeployment(c, rawDesired.Deployment, rawNew.Deployment)
	}

	if dcl.IsNotReturnedByServer(rawNew.Attestation) && dcl.IsNotReturnedByServer(rawDesired.Attestation) {
		rawNew.Attestation = rawDesired.Attestation
	} else {
		rawNew.Attestation = canonicalizeNewNoteAttestation(c, rawDesired.Attestation, rawNew.Attestation)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeNoteRelatedUrl(des, initial *NoteRelatedUrl, opts ...dcl.ApplyOption) *NoteRelatedUrl {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteRelatedUrl{}

	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		cDes.Url = initial.Url
	} else {
		cDes.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Label, initial.Label) || dcl.IsZeroValue(des.Label) {
		cDes.Label = initial.Label
	} else {
		cDes.Label = des.Label
	}

	return cDes
}

func canonicalizeNoteRelatedUrlSlice(des, initial []NoteRelatedUrl, opts ...dcl.ApplyOption) []NoteRelatedUrl {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteRelatedUrl, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteRelatedUrl(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteRelatedUrl, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteRelatedUrl(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteRelatedUrl(c *Client, des, nw *NoteRelatedUrl) *NoteRelatedUrl {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteRelatedUrl while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Label, nw.Label) {
		nw.Label = des.Label
	}

	return nw
}

func canonicalizeNewNoteRelatedUrlSet(c *Client, des, nw []NoteRelatedUrl) []NoteRelatedUrl {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteRelatedUrl
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNoteRelatedUrlNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNoteRelatedUrlSlice(c *Client, des, nw []NoteRelatedUrl) []NoteRelatedUrl {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteRelatedUrl
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteRelatedUrl(c, &d, &n))
	}

	return items
}

func canonicalizeNoteImage(des, initial *NoteImage, opts ...dcl.ApplyOption) *NoteImage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteImage{}

	if dcl.StringCanonicalize(des.ResourceUrl, initial.ResourceUrl) || dcl.IsZeroValue(des.ResourceUrl) {
		cDes.ResourceUrl = initial.ResourceUrl
	} else {
		cDes.ResourceUrl = des.ResourceUrl
	}
	cDes.Fingerprint = canonicalizeNoteImageFingerprint(des.Fingerprint, initial.Fingerprint, opts...)

	return cDes
}

func canonicalizeNoteImageSlice(des, initial []NoteImage, opts ...dcl.ApplyOption) []NoteImage {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteImage, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteImage(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteImage, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteImage(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteImage(c *Client, des, nw *NoteImage) *NoteImage {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteImage while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ResourceUrl, nw.ResourceUrl) {
		nw.ResourceUrl = des.ResourceUrl
	}
	nw.Fingerprint = canonicalizeNewNoteImageFingerprint(c, des.Fingerprint, nw.Fingerprint)

	return nw
}

func canonicalizeNewNoteImageSet(c *Client, des, nw []NoteImage) []NoteImage {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteImage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNoteImageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNoteImageSlice(c *Client, des, nw []NoteImage) []NoteImage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteImage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteImage(c, &d, &n))
	}

	return items
}

func canonicalizeNoteImageFingerprint(des, initial *NoteImageFingerprint, opts ...dcl.ApplyOption) *NoteImageFingerprint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteImageFingerprint{}

	if dcl.StringCanonicalize(des.V1Name, initial.V1Name) || dcl.IsZeroValue(des.V1Name) {
		cDes.V1Name = initial.V1Name
	} else {
		cDes.V1Name = des.V1Name
	}
	if dcl.StringArrayCanonicalize(des.V2Blob, initial.V2Blob) || dcl.IsZeroValue(des.V2Blob) {
		cDes.V2Blob = initial.V2Blob
	} else {
		cDes.V2Blob = des.V2Blob
	}

	return cDes
}

func canonicalizeNoteImageFingerprintSlice(des, initial []NoteImageFingerprint, opts ...dcl.ApplyOption) []NoteImageFingerprint {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteImageFingerprint, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteImageFingerprint(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteImageFingerprint, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteImageFingerprint(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteImageFingerprint(c *Client, des, nw *NoteImageFingerprint) *NoteImageFingerprint {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteImageFingerprint while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.V1Name, nw.V1Name) {
		nw.V1Name = des.V1Name
	}
	if dcl.StringArrayCanonicalize(des.V2Blob, nw.V2Blob) {
		nw.V2Blob = des.V2Blob
	}
	if dcl.StringCanonicalize(des.V2Name, nw.V2Name) {
		nw.V2Name = des.V2Name
	}

	return nw
}

func canonicalizeNewNoteImageFingerprintSet(c *Client, des, nw []NoteImageFingerprint) []NoteImageFingerprint {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteImageFingerprint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNoteImageFingerprintNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNoteImageFingerprintSlice(c *Client, des, nw []NoteImageFingerprint) []NoteImageFingerprint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteImageFingerprint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteImageFingerprint(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackage(des, initial *NotePackage, opts ...dcl.ApplyOption) *NotePackage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NotePackage{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.Distribution = canonicalizeNotePackageDistributionSlice(des.Distribution, initial.Distribution, opts...)

	return cDes
}

func canonicalizeNotePackageSlice(des, initial []NotePackage, opts ...dcl.ApplyOption) []NotePackage {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NotePackage, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNotePackage(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NotePackage, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNotePackage(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNotePackage(c *Client, des, nw *NotePackage) *NotePackage {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NotePackage while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Distribution = canonicalizeNewNotePackageDistributionSlice(c, des.Distribution, nw.Distribution)

	return nw
}

func canonicalizeNewNotePackageSet(c *Client, des, nw []NotePackage) []NotePackage {
	if des == nil {
		return nw
	}
	var reorderedNew []NotePackage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNotePackageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNotePackageSlice(c *Client, des, nw []NotePackage) []NotePackage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NotePackage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackage(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackageDistribution(des, initial *NotePackageDistribution, opts ...dcl.ApplyOption) *NotePackageDistribution {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NotePackageDistribution{}

	if dcl.StringCanonicalize(des.CpeUri, initial.CpeUri) || dcl.IsZeroValue(des.CpeUri) {
		cDes.CpeUri = initial.CpeUri
	} else {
		cDes.CpeUri = des.CpeUri
	}
	if dcl.IsZeroValue(des.Architecture) {
		des.Architecture = initial.Architecture
	} else {
		cDes.Architecture = des.Architecture
	}
	cDes.LatestVersion = canonicalizeNotePackageDistributionLatestVersion(des.LatestVersion, initial.LatestVersion, opts...)
	if dcl.StringCanonicalize(des.Maintainer, initial.Maintainer) || dcl.IsZeroValue(des.Maintainer) {
		cDes.Maintainer = initial.Maintainer
	} else {
		cDes.Maintainer = des.Maintainer
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		cDes.Url = initial.Url
	} else {
		cDes.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}

	return cDes
}

func canonicalizeNotePackageDistributionSlice(des, initial []NotePackageDistribution, opts ...dcl.ApplyOption) []NotePackageDistribution {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NotePackageDistribution, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNotePackageDistribution(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NotePackageDistribution, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNotePackageDistribution(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNotePackageDistribution(c *Client, des, nw *NotePackageDistribution) *NotePackageDistribution {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NotePackageDistribution while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.CpeUri, nw.CpeUri) {
		nw.CpeUri = des.CpeUri
	}
	nw.LatestVersion = canonicalizeNewNotePackageDistributionLatestVersion(c, des.LatestVersion, nw.LatestVersion)
	if dcl.StringCanonicalize(des.Maintainer, nw.Maintainer) {
		nw.Maintainer = des.Maintainer
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewNotePackageDistributionSet(c *Client, des, nw []NotePackageDistribution) []NotePackageDistribution {
	if des == nil {
		return nw
	}
	var reorderedNew []NotePackageDistribution
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNotePackageDistributionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNotePackageDistributionSlice(c *Client, des, nw []NotePackageDistribution) []NotePackageDistribution {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NotePackageDistribution
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackageDistribution(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackageDistributionLatestVersion(des, initial *NotePackageDistributionLatestVersion, opts ...dcl.ApplyOption) *NotePackageDistributionLatestVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NotePackageDistributionLatestVersion{}

	if dcl.IsZeroValue(des.Epoch) {
		des.Epoch = initial.Epoch
	} else {
		cDes.Epoch = des.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	} else {
		cDes.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		cDes.FullName = initial.FullName
	} else {
		cDes.FullName = des.FullName
	}

	return cDes
}

func canonicalizeNotePackageDistributionLatestVersionSlice(des, initial []NotePackageDistributionLatestVersion, opts ...dcl.ApplyOption) []NotePackageDistributionLatestVersion {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NotePackageDistributionLatestVersion, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNotePackageDistributionLatestVersion(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NotePackageDistributionLatestVersion, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNotePackageDistributionLatestVersion(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNotePackageDistributionLatestVersion(c *Client, des, nw *NotePackageDistributionLatestVersion) *NotePackageDistributionLatestVersion {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NotePackageDistributionLatestVersion while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNotePackageDistributionLatestVersionSet(c *Client, des, nw []NotePackageDistributionLatestVersion) []NotePackageDistributionLatestVersion {
	if des == nil {
		return nw
	}
	var reorderedNew []NotePackageDistributionLatestVersion
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNotePackageDistributionLatestVersionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNotePackageDistributionLatestVersionSlice(c *Client, des, nw []NotePackageDistributionLatestVersion) []NotePackageDistributionLatestVersion {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NotePackageDistributionLatestVersion
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackageDistributionLatestVersion(c, &d, &n))
	}

	return items
}

func canonicalizeNoteDiscovery(des, initial *NoteDiscovery, opts ...dcl.ApplyOption) *NoteDiscovery {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteDiscovery{}

	if dcl.IsZeroValue(des.AnalysisKind) {
		des.AnalysisKind = initial.AnalysisKind
	} else {
		cDes.AnalysisKind = des.AnalysisKind
	}

	return cDes
}

func canonicalizeNoteDiscoverySlice(des, initial []NoteDiscovery, opts ...dcl.ApplyOption) []NoteDiscovery {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteDiscovery, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteDiscovery(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteDiscovery, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteDiscovery(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteDiscovery(c *Client, des, nw *NoteDiscovery) *NoteDiscovery {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteDiscovery while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewNoteDiscoverySet(c *Client, des, nw []NoteDiscovery) []NoteDiscovery {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteDiscovery
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNoteDiscoveryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNoteDiscoverySlice(c *Client, des, nw []NoteDiscovery) []NoteDiscovery {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteDiscovery
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteDiscovery(c, &d, &n))
	}

	return items
}

func canonicalizeNoteDeployment(des, initial *NoteDeployment, opts ...dcl.ApplyOption) *NoteDeployment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteDeployment{}

	if dcl.StringArrayCanonicalize(des.ResourceUri, initial.ResourceUri) || dcl.IsZeroValue(des.ResourceUri) {
		cDes.ResourceUri = initial.ResourceUri
	} else {
		cDes.ResourceUri = des.ResourceUri
	}

	return cDes
}

func canonicalizeNoteDeploymentSlice(des, initial []NoteDeployment, opts ...dcl.ApplyOption) []NoteDeployment {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteDeployment, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteDeployment(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteDeployment, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteDeployment(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteDeployment(c *Client, des, nw *NoteDeployment) *NoteDeployment {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteDeployment while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.ResourceUri, nw.ResourceUri) {
		nw.ResourceUri = des.ResourceUri
	}

	return nw
}

func canonicalizeNewNoteDeploymentSet(c *Client, des, nw []NoteDeployment) []NoteDeployment {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteDeployment
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNoteDeploymentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNoteDeploymentSlice(c *Client, des, nw []NoteDeployment) []NoteDeployment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteDeployment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteDeployment(c, &d, &n))
	}

	return items
}

func canonicalizeNoteAttestation(des, initial *NoteAttestation, opts ...dcl.ApplyOption) *NoteAttestation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteAttestation{}

	cDes.Hint = canonicalizeNoteAttestationHint(des.Hint, initial.Hint, opts...)

	return cDes
}

func canonicalizeNoteAttestationSlice(des, initial []NoteAttestation, opts ...dcl.ApplyOption) []NoteAttestation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteAttestation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteAttestation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteAttestation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteAttestation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteAttestation(c *Client, des, nw *NoteAttestation) *NoteAttestation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteAttestation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Hint = canonicalizeNewNoteAttestationHint(c, des.Hint, nw.Hint)

	return nw
}

func canonicalizeNewNoteAttestationSet(c *Client, des, nw []NoteAttestation) []NoteAttestation {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteAttestation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNoteAttestationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNoteAttestationSlice(c *Client, des, nw []NoteAttestation) []NoteAttestation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteAttestation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteAttestation(c, &d, &n))
	}

	return items
}

func canonicalizeNoteAttestationHint(des, initial *NoteAttestationHint, opts ...dcl.ApplyOption) *NoteAttestationHint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteAttestationHint{}

	if dcl.StringCanonicalize(des.HumanReadableName, initial.HumanReadableName) || dcl.IsZeroValue(des.HumanReadableName) {
		cDes.HumanReadableName = initial.HumanReadableName
	} else {
		cDes.HumanReadableName = des.HumanReadableName
	}

	return cDes
}

func canonicalizeNoteAttestationHintSlice(des, initial []NoteAttestationHint, opts ...dcl.ApplyOption) []NoteAttestationHint {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteAttestationHint, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteAttestationHint(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteAttestationHint, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteAttestationHint(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteAttestationHint(c *Client, des, nw *NoteAttestationHint) *NoteAttestationHint {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteAttestationHint while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.HumanReadableName, nw.HumanReadableName) {
		nw.HumanReadableName = des.HumanReadableName
	}

	return nw
}

func canonicalizeNewNoteAttestationHintSet(c *Client, des, nw []NoteAttestationHint) []NoteAttestationHint {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteAttestationHint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNoteAttestationHintNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNoteAttestationHintSlice(c *Client, des, nw []NoteAttestationHint) []NoteAttestationHint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteAttestationHint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteAttestationHint(c, &d, &n))
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
func diffNote(c *Client, desired, actual *Note, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ShortDescription, actual.ShortDescription, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ShortDescription")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LongDescription, actual.LongDescription, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("LongDescription")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RelatedUrl, actual.RelatedUrl, dcl.Info{ObjectFunction: compareNoteRelatedUrlNewStyle, EmptyObject: EmptyNoteRelatedUrl, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("RelatedUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpirationTime, actual.ExpirationTime, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ExpirationTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.Info{ObjectFunction: compareNoteImageNewStyle, EmptyObject: EmptyNoteImage, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Image")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Package, actual.Package, dcl.Info{ObjectFunction: compareNotePackageNewStyle, EmptyObject: EmptyNotePackage, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Package")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Discovery, actual.Discovery, dcl.Info{ObjectFunction: compareNoteDiscoveryNewStyle, EmptyObject: EmptyNoteDiscovery, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Discovery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deployment, actual.Deployment, dcl.Info{ObjectFunction: compareNoteDeploymentNewStyle, EmptyObject: EmptyNoteDeployment, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Deployment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Attestation, actual.Attestation, dcl.Info{ObjectFunction: compareNoteAttestationNewStyle, EmptyObject: EmptyNoteAttestation, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Attestation")); len(ds) != 0 || err != nil {
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

	newDiffs = analyzeNoteDiff(desired, actual, newDiffs)
	return newDiffs, nil
}
func compareNoteRelatedUrlNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteRelatedUrl)
	if !ok {
		desiredNotPointer, ok := d.(NoteRelatedUrl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteRelatedUrl or *NoteRelatedUrl", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteRelatedUrl)
	if !ok {
		actualNotPointer, ok := a.(NoteRelatedUrl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteRelatedUrl", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Label, actual.Label, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Label")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteImageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteImage)
	if !ok {
		desiredNotPointer, ok := d.(NoteImage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteImage or *NoteImage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteImage)
	if !ok {
		actualNotPointer, ok := a.(NoteImage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteImage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceUrl, actual.ResourceUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Fingerprint, actual.Fingerprint, dcl.Info{ObjectFunction: compareNoteImageFingerprintNewStyle, EmptyObject: EmptyNoteImageFingerprint, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Fingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteImageFingerprintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteImageFingerprint)
	if !ok {
		desiredNotPointer, ok := d.(NoteImageFingerprint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteImageFingerprint or *NoteImageFingerprint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteImageFingerprint)
	if !ok {
		actualNotPointer, ok := a.(NoteImageFingerprint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteImageFingerprint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.V1Name, actual.V1Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("V1Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.V2Blob, actual.V2Blob, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("V2Blob")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.V2Name, actual.V2Name, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("V2Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNotePackageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NotePackage)
	if !ok {
		desiredNotPointer, ok := d.(NotePackage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackage or *NotePackage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NotePackage)
	if !ok {
		actualNotPointer, ok := a.(NotePackage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Distribution, actual.Distribution, dcl.Info{ObjectFunction: compareNotePackageDistributionNewStyle, EmptyObject: EmptyNotePackageDistribution, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Distribution")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNotePackageDistributionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NotePackageDistribution)
	if !ok {
		desiredNotPointer, ok := d.(NotePackageDistribution)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackageDistribution or *NotePackageDistribution", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NotePackageDistribution)
	if !ok {
		actualNotPointer, ok := a.(NotePackageDistribution)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackageDistribution", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CpeUri, actual.CpeUri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CpeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Architecture, actual.Architecture, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Architecture")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestVersion, actual.LatestVersion, dcl.Info{ObjectFunction: compareNotePackageDistributionLatestVersionNewStyle, EmptyObject: EmptyNotePackageDistributionLatestVersion, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LatestVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Maintainer, actual.Maintainer, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Maintainer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNotePackageDistributionLatestVersionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NotePackageDistributionLatestVersion)
	if !ok {
		desiredNotPointer, ok := d.(NotePackageDistributionLatestVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackageDistributionLatestVersion or *NotePackageDistributionLatestVersion", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NotePackageDistributionLatestVersion)
	if !ok {
		actualNotPointer, ok := a.(NotePackageDistributionLatestVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackageDistributionLatestVersion", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Epoch, actual.Epoch, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Epoch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullName, actual.FullName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FullName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteDiscoveryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteDiscovery)
	if !ok {
		desiredNotPointer, ok := d.(NoteDiscovery)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteDiscovery or *NoteDiscovery", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteDiscovery)
	if !ok {
		actualNotPointer, ok := a.(NoteDiscovery)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteDiscovery", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AnalysisKind, actual.AnalysisKind, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AnalysisKind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteDeploymentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteDeployment)
	if !ok {
		desiredNotPointer, ok := d.(NoteDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteDeployment or *NoteDeployment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteDeployment)
	if !ok {
		actualNotPointer, ok := a.(NoteDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteDeployment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceUri, actual.ResourceUri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteAttestationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteAttestation)
	if !ok {
		desiredNotPointer, ok := d.(NoteAttestation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteAttestation or *NoteAttestation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteAttestation)
	if !ok {
		actualNotPointer, ok := a.(NoteAttestation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteAttestation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Hint, actual.Hint, dcl.Info{ObjectFunction: compareNoteAttestationHintNewStyle, EmptyObject: EmptyNoteAttestationHint, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Hint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteAttestationHintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteAttestationHint)
	if !ok {
		desiredNotPointer, ok := d.(NoteAttestationHint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteAttestationHint or *NoteAttestationHint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteAttestationHint)
	if !ok {
		actualNotPointer, ok := a.(NoteAttestationHint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteAttestationHint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HumanReadableName, actual.HumanReadableName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HumanReadableName")); len(ds) != 0 || err != nil {
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
func (r *Note) urlNormalized() *Note {
	normalized := dcl.Copy(*r).(Note)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.ShortDescription = dcl.SelfLinkToName(r.ShortDescription)
	normalized.LongDescription = dcl.SelfLinkToName(r.LongDescription)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Note) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateNote" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/notes/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Note resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Note) marshal(c *Client) ([]byte, error) {
	m, err := expandNote(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Note: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalNote decodes JSON responses into the Note resource schema.
func unmarshalNote(b []byte, c *Client) (*Note, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNote(m, c)
}

func unmarshalMapNote(m map[string]interface{}, c *Client) (*Note, error) {

	flattened := flattenNote(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandNote expands Note into a JSON request object.
func expandNote(c *Client, f *Note) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/notes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.ShortDescription; dcl.ValueShouldBeSent(v) {
		m["shortDescription"] = v
	}
	if v := f.LongDescription; dcl.ValueShouldBeSent(v) {
		m["longDescription"] = v
	}
	if v, err := expandNoteRelatedUrlSlice(c, f.RelatedUrl); err != nil {
		return nil, fmt.Errorf("error expanding RelatedUrl into relatedUrl: %w", err)
	} else {
		m["relatedUrl"] = v
	}
	if v := f.ExpirationTime; dcl.ValueShouldBeSent(v) {
		m["expirationTime"] = v
	}
	if v, err := expandNoteImage(c, f.Image); err != nil {
		return nil, fmt.Errorf("error expanding Image into image: %w", err)
	} else if v != nil {
		m["image"] = v
	}
	if v, err := expandNotePackage(c, f.Package); err != nil {
		return nil, fmt.Errorf("error expanding Package into package: %w", err)
	} else if v != nil {
		m["package"] = v
	}
	if v, err := expandNoteDiscovery(c, f.Discovery); err != nil {
		return nil, fmt.Errorf("error expanding Discovery into discovery: %w", err)
	} else if v != nil {
		m["discovery"] = v
	}
	if v, err := expandNoteDeployment(c, f.Deployment); err != nil {
		return nil, fmt.Errorf("error expanding Deployment into deployment: %w", err)
	} else if v != nil {
		m["deployment"] = v
	}
	if v, err := expandNoteAttestation(c, f.Attestation); err != nil {
		return nil, fmt.Errorf("error expanding Attestation into attestation: %w", err)
	} else if v != nil {
		m["attestation"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenNote flattens Note from a JSON request object into the
// Note type.
func flattenNote(c *Client, i interface{}) *Note {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Note{}
	res.Name = dcl.FlattenString(m["name"])
	res.ShortDescription = dcl.FlattenString(m["shortDescription"])
	res.LongDescription = dcl.FlattenString(m["longDescription"])
	res.RelatedUrl = flattenNoteRelatedUrlSlice(c, m["relatedUrl"])
	res.ExpirationTime = dcl.FlattenString(m["expirationTime"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Image = flattenNoteImage(c, m["image"])
	res.Package = flattenNotePackage(c, m["package"])
	res.Discovery = flattenNoteDiscovery(c, m["discovery"])
	res.Deployment = flattenNoteDeployment(c, m["deployment"])
	res.Attestation = flattenNoteAttestation(c, m["attestation"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandNoteRelatedUrlMap expands the contents of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrlMap(c *Client, f map[string]NoteRelatedUrl) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteRelatedUrl(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteRelatedUrlSlice expands the contents of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrlSlice(c *Client, f []NoteRelatedUrl) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteRelatedUrl(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteRelatedUrlMap flattens the contents of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrlMap(c *Client, i interface{}) map[string]NoteRelatedUrl {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteRelatedUrl{}
	}

	if len(a) == 0 {
		return map[string]NoteRelatedUrl{}
	}

	items := make(map[string]NoteRelatedUrl)
	for k, item := range a {
		items[k] = *flattenNoteRelatedUrl(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteRelatedUrlSlice flattens the contents of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrlSlice(c *Client, i interface{}) []NoteRelatedUrl {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteRelatedUrl{}
	}

	if len(a) == 0 {
		return []NoteRelatedUrl{}
	}

	items := make([]NoteRelatedUrl, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteRelatedUrl(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteRelatedUrl expands an instance of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrl(c *Client, f *NoteRelatedUrl) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v := f.Label; !dcl.IsEmptyValueIndirect(v) {
		m["label"] = v
	}

	return m, nil
}

// flattenNoteRelatedUrl flattens an instance of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrl(c *Client, i interface{}) *NoteRelatedUrl {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteRelatedUrl{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteRelatedUrl
	}
	r.Url = dcl.FlattenString(m["url"])
	r.Label = dcl.FlattenString(m["label"])

	return r
}

// expandNoteImageMap expands the contents of NoteImage into a JSON
// request object.
func expandNoteImageMap(c *Client, f map[string]NoteImage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteImage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteImageSlice expands the contents of NoteImage into a JSON
// request object.
func expandNoteImageSlice(c *Client, f []NoteImage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteImage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteImageMap flattens the contents of NoteImage from a JSON
// response object.
func flattenNoteImageMap(c *Client, i interface{}) map[string]NoteImage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteImage{}
	}

	if len(a) == 0 {
		return map[string]NoteImage{}
	}

	items := make(map[string]NoteImage)
	for k, item := range a {
		items[k] = *flattenNoteImage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteImageSlice flattens the contents of NoteImage from a JSON
// response object.
func flattenNoteImageSlice(c *Client, i interface{}) []NoteImage {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteImage{}
	}

	if len(a) == 0 {
		return []NoteImage{}
	}

	items := make([]NoteImage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteImage(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteImage expands an instance of NoteImage into a JSON
// request object.
func expandNoteImage(c *Client, f *NoteImage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceUrl; !dcl.IsEmptyValueIndirect(v) {
		m["resourceUrl"] = v
	}
	if v, err := expandNoteImageFingerprint(c, f.Fingerprint); err != nil {
		return nil, fmt.Errorf("error expanding Fingerprint into fingerprint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fingerprint"] = v
	}

	return m, nil
}

// flattenNoteImage flattens an instance of NoteImage from a JSON
// response object.
func flattenNoteImage(c *Client, i interface{}) *NoteImage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteImage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteImage
	}
	r.ResourceUrl = dcl.FlattenString(m["resourceUrl"])
	r.Fingerprint = flattenNoteImageFingerprint(c, m["fingerprint"])

	return r
}

// expandNoteImageFingerprintMap expands the contents of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprintMap(c *Client, f map[string]NoteImageFingerprint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteImageFingerprint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteImageFingerprintSlice expands the contents of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprintSlice(c *Client, f []NoteImageFingerprint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteImageFingerprint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteImageFingerprintMap flattens the contents of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprintMap(c *Client, i interface{}) map[string]NoteImageFingerprint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteImageFingerprint{}
	}

	if len(a) == 0 {
		return map[string]NoteImageFingerprint{}
	}

	items := make(map[string]NoteImageFingerprint)
	for k, item := range a {
		items[k] = *flattenNoteImageFingerprint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteImageFingerprintSlice flattens the contents of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprintSlice(c *Client, i interface{}) []NoteImageFingerprint {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteImageFingerprint{}
	}

	if len(a) == 0 {
		return []NoteImageFingerprint{}
	}

	items := make([]NoteImageFingerprint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteImageFingerprint(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteImageFingerprint expands an instance of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprint(c *Client, f *NoteImageFingerprint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.V1Name; !dcl.IsEmptyValueIndirect(v) {
		m["v1Name"] = v
	}
	if v := f.V2Blob; v != nil {
		m["v2Blob"] = v
	}

	return m, nil
}

// flattenNoteImageFingerprint flattens an instance of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprint(c *Client, i interface{}) *NoteImageFingerprint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteImageFingerprint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteImageFingerprint
	}
	r.V1Name = dcl.FlattenString(m["v1Name"])
	r.V2Blob = dcl.FlattenStringSlice(m["v2Blob"])
	r.V2Name = dcl.FlattenString(m["v2Name"])

	return r
}

// expandNotePackageMap expands the contents of NotePackage into a JSON
// request object.
func expandNotePackageMap(c *Client, f map[string]NotePackage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageSlice expands the contents of NotePackage into a JSON
// request object.
func expandNotePackageSlice(c *Client, f []NotePackage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageMap flattens the contents of NotePackage from a JSON
// response object.
func flattenNotePackageMap(c *Client, i interface{}) map[string]NotePackage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackage{}
	}

	if len(a) == 0 {
		return map[string]NotePackage{}
	}

	items := make(map[string]NotePackage)
	for k, item := range a {
		items[k] = *flattenNotePackage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNotePackageSlice flattens the contents of NotePackage from a JSON
// response object.
func flattenNotePackageSlice(c *Client, i interface{}) []NotePackage {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackage{}
	}

	if len(a) == 0 {
		return []NotePackage{}
	}

	items := make([]NotePackage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackage(c, item.(map[string]interface{})))
	}

	return items
}

// expandNotePackage expands an instance of NotePackage into a JSON
// request object.
func expandNotePackage(c *Client, f *NotePackage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandNotePackageDistributionSlice(c, f.Distribution); err != nil {
		return nil, fmt.Errorf("error expanding Distribution into distribution: %w", err)
	} else if v != nil {
		m["distribution"] = v
	}

	return m, nil
}

// flattenNotePackage flattens an instance of NotePackage from a JSON
// response object.
func flattenNotePackage(c *Client, i interface{}) *NotePackage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNotePackage
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Distribution = flattenNotePackageDistributionSlice(c, m["distribution"])

	return r
}

// expandNotePackageDistributionMap expands the contents of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistributionMap(c *Client, f map[string]NotePackageDistribution) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackageDistribution(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageDistributionSlice expands the contents of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistributionSlice(c *Client, f []NotePackageDistribution) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackageDistribution(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageDistributionMap flattens the contents of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistributionMap(c *Client, i interface{}) map[string]NotePackageDistribution {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistribution{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistribution{}
	}

	items := make(map[string]NotePackageDistribution)
	for k, item := range a {
		items[k] = *flattenNotePackageDistribution(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNotePackageDistributionSlice flattens the contents of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistributionSlice(c *Client, i interface{}) []NotePackageDistribution {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistribution{}
	}

	if len(a) == 0 {
		return []NotePackageDistribution{}
	}

	items := make([]NotePackageDistribution, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistribution(c, item.(map[string]interface{})))
	}

	return items
}

// expandNotePackageDistribution expands an instance of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistribution(c *Client, f *NotePackageDistribution) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["cpeUri"] = v
	}
	if v := f.Architecture; !dcl.IsEmptyValueIndirect(v) {
		m["architecture"] = v
	}
	if v, err := expandNotePackageDistributionLatestVersion(c, f.LatestVersion); err != nil {
		return nil, fmt.Errorf("error expanding LatestVersion into latestVersion: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["latestVersion"] = v
	}
	if v := f.Maintainer; !dcl.IsEmptyValueIndirect(v) {
		m["maintainer"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenNotePackageDistribution flattens an instance of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistribution(c *Client, i interface{}) *NotePackageDistribution {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackageDistribution{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNotePackageDistribution
	}
	r.CpeUri = dcl.FlattenString(m["cpeUri"])
	r.Architecture = flattenNotePackageDistributionArchitectureEnum(m["architecture"])
	r.LatestVersion = flattenNotePackageDistributionLatestVersion(c, m["latestVersion"])
	r.Maintainer = dcl.FlattenString(m["maintainer"])
	r.Url = dcl.FlattenString(m["url"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandNotePackageDistributionLatestVersionMap expands the contents of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersionMap(c *Client, f map[string]NotePackageDistributionLatestVersion) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackageDistributionLatestVersion(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageDistributionLatestVersionSlice expands the contents of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersionSlice(c *Client, f []NotePackageDistributionLatestVersion) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackageDistributionLatestVersion(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageDistributionLatestVersionMap flattens the contents of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionMap(c *Client, i interface{}) map[string]NotePackageDistributionLatestVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistributionLatestVersion{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistributionLatestVersion{}
	}

	items := make(map[string]NotePackageDistributionLatestVersion)
	for k, item := range a {
		items[k] = *flattenNotePackageDistributionLatestVersion(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNotePackageDistributionLatestVersionSlice flattens the contents of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionSlice(c *Client, i interface{}) []NotePackageDistributionLatestVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionLatestVersion{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionLatestVersion{}
	}

	items := make([]NotePackageDistributionLatestVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionLatestVersion(c, item.(map[string]interface{})))
	}

	return items
}

// expandNotePackageDistributionLatestVersion expands an instance of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersion(c *Client, f *NotePackageDistributionLatestVersion) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNotePackageDistributionLatestVersion flattens an instance of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersion(c *Client, i interface{}) *NotePackageDistributionLatestVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackageDistributionLatestVersion{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNotePackageDistributionLatestVersion
	}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNotePackageDistributionLatestVersionKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteDiscoveryMap expands the contents of NoteDiscovery into a JSON
// request object.
func expandNoteDiscoveryMap(c *Client, f map[string]NoteDiscovery) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteDiscovery(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteDiscoverySlice expands the contents of NoteDiscovery into a JSON
// request object.
func expandNoteDiscoverySlice(c *Client, f []NoteDiscovery) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteDiscovery(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteDiscoveryMap flattens the contents of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscoveryMap(c *Client, i interface{}) map[string]NoteDiscovery {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteDiscovery{}
	}

	if len(a) == 0 {
		return map[string]NoteDiscovery{}
	}

	items := make(map[string]NoteDiscovery)
	for k, item := range a {
		items[k] = *flattenNoteDiscovery(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteDiscoverySlice flattens the contents of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscoverySlice(c *Client, i interface{}) []NoteDiscovery {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDiscovery{}
	}

	if len(a) == 0 {
		return []NoteDiscovery{}
	}

	items := make([]NoteDiscovery, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDiscovery(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteDiscovery expands an instance of NoteDiscovery into a JSON
// request object.
func expandNoteDiscovery(c *Client, f *NoteDiscovery) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AnalysisKind; !dcl.IsEmptyValueIndirect(v) {
		m["analysisKind"] = v
	}

	return m, nil
}

// flattenNoteDiscovery flattens an instance of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscovery(c *Client, i interface{}) *NoteDiscovery {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteDiscovery{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteDiscovery
	}
	r.AnalysisKind = flattenNoteDiscoveryAnalysisKindEnum(m["analysisKind"])

	return r
}

// expandNoteDeploymentMap expands the contents of NoteDeployment into a JSON
// request object.
func expandNoteDeploymentMap(c *Client, f map[string]NoteDeployment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteDeployment(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteDeploymentSlice expands the contents of NoteDeployment into a JSON
// request object.
func expandNoteDeploymentSlice(c *Client, f []NoteDeployment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteDeployment(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteDeploymentMap flattens the contents of NoteDeployment from a JSON
// response object.
func flattenNoteDeploymentMap(c *Client, i interface{}) map[string]NoteDeployment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteDeployment{}
	}

	if len(a) == 0 {
		return map[string]NoteDeployment{}
	}

	items := make(map[string]NoteDeployment)
	for k, item := range a {
		items[k] = *flattenNoteDeployment(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteDeploymentSlice flattens the contents of NoteDeployment from a JSON
// response object.
func flattenNoteDeploymentSlice(c *Client, i interface{}) []NoteDeployment {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDeployment{}
	}

	if len(a) == 0 {
		return []NoteDeployment{}
	}

	items := make([]NoteDeployment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDeployment(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteDeployment expands an instance of NoteDeployment into a JSON
// request object.
func expandNoteDeployment(c *Client, f *NoteDeployment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceUri; v != nil {
		m["resourceUri"] = v
	}

	return m, nil
}

// flattenNoteDeployment flattens an instance of NoteDeployment from a JSON
// response object.
func flattenNoteDeployment(c *Client, i interface{}) *NoteDeployment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteDeployment{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteDeployment
	}
	r.ResourceUri = dcl.FlattenStringSlice(m["resourceUri"])

	return r
}

// expandNoteAttestationMap expands the contents of NoteAttestation into a JSON
// request object.
func expandNoteAttestationMap(c *Client, f map[string]NoteAttestation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteAttestation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteAttestationSlice expands the contents of NoteAttestation into a JSON
// request object.
func expandNoteAttestationSlice(c *Client, f []NoteAttestation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteAttestation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteAttestationMap flattens the contents of NoteAttestation from a JSON
// response object.
func flattenNoteAttestationMap(c *Client, i interface{}) map[string]NoteAttestation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteAttestation{}
	}

	if len(a) == 0 {
		return map[string]NoteAttestation{}
	}

	items := make(map[string]NoteAttestation)
	for k, item := range a {
		items[k] = *flattenNoteAttestation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteAttestationSlice flattens the contents of NoteAttestation from a JSON
// response object.
func flattenNoteAttestationSlice(c *Client, i interface{}) []NoteAttestation {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteAttestation{}
	}

	if len(a) == 0 {
		return []NoteAttestation{}
	}

	items := make([]NoteAttestation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteAttestation(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteAttestation expands an instance of NoteAttestation into a JSON
// request object.
func expandNoteAttestation(c *Client, f *NoteAttestation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandNoteAttestationHint(c, f.Hint); err != nil {
		return nil, fmt.Errorf("error expanding Hint into hint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hint"] = v
	}

	return m, nil
}

// flattenNoteAttestation flattens an instance of NoteAttestation from a JSON
// response object.
func flattenNoteAttestation(c *Client, i interface{}) *NoteAttestation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteAttestation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteAttestation
	}
	r.Hint = flattenNoteAttestationHint(c, m["hint"])

	return r
}

// expandNoteAttestationHintMap expands the contents of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHintMap(c *Client, f map[string]NoteAttestationHint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteAttestationHint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteAttestationHintSlice expands the contents of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHintSlice(c *Client, f []NoteAttestationHint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteAttestationHint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteAttestationHintMap flattens the contents of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHintMap(c *Client, i interface{}) map[string]NoteAttestationHint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteAttestationHint{}
	}

	if len(a) == 0 {
		return map[string]NoteAttestationHint{}
	}

	items := make(map[string]NoteAttestationHint)
	for k, item := range a {
		items[k] = *flattenNoteAttestationHint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteAttestationHintSlice flattens the contents of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHintSlice(c *Client, i interface{}) []NoteAttestationHint {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteAttestationHint{}
	}

	if len(a) == 0 {
		return []NoteAttestationHint{}
	}

	items := make([]NoteAttestationHint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteAttestationHint(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteAttestationHint expands an instance of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHint(c *Client, f *NoteAttestationHint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HumanReadableName; !dcl.IsEmptyValueIndirect(v) {
		m["humanReadableName"] = v
	}

	return m, nil
}

// flattenNoteAttestationHint flattens an instance of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHint(c *Client, i interface{}) *NoteAttestationHint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteAttestationHint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteAttestationHint
	}
	r.HumanReadableName = dcl.FlattenString(m["humanReadableName"])

	return r
}

// flattenNotePackageDistributionArchitectureEnumMap flattens the contents of NotePackageDistributionArchitectureEnum from a JSON
// response object.
func flattenNotePackageDistributionArchitectureEnumMap(c *Client, i interface{}) map[string]NotePackageDistributionArchitectureEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistributionArchitectureEnum{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistributionArchitectureEnum{}
	}

	items := make(map[string]NotePackageDistributionArchitectureEnum)
	for k, item := range a {
		items[k] = *flattenNotePackageDistributionArchitectureEnum(item.(interface{}))
	}

	return items
}

// flattenNotePackageDistributionArchitectureEnumSlice flattens the contents of NotePackageDistributionArchitectureEnum from a JSON
// response object.
func flattenNotePackageDistributionArchitectureEnumSlice(c *Client, i interface{}) []NotePackageDistributionArchitectureEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionArchitectureEnum{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionArchitectureEnum{}
	}

	items := make([]NotePackageDistributionArchitectureEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionArchitectureEnum(item.(interface{})))
	}

	return items
}

// flattenNotePackageDistributionArchitectureEnum asserts that an interface is a string, and returns a
// pointer to a *NotePackageDistributionArchitectureEnum with the same value as that string.
func flattenNotePackageDistributionArchitectureEnum(i interface{}) *NotePackageDistributionArchitectureEnum {
	s, ok := i.(string)
	if !ok {
		return NotePackageDistributionArchitectureEnumRef("")
	}

	return NotePackageDistributionArchitectureEnumRef(s)
}

// flattenNotePackageDistributionLatestVersionKindEnumMap flattens the contents of NotePackageDistributionLatestVersionKindEnum from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionKindEnumMap(c *Client, i interface{}) map[string]NotePackageDistributionLatestVersionKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistributionLatestVersionKindEnum{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistributionLatestVersionKindEnum{}
	}

	items := make(map[string]NotePackageDistributionLatestVersionKindEnum)
	for k, item := range a {
		items[k] = *flattenNotePackageDistributionLatestVersionKindEnum(item.(interface{}))
	}

	return items
}

// flattenNotePackageDistributionLatestVersionKindEnumSlice flattens the contents of NotePackageDistributionLatestVersionKindEnum from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionKindEnumSlice(c *Client, i interface{}) []NotePackageDistributionLatestVersionKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionLatestVersionKindEnum{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionLatestVersionKindEnum{}
	}

	items := make([]NotePackageDistributionLatestVersionKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionLatestVersionKindEnum(item.(interface{})))
	}

	return items
}

// flattenNotePackageDistributionLatestVersionKindEnum asserts that an interface is a string, and returns a
// pointer to a *NotePackageDistributionLatestVersionKindEnum with the same value as that string.
func flattenNotePackageDistributionLatestVersionKindEnum(i interface{}) *NotePackageDistributionLatestVersionKindEnum {
	s, ok := i.(string)
	if !ok {
		return NotePackageDistributionLatestVersionKindEnumRef("")
	}

	return NotePackageDistributionLatestVersionKindEnumRef(s)
}

// flattenNoteDiscoveryAnalysisKindEnumMap flattens the contents of NoteDiscoveryAnalysisKindEnum from a JSON
// response object.
func flattenNoteDiscoveryAnalysisKindEnumMap(c *Client, i interface{}) map[string]NoteDiscoveryAnalysisKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteDiscoveryAnalysisKindEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteDiscoveryAnalysisKindEnum{}
	}

	items := make(map[string]NoteDiscoveryAnalysisKindEnum)
	for k, item := range a {
		items[k] = *flattenNoteDiscoveryAnalysisKindEnum(item.(interface{}))
	}

	return items
}

// flattenNoteDiscoveryAnalysisKindEnumSlice flattens the contents of NoteDiscoveryAnalysisKindEnum from a JSON
// response object.
func flattenNoteDiscoveryAnalysisKindEnumSlice(c *Client, i interface{}) []NoteDiscoveryAnalysisKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDiscoveryAnalysisKindEnum{}
	}

	if len(a) == 0 {
		return []NoteDiscoveryAnalysisKindEnum{}
	}

	items := make([]NoteDiscoveryAnalysisKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDiscoveryAnalysisKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteDiscoveryAnalysisKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteDiscoveryAnalysisKindEnum with the same value as that string.
func flattenNoteDiscoveryAnalysisKindEnum(i interface{}) *NoteDiscoveryAnalysisKindEnum {
	s, ok := i.(string)
	if !ok {
		return NoteDiscoveryAnalysisKindEnumRef("")
	}

	return NoteDiscoveryAnalysisKindEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Note) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNote(b, c)
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

type noteDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         noteApiOperation
}

func convertFieldDiffsToNoteDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]noteDiff, error) {
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
	var diffs []noteDiff
	// For each operation name, create a noteDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := noteDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToNoteApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToNoteApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (noteApiOperation, error) {
	switch opName {

	case "updateNoteUpdateNoteOperation":
		return &updateNoteUpdateNoteOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractNoteFields(r *Note) error {
	vImage := r.Image
	if vImage == nil {
		// note: explicitly not the empty object.
		vImage = &NoteImage{}
	}
	if err := extractNoteImageFields(r, vImage); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vImage) {
		r.Image = vImage
	}
	vPackage := r.Package
	if vPackage == nil {
		// note: explicitly not the empty object.
		vPackage = &NotePackage{}
	}
	if err := extractNotePackageFields(r, vPackage); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPackage) {
		r.Package = vPackage
	}
	vDiscovery := r.Discovery
	if vDiscovery == nil {
		// note: explicitly not the empty object.
		vDiscovery = &NoteDiscovery{}
	}
	if err := extractNoteDiscoveryFields(r, vDiscovery); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDiscovery) {
		r.Discovery = vDiscovery
	}
	vDeployment := r.Deployment
	if vDeployment == nil {
		// note: explicitly not the empty object.
		vDeployment = &NoteDeployment{}
	}
	if err := extractNoteDeploymentFields(r, vDeployment); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDeployment) {
		r.Deployment = vDeployment
	}
	vAttestation := r.Attestation
	if vAttestation == nil {
		// note: explicitly not the empty object.
		vAttestation = &NoteAttestation{}
	}
	if err := extractNoteAttestationFields(r, vAttestation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAttestation) {
		r.Attestation = vAttestation
	}
	return nil
}
func extractNoteRelatedUrlFields(r *Note, o *NoteRelatedUrl) error {
	return nil
}
func extractNoteImageFields(r *Note, o *NoteImage) error {
	vFingerprint := o.Fingerprint
	if vFingerprint == nil {
		// note: explicitly not the empty object.
		vFingerprint = &NoteImageFingerprint{}
	}
	if err := extractNoteImageFingerprintFields(r, vFingerprint); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vFingerprint) {
		o.Fingerprint = vFingerprint
	}
	return nil
}
func extractNoteImageFingerprintFields(r *Note, o *NoteImageFingerprint) error {
	return nil
}
func extractNotePackageFields(r *Note, o *NotePackage) error {
	return nil
}
func extractNotePackageDistributionFields(r *Note, o *NotePackageDistribution) error {
	vLatestVersion := o.LatestVersion
	if vLatestVersion == nil {
		// note: explicitly not the empty object.
		vLatestVersion = &NotePackageDistributionLatestVersion{}
	}
	if err := extractNotePackageDistributionLatestVersionFields(r, vLatestVersion); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vLatestVersion) {
		o.LatestVersion = vLatestVersion
	}
	return nil
}
func extractNotePackageDistributionLatestVersionFields(r *Note, o *NotePackageDistributionLatestVersion) error {
	return nil
}
func extractNoteDiscoveryFields(r *Note, o *NoteDiscovery) error {
	return nil
}
func extractNoteDeploymentFields(r *Note, o *NoteDeployment) error {
	return nil
}
func extractNoteAttestationFields(r *Note, o *NoteAttestation) error {
	vHint := o.Hint
	if vHint == nil {
		// note: explicitly not the empty object.
		vHint = &NoteAttestationHint{}
	}
	if err := extractNoteAttestationHintFields(r, vHint); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vHint) {
		o.Hint = vHint
	}
	return nil
}
func extractNoteAttestationHintFields(r *Note, o *NoteAttestationHint) error {
	return nil
}

func postReadExtractNoteFields(r *Note) error {
	vImage := r.Image
	if vImage == nil {
		// note: explicitly not the empty object.
		vImage = &NoteImage{}
	}
	if err := postReadExtractNoteImageFields(r, vImage); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vImage) {
		r.Image = vImage
	}
	vPackage := r.Package
	if vPackage == nil {
		// note: explicitly not the empty object.
		vPackage = &NotePackage{}
	}
	if err := postReadExtractNotePackageFields(r, vPackage); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPackage) {
		r.Package = vPackage
	}
	vDiscovery := r.Discovery
	if vDiscovery == nil {
		// note: explicitly not the empty object.
		vDiscovery = &NoteDiscovery{}
	}
	if err := postReadExtractNoteDiscoveryFields(r, vDiscovery); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDiscovery) {
		r.Discovery = vDiscovery
	}
	vDeployment := r.Deployment
	if vDeployment == nil {
		// note: explicitly not the empty object.
		vDeployment = &NoteDeployment{}
	}
	if err := postReadExtractNoteDeploymentFields(r, vDeployment); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDeployment) {
		r.Deployment = vDeployment
	}
	vAttestation := r.Attestation
	if vAttestation == nil {
		// note: explicitly not the empty object.
		vAttestation = &NoteAttestation{}
	}
	if err := postReadExtractNoteAttestationFields(r, vAttestation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAttestation) {
		r.Attestation = vAttestation
	}
	return nil
}
func postReadExtractNoteRelatedUrlFields(r *Note, o *NoteRelatedUrl) error {
	return nil
}
func postReadExtractNoteImageFields(r *Note, o *NoteImage) error {
	vFingerprint := o.Fingerprint
	if vFingerprint == nil {
		// note: explicitly not the empty object.
		vFingerprint = &NoteImageFingerprint{}
	}
	if err := extractNoteImageFingerprintFields(r, vFingerprint); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vFingerprint) {
		o.Fingerprint = vFingerprint
	}
	return nil
}
func postReadExtractNoteImageFingerprintFields(r *Note, o *NoteImageFingerprint) error {
	return nil
}
func postReadExtractNotePackageFields(r *Note, o *NotePackage) error {
	return nil
}
func postReadExtractNotePackageDistributionFields(r *Note, o *NotePackageDistribution) error {
	vLatestVersion := o.LatestVersion
	if vLatestVersion == nil {
		// note: explicitly not the empty object.
		vLatestVersion = &NotePackageDistributionLatestVersion{}
	}
	if err := extractNotePackageDistributionLatestVersionFields(r, vLatestVersion); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vLatestVersion) {
		o.LatestVersion = vLatestVersion
	}
	return nil
}
func postReadExtractNotePackageDistributionLatestVersionFields(r *Note, o *NotePackageDistributionLatestVersion) error {
	return nil
}
func postReadExtractNoteDiscoveryFields(r *Note, o *NoteDiscovery) error {
	return nil
}
func postReadExtractNoteDeploymentFields(r *Note, o *NoteDeployment) error {
	return nil
}
func postReadExtractNoteAttestationFields(r *Note, o *NoteAttestation) error {
	vHint := o.Hint
	if vHint == nil {
		// note: explicitly not the empty object.
		vHint = &NoteAttestationHint{}
	}
	if err := extractNoteAttestationHintFields(r, vHint); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vHint) {
		o.Hint = vHint
	}
	return nil
}
func postReadExtractNoteAttestationHintFields(r *Note, o *NoteAttestationHint) error {
	return nil
}
