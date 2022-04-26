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
// 		_, err := fortios.NewLogFortianalyzerOverrideSetting(ctx, "trname", &fortios.LogFortianalyzerOverrideSettingArgs{
// 			__changeIp:                pulumi.Int(0),
// 			ConnTimeout:               pulumi.Int(10),
// 			EncAlgorithm:              pulumi.String("high"),
// 			FazType:                   pulumi.Int(4),
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
// LogFortianalyzer OverrideSetting can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzerOverrideSetting:LogFortianalyzerOverrideSetting labelname LogFortianalyzerOverrideSetting
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzerOverrideSetting:LogFortianalyzerOverrideSetting labelname LogFortianalyzerOverrideSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzerOverrideSetting struct {
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
	Serials LogFortianalyzerOverrideSettingSerialArrayOutput `pulumi:"serials"`
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

// NewLogFortianalyzerOverrideSetting registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzerOverrideSetting(ctx *pulumi.Context,
	name string, args *LogFortianalyzerOverrideSettingArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzerOverrideSetting, error) {
	if args == nil {
		args = &LogFortianalyzerOverrideSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogFortianalyzerOverrideSetting
	err := ctx.RegisterResource("fortios:index/logFortianalyzerOverrideSetting:LogFortianalyzerOverrideSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzerOverrideSetting gets an existing LogFortianalyzerOverrideSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzerOverrideSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzerOverrideSettingState, opts ...pulumi.ResourceOption) (*LogFortianalyzerOverrideSetting, error) {
	var resource LogFortianalyzerOverrideSetting
	err := ctx.ReadResource("fortios:index/logFortianalyzerOverrideSetting:LogFortianalyzerOverrideSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzerOverrideSetting resources.
type logFortianalyzerOverrideSettingState struct {
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
	Serials []LogFortianalyzerOverrideSettingSerial `pulumi:"serials"`
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

type LogFortianalyzerOverrideSettingState struct {
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
	Serials LogFortianalyzerOverrideSettingSerialArrayInput
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

func (LogFortianalyzerOverrideSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerOverrideSettingState)(nil)).Elem()
}

type logFortianalyzerOverrideSettingArgs struct {
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
	Serials []LogFortianalyzerOverrideSettingSerial `pulumi:"serials"`
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

// The set of arguments for constructing a LogFortianalyzerOverrideSetting resource.
type LogFortianalyzerOverrideSettingArgs struct {
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
	Serials LogFortianalyzerOverrideSettingSerialArrayInput
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

func (LogFortianalyzerOverrideSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerOverrideSettingArgs)(nil)).Elem()
}

type LogFortianalyzerOverrideSettingInput interface {
	pulumi.Input

	ToLogFortianalyzerOverrideSettingOutput() LogFortianalyzerOverrideSettingOutput
	ToLogFortianalyzerOverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzerOverrideSettingOutput
}

func (*LogFortianalyzerOverrideSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerOverrideSetting)(nil)).Elem()
}

func (i *LogFortianalyzerOverrideSetting) ToLogFortianalyzerOverrideSettingOutput() LogFortianalyzerOverrideSettingOutput {
	return i.ToLogFortianalyzerOverrideSettingOutputWithContext(context.Background())
}

func (i *LogFortianalyzerOverrideSetting) ToLogFortianalyzerOverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzerOverrideSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerOverrideSettingOutput)
}

// LogFortianalyzerOverrideSettingArrayInput is an input type that accepts LogFortianalyzerOverrideSettingArray and LogFortianalyzerOverrideSettingArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzerOverrideSettingArrayInput` via:
//
//          LogFortianalyzerOverrideSettingArray{ LogFortianalyzerOverrideSettingArgs{...} }
type LogFortianalyzerOverrideSettingArrayInput interface {
	pulumi.Input

	ToLogFortianalyzerOverrideSettingArrayOutput() LogFortianalyzerOverrideSettingArrayOutput
	ToLogFortianalyzerOverrideSettingArrayOutputWithContext(context.Context) LogFortianalyzerOverrideSettingArrayOutput
}

type LogFortianalyzerOverrideSettingArray []LogFortianalyzerOverrideSettingInput

func (LogFortianalyzerOverrideSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerOverrideSetting)(nil)).Elem()
}

