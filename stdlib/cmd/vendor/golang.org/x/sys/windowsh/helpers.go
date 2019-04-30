package windowsh

import (
	"cmd/vendor/golang.org/x/sys/windows"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	AcceptKey = "Accept"

	AcceptExKey = "AcceptEx"

	AllocateAndInitializeSidKey = "AllocateAndInitializeSid"

	BindKey = "Bind"

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

	ChangeServiceConfigKey = "ChangeServiceConfig"

	ChangeServiceConfig2Key = "ChangeServiceConfig2"

	ChdirKey = "Chdir"

	ChmodKey = "Chmod"

	ChownKey = "Chown"

	ClearenvKey = "Clearenv"

	CloseKey = "Close"

	CloseHandleKey = "CloseHandle"

	CloseOnExecKey = "CloseOnExec"

	CloseServiceHandleKey = "CloseServiceHandle"

	ClosesocketKey = "Closesocket"

	CommandLineToArgvKey = "CommandLineToArgv"

	ComputerNameKey = "ComputerName"

	ConnectKey = "Connect"

	ConnectExKey = "ConnectEx"

	ControlServiceKey = "ControlService"

	ConvertSidToStringSidKey = "ConvertSidToStringSid"

	ConvertStringSidToSidKey = "ConvertStringSidToSid"

	CopySidKey = "CopySid"

	CreateDirectoryKey = "CreateDirectory"

	CreateEventKey = "CreateEvent"

	CreateEventExKey = "CreateEventEx"

	CreateFileKey = "CreateFile"

	CreateFileMappingKey = "CreateFileMapping"

	CreateHardLinkKey = "CreateHardLink"

	CreateIoCompletionPortKey = "CreateIoCompletionPort"

	CreatePipeKey = "CreatePipe"

	CreateProcessKey = "CreateProcess"

	CreateServiceKey = "CreateService"

	CreateSymbolicLinkKey = "CreateSymbolicLink"

	CreateToolhelp32SnapshotKey = "CreateToolhelp32Snapshot"

	CryptAcquireContextKey = "CryptAcquireContext"

	CryptGenRandomKey = "CryptGenRandom"

	CryptReleaseContextKey = "CryptReleaseContext"

	DefineDosDeviceKey = "DefineDosDevice"

	DeleteFileKey = "DeleteFile"

	DeleteServiceKey = "DeleteService"

	DeleteVolumeMountPointKey = "DeleteVolumeMountPoint"

	DeregisterEventSourceKey = "DeregisterEventSource"

	DeviceIoControlKey = "DeviceIoControl"

	DnsNameCompareKey = "DnsNameCompare"

	DnsQueryKey = "DnsQuery"

	DnsRecordListFreeKey = "DnsRecordListFree"

	DuplicateHandleKey = "DuplicateHandle"

	EnumServicesStatusExKey = "EnumServicesStatusEx"

	EnvironKey = "Environ"

	EqualSidKey = "EqualSid"

	EscapeArgKey = "EscapeArg"

	ExitKey = "Exit"

	ExitProcessKey = "ExitProcess"

	FchdirKey = "Fchdir"

	FchmodKey = "Fchmod"

	FchownKey = "Fchown"

	FindCloseKey = "FindClose"

	FindFirstFileKey = "FindFirstFile"

	FindFirstVolumeKey = "FindFirstVolume"

	FindFirstVolumeMountPointKey = "FindFirstVolumeMountPoint"

	FindNextFileKey = "FindNextFile"

	FindNextVolumeKey = "FindNextVolume"

	FindNextVolumeMountPointKey = "FindNextVolumeMountPoint"

	FindVolumeCloseKey = "FindVolumeClose"

	FindVolumeMountPointCloseKey = "FindVolumeMountPointClose"

	FlushFileBuffersKey = "FlushFileBuffers"

	FlushViewOfFileKey = "FlushViewOfFile"

	FormatMessageKey = "FormatMessage"

	FreeAddrInfoWKey = "FreeAddrInfoW"

	FreeEnvironmentStringsKey = "FreeEnvironmentStrings"

	FreeLibraryKey = "FreeLibrary"

	FreeSidKey = "FreeSid"

	FsyncKey = "Fsync"

	FtruncateKey = "Ftruncate"

	FullPathKey = "FullPath"

	GetACPKey = "GetACP"

	GetAcceptExSockaddrsKey = "GetAcceptExSockaddrs"

	GetAdaptersAddressesKey = "GetAdaptersAddresses"

	GetAdaptersInfoKey = "GetAdaptersInfo"

	GetAddrInfoWKey = "GetAddrInfoW"

	GetCommandLineKey = "GetCommandLine"

	GetComputerNameKey = "GetComputerName"

	GetComputerNameExKey = "GetComputerNameEx"

	GetConsoleModeKey = "GetConsoleMode"

	GetConsoleScreenBufferInfoKey = "GetConsoleScreenBufferInfo"

	GetCurrentDirectoryKey = "GetCurrentDirectory"

	GetCurrentProcessKey = "GetCurrentProcess"

	GetCurrentThreadIdKey = "GetCurrentThreadId"

	GetDriveTypeKey = "GetDriveType"

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

	GetLogicalDriveStringsKey = "GetLogicalDriveStrings"

	GetLogicalDrivesKey = "GetLogicalDrives"

	GetLongPathNameKey = "GetLongPathName"

	GetProcAddressKey = "GetProcAddress"

	GetProcAddressByOrdinalKey = "GetProcAddressByOrdinal"

	GetProcessTimesKey = "GetProcessTimes"

	GetProtoByNameKey = "GetProtoByName"

	GetQueuedCompletionStatusKey = "GetQueuedCompletionStatus"

	GetServByNameKey = "GetServByName"

	GetShortPathNameKey = "GetShortPathName"

	GetStartupInfoKey = "GetStartupInfo"

	GetStdHandleKey = "GetStdHandle"

	GetSystemTimeAsFileTimeKey = "GetSystemTimeAsFileTime"

	GetSystemTimePreciseAsFileTimeKey = "GetSystemTimePreciseAsFileTime"

	GetTempPathKey = "GetTempPath"

	GetTimeZoneInformationKey = "GetTimeZoneInformation"

	GetTokenInformationKey = "GetTokenInformation"

	GetUserNameExKey = "GetUserNameEx"

	GetUserProfileDirectoryKey = "GetUserProfileDirectory"

	GetVersionKey = "GetVersion"

	GetVolumeInformationKey = "GetVolumeInformation"

	GetVolumeInformationByHandleKey = "GetVolumeInformationByHandle"

	GetVolumeNameForVolumeMountPointKey = "GetVolumeNameForVolumeMountPoint"

	GetVolumePathNameKey = "GetVolumePathName"

	GetVolumePathNamesForVolumeNameKey = "GetVolumePathNamesForVolumeName"

	GetegidKey = "Getegid"

	GetenvKey = "Getenv"

	GeteuidKey = "Geteuid"

	GetgidKey = "Getgid"

	GetgroupsKey = "Getgroups"

	GetpagesizeKey = "Getpagesize"

	GetpeernameKey = "Getpeername"

	GetpidKey = "Getpid"

	GetppidKey = "Getppid"

	GetsocknameKey = "Getsockname"

	GetsockoptKey = "Getsockopt"

	GetsockoptIntKey = "GetsockoptInt"

	GettimeofdayKey = "Gettimeofday"

	GetuidKey = "Getuid"

	GetwdKey = "Getwd"

	LchownKey = "Lchown"

	LinkKey = "Link"

	ListenKey = "Listen"

	LoadCancelIoExKey = "LoadCancelIoEx"

	LoadConnectExKey = "LoadConnectEx"

	LoadCreateSymbolicLinkKey = "LoadCreateSymbolicLink"

	LoadDLLKey = "LoadDLL"

	LoadGetAddrInfoKey = "LoadGetAddrInfo"

	LoadGetSystemTimePreciseAsFileTimeKey = "LoadGetSystemTimePreciseAsFileTime"

	LoadLibraryKey = "LoadLibrary"

	LoadLibraryExKey = "LoadLibraryEx"

	LoadSetFileCompletionNotificationModesKey = "LoadSetFileCompletionNotificationModes"

	LocalFreeKey = "LocalFree"

	LookupAccountNameKey = "LookupAccountName"

	LookupAccountSidKey = "LookupAccountSid"

	LookupSIDKey = "LookupSID"

	MapViewOfFileKey = "MapViewOfFile"

	MkdirKey = "Mkdir"

	MoveFileKey = "MoveFile"

	MoveFileExKey = "MoveFileEx"

	MultiByteToWideCharKey = "MultiByteToWideChar"

	MustLoadDLLKey = "MustLoadDLL"

	NetApiBufferFreeKey = "NetApiBufferFree"

	NetGetJoinInformationKey = "NetGetJoinInformation"

	NetUserGetInfoKey = "NetUserGetInfo"

	NewCallbackKey = "NewCallback"

	NewCallbackCDeclKey = "NewCallbackCDecl"

	NewLazyDLLKey = "NewLazyDLL"

	NewLazySystemDLLKey = "NewLazySystemDLL"

	NsecToFiletimeKey = "NsecToFiletime"

	NsecToTimespecKey = "NsecToTimespec"

	NsecToTimevalKey = "NsecToTimeval"

	NtohsKey = "Ntohs"

	OpenKey = "Open"

	OpenCurrentProcessTokenKey = "OpenCurrentProcessToken"

	OpenEventKey = "OpenEvent"

	OpenProcessKey = "OpenProcess"

	OpenProcessTokenKey = "OpenProcessToken"

	OpenSCManagerKey = "OpenSCManager"

	OpenServiceKey = "OpenService"

	PipeKey = "Pipe"

	PostQueuedCompletionStatusKey = "PostQueuedCompletionStatus"

	Process32FirstKey = "Process32First"

	Process32NextKey = "Process32Next"

	PulseEventKey = "PulseEvent"

	QueryDosDeviceKey = "QueryDosDevice"

	QueryServiceConfigKey = "QueryServiceConfig"

	QueryServiceConfig2Key = "QueryServiceConfig2"

	QueryServiceStatusKey = "QueryServiceStatus"

	QueryServiceStatusExKey = "QueryServiceStatusEx"

	ReadKey = "Read"

	ReadConsoleKey = "ReadConsole"

	ReadDirectoryChangesKey = "ReadDirectoryChanges"

	ReadFileKey = "ReadFile"

	ReadlinkKey = "Readlink"

	RecvfromKey = "Recvfrom"

	RegCloseKeyKey = "RegCloseKey"

	RegEnumKeyExKey = "RegEnumKeyEx"

	RegOpenKeyExKey = "RegOpenKeyEx"

	RegQueryInfoKeyKey = "RegQueryInfoKey"

	RegQueryValueExKey = "RegQueryValueEx"

	RegisterEventSourceKey = "RegisterEventSource"

	RemoveDirectoryKey = "RemoveDirectory"

	RenameKey = "Rename"

	ReportEventKey = "ReportEvent"

	ResetEventKey = "ResetEvent"

	RmdirKey = "Rmdir"

	SeekKey = "Seek"

	SendtoKey = "Sendto"

	SetConsoleModeKey = "SetConsoleMode"

	SetCurrentDirectoryKey = "SetCurrentDirectory"

	SetEndOfFileKey = "SetEndOfFile"

	SetEnvironmentVariableKey = "SetEnvironmentVariable"

	SetEventKey = "SetEvent"

	SetFileAttributesKey = "SetFileAttributes"

	SetFileCompletionNotificationModesKey = "SetFileCompletionNotificationModes"

	SetFilePointerKey = "SetFilePointer"

	SetFileTimeKey = "SetFileTime"

	SetHandleInformationKey = "SetHandleInformation"

	SetServiceStatusKey = "SetServiceStatus"

	SetStdHandleKey = "SetStdHandle"

	SetVolumeLabelKey = "SetVolumeLabel"

	SetVolumeMountPointKey = "SetVolumeMountPoint"

	SetenvKey = "Setenv"

	SetsockoptKey = "Setsockopt"

	SetsockoptIPMreqKey = "SetsockoptIPMreq"

	SetsockoptIPv6MreqKey = "SetsockoptIPv6Mreq"

	SetsockoptInet4AddrKey = "SetsockoptInet4Addr"

	SetsockoptIntKey = "SetsockoptInt"

	SetsockoptLingerKey = "SetsockoptLinger"

	SetsockoptTimevalKey = "SetsockoptTimeval"

	ShutdownKey = "Shutdown"

	SocketKey = "Socket"

	StartServiceKey = "StartService"

	StartServiceCtrlDispatcherKey = "StartServiceCtrlDispatcher"

	StringToSidKey = "StringToSid"

	StringToUTF16Key = "StringToUTF16"

	StringToUTF16PtrKey = "StringToUTF16Ptr"

	SymlinkKey = "Symlink"

	TerminateProcessKey = "TerminateProcess"

	TimespecToNsecKey = "TimespecToNsec"

	TranslateAccountNameKey = "TranslateAccountName"

	TranslateNameKey = "TranslateName"

	TransmitFileKey = "TransmitFile"

	UTF16FromStringKey = "UTF16FromString"

	UTF16PtrFromStringKey = "UTF16PtrFromString"

	UTF16ToStringKey = "UTF16ToString"

	UnlinkKey = "Unlink"

	UnmapViewOfFileKey = "UnmapViewOfFile"

	UnsetenvKey = "Unsetenv"

	UtimesKey = "Utimes"

	UtimesNanoKey = "UtimesNano"

	VirtualAllocKey = "VirtualAlloc"

	VirtualFreeKey = "VirtualFree"

	VirtualLockKey = "VirtualLock"

	VirtualProtectKey = "VirtualProtect"

	VirtualUnlockKey = "VirtualUnlock"

	WSACleanupKey = "WSACleanup"

	WSAEnumProtocolsKey = "WSAEnumProtocols"

	WSAIoctlKey = "WSAIoctl"

	WSARecvKey = "WSARecv"

	WSARecvFromKey = "WSARecvFrom"

	WSARecvMsgKey = "WSARecvMsg"

	WSASendKey = "WSASend"

	WSASendMsgKey = "WSASendMsg"

	WSASendToKey = "WSASendTo"

	WSASendtoKey = "WSASendto"

	WSAStartupKey = "WSAStartup"

	WaitForSingleObjectKey = "WaitForSingleObject"

	WriteKey = "Write"

	WriteConsoleKey = "WriteConsole"

	WriteFileKey = "WriteFile"
)

