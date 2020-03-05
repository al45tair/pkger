package mem

import (
	"os"
	"testing"

	"github.com/al45tair/pkger/pkging/pkgtest"
	"github.com/stretchr/testify/require"
)

func Test_Pkger_Remove(t *testing.T) {
	r := require.New(t)

	ref, err := pkgtest.NewRef()
	r.NoError(err)
	defer os.RemoveAll(ref.Dir)

	pkg, err := New(ref.Info)
	r.NoError(err)

	pkgtest.RemoveTest(t, ref, pkg)
}
