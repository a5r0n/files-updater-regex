package updater

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegexUpdater(t *testing.T) {
	require := require.New(t)

	updater := &Updater{}
	opts := make(map[string]string)
	opts["pattern"] = `(?m)v[0-9]+\.[0-9]+\.[0-9]+[^ ]*`

	err := updater.Init(opts)
	require.NoError(err)

	verStr := "v1.2.3"
	versionPath := "../../test/.version"

	err = updater.Apply(versionPath, verStr)
	require.NoError(err)

	f, err := os.OpenFile(versionPath, os.O_RDONLY, 0)
	require.NoError(err)
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	require.NoError(err)
	require.Equal(`appVersion: "`+verStr+`"`, string(b))
}
