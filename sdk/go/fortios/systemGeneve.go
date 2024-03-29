// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemGeneve struct {
	pulumi.CustomResourceState

	Dstport   pulumi.IntOutput       `pulumi:"dstport"`
	Interface pulumi.StringOutput    `pulumi:"interface"`
	IpVersion pulumi.StringOutput    `pulumi:"ipVersion"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	RemoteIp  pulumi.StringOutput    `pulumi:"remoteIp"`
	RemoteIp6 pulumi.StringOutput    `pulumi:"remoteIp6"`
	Type      pulumi.StringOutput    `pulumi:"type"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	Vni       pulumi.IntOutput       `pulumi:"vni"`
}

// NewSystemGeneve registers a new resource with the given unique name, arguments, and options.
func NewSystemGeneve(ctx *pulumi.Context,
	name string, args *SystemGeneveArgs, opts ...pulumi.ResourceOption) (*SystemGeneve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.IpVersion == nil {
		return nil, errors.New("invalid value for required argument 'IpVersion'")
	}
	if args.RemoteIp == nil {
		return nil, errors.New("invalid value for required argument 'RemoteIp'")
	}
	if args.Vni == nil {
		return nil, errors.New("invalid value for required argument 'Vni'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemGeneve
	err := ctx.RegisterResource("fortios:index/systemGeneve:SystemGeneve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemGeneve gets an existing SystemGeneve resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemGeneve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemGeneveState, opts ...pulumi.ResourceOption) (*SystemGeneve, error) {
	var resource SystemGeneve
	err := ctx.ReadResource("fortios:index/systemGeneve:SystemGeneve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemGeneve resources.
type systemGeneveState struct {
	Dstport   *int    `pulumi:"dstport"`
	Interface *string `pulumi:"interface"`
	IpVersion *string `pulumi:"ipVersion"`
	Name      *string `pulumi:"name"`
	RemoteIp  *string `pulumi:"remoteIp"`
	RemoteIp6 *string `pulumi:"remoteIp6"`
	Type      *string `pulumi:"type"`
	Vdomparam *string `pulumi:"vdomparam"`
	Vni       *int    `pulumi:"vni"`
}

type SystemGeneveState struct {
	Dstport   pulumi.IntPtrInput
	Interface pulumi.StringPtrInput
	IpVersion pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	RemoteIp  pulumi.StringPtrInput
	RemoteIp6 pulumi.StringPtrInput
	Type      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
	Vni       pulumi.IntPtrInput
}

func (SystemGeneveState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGeneveState)(nil)).Elem()
}

type systemGeneveArgs struct {
	Dstport   *int    `pulumi:"dstport"`
	Interface string  `pulumi:"interface"`
	IpVersion string  `pulumi:"ipVersion"`
	Name      *string `pulumi:"name"`
	RemoteIp  string  `pulumi:"remoteIp"`
	RemoteIp6 *string `pulumi:"remoteIp6"`
	Type      *string `pulumi:"type"`
	Vdomparam *string `pulumi:"vdomparam"`
	Vni       int     `pulumi:"vni"`
}

// The set of arguments for constructing a SystemGeneve resource.
type SystemGeneveArgs struct {
	Dstport   pulumi.IntPtrInput
	Interface pulumi.StringInput
	IpVersion pulumi.StringInput
	Name      pulumi.StringPtrInput
	RemoteIp  pulumi.StringInput
	RemoteIp6 pulumi.StringPtrInput
	Type      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
	Vni       pulumi.IntInput
}

func (SystemGeneveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGeneveArgs)(nil)).Elem()
}

type SystemGeneveInput interface {
	pulumi.Input

	ToSystemGeneveOutput() SystemGeneveOutput
	ToSystemGeneveOutputWithContext(ctx context.Context) SystemGeneveOutput
}

func (*SystemGeneve) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGeneve)(nil)).Elem()
}

func (i *SystemGeneve) ToSystemGeneveOutput() SystemGeneveOutput {
	return i.ToSystemGeneveOutputWithContext(context.Background())
}

func (i *SystemGeneve) ToSystemGeneveOutputWithContext(ctx context.Context) SystemGeneveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeneveOutput)
}

// SystemGeneveArrayInput is an input type that accepts SystemGeneveArray and SystemGeneveArrayOutput values.
// You can construct a concrete instance of `SystemGeneveArrayInput` via:
//
//	SystemGeneveArray{ SystemGeneveArgs{...} }
type SystemGeneveArrayInput interface {
	pulumi.Input

	ToSystemGeneveArrayOutput() SystemGeneveArrayOutput
	ToSystemGeneveArrayOutputWithContext(context.Context) SystemGeneveArrayOutput
}

