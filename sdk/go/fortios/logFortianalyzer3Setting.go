// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogFortianalyzer3Setting struct {
	pulumi.CustomResourceState

	__changeIp                pulumi.IntOutput                          `pulumi:"__changeIp"`
	AccessConfig              pulumi.StringOutput                       `pulumi:"accessConfig"`
	Certificate               pulumi.StringOutput                       `pulumi:"certificate"`
	CertificateVerification   pulumi.StringOutput                       `pulumi:"certificateVerification"`
	ConnTimeout               pulumi.IntOutput                          `pulumi:"connTimeout"`
	DynamicSortSubtable       pulumi.StringPtrOutput                    `pulumi:"dynamicSortSubtable"`
	EncAlgorithm              pulumi.StringOutput                       `pulumi:"encAlgorithm"`
	FazType                   pulumi.IntOutput                          `pulumi:"fazType"`
	GetAllTables              pulumi.StringPtrOutput                    `pulumi:"getAllTables"`
	HmacAlgorithm             pulumi.StringOutput                       `pulumi:"hmacAlgorithm"`
	Interface                 pulumi.StringOutput                       `pulumi:"interface"`
	InterfaceSelectMethod     pulumi.StringOutput                       `pulumi:"interfaceSelectMethod"`
	IpsArchive                pulumi.StringOutput                       `pulumi:"ipsArchive"`
	MaxLogRate                pulumi.IntOutput                          `pulumi:"maxLogRate"`
	MgmtName                  pulumi.StringOutput                       `pulumi:"mgmtName"`
	MonitorFailureRetryPeriod pulumi.IntOutput                          `pulumi:"monitorFailureRetryPeriod"`
	MonitorKeepalivePeriod    pulumi.IntOutput                          `pulumi:"monitorKeepalivePeriod"`
	PresharedKey              pulumi.StringOutput                       `pulumi:"presharedKey"`
	Priority                  pulumi.StringOutput                       `pulumi:"priority"`
	Reliable                  pulumi.StringOutput                       `pulumi:"reliable"`
	Serials                   LogFortianalyzer3SettingSerialArrayOutput `pulumi:"serials"`
	Server                    pulumi.StringOutput                       `pulumi:"server"`
	SourceIp                  pulumi.StringOutput                       `pulumi:"sourceIp"`
	SslMinProtoVersion        pulumi.StringOutput                       `pulumi:"sslMinProtoVersion"`
	Status                    pulumi.StringOutput                       `pulumi:"status"`
	UploadDay                 pulumi.StringOutput                       `pulumi:"uploadDay"`
	UploadInterval            pulumi.StringOutput                       `pulumi:"uploadInterval"`
	UploadOption              pulumi.StringOutput                       `pulumi:"uploadOption"`
	UploadTime                pulumi.StringOutput                       `pulumi:"uploadTime"`
	Vdomparam                 pulumi.StringPtrOutput                    `pulumi:"vdomparam"`
}

