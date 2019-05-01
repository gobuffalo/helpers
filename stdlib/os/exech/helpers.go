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
//
// On Windows, processes receive the whole command line as a single string
// and do their own parsing. Command combines and quotes Args into a command
// line string with an algorithm compatible with applications using
// CommandLineToArgvW (which is the most common way). Notable exceptions are
// msiexec.exe and cmd.exe (and thus, all batch files), which have a different
// unquoting algorithm. In these or other similar cases, you can do the
// quoting yourself and provide the full command line in SysProcAttr.CmdLine,
// leaving Args empty.
var Command = exec.Command

// CommandContext is like Command but includes a context.
//
// The provided context is used to kill the process (by calling
// os.Process.Kill) if the context becomes done before the command
// completes on its own.
var CommandContext = exec.CommandContext

// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
var LookPath = exec.LookPath
