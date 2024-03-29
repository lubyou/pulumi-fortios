// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemFssoPolling struct {
	pulumi.CustomResourceState

	AuthPassword   pulumi.StringPtrOutput `pulumi:"authPassword"`
	Authentication pulumi.StringOutput    `pulumi:"authentication"`
	ListeningPort  pulumi.IntOutput       `pulumi:"listeningPort"`
	Status         pulumi.StringOutput    `pulumi:"status"`
	Vdomparam      pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFssoPolling registers a new resource with the given unique name, arguments, and options.
func NewSystemFssoPolling(ctx *pulumi.Context,
	name string, args *SystemFssoPollingArgs, opts ...pulumi.ResourceOption) (*SystemFssoPolling, error) {
	if args == nil {
		args = &SystemFssoPollingArgs{}
	}

	if args.AuthPassword != nil {
		args.AuthPassword = pulumi.ToSecret(args.AuthPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemFssoPolling
	err := ctx.RegisterResource("fortios:index/systemFssoPolling:SystemFssoPolling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFssoPolling gets an existing SystemFssoPolling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFssoPolling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFssoPollingState, opts ...pulumi.ResourceOption) (*SystemFssoPolling, error) {
	var resource SystemFssoPolling
	err := ctx.ReadResource("fortios:index/systemFssoPolling:SystemFssoPolling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFssoPolling resources.
type systemFssoPollingState struct {
	AuthPassword   *string `pulumi:"authPassword"`
	Authentication *string `pulumi:"authentication"`
	ListeningPort  *int    `pulumi:"listeningPort"`
	Status         *string `pulumi:"status"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

type SystemFssoPollingState struct {
	AuthPassword   pulumi.StringPtrInput
	Authentication pulumi.StringPtrInput
	ListeningPort  pulumi.IntPtrInput
	Status         pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (SystemFssoPollingState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFssoPollingState)(nil)).Elem()
}

type systemFssoPollingArgs struct {
	AuthPassword   *string `pulumi:"authPassword"`
	Authentication *string `pulumi:"authentication"`
	ListeningPort  *int    `pulumi:"listeningPort"`
	Status         *string `pulumi:"status"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFssoPolling resource.
type SystemFssoPollingArgs struct {
	AuthPassword   pulumi.StringPtrInput
	Authentication pulumi.StringPtrInput
	ListeningPort  pulumi.IntPtrInput
	Status         pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (SystemFssoPollingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFssoPollingArgs)(nil)).Elem()
}

type SystemFssoPollingInput interface {
	pulumi.Input

	ToSystemFssoPollingOutput() SystemFssoPollingOutput
	ToSystemFssoPollingOutputWithContext(ctx context.Context) SystemFssoPollingOutput
}

func (*SystemFssoPolling) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFssoPolling)(nil)).Elem()
}

func (i *SystemFssoPolling) ToSystemFssoPollingOutput() SystemFssoPollingOutput {
	return i.ToSystemFssoPollingOutputWithContext(context.Background())
}

func (i *SystemFssoPolling) ToSystemFssoPollingOutputWithContext(ctx context.Context) SystemFssoPollingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFssoPollingOutput)
}

// SystemFssoPollingArrayInput is an input type that accepts SystemFssoPollingArray and SystemFssoPollingArrayOutput values.
// You can construct a concrete instance of `SystemFssoPollingArrayInput` via:
//
//	SystemFssoPollingArray{ SystemFssoPollingArgs{...} }
type SystemFssoPollingArrayInput interface {
	pulumi.Input

	ToSystemFssoPollingArrayOutput() SystemFssoPollingArrayOutput
	ToSystemFssoPollingArrayOutputWithContext(context.Context) SystemFssoPollingArrayOutput
}

type SystemFssoPollingArray []SystemFssoPollingInput

func (SystemFssoPollingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFssoPolling)(nil)).Elem()
}

func (i SystemFssoPollingArray) ToSystemFssoPollingArrayOutput() SystemFssoPollingArrayOutput {
	return i.ToSystemFssoPollingArrayOutputWithContext(context.Background())
}

