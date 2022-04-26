// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemSettingGlobal struct {
	pulumi.CustomResourceState

	// Enable SCP over SSH
	AdminScp pulumi.StringOutput `pulumi:"adminScp"`
	// Administrative access port for HTTPS.
	AdminSport pulumi.StringOutput `pulumi:"adminSport"`
	// Administrative access port for SSH.
	AdminSshPort pulumi.StringOutput `pulumi:"adminSshPort"`
	// Number of minutes before an idle administrator session time out.
	Admintimeout pulumi.StringOutput `pulumi:"admintimeout"`
	// FortiGate unit's hostname.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Number corresponding to your time zone from 00 to 86.
	Timezone pulumi.StringOutput `pulumi:"timezone"`
}

// NewSystemSettingGlobal registers a new resource with the given unique name, arguments, and options.
func NewSystemSettingGlobal(ctx *pulumi.Context,
	name string, args *SystemSettingGlobalArgs, opts ...pulumi.ResourceOption) (*SystemSettingGlobal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSettingGlobal
	err := ctx.RegisterResource("fortios:index/systemSettingGlobal:SystemSettingGlobal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSettingGlobal gets an existing SystemSettingGlobal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSettingGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSettingGlobalState, opts ...pulumi.ResourceOption) (*SystemSettingGlobal, error) {
	var resource SystemSettingGlobal
	err := ctx.ReadResource("fortios:index/systemSettingGlobal:SystemSettingGlobal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSettingGlobal resources.
type systemSettingGlobalState struct {
	// Enable SCP over SSH
	AdminScp *string `pulumi:"adminScp"`
	// Administrative access port for HTTPS.
	AdminSport *string `pulumi:"adminSport"`
	// Administrative access port for SSH.
	AdminSshPort *string `pulumi:"adminSshPort"`
	// Number of minutes before an idle administrator session time out.
	Admintimeout *string `pulumi:"admintimeout"`
	// FortiGate unit's hostname.
	Hostname *string `pulumi:"hostname"`
	// Number corresponding to your time zone from 00 to 86.
	Timezone *string `pulumi:"timezone"`
}

type SystemSettingGlobalState struct {
	// Enable SCP over SSH
	AdminScp pulumi.StringPtrInput
	// Administrative access port for HTTPS.
	AdminSport pulumi.StringPtrInput
	// Administrative access port for SSH.
	AdminSshPort pulumi.StringPtrInput
	// Number of minutes before an idle administrator session time out.
	Admintimeout pulumi.StringPtrInput
	// FortiGate unit's hostname.
	Hostname pulumi.StringPtrInput
	// Number corresponding to your time zone from 00 to 86.
	Timezone pulumi.StringPtrInput
}

func (SystemSettingGlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSettingGlobalState)(nil)).Elem()
}

type systemSettingGlobalArgs struct {
	// Enable SCP over SSH
	AdminScp *string `pulumi:"adminScp"`
	// Administrative access port for HTTPS.
	AdminSport *string `pulumi:"adminSport"`
	// Administrative access port for SSH.
	AdminSshPort *string `pulumi:"adminSshPort"`
	// Number of minutes before an idle administrator session time out.
	Admintimeout *string `pulumi:"admintimeout"`
	// FortiGate unit's hostname.
	Hostname string `pulumi:"hostname"`
	// Number corresponding to your time zone from 00 to 86.
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a SystemSettingGlobal resource.
type SystemSettingGlobalArgs struct {
	// Enable SCP over SSH
	AdminScp pulumi.StringPtrInput
	// Administrative access port for HTTPS.
	AdminSport pulumi.StringPtrInput
	// Administrative access port for SSH.
	AdminSshPort pulumi.StringPtrInput
	// Number of minutes before an idle administrator session time out.
	Admintimeout pulumi.StringPtrInput
	// FortiGate unit's hostname.
	Hostname pulumi.StringInput
	// Number corresponding to your time zone from 00 to 86.
	Timezone pulumi.StringPtrInput
}

func (SystemSettingGlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSettingGlobalArgs)(nil)).Elem()
}

type SystemSettingGlobalInput interface {
	pulumi.Input

	ToSystemSettingGlobalOutput() SystemSettingGlobalOutput
	ToSystemSettingGlobalOutputWithContext(ctx context.Context) SystemSettingGlobalOutput
}

func (*SystemSettingGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSettingGlobal)(nil)).Elem()
}

func (i *SystemSettingGlobal) ToSystemSettingGlobalOutput() SystemSettingGlobalOutput {
	return i.ToSystemSettingGlobalOutputWithContext(context.Background())
}

func (i *SystemSettingGlobal) ToSystemSettingGlobalOutputWithContext(ctx context.Context) SystemSettingGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingGlobalOutput)
}

