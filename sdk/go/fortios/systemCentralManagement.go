// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type SystemCentralManagement struct {
	pulumi.CustomResourceState

	AllowMonitor                    pulumi.StringOutput                          `pulumi:"allowMonitor"`
	AllowPushConfiguration          pulumi.StringOutput                          `pulumi:"allowPushConfiguration"`
	AllowPushFirmware               pulumi.StringOutput                          `pulumi:"allowPushFirmware"`
	AllowRemoteFirmwareUpgrade      pulumi.StringOutput                          `pulumi:"allowRemoteFirmwareUpgrade"`
	CaCert                          pulumi.StringOutput                          `pulumi:"caCert"`
	DynamicSortSubtable             pulumi.StringPtrOutput                       `pulumi:"dynamicSortSubtable"`
	EncAlgorithm                    pulumi.StringOutput                          `pulumi:"encAlgorithm"`
	Fmg                             pulumi.StringOutput                          `pulumi:"fmg"`
	FmgSourceIp                     pulumi.StringOutput                          `pulumi:"fmgSourceIp"`
	FmgSourceIp6                    pulumi.StringOutput                          `pulumi:"fmgSourceIp6"`
	FmgUpdatePort                   pulumi.StringOutput                          `pulumi:"fmgUpdatePort"`
	FortigateCloudSsoDefaultProfile pulumi.StringOutput                          `pulumi:"fortigateCloudSsoDefaultProfile"`
	GetAllTables                    pulumi.StringPtrOutput                       `pulumi:"getAllTables"`
	IncludeDefaultServers           pulumi.StringOutput                          `pulumi:"includeDefaultServers"`
	Interface                       pulumi.StringOutput                          `pulumi:"interface"`
	InterfaceSelectMethod           pulumi.StringOutput                          `pulumi:"interfaceSelectMethod"`
	LocalCert                       pulumi.StringOutput                          `pulumi:"localCert"`
	Mode                            pulumi.StringOutput                          `pulumi:"mode"`
	ScheduleConfigRestore           pulumi.StringOutput                          `pulumi:"scheduleConfigRestore"`
	ScheduleScriptRestore           pulumi.StringOutput                          `pulumi:"scheduleScriptRestore"`
	SerialNumber                    pulumi.StringOutput                          `pulumi:"serialNumber"`
	ServerLists                     SystemCentralManagementServerListArrayOutput `pulumi:"serverLists"`
	Type                            pulumi.StringOutput                          `pulumi:"type"`
	Vdom                            pulumi.StringOutput                          `pulumi:"vdom"`
	Vdomparam                       pulumi.StringPtrOutput                       `pulumi:"vdomparam"`
}

