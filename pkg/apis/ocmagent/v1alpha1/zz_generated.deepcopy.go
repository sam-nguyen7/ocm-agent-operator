// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCMAgent) DeepCopyInto(out *OCMAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCMAgent.
func (in *OCMAgent) DeepCopy() *OCMAgent {
	if in == nil {
		return nil
	}
	out := new(OCMAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCMAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCMAgentList) DeepCopyInto(out *OCMAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCMAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCMAgentList.
func (in *OCMAgentList) DeepCopy() *OCMAgentList {
	if in == nil {
		return nil
	}
	out := new(OCMAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCMAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCMAgentSpec) DeepCopyInto(out *OCMAgentSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCMAgentSpec.
func (in *OCMAgentSpec) DeepCopy() *OCMAgentSpec {
	if in == nil {
		return nil
	}
	out := new(OCMAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCMAgentStatus) DeepCopyInto(out *OCMAgentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCMAgentStatus.
func (in *OCMAgentStatus) DeepCopy() *OCMAgentStatus {
	if in == nil {
		return nil
	}
	out := new(OCMAgentStatus)
	in.DeepCopyInto(out)
	return out
}