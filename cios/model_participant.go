/*
 * Cios Openapi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cios

import (
	"encoding/json"
)

// Participant struct for Participant
type Participant struct {
	ParticipantId *string `json:"participant_id,omitempty"`
	Token *string `json:"token,omitempty"`
	Expires *int32 `json:"expires,omitempty"`
	Target *int32 `json:"target,omitempty"`
	Used *int32 `json:"used,omitempty"`
}

// NewParticipant instantiates a new Participant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipant() *Participant {
	this := Participant{}
	return &this
}

// NewParticipantWithDefaults instantiates a new Participant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantWithDefaults() *Participant {
	this := Participant{}
	return &this
}

// GetParticipantId returns the ParticipantId field value if set, zero value otherwise.
func (o *Participant) GetParticipantId() string {
	if o == nil || o.ParticipantId == nil {
		var ret string
		return ret
	}
	return *o.ParticipantId
}

// GetParticipantIdOk returns a tuple with the ParticipantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetParticipantIdOk() (*string, bool) {
	if o == nil || o.ParticipantId == nil {
		return nil, false
	}
	return o.ParticipantId, true
}

// HasParticipantId returns a boolean if a field has been set.
func (o *Participant) HasParticipantId() bool {
	if o != nil && o.ParticipantId != nil {
		return true
	}

	return false
}

// SetParticipantId gets a reference to the given string and assigns it to the ParticipantId field.
func (o *Participant) SetParticipantId(v string) {
	o.ParticipantId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Participant) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Participant) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Participant) SetToken(v string) {
	o.Token = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *Participant) GetExpires() int32 {
	if o == nil || o.Expires == nil {
		var ret int32
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetExpiresOk() (*int32, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *Participant) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given int32 and assigns it to the Expires field.
func (o *Participant) SetExpires(v int32) {
	o.Expires = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Participant) GetTarget() int32 {
	if o == nil || o.Target == nil {
		var ret int32
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetTargetOk() (*int32, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Participant) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given int32 and assigns it to the Target field.
func (o *Participant) SetTarget(v int32) {
	o.Target = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *Participant) GetUsed() int32 {
	if o == nil || o.Used == nil {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetUsedOk() (*int32, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *Participant) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *Participant) SetUsed(v int32) {
	o.Used = &v
}

func (o Participant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParticipantId != nil {
		toSerialize["participant_id"] = o.ParticipantId
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableParticipant struct {
	value *Participant
	isSet bool
}

func (v NullableParticipant) Get() *Participant {
	return v.value
}

func (v *NullableParticipant) Set(val *Participant) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipant) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipant(val *Participant) *NullableParticipant {
	return &NullableParticipant{value: val, isSet: true}
}

func (v NullableParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


