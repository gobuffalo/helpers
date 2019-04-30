package syscallh

import (
	"syscall"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AcceptKey = "Accept"

	AcceptKey = "Accept"

	AcceptKey = "Accept"

	AcceptKey = "Accept"

	AcceptKey = "Accept"

	AcceptKey = "Accept"

	AcceptKey = "Accept"

	Accept4Key = "Accept4"

	Accept4Key = "Accept4"

	Accept4Key = "Accept4"

	Accept4Key = "Accept4"

	Accept4Key = "Accept4"

	AcceptExKey = "AcceptEx"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AccessKey = "Access"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimeKey = "Adjtime"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	AttachLsfKey = "AttachLsf"

	AwaitKey = "Await"

	BindKey = "Bind"

	BindKey = "Bind"

	BindKey = "Bind"

	BindKey = "Bind"

	BindKey = "Bind"

	BindToDeviceKey = "BindToDevice"

	BpfBuflenKey = "BpfBuflen"

	BpfBuflenKey = "BpfBuflen"

	BpfDatalinkKey = "BpfDatalink"

	BpfDatalinkKey = "BpfDatalink"

	BpfHeadercmplKey = "BpfHeadercmpl"

	BpfHeadercmplKey = "BpfHeadercmpl"

	BpfInterfaceKey = "BpfInterface"

	BpfInterfaceKey = "BpfInterface"

	BpfJumpKey = "BpfJump"

	BpfJumpKey = "BpfJump"

	BpfStatsKey = "BpfStats"

	BpfStatsKey = "BpfStats"

	BpfStmtKey = "BpfStmt"

	BpfStmtKey = "BpfStmt"

	BpfTimeoutKey = "BpfTimeout"

	BpfTimeoutKey = "BpfTimeout"

	BytePtrFromStringKey = "BytePtrFromString"

	ByteSliceFromStringKey = "ByteSliceFromString"

	CancelIoKey = "CancelIo"

	CancelIoExKey = "CancelIoEx"

	CertAddCertificateContextToStoreKey = "CertAddCertificateContextToStore"

	CertCloseStoreKey = "CertCloseStore"

	CertCreateCertificateContextKey = "CertCreateCertificateContext"

	CertEnumCertificatesInStoreKey = "CertEnumCertificatesInStore"

	CertFreeCertificateChainKey = "CertFreeCertificateChain"

	CertFreeCertificateContextKey = "CertFreeCertificateContext"

	CertGetCertificateChainKey = "CertGetCertificateChain"

	CertOpenStoreKey = "CertOpenStore"

	CertOpenSystemStoreKey = "CertOpenSystemStore"

	CertVerifyCertificateChainPolicyKey = "CertVerifyCertificateChainPolicy"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	CheckBpfVersionKey = "CheckBpfVersion"

	CheckBpfVersionKey = "CheckBpfVersion"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChflagsKey = "Chflags"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChmodKey = "Chmod"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChownKey = "Chown"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ChrootKey = "Chroot"

	ClearenvKey = "Clearenv"

	ClearenvKey = "Clearenv"

	ClearenvKey = "Clearenv"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseKey = "Close"

	CloseHandleKey = "CloseHandle"

	CloseOnExecKey = "CloseOnExec"

	CloseOnExecKey = "CloseOnExec"

	CloseOnExecKey = "CloseOnExec"

	CloseOnExecKey = "CloseOnExec"

	ClosesocketKey = "Closesocket"

	CmsgLenKey = "CmsgLen"

	CmsgSpaceKey = "CmsgSpace"

	CommandLineToArgvKey = "CommandLineToArgv"

	ComputerNameKey = "ComputerName"

	ConnectKey = "Connect"

	ConnectKey = "Connect"

	ConnectKey = "Connect"

	ConnectKey = "Connect"

	ConnectExKey = "ConnectEx"

	ConvertSidToStringSidKey = "ConvertSidToStringSid"

	ConvertStringSidToSidKey = "ConvertStringSidToSid"

	CopySidKey = "CopySid"

	CreatKey = "Creat"

	CreateKey = "Create"

	CreateDirectoryKey = "CreateDirectory"

	CreateFileKey = "CreateFile"

	CreateFileMappingKey = "CreateFileMapping"

	CreateHardLinkKey = "CreateHardLink"

	CreateIoCompletionPortKey = "CreateIoCompletionPort"

	CreatePipeKey = "CreatePipe"

	CreateProcessKey = "CreateProcess"

	CreateProcessAsUserKey = "CreateProcessAsUser"

	CreateSymbolicLinkKey = "CreateSymbolicLink"

	CreateToolhelp32SnapshotKey = "CreateToolhelp32Snapshot"

	CryptAcquireContextKey = "CryptAcquireContext"

	CryptGenRandomKey = "CryptGenRandom"

	CryptReleaseContextKey = "CryptReleaseContext"

	DeleteFileKey = "DeleteFile"

	DetachLsfKey = "DetachLsf"

	DeviceIoControlKey = "DeviceIoControl"

	DnsNameCompareKey = "DnsNameCompare"

	DnsQueryKey = "DnsQuery"

	DnsRecordListFreeKey = "DnsRecordListFree"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	DupKey = "Dup"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup2Key = "Dup2"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

	DuplicateHandleKey = "DuplicateHandle"

	EnvironKey = "Environ"

	EnvironKey = "Environ"

	EnvironKey = "Environ"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreateKey = "EpollCreate"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCreate1Key = "EpollCreate1"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollCtlKey = "EpollCtl"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	EscapeArgKey = "EscapeArg"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExecKey = "Exec"

	ExecKey = "Exec"

	ExecKey = "Exec"

	ExitKey = "Exit"

	ExitProcessKey = "ExitProcess"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FallocateKey = "Fallocate"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchdirKey = "Fchdir"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchflagsKey = "Fchflags"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownKey = "Fchown"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FchownatKey = "Fchownat"

	FcntlFlockKey = "FcntlFlock"

	FcntlFlockKey = "FcntlFlock"

	FcntlFlockKey = "FcntlFlock"

	FcntlFlockKey = "FcntlFlock"

	Fd2pathKey = "Fd2path"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FindCloseKey = "FindClose"

	FindFirstFileKey = "FindFirstFile"

	FindNextFileKey = "FindNextFile"

	FixwdKey = "Fixwd"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlushBpfKey = "FlushBpf"

	FlushBpfKey = "FlushBpf"

	FlushFileBuffersKey = "FlushFileBuffers"

	FlushViewOfFileKey = "FlushViewOfFile"

	ForkExecKey = "ForkExec"

	ForkExecKey = "ForkExec"

	FormatMessageKey = "FormatMessage"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FreeAddrInfoWKey = "FreeAddrInfoW"

	FreeEnvironmentStringsKey = "FreeEnvironmentStrings"

	FreeLibraryKey = "FreeLibrary"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FsyncKey = "Fsync"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FullPathKey = "FullPath"

	FutimesKey = "Futimes"

	FutimesKey = "Futimes"

	FutimesatKey = "Futimesat"

	FwstatKey = "Fwstat"

	FwstatKey = "Fwstat"

	FwstatKey = "Fwstat"

	GetAcceptExSockaddrsKey = "GetAcceptExSockaddrs"

	GetAdaptersInfoKey = "GetAdaptersInfo"

	GetAddrInfoWKey = "GetAddrInfoW"

	GetCommandLineKey = "GetCommandLine"

	GetComputerNameKey = "GetComputerName"

	GetConsoleModeKey = "GetConsoleMode"

	GetCurrentDirectoryKey = "GetCurrentDirectory"

	GetCurrentProcessKey = "GetCurrentProcess"

	GetEnvironmentStringsKey = "GetEnvironmentStrings"

	GetEnvironmentVariableKey = "GetEnvironmentVariable"

	GetExitCodeProcessKey = "GetExitCodeProcess"

	GetFileAttributesKey = "GetFileAttributes"

	GetFileAttributesExKey = "GetFileAttributesEx"

	GetFileInformationByHandleKey = "GetFileInformationByHandle"

	GetFileTypeKey = "GetFileType"

	GetFullPathNameKey = "GetFullPathName"

	GetHostByNameKey = "GetHostByName"

	GetIfEntryKey = "GetIfEntry"

	GetLastErrorKey = "GetLastError"

	GetLengthSidKey = "GetLengthSid"

	GetLongPathNameKey = "GetLongPathName"

	GetProcAddressKey = "GetProcAddress"

	GetProcessTimesKey = "GetProcessTimes"

	GetProtoByNameKey = "GetProtoByName"

	GetQueuedCompletionStatusKey = "GetQueuedCompletionStatus"

	GetServByNameKey = "GetServByName"

	GetShortPathNameKey = "GetShortPathName"

	GetStartupInfoKey = "GetStartupInfo"

	GetStdHandleKey = "GetStdHandle"

	GetSystemTimeAsFileTimeKey = "GetSystemTimeAsFileTime"

	GetTempPathKey = "GetTempPath"

	GetTimeZoneInformationKey = "GetTimeZoneInformation"

	GetTokenInformationKey = "GetTokenInformation"

	GetUserNameExKey = "GetUserNameEx"

	GetUserProfileDirectoryKey = "GetUserProfileDirectory"

	GetVersionKey = "GetVersion"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetcwdKey = "Getcwd"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdentsKey = "Getdents"

	GetdirentriesKey = "Getdirentries"

	GetdirentriesKey = "Getdirentries"

	GetdirentriesKey = "Getdirentries"

	GetdirentriesKey = "Getdirentries"

	GetdirentriesKey = "Getdirentries"

	GetdirentriesKey = "Getdirentries"

	GetdirentriesKey = "Getdirentries"

	GetdirentriesKey = "Getdirentries"

	GetdtablesizeKey = "Getdtablesize"

	GetdtablesizeKey = "Getdtablesize"

	GetdtablesizeKey = "Getdtablesize"

	GetdtablesizeKey = "Getdtablesize"

	GetdtablesizeKey = "Getdtablesize"

	GetdtablesizeKey = "Getdtablesize"

	GetdtablesizeKey = "Getdtablesize"

	GetdtablesizeKey = "Getdtablesize"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetenvKey = "Getenv"

	GetenvKey = "Getenv"

	GetenvKey = "Getenv"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GetexecnameKey = "Getexecname"

	GetfsstatKey = "Getfsstat"

	GetfsstatKey = "Getfsstat"

	GetfsstatKey = "Getfsstat"

	GetfsstatKey = "Getfsstat"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GethostnameKey = "Gethostname"

	GetkerninfoKey = "Getkerninfo"

	GetpagesizeKey = "Getpagesize"

	GetpeernameKey = "Getpeername"

	GetpeernameKey = "Getpeername"

	GetpeernameKey = "Getpeername"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgidKey = "Getpgid"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpgrpKey = "Getpgrp"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetpidKey = "Getpid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetppidKey = "Getppid"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetrusageKey = "Getrusage"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsidKey = "Getsid"

	GetsocknameKey = "Getsockname"

	GetsocknameKey = "Getsockname"

	GetsocknameKey = "Getsockname"

	GetsocknameKey = "Getsockname"

	GetsocknameKey = "Getsockname"

	GetsocknameKey = "Getsockname"

	GetsockoptKey = "Getsockopt"

	GetsockoptByteKey = "GetsockoptByte"

	GetsockoptICMPv6FilterKey = "GetsockoptICMPv6Filter"

	GetsockoptICMPv6FilterKey = "GetsockoptICMPv6Filter"

	GetsockoptIPMreqKey = "GetsockoptIPMreq"

	GetsockoptIPMreqKey = "GetsockoptIPMreq"

	GetsockoptIPMreqnKey = "GetsockoptIPMreqn"

	GetsockoptIPMreqnKey = "GetsockoptIPMreqn"

	GetsockoptIPv6MTUInfoKey = "GetsockoptIPv6MTUInfo"

	GetsockoptIPv6MTUInfoKey = "GetsockoptIPv6MTUInfo"

	GetsockoptIPv6MreqKey = "GetsockoptIPv6Mreq"

	GetsockoptIPv6MreqKey = "GetsockoptIPv6Mreq"

	GetsockoptInet4AddrKey = "GetsockoptInet4Addr"

	GetsockoptInet4AddrKey = "GetsockoptInet4Addr"

	GetsockoptIntKey = "GetsockoptInt"

	GetsockoptIntKey = "GetsockoptInt"

	GetsockoptIntKey = "GetsockoptInt"

	GetsockoptIntKey = "GetsockoptInt"

	GetsockoptUcredKey = "GetsockoptUcred"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettidKey = "Gettid"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GettimeofdayKey = "Gettimeofday"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetuidKey = "Getuid"

	GetwdKey = "Getwd"

	GetwdKey = "Getwd"

	GetwdKey = "Getwd"

	GetwdKey = "Getwd"

	GetwdKey = "Getwd"

	GetwdKey = "Getwd"

	GetwdKey = "Getwd"

	GetwdKey = "Getwd"

	GetwdKey = "Getwd"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyAddWatchKey = "InotifyAddWatch"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInitKey = "InotifyInit"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyInit1Key = "InotifyInit1"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	IopermKey = "Ioperm"

	IopermKey = "Ioperm"

	IopermKey = "Ioperm"

	IopermKey = "Ioperm"

	IopermKey = "Ioperm"

	IopermKey = "Ioperm"

	IopermKey = "Ioperm"

	IoplKey = "Iopl"

	IoplKey = "Iopl"

	IoplKey = "Iopl"

	IoplKey = "Iopl"

	IoplKey = "Iopl"

	IoplKey = "Iopl"

	IoplKey = "Iopl"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	KeventKey = "Kevent"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KillKey = "Kill"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KlogctlKey = "Klogctl"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	KqueueKey = "Kqueue"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	LinkKey = "Link"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListenKey = "Listen"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	LoadCancelIoExKey = "LoadCancelIoEx"

	LoadConnectExKey = "LoadConnectEx"

	LoadCreateSymbolicLinkKey = "LoadCreateSymbolicLink"

	LoadDLLKey = "LoadDLL"

	LoadGetAddrInfoKey = "LoadGetAddrInfo"

	LoadLibraryKey = "LoadLibrary"

	LoadSetFileCompletionNotificationModesKey = "LoadSetFileCompletionNotificationModes"

	LocalFreeKey = "LocalFree"

	LookupAccountNameKey = "LookupAccountName"

	LookupAccountSidKey = "LookupAccountSid"

	LookupSIDKey = "LookupSID"

	LsfJumpKey = "LsfJump"

	LsfSocketKey = "LsfSocket"

	LsfStmtKey = "LsfStmt"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	LstatKey = "Lstat"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MapViewOfFileKey = "MapViewOfFile"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdirKey = "Mkdir"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkdiratKey = "Mkdirat"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodKey = "Mknod"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MknodatKey = "Mknodat"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockKey = "Mlock"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MmapKey = "Mmap"

	MmapKey = "Mmap"

	MmapKey = "Mmap"

	MountKey = "Mount"

	MountKey = "Mount"

	MoveFileKey = "MoveFile"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockKey = "Munlock"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunmapKey = "Munmap"

	MunmapKey = "Munmap"

	MunmapKey = "Munmap"

	MustLoadDLLKey = "MustLoadDLL"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NetApiBufferFreeKey = "NetApiBufferFree"

	NetGetJoinInformationKey = "NetGetJoinInformation"

	NetUserGetInfoKey = "NetUserGetInfo"

	NetlinkRIBKey = "NetlinkRIB"

	NewCallbackKey = "NewCallback"

	NewCallbackCDeclKey = "NewCallbackCDecl"

	NewErrorKey = "NewError"

	NewLazyDLLKey = "NewLazyDLL"

	NsecToFiletimeKey = "NsecToFiletime"

	NsecToTimespecKey = "NsecToTimespec"

	NsecToTimespecKey = "NsecToTimespec"

	NsecToTimevalKey = "NsecToTimeval"

	NsecToTimevalKey = "NsecToTimeval"

	NsecToTimevalKey = "NsecToTimeval"

	NtohsKey = "Ntohs"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenKey = "Open"

	OpenCurrentProcessTokenKey = "OpenCurrentProcessToken"

	OpenProcessKey = "OpenProcess"

	OpenProcessTokenKey = "OpenProcessToken"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	ParseDirentKey = "ParseDirent"

	ParseFilesKey = "ParseFiles"

	ParseNetlinkMessageKey = "ParseNetlinkMessage"

	ParseNetlinkRouteAttrKey = "ParseNetlinkRouteAttr"

	ParseRoutingMessageKey = "ParseRoutingMessage"

	ParseRoutingMessageKey = "ParseRoutingMessage"

	ParseRoutingSockaddrKey = "ParseRoutingSockaddr"

	ParseRoutingSockaddrKey = "ParseRoutingSockaddr"

	ParseSocketControlMessageKey = "ParseSocketControlMessage"

	ParseUnixCredentialsKey = "ParseUnixCredentials"

	ParseUnixRightsKey = "ParseUnixRights"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PathconfKey = "Pathconf"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	PipeKey = "Pipe"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	Pipe2Key = "Pipe2"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PostQueuedCompletionStatusKey = "PostQueuedCompletionStatus"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	Process32FirstKey = "Process32First"

	Process32NextKey = "Process32Next"

	PtraceAttachKey = "PtraceAttach"

	PtraceAttachKey = "PtraceAttach"

	PtraceAttachKey = "PtraceAttach"

	PtraceContKey = "PtraceCont"

	PtraceContKey = "PtraceCont"

	PtraceDetachKey = "PtraceDetach"

	PtraceDetachKey = "PtraceDetach"

	PtraceDetachKey = "PtraceDetach"

	PtraceGetEventMsgKey = "PtraceGetEventMsg"

	PtraceGetRegsKey = "PtraceGetRegs"

	PtracePeekDataKey = "PtracePeekData"

	PtracePeekDataKey = "PtracePeekData"

	PtracePeekTextKey = "PtracePeekText"

	PtracePeekTextKey = "PtracePeekText"

	PtracePokeDataKey = "PtracePokeData"

	PtracePokeDataKey = "PtracePokeData"

	PtracePokeTextKey = "PtracePokeText"

	PtracePokeTextKey = "PtracePokeText"

	PtraceSetOptionsKey = "PtraceSetOptions"

	PtraceSetRegsKey = "PtraceSetRegs"

	PtraceSingleStepKey = "PtraceSingleStep"

	PtraceSingleStepKey = "PtraceSingleStep"

	PtraceSyscallKey = "PtraceSyscall"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	RawSyscallKey = "RawSyscall"

	RawSyscallKey = "RawSyscall"

	RawSyscallKey = "RawSyscall"

	RawSyscallKey = "RawSyscall"

	RawSyscall6Key = "RawSyscall6"

	RawSyscall6Key = "RawSyscall6"

	RawSyscall6Key = "RawSyscall6"

	RawSyscall6Key = "RawSyscall6"

	ReadKey = "Read"

	ReadKey = "Read"

	ReadKey = "Read"

	ReadKey = "Read"

	ReadKey = "Read"

	ReadConsoleKey = "ReadConsole"

	ReadDirectoryChangesKey = "ReadDirectoryChanges"

	ReadDirentKey = "ReadDirent"

	ReadDirentKey = "ReadDirent"

	ReadDirentKey = "ReadDirent"

	ReadDirentKey = "ReadDirent"

	ReadDirentKey = "ReadDirent"

	ReadDirentKey = "ReadDirent"

	ReadFileKey = "ReadFile"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	RebootKey = "Reboot"

	RebootKey = "Reboot"

	RecvfromKey = "Recvfrom"

	RecvfromKey = "Recvfrom"

	RecvfromKey = "Recvfrom"

	RecvfromKey = "Recvfrom"

	RecvmsgKey = "Recvmsg"

	RecvmsgKey = "Recvmsg"

	RecvmsgKey = "Recvmsg"

	RecvmsgKey = "Recvmsg"

	RecvmsgKey = "Recvmsg"

	RecvmsgKey = "Recvmsg"

	RegCloseKeyKey = "RegCloseKey"

	RegEnumKeyExKey = "RegEnumKeyEx"

	RegOpenKeyExKey = "RegOpenKeyEx"

	RegQueryInfoKeyKey = "RegQueryInfoKey"

	RegQueryValueExKey = "RegQueryValueEx"

	RemoveKey = "Remove"

	RemoveDirectoryKey = "RemoveDirectory"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RemovexattrKey = "Removexattr"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameKey = "Rename"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RenameatKey = "Renameat"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RevokeKey = "Revoke"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RmdirKey = "Rmdir"

	RouteRIBKey = "RouteRIB"

	RouteRIBKey = "RouteRIB"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SeekKey = "Seek"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SelectKey = "Select"

	SendfileKey = "Sendfile"

	SendfileKey = "Sendfile"

	SendfileKey = "Sendfile"

	SendmsgKey = "Sendmsg"

	SendmsgKey = "Sendmsg"

	SendmsgKey = "Sendmsg"

	SendmsgKey = "Sendmsg"

	SendmsgKey = "Sendmsg"

	SendmsgNKey = "SendmsgN"

	SendmsgNKey = "SendmsgN"

	SendmsgNKey = "SendmsgN"

	SendmsgNKey = "SendmsgN"

	SendmsgNKey = "SendmsgN"

	SendmsgNKey = "SendmsgN"

	SendtoKey = "Sendto"

	SendtoKey = "Sendto"

	SendtoKey = "Sendto"

	SendtoKey = "Sendto"

	SetBpfKey = "SetBpf"

	SetBpfKey = "SetBpf"

	SetBpfBuflenKey = "SetBpfBuflen"

	SetBpfBuflenKey = "SetBpfBuflen"

	SetBpfDatalinkKey = "SetBpfDatalink"

	SetBpfDatalinkKey = "SetBpfDatalink"

	SetBpfHeadercmplKey = "SetBpfHeadercmpl"

	SetBpfHeadercmplKey = "SetBpfHeadercmpl"

	SetBpfImmediateKey = "SetBpfImmediate"

	SetBpfImmediateKey = "SetBpfImmediate"

	SetBpfInterfaceKey = "SetBpfInterface"

	SetBpfInterfaceKey = "SetBpfInterface"

	SetBpfPromiscKey = "SetBpfPromisc"

	SetBpfPromiscKey = "SetBpfPromisc"

	SetBpfTimeoutKey = "SetBpfTimeout"

	SetBpfTimeoutKey = "SetBpfTimeout"

	SetCurrentDirectoryKey = "SetCurrentDirectory"

	SetEndOfFileKey = "SetEndOfFile"

	SetEnvironmentVariableKey = "SetEnvironmentVariable"

	SetFileAttributesKey = "SetFileAttributes"

	SetFileCompletionNotificationModesKey = "SetFileCompletionNotificationModes"

	SetFilePointerKey = "SetFilePointer"

	SetFileTimeKey = "SetFileTime"

	SetHandleInformationKey = "SetHandleInformation"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetKeventKey = "SetKevent"

	SetLsfPromiscKey = "SetLsfPromisc"

	SetNonblockKey = "SetNonblock"

	SetNonblockKey = "SetNonblock"

	SetNonblockKey = "SetNonblock"

	SetNonblockKey = "SetNonblock"

	SetReadDeadlineKey = "SetReadDeadline"

	SetReadDeadlineKey = "SetReadDeadline"

	SetWriteDeadlineKey = "SetWriteDeadline"

	SetWriteDeadlineKey = "SetWriteDeadline"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetdomainnameKey = "Setdomainname"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetenvKey = "Setenv"

	SetenvKey = "Setenv"

	SetenvKey = "Setenv"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SeteuidKey = "Seteuid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsgidKey = "Setfsgid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetfsuidKey = "Setfsuid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgidKey = "Setgid"

	SetgroupsKey = "Setgroups"

	SetgroupsKey = "Setgroups"

	SetgroupsKey = "Setgroups"

	SetgroupsKey = "Setgroups"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SethostnameKey = "Sethostname"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpgidKey = "Setpgid"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetpriorityKey = "Setpriority"

	SetprivexecKey = "Setprivexec"

	SetprivexecKey = "Setprivexec"

	SetprivexecKey = "Setprivexec"

	SetprivexecKey = "Setprivexec"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetregidKey = "Setregid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresgidKey = "Setresgid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetresuidKey = "Setresuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetreuidKey = "Setreuid"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsockoptKey = "Setsockopt"

	SetsockoptByteKey = "SetsockoptByte"

	SetsockoptByteKey = "SetsockoptByte"

	SetsockoptICMPv6FilterKey = "SetsockoptICMPv6Filter"

	SetsockoptICMPv6FilterKey = "SetsockoptICMPv6Filter"

	SetsockoptIPMreqKey = "SetsockoptIPMreq"

	SetsockoptIPMreqKey = "SetsockoptIPMreq"

	SetsockoptIPMreqKey = "SetsockoptIPMreq"

	SetsockoptIPMreqnKey = "SetsockoptIPMreqn"

	SetsockoptIPMreqnKey = "SetsockoptIPMreqn"

	SetsockoptIPv6MreqKey = "SetsockoptIPv6Mreq"

	SetsockoptIPv6MreqKey = "SetsockoptIPv6Mreq"

	SetsockoptIPv6MreqKey = "SetsockoptIPv6Mreq"

	SetsockoptInet4AddrKey = "SetsockoptInet4Addr"

	SetsockoptInet4AddrKey = "SetsockoptInet4Addr"

	SetsockoptInet4AddrKey = "SetsockoptInet4Addr"

	SetsockoptIntKey = "SetsockoptInt"

	SetsockoptIntKey = "SetsockoptInt"

	SetsockoptIntKey = "SetsockoptInt"

	SetsockoptIntKey = "SetsockoptInt"

	SetsockoptLingerKey = "SetsockoptLinger"

	SetsockoptLingerKey = "SetsockoptLinger"

	SetsockoptLingerKey = "SetsockoptLinger"

	SetsockoptStringKey = "SetsockoptString"

	SetsockoptStringKey = "SetsockoptString"

	SetsockoptTimevalKey = "SetsockoptTimeval"

	SetsockoptTimevalKey = "SetsockoptTimeval"

	SetsockoptTimevalKey = "SetsockoptTimeval"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SettimeofdayKey = "Settimeofday"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetuidKey = "Setuid"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	SetxattrKey = "Setxattr"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	SlicePtrFromStringsKey = "SlicePtrFromStrings"

	SlicePtrFromStringsKey = "SlicePtrFromStrings"

	SocketKey = "Socket"

	SocketKey = "Socket"

	SocketKey = "Socket"

	SocketKey = "Socket"

	SocketpairKey = "Socketpair"

	SocketpairKey = "Socketpair"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	StartProcessKey = "StartProcess"

	StartProcessKey = "StartProcess"

	StartProcessKey = "StartProcess"

	StartProcessKey = "StartProcess"

	StartProcessKey = "StartProcess"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatKey = "Stat"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StopIOKey = "StopIO"

	StopIOKey = "StopIO"

	StringBytePtrKey = "StringBytePtr"

	StringByteSliceKey = "StringByteSlice"

	StringSlicePtrKey = "StringSlicePtr"

	StringSlicePtrKey = "StringSlicePtr"

	StringToSidKey = "StringToSid"

	StringToUTF16Key = "StringToUTF16"

	StringToUTF16PtrKey = "StringToUTF16Ptr"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SymlinkKey = "Symlink"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncKey = "Sync"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyscallKey = "Syscall"

	SyscallKey = "Syscall"

	SyscallKey = "Syscall"

	SyscallKey = "Syscall"

	SyscallKey = "Syscall"

	Syscall12Key = "Syscall12"

	Syscall15Key = "Syscall15"

	Syscall18Key = "Syscall18"

	Syscall6Key = "Syscall6"

	Syscall6Key = "Syscall6"

	Syscall6Key = "Syscall6"

	Syscall6Key = "Syscall6"

	Syscall6Key = "Syscall6"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	Syscall9Key = "Syscall9"

	SysctlKey = "Sysctl"

	SysctlKey = "Sysctl"

	SysctlKey = "Sysctl"

	SysctlUint32Key = "SysctlUint32"

	SysctlUint32Key = "SysctlUint32"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	SysinfoKey = "Sysinfo"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TeeKey = "Tee"

	TerminateProcessKey = "TerminateProcess"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TgkillKey = "Tgkill"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimespecToNsecKey = "TimespecToNsec"

	TimespecToNsecKey = "TimespecToNsec"

	TimevalToNsecKey = "TimevalToNsec"

	TranslateAccountNameKey = "TranslateAccountName"

	TranslateNameKey = "TranslateName"

	TransmitFileKey = "TransmitFile"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	UTF16FromStringKey = "UTF16FromString"

	UTF16PtrFromStringKey = "UTF16PtrFromString"

	UTF16ToStringKey = "UTF16ToString"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UmaskKey = "Umask"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UnameKey = "Uname"

	UndeleteKey = "Undelete"

	UndeleteKey = "Undelete"

	UndeleteKey = "Undelete"

	UndeleteKey = "Undelete"

	UndeleteKey = "Undelete"

	UndeleteKey = "Undelete"

	UndeleteKey = "Undelete"

	UndeleteKey = "Undelete"

	UnixCredentialsKey = "UnixCredentials"

	UnixRightsKey = "UnixRights"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnmapViewOfFileKey = "UnmapViewOfFile"

	UnmarshalDirKey = "UnmarshalDir"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnsetenvKey = "Unsetenv"

	UnsetenvKey = "Unsetenv"

	UnsetenvKey = "Unsetenv"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UstatKey = "Ustat"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimesKey = "Utimes"

	UtimesKey = "Utimes"

	UtimesKey = "Utimes"

	UtimesKey = "Utimes"

	UtimesKey = "Utimes"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	VirtualLockKey = "VirtualLock"

	VirtualUnlockKey = "VirtualUnlock"

	WSACleanupKey = "WSACleanup"

	WSAEnumProtocolsKey = "WSAEnumProtocols"

	WSAIoctlKey = "WSAIoctl"

	WSARecvKey = "WSARecv"

	WSARecvFromKey = "WSARecvFrom"

	WSASendKey = "WSASend"

	WSASendToKey = "WSASendTo"

	WSASendtoKey = "WSASendto"

	WSAStartupKey = "WSAStartup"

	Wait4Key = "Wait4"

	Wait4Key = "Wait4"

	Wait4Key = "Wait4"

	Wait4Key = "Wait4"

	Wait4Key = "Wait4"

	Wait4Key = "Wait4"

	WaitForSingleObjectKey = "WaitForSingleObject"

	WaitProcessKey = "WaitProcess"

	WriteKey = "Write"

	WriteKey = "Write"

	WriteKey = "Write"

	WriteKey = "Write"

	WriteKey = "Write"

	WriteConsoleKey = "WriteConsole"

	WriteFileKey = "WriteFile"

	WstatKey = "Wstat"
)

