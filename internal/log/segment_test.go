package log

import (
	"io/ioutil"
	"os"
	"testing"

	api "github.com/DoppleDankster/golog/api/v1"
	"github.com/stretchr/testify/require"
)

func TestSegment(t *testing.T) {
	dir, _ := ioutil.TempDir("", "segment-test")
	defer os.Remove(dir)
	want := &api.Record{Value: []byte("hello world")}

	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = entWidth * 3
	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
}
