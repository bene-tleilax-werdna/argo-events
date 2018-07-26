// +build !ignore_autogenerated

/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactLocation) DeepCopyInto(out *ArtifactLocation) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3Artifact)
		(*in).DeepCopyInto(*out)
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileArtifact)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(URLArtifact)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactLocation.
func (in *ArtifactLocation) DeepCopy() *ArtifactLocation {
	if in == nil {
		return nil
	}
	out := new(ArtifactLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactSignal) DeepCopyInto(out *ArtifactSignal) {
	*out = *in
	in.ArtifactLocation.DeepCopyInto(&out.ArtifactLocation)
	in.Target.DeepCopyInto(&out.Target)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactSignal.
func (in *ArtifactSignal) DeepCopy() *ArtifactSignal {
	if in == nil {
		return nil
	}
	out := new(ArtifactSignal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalendarSignal) DeepCopyInto(out *CalendarSignal) {
	*out = *in
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalendarSignal.
func (in *CalendarSignal) DeepCopy() *CalendarSignal {
	if in == nil {
		return nil
	}
	out := new(CalendarSignal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataFilter) DeepCopyInto(out *DataFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataFilter.
func (in *DataFilter) DeepCopy() *DataFilter {
	if in == nil {
		return nil
	}
	out := new(DataFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EscalationPolicy) DeepCopyInto(out *EscalationPolicy) {
	*out = *in
	in.Message.DeepCopyInto(&out.Message)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EscalationPolicy.
func (in *EscalationPolicy) DeepCopy() *EscalationPolicy {
	if in == nil {
		return nil
	}
	out := new(EscalationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Event) DeepCopyInto(out *Event) {
	*out = *in
	in.Context.DeepCopyInto(&out.Context)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Event.
func (in *Event) DeepCopy() *Event {
	if in == nil {
		return nil
	}
	out := new(Event)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventContext) DeepCopyInto(out *EventContext) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(URI)
		**out = **in
	}
	in.EventTime.DeepCopyInto(&out.EventTime)
	if in.SchemaURL != nil {
		in, out := &in.SchemaURL, &out.SchemaURL
		*out = new(URI)
		**out = **in
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventContext.
func (in *EventContext) DeepCopy() *EventContext {
	if in == nil {
		return nil
	}
	out := new(EventContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventWrapper) DeepCopyInto(out *EventWrapper) {
	*out = *in
	in.Event.DeepCopyInto(&out.Event)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventWrapper.
func (in *EventWrapper) DeepCopy() *EventWrapper {
	if in == nil {
		return nil
	}
	out := new(EventWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileArtifact) DeepCopyInto(out *FileArtifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileArtifact.
func (in *FileArtifact) DeepCopy() *FileArtifact {
	if in == nil {
		return nil
	}
	out := new(FileArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionKind) DeepCopyInto(out *GroupVersionKind) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionKind.
func (in *GroupVersionKind) DeepCopy() *GroupVersionKind {
	if in == nil {
		return nil
	}
	out := new(GroupVersionKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Message) DeepCopyInto(out *Message) {
	*out = *in
	in.Stream.DeepCopyInto(&out.Stream)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Message.
func (in *Message) DeepCopy() *Message {
	if in == nil {
		return nil
	}
	out := new(Message)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.CompletedAt.DeepCopyInto(&out.CompletedAt)
	if in.LatestEvent != nil {
		in, out := &in.LatestEvent, &out.LatestEvent
		*out = new(EventWrapper)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFilter) DeepCopyInto(out *ResourceFilter) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.CreatedBy.DeepCopyInto(&out.CreatedBy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFilter.
func (in *ResourceFilter) DeepCopy() *ResourceFilter {
	if in == nil {
		return nil
	}
	out := new(ResourceFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceObject) DeepCopyInto(out *ResourceObject) {
	*out = *in
	out.GroupVersionKind = in.GroupVersionKind
	in.Source.DeepCopyInto(&out.Source)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ResourceParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceObject.
func (in *ResourceObject) DeepCopy() *ResourceObject {
	if in == nil {
		return nil
	}
	out := new(ResourceObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceParameter) DeepCopyInto(out *ResourceParameter) {
	*out = *in
	if in.Src != nil {
		in, out := &in.Src, &out.Src
		*out = new(ResourceParameterSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceParameter.
func (in *ResourceParameter) DeepCopy() *ResourceParameter {
	if in == nil {
		return nil
	}
	out := new(ResourceParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceParameterSource) DeepCopyInto(out *ResourceParameterSource) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceParameterSource.
func (in *ResourceParameterSource) DeepCopy() *ResourceParameterSource {
	if in == nil {
		return nil
	}
	out := new(ResourceParameterSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSignal) DeepCopyInto(out *ResourceSignal) {
	*out = *in
	out.GroupVersionKind = in.GroupVersionKind
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(ResourceFilter)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSignal.
func (in *ResourceSignal) DeepCopy() *ResourceSignal {
	if in == nil {
		return nil
	}
	out := new(ResourceSignal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryStrategy) DeepCopyInto(out *RetryStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryStrategy.
func (in *RetryStrategy) DeepCopy() *RetryStrategy {
	if in == nil {
		return nil
	}
	out := new(RetryStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Artifact) DeepCopyInto(out *S3Artifact) {
	*out = *in
	in.S3Bucket.DeepCopyInto(&out.S3Bucket)
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(S3Filter)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Artifact.
func (in *S3Artifact) DeepCopy() *S3Artifact {
	if in == nil {
		return nil
	}
	out := new(S3Artifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
	in.AccessKey.DeepCopyInto(&out.AccessKey)
	in.SecretKey.DeepCopyInto(&out.SecretKey)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Filter) DeepCopyInto(out *S3Filter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Filter.
func (in *S3Filter) DeepCopy() *S3Filter {
	if in == nil {
		return nil
	}
	out := new(S3Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sensor) DeepCopyInto(out *Sensor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sensor.
func (in *Sensor) DeepCopy() *Sensor {
	if in == nil {
		return nil
	}
	out := new(Sensor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sensor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorList) DeepCopyInto(out *SensorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sensor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorList.
func (in *SensorList) DeepCopy() *SensorList {
	if in == nil {
		return nil
	}
	out := new(SensorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SensorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorSpec) DeepCopyInto(out *SensorSpec) {
	*out = *in
	if in.Signals != nil {
		in, out := &in.Signals, &out.Signals
		*out = make([]Signal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Escalation.DeepCopyInto(&out.Escalation)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorSpec.
func (in *SensorSpec) DeepCopy() *SensorSpec {
	if in == nil {
		return nil
	}
	out := new(SensorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorStatus) DeepCopyInto(out *SensorStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.CompletedAt.DeepCopyInto(&out.CompletedAt)
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]NodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorStatus.
func (in *SensorStatus) DeepCopy() *SensorStatus {
	if in == nil {
		return nil
	}
	out := new(SensorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Signal) DeepCopyInto(out *Signal) {
	*out = *in
	if in.Stream != nil {
		in, out := &in.Stream, &out.Stream
		*out = new(Stream)
		(*in).DeepCopyInto(*out)
	}
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		*out = new(ArtifactSignal)
		(*in).DeepCopyInto(*out)
	}
	if in.Calendar != nil {
		in, out := &in.Calendar, &out.Calendar
		*out = new(CalendarSignal)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceSignal)
		(*in).DeepCopyInto(*out)
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookSignal)
		**out = **in
	}
	in.Filters.DeepCopyInto(&out.Filters)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Signal.
func (in *Signal) DeepCopy() *Signal {
	if in == nil {
		return nil
	}
	out := new(Signal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalFilter) DeepCopyInto(out *SignalFilter) {
	*out = *in
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(TimeFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(EventContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]*DataFilter, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DataFilter)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalFilter.
func (in *SignalFilter) DeepCopy() *SignalFilter {
	if in == nil {
		return nil
	}
	out := new(SignalFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stream) DeepCopyInto(out *Stream) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stream.
func (in *Stream) DeepCopy() *Stream {
	if in == nil {
		return nil
	}
	out := new(Stream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeFilter) DeepCopyInto(out *TimeFilter) {
	*out = *in
	in.Start.DeepCopyInto(&out.Start)
	in.Stop.DeepCopyInto(&out.Stop)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeFilter.
func (in *TimeFilter) DeepCopy() *TimeFilter {
	if in == nil {
		return nil
	}
	out := new(TimeFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceObject)
		(*in).DeepCopyInto(*out)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(Message)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryStrategy != nil {
		in, out := &in.RetryStrategy, &out.RetryStrategy
		*out = new(RetryStrategy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URI) DeepCopyInto(out *URI) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URI.
func (in *URI) DeepCopy() *URI {
	if in == nil {
		return nil
	}
	out := new(URI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URLArtifact) DeepCopyInto(out *URLArtifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URLArtifact.
func (in *URLArtifact) DeepCopy() *URLArtifact {
	if in == nil {
		return nil
	}
	out := new(URLArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSignal) DeepCopyInto(out *WebhookSignal) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSignal.
func (in *WebhookSignal) DeepCopy() *WebhookSignal {
	if in == nil {
		return nil
	}
	out := new(WebhookSignal)
	in.DeepCopyInto(out)
	return out
}
