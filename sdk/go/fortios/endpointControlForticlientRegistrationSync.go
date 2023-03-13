// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EndpointControlForticlientRegistrationSync struct {
	pulumi.CustomResourceState

	PeerIp    pulumi.StringOutput    `pulumi:"peerIp"`
	PeerName  pulumi.StringOutput    `pulumi:"peerName"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEndpointControlForticlientRegistrationSync registers a new resource with the given unique name, arguments, and options.
func NewEndpointControlForticlientRegistrationSync(ctx *pulumi.Context,
	name string, args *EndpointControlForticlientRegistrationSyncArgs, opts ...pulumi.ResourceOption) (*EndpointControlForticlientRegistrationSync, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerIp == nil {
		return nil, errors.New("invalid value for required argument 'PeerIp'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EndpointControlForticlientRegistrationSync
	err := ctx.RegisterResource("fortios:index/endpointControlForticlientRegistrationSync:EndpointControlForticlientRegistrationSync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointControlForticlientRegistrationSync gets an existing EndpointControlForticlientRegistrationSync resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointControlForticlientRegistrationSync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointControlForticlientRegistrationSyncState, opts ...pulumi.ResourceOption) (*EndpointControlForticlientRegistrationSync, error) {
	var resource EndpointControlForticlientRegistrationSync
	err := ctx.ReadResource("fortios:index/endpointControlForticlientRegistrationSync:EndpointControlForticlientRegistrationSync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointControlForticlientRegistrationSync resources.
type endpointControlForticlientRegistrationSyncState struct {
	PeerIp    *string `pulumi:"peerIp"`
	PeerName  *string `pulumi:"peerName"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type EndpointControlForticlientRegistrationSyncState struct {
	PeerIp    pulumi.StringPtrInput
	PeerName  pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (EndpointControlForticlientRegistrationSyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlForticlientRegistrationSyncState)(nil)).Elem()
}

