package unixh

import (
	"cmd/vendor/golang.org/x/sys/unix"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	AcceptKey = "Accept"

	AcceptKey = "Accept"

	AcceptKey = "Accept"

	AcceptKey = "Accept"

	Accept4Key = "Accept4"

	Accept4Key = "Accept4"

	Accept4Key = "Accept4"

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

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AcctKey = "Acct"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

	AddKeyKey = "AddKey"

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

	AdjtimexKey = "Adjtimex"

	AdjtimexKey = "Adjtimex"

	BindKey = "Bind"

	BindToDeviceKey = "BindToDevice"

	BytePtrFromStringKey = "BytePtrFromString"

	ByteSliceFromStringKey = "ByteSliceFromString"

	CapEnterKey = "CapEnter"

	CapEnterKey = "CapEnter"

	CapEnterKey = "CapEnter"

	CapEnterKey = "CapEnter"

	CapRightsClearKey = "CapRightsClear"

	CapRightsGetKey = "CapRightsGet"

	CapRightsInitKey = "CapRightsInit"

	CapRightsIsSetKey = "CapRightsIsSet"

	CapRightsLimitKey = "CapRightsLimit"

	CapRightsSetKey = "CapRightsSet"

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

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

	ChdirKey = "Chdir"

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

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGetresKey = "ClockGetres"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

	ClockGettimeKey = "ClockGettime"

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

	CloseKey = "Close"

	CloseKey = "Close"

	CloseOnExecKey = "CloseOnExec"

	CmsgLenKey = "CmsgLen"

	CmsgSpaceKey = "CmsgSpace"

	ConnectKey = "Connect"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CopyFileRangeKey = "CopyFileRange"

	CreatKey = "Creat"

	CreatKey = "Creat"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

	DeleteModuleKey = "DeleteModule"

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

	Dup3Key = "Dup3"

	Dup3Key = "Dup3"

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

	EpollWaitKey = "EpollWait"

	EpollWaitKey = "EpollWait"

	ErrnoNameKey = "ErrnoName"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	EventfdKey = "Eventfd"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExchangedataKey = "Exchangedata"

	ExecKey = "Exec"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExitKey = "Exit"

	ExtattrDeleteFdKey = "ExtattrDeleteFd"

	ExtattrDeleteFdKey = "ExtattrDeleteFd"

	ExtattrDeleteFdKey = "ExtattrDeleteFd"

	ExtattrDeleteFdKey = "ExtattrDeleteFd"

	ExtattrDeleteFdKey = "ExtattrDeleteFd"

	ExtattrDeleteFdKey = "ExtattrDeleteFd"

	ExtattrDeleteFdKey = "ExtattrDeleteFd"

	ExtattrDeleteFileKey = "ExtattrDeleteFile"

	ExtattrDeleteFileKey = "ExtattrDeleteFile"

	ExtattrDeleteFileKey = "ExtattrDeleteFile"

	ExtattrDeleteFileKey = "ExtattrDeleteFile"

	ExtattrDeleteFileKey = "ExtattrDeleteFile"

	ExtattrDeleteFileKey = "ExtattrDeleteFile"

	ExtattrDeleteFileKey = "ExtattrDeleteFile"

	ExtattrDeleteLinkKey = "ExtattrDeleteLink"

	ExtattrDeleteLinkKey = "ExtattrDeleteLink"

	ExtattrDeleteLinkKey = "ExtattrDeleteLink"

	ExtattrDeleteLinkKey = "ExtattrDeleteLink"

	ExtattrDeleteLinkKey = "ExtattrDeleteLink"

	ExtattrDeleteLinkKey = "ExtattrDeleteLink"

	ExtattrDeleteLinkKey = "ExtattrDeleteLink"

	ExtattrGetFdKey = "ExtattrGetFd"

	ExtattrGetFdKey = "ExtattrGetFd"

	ExtattrGetFdKey = "ExtattrGetFd"

	ExtattrGetFdKey = "ExtattrGetFd"

	ExtattrGetFdKey = "ExtattrGetFd"

	ExtattrGetFdKey = "ExtattrGetFd"

	ExtattrGetFdKey = "ExtattrGetFd"

	ExtattrGetFileKey = "ExtattrGetFile"

	ExtattrGetFileKey = "ExtattrGetFile"

	ExtattrGetFileKey = "ExtattrGetFile"

	ExtattrGetFileKey = "ExtattrGetFile"

	ExtattrGetFileKey = "ExtattrGetFile"

	ExtattrGetFileKey = "ExtattrGetFile"

	ExtattrGetFileKey = "ExtattrGetFile"

	ExtattrGetLinkKey = "ExtattrGetLink"

	ExtattrGetLinkKey = "ExtattrGetLink"

	ExtattrGetLinkKey = "ExtattrGetLink"

	ExtattrGetLinkKey = "ExtattrGetLink"

	ExtattrGetLinkKey = "ExtattrGetLink"

	ExtattrGetLinkKey = "ExtattrGetLink"

	ExtattrGetLinkKey = "ExtattrGetLink"

	ExtattrListFdKey = "ExtattrListFd"

	ExtattrListFdKey = "ExtattrListFd"

	ExtattrListFdKey = "ExtattrListFd"

	ExtattrListFdKey = "ExtattrListFd"

	ExtattrListFdKey = "ExtattrListFd"

	ExtattrListFdKey = "ExtattrListFd"

	ExtattrListFdKey = "ExtattrListFd"

	ExtattrListFileKey = "ExtattrListFile"

	ExtattrListFileKey = "ExtattrListFile"

	ExtattrListFileKey = "ExtattrListFile"

	ExtattrListFileKey = "ExtattrListFile"

	ExtattrListFileKey = "ExtattrListFile"

	ExtattrListFileKey = "ExtattrListFile"

	ExtattrListFileKey = "ExtattrListFile"

	ExtattrListLinkKey = "ExtattrListLink"

	ExtattrListLinkKey = "ExtattrListLink"

	ExtattrListLinkKey = "ExtattrListLink"

	ExtattrListLinkKey = "ExtattrListLink"

	ExtattrListLinkKey = "ExtattrListLink"

	ExtattrListLinkKey = "ExtattrListLink"

	ExtattrListLinkKey = "ExtattrListLink"

	ExtattrSetFdKey = "ExtattrSetFd"

	ExtattrSetFdKey = "ExtattrSetFd"

	ExtattrSetFdKey = "ExtattrSetFd"

	ExtattrSetFdKey = "ExtattrSetFd"

	ExtattrSetFdKey = "ExtattrSetFd"

	ExtattrSetFdKey = "ExtattrSetFd"

	ExtattrSetFdKey = "ExtattrSetFd"

	ExtattrSetFileKey = "ExtattrSetFile"

	ExtattrSetFileKey = "ExtattrSetFile"

	ExtattrSetFileKey = "ExtattrSetFile"

	ExtattrSetFileKey = "ExtattrSetFile"

	ExtattrSetFileKey = "ExtattrSetFile"

	ExtattrSetFileKey = "ExtattrSetFile"

	ExtattrSetFileKey = "ExtattrSetFile"

	ExtattrSetLinkKey = "ExtattrSetLink"

	ExtattrSetLinkKey = "ExtattrSetLink"

	ExtattrSetLinkKey = "ExtattrSetLink"

	ExtattrSetLinkKey = "ExtattrSetLink"

	ExtattrSetLinkKey = "ExtattrSetLink"

	ExtattrSetLinkKey = "ExtattrSetLink"

	ExtattrSetLinkKey = "ExtattrSetLink"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FaccessatKey = "Faccessat"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

	FadviseKey = "Fadvise"

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

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodKey = "Fchmod"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

	FchmodatKey = "Fchmodat"

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

	FcntlFlockKey = "FcntlFlock"

	FcntlIntKey = "FcntlInt"

	FcntlIntKey = "FcntlInt"

	FcntlIntKey = "FcntlInt"

	FcntlIntKey = "FcntlInt"

	FcntlIntKey = "FcntlInt"

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

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FdatasyncKey = "Fdatasync"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FgetxattrKey = "Fgetxattr"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FinitModuleKey = "FinitModule"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

	FlistxattrKey = "Flistxattr"

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

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

	FlockKey = "Flock"

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

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FpathconfKey = "Fpathconf"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FremovexattrKey = "Fremovexattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

	FsetxattrKey = "Fsetxattr"

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

	FstatKey = "Fstat"

	FstatKey = "Fstat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

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

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatfsKey = "Fstatfs"

	FstatvfsKey = "Fstatvfs"

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

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FtruncateKey = "Ftruncate"

	FutimesKey = "Futimes"

	FutimesKey = "Futimes"

	FutimesKey = "Futimes"

	FutimesatKey = "Futimesat"

	FutimesatKey = "Futimesat"

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

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

	GetegidKey = "Getegid"

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

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

	GeteuidKey = "Geteuid"

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

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgidKey = "Getgid"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GetgroupsKey = "Getgroups"

	GethostnameKey = "Gethostname"

	GetpagesizeKey = "Getpagesize"

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

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetpriorityKey = "Getpriority"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

	GetrandomKey = "Getrandom"

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

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrlimitKey = "Getrlimit"

	GetrtableKey = "Getrtable"

	GetrtableKey = "Getrtable"

	GetrtableKey = "Getrtable"

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

	GetsockoptByteKey = "GetsockoptByte"

	GetsockoptICMPv6FilterKey = "GetsockoptICMPv6Filter"

	GetsockoptIPMreqKey = "GetsockoptIPMreq"

	GetsockoptIPMreqnKey = "GetsockoptIPMreqn"

	GetsockoptIPMreqnKey = "GetsockoptIPMreqn"

	GetsockoptIPv6MTUInfoKey = "GetsockoptIPv6MTUInfo"

	GetsockoptIPv6MreqKey = "GetsockoptIPv6Mreq"

	GetsockoptInet4AddrKey = "GetsockoptInet4Addr"

	GetsockoptIntKey = "GetsockoptInt"

	GetsockoptLingerKey = "GetsockoptLinger"

	GetsockoptStringKey = "GetsockoptString"

	GetsockoptStringKey = "GetsockoptString"

	GetsockoptStringKey = "GetsockoptString"

	GetsockoptTCPInfoKey = "GetsockoptTCPInfo"

	GetsockoptTimevalKey = "GetsockoptTimeval"

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

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	GetxattrKey = "Getxattr"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

	InitModuleKey = "InitModule"

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

	InotifyRmWatchKey = "InotifyRmWatch"

	InotifyRmWatchKey = "InotifyRmWatch"

	IoctlGetIntKey = "IoctlGetInt"

	IoctlGetIntKey = "IoctlGetInt"

	IoctlGetIntKey = "IoctlGetInt"

	IoctlGetIntKey = "IoctlGetInt"

	IoctlGetIntKey = "IoctlGetInt"

	IoctlGetIntKey = "IoctlGetInt"

	IoctlGetIntKey = "IoctlGetInt"

	IoctlGetIntKey = "IoctlGetInt"

	IoctlGetPtmgetKey = "IoctlGetPtmget"

	IoctlGetTermioKey = "IoctlGetTermio"

	IoctlGetTermiosKey = "IoctlGetTermios"

	IoctlGetTermiosKey = "IoctlGetTermios"

	IoctlGetTermiosKey = "IoctlGetTermios"

	IoctlGetTermiosKey = "IoctlGetTermios"

	IoctlGetTermiosKey = "IoctlGetTermios"

	IoctlGetTermiosKey = "IoctlGetTermios"

	IoctlGetTermiosKey = "IoctlGetTermios"

	IoctlGetTermiosKey = "IoctlGetTermios"

	IoctlGetWinsizeKey = "IoctlGetWinsize"

	IoctlGetWinsizeKey = "IoctlGetWinsize"

	IoctlGetWinsizeKey = "IoctlGetWinsize"

	IoctlGetWinsizeKey = "IoctlGetWinsize"

	IoctlGetWinsizeKey = "IoctlGetWinsize"

	IoctlGetWinsizeKey = "IoctlGetWinsize"

	IoctlGetWinsizeKey = "IoctlGetWinsize"

	IoctlGetWinsizeKey = "IoctlGetWinsize"

	IoctlSetIntKey = "IoctlSetInt"

	IoctlSetIntKey = "IoctlSetInt"

	IoctlSetIntKey = "IoctlSetInt"

	IoctlSetIntKey = "IoctlSetInt"

	IoctlSetIntKey = "IoctlSetInt"

	IoctlSetIntKey = "IoctlSetInt"

	IoctlSetIntKey = "IoctlSetInt"

	IoctlSetIntKey = "IoctlSetInt"

	IoctlSetPointerIntKey = "IoctlSetPointerInt"

	IoctlSetTermioKey = "IoctlSetTermio"

	IoctlSetTermiosKey = "IoctlSetTermios"

	IoctlSetWinsizeKey = "IoctlSetWinsize"

	IopermKey = "Ioperm"

	IopermKey = "Ioperm"

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

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	IssetugidKey = "Issetugid"

	KeventKey = "Kevent"

	KexecFileLoadKey = "KexecFileLoad"

	KexecFileLoadKey = "KexecFileLoad"

	KexecFileLoadKey = "KexecFileLoad"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlBufferKey = "KeyctlBuffer"

	KeyctlDHComputeKey = "KeyctlDHCompute"

	KeyctlGetKeyringIDKey = "KeyctlGetKeyringID"

	KeyctlInstantiateIOVKey = "KeyctlInstantiateIOV"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlIntKey = "KeyctlInt"

	KeyctlJoinSessionKeyringKey = "KeyctlJoinSessionKeyring"

	KeyctlSearchKey = "KeyctlSearch"

	KeyctlSetpermKey = "KeyctlSetperm"

	KeyctlStringKey = "KeyctlString"

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

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LchownKey = "Lchown"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

	LgetxattrKey = "Lgetxattr"

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

	LinkKey = "Link"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

	LinkatKey = "Linkat"

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

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	ListxattrKey = "Listxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LlistxattrKey = "Llistxattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LremovexattrKey = "Lremovexattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

	LsetxattrKey = "Lsetxattr"

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

	MadviseKey = "Madvise"

	MadviseKey = "Madvise"

	MajorKey = "Major"

	MajorKey = "Major"

	MajorKey = "Major"

	MajorKey = "Major"

	MajorKey = "Major"

	MajorKey = "Major"

	MajorKey = "Major"

	MajorKey = "Major"

	MajorKey = "Major"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MemfdCreateKey = "MemfdCreate"

	MinorKey = "Minor"

	MinorKey = "Minor"

	MinorKey = "Minor"

	MinorKey = "Minor"

	MinorKey = "Minor"

	MinorKey = "Minor"

	MinorKey = "Minor"

	MinorKey = "Minor"

	MinorKey = "Minor"

	MkdevKey = "Mkdev"

	MkdevKey = "Mkdev"

	MkdevKey = "Mkdev"

	MkdevKey = "Mkdev"

	MkdevKey = "Mkdev"

	MkdevKey = "Mkdev"

	MkdevKey = "Mkdev"

	MkdevKey = "Mkdev"

	MkdevKey = "Mkdev"

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

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoKey = "Mkfifo"

	MkfifoatKey = "Mkfifoat"

	MkfifoatKey = "Mkfifoat"

	MkfifoatKey = "Mkfifoat"

	MkfifoatKey = "Mkfifoat"

	MkfifoatKey = "Mkfifoat"

	MkfifoatKey = "Mkfifoat"

	MkfifoatKey = "Mkfifoat"

	MkfifoatKey = "Mkfifoat"

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

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MlockallKey = "Mlockall"

	MmapKey = "Mmap"

	MmapKey = "Mmap"

	MmapKey = "Mmap"

	MmapKey = "Mmap"

	MountKey = "Mount"

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

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MprotectKey = "Mprotect"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

	MsyncKey = "Msync"

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

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunlockallKey = "Munlockall"

	MunmapKey = "Munmap"

	MunmapKey = "Munmap"

	MunmapKey = "Munmap"

	MunmapKey = "Munmap"

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

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NanosleepKey = "Nanosleep"

	NsecToTimespecKey = "NsecToTimespec"

	NsecToTimevalKey = "NsecToTimeval"

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

	OpenKey = "Open"

	OpenKey = "Open"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	ParseDirentKey = "ParseDirent"

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

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PauseKey = "Pause"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

	PerfEventOpenKey = "PerfEventOpen"

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

	PivotRootKey = "PivotRoot"

	PivotRootKey = "PivotRoot"

	PledgeKey = "Pledge"

	PledgeExecpromisesKey = "PledgeExecpromises"

	PledgePromisesKey = "PledgePromises"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PollKey = "Poll"

	PpollKey = "Ppoll"

	PpollKey = "Ppoll"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

	PrctlKey = "Prctl"

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

	PreadKey = "Pread"

	PreadKey = "Pread"

	PreadKey = "Pread"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PselectKey = "Pselect"

	PtraceAttachKey = "PtraceAttach"

	PtraceAttachKey = "PtraceAttach"

	PtraceContKey = "PtraceCont"

	PtraceDetachKey = "PtraceDetach"

	PtraceDetachKey = "PtraceDetach"

	PtraceGetEventMsgKey = "PtraceGetEventMsg"

	PtraceGetRegsKey = "PtraceGetRegs"

	PtraceGetRegs386Key = "PtraceGetRegs386"

	PtraceGetRegsAmd64Key = "PtraceGetRegsAmd64"

	PtraceGetRegsArmKey = "PtraceGetRegsArm"

	PtraceGetRegsArm64Key = "PtraceGetRegsArm64"

	PtraceGetRegsMipsKey = "PtraceGetRegsMips"

	PtraceGetRegsMips64Key = "PtraceGetRegsMips64"

	PtraceGetRegsMips64leKey = "PtraceGetRegsMips64le"

	PtraceGetRegsMipsleKey = "PtraceGetRegsMipsle"

	PtracePeekDataKey = "PtracePeekData"

	PtracePeekTextKey = "PtracePeekText"

	PtracePeekUserKey = "PtracePeekUser"

	PtracePokeDataKey = "PtracePokeData"

	PtracePokeTextKey = "PtracePokeText"

	PtracePokeUserKey = "PtracePokeUser"

	PtraceSetOptionsKey = "PtraceSetOptions"

	PtraceSetRegsKey = "PtraceSetRegs"

	PtraceSetRegs386Key = "PtraceSetRegs386"

	PtraceSetRegsAmd64Key = "PtraceSetRegsAmd64"

	PtraceSetRegsArmKey = "PtraceSetRegsArm"

	PtraceSetRegsArm64Key = "PtraceSetRegsArm64"

	PtraceSetRegsMipsKey = "PtraceSetRegsMips"

	PtraceSetRegsMips64Key = "PtraceSetRegsMips64"

	PtraceSetRegsMips64leKey = "PtraceSetRegsMips64le"

	PtraceSetRegsMipsleKey = "PtraceSetRegsMipsle"

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

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	PwriteKey = "Pwrite"

	RawSyscallKey = "RawSyscall"

	RawSyscallKey = "RawSyscall"

	RawSyscallKey = "RawSyscall"

	RawSyscall6Key = "RawSyscall6"

	RawSyscall6Key = "RawSyscall6"

	RawSyscall6Key = "RawSyscall6"

	RawSyscallNoErrorKey = "RawSyscallNoError"

	RawSyscallNoErrorKey = "RawSyscallNoError"

	ReadKey = "Read"

	ReadDirentKey = "ReadDirent"

	ReadDirentKey = "ReadDirent"

	ReadDirentKey = "ReadDirent"

	ReadDirentKey = "ReadDirent"

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

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkKey = "Readlink"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	ReadlinkatKey = "Readlinkat"

	RebootKey = "Reboot"

	RecvfromKey = "Recvfrom"

	RecvmsgKey = "Recvmsg"

	RecvmsgKey = "Recvmsg"

	RecvmsgKey = "Recvmsg"

	RecvmsgKey = "Recvmsg"

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

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	Renameat2Key = "Renameat2"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

	RequestKeyKey = "RequestKey"

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

	RmdirKey = "Rmdir"

	SchedGetaffinityKey = "SchedGetaffinity"

	SchedSetaffinityKey = "SchedSetaffinity"

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

	SendfileKey = "Sendfile"

	SendfileKey = "Sendfile"

	SendfileKey = "Sendfile"

	SendfileKey = "Sendfile"

	SendfileKey = "Sendfile"

	SendmsgKey = "Sendmsg"

	SendmsgKey = "Sendmsg"

	SendmsgKey = "Sendmsg"

	SendmsgKey = "Sendmsg"

	SendmsgNKey = "SendmsgN"

	SendmsgNKey = "SendmsgN"

	SendmsgNKey = "SendmsgN"

	SendmsgNKey = "SendmsgN"

	SendtoKey = "Sendto"

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

	SetKeventKey = "SetKevent"

	SetNonblockKey = "SetNonblock"

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

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

	SetegidKey = "Setegid"

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

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetloginKey = "Setlogin"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

	SetnsKey = "Setns"

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

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrlimitKey = "Setrlimit"

	SetrtableKey = "Setrtable"

	SetrtableKey = "Setrtable"

	SetrtableKey = "Setrtable"

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

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsidKey = "Setsid"

	SetsockoptByteKey = "SetsockoptByte"

	SetsockoptICMPv6FilterKey = "SetsockoptICMPv6Filter"

	SetsockoptIPMreqKey = "SetsockoptIPMreq"

	SetsockoptIPMreqnKey = "SetsockoptIPMreqn"

	SetsockoptIPMreqnKey = "SetsockoptIPMreqn"

	SetsockoptIPv6MreqKey = "SetsockoptIPv6Mreq"

	SetsockoptInet4AddrKey = "SetsockoptInet4Addr"

	SetsockoptIntKey = "SetsockoptInt"

	SetsockoptLingerKey = "SetsockoptLinger"

	SetsockoptStringKey = "SetsockoptString"

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

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	ShutdownKey = "Shutdown"

	SignalNameKey = "SignalName"

	SocketKey = "Socket"

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

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

	SpliceKey = "Splice"

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

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatfsKey = "Statfs"

	StatvfsKey = "Statvfs"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

	StatxKey = "Statx"

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

	SymlinkKey = "Symlink"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

	SymlinkatKey = "Symlinkat"

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

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncFileRangeKey = "SyncFileRange"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyncfsKey = "Syncfs"

	SyscallKey = "Syscall"

	SyscallKey = "Syscall"

	SyscallKey = "Syscall"

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

	Syscall9Key = "Syscall9"

	SyscallNoErrorKey = "SyscallNoError"

	SyscallNoErrorKey = "SyscallNoError"

	SysctlKey = "Sysctl"

	SysctlArgsKey = "SysctlArgs"

	SysctlClockinfoKey = "SysctlClockinfo"

	SysctlRawKey = "SysctlRaw"

	SysctlUint32Key = "SysctlUint32"

	SysctlUint32ArgsKey = "SysctlUint32Args"

	SysctlUint64Key = "SysctlUint64"

	SysctlUvmexpKey = "SysctlUvmexp"

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

	TeeKey = "Tee"

	TeeKey = "Tee"

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

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeKey = "Time"

	TimeToTimespecKey = "TimeToTimespec"

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

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimesKey = "Times"

	TimespecToNsecKey = "TimespecToNsec"

	TimevalToNsecKey = "TimevalToNsec"

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

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

	TruncateKey = "Truncate"

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

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkKey = "Unlink"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

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

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

	UnmountKey = "Unmount"

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

	UnshareKey = "Unshare"

	UnshareKey = "Unshare"

	UnveilKey = "Unveil"

	UnveilBlockKey = "UnveilBlock"

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

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimeKey = "Utime"

	UtimesKey = "Utimes"

	UtimesKey = "Utimes"

	UtimesKey = "Utimes"

	UtimesKey = "Utimes"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoKey = "UtimesNano"

	UtimesNanoAtKey = "UtimesNanoAt"

	UtimesNanoAtKey = "UtimesNanoAt"

	UtimesNanoAtKey = "UtimesNanoAt"

	UtimesNanoAtKey = "UtimesNanoAt"

	VmspliceKey = "Vmsplice"

	Wait4Key = "Wait4"

	Wait4Key = "Wait4"

	Wait4Key = "Wait4"

	Wait4Key = "Wait4"

	WriteKey = "Write"
)

