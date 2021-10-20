// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure logging to FortiCloud.
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
// 		_, err := fortios.NewLogFortiguardSetting(ctx, "trname", &fortios.LogFortiguardSettingArgs{
// 			EncAlgorithm:       pulumi.String("high"),
// 			SourceIp:           pulumi.String("0.0.0.0"),
// 			SslMinProtoVersion: pulumi.String("default"),
// 			Status:             pulumi.String("disable"),
// 			UploadInterval:     pulumi.String("daily"),
// 			UploadOption:       pulumi.String("5-minute"),
// 			UploadTime:         pulumi.String("00:00"),
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
// LogFortiguard Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logFortiguardSetting:LogFortiguardSetting labelname LogFortiguardSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortiguardSetting struct {
	pulumi.CustomResourceState

	// FortiGate Cloud connection timeout in seconds.
	ConnTimeout pulumi.IntOutput `pulumi:"connTimeout"`
	// Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringOutput `pulumi:"encAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntOutput `pulumi:"maxLogRate"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringOutput `pulumi:"priority"`
	// Source IP address used to connect FortiCloud.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Day of week to roll logs.
	UploadDay pulumi.StringOutput `pulumi:"uploadDay"`
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringOutput `pulumi:"uploadInterval"`
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringOutput `pulumi:"uploadOption"`
	// Time of day to roll logs (hh:mm).
	UploadTime pulumi.StringOutput `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogFortiguardSetting registers a new resource with the given unique name, arguments, and options.
func NewLogFortiguardSetting(ctx *pulumi.Context,
	name string, args *LogFortiguardSettingArgs, opts ...pulumi.ResourceOption) (*LogFortiguardSetting, error) {
	if args == nil {
		args = &LogFortiguardSettingArgs{}
	}

	var resource LogFortiguardSetting
	err := ctx.RegisterResource("fortios:index/logFortiguardSetting:LogFortiguardSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortiguardSetting gets an existing LogFortiguardSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortiguardSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortiguardSettingState, opts ...pulumi.ResourceOption) (*LogFortiguardSetting, error) {
	var resource LogFortiguardSetting
	err := ctx.ReadResource("fortios:index/logFortiguardSetting:LogFortiguardSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortiguardSetting resources.
type logFortiguardSettingState struct {
	// FortiGate Cloud connection timeout in seconds.
	ConnTimeout *int `pulumi:"connTimeout"`
	// Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Source IP address used to connect FortiCloud.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Day of week to roll logs.
	UploadDay *string `pulumi:"uploadDay"`
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval *string `pulumi:"uploadInterval"`
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption *string `pulumi:"uploadOption"`
	// Time of day to roll logs (hh:mm).
	UploadTime *string `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogFortiguardSettingState struct {
	// FortiGate Cloud connection timeout in seconds.
	ConnTimeout pulumi.IntPtrInput
	// Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Source IP address used to connect FortiCloud.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Day of week to roll logs.
	UploadDay pulumi.StringPtrInput
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringPtrInput
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringPtrInput
	// Time of day to roll logs (hh:mm).
	UploadTime pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogFortiguardSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortiguardSettingState)(nil)).Elem()
}

