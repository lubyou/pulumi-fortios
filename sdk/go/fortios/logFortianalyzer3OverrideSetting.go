// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogFortianalyzer3OverrideSetting struct {
	pulumi.CustomResourceState

	__changeIp                pulumi.IntOutput                                  `pulumi:"__changeIp"`
	AccessConfig              pulumi.StringOutput                               `pulumi:"accessConfig"`
	AltServer                 pulumi.StringOutput                               `pulumi:"altServer"`
	Certificate               pulumi.StringOutput                               `pulumi:"certificate"`
	CertificateVerification   pulumi.StringOutput                               `pulumi:"certificateVerification"`
	ConnTimeout               pulumi.IntOutput                                  `pulumi:"connTimeout"`
	DynamicSortSubtable       pulumi.StringPtrOutput                            `pulumi:"dynamicSortSubtable"`
	EncAlgorithm              pulumi.StringOutput                               `pulumi:"encAlgorithm"`
	FallbackToPrimary         pulumi.StringOutput                               `pulumi:"fallbackToPrimary"`
	FazType                   pulumi.IntOutput                                  `pulumi:"fazType"`
	GetAllTables              pulumi.StringPtrOutput                            `pulumi:"getAllTables"`
	HmacAlgorithm             pulumi.StringOutput                               `pulumi:"hmacAlgorithm"`
	Interface                 pulumi.StringOutput                               `pulumi:"interface"`
	InterfaceSelectMethod     pulumi.StringOutput                               `pulumi:"interfaceSelectMethod"`
	IpsArchive                pulumi.StringOutput                               `pulumi:"ipsArchive"`
	MaxLogRate                pulumi.IntOutput                                  `pulumi:"maxLogRate"`
	MgmtName                  pulumi.StringOutput                               `pulumi:"mgmtName"`
	MonitorFailureRetryPeriod pulumi.IntOutput                                  `pulumi:"monitorFailureRetryPeriod"`
	MonitorKeepalivePeriod    pulumi.IntOutput                                  `pulumi:"monitorKeepalivePeriod"`
	Override                  pulumi.StringOutput                               `pulumi:"override"`
	PresharedKey              pulumi.StringOutput                               `pulumi:"presharedKey"`
	Priority                  pulumi.StringOutput                               `pulumi:"priority"`
	Reliable                  pulumi.StringOutput                               `pulumi:"reliable"`
	Serials                   LogFortianalyzer3OverrideSettingSerialArrayOutput `pulumi:"serials"`
	Server                    pulumi.StringOutput                               `pulumi:"server"`
	SourceIp                  pulumi.StringOutput                               `pulumi:"sourceIp"`
	SslMinProtoVersion        pulumi.StringOutput                               `pulumi:"sslMinProtoVersion"`
	Status                    pulumi.StringOutput                               `pulumi:"status"`
	UploadDay                 pulumi.StringOutput                               `pulumi:"uploadDay"`
	UploadInterval            pulumi.StringOutput                               `pulumi:"uploadInterval"`
	UploadOption              pulumi.StringOutput                               `pulumi:"uploadOption"`
	UploadTime                pulumi.StringOutput                               `pulumi:"uploadTime"`
	UseManagementVdom         pulumi.StringOutput                               `pulumi:"useManagementVdom"`
	Vdomparam                 pulumi.StringPtrOutput                            `pulumi:"vdomparam"`
}