func New() hctx.Map {
	return hctx.Map{

		AcceptKey: Accept,

		AcceptExKey: AcceptEx,

		AllocateAndInitializeSidKey: AllocateAndInitializeSid,

		BindKey: Bind,

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

		ChangeServiceConfigKey: ChangeServiceConfig,

		ChangeServiceConfig2Key: ChangeServiceConfig2,

		ChdirKey: Chdir,

		ChmodKey: Chmod,

		ChownKey: Chown,

		ClearenvKey: Clearenv,

		CloseKey: Close,

		CloseHandleKey: CloseHandle,

		CloseOnExecKey: CloseOnExec,

		CloseServiceHandleKey: CloseServiceHandle,

		ClosesocketKey: Closesocket,

		CommandLineToArgvKey: CommandLineToArgv,

		ComputerNameKey: ComputerName,

		ConnectKey: Connect,

		ConnectExKey: ConnectEx,

		ControlServiceKey: ControlService,

		ConvertSidToStringSidKey: ConvertSidToStringSid,

		ConvertStringSidToSidKey: ConvertStringSidToSid,

		CopySidKey: CopySid,

		CreateDirectoryKey: CreateDirectory,

		CreateEventKey: CreateEvent,

		CreateEventExKey: CreateEventEx,

		CreateFileKey: CreateFile,

		CreateFileMappingKey: CreateFileMapping,

		CreateHardLinkKey: CreateHardLink,

		CreateIoCompletionPortKey: CreateIoCompletionPort,

		CreatePipeKey: CreatePipe,

		CreateProcessKey: CreateProcess,

		CreateServiceKey: CreateService,

		CreateSymbolicLinkKey: CreateSymbolicLink,

		CreateToolhelp32SnapshotKey: CreateToolhelp32Snapshot,

		CryptAcquireContextKey: CryptAcquireContext,

		CryptGenRandomKey: CryptGenRandom,

		CryptReleaseContextKey: CryptReleaseContext,

		DefineDosDeviceKey: DefineDosDevice,

		DeleteFileKey: DeleteFile,

		DeleteServiceKey: DeleteService,

		DeleteVolumeMountPointKey: DeleteVolumeMountPoint,

		DeregisterEventSourceKey: DeregisterEventSource,

		DeviceIoControlKey: DeviceIoControl,

		DnsNameCompareKey: DnsNameCompare,

		DnsQueryKey: DnsQuery,

		DnsRecordListFreeKey: DnsRecordListFree,

		DuplicateHandleKey: DuplicateHandle,

		EnumServicesStatusExKey: EnumServicesStatusEx,

		EnvironKey: Environ,

		EqualSidKey: EqualSid,

		EscapeArgKey: EscapeArg,

		ExitKey: Exit,

		ExitProcessKey: ExitProcess,

		FchdirKey: Fchdir,

		FchmodKey: Fchmod,

		FchownKey: Fchown,

		FindCloseKey: FindClose,

		FindFirstFileKey: FindFirstFile,

		FindFirstVolumeKey: FindFirstVolume,

		FindFirstVolumeMountPointKey: FindFirstVolumeMountPoint,

		FindNextFileKey: FindNextFile,

		FindNextVolumeKey: FindNextVolume,

		FindNextVolumeMountPointKey: FindNextVolumeMountPoint,

		FindVolumeCloseKey: FindVolumeClose,

		FindVolumeMountPointCloseKey: FindVolumeMountPointClose,

		FlushFileBuffersKey: FlushFileBuffers,

		FlushViewOfFileKey: FlushViewOfFile,

		FormatMessageKey: FormatMessage,

		FreeAddrInfoWKey: FreeAddrInfoW,

		FreeEnvironmentStringsKey: FreeEnvironmentStrings,

		FreeLibraryKey: FreeLibrary,

		FreeSidKey: FreeSid,

		FsyncKey: Fsync,

		FtruncateKey: Ftruncate,

		FullPathKey: FullPath,

		GetACPKey: GetACP,

		GetAcceptExSockaddrsKey: GetAcceptExSockaddrs,

		GetAdaptersAddressesKey: GetAdaptersAddresses,

		GetAdaptersInfoKey: GetAdaptersInfo,

		GetAddrInfoWKey: GetAddrInfoW,

		GetCommandLineKey: GetCommandLine,

		GetComputerNameKey: GetComputerName,

		GetComputerNameExKey: GetComputerNameEx,

		GetConsoleModeKey: GetConsoleMode,

		GetConsoleScreenBufferInfoKey: GetConsoleScreenBufferInfo,

		GetCurrentDirectoryKey: GetCurrentDirectory,

		GetCurrentProcessKey: GetCurrentProcess,

		GetCurrentThreadIdKey: GetCurrentThreadId,

		GetDriveTypeKey: GetDriveType,

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

		GetLogicalDriveStringsKey: GetLogicalDriveStrings,

		GetLogicalDrivesKey: GetLogicalDrives,

		GetLongPathNameKey: GetLongPathName,

		GetProcAddressKey: GetProcAddress,

		GetProcAddressByOrdinalKey: GetProcAddressByOrdinal,

		GetProcessTimesKey: GetProcessTimes,

		GetProtoByNameKey: GetProtoByName,

		GetQueuedCompletionStatusKey: GetQueuedCompletionStatus,

		GetServByNameKey: GetServByName,

		GetShortPathNameKey: GetShortPathName,

		GetStartupInfoKey: GetStartupInfo,

		GetStdHandleKey: GetStdHandle,

		GetSystemTimeAsFileTimeKey: GetSystemTimeAsFileTime,

		GetSystemTimePreciseAsFileTimeKey: GetSystemTimePreciseAsFileTime,

		GetTempPathKey: GetTempPath,

		GetTimeZoneInformationKey: GetTimeZoneInformation,

		GetTokenInformationKey: GetTokenInformation,

		GetUserNameExKey: GetUserNameEx,

		GetUserProfileDirectoryKey: GetUserProfileDirectory,

		GetVersionKey: GetVersion,

		GetVolumeInformationKey: GetVolumeInformation,

		GetVolumeInformationByHandleKey: GetVolumeInformationByHandle,

		GetVolumeNameForVolumeMountPointKey: GetVolumeNameForVolumeMountPoint,

		GetVolumePathNameKey: GetVolumePathName,

		GetVolumePathNamesForVolumeNameKey: GetVolumePathNamesForVolumeName,

		GetegidKey: Getegid,

		GetenvKey: Getenv,

		GeteuidKey: Geteuid,

		GetgidKey: Getgid,

		GetgroupsKey: Getgroups,

		GetpagesizeKey: Getpagesize,

		GetpeernameKey: Getpeername,

		GetpidKey: Getpid,

		GetppidKey: Getppid,

		GetsocknameKey: Getsockname,

		GetsockoptKey: Getsockopt,

		GetsockoptIntKey: GetsockoptInt,

		GettimeofdayKey: Gettimeofday,

		GetuidKey: Getuid,

		GetwdKey: Getwd,

		LchownKey: Lchown,

		LinkKey: Link,

		ListenKey: Listen,

		LoadCancelIoExKey: LoadCancelIoEx,

		LoadConnectExKey: LoadConnectEx,

		LoadCreateSymbolicLinkKey: LoadCreateSymbolicLink,

		LoadDLLKey: LoadDLL,

		LoadGetAddrInfoKey: LoadGetAddrInfo,

		LoadGetSystemTimePreciseAsFileTimeKey: LoadGetSystemTimePreciseAsFileTime,

		LoadLibraryKey: LoadLibrary,

		LoadLibraryExKey: LoadLibraryEx,

		LoadSetFileCompletionNotificationModesKey: LoadSetFileCompletionNotificationModes,

		LocalFreeKey: LocalFree,

		LookupAccountNameKey: LookupAccountName,

		LookupAccountSidKey: LookupAccountSid,

		LookupSIDKey: LookupSID,

		MapViewOfFileKey: MapViewOfFile,

		MkdirKey: Mkdir,

		MoveFileKey: MoveFile,

		MoveFileExKey: MoveFileEx,

		MultiByteToWideCharKey: MultiByteToWideChar,

		MustLoadDLLKey: MustLoadDLL,

		NetApiBufferFreeKey: NetApiBufferFree,

		NetGetJoinInformationKey: NetGetJoinInformation,

		NetUserGetInfoKey: NetUserGetInfo,

		NewCallbackKey: NewCallback,

		NewCallbackCDeclKey: NewCallbackCDecl,

		NewLazyDLLKey: NewLazyDLL,

		NewLazySystemDLLKey: NewLazySystemDLL,

		NsecToFiletimeKey: NsecToFiletime,

		NsecToTimespecKey: NsecToTimespec,

		NsecToTimevalKey: NsecToTimeval,

		NtohsKey: Ntohs,

		OpenKey: Open,

		OpenCurrentProcessTokenKey: OpenCurrentProcessToken,

		OpenEventKey: OpenEvent,

		OpenProcessKey: OpenProcess,

		OpenProcessTokenKey: OpenProcessToken,

		OpenSCManagerKey: OpenSCManager,

		OpenServiceKey: OpenService,

		PipeKey: Pipe,

		PostQueuedCompletionStatusKey: PostQueuedCompletionStatus,

		Process32FirstKey: Process32First,

		Process32NextKey: Process32Next,

		PulseEventKey: PulseEvent,

		QueryDosDeviceKey: QueryDosDevice,

		QueryServiceConfigKey: QueryServiceConfig,

		QueryServiceConfig2Key: QueryServiceConfig2,

		QueryServiceStatusKey: QueryServiceStatus,

		QueryServiceStatusExKey: QueryServiceStatusEx,

		ReadKey: Read,

		ReadConsoleKey: ReadConsole,

		ReadDirectoryChangesKey: ReadDirectoryChanges,

		ReadFileKey: ReadFile,

		ReadlinkKey: Readlink,

		RecvfromKey: Recvfrom,

		RegCloseKeyKey: RegCloseKey,

		RegEnumKeyExKey: RegEnumKeyEx,

		RegOpenKeyExKey: RegOpenKeyEx,

		RegQueryInfoKeyKey: RegQueryInfoKey,

		RegQueryValueExKey: RegQueryValueEx,

		RegisterEventSourceKey: RegisterEventSource,

		RemoveDirectoryKey: RemoveDirectory,

		RenameKey: Rename,

		ReportEventKey: ReportEvent,

		ResetEventKey: ResetEvent,

		RmdirKey: Rmdir,

		SeekKey: Seek,

		SendtoKey: Sendto,

		SetConsoleModeKey: SetConsoleMode,

		SetCurrentDirectoryKey: SetCurrentDirectory,

		SetEndOfFileKey: SetEndOfFile,

		SetEnvironmentVariableKey: SetEnvironmentVariable,

		SetEventKey: SetEvent,

		SetFileAttributesKey: SetFileAttributes,

		SetFileCompletionNotificationModesKey: SetFileCompletionNotificationModes,

		SetFilePointerKey: SetFilePointer,

		SetFileTimeKey: SetFileTime,

		SetHandleInformationKey: SetHandleInformation,

		SetServiceStatusKey: SetServiceStatus,

		SetStdHandleKey: SetStdHandle,

		SetVolumeLabelKey: SetVolumeLabel,

		SetVolumeMountPointKey: SetVolumeMountPoint,

		SetenvKey: Setenv,

		SetsockoptKey: Setsockopt,

		SetsockoptIPMreqKey: SetsockoptIPMreq,

		SetsockoptIPv6MreqKey: SetsockoptIPv6Mreq,

		SetsockoptInet4AddrKey: SetsockoptInet4Addr,

		SetsockoptIntKey: SetsockoptInt,

		SetsockoptLingerKey: SetsockoptLinger,

		SetsockoptTimevalKey: SetsockoptTimeval,

		ShutdownKey: Shutdown,

		SocketKey: Socket,

		StartServiceKey: StartService,

		StartServiceCtrlDispatcherKey: StartServiceCtrlDispatcher,

		StringToSidKey: StringToSid,

		StringToUTF16Key: StringToUTF16,

		StringToUTF16PtrKey: StringToUTF16Ptr,

		SymlinkKey: Symlink,

		TerminateProcessKey: TerminateProcess,

		TimespecToNsecKey: TimespecToNsec,

		TranslateAccountNameKey: TranslateAccountName,

		TranslateNameKey: TranslateName,

		TransmitFileKey: TransmitFile,

		UTF16FromStringKey: UTF16FromString,

		UTF16PtrFromStringKey: UTF16PtrFromString,

		UTF16ToStringKey: UTF16ToString,

		UnlinkKey: Unlink,

		UnmapViewOfFileKey: UnmapViewOfFile,

		UnsetenvKey: Unsetenv,

		UtimesKey: Utimes,

		UtimesNanoKey: UtimesNano,

		VirtualAllocKey: VirtualAlloc,

		VirtualFreeKey: VirtualFree,

		VirtualLockKey: VirtualLock,

		VirtualProtectKey: VirtualProtect,

		VirtualUnlockKey: VirtualUnlock,

		WSACleanupKey: WSACleanup,

		WSAEnumProtocolsKey: WSAEnumProtocols,

		WSAIoctlKey: WSAIoctl,

		WSARecvKey: WSARecv,

		WSARecvFromKey: WSARecvFrom,

		WSARecvMsgKey: WSARecvMsg,

		WSASendKey: WSASend,

		WSASendMsgKey: WSASendMsg,

		WSASendToKey: WSASendTo,

		WSASendtoKey: WSASendto,

		WSAStartupKey: WSAStartup,

		WaitForSingleObjectKey: WaitForSingleObject,

		WriteKey: Write,

		WriteConsoleKey: WriteConsole,

		WriteFileKey: WriteFile,
	}
}