type endpointControlForticlientRegistrationSyncArgs struct {
	PeerIp    string  `pulumi:"peerIp"`
	PeerName  *string `pulumi:"peerName"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EndpointControlForticlientRegistrationSync resource.
type EndpointControlForticlientRegistrationSyncArgs struct {
	PeerIp    pulumi.StringInput
	PeerName  pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (EndpointControlForticlientRegistrationSyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlForticlientRegistrationSyncArgs)(nil)).Elem()
}

type EndpointControlForticlientRegistrationSyncInput interface {
	pulumi.Input

	ToEndpointControlForticlientRegistrationSyncOutput() EndpointControlForticlientRegistrationSyncOutput
	ToEndpointControlForticlientRegistrationSyncOutputWithContext(ctx context.Context) EndpointControlForticlientRegistrationSyncOutput
}

func (*EndpointControlForticlientRegistrationSync) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlForticlientRegistrationSync)(nil)).Elem()
}

func (i *EndpointControlForticlientRegistrationSync) ToEndpointControlForticlientRegistrationSyncOutput() EndpointControlForticlientRegistrationSyncOutput {
	return i.ToEndpointControlForticlientRegistrationSyncOutputWithContext(context.Background())
}

func (i *EndpointControlForticlientRegistrationSync) ToEndpointControlForticlientRegistrationSyncOutputWithContext(ctx context.Context) EndpointControlForticlientRegistrationSyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlForticlientRegistrationSyncOutput)
}

// EndpointControlForticlientRegistrationSyncArrayInput is an input type that accepts EndpointControlForticlientRegistrationSyncArray and EndpointControlForticlientRegistrationSyncArrayOutput values.
// You can construct a concrete instance of `EndpointControlForticlientRegistrationSyncArrayInput` via:
//
//	EndpointControlForticlientRegistrationSyncArray{ EndpointControlForticlientRegistrationSyncArgs{...} }
type EndpointControlForticlientRegistrationSyncArrayInput interface {
	pulumi.Input

	ToEndpointControlForticlientRegistrationSyncArrayOutput() EndpointControlForticlientRegistrationSyncArrayOutput
	ToEndpointControlForticlientRegistrationSyncArrayOutputWithContext(context.Context) EndpointControlForticlientRegistrationSyncArrayOutput
}

type EndpointControlForticlientRegistrationSyncArray []EndpointControlForticlientRegistrationSyncInput

func (EndpointControlForticlientRegistrationSyncArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointControlForticlientRegistrationSync)(nil)).Elem()
}

func (i EndpointControlForticlientRegistrationSyncArray) ToEndpointControlForticlientRegistrationSyncArrayOutput() EndpointControlForticlientRegistrationSyncArrayOutput {
	return i.ToEndpointControlForticlientRegistrationSyncArrayOutputWithContext(context.Background())
}

func (i EndpointControlForticlientRegistrationSyncArray) ToEndpointControlForticlientRegistrationSyncArrayOutputWithContext(ctx context.Context) EndpointControlForticlientRegistrationSyncArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlForticlientRegistrationSyncArrayOutput)
}

// EndpointControlForticlientRegistrationSyncMapInput is an input type that accepts EndpointControlForticlientRegistrationSyncMap and EndpointControlForticlientRegistrationSyncMapOutput values.
// You can construct a concrete instance of `EndpointControlForticlientRegistrationSyncMapInput` via:
//
//	EndpointControlForticlientRegistrationSyncMap{ "key": EndpointControlForticlientRegistrationSyncArgs{...} }
type EndpointControlForticlientRegistrationSyncMapInput interface {
	pulumi.Input

	ToEndpointControlForticlientRegistrationSyncMapOutput() EndpointControlForticlientRegistrationSyncMapOutput
	ToEndpointControlForticlientRegistrationSyncMapOutputWithContext(context.Context) EndpointControlForticlientRegistrationSyncMapOutput
}

type EndpointControlForticlientRegistrationSyncMap map[string]EndpointControlForticlientRegistrationSyncInput

func (EndpointControlForticlientRegistrationSyncMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointControlForticlientRegistrationSync)(nil)).Elem()
}

func (i EndpointControlForticlientRegistrationSyncMap) ToEndpointControlForticlientRegistrationSyncMapOutput() EndpointControlForticlientRegistrationSyncMapOutput {
	return i.ToEndpointControlForticlientRegistrationSyncMapOutputWithContext(context.Background())
}

func (i EndpointControlForticlientRegistrationSyncMap) ToEndpointControlForticlientRegistrationSyncMapOutputWithContext(ctx context.Context) EndpointControlForticlientRegistrationSyncMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlForticlientRegistrationSyncMapOutput)
}

type EndpointControlForticlientRegistrationSyncOutput struct{ *pulumi.OutputState }

func (EndpointControlForticlientRegistrationSyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlForticlientRegistrationSync)(nil)).Elem()
}

func (o EndpointControlForticlientRegistrationSyncOutput) ToEndpointControlForticlientRegistrationSyncOutput() EndpointControlForticlientRegistrationSyncOutput {
	return o
}

func (o EndpointControlForticlientRegistrationSyncOutput) ToEndpointControlForticlientRegistrationSyncOutputWithContext(ctx context.Context) EndpointControlForticlientRegistrationSyncOutput {
	return o
}

func (o EndpointControlForticlientRegistrationSyncOutput) PeerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlForticlientRegistrationSync) pulumi.StringOutput { return v.PeerIp }).(pulumi.StringOutput)
}

func (o EndpointControlForticlientRegistrationSyncOutput) PeerName() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlForticlientRegistrationSync) pulumi.StringOutput { return v.PeerName }).(pulumi.StringOutput)
}

func (o EndpointControlForticlientRegistrationSyncOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointControlForticlientRegistrationSync) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type EndpointControlForticlientRegistrationSyncArrayOutput struct{ *pulumi.OutputState }

func (EndpointControlForticlientRegistrationSyncArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointControlForticlientRegistrationSync)(nil)).Elem()
}

func (o EndpointControlForticlientRegistrationSyncArrayOutput) ToEndpointControlForticlientRegistrationSyncArrayOutput() EndpointControlForticlientRegistrationSyncArrayOutput {
	return o
}

func (o EndpointControlForticlientRegistrationSyncArrayOutput) ToEndpointControlForticlientRegistrationSyncArrayOutputWithContext(ctx context.Context) EndpointControlForticlientRegistrationSyncArrayOutput {
	return o
}

func (o EndpointControlForticlientRegistrationSyncArrayOutput) Index(i pulumi.IntInput) EndpointControlForticlientRegistrationSyncOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointControlForticlientRegistrationSync {
		return vs[0].([]*EndpointControlForticlientRegistrationSync)[vs[1].(int)]
	}).(EndpointControlForticlientRegistrationSyncOutput)
}

type EndpointControlForticlientRegistrationSyncMapOutput struct{ *pulumi.OutputState }

func (EndpointControlForticlientRegistrationSyncMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointControlForticlientRegistrationSync)(nil)).Elem()
}

func (o EndpointControlForticlientRegistrationSyncMapOutput) ToEndpointControlForticlientRegistrationSyncMapOutput() EndpointControlForticlientRegistrationSyncMapOutput {
	return o
}

func (o EndpointControlForticlientRegistrationSyncMapOutput) ToEndpointControlForticlientRegistrationSyncMapOutputWithContext(ctx context.Context) EndpointControlForticlientRegistrationSyncMapOutput {
	return o
}

func (o EndpointControlForticlientRegistrationSyncMapOutput) MapIndex(k pulumi.StringInput) EndpointControlForticlientRegistrationSyncOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointControlForticlientRegistrationSync {
		return vs[0].(map[string]*EndpointControlForticlientRegistrationSync)[vs[1].(string)]
	}).(EndpointControlForticlientRegistrationSyncOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlForticlientRegistrationSyncInput)(nil)).Elem(), &EndpointControlForticlientRegistrationSync{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlForticlientRegistrationSyncArrayInput)(nil)).Elem(), EndpointControlForticlientRegistrationSyncArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlForticlientRegistrationSyncMapInput)(nil)).Elem(), EndpointControlForticlientRegistrationSyncMap{})
	pulumi.RegisterOutputType(EndpointControlForticlientRegistrationSyncOutput{})
	pulumi.RegisterOutputType(EndpointControlForticlientRegistrationSyncArrayOutput{})
	pulumi.RegisterOutputType(EndpointControlForticlientRegistrationSyncMapOutput{})
}