func (i SystemFssoPollingArray) ToSystemFssoPollingArrayOutputWithContext(ctx context.Context) SystemFssoPollingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFssoPollingArrayOutput)
}

// SystemFssoPollingMapInput is an input type that accepts SystemFssoPollingMap and SystemFssoPollingMapOutput values.
// You can construct a concrete instance of `SystemFssoPollingMapInput` via:
//
//	SystemFssoPollingMap{ "key": SystemFssoPollingArgs{...} }
type SystemFssoPollingMapInput interface {
	pulumi.Input

	ToSystemFssoPollingMapOutput() SystemFssoPollingMapOutput
	ToSystemFssoPollingMapOutputWithContext(context.Context) SystemFssoPollingMapOutput
}

type SystemFssoPollingMap map[string]SystemFssoPollingInput

func (SystemFssoPollingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFssoPolling)(nil)).Elem()
}

func (i SystemFssoPollingMap) ToSystemFssoPollingMapOutput() SystemFssoPollingMapOutput {
	return i.ToSystemFssoPollingMapOutputWithContext(context.Background())
}

func (i SystemFssoPollingMap) ToSystemFssoPollingMapOutputWithContext(ctx context.Context) SystemFssoPollingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFssoPollingMapOutput)
}

type SystemFssoPollingOutput struct{ *pulumi.OutputState }

func (SystemFssoPollingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFssoPolling)(nil)).Elem()
}

func (o SystemFssoPollingOutput) ToSystemFssoPollingOutput() SystemFssoPollingOutput {
	return o
}

func (o SystemFssoPollingOutput) ToSystemFssoPollingOutputWithContext(ctx context.Context) SystemFssoPollingOutput {
	return o
}

func (o SystemFssoPollingOutput) AuthPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFssoPolling) pulumi.StringPtrOutput { return v.AuthPassword }).(pulumi.StringPtrOutput)
}

func (o SystemFssoPollingOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFssoPolling) pulumi.StringOutput { return v.Authentication }).(pulumi.StringOutput)
}

func (o SystemFssoPollingOutput) ListeningPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemFssoPolling) pulumi.IntOutput { return v.ListeningPort }).(pulumi.IntOutput)
}

func (o SystemFssoPollingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFssoPolling) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SystemFssoPollingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFssoPolling) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemFssoPollingArrayOutput struct{ *pulumi.OutputState }

func (SystemFssoPollingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFssoPolling)(nil)).Elem()
}

func (o SystemFssoPollingArrayOutput) ToSystemFssoPollingArrayOutput() SystemFssoPollingArrayOutput {
	return o
}

func (o SystemFssoPollingArrayOutput) ToSystemFssoPollingArrayOutputWithContext(ctx context.Context) SystemFssoPollingArrayOutput {
	return o
}

func (o SystemFssoPollingArrayOutput) Index(i pulumi.IntInput) SystemFssoPollingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFssoPolling {
		return vs[0].([]*SystemFssoPolling)[vs[1].(int)]
	}).(SystemFssoPollingOutput)
}

type SystemFssoPollingMapOutput struct{ *pulumi.OutputState }

func (SystemFssoPollingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFssoPolling)(nil)).Elem()
}

func (o SystemFssoPollingMapOutput) ToSystemFssoPollingMapOutput() SystemFssoPollingMapOutput {
	return o
}

func (o SystemFssoPollingMapOutput) ToSystemFssoPollingMapOutputWithContext(ctx context.Context) SystemFssoPollingMapOutput {
	return o
}

func (o SystemFssoPollingMapOutput) MapIndex(k pulumi.StringInput) SystemFssoPollingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFssoPolling {
		return vs[0].(map[string]*SystemFssoPolling)[vs[1].(string)]
	}).(SystemFssoPollingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFssoPollingInput)(nil)).Elem(), &SystemFssoPolling{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFssoPollingArrayInput)(nil)).Elem(), SystemFssoPollingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFssoPollingMapInput)(nil)).Elem(), SystemFssoPollingMap{})
	pulumi.RegisterOutputType(SystemFssoPollingOutput{})
	pulumi.RegisterOutputType(SystemFssoPollingArrayOutput{})
	pulumi.RegisterOutputType(SystemFssoPollingMapOutput{})
}
