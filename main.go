package main

import (
	cm "github.com/mongodb/mongodb-kubernetes-operator/pkg/kube/configmap"
	"fmt"
)

func main() {
	fmt.Println("%v", cm.Client{
	})

}
