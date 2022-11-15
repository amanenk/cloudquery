// Code generated by codegen; DO NOT EDIT.

package rbac

import (
	"testing"

	client "github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/rbac/v1"
	// k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createClusterRoleBindings(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.ClusterRoleBinding{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockClusterRoleBindingInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.ClusterRoleBindingList{Items: []resource.ClusterRoleBinding{r}}, nil,
	)

	serviceClient := resourcemock.NewMockRbacV1Interface(ctrl)

	serviceClient.EXPECT().ClusterRoleBindings().Return(resourceClient)

	client := mocks.NewMockInterface(ctrl)
	client.EXPECT().RbacV1().Return(serviceClient)

	return client
}

func TestClusterRoleBindings(t *testing.T) {
	client.K8sMockTestHelper(t, ClusterRoleBindings(), createClusterRoleBindings)
}
