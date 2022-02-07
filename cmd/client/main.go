// Copyright 2022 Leonard Vincent Simon Pahlke
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

package main

import (
	"whopper/cmd/client/cmd"
	"whopper/pkg/util"

	"go.uber.org/zap"
)

// Entry point for the wclient cli
// The cli uses cobra and therefore cobra project structuring conventions
// und the cmd/ folder the main root command (root.go) and sub commands are defined
func main() {
	if err := cmd.Execute(); err != nil {
		util.GetLogger(zap.InfoLevel).Fatalw("could not execute command", "error", err)
	}
}
