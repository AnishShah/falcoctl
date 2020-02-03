/*
Copyright © 2020 The Falco Authors

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
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// ReloadOptions represents the reload command options
type ReloadOptions struct {
	genericclioptions.IOStreams
}

// Validate validates the `reload` command options
func (o ReloadOptions) Validate(c *cobra.Command, args []string) error {
	return nil
}

// NewReloadOptions instantiates the `reload` command options
func NewReloadOptions(streams genericclioptions.IOStreams) CommandOptions {
	return &ReloadOptions{
		IOStreams: streams,
	}
}

// Reload creates the `reload` command
func Reload(streams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "reload",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		Short:                 "Reloads Falco rules",
		Long:                  `Reloads Falco rules`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
