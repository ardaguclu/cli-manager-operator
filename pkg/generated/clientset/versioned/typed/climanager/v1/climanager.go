// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/cli-manager-operator/pkg/apis/climanager/v1"
	climanagerv1 "github.com/openshift/cli-manager-operator/pkg/generated/applyconfiguration/climanager/v1"
	scheme "github.com/openshift/cli-manager-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// CliManagersGetter has a method to return a CliManagerInterface.
// A group's client should implement this interface.
type CliManagersGetter interface {
	CliManagers(namespace string) CliManagerInterface
}

// CliManagerInterface has methods to work with CliManager resources.
type CliManagerInterface interface {
	Create(ctx context.Context, cliManager *v1.CliManager, opts metav1.CreateOptions) (*v1.CliManager, error)
	Update(ctx context.Context, cliManager *v1.CliManager, opts metav1.UpdateOptions) (*v1.CliManager, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, cliManager *v1.CliManager, opts metav1.UpdateOptions) (*v1.CliManager, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CliManager, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CliManagerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CliManager, err error)
	Apply(ctx context.Context, cliManager *climanagerv1.CliManagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CliManager, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, cliManager *climanagerv1.CliManagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CliManager, err error)
	CliManagerExpansion
}

// cliManagers implements CliManagerInterface
type cliManagers struct {
	*gentype.ClientWithListAndApply[*v1.CliManager, *v1.CliManagerList, *climanagerv1.CliManagerApplyConfiguration]
}

// newCliManagers returns a CliManagers
func newCliManagers(c *ClimanagersV1Client, namespace string) *cliManagers {
	return &cliManagers{
		gentype.NewClientWithListAndApply[*v1.CliManager, *v1.CliManagerList, *climanagerv1.CliManagerApplyConfiguration](
			"climanagers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.CliManager { return &v1.CliManager{} },
			func() *v1.CliManagerList { return &v1.CliManagerList{} }),
	}
}
