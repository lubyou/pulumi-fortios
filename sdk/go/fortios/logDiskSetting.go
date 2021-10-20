// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for local disk logging.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewLogDiskSetting(ctx, "trname", &fortios.LogDiskSettingArgs{
// 			Diskfull:                   pulumi.String("overwrite"),
// 			DlpArchiveQuota:            pulumi.Int(0),
// 			FullFinalWarningThreshold:  pulumi.Int(95),
// 			FullFirstWarningThreshold:  pulumi.Int(75),
// 			FullSecondWarningThreshold: pulumi.Int(90),
// 			IpsArchive:                 pulumi.String("enable"),
// 			LogQuota:                   pulumi.Int(0),
// 			MaxLogFileSize:             pulumi.Int(20),
// 			MaxPolicyPacketCaptureSize: pulumi.Int(100),
// 			MaximumLogAge:              pulumi.Int(7),
// 			ReportQuota:                pulumi.Int(0),
// 			RollDay:                    pulumi.String("sunday"),
// 			RollSchedule:               pulumi.String("daily"),
// 			RollTime:                   pulumi.String("00:00"),
// 			SourceIp:                   pulumi.String("0.0.0.0"),
// 			Status:                     pulumi.String("enable"),
// 			Upload:                     pulumi.String("disable"),
// 			UploadDeleteFiles:          pulumi.String("enable"),
// 			UploadDestination:          pulumi.String("ftp-server"),
// 			UploadSslConn:              pulumi.String("default"),
// 			Uploadip:                   pulumi.String("0.0.0.0"),
// 			Uploadport:                 pulumi.Int(21),
// 			Uploadsched:                pulumi.String("disable"),
// 			Uploadtime:                 pulumi.String("00:00"),
// 			Uploadtype:                 pulumi.String("traffic event virus webfilter IPS spamfilter dlp-archive anomaly voip dlp app-ctrl waf netscan gtp dns"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// LogDisk Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logDiskSetting:LogDiskSetting labelname LogDiskSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogDiskSetting struct {
	pulumi.CustomResourceState

	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull pulumi.StringOutput `pulumi:"diskfull"`
	// DLP archive quota (MB).
	DlpArchiveQuota pulumi.IntOutput `pulumi:"dlpArchiveQuota"`
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold pulumi.IntOutput `pulumi:"fullFinalWarningThreshold"`
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold pulumi.IntOutput `pulumi:"fullFirstWarningThreshold"`
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold pulumi.IntOutput `pulumi:"fullSecondWarningThreshold"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringOutput `pulumi:"ipsArchive"`
	// Disk log quota (MB).
	LogQuota pulumi.IntOutput `pulumi:"logQuota"`
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize pulumi.IntOutput `pulumi:"maxLogFileSize"`
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize pulumi.IntOutput `pulumi:"maxPolicyPacketCaptureSize"`
	// Delete log files older than (days).
	MaximumLogAge pulumi.IntOutput `pulumi:"maximumLogAge"`
	// Report quota (MB).
	ReportQuota pulumi.IntOutput `pulumi:"reportQuota"`
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay pulumi.StringOutput `pulumi:"rollDay"`
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule pulumi.StringOutput `pulumi:"rollSchedule"`
	// Time of day to roll the log file (hh:mm).
	RollTime pulumi.StringOutput `pulumi:"rollTime"`
	// Source IP address to use for uploading disk log files.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload pulumi.StringOutput `pulumi:"upload"`
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles pulumi.StringOutput `pulumi:"uploadDeleteFiles"`
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination pulumi.StringOutput `pulumi:"uploadDestination"`
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn pulumi.StringOutput `pulumi:"uploadSslConn"`
	// The remote directory on the FTP server to upload log files to.
	Uploaddir pulumi.StringOutput `pulumi:"uploaddir"`
	// IP address of the FTP server to upload log files to.
	Uploadip pulumi.StringOutput `pulumi:"uploadip"`
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass pulumi.StringPtrOutput `pulumi:"uploadpass"`
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport pulumi.IntOutput `pulumi:"uploadport"`
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched pulumi.StringOutput `pulumi:"uploadsched"`
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime pulumi.StringOutput `pulumi:"uploadtime"`
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype pulumi.StringOutput `pulumi:"uploadtype"`
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser pulumi.StringOutput `pulumi:"uploaduser"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogDiskSetting registers a new resource with the given unique name, arguments, and options.
func NewLogDiskSetting(ctx *pulumi.Context,
	name string, args *LogDiskSettingArgs, opts ...pulumi.ResourceOption) (*LogDiskSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource LogDiskSetting
	err := ctx.RegisterResource("fortios:index/logDiskSetting:LogDiskSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogDiskSetting gets an existing LogDiskSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogDiskSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogDiskSettingState, opts ...pulumi.ResourceOption) (*LogDiskSetting, error) {
	var resource LogDiskSetting
	err := ctx.ReadResource("fortios:index/logDiskSetting:LogDiskSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogDiskSetting resources.
type logDiskSettingState struct {
	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull *string `pulumi:"diskfull"`
	// DLP archive quota (MB).
	DlpArchiveQuota *int `pulumi:"dlpArchiveQuota"`
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold *int `pulumi:"fullFinalWarningThreshold"`
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold *int `pulumi:"fullFirstWarningThreshold"`
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold *int `pulumi:"fullSecondWarningThreshold"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive *string `pulumi:"ipsArchive"`
	// Disk log quota (MB).
	LogQuota *int `pulumi:"logQuota"`
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize *int `pulumi:"maxLogFileSize"`
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize *int `pulumi:"maxPolicyPacketCaptureSize"`
	// Delete log files older than (days).
	MaximumLogAge *int `pulumi:"maximumLogAge"`
	// Report quota (MB).
	ReportQuota *int `pulumi:"reportQuota"`
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay *string `pulumi:"rollDay"`
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule *string `pulumi:"rollSchedule"`
	// Time of day to roll the log file (hh:mm).
	RollTime *string `pulumi:"rollTime"`
	// Source IP address to use for uploading disk log files.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload *string `pulumi:"upload"`
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles *string `pulumi:"uploadDeleteFiles"`
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination *string `pulumi:"uploadDestination"`
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn *string `pulumi:"uploadSslConn"`
	// The remote directory on the FTP server to upload log files to.
	Uploaddir *string `pulumi:"uploaddir"`
	// IP address of the FTP server to upload log files to.
	Uploadip *string `pulumi:"uploadip"`
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass *string `pulumi:"uploadpass"`
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport *int `pulumi:"uploadport"`
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched *string `pulumi:"uploadsched"`
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime *string `pulumi:"uploadtime"`
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype *string `pulumi:"uploadtype"`
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser *string `pulumi:"uploaduser"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogDiskSettingState struct {
	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull pulumi.StringPtrInput
	// DLP archive quota (MB).
	DlpArchiveQuota pulumi.IntPtrInput
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold pulumi.IntPtrInput
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold pulumi.IntPtrInput
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringPtrInput
	// Disk log quota (MB).
	LogQuota pulumi.IntPtrInput
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize pulumi.IntPtrInput
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize pulumi.IntPtrInput
	// Delete log files older than (days).
	MaximumLogAge pulumi.IntPtrInput
	// Report quota (MB).
	ReportQuota pulumi.IntPtrInput
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay pulumi.StringPtrInput
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule pulumi.StringPtrInput
	// Time of day to roll the log file (hh:mm).
	RollTime pulumi.StringPtrInput
	// Source IP address to use for uploading disk log files.
	SourceIp pulumi.StringPtrInput
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload pulumi.StringPtrInput
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles pulumi.StringPtrInput
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination pulumi.StringPtrInput
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn pulumi.StringPtrInput
	// The remote directory on the FTP server to upload log files to.
	Uploaddir pulumi.StringPtrInput
	// IP address of the FTP server to upload log files to.
	Uploadip pulumi.StringPtrInput
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass pulumi.StringPtrInput
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport pulumi.IntPtrInput
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched pulumi.StringPtrInput
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime pulumi.StringPtrInput
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype pulumi.StringPtrInput
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogDiskSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logDiskSettingState)(nil)).Elem()
}

type logDiskSettingArgs struct {
	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull *string `pulumi:"diskfull"`
	// DLP archive quota (MB).
	DlpArchiveQuota *int `pulumi:"dlpArchiveQuota"`
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold *int `pulumi:"fullFinalWarningThreshold"`
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold *int `pulumi:"fullFirstWarningThreshold"`
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold *int `pulumi:"fullSecondWarningThreshold"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive *string `pulumi:"ipsArchive"`
	// Disk log quota (MB).
	LogQuota *int `pulumi:"logQuota"`
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize *int `pulumi:"maxLogFileSize"`
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize *int `pulumi:"maxPolicyPacketCaptureSize"`
	// Delete log files older than (days).
	MaximumLogAge *int `pulumi:"maximumLogAge"`
	// Report quota (MB).
	ReportQuota *int `pulumi:"reportQuota"`
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay *string `pulumi:"rollDay"`
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule *string `pulumi:"rollSchedule"`
	// Time of day to roll the log file (hh:mm).
	RollTime *string `pulumi:"rollTime"`
	// Source IP address to use for uploading disk log files.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload *string `pulumi:"upload"`
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles *string `pulumi:"uploadDeleteFiles"`
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination *string `pulumi:"uploadDestination"`
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn *string `pulumi:"uploadSslConn"`
	// The remote directory on the FTP server to upload log files to.
	Uploaddir *string `pulumi:"uploaddir"`
	// IP address of the FTP server to upload log files to.
	Uploadip *string `pulumi:"uploadip"`
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass *string `pulumi:"uploadpass"`
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport *int `pulumi:"uploadport"`
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched *string `pulumi:"uploadsched"`
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime *string `pulumi:"uploadtime"`
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype *string `pulumi:"uploadtype"`
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser *string `pulumi:"uploaduser"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogDiskSetting resource.
type LogDiskSettingArgs struct {
	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull pulumi.StringPtrInput
	// DLP archive quota (MB).
	DlpArchiveQuota pulumi.IntPtrInput
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold pulumi.IntPtrInput
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold pulumi.IntPtrInput
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringPtrInput
	// Disk log quota (MB).
	LogQuota pulumi.IntPtrInput
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize pulumi.IntPtrInput
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize pulumi.IntPtrInput
	// Delete log files older than (days).
	MaximumLogAge pulumi.IntPtrInput
	// Report quota (MB).
	ReportQuota pulumi.IntPtrInput
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay pulumi.StringPtrInput
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule pulumi.StringPtrInput
	// Time of day to roll the log file (hh:mm).
	RollTime pulumi.StringPtrInput
	// Source IP address to use for uploading disk log files.
	SourceIp pulumi.StringPtrInput
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload pulumi.StringPtrInput
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles pulumi.StringPtrInput
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination pulumi.StringPtrInput
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn pulumi.StringPtrInput
	// The remote directory on the FTP server to upload log files to.
	Uploaddir pulumi.StringPtrInput
	// IP address of the FTP server to upload log files to.
	Uploadip pulumi.StringPtrInput
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass pulumi.StringPtrInput
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport pulumi.IntPtrInput
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched pulumi.StringPtrInput
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime pulumi.StringPtrInput
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype pulumi.StringPtrInput
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogDiskSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logDiskSettingArgs)(nil)).Elem()
}

type LogDiskSettingInput interface {
	pulumi.Input

	ToLogDiskSettingOutput() LogDiskSettingOutput
	ToLogDiskSettingOutputWithContext(ctx context.Context) LogDiskSettingOutput
}

func (*LogDiskSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*LogDiskSetting)(nil))
}

