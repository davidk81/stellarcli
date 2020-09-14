//
// Copyright © 2020 Infostellar, Inc.
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

package flag

import (
	"time"

	"github.com/spf13/cobra"
)

type CpuProfile struct {
	FileName string
}

// AddFlags Add a flag to the command.
func (f *CpuProfile) AddFlags(cmd *cobra.Command) {
	val := "cpu-profile-" + time.Now().Format("01-02-2006") + ".pprof"
	cmd.Flags().StringVarP(&f.FileName, "cpu-profile", "", val, "[Alpha feature] The file to write cpu profile to. Overwrites to file if it already exists. default: "+val)
}

// Validate flag values.
func (f *CpuProfile) Validate() error {
	return nil
}

// NewCpuProfileFlag Create a new CpuProfile with default values set.
func NewCpuProfileFlag() *CpuProfile {
	return &CpuProfile{}
}
