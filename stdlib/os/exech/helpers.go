package exech

import (
	"os/exec"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CommandKey = "Command"

	CommandContextKey = "CommandContext"

	LookPathKey = "LookPath"
)

func New() hctx.Map {
	return hctx.Map{

		CommandKey: Command,

		CommandContextKey: CommandContext,

		LookPathKey: LookPath,
	}
}

// Command returns the Cmd struct to execute the named program with
// the given arguments.
//
// It sets only the Path and Args in the returned structure.
//
// If name contains no path separators, Command uses LookPath to
// resolve name to a complete path if possible. Otherwise it uses name
// directly as Path.
//
// The returned Cmd&#39;s Args field is constructed from the command name
// followed by the elements of arg, so arg should not include the
// command name itself. For example, Command(&#34;echo&#34;, &#34;hello&#34;).
// Args[0] is always name, not the possibly resolved Path.
var Command = exec.Command

// CommandContext is like Command but includes a context.
//
// The provided context is used to kill the process (by calling
// os.Process.Kill) if the context becomes done before the command
// completes on its own.
var CommandContext = exec.CommandContext

// LookPath searches for an executable binary named file
// in the directories named by the path environment variable.
// If file begins with &#34;/&#34;, &#34;#&#34;, &#34;./&#34;, or &#34;../&#34;, it is tried
// directly and the path is not consulted.
// The result may be an absolute path or a path relative to the current directory.
var LookPath = exec.LookPath