var Accept = windows.Accept

var AcceptEx = windows.AcceptEx

var AllocateAndInitializeSid = windows.AllocateAndInitializeSid

var Bind = windows.Bind

// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
var BytePtrFromString = windows.BytePtrFromString

// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
var ByteSliceFromString = windows.ByteSliceFromString

var CancelIo = windows.CancelIo

var CancelIoEx = windows.CancelIoEx

var CertAddCertificateContextToStore = windows.CertAddCertificateContextToStore

var CertCloseStore = windows.CertCloseStore

var CertCreateCertificateContext = windows.CertCreateCertificateContext

var CertEnumCertificatesInStore = windows.CertEnumCertificatesInStore

var CertFreeCertificateChain = windows.CertFreeCertificateChain

var CertFreeCertificateContext = windows.CertFreeCertificateContext

var CertGetCertificateChain = windows.CertGetCertificateChain

var CertOpenStore = windows.CertOpenStore

var CertOpenSystemStore = windows.CertOpenSystemStore

var CertVerifyCertificateChainPolicy = windows.CertVerifyCertificateChainPolicy

var ChangeServiceConfig = windows.ChangeServiceConfig

var ChangeServiceConfig2 = windows.ChangeServiceConfig2

var Chdir = windows.Chdir

var Chmod = windows.Chmod

