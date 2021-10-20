// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Override FortiAnalyzer settings.
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
// 		_, err := fortios.NewLogFortianalyzer2OverrideSetting(ctx, "trname", &fortios.LogFortianalyzer2OverrideSettingArgs{
// 			__changeIp:                pulumi.Int(0),
// 			ConnTimeout:               pulumi.Int(10),
// 			EncAlgorithm:              pulumi.String("high"),
// 			FazType:                   pulumi.Int(5),
// 			HmacAlgorithm:             pulumi.String("sha256"),
// 			IpsArchive:                pulumi.String("enable"),
// 			MonitorFailureRetryPeriod: pulumi.Int(5),
// 			MonitorKeepalivePeriod:    pulumi.Int(5),
// 			Override:                  pulumi.String("disable"),
// 			Reliable:                  pulumi.String("disable"),
// 			SslMinProtoVersion:        pulumi.String("default"),
// 			Status:                    pulumi.String("disable"),
// 			UploadInterval:            pulumi.String("daily"),
// 			UploadOption:              pulumi.String("5-minute"),
// 			UploadTime:                pulumi.String("00:59"),
// 			UseManagementVdom:         pulumi.String("disable"),
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
// LogFortianalyzer2 OverrideSetting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzer2OverrideSetting:LogFortianalyzer2OverrideSetting labelname LogFortianalyzer2OverrideSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzer2OverrideSetting struct {
	pulumi.CustomResourceState

	// Hidden attribute.
	__changeIp pulumi.IntOutput `pulumi:"__changeIp"`
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringOutput `pulumi:"accessConfig"`
	// Certificate used to communicate with FortiAnalyzer.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification pulumi.StringOutput `pulumi:"certificateVerification"`
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout pulumi.IntOutput `pulumi:"connTimeout"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringOutput `pulumi:"encAlgorithm"`
	// Hidden setting index of FortiAnalyzer.
	FazType pulumi.IntOutput `pulumi:"fazType"`
	// FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
	HmacAlgorithm pulumi.StringOutput `pulumi:"hmacAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringOutput `pulumi:"ipsArchive"`
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntOutput `pulumi:"maxLogRate"`
	// Hidden management name of FortiAnalyzer.
	MgmtName pulumi.StringOutput `pulumi:"mgmtName"`
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod pulumi.IntOutput `pulumi:"monitorFailureRetryPeriod"`
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod pulumi.IntOutput `pulumi:"monitorKeepalivePeriod"`
	// Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
	Override pulumi.StringOutput `pulumi:"override"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringOutput `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringOutput `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials LogFortianalyzer2OverrideSettingSerialArrayOutput `pulumi:"serials"`
	// The remote FortiAnalyzer.
	Server pulumi.StringOutput `pulumi:"server"`
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Day of week (month) to upload logs.
	UploadDay pulumi.StringOutput `pulumi:"uploadDay"`
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringOutput `pulumi:"uploadInterval"`
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringOutput `pulumi:"uploadOption"`
	// Time to upload logs (hh:mm).
	UploadTime pulumi.StringOutput `pulumi:"uploadTime"`
	// Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
	UseManagementVdom pulumi.StringOutput `pulumi:"useManagementVdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogFortianalyzer2OverrideSetting registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer2OverrideSetting(ctx *pulumi.Context,
	name string, args *LogFortianalyzer2OverrideSettingArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer2OverrideSetting, error) {
	if args == nil {
		args = &LogFortianalyzer2OverrideSettingArgs{}
	}

	var resource LogFortianalyzer2OverrideSetting
	err := ctx.RegisterResource("fortios:index/logFortianalyzer2OverrideSetting:LogFortianalyzer2OverrideSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzer2OverrideSetting gets an existing LogFortianalyzer2OverrideSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzer2OverrideSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzer2OverrideSettingState, opts ...pulumi.ResourceOption) (*LogFortianalyzer2OverrideSetting, error) {
	var resource LogFortianalyzer2OverrideSetting
	err := ctx.ReadResource("fortios:index/logFortianalyzer2OverrideSetting:LogFortianalyzer2OverrideSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzer2OverrideSetting resources.
type logFortianalyzer2OverrideSettingState struct {
	// Hidden attribute.
	__changeIp *int `pulumi:"__changeIp"`
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig *string `pulumi:"accessConfig"`
	// Certificate used to communicate with FortiAnalyzer.
	Certificate *string `pulumi:"certificate"`
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification *string `pulumi:"certificateVerification"`
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout *int `pulumi:"connTimeout"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Hidden setting index of FortiAnalyzer.
	FazType *int `pulumi:"fazType"`
	// FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
	HmacAlgorithm *string `pulumi:"hmacAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive *string `pulumi:"ipsArchive"`
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Hidden management name of FortiAnalyzer.
	MgmtName *string `pulumi:"mgmtName"`
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod *int `pulumi:"monitorFailureRetryPeriod"`
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod *int `pulumi:"monitorKeepalivePeriod"`
	// Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable *string `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials []LogFortianalyzer2OverrideSettingSerial `pulumi:"serials"`
	// The remote FortiAnalyzer.
	Server *string `pulumi:"server"`
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Day of week (month) to upload logs.
	UploadDay *string `pulumi:"uploadDay"`
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval *string `pulumi:"uploadInterval"`
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption *string `pulumi:"uploadOption"`
	// Time to upload logs (hh:mm).
	UploadTime *string `pulumi:"uploadTime"`
	// Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
	UseManagementVdom *string `pulumi:"useManagementVdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogFortianalyzer2OverrideSettingState struct {
	// Hidden attribute.
	__changeIp pulumi.IntPtrInput
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringPtrInput
	// Certificate used to communicate with FortiAnalyzer.
	Certificate pulumi.StringPtrInput
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification pulumi.StringPtrInput
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringPtrInput
	// Hidden setting index of FortiAnalyzer.
	FazType pulumi.IntPtrInput
	// FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
	HmacAlgorithm pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringPtrInput
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Hidden management name of FortiAnalyzer.
	MgmtName pulumi.StringPtrInput
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod pulumi.IntPtrInput
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod pulumi.IntPtrInput
	// Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringPtrInput
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials LogFortianalyzer2OverrideSettingSerialArrayInput
	// The remote FortiAnalyzer.
	Server pulumi.StringPtrInput
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Day of week (month) to upload logs.
	UploadDay pulumi.StringPtrInput
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringPtrInput
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringPtrInput
	// Time to upload logs (hh:mm).
	UploadTime pulumi.StringPtrInput
	// Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
	UseManagementVdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogFortianalyzer2OverrideSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer2OverrideSettingState)(nil)).Elem()
}

type logFortianalyzer2OverrideSettingArgs struct {
	// Hidden attribute.
	__changeIp *int `pulumi:"__changeIp"`
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig *string `pulumi:"accessConfig"`
	// Certificate used to communicate with FortiAnalyzer.
	Certificate *string `pulumi:"certificate"`
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification *string `pulumi:"certificateVerification"`
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout *int `pulumi:"connTimeout"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Hidden setting index of FortiAnalyzer.
	FazType *int `pulumi:"fazType"`
	// FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
	HmacAlgorithm *string `pulumi:"hmacAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive *string `pulumi:"ipsArchive"`
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Hidden management name of FortiAnalyzer.
	MgmtName *string `pulumi:"mgmtName"`
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod *int `pulumi:"monitorFailureRetryPeriod"`
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod *int `pulumi:"monitorKeepalivePeriod"`
	// Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable *string `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials []LogFortianalyzer2OverrideSettingSerial `pulumi:"serials"`
	// The remote FortiAnalyzer.
	Server *string `pulumi:"server"`
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Day of week (month) to upload logs.
	UploadDay *string `pulumi:"uploadDay"`
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval *string `pulumi:"uploadInterval"`
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption *string `pulumi:"uploadOption"`
	// Time to upload logs (hh:mm).
	UploadTime *string `pulumi:"uploadTime"`
	// Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
	UseManagementVdom *string `pulumi:"useManagementVdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogFortianalyzer2OverrideSetting resource.
type LogFortianalyzer2OverrideSettingArgs struct {
	// Hidden attribute.
	__changeIp pulumi.IntPtrInput
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringPtrInput
	// Certificate used to communicate with FortiAnalyzer.
	Certificate pulumi.StringPtrInput
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification pulumi.StringPtrInput
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringPtrInput
	// Hidden setting index of FortiAnalyzer.
	FazType pulumi.IntPtrInput
	// FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
	HmacAlgorithm pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringPtrInput
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Hidden management name of FortiAnalyzer.
	MgmtName pulumi.StringPtrInput
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod pulumi.IntPtrInput
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod pulumi.IntPtrInput
	// Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringPtrInput
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials LogFortianalyzer2OverrideSettingSerialArrayInput
	// The remote FortiAnalyzer.
	Server pulumi.StringPtrInput
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Day of week (month) to upload logs.
	UploadDay pulumi.StringPtrInput
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringPtrInput
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringPtrInput
	// Time to upload logs (hh:mm).
	UploadTime pulumi.StringPtrInput
	// Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
	UseManagementVdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogFortianalyzer2OverrideSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer2OverrideSettingArgs)(nil)).Elem()
}

type LogFortianalyzer2OverrideSettingInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideSettingOutput() LogFortianalyzer2OverrideSettingOutput
	ToLogFortianalyzer2OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingOutput
}

func (*LogFortianalyzer2OverrideSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortianalyzer2OverrideSetting)(nil))
}

func (i *LogFortianalyzer2OverrideSetting) ToLogFortianalyzer2OverrideSettingOutput() LogFortianalyzer2OverrideSettingOutput {
	return i.ToLogFortianalyzer2OverrideSettingOutputWithContext(context.Background())
}

func (i *LogFortianalyzer2OverrideSetting) ToLogFortianalyzer2OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideSettingOutput)
}

func (i *LogFortianalyzer2OverrideSetting) ToLogFortianalyzer2OverrideSettingPtrOutput() LogFortianalyzer2OverrideSettingPtrOutput {
	return i.ToLogFortianalyzer2OverrideSettingPtrOutputWithContext(context.Background())
}

func (i *LogFortianalyzer2OverrideSetting) ToLogFortianalyzer2OverrideSettingPtrOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideSettingPtrOutput)
}

type LogFortianalyzer2OverrideSettingPtrInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideSettingPtrOutput() LogFortianalyzer2OverrideSettingPtrOutput
	ToLogFortianalyzer2OverrideSettingPtrOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingPtrOutput
}

type logFortianalyzer2OverrideSettingPtrType LogFortianalyzer2OverrideSettingArgs

func (*logFortianalyzer2OverrideSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer2OverrideSetting)(nil))
}

func (i *logFortianalyzer2OverrideSettingPtrType) ToLogFortianalyzer2OverrideSettingPtrOutput() LogFortianalyzer2OverrideSettingPtrOutput {
	return i.ToLogFortianalyzer2OverrideSettingPtrOutputWithContext(context.Background())
}

func (i *logFortianalyzer2OverrideSettingPtrType) ToLogFortianalyzer2OverrideSettingPtrOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideSettingPtrOutput)
}

// LogFortianalyzer2OverrideSettingArrayInput is an input type that accepts LogFortianalyzer2OverrideSettingArray and LogFortianalyzer2OverrideSettingArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzer2OverrideSettingArrayInput` via:
//
//          LogFortianalyzer2OverrideSettingArray{ LogFortianalyzer2OverrideSettingArgs{...} }
type LogFortianalyzer2OverrideSettingArrayInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideSettingArrayOutput() LogFortianalyzer2OverrideSettingArrayOutput
	ToLogFortianalyzer2OverrideSettingArrayOutputWithContext(context.Context) LogFortianalyzer2OverrideSettingArrayOutput
}

type LogFortianalyzer2OverrideSettingArray []LogFortianalyzer2OverrideSettingInput

func (LogFortianalyzer2OverrideSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogFortianalyzer2OverrideSetting)(nil))
}

func (i LogFortianalyzer2OverrideSettingArray) ToLogFortianalyzer2OverrideSettingArrayOutput() LogFortianalyzer2OverrideSettingArrayOutput {
	return i.ToLogFortianalyzer2OverrideSettingArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzer2OverrideSettingArray) ToLogFortianalyzer2OverrideSettingArrayOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideSettingArrayOutput)
}

