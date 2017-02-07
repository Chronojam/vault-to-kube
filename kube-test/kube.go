package main

import (
	"fmt"
	"reflect"
//	types "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
//	"k8s.io/client-go/pkg/api/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



func main() {
	CreateNamespace("Hello")
}

func CreateNamespace(name string) {
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
		// Namespace doesnt exist
		fmt.Println(err)
		fmt.Println(reflect.TypeOf(err).String())
	}
	//namespaces,  _ := clientset.CoreV1().Namespaces().Create(v1.Namespace{Name: name})
	//fmt.Println(namespaces)
}
