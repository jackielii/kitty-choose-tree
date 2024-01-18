package kitty

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"

	"golang.org/x/sync/errgroup"
)

func RunKitten(body string) (string, error) {
	f, err := os.CreateTemp("", "kitten_*.sock")
	if err != nil {
		return "", err
	}
	f.Close()
	os.Remove(f.Name())
	socket := f.Name() + ".sock"
	defer os.Remove(socket)

	g := &errgroup.Group{}
	resultCh := make(chan string)
	sendResult := strings.Contains(body, "answer") // TODO: better way to detect this
	if sendResult {
		g.Go(func() error {
			defer close(resultCh)

			listener, err := net.Listen("unix", socket)
			if err != nil {
				return err
			}
			defer listener.Close()

			conn, err := listener.Accept()
			if err != nil {
				return err
			}
			defer conn.Close()
			buffer := make([]byte, 10)
			res := &bytes.Buffer{}
			for {
				n, err := conn.Read(buffer)
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}
				res.Write(buffer[:n])
			}

			resultCh <- res.String()
			return nil
		})
	} else {
		close(resultCh)
	}

	script, err := renderPyToFile(scriptData{
		Script:     body,
		SendResult: sendResult,
		Socket:     socket,
	})
	if err != nil {
		return "", err
	}
	defer os.Remove(script)

	out, err := exec.Command("kitten", "@", "kitten", script).CombinedOutput()
	if err != nil {
		err = fmt.Errorf("failed to run kitten:\n%s\nerr: %w, ", string(out), err)
		return "", err
	}

	result := <-resultCh
	err = g.Wait()
	if err != nil {
		return "", err
	}

	return result, nil
}

//go:embed kitten.py
var kittenPy string

type scriptData struct {
	Script     string
	SendResult bool
	Socket     string
}

func renderPyToFile(data scriptData) (string, error) {
	dir := ""
	if runtime.GOOS == "darwin" {
		dir = "/tmp" // default doesn't allow read..
	}
	f, err := os.CreateTemp(dir, "kitten_*.py")
	if err != nil {
		return "", err
	}
	defer f.Close()
	funcMap := template.FuncMap{
		"indent": func(s string) string {
			b := &strings.Builder{}
			for _, line := range strings.Split(s, "\n") {
				if strings.TrimSpace(line) != "" {
					b.WriteString(strings.Repeat(" ", 4) + line)
				}
				b.WriteRune('\n')
			}
			return b.String()
		},
	}
	err = template.Must(template.New("kitten").Funcs(funcMap).Parse(kittenPy)).Execute(f, data)
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}
