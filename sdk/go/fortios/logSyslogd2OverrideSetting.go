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
// LogSyslogd2 OverrideSetting can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd2OverrideSetting:LogSyslogd2OverrideSetting labelname LogSyslogd2OverrideSetting
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd2OverrideSetting:LogSyslogd2OverrideSetting labelname LogSyslogd2OverrideSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogd2OverrideSetting struct {
	pulumi.CustomResourceState

	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd2OverrideSettingCustomFieldNameArrayOutput `pulumi:"customFieldNames"`
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

// NewLogSyslogd2OverrideSetting registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd2OverrideSetting(ctx *pulumi.Context,
	name string, args *LogSyslogd2OverrideSettingArgs, opts ...pulumi.ResourceOption) (*LogSyslogd2OverrideSetting, error) {
	if args == nil {
		args = &LogSyslogd2OverrideSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogd2OverrideSetting
	err := ctx.RegisterResource("fortios:index/logSyslogd2OverrideSetting:LogSyslogd2OverrideSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd2OverrideSetting gets an existing LogSyslogd2OverrideSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd2OverrideSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd2OverrideSettingState, opts ...pulumi.ResourceOption) (*LogSyslogd2OverrideSetting, error) {
	var resource LogSyslogd2OverrideSetting
	err := ctx.ReadResource("fortios:index/logSyslogd2OverrideSetting:LogSyslogd2OverrideSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd2OverrideSetting resources.
type logSyslogd2OverrideSettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []LogSyslogd2OverrideSettingCustomFieldName `pulumi:"customFieldNames"`
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

type LogSyslogd2OverrideSettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd2OverrideSettingCustomFieldNameArrayInput
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

func (LogSyslogd2OverrideSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd2OverrideSettingState)(nil)).Elem()
}

type logSyslogd2OverrideSettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []LogSyslogd2OverrideSettingCustomFieldName `pulumi:"customFieldNames"`
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

// The set of arguments for constructing a LogSyslogd2OverrideSetting resource.
type LogSyslogd2OverrideSettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames LogSyslogd2OverrideSettingCustomFieldNameArrayInput
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

func (LogSyslogd2OverrideSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd2OverrideSettingArgs)(nil)).Elem()
}

type LogSyslogd2OverrideSettingInput interface {
	pulumi.Input

	ToLogSyslogd2OverrideSettingOutput() LogSyslogd2OverrideSettingOutput
	ToLogSyslogd2OverrideSettingOutputWithContext(ctx context.Context) LogSyslogd2OverrideSettingOutput
}

func (*LogSyslogd2OverrideSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd2OverrideSetting)(nil)).Elem()
}

func (i *LogSyslogd2OverrideSetting) ToLogSyslogd2OverrideSettingOutput() LogSyslogd2OverrideSettingOutput {
	return i.ToLogSyslogd2OverrideSettingOutputWithContext(context.Background())
}

func (i *LogSyslogd2OverrideSetting) ToLogSyslogd2OverrideSettingOutputWithContext(ctx context.Context) LogSyslogd2OverrideSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2OverrideSettingOutput)
}

// LogSyslogd2OverrideSettingArrayInput is an input type that accepts LogSyslogd2OverrideSettingArray and LogSyslogd2OverrideSettingArrayOutput values.
// You can construct a concrete instance of `LogSyslogd2OverrideSettingArrayInput` via:
//
//          LogSyslogd2OverrideSettingArray{ LogSyslogd2OverrideSettingArgs{...} }
type LogSyslogd2OverrideSettingArrayInput interface {
	pulumi.Input

	ToLogSyslogd2OverrideSettingArrayOutput() LogSyslogd2OverrideSettingArrayOutput
	ToLogSyslogd2OverrideSettingArrayOutputWithContext(context.Context) LogSyslogd2OverrideSettingArrayOutput
}

type LogSyslogd2OverrideSettingArray []LogSyslogd2OverrideSettingInput

func (LogSyslogd2OverrideSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd2OverrideSetting)(nil)).Elem()
}

func (i LogSyslogd2OverrideSettingArray) ToLogSyslogd2OverrideSettingArrayOutput() LogSyslogd2OverrideSettingArrayOutput {
	return i.ToLogSyslogd2OverrideSettingArrayOutputWithContext(context.Background())
}