func New() hctx.Map {
	return hctx.Map{

		AcceptKey: Accept,

		AcceptKey: Accept,

		AcceptKey: Accept,

		AcceptKey: Accept,

		AcceptKey: Accept,

		AcceptKey: Accept,

		AcceptKey: Accept,

		Accept4Key: Accept4,

		Accept4Key: Accept4,

		Accept4Key: Accept4,

		Accept4Key: Accept4,

		Accept4Key: Accept4,

		AcceptExKey: AcceptEx,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AccessKey: Access,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimeKey: Adjtime,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		AttachLsfKey: AttachLsf,

		AwaitKey: Await,

		BindKey: Bind,

		BindKey: Bind,

		BindKey: Bind,

		BindKey: Bind,

		BindKey: Bind,

		BindToDeviceKey: BindToDevice,

		BpfBuflenKey: BpfBuflen,

		BpfBuflenKey: BpfBuflen,

		BpfDatalinkKey: BpfDatalink,

		BpfDatalinkKey: BpfDatalink,

		BpfHeadercmplKey: BpfHeadercmpl,

		BpfHeadercmplKey: BpfHeadercmpl,

		BpfInterfaceKey: BpfInterface,

		BpfInterfaceKey: BpfInterface,

		BpfJumpKey: BpfJump,

		BpfJumpKey: BpfJump,

		BpfStatsKey: BpfStats,

		BpfStatsKey: BpfStats,

		BpfStmtKey: BpfStmt,

		BpfStmtKey: BpfStmt,

		BpfTimeoutKey: BpfTimeout,

		BpfTimeoutKey: BpfTimeout,

		BytePtrFromStringKey: BytePtrFromString,

		ByteSliceFromStringKey: ByteSliceFromString,

		CancelIoKey: CancelIo,

		CancelIoExKey: CancelIoEx,

		CertAddCertificateContextToStoreKey: CertAddCertificateContextToStore,

		CertCloseStoreKey: CertCloseStore,

		CertCreateCertificateContextKey: CertCreateCertificateContext,

		CertEnumCertificatesInStoreKey: CertEnumCertificatesInStore,

		CertFreeCertificateChainKey: CertFreeCertificateChain,

		CertFreeCertificateContextKey: CertFreeCertificateContext,

		CertGetCertificateChainKey: CertGetCertificateChain,

		CertOpenStoreKey: CertOpenStore,

		CertOpenSystemStoreKey: CertOpenSystemStore,

		CertVerifyCertificateChainPolicyKey: CertVerifyCertificateChainPolicy,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		CheckBpfVersionKey: CheckBpfVersion,

		CheckBpfVersionKey: CheckBpfVersion,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChflagsKey: Chflags,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChmodKey: Chmod,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChownKey: Chown,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ChrootKey: Chroot,

		ClearenvKey: Clearenv,

		ClearenvKey: Clearenv,

		ClearenvKey: Clearenv,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseKey: Close,

		CloseHandleKey: CloseHandle,

		CloseOnExecKey: CloseOnExec,

		CloseOnExecKey: CloseOnExec,

		CloseOnExecKey: CloseOnExec,

		CloseOnExecKey: CloseOnExec,

		ClosesocketKey: Closesocket,

		CmsgLenKey: CmsgLen,

		CmsgSpaceKey: CmsgSpace,

		CommandLineToArgvKey: CommandLineToArgv,

		ComputerNameKey: ComputerName,

		ConnectKey: Connect,

		ConnectKey: Connect,

		ConnectKey: Connect,

		ConnectKey: Connect,

		ConnectExKey: ConnectEx,

		ConvertSidToStringSidKey: ConvertSidToStringSid,

		ConvertStringSidToSidKey: ConvertStringSidToSid,

		CopySidKey: CopySid,

		CreatKey: Creat,

		CreateKey: Create,

		CreateDirectoryKey: CreateDirectory,

		CreateFileKey: CreateFile,

		CreateFileMappingKey: CreateFileMapping,

		CreateHardLinkKey: CreateHardLink,

		CreateIoCompletionPortKey: CreateIoCompletionPort,

		CreatePipeKey: CreatePipe,

		CreateProcessKey: CreateProcess,

		CreateProcessAsUserKey: CreateProcessAsUser,

		CreateSymbolicLinkKey: CreateSymbolicLink,

		CreateToolhelp32SnapshotKey: CreateToolhelp32Snapshot,

		CryptAcquireContextKey: CryptAcquireContext,

		CryptGenRandomKey: CryptGenRandom,

		CryptReleaseContextKey: CryptReleaseContext,

		DeleteFileKey: DeleteFile,

		DetachLsfKey: DetachLsf,

		DeviceIoControlKey: DeviceIoControl,

		DnsNameCompareKey: DnsNameCompare,

		DnsQueryKey: DnsQuery,

		DnsRecordListFreeKey: DnsRecordListFree,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		DupKey: Dup,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup2Key: Dup2,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		Dup3Key: Dup3,

		DuplicateHandleKey: DuplicateHandle,

		EnvironKey: Environ,

		EnvironKey: Environ,

		EnvironKey: Environ,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreateKey: EpollCreate,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCreate1Key: EpollCreate1,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollCtlKey: EpollCtl,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		EscapeArgKey: EscapeArg,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExecKey: Exec,

		ExecKey: Exec,

		ExecKey: Exec,

		ExitKey: Exit,

		ExitProcessKey: ExitProcess,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FallocateKey: Fallocate,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchdirKey: Fchdir,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchflagsKey: Fchflags,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownKey: Fchown,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FchownatKey: Fchownat,

		FcntlFlockKey: FcntlFlock,

		FcntlFlockKey: FcntlFlock,

		FcntlFlockKey: FcntlFlock,

		FcntlFlockKey: FcntlFlock,

		Fd2pathKey: Fd2path,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FindCloseKey: FindClose,

		FindFirstFileKey: FindFirstFile,

		FindNextFileKey: FindNextFile,

		FixwdKey: Fixwd,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlushBpfKey: FlushBpf,

		FlushBpfKey: FlushBpf,

		FlushFileBuffersKey: FlushFileBuffers,

		FlushViewOfFileKey: FlushViewOfFile,

		ForkExecKey: ForkExec,

		ForkExecKey: ForkExec,

		FormatMessageKey: FormatMessage,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FreeAddrInfoWKey: FreeAddrInfoW,

		FreeEnvironmentStringsKey: FreeEnvironmentStrings,

		FreeLibraryKey: FreeLibrary,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FsyncKey: Fsync,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FullPathKey: FullPath,

		FutimesKey: Futimes,

		FutimesKey: Futimes,

		FutimesatKey: Futimesat,

		FwstatKey: Fwstat,

		FwstatKey: Fwstat,

		FwstatKey: Fwstat,

		GetAcceptExSockaddrsKey: GetAcceptExSockaddrs,

		GetAdaptersInfoKey: GetAdaptersInfo,

		GetAddrInfoWKey: GetAddrInfoW,

		GetCommandLineKey: GetCommandLine,

		GetComputerNameKey: GetComputerName,

		GetConsoleModeKey: GetConsoleMode,

		GetCurrentDirectoryKey: GetCurrentDirectory,

		GetCurrentProcessKey: GetCurrentProcess,

		GetEnvironmentStringsKey: GetEnvironmentStrings,

		GetEnvironmentVariableKey: GetEnvironmentVariable,

		GetExitCodeProcessKey: GetExitCodeProcess,

		GetFileAttributesKey: GetFileAttributes,

		GetFileAttributesExKey: GetFileAttributesEx,

		GetFileInformationByHandleKey: GetFileInformationByHandle,

		GetFileTypeKey: GetFileType,

		GetFullPathNameKey: GetFullPathName,

		GetHostByNameKey: GetHostByName,

		GetIfEntryKey: GetIfEntry,

		GetLastErrorKey: GetLastError,

		GetLengthSidKey: GetLengthSid,

		GetLongPathNameKey: GetLongPathName,

		GetProcAddressKey: GetProcAddress,

		GetProcessTimesKey: GetProcessTimes,

		GetProtoByNameKey: GetProtoByName,

		GetQueuedCompletionStatusKey: GetQueuedCompletionStatus,

		GetServByNameKey: GetServByName,

		GetShortPathNameKey: GetShortPathName,

		GetStartupInfoKey: GetStartupInfo,

		GetStdHandleKey: GetStdHandle,

		GetSystemTimeAsFileTimeKey: GetSystemTimeAsFileTime,

		GetTempPathKey: GetTempPath,

		GetTimeZoneInformationKey: GetTimeZoneInformation,

		GetTokenInformationKey: GetTokenInformation,

		GetUserNameExKey: GetUserNameEx,

		GetUserProfileDirectoryKey: GetUserProfileDirectory,

		GetVersionKey: GetVersion,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetcwdKey: Getcwd,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdentsKey: Getdents,

		GetdirentriesKey: Getdirentries,

		GetdirentriesKey: Getdirentries,

		GetdirentriesKey: Getdirentries,

		GetdirentriesKey: Getdirentries,

		GetdirentriesKey: Getdirentries,

		GetdirentriesKey: Getdirentries,

		GetdirentriesKey: Getdirentries,

		GetdirentriesKey: Getdirentries,

		GetdtablesizeKey: Getdtablesize,

		GetdtablesizeKey: Getdtablesize,

		GetdtablesizeKey: Getdtablesize,

		GetdtablesizeKey: Getdtablesize,

		GetdtablesizeKey: Getdtablesize,

		GetdtablesizeKey: Getdtablesize,

		GetdtablesizeKey: Getdtablesize,

		GetdtablesizeKey: Getdtablesize,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetenvKey: Getenv,

		GetenvKey: Getenv,

		GetenvKey: Getenv,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GetexecnameKey: Getexecname,

		GetfsstatKey: Getfsstat,

		GetfsstatKey: Getfsstat,

		GetfsstatKey: Getfsstat,

		GetfsstatKey: Getfsstat,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GethostnameKey: Gethostname,

		GetkerninfoKey: Getkerninfo,

		GetpagesizeKey: Getpagesize,

		GetpeernameKey: Getpeername,

		GetpeernameKey: Getpeername,

		GetpeernameKey: Getpeername,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgidKey: Getpgid,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpgrpKey: Getpgrp,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetpidKey: Getpid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetppidKey: Getppid,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetrusageKey: Getrusage,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsidKey: Getsid,

		GetsocknameKey: Getsockname,

		GetsocknameKey: Getsockname,

		GetsocknameKey: Getsockname,

		GetsocknameKey: Getsockname,

		GetsocknameKey: Getsockname,

		GetsocknameKey: Getsockname,

		GetsockoptKey: Getsockopt,

		GetsockoptByteKey: GetsockoptByte,

		GetsockoptICMPv6FilterKey: GetsockoptICMPv6Filter,

		GetsockoptICMPv6FilterKey: GetsockoptICMPv6Filter,

		GetsockoptIPMreqKey: GetsockoptIPMreq,

		GetsockoptIPMreqKey: GetsockoptIPMreq,

		GetsockoptIPMreqnKey: GetsockoptIPMreqn,

		GetsockoptIPMreqnKey: GetsockoptIPMreqn,

		GetsockoptIPv6MTUInfoKey: GetsockoptIPv6MTUInfo,

		GetsockoptIPv6MTUInfoKey: GetsockoptIPv6MTUInfo,

		GetsockoptIPv6MreqKey: GetsockoptIPv6Mreq,

		GetsockoptIPv6MreqKey: GetsockoptIPv6Mreq,

		GetsockoptInet4AddrKey: GetsockoptInet4Addr,

		GetsockoptInet4AddrKey: GetsockoptInet4Addr,

		GetsockoptIntKey: GetsockoptInt,

		GetsockoptIntKey: GetsockoptInt,

		GetsockoptIntKey: GetsockoptInt,

		GetsockoptIntKey: GetsockoptInt,

		GetsockoptUcredKey: GetsockoptUcred,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettidKey: Gettid,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GettimeofdayKey: Gettimeofday,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetuidKey: Getuid,

		GetwdKey: Getwd,

		GetwdKey: Getwd,

		GetwdKey: Getwd,

		GetwdKey: Getwd,

		GetwdKey: Getwd,

		GetwdKey: Getwd,

		GetwdKey: Getwd,

		GetwdKey: Getwd,

		GetwdKey: Getwd,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyAddWatchKey: InotifyAddWatch,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInitKey: InotifyInit,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyInit1Key: InotifyInit1,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		IopermKey: Ioperm,

		IopermKey: Ioperm,

		IopermKey: Ioperm,

		IopermKey: Ioperm,

		IopermKey: Ioperm,

		IopermKey: Ioperm,

		IopermKey: Ioperm,

		IoplKey: Iopl,

		IoplKey: Iopl,

		IoplKey: Iopl,

		IoplKey: Iopl,

		IoplKey: Iopl,

		IoplKey: Iopl,

		IoplKey: Iopl,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		KeventKey: Kevent,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KillKey: Kill,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KlogctlKey: Klogctl,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		KqueueKey: Kqueue,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		LinkKey: Link,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListenKey: Listen,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		LoadCancelIoExKey: LoadCancelIoEx,

		LoadConnectExKey: LoadConnectEx,

		LoadCreateSymbolicLinkKey: LoadCreateSymbolicLink,

		LoadDLLKey: LoadDLL,

		LoadGetAddrInfoKey: LoadGetAddrInfo,

		LoadLibraryKey: LoadLibrary,

		LoadSetFileCompletionNotificationModesKey: LoadSetFileCompletionNotificationModes,

		LocalFreeKey: LocalFree,

		LookupAccountNameKey: LookupAccountName,

		LookupAccountSidKey: LookupAccountSid,

		LookupSIDKey: LookupSID,

		LsfJumpKey: LsfJump,

		LsfSocketKey: LsfSocket,

		LsfStmtKey: LsfStmt,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		LstatKey: Lstat,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MapViewOfFileKey: MapViewOfFile,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdirKey: Mkdir,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkdiratKey: Mkdirat,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodKey: Mknod,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MknodatKey: Mknodat,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockKey: Mlock,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MmapKey: Mmap,

		MmapKey: Mmap,

		MmapKey: Mmap,

		MountKey: Mount,

		MountKey: Mount,

		MoveFileKey: MoveFile,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockKey: Munlock,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunmapKey: Munmap,

		MunmapKey: Munmap,

		MunmapKey: Munmap,

		MustLoadDLLKey: MustLoadDLL,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NetApiBufferFreeKey: NetApiBufferFree,

		NetGetJoinInformationKey: NetGetJoinInformation,

		NetUserGetInfoKey: NetUserGetInfo,

		NetlinkRIBKey: NetlinkRIB,

		NewCallbackKey: NewCallback,

		NewCallbackCDeclKey: NewCallbackCDecl,

		NewErrorKey: NewError,

		NewLazyDLLKey: NewLazyDLL,

		NsecToFiletimeKey: NsecToFiletime,

		NsecToTimespecKey: NsecToTimespec,

		NsecToTimespecKey: NsecToTimespec,

		NsecToTimevalKey: NsecToTimeval,

		NsecToTimevalKey: NsecToTimeval,

		NsecToTimevalKey: NsecToTimeval,

		NtohsKey: Ntohs,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenKey: Open,

		OpenCurrentProcessTokenKey: OpenCurrentProcessToken,

		OpenProcessKey: OpenProcess,

		OpenProcessTokenKey: OpenProcessToken,

		OpenatKey: Openat,

		OpenatKey: Openat,

		ParseDirentKey: ParseDirent,

		ParseFilesKey: ParseFiles,

		ParseNetlinkMessageKey: ParseNetlinkMessage,

		ParseNetlinkRouteAttrKey: ParseNetlinkRouteAttr,

		ParseRoutingMessageKey: ParseRoutingMessage,

		ParseRoutingMessageKey: ParseRoutingMessage,

		ParseRoutingSockaddrKey: ParseRoutingSockaddr,

		ParseRoutingSockaddrKey: ParseRoutingSockaddr,

		ParseSocketControlMessageKey: ParseSocketControlMessage,

		ParseUnixCredentialsKey: ParseUnixCredentials,

		ParseUnixRightsKey: ParseUnixRights,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PathconfKey: Pathconf,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		PipeKey: Pipe,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		Pipe2Key: Pipe2,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PostQueuedCompletionStatusKey: PostQueuedCompletionStatus,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		Process32FirstKey: Process32First,

		Process32NextKey: Process32Next,

		PtraceAttachKey: PtraceAttach,

		PtraceAttachKey: PtraceAttach,

		PtraceAttachKey: PtraceAttach,

		PtraceContKey: PtraceCont,

		PtraceContKey: PtraceCont,

		PtraceDetachKey: PtraceDetach,

		PtraceDetachKey: PtraceDetach,

		PtraceDetachKey: PtraceDetach,

		PtraceGetEventMsgKey: PtraceGetEventMsg,

		PtraceGetRegsKey: PtraceGetRegs,

		PtracePeekDataKey: PtracePeekData,

		PtracePeekDataKey: PtracePeekData,

		PtracePeekTextKey: PtracePeekText,

		PtracePeekTextKey: PtracePeekText,

		PtracePokeDataKey: PtracePokeData,

		PtracePokeDataKey: PtracePokeData,

		PtracePokeTextKey: PtracePokeText,

		PtracePokeTextKey: PtracePokeText,

		PtraceSetOptionsKey: PtraceSetOptions,

		PtraceSetRegsKey: PtraceSetRegs,

		PtraceSingleStepKey: PtraceSingleStep,

		PtraceSingleStepKey: PtraceSingleStep,

		PtraceSyscallKey: PtraceSyscall,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		RawSyscallKey: RawSyscall,

		RawSyscallKey: RawSyscall,

		RawSyscallKey: RawSyscall,

		RawSyscallKey: RawSyscall,

		RawSyscall6Key: RawSyscall6,

		RawSyscall6Key: RawSyscall6,

		RawSyscall6Key: RawSyscall6,

		RawSyscall6Key: RawSyscall6,

		ReadKey: Read,

		ReadKey: Read,

		ReadKey: Read,

		ReadKey: Read,

		ReadKey: Read,

		ReadConsoleKey: ReadConsole,

		ReadDirectoryChangesKey: ReadDirectoryChanges,

		ReadDirentKey: ReadDirent,

		ReadDirentKey: ReadDirent,

		ReadDirentKey: ReadDirent,

		ReadDirentKey: ReadDirent,

		ReadDirentKey: ReadDirent,

		ReadDirentKey: ReadDirent,

		ReadFileKey: ReadFile,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		RebootKey: Reboot,

		RebootKey: Reboot,

		RecvfromKey: Recvfrom,

		RecvfromKey: Recvfrom,

		RecvfromKey: Recvfrom,

		RecvfromKey: Recvfrom,

		RecvmsgKey: Recvmsg,

		RecvmsgKey: Recvmsg,

		RecvmsgKey: Recvmsg,

		RecvmsgKey: Recvmsg,

		RecvmsgKey: Recvmsg,

		RecvmsgKey: Recvmsg,

		RegCloseKeyKey: RegCloseKey,

		RegEnumKeyExKey: RegEnumKeyEx,

		RegOpenKeyExKey: RegOpenKeyEx,

		RegQueryInfoKeyKey: RegQueryInfoKey,

		RegQueryValueExKey: RegQueryValueEx,

		RemoveKey: Remove,

		RemoveDirectoryKey: RemoveDirectory,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RemovexattrKey: Removexattr,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameKey: Rename,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RenameatKey: Renameat,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RevokeKey: Revoke,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RmdirKey: Rmdir,

		RouteRIBKey: RouteRIB,

		RouteRIBKey: RouteRIB,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SeekKey: Seek,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SelectKey: Select,

		SendfileKey: Sendfile,

		SendfileKey: Sendfile,

		SendfileKey: Sendfile,

		SendmsgKey: Sendmsg,

		SendmsgKey: Sendmsg,

		SendmsgKey: Sendmsg,

		SendmsgKey: Sendmsg,

		SendmsgKey: Sendmsg,

		SendmsgNKey: SendmsgN,

		SendmsgNKey: SendmsgN,

		SendmsgNKey: SendmsgN,

		SendmsgNKey: SendmsgN,

		SendmsgNKey: SendmsgN,

		SendmsgNKey: SendmsgN,

		SendtoKey: Sendto,

		SendtoKey: Sendto,

		SendtoKey: Sendto,

		SendtoKey: Sendto,

		SetBpfKey: SetBpf,

		SetBpfKey: SetBpf,

		SetBpfBuflenKey: SetBpfBuflen,

		SetBpfBuflenKey: SetBpfBuflen,

		SetBpfDatalinkKey: SetBpfDatalink,

		SetBpfDatalinkKey: SetBpfDatalink,

		SetBpfHeadercmplKey: SetBpfHeadercmpl,

		SetBpfHeadercmplKey: SetBpfHeadercmpl,

		SetBpfImmediateKey: SetBpfImmediate,

		SetBpfImmediateKey: SetBpfImmediate,

		SetBpfInterfaceKey: SetBpfInterface,

		SetBpfInterfaceKey: SetBpfInterface,

		SetBpfPromiscKey: SetBpfPromisc,

		SetBpfPromiscKey: SetBpfPromisc,

		SetBpfTimeoutKey: SetBpfTimeout,

		SetBpfTimeoutKey: SetBpfTimeout,

		SetCurrentDirectoryKey: SetCurrentDirectory,

		SetEndOfFileKey: SetEndOfFile,

		SetEnvironmentVariableKey: SetEnvironmentVariable,

		SetFileAttributesKey: SetFileAttributes,

		SetFileCompletionNotificationModesKey: SetFileCompletionNotificationModes,

		SetFilePointerKey: SetFilePointer,

		SetFileTimeKey: SetFileTime,

		SetHandleInformationKey: SetHandleInformation,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetKeventKey: SetKevent,

		SetLsfPromiscKey: SetLsfPromisc,

		SetNonblockKey: SetNonblock,

		SetNonblockKey: SetNonblock,

		SetNonblockKey: SetNonblock,

		SetNonblockKey: SetNonblock,

		SetReadDeadlineKey: SetReadDeadline,

		SetReadDeadlineKey: SetReadDeadline,

		SetWriteDeadlineKey: SetWriteDeadline,

		SetWriteDeadlineKey: SetWriteDeadline,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetdomainnameKey: Setdomainname,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetenvKey: Setenv,

		SetenvKey: Setenv,

		SetenvKey: Setenv,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SeteuidKey: Seteuid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsgidKey: Setfsgid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetfsuidKey: Setfsuid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgidKey: Setgid,

		SetgroupsKey: Setgroups,

		SetgroupsKey: Setgroups,

		SetgroupsKey: Setgroups,

		SetgroupsKey: Setgroups,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SethostnameKey: Sethostname,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpgidKey: Setpgid,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetpriorityKey: Setpriority,

		SetprivexecKey: Setprivexec,

		SetprivexecKey: Setprivexec,

		SetprivexecKey: Setprivexec,

		SetprivexecKey: Setprivexec,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetregidKey: Setregid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresgidKey: Setresgid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetresuidKey: Setresuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetreuidKey: Setreuid,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsockoptKey: Setsockopt,

		SetsockoptByteKey: SetsockoptByte,

		SetsockoptByteKey: SetsockoptByte,

		SetsockoptICMPv6FilterKey: SetsockoptICMPv6Filter,

		SetsockoptICMPv6FilterKey: SetsockoptICMPv6Filter,

		SetsockoptIPMreqKey: SetsockoptIPMreq,

		SetsockoptIPMreqKey: SetsockoptIPMreq,

		SetsockoptIPMreqKey: SetsockoptIPMreq,

		SetsockoptIPMreqnKey: SetsockoptIPMreqn,

		SetsockoptIPMreqnKey: SetsockoptIPMreqn,

		SetsockoptIPv6MreqKey: SetsockoptIPv6Mreq,

		SetsockoptIPv6MreqKey: SetsockoptIPv6Mreq,

		SetsockoptIPv6MreqKey: SetsockoptIPv6Mreq,

		SetsockoptInet4AddrKey: SetsockoptInet4Addr,

		SetsockoptInet4AddrKey: SetsockoptInet4Addr,

		SetsockoptInet4AddrKey: SetsockoptInet4Addr,

		SetsockoptIntKey: SetsockoptInt,

		SetsockoptIntKey: SetsockoptInt,

		SetsockoptIntKey: SetsockoptInt,

		SetsockoptIntKey: SetsockoptInt,

		SetsockoptLingerKey: SetsockoptLinger,

		SetsockoptLingerKey: SetsockoptLinger,

		SetsockoptLingerKey: SetsockoptLinger,

		SetsockoptStringKey: SetsockoptString,

		SetsockoptStringKey: SetsockoptString,

		SetsockoptTimevalKey: SetsockoptTimeval,

		SetsockoptTimevalKey: SetsockoptTimeval,

		SetsockoptTimevalKey: SetsockoptTimeval,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SettimeofdayKey: Settimeofday,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetuidKey: Setuid,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		SetxattrKey: Setxattr,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		SlicePtrFromStringsKey: SlicePtrFromStrings,

		SlicePtrFromStringsKey: SlicePtrFromStrings,

		SocketKey: Socket,

		SocketKey: Socket,

		SocketKey: Socket,

		SocketKey: Socket,

		SocketpairKey: Socketpair,

		SocketpairKey: Socketpair,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		StartProcessKey: StartProcess,

		StartProcessKey: StartProcess,

		StartProcessKey: StartProcess,

		StartProcessKey: StartProcess,

		StartProcessKey: StartProcess,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatKey: Stat,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StopIOKey: StopIO,

		StopIOKey: StopIO,

		StringBytePtrKey: StringBytePtr,

		StringByteSliceKey: StringByteSlice,

		StringSlicePtrKey: StringSlicePtr,

		StringSlicePtrKey: StringSlicePtr,

		StringToSidKey: StringToSid,

		StringToUTF16Key: StringToUTF16,

		StringToUTF16PtrKey: StringToUTF16Ptr,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SymlinkKey: Symlink,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncKey: Sync,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyscallKey: Syscall,

		SyscallKey: Syscall,

		SyscallKey: Syscall,

		SyscallKey: Syscall,

		SyscallKey: Syscall,

		Syscall12Key: Syscall12,

		Syscall15Key: Syscall15,

		Syscall18Key: Syscall18,

		Syscall6Key: Syscall6,

		Syscall6Key: Syscall6,

		Syscall6Key: Syscall6,

		Syscall6Key: Syscall6,

		Syscall6Key: Syscall6,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		Syscall9Key: Syscall9,

		SysctlKey: Sysctl,

		SysctlKey: Sysctl,

		SysctlKey: Sysctl,

		SysctlUint32Key: SysctlUint32,

		SysctlUint32Key: SysctlUint32,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		SysinfoKey: Sysinfo,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TeeKey: Tee,

		TerminateProcessKey: TerminateProcess,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TgkillKey: Tgkill,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimespecToNsecKey: TimespecToNsec,

		TimespecToNsecKey: TimespecToNsec,

		TimevalToNsecKey: TimevalToNsec,

		TranslateAccountNameKey: TranslateAccountName,

		TranslateNameKey: TranslateName,

		TransmitFileKey: TransmitFile,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		UTF16FromStringKey: UTF16FromString,

		UTF16PtrFromStringKey: UTF16PtrFromString,

		UTF16ToStringKey: UTF16ToString,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UmaskKey: Umask,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UnameKey: Uname,

		UndeleteKey: Undelete,

		UndeleteKey: Undelete,

		UndeleteKey: Undelete,

		UndeleteKey: Undelete,

		UndeleteKey: Undelete,

		UndeleteKey: Undelete,

		UndeleteKey: Undelete,

		UndeleteKey: Undelete,

		UnixCredentialsKey: UnixCredentials,

		UnixRightsKey: UnixRights,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnmapViewOfFileKey: UnmapViewOfFile,

		UnmarshalDirKey: UnmarshalDir,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnsetenvKey: Unsetenv,

		UnsetenvKey: Unsetenv,

		UnsetenvKey: Unsetenv,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UstatKey: Ustat,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimesKey: Utimes,

		UtimesKey: Utimes,

		UtimesKey: Utimes,

		UtimesKey: Utimes,

		UtimesKey: Utimes,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		VirtualLockKey: VirtualLock,

		VirtualUnlockKey: VirtualUnlock,

		WSACleanupKey: WSACleanup,

		WSAEnumProtocolsKey: WSAEnumProtocols,

		WSAIoctlKey: WSAIoctl,

		WSARecvKey: WSARecv,

		WSARecvFromKey: WSARecvFrom,

		WSASendKey: WSASend,

		WSASendToKey: WSASendTo,

		WSASendtoKey: WSASendto,

		WSAStartupKey: WSAStartup,

		Wait4Key: Wait4,

		Wait4Key: Wait4,

		Wait4Key: Wait4,

		Wait4Key: Wait4,

		Wait4Key: Wait4,

		Wait4Key: Wait4,

		WaitForSingleObjectKey: WaitForSingleObject,

		WaitProcessKey: WaitProcess,

		WriteKey: Write,

		WriteKey: Write,

		WriteKey: Write,

		WriteKey: Write,

		WriteKey: Write,

		WriteConsoleKey: WriteConsole,

		WriteFileKey: WriteFile,

		WstatKey: Wstat,
	}
}

