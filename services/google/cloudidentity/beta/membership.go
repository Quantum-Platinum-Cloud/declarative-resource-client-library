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
package beta

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Membership struct {
	Name               *string                        `json:"name"`
	PreferredMemberKey *MembershipPreferredMemberKey  `json:"preferredMemberKey"`
	CreateTime         *string                        `json:"createTime"`
	UpdateTime         *string                        `json:"updateTime"`
	Roles              []MembershipRoles              `json:"roles"`
	Type               *MembershipTypeEnum            `json:"type"`
	DeliverySetting    *MembershipDeliverySettingEnum `json:"deliverySetting"`
	DisplayName        *MembershipDisplayName         `json:"displayName"`
	MemberKey          *MembershipMemberKey           `json:"memberKey"`
	Group              *string                        `json:"group"`
}

func (r *Membership) String() string {
	return dcl.SprintResource(r)
}

// The enum MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum.
type MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum string

// MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumRef returns a *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumRef(s string) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	v := MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(s)
	return &v
}

func (v MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ENCRYPTION_STATE_UNSPECIFIED", "UNSUPPORTED_BY_DEVICE", "ENCRYPTED", "NOT_ENCRYPTED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum MembershipTypeEnum.
type MembershipTypeEnum string

// MembershipTypeEnumRef returns a *MembershipTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func MembershipTypeEnumRef(s string) *MembershipTypeEnum {
	v := MembershipTypeEnum(s)
	return &v
}

func (v MembershipTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"OWNER_TYPE_UNSPECIFIED", "OWNER_TYPE_CUSTOMER", "OWNER_TYPE_PARTNER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MembershipTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum MembershipDeliverySettingEnum.
type MembershipDeliverySettingEnum string

// MembershipDeliverySettingEnumRef returns a *MembershipDeliverySettingEnum with the value of string s
// If the empty string is provided, nil is returned.
func MembershipDeliverySettingEnumRef(s string) *MembershipDeliverySettingEnum {
	v := MembershipDeliverySettingEnum(s)
	return &v
}

func (v MembershipDeliverySettingEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"DELIVERY_SETTING_UNSPECIFIED", "ALL_MAIL", "DIGEST", "DAILY", "NONE", "DISABLED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MembershipDeliverySettingEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type MembershipPreferredMemberKey struct {
	empty     bool    `json:"-"`
	Id        *string `json:"id"`
	Namespace *string `json:"namespace"`
}

type jsonMembershipPreferredMemberKey MembershipPreferredMemberKey

func (r *MembershipPreferredMemberKey) UnmarshalJSON(data []byte) error {
	var res jsonMembershipPreferredMemberKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipPreferredMemberKey
	} else {

		r.Id = res.Id

		r.Namespace = res.Namespace

	}
	return nil
}

// This object is used to assert a desired state where this MembershipPreferredMemberKey is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipPreferredMemberKey *MembershipPreferredMemberKey = &MembershipPreferredMemberKey{empty: true}

func (r *MembershipPreferredMemberKey) Empty() bool {
	return r.empty
}

func (r *MembershipPreferredMemberKey) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipPreferredMemberKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipRoles struct {
	empty                  bool                                   `json:"-"`
	Name                   *string                                `json:"name"`
	ExpiryDetail           *MembershipRolesExpiryDetail           `json:"expiryDetail"`
	RestrictionEvaluations *MembershipRolesRestrictionEvaluations `json:"restrictionEvaluations"`
}

type jsonMembershipRoles MembershipRoles

func (r *MembershipRoles) UnmarshalJSON(data []byte) error {
	var res jsonMembershipRoles
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipRoles
	} else {

		r.Name = res.Name

		r.ExpiryDetail = res.ExpiryDetail

		r.RestrictionEvaluations = res.RestrictionEvaluations

	}
	return nil
}

// This object is used to assert a desired state where this MembershipRoles is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipRoles *MembershipRoles = &MembershipRoles{empty: true}

func (r *MembershipRoles) Empty() bool {
	return r.empty
}

func (r *MembershipRoles) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipRoles) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipRolesExpiryDetail struct {
	empty      bool    `json:"-"`
	ExpireTime *string `json:"expireTime"`
}

type jsonMembershipRolesExpiryDetail MembershipRolesExpiryDetail

func (r *MembershipRolesExpiryDetail) UnmarshalJSON(data []byte) error {
	var res jsonMembershipRolesExpiryDetail
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipRolesExpiryDetail
	} else {

		r.ExpireTime = res.ExpireTime

	}
	return nil
}

// This object is used to assert a desired state where this MembershipRolesExpiryDetail is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipRolesExpiryDetail *MembershipRolesExpiryDetail = &MembershipRolesExpiryDetail{empty: true}

func (r *MembershipRolesExpiryDetail) Empty() bool {
	return r.empty
}

func (r *MembershipRolesExpiryDetail) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipRolesExpiryDetail) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipRolesRestrictionEvaluations struct {
	empty                       bool                                                              `json:"-"`
	MemberRestrictionEvaluation *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation `json:"memberRestrictionEvaluation"`
}

type jsonMembershipRolesRestrictionEvaluations MembershipRolesRestrictionEvaluations

func (r *MembershipRolesRestrictionEvaluations) UnmarshalJSON(data []byte) error {
	var res jsonMembershipRolesRestrictionEvaluations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipRolesRestrictionEvaluations
	} else {

		r.MemberRestrictionEvaluation = res.MemberRestrictionEvaluation

	}
	return nil
}