var Chown = windows.Chown

var Clearenv = windows.Clearenv

var Close = windows.Close

var CloseHandle = windows.CloseHandle

var CloseOnExec = windows.CloseOnExec

var CloseServiceHandle = windows.CloseServiceHandle

var Closesocket = windows.Closesocket

var CommandLineToArgv = windows.CommandLineToArgv

var ComputerName = windows.ComputerName

var Connect = windows.Connect

var ConnectEx = windows.ConnectEx

var ControlService = windows.ControlService

var ConvertSidToStringSid = windows.ConvertSidToStringSid

var ConvertStringSidToSid = windows.ConvertStringSidToSid

var CopySid = windows.CopySid

var CreateDirectory = windows.CreateDirectory

var CreateEvent = windows.CreateEvent

var CreateEventEx = windows.CreateEventEx

var CreateFile = windows.CreateFile

var CreateFileMapping = windows.CreateFileMapping

var CreateHardLink = windows.CreateHardLink

var CreateIoCompletionPort = windows.CreateIoCompletionPort

var CreatePipe = windows.CreatePipe

var CreateProcess = windows.CreateProcess

var CreateService = windows.CreateService

var CreateSymbolicLink = windows.CreateSymbolicLink

var CreateToolhelp32Snapshot = windows.CreateToolhelp32Snapshot

