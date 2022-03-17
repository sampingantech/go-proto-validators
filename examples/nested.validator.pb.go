// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/nested.proto

package validator_examples

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/sampingantech/go-proto-validators"
	regexp "regexp"
	github_com_sampingantech_go_proto_validators "github.com/sampingantech/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *InnerMessage) Validate() *github_com_sampingantech_go_proto_validators.MultiError {
	multiError := github_com_sampingantech_go_proto_validators.NewMultiError()
	if !(this.SomeInteger > 0) {
		multiError.Append("SomeInteger", github_com_sampingantech_go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be greater than '0'`, this.SomeInteger)))
	}
	if !(this.SomeInteger < 100) {
		multiError.Append("SomeInteger", github_com_sampingantech_go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be less than '100'`, this.SomeInteger)))
	}
	if !multiError.HasError() {
		return nil
	}
	return multiError
}

var _regex_OuterMessage_ImportantString = regexp.MustCompile(`^[a-z]{2,5}$`)

func (this *OuterMessage) Validate() *github_com_sampingantech_go_proto_validators.MultiError {
	multiError := github_com_sampingantech_go_proto_validators.NewMultiError()
	if !_regex_OuterMessage_ImportantString.MatchString(this.ImportantString) {
		multiError.Append("ImportantString", github_com_sampingantech_go_proto_validators.FieldError("ImportantString", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z]{2,5}$"`, this.ImportantString)))
	}
	if nil == this.Inner {
		multiError.Append("Inner", github_com_sampingantech_go_proto_validators.FieldError("Inner", fmt.Errorf("message must exist")))
	}
	if this.Inner != nil {
		if err := github_com_sampingantech_go_proto_validators.CallValidatorIfExists(this.Inner); err != nil && err.HasError() {
			getErr := err.ToMap()
			multiError.Append("Inner", github_com_sampingantech_go_proto_validators.FieldError("Inner", fmt.Errorf(getErr["Inner"])))
		}
	}
	if !multiError.HasError() {
		return nil
	}
	return multiError
}
