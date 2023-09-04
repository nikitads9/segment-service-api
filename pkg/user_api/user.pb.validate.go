// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package user_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ModifySegmentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ModifySegmentsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ModifySegmentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ModifySegmentsRequestMultiError, or nil if none found.
func (m *ModifySegmentsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ModifySegmentsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := ModifySegmentsRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_ModifySegmentsRequest_SlugsToAdd_Unique := make(map[string]struct{}, len(m.GetSlugsToAdd()))

	for idx, item := range m.GetSlugsToAdd() {
		_, _ = idx, item

		if _, exists := _ModifySegmentsRequest_SlugsToAdd_Unique[item]; exists {
			err := ModifySegmentsRequestValidationError{
				field:  fmt.Sprintf("SlugsToAdd[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_ModifySegmentsRequest_SlugsToAdd_Unique[item] = struct{}{}
		}

		if l := utf8.RuneCountInString(item); l < 1 || l > 100 {
			err := ModifySegmentsRequestValidationError{
				field:  fmt.Sprintf("SlugsToAdd[%v]", idx),
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	_ModifySegmentsRequest_SlugsToRemove_Unique := make(map[string]struct{}, len(m.GetSlugsToRemove()))

	for idx, item := range m.GetSlugsToRemove() {
		_, _ = idx, item

		if _, exists := _ModifySegmentsRequest_SlugsToRemove_Unique[item]; exists {
			err := ModifySegmentsRequestValidationError{
				field:  fmt.Sprintf("SlugsToRemove[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_ModifySegmentsRequest_SlugsToRemove_Unique[item] = struct{}{}
		}

		if l := utf8.RuneCountInString(item); l < 1 || l > 100 {
			err := ModifySegmentsRequestValidationError{
				field:  fmt.Sprintf("SlugsToRemove[%v]", idx),
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ModifySegmentsRequestMultiError(errors)
	}

	return nil
}

// ModifySegmentsRequestMultiError is an error wrapping multiple validation
// errors returned by ModifySegmentsRequest.ValidateAll() if the designated
// constraints aren't met.
type ModifySegmentsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ModifySegmentsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ModifySegmentsRequestMultiError) AllErrors() []error { return m }

// ModifySegmentsRequestValidationError is the validation error returned by
// ModifySegmentsRequest.Validate if the designated constraints aren't met.
type ModifySegmentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ModifySegmentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ModifySegmentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ModifySegmentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ModifySegmentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ModifySegmentsRequestValidationError) ErrorName() string {
	return "ModifySegmentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ModifySegmentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sModifySegmentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ModifySegmentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ModifySegmentsRequestValidationError{}

// Validate checks the field values on GetSegmentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSegmentsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSegmentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSegmentsRequestMultiError, or nil if none found.
func (m *GetSegmentsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSegmentsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetSegmentsRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetSegmentsRequestMultiError(errors)
	}

	return nil
}

// GetSegmentsRequestMultiError is an error wrapping multiple validation errors
// returned by GetSegmentsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetSegmentsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSegmentsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSegmentsRequestMultiError) AllErrors() []error { return m }

// GetSegmentsRequestValidationError is the validation error returned by
// GetSegmentsRequest.Validate if the designated constraints aren't met.
type GetSegmentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSegmentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSegmentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSegmentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSegmentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSegmentsRequestValidationError) ErrorName() string {
	return "GetSegmentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSegmentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSegmentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSegmentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSegmentsRequestValidationError{}

// Validate checks the field values on GetSegmentsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSegmentsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSegmentsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSegmentsResponseMultiError, or nil if none found.
func (m *GetSegmentsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSegmentsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetSegmentsResponseMultiError(errors)
	}

	return nil
}

// GetSegmentsResponseMultiError is an error wrapping multiple validation
// errors returned by GetSegmentsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetSegmentsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSegmentsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSegmentsResponseMultiError) AllErrors() []error { return m }

// GetSegmentsResponseValidationError is the validation error returned by
// GetSegmentsResponse.Validate if the designated constraints aren't met.
type GetSegmentsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSegmentsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSegmentsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSegmentsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSegmentsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSegmentsResponseValidationError) ErrorName() string {
	return "GetSegmentsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSegmentsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSegmentsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSegmentsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSegmentsResponseValidationError{}

// Validate checks the field values on SetExpireTimeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SetExpireTimeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetExpireTimeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetExpireTimeRequestMultiError, or nil if none found.
func (m *SetExpireTimeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SetExpireTimeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := SetExpireTimeRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetSlug()); l < 1 || l > 100 {
		err := SetExpireTimeRequestValidationError{
			field:  "Slug",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetExpirationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SetExpireTimeRequestValidationError{
					field:  "ExpirationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SetExpireTimeRequestValidationError{
					field:  "ExpirationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpirationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetExpireTimeRequestValidationError{
				field:  "ExpirationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SetExpireTimeRequestMultiError(errors)
	}

	return nil
}

// SetExpireTimeRequestMultiError is an error wrapping multiple validation
// errors returned by SetExpireTimeRequest.ValidateAll() if the designated
// constraints aren't met.
type SetExpireTimeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetExpireTimeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetExpireTimeRequestMultiError) AllErrors() []error { return m }

// SetExpireTimeRequestValidationError is the validation error returned by
// SetExpireTimeRequest.Validate if the designated constraints aren't met.
type SetExpireTimeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetExpireTimeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetExpireTimeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetExpireTimeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetExpireTimeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetExpireTimeRequestValidationError) ErrorName() string {
	return "SetExpireTimeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetExpireTimeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetExpireTimeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetExpireTimeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetExpireTimeRequestValidationError{}

// Validate checks the field values on AddUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddUserRequestMultiError,
// or nil if none found.
func (m *AddUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUserName()); l < 1 || l > 100 {
		err := AddUserRequestValidationError{
			field:  "UserName",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddUserRequestMultiError(errors)
	}

	return nil
}

// AddUserRequestMultiError is an error wrapping multiple validation errors
// returned by AddUserRequest.ValidateAll() if the designated constraints
// aren't met.
type AddUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddUserRequestMultiError) AllErrors() []error { return m }

// AddUserRequestValidationError is the validation error returned by
// AddUserRequest.Validate if the designated constraints aren't met.
type AddUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddUserRequestValidationError) ErrorName() string { return "AddUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddUserRequestValidationError{}

// Validate checks the field values on AddUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddUserResponseMultiError, or nil if none found.
func (m *AddUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddUserResponseMultiError(errors)
	}

	return nil
}

// AddUserResponseMultiError is an error wrapping multiple validation errors
// returned by AddUserResponse.ValidateAll() if the designated constraints
// aren't met.
type AddUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddUserResponseMultiError) AllErrors() []error { return m }

// AddUserResponseValidationError is the validation error returned by
// AddUserResponse.Validate if the designated constraints aren't met.
type AddUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddUserResponseValidationError) ErrorName() string { return "AddUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddUserResponseValidationError{}

// Validate checks the field values on RemoveUsertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveUsertRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveUsertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveUsertRequestMultiError, or nil if none found.
func (m *RemoveUsertRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveUsertRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := RemoveUsertRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RemoveUsertRequestMultiError(errors)
	}

	return nil
}

// RemoveUsertRequestMultiError is an error wrapping multiple validation errors
// returned by RemoveUsertRequest.ValidateAll() if the designated constraints
// aren't met.
type RemoveUsertRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveUsertRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveUsertRequestMultiError) AllErrors() []error { return m }

// RemoveUsertRequestValidationError is the validation error returned by
// RemoveUsertRequest.Validate if the designated constraints aren't met.
type RemoveUsertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveUsertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveUsertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveUsertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveUsertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveUsertRequestValidationError) ErrorName() string {
	return "RemoveUsertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveUsertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveUsertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveUsertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveUsertRequestValidationError{}

// Validate checks the field values on GetUserHistoryCsvRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserHistoryCsvRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserHistoryCsvRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserHistoryCsvRequestMultiError, or nil if none found.
func (m *GetUserHistoryCsvRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserHistoryCsvRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetUserHistoryCsvRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetUserHistoryCsvRequestMultiError(errors)
	}

	return nil
}

