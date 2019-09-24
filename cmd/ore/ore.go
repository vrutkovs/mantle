// Copyright 2014 CoreOS, Inc.
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
	"github.com/coreos/pkg/capnslog"
	"github.com/spf13/cobra"

	"github.com/coreos/mantle/cli"
)

var (
	root = &cobra.Command{
		Use:   "ore [command]",
		Short: "cloud image creation and upload tools",
	}

	logDebug   bool
	logVerbose bool
	logLevel   = capnslog.NOTICE
	plog = capnslog.NewPackageLogger("github.com/coreos/mantle", "ore")
)

func main() {
	root.PersistentFlags().Var(&logLevel, "log-level",
		"Set global log level.")
	root.PersistentFlags().BoolVarP(&logVerbose, "verbose", "v", false,
		"Alias for --log-level=INFO")
	root.PersistentFlags().BoolVarP(&logDebug, "debug", "d", false,
		"Alias for --log-level=DEBUG")
	cli.WrapPreRun(root, func(cmd *cobra.Command, args []string) error {
		cli.StartLogging(cmd)
		return nil
	})
	cli.Execute(root)
}
