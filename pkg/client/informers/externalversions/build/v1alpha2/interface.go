/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "github.com/pivotal/kpack/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Builds returns a BuildInformer.
	Builds() BuildInformer
	// Builders returns a BuilderInformer.
	Builders() BuilderInformer
	// Buildpacks returns a BuildpackInformer.
	Buildpacks() BuildpackInformer
	// ClusterBuilders returns a ClusterBuilderInformer.
	ClusterBuilders() ClusterBuilderInformer
	// ClusterBuildpacks returns a ClusterBuildpackInformer.
	ClusterBuildpacks() ClusterBuildpackInformer
	// ClusterStacks returns a ClusterStackInformer.
	ClusterStacks() ClusterStackInformer
	// ClusterStores returns a ClusterStoreInformer.
	ClusterStores() ClusterStoreInformer
	// Images returns a ImageInformer.
	Images() ImageInformer
	// SourceResolvers returns a SourceResolverInformer.
	SourceResolvers() SourceResolverInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Builds returns a BuildInformer.
func (v *version) Builds() BuildInformer {
	return &buildInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Builders returns a BuilderInformer.
func (v *version) Builders() BuilderInformer {
	return &builderInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Buildpacks returns a BuildpackInformer.
func (v *version) Buildpacks() BuildpackInformer {
	return &buildpackInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterBuilders returns a ClusterBuilderInformer.
func (v *version) ClusterBuilders() ClusterBuilderInformer {
	return &clusterBuilderInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterBuildpacks returns a ClusterBuildpackInformer.
func (v *version) ClusterBuildpacks() ClusterBuildpackInformer {
	return &clusterBuildpackInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterStacks returns a ClusterStackInformer.
func (v *version) ClusterStacks() ClusterStackInformer {
	return &clusterStackInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterStores returns a ClusterStoreInformer.
func (v *version) ClusterStores() ClusterStoreInformer {
	return &clusterStoreInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Images returns a ImageInformer.
func (v *version) Images() ImageInformer {
	return &imageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SourceResolvers returns a SourceResolverInformer.
func (v *version) SourceResolvers() SourceResolverInformer {
	return &sourceResolverInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
