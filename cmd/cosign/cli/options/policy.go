//
// Copyright 2021 The Sigstore Authors.
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

package options

import (
	"github.com/spf13/cobra"
)

// PolicyInitOptions is the top level wrapper for the policy-init command.
type PolicyInitOptions struct {
	ImageRef    string
	Maintainers []string
	Threshold   int
	OutFile     string
}

var _ Interface = (*PolicyInitOptions)(nil)

// AddFlags implements Interface
func (o *PolicyInitOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.ImageRef, "namespace", "ns",
		"registry namespace that the root policy belongs to")

	cmd.Flags().StringVar(&o.OutFile, "out", "o",
		"output policy locally")

	cmd.Flags().IntVar(&o.Threshold, "threshold", 1,
		"threshold for root policy signers")

	cmd.Flags().StringSliceVarP(&o.Maintainers, "maintainers", "m", nil,
		"list of maintainers to add to the root policy")
}