func (i LogSyslogd2OverrideSettingArray) ToLogSyslogd2OverrideSettingArrayOutputWithContext(ctx context.Context) LogSyslogd2OverrideSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2OverrideSettingArrayOutput)
}

// LogSyslogd2OverrideSettingMapInput is an input type that accepts LogSyslogd2OverrideSettingMap and LogSyslogd2OverrideSettingMapOutput values.
// You can construct a concrete instance of `LogSyslogd2OverrideSettingMapInput` via:
//
//          LogSyslogd2OverrideSettingMap{ "key": LogSyslogd2OverrideSettingArgs{...} }
type LogSyslogd2OverrideSettingMapInput interface {
	pulumi.Input

	ToLogSyslogd2OverrideSettingMapOutput() LogSyslogd2OverrideSettingMapOutput
	ToLogSyslogd2OverrideSettingMapOutputWithContext(context.Context) LogSyslogd2OverrideSettingMapOutput
}

type LogSyslogd2OverrideSettingMap map[string]LogSyslogd2OverrideSettingInput

func (LogSyslogd2OverrideSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd2OverrideSetting)(nil)).Elem()
}

func (i LogSyslogd2OverrideSettingMap) ToLogSyslogd2OverrideSettingMapOutput() LogSyslogd2OverrideSettingMapOutput {
	return i.ToLogSyslogd2OverrideSettingMapOutputWithContext(context.Background())
}

func (i LogSyslogd2OverrideSettingMap) ToLogSyslogd2OverrideSettingMapOutputWithContext(ctx context.Context) LogSyslogd2OverrideSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2OverrideSettingMapOutput)
}

type LogSyslogd2OverrideSettingOutput struct{ *pulumi.OutputState }

func (LogSyslogd2OverrideSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd2OverrideSetting)(nil)).Elem()
}

func (o LogSyslogd2OverrideSettingOutput) ToLogSyslogd2OverrideSettingOutput() LogSyslogd2OverrideSettingOutput {
	return o
}

func (o LogSyslogd2OverrideSettingOutput) ToLogSyslogd2OverrideSettingOutputWithContext(ctx context.Context) LogSyslogd2OverrideSettingOutput {
	return o
}

type LogSyslogd2OverrideSettingArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd2OverrideSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd2OverrideSetting)(nil)).Elem()
}

func (o LogSyslogd2OverrideSettingArrayOutput) ToLogSyslogd2OverrideSettingArrayOutput() LogSyslogd2OverrideSettingArrayOutput {
	return o
}

func (o LogSyslogd2OverrideSettingArrayOutput) ToLogSyslogd2OverrideSettingArrayOutputWithContext(ctx context.Context) LogSyslogd2OverrideSettingArrayOutput {
	return o
}

func (o LogSyslogd2OverrideSettingArrayOutput) Index(i pulumi.IntInput) LogSyslogd2OverrideSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogd2OverrideSetting {
		return vs[0].([]*LogSyslogd2OverrideSetting)[vs[1].(int)]
	}).(LogSyslogd2OverrideSettingOutput)
}

type LogSyslogd2OverrideSettingMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd2OverrideSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd2OverrideSetting)(nil)).Elem()
}

func (o LogSyslogd2OverrideSettingMapOutput) ToLogSyslogd2OverrideSettingMapOutput() LogSyslogd2OverrideSettingMapOutput {
	return o
}

func (o LogSyslogd2OverrideSettingMapOutput) ToLogSyslogd2OverrideSettingMapOutputWithContext(ctx context.Context) LogSyslogd2OverrideSettingMapOutput {
	return o
}

func (o LogSyslogd2OverrideSettingMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd2OverrideSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogd2OverrideSetting {
		return vs[0].(map[string]*LogSyslogd2OverrideSetting)[vs[1].(string)]
	}).(LogSyslogd2OverrideSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2OverrideSettingInput)(nil)).Elem(), &LogSyslogd2OverrideSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2OverrideSettingArrayInput)(nil)).Elem(), LogSyslogd2OverrideSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2OverrideSettingMapInput)(nil)).Elem(), LogSyslogd2OverrideSettingMap{})
	pulumi.RegisterOutputType(LogSyslogd2OverrideSettingOutput{})
	pulumi.RegisterOutputType(LogSyslogd2OverrideSettingArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd2OverrideSettingMapOutput{})
}