// sys	accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error)
var Accept = syscall.Accept

var Accept = syscall.Accept

var Accept = syscall.Accept

var Accept = syscall.Accept

var Accept = syscall.Accept

var Accept = syscall.Accept

var Accept = syscall.Accept

var Accept4 = syscall.Accept4

var Accept4 = syscall.Accept4

// sys paccept(fd int, rsa *RawSockaddrAny, addrlen *_Socklen, sigmask *sigset, flags int) (nfd int, err error)
var Accept4 = syscall.Accept4

var Accept4 = syscall.Accept4

// sys	accept4(fd int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (nfd int, err error)
var Accept4 = syscall.Accept4

var AcceptEx = syscall.AcceptEx

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Access = syscall.Access

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Acct = syscall.Acct

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtime = syscall.Adjtime

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

var Adjtimex = syscall.Adjtimex

// Deprecated: Use golang.org/x/net/bpf instead.
var AttachLsf = syscall.AttachLsf

// sys	await(s []byte) (n int, err error)
var Await = syscall.Await

var Bind = syscall.Bind

var Bind = syscall.Bind

var Bind = syscall.Bind

var Bind = syscall.Bind

// sys	bind(name string, old string, flag int) (err error)
var Bind = syscall.Bind

// BindToDevice binds the socket associated with fd to device.
var BindToDevice = syscall.BindToDevice

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfBuflen = syscall.BpfBuflen

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfBuflen = syscall.BpfBuflen

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfDatalink = syscall.BpfDatalink

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfDatalink = syscall.BpfDatalink

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfHeadercmpl = syscall.BpfHeadercmpl

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfHeadercmpl = syscall.BpfHeadercmpl

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfInterface = syscall.BpfInterface

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfInterface = syscall.BpfInterface

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfJump = syscall.BpfJump

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfJump = syscall.BpfJump

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfStats = syscall.BpfStats

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfStats = syscall.BpfStats

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfStmt = syscall.BpfStmt

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfStmt = syscall.BpfStmt

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfTimeout = syscall.BpfTimeout

