/*
Copyright © 2020 Red Hat, Inc.

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

package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"ocm-support-cli/cmd/ocm-support/accounts"
	"ocm-support-cli/cmd/ocm-support/organizations"
	registrycredentials "ocm-support-cli/cmd/ocm-support/registryCredentials"
	"ocm-support-cli/cmd/ocm-support/version"
)

// Created so that multiple inputs can be accepted
type levelFlag logrus.Level

func (l *levelFlag) String() string {
	// change this, this is just can example to satisfy the interface
	return logrus.Level(*l).String()
}

func (l *levelFlag) Set(value string) error {
	lvl, err := log.ParseLevel(strings.TrimSpace(value))
	if err == nil {
		*l = levelFlag(lvl)
	}
	return err
}

func (l *levelFlag) Type() string {
	return "string"
}

var (
	// some defaults for configuration
	defaultLogLevel = log.WarnLevel.String()
	logLevel        levelFlag
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "ocm-support",
	Short:         "support plugin for OCM",
	Long:          `This is a binary for ocm support plugin.`,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorln(err.Error())
		os.Exit(1)
	}
}

func init() {
	// Set default log level
	_ = logLevel.Set(defaultLogLevel)
	logLevelFlag := rootCmd.PersistentFlags().VarPF(&logLevel, "verbosity", "v", "Verbosity level: panic, fatal, error, warn, info, debug. Providing no level string will select info.")
	logLevelFlag.NoOptDefVal = logrus.InfoLevel.String()

	// Register sub-commands
	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(accounts.Cmd)
	rootCmd.AddCommand(organizations.Cmd)
	rootCmd.AddCommand(registrycredentials.Cmd)

	// Set the log level before each command runs.
	cobra.OnInitialize(initLogLevel)
}

func initLogLevel() {
	logrus.SetLevel(logrus.Level(logLevel))
}