// NewLogFortianalyzer3Setting registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer3Setting(ctx *pulumi.Context,
	name string, args *LogFortianalyzer3SettingArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer3Setting, error) {
	if args == nil {
		args = &LogFortianalyzer3SettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogFortianalyzer3Setting
	err := ctx.RegisterResource("fortios:index/logFortianalyzer3Setting:LogFortianalyzer3Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzer3Setting gets an existing LogFortianalyzer3Setting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzer3Setting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzer3SettingState, opts ...pulumi.ResourceOption) (*LogFortianalyzer3Setting, error) {
	var resource LogFortianalyzer3Setting
	err := ctx.ReadResource("fortios:index/logFortianalyzer3Setting:LogFortianalyzer3Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzer3Setting resources.
type logFortianalyzer3SettingState struct {
	__changeIp                *int                             `pulumi:"__changeIp"`
	AccessConfig              *string                          `pulumi:"accessConfig"`
	Certificate               *string                          `pulumi:"certificate"`
	CertificateVerification   *string                          `pulumi:"certificateVerification"`
	ConnTimeout               *int                             `pulumi:"connTimeout"`
	DynamicSortSubtable       *string                          `pulumi:"dynamicSortSubtable"`
	EncAlgorithm              *string                          `pulumi:"encAlgorithm"`
	FazType                   *int                             `pulumi:"fazType"`
	GetAllTables              *string                          `pulumi:"getAllTables"`
	HmacAlgorithm             *string                          `pulumi:"hmacAlgorithm"`
	Interface                 *string                          `pulumi:"interface"`
	InterfaceSelectMethod     *string                          `pulumi:"interfaceSelectMethod"`
	IpsArchive                *string                          `pulumi:"ipsArchive"`
	MaxLogRate                *int                             `pulumi:"maxLogRate"`
	MgmtName                  *string                          `pulumi:"mgmtName"`
	MonitorFailureRetryPeriod *int                             `pulumi:"monitorFailureRetryPeriod"`
	MonitorKeepalivePeriod    *int                             `pulumi:"monitorKeepalivePeriod"`
	PresharedKey              *string                          `pulumi:"presharedKey"`
	Priority                  *string                          `pulumi:"priority"`
	Reliable                  *string                          `pulumi:"reliable"`
	Serials                   []LogFortianalyzer3SettingSerial `pulumi:"serials"`
	Server                    *string                          `pulumi:"server"`
	SourceIp                  *string                          `pulumi:"sourceIp"`
	SslMinProtoVersion        *string                          `pulumi:"sslMinProtoVersion"`
	Status                    *string                          `pulumi:"status"`
	UploadDay                 *string                          `pulumi:"uploadDay"`
	UploadInterval            *string                          `pulumi:"uploadInterval"`
	UploadOption              *string                          `pulumi:"uploadOption"`
	UploadTime                *string                          `pulumi:"uploadTime"`
	Vdomparam                 *string                          `pulumi:"vdomparam"`
}

type LogFortianalyzer3SettingState struct {
	__changeIp                pulumi.IntPtrInput
	AccessConfig              pulumi.StringPtrInput
	Certificate               pulumi.StringPtrInput
	CertificateVerification   pulumi.StringPtrInput
	ConnTimeout               pulumi.IntPtrInput
	DynamicSortSubtable       pulumi.StringPtrInput
	EncAlgorithm              pulumi.StringPtrInput
	FazType                   pulumi.IntPtrInput
	GetAllTables              pulumi.StringPtrInput
	HmacAlgorithm             pulumi.StringPtrInput
	Interface                 pulumi.StringPtrInput
	InterfaceSelectMethod     pulumi.StringPtrInput
	IpsArchive                pulumi.StringPtrInput
	MaxLogRate                pulumi.IntPtrInput
	MgmtName                  pulumi.StringPtrInput
	MonitorFailureRetryPeriod pulumi.IntPtrInput
	MonitorKeepalivePeriod    pulumi.IntPtrInput
	PresharedKey              pulumi.StringPtrInput
	Priority                  pulumi.StringPtrInput
	Reliable                  pulumi.StringPtrInput
	Serials                   LogFortianalyzer3SettingSerialArrayInput
	Server                    pulumi.StringPtrInput
	SourceIp                  pulumi.StringPtrInput
	SslMinProtoVersion        pulumi.StringPtrInput
	Status                    pulumi.StringPtrInput
	UploadDay                 pulumi.StringPtrInput
	UploadInterval            pulumi.StringPtrInput
	UploadOption              pulumi.StringPtrInput
	UploadTime                pulumi.StringPtrInput
	Vdomparam                 pulumi.StringPtrInput
}

func (LogFortianalyzer3SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3SettingState)(nil)).Elem()
}

type logFortianalyzer3SettingArgs struct {
	__changeIp                *int                             `pulumi:"__changeIp"`
	AccessConfig              *string                          `pulumi:"accessConfig"`
	Certificate               *string                          `pulumi:"certificate"`
	CertificateVerification   *string                          `pulumi:"certificateVerification"`
	ConnTimeout               *int                             `pulumi:"connTimeout"`
	DynamicSortSubtable       *string                          `pulumi:"dynamicSortSubtable"`
	EncAlgorithm              *string                          `pulumi:"encAlgorithm"`
	FazType                   *int                             `pulumi:"fazType"`
	GetAllTables              *string                          `pulumi:"getAllTables"`
	HmacAlgorithm             *string                          `pulumi:"hmacAlgorithm"`
	Interface                 *string                          `pulumi:"interface"`
	InterfaceSelectMethod     *string                          `pulumi:"interfaceSelectMethod"`
	IpsArchive                *string                          `pulumi:"ipsArchive"`
	MaxLogRate                *int                             `pulumi:"maxLogRate"`
	MgmtName                  *string                          `pulumi:"mgmtName"`
	MonitorFailureRetryPeriod *int                             `pulumi:"monitorFailureRetryPeriod"`
	MonitorKeepalivePeriod    *int                             `pulumi:"monitorKeepalivePeriod"`
	PresharedKey              *string                          `pulumi:"presharedKey"`
	Priority                  *string                          `pulumi:"priority"`
	Reliable                  *string                          `pulumi:"reliable"`
	Serials                   []LogFortianalyzer3SettingSerial `pulumi:"serials"`
	Server                    *string                          `pulumi:"server"`
	SourceIp                  *string                          `pulumi:"sourceIp"`
	SslMinProtoVersion        *string                          `pulumi:"sslMinProtoVersion"`
	Status                    *string                          `pulumi:"status"`
	UploadDay                 *string                          `pulumi:"uploadDay"`
	UploadInterval            *string                          `pulumi:"uploadInterval"`
	UploadOption              *string                          `pulumi:"uploadOption"`
	UploadTime                *string                          `pulumi:"uploadTime"`
	Vdomparam                 *string                          `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogFortianalyzer3Setting resource.
type LogFortianalyzer3SettingArgs struct {
	__changeIp                pulumi.IntPtrInput
	AccessConfig              pulumi.StringPtrInput
	Certificate               pulumi.StringPtrInput
	CertificateVerification   pulumi.StringPtrInput
	ConnTimeout               pulumi.IntPtrInput
	DynamicSortSubtable       pulumi.StringPtrInput
	EncAlgorithm              pulumi.StringPtrInput
	FazType                   pulumi.IntPtrInput
	GetAllTables              pulumi.StringPtrInput
	HmacAlgorithm             pulumi.StringPtrInput
	Interface                 pulumi.StringPtrInput
	InterfaceSelectMethod     pulumi.StringPtrInput
	IpsArchive                pulumi.StringPtrInput
	MaxLogRate                pulumi.IntPtrInput
	MgmtName                  pulumi.StringPtrInput
	MonitorFailureRetryPeriod pulumi.IntPtrInput
	MonitorKeepalivePeriod    pulumi.IntPtrInput
	PresharedKey              pulumi.StringPtrInput
	Priority                  pulumi.StringPtrInput
	Reliable                  pulumi.StringPtrInput
	Serials                   LogFortianalyzer3SettingSerialArrayInput
	Server                    pulumi.StringPtrInput
	SourceIp                  pulumi.StringPtrInput
	SslMinProtoVersion        pulumi.StringPtrInput
	Status                    pulumi.StringPtrInput
	UploadDay                 pulumi.StringPtrInput
	UploadInterval            pulumi.StringPtrInput
	UploadOption              pulumi.StringPtrInput
	UploadTime                pulumi.StringPtrInput
	Vdomparam                 pulumi.StringPtrInput
}

func (LogFortianalyzer3SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3SettingArgs)(nil)).Elem()
}

type LogFortianalyzer3SettingInput interface {
	pulumi.Input

	ToLogFortianalyzer3SettingOutput() LogFortianalyzer3SettingOutput
	ToLogFortianalyzer3SettingOutputWithContext(ctx context.Context) LogFortianalyzer3SettingOutput
}

func (*LogFortianalyzer3Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3Setting)(nil)).Elem()
}

func (i *LogFortianalyzer3Setting) ToLogFortianalyzer3SettingOutput() LogFortianalyzer3SettingOutput {
	return i.ToLogFortianalyzer3SettingOutputWithContext(context.Background())
}

func (i *LogFortianalyzer3Setting) ToLogFortianalyzer3SettingOutputWithContext(ctx context.Context) LogFortianalyzer3SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3SettingOutput)
}

// LogFortianalyzer3SettingArrayInput is an input type that accepts LogFortianalyzer3SettingArray and LogFortianalyzer3SettingArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzer3SettingArrayInput` via:
//
//	LogFortianalyzer3SettingArray{ LogFortianalyzer3SettingArgs{...} }
type LogFortianalyzer3SettingArrayInput interface {
	pulumi.Input

	ToLogFortianalyzer3SettingArrayOutput() LogFortianalyzer3SettingArrayOutput
	ToLogFortianalyzer3SettingArrayOutputWithContext(context.Context) LogFortianalyzer3SettingArrayOutput
}

type LogFortianalyzer3SettingArray []LogFortianalyzer3SettingInput

func (LogFortianalyzer3SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer3Setting)(nil)).Elem()
}

func (i LogFortianalyzer3SettingArray) ToLogFortianalyzer3SettingArrayOutput() LogFortianalyzer3SettingArrayOutput {
	return i.ToLogFortianalyzer3SettingArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzer3SettingArray) ToLogFortianalyzer3SettingArrayOutputWithContext(ctx context.Context) LogFortianalyzer3SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3SettingArrayOutput)
}