func (i LogFortianalyzerOverrideSettingArray) ToLogFortianalyzerOverrideSettingArrayOutput() LogFortianalyzerOverrideSettingArrayOutput {
	return i.ToLogFortianalyzerOverrideSettingArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzerOverrideSettingArray) ToLogFortianalyzerOverrideSettingArrayOutputWithContext(ctx context.Context) LogFortianalyzerOverrideSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerOverrideSettingArrayOutput)
}

// LogFortianalyzerOverrideSettingMapInput is an input type that accepts LogFortianalyzerOverrideSettingMap and LogFortianalyzerOverrideSettingMapOutput values.
// You can construct a concrete instance of `LogFortianalyzerOverrideSettingMapInput` via:
//
//          LogFortianalyzerOverrideSettingMap{ "key": LogFortianalyzerOverrideSettingArgs{...} }
type LogFortianalyzerOverrideSettingMapInput interface {
	pulumi.Input

	ToLogFortianalyzerOverrideSettingMapOutput() LogFortianalyzerOverrideSettingMapOutput
	ToLogFortianalyzerOverrideSettingMapOutputWithContext(context.Context) LogFortianalyzerOverrideSettingMapOutput
}

type LogFortianalyzerOverrideSettingMap map[string]LogFortianalyzerOverrideSettingInput

func (LogFortianalyzerOverrideSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerOverrideSetting)(nil)).Elem()
}

func (i LogFortianalyzerOverrideSettingMap) ToLogFortianalyzerOverrideSettingMapOutput() LogFortianalyzerOverrideSettingMapOutput {
	return i.ToLogFortianalyzerOverrideSettingMapOutputWithContext(context.Background())
}

func (i LogFortianalyzerOverrideSettingMap) ToLogFortianalyzerOverrideSettingMapOutputWithContext(ctx context.Context) LogFortianalyzerOverrideSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerOverrideSettingMapOutput)
}

type LogFortianalyzerOverrideSettingOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerOverrideSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerOverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzerOverrideSettingOutput) ToLogFortianalyzerOverrideSettingOutput() LogFortianalyzerOverrideSettingOutput {
	return o
}

func (o LogFortianalyzerOverrideSettingOutput) ToLogFortianalyzerOverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzerOverrideSettingOutput {
	return o
}

type LogFortianalyzerOverrideSettingArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerOverrideSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerOverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzerOverrideSettingArrayOutput) ToLogFortianalyzerOverrideSettingArrayOutput() LogFortianalyzerOverrideSettingArrayOutput {
	return o
}

func (o LogFortianalyzerOverrideSettingArrayOutput) ToLogFortianalyzerOverrideSettingArrayOutputWithContext(ctx context.Context) LogFortianalyzerOverrideSettingArrayOutput {
	return o
}

func (o LogFortianalyzerOverrideSettingArrayOutput) Index(i pulumi.IntInput) LogFortianalyzerOverrideSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzerOverrideSetting {
		return vs[0].([]*LogFortianalyzerOverrideSetting)[vs[1].(int)]
	}).(LogFortianalyzerOverrideSettingOutput)
}

type LogFortianalyzerOverrideSettingMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerOverrideSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerOverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzerOverrideSettingMapOutput) ToLogFortianalyzerOverrideSettingMapOutput() LogFortianalyzerOverrideSettingMapOutput {
	return o
}

func (o LogFortianalyzerOverrideSettingMapOutput) ToLogFortianalyzerOverrideSettingMapOutputWithContext(ctx context.Context) LogFortianalyzerOverrideSettingMapOutput {
	return o
}

func (o LogFortianalyzerOverrideSettingMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzerOverrideSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzerOverrideSetting {
		return vs[0].(map[string]*LogFortianalyzerOverrideSetting)[vs[1].(string)]
	}).(LogFortianalyzerOverrideSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerOverrideSettingInput)(nil)).Elem(), &LogFortianalyzerOverrideSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerOverrideSettingArrayInput)(nil)).Elem(), LogFortianalyzerOverrideSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerOverrideSettingMapInput)(nil)).Elem(), LogFortianalyzerOverrideSettingMap{})
	pulumi.RegisterOutputType(LogFortianalyzerOverrideSettingOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerOverrideSettingArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerOverrideSettingMapOutput{})
}
