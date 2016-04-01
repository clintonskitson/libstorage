/*
Package libstorage provides a vendor agnostic storage orchestration model, API,
and reference client and server implementations. libStorage enables storage
consumption by leveraging methods commonly available, locally and/or externally,
to an operating system (OS).

The Past

The libStorage project and its architecture represents a culmination of
experience gained from the project authors' building of
several (http://bit.ly/1HIAet6) different storage (http://bit.ly/1Ya9Uft)
orchestration tools (https://github.com/emccode/rexray). While created using
different languages and targeting disparate storage platforms, all the tools
were architecturally aligned and embedded functionality directly inside the
tools and affected storage platforms.

This shared design goal enabled tools that natively consumed storage, sans
external dependencies.

The Present

Today libStorage focuses on adding value to container runtimes and storage
orchestration tools such as Docker and Mesos, however the libStorage
framework is available abstractly for more general usage across:

  * Operating systems
  * Storage platforms
  * Hardware platforms
  * Virtualization platforms

The client side implementation, focused on operating system activities,
has a minimal set of dependencies in order to avoid a large, runtime footprint.
*/
package libstorage

import (
	"io"

	"github.com/akutz/gofig"

	"github.com/emccode/libstorage/api/client"
	"github.com/emccode/libstorage/api/registry"
	"github.com/emccode/libstorage/api/server"
	"github.com/emccode/libstorage/api/types/context"
	"github.com/emccode/libstorage/api/types/drivers"
)

func init() {
	registerGofigDefaults()
}

// RegisterStorageDriver registers a new StorageDriver with the libStorage
// service.
func RegisterStorageDriver(name string, ctor drivers.NewStorageDriver) {
	registry.RegisterStorageDriver(name, ctor)
}

// RegisterOSDriver registers a new StorageDriver with the libStorage
// service.
func RegisterOSDriver(name string, ctor drivers.NewOSDriver) {
	registry.RegisterOSDriver(name, ctor)
}

// RegisterIntegrationDriver registers a new IntegrationDriver with the
// libStorage service.
func RegisterIntegrationDriver(name string, ctor drivers.NewIntegrationDriver) {
	registry.RegisterIntegrationDriver(name, ctor)
}

/*
Serve starts the reference implementation of a server hosting an
HTTP/JSON service that implements the libStorage API endpoint.

If the config parameter is nil a default instance is created. The
libStorage service is served at the address specified by the configuration
property libstorage.host.
*/
func Serve(config gofig.Config) (io.Closer, <-chan error) {
	return server.Serve(config)
}

// Dial opens a connection to a remote libStorage serice and returns the client
// that can be used to communicate with said endpoint.
//
// If the config parameter is nil a default instance is created. The
// function dials the libStorage service specified by the configuration
// property libstorage.host.
func Dial(ctx context.Context, config gofig.Config) (client.Client, error) {
	return client.Dial(ctx, config)
}

func registerGofigDefaults() {
	r := gofig.NewRegistration("libStorage")
	r.Key(gofig.String, "", "", "", "libstorage.host")
	r.Key(gofig.String, "", "", "", "libstorage.service")
	r.Key(gofig.String, "", "", "", "libstorage.driver")
	r.Key(gofig.Bool, "", false, "", "libstorage.profiles.enabled")
	r.Key(gofig.Bool, "", false, "", "libstorage.profiles.client")
	r.Key(gofig.String, "", "local=127.0.0.1", "", "libstorage.profiles.groups")

	r.Key(gofig.String, "",
		"/proc/partitions", "", "libstorage.client.localdevicesfile")

	r.Key(gofig.String, "", "/usr/local/bin", "", "libstorage.client.tooldir")
	r.Key(gofig.Bool, "", false, "", "libstorage.client.http.logging.enabled")
	r.Key(gofig.String, "", "", "", "libstorage.client.http.logging.out")
	r.Key(gofig.String, "", "", "", "libstorage.client.http.logging.err")
	r.Key(gofig.Bool, "",
		false, "", "libstorage.client.http.logging.logrequest")
	r.Key(gofig.Bool, "",
		false, "", "libstorage.client.http.logging.logresponse")

	r.Key(gofig.Int, "", 60, "", "libstorage.server.readtimeout")
	r.Key(gofig.Int, "", 60, "", "libstorage.server.writetimeout")
	r.Key(gofig.Bool, "", false, "", "libstorage.server.http.logging.enabled")
	r.Key(gofig.String, "", "", "", "libstorage.server.http.logging.out")
	r.Key(gofig.String, "", "", "", "libstorage.server.http.logging.err")

	r.Key(gofig.Bool, "",
		false, "", "libstorage.server.http.logging.logrequest")
	r.Key(gofig.Bool, "",
		false, "", "libstorage.server.http.logging.logresponse")

	gofig.Register(r)
}