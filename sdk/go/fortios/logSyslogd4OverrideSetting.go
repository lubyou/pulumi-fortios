// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Override settings for remote syslog server.
//
// ## Import
//
// LogSyslogd4 OverrideSetting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd4OverrideSetting:LogSyslogd4OverrideSetting labelname LogSyslogd4OverrideSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogd4OverrideSetting struct {
	pulumi.CustomResourceState

	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd4OverrideSettingCustomFieldNameArrayOutput `pulumi:"customFieldNames"`
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
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override pulumi.StringOutput `pulumi:"override"`
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

// NewLogSyslogd4OverrideSetting registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd4OverrideSetting(ctx *pulumi.Context,
	name string, args *LogSyslogd4OverrideSettingArgs, opts ...pulumi.ResourceOption) (*LogSyslogd4OverrideSetting, error) {
	if args == nil {
		args = &LogSyslogd4OverrideSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogd4OverrideSetting
	err := ctx.RegisterResource("fortios:index/logSyslogd4OverrideSetting:LogSyslogd4OverrideSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd4OverrideSetting gets an existing LogSyslogd4OverrideSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd4OverrideSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd4OverrideSettingState, opts ...pulumi.ResourceOption) (*LogSyslogd4OverrideSetting, error) {
	var resource LogSyslogd4OverrideSetting
	err := ctx.ReadResource("fortios:index/logSyslogd4OverrideSetting:LogSyslogd4OverrideSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd4OverrideSetting resources.
type logSyslogd4OverrideSettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []LogSyslogd4OverrideSettingCustomFieldName `pulumi:"customFieldNames"`
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
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
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

type LogSyslogd4OverrideSettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd4OverrideSettingCustomFieldNameArrayInput
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
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
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

func (LogSyslogd4OverrideSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4OverrideSettingState)(nil)).Elem()
}

type logSyslogd4OverrideSettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []LogSyslogd4OverrideSettingCustomFieldName `pulumi:"customFieldNames"`
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
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
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

// The set of arguments for constructing a LogSyslogd4OverrideSetting resource.
type LogSyslogd4OverrideSettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd4OverrideSettingCustomFieldNameArrayInput
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
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
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

func (LogSyslogd4OverrideSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4OverrideSettingArgs)(nil)).Elem()
}

type LogSyslogd4OverrideSettingInput interface {
	pulumi.Input

	ToLogSyslogd4OverrideSettingOutput() LogSyslogd4OverrideSettingOutput
	ToLogSyslogd4OverrideSettingOutputWithContext(ctx context.Context) LogSyslogd4OverrideSettingOutput
}

func (*LogSyslogd4OverrideSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd4OverrideSetting)(nil)).Elem()
}

func (i *LogSyslogd4OverrideSetting) ToLogSyslogd4OverrideSettingOutput() LogSyslogd4OverrideSettingOutput {
	return i.ToLogSyslogd4OverrideSettingOutputWithContext(context.Background())
}

func (i *LogSyslogd4OverrideSetting) ToLogSyslogd4OverrideSettingOutputWithContext(ctx context.Context) LogSyslogd4OverrideSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4OverrideSettingOutput)
}

// LogSyslogd4OverrideSettingArrayInput is an input type that accepts LogSyslogd4OverrideSettingArray and LogSyslogd4OverrideSettingArrayOutput values.
// You can construct a concrete instance of `LogSyslogd4OverrideSettingArrayInput` via:
//
//          LogSyslogd4OverrideSettingArray{ LogSyslogd4OverrideSettingArgs{...} }
type LogSyslogd4OverrideSettingArrayInput interface {
	pulumi.Input

	ToLogSyslogd4OverrideSettingArrayOutput() LogSyslogd4OverrideSettingArrayOutput
	ToLogSyslogd4OverrideSettingArrayOutputWithContext(context.Context) LogSyslogd4OverrideSettingArrayOutput
}

type LogSyslogd4OverrideSettingArray []LogSyslogd4OverrideSettingInput

func (LogSyslogd4OverrideSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd4OverrideSetting)(nil)).Elem()
}

func (i LogSyslogd4OverrideSettingArray) ToLogSyslogd4OverrideSettingArrayOutput() LogSyslogd4OverrideSettingArrayOutput {
	return i.ToLogSyslogd4OverrideSettingArrayOutputWithContext(context.Background())
}

