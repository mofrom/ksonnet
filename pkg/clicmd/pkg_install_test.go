// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package clicmd

import (
	"testing"

	"github.com/ksonnet/ksonnet/pkg/actions"
)

func Test_pkgInstallCmd(t *testing.T) {
	cases := []cmdTestCase{
		{
			name:   "in general",
			args:   []string{"pkg", "install", "package-name"},
			action: actionPkgInstall,
			expected: map[string]interface{}{
				actions.OptionApp:           nil,
				actions.OptionPkgName:       "package-name",
				actions.OptionName:          "",
				actions.OptionEnvName:       "",
				actions.OptionForce:         false,
				actions.OptionTLSSkipVerify: false,
			},
		},
		{
			name:   "with env flag",
			args:   []string{"pkg", "install", "--env", "production", "package-name"},
			action: actionPkgInstall,
			expected: map[string]interface{}{
				actions.OptionApp:           nil,
				actions.OptionPkgName:       "package-name",
				actions.OptionName:          "",
				actions.OptionEnvName:       "production",
				actions.OptionForce:         false,
				actions.OptionTLSSkipVerify: false,
			},
		},
		{
			name:   "force install",
			args:   []string{"pkg", "install", "package-name", "--force"},
			action: actionPkgInstall,
			expected: map[string]interface{}{
				actions.OptionApp:           nil,
				actions.OptionPkgName:       "package-name",
				actions.OptionName:          "",
				actions.OptionEnvName:       "",
				actions.OptionForce:         true,
				actions.OptionTLSSkipVerify: false,
			},
		},
		{
			name:  "invalid args",
			args:  []string{"pkg", "install"},
			isErr: true,
		},
	}

	runTestCmd(t, cases)
}