func New() hctx.Map {
	return hctx.Map{

		AcceptKey: Accept,

		AcceptKey: Accept,

		AcceptKey: Accept,

		AcceptKey: Accept,

		Accept4Key: Accept4,

		Accept4Key: Accept4,

		Accept4Key: Accept4,

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

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AcctKey: Acct,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

		AddKeyKey: AddKey,

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

		AdjtimexKey: Adjtimex,

		AdjtimexKey: Adjtimex,

		BindKey: Bind,

		BindToDeviceKey: BindToDevice,

		BytePtrFromStringKey: BytePtrFromString,

		ByteSliceFromStringKey: ByteSliceFromString,

		CapEnterKey: CapEnter,

		CapEnterKey: CapEnter,

		CapEnterKey: CapEnter,

		CapEnterKey: CapEnter,

		CapRightsClearKey: CapRightsClear,

		CapRightsGetKey: CapRightsGet,

		CapRightsInitKey: CapRightsInit,

		CapRightsIsSetKey: CapRightsIsSet,

		CapRightsLimitKey: CapRightsLimit,

		CapRightsSetKey: CapRightsSet,

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

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

		ChdirKey: Chdir,

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

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGetresKey: ClockGetres,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

		ClockGettimeKey: ClockGettime,

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

		CloseKey: Close,

		CloseKey: Close,

		CloseOnExecKey: CloseOnExec,

		CmsgLenKey: CmsgLen,

		CmsgSpaceKey: CmsgSpace,

		ConnectKey: Connect,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CopyFileRangeKey: CopyFileRange,

		CreatKey: Creat,

		CreatKey: Creat,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

		DeleteModuleKey: DeleteModule,

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

		Dup3Key: Dup3,

		Dup3Key: Dup3,

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

		EpollWaitKey: EpollWait,

		EpollWaitKey: EpollWait,

		ErrnoNameKey: ErrnoName,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		EventfdKey: Eventfd,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExchangedataKey: Exchangedata,

		ExecKey: Exec,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExitKey: Exit,

		ExtattrDeleteFdKey: ExtattrDeleteFd,

		ExtattrDeleteFdKey: ExtattrDeleteFd,

		ExtattrDeleteFdKey: ExtattrDeleteFd,

		ExtattrDeleteFdKey: ExtattrDeleteFd,

		ExtattrDeleteFdKey: ExtattrDeleteFd,

		ExtattrDeleteFdKey: ExtattrDeleteFd,

		ExtattrDeleteFdKey: ExtattrDeleteFd,

		ExtattrDeleteFileKey: ExtattrDeleteFile,

		ExtattrDeleteFileKey: ExtattrDeleteFile,

		ExtattrDeleteFileKey: ExtattrDeleteFile,

		ExtattrDeleteFileKey: ExtattrDeleteFile,

		ExtattrDeleteFileKey: ExtattrDeleteFile,

		ExtattrDeleteFileKey: ExtattrDeleteFile,

		ExtattrDeleteFileKey: ExtattrDeleteFile,

		ExtattrDeleteLinkKey: ExtattrDeleteLink,

		ExtattrDeleteLinkKey: ExtattrDeleteLink,

		ExtattrDeleteLinkKey: ExtattrDeleteLink,

		ExtattrDeleteLinkKey: ExtattrDeleteLink,

		ExtattrDeleteLinkKey: ExtattrDeleteLink,

		ExtattrDeleteLinkKey: ExtattrDeleteLink,

		ExtattrDeleteLinkKey: ExtattrDeleteLink,

		ExtattrGetFdKey: ExtattrGetFd,

		ExtattrGetFdKey: ExtattrGetFd,

		ExtattrGetFdKey: ExtattrGetFd,

		ExtattrGetFdKey: ExtattrGetFd,

		ExtattrGetFdKey: ExtattrGetFd,

		ExtattrGetFdKey: ExtattrGetFd,

		ExtattrGetFdKey: ExtattrGetFd,

		ExtattrGetFileKey: ExtattrGetFile,

		ExtattrGetFileKey: ExtattrGetFile,

		ExtattrGetFileKey: ExtattrGetFile,

		ExtattrGetFileKey: ExtattrGetFile,

		ExtattrGetFileKey: ExtattrGetFile,

		ExtattrGetFileKey: ExtattrGetFile,

		ExtattrGetFileKey: ExtattrGetFile,

		ExtattrGetLinkKey: ExtattrGetLink,

		ExtattrGetLinkKey: ExtattrGetLink,

		ExtattrGetLinkKey: ExtattrGetLink,

		ExtattrGetLinkKey: ExtattrGetLink,

		ExtattrGetLinkKey: ExtattrGetLink,

		ExtattrGetLinkKey: ExtattrGetLink,

		ExtattrGetLinkKey: ExtattrGetLink,

		ExtattrListFdKey: ExtattrListFd,

		ExtattrListFdKey: ExtattrListFd,

		ExtattrListFdKey: ExtattrListFd,

		ExtattrListFdKey: ExtattrListFd,

		ExtattrListFdKey: ExtattrListFd,

		ExtattrListFdKey: ExtattrListFd,

		ExtattrListFdKey: ExtattrListFd,

		ExtattrListFileKey: ExtattrListFile,

		ExtattrListFileKey: ExtattrListFile,

		ExtattrListFileKey: ExtattrListFile,

		ExtattrListFileKey: ExtattrListFile,

		ExtattrListFileKey: ExtattrListFile,

		ExtattrListFileKey: ExtattrListFile,

		ExtattrListFileKey: ExtattrListFile,

		ExtattrListLinkKey: ExtattrListLink,

		ExtattrListLinkKey: ExtattrListLink,

		ExtattrListLinkKey: ExtattrListLink,

		ExtattrListLinkKey: ExtattrListLink,

		ExtattrListLinkKey: ExtattrListLink,

		ExtattrListLinkKey: ExtattrListLink,

		ExtattrListLinkKey: ExtattrListLink,

		ExtattrSetFdKey: ExtattrSetFd,

		ExtattrSetFdKey: ExtattrSetFd,

		ExtattrSetFdKey: ExtattrSetFd,

		ExtattrSetFdKey: ExtattrSetFd,

		ExtattrSetFdKey: ExtattrSetFd,

		ExtattrSetFdKey: ExtattrSetFd,

		ExtattrSetFdKey: ExtattrSetFd,

		ExtattrSetFileKey: ExtattrSetFile,

		ExtattrSetFileKey: ExtattrSetFile,

		ExtattrSetFileKey: ExtattrSetFile,

		ExtattrSetFileKey: ExtattrSetFile,

		ExtattrSetFileKey: ExtattrSetFile,

		ExtattrSetFileKey: ExtattrSetFile,

		ExtattrSetFileKey: ExtattrSetFile,

		ExtattrSetLinkKey: ExtattrSetLink,

		ExtattrSetLinkKey: ExtattrSetLink,

		ExtattrSetLinkKey: ExtattrSetLink,

		ExtattrSetLinkKey: ExtattrSetLink,

		ExtattrSetLinkKey: ExtattrSetLink,

		ExtattrSetLinkKey: ExtattrSetLink,

		ExtattrSetLinkKey: ExtattrSetLink,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FaccessatKey: Faccessat,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

		FadviseKey: Fadvise,

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

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodKey: Fchmod,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

		FchmodatKey: Fchmodat,

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

		FcntlFlockKey: FcntlFlock,

		FcntlIntKey: FcntlInt,

		FcntlIntKey: FcntlInt,

		FcntlIntKey: FcntlInt,

		FcntlIntKey: FcntlInt,

		FcntlIntKey: FcntlInt,

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

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FdatasyncKey: Fdatasync,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FgetxattrKey: Fgetxattr,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FinitModuleKey: FinitModule,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

		FlistxattrKey: Flistxattr,

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

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

		FlockKey: Flock,

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

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FpathconfKey: Fpathconf,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FremovexattrKey: Fremovexattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

		FsetxattrKey: Fsetxattr,

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

		FstatKey: Fstat,

		FstatKey: Fstat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

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

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatfsKey: Fstatfs,

		FstatvfsKey: Fstatvfs,

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

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FtruncateKey: Ftruncate,

		FutimesKey: Futimes,

		FutimesKey: Futimes,

		FutimesKey: Futimes,

		FutimesatKey: Futimesat,

		FutimesatKey: Futimesat,

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

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

		GetegidKey: Getegid,

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

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

		GeteuidKey: Geteuid,

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

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgidKey: Getgid,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GetgroupsKey: Getgroups,

		GethostnameKey: Gethostname,

		GetpagesizeKey: Getpagesize,

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

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetpriorityKey: Getpriority,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

		GetrandomKey: Getrandom,

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

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrlimitKey: Getrlimit,

		GetrtableKey: Getrtable,

		GetrtableKey: Getrtable,

		GetrtableKey: Getrtable,

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

		GetsockoptByteKey: GetsockoptByte,

		GetsockoptICMPv6FilterKey: GetsockoptICMPv6Filter,

		GetsockoptIPMreqKey: GetsockoptIPMreq,

		GetsockoptIPMreqnKey: GetsockoptIPMreqn,

		GetsockoptIPMreqnKey: GetsockoptIPMreqn,

		GetsockoptIPv6MTUInfoKey: GetsockoptIPv6MTUInfo,

		GetsockoptIPv6MreqKey: GetsockoptIPv6Mreq,

		GetsockoptInet4AddrKey: GetsockoptInet4Addr,

		GetsockoptIntKey: GetsockoptInt,

		GetsockoptLingerKey: GetsockoptLinger,

		GetsockoptStringKey: GetsockoptString,

		GetsockoptStringKey: GetsockoptString,

		GetsockoptStringKey: GetsockoptString,

		GetsockoptTCPInfoKey: GetsockoptTCPInfo,

		GetsockoptTimevalKey: GetsockoptTimeval,

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

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		GetxattrKey: Getxattr,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

		InitModuleKey: InitModule,

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

		InotifyRmWatchKey: InotifyRmWatch,

		InotifyRmWatchKey: InotifyRmWatch,

		IoctlGetIntKey: IoctlGetInt,

		IoctlGetIntKey: IoctlGetInt,

		IoctlGetIntKey: IoctlGetInt,

		IoctlGetIntKey: IoctlGetInt,

		IoctlGetIntKey: IoctlGetInt,

		IoctlGetIntKey: IoctlGetInt,

		IoctlGetIntKey: IoctlGetInt,

		IoctlGetIntKey: IoctlGetInt,

		IoctlGetPtmgetKey: IoctlGetPtmget,

		IoctlGetTermioKey: IoctlGetTermio,

		IoctlGetTermiosKey: IoctlGetTermios,

		IoctlGetTermiosKey: IoctlGetTermios,

		IoctlGetTermiosKey: IoctlGetTermios,

		IoctlGetTermiosKey: IoctlGetTermios,

		IoctlGetTermiosKey: IoctlGetTermios,

		IoctlGetTermiosKey: IoctlGetTermios,

		IoctlGetTermiosKey: IoctlGetTermios,

		IoctlGetTermiosKey: IoctlGetTermios,

		IoctlGetWinsizeKey: IoctlGetWinsize,

		IoctlGetWinsizeKey: IoctlGetWinsize,

		IoctlGetWinsizeKey: IoctlGetWinsize,

		IoctlGetWinsizeKey: IoctlGetWinsize,

		IoctlGetWinsizeKey: IoctlGetWinsize,

		IoctlGetWinsizeKey: IoctlGetWinsize,

		IoctlGetWinsizeKey: IoctlGetWinsize,

		IoctlGetWinsizeKey: IoctlGetWinsize,

		IoctlSetIntKey: IoctlSetInt,

		IoctlSetIntKey: IoctlSetInt,

		IoctlSetIntKey: IoctlSetInt,

		IoctlSetIntKey: IoctlSetInt,

		IoctlSetIntKey: IoctlSetInt,

		IoctlSetIntKey: IoctlSetInt,

		IoctlSetIntKey: IoctlSetInt,

		IoctlSetIntKey: IoctlSetInt,

		IoctlSetPointerIntKey: IoctlSetPointerInt,

		IoctlSetTermioKey: IoctlSetTermio,

		IoctlSetTermiosKey: IoctlSetTermios,

		IoctlSetWinsizeKey: IoctlSetWinsize,

		IopermKey: Ioperm,

		IopermKey: Ioperm,

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

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		IssetugidKey: Issetugid,

		KeventKey: Kevent,

		KexecFileLoadKey: KexecFileLoad,

		KexecFileLoadKey: KexecFileLoad,

		KexecFileLoadKey: KexecFileLoad,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlBufferKey: KeyctlBuffer,

		KeyctlDHComputeKey: KeyctlDHCompute,

		KeyctlGetKeyringIDKey: KeyctlGetKeyringID,

		KeyctlInstantiateIOVKey: KeyctlInstantiateIOV,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlIntKey: KeyctlInt,

		KeyctlJoinSessionKeyringKey: KeyctlJoinSessionKeyring,

		KeyctlSearchKey: KeyctlSearch,

		KeyctlSetpermKey: KeyctlSetperm,

		KeyctlStringKey: KeyctlString,

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

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LchownKey: Lchown,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

		LgetxattrKey: Lgetxattr,

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

		LinkKey: Link,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

		LinkatKey: Linkat,

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

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		ListxattrKey: Listxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LlistxattrKey: Llistxattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LremovexattrKey: Lremovexattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

		LsetxattrKey: Lsetxattr,

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

		MadviseKey: Madvise,

		MadviseKey: Madvise,

		MajorKey: Major,

		MajorKey: Major,

		MajorKey: Major,

		MajorKey: Major,

		MajorKey: Major,

		MajorKey: Major,

		MajorKey: Major,

		MajorKey: Major,

		MajorKey: Major,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MemfdCreateKey: MemfdCreate,

		MinorKey: Minor,

		MinorKey: Minor,

		MinorKey: Minor,

		MinorKey: Minor,

		MinorKey: Minor,

		MinorKey: Minor,

		MinorKey: Minor,

		MinorKey: Minor,

		MinorKey: Minor,

		MkdevKey: Mkdev,

		MkdevKey: Mkdev,

		MkdevKey: Mkdev,

		MkdevKey: Mkdev,

		MkdevKey: Mkdev,

		MkdevKey: Mkdev,

		MkdevKey: Mkdev,

		MkdevKey: Mkdev,

		MkdevKey: Mkdev,

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

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoKey: Mkfifo,

		MkfifoatKey: Mkfifoat,

		MkfifoatKey: Mkfifoat,

		MkfifoatKey: Mkfifoat,

		MkfifoatKey: Mkfifoat,

		MkfifoatKey: Mkfifoat,

		MkfifoatKey: Mkfifoat,

		MkfifoatKey: Mkfifoat,

		MkfifoatKey: Mkfifoat,

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

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MlockallKey: Mlockall,

		MmapKey: Mmap,

		MmapKey: Mmap,

		MmapKey: Mmap,

		MmapKey: Mmap,

		MountKey: Mount,

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

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MprotectKey: Mprotect,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

		MsyncKey: Msync,

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

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunlockallKey: Munlockall,

		MunmapKey: Munmap,

		MunmapKey: Munmap,

		MunmapKey: Munmap,

		MunmapKey: Munmap,

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

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NanosleepKey: Nanosleep,

		NsecToTimespecKey: NsecToTimespec,

		NsecToTimevalKey: NsecToTimeval,

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

		OpenKey: Open,

		OpenKey: Open,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		ParseDirentKey: ParseDirent,

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

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PauseKey: Pause,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

		PerfEventOpenKey: PerfEventOpen,

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

		PivotRootKey: PivotRoot,

		PivotRootKey: PivotRoot,

		PledgeKey: Pledge,

		PledgeExecpromisesKey: PledgeExecpromises,

		PledgePromisesKey: PledgePromises,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PollKey: Poll,

		PpollKey: Ppoll,

		PpollKey: Ppoll,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

		PrctlKey: Prctl,

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

		PreadKey: Pread,

		PreadKey: Pread,

		PreadKey: Pread,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PselectKey: Pselect,

		PtraceAttachKey: PtraceAttach,

		PtraceAttachKey: PtraceAttach,

		PtraceContKey: PtraceCont,

		PtraceDetachKey: PtraceDetach,

		PtraceDetachKey: PtraceDetach,

		PtraceGetEventMsgKey: PtraceGetEventMsg,

		PtraceGetRegsKey: PtraceGetRegs,

		PtraceGetRegs386Key: PtraceGetRegs386,

		PtraceGetRegsAmd64Key: PtraceGetRegsAmd64,

		PtraceGetRegsArmKey: PtraceGetRegsArm,

		PtraceGetRegsArm64Key: PtraceGetRegsArm64,

		PtraceGetRegsMipsKey: PtraceGetRegsMips,

		PtraceGetRegsMips64Key: PtraceGetRegsMips64,

		PtraceGetRegsMips64leKey: PtraceGetRegsMips64le,

		PtraceGetRegsMipsleKey: PtraceGetRegsMipsle,

		PtracePeekDataKey: PtracePeekData,

		PtracePeekTextKey: PtracePeekText,

		PtracePeekUserKey: PtracePeekUser,

		PtracePokeDataKey: PtracePokeData,

		PtracePokeTextKey: PtracePokeText,

		PtracePokeUserKey: PtracePokeUser,

		PtraceSetOptionsKey: PtraceSetOptions,

		PtraceSetRegsKey: PtraceSetRegs,

		PtraceSetRegs386Key: PtraceSetRegs386,

		PtraceSetRegsAmd64Key: PtraceSetRegsAmd64,

		PtraceSetRegsArmKey: PtraceSetRegsArm,

		PtraceSetRegsArm64Key: PtraceSetRegsArm64,

		PtraceSetRegsMipsKey: PtraceSetRegsMips,

		PtraceSetRegsMips64Key: PtraceSetRegsMips64,

		PtraceSetRegsMips64leKey: PtraceSetRegsMips64le,

		PtraceSetRegsMipsleKey: PtraceSetRegsMipsle,

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

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		PwriteKey: Pwrite,

		RawSyscallKey: RawSyscall,

		RawSyscallKey: RawSyscall,

		RawSyscallKey: RawSyscall,

		RawSyscall6Key: RawSyscall6,

		RawSyscall6Key: RawSyscall6,

		RawSyscall6Key: RawSyscall6,

		RawSyscallNoErrorKey: RawSyscallNoError,

		RawSyscallNoErrorKey: RawSyscallNoError,

		ReadKey: Read,

		ReadDirentKey: ReadDirent,

		ReadDirentKey: ReadDirent,

		ReadDirentKey: ReadDirent,

		ReadDirentKey: ReadDirent,

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

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkKey: Readlink,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		ReadlinkatKey: Readlinkat,

		RebootKey: Reboot,

		RecvfromKey: Recvfrom,

		RecvmsgKey: Recvmsg,

		RecvmsgKey: Recvmsg,

		RecvmsgKey: Recvmsg,

		RecvmsgKey: Recvmsg,

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

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		Renameat2Key: Renameat2,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

		RequestKeyKey: RequestKey,

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

		RmdirKey: Rmdir,

		SchedGetaffinityKey: SchedGetaffinity,

		SchedSetaffinityKey: SchedSetaffinity,

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

		SendfileKey: Sendfile,

		SendfileKey: Sendfile,

		SendfileKey: Sendfile,

		SendfileKey: Sendfile,

		SendfileKey: Sendfile,

		SendmsgKey: Sendmsg,

		SendmsgKey: Sendmsg,

		SendmsgKey: Sendmsg,

		SendmsgKey: Sendmsg,

		SendmsgNKey: SendmsgN,

		SendmsgNKey: SendmsgN,

		SendmsgNKey: SendmsgN,

		SendmsgNKey: SendmsgN,

		SendtoKey: Sendto,

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

		SetKeventKey: SetKevent,

		SetNonblockKey: SetNonblock,

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

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

		SetegidKey: Setegid,

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

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetloginKey: Setlogin,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

		SetnsKey: Setns,

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

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrlimitKey: Setrlimit,

		SetrtableKey: Setrtable,

		SetrtableKey: Setrtable,

		SetrtableKey: Setrtable,

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

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsidKey: Setsid,

		SetsockoptByteKey: SetsockoptByte,

		SetsockoptICMPv6FilterKey: SetsockoptICMPv6Filter,

		SetsockoptIPMreqKey: SetsockoptIPMreq,

		SetsockoptIPMreqnKey: SetsockoptIPMreqn,

		SetsockoptIPMreqnKey: SetsockoptIPMreqn,

		SetsockoptIPv6MreqKey: SetsockoptIPv6Mreq,

		SetsockoptInet4AddrKey: SetsockoptInet4Addr,

		SetsockoptIntKey: SetsockoptInt,

		SetsockoptLingerKey: SetsockoptLinger,

		SetsockoptStringKey: SetsockoptString,

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

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		ShutdownKey: Shutdown,

		SignalNameKey: SignalName,

		SocketKey: Socket,

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

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

		SpliceKey: Splice,

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

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatfsKey: Statfs,

		StatvfsKey: Statvfs,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

		StatxKey: Statx,

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

		SymlinkKey: Symlink,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

		SymlinkatKey: Symlinkat,

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

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncFileRangeKey: SyncFileRange,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyncfsKey: Syncfs,

		SyscallKey: Syscall,

		SyscallKey: Syscall,

		SyscallKey: Syscall,

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

		Syscall9Key: Syscall9,

		SyscallNoErrorKey: SyscallNoError,

		SyscallNoErrorKey: SyscallNoError,

		SysctlKey: Sysctl,

		SysctlArgsKey: SysctlArgs,

		SysctlClockinfoKey: SysctlClockinfo,

		SysctlRawKey: SysctlRaw,

		SysctlUint32Key: SysctlUint32,

		SysctlUint32ArgsKey: SysctlUint32Args,

		SysctlUint64Key: SysctlUint64,

		SysctlUvmexpKey: SysctlUvmexp,

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

		TeeKey: Tee,

		TeeKey: Tee,

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

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeKey: Time,

		TimeToTimespecKey: TimeToTimespec,

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

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimesKey: Times,

		TimespecToNsecKey: TimespecToNsec,

		TimevalToNsecKey: TimevalToNsec,

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

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

		TruncateKey: Truncate,

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

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkKey: Unlink,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

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

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

		UnmountKey: Unmount,

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

		UnshareKey: Unshare,

		UnshareKey: Unshare,

		UnveilKey: Unveil,

		UnveilBlockKey: UnveilBlock,

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

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimeKey: Utime,

		UtimesKey: Utimes,

		UtimesKey: Utimes,

		UtimesKey: Utimes,

		UtimesKey: Utimes,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		UtimesNanoKey: UtimesNano,

		UtimesNanoAtKey: UtimesNanoAt,

		UtimesNanoAtKey: UtimesNanoAt,

		UtimesNanoAtKey: UtimesNanoAt,

		UtimesNanoAtKey: UtimesNanoAt,

		VmspliceKey: Vmsplice,

		Wait4Key: Wait4,

		Wait4Key: Wait4,

		Wait4Key: Wait4,

		Wait4Key: Wait4,

		WriteKey: Write,
	}
}

