// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiClient registration synchronization settings. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewEndpointControlForticlientRegistrationSync(ctx, "trname", &fortios.EndpointControlForticlientRegistrationSyncArgs{
// 			PeerIp:   pulumi.String("1.1.1.1"),
// 			PeerName: pulumi.String("1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// EndpointControl ForticlientRegistrationSync can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/endpointControlForticlientRegistrationSync:EndpointControlForticlientRegistrationSync labelname {{peer_name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type EndpointControlForticlientRegistrationSync struct {
	pulumi.CustomResourceState

	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp pulumi.StringOutput `pulumi:"peerIp"`
	// Peer name.
	PeerName pulumi.StringOutput `pulumi:"peerName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
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
	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp *string `pulumi:"peerIp"`
	// Peer name.
	PeerName *string `pulumi:"peerName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EndpointControlForticlientRegistrationSyncState struct {
	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp pulumi.StringPtrInput
	// Peer name.
	PeerName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EndpointControlForticlientRegistrationSyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlForticlientRegistrationSyncState)(nil)).Elem()
}

type endpointControlForticlientRegistrationSyncArgs struct {
	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp string `pulumi:"peerIp"`
	// Peer name.
	PeerName *string `pulumi:"peerName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EndpointControlForticlientRegistrationSync resource.
type EndpointControlForticlientRegistrationSyncArgs struct {
	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp pulumi.StringInput
	// Peer name.
	PeerName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
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
//          EndpointControlForticlientRegistrationSyncArray{ EndpointControlForticlientRegistrationSyncArgs{...} }
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
//          EndpointControlForticlientRegistrationSyncMap{ "key": EndpointControlForticlientRegistrationSyncArgs{...} }
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
