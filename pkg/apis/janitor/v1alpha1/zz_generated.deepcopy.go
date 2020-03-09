// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sweep) DeepCopyInto(out *Sweep) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sweep.
func (in *Sweep) DeepCopy() *Sweep {
	if in == nil {
		return nil
	}
	out := new(Sweep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sweep) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SweepList) DeepCopyInto(out *SweepList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sweep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SweepList.
func (in *SweepList) DeepCopy() *SweepList {
	if in == nil {
		return nil
	}
	out := new(SweepList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SweepList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SweepSpec) DeepCopyInto(out *SweepSpec) {
	*out = *in
	if in.IgnoreProjects != nil {
		in, out := &in.IgnoreProjects, &out.IgnoreProjects
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnoreAnnotations != nil {
		in, out := &in.IgnoreAnnotations, &out.IgnoreAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SweepSpec.
func (in *SweepSpec) DeepCopy() *SweepSpec {
	if in == nil {
		return nil
	}
	out := new(SweepSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SweepStatus) DeepCopyInto(out *SweepStatus) {
	*out = *in
	if in.Started != nil {
		in, out := &in.Started, &out.Started
		*out = (*in).DeepCopy()
	}
	if in.Finished != nil {
		in, out := &in.Finished, &out.Finished
		*out = (*in).DeepCopy()
	}
	if in.ProjectsDeleted != nil {
		in, out := &in.ProjectsDeleted, &out.ProjectsDeleted
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SweepStatus.
func (in *SweepStatus) DeepCopy() *SweepStatus {
	if in == nil {
		return nil
	}
	out := new(SweepStatus)
	in.DeepCopyInto(out)
	return out
}