var CryptAcquireContext = windows.CryptAcquireContext

var CryptGenRandom = windows.CryptGenRandom

var CryptReleaseContext = windows.CryptReleaseContext

var DefineDosDevice = windows.DefineDosDevice

var DeleteFile = windows.DeleteFile

var DeleteService = windows.DeleteService

var DeleteVolumeMountPoint = windows.DeleteVolumeMountPoint

var DeregisterEventSource = windows.DeregisterEventSource

var DeviceIoControl = windows.DeviceIoControl

var DnsNameCompare = windows.DnsNameCompare

var DnsQuery = windows.DnsQuery

var DnsRecordListFree = windows.DnsRecordListFree

var DuplicateHandle = windows.DuplicateHandle

var EnumServicesStatusEx = windows.EnumServicesStatusEx

var Environ = windows.Environ

var EqualSid = windows.EqualSid

// EscapeArg rewrites command line argument s as prescribed
// in http://msdn.microsoft.com/en-us/library/ms880421.
// This function returns &#34;&#34; (2 double quotes) if s is empty.
// Alternatively, these transformations are done:
// - every back slash (\) is doubled, but only if immediately
//   followed by double quote (&#34;);
// - every double quote (&#34;) is escaped by back slash (\);
// - finally, s is wrapped with double quotes (arg -&gt; &#34;arg&#34;),
//   but only if there is space or tab inside s.
var EscapeArg = windows.EscapeArg

