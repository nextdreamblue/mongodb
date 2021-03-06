/*
Copyright The KubeDB Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubedb.dev/apimachinery/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ElasticsearchModificationRequests returns a ElasticsearchModificationRequestInformer.
	ElasticsearchModificationRequests() ElasticsearchModificationRequestInformer
	// EtcdModificationRequests returns a EtcdModificationRequestInformer.
	EtcdModificationRequests() EtcdModificationRequestInformer
	// MemcachedModificationRequests returns a MemcachedModificationRequestInformer.
	MemcachedModificationRequests() MemcachedModificationRequestInformer
	// MongoDBModificationRequests returns a MongoDBModificationRequestInformer.
	MongoDBModificationRequests() MongoDBModificationRequestInformer
	// MySQLModificationRequests returns a MySQLModificationRequestInformer.
	MySQLModificationRequests() MySQLModificationRequestInformer
	// PerconaXtraDBModificationRequests returns a PerconaXtraDBModificationRequestInformer.
	PerconaXtraDBModificationRequests() PerconaXtraDBModificationRequestInformer
	// PgBouncerModificationRequests returns a PgBouncerModificationRequestInformer.
	PgBouncerModificationRequests() PgBouncerModificationRequestInformer
	// PostgresModificationRequests returns a PostgresModificationRequestInformer.
	PostgresModificationRequests() PostgresModificationRequestInformer
	// ProxySQLModificationRequests returns a ProxySQLModificationRequestInformer.
	ProxySQLModificationRequests() ProxySQLModificationRequestInformer
	// RedisModificationRequests returns a RedisModificationRequestInformer.
	RedisModificationRequests() RedisModificationRequestInformer
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

// ElasticsearchModificationRequests returns a ElasticsearchModificationRequestInformer.
func (v *version) ElasticsearchModificationRequests() ElasticsearchModificationRequestInformer {
	return &elasticsearchModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// EtcdModificationRequests returns a EtcdModificationRequestInformer.
func (v *version) EtcdModificationRequests() EtcdModificationRequestInformer {
	return &etcdModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MemcachedModificationRequests returns a MemcachedModificationRequestInformer.
func (v *version) MemcachedModificationRequests() MemcachedModificationRequestInformer {
	return &memcachedModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MongoDBModificationRequests returns a MongoDBModificationRequestInformer.
func (v *version) MongoDBModificationRequests() MongoDBModificationRequestInformer {
	return &mongoDBModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MySQLModificationRequests returns a MySQLModificationRequestInformer.
func (v *version) MySQLModificationRequests() MySQLModificationRequestInformer {
	return &mySQLModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PerconaXtraDBModificationRequests returns a PerconaXtraDBModificationRequestInformer.
func (v *version) PerconaXtraDBModificationRequests() PerconaXtraDBModificationRequestInformer {
	return &perconaXtraDBModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PgBouncerModificationRequests returns a PgBouncerModificationRequestInformer.
func (v *version) PgBouncerModificationRequests() PgBouncerModificationRequestInformer {
	return &pgBouncerModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PostgresModificationRequests returns a PostgresModificationRequestInformer.
func (v *version) PostgresModificationRequests() PostgresModificationRequestInformer {
	return &postgresModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ProxySQLModificationRequests returns a ProxySQLModificationRequestInformer.
func (v *version) ProxySQLModificationRequests() ProxySQLModificationRequestInformer {
	return &proxySQLModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// RedisModificationRequests returns a RedisModificationRequestInformer.
func (v *version) RedisModificationRequests() RedisModificationRequestInformer {
	return &redisModificationRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