var Accept = unix.Accept

var Accept = unix.Accept

var Accept = unix.Accept

var Accept = unix.Accept

var Accept4 = unix.Accept4

var Accept4 = unix.Accept4

var Accept4 = unix.Accept4

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Access = unix.Access

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var Acct = unix.Acct

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var AddKey = unix.AddKey

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtime = unix.Adjtime

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Adjtimex = unix.Adjtimex

var Bind = unix.Bind

// BindToDevice binds the socket associated with fd to device.
var BindToDevice = unix.BindToDevice

// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
var BytePtrFromString = unix.BytePtrFromString

// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
var ByteSliceFromString = unix.ByteSliceFromString

var CapEnter = unix.CapEnter

var CapEnter = unix.CapEnter

var CapEnter = unix.CapEnter

var CapEnter = unix.CapEnter

// CapRightsClear clears the permissions in clearrights from rights.
var CapRightsClear = unix.CapRightsClear

// CapRightsGet returns a CapRights structure containing the operations permitted on fd.
// See man cap_rights_get(3) and rights(4).
var CapRightsGet = unix.CapRightsGet

// CapRightsInit returns a pointer to an initialised CapRights structure filled with rights.
// See man cap_rights_init(3) and rights(4).
var CapRightsInit = unix.CapRightsInit

