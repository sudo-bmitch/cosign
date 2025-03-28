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

// AttestOptions is the top level wrapper for the attest command.
type AttestOptions struct {
	Key       string
	Cert      string
	Upload    bool
	Force     bool
	Recursive bool

	Rekor       RekorOptions
	Fulcio      FulcioOptions
	SecurityKey SecurityKeyOptions
	Predicate   PredicateOptions
	Registry    RegistryOptions
}

var _ Interface = (*AttestOptions)(nil)

// AddFlags implements Interface
func (o *AttestOptions) AddFlags(cmd *cobra.Command) {
	o.SecurityKey.AddFlags(cmd)
	o.Predicate.AddFlags(cmd)
	o.Fulcio.AddFlags(cmd)
	o.Rekor.AddFlags(cmd)
	o.Registry.AddFlags(cmd)

	cmd.Flags().StringVar(&o.Key, "key", "",
		"path to the private key file, KMS URI or Kubernetes Secret")

	cmd.Flags().StringVar(&o.Cert, "cert", "",
		"path to the x509 certificate to include in the Signature")

	cmd.Flags().BoolVar(&o.Upload, "upload", true,
		"whether to upload the signature")

	cmd.Flags().BoolVarP(&o.Force, "force", "f", false,
		"skip warnings and confirmations")

	cmd.Flags().BoolVarP(&o.Recursive, "recursive", "r", false,
		"if a multi-arch image is specified, additionally sign each discrete image")
}
