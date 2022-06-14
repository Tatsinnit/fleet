/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package utils

import (
	"fmt"
	"strconv"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

const (
	// TestCaseMsg is used in the table driven test
	TestCaseMsg string = "\nTest case:  %s"
)

// NewFakeRecorder makes a new fake event recorder that prints the object.
func NewFakeRecorder(bufferSize int) *record.FakeRecorder {
	recorder := record.NewFakeRecorder(bufferSize)
	recorder.IncludeObject = true
	return recorder
}

// GetEventString get the exact string literal of the event created by the fake event library.
func GetEventString(object runtime.Object, eventtype, reason, messageFmt string, args ...interface{}) string {
	return fmt.Sprintf(eventtype+" "+reason+" "+messageFmt, args...) +
		fmt.Sprintf(" involvedObject{kind=%s,apiVersion=%s}",
			object.GetObjectKind().GroupVersionKind().Kind, object.GetObjectKind().GroupVersionKind().GroupVersion())
}

// NewTestNodes return a set of nodes for test purpose. Those nodes have random names and capacities/ allocatable.
func NewTestNodes(ns string) []v1.Node {
	numOfNodes := RandSecureInt(10)
	var nodes []v1.Node
	for i := int64(0); i < numOfNodes; i++ {
		allocatableCPU := RandSecureInt(100)
		allocatableMemory := RandSecureInt(5)
		capacityCPU := RandSecureInt(100)
		capacityMemory := RandSecureInt(5)
		node := v1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "rand-" + strings.ToLower(RandStr()) + "-node",
				Namespace: ns,
			},
			Status: v1.NodeStatus{Allocatable: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse(strconv.FormatInt(allocatableCPU, 10) + "m"),
				v1.ResourceMemory: resource.MustParse(strconv.FormatInt(allocatableMemory, 10) + "G"),
			}, Capacity: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse(strconv.FormatInt(capacityCPU, 10) + "m"),
				v1.ResourceMemory: resource.MustParse(strconv.FormatInt(capacityMemory, 10) + "G"),
			}},
		}
		nodes = append(nodes, node)
	}
	return nodes
}
