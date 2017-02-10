package ctx

import (
	"github.com/docker/docker/cliconfig"
	"github.com/docker/docker/cliconfig/configfile"
	"github.com/docker/libcompose/docker/auth"
	"github.com/docker/libcompose/docker/client"
	"github.com/rancher/rancher-compose/project"
)

// Context holds context meta information about a libcompose project and docker
// client information (like configuration file, builder to use, …)
type Context struct {
	project.Context
	ClientFactory client.Factory
	ConfigDir     string
	ConfigFile    *configfile.ConfigFile
	AuthLookup    auth.Lookup
}

// LookupConfig tries to load the docker configuration files, if any.
func (c *Context) LookupConfig() error {
	if c.ConfigFile != nil {
		return nil
	}

	config, err := cliconfig.Load(c.ConfigDir)
	if err != nil {
		return err
	}

	c.ConfigFile = config

	return nil
}