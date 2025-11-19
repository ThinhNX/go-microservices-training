package features
import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

func InitApiV1Client() {
	kubeconfig := flag.String("kubeconfig", getKubeconfigPath(), "Path to the kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	fmt.Printf("Error building Kubernetes client: %v\n", err)
	// 	os.Exit(1)
	// }

	apiextensionsClient, err := apiextensionsv1.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error building ApiextensionsV1 client: %v\n", err)
		os.Exit(1)
	}
	namespace := "default" // Replace with your actual namespace

	// List Custom Resource Definitions (CRDs)
	crdList, err := apiextensionsClient.ApiextensionsV1().CustomResourceDefinitions().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing CRDs: %v\n", err)
		os.Exit(1)
	}

	// Print information about each CRD
	fmt.Printf("Custom Resource Definitions in Namespace %s:\n", namespace)
	for _, crd := range crdList.Items {
		fmt.Printf("Name: %s\n", crd.GetName())
		// Add more fields or information as needed
	}

	// Retrieve and print instances of a custom resource
	// Replace "example.com" and "v1" with the API group and version of your CRD
	// resourceType := "chaosresults"
	// customResourceList, err := clientset.Resource(metav1.GroupVersionResource{
	// 	Group:    "litmuschaos.io",
	// 	Version:  "v1alpha1",
	// 	Resource: resourceType,
	// }).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	fmt.Printf("Error listing custom resources: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("Custom Resources of Type %s in Namespace %s:\n", resourceType, namespace)
	// for _, resource := range customResourceList.Items {
	// 	fmt.Printf("Name: %s\n", resource.GetName())
	// 	// Add more fields or information as needed
	// }
}

func getKubeconfigPath() string {
	home := homedir.HomeDir()
	return filepath.Join(home, ".kube", "config")
}