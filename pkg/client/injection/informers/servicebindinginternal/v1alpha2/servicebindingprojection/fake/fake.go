/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/vmware-labs/service-bindings/pkg/client/injection/informers/factory/fake"
	servicebindingprojection "github.com/vmware-labs/service-bindings/pkg/client/injection/informers/servicebindinginternal/v1alpha2/servicebindingprojection"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = servicebindingprojection.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Internal().V1alpha2().ServiceBindingProjections()
	return context.WithValue(ctx, servicebindingprojection.Key{}, inf), inf.Informer()
}