// CapRightsIsSet checks whether all the permissions in setrights are present in rights.
var CapRightsIsSet = unix.CapRightsIsSet

// CapRightsLimit reduces the operations permitted on fd to at most those contained in rights.
// The capability rights on fd can never be increased by CapRightsLimit.
// See man cap_rights_limit(2) and rights(4).
var CapRightsLimit = unix.CapRightsLimit

// CapRightsSet sets the permissions in setrights in rights.
var CapRightsSet = unix.CapRightsSet

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chdir = unix.Chdir

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chflags = unix.Chflags

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chmod = unix.Chmod

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chown = unix.Chown

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Chroot = unix.Chroot

var Clearenv = unix.Clearenv

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGetres = unix.ClockGetres

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var ClockGettime = unix.ClockGettime

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var Close = unix.Close

var CloseOnExec = unix.CloseOnExec

// CmsgLen returns the value to store in the Len field of the Cmsghdr
// structure, taking into account any necessary alignment.
var CmsgLen = unix.CmsgLen

// CmsgSpace returns the number of bytes an ancillary element with
// payload of the passed data length occupies.
var CmsgSpace = unix.CmsgSpace

var Connect = unix.Connect

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var CopyFileRange = unix.CopyFileRange

var Creat = unix.Creat

var Creat = unix.Creat

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var DeleteModule = unix.DeleteModule

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup = unix.Dup

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup2 = unix.Dup2

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Dup3 = unix.Dup3

