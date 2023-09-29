// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package meshv2beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using GRPCRoute within kubernetes types, where deepcopy-gen is used.
func (in *GRPCRoute) DeepCopyInto(out *GRPCRoute) {
	p := proto.Clone(in).(*GRPCRoute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRoute. Required by controller-gen.
func (in *GRPCRoute) DeepCopy() *GRPCRoute {
	if in == nil {
		return nil
	}
	out := new(GRPCRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRoute. Required by controller-gen.
func (in *GRPCRoute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GRPCRouteRule within kubernetes types, where deepcopy-gen is used.
func (in *GRPCRouteRule) DeepCopyInto(out *GRPCRouteRule) {
	p := proto.Clone(in).(*GRPCRouteRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRouteRule. Required by controller-gen.
func (in *GRPCRouteRule) DeepCopy() *GRPCRouteRule {
	if in == nil {
		return nil
	}
	out := new(GRPCRouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRouteRule. Required by controller-gen.
func (in *GRPCRouteRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GRPCRouteMatch within kubernetes types, where deepcopy-gen is used.
func (in *GRPCRouteMatch) DeepCopyInto(out *GRPCRouteMatch) {
	p := proto.Clone(in).(*GRPCRouteMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRouteMatch. Required by controller-gen.
func (in *GRPCRouteMatch) DeepCopy() *GRPCRouteMatch {
	if in == nil {
		return nil
	}
	out := new(GRPCRouteMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRouteMatch. Required by controller-gen.
func (in *GRPCRouteMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GRPCMethodMatch within kubernetes types, where deepcopy-gen is used.
func (in *GRPCMethodMatch) DeepCopyInto(out *GRPCMethodMatch) {
	p := proto.Clone(in).(*GRPCMethodMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCMethodMatch. Required by controller-gen.
func (in *GRPCMethodMatch) DeepCopy() *GRPCMethodMatch {
	if in == nil {
		return nil
	}
	out := new(GRPCMethodMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GRPCMethodMatch. Required by controller-gen.
func (in *GRPCMethodMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GRPCHeaderMatch within kubernetes types, where deepcopy-gen is used.
func (in *GRPCHeaderMatch) DeepCopyInto(out *GRPCHeaderMatch) {
	p := proto.Clone(in).(*GRPCHeaderMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCHeaderMatch. Required by controller-gen.
func (in *GRPCHeaderMatch) DeepCopy() *GRPCHeaderMatch {
	if in == nil {
		return nil
	}
	out := new(GRPCHeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GRPCHeaderMatch. Required by controller-gen.
func (in *GRPCHeaderMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GRPCRouteFilter within kubernetes types, where deepcopy-gen is used.
func (in *GRPCRouteFilter) DeepCopyInto(out *GRPCRouteFilter) {
	p := proto.Clone(in).(*GRPCRouteFilter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRouteFilter. Required by controller-gen.
func (in *GRPCRouteFilter) DeepCopy() *GRPCRouteFilter {
	if in == nil {
		return nil
	}
	out := new(GRPCRouteFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRouteFilter. Required by controller-gen.
func (in *GRPCRouteFilter) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GRPCBackendRef within kubernetes types, where deepcopy-gen is used.
func (in *GRPCBackendRef) DeepCopyInto(out *GRPCBackendRef) {
	p := proto.Clone(in).(*GRPCBackendRef)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCBackendRef. Required by controller-gen.
func (in *GRPCBackendRef) DeepCopy() *GRPCBackendRef {
	if in == nil {
		return nil
	}
	out := new(GRPCBackendRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GRPCBackendRef. Required by controller-gen.
func (in *GRPCBackendRef) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