// Deprecated: Use golang.org/x/net/bpf instead.
var BpfTimeout = syscall.BpfTimeout

// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
var BytePtrFromString = syscall.BytePtrFromString

// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
var ByteSliceFromString = syscall.ByteSliceFromString

var CancelIo = syscall.CancelIo

var CancelIoEx = syscall.CancelIoEx

var CertAddCertificateContextToStore = syscall.CertAddCertificateContextToStore

var CertCloseStore = syscall.CertCloseStore

var CertCreateCertificateContext = syscall.CertCreateCertificateContext

var CertEnumCertificatesInStore = syscall.CertEnumCertificatesInStore

var CertFreeCertificateChain = syscall.CertFreeCertificateChain

var CertFreeCertificateContext = syscall.CertFreeCertificateContext

var CertGetCertificateChain = syscall.CertGetCertificateChain

var CertOpenStore = syscall.CertOpenStore

var CertOpenSystemStore = syscall.CertOpenSystemStore

var CertVerifyCertificateChainPolicy = syscall.CertVerifyCertificateChainPolicy

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

var Chdir = syscall.Chdir

// Deprecated: Use golang.org/x/net/bpf instead.
var CheckBpfVersion = syscall.CheckBpfVersion

// Deprecated: Use golang.org/x/net/bpf instead.
var CheckBpfVersion = syscall.CheckBpfVersion

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chflags = syscall.Chflags

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chmod = syscall.Chmod

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chown = syscall.Chown

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Chroot = syscall.Chroot