var Environ = unix.Environ

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate = unix.EpollCreate

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCreate1 = unix.EpollCreate1

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollCtl = unix.EpollCtl

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

var EpollWait = unix.EpollWait

// ErrnoName returns the error name for error number e.
var ErrnoName = unix.ErrnoName

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Eventfd = unix.Eventfd

var Exchangedata = unix.Exchangedata

var Exchangedata = unix.Exchangedata

var Exchangedata = unix.Exchangedata

var Exchangedata = unix.Exchangedata

var Exchangedata = unix.Exchangedata

var Exchangedata = unix.Exchangedata

var Exchangedata = unix.Exchangedata

var Exchangedata = unix.Exchangedata

// Exec calls execve(2), which replaces the calling executable in the process
// tree. argv0 should be the full path to an executable (&#34;/bin/ls&#34;) and the
// executable name should also be the first argument in argv ([&#34;ls&#34;, &#34;-l&#34;]).
// envv are the environment variables that should be passed to the new
// process ([&#34;USER=go&#34;, &#34;PWD=/tmp&#34;]).
var Exec = unix.Exec

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var Exit = unix.Exit

var ExtattrDeleteFd = unix.ExtattrDeleteFd

var ExtattrDeleteFd = unix.ExtattrDeleteFd

var ExtattrDeleteFd = unix.ExtattrDeleteFd

var ExtattrDeleteFd = unix.ExtattrDeleteFd

var ExtattrDeleteFd = unix.ExtattrDeleteFd

var ExtattrDeleteFd = unix.ExtattrDeleteFd

var ExtattrDeleteFd = unix.ExtattrDeleteFd

var ExtattrDeleteFile = unix.ExtattrDeleteFile

var ExtattrDeleteFile = unix.ExtattrDeleteFile

var ExtattrDeleteFile = unix.ExtattrDeleteFile

var ExtattrDeleteFile = unix.ExtattrDeleteFile

var ExtattrDeleteFile = unix.ExtattrDeleteFile

var ExtattrDeleteFile = unix.ExtattrDeleteFile

var ExtattrDeleteFile = unix.ExtattrDeleteFile

var ExtattrDeleteLink = unix.ExtattrDeleteLink

var ExtattrDeleteLink = unix.ExtattrDeleteLink

var ExtattrDeleteLink = unix.ExtattrDeleteLink

var ExtattrDeleteLink = unix.ExtattrDeleteLink

var ExtattrDeleteLink = unix.ExtattrDeleteLink

var ExtattrDeleteLink = unix.ExtattrDeleteLink

var ExtattrDeleteLink = unix.ExtattrDeleteLink

var ExtattrGetFd = unix.ExtattrGetFd

var ExtattrGetFd = unix.ExtattrGetFd

var ExtattrGetFd = unix.ExtattrGetFd

var ExtattrGetFd = unix.ExtattrGetFd

var ExtattrGetFd = unix.ExtattrGetFd

var ExtattrGetFd = unix.ExtattrGetFd

var ExtattrGetFd = unix.ExtattrGetFd

var ExtattrGetFile = unix.ExtattrGetFile

var ExtattrGetFile = unix.ExtattrGetFile

var ExtattrGetFile = unix.ExtattrGetFile

var ExtattrGetFile = unix.ExtattrGetFile

var ExtattrGetFile = unix.ExtattrGetFile

var ExtattrGetFile = unix.ExtattrGetFile

var ExtattrGetFile = unix.ExtattrGetFile

var ExtattrGetLink = unix.ExtattrGetLink

var ExtattrGetLink = unix.ExtattrGetLink

var ExtattrGetLink = unix.ExtattrGetLink

var ExtattrGetLink = unix.ExtattrGetLink

var ExtattrGetLink = unix.ExtattrGetLink

var ExtattrGetLink = unix.ExtattrGetLink

var ExtattrGetLink = unix.ExtattrGetLink

var ExtattrListFd = unix.ExtattrListFd

var ExtattrListFd = unix.ExtattrListFd

var ExtattrListFd = unix.ExtattrListFd

var ExtattrListFd = unix.ExtattrListFd

var ExtattrListFd = unix.ExtattrListFd

var ExtattrListFd = unix.ExtattrListFd

var ExtattrListFd = unix.ExtattrListFd

var ExtattrListFile = unix.ExtattrListFile

var ExtattrListFile = unix.ExtattrListFile

var ExtattrListFile = unix.ExtattrListFile

var ExtattrListFile = unix.ExtattrListFile

var ExtattrListFile = unix.ExtattrListFile

