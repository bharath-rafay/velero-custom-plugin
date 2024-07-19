package main

import (
	"fmt"
	"os"

	"github.com/replicatedhq/local-volume-provider/pkg/plugin"
	"github.com/replicatedhq/local-volume-provider/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	veleroplugin "github.com/vmware-tanzu/velero/pkg/plugin/framework"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version.Get())
		os.Exit(0)
	}

	veleroplugin.NewServer().
		BindFlags(pflag.CommandLine).
		RegisterObjectStore("neridio.io/hostpath", newHostPathObjectStorePlugin).
		RegisterObjectStore("neridio.io/nfs", newNFSObjectStorePlugin).
		RegisterObjectStore("neridio.io/pvc", newPVCObjectStorePlugin).
		Serve()
}

func newHostPathObjectStorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return plugin.NewLocalVolumeObjectStore(logger, plugin.Hostpath), nil
}

func newNFSObjectStorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return plugin.NewLocalVolumeObjectStore(logger, plugin.NFS), nil
}

func newPVCObjectStorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return plugin.NewLocalVolumeObjectStore(logger, plugin.PVC), nil
}