var Clearenv = syscall.Clearenv

var Clearenv = syscall.Clearenv

var Clearenv = syscall.Clearenv

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var Close = syscall.Close

var CloseHandle = syscall.CloseHandle

var CloseOnExec = syscall.CloseOnExec

var CloseOnExec = syscall.CloseOnExec

var CloseOnExec = syscall.CloseOnExec

var CloseOnExec = syscall.CloseOnExec

var Closesocket = syscall.Closesocket

// CmsgLen returns the value to store in the Len field of the Cmsghdr
// structure, taking into account any necessary alignment.
var CmsgLen = syscall.CmsgLen

// CmsgSpace returns the number of bytes an ancillary element with
// payload of the passed data length occupies.
var CmsgSpace = syscall.CmsgSpace

var CommandLineToArgv = syscall.CommandLineToArgv

var ComputerName = syscall.ComputerName

var Connect = syscall.Connect

var Connect = syscall.Connect

var Connect = syscall.Connect

var Connect = syscall.Connect

var ConnectEx = syscall.ConnectEx

var ConvertSidToStringSid = syscall.ConvertSidToStringSid

var ConvertStringSidToSid = syscall.ConvertStringSidToSid

var CopySid = syscall.CopySid

var Creat = syscall.Creat

// sys	create(path string, mode int, perm uint32) (fd int, err error)
var Create = syscall.Create

var CreateDirectory = syscall.CreateDirectory

var CreateFile = syscall.CreateFile

var CreateFileMapping = syscall.CreateFileMapping

var CreateHardLink = syscall.CreateHardLink

var CreateIoCompletionPort = syscall.CreateIoCompletionPort