// LogFortianalyzer2OverrideSettingMapInput is an input type that accepts LogFortianalyzer2OverrideSettingMap and LogFortianalyzer2OverrideSettingMapOutput values.
// You can construct a concrete instance of `LogFortianalyzer2OverrideSettingMapInput` via:
//
//          LogFortianalyzer2OverrideSettingMap{ "key": LogFortianalyzer2OverrideSettingArgs{...} }
type LogFortianalyzer2OverrideSettingMapInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideSettingMapOutput() LogFortianalyzer2OverrideSettingMapOutput
	ToLogFortianalyzer2OverrideSettingMapOutputWithContext(context.Context) LogFortianalyzer2OverrideSettingMapOutput
}

type LogFortianalyzer2OverrideSettingMap map[string]LogFortianalyzer2OverrideSettingInput

func (LogFortianalyzer2OverrideSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogFortianalyzer2OverrideSetting)(nil))
}

func (i LogFortianalyzer2OverrideSettingMap) ToLogFortianalyzer2OverrideSettingMapOutput() LogFortianalyzer2OverrideSettingMapOutput {
	return i.ToLogFortianalyzer2OverrideSettingMapOutputWithContext(context.Background())
}

func (i LogFortianalyzer2OverrideSettingMap) ToLogFortianalyzer2OverrideSettingMapOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideSettingMapOutput)
}

