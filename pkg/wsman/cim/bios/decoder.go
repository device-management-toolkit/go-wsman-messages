/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package bios

const (
	CIMBIOSElement string = "CIM_BIOSElement"
	CIMBIOSFeature string = "CIM_BIOSFeature"
	ValueNotFound  string = "Value not found in map"
)

const (
	Unknown TargetOperatingSystem = iota
	Other
	MACOS
	ATTUNIX
	DGUX
	DECNT
	Tru64UNIX
	OpenVMS
	HPUX
	AIX
	MVS
	OS400
	OS2
	JavaVM
	MSDOS
	WIN3x
	WIN95
	WIN98
	WINNT
	WINCE
	NCR3000
	NetWare
	OSF
	DCOS
	ReliantUNIX
	SCOUnixWare
	SCOOpenServer
	Sequent
	IRIX
	Solaris
	SunOS
	U6000
	ASERIES
	HPNonStopOS
	HPNonStopOSS
	BS2000
	LINUX
	Lynx
	XENIX
	VM
	InteractiveUNIX
	BSDUNIX
	FreeBSD
	NetBSD
	GNUHurd
	OS9
	MACHKernel
	Inferno
	QNX
	EPOC
	IxWorks
	VxWorks
	MiNT
	BeOS
	HPMPE
	NextStep
	PalmPilot
	Rhapsody
	Windows2000
	Dedicated
	OS390
	VSE
	TPF
	WindowsMe
	CalderaOpenUNIX
	OpenBSD
	NotApplicable
	WindowsXP
	ZOS
	MicrosoftWindowsServer2003
	MicrosoftWindowsServer200364Bit
	WindowsXP64Bit
	WindowsXPEmbedded
	WindowsVista
	WindowsVista64Bit
	WindowsEmbeddedForPointOfService
	MicrosoftWindowsServer2008
	MicrosoftWindowsServer200864Bit
	FreeBSD64Bit
	RedHatEnterpriseLinux
	RedHatEnterpriseLinux64Bit
	Solaris64Bit
	SUSE
	SUSE64Bit
	SLES
	SLES64Bit
	NovellOES
	NovellLinuxDesktop
	SunJavaDesktopSystem
	Mandriva
	Mandriva64Bit
	TurboLinux
	TurboLinux64Bit
	Ubuntu
	Ubuntu64Bit
	Debian
	Debian64Bit
	Linux24x
	Linux24x64Bit
	Linux26x
	Linux26x64Bit
	Linux64Bit
	Other64Bit
	MicrosoftWindowsServer2008R2
	VMwareESXi
	MicrosoftWindows7
	CentOS32bit
	CentOS64bit
	OracleEnterpriseLinux32bit
	OracleEnterpriseLinux64bit
	EComStation32BitX
	MicrosoftWindowsServer2011
	MicrosoftWindowsServer201164Bit
	MicrosoftWindowsServer8
)

