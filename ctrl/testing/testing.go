/*
Copyright 2018 The Kubernetes Authors.

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

package testing

import (
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/workqueue"
)

var _ runtime.Object = &ErrorType{}

// ErrorType implements runtime.Object but isn't registered in any Scheme and should cause errors in tests as a result.
type ErrorType struct{}

func (ErrorType) GetObjectKind() schema.ObjectKind { return nil }
func (ErrorType) DeepCopyObject() runtime.Object   { return nil }

var _ workqueue.RateLimitingInterface = Queue{}

// Queue implements a RateLimiting queue as a non-ratelimited queue for testing.
// This helps testing by having functions that use a RateLimiting queue synchronously add items to the queue.
type Queue struct {
	workqueue.Interface
}

func (q Queue) AddAfter(item interface{}, duration time.Duration) {
	q.Add(item)
}

func (q Queue) AddRateLimited(item interface{}) {
	q.Add(item)
}

func (q Queue) Forget(item interface{}) {}

func (q Queue) NumRequeues(item interface{}) int {
	return 0
}