var ExtattrListFile = unix.ExtattrListFile

var ExtattrListFile = unix.ExtattrListFile

var ExtattrListLink = unix.ExtattrListLink

var ExtattrListLink = unix.ExtattrListLink

var ExtattrListLink = unix.ExtattrListLink

var ExtattrListLink = unix.ExtattrListLink

var ExtattrListLink = unix.ExtattrListLink

var ExtattrListLink = unix.ExtattrListLink

var ExtattrListLink = unix.ExtattrListLink

var ExtattrSetFd = unix.ExtattrSetFd

var ExtattrSetFd = unix.ExtattrSetFd

var ExtattrSetFd = unix.ExtattrSetFd

var ExtattrSetFd = unix.ExtattrSetFd

var ExtattrSetFd = unix.ExtattrSetFd

var ExtattrSetFd = unix.ExtattrSetFd

var ExtattrSetFd = unix.ExtattrSetFd

var ExtattrSetFile = unix.ExtattrSetFile

var ExtattrSetFile = unix.ExtattrSetFile

var ExtattrSetFile = unix.ExtattrSetFile

var ExtattrSetFile = unix.ExtattrSetFile

var ExtattrSetFile = unix.ExtattrSetFile

var ExtattrSetFile = unix.ExtattrSetFile

var ExtattrSetFile = unix.ExtattrSetFile

var ExtattrSetLink = unix.ExtattrSetLink

var ExtattrSetLink = unix.ExtattrSetLink

var ExtattrSetLink = unix.ExtattrSetLink

var ExtattrSetLink = unix.ExtattrSetLink

var ExtattrSetLink = unix.ExtattrSetLink

var ExtattrSetLink = unix.ExtattrSetLink

var ExtattrSetLink = unix.ExtattrSetLink

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Faccessat = unix.Faccessat

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fadvise = unix.Fadvise

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fallocate = unix.Fallocate

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchdir = unix.Fchdir

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchflags = unix.Fchflags

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmod = unix.Fchmod

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchmodat = unix.Fchmodat

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchown = unix.Fchown

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var Fchownat = unix.Fchownat

var FcntlFlock = unix.FcntlFlock

var FcntlFlock = unix.FcntlFlock

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
var FcntlFlock = unix.FcntlFlock

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
var FcntlFlock = unix.FcntlFlock

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
var FcntlFlock = unix.FcntlFlock

var FcntlInt = unix.FcntlInt

var FcntlInt = unix.FcntlInt

// FcntlInt performs a fcntl syscall on fd with the provided command and argument.
var FcntlInt = unix.FcntlInt

// FcntlInt performs a fcntl syscall on fd with the provided command and argument.
var FcntlInt = unix.FcntlInt

// FcntlInt performs a fcntl syscall on fd with the provided command and argument.
var FcntlInt = unix.FcntlInt

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fdatasync = unix.Fdatasync

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var Fgetxattr = unix.Fgetxattr

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var FinitModule = unix.FinitModule

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flistxattr = unix.Flistxattr

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Flock = unix.Flock

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fpathconf = unix.Fpathconf

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fremovexattr = unix.Fremovexattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fsetxattr = unix.Fsetxattr

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstat = unix.Fstat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatfs = unix.Fstatfs

var Fstatvfs = unix.Fstatvfs

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Fsync = unix.Fsync

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Ftruncate = unix.Ftruncate

var Futimes = unix.Futimes

var Futimes = unix.Futimes

// Solaris doesn&#39;t have an futimes function because it allows NULL to be
// specified as the path for futimesat. However, Go doesn&#39;t like
// NULL-style string interfaces, so this simple wrapper is provided.
var Futimes = unix.Futimes

var Futimesat = unix.Futimesat

var Futimesat = unix.Futimesat

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getcwd = unix.Getcwd

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdents = unix.Getdents

var Getdirentries = unix.Getdirentries

// sys getdents(fd int, buf []byte) (n int, err error)
var Getdirentries = unix.Getdirentries

var Getdirentries = unix.Getdirentries

var Getdirentries = unix.Getdirentries

var Getdirentries = unix.Getdirentries

var Getdirentries = unix.Getdirentries

var Getdirentries = unix.Getdirentries

// sys getdents(fd int, buf []byte) (n int, err error)
var Getdirentries = unix.Getdirentries

var Getdirentries = unix.Getdirentries

var Getdirentries = unix.Getdirentries

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getdtablesize = unix.Getdtablesize

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getegid = unix.Getegid

var Getenv = unix.Getenv

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Geteuid = unix.Geteuid

var Getfsstat = unix.Getfsstat

var Getfsstat = unix.Getfsstat

var Getfsstat = unix.Getfsstat

var Getfsstat = unix.Getfsstat

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgid = unix.Getgid

var Getgroups = unix.Getgroups

var Getgroups = unix.Getgroups

var Getgroups = unix.Getgroups

var Getgroups = unix.Getgroups

var Gethostname = unix.Gethostname

var Getpagesize = unix.Getpagesize

var Getpeername = unix.Getpeername

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgid = unix.Getpgid

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpgrp = unix.Getpgrp

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getpid = unix.Getpid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getppid = unix.Getppid

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getpriority = unix.Getpriority

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrandom = unix.Getrandom

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrlimit = unix.Getrlimit

var Getrtable = unix.Getrtable

var Getrtable = unix.Getrtable

var Getrtable = unix.Getrtable

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getrusage = unix.Getrusage

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsid = unix.Getsid

var Getsockname = unix.Getsockname

var Getsockname = unix.Getsockname

var Getsockname = unix.Getsockname

var Getsockname = unix.Getsockname

var GetsockoptByte = unix.GetsockoptByte

var GetsockoptICMPv6Filter = unix.GetsockoptICMPv6Filter

var GetsockoptIPMreq = unix.GetsockoptIPMreq

var GetsockoptIPMreqn = unix.GetsockoptIPMreqn

var GetsockoptIPMreqn = unix.GetsockoptIPMreqn

var GetsockoptIPv6MTUInfo = unix.GetsockoptIPv6MTUInfo

var GetsockoptIPv6Mreq = unix.GetsockoptIPv6Mreq

var GetsockoptInet4Addr = unix.GetsockoptInet4Addr

var GetsockoptInt = unix.GetsockoptInt

var GetsockoptLinger = unix.GetsockoptLinger

// GetsockoptString returns the string value of the socket option opt for the
// socket associated with fd at the given socket level.
var GetsockoptString = unix.GetsockoptString

// GetsockoptString returns the string value of the socket option opt for the
// socket associated with fd at the given socket level.
var GetsockoptString = unix.GetsockoptString

// GetsockoptString returns the string value of the socket option opt for the
// socket associated with fd at the given socket level.
var GetsockoptString = unix.GetsockoptString

var GetsockoptTCPInfo = unix.GetsockoptTCPInfo

var GetsockoptTimeval = unix.GetsockoptTimeval

var GetsockoptUcred = unix.GetsockoptUcred

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettid = unix.Gettid

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

// sysnb	gettimeofday(tp *Timeval) (sec int64, usec int32, err error)
var Gettimeofday = unix.Gettimeofday

// sysnb	gettimeofday(tp *Timeval) (sec int32, usec int32, err error)
var Gettimeofday = unix.Gettimeofday

// sysnb	gettimeofday(tp *Timeval) (sec int64, usec int32, err error)
var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

// sysnb	gettimeofday(tp *Timeval) (sec int32, usec int32, err error)
var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Gettimeofday = unix.Gettimeofday

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getuid = unix.Getuid

var Getwd = unix.Getwd

var Getwd = unix.Getwd

var Getwd = unix.Getwd

var Getwd = unix.Getwd

var Getwd = unix.Getwd

var Getwd = unix.Getwd

var Getwd = unix.Getwd

var Getwd = unix.Getwd

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var Getxattr = unix.Getxattr

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InitModule = unix.InitModule

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyAddWatch = unix.InotifyAddWatch

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit = unix.InotifyInit

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyInit1 = unix.InotifyInit1

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

var InotifyRmWatch = unix.InotifyRmWatch

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
var IoctlGetInt = unix.IoctlGetInt

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
var IoctlGetInt = unix.IoctlGetInt

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
var IoctlGetInt = unix.IoctlGetInt

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
var IoctlGetInt = unix.IoctlGetInt

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
var IoctlGetInt = unix.IoctlGetInt

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
var IoctlGetInt = unix.IoctlGetInt

var IoctlGetInt = unix.IoctlGetInt

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
var IoctlGetInt = unix.IoctlGetInt

var IoctlGetPtmget = unix.IoctlGetPtmget

var IoctlGetTermio = unix.IoctlGetTermio

var IoctlGetTermios = unix.IoctlGetTermios

var IoctlGetTermios = unix.IoctlGetTermios

var IoctlGetTermios = unix.IoctlGetTermios

var IoctlGetTermios = unix.IoctlGetTermios

var IoctlGetTermios = unix.IoctlGetTermios

var IoctlGetTermios = unix.IoctlGetTermios

var IoctlGetTermios = unix.IoctlGetTermios

var IoctlGetTermios = unix.IoctlGetTermios

var IoctlGetWinsize = unix.IoctlGetWinsize

var IoctlGetWinsize = unix.IoctlGetWinsize

var IoctlGetWinsize = unix.IoctlGetWinsize

var IoctlGetWinsize = unix.IoctlGetWinsize

var IoctlGetWinsize = unix.IoctlGetWinsize

var IoctlGetWinsize = unix.IoctlGetWinsize

var IoctlGetWinsize = unix.IoctlGetWinsize

var IoctlGetWinsize = unix.IoctlGetWinsize

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
var IoctlSetInt = unix.IoctlSetInt

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
var IoctlSetInt = unix.IoctlSetInt

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
var IoctlSetInt = unix.IoctlSetInt

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
var IoctlSetInt = unix.IoctlSetInt

var IoctlSetInt = unix.IoctlSetInt

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
var IoctlSetInt = unix.IoctlSetInt

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
var IoctlSetInt = unix.IoctlSetInt

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
var IoctlSetInt = unix.IoctlSetInt

// IoctlSetPointerInt performs an ioctl operation which sets an
// integer value on fd, using the specified request number. The ioctl
// argument is called with a pointer to the integer value, rather than
// passing the integer value directly.
var IoctlSetPointerInt = unix.IoctlSetPointerInt

var IoctlSetTermio = unix.IoctlSetTermio

// IoctlSetTermios performs an ioctl on fd with a *Termios.
//
// The req value will usually be TCSETA or TIOCSETA.
var IoctlSetTermios = unix.IoctlSetTermios

// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
//
// To change fd&#39;s window size, the req argument should be TIOCSWINSZ.
var IoctlSetWinsize = unix.IoctlSetWinsize

var Ioperm = unix.Ioperm

var Ioperm = unix.Ioperm

var Ioperm = unix.Ioperm

var Ioperm = unix.Ioperm

var Ioperm = unix.Ioperm

var Ioperm = unix.Ioperm

var Ioperm = unix.Ioperm