// LogFortianalyzer3SettingMapInput is an input type that accepts LogFortianalyzer3SettingMap and LogFortianalyzer3SettingMapOutput values.
// You can construct a concrete instance of `LogFortianalyzer3SettingMapInput` via:
//
//	LogFortianalyzer3SettingMap{ "key": LogFortianalyzer3SettingArgs{...} }
type LogFortianalyzer3SettingMapInput interface {
	pulumi.Input

	ToLogFortianalyzer3SettingMapOutput() LogFortianalyzer3SettingMapOutput
	ToLogFortianalyzer3SettingMapOutputWithContext(context.Context) LogFortianalyzer3SettingMapOutput
}

type LogFortianalyzer3SettingMap map[string]LogFortianalyzer3SettingInput

func (LogFortianalyzer3SettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer3Setting)(nil)).Elem()
}

func (i LogFortianalyzer3SettingMap) ToLogFortianalyzer3SettingMapOutput() LogFortianalyzer3SettingMapOutput {
	return i.ToLogFortianalyzer3SettingMapOutputWithContext(context.Background())
}

func (i LogFortianalyzer3SettingMap) ToLogFortianalyzer3SettingMapOutputWithContext(ctx context.Context) LogFortianalyzer3SettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3SettingMapOutput)
}