// GetUserHistoryCsvRequestMultiError is an error wrapping multiple validation
// errors returned by GetUserHistoryCsvRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserHistoryCsvRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserHistoryCsvRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserHistoryCsvRequestMultiError) AllErrors() []error { return m }

// GetUserHistoryCsvRequestValidationError is the validation error returned by
// GetUserHistoryCsvRequest.Validate if the designated constraints aren't met.
type GetUserHistoryCsvRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserHistoryCsvRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserHistoryCsvRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserHistoryCsvRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserHistoryCsvRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserHistoryCsvRequestValidationError) ErrorName() string {
	return "GetUserHistoryCsvRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserHistoryCsvRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserHistoryCsvRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserHistoryCsvRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserHistoryCsvRequestValidationError{}

// Validate checks the field values on GetUserHistoryCsvResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserHistoryCsvResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserHistoryCsvResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserHistoryCsvResponseMultiError, or nil if none found.
func (m *GetUserHistoryCsvResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserHistoryCsvResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Chunk

	if len(errors) > 0 {
		return GetUserHistoryCsvResponseMultiError(errors)
	}

	return nil
}

// GetUserHistoryCsvResponseMultiError is an error wrapping multiple validation
// errors returned by GetUserHistoryCsvResponse.ValidateAll() if the
// designated constraints aren't met.
type GetUserHistoryCsvResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserHistoryCsvResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserHistoryCsvResponseMultiError) AllErrors() []error { return m }

// GetUserHistoryCsvResponseValidationError is the validation error returned by
// GetUserHistoryCsvResponse.Validate if the designated constraints aren't met.
type GetUserHistoryCsvResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserHistoryCsvResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserHistoryCsvResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserHistoryCsvResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserHistoryCsvResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserHistoryCsvResponseValidationError) ErrorName() string {
	return "GetUserHistoryCsvResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserHistoryCsvResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserHistoryCsvResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserHistoryCsvResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserHistoryCsvResponseValidationError{}
