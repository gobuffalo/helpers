package codehosth

import (
	"cmd/go/internal/modfetch/codehost"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AllHexKey = "AllHex"

	GitRepoKey = "GitRepo"

	LocalGitRepoKey = "LocalGitRepo"

	NewRepoKey = "NewRepo"

	RunKey = "Run"

	RunWithStdinKey = "RunWithStdin"

	ShortenSHA1Key = "ShortenSHA1"

	WorkDirKey = "WorkDir"
)

func New() hctx.Map {
	return hctx.Map{

		AllHexKey: AllHex,

		GitRepoKey: GitRepo,

		LocalGitRepoKey: LocalGitRepo,

		NewRepoKey: NewRepo,

		RunKey: Run,

		RunWithStdinKey: RunWithStdin,

		ShortenSHA1Key: ShortenSHA1,

		WorkDirKey: WorkDir,
	}
}

// AllHex reports whether the revision rev is entirely lower-case hexadecimal digits.
var AllHex = codehost.AllHex

// GitRepo returns the code repository at the given Git remote reference.
var GitRepo = codehost.GitRepo

// LocalGitRepo is like Repo but accepts both Git remote references
// and paths to repositories on the local file system.
var LocalGitRepo = codehost.LocalGitRepo

var NewRepo = codehost.NewRepo

// Run runs the command line in the given directory
// (an empty dir means the current directory).
// It returns the standard output and, for a non-zero exit,
// a *RunError indicating the command, exit status, and standard error.
// Standard error is unavailable for commands that exit successfully.
var Run = codehost.Run

var RunWithStdin = codehost.RunWithStdin

// ShortenSHA1 shortens a SHA1 hash (40 hex digits) to the canonical length
// used in pseudo-versions (12 hex digits).
var ShortenSHA1 = codehost.ShortenSHA1

// WorkDir returns the name of the cached work directory to use for the
// given repository type and name.
var WorkDir = codehost.WorkDir
