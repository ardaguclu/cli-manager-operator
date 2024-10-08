// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// RouteHTTPHeaderApplyConfiguration represents a declarative configuration of the RouteHTTPHeader type for use
// with apply.
type RouteHTTPHeaderApplyConfiguration struct {
	Name   *string                                       `json:"name,omitempty"`
	Action *RouteHTTPHeaderActionUnionApplyConfiguration `json:"action,omitempty"`
}

// RouteHTTPHeaderApplyConfiguration constructs a declarative configuration of the RouteHTTPHeader type for use with
// apply.
func RouteHTTPHeader() *RouteHTTPHeaderApplyConfiguration {
	return &RouteHTTPHeaderApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *RouteHTTPHeaderApplyConfiguration) WithName(value string) *RouteHTTPHeaderApplyConfiguration {
	b.Name = &value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *RouteHTTPHeaderApplyConfiguration) WithAction(value *RouteHTTPHeaderActionUnionApplyConfiguration) *RouteHTTPHeaderApplyConfiguration {
	b.Action = value
	return b
}