// NewLogFortianalyzer3OverrideSetting registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer3OverrideSetting(ctx *pulumi.Context,
	name string, args *LogFortianalyzer3OverrideSettingArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer3OverrideSetting, error) {
	if args == nil {
		args = &LogFortianalyzer3OverrideSettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	__changeIp                *int                                     `pulumi:"__changeIp"`
	AccessConfig              *string                                  `pulumi:"accessConfig"`
	AltServer                 *string                                  `pulumi:"altServer"`
	Certificate               *string                                  `pulumi:"certificate"`
	CertificateVerification   *string                                  `pulumi:"certificateVerification"`
	ConnTimeout               *int                                     `pulumi:"connTimeout"`
	DynamicSortSubtable       *string                                  `pulumi:"dynamicSortSubtable"`
	EncAlgorithm              *string                                  `pulumi:"encAlgorithm"`
	FallbackToPrimary         *string                                  `pulumi:"fallbackToPrimary"`
	FazType                   *int                                     `pulumi:"fazType"`
	GetAllTables              *string                                  `pulumi:"getAllTables"`
	HmacAlgorithm             *string                                  `pulumi:"hmacAlgorithm"`
	Interface                 *string                                  `pulumi:"interface"`
	InterfaceSelectMethod     *string                                  `pulumi:"interfaceSelectMethod"`
	IpsArchive                *string                                  `pulumi:"ipsArchive"`
	MaxLogRate                *int                                     `pulumi:"maxLogRate"`
	MgmtName                  *string                                  `pulumi:"mgmtName"`
	MonitorFailureRetryPeriod *int                                     `pulumi:"monitorFailureRetryPeriod"`
	MonitorKeepalivePeriod    *int                                     `pulumi:"monitorKeepalivePeriod"`
	Override                  *string                                  `pulumi:"override"`
	PresharedKey              *string                                  `pulumi:"presharedKey"`
	Priority                  *string                                  `pulumi:"priority"`
	Reliable                  *string                                  `pulumi:"reliable"`
	Serials                   []LogFortianalyzer3OverrideSettingSerial `pulumi:"serials"`
	Server                    *string                                  `pulumi:"server"`
	SourceIp                  *string                                  `pulumi:"sourceIp"`
	SslMinProtoVersion        *string                                  `pulumi:"sslMinProtoVersion"`
	Status                    *string                                  `pulumi:"status"`
	UploadDay                 *string                                  `pulumi:"uploadDay"`
	UploadInterval            *string                                  `pulumi:"uploadInterval"`
	UploadOption              *string                                  `pulumi:"uploadOption"`
	UploadTime                *string                                  `pulumi:"uploadTime"`
	UseManagementVdom         *string                                  `pulumi:"useManagementVdom"`
	Vdomparam                 *string                                  `pulumi:"vdomparam"`
}

type LogFortianalyzer3OverrideSettingState struct {
	__changeIp                pulumi.IntPtrInput
	AccessConfig              pulumi.StringPtrInput
	AltServer                 pulumi.StringPtrInput
	Certificate               pulumi.StringPtrInput
	CertificateVerification   pulumi.StringPtrInput
	ConnTimeout               pulumi.IntPtrInput
	DynamicSortSubtable       pulumi.StringPtrInput
	EncAlgorithm              pulumi.StringPtrInput
	FallbackToPrimary         pulumi.StringPtrInput
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
	Override                  pulumi.StringPtrInput
	PresharedKey              pulumi.StringPtrInput
	Priority                  pulumi.StringPtrInput
	Reliable                  pulumi.StringPtrInput
	Serials                   LogFortianalyzer3OverrideSettingSerialArrayInput
	Server                    pulumi.StringPtrInput
	SourceIp                  pulumi.StringPtrInput
	SslMinProtoVersion        pulumi.StringPtrInput
	Status                    pulumi.StringPtrInput
	UploadDay                 pulumi.StringPtrInput
	UploadInterval            pulumi.StringPtrInput
	UploadOption              pulumi.StringPtrInput
	UploadTime                pulumi.StringPtrInput
	UseManagementVdom         pulumi.StringPtrInput
	Vdomparam                 pulumi.StringPtrInput
}

func (LogFortianalyzer3OverrideSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3OverrideSettingState)(nil)).Elem()
}

