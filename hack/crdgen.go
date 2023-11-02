package main

import (
	"github.com/rancher/wrangler/pkg/crd"
	_ "github.com/spectrocloud/system-upgrade-controller/pkg/generated/controllers/upgrade.cattle.io/v1"
	"github.com/spectrocloud/system-upgrade-controller/pkg/upgrade/plan"
	"os"
)

func main() {
	planCrd, err := plan.CRD()
	if err != nil {
		print(err)
		return
	}
	crd.Print(os.Stdout, []crd.CRD{*planCrd})
}
