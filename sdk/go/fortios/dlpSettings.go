// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DlpSettings struct {
	pulumi.CustomResourceState

	CacheMemPercent pulumi.IntOutput       `pulumi:"cacheMemPercent"`
	ChunkSize       pulumi.IntOutput       `pulumi:"chunkSize"`
	DbMode          pulumi.StringOutput    `pulumi:"dbMode"`
	Size            pulumi.IntOutput       `pulumi:"size"`
	StorageDevice   pulumi.StringOutput    `pulumi:"storageDevice"`
	Vdomparam       pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDlpSettings registers a new resource with the given unique name, arguments, and options.
func NewDlpSettings(ctx *pulumi.Context,
	name string, args *DlpSettingsArgs, opts ...pulumi.ResourceOption) (*DlpSettings, error) {
	if args == nil {
		args = &DlpSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DlpSettings
	err := ctx.RegisterResource("fortios:index/dlpSettings:DlpSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpSettings gets an existing DlpSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpSettingsState, opts ...pulumi.ResourceOption) (*DlpSettings, error) {
	var resource DlpSettings
	err := ctx.ReadResource("fortios:index/dlpSettings:DlpSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpSettings resources.
type dlpSettingsState struct {
	CacheMemPercent *int    `pulumi:"cacheMemPercent"`
	ChunkSize       *int    `pulumi:"chunkSize"`
	DbMode          *string `pulumi:"dbMode"`
	Size            *int    `pulumi:"size"`
	StorageDevice   *string `pulumi:"storageDevice"`
	Vdomparam       *string `pulumi:"vdomparam"`
}

type DlpSettingsState struct {
	CacheMemPercent pulumi.IntPtrInput
	ChunkSize       pulumi.IntPtrInput
	DbMode          pulumi.StringPtrInput
	Size            pulumi.IntPtrInput
	StorageDevice   pulumi.StringPtrInput
	Vdomparam       pulumi.StringPtrInput
}

func (DlpSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpSettingsState)(nil)).Elem()
}

type dlpSettingsArgs struct {
	CacheMemPercent *int    `pulumi:"cacheMemPercent"`
	ChunkSize       *int    `pulumi:"chunkSize"`
	DbMode          *string `pulumi:"dbMode"`
	Size            *int    `pulumi:"size"`
	StorageDevice   *string `pulumi:"storageDevice"`
	Vdomparam       *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpSettings resource.
type DlpSettingsArgs struct {
	CacheMemPercent pulumi.IntPtrInput
	ChunkSize       pulumi.IntPtrInput
	DbMode          pulumi.StringPtrInput
	Size            pulumi.IntPtrInput
	StorageDevice   pulumi.StringPtrInput
	Vdomparam       pulumi.StringPtrInput
}

func (DlpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpSettingsArgs)(nil)).Elem()
}

type DlpSettingsInput interface {
	pulumi.Input

	ToDlpSettingsOutput() DlpSettingsOutput
	ToDlpSettingsOutputWithContext(ctx context.Context) DlpSettingsOutput
}

func (*DlpSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpSettings)(nil)).Elem()
}

func (i *DlpSettings) ToDlpSettingsOutput() DlpSettingsOutput {
	return i.ToDlpSettingsOutputWithContext(context.Background())
}

func (i *DlpSettings) ToDlpSettingsOutputWithContext(ctx context.Context) DlpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSettingsOutput)
}

// DlpSettingsArrayInput is an input type that accepts DlpSettingsArray and DlpSettingsArrayOutput values.
// You can construct a concrete instance of `DlpSettingsArrayInput` via:
//
//	DlpSettingsArray{ DlpSettingsArgs{...} }
type DlpSettingsArrayInput interface {
	pulumi.Input

	ToDlpSettingsArrayOutput() DlpSettingsArrayOutput
	ToDlpSettingsArrayOutputWithContext(context.Context) DlpSettingsArrayOutput
}

type DlpSettingsArray []DlpSettingsInput

func (DlpSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpSettings)(nil)).Elem()
}

func (i DlpSettingsArray) ToDlpSettingsArrayOutput() DlpSettingsArrayOutput {
	return i.ToDlpSettingsArrayOutputWithContext(context.Background())
}

