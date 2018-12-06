// Copyright © 2018 Infostellar, Inc.
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

package groundstation

import (
	"github.com/spf13/cobra"

	"github.com/infostellarinc/stellarcli/cmd/util"
)

var (
	gsUsage = util.Normalize("ground-station")
	gsShort = util.Normalize("Commands for working with ground stations.")
)

// Create ground station command.
func NewGroundStationCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     gsUsage,
		Aliases: []string{"gs"},
		Short:   gsShort,
	}

	command.AddCommand(NewAddUWCommand())
	command.AddCommand(NewDeleteUWCommand())
	command.AddCommand(NewListPlansCommand())
	command.AddCommand(NewListUWCommand())

	return command
}