var Ioperm = unix.Ioperm

var Ioperm = unix.Ioperm

var Iopl = unix.Iopl

var Iopl = unix.Iopl

var Iopl = unix.Iopl

var Iopl = unix.Iopl

var Iopl = unix.Iopl

var Iopl = unix.Iopl

var Iopl = unix.Iopl

var Iopl = unix.Iopl

var Iopl = unix.Iopl

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Issetugid = unix.Issetugid

var Kevent = unix.Kevent

var KexecFileLoad = unix.KexecFileLoad

var KexecFileLoad = unix.KexecFileLoad

var KexecFileLoad = unix.KexecFileLoad

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

var KeyctlBuffer = unix.KeyctlBuffer

// KeyctlDHCompute implements the KEYCTL_DH_COMPUTE command. This command
// computes a Diffie-Hellman shared secret based on the provide params. The
// secret is written to the provided buffer and the returned size is the number
// of bytes written (returning an error if there is insufficient space in the
// buffer). If a nil buffer is passed in, this function returns the minimum
// buffer length needed to store the appropriate data. Note that this differs
// from KEYCTL_READ&#39;s behavior which always returns the requested payload size.
// See the full documentation at:
// http://man7.org/linux/man-pages/man3/keyctl_dh_compute.3.html
var KeyctlDHCompute = unix.KeyctlDHCompute

// KeyctlGetKeyringID implements the KEYCTL_GET_KEYRING_ID command.
// See the full documentation at:
// http://man7.org/linux/man-pages/man3/keyctl_get_keyring_ID.3.html
var KeyctlGetKeyringID = unix.KeyctlGetKeyringID

// KeyctlInstantiateIOV implements the KEYCTL_INSTANTIATE_IOV command. This
// command is similar to KEYCTL_INSTANTIATE, except that the payload is a slice
// of Iovec (each of which represents a buffer) instead of a single buffer.
// See the full documentation at:
// http://man7.org/linux/man-pages/man3/keyctl_instantiate_iov.3.html
var KeyctlInstantiateIOV = unix.KeyctlInstantiateIOV

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

var KeyctlInt = unix.KeyctlInt

// KeyctlJoinSessionKeyring implements the KEYCTL_JOIN_SESSION_KEYRING command.
// See the full documentation at:
// http://man7.org/linux/man-pages/man3/keyctl_join_session_keyring.3.html
var KeyctlJoinSessionKeyring = unix.KeyctlJoinSessionKeyring

// KeyctlSearch implements the KEYCTL_SEARCH command.
// See the full documentation at:
// http://man7.org/linux/man-pages/man3/keyctl_search.3.html
var KeyctlSearch = unix.KeyctlSearch

// KeyctlSetperm implements the KEYCTL_SETPERM command. The perm value is the
// key handle permission mask as described in the &#34;keyctl setperm&#34; section of
// http://man7.org/linux/man-pages/man1/keyctl.1.html.
// See the full documentation at:
// http://man7.org/linux/man-pages/man3/keyctl_setperm.3.html
var KeyctlSetperm = unix.KeyctlSetperm

// KeyctlString calls keyctl commands which return a string.
// These commands are KEYCTL_DESCRIBE and KEYCTL_GET_SECURITY.
var KeyctlString = unix.KeyctlString

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Kill = unix.Kill

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Klogctl = unix.Klogctl

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Kqueue = unix.Kqueue

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lchown = unix.Lchown

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Lgetxattr = unix.Lgetxattr

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Link = unix.Link

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Linkat = unix.Linkat

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listen = unix.Listen

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Listxattr = unix.Listxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Llistxattr = unix.Llistxattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lremovexattr = unix.Lremovexattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lsetxattr = unix.Lsetxattr

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Lstat = unix.Lstat

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

var Madvise = unix.Madvise

// Major returns the major component of a Darwin device number.
var Major = unix.Major

// Major returns the major component of an OpenBSD device number.
var Major = unix.Major

// Major returns the major component of a Linux device number.
var Major = unix.Major

// Major returns the major component of a NetBSD device number.
var Major = unix.Major

// Major returns the major component of a Linux device number.
var Major = unix.Major

var Major = unix.Major

// Major returns the major component of a FreeBSD device number.
var Major = unix.Major

// Major returns the major component of a DragonFlyBSD device number.
var Major = unix.Major

// Major returns the major component of a Linux device number.
var Major = unix.Major

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

var MemfdCreate = unix.MemfdCreate

// Minor returns the minor component of a Linux device number.
var Minor = unix.Minor

// Minor returns the minor component of a Linux device number.
var Minor = unix.Minor

// Minor returns the minor component of an OpenBSD device number.
var Minor = unix.Minor

// Minor returns the minor component of a Linux device number.
var Minor = unix.Minor

var Minor = unix.Minor

// Minor returns the minor component of a DragonFlyBSD device number.
var Minor = unix.Minor

// Minor returns the minor component of a NetBSD device number.
var Minor = unix.Minor

// Minor returns the minor component of a FreeBSD device number.
var Minor = unix.Minor

// Minor returns the minor component of a Darwin device number.
var Minor = unix.Minor

// Mkdev returns a FreeBSD device number generated from the given major and
// minor components.
var Mkdev = unix.Mkdev

// Mkdev returns an OpenBSD device number generated from the given major and minor
// components.
var Mkdev = unix.Mkdev

// Mkdev returns a Linux device number generated from the given major and minor
// components.
var Mkdev = unix.Mkdev

// Mkdev returns a Linux device number generated from the given major and minor
// components.
var Mkdev = unix.Mkdev

// Mkdev returns a Linux device number generated from the given major and minor
// components.
var Mkdev = unix.Mkdev

// Mkdev returns a Darwin device number generated from the given major and minor
// components.
var Mkdev = unix.Mkdev

var Mkdev = unix.Mkdev

// Mkdev returns a DragonFlyBSD device number generated from the given major and
// minor components.
var Mkdev = unix.Mkdev

// Mkdev returns a NetBSD device number generated from the given major and minor
// components.
var Mkdev = unix.Mkdev

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdir = unix.Mkdir

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkdirat = unix.Mkdirat

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifo = unix.Mkfifo

var Mkfifoat = unix.Mkfifoat

var Mkfifoat = unix.Mkfifoat

var Mkfifoat = unix.Mkfifoat

var Mkfifoat = unix.Mkfifoat

var Mkfifoat = unix.Mkfifoat

var Mkfifoat = unix.Mkfifoat

var Mkfifoat = unix.Mkfifoat

var Mkfifoat = unix.Mkfifoat

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknod = unix.Mknod

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mknodat = unix.Mknodat

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlock = unix.Mlock

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mlockall = unix.Mlockall

var Mmap = unix.Mmap

var Mmap = unix.Mmap

var Mmap = unix.Mmap

var Mmap = unix.Mmap

var Mount = unix.Mount

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Mprotect = unix.Mprotect

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Msync = unix.Msync

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlock = unix.Munlock

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munlockall = unix.Munlockall

var Munmap = unix.Munmap

var Munmap = unix.Munmap

var Munmap = unix.Munmap

var Munmap = unix.Munmap

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

var Nanosleep = unix.Nanosleep

// NsecToTimespec takes a number of nanoseconds since the Unix epoch
// and returns the corresponding Timespec value.
var NsecToTimespec = unix.NsecToTimespec

// NsecToTimeval takes a number of nanoseconds since the Unix epoch
// and returns the corresponding Timeval value.
var NsecToTimeval = unix.NsecToTimeval

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Open = unix.Open

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

// ParseDirent parses up to max directory entries in buf,
// appending the names to names. It returns the number of
// bytes consumed from buf, the number of entries added
// to names, and the new names slice.
var ParseDirent = unix.ParseDirent

// ParseSocketControlMessage parses b as an array of socket control
// messages.
var ParseSocketControlMessage = unix.ParseSocketControlMessage

// ParseUnixCredentials decodes a socket control message that contains
// credentials in a Ucred structure. To receive such a message, the
// SO_PASSCRED option must be enabled on the socket.
var ParseUnixCredentials = unix.ParseUnixCredentials

// ParseUnixRights decodes a socket control message that contains an
// integer array of open file descriptors from another process.
var ParseUnixRights = unix.ParseUnixRights

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pathconf = unix.Pathconf

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var Pause = unix.Pause

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var PerfEventOpen = unix.PerfEventOpen

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

// sysnb pipe(p *[2]_C_int) (err error)
var Pipe = unix.Pipe

// sysnb pipe() (fd1 int, fd2 int, err error)
var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe = unix.Pipe

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var Pipe2 = unix.Pipe2

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

var PivotRoot = unix.PivotRoot

// Pledge implements the pledge syscall.
//
// The pledge syscall does not accept execpromises on OpenBSD releases
// before 6.3.
//
// execpromises must be empty when Pledge is called on OpenBSD
// releases predating 6.3, otherwise an error will be returned.
//
// For more information see pledge(2).
var Pledge = unix.Pledge

// PledgeExecpromises implements the pledge syscall.
//
// This changes the execpromises and leaves the promises untouched.
//
// For more information see pledge(2).
var PledgeExecpromises = unix.PledgeExecpromises

// PledgePromises implements the pledge syscall.
//
// This changes the promises and leaves the execpromises untouched.
//
// For more information see pledge(2).
var PledgePromises = unix.PledgePromises

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Poll = unix.Poll

var Ppoll = unix.Ppoll

var Ppoll = unix.Ppoll

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Prctl = unix.Prctl

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

// sys	extpread(fd int, p []byte, flags int, offset int64) (n int, err error)
var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pread = unix.Pread

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

var Pselect = unix.Pselect

// sys   ptrace(request int, pid int, addr uintptr, data uintptr) (err error)
var PtraceAttach = unix.PtraceAttach

var PtraceAttach = unix.PtraceAttach

var PtraceCont = unix.PtraceCont

var PtraceDetach = unix.PtraceDetach

var PtraceDetach = unix.PtraceDetach

var PtraceGetEventMsg = unix.PtraceGetEventMsg

var PtraceGetRegs = unix.PtraceGetRegs

// PtraceGetRegs386 fetches the registers used by 386 binaries.
var PtraceGetRegs386 = unix.PtraceGetRegs386

// PtraceGetRegsAmd64 fetches the registers used by amd64 binaries.
var PtraceGetRegsAmd64 = unix.PtraceGetRegsAmd64

// PtraceGetRegsArm fetches the registers used by arm binaries.
var PtraceGetRegsArm = unix.PtraceGetRegsArm

// PtraceGetRegsArm64 fetches the registers used by arm64 binaries.
var PtraceGetRegsArm64 = unix.PtraceGetRegsArm64

// PtraceGetRegsMips fetches the registers used by mips binaries.
var PtraceGetRegsMips = unix.PtraceGetRegsMips

// PtraceGetRegsMips64 fetches the registers used by mips64 binaries.
var PtraceGetRegsMips64 = unix.PtraceGetRegsMips64