var CreatePipe = syscall.CreatePipe

var CreateProcess = syscall.CreateProcess

var CreateProcessAsUser = syscall.CreateProcessAsUser

var CreateSymbolicLink = syscall.CreateSymbolicLink

var CreateToolhelp32Snapshot = syscall.CreateToolhelp32Snapshot

var CryptAcquireContext = syscall.CryptAcquireContext

var CryptGenRandom = syscall.CryptGenRandom

var CryptReleaseContext = syscall.CryptReleaseContext

var DeleteFile = syscall.DeleteFile

// Deprecated: Use golang.org/x/net/bpf instead.
var DetachLsf = syscall.DetachLsf

var DeviceIoControl = syscall.DeviceIoControl

var DnsNameCompare = syscall.DnsNameCompare

var DnsQuery = syscall.DnsQuery

var DnsRecordListFree = syscall.DnsRecordListFree

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup = syscall.Dup

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup2 = syscall.Dup2

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var Dup3 = syscall.Dup3

var DuplicateHandle = syscall.DuplicateHandle

var Environ = syscall.Environ

var Environ = syscall.Environ

var Environ = syscall.Environ

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate = syscall.EpollCreate

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCreate1 = syscall.EpollCreate1

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollCtl = syscall.EpollCtl

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

var EpollWait = syscall.EpollWait

// EscapeArg rewrites command line argument s as prescribed
// in https://msdn.microsoft.com/en-us/library/ms880421.
// This function returns &#34;&#34; (2 double quotes) if s is empty.
// Alternatively, these transformations are done:
// - every back slash (\) is doubled, but only if immediately
//   followed by double quote (&#34;);
// - every double quote (&#34;) is escaped by back slash (\);
// - finally, s is wrapped with double quotes (arg -&gt; &#34;arg&#34;),
//   but only if there is space or tab inside s.
var EscapeArg = syscall.EscapeArg

var Exchangedata = syscall.Exchangedata

var Exchangedata = syscall.Exchangedata

var Exchangedata = syscall.Exchangedata

var Exchangedata = syscall.Exchangedata

// Ordinary exec.
var Exec = syscall.Exec

var Exec = syscall.Exec

// Exec invokes the execve(2) system call.
var Exec = syscall.Exec

var Exit = syscall.Exit

var ExitProcess = syscall.ExitProcess

var Faccessat = syscall.Faccessat

var Faccessat = syscall.Faccessat

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fallocate = syscall.Fallocate

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

// TODO(brainman): fix all needed for os
var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchdir = syscall.Fchdir

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchflags = syscall.Fchflags

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmod = syscall.Fchmod

var Fchmodat = syscall.Fchmodat

var Fchmodat = syscall.Fchmodat

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchown = syscall.Fchown

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

var Fchownat = syscall.Fchownat

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
var FcntlFlock = syscall.FcntlFlock

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
var FcntlFlock = syscall.FcntlFlock

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
var FcntlFlock = syscall.FcntlFlock

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
var FcntlFlock = syscall.FcntlFlock

// sys	fd2path(fd int, buf []byte) (err error)
var Fd2path = syscall.Fd2path

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var Fdatasync = syscall.Fdatasync

var FindClose = syscall.FindClose

var FindFirstFile = syscall.FindFirstFile

var FindNextFile = syscall.FindNextFile

var Fixwd = syscall.Fixwd

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

var Flock = syscall.Flock

// Deprecated: Use golang.org/x/net/bpf instead.
var FlushBpf = syscall.FlushBpf

// Deprecated: Use golang.org/x/net/bpf instead.
var FlushBpf = syscall.FlushBpf

var FlushFileBuffers = syscall.FlushFileBuffers

var FlushViewOfFile = syscall.FlushViewOfFile

// Combination of fork and exec, careful to be thread safe.
var ForkExec = syscall.ForkExec

// Combination of fork and exec, careful to be thread safe.
var ForkExec = syscall.ForkExec

// FormatMessage is deprecated (msgsrc should be uintptr, not uint32, but can
// not be changed due to the Go 1 compatibility guarantee).
//
// Deprecated: Use FormatMessage from golang.org/x/sys/windows instead.
var FormatMessage = syscall.FormatMessage

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var Fpathconf = syscall.Fpathconf

var FreeAddrInfoW = syscall.FreeAddrInfoW

var FreeEnvironmentStrings = syscall.FreeEnvironmentStrings

var FreeLibrary = syscall.FreeLibrary

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstat = syscall.Fstat

var Fstatat = syscall.Fstatat

var Fstatat = syscall.Fstatat

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fstatfs = syscall.Fstatfs

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Fsync = syscall.Fsync

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

var Ftruncate = syscall.Ftruncate

// FullPath retrieves the full path of the specified file.
var FullPath = syscall.FullPath

var Futimes = syscall.Futimes

var Futimes = syscall.Futimes

var Futimesat = syscall.Futimesat

var Fwstat = syscall.Fwstat

var Fwstat = syscall.Fwstat

var Fwstat = syscall.Fwstat

var GetAcceptExSockaddrs = syscall.GetAcceptExSockaddrs

var GetAdaptersInfo = syscall.GetAdaptersInfo

var GetAddrInfoW = syscall.GetAddrInfoW

var GetCommandLine = syscall.GetCommandLine

var GetComputerName = syscall.GetComputerName

var GetConsoleMode = syscall.GetConsoleMode

var GetCurrentDirectory = syscall.GetCurrentDirectory

var GetCurrentProcess = syscall.GetCurrentProcess

var GetEnvironmentStrings = syscall.GetEnvironmentStrings

var GetEnvironmentVariable = syscall.GetEnvironmentVariable

var GetExitCodeProcess = syscall.GetExitCodeProcess

var GetFileAttributes = syscall.GetFileAttributes

var GetFileAttributesEx = syscall.GetFileAttributesEx

var GetFileInformationByHandle = syscall.GetFileInformationByHandle

var GetFileType = syscall.GetFileType

var GetFullPathName = syscall.GetFullPathName

var GetHostByName = syscall.GetHostByName

var GetIfEntry = syscall.GetIfEntry

var GetLastError = syscall.GetLastError

var GetLengthSid = syscall.GetLengthSid

var GetLongPathName = syscall.GetLongPathName

var GetProcAddress = syscall.GetProcAddress

var GetProcessTimes = syscall.GetProcessTimes

var GetProtoByName = syscall.GetProtoByName

var GetQueuedCompletionStatus = syscall.GetQueuedCompletionStatus

var GetServByName = syscall.GetServByName

var GetShortPathName = syscall.GetShortPathName

var GetStartupInfo = syscall.GetStartupInfo

var GetStdHandle = syscall.GetStdHandle

var GetSystemTimeAsFileTime = syscall.GetSystemTimeAsFileTime

var GetTempPath = syscall.GetTempPath

var GetTimeZoneInformation = syscall.GetTimeZoneInformation

var GetTokenInformation = syscall.GetTokenInformation

var GetUserNameEx = syscall.GetUserNameEx

var GetUserProfileDirectory = syscall.GetUserProfileDirectory

var GetVersion = syscall.GetVersion

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getcwd = syscall.Getcwd

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdents = syscall.Getdents

var Getdirentries = syscall.Getdirentries

// sys getdents(fd int, buf []byte) (n int, err error)
var Getdirentries = syscall.Getdirentries

var Getdirentries = syscall.Getdirentries

var Getdirentries = syscall.Getdirentries

var Getdirentries = syscall.Getdirentries

var Getdirentries = syscall.Getdirentries

// sys getdents(fd int, buf []byte) (n int, err error)
var Getdirentries = syscall.Getdirentries

var Getdirentries = syscall.Getdirentries

var Getdtablesize = syscall.Getdtablesize

var Getdtablesize = syscall.Getdtablesize

var Getdtablesize = syscall.Getdtablesize

var Getdtablesize = syscall.Getdtablesize

var Getdtablesize = syscall.Getdtablesize

var Getdtablesize = syscall.Getdtablesize

var Getdtablesize = syscall.Getdtablesize

var Getdtablesize = syscall.Getdtablesize

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getegid = syscall.Getegid

var Getenv = syscall.Getenv

var Getenv = syscall.Getenv

var Getenv = syscall.Getenv

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Geteuid = syscall.Geteuid

var Getexecname = syscall.Getexecname

var Getfsstat = syscall.Getfsstat

var Getfsstat = syscall.Getfsstat

var Getfsstat = syscall.Getfsstat

var Getfsstat = syscall.Getfsstat

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgid = syscall.Getgid

var Getgroups = syscall.Getgroups

var Getgroups = syscall.Getgroups

var Getgroups = syscall.Getgroups

var Getgroups = syscall.Getgroups

var Getgroups = syscall.Getgroups

var Getgroups = syscall.Getgroups

var Getgroups = syscall.Getgroups

var Getgroups = syscall.Getgroups

var Gethostname = syscall.Gethostname

var Getkerninfo = syscall.Getkerninfo

var Getpagesize = syscall.Getpagesize

var Getpeername = syscall.Getpeername

var Getpeername = syscall.Getpeername

var Getpeername = syscall.Getpeername

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgid = syscall.Getpgid

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpgrp = syscall.Getpgrp

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getpid = syscall.Getpid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getppid = syscall.Getppid

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getpriority = syscall.Getpriority

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrlimit = syscall.Getrlimit

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getrusage = syscall.Getrusage

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsid = syscall.Getsid

var Getsockname = syscall.Getsockname

var Getsockname = syscall.Getsockname

var Getsockname = syscall.Getsockname

var Getsockname = syscall.Getsockname

var Getsockname = syscall.Getsockname

var Getsockname = syscall.Getsockname

var Getsockopt = syscall.Getsockopt

var GetsockoptByte = syscall.GetsockoptByte

var GetsockoptICMPv6Filter = syscall.GetsockoptICMPv6Filter

var GetsockoptICMPv6Filter = syscall.GetsockoptICMPv6Filter

var GetsockoptIPMreq = syscall.GetsockoptIPMreq

var GetsockoptIPMreq = syscall.GetsockoptIPMreq

var GetsockoptIPMreqn = syscall.GetsockoptIPMreqn

var GetsockoptIPMreqn = syscall.GetsockoptIPMreqn

var GetsockoptIPv6MTUInfo = syscall.GetsockoptIPv6MTUInfo

var GetsockoptIPv6MTUInfo = syscall.GetsockoptIPv6MTUInfo

var GetsockoptIPv6Mreq = syscall.GetsockoptIPv6Mreq

var GetsockoptIPv6Mreq = syscall.GetsockoptIPv6Mreq

var GetsockoptInet4Addr = syscall.GetsockoptInet4Addr

var GetsockoptInet4Addr = syscall.GetsockoptInet4Addr

var GetsockoptInt = syscall.GetsockoptInt

var GetsockoptInt = syscall.GetsockoptInt

var GetsockoptInt = syscall.GetsockoptInt

var GetsockoptInt = syscall.GetsockoptInt

var GetsockoptUcred = syscall.GetsockoptUcred

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettid = syscall.Gettid

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Gettimeofday = syscall.Gettimeofday

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getuid = syscall.Getuid

var Getwd = syscall.Getwd

var Getwd = syscall.Getwd

var Getwd = syscall.Getwd

var Getwd = syscall.Getwd

var Getwd = syscall.Getwd

var Getwd = syscall.Getwd

var Getwd = syscall.Getwd

var Getwd = syscall.Getwd

var Getwd = syscall.Getwd

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var Getxattr = syscall.Getxattr

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyAddWatch = syscall.InotifyAddWatch

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit = syscall.InotifyInit

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyInit1 = syscall.InotifyInit1

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var InotifyRmWatch = syscall.InotifyRmWatch

var Ioperm = syscall.Ioperm

var Ioperm = syscall.Ioperm

var Ioperm = syscall.Ioperm

var Ioperm = syscall.Ioperm

var Ioperm = syscall.Ioperm

var Ioperm = syscall.Ioperm

var Ioperm = syscall.Ioperm

var Iopl = syscall.Iopl

var Iopl = syscall.Iopl

var Iopl = syscall.Iopl

var Iopl = syscall.Iopl

var Iopl = syscall.Iopl

var Iopl = syscall.Iopl

var Iopl = syscall.Iopl

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Issetugid = syscall.Issetugid

var Kevent = syscall.Kevent

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Kill = syscall.Kill

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Klogctl = syscall.Klogctl

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Kqueue = syscall.Kqueue

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Lchown = syscall.Lchown

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Link = syscall.Link

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listen = syscall.Listen

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var Listxattr = syscall.Listxattr

var LoadCancelIoEx = syscall.LoadCancelIoEx

var LoadConnectEx = syscall.LoadConnectEx

var LoadCreateSymbolicLink = syscall.LoadCreateSymbolicLink

// LoadDLL loads the named DLL file into memory.
//
// If name is not an absolute path and is not a known system DLL used by
// Go, Windows will search for the named DLL in many locations, causing
// potential DLL preloading attacks.
//
// Use LazyDLL in golang.org/x/sys/windows for a secure way to
// load system DLLs.
var LoadDLL = syscall.LoadDLL

