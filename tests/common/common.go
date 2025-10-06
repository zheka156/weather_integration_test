package common

import (
	"integration-tests/pkg/di"
	"testing"
	"time"

	"github.com/onsi/gomega"
)

func Before(t *testing.T) (*di.Components, gomega.Gomega) {
	c := di.InitComponents()
	g := SetGomegaParameters(t)
	return c, g
}

func SetGomegaParameters(t *testing.T) gomega.Gomega {
	g := gomega.NewWithT(t)
	g.SetDefaultEventuallyTimeout(time.Second * 15)
	g.SetDefaultEventuallyPollingInterval(time.Second * 300)
	return g
}
