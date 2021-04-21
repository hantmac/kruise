/*
Copyright 2019 The Kruise Authors.

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

package util

import (
	"encoding/json"
	"reflect"
)

// DumpJSON returns the JSON encoding
func DumpJSON(o interface{}) string {
	j, _ := json.Marshal(o)
	return string(j)
}

// IsJSONEqual checks if two objects are equal after encoding json
func IsJSONEqual(o1, o2 interface{}) bool {
	oj1, _ := json.Marshal(o1)
	oj2, _ := json.Marshal(o2)

	om1 := make(map[string]interface{})
	om2 := make(map[string]interface{})
	_ = json.Unmarshal(oj1, &om1)
	_ = json.Unmarshal(oj2, &om2)

	return reflect.DeepEqual(om1, om2)
}
