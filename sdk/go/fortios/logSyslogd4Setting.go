// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

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
// 		_, err := fortios.NewLogSyslogd4Setting(ctx, "trname", &fortios.LogSyslogd4SettingArgs{
// 			EncAlgorithm:       pulumi.String("disable"),
// 			Facility:           pulumi.String("local7"),
// 			Format:             pulumi.String("default"),
// 			Mode:               pulumi.String("udp"),
// 			Port:               pulumi.Int(514),
// 			SslMinProtoVersion: pulumi.String("default"),
// 			Status:             pulumi.String("disable"),
// 			SyslogType:         pulumi.Int(4),
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
// LogSyslogd4 Setting can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd4Setting:LogSyslogd4Setting labelname LogSyslogd4Setting
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd4Setting:LogSyslogd4Setting labelname LogSyslogd4Setting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogd4Setting struct {
	pulumi.CustomResourceState

	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd4SettingCustomFieldNameArrayOutput `pulumi:"customFieldNames"`
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

// NewLogSyslogd4Setting registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd4Setting(ctx *pulumi.Context,
	name string, args *LogSyslogd4SettingArgs, opts ...pulumi.ResourceOption) (*LogSyslogd4Setting, error) {
	if args == nil {
		args = &LogSyslogd4SettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogd4Setting
	err := ctx.RegisterResource("fortios:index/logSyslogd4Setting:LogSyslogd4Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd4Setting gets an existing LogSyslogd4Setting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd4Setting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd4SettingState, opts ...pulumi.ResourceOption) (*LogSyslogd4Setting, error) {
	var resource LogSyslogd4Setting
	err := ctx.ReadResource("fortios:index/logSyslogd4Setting:LogSyslogd4Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd4Setting resources.
type logSyslogd4SettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []LogSyslogd4SettingCustomFieldName `pulumi:"customFieldNames"`
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

type LogSyslogd4SettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd4SettingCustomFieldNameArrayInput
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

func (LogSyslogd4SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4SettingState)(nil)).Elem()
}

type logSyslogd4SettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []LogSyslogd4SettingCustomFieldName `pulumi:"customFieldNames"`
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

// The set of arguments for constructing a LogSyslogd4Setting resource.
type LogSyslogd4SettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd4SettingCustomFieldNameArrayInput
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

func (LogSyslogd4SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4SettingArgs)(nil)).Elem()
}

type LogSyslogd4SettingInput interface {
	pulumi.Input

	ToLogSyslogd4SettingOutput() LogSyslogd4SettingOutput
	ToLogSyslogd4SettingOutputWithContext(ctx context.Context) LogSyslogd4SettingOutput
}

func (*LogSyslogd4Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd4Setting)(nil)).Elem()
}

func (i *LogSyslogd4Setting) ToLogSyslogd4SettingOutput() LogSyslogd4SettingOutput {
	return i.ToLogSyslogd4SettingOutputWithContext(context.Background())
}

func (i *LogSyslogd4Setting) ToLogSyslogd4SettingOutputWithContext(ctx context.Context) LogSyslogd4SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4SettingOutput)
}

// LogSyslogd4SettingArrayInput is an input type that accepts LogSyslogd4SettingArray and LogSyslogd4SettingArrayOutput values.
// You can construct a concrete instance of `LogSyslogd4SettingArrayInput` via:
//
//          LogSyslogd4SettingArray{ LogSyslogd4SettingArgs{...} }
type LogSyslogd4SettingArrayInput interface {
	pulumi.Input

	ToLogSyslogd4SettingArrayOutput() LogSyslogd4SettingArrayOutput
	ToLogSyslogd4SettingArrayOutputWithContext(context.Context) LogSyslogd4SettingArrayOutput
}

type LogSyslogd4SettingArray []LogSyslogd4SettingInput

func (LogSyslogd4SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd4Setting)(nil)).Elem()
}

func (i LogSyslogd4SettingArray) ToLogSyslogd4SettingArrayOutput() LogSyslogd4SettingArrayOutput {
	return i.ToLogSyslogd4SettingArrayOutputWithContext(context.Background())
}

func (i LogSyslogd4SettingArray) ToLogSyslogd4SettingArrayOutputWithContext(ctx context.Context) LogSyslogd4SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4SettingArrayOutput)
}

// LogSyslogd4SettingMapInput is an input type that accepts LogSyslogd4SettingMap and LogSyslogd4SettingMapOutput values.
// You can construct a concrete instance of `LogSyslogd4SettingMapInput` via:
//
//          LogSyslogd4SettingMap{ "key": LogSyslogd4SettingArgs{...} }
type LogSyslogd4SettingMapInput interface {
	pulumi.Input

	ToLogSyslogd4SettingMapOutput() LogSyslogd4SettingMapOutput
	ToLogSyslogd4SettingMapOutputWithContext(context.Context) LogSyslogd4SettingMapOutput
}

type LogSyslogd4SettingMap map[string]LogSyslogd4SettingInput

func (LogSyslogd4SettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd4Setting)(nil)).Elem()
}

func (i LogSyslogd4SettingMap) ToLogSyslogd4SettingMapOutput() LogSyslogd4SettingMapOutput {
	return i.ToLogSyslogd4SettingMapOutputWithContext(context.Background())
}

func (i LogSyslogd4SettingMap) ToLogSyslogd4SettingMapOutputWithContext(ctx context.Context) LogSyslogd4SettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4SettingMapOutput)
}

type LogSyslogd4SettingOutput struct{ *pulumi.OutputState }

func (LogSyslogd4SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd4Setting)(nil)).Elem()
}

func (o LogSyslogd4SettingOutput) ToLogSyslogd4SettingOutput() LogSyslogd4SettingOutput {
	return o
}

func (o LogSyslogd4SettingOutput) ToLogSyslogd4SettingOutputWithContext(ctx context.Context) LogSyslogd4SettingOutput {
	return o
}

type LogSyslogd4SettingArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd4SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd4Setting)(nil)).Elem()
}

func (o LogSyslogd4SettingArrayOutput) ToLogSyslogd4SettingArrayOutput() LogSyslogd4SettingArrayOutput {
	return o
}

func (o LogSyslogd4SettingArrayOutput) ToLogSyslogd4SettingArrayOutputWithContext(ctx context.Context) LogSyslogd4SettingArrayOutput {
	return o
}

func (o LogSyslogd4SettingArrayOutput) Index(i pulumi.IntInput) LogSyslogd4SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogd4Setting {
		return vs[0].([]*LogSyslogd4Setting)[vs[1].(int)]
	}).(LogSyslogd4SettingOutput)
}

type LogSyslogd4SettingMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd4SettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd4Setting)(nil)).Elem()
}

func (o LogSyslogd4SettingMapOutput) ToLogSyslogd4SettingMapOutput() LogSyslogd4SettingMapOutput {
	return o
}

func (o LogSyslogd4SettingMapOutput) ToLogSyslogd4SettingMapOutputWithContext(ctx context.Context) LogSyslogd4SettingMapOutput {
	return o
}

func (o LogSyslogd4SettingMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd4SettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogd4Setting {
		return vs[0].(map[string]*LogSyslogd4Setting)[vs[1].(string)]
	}).(LogSyslogd4SettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4SettingInput)(nil)).Elem(), &LogSyslogd4Setting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4SettingArrayInput)(nil)).Elem(), LogSyslogd4SettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4SettingMapInput)(nil)).Elem(), LogSyslogd4SettingMap{})
	pulumi.RegisterOutputType(LogSyslogd4SettingOutput{})
	pulumi.RegisterOutputType(LogSyslogd4SettingArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd4SettingMapOutput{})
}
