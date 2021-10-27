// +build linux

/*
Copyright © 2021 SUSE LLC
1
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
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rancher-sandbox/rancher-desktop/src/wsl-helper/pkg/dockerproxy"
)

var dockerproxyStartViper = viper.New()

// dockerproxyStartCmd is the `wsl-helper docker-proxy start` command.
// This command is Linux-only.
var dockerproxyStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the docker daemon using vsock",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dockerproxy.Start(dockerproxyStartViper.GetUint32("port"), args)
	},
}

func init() {
	dockerproxyStartCmd.Flags().Uint32("port", 0, "Vsock port to listen on")
	dockerproxyStartViper.AutomaticEnv()
	dockerproxyStartViper.BindPFlags(dockerproxyStartCmd.Flags())
	dockerproxyCmd.AddCommand(dockerproxyStartCmd)
}
