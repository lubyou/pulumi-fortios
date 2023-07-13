// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IpsSensor struct {
	pulumi.CustomResourceState

	BlockMaliciousUrl     pulumi.StringOutput          `pulumi:"blockMaliciousUrl"`
	Comment               pulumi.StringPtrOutput       `pulumi:"comment"`
	DynamicSortSubtable   pulumi.StringPtrOutput       `pulumi:"dynamicSortSubtable"`
	Entries               IpsSensorEntryArrayOutput    `pulumi:"entries"`
	ExtendedLog           pulumi.StringOutput          `pulumi:"extendedLog"`
	Filters               IpsSensorFilterArrayOutput   `pulumi:"filters"`
	GetAllTables          pulumi.StringPtrOutput       `pulumi:"getAllTables"`
	Name                  pulumi.StringOutput          `pulumi:"name"`
	Overrides             IpsSensorOverrideArrayOutput `pulumi:"overrides"`
	ReplacemsgGroup       pulumi.StringOutput          `pulumi:"replacemsgGroup"`
	ScanBotnetConnections pulumi.StringOutput          `pulumi:"scanBotnetConnections"`
	Vdomparam             pulumi.StringPtrOutput       `pulumi:"vdomparam"`
}

// NewIpsSensor registers a new resource with the given unique name, arguments, and options.
func NewIpsSensor(ctx *pulumi.Context,
	name string, args *IpsSensorArgs, opts ...pulumi.ResourceOption) (*IpsSensor, error) {
	if args == nil {
		args = &IpsSensorArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpsSensor
	err := ctx.RegisterResource("fortios:index/ipsSensor:IpsSensor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsSensor gets an existing IpsSensor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsSensor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsSensorState, opts ...pulumi.ResourceOption) (*IpsSensor, error) {
	var resource IpsSensor
	err := ctx.ReadResource("fortios:index/ipsSensor:IpsSensor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsSensor resources.
type ipsSensorState struct {
	BlockMaliciousUrl     *string             `pulumi:"blockMaliciousUrl"`
	Comment               *string             `pulumi:"comment"`
	DynamicSortSubtable   *string             `pulumi:"dynamicSortSubtable"`
	Entries               []IpsSensorEntry    `pulumi:"entries"`
	ExtendedLog           *string             `pulumi:"extendedLog"`
	Filters               []IpsSensorFilter   `pulumi:"filters"`
	GetAllTables          *string             `pulumi:"getAllTables"`
	Name                  *string             `pulumi:"name"`
	Overrides             []IpsSensorOverride `pulumi:"overrides"`
	ReplacemsgGroup       *string             `pulumi:"replacemsgGroup"`
	ScanBotnetConnections *string             `pulumi:"scanBotnetConnections"`
	Vdomparam             *string             `pulumi:"vdomparam"`
}

type IpsSensorState struct {
	BlockMaliciousUrl     pulumi.StringPtrInput
	Comment               pulumi.StringPtrInput
	DynamicSortSubtable   pulumi.StringPtrInput
	Entries               IpsSensorEntryArrayInput
	ExtendedLog           pulumi.StringPtrInput
	Filters               IpsSensorFilterArrayInput
	GetAllTables          pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	Overrides             IpsSensorOverrideArrayInput
	ReplacemsgGroup       pulumi.StringPtrInput
	ScanBotnetConnections pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (IpsSensorState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsSensorState)(nil)).Elem()
}

type ipsSensorArgs struct {
	BlockMaliciousUrl     *string             `pulumi:"blockMaliciousUrl"`
	Comment               *string             `pulumi:"comment"`
	DynamicSortSubtable   *string             `pulumi:"dynamicSortSubtable"`
	Entries               []IpsSensorEntry    `pulumi:"entries"`
	ExtendedLog           *string             `pulumi:"extendedLog"`
	Filters               []IpsSensorFilter   `pulumi:"filters"`
	GetAllTables          *string             `pulumi:"getAllTables"`
	Name                  *string             `pulumi:"name"`
	Overrides             []IpsSensorOverride `pulumi:"overrides"`
	ReplacemsgGroup       *string             `pulumi:"replacemsgGroup"`
	ScanBotnetConnections *string             `pulumi:"scanBotnetConnections"`
	Vdomparam             *string             `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsSensor resource.
type IpsSensorArgs struct {
	BlockMaliciousUrl     pulumi.StringPtrInput
	Comment               pulumi.StringPtrInput
	DynamicSortSubtable   pulumi.StringPtrInput
	Entries               IpsSensorEntryArrayInput
	ExtendedLog           pulumi.StringPtrInput
	Filters               IpsSensorFilterArrayInput
	GetAllTables          pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	Overrides             IpsSensorOverrideArrayInput
	ReplacemsgGroup       pulumi.StringPtrInput
	ScanBotnetConnections pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (IpsSensorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsSensorArgs)(nil)).Elem()
}

type IpsSensorInput interface {
	pulumi.Input

	ToIpsSensorOutput() IpsSensorOutput
	ToIpsSensorOutputWithContext(ctx context.Context) IpsSensorOutput
}

func (*IpsSensor) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsSensor)(nil)).Elem()
}

func (i *IpsSensor) ToIpsSensorOutput() IpsSensorOutput {
	return i.ToIpsSensorOutputWithContext(context.Background())
}

func (i *IpsSensor) ToIpsSensorOutputWithContext(ctx context.Context) IpsSensorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSensorOutput)
}

// IpsSensorArrayInput is an input type that accepts IpsSensorArray and IpsSensorArrayOutput values.
// You can construct a concrete instance of `IpsSensorArrayInput` via:
//
//	IpsSensorArray{ IpsSensorArgs{...} }
type IpsSensorArrayInput interface {
	pulumi.Input

	ToIpsSensorArrayOutput() IpsSensorArrayOutput
	ToIpsSensorArrayOutputWithContext(context.Context) IpsSensorArrayOutput
}

type IpsSensorArray []IpsSensorInput

func (IpsSensorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsSensor)(nil)).Elem()
}

func (i IpsSensorArray) ToIpsSensorArrayOutput() IpsSensorArrayOutput {
	return i.ToIpsSensorArrayOutputWithContext(context.Background())
}

func (i IpsSensorArray) ToIpsSensorArrayOutputWithContext(ctx context.Context) IpsSensorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSensorArrayOutput)
}

// IpsSensorMapInput is an input type that accepts IpsSensorMap and IpsSensorMapOutput values.
// You can construct a concrete instance of `IpsSensorMapInput` via:
//
//	IpsSensorMap{ "key": IpsSensorArgs{...} }
type IpsSensorMapInput interface {
	pulumi.Input

	ToIpsSensorMapOutput() IpsSensorMapOutput
	ToIpsSensorMapOutputWithContext(context.Context) IpsSensorMapOutput
}

type IpsSensorMap map[string]IpsSensorInput

func (IpsSensorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsSensor)(nil)).Elem()
}

func (i IpsSensorMap) ToIpsSensorMapOutput() IpsSensorMapOutput {
	return i.ToIpsSensorMapOutputWithContext(context.Background())
}

func (i IpsSensorMap) ToIpsSensorMapOutputWithContext(ctx context.Context) IpsSensorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSensorMapOutput)
}

type IpsSensorOutput struct{ *pulumi.OutputState }

func (IpsSensorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsSensor)(nil)).Elem()
}

func (o IpsSensorOutput) ToIpsSensorOutput() IpsSensorOutput {
	return o
}

func (o IpsSensorOutput) ToIpsSensorOutputWithContext(ctx context.Context) IpsSensorOutput {
	return o
}

func (o IpsSensorOutput) BlockMaliciousUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringOutput { return v.BlockMaliciousUrl }).(pulumi.StringOutput)
}

func (o IpsSensorOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o IpsSensorOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o IpsSensorOutput) Entries() IpsSensorEntryArrayOutput {
	return o.ApplyT(func(v *IpsSensor) IpsSensorEntryArrayOutput { return v.Entries }).(IpsSensorEntryArrayOutput)
}

func (o IpsSensorOutput) ExtendedLog() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringOutput { return v.ExtendedLog }).(pulumi.StringOutput)
}

func (o IpsSensorOutput) Filters() IpsSensorFilterArrayOutput {
	return o.ApplyT(func(v *IpsSensor) IpsSensorFilterArrayOutput { return v.Filters }).(IpsSensorFilterArrayOutput)
}

func (o IpsSensorOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o IpsSensorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IpsSensorOutput) Overrides() IpsSensorOverrideArrayOutput {
	return o.ApplyT(func(v *IpsSensor) IpsSensorOverrideArrayOutput { return v.Overrides }).(IpsSensorOverrideArrayOutput)
}

func (o IpsSensorOutput) ReplacemsgGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringOutput { return v.ReplacemsgGroup }).(pulumi.StringOutput)
}

func (o IpsSensorOutput) ScanBotnetConnections() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringOutput { return v.ScanBotnetConnections }).(pulumi.StringOutput)
}

func (o IpsSensorOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsSensor) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IpsSensorArrayOutput struct{ *pulumi.OutputState }

func (IpsSensorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsSensor)(nil)).Elem()
}

func (o IpsSensorArrayOutput) ToIpsSensorArrayOutput() IpsSensorArrayOutput {
	return o
}

func (o IpsSensorArrayOutput) ToIpsSensorArrayOutputWithContext(ctx context.Context) IpsSensorArrayOutput {
	return o
}

func (o IpsSensorArrayOutput) Index(i pulumi.IntInput) IpsSensorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsSensor {
		return vs[0].([]*IpsSensor)[vs[1].(int)]
	}).(IpsSensorOutput)
}

type IpsSensorMapOutput struct{ *pulumi.OutputState }

func (IpsSensorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsSensor)(nil)).Elem()
}

func (o IpsSensorMapOutput) ToIpsSensorMapOutput() IpsSensorMapOutput {
	return o
}

func (o IpsSensorMapOutput) ToIpsSensorMapOutputWithContext(ctx context.Context) IpsSensorMapOutput {
	return o
}

func (o IpsSensorMapOutput) MapIndex(k pulumi.StringInput) IpsSensorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsSensor {
		return vs[0].(map[string]*IpsSensor)[vs[1].(string)]
	}).(IpsSensorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSensorInput)(nil)).Elem(), &IpsSensor{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSensorArrayInput)(nil)).Elem(), IpsSensorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSensorMapInput)(nil)).Elem(), IpsSensorMap{})
	pulumi.RegisterOutputType(IpsSensorOutput{})
	pulumi.RegisterOutputType(IpsSensorArrayOutput{})
	pulumi.RegisterOutputType(IpsSensorMapOutput{})
}