// targetOperatingSystemToString is a mapping of the TargetOperatingSystem value to a string.
var targetOperatingSystemToString = map[TargetOperatingSystem]string{
	Unknown:                          "Unknown",
	Other:                            "Other",
	MACOS:                            "MACOS",
	ATTUNIX:                          "ATTUNIX",
	DGUX:                             "DGUX",
	DECNT:                            "DECNT",
	Tru64UNIX:                        "Tru64UNIX",
	OpenVMS:                          "OpenVMS",
	HPUX:                             "HPUX",
	AIX:                              "AIX",
	MVS:                              "MVS",
	OS400:                            "OS400",
	OS2:                              "OS2",
	JavaVM:                           "JavaVM",
	MSDOS:                            "MSDOS",
	WIN3x:                            "WIN3x",
	WIN95:                            "WIN95",
	WIN98:                            "WIN98",
	WINNT:                            "WINNT",
	WINCE:                            "WINCE",
	NCR3000:                          "NCR3000",
	NetWare:                          "NetWare",
	OSF:                              "OSF",
	DCOS:                             "DCOS",
	ReliantUNIX:                      "ReliantUNIX",
	SCOUnixWare:                      "SCOUnixWare",
	SCOOpenServer:                    "SCOOpenServer",
	Sequent:                          "Sequent",
	IRIX:                             "IRIX",
	Solaris:                          "Solaris",
	SunOS:                            "SunOS",
	U6000:                            "U6000",
	ASERIES:                          "ASERIES",
	HPNonStopOS:                      "HPNonStopOS",
	HPNonStopOSS:                     "HPNonStopOSS",
	BS2000:                           "BS2000",
	LINUX:                            "LINUX",
	Lynx:                             "Lynx",
	XENIX:                            "XENIX",
	VM:                               "VM",
	InteractiveUNIX:                  "InteractiveUNIX",
	BSDUNIX:                          "BSDUNIX",
	FreeBSD:                          "FreeBSD",
	NetBSD:                           "NetBSD",
	GNUHurd:                          "GNUHurd",
	OS9:                              "OS9",
	MACHKernel:                       "MACHKernel",
	Inferno:                          "Inferno",
	QNX:                              "QNX",
	EPOC:                             "EPOC",
	IxWorks:                          "IxWorks",
	VxWorks:                          "VxWorks",
	MiNT:                             "MiNT",
	BeOS:                             "BeOS",
	HPMPE:                            "HPMPE",
	NextStep:                         "NextStep",
	PalmPilot:                        "PalmPilot",
	Rhapsody:                         "Rhapsody",
	Windows2000:                      "Windows2000",
	Dedicated:                        "Dedicated",
	OS390:                            "OS390",
	VSE:                              "VSE",
	TPF:                              "TPF",
	WindowsMe:                        "WindowsMe",
	CalderaOpenUNIX:                  "CalderaOpenUNIX",
	OpenBSD:                          "OpenBSD",
	NotApplicable:                    "NotApplicable",
	WindowsXP:                        "WindowsXP",
	ZOS:                              "ZOS",
	MicrosoftWindowsServer2003:       "MicrosoftWindowsServer2003",
	MicrosoftWindowsServer200364Bit:  "MicrosoftWindowsServer200364Bit",
	WindowsXP64Bit:                   "WindowsXP64Bit",
	WindowsXPEmbedded:                "WindowsXPEmbedded",
	WindowsVista:                     "WindowsVista",
	WindowsVista64Bit:                "WindowsVista64Bit",
	WindowsEmbeddedForPointOfService: "WindowsEmbeddedForPointOfService",
	MicrosoftWindowsServer2008:       "MicrosoftWindowsServer2008",
	MicrosoftWindowsServer200864Bit:  "MicrosoftWindowsServer200864Bit",
	FreeBSD64Bit:                     "FreeBSD64Bit",
	RedHatEnterpriseLinux:            "RedHatEnterpriseLinux",
	RedHatEnterpriseLinux64Bit:       "RedHatEnterpriseLinux64Bit",
	Solaris64Bit:                     "Solaris64Bit",
	SUSE:                             "SUSE",
	SUSE64Bit:                        "SUSE64Bit",
	SLES:                             "SLES",
	SLES64Bit:                        "SLES64Bit",
	NovellOES:                        "NovellOES",
	NovellLinuxDesktop:               "NovellLinuxDesktop",
	SunJavaDesktopSystem:             "SunJavaDesktopSystem",
	Mandriva:                         "Mandriva",
	Mandriva64Bit:                    "Mandriva64Bit",
	TurboLinux:                       "TurboLinux",
	TurboLinux64Bit:                  "TurboLinux64Bit",
	Ubuntu:                           "Ubuntu",
	Ubuntu64Bit:                      "Ubuntu64Bit",
	Debian:                           "Debian",
	Debian64Bit:                      "Debian64Bit",
	Linux24x:                         "Linux24x",
	Linux24x64Bit:                    "Linux24x64Bit",
	Linux26x:                         "Linux26x",
	Linux26x64Bit:                    "Linux26x64Bit",
	Linux64Bit:                       "Linux64Bit",
	Other64Bit:                       "Other64Bit",
	MicrosoftWindowsServer2008R2:     "MicrosoftWindowsServer2008R2",
	VMwareESXi:                       "VMwareESXi",
	MicrosoftWindows7:                "MicrosoftWindows7",
	CentOS32bit:                      "CentOS32bit",
	CentOS64bit:                      "CentOS64bit",
	OracleEnterpriseLinux32bit:       "OracleEnterpriseLinux32bit",
	OracleEnterpriseLinux64bit:       "OracleEnterpriseLinux64bit",
	EComStation32BitX:                "EComStation32BitX",
	MicrosoftWindowsServer2011:       "MicrosoftWindowsServer2011",
	MicrosoftWindowsServer201164Bit:  "MicrosoftWindowsServer201164Bit",
	MicrosoftWindowsServer8:          "MicrosoftWindowsServer8",
}

// String returns the string representation of the TargetOperatingSystem value.
func (t TargetOperatingSystem) String() string {
	if value, exists := targetOperatingSystemToString[t]; exists {
		return value
	}

	return ValueNotFound
}

const (
	SoftwareElementStateDeployable SoftwareElementState = iota
	SoftwareElementStateInstallable
	SoftwareElementStateExecutable
	SoftwareElementStateRunning
)

// SoftwareElementStateToString is a mapping of the SoftwareElementState value to a string.
var SoftwareElementStateToString = map[SoftwareElementState]string{
	SoftwareElementStateDeployable:  "Deployable",
	SoftwareElementStateInstallable: "Installable",
	SoftwareElementStateExecutable:  "Executable",
	SoftwareElementStateRunning:     "Running",
}

// String returns the string representation of the SoftwareElementState value.
func (s SoftwareElementState) String() string {
	if value, exists := SoftwareElementStateToString[s]; exists {
		return value
	}

	return ValueNotFound
}

const (
	OperationalStatusUnknown OperationalStatus = iota
	OperationalStatusOther
	OperationalStatusOK
	OperationalStatusDegraded
	OperationalStatusStressed
	OperationalStatusPredictiveFailure
	OperationalStatusError
	OperationalStatusNonRecoverableError
	OperationalStatusStarting
	OperationalStatusStopping
	OperationalStatusStopped
	OperationalStatusInService
	OperationalStatusNoContact
	OperationalStatusLostCommunication
	OperationalStatusAborted
	OperationalStatusDormant
	OperationalStatusSupportingEntityInError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
)

// operationalStatusToString is a mapping of the OperationalStatus value to a string.
var operationalStatusToString = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "PredictiveFailure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "NonRecoverableError",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "InService",
	OperationalStatusNoContact:               "NoContact",
	OperationalStatusLostCommunication:       "LostCommunication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "SupportingEntityInError",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "PowerMode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns the string representation of the OperationalStatus value.
func (o OperationalStatus) String() string {
	if value, exists := operationalStatusToString[o]; exists {
		return value
	}

	return ValueNotFound
}