type logFortiguardSettingArgs struct {
	// FortiGate Cloud connection timeout in seconds.
	ConnTimeout *int `pulumi:"connTimeout"`
	// Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Source IP address used to connect FortiCloud.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Day of week to roll logs.
	UploadDay *string `pulumi:"uploadDay"`
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval *string `pulumi:"uploadInterval"`
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption *string `pulumi:"uploadOption"`
	// Time of day to roll logs (hh:mm).
	UploadTime *string `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogFortiguardSetting resource.
type LogFortiguardSettingArgs struct {
	// FortiGate Cloud connection timeout in seconds.
	ConnTimeout pulumi.IntPtrInput
	// Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Source IP address used to connect FortiCloud.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Day of week to roll logs.
	UploadDay pulumi.StringPtrInput
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringPtrInput
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringPtrInput
	// Time of day to roll logs (hh:mm).
	UploadTime pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogFortiguardSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortiguardSettingArgs)(nil)).Elem()
}

type LogFortiguardSettingInput interface {
	pulumi.Input

	ToLogFortiguardSettingOutput() LogFortiguardSettingOutput
	ToLogFortiguardSettingOutputWithContext(ctx context.Context) LogFortiguardSettingOutput
}

func (*LogFortiguardSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortiguardSetting)(nil))
}

func (i *LogFortiguardSetting) ToLogFortiguardSettingOutput() LogFortiguardSettingOutput {
	return i.ToLogFortiguardSettingOutputWithContext(context.Background())
}

func (i *LogFortiguardSetting) ToLogFortiguardSettingOutputWithContext(ctx context.Context) LogFortiguardSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardSettingOutput)
}

func (i *LogFortiguardSetting) ToLogFortiguardSettingPtrOutput() LogFortiguardSettingPtrOutput {
	return i.ToLogFortiguardSettingPtrOutputWithContext(context.Background())
}

func (i *LogFortiguardSetting) ToLogFortiguardSettingPtrOutputWithContext(ctx context.Context) LogFortiguardSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardSettingPtrOutput)
}

type LogFortiguardSettingPtrInput interface {
	pulumi.Input

	ToLogFortiguardSettingPtrOutput() LogFortiguardSettingPtrOutput
	ToLogFortiguardSettingPtrOutputWithContext(ctx context.Context) LogFortiguardSettingPtrOutput
}

type logFortiguardSettingPtrType LogFortiguardSettingArgs

func (*logFortiguardSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortiguardSetting)(nil))
}

func (i *logFortiguardSettingPtrType) ToLogFortiguardSettingPtrOutput() LogFortiguardSettingPtrOutput {
	return i.ToLogFortiguardSettingPtrOutputWithContext(context.Background())
}

func (i *logFortiguardSettingPtrType) ToLogFortiguardSettingPtrOutputWithContext(ctx context.Context) LogFortiguardSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardSettingPtrOutput)
}

// LogFortiguardSettingArrayInput is an input type that accepts LogFortiguardSettingArray and LogFortiguardSettingArrayOutput values.
// You can construct a concrete instance of `LogFortiguardSettingArrayInput` via:
//
//          LogFortiguardSettingArray{ LogFortiguardSettingArgs{...} }
type LogFortiguardSettingArrayInput interface {
	pulumi.Input

	ToLogFortiguardSettingArrayOutput() LogFortiguardSettingArrayOutput
	ToLogFortiguardSettingArrayOutputWithContext(context.Context) LogFortiguardSettingArrayOutput
}

type LogFortiguardSettingArray []LogFortiguardSettingInput

func (LogFortiguardSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogFortiguardSetting)(nil))
}

func (i LogFortiguardSettingArray) ToLogFortiguardSettingArrayOutput() LogFortiguardSettingArrayOutput {
	return i.ToLogFortiguardSettingArrayOutputWithContext(context.Background())
}

func (i LogFortiguardSettingArray) ToLogFortiguardSettingArrayOutputWithContext(ctx context.Context) LogFortiguardSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardSettingArrayOutput)
}

// LogFortiguardSettingMapInput is an input type that accepts LogFortiguardSettingMap and LogFortiguardSettingMapOutput values.
// You can construct a concrete instance of `LogFortiguardSettingMapInput` via:
//
//          LogFortiguardSettingMap{ "key": LogFortiguardSettingArgs{...} }
type LogFortiguardSettingMapInput interface {
	pulumi.Input

	ToLogFortiguardSettingMapOutput() LogFortiguardSettingMapOutput
	ToLogFortiguardSettingMapOutputWithContext(context.Context) LogFortiguardSettingMapOutput
}

type LogFortiguardSettingMap map[string]LogFortiguardSettingInput

func (LogFortiguardSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogFortiguardSetting)(nil))
}

func (i LogFortiguardSettingMap) ToLogFortiguardSettingMapOutput() LogFortiguardSettingMapOutput {
	return i.ToLogFortiguardSettingMapOutputWithContext(context.Background())
}

func (i LogFortiguardSettingMap) ToLogFortiguardSettingMapOutputWithContext(ctx context.Context) LogFortiguardSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardSettingMapOutput)
}

type LogFortiguardSettingOutput struct {
	*pulumi.OutputState
}

func (LogFortiguardSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortiguardSetting)(nil))
}

func (o LogFortiguardSettingOutput) ToLogFortiguardSettingOutput() LogFortiguardSettingOutput {
	return o
}

func (o LogFortiguardSettingOutput) ToLogFortiguardSettingOutputWithContext(ctx context.Context) LogFortiguardSettingOutput {
	return o
}

func (o LogFortiguardSettingOutput) ToLogFortiguardSettingPtrOutput() LogFortiguardSettingPtrOutput {
	return o.ToLogFortiguardSettingPtrOutputWithContext(context.Background())
}

func (o LogFortiguardSettingOutput) ToLogFortiguardSettingPtrOutputWithContext(ctx context.Context) LogFortiguardSettingPtrOutput {
	return o.ApplyT(func(v LogFortiguardSetting) *LogFortiguardSetting {
		return &v
	}).(LogFortiguardSettingPtrOutput)
}

type LogFortiguardSettingPtrOutput struct {
	*pulumi.OutputState
}

func (LogFortiguardSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortiguardSetting)(nil))
}

func (o LogFortiguardSettingPtrOutput) ToLogFortiguardSettingPtrOutput() LogFortiguardSettingPtrOutput {
	return o
}

func (o LogFortiguardSettingPtrOutput) ToLogFortiguardSettingPtrOutputWithContext(ctx context.Context) LogFortiguardSettingPtrOutput {
	return o
}

type LogFortiguardSettingArrayOutput struct{ *pulumi.OutputState }

func (LogFortiguardSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogFortiguardSetting)(nil))
}

func (o LogFortiguardSettingArrayOutput) ToLogFortiguardSettingArrayOutput() LogFortiguardSettingArrayOutput {
	return o
}

func (o LogFortiguardSettingArrayOutput) ToLogFortiguardSettingArrayOutputWithContext(ctx context.Context) LogFortiguardSettingArrayOutput {
	return o
}

func (o LogFortiguardSettingArrayOutput) Index(i pulumi.IntInput) LogFortiguardSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogFortiguardSetting {
		return vs[0].([]LogFortiguardSetting)[vs[1].(int)]
	}).(LogFortiguardSettingOutput)
}

type LogFortiguardSettingMapOutput struct{ *pulumi.OutputState }

func (LogFortiguardSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogFortiguardSetting)(nil))
}

func (o LogFortiguardSettingMapOutput) ToLogFortiguardSettingMapOutput() LogFortiguardSettingMapOutput {
	return o
}

func (o LogFortiguardSettingMapOutput) ToLogFortiguardSettingMapOutputWithContext(ctx context.Context) LogFortiguardSettingMapOutput {
	return o
}

func (o LogFortiguardSettingMapOutput) MapIndex(k pulumi.StringInput) LogFortiguardSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogFortiguardSetting {
		return vs[0].(map[string]LogFortiguardSetting)[vs[1].(string)]
	}).(LogFortiguardSettingOutput)
}

func init() {
	pulumi.RegisterOutputType(LogFortiguardSettingOutput{})
	pulumi.RegisterOutputType(LogFortiguardSettingPtrOutput{})
	pulumi.RegisterOutputType(LogFortiguardSettingArrayOutput{})
	pulumi.RegisterOutputType(LogFortiguardSettingMapOutput{})
}