var Exit = windows.Exit

var ExitProcess = windows.ExitProcess

// TODO(brainman): fix all needed for os
var Fchdir = windows.Fchdir

var Fchmod = windows.Fchmod

var Fchown = windows.Fchown

var FindClose = windows.FindClose

var FindFirstFile = windows.FindFirstFile

var FindFirstVolume = windows.FindFirstVolume

var FindFirstVolumeMountPoint = windows.FindFirstVolumeMountPoint

var FindNextFile = windows.FindNextFile

var FindNextVolume = windows.FindNextVolume

var FindNextVolumeMountPoint = windows.FindNextVolumeMountPoint

var FindVolumeClose = windows.FindVolumeClose

var FindVolumeMountPointClose = windows.FindVolumeMountPointClose

var FlushFileBuffers = windows.FlushFileBuffers

var FlushViewOfFile = windows.FlushViewOfFile

var FormatMessage = windows.FormatMessage

var FreeAddrInfoW = windows.FreeAddrInfoW

var FreeEnvironmentStrings = windows.FreeEnvironmentStrings

var FreeLibrary = windows.FreeLibrary

var FreeSid = windows.FreeSid

var Fsync = windows.Fsync

var Ftruncate = windows.Ftruncate

// FullPath retrieves the full path of the specified file.
var FullPath = windows.FullPath

var GetACP = windows.GetACP

var GetAcceptExSockaddrs = windows.GetAcceptExSockaddrs

var GetAdaptersAddresses = windows.GetAdaptersAddresses

var GetAdaptersInfo = windows.GetAdaptersInfo

var GetAddrInfoW = windows.GetAddrInfoW

var GetCommandLine = windows.GetCommandLine

var GetComputerName = windows.GetComputerName

var GetComputerNameEx = windows.GetComputerNameEx

var GetConsoleMode = windows.GetConsoleMode

var GetConsoleScreenBufferInfo = windows.GetConsoleScreenBufferInfo

var GetCurrentDirectory = windows.GetCurrentDirectory

var GetCurrentProcess = windows.GetCurrentProcess

var GetCurrentThreadId = windows.GetCurrentThreadId