type logFortianalyzer3OverrideSettingArgs struct {
	__changeIp                *int                                     `pulumi:"__changeIp"`
	AccessConfig              *string                                  `pulumi:"accessConfig"`
	AltServer                 *string                                  `pulumi:"altServer"`
	Certificate               *string                                  `pulumi:"certificate"`
	CertificateVerification   *string                                  `pulumi:"certificateVerification"`
	ConnTimeout               *int                                     `pulumi:"connTimeout"`
	DynamicSortSubtable       *string                                  `pulumi:"dynamicSortSubtable"`
	EncAlgorithm              *string                                  `pulumi:"encAlgorithm"`
	FallbackToPrimary         *string                                  `pulumi:"fallbackToPrimary"`
	FazType                   *int                                     `pulumi:"fazType"`
	GetAllTables              *string                                  `pulumi:"getAllTables"`
	HmacAlgorithm             *string                                  `pulumi:"hmacAlgorithm"`
	Interface                 *string                                  `pulumi:"interface"`
	InterfaceSelectMethod     *string                                  `pulumi:"interfaceSelectMethod"`
	IpsArchive                *string                                  `pulumi:"ipsArchive"`
	MaxLogRate                *int                                     `pulumi:"maxLogRate"`
	MgmtName                  *string                                  `pulumi:"mgmtName"`
	MonitorFailureRetryPeriod *int                                     `pulumi:"monitorFailureRetryPeriod"`
	MonitorKeepalivePeriod    *int                                     `pulumi:"monitorKeepalivePeriod"`
	Override                  *string                                  `pulumi:"override"`
	PresharedKey              *string                                  `pulumi:"presharedKey"`
	Priority                  *string                                  `pulumi:"priority"`
	Reliable                  *string                                  `pulumi:"reliable"`
	Serials                   []LogFortianalyzer3OverrideSettingSerial `pulumi:"serials"`
	Server                    *string                                  `pulumi:"server"`
	SourceIp                  *string                                  `pulumi:"sourceIp"`
	SslMinProtoVersion        *string                                  `pulumi:"sslMinProtoVersion"`
	Status                    *string                                  `pulumi:"status"`
	UploadDay                 *string                                  `pulumi:"uploadDay"`
	UploadInterval            *string                                  `pulumi:"uploadInterval"`
	UploadOption              *string                                  `pulumi:"uploadOption"`
	UploadTime                *string                                  `pulumi:"uploadTime"`
	UseManagementVdom         *string                                  `pulumi:"useManagementVdom"`
	Vdomparam                 *string                                  `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogFortianalyzer3OverrideSetting resource.
type LogFortianalyzer3OverrideSettingArgs struct {
	__changeIp                pulumi.IntPtrInput
	AccessConfig              pulumi.StringPtrInput
	AltServer                 pulumi.StringPtrInput
	Certificate               pulumi.StringPtrInput
	CertificateVerification   pulumi.StringPtrInput
	ConnTimeout               pulumi.IntPtrInput
	DynamicSortSubtable       pulumi.StringPtrInput
	EncAlgorithm              pulumi.StringPtrInput
	FallbackToPrimary         pulumi.StringPtrInput
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
	Override                  pulumi.StringPtrInput
	PresharedKey              pulumi.StringPtrInput
	Priority                  pulumi.StringPtrInput
	Reliable                  pulumi.StringPtrInput
	Serials                   LogFortianalyzer3OverrideSettingSerialArrayInput
	Server                    pulumi.StringPtrInput
	SourceIp                  pulumi.StringPtrInput
	SslMinProtoVersion        pulumi.StringPtrInput
	Status                    pulumi.StringPtrInput
	UploadDay                 pulumi.StringPtrInput
	UploadInterval            pulumi.StringPtrInput
	UploadOption              pulumi.StringPtrInput
	UploadTime                pulumi.StringPtrInput
	UseManagementVdom         pulumi.StringPtrInput
	Vdomparam                 pulumi.StringPtrInput
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
//	LogFortianalyzer3OverrideSettingArray{ LogFortianalyzer3OverrideSettingArgs{...} }
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
//	LogFortianalyzer3OverrideSettingMap{ "key": LogFortianalyzer3OverrideSettingArgs{...} }
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

func (o LogFortianalyzer3OverrideSettingOutput) __changeIp() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.IntOutput { return v.__changeIp }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) AccessConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.AccessConfig }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) AltServer() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.AltServer }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) CertificateVerification() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.CertificateVerification }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) ConnTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.IntOutput { return v.ConnTimeout }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) FallbackToPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.FallbackToPrimary }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) FazType() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.IntOutput { return v.FazType }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) HmacAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.HmacAlgorithm }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) IpsArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.IpsArchive }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) MaxLogRate() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.IntOutput { return v.MaxLogRate }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) MgmtName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.MgmtName }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) MonitorFailureRetryPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.IntOutput { return v.MonitorFailureRetryPeriod }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) MonitorKeepalivePeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.IntOutput { return v.MonitorKeepalivePeriod }).(pulumi.IntOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Override() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.Override }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) PresharedKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.PresharedKey }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Reliable() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.Reliable }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Serials() LogFortianalyzer3OverrideSettingSerialArrayOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) LogFortianalyzer3OverrideSettingSerialArrayOutput {
		return v.Serials
	}).(LogFortianalyzer3OverrideSettingSerialArrayOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) UploadDay() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.UploadDay }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) UploadInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.UploadInterval }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) UploadOption() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.UploadOption }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) UploadTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.UploadTime }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) UseManagementVdom() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringOutput { return v.UseManagementVdom }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3OverrideSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3OverrideSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
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