func (i DlpSettingsArray) ToDlpSettingsArrayOutputWithContext(ctx context.Context) DlpSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSettingsArrayOutput)
}

// DlpSettingsMapInput is an input type that accepts DlpSettingsMap and DlpSettingsMapOutput values.
// You can construct a concrete instance of `DlpSettingsMapInput` via:
//
//	DlpSettingsMap{ "key": DlpSettingsArgs{...} }
type DlpSettingsMapInput interface {
	pulumi.Input

	ToDlpSettingsMapOutput() DlpSettingsMapOutput
	ToDlpSettingsMapOutputWithContext(context.Context) DlpSettingsMapOutput
}

type DlpSettingsMap map[string]DlpSettingsInput

func (DlpSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpSettings)(nil)).Elem()
}

func (i DlpSettingsMap) ToDlpSettingsMapOutput() DlpSettingsMapOutput {
	return i.ToDlpSettingsMapOutputWithContext(context.Background())
}

func (i DlpSettingsMap) ToDlpSettingsMapOutputWithContext(ctx context.Context) DlpSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSettingsMapOutput)
}

type DlpSettingsOutput struct{ *pulumi.OutputState }

func (DlpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpSettings)(nil)).Elem()
}

func (o DlpSettingsOutput) ToDlpSettingsOutput() DlpSettingsOutput {
	return o
}

func (o DlpSettingsOutput) ToDlpSettingsOutputWithContext(ctx context.Context) DlpSettingsOutput {
	return o
}

func (o DlpSettingsOutput) CacheMemPercent() pulumi.IntOutput {
	return o.ApplyT(func(v *DlpSettings) pulumi.IntOutput { return v.CacheMemPercent }).(pulumi.IntOutput)
}

func (o DlpSettingsOutput) ChunkSize() pulumi.IntOutput {
	return o.ApplyT(func(v *DlpSettings) pulumi.IntOutput { return v.ChunkSize }).(pulumi.IntOutput)
}

func (o DlpSettingsOutput) DbMode() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpSettings) pulumi.StringOutput { return v.DbMode }).(pulumi.StringOutput)
}

func (o DlpSettingsOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *DlpSettings) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

func (o DlpSettingsOutput) StorageDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpSettings) pulumi.StringOutput { return v.StorageDevice }).(pulumi.StringOutput)
}

func (o DlpSettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpSettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DlpSettingsArrayOutput struct{ *pulumi.OutputState }

func (DlpSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpSettings)(nil)).Elem()
}

func (o DlpSettingsArrayOutput) ToDlpSettingsArrayOutput() DlpSettingsArrayOutput {
	return o
}

func (o DlpSettingsArrayOutput) ToDlpSettingsArrayOutputWithContext(ctx context.Context) DlpSettingsArrayOutput {
	return o
}

func (o DlpSettingsArrayOutput) Index(i pulumi.IntInput) DlpSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DlpSettings {
		return vs[0].([]*DlpSettings)[vs[1].(int)]
	}).(DlpSettingsOutput)
}

type DlpSettingsMapOutput struct{ *pulumi.OutputState }

func (DlpSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpSettings)(nil)).Elem()
}

func (o DlpSettingsMapOutput) ToDlpSettingsMapOutput() DlpSettingsMapOutput {
	return o
}

func (o DlpSettingsMapOutput) ToDlpSettingsMapOutputWithContext(ctx context.Context) DlpSettingsMapOutput {
	return o
}

func (o DlpSettingsMapOutput) MapIndex(k pulumi.StringInput) DlpSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DlpSettings {
		return vs[0].(map[string]*DlpSettings)[vs[1].(string)]
	}).(DlpSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DlpSettingsInput)(nil)).Elem(), &DlpSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpSettingsArrayInput)(nil)).Elem(), DlpSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpSettingsMapInput)(nil)).Elem(), DlpSettingsMap{})
	pulumi.RegisterOutputType(DlpSettingsOutput{})
	pulumi.RegisterOutputType(DlpSettingsArrayOutput{})
	pulumi.RegisterOutputType(DlpSettingsMapOutput{})
}
