package functions

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

type FunctionsImpl struct {
	shell *shell.Shell
	cache map[string]*Function
}

func New(shell *shell.Shell) FunctionInterface {
	impl := FunctionsImpl{
		shell: shell,
		cache: make(map[string]*Function),
	}

	return impl
}

func (fi FunctionsImpl) Store(f *Function) (string, error) {
	b, err := f.Marshal()
	if err != nil {
		return "", err
	}
	cid, err := fi.shell.Add(bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	fi.cache[cid] = f

	return cid, nil
}

func (fi FunctionsImpl) GetAndExecute(path string, writer *strings.Builder) error {
	var f Function
	if fi.cache[path] == nil {
		err := fi.shell.Get(path, os.TempDir()+path)
		if err != nil {
			return err
		}
		defer os.Remove(os.TempDir() + path)
		data, err := ioutil.ReadFile(os.TempDir() + path)
		if err != nil {
			return err
		}

		f = Function{}

		err = f.Unmarshal(data)
		if err != nil {
			return err
		}

		fi.cache[path] = &f
	} else {
		f = *fi.cache[path]
	}

	return fi.Execute(&f, writer)
}

func (fi FunctionsImpl) Execute(function *Function, writer *strings.Builder) error {
	ts := fmt.Sprint(time.Now().Unix())
	b := make([]byte, function.file.Len())

	_, err := function.file.Read(b)
	if err != nil {
		return err
	}
	defer os.Remove(os.TempDir() + "/" + ts)
	err = ioutil.WriteFile(os.TempDir()+"/"+ts, b, 0777)
	if err != nil {
		return err
	}

	out, err := exec.Command(os.TempDir() + ts).Output()
	if err != nil {
		return err
	}

	writer.WriteString(string(out))
	return nil
}