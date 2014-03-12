package cmd

import (
	"fmt"
	"os"

	"github.com/coreos/etcd-ca/third_party/github.com/codegangsta/cli"

	"github.com/coreos/etcd-ca/depot"
)

func NewChainCommand() cli.Command {
	return cli.Command{
		Name:        "chain",
		Usage:       "Export certificate chain",
		Description: "Export the certificate chain for host1. With no args it exports this CA's certificate.",
		Action:      newChainAction,
	}
}

func newChainAction(c *cli.Context) {
	crt, err := depot.GetCertificateAuthority(d)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Got CA certificate error:", err)
		os.Exit(1)
	}
	// Should not fail if creating from depot
	crtBytes, _ := crt.Export()

	if len(c.Args()) == 0 {
		fmt.Printf("%s", crtBytes)
		return
	}
	name := c.Args()[0]

	crtHost, err := depot.GetCertificateHost(d, name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Got certificate error:", err)
		os.Exit(1)
	}
	crtHostBytes, _ := crtHost.Export()

	if err = crt.VerifyHost(crtHost, name); err != nil {
		fmt.Fprintln(os.Stderr, "Failed verifying certificate chain:", err)
		os.Exit(1)
	}

	fmt.Printf("%s%s", crtBytes, crtHostBytes)
}