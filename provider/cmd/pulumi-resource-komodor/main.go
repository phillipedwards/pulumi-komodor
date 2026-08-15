// Copyright 2016-2018, Pulumi Corporation.
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
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"

	komodor "github.com/phillipedwards/pulumi-komodor/provider"
	"github.com/phillipedwards/pulumi-komodor/provider/pkg/version"
)

//go:embed schema.json
var pulumiSchema []byte

func main() {
	tfbridge.Main("komodor", version.Version, komodor.Provider(), withVersion(pulumiSchema, version.Version))
}

// withVersion stamps the linker-supplied version into the embedded schema.
//
// tfgen deliberately omits "version" from schema.json so the committed file stays
// stable across builds (the version is often derived from a git describe). Consumers
// do need it though: `pulumi package add` refuses to generate an SDK from a schema
// with no version, so serving the file verbatim makes the provider uninstallable.
func withVersion(schema []byte, v string) []byte {
	if v == "" {
		// Unversioned dev build; nothing useful to stamp.
		return schema
	}

	var spec map[string]any
	if err := json.Unmarshal(schema, &spec); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse embedded schema: %v\n", err)
		os.Exit(1)
	}
	spec["version"] = v

	stamped, err := json.Marshal(spec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to re-encode embedded schema: %v\n", err)
		os.Exit(1)
	}
	return stamped
}