var GetDriveType = windows.GetDriveType

var GetEnvironmentStrings = windows.GetEnvironmentStrings

var GetEnvironmentVariable = windows.GetEnvironmentVariable

var GetExitCodeProcess = windows.GetExitCodeProcess

var GetFileAttributes = windows.GetFileAttributes

var GetFileAttributesEx = windows.GetFileAttributesEx

var GetFileInformationByHandle = windows.GetFileInformationByHandle

var GetFileType = windows.GetFileType

var GetFullPathName = windows.GetFullPathName

var GetHostByName = windows.GetHostByName

var GetIfEntry = windows.GetIfEntry

var GetLastError = windows.GetLastError

var GetLengthSid = windows.GetLengthSid

var GetLogicalDriveStrings = windows.GetLogicalDriveStrings

var GetLogicalDrives = windows.GetLogicalDrives

var GetLongPathName = windows.GetLongPathName

var GetProcAddress = windows.GetProcAddress

// GetProcAddressByOrdinal retrieves the address of the exported
// function from module by ordinal.
var GetProcAddressByOrdinal = windows.GetProcAddressByOrdinal

var GetProcessTimes = windows.GetProcessTimes

var GetProtoByName = windows.GetProtoByName

var GetQueuedCompletionStatus = windows.GetQueuedCompletionStatus

var GetServByName = windows.GetServByName

var GetShortPathName = windows.GetShortPathName

var GetStartupInfo = windows.GetStartupInfo

var GetStdHandle = windows.GetStdHandle

var GetSystemTimeAsFileTime = windows.GetSystemTimeAsFileTime

var GetSystemTimePreciseAsFileTime = windows.GetSystemTimePreciseAsFileTime

var GetTempPath = windows.GetTempPath

var GetTimeZoneInformation = windows.GetTimeZoneInformation

var GetTokenInformation = windows.GetTokenInformation

var GetUserNameEx = windows.GetUserNameEx

var GetUserProfileDirectory = windows.GetUserProfileDirectory

var GetVersion = windows.GetVersion

var GetVolumeInformation = windows.GetVolumeInformation

var GetVolumeInformationByHandle = windows.GetVolumeInformationByHandle

var GetVolumeNameForVolumeMountPoint = windows.GetVolumeNameForVolumeMountPoint

var GetVolumePathName = windows.GetVolumePathName

var GetVolumePathNamesForVolumeName = windows.GetVolumePathNamesForVolumeName

var Getegid = windows.Getegid

var Getenv = windows.Getenv

var Geteuid = windows.Geteuid

var Getgid = windows.Getgid

var Getgroups = windows.Getgroups

var Getpagesize = windows.Getpagesize

var Getpeername = windows.Getpeername

var Getpid = windows.Getpid

var Getppid = windows.Getppid

var Getsockname = windows.Getsockname

var Getsockopt = windows.Getsockopt

var GetsockoptInt = windows.GetsockoptInt

var Gettimeofday = windows.Gettimeofday

var Getuid = windows.Getuid

var Getwd = windows.Getwd

var Lchown = windows.Lchown

var Link = windows.Link

var Listen = windows.Listen

var LoadCancelIoEx = windows.LoadCancelIoEx

var LoadConnectEx = windows.LoadConnectEx

var LoadCreateSymbolicLink = windows.LoadCreateSymbolicLink

// LoadDLL loads DLL file into memory.
//
// Warning: using LoadDLL without an absolute path name is subject to
// DLL preloading attacks. To safely load a system DLL, use LazyDLL
// with System set to true, or use LoadLibraryEx directly.
var LoadDLL = windows.LoadDLL

var LoadGetAddrInfo = windows.LoadGetAddrInfo

var LoadGetSystemTimePreciseAsFileTime = windows.LoadGetSystemTimePreciseAsFileTime

var LoadLibrary = windows.LoadLibrary

var LoadLibraryEx = windows.LoadLibraryEx

var LoadSetFileCompletionNotificationModes = windows.LoadSetFileCompletionNotificationModes

var LocalFree = windows.LocalFree

var LookupAccountName = windows.LookupAccountName

var LookupAccountSid = windows.LookupAccountSid

// LookupSID retrieves a security identifier sid for the account
// and the name of the domain on which the account was found.
// System specify target computer to search.
var LookupSID = windows.LookupSID

var MapViewOfFile = windows.MapViewOfFile

var Mkdir = windows.Mkdir

var MoveFile = windows.MoveFile

var MoveFileEx = windows.MoveFileEx

var MultiByteToWideChar = windows.MultiByteToWideChar

// MustLoadDLL is like LoadDLL but panics if load operation failes.
var MustLoadDLL = windows.MustLoadDLL

var NetApiBufferFree = windows.NetApiBufferFree

var NetGetJoinInformation = windows.NetGetJoinInformation

var NetUserGetInfo = windows.NetUserGetInfo

// NewCallback converts a Go function to a function pointer conforming to the stdcall calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
var NewCallback = windows.NewCallback

// NewCallbackCDecl converts a Go function to a function pointer conforming to the cdecl calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
var NewCallbackCDecl = windows.NewCallbackCDecl

// NewLazyDLL creates new LazyDLL associated with DLL file.
var NewLazyDLL = windows.NewLazyDLL

// NewLazySystemDLL is like NewLazyDLL, but will only
// search Windows System directory for the DLL if name is
// a base name (like &#34;advapi32.dll&#34;).
var NewLazySystemDLL = windows.NewLazySystemDLL

var NsecToFiletime = windows.NsecToFiletime

var NsecToTimespec = windows.NsecToTimespec

var NsecToTimeval = windows.NsecToTimeval

var Ntohs = windows.Ntohs

var Open = windows.Open

// OpenCurrentProcessToken opens the access token
// associated with current process.
var OpenCurrentProcessToken = windows.OpenCurrentProcessToken

