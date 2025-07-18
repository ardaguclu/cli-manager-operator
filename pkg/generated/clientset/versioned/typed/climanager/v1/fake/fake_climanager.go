// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/cli-manager-operator/pkg/apis/climanager/v1"
	climanagerv1 "github.com/openshift/cli-manager-operator/pkg/generated/applyconfiguration/climanager/v1"
	typedclimanagerv1 "github.com/openshift/cli-manager-operator/pkg/generated/clientset/versioned/typed/climanager/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeCliManagers implements CliManagerInterface
type fakeCliManagers struct {
	*gentype.FakeClientWithListAndApply[*v1.CliManager, *v1.CliManagerList, *climanagerv1.CliManagerApplyConfiguration]
	Fake *FakeClimanagersV1
}

func newFakeCliManagers(fake *FakeClimanagersV1, namespace string) typedclimanagerv1.CliManagerInterface {
	return &fakeCliManagers{
		gentype.NewFakeClientWithListAndApply[*v1.CliManager, *v1.CliManagerList, *climanagerv1.CliManagerApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("climanagers"),
			v1.SchemeGroupVersion.WithKind("CliManager"),
			func() *v1.CliManager { return &v1.CliManager{} },
			func() *v1.CliManagerList { return &v1.CliManagerList{} },
			func(dst, src *v1.CliManagerList) { dst.ListMeta = src.ListMeta },
			func(list *v1.CliManagerList) []*v1.CliManager { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.CliManagerList, items []*v1.CliManager) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
