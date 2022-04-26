// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

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
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewLogFortianalyzer3OverrideSetting(ctx, "trname", &fortios.LogFortianalyzer3OverrideSettingArgs{
// 			__changeIp:                pulumi.Int(0),
// 			ConnTimeout:               pulumi.Int(10),
// 			EncAlgorithm:              pulumi.String("high"),
// 			FazType:                   pulumi.Int(6),
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
// LogFortianalyzer3 OverrideSetting can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzer3OverrideSetting:LogFortianalyzer3OverrideSetting labelname LogFortianalyzer3OverrideSetting
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzer3OverrideSetting:LogFortianalyzer3OverrideSetting labelname LogFortianalyzer3OverrideSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzer3OverrideSetting struct {
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
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey pulumi.StringOutput `pulumi:"presharedKey"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringOutput `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringOutput `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials LogFortianalyzer3OverrideSettingSerialArrayOutput `pulumi:"serials"`
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

// NewLogFortianalyzer3OverrideSetting registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer3OverrideSetting(ctx *pulumi.Context,
	name string, args *LogFortianalyzer3OverrideSettingArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer3OverrideSetting, error) {
	if args == nil {
		args = &LogFortianalyzer3OverrideSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogFortianalyzer3OverrideSetting
	err := ctx.RegisterResource("fortios:index/logFortianalyzer3OverrideSetting:LogFortianalyzer3OverrideSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzer3OverrideSetting gets an existing LogFortianalyzer3OverrideSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzer3OverrideSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzer3OverrideSettingState, opts ...pulumi.ResourceOption) (*LogFortianalyzer3OverrideSetting, error) {
	var resource LogFortianalyzer3OverrideSetting
	err := ctx.ReadResource("fortios:index/logFortianalyzer3OverrideSetting:LogFortianalyzer3OverrideSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzer3OverrideSetting resources.
type logFortianalyzer3OverrideSettingState struct {
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
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey *string `pulumi:"presharedKey"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable *string `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials []LogFortianalyzer3OverrideSettingSerial `pulumi:"serials"`
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

type LogFortianalyzer3OverrideSettingState struct {
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
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey pulumi.StringPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringPtrInput
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials LogFortianalyzer3OverrideSettingSerialArrayInput
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

func (LogFortianalyzer3OverrideSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3OverrideSettingState)(nil)).Elem()
}

type logFortianalyzer3OverrideSettingArgs struct {
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
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey *string `pulumi:"presharedKey"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable *string `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials []LogFortianalyzer3OverrideSettingSerial `pulumi:"serials"`
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

// The set of arguments for constructing a LogFortianalyzer3OverrideSetting resource.
type LogFortianalyzer3OverrideSettingArgs struct {
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
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey pulumi.StringPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringPtrInput
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials LogFortianalyzer3OverrideSettingSerialArrayInput
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

func (LogFortianalyzer3OverrideSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3OverrideSettingArgs)(nil)).Elem()
}

type LogFortianalyzer3OverrideSettingInput interface {
	pulumi.Input

	ToLogFortianalyzer3OverrideSettingOutput() LogFortianalyzer3OverrideSettingOutput
	ToLogFortianalyzer3OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideSettingOutput
}

func (*LogFortianalyzer3OverrideSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3OverrideSetting)(nil)).Elem()
}

func (i *LogFortianalyzer3OverrideSetting) ToLogFortianalyzer3OverrideSettingOutput() LogFortianalyzer3OverrideSettingOutput {
	return i.ToLogFortianalyzer3OverrideSettingOutputWithContext(context.Background())
}

func (i *LogFortianalyzer3OverrideSetting) ToLogFortianalyzer3OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3OverrideSettingOutput)
}

// LogFortianalyzer3OverrideSettingArrayInput is an input type that accepts LogFortianalyzer3OverrideSettingArray and LogFortianalyzer3OverrideSettingArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzer3OverrideSettingArrayInput` via:
//
//          LogFortianalyzer3OverrideSettingArray{ LogFortianalyzer3OverrideSettingArgs{...} }
type LogFortianalyzer3OverrideSettingArrayInput interface {
	pulumi.Input

	ToLogFortianalyzer3OverrideSettingArrayOutput() LogFortianalyzer3OverrideSettingArrayOutput
	ToLogFortianalyzer3OverrideSettingArrayOutputWithContext(context.Context) LogFortianalyzer3OverrideSettingArrayOutput
}

type LogFortianalyzer3OverrideSettingArray []LogFortianalyzer3OverrideSettingInput

func (LogFortianalyzer3OverrideSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer3OverrideSetting)(nil)).Elem()
}

func (i LogFortianalyzer3OverrideSettingArray) ToLogFortianalyzer3OverrideSettingArrayOutput() LogFortianalyzer3OverrideSettingArrayOutput {
	return i.ToLogFortianalyzer3OverrideSettingArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzer3OverrideSettingArray) ToLogFortianalyzer3OverrideSettingArrayOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3OverrideSettingArrayOutput)
}