// PtraceGetRegsMips64le fetches the registers used by mips64le binaries.
var PtraceGetRegsMips64le = unix.PtraceGetRegsMips64le

// PtraceGetRegsMipsle fetches the registers used by mipsle binaries.
var PtraceGetRegsMipsle = unix.PtraceGetRegsMipsle

var PtracePeekData = unix.PtracePeekData

var PtracePeekText = unix.PtracePeekText

var PtracePeekUser = unix.PtracePeekUser

var PtracePokeData = unix.PtracePokeData

var PtracePokeText = unix.PtracePokeText

var PtracePokeUser = unix.PtracePokeUser

var PtraceSetOptions = unix.PtraceSetOptions

var PtraceSetRegs = unix.PtraceSetRegs

// PtraceSetRegs386 sets the registers used by 386 binaries.
var PtraceSetRegs386 = unix.PtraceSetRegs386

// PtraceSetRegsAmd64 sets the registers used by amd64 binaries.
var PtraceSetRegsAmd64 = unix.PtraceSetRegsAmd64

// PtraceSetRegsArm sets the registers used by arm binaries.
var PtraceSetRegsArm = unix.PtraceSetRegsArm

// PtraceSetRegsArm64 sets the registers used by arm64 binaries.
var PtraceSetRegsArm64 = unix.PtraceSetRegsArm64

// PtraceSetRegsMips sets the registers used by mips binaries.
var PtraceSetRegsMips = unix.PtraceSetRegsMips

// PtraceSetRegsMips64 sets the registers used by mips64 binaries.
var PtraceSetRegsMips64 = unix.PtraceSetRegsMips64

// PtraceSetRegsMips64le sets the registers used by mips64le binaries.
var PtraceSetRegsMips64le = unix.PtraceSetRegsMips64le

// PtraceSetRegsMipsle sets the registers used by mipsle binaries.
var PtraceSetRegsMipsle = unix.PtraceSetRegsMipsle

var PtraceSingleStep = unix.PtraceSingleStep

var PtraceSyscall = unix.PtraceSyscall

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

// sys	extpwrite(fd int, p []byte, flags int, offset int64) (n int, err error)
var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var Pwrite = unix.Pwrite

var RawSyscall = unix.RawSyscall

var RawSyscall = unix.RawSyscall

var RawSyscall = unix.RawSyscall

var RawSyscall6 = unix.RawSyscall6

var RawSyscall6 = unix.RawSyscall6

var RawSyscall6 = unix.RawSyscall6

var RawSyscallNoError = unix.RawSyscallNoError

// RawSyscallNoError may be used instead of RawSyscall for syscalls that don&#39;t
// fail.
var RawSyscallNoError = unix.RawSyscallNoError

var Read = unix.Read

var ReadDirent = unix.ReadDirent

var ReadDirent = unix.ReadDirent

// sys	getdirent(fd int, buf []byte) (n int, err error)
var ReadDirent = unix.ReadDirent

var ReadDirent = unix.ReadDirent

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlink = unix.Readlink

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Readlinkat = unix.Readlinkat

var Reboot = unix.Reboot

var Recvfrom = unix.Recvfrom

var Recvmsg = unix.Recvmsg

var Recvmsg = unix.Recvmsg

var Recvmsg = unix.Recvmsg

var Recvmsg = unix.Recvmsg

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Removexattr = unix.Removexattr

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Rename = unix.Rename

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat = unix.Renameat

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var Renameat2 = unix.Renameat2

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var RequestKey = unix.RequestKey

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Revoke = unix.Revoke

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

var Rmdir = unix.Rmdir

// SchedGetaffinity gets the CPU affinity mask of the thread specified by pid.
// If pid is 0 the calling thread is used.
var SchedGetaffinity = unix.SchedGetaffinity

// SchedSetaffinity sets the CPU affinity mask of the thread specified by pid.
// If pid is 0 the calling thread is used.
var SchedSetaffinity = unix.SchedSetaffinity

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Seek = unix.Seek

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Select = unix.Select

var Sendfile = unix.Sendfile

var Sendfile = unix.Sendfile

var Sendfile = unix.Sendfile

var Sendfile = unix.Sendfile

var Sendfile = unix.Sendfile

var Sendfile = unix.Sendfile

var Sendfile = unix.Sendfile

var Sendfile = unix.Sendfile

var Sendmsg = unix.Sendmsg

var Sendmsg = unix.Sendmsg

var Sendmsg = unix.Sendmsg

var Sendmsg = unix.Sendmsg

var SendmsgN = unix.SendmsgN

var SendmsgN = unix.SendmsgN

var SendmsgN = unix.SendmsgN

var SendmsgN = unix.SendmsgN

var Sendto = unix.Sendto

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetKevent = unix.SetKevent

var SetNonblock = unix.SetNonblock

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setdomainname = unix.Setdomainname

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setegid = unix.Setegid

var Setenv = unix.Setenv

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Seteuid = unix.Seteuid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsgid = unix.Setfsgid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setfsuid = unix.Setfsuid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgid = unix.Setgid

var Setgroups = unix.Setgroups

var Setgroups = unix.Setgroups

var Setgroups = unix.Setgroups

var Setgroups = unix.Setgroups

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Sethostname = unix.Sethostname

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setlogin = unix.Setlogin

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setns = unix.Setns

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpgid = unix.Setpgid

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setpriority = unix.Setpriority

var Setprivexec = unix.Setprivexec

var Setprivexec = unix.Setprivexec

var Setprivexec = unix.Setprivexec

var Setprivexec = unix.Setprivexec

var Setprivexec = unix.Setprivexec

var Setprivexec = unix.Setprivexec

var Setprivexec = unix.Setprivexec

var Setprivexec = unix.Setprivexec

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setregid = unix.Setregid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresgid = unix.Setresgid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setresuid = unix.Setresuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setreuid = unix.Setreuid

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrlimit = unix.Setrlimit

var Setrtable = unix.Setrtable

var Setrtable = unix.Setrtable

var Setrtable = unix.Setrtable

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var Setsid = unix.Setsid

var SetsockoptByte = unix.SetsockoptByte

var SetsockoptICMPv6Filter = unix.SetsockoptICMPv6Filter

var SetsockoptIPMreq = unix.SetsockoptIPMreq

var SetsockoptIPMreqn = unix.SetsockoptIPMreqn

var SetsockoptIPMreqn = unix.SetsockoptIPMreqn

var SetsockoptIPv6Mreq = unix.SetsockoptIPv6Mreq

var SetsockoptInet4Addr = unix.SetsockoptInet4Addr

var SetsockoptInt = unix.SetsockoptInt

var SetsockoptLinger = unix.SetsockoptLinger

var SetsockoptString = unix.SetsockoptString

var SetsockoptTimeval = unix.SetsockoptTimeval

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Settimeofday = unix.Settimeofday

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setuid = unix.Setuid

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Setxattr = unix.Setxattr

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

var Shutdown = unix.Shutdown

// SignalName returns the signal name for signal number s.
var SignalName = unix.SignalName

var Socket = unix.Socket

var Socketpair = unix.Socketpair

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Splice = unix.Splice

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Stat = unix.Stat

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statfs = unix.Statfs

var Statvfs = unix.Statvfs

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Statx = unix.Statx

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlink = unix.Symlink

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Symlinkat = unix.Symlinkat

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var Sync = unix.Sync

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var SyncFileRange = unix.SyncFileRange

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syncfs = unix.Syncfs

var Syscall = unix.Syscall

var Syscall = unix.Syscall

var Syscall = unix.Syscall

var Syscall6 = unix.Syscall6

var Syscall6 = unix.Syscall6

var Syscall6 = unix.Syscall6

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

var Syscall9 = unix.Syscall9

// SyscallNoError may be used instead of Syscall for syscalls that don&#39;t fail.
var SyscallNoError = unix.SyscallNoError

var SyscallNoError = unix.SyscallNoError

var Sysctl = unix.Sysctl

var SysctlArgs = unix.SysctlArgs

var SysctlClockinfo = unix.SysctlClockinfo

var SysctlRaw = unix.SysctlRaw

var SysctlUint32 = unix.SysctlUint32

var SysctlUint32Args = unix.SysctlUint32Args

var SysctlUint64 = unix.SysctlUint64

var SysctlUvmexp = unix.SysctlUvmexp

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Sysinfo = unix.Sysinfo

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tee = unix.Tee

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Tgkill = unix.Tgkill

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

var Time = unix.Time

// TimeToTimespec converts t into a Timespec.
// On some 32-bit systems the range of valid Timespec values are smaller
// than that of time.Time values.  So if t is out of the valid range of
// Timespec, it returns a zero Timespec and ERANGE.
var TimeToTimespec = unix.TimeToTimespec

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

var Times = unix.Times

// TimespecToNsec converts a Timespec value into a number of
// nanoseconds since the Unix epoch.
var TimespecToNsec = unix.TimespecToNsec

// TimevalToNsec converts a Timeval value into a number of nanoseconds
// since the Unix epoch.
var TimevalToNsec = unix.TimevalToNsec

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Truncate = unix.Truncate

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Umask = unix.Umask

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Uname = unix.Uname

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

var Undelete = unix.Undelete

// UnixCredentials encodes credentials into a socket control message
// for sending to another process. This can be used for
// authentication.
var UnixCredentials = unix.UnixCredentials

// UnixRights encodes a set of open file descriptors into a socket
// control message for sending to another process.
var UnixRights = unix.UnixRights

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlink = unix.Unlink

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unmount = unix.Unmount

var Unsetenv = unix.Unsetenv

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

var Unshare = unix.Unshare

// Unveil implements the unveil syscall.
// For more information see unveil(2).
// Note that the special case of blocking further
// unveil calls is handled by UnveilBlock.
var Unveil = unix.Unveil

// UnveilBlock blocks future unveil calls.
// For more information see unveil(2).
var UnveilBlock = unix.UnveilBlock

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Ustat = unix.Ustat

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

var Utime = unix.Utime

// sys	utimes(path string, times *[2]Timeval) (err error)
var Utimes = unix.Utimes

var Utimes = unix.Utimes

var Utimes = unix.Utimes

var Utimes = unix.Utimes

var UtimesNano = unix.UtimesNano

var UtimesNano = unix.UtimesNano

// sys	utimensat(dirfd int, path string, times *[2]Timespec, flag int) (err error)
var UtimesNano = unix.UtimesNano

var UtimesNano = unix.UtimesNano

var UtimesNanoAt = unix.UtimesNanoAt

var UtimesNanoAt = unix.UtimesNanoAt

var UtimesNanoAt = unix.UtimesNanoAt

var UtimesNanoAt = unix.UtimesNanoAt

// Vmsplice splices user pages from a slice of Iovecs into a pipe specified by fd,
// using the specified flags.
var Vmsplice = unix.Vmsplice

var Wait4 = unix.Wait4

var Wait4 = unix.Wait4

var Wait4 = unix.Wait4

// sys	wait4(pid Pid_t, status *_C_int, options int, rusage *Rusage) (wpid Pid_t, err error)
var Wait4 = unix.Wait4

var Write = unix.Write