func (i *LogDiskSetting) ToLogDiskSettingOutput() LogDiskSettingOutput {
	return i.ToLogDiskSettingOutputWithContext(context.Background())
}

func (i *LogDiskSetting) ToLogDiskSettingOutputWithContext(ctx context.Context) LogDiskSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDiskSettingOutput)
}

func (i *LogDiskSetting) ToLogDiskSettingPtrOutput() LogDiskSettingPtrOutput {
	return i.ToLogDiskSettingPtrOutputWithContext(context.Background())
}

func (i *LogDiskSetting) ToLogDiskSettingPtrOutputWithContext(ctx context.Context) LogDiskSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDiskSettingPtrOutput)
}

type LogDiskSettingPtrInput interface {
	pulumi.Input

	ToLogDiskSettingPtrOutput() LogDiskSettingPtrOutput
	ToLogDiskSettingPtrOutputWithContext(ctx context.Context) LogDiskSettingPtrOutput
}

type logDiskSettingPtrType LogDiskSettingArgs

func (*logDiskSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDiskSetting)(nil))
}

func (i *logDiskSettingPtrType) ToLogDiskSettingPtrOutput() LogDiskSettingPtrOutput {
	return i.ToLogDiskSettingPtrOutputWithContext(context.Background())
}

func (i *logDiskSettingPtrType) ToLogDiskSettingPtrOutputWithContext(ctx context.Context) LogDiskSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDiskSettingPtrOutput)
}

