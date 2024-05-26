// Copyright 2024 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import "fmt"

// Must if err != nil just panic which err
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Must if err != nil just panic which msg and err
func Mustf(err error, format string, a ...any) {
	if err != nil {
		msg := fmt.Sprintf(format, a...)
		panic(fmt.Errorf("%s, %w", msg, err))
	}
}
