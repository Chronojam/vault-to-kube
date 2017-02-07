package main

import (
	"fmt"
///	"reflect"
//	types "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/pkg/api/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



func main() {
	ApplicationName := "helloworld"
	UpdateSecret(ApplicationName, "muh-secrets", map[string][]byte{})
}

func UpdateSecret(applicationName string, name string, content map[string][]byte) {
	notFoundErrorString := fmt.Sprintf("secrets \"%s\" not found", name)
	// First we'll make sure the namespace exists, and create it if not.
	CreateNamespace(applicationName)

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	_, err = clientset.CoreV1().Secrets(applicationName).Get(name, meta_v1.GetOptions{})
	if err != nil {
		if string(err.Error()) != notFoundErrorString {
			panic(err)
		}

		_, err := clientset.CoreV1().Secrets(applicationName).Create(&v1.Secret{ObjectMeta: meta_v1.ObjectMeta{Name: name}, Data: content})
		if err != nil {
			panic(err)
		}
		return
	}

}

func CreateNamespace(name string) {
	notFoundErrorString := fmt.Sprintf("namespaces \"%s\" not found", name)
	// We're inside the cluster, of course we're inside the cluster.
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	_, err = clientset.CoreV1().Namespaces().Get(name, meta_v1.GetOptions{})
	if err != nil {
		// Does not exist error is too generic for us to use, so we'll have to string match.
		if string(err.Error()) != notFoundErrorString {
			// Actual error.
			panic(err)
		}
		_, err := clientset.CoreV1().Namespaces().Create(&v1.Namespace{ObjectMeta: meta_v1.ObjectMeta{Name: name}})
		if err != nil {
			// Actual error.
			panic(err)
		}
		fmt.Println("Created namespace: ", name)
	}
}