// NewSystemCentralManagement registers a new resource with the given unique name, arguments, and options.
func NewSystemCentralManagement(ctx *pulumi.Context,
	name string, args *SystemCentralManagementArgs, opts ...pulumi.ResourceOption) (*SystemCentralManagement, error) {
	if args == nil {
		args = &SystemCentralManagementArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemCentralManagement
	err := ctx.RegisterResource("fortios:index/systemCentralManagement:SystemCentralManagement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemCentralManagement gets an existing SystemCentralManagement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemCentralManagement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemCentralManagementState, opts ...pulumi.ResourceOption) (*SystemCentralManagement, error) {
	var resource SystemCentralManagement
	err := ctx.ReadResource("fortios:index/systemCentralManagement:SystemCentralManagement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemCentralManagement resources.
type systemCentralManagementState struct {
	AllowMonitor                    *string                             `pulumi:"allowMonitor"`
	AllowPushConfiguration          *string                             `pulumi:"allowPushConfiguration"`
	AllowPushFirmware               *string                             `pulumi:"allowPushFirmware"`
	AllowRemoteFirmwareUpgrade      *string                             `pulumi:"allowRemoteFirmwareUpgrade"`
	CaCert                          *string                             `pulumi:"caCert"`
	DynamicSortSubtable             *string                             `pulumi:"dynamicSortSubtable"`
	EncAlgorithm                    *string                             `pulumi:"encAlgorithm"`
	Fmg                             *string                             `pulumi:"fmg"`
	FmgSourceIp                     *string                             `pulumi:"fmgSourceIp"`
	FmgSourceIp6                    *string                             `pulumi:"fmgSourceIp6"`
	FmgUpdatePort                   *string                             `pulumi:"fmgUpdatePort"`
	FortigateCloudSsoDefaultProfile *string                             `pulumi:"fortigateCloudSsoDefaultProfile"`
	GetAllTables                    *string                             `pulumi:"getAllTables"`
	IncludeDefaultServers           *string                             `pulumi:"includeDefaultServers"`
	Interface                       *string                             `pulumi:"interface"`
	InterfaceSelectMethod           *string                             `pulumi:"interfaceSelectMethod"`
	LocalCert                       *string                             `pulumi:"localCert"`
	Mode                            *string                             `pulumi:"mode"`
	ScheduleConfigRestore           *string                             `pulumi:"scheduleConfigRestore"`
	ScheduleScriptRestore           *string                             `pulumi:"scheduleScriptRestore"`
	SerialNumber                    *string                             `pulumi:"serialNumber"`
	ServerLists                     []SystemCentralManagementServerList `pulumi:"serverLists"`
	Type                            *string                             `pulumi:"type"`
	Vdom                            *string                             `pulumi:"vdom"`
	Vdomparam                       *string                             `pulumi:"vdomparam"`
}

type SystemCentralManagementState struct {
	AllowMonitor                    pulumi.StringPtrInput
	AllowPushConfiguration          pulumi.StringPtrInput
	AllowPushFirmware               pulumi.StringPtrInput
	AllowRemoteFirmwareUpgrade      pulumi.StringPtrInput
	CaCert                          pulumi.StringPtrInput
	DynamicSortSubtable             pulumi.StringPtrInput
	EncAlgorithm                    pulumi.StringPtrInput
	Fmg                             pulumi.StringPtrInput
	FmgSourceIp                     pulumi.StringPtrInput
	FmgSourceIp6                    pulumi.StringPtrInput
	FmgUpdatePort                   pulumi.StringPtrInput
	FortigateCloudSsoDefaultProfile pulumi.StringPtrInput
	GetAllTables                    pulumi.StringPtrInput
	IncludeDefaultServers           pulumi.StringPtrInput
	Interface                       pulumi.StringPtrInput
	InterfaceSelectMethod           pulumi.StringPtrInput
	LocalCert                       pulumi.StringPtrInput
	Mode                            pulumi.StringPtrInput
	ScheduleConfigRestore           pulumi.StringPtrInput
	ScheduleScriptRestore           pulumi.StringPtrInput
	SerialNumber                    pulumi.StringPtrInput
	ServerLists                     SystemCentralManagementServerListArrayInput
	Type                            pulumi.StringPtrInput
	Vdom                            pulumi.StringPtrInput
	Vdomparam                       pulumi.StringPtrInput
}

func (SystemCentralManagementState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemCentralManagementState)(nil)).Elem()
}

type systemCentralManagementArgs struct {
	AllowMonitor                    *string                             `pulumi:"allowMonitor"`
	AllowPushConfiguration          *string                             `pulumi:"allowPushConfiguration"`
	AllowPushFirmware               *string                             `pulumi:"allowPushFirmware"`
	AllowRemoteFirmwareUpgrade      *string                             `pulumi:"allowRemoteFirmwareUpgrade"`
	CaCert                          *string                             `pulumi:"caCert"`
	DynamicSortSubtable             *string                             `pulumi:"dynamicSortSubtable"`
	EncAlgorithm                    *string                             `pulumi:"encAlgorithm"`
	Fmg                             *string                             `pulumi:"fmg"`
	FmgSourceIp                     *string                             `pulumi:"fmgSourceIp"`
	FmgSourceIp6                    *string                             `pulumi:"fmgSourceIp6"`
	FmgUpdatePort                   *string                             `pulumi:"fmgUpdatePort"`
	FortigateCloudSsoDefaultProfile *string                             `pulumi:"fortigateCloudSsoDefaultProfile"`
	GetAllTables                    *string                             `pulumi:"getAllTables"`
	IncludeDefaultServers           *string                             `pulumi:"includeDefaultServers"`
	Interface                       *string                             `pulumi:"interface"`
	InterfaceSelectMethod           *string                             `pulumi:"interfaceSelectMethod"`
	LocalCert                       *string                             `pulumi:"localCert"`
	Mode                            *string                             `pulumi:"mode"`
	ScheduleConfigRestore           *string                             `pulumi:"scheduleConfigRestore"`
	ScheduleScriptRestore           *string                             `pulumi:"scheduleScriptRestore"`
	SerialNumber                    *string                             `pulumi:"serialNumber"`
	ServerLists                     []SystemCentralManagementServerList `pulumi:"serverLists"`
	Type                            *string                             `pulumi:"type"`
	Vdom                            *string                             `pulumi:"vdom"`
	Vdomparam                       *string                             `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemCentralManagement resource.
type SystemCentralManagementArgs struct {
	AllowMonitor                    pulumi.StringPtrInput
	AllowPushConfiguration          pulumi.StringPtrInput
	AllowPushFirmware               pulumi.StringPtrInput
	AllowRemoteFirmwareUpgrade      pulumi.StringPtrInput
	CaCert                          pulumi.StringPtrInput
	DynamicSortSubtable             pulumi.StringPtrInput
	EncAlgorithm                    pulumi.StringPtrInput
	Fmg                             pulumi.StringPtrInput
	FmgSourceIp                     pulumi.StringPtrInput
	FmgSourceIp6                    pulumi.StringPtrInput
	FmgUpdatePort                   pulumi.StringPtrInput
	FortigateCloudSsoDefaultProfile pulumi.StringPtrInput
	GetAllTables                    pulumi.StringPtrInput
	IncludeDefaultServers           pulumi.StringPtrInput
	Interface                       pulumi.StringPtrInput
	InterfaceSelectMethod           pulumi.StringPtrInput
	LocalCert                       pulumi.StringPtrInput
	Mode                            pulumi.StringPtrInput
	ScheduleConfigRestore           pulumi.StringPtrInput
	ScheduleScriptRestore           pulumi.StringPtrInput
	SerialNumber                    pulumi.StringPtrInput
	ServerLists                     SystemCentralManagementServerListArrayInput
	Type                            pulumi.StringPtrInput
	Vdom                            pulumi.StringPtrInput
	Vdomparam                       pulumi.StringPtrInput
}

func (SystemCentralManagementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemCentralManagementArgs)(nil)).Elem()
}

type SystemCentralManagementInput interface {
	pulumi.Input

	ToSystemCentralManagementOutput() SystemCentralManagementOutput
	ToSystemCentralManagementOutputWithContext(ctx context.Context) SystemCentralManagementOutput
}

func (*SystemCentralManagement) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemCentralManagement)(nil)).Elem()
}

func (i *SystemCentralManagement) ToSystemCentralManagementOutput() SystemCentralManagementOutput {
	return i.ToSystemCentralManagementOutputWithContext(context.Background())
}

func (i *SystemCentralManagement) ToSystemCentralManagementOutputWithContext(ctx context.Context) SystemCentralManagementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemCentralManagementOutput)
}

func (i *SystemCentralManagement) ToOutput(ctx context.Context) pulumix.Output[*SystemCentralManagement] {
	return pulumix.Output[*SystemCentralManagement]{
		OutputState: i.ToSystemCentralManagementOutputWithContext(ctx).OutputState,
	}
}

// SystemCentralManagementArrayInput is an input type that accepts SystemCentralManagementArray and SystemCentralManagementArrayOutput values.
// You can construct a concrete instance of `SystemCentralManagementArrayInput` via:
//
//	SystemCentralManagementArray{ SystemCentralManagementArgs{...} }
type SystemCentralManagementArrayInput interface {
	pulumi.Input

	ToSystemCentralManagementArrayOutput() SystemCentralManagementArrayOutput
	ToSystemCentralManagementArrayOutputWithContext(context.Context) SystemCentralManagementArrayOutput
}

type SystemCentralManagementArray []SystemCentralManagementInput

func (SystemCentralManagementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemCentralManagement)(nil)).Elem()
}

func (i SystemCentralManagementArray) ToSystemCentralManagementArrayOutput() SystemCentralManagementArrayOutput {
	return i.ToSystemCentralManagementArrayOutputWithContext(context.Background())
}

func (i SystemCentralManagementArray) ToSystemCentralManagementArrayOutputWithContext(ctx context.Context) SystemCentralManagementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemCentralManagementArrayOutput)
}

func (i SystemCentralManagementArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemCentralManagement] {
	return pulumix.Output[[]*SystemCentralManagement]{
		OutputState: i.ToSystemCentralManagementArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemCentralManagementMapInput is an input type that accepts SystemCentralManagementMap and SystemCentralManagementMapOutput values.
// You can construct a concrete instance of `SystemCentralManagementMapInput` via:
//
//	SystemCentralManagementMap{ "key": SystemCentralManagementArgs{...} }
type SystemCentralManagementMapInput interface {
	pulumi.Input

	ToSystemCentralManagementMapOutput() SystemCentralManagementMapOutput
	ToSystemCentralManagementMapOutputWithContext(context.Context) SystemCentralManagementMapOutput
}

type SystemCentralManagementMap map[string]SystemCentralManagementInput

func (SystemCentralManagementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemCentralManagement)(nil)).Elem()
}

func (i SystemCentralManagementMap) ToSystemCentralManagementMapOutput() SystemCentralManagementMapOutput {
	return i.ToSystemCentralManagementMapOutputWithContext(context.Background())
}

func (i SystemCentralManagementMap) ToSystemCentralManagementMapOutputWithContext(ctx context.Context) SystemCentralManagementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemCentralManagementMapOutput)
}

func (i SystemCentralManagementMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemCentralManagement] {
	return pulumix.Output[map[string]*SystemCentralManagement]{
		OutputState: i.ToSystemCentralManagementMapOutputWithContext(ctx).OutputState,
	}
}

type SystemCentralManagementOutput struct{ *pulumi.OutputState }

func (SystemCentralManagementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemCentralManagement)(nil)).Elem()
}

func (o SystemCentralManagementOutput) ToSystemCentralManagementOutput() SystemCentralManagementOutput {
	return o
}

func (o SystemCentralManagementOutput) ToSystemCentralManagementOutputWithContext(ctx context.Context) SystemCentralManagementOutput {
	return o
}

func (o SystemCentralManagementOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemCentralManagement] {
	return pulumix.Output[*SystemCentralManagement]{
		OutputState: o.OutputState,
	}
}

func (o SystemCentralManagementOutput) AllowMonitor() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.AllowMonitor }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) AllowPushConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.AllowPushConfiguration }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) AllowPushFirmware() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.AllowPushFirmware }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) AllowRemoteFirmwareUpgrade() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.AllowRemoteFirmwareUpgrade }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) CaCert() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.CaCert }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SystemCentralManagementOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) Fmg() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.Fmg }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) FmgSourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.FmgSourceIp }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) FmgSourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.FmgSourceIp6 }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) FmgUpdatePort() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.FmgUpdatePort }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) FortigateCloudSsoDefaultProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.FortigateCloudSsoDefaultProfile }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SystemCentralManagementOutput) IncludeDefaultServers() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.IncludeDefaultServers }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) LocalCert() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.LocalCert }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) ScheduleConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.ScheduleConfigRestore }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) ScheduleScriptRestore() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.ScheduleScriptRestore }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) ServerLists() SystemCentralManagementServerListArrayOutput {
	return o.ApplyT(func(v *SystemCentralManagement) SystemCentralManagementServerListArrayOutput { return v.ServerLists }).(SystemCentralManagementServerListArrayOutput)
}

func (o SystemCentralManagementOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringOutput { return v.Vdom }).(pulumi.StringOutput)
}

func (o SystemCentralManagementOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemCentralManagement) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemCentralManagementArrayOutput struct{ *pulumi.OutputState }

func (SystemCentralManagementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemCentralManagement)(nil)).Elem()
}

func (o SystemCentralManagementArrayOutput) ToSystemCentralManagementArrayOutput() SystemCentralManagementArrayOutput {
	return o
}

func (o SystemCentralManagementArrayOutput) ToSystemCentralManagementArrayOutputWithContext(ctx context.Context) SystemCentralManagementArrayOutput {
	return o
}

func (o SystemCentralManagementArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemCentralManagement] {
	return pulumix.Output[[]*SystemCentralManagement]{
		OutputState: o.OutputState,
	}
}

func (o SystemCentralManagementArrayOutput) Index(i pulumi.IntInput) SystemCentralManagementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemCentralManagement {
		return vs[0].([]*SystemCentralManagement)[vs[1].(int)]
	}).(SystemCentralManagementOutput)
}

type SystemCentralManagementMapOutput struct{ *pulumi.OutputState }

func (SystemCentralManagementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemCentralManagement)(nil)).Elem()
}

func (o SystemCentralManagementMapOutput) ToSystemCentralManagementMapOutput() SystemCentralManagementMapOutput {
	return o
}

func (o SystemCentralManagementMapOutput) ToSystemCentralManagementMapOutputWithContext(ctx context.Context) SystemCentralManagementMapOutput {
	return o
}

func (o SystemCentralManagementMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemCentralManagement] {
	return pulumix.Output[map[string]*SystemCentralManagement]{
		OutputState: o.OutputState,
	}
}

func (o SystemCentralManagementMapOutput) MapIndex(k pulumi.StringInput) SystemCentralManagementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemCentralManagement {
		return vs[0].(map[string]*SystemCentralManagement)[vs[1].(string)]
	}).(SystemCentralManagementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemCentralManagementInput)(nil)).Elem(), &SystemCentralManagement{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemCentralManagementArrayInput)(nil)).Elem(), SystemCentralManagementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemCentralManagementMapInput)(nil)).Elem(), SystemCentralManagementMap{})
	pulumi.RegisterOutputType(SystemCentralManagementOutput{})
	pulumi.RegisterOutputType(SystemCentralManagementArrayOutput{})
	pulumi.RegisterOutputType(SystemCentralManagementMapOutput{})
}
