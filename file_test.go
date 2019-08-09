package pkger

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_File_Open(t *testing.T) {
	r := require.New(t)

	f, err := Open("/file_test.go")
	r.NoError(err)

	r.Equal("/file_test.go", f.Name())

	b, err := ioutil.ReadAll(f)
	r.NoError(err)
	r.Contains(string(b), "Test_File_Open")
	r.NoError(f.Close())
}

func Test_File_Open_Dir(t *testing.T) {
	r := require.New(t)

	f, err := Open("/cmd")
	r.NoError(err)

	r.Equal("/cmd", f.Name())

	r.NoError(f.Close())
}

func Test_File_Read_Memory(t *testing.T) {
	r := require.New(t)

	f, err := Open("/file_test.go")
	r.NoError(err)
	f.data = []byte("hi!")

	r.Equal("/file_test.go", f.Name())

	b, err := ioutil.ReadAll(f)
	r.NoError(err)
	r.Equal(string(b), "hi!")
	r.NoError(f.Close())
}

func Test_File_Write(t *testing.T) {
	r := require.New(t)

	f, err := Create("/hello.txt")
	r.NoError(err)
	r.NotNil(f)

	fi, err := f.Stat()
	r.NoError(err)
	r.Zero(fi.Size())

	r.Equal("/hello.txt", fi.Name())

	mt := fi.ModTime()
	r.NotZero(mt)

	sz, err := io.Copy(f, strings.NewReader(radio))
	r.NoError(err)
	r.Equal(int64(1381), sz)

	// because windows can't handle the time precisely
	// enough, we have to *force* just a smidge of time
	// to ensure the two ModTime's are different.
	// i know, i hate it too.
	time.Sleep(time.Millisecond)
	r.NoError(f.Close())
	r.Equal(int64(1381), fi.Size())
	r.NotZero(fi.ModTime())
	r.NotEqual(mt, fi.ModTime())
}

func Test_File_JSON(t *testing.T) {
	r := require.New(t)

	f, err := createFile("/radio.radio")
	r.NoError(err)
	r.NotNil(f)

	bi, err := f.Stat()
	r.NoError(err)

	mj, err := json.Marshal(f)
	r.NoError(err)

	f2 := &File{}

	r.NoError(json.Unmarshal(mj, f2))

	ai, err := f2.Stat()
	r.NoError(err)

	r.Equal(bi.Size(), ai.Size())

	r.Equal(string(f.data), string(f2.data))
}
