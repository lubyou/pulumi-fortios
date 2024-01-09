// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogFortianalyzer2OverrideSetting struct {
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
	Serials                   LogFortianalyzer2OverrideSettingSerialArrayOutput `pulumi:"serials"`
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

// NewLogFortianalyzer2OverrideSetting registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer2OverrideSetting(ctx *pulumi.Context,
	name string, args *LogFortianalyzer2OverrideSettingArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer2OverrideSetting, error) {
	if args == nil {
		args = &LogFortianalyzer2OverrideSettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	Serials                   []LogFortianalyzer2OverrideSettingSerial `pulumi:"serials"`
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

type LogFortianalyzer2OverrideSettingState struct {
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
	Serials                   LogFortianalyzer2OverrideSettingSerialArrayInput
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

func (LogFortianalyzer2OverrideSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer2OverrideSettingState)(nil)).Elem()
}

type logFortianalyzer2OverrideSettingArgs struct {
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
	Serials                   []LogFortianalyzer2OverrideSettingSerial `pulumi:"serials"`
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

// The set of arguments for constructing a LogFortianalyzer2OverrideSetting resource.
type LogFortianalyzer2OverrideSettingArgs struct {
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
	Serials                   LogFortianalyzer2OverrideSettingSerialArrayInput
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

func (LogFortianalyzer2OverrideSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer2OverrideSettingArgs)(nil)).Elem()
}

type LogFortianalyzer2OverrideSettingInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideSettingOutput() LogFortianalyzer2OverrideSettingOutput
	ToLogFortianalyzer2OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingOutput
}

func (*LogFortianalyzer2OverrideSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer2OverrideSetting)(nil)).Elem()
}

func (i *LogFortianalyzer2OverrideSetting) ToLogFortianalyzer2OverrideSettingOutput() LogFortianalyzer2OverrideSettingOutput {
	return i.ToLogFortianalyzer2OverrideSettingOutputWithContext(context.Background())
}

func (i *LogFortianalyzer2OverrideSetting) ToLogFortianalyzer2OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideSettingOutput)
}

// LogFortianalyzer2OverrideSettingArrayInput is an input type that accepts LogFortianalyzer2OverrideSettingArray and LogFortianalyzer2OverrideSettingArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzer2OverrideSettingArrayInput` via:
//
//	LogFortianalyzer2OverrideSettingArray{ LogFortianalyzer2OverrideSettingArgs{...} }
type LogFortianalyzer2OverrideSettingArrayInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideSettingArrayOutput() LogFortianalyzer2OverrideSettingArrayOutput
	ToLogFortianalyzer2OverrideSettingArrayOutputWithContext(context.Context) LogFortianalyzer2OverrideSettingArrayOutput
}

type LogFortianalyzer2OverrideSettingArray []LogFortianalyzer2OverrideSettingInput

func (LogFortianalyzer2OverrideSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer2OverrideSetting)(nil)).Elem()
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
//	LogFortianalyzer2OverrideSettingMap{ "key": LogFortianalyzer2OverrideSettingArgs{...} }
type LogFortianalyzer2OverrideSettingMapInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideSettingMapOutput() LogFortianalyzer2OverrideSettingMapOutput
	ToLogFortianalyzer2OverrideSettingMapOutputWithContext(context.Context) LogFortianalyzer2OverrideSettingMapOutput
}

type LogFortianalyzer2OverrideSettingMap map[string]LogFortianalyzer2OverrideSettingInput

func (LogFortianalyzer2OverrideSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer2OverrideSetting)(nil)).Elem()
}

func (i LogFortianalyzer2OverrideSettingMap) ToLogFortianalyzer2OverrideSettingMapOutput() LogFortianalyzer2OverrideSettingMapOutput {
	return i.ToLogFortianalyzer2OverrideSettingMapOutputWithContext(context.Background())
}

func (i LogFortianalyzer2OverrideSettingMap) ToLogFortianalyzer2OverrideSettingMapOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideSettingMapOutput)
}

type LogFortianalyzer2OverrideSettingOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer2OverrideSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer2OverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzer2OverrideSettingOutput) ToLogFortianalyzer2OverrideSettingOutput() LogFortianalyzer2OverrideSettingOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingOutput) ToLogFortianalyzer2OverrideSettingOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingOutput) __changeIp() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.IntOutput { return v.__changeIp }).(pulumi.IntOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) AccessConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.AccessConfig }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) AltServer() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.AltServer }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) CertificateVerification() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.CertificateVerification }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) ConnTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.IntOutput { return v.ConnTimeout }).(pulumi.IntOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) FallbackToPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.FallbackToPrimary }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) FazType() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.IntOutput { return v.FazType }).(pulumi.IntOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) HmacAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.HmacAlgorithm }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) IpsArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.IpsArchive }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) MaxLogRate() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.IntOutput { return v.MaxLogRate }).(pulumi.IntOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) MgmtName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.MgmtName }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) MonitorFailureRetryPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.IntOutput { return v.MonitorFailureRetryPeriod }).(pulumi.IntOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) MonitorKeepalivePeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.IntOutput { return v.MonitorKeepalivePeriod }).(pulumi.IntOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Override() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.Override }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) PresharedKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.PresharedKey }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Reliable() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.Reliable }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Serials() LogFortianalyzer2OverrideSettingSerialArrayOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) LogFortianalyzer2OverrideSettingSerialArrayOutput {
		return v.Serials
	}).(LogFortianalyzer2OverrideSettingSerialArrayOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) UploadDay() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.UploadDay }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) UploadInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.UploadInterval }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) UploadOption() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.UploadOption }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) UploadTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.UploadTime }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) UseManagementVdom() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringOutput { return v.UseManagementVdom }).(pulumi.StringOutput)
}

func (o LogFortianalyzer2OverrideSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer2OverrideSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LogFortianalyzer2OverrideSettingArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer2OverrideSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer2OverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzer2OverrideSettingArrayOutput) ToLogFortianalyzer2OverrideSettingArrayOutput() LogFortianalyzer2OverrideSettingArrayOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingArrayOutput) ToLogFortianalyzer2OverrideSettingArrayOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingArrayOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingArrayOutput) Index(i pulumi.IntInput) LogFortianalyzer2OverrideSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzer2OverrideSetting {
		return vs[0].([]*LogFortianalyzer2OverrideSetting)[vs[1].(int)]
	}).(LogFortianalyzer2OverrideSettingOutput)
}

type LogFortianalyzer2OverrideSettingMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer2OverrideSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer2OverrideSetting)(nil)).Elem()
}

func (o LogFortianalyzer2OverrideSettingMapOutput) ToLogFortianalyzer2OverrideSettingMapOutput() LogFortianalyzer2OverrideSettingMapOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingMapOutput) ToLogFortianalyzer2OverrideSettingMapOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideSettingMapOutput {
	return o
}

func (o LogFortianalyzer2OverrideSettingMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzer2OverrideSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzer2OverrideSetting {
		return vs[0].(map[string]*LogFortianalyzer2OverrideSetting)[vs[1].(string)]
	}).(LogFortianalyzer2OverrideSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer2OverrideSettingInput)(nil)).Elem(), &LogFortianalyzer2OverrideSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer2OverrideSettingArrayInput)(nil)).Elem(), LogFortianalyzer2OverrideSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer2OverrideSettingMapInput)(nil)).Elem(), LogFortianalyzer2OverrideSettingMap{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideSettingOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideSettingArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideSettingMapOutput{})
}