func (i LogSyslogd4OverrideSettingArray) ToLogSyslogd4OverrideSettingArrayOutputWithContext(ctx context.Context) LogSyslogd4OverrideSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4OverrideSettingArrayOutput)
}

// LogSyslogd4OverrideSettingMapInput is an input type that accepts LogSyslogd4OverrideSettingMap and LogSyslogd4OverrideSettingMapOutput values.
// You can construct a concrete instance of `LogSyslogd4OverrideSettingMapInput` via:
//
//          LogSyslogd4OverrideSettingMap{ "key": LogSyslogd4OverrideSettingArgs{...} }
type LogSyslogd4OverrideSettingMapInput interface {
	pulumi.Input

	ToLogSyslogd4OverrideSettingMapOutput() LogSyslogd4OverrideSettingMapOutput
	ToLogSyslogd4OverrideSettingMapOutputWithContext(context.Context) LogSyslogd4OverrideSettingMapOutput
}

type LogSyslogd4OverrideSettingMap map[string]LogSyslogd4OverrideSettingInput

func (LogSyslogd4OverrideSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd4OverrideSetting)(nil)).Elem()
}

func (i LogSyslogd4OverrideSettingMap) ToLogSyslogd4OverrideSettingMapOutput() LogSyslogd4OverrideSettingMapOutput {
	return i.ToLogSyslogd4OverrideSettingMapOutputWithContext(context.Background())
}

func (i LogSyslogd4OverrideSettingMap) ToLogSyslogd4OverrideSettingMapOutputWithContext(ctx context.Context) LogSyslogd4OverrideSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4OverrideSettingMapOutput)
}

type LogSyslogd4OverrideSettingOutput struct{ *pulumi.OutputState }

func (LogSyslogd4OverrideSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd4OverrideSetting)(nil)).Elem()
}

func (o LogSyslogd4OverrideSettingOutput) ToLogSyslogd4OverrideSettingOutput() LogSyslogd4OverrideSettingOutput {
	return o
}

func (o LogSyslogd4OverrideSettingOutput) ToLogSyslogd4OverrideSettingOutputWithContext(ctx context.Context) LogSyslogd4OverrideSettingOutput {
	return o
}

type LogSyslogd4OverrideSettingArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd4OverrideSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd4OverrideSetting)(nil)).Elem()
}

func (o LogSyslogd4OverrideSettingArrayOutput) ToLogSyslogd4OverrideSettingArrayOutput() LogSyslogd4OverrideSettingArrayOutput {
	return o
}

func (o LogSyslogd4OverrideSettingArrayOutput) ToLogSyslogd4OverrideSettingArrayOutputWithContext(ctx context.Context) LogSyslogd4OverrideSettingArrayOutput {
	return o
}

func (o LogSyslogd4OverrideSettingArrayOutput) Index(i pulumi.IntInput) LogSyslogd4OverrideSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogd4OverrideSetting {
		return vs[0].([]*LogSyslogd4OverrideSetting)[vs[1].(int)]
	}).(LogSyslogd4OverrideSettingOutput)
}

type LogSyslogd4OverrideSettingMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd4OverrideSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd4OverrideSetting)(nil)).Elem()
}

func (o LogSyslogd4OverrideSettingMapOutput) ToLogSyslogd4OverrideSettingMapOutput() LogSyslogd4OverrideSettingMapOutput {
	return o
}

func (o LogSyslogd4OverrideSettingMapOutput) ToLogSyslogd4OverrideSettingMapOutputWithContext(ctx context.Context) LogSyslogd4OverrideSettingMapOutput {
	return o
}

func (o LogSyslogd4OverrideSettingMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd4OverrideSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogd4OverrideSetting {
		return vs[0].(map[string]*LogSyslogd4OverrideSetting)[vs[1].(string)]
	}).(LogSyslogd4OverrideSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4OverrideSettingInput)(nil)).Elem(), &LogSyslogd4OverrideSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4OverrideSettingArrayInput)(nil)).Elem(), LogSyslogd4OverrideSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4OverrideSettingMapInput)(nil)).Elem(), LogSyslogd4OverrideSettingMap{})
	pulumi.RegisterOutputType(LogSyslogd4OverrideSettingOutput{})
	pulumi.RegisterOutputType(LogSyslogd4OverrideSettingArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd4OverrideSettingMapOutput{})
}
