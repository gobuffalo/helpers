package osh

import (
	"os"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ChdirKey = "Chdir"

	ChmodKey = "Chmod"

	ChownKey = "Chown"

	ChtimesKey = "Chtimes"

	ClearenvKey = "Clearenv"

	CreateKey = "Create"

	EnvironKey = "Environ"

	ErrNoDeadlineKey = "ErrNoDeadline"

	ExecutableKey = "Executable"

	ExitKey = "Exit"

	ExpandKey = "Expand"

	ExpandEnvKey = "ExpandEnv"

	FindProcessKey = "FindProcess"

	GetegidKey = "Getegid"

	GetenvKey = "Getenv"

	GeteuidKey = "Geteuid"

	GetgidKey = "Getgid"

	GetgroupsKey = "Getgroups"

	GetpagesizeKey = "Getpagesize"

	GetpidKey = "Getpid"

	GetppidKey = "Getppid"

	GetuidKey = "Getuid"

	GetwdKey = "Getwd"

	HostnameKey = "Hostname"

	InterruptKey = "Interrupt"

	IsExistKey = "IsExist"

	IsNotExistKey = "IsNotExist"

	IsPathSeparatorKey = "IsPathSeparator"

	IsPermissionKey = "IsPermission"

	IsTimeoutKey = "IsTimeout"

	KillKey = "Kill"

	LchownKey = "Lchown"

	LinkKey = "Link"

	LookupEnvKey = "LookupEnv"

	LstatKey = "Lstat"

	MkdirKey = "Mkdir"

	MkdirAllKey = "MkdirAll"

	NewFileKey = "NewFile"

	NewSyscallErrorKey = "NewSyscallError"

	O_APPENDKey = "O_APPEND"

	O_CREATEKey = "O_CREATE"

	O_EXCLKey = "O_EXCL"

	O_RDONLYKey = "O_RDONLY"

	O_RDWRKey = "O_RDWR"

	O_SYNCKey = "O_SYNC"

	O_TRUNCKey = "O_TRUNC"

	O_WRONLYKey = "O_WRONLY"

	OpenKey = "Open"

	OpenFileKey = "OpenFile"

	PipeKey = "Pipe"

	ReadlinkKey = "Readlink"

	RemoveKey = "Remove"

	RemoveAllKey = "RemoveAll"

	RenameKey = "Rename"

	SameFileKey = "SameFile"

	SetenvKey = "Setenv"

	StartProcessKey = "StartProcess"

	StatKey = "Stat"

	SymlinkKey = "Symlink"

	TempDirKey = "TempDir"

	TruncateKey = "Truncate"

	UnsetenvKey = "Unsetenv"
)

func New() hctx.Map {
	return hctx.Map{

		ChdirKey: Chdir,

		ChmodKey: Chmod,

		ChownKey: Chown,

		ChtimesKey: Chtimes,

		ClearenvKey: Clearenv,

		CreateKey: Create,

		EnvironKey: Environ,

		ErrNoDeadlineKey: ErrNoDeadline,

		ExecutableKey: Executable,

		ExitKey: Exit,

		ExpandKey: Expand,

		ExpandEnvKey: ExpandEnv,

		FindProcessKey: FindProcess,

		GetegidKey: Getegid,

		GetenvKey: Getenv,

		GeteuidKey: Geteuid,

		GetgidKey: Getgid,

		GetgroupsKey: Getgroups,

		GetpagesizeKey: Getpagesize,

		GetpidKey: Getpid,

		GetppidKey: Getppid,

		GetuidKey: Getuid,

		GetwdKey: Getwd,

		HostnameKey: Hostname,

		InterruptKey: Interrupt,

		IsExistKey: IsExist,

		IsNotExistKey: IsNotExist,

		IsPathSeparatorKey: IsPathSeparator,

		IsPermissionKey: IsPermission,

		IsTimeoutKey: IsTimeout,

		KillKey: Kill,

		LchownKey: Lchown,

		LinkKey: Link,

		LookupEnvKey: LookupEnv,

		LstatKey: Lstat,

		MkdirKey: Mkdir,

		MkdirAllKey: MkdirAll,

		NewFileKey: NewFile,

		NewSyscallErrorKey: NewSyscallError,

		O_APPENDKey: O_APPEND,

		O_CREATEKey: O_CREATE,

		O_EXCLKey: O_EXCL,

		O_RDONLYKey: O_RDONLY,

		O_RDWRKey: O_RDWR,

		O_SYNCKey: O_SYNC,

		O_TRUNCKey: O_TRUNC,

		O_WRONLYKey: O_WRONLY,

		OpenKey: Open,

		OpenFileKey: OpenFile,

		PipeKey: Pipe,

		ReadlinkKey: Readlink,

		RemoveKey: Remove,

		RemoveAllKey: RemoveAll,

		RenameKey: Rename,

		SameFileKey: SameFile,

		SetenvKey: Setenv,

		StartProcessKey: StartProcess,

		StatKey: Stat,

		SymlinkKey: Symlink,

		TempDirKey: TempDir,

		TruncateKey: Truncate,

		UnsetenvKey: Unsetenv,
	}
}

// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
var Chdir = os.Chdir

// Chmod changes the mode of the named file to mode.
// If the file is a symbolic link, it changes the mode of the link&#39;s target.
// If there is an error, it will be of type *PathError.
//
// A different subset of the mode bits are used, depending on the
// operating system.
//
// On Unix, the mode&#39;s permission bits, ModeSetuid, ModeSetgid, and
// ModeSticky are used.
//
// On Windows, the mode must be non-zero but otherwise only the 0200
// bit (owner writable) of mode is used; it controls whether the
// file&#39;s read-only attribute is set or cleared. attribute. The other
// bits are currently unused. Use mode 0400 for a read-only file and
// 0600 for a readable+writable file.
//
// On Plan 9, the mode&#39;s permission bits, ModeAppend, ModeExclusive,
// and ModeTemporary are used.
var Chmod = os.Chmod

// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link&#39;s target.
// If there is an error, it will be of type *PathError.
var Chown = os.Chown

// Chtimes changes the access and modification times of the named
// file, similar to the Unix utime() or utimes() functions.
//
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
// If there is an error, it will be of type *PathError.
var Chtimes = os.Chtimes

// Clearenv deletes all environment variables.
var Clearenv = os.Clearenv

// Create creates the named file with mode 0666 (before umask), truncating
// it if it already exists. If successful, methods on the returned
// File can be used for I/O; the associated file descriptor has mode
// O_RDWR.
// If there is an error, it will be of type *PathError.
var Create = os.Create

// Environ returns a copy of strings representing the environment,
// in the form &#34;key=value&#34;.
var Environ = os.Environ

var ErrNoDeadline = os.ErrNoDeadline

// Executable returns the path name for the executable that started
// the current process. There is no guarantee that the path is still
// pointing to the correct executable. If a symlink was used to start
// the process, depending on the operating system, the result might
// be the symlink or the path it pointed to. If a stable result is
// needed, path/filepath.EvalSymlinks might help.
//
// Executable returns an absolute path unless an error occurred.
//
// The main use case is finding resources located relative to an
// executable.
//
// Executable is not supported on nacl.
var Executable = os.Executable

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
var Exit = os.Exit

// Expand replaces ${var} or $var in the string based on the mapping function.
// For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).
var Expand = os.Expand

// ExpandEnv replaces ${var} or $var in the string according to the values
// of the current environment variables. References to undefined
// variables are replaced by the empty string.
var ExpandEnv = os.ExpandEnv

// FindProcess looks for a running process by its pid.
//
// The Process it returns can be used to obtain information
// about the underlying operating system process.
//
// On Unix systems, FindProcess always succeeds and returns a Process
// for the given pid, regardless of whether the process exists.
var FindProcess = os.FindProcess

// Getegid returns the numeric effective group id of the caller.
//
// On Windows, it returns -1.
var Getegid = os.Getegid

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// To distinguish between an empty value and an unset value, use LookupEnv.
var Getenv = os.Getenv

// Geteuid returns the numeric effective user id of the caller.
//
// On Windows, it returns -1.
var Geteuid = os.Geteuid

// Getgid returns the numeric group id of the caller.
//
// On Windows, it returns -1.
var Getgid = os.Getgid

// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
//
// On Windows, it returns syscall.EWINDOWS. See the os/user package
// for a possible alternative.
var Getgroups = os.Getgroups

// Getpagesize returns the underlying system&#39;s memory page size.
var Getpagesize = os.Getpagesize

// Getpid returns the process id of the caller.
var Getpid = os.Getpid

// Getppid returns the process id of the caller&#39;s parent.
var Getppid = os.Getppid

// Getuid returns the numeric user id of the caller.
//
// On Windows, it returns -1.
var Getuid = os.Getuid

// Getwd returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
var Getwd = os.Getwd

// Hostname returns the host name reported by the kernel.
var Hostname = os.Hostname

var Interrupt = os.Interrupt

// IsExist returns a boolean indicating whether the error is known to report
// that a file or directory already exists. It is satisfied by ErrExist as
// well as some syscall errors.
var IsExist = os.IsExist

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
var IsNotExist = os.IsNotExist