// This object is used to assert a desired state where this MembershipRolesRestrictionEvaluations is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipRolesRestrictionEvaluations *MembershipRolesRestrictionEvaluations = &MembershipRolesRestrictionEvaluations{empty: true}

func (r *MembershipRolesRestrictionEvaluations) Empty() bool {
	return r.empty
}

func (r *MembershipRolesRestrictionEvaluations) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipRolesRestrictionEvaluations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation struct {
	empty bool                                                                       `json:"-"`
	State *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum `json:"state"`
}

type jsonMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation

func (r *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) UnmarshalJSON(data []byte) error {
	var res jsonMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation
	} else {

		r.State = res.State

	}
	return nil
}

// This object is used to assert a desired state where this MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation = &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{empty: true}

func (r *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) Empty() bool {
	return r.empty
}

func (r *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipDisplayName struct {
	empty      bool    `json:"-"`
	GivenName  *string `json:"givenName"`
	FamilyName *string `json:"familyName"`
	FullName   *string `json:"fullName"`
}

type jsonMembershipDisplayName MembershipDisplayName

func (r *MembershipDisplayName) UnmarshalJSON(data []byte) error {
	var res jsonMembershipDisplayName
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipDisplayName
	} else {

		r.GivenName = res.GivenName

		r.FamilyName = res.FamilyName

		r.FullName = res.FullName

	}
	return nil
}

// This object is used to assert a desired state where this MembershipDisplayName is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipDisplayName *MembershipDisplayName = &MembershipDisplayName{empty: true}

func (r *MembershipDisplayName) Empty() bool {
	return r.empty
}

func (r *MembershipDisplayName) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipDisplayName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipMemberKey struct {
	empty     bool    `json:"-"`
	Id        *string `json:"id"`
	Namespace *string `json:"namespace"`
}

type jsonMembershipMemberKey MembershipMemberKey

func (r *MembershipMemberKey) UnmarshalJSON(data []byte) error {
	var res jsonMembershipMemberKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipMemberKey
	} else {

		r.Id = res.Id

		r.Namespace = res.Namespace

	}
	return nil
}

// This object is used to assert a desired state where this MembershipMemberKey is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipMemberKey *MembershipMemberKey = &MembershipMemberKey{empty: true}

func (r *MembershipMemberKey) Empty() bool {
	return r.empty
}

func (r *MembershipMemberKey) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipMemberKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Membership) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloudidentity",
		Type:    "Membership",
		Version: "beta",
	}
}

func (r *Membership) ID() (string, error) {
	if err := extractMembershipFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":               dcl.ValueOrEmptyString(nr.Name),
		"preferredMemberKey": dcl.ValueOrEmptyString(nr.PreferredMemberKey),
		"createTime":         dcl.ValueOrEmptyString(nr.CreateTime),
		"updateTime":         dcl.ValueOrEmptyString(nr.UpdateTime),
		"roles":              dcl.ValueOrEmptyString(nr.Roles),
		"type":               dcl.ValueOrEmptyString(nr.Type),
		"deliverySetting":    dcl.ValueOrEmptyString(nr.DeliverySetting),
		"displayName":        dcl.ValueOrEmptyString(nr.DisplayName),
		"memberKey":          dcl.ValueOrEmptyString(nr.MemberKey),
		"group":              dcl.ValueOrEmptyString(nr.Group),
	}
	return dcl.Nprintf("groups/{{group}}/memberships/{{name}}", params), nil
}

const MembershipMaxPage = -1

type MembershipList struct {
	Items []*Membership

	nextToken string

	pageSize int32

	resource *Membership
}

func (l *MembershipList) HasNext() bool {
	return l.nextToken != ""
}

func (l *MembershipList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listMembership(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListMembership(ctx context.Context, group string) (*MembershipList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListMembershipWithMaxResults(ctx, group, MembershipMaxPage)

}

func (c *Client) ListMembershipWithMaxResults(ctx context.Context, group string, pageSize int32) (*MembershipList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Membership{
		Group: &group,
	}
	items, token, err := c.listMembership(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &MembershipList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetMembership(ctx context.Context, r *Membership) (*Membership, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractMembershipFields(r)

	b, err := c.getMembershipRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalMembership(b, c, r)
	if err != nil {
		return nil, err
	}
	nr := r.urlNormalized()
	result.Group = nr.Group
	result.Name = nr.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeMembershipNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractMembershipFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteMembership(ctx context.Context, r *Membership) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Membership resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Membership...")
	deleteOp := deleteMembershipOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllMembership deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllMembership(ctx context.Context, group string, filter func(*Membership) bool) error {
	listObj, err := c.ListMembership(ctx, group)
	if err != nil {
		return err
	}

	err = c.deleteAllMembership(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllMembership(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyMembership(ctx context.Context, rawDesired *Membership, opts ...dcl.ApplyOption) (*Membership, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Membership
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyMembershipHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyMembershipHelper(c *Client, ctx context.Context, rawDesired *Membership, opts ...dcl.ApplyOption) (*Membership, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyMembership...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractMembershipFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.membershipDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToMembershipDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []membershipApiOperation
	if create {
		ops = append(ops, &createMembershipOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyMembershipDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyMembershipDiff(c *Client, ctx context.Context, desired *Membership, rawDesired *Membership, ops []membershipApiOperation, opts ...dcl.ApplyOption) (*Membership, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetMembership(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createMembershipOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapMembership(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeMembershipNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeMembershipNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeMembershipDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractMembershipFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractMembershipFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffMembership(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