// LogDiskSettingArrayInput is an input type that accepts LogDiskSettingArray and LogDiskSettingArrayOutput values.
// You can construct a concrete instance of `LogDiskSettingArrayInput` via:
//
//          LogDiskSettingArray{ LogDiskSettingArgs{...} }
type LogDiskSettingArrayInput interface {
	pulumi.Input

	ToLogDiskSettingArrayOutput() LogDiskSettingArrayOutput
	ToLogDiskSettingArrayOutputWithContext(context.Context) LogDiskSettingArrayOutput
}

type LogDiskSettingArray []LogDiskSettingInput

func (LogDiskSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogDiskSetting)(nil))
}

func (i LogDiskSettingArray) ToLogDiskSettingArrayOutput() LogDiskSettingArrayOutput {
	return i.ToLogDiskSettingArrayOutputWithContext(context.Background())
}

func (i LogDiskSettingArray) ToLogDiskSettingArrayOutputWithContext(ctx context.Context) LogDiskSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDiskSettingArrayOutput)
}

// LogDiskSettingMapInput is an input type that accepts LogDiskSettingMap and LogDiskSettingMapOutput values.
// You can construct a concrete instance of `LogDiskSettingMapInput` via:
//
//          LogDiskSettingMap{ "key": LogDiskSettingArgs{...} }
type LogDiskSettingMapInput interface {
	pulumi.Input

	ToLogDiskSettingMapOutput() LogDiskSettingMapOutput
	ToLogDiskSettingMapOutputWithContext(context.Context) LogDiskSettingMapOutput
}