// SystemSettingGlobalArrayInput is an input type that accepts SystemSettingGlobalArray and SystemSettingGlobalArrayOutput values.
// You can construct a concrete instance of `SystemSettingGlobalArrayInput` via:
//
//          SystemSettingGlobalArray{ SystemSettingGlobalArgs{...} }
type SystemSettingGlobalArrayInput interface {
	pulumi.Input

	ToSystemSettingGlobalArrayOutput() SystemSettingGlobalArrayOutput
	ToSystemSettingGlobalArrayOutputWithContext(context.Context) SystemSettingGlobalArrayOutput
}

type SystemSettingGlobalArray []SystemSettingGlobalInput

func (SystemSettingGlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSettingGlobal)(nil)).Elem()
}

func (i SystemSettingGlobalArray) ToSystemSettingGlobalArrayOutput() SystemSettingGlobalArrayOutput {
	return i.ToSystemSettingGlobalArrayOutputWithContext(context.Background())
}

func (i SystemSettingGlobalArray) ToSystemSettingGlobalArrayOutputWithContext(ctx context.Context) SystemSettingGlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingGlobalArrayOutput)
}

// SystemSettingGlobalMapInput is an input type that accepts SystemSettingGlobalMap and SystemSettingGlobalMapOutput values.
// You can construct a concrete instance of `SystemSettingGlobalMapInput` via:
//
//          SystemSettingGlobalMap{ "key": SystemSettingGlobalArgs{...} }
type SystemSettingGlobalMapInput interface {
	pulumi.Input

	ToSystemSettingGlobalMapOutput() SystemSettingGlobalMapOutput
	ToSystemSettingGlobalMapOutputWithContext(context.Context) SystemSettingGlobalMapOutput
}

type SystemSettingGlobalMap map[string]SystemSettingGlobalInput

func (SystemSettingGlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSettingGlobal)(nil)).Elem()
}

func (i SystemSettingGlobalMap) ToSystemSettingGlobalMapOutput() SystemSettingGlobalMapOutput {
	return i.ToSystemSettingGlobalMapOutputWithContext(context.Background())
}

func (i SystemSettingGlobalMap) ToSystemSettingGlobalMapOutputWithContext(ctx context.Context) SystemSettingGlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingGlobalMapOutput)
}

type SystemSettingGlobalOutput struct{ *pulumi.OutputState }

func (SystemSettingGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSettingGlobal)(nil)).Elem()
}

func (o SystemSettingGlobalOutput) ToSystemSettingGlobalOutput() SystemSettingGlobalOutput {
	return o
}

func (o SystemSettingGlobalOutput) ToSystemSettingGlobalOutputWithContext(ctx context.Context) SystemSettingGlobalOutput {
	return o
}

type SystemSettingGlobalArrayOutput struct{ *pulumi.OutputState }

func (SystemSettingGlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSettingGlobal)(nil)).Elem()
}

func (o SystemSettingGlobalArrayOutput) ToSystemSettingGlobalArrayOutput() SystemSettingGlobalArrayOutput {
	return o
}

func (o SystemSettingGlobalArrayOutput) ToSystemSettingGlobalArrayOutputWithContext(ctx context.Context) SystemSettingGlobalArrayOutput {
	return o
}

func (o SystemSettingGlobalArrayOutput) Index(i pulumi.IntInput) SystemSettingGlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSettingGlobal {
		return vs[0].([]*SystemSettingGlobal)[vs[1].(int)]
	}).(SystemSettingGlobalOutput)
}

type SystemSettingGlobalMapOutput struct{ *pulumi.OutputState }

func (SystemSettingGlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSettingGlobal)(nil)).Elem()
}

func (o SystemSettingGlobalMapOutput) ToSystemSettingGlobalMapOutput() SystemSettingGlobalMapOutput {
	return o
}

func (o SystemSettingGlobalMapOutput) ToSystemSettingGlobalMapOutputWithContext(ctx context.Context) SystemSettingGlobalMapOutput {
	return o
}

func (o SystemSettingGlobalMapOutput) MapIndex(k pulumi.StringInput) SystemSettingGlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSettingGlobal {
		return vs[0].(map[string]*SystemSettingGlobal)[vs[1].(string)]
	}).(SystemSettingGlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingGlobalInput)(nil)).Elem(), &SystemSettingGlobal{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingGlobalArrayInput)(nil)).Elem(), SystemSettingGlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingGlobalMapInput)(nil)).Elem(), SystemSettingGlobalMap{})
	pulumi.RegisterOutputType(SystemSettingGlobalOutput{})
	pulumi.RegisterOutputType(SystemSettingGlobalArrayOutput{})
	pulumi.RegisterOutputType(SystemSettingGlobalMapOutput{})
}