type SystemGeneveArray []SystemGeneveInput

func (SystemGeneveArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGeneve)(nil)).Elem()
}

func (i SystemGeneveArray) ToSystemGeneveArrayOutput() SystemGeneveArrayOutput {
	return i.ToSystemGeneveArrayOutputWithContext(context.Background())
}

func (i SystemGeneveArray) ToSystemGeneveArrayOutputWithContext(ctx context.Context) SystemGeneveArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeneveArrayOutput)
}

// SystemGeneveMapInput is an input type that accepts SystemGeneveMap and SystemGeneveMapOutput values.
// You can construct a concrete instance of `SystemGeneveMapInput` via:
//
//	SystemGeneveMap{ "key": SystemGeneveArgs{...} }
type SystemGeneveMapInput interface {
	pulumi.Input

	ToSystemGeneveMapOutput() SystemGeneveMapOutput
	ToSystemGeneveMapOutputWithContext(context.Context) SystemGeneveMapOutput
}

type SystemGeneveMap map[string]SystemGeneveInput

func (SystemGeneveMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGeneve)(nil)).Elem()
}

func (i SystemGeneveMap) ToSystemGeneveMapOutput() SystemGeneveMapOutput {
	return i.ToSystemGeneveMapOutputWithContext(context.Background())
}

func (i SystemGeneveMap) ToSystemGeneveMapOutputWithContext(ctx context.Context) SystemGeneveMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeneveMapOutput)
}

type SystemGeneveOutput struct{ *pulumi.OutputState }

func (SystemGeneveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGeneve)(nil)).Elem()
}

func (o SystemGeneveOutput) ToSystemGeneveOutput() SystemGeneveOutput {
	return o
}

func (o SystemGeneveOutput) ToSystemGeneveOutputWithContext(ctx context.Context) SystemGeneveOutput {
	return o
}

func (o SystemGeneveOutput) Dstport() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.IntOutput { return v.Dstport }).(pulumi.IntOutput)
}

func (o SystemGeneveOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemGeneveOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

func (o SystemGeneveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemGeneveOutput) RemoteIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.StringOutput { return v.RemoteIp }).(pulumi.StringOutput)
}

func (o SystemGeneveOutput) RemoteIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.StringOutput { return v.RemoteIp6 }).(pulumi.StringOutput)
}

func (o SystemGeneveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SystemGeneveOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o SystemGeneveOutput) Vni() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGeneve) pulumi.IntOutput { return v.Vni }).(pulumi.IntOutput)
}

type SystemGeneveArrayOutput struct{ *pulumi.OutputState }

func (SystemGeneveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGeneve)(nil)).Elem()
}

func (o SystemGeneveArrayOutput) ToSystemGeneveArrayOutput() SystemGeneveArrayOutput {
	return o
}

func (o SystemGeneveArrayOutput) ToSystemGeneveArrayOutputWithContext(ctx context.Context) SystemGeneveArrayOutput {
	return o
}

func (o SystemGeneveArrayOutput) Index(i pulumi.IntInput) SystemGeneveOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemGeneve {
		return vs[0].([]*SystemGeneve)[vs[1].(int)]
	}).(SystemGeneveOutput)
}

type SystemGeneveMapOutput struct{ *pulumi.OutputState }

func (SystemGeneveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGeneve)(nil)).Elem()
}

func (o SystemGeneveMapOutput) ToSystemGeneveMapOutput() SystemGeneveMapOutput {
	return o
}

func (o SystemGeneveMapOutput) ToSystemGeneveMapOutputWithContext(ctx context.Context) SystemGeneveMapOutput {
	return o
}

func (o SystemGeneveMapOutput) MapIndex(k pulumi.StringInput) SystemGeneveOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemGeneve {
		return vs[0].(map[string]*SystemGeneve)[vs[1].(string)]
	}).(SystemGeneveOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeneveInput)(nil)).Elem(), &SystemGeneve{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeneveArrayInput)(nil)).Elem(), SystemGeneveArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeneveMapInput)(nil)).Elem(), SystemGeneveMap{})
	pulumi.RegisterOutputType(SystemGeneveOutput{})
	pulumi.RegisterOutputType(SystemGeneveArrayOutput{})
	pulumi.RegisterOutputType(SystemGeneveMapOutput{})
}
