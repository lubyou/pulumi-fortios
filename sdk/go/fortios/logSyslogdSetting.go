// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Global settings for remote syslog server.
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
// 		_, err := fortios.NewLogSyslogdSetting(ctx, "trname", &fortios.LogSyslogdSettingArgs{
// 			EncAlgorithm:       pulumi.String("disable"),
// 			Facility:           pulumi.String("local7"),
// 			Format:             pulumi.String("default"),
// 			Mode:               pulumi.String("udp"),
// 			Port:               pulumi.Int(514),
// 			SslMinProtoVersion: pulumi.String("default"),
// 			Status:             pulumi.String("disable"),
// 			SyslogType:         pulumi.Int(1),
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
// LogSyslogd Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogdSetting:LogSyslogdSetting labelname LogSyslogdSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogdSetting struct {
	pulumi.CustomResourceState

	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogdSettingCustomFieldNameArrayOutput `pulumi:"customFieldNames"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm pulumi.StringOutput `pulumi:"encAlgorithm"`
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility pulumi.StringOutput `pulumi:"facility"`
	// Log format.
	Format pulumi.StringOutput `pulumi:"format"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntOutput `pulumi:"maxLogRate"`
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Server listen port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringOutput `pulumi:"priority"`
	// Address of remote syslog server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Hidden setting index of Syslog.
	SyslogType pulumi.IntOutput `pulumi:"syslogType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogSyslogdSetting registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogdSetting(ctx *pulumi.Context,
	name string, args *LogSyslogdSettingArgs, opts ...pulumi.ResourceOption) (*LogSyslogdSetting, error) {
	if args == nil {
		args = &LogSyslogdSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogdSetting
	err := ctx.RegisterResource("fortios:index/logSyslogdSetting:LogSyslogdSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogdSetting gets an existing LogSyslogdSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogdSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogdSettingState, opts ...pulumi.ResourceOption) (*LogSyslogdSetting, error) {
	var resource LogSyslogdSetting
	err := ctx.ReadResource("fortios:index/logSyslogdSetting:LogSyslogdSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogdSetting resources.
type logSyslogdSettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []LogSyslogdSettingCustomFieldName `pulumi:"customFieldNames"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility *string `pulumi:"facility"`
	// Log format.
	Format *string `pulumi:"format"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode *string `pulumi:"mode"`
	// Server listen port.
	Port *int `pulumi:"port"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Address of remote syslog server.
	Server *string `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Hidden setting index of Syslog.
	SyslogType *int `pulumi:"syslogType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogSyslogdSettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogdSettingCustomFieldNameArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm pulumi.StringPtrInput
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility pulumi.StringPtrInput
	// Log format.
	Format pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode pulumi.StringPtrInput
	// Server listen port.
	Port pulumi.IntPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Address of remote syslog server.
	Server pulumi.StringPtrInput
	// Source IP address of syslog.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Hidden setting index of Syslog.
	SyslogType pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogSyslogdSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdSettingState)(nil)).Elem()
}

type logSyslogdSettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []LogSyslogdSettingCustomFieldName `pulumi:"customFieldNames"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility *string `pulumi:"facility"`
	// Log format.
	Format *string `pulumi:"format"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode *string `pulumi:"mode"`
	// Server listen port.
	Port *int `pulumi:"port"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Address of remote syslog server.
	Server *string `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Hidden setting index of Syslog.
	SyslogType *int `pulumi:"syslogType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogSyslogdSetting resource.
type LogSyslogdSettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogdSettingCustomFieldNameArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm pulumi.StringPtrInput
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility pulumi.StringPtrInput
	// Log format.
	Format pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode pulumi.StringPtrInput
	// Server listen port.
	Port pulumi.IntPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Address of remote syslog server.
	Server pulumi.StringPtrInput
	// Source IP address of syslog.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Hidden setting index of Syslog.
	SyslogType pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogSyslogdSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdSettingArgs)(nil)).Elem()
}

type LogSyslogdSettingInput interface {
	pulumi.Input

	ToLogSyslogdSettingOutput() LogSyslogdSettingOutput
	ToLogSyslogdSettingOutputWithContext(ctx context.Context) LogSyslogdSettingOutput
}

func (*LogSyslogdSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdSetting)(nil)).Elem()
}

func (i *LogSyslogdSetting) ToLogSyslogdSettingOutput() LogSyslogdSettingOutput {
	return i.ToLogSyslogdSettingOutputWithContext(context.Background())
}

func (i *LogSyslogdSetting) ToLogSyslogdSettingOutputWithContext(ctx context.Context) LogSyslogdSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdSettingOutput)
}

// LogSyslogdSettingArrayInput is an input type that accepts LogSyslogdSettingArray and LogSyslogdSettingArrayOutput values.
// You can construct a concrete instance of `LogSyslogdSettingArrayInput` via:
//
//          LogSyslogdSettingArray{ LogSyslogdSettingArgs{...} }
type LogSyslogdSettingArrayInput interface {
	pulumi.Input

	ToLogSyslogdSettingArrayOutput() LogSyslogdSettingArrayOutput
	ToLogSyslogdSettingArrayOutputWithContext(context.Context) LogSyslogdSettingArrayOutput
}

type LogSyslogdSettingArray []LogSyslogdSettingInput

func (LogSyslogdSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdSetting)(nil)).Elem()
}

func (i LogSyslogdSettingArray) ToLogSyslogdSettingArrayOutput() LogSyslogdSettingArrayOutput {
	return i.ToLogSyslogdSettingArrayOutputWithContext(context.Background())
}

func (i LogSyslogdSettingArray) ToLogSyslogdSettingArrayOutputWithContext(ctx context.Context) LogSyslogdSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdSettingArrayOutput)
}

// LogSyslogdSettingMapInput is an input type that accepts LogSyslogdSettingMap and LogSyslogdSettingMapOutput values.
// You can construct a concrete instance of `LogSyslogdSettingMapInput` via:
//
//          LogSyslogdSettingMap{ "key": LogSyslogdSettingArgs{...} }
type LogSyslogdSettingMapInput interface {
	pulumi.Input

	ToLogSyslogdSettingMapOutput() LogSyslogdSettingMapOutput
	ToLogSyslogdSettingMapOutputWithContext(context.Context) LogSyslogdSettingMapOutput
}

type LogSyslogdSettingMap map[string]LogSyslogdSettingInput

func (LogSyslogdSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdSetting)(nil)).Elem()
}

func (i LogSyslogdSettingMap) ToLogSyslogdSettingMapOutput() LogSyslogdSettingMapOutput {
	return i.ToLogSyslogdSettingMapOutputWithContext(context.Background())
}

func (i LogSyslogdSettingMap) ToLogSyslogdSettingMapOutputWithContext(ctx context.Context) LogSyslogdSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdSettingMapOutput)
}

type LogSyslogdSettingOutput struct{ *pulumi.OutputState }

func (LogSyslogdSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdSetting)(nil)).Elem()
}

func (o LogSyslogdSettingOutput) ToLogSyslogdSettingOutput() LogSyslogdSettingOutput {
	return o
}

func (o LogSyslogdSettingOutput) ToLogSyslogdSettingOutputWithContext(ctx context.Context) LogSyslogdSettingOutput {
	return o
}

type LogSyslogdSettingArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogdSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdSetting)(nil)).Elem()
}

func (o LogSyslogdSettingArrayOutput) ToLogSyslogdSettingArrayOutput() LogSyslogdSettingArrayOutput {
	return o
}

func (o LogSyslogdSettingArrayOutput) ToLogSyslogdSettingArrayOutputWithContext(ctx context.Context) LogSyslogdSettingArrayOutput {
	return o
}

func (o LogSyslogdSettingArrayOutput) Index(i pulumi.IntInput) LogSyslogdSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogdSetting {
		return vs[0].([]*LogSyslogdSetting)[vs[1].(int)]
	}).(LogSyslogdSettingOutput)
}

type LogSyslogdSettingMapOutput struct{ *pulumi.OutputState }

func (LogSyslogdSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdSetting)(nil)).Elem()
}

func (o LogSyslogdSettingMapOutput) ToLogSyslogdSettingMapOutput() LogSyslogdSettingMapOutput {
	return o
}

func (o LogSyslogdSettingMapOutput) ToLogSyslogdSettingMapOutputWithContext(ctx context.Context) LogSyslogdSettingMapOutput {
	return o
}

func (o LogSyslogdSettingMapOutput) MapIndex(k pulumi.StringInput) LogSyslogdSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogdSetting {
		return vs[0].(map[string]*LogSyslogdSetting)[vs[1].(string)]
	}).(LogSyslogdSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdSettingInput)(nil)).Elem(), &LogSyslogdSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdSettingArrayInput)(nil)).Elem(), LogSyslogdSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdSettingMapInput)(nil)).Elem(), LogSyslogdSettingMap{})
	pulumi.RegisterOutputType(LogSyslogdSettingOutput{})
	pulumi.RegisterOutputType(LogSyslogdSettingArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogdSettingMapOutput{})
}
