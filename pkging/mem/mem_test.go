package mem_test

import (
	"testing"

	"github.com/al45tair/pkger/pkging"
	"github.com/al45tair/pkger/pkging/mem"
	"github.com/al45tair/pkger/pkging/pkgtest"
)

func Test_Pkger(t *testing.T) {
	pkgtest.All(t, func(ref *pkgtest.Ref) (pkging.Pkger, error) {
		return mem.New(ref.Info)
	})
}