var LoadGetAddrInfo = syscall.LoadGetAddrInfo

var LoadLibrary = syscall.LoadLibrary

var LoadSetFileCompletionNotificationModes = syscall.LoadSetFileCompletionNotificationModes

var LocalFree = syscall.LocalFree

var LookupAccountName = syscall.LookupAccountName

var LookupAccountSid = syscall.LookupAccountSid

// LookupSID retrieves a security identifier sid for the account
// and the name of the domain on which the account was found.
// System specify target computer to search.
var LookupSID = syscall.LookupSID

// Deprecated: Use golang.org/x/net/bpf instead.
var LsfJump = syscall.LsfJump

// Deprecated: Use golang.org/x/net/bpf instead.
var LsfSocket = syscall.LsfSocket

// Deprecated: Use golang.org/x/net/bpf instead.
var LsfStmt = syscall.LsfStmt

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Lstat = syscall.Lstat

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var Madvise = syscall.Madvise

var MapViewOfFile = syscall.MapViewOfFile

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdir = syscall.Mkdir

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkdirat = syscall.Mkdirat

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mkfifo = syscall.Mkfifo

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknod = syscall.Mknod

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mknodat = syscall.Mknodat

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlock = syscall.Mlock

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mlockall = syscall.Mlockall

var Mmap = syscall.Mmap

var Mmap = syscall.Mmap

var Mmap = syscall.Mmap

var Mount = syscall.Mount

// sys	mount(fd int, afd int, old string, flag int, aname string) (err error)
var Mount = syscall.Mount

var MoveFile = syscall.MoveFile

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Mprotect = syscall.Mprotect

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlock = syscall.Munlock

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munlockall = syscall.Munlockall

var Munmap = syscall.Munmap

var Munmap = syscall.Munmap

var Munmap = syscall.Munmap

// MustLoadDLL is like LoadDLL but panics if load operation fails.
var MustLoadDLL = syscall.MustLoadDLL

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var Nanosleep = syscall.Nanosleep

var NetApiBufferFree = syscall.NetApiBufferFree

var NetGetJoinInformation = syscall.NetGetJoinInformation

var NetUserGetInfo = syscall.NetUserGetInfo

// NetlinkRIB returns routing information base, as known as RIB, which
// consists of network facility information, states and parameters.
var NetlinkRIB = syscall.NetlinkRIB

// NewCallback converts a Go function to a function pointer conforming to the stdcall calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
var NewCallback = syscall.NewCallback

// NewCallbackCDecl converts a Go function to a function pointer conforming to the cdecl calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
var NewCallbackCDecl = syscall.NewCallbackCDecl

// NewError converts s to an ErrorString, which satisfies the Error interface.
var NewError = syscall.NewError

// NewLazyDLL creates new LazyDLL associated with DLL file.
var NewLazyDLL = syscall.NewLazyDLL

var NsecToFiletime = syscall.NsecToFiletime

// NsecToTimespec takes a number of nanoseconds since the Unix epoch
// and returns the corresponding Timespec value.
var NsecToTimespec = syscall.NsecToTimespec

var NsecToTimespec = syscall.NsecToTimespec

// NsecToTimeval takes a number of nanoseconds since the Unix epoch
// and returns the corresponding Timeval value.
var NsecToTimeval = syscall.NsecToTimeval

var NsecToTimeval = syscall.NsecToTimeval

var NsecToTimeval = syscall.NsecToTimeval

var Ntohs = syscall.Ntohs

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

// sys	open(path string, mode int) (fd int, err error)
var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

var Open = syscall.Open

// OpenCurrentProcessToken opens the access token
// associated with current process.
var OpenCurrentProcessToken = syscall.OpenCurrentProcessToken

var OpenProcess = syscall.OpenProcess

var OpenProcessToken = syscall.OpenProcessToken

var Openat = syscall.Openat

var Openat = syscall.Openat

// ParseDirent parses up to max directory entries in buf,
// appending the names to names. It returns the number of
// bytes consumed from buf, the number of entries added
// to names, and the new names slice.
var ParseDirent = syscall.ParseDirent

// ParseFiles parses files listed in fs and extracts all syscall
// functions listed in sys comments. It returns source files
// and functions collection *Source if successful.
var ParseFiles = syscall.ParseFiles

// ParseNetlinkMessage parses b as an array of netlink messages and
// returns the slice containing the NetlinkMessage structures.
var ParseNetlinkMessage = syscall.ParseNetlinkMessage

// ParseNetlinkRouteAttr parses m&#39;s payload as an array of netlink
// route attributes and returns the slice containing the
// NetlinkRouteAttr structures.
var ParseNetlinkRouteAttr = syscall.ParseNetlinkRouteAttr

var ParseRoutingMessage = syscall.ParseRoutingMessage

// ParseRoutingMessage parses b as routing messages and returns the
// slice containing the RoutingMessage interfaces.
//
// Deprecated: Use golang.org/x/net/route instead.
var ParseRoutingMessage = syscall.ParseRoutingMessage

var ParseRoutingSockaddr = syscall.ParseRoutingSockaddr

// ParseRoutingSockaddr parses msg&#39;s payload as raw sockaddrs and
// returns the slice containing the Sockaddr interfaces.
//
// Deprecated: Use golang.org/x/net/route instead.
var ParseRoutingSockaddr = syscall.ParseRoutingSockaddr

// ParseSocketControlMessage parses b as an array of socket control
// messages.
var ParseSocketControlMessage = syscall.ParseSocketControlMessage

// ParseUnixCredentials decodes a socket control message that contains
// credentials in a Ucred structure. To receive such a message, the
// SO_PASSCRED option must be enabled on the socket.
var ParseUnixCredentials = syscall.ParseUnixCredentials

// ParseUnixRights decodes a socket control message that contains an
// integer array of open file descriptors from another process.
var ParseUnixRights = syscall.ParseUnixRights

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pathconf = syscall.Pathconf

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pause = syscall.Pause

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

// sys	pipe(p *[2]int32) (err error)
var Pipe = syscall.Pipe

// sysnb pipe(p *[2]_C_int) (err error)
var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

var Pipe = syscall.Pipe

// sysnb pipe2(p *[2]_C_int, flags int) (err error)
var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

// sysnb pipe2(p *[2]_C_int, flags int) (err error)
var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

var Pipe2 = syscall.Pipe2

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PivotRoot = syscall.PivotRoot

var PostQueuedCompletionStatus = syscall.PostQueuedCompletionStatus

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

// sys	extpread(fd int, p []byte, flags int, offset int64) (n int, err error)
var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Pread = syscall.Pread

var Process32First = syscall.Process32First

var Process32Next = syscall.Process32Next

var PtraceAttach = syscall.PtraceAttach

var PtraceAttach = syscall.PtraceAttach

// sys   ptrace(request int, pid int, addr uintptr, data uintptr) (err error)
var PtraceAttach = syscall.PtraceAttach

var PtraceCont = syscall.PtraceCont

var PtraceCont = syscall.PtraceCont

var PtraceDetach = syscall.PtraceDetach

var PtraceDetach = syscall.PtraceDetach

var PtraceDetach = syscall.PtraceDetach

var PtraceGetEventMsg = syscall.PtraceGetEventMsg

var PtraceGetRegs = syscall.PtraceGetRegs

var PtracePeekData = syscall.PtracePeekData

var PtracePeekData = syscall.PtracePeekData

var PtracePeekText = syscall.PtracePeekText

var PtracePeekText = syscall.PtracePeekText

var PtracePokeData = syscall.PtracePokeData

var PtracePokeData = syscall.PtracePokeData

var PtracePokeText = syscall.PtracePokeText

var PtracePokeText = syscall.PtracePokeText

var PtraceSetOptions = syscall.PtraceSetOptions

var PtraceSetRegs = syscall.PtraceSetRegs

var PtraceSingleStep = syscall.PtraceSingleStep

var PtraceSingleStep = syscall.PtraceSingleStep

var PtraceSyscall = syscall.PtraceSyscall

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

// sys	extpwrite(fd int, p []byte, flags int, offset int64) (n int, err error)
var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var Pwrite = syscall.Pwrite

var RawSyscall = syscall.RawSyscall

var RawSyscall = syscall.RawSyscall

var RawSyscall = syscall.RawSyscall

var RawSyscall = syscall.RawSyscall

var RawSyscall6 = syscall.RawSyscall6

var RawSyscall6 = syscall.RawSyscall6

var RawSyscall6 = syscall.RawSyscall6

var RawSyscall6 = syscall.RawSyscall6

var Read = syscall.Read

var Read = syscall.Read

var Read = syscall.Read

var Read = syscall.Read

var Read = syscall.Read

var ReadConsole = syscall.ReadConsole

var ReadDirectoryChanges = syscall.ReadDirectoryChanges

// sys	getdirent(fd int, buf []byte) (n int, err error)
var ReadDirent = syscall.ReadDirent

var ReadDirent = syscall.ReadDirent

var ReadDirent = syscall.ReadDirent

var ReadDirent = syscall.ReadDirent

var ReadDirent = syscall.ReadDirent

var ReadDirent = syscall.ReadDirent

var ReadFile = syscall.ReadFile

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

// sys	readlink(path string, buf []byte, bufSize uint64) (n int, err error)
var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

// Readlink returns the destination of the named symbolic link.
var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Readlink = syscall.Readlink

var Reboot = syscall.Reboot

var Reboot = syscall.Reboot

var Recvfrom = syscall.Recvfrom

var Recvfrom = syscall.Recvfrom

var Recvfrom = syscall.Recvfrom

var Recvfrom = syscall.Recvfrom

var Recvmsg = syscall.Recvmsg

var Recvmsg = syscall.Recvmsg

var Recvmsg = syscall.Recvmsg

var Recvmsg = syscall.Recvmsg

var Recvmsg = syscall.Recvmsg

var Recvmsg = syscall.Recvmsg

var RegCloseKey = syscall.RegCloseKey

var RegEnumKeyEx = syscall.RegEnumKeyEx

var RegOpenKeyEx = syscall.RegOpenKeyEx

var RegQueryInfoKey = syscall.RegQueryInfoKey

var RegQueryValueEx = syscall.RegQueryValueEx

// sys	remove(path string) (err error)
var Remove = syscall.Remove

var RemoveDirectory = syscall.RemoveDirectory

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Removexattr = syscall.Removexattr

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Rename = syscall.Rename

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Renameat = syscall.Renameat

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Revoke = syscall.Revoke

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var Rmdir = syscall.Rmdir

var RouteRIB = syscall.RouteRIB

// RouteRIB returns routing information base, as known as RIB,
// which consists of network facility information, states and
// parameters.
//
// Deprecated: Use golang.org/x/net/route instead.
var RouteRIB = syscall.RouteRIB

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Seek = syscall.Seek

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Select = syscall.Select

var Sendfile = syscall.Sendfile

var Sendfile = syscall.Sendfile

var Sendfile = syscall.Sendfile

var Sendmsg = syscall.Sendmsg

var Sendmsg = syscall.Sendmsg

var Sendmsg = syscall.Sendmsg

var Sendmsg = syscall.Sendmsg

var Sendmsg = syscall.Sendmsg

var SendmsgN = syscall.SendmsgN

var SendmsgN = syscall.SendmsgN

var SendmsgN = syscall.SendmsgN

var SendmsgN = syscall.SendmsgN

var SendmsgN = syscall.SendmsgN

var SendmsgN = syscall.SendmsgN

var Sendto = syscall.Sendto

var Sendto = syscall.Sendto

var Sendto = syscall.Sendto

var Sendto = syscall.Sendto

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpf = syscall.SetBpf

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpf = syscall.SetBpf

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfBuflen = syscall.SetBpfBuflen

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfBuflen = syscall.SetBpfBuflen

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfDatalink = syscall.SetBpfDatalink

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfDatalink = syscall.SetBpfDatalink

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfHeadercmpl = syscall.SetBpfHeadercmpl

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfHeadercmpl = syscall.SetBpfHeadercmpl

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfImmediate = syscall.SetBpfImmediate

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfImmediate = syscall.SetBpfImmediate

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfInterface = syscall.SetBpfInterface

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfInterface = syscall.SetBpfInterface

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfPromisc = syscall.SetBpfPromisc

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfPromisc = syscall.SetBpfPromisc

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfTimeout = syscall.SetBpfTimeout

// Deprecated: Use golang.org/x/net/bpf instead.
var SetBpfTimeout = syscall.SetBpfTimeout

var SetCurrentDirectory = syscall.SetCurrentDirectory

var SetEndOfFile = syscall.SetEndOfFile

var SetEnvironmentVariable = syscall.SetEnvironmentVariable

var SetFileAttributes = syscall.SetFileAttributes

var SetFileCompletionNotificationModes = syscall.SetFileCompletionNotificationModes

var SetFilePointer = syscall.SetFilePointer

var SetFileTime = syscall.SetFileTime

var SetHandleInformation = syscall.SetHandleInformation

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

var SetKevent = syscall.SetKevent

// Deprecated: Use golang.org/x/net/bpf instead.
var SetLsfPromisc = syscall.SetLsfPromisc

var SetNonblock = syscall.SetNonblock

var SetNonblock = syscall.SetNonblock