// LogFortianalyzer3OverrideSettingMapInput is an input type that accepts LogFortianalyzer3OverrideSettingMap and LogFortianalyzer3OverrideSettingMapOutput values.
// You can construct a concrete instance of `LogFortianalyzer3OverrideSettingMapInput` via:
//
//          LogFortianalyzer3OverrideSettingMap{ "key": LogFortianalyzer3OverrideSettingArgs{...} }
type LogFortianalyzer3OverrideSettingMapInput interface {
	pulumi.Input

	ToLogFortianalyzer3OverrideSettingMapOutput() LogFortianalyzer3OverrideSettingMapOutput
	ToLogFortianalyzer3OverrideSettingMapOutputWithContext(context.Context) LogFortianalyzer3OverrideSettingMapOutput
}

type LogFortianalyzer3OverrideSettingMap map[string]LogFortianalyzer3OverrideSettingInput

func (LogFortianalyzer3OverrideSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer3OverrideSetting)(nil)).Elem()
}

func (i LogFortianalyzer3OverrideSettingMap) ToLogFortianalyzer3OverrideSettingMapOutput() LogFortianalyzer3OverrideSettingMapOutput {
	return i.ToLogFortianalyzer3OverrideSettingMapOutputWithContext(context.Background())
}

func (i LogFortianalyzer3OverrideSettingMap) ToLogFortianalyzer3OverrideSettingMapOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3OverrideSettingMapOutput)
}

type LogFortianalyzer3OverrideSettingOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3OverrideSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3OverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzer3OverrideSettingOutput) ToLogFortianalyzer3OverrideSettingOutput() LogFortianalyzer3OverrideSettingOutput {
	return o
}

func (o LogFortianalyzer3OverrideSettingOutput) ToLogFortianalyzer3OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideSettingOutput {
	return o
}

type LogFortianalyzer3OverrideSettingArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3OverrideSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer3OverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzer3OverrideSettingArrayOutput) ToLogFortianalyzer3OverrideSettingArrayOutput() LogFortianalyzer3OverrideSettingArrayOutput {
	return o
}

func (o LogFortianalyzer3OverrideSettingArrayOutput) ToLogFortianalyzer3OverrideSettingArrayOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideSettingArrayOutput {
	return o
}

func (o LogFortianalyzer3OverrideSettingArrayOutput) Index(i pulumi.IntInput) LogFortianalyzer3OverrideSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzer3OverrideSetting {
		return vs[0].([]*LogFortianalyzer3OverrideSetting)[vs[1].(int)]
	}).(LogFortianalyzer3OverrideSettingOutput)
}

type LogFortianalyzer3OverrideSettingMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3OverrideSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer3OverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzer3OverrideSettingMapOutput) ToLogFortianalyzer3OverrideSettingMapOutput() LogFortianalyzer3OverrideSettingMapOutput {
	return o
}

func (o LogFortianalyzer3OverrideSettingMapOutput) ToLogFortianalyzer3OverrideSettingMapOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideSettingMapOutput {
	return o
}

func (o LogFortianalyzer3OverrideSettingMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzer3OverrideSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzer3OverrideSetting {
		return vs[0].(map[string]*LogFortianalyzer3OverrideSetting)[vs[1].(string)]
	}).(LogFortianalyzer3OverrideSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3OverrideSettingInput)(nil)).Elem(), &LogFortianalyzer3OverrideSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3OverrideSettingArrayInput)(nil)).Elem(), LogFortianalyzer3OverrideSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3OverrideSettingMapInput)(nil)).Elem(), LogFortianalyzer3OverrideSettingMap{})
	pulumi.RegisterOutputType(LogFortianalyzer3OverrideSettingOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3OverrideSettingArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3OverrideSettingMapOutput{})
}
