package windowsh

import (
	"internal/syscall/windows"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AdjustTokenPrivilegesKey = "AdjustTokenPrivileges"

	DuplicateTokenExKey = "DuplicateTokenEx"

	GetACPKey = "GetACP"

	GetAdaptersAddressesKey = "GetAdaptersAddresses"

	GetComputerNameExKey = "GetComputerNameEx"

	GetConsoleCPKey = "GetConsoleCP"

	GetCurrentThreadKey = "GetCurrentThread"

	GetFileInformationByHandleExKey = "GetFileInformationByHandleEx"

	GetFinalPathNameByHandleKey = "GetFinalPathNameByHandle"

	GetModuleFileNameKey = "GetModuleFileName"

	GetProcessMemoryInfoKey = "GetProcessMemoryInfo"

	GetProfilesDirectoryKey = "GetProfilesDirectory"

	ImpersonateSelfKey = "ImpersonateSelf"

	LoadGetFinalPathNameByHandleKey = "LoadGetFinalPathNameByHandle"

	LockFileExKey = "LockFileEx"

	LookupPrivilegeValueKey = "LookupPrivilegeValue"

	MoveFileExKey = "MoveFileEx"

	MultiByteToWideCharKey = "MultiByteToWideChar"

	NetShareAddKey = "NetShareAdd"

	NetShareDelKey = "NetShareDel"

	NetUserGetLocalGroupsKey = "NetUserGetLocalGroups"

	OpenThreadTokenKey = "OpenThreadToken"

	RenameKey = "Rename"

	RevertToSelfKey = "RevertToSelf"

	SetTokenInformationKey = "SetTokenInformation"

	UnlockFileExKey = "UnlockFileEx"

	WSARecvMsgKey = "WSARecvMsg"

	WSASendMsgKey = "WSASendMsg"

	WSASocketKey = "WSASocket"
)

func New() hctx.Map {
	return hctx.Map{

		AdjustTokenPrivilegesKey: AdjustTokenPrivileges,

		DuplicateTokenExKey: DuplicateTokenEx,

		GetACPKey: GetACP,

		GetAdaptersAddressesKey: GetAdaptersAddresses,

		GetComputerNameExKey: GetComputerNameEx,

		GetConsoleCPKey: GetConsoleCP,

		GetCurrentThreadKey: GetCurrentThread,

		GetFileInformationByHandleExKey: GetFileInformationByHandleEx,

		GetFinalPathNameByHandleKey: GetFinalPathNameByHandle,

		GetModuleFileNameKey: GetModuleFileName,

		GetProcessMemoryInfoKey: GetProcessMemoryInfo,

		GetProfilesDirectoryKey: GetProfilesDirectory,

		ImpersonateSelfKey: ImpersonateSelf,

		LoadGetFinalPathNameByHandleKey: LoadGetFinalPathNameByHandle,

		LockFileExKey: LockFileEx,

		LookupPrivilegeValueKey: LookupPrivilegeValue,

		MoveFileExKey: MoveFileEx,

		MultiByteToWideCharKey: MultiByteToWideChar,

		NetShareAddKey: NetShareAdd,

		NetShareDelKey: NetShareDel,

		NetUserGetLocalGroupsKey: NetUserGetLocalGroups,

		OpenThreadTokenKey: OpenThreadToken,

		RenameKey: Rename,

		RevertToSelfKey: RevertToSelf,

		SetTokenInformationKey: SetTokenInformation,

		UnlockFileExKey: UnlockFileEx,

		WSARecvMsgKey: WSARecvMsg,

		WSASendMsgKey: WSASendMsg,

		WSASocketKey: WSASocket,
	}
}

var AdjustTokenPrivileges = windows.AdjustTokenPrivileges

var DuplicateTokenEx = windows.DuplicateTokenEx

var GetACP = windows.GetACP

var GetAdaptersAddresses = windows.GetAdaptersAddresses

var GetComputerNameEx = windows.GetComputerNameEx

var GetConsoleCP = windows.GetConsoleCP

var GetCurrentThread = windows.GetCurrentThread

var GetFileInformationByHandleEx = windows.GetFileInformationByHandleEx

var GetFinalPathNameByHandle = windows.GetFinalPathNameByHandle

var GetModuleFileName = windows.GetModuleFileName

var GetProcessMemoryInfo = windows.GetProcessMemoryInfo

var GetProfilesDirectory = windows.GetProfilesDirectory

var ImpersonateSelf = windows.ImpersonateSelf

var LoadGetFinalPathNameByHandle = windows.LoadGetFinalPathNameByHandle

var LockFileEx = windows.LockFileEx

var LookupPrivilegeValue = windows.LookupPrivilegeValue

var MoveFileEx = windows.MoveFileEx

var MultiByteToWideChar = windows.MultiByteToWideChar

var NetShareAdd = windows.NetShareAdd

var NetShareDel = windows.NetShareDel

var NetUserGetLocalGroups = windows.NetUserGetLocalGroups

var OpenThreadToken = windows.OpenThreadToken

var Rename = windows.Rename

var RevertToSelf = windows.RevertToSelf

var SetTokenInformation = windows.SetTokenInformation

var UnlockFileEx = windows.UnlockFileEx

var WSARecvMsg = windows.WSARecvMsg

var WSASendMsg = windows.WSASendMsg

var WSASocket = windows.WSASocket
