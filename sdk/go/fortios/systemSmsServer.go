// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type SystemSmsServer struct {
	pulumi.CustomResourceState

	MailServer pulumi.StringOutput    `pulumi:"mailServer"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSmsServer registers a new resource with the given unique name, arguments, and options.
func NewSystemSmsServer(ctx *pulumi.Context,
	name string, args *SystemSmsServerArgs, opts ...pulumi.ResourceOption) (*SystemSmsServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MailServer == nil {
		return nil, errors.New("invalid value for required argument 'MailServer'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemSmsServer
	err := ctx.RegisterResource("fortios:index/systemSmsServer:SystemSmsServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSmsServer gets an existing SystemSmsServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSmsServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSmsServerState, opts ...pulumi.ResourceOption) (*SystemSmsServer, error) {
	var resource SystemSmsServer
	err := ctx.ReadResource("fortios:index/systemSmsServer:SystemSmsServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSmsServer resources.
type systemSmsServerState struct {
	MailServer *string `pulumi:"mailServer"`
	Name       *string `pulumi:"name"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

type SystemSmsServerState struct {
	MailServer pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (SystemSmsServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSmsServerState)(nil)).Elem()
}

type systemSmsServerArgs struct {
	MailServer string  `pulumi:"mailServer"`
	Name       *string `pulumi:"name"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSmsServer resource.
type SystemSmsServerArgs struct {
	MailServer pulumi.StringInput
	Name       pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (SystemSmsServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSmsServerArgs)(nil)).Elem()
}

type SystemSmsServerInput interface {
	pulumi.Input

	ToSystemSmsServerOutput() SystemSmsServerOutput
	ToSystemSmsServerOutputWithContext(ctx context.Context) SystemSmsServerOutput
}

func (*SystemSmsServer) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSmsServer)(nil)).Elem()
}

func (i *SystemSmsServer) ToSystemSmsServerOutput() SystemSmsServerOutput {
	return i.ToSystemSmsServerOutputWithContext(context.Background())
}

func (i *SystemSmsServer) ToSystemSmsServerOutputWithContext(ctx context.Context) SystemSmsServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSmsServerOutput)
}

func (i *SystemSmsServer) ToOutput(ctx context.Context) pulumix.Output[*SystemSmsServer] {
	return pulumix.Output[*SystemSmsServer]{
		OutputState: i.ToSystemSmsServerOutputWithContext(ctx).OutputState,
	}
}

// SystemSmsServerArrayInput is an input type that accepts SystemSmsServerArray and SystemSmsServerArrayOutput values.
// You can construct a concrete instance of `SystemSmsServerArrayInput` via:
//
//	SystemSmsServerArray{ SystemSmsServerArgs{...} }
type SystemSmsServerArrayInput interface {
	pulumi.Input

	ToSystemSmsServerArrayOutput() SystemSmsServerArrayOutput
	ToSystemSmsServerArrayOutputWithContext(context.Context) SystemSmsServerArrayOutput
}

type SystemSmsServerArray []SystemSmsServerInput

func (SystemSmsServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSmsServer)(nil)).Elem()
}

func (i SystemSmsServerArray) ToSystemSmsServerArrayOutput() SystemSmsServerArrayOutput {
	return i.ToSystemSmsServerArrayOutputWithContext(context.Background())
}

func (i SystemSmsServerArray) ToSystemSmsServerArrayOutputWithContext(ctx context.Context) SystemSmsServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSmsServerArrayOutput)
}

func (i SystemSmsServerArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemSmsServer] {
	return pulumix.Output[[]*SystemSmsServer]{
		OutputState: i.ToSystemSmsServerArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemSmsServerMapInput is an input type that accepts SystemSmsServerMap and SystemSmsServerMapOutput values.
// You can construct a concrete instance of `SystemSmsServerMapInput` via:
//
//	SystemSmsServerMap{ "key": SystemSmsServerArgs{...} }
type SystemSmsServerMapInput interface {
	pulumi.Input

	ToSystemSmsServerMapOutput() SystemSmsServerMapOutput
	ToSystemSmsServerMapOutputWithContext(context.Context) SystemSmsServerMapOutput
}

type SystemSmsServerMap map[string]SystemSmsServerInput

func (SystemSmsServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSmsServer)(nil)).Elem()
}

func (i SystemSmsServerMap) ToSystemSmsServerMapOutput() SystemSmsServerMapOutput {
	return i.ToSystemSmsServerMapOutputWithContext(context.Background())
}

func (i SystemSmsServerMap) ToSystemSmsServerMapOutputWithContext(ctx context.Context) SystemSmsServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSmsServerMapOutput)
}

func (i SystemSmsServerMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemSmsServer] {
	return pulumix.Output[map[string]*SystemSmsServer]{
		OutputState: i.ToSystemSmsServerMapOutputWithContext(ctx).OutputState,
	}
}

type SystemSmsServerOutput struct{ *pulumi.OutputState }

func (SystemSmsServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSmsServer)(nil)).Elem()
}

func (o SystemSmsServerOutput) ToSystemSmsServerOutput() SystemSmsServerOutput {
	return o
}

func (o SystemSmsServerOutput) ToSystemSmsServerOutputWithContext(ctx context.Context) SystemSmsServerOutput {
	return o
}

func (o SystemSmsServerOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemSmsServer] {
	return pulumix.Output[*SystemSmsServer]{
		OutputState: o.OutputState,
	}
}

func (o SystemSmsServerOutput) MailServer() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSmsServer) pulumi.StringOutput { return v.MailServer }).(pulumi.StringOutput)
}

func (o SystemSmsServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSmsServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemSmsServerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSmsServer) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemSmsServerArrayOutput struct{ *pulumi.OutputState }

func (SystemSmsServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSmsServer)(nil)).Elem()
}

func (o SystemSmsServerArrayOutput) ToSystemSmsServerArrayOutput() SystemSmsServerArrayOutput {
	return o
}

func (o SystemSmsServerArrayOutput) ToSystemSmsServerArrayOutputWithContext(ctx context.Context) SystemSmsServerArrayOutput {
	return o
}

func (o SystemSmsServerArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemSmsServer] {
	return pulumix.Output[[]*SystemSmsServer]{
		OutputState: o.OutputState,
	}
}

func (o SystemSmsServerArrayOutput) Index(i pulumi.IntInput) SystemSmsServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSmsServer {
		return vs[0].([]*SystemSmsServer)[vs[1].(int)]
	}).(SystemSmsServerOutput)
}

type SystemSmsServerMapOutput struct{ *pulumi.OutputState }

func (SystemSmsServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSmsServer)(nil)).Elem()
}

func (o SystemSmsServerMapOutput) ToSystemSmsServerMapOutput() SystemSmsServerMapOutput {
	return o
}

func (o SystemSmsServerMapOutput) ToSystemSmsServerMapOutputWithContext(ctx context.Context) SystemSmsServerMapOutput {
	return o
}

func (o SystemSmsServerMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemSmsServer] {
	return pulumix.Output[map[string]*SystemSmsServer]{
		OutputState: o.OutputState,
	}
}

func (o SystemSmsServerMapOutput) MapIndex(k pulumi.StringInput) SystemSmsServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSmsServer {
		return vs[0].(map[string]*SystemSmsServer)[vs[1].(string)]
	}).(SystemSmsServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSmsServerInput)(nil)).Elem(), &SystemSmsServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSmsServerArrayInput)(nil)).Elem(), SystemSmsServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSmsServerMapInput)(nil)).Elem(), SystemSmsServerMap{})
	pulumi.RegisterOutputType(SystemSmsServerOutput{})
	pulumi.RegisterOutputType(SystemSmsServerArrayOutput{})
	pulumi.RegisterOutputType(SystemSmsServerMapOutput{})
}
