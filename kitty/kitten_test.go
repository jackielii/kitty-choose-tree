package kitty

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKitten(t *testing.T) {
	d := scriptData{
		Script: `
print("hello world")
print("hello world")
`,
		SendResult: true,
	}
	s, err := renderPyToFile(d)
	require.NoError(t, err)
	c, err := os.ReadFile(s)
	require.NoError(t, err)
	fmt.Println(string(c))
}

func TestRunKitten(t *testing.T) {
	expected := "foo bar hello world"
	script := fmt.Sprintf(`
os_window_id = last_focused_os_window_id()
set_os_window_title(os_window_id, "%s")
answer = get_os_window_title(os_window_id)
	`, expected)
	result, err := RunKitten(script)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestRunKittenFail(t *testing.T) {
	script := `print(xxx)`
	result, err := RunKitten(script)
	assert.Empty(t, result)
	assert.Contains(t, err.Error(), "NameError")
}

func TestOSWindowTitles(t *testing.T) {
	s, err := osWindowTitles()
	require.NoError(t, err)
	_ = s
}
