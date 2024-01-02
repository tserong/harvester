/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	loggingv1beta1 "github.com/banzaicloud/logging-operator/pkg/sdk/logging/api/v1beta1"
	harvesterhciv1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	k8scnicncfiov1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	longhornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	catalogv1 "github.com/rancher/rancher/pkg/apis/catalog.cattle.io/v1"
	managementv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	upgradev1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	networkingv1 "k8s.io/api/networking/v1"
	storagev1 "k8s.io/api/storage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubevirtv1 "kubevirt.io/api/core/v1"
	clusterv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	catalogv1.AddToScheme,
	clusterv1alpha4.AddToScheme,
	harvesterhciv1beta1.AddToScheme,
	k8scnicncfiov1.AddToScheme,
	kubevirtv1.AddToScheme,
	loggingv1beta1.AddToScheme,
	longhornv1beta2.AddToScheme,
	managementv3.AddToScheme,
	monitoringv1.AddToScheme,
	networkingv1.AddToScheme,
	snapshotv1.AddToScheme,
	storagev1.AddToScheme,
	upgradev1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//	import (
//	  "k8s.io/client-go/kubernetes"
//	  clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//	  aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//	)
//
//	kclientset, _ := kubernetes.NewForConfig(c)
//	_ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