type LogFortianalyzer2OverrideSettingOutput struct {
	*pulumi.OutputState
}

func (LogFortianalyzer2OverrideSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortianalyzer2OverrideSetting)(nil))
}

func (o LogFortianalyzer2OverrideSettingOutput) ToLogFortianalyzer2OverrideSettingOutput() LogFortianalyzer2OverrideSettingOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingOutput) ToLogFortianalyzer2OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingOutput) ToLogFortianalyzer2OverrideSettingPtrOutput() LogFortianalyzer2OverrideSettingPtrOutput {
	return o.ToLogFortianalyzer2OverrideSettingPtrOutputWithContext(context.Background())
}

func (o LogFortianalyzer2OverrideSettingOutput) ToLogFortianalyzer2OverrideSettingPtrOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingPtrOutput {
	return o.ApplyT(func(v LogFortianalyzer2OverrideSetting) *LogFortianalyzer2OverrideSetting {
		return &v
	}).(LogFortianalyzer2OverrideSettingPtrOutput)
}

type LogFortianalyzer2OverrideSettingPtrOutput struct {
	*pulumi.OutputState
}

func (LogFortianalyzer2OverrideSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer2OverrideSetting)(nil))
}

func (o LogFortianalyzer2OverrideSettingPtrOutput) ToLogFortianalyzer2OverrideSettingPtrOutput() LogFortianalyzer2OverrideSettingPtrOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingPtrOutput) ToLogFortianalyzer2OverrideSettingPtrOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingPtrOutput {
	return o
}