// IsPathSeparator reports whether c is a directory separator character.
var IsPathSeparator = os.IsPathSeparator

// IsPermission returns a boolean indicating whether the error is known to
// report that permission is denied. It is satisfied by ErrPermission as well
// as some syscall errors.
var IsPermission = os.IsPermission

// IsTimeout returns a boolean indicating whether the error is known
// to report that a timeout occurred.
var IsTimeout = os.IsTimeout

var Kill = os.Kill

// Lchown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
// If there is an error, it will be of type *PathError.
var Lchown = os.Lchown

// Link creates newname as a hard link to the oldname file.
// If there is an error, it will be of type *LinkError.
var Link = os.Link

// LookupEnv retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
var LookupEnv = os.LookupEnv

// Lstat returns a FileInfo describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link. Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
var Lstat = os.Lstat

// Mkdir creates a new directory with the specified name and permission
// bits (before umask).
// If there is an error, it will be of type *PathError.
var Mkdir = os.Mkdir

// MkdirAll creates a directory named path,
// along with any necessary parents, and returns nil,
// or else returns an error.
// The permission bits perm (before umask) are used for all
// directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing
// and returns nil.
var MkdirAll = os.MkdirAll

// NewFile returns a new File with the given file descriptor and
// name. The returned value will be nil if fd is not a valid file
// descriptor.
var NewFile = os.NewFile

// NewSyscallError returns, as an error, a new SyscallError
// with the given system call name and error details.
// As a convenience, if err is nil, NewSyscallError returns nil.
var NewSyscallError = os.NewSyscallError

// The remaining values may be or&#39;ed in to control behavior.
var O_APPEND = os.O_APPEND

var O_CREATE = os.O_CREATE

var O_EXCL = os.O_EXCL

// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
var O_RDONLY = os.O_RDONLY

var O_RDWR = os.O_RDWR

var O_SYNC = os.O_SYNC

var O_TRUNC = os.O_TRUNC

var O_WRONLY = os.O_WRONLY

// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
var Open = os.Open

// OpenFile is the generalized open call; most users will use Open
// or Create instead. It opens the named file with specified flag
// (O_RDONLY etc.) and perm (before umask), if applicable. If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type *PathError.
var OpenFile = os.OpenFile

// Pipe returns a connected pair of Files; reads from r return bytes
// written to w. It returns the files and an error, if any.
var Pipe = os.Pipe

// Readlink returns the destination of the named symbolic link.
// If there is an error, it will be of type *PathError.
var Readlink = os.Readlink

// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
var Remove = os.Remove

// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first error
// it encounters. If the path does not exist, RemoveAll
// returns nil (no error).
var RemoveAll = os.RemoveAll

// Rename renames (moves) oldpath to newpath.
// If newpath already exists and is not a directory, Rename replaces it.
// OS-specific restrictions may apply when oldpath and newpath are in different directories.
// If there is an error, it will be of type *LinkError.
var Rename = os.Rename

// SameFile reports whether fi1 and fi2 describe the same file.
// For example, on Unix this means that the device and inode fields
// of the two underlying structures are identical; on other systems
// the decision may be based on the path names.
// SameFile only applies to results returned by this package&#39;s Stat.
// It returns false in other cases.
var SameFile = os.SameFile

// Setenv sets the value of the environment variable named by the key.
// It returns an error, if any.
var Setenv = os.Setenv

// StartProcess starts a new process with the program, arguments and attributes
// specified by name, argv and attr. The argv slice will become os.Args in the
// new process, so it normally starts with the program name.
//
// If the calling goroutine has locked the operating system thread
// with runtime.LockOSThread and modified any inheritable OS-level
// thread state (for example, Linux or Plan 9 name spaces), the new
// process will inherit the caller&#39;s thread state.
//
// StartProcess is a low-level interface. The os/exec package provides
// higher-level interfaces.
//
// If there is an error, it will be of type *PathError.
var StartProcess = os.StartProcess

// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
var Stat = os.Stat

// Symlink creates newname as a symbolic link to oldname.
// If there is an error, it will be of type *LinkError.
var Symlink = os.Symlink

// TempDir returns the default directory to use for temporary files.
//
// On Unix systems, it returns $TMPDIR if non-empty, else /tmp.
// On Windows, it uses GetTempPath, returning the first non-empty
// value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory.
// On Plan 9, it returns /tmp.
//
// The directory is neither guaranteed to exist nor have accessible
// permissions.
var TempDir = os.TempDir

// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link&#39;s target.
// If there is an error, it will be of type *PathError.
var Truncate = os.Truncate

// Unsetenv unsets a single environment variable.
var Unsetenv = os.Unsetenv