type LogDiskSettingMap map[string]LogDiskSettingInput

func (LogDiskSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogDiskSetting)(nil))
}

func (i LogDiskSettingMap) ToLogDiskSettingMapOutput() LogDiskSettingMapOutput {
	return i.ToLogDiskSettingMapOutputWithContext(context.Background())
}

func (i LogDiskSettingMap) ToLogDiskSettingMapOutputWithContext(ctx context.Context) LogDiskSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDiskSettingMapOutput)
}

type LogDiskSettingOutput struct {
	*pulumi.OutputState
}

func (LogDiskSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogDiskSetting)(nil))
}

func (o LogDiskSettingOutput) ToLogDiskSettingOutput() LogDiskSettingOutput {
	return o
}

func (o LogDiskSettingOutput) ToLogDiskSettingOutputWithContext(ctx context.Context) LogDiskSettingOutput {
	return o
}

func (o LogDiskSettingOutput) ToLogDiskSettingPtrOutput() LogDiskSettingPtrOutput {
	return o.ToLogDiskSettingPtrOutputWithContext(context.Background())
}

func (o LogDiskSettingOutput) ToLogDiskSettingPtrOutputWithContext(ctx context.Context) LogDiskSettingPtrOutput {
	return o.ApplyT(func(v LogDiskSetting) *LogDiskSetting {
		return &v
	}).(LogDiskSettingPtrOutput)
}

type LogDiskSettingPtrOutput struct {
	*pulumi.OutputState
}

func (LogDiskSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDiskSetting)(nil))
}

func (o LogDiskSettingPtrOutput) ToLogDiskSettingPtrOutput() LogDiskSettingPtrOutput {
	return o
}

func (o LogDiskSettingPtrOutput) ToLogDiskSettingPtrOutputWithContext(ctx context.Context) LogDiskSettingPtrOutput {
	return o
}

type LogDiskSettingArrayOutput struct{ *pulumi.OutputState }

func (LogDiskSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogDiskSetting)(nil))
}

func (o LogDiskSettingArrayOutput) ToLogDiskSettingArrayOutput() LogDiskSettingArrayOutput {
	return o
}

func (o LogDiskSettingArrayOutput) ToLogDiskSettingArrayOutputWithContext(ctx context.Context) LogDiskSettingArrayOutput {
	return o
}

func (o LogDiskSettingArrayOutput) Index(i pulumi.IntInput) LogDiskSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogDiskSetting {
		return vs[0].([]LogDiskSetting)[vs[1].(int)]
	}).(LogDiskSettingOutput)
}

type LogDiskSettingMapOutput struct{ *pulumi.OutputState }

func (LogDiskSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogDiskSetting)(nil))
}

func (o LogDiskSettingMapOutput) ToLogDiskSettingMapOutput() LogDiskSettingMapOutput {
	return o
}

func (o LogDiskSettingMapOutput) ToLogDiskSettingMapOutputWithContext(ctx context.Context) LogDiskSettingMapOutput {
	return o
}

func (o LogDiskSettingMapOutput) MapIndex(k pulumi.StringInput) LogDiskSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogDiskSetting {
		return vs[0].(map[string]LogDiskSetting)[vs[1].(string)]
	}).(LogDiskSettingOutput)
}

func init() {
	pulumi.RegisterOutputType(LogDiskSettingOutput{})
	pulumi.RegisterOutputType(LogDiskSettingPtrOutput{})
	pulumi.RegisterOutputType(LogDiskSettingArrayOutput{})
	pulumi.RegisterOutputType(LogDiskSettingMapOutput{})
}