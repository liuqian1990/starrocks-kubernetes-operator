//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingPolicy) DeepCopyInto(out *AutoScalingPolicy) {
	*out = *in
	if in.HPAPolicy != nil {
		in, out := &in.HPAPolicy, &out.HPAPolicy
		*out = new(HPAPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingPolicy.
func (in *AutoScalingPolicy) DeepCopy() *AutoScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(AutoScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapInfo) DeepCopyInto(out *ConfigMapInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapInfo.
func (in *ConfigMapInfo) DeepCopy() *ConfigMapInfo {
	if in == nil {
		return nil
	}
	out := new(ConfigMapInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerResourceMetricSource) DeepCopyInto(out *ContainerResourceMetricSource) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerResourceMetricSource.
func (in *ContainerResourceMetricSource) DeepCopy() *ContainerResourceMetricSource {
	if in == nil {
		return nil
	}
	out := new(ContainerResourceMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossVersionObjectReference) DeepCopyInto(out *CrossVersionObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossVersionObjectReference.
func (in *CrossVersionObjectReference) DeepCopy() *CrossVersionObjectReference {
	if in == nil {
		return nil
	}
	out := new(CrossVersionObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetricSource) DeepCopyInto(out *ExternalMetricSource) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricSource.
func (in *ExternalMetricSource) DeepCopy() *ExternalMetricSource {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAPolicy) DeepCopyInto(out *HPAPolicy) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAPolicy.
func (in *HPAPolicy) DeepCopy() *HPAPolicy {
	if in == nil {
		return nil
	}
	out := new(HPAPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAScalingPolicy) DeepCopyInto(out *HPAScalingPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAScalingPolicy.
func (in *HPAScalingPolicy) DeepCopy() *HPAScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(HPAScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAScalingRules) DeepCopyInto(out *HPAScalingRules) {
	*out = *in
	if in.StabilizationWindowSeconds != nil {
		in, out := &in.StabilizationWindowSeconds, &out.StabilizationWindowSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SelectPolicy != nil {
		in, out := &in.SelectPolicy, &out.SelectPolicy
		*out = new(ScalingPolicySelect)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]HPAScalingPolicy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAScalingRules.
func (in *HPAScalingRules) DeepCopy() *HPAScalingRules {
	if in == nil {
		return nil
	}
	out := new(HPAScalingRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerBehavior) DeepCopyInto(out *HorizontalPodAutoscalerBehavior) {
	*out = *in
	if in.ScaleUp != nil {
		in, out := &in.ScaleUp, &out.ScaleUp
		*out = new(HPAScalingRules)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleDown != nil {
		in, out := &in.ScaleDown, &out.ScaleDown
		*out = new(HPAScalingRules)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerBehavior.
func (in *HorizontalPodAutoscalerBehavior) DeepCopy() *HorizontalPodAutoscalerBehavior {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerBehavior)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalScaler) DeepCopyInto(out *HorizontalScaler) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalScaler.
func (in *HorizontalScaler) DeepCopy() *HorizontalScaler {
	if in == nil {
		return nil
	}
	out := new(HorizontalScaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricIdentifier) DeepCopyInto(out *MetricIdentifier) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricIdentifier.
func (in *MetricIdentifier) DeepCopy() *MetricIdentifier {
	if in == nil {
		return nil
	}
	out := new(MetricIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(ObjectMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(PodsMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerResource != nil {
		in, out := &in.ContainerResource, &out.ContainerResource
		*out = new(ContainerResourceMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalMetricSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTarget) DeepCopyInto(out *MetricTarget) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AverageValue != nil {
		in, out := &in.AverageValue, &out.AverageValue
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AverageUtilization != nil {
		in, out := &in.AverageUtilization, &out.AverageUtilization
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTarget.
func (in *MetricTarget) DeepCopy() *MetricTarget {
	if in == nil {
		return nil
	}
	out := new(MetricTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMetricSource) DeepCopyInto(out *ObjectMetricSource) {
	*out = *in
	out.DescribedObject = in.DescribedObject
	in.Target.DeepCopyInto(&out.Target)
	in.Metric.DeepCopyInto(&out.Metric)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMetricSource.
func (in *ObjectMetricSource) DeepCopy() *ObjectMetricSource {
	if in == nil {
		return nil
	}
	out := new(ObjectMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodsMetricSource) DeepCopyInto(out *PodsMetricSource) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodsMetricSource.
func (in *PodsMetricSource) DeepCopy() *PodsMetricSource {
	if in == nil {
		return nil
	}
	out := new(PodsMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetricSource) DeepCopyInto(out *ResourceMetricSource) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetricSource.
func (in *ResourceMetricSource) DeepCopy() *ResourceMetricSource {
	if in == nil {
		return nil
	}
	out := new(ResourceMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretInfo) DeepCopyInto(out *SecretInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretInfo.
func (in *SecretInfo) DeepCopy() *SecretInfo {
	if in == nil {
		return nil
	}
	out := new(SecretInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksBeSpec) DeepCopyInto(out *StarRocksBeSpec) {
	*out = *in
	in.StarRocksComponentSpec.DeepCopyInto(&out.StarRocksComponentSpec)
	if in.BeEnvVars != nil {
		in, out := &in.BeEnvVars, &out.BeEnvVars
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksBeSpec.
func (in *StarRocksBeSpec) DeepCopy() *StarRocksBeSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksBeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksBeStatus) DeepCopyInto(out *StarRocksBeStatus) {
	*out = *in
	in.StarRocksComponentStatus.DeepCopyInto(&out.StarRocksComponentStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksBeStatus.
func (in *StarRocksBeStatus) DeepCopy() *StarRocksBeStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksBeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksCluster) DeepCopyInto(out *StarRocksCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksCluster.
func (in *StarRocksCluster) DeepCopy() *StarRocksCluster {
	if in == nil {
		return nil
	}
	out := new(StarRocksCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StarRocksCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksClusterList) DeepCopyInto(out *StarRocksClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StarRocksCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksClusterList.
func (in *StarRocksClusterList) DeepCopy() *StarRocksClusterList {
	if in == nil {
		return nil
	}
	out := new(StarRocksClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StarRocksClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksClusterSpec) DeepCopyInto(out *StarRocksClusterSpec) {
	*out = *in
	if in.StarRocksFeSpec != nil {
		in, out := &in.StarRocksFeSpec, &out.StarRocksFeSpec
		*out = new(StarRocksFeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StarRocksBeSpec != nil {
		in, out := &in.StarRocksBeSpec, &out.StarRocksBeSpec
		*out = new(StarRocksBeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StarRocksCnSpec != nil {
		in, out := &in.StarRocksCnSpec, &out.StarRocksCnSpec
		*out = new(StarRocksCnSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksClusterSpec.
func (in *StarRocksClusterSpec) DeepCopy() *StarRocksClusterSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksClusterStatus) DeepCopyInto(out *StarRocksClusterStatus) {
	*out = *in
	if in.StarRocksFeStatus != nil {
		in, out := &in.StarRocksFeStatus, &out.StarRocksFeStatus
		*out = new(StarRocksFeStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.StarRocksBeStatus != nil {
		in, out := &in.StarRocksBeStatus, &out.StarRocksBeStatus
		*out = new(StarRocksBeStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.StarRocksCnStatus != nil {
		in, out := &in.StarRocksCnStatus, &out.StarRocksCnStatus
		*out = new(StarRocksCnStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksClusterStatus.
func (in *StarRocksClusterStatus) DeepCopy() *StarRocksClusterStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksCnSpec) DeepCopyInto(out *StarRocksCnSpec) {
	*out = *in
	in.StarRocksComponentSpec.DeepCopyInto(&out.StarRocksComponentSpec)
	if in.CnEnvVars != nil {
		in, out := &in.CnEnvVars, &out.CnEnvVars
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoScalingPolicy != nil {
		in, out := &in.AutoScalingPolicy, &out.AutoScalingPolicy
		*out = new(AutoScalingPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksCnSpec.
func (in *StarRocksCnSpec) DeepCopy() *StarRocksCnSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksCnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksCnStatus) DeepCopyInto(out *StarRocksCnStatus) {
	*out = *in
	in.StarRocksComponentStatus.DeepCopyInto(&out.StarRocksComponentStatus)
	out.HorizontalScaler = in.HorizontalScaler
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksCnStatus.
func (in *StarRocksCnStatus) DeepCopy() *StarRocksCnStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksCnStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksComponentSpec) DeepCopyInto(out *StarRocksComponentSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FsGroup != nil {
		in, out := &in.FsGroup, &out.FsGroup
		*out = new(int64)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(StarRocksService)
		(*in).DeepCopyInto(*out)
	}
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	out.ConfigMapInfo = in.ConfigMapInfo
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]SecretInfo, len(*in))
		copy(*out, *in)
	}
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(StarRocksProbe)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]corev1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageVolumes != nil {
		in, out := &in.StorageVolumes, &out.StorageVolumes
		*out = make([]StorageVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksComponentSpec.
func (in *StarRocksComponentSpec) DeepCopy() *StarRocksComponentSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksComponentStatus) DeepCopyInto(out *StarRocksComponentStatus) {
	*out = *in
	if in.FailedInstances != nil {
		in, out := &in.FailedInstances, &out.FailedInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreatingInstances != nil {
		in, out := &in.CreatingInstances, &out.CreatingInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RunningInstances != nil {
		in, out := &in.RunningInstances, &out.RunningInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksComponentStatus.
func (in *StarRocksComponentStatus) DeepCopy() *StarRocksComponentStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksFeSpec) DeepCopyInto(out *StarRocksFeSpec) {
	*out = *in
	in.StarRocksComponentSpec.DeepCopyInto(&out.StarRocksComponentSpec)
	if in.FeEnvVars != nil {
		in, out := &in.FeEnvVars, &out.FeEnvVars
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksFeSpec.
func (in *StarRocksFeSpec) DeepCopy() *StarRocksFeSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksFeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksFeStatus) DeepCopyInto(out *StarRocksFeStatus) {
	*out = *in
	in.StarRocksComponentStatus.DeepCopyInto(&out.StarRocksComponentStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksFeStatus.
func (in *StarRocksFeStatus) DeepCopy() *StarRocksFeStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksFeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksProbe) DeepCopyInto(out *StarRocksProbe) {
	*out = *in
	if in.InitialDelaySeconds != nil {
		in, out := &in.InitialDelaySeconds, &out.InitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksProbe.
func (in *StarRocksProbe) DeepCopy() *StarRocksProbe {
	if in == nil {
		return nil
	}
	out := new(StarRocksProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksService) DeepCopyInto(out *StarRocksService) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]StarRocksServicePort, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksService.
func (in *StarRocksService) DeepCopy() *StarRocksService {
	if in == nil {
		return nil
	}
	out := new(StarRocksService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksServicePort) DeepCopyInto(out *StarRocksServicePort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksServicePort.
func (in *StarRocksServicePort) DeepCopy() *StarRocksServicePort {
	if in == nil {
		return nil
	}
	out := new(StarRocksServicePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVolume) DeepCopyInto(out *StorageVolume) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVolume.
func (in *StorageVolume) DeepCopy() *StorageVolume {
	if in == nil {
		return nil
	}
	out := new(StorageVolume)
	in.DeepCopyInto(out)
	return out
}
