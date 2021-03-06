/*
Copyright 2019 HAProxy Technologies

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

package actions

import (
	"fmt"
	"strings"
)

type ScIncGpc0 struct {
	ScID string
}

func (f *ScIncGpc0) Parse(parts []string) error {

	if len(parts) == 1 {

		data := strings.TrimPrefix(parts[0], "sc-inc-gpc0(")

		data = strings.TrimRight(data, ")")

		f.ScID = data

		return nil
	}

	return fmt.Errorf("not enough params")
}

func (f *ScIncGpc0) String() string {
	return fmt.Sprintf("sc-inc-gpc0(%s)", f.ScID)
}