type LogFortianalyzer2OverrideSettingArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer2OverrideSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogFortianalyzer2OverrideSetting)(nil))
}

func (o LogFortianalyzer2OverrideSettingArrayOutput) ToLogFortianalyzer2OverrideSettingArrayOutput() LogFortianalyzer2OverrideSettingArrayOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingArrayOutput) ToLogFortianalyzer2OverrideSettingArrayOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingArrayOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingArrayOutput) Index(i pulumi.IntInput) LogFortianalyzer2OverrideSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogFortianalyzer2OverrideSetting {
		return vs[0].([]LogFortianalyzer2OverrideSetting)[vs[1].(int)]
	}).(LogFortianalyzer2OverrideSettingOutput)
}

type LogFortianalyzer2OverrideSettingMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer2OverrideSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogFortianalyzer2OverrideSetting)(nil))
}

func (o LogFortianalyzer2OverrideSettingMapOutput) ToLogFortianalyzer2OverrideSettingMapOutput() LogFortianalyzer2OverrideSettingMapOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingMapOutput) ToLogFortianalyzer2OverrideSettingMapOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingMapOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzer2OverrideSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogFortianalyzer2OverrideSetting {
		return vs[0].(map[string]LogFortianalyzer2OverrideSetting)[vs[1].(string)]
	}).(LogFortianalyzer2OverrideSettingOutput)
}

func init() {
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideSettingOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideSettingPtrOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideSettingArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideSettingMapOutput{})
}