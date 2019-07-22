package main

import (
	"flag"
	"testing"

	"github.com/yard-turkey/csi-certify/pkg/certify"
	"k8s.io/kubernetes/test/e2e/framework"
)

var customTestDriver string

func init() {
	flag.StringVar(&customTestDriver, "testdriver", "", "the testdriver implementation that you want to run (should be one of the implementations defined in CSITestDrivers)")
	framework.HandleFlags()
	framework.AfterReadingAllFlags(&framework.TestContext)
}

func Test(t *testing.T) {
	certify.Test(t, customTestDriver)
}