type LogFortianalyzer3SettingOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3Setting)(nil)).Elem()
}

func (o LogFortianalyzer3SettingOutput) ToLogFortianalyzer3SettingOutput() LogFortianalyzer3SettingOutput {
	return o
}

func (o LogFortianalyzer3SettingOutput) ToLogFortianalyzer3SettingOutputWithContext(ctx context.Context) LogFortianalyzer3SettingOutput {
	return o
}

func (o LogFortianalyzer3SettingOutput) __changeIp() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.IntOutput { return v.__changeIp }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3SettingOutput) AccessConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.AccessConfig }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) CertificateVerification() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.CertificateVerification }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) ConnTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.IntOutput { return v.ConnTimeout }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3SettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer3SettingOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) FazType() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.IntOutput { return v.FazType }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3SettingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer3SettingOutput) HmacAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.HmacAlgorithm }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) IpsArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.IpsArchive }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) MaxLogRate() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.IntOutput { return v.MaxLogRate }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3SettingOutput) MgmtName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.MgmtName }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) MonitorFailureRetryPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.IntOutput { return v.MonitorFailureRetryPeriod }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3SettingOutput) MonitorKeepalivePeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.IntOutput { return v.MonitorKeepalivePeriod }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3SettingOutput) PresharedKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.PresharedKey }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) Reliable() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.Reliable }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) Serials() LogFortianalyzer3SettingSerialArrayOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) LogFortianalyzer3SettingSerialArrayOutput { return v.Serials }).(LogFortianalyzer3SettingSerialArrayOutput)
}

func (o LogFortianalyzer3SettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) UploadDay() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.UploadDay }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) UploadInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.UploadInterval }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) UploadOption() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.UploadOption }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) UploadTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringOutput { return v.UploadTime }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3SettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Setting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LogFortianalyzer3SettingArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer3Setting)(nil)).Elem()
}

func (o LogFortianalyzer3SettingArrayOutput) ToLogFortianalyzer3SettingArrayOutput() LogFortianalyzer3SettingArrayOutput {
	return o
}

func (o LogFortianalyzer3SettingArrayOutput) ToLogFortianalyzer3SettingArrayOutputWithContext(ctx context.Context) LogFortianalyzer3SettingArrayOutput {
	return o
}

func (o LogFortianalyzer3SettingArrayOutput) Index(i pulumi.IntInput) LogFortianalyzer3SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzer3Setting {
		return vs[0].([]*LogFortianalyzer3Setting)[vs[1].(int)]
	}).(LogFortianalyzer3SettingOutput)
}

type LogFortianalyzer3SettingMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3SettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer3Setting)(nil)).Elem()
}

func (o LogFortianalyzer3SettingMapOutput) ToLogFortianalyzer3SettingMapOutput() LogFortianalyzer3SettingMapOutput {
	return o
}

func (o LogFortianalyzer3SettingMapOutput) ToLogFortianalyzer3SettingMapOutputWithContext(ctx context.Context) LogFortianalyzer3SettingMapOutput {
	return o
}

func (o LogFortianalyzer3SettingMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzer3SettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzer3Setting {
		return vs[0].(map[string]*LogFortianalyzer3Setting)[vs[1].(string)]
	}).(LogFortianalyzer3SettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3SettingInput)(nil)).Elem(), &LogFortianalyzer3Setting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3SettingArrayInput)(nil)).Elem(), LogFortianalyzer3SettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3SettingMapInput)(nil)).Elem(), LogFortianalyzer3SettingMap{})
	pulumi.RegisterOutputType(LogFortianalyzer3SettingOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3SettingArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3SettingMapOutput{})
}