var SetNonblock = syscall.SetNonblock

var SetNonblock = syscall.SetNonblock

var SetReadDeadline = syscall.SetReadDeadline

var SetReadDeadline = syscall.SetReadDeadline

var SetWriteDeadline = syscall.SetWriteDeadline

var SetWriteDeadline = syscall.SetWriteDeadline

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setdomainname = syscall.Setdomainname

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setegid = syscall.Setegid

var Setenv = syscall.Setenv

var Setenv = syscall.Setenv

var Setenv = syscall.Setenv

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Seteuid = syscall.Seteuid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsgid = syscall.Setfsgid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setfsuid = syscall.Setfsuid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgid = syscall.Setgid

var Setgroups = syscall.Setgroups

var Setgroups = syscall.Setgroups

var Setgroups = syscall.Setgroups

var Setgroups = syscall.Setgroups

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Sethostname = syscall.Sethostname

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setlogin = syscall.Setlogin

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpgid = syscall.Setpgid

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setpriority = syscall.Setpriority

var Setprivexec = syscall.Setprivexec

var Setprivexec = syscall.Setprivexec

var Setprivexec = syscall.Setprivexec

var Setprivexec = syscall.Setprivexec

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setregid = syscall.Setregid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresgid = syscall.Setresgid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setresuid = syscall.Setresuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setreuid = syscall.Setreuid

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setrlimit = syscall.Setrlimit

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsid = syscall.Setsid

var Setsockopt = syscall.Setsockopt

var SetsockoptByte = syscall.SetsockoptByte

var SetsockoptByte = syscall.SetsockoptByte

var SetsockoptICMPv6Filter = syscall.SetsockoptICMPv6Filter

var SetsockoptICMPv6Filter = syscall.SetsockoptICMPv6Filter

var SetsockoptIPMreq = syscall.SetsockoptIPMreq

var SetsockoptIPMreq = syscall.SetsockoptIPMreq

var SetsockoptIPMreq = syscall.SetsockoptIPMreq

var SetsockoptIPMreqn = syscall.SetsockoptIPMreqn

var SetsockoptIPMreqn = syscall.SetsockoptIPMreqn

var SetsockoptIPv6Mreq = syscall.SetsockoptIPv6Mreq

var SetsockoptIPv6Mreq = syscall.SetsockoptIPv6Mreq

var SetsockoptIPv6Mreq = syscall.SetsockoptIPv6Mreq

var SetsockoptInet4Addr = syscall.SetsockoptInet4Addr

var SetsockoptInet4Addr = syscall.SetsockoptInet4Addr

var SetsockoptInet4Addr = syscall.SetsockoptInet4Addr

var SetsockoptInt = syscall.SetsockoptInt

var SetsockoptInt = syscall.SetsockoptInt

var SetsockoptInt = syscall.SetsockoptInt

var SetsockoptInt = syscall.SetsockoptInt

var SetsockoptLinger = syscall.SetsockoptLinger

var SetsockoptLinger = syscall.SetsockoptLinger

var SetsockoptLinger = syscall.SetsockoptLinger

var SetsockoptString = syscall.SetsockoptString

var SetsockoptString = syscall.SetsockoptString

var SetsockoptTimeval = syscall.SetsockoptTimeval

var SetsockoptTimeval = syscall.SetsockoptTimeval

var SetsockoptTimeval = syscall.SetsockoptTimeval

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Settimeofday = syscall.Settimeofday

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setuid = syscall.Setuid

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Setxattr = syscall.Setxattr

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

var Shutdown = syscall.Shutdown

// SlicePtrFromStrings converts a slice of strings to a slice of
// pointers to NUL-terminated byte arrays. If any string contains
// a NUL byte, it returns (nil, EINVAL).
var SlicePtrFromStrings = syscall.SlicePtrFromStrings

// SlicePtrFromStrings converts a slice of strings to a slice of
// pointers to NUL-terminated byte arrays. If any string contains
// a NUL byte, it returns (nil, EINVAL).
var SlicePtrFromStrings = syscall.SlicePtrFromStrings

var Socket = syscall.Socket

var Socket = syscall.Socket

var Socket = syscall.Socket

var Socket = syscall.Socket

var Socketpair = syscall.Socketpair

var Socketpair = syscall.Socketpair

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var Splice = syscall.Splice

var StartProcess = syscall.StartProcess

var StartProcess = syscall.StartProcess

// StartProcess wraps ForkExec for package os.
var StartProcess = syscall.StartProcess

var StartProcess = syscall.StartProcess

// StartProcess wraps ForkExec for package os.
var StartProcess = syscall.StartProcess

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

// sys	stat(path string, edir []byte) (n int, err error)
var Stat = syscall.Stat

var Stat = syscall.Stat

var Stat = syscall.Stat

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var Statfs = syscall.Statfs

var StopIO = syscall.StopIO

var StopIO = syscall.StopIO

// StringBytePtr returns a pointer to a NUL-terminated array of bytes.
// If s contains a NUL byte this function panics instead of returning
// an error.
//
// Deprecated: Use BytePtrFromString instead.
var StringBytePtr = syscall.StringBytePtr

// StringByteSlice converts a string to a NUL-terminated []byte,
// If s contains a NUL byte this function panics instead of
// returning an error.
//
// Deprecated: Use ByteSliceFromString instead.
var StringByteSlice = syscall.StringByteSlice

// StringSlicePtr converts a slice of strings to a slice of pointers
// to NUL-terminated byte arrays. If any string contains a NUL byte
// this function panics instead of returning an error.
//
// Deprecated: Use SlicePtrFromStrings instead.
var StringSlicePtr = syscall.StringSlicePtr

// StringSlicePtr converts a slice of strings to a slice of pointers
// to NUL-terminated byte arrays. If any string contains a NUL byte
// this function panics instead of returning an error.
//
// Deprecated: Use SlicePtrFromStrings instead.
var StringSlicePtr = syscall.StringSlicePtr

// StringToSid converts a string-format security identifier
// sid into a valid, functional sid.
var StringToSid = syscall.StringToSid

// StringToUTF16 returns the UTF-16 encoding of the UTF-8 string s,
// with a terminating NUL added. If s contains a NUL byte this
// function panics instead of returning an error.
//
// Deprecated: Use UTF16FromString instead.
var StringToUTF16 = syscall.StringToUTF16

// StringToUTF16Ptr returns pointer to the UTF-16 encoding of
// the UTF-8 string s, with a terminating NUL added. If s
// contains a NUL byte this function panics instead of
// returning an error.
//
// Deprecated: Use UTF16PtrFromString instead.
var StringToUTF16Ptr = syscall.StringToUTF16Ptr

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Symlink = syscall.Symlink

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var Sync = syscall.Sync

var SyncFileRange = syscall.SyncFileRange

var SyncFileRange = syscall.SyncFileRange

var SyncFileRange = syscall.SyncFileRange

var SyncFileRange = syscall.SyncFileRange

var SyncFileRange = syscall.SyncFileRange

var SyncFileRange = syscall.SyncFileRange

var SyncFileRange = syscall.SyncFileRange

var SyncFileRange = syscall.SyncFileRange

var SyncFileRange = syscall.SyncFileRange

// Implemented in ../runtime/syscall_windows.go.
var Syscall = syscall.Syscall

var Syscall = syscall.Syscall

var Syscall = syscall.Syscall

var Syscall = syscall.Syscall

var Syscall = syscall.Syscall

var Syscall12 = syscall.Syscall12

var Syscall15 = syscall.Syscall15

var Syscall18 = syscall.Syscall18

var Syscall6 = syscall.Syscall6

var Syscall6 = syscall.Syscall6

var Syscall6 = syscall.Syscall6

var Syscall6 = syscall.Syscall6

var Syscall6 = syscall.Syscall6

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Syscall9 = syscall.Syscall9

var Sysctl = syscall.Sysctl

var Sysctl = syscall.Sysctl

var Sysctl = syscall.Sysctl

var SysctlUint32 = syscall.SysctlUint32

var SysctlUint32 = syscall.SysctlUint32

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Sysinfo = syscall.Sysinfo

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var Tee = syscall.Tee

var TerminateProcess = syscall.TerminateProcess

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Tgkill = syscall.Tgkill

var Time = syscall.Time

var Time = syscall.Time

var Time = syscall.Time

var Time = syscall.Time

var Time = syscall.Time

var Time = syscall.Time

var Time = syscall.Time

var Time = syscall.Time

var Time = syscall.Time

var Time = syscall.Time

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var Times = syscall.Times

var TimespecToNsec = syscall.TimespecToNsec

// TimespecToNsec converts a Timespec value into a number of
// nanoseconds since the Unix epoch.
var TimespecToNsec = syscall.TimespecToNsec

// TimevalToNsec converts a Timeval value into a number of nanoseconds
// since the Unix epoch.
var TimevalToNsec = syscall.TimevalToNsec

// TranslateAccountName converts a directory service
// object name from one format to another.
var TranslateAccountName = syscall.TranslateAccountName

var TranslateName = syscall.TranslateName

var TransmitFile = syscall.TransmitFile

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

var Truncate = syscall.Truncate

// UTF16FromString returns the UTF-16 encoding of the UTF-8 string
// s, with a terminating NUL added. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
var UTF16FromString = syscall.UTF16FromString

// UTF16PtrFromString returns pointer to the UTF-16 encoding of
// the UTF-8 string s, with a terminating NUL added. If s
// contains a NUL byte at any location, it returns (nil, EINVAL).
var UTF16PtrFromString = syscall.UTF16PtrFromString

// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL removed.
var UTF16ToString = syscall.UTF16ToString

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Umask = syscall.Umask

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Uname = syscall.Uname

var Undelete = syscall.Undelete

var Undelete = syscall.Undelete

var Undelete = syscall.Undelete

var Undelete = syscall.Undelete

var Undelete = syscall.Undelete

var Undelete = syscall.Undelete

var Undelete = syscall.Undelete

var Undelete = syscall.Undelete

// UnixCredentials encodes credentials into a socket control message
// for sending to another process. This can be used for
// authentication.
var UnixCredentials = syscall.UnixCredentials

// UnixRights encodes a set of open file descriptors into a socket
// control message for sending to another process.
var UnixRights = syscall.UnixRights

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlink = syscall.Unlink

var Unlinkat = syscall.Unlinkat

// sys	unlinkat(dirfd int, path string, flags int) (err error)
var Unlinkat = syscall.Unlinkat

var UnmapViewOfFile = syscall.UnmapViewOfFile

// UnmarshalDir decodes a single 9P stat message from b and returns the resulting Dir.
//
// If b is too small to hold a valid stat message, ErrShortStat is returned.
//
// If the stat message itself is invalid, ErrBadStat is returned.
var UnmarshalDir = syscall.UnmarshalDir

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unmount = syscall.Unmount

var Unsetenv = syscall.Unsetenv

var Unsetenv = syscall.Unsetenv

var Unsetenv = syscall.Unsetenv

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Unshare = syscall.Unshare

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Ustat = syscall.Ustat

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utime = syscall.Utime

var Utimes = syscall.Utimes

var Utimes = syscall.Utimes

var Utimes = syscall.Utimes

// sys	utimes(path string, times *[2]Timeval) (err error)
var Utimes = syscall.Utimes

var Utimes = syscall.Utimes

var UtimesNano = syscall.UtimesNano

var UtimesNano = syscall.UtimesNano

// sys	utimensat(dirfd int, path string, times *[2]Timespec, flag int) (err error)
var UtimesNano = syscall.UtimesNano

var UtimesNano = syscall.UtimesNano

var UtimesNano = syscall.UtimesNano

var UtimesNano = syscall.UtimesNano

var UtimesNano = syscall.UtimesNano

var VirtualLock = syscall.VirtualLock

var VirtualUnlock = syscall.VirtualUnlock

var WSACleanup = syscall.WSACleanup

var WSAEnumProtocols = syscall.WSAEnumProtocols

var WSAIoctl = syscall.WSAIoctl

var WSARecv = syscall.WSARecv

var WSARecvFrom = syscall.WSARecvFrom

var WSASend = syscall.WSASend

var WSASendTo = syscall.WSASendTo

var WSASendto = syscall.WSASendto

var WSAStartup = syscall.WSAStartup

var Wait4 = syscall.Wait4

var Wait4 = syscall.Wait4

var Wait4 = syscall.Wait4

var Wait4 = syscall.Wait4

// sys  wait4(pid _Pid_t, status *_C_int, options int, rusage *Rusage) (wpid _Pid_t, err error)
var Wait4 = syscall.Wait4

var Wait4 = syscall.Wait4

var WaitForSingleObject = syscall.WaitForSingleObject

// WaitProcess waits until the pid of a
// running process is found in the queue of
// wait messages. It is used in conjunction
// with ForkExec/StartProcess to wait for a
// running process to exit.
var WaitProcess = syscall.WaitProcess

var Write = syscall.Write

var Write = syscall.Write

var Write = syscall.Write

var Write = syscall.Write

var Write = syscall.Write

var WriteConsole = syscall.WriteConsole

var WriteFile = syscall.WriteFile

// sys	wstat(path string, edir []byte) (err error)
var Wstat = syscall.Wstat