var OpenEvent = windows.OpenEvent

var OpenProcess = windows.OpenProcess

var OpenProcessToken = windows.OpenProcessToken

var OpenSCManager = windows.OpenSCManager

var OpenService = windows.OpenService

var Pipe = windows.Pipe

var PostQueuedCompletionStatus = windows.PostQueuedCompletionStatus

var Process32First = windows.Process32First

var Process32Next = windows.Process32Next

var PulseEvent = windows.PulseEvent

var QueryDosDevice = windows.QueryDosDevice

var QueryServiceConfig = windows.QueryServiceConfig

var QueryServiceConfig2 = windows.QueryServiceConfig2

var QueryServiceStatus = windows.QueryServiceStatus

var QueryServiceStatusEx = windows.QueryServiceStatusEx

var Read = windows.Read

var ReadConsole = windows.ReadConsole

var ReadDirectoryChanges = windows.ReadDirectoryChanges

var ReadFile = windows.ReadFile

// Readlink returns the destination of the named symbolic link.
var Readlink = windows.Readlink

var Recvfrom = windows.Recvfrom

var RegCloseKey = windows.RegCloseKey

var RegEnumKeyEx = windows.RegEnumKeyEx

var RegOpenKeyEx = windows.RegOpenKeyEx

var RegQueryInfoKey = windows.RegQueryInfoKey

var RegQueryValueEx = windows.RegQueryValueEx

var RegisterEventSource = windows.RegisterEventSource

var RemoveDirectory = windows.RemoveDirectory

var Rename = windows.Rename

var ReportEvent = windows.ReportEvent

var ResetEvent = windows.ResetEvent

var Rmdir = windows.Rmdir

var Seek = windows.Seek

var Sendto = windows.Sendto

var SetConsoleMode = windows.SetConsoleMode

var SetCurrentDirectory = windows.SetCurrentDirectory

var SetEndOfFile = windows.SetEndOfFile

var SetEnvironmentVariable = windows.SetEnvironmentVariable

var SetEvent = windows.SetEvent

var SetFileAttributes = windows.SetFileAttributes

var SetFileCompletionNotificationModes = windows.SetFileCompletionNotificationModes

var SetFilePointer = windows.SetFilePointer

var SetFileTime = windows.SetFileTime

var SetHandleInformation = windows.SetHandleInformation

var SetServiceStatus = windows.SetServiceStatus

var SetStdHandle = windows.SetStdHandle

var SetVolumeLabel = windows.SetVolumeLabel

var SetVolumeMountPoint = windows.SetVolumeMountPoint

var Setenv = windows.Setenv

var Setsockopt = windows.Setsockopt

var SetsockoptIPMreq = windows.SetsockoptIPMreq

var SetsockoptIPv6Mreq = windows.SetsockoptIPv6Mreq

var SetsockoptInet4Addr = windows.SetsockoptInet4Addr

var SetsockoptInt = windows.SetsockoptInt

var SetsockoptLinger = windows.SetsockoptLinger

var SetsockoptTimeval = windows.SetsockoptTimeval

var Shutdown = windows.Shutdown

var Socket = windows.Socket

var StartService = windows.StartService

var StartServiceCtrlDispatcher = windows.StartServiceCtrlDispatcher

// StringToSid converts a string-format security identifier
// sid into a valid, functional sid.
var StringToSid = windows.StringToSid

// StringToUTF16 is deprecated. Use UTF16FromString instead.
// If s contains a NUL byte this function panics instead of
// returning an error.
var StringToUTF16 = windows.StringToUTF16

// StringToUTF16Ptr is deprecated. Use UTF16PtrFromString instead.
// If s contains a NUL byte this function panics instead of
// returning an error.
var StringToUTF16Ptr = windows.StringToUTF16Ptr

var Symlink = windows.Symlink

var TerminateProcess = windows.TerminateProcess

var TimespecToNsec = windows.TimespecToNsec

// TranslateAccountName converts a directory service
// object name from one format to another.
var TranslateAccountName = windows.TranslateAccountName

var TranslateName = windows.TranslateName

var TransmitFile = windows.TransmitFile

// UTF16FromString returns the UTF-16 encoding of the UTF-8 string
// s, with a terminating NUL added. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
var UTF16FromString = windows.UTF16FromString

// UTF16PtrFromString returns pointer to the UTF-16 encoding of
// the UTF-8 string s, with a terminating NUL added. If s
// contains a NUL byte at any location, it returns (nil, syscall.EINVAL).
var UTF16PtrFromString = windows.UTF16PtrFromString

// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL removed.
var UTF16ToString = windows.UTF16ToString

var Unlink = windows.Unlink

var UnmapViewOfFile = windows.UnmapViewOfFile

var Unsetenv = windows.Unsetenv

var Utimes = windows.Utimes

var UtimesNano = windows.UtimesNano

var VirtualAlloc = windows.VirtualAlloc

var VirtualFree = windows.VirtualFree

var VirtualLock = windows.VirtualLock

var VirtualProtect = windows.VirtualProtect

var VirtualUnlock = windows.VirtualUnlock

var WSACleanup = windows.WSACleanup

var WSAEnumProtocols = windows.WSAEnumProtocols

var WSAIoctl = windows.WSAIoctl

var WSARecv = windows.WSARecv

var WSARecvFrom = windows.WSARecvFrom

var WSARecvMsg = windows.WSARecvMsg

var WSASend = windows.WSASend

var WSASendMsg = windows.WSASendMsg

var WSASendTo = windows.WSASendTo

var WSASendto = windows.WSASendto

var WSAStartup = windows.WSAStartup

var WaitForSingleObject = windows.WaitForSingleObject

var Write = windows.Write

var WriteConsole = windows.WriteConsole

var WriteFile = windows.WriteFile
