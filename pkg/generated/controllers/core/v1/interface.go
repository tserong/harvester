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

package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"github.com/rancher/wrangler/v3/pkg/schemes"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	PersistentVolume() PersistentVolumeController
	ResourceQuota() ResourceQuotaController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) PersistentVolume() PersistentVolumeController {
	return generic.NewNonNamespacedController[*v1.PersistentVolume, *v1.PersistentVolumeList](schema.GroupVersionKind{Group: "", Version: "v1", Kind: "PersistentVolume"}, "persistentvolumes", v.controllerFactory)
}

func (v *version) ResourceQuota() ResourceQuotaController {
	return generic.NewController[*v1.ResourceQuota, *v1.ResourceQuotaList](schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ResourceQuota"}, "resourcequotas", true, v.controllerFactory)
}
