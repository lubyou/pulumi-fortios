// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.
//
// ## Import
//
// System VdomRadiusServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemVdomRadiusServer:SystemVdomRadiusServer labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemVdomRadiusServer struct {
	pulumi.CustomResourceState

	// Name of the VDOM that you are adding the RADIUS server to.
	Name pulumi.StringOutput `pulumi:"name"`
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom pulumi.StringOutput `pulumi:"radiusServerVdom"`
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdomRadiusServer registers a new resource with the given unique name, arguments, and options.
func NewSystemVdomRadiusServer(ctx *pulumi.Context,
	name string, args *SystemVdomRadiusServerArgs, opts ...pulumi.ResourceOption) (*SystemVdomRadiusServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RadiusServerVdom == nil {
		return nil, errors.New("invalid value for required argument 'RadiusServerVdom'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemVdomRadiusServer
	err := ctx.RegisterResource("fortios:index/systemVdomRadiusServer:SystemVdomRadiusServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdomRadiusServer gets an existing SystemVdomRadiusServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdomRadiusServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdomRadiusServerState, opts ...pulumi.ResourceOption) (*SystemVdomRadiusServer, error) {
	var resource SystemVdomRadiusServer
	err := ctx.ReadResource("fortios:index/systemVdomRadiusServer:SystemVdomRadiusServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdomRadiusServer resources.
type systemVdomRadiusServerState struct {
	// Name of the VDOM that you are adding the RADIUS server to.
	Name *string `pulumi:"name"`
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom *string `pulumi:"radiusServerVdom"`
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemVdomRadiusServerState struct {
	// Name of the VDOM that you are adding the RADIUS server to.
	Name pulumi.StringPtrInput
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom pulumi.StringPtrInput
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomRadiusServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomRadiusServerState)(nil)).Elem()
}

type systemVdomRadiusServerArgs struct {
	// Name of the VDOM that you are adding the RADIUS server to.
	Name *string `pulumi:"name"`
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom string `pulumi:"radiusServerVdom"`
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdomRadiusServer resource.
type SystemVdomRadiusServerArgs struct {
	// Name of the VDOM that you are adding the RADIUS server to.
	Name pulumi.StringPtrInput
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom pulumi.StringInput
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomRadiusServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomRadiusServerArgs)(nil)).Elem()
}

type SystemVdomRadiusServerInput interface {
	pulumi.Input

	ToSystemVdomRadiusServerOutput() SystemVdomRadiusServerOutput
	ToSystemVdomRadiusServerOutputWithContext(ctx context.Context) SystemVdomRadiusServerOutput
}

func (*SystemVdomRadiusServer) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomRadiusServer)(nil)).Elem()
}

func (i *SystemVdomRadiusServer) ToSystemVdomRadiusServerOutput() SystemVdomRadiusServerOutput {
	return i.ToSystemVdomRadiusServerOutputWithContext(context.Background())
}

func (i *SystemVdomRadiusServer) ToSystemVdomRadiusServerOutputWithContext(ctx context.Context) SystemVdomRadiusServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomRadiusServerOutput)
}

// SystemVdomRadiusServerArrayInput is an input type that accepts SystemVdomRadiusServerArray and SystemVdomRadiusServerArrayOutput values.
// You can construct a concrete instance of `SystemVdomRadiusServerArrayInput` via:
//
//          SystemVdomRadiusServerArray{ SystemVdomRadiusServerArgs{...} }
type SystemVdomRadiusServerArrayInput interface {
	pulumi.Input

	ToSystemVdomRadiusServerArrayOutput() SystemVdomRadiusServerArrayOutput
	ToSystemVdomRadiusServerArrayOutputWithContext(context.Context) SystemVdomRadiusServerArrayOutput
}

type SystemVdomRadiusServerArray []SystemVdomRadiusServerInput

func (SystemVdomRadiusServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomRadiusServer)(nil)).Elem()
}

func (i SystemVdomRadiusServerArray) ToSystemVdomRadiusServerArrayOutput() SystemVdomRadiusServerArrayOutput {
	return i.ToSystemVdomRadiusServerArrayOutputWithContext(context.Background())
}

func (i SystemVdomRadiusServerArray) ToSystemVdomRadiusServerArrayOutputWithContext(ctx context.Context) SystemVdomRadiusServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomRadiusServerArrayOutput)
}

// SystemVdomRadiusServerMapInput is an input type that accepts SystemVdomRadiusServerMap and SystemVdomRadiusServerMapOutput values.
// You can construct a concrete instance of `SystemVdomRadiusServerMapInput` via:
//
//          SystemVdomRadiusServerMap{ "key": SystemVdomRadiusServerArgs{...} }
type SystemVdomRadiusServerMapInput interface {
	pulumi.Input

	ToSystemVdomRadiusServerMapOutput() SystemVdomRadiusServerMapOutput
	ToSystemVdomRadiusServerMapOutputWithContext(context.Context) SystemVdomRadiusServerMapOutput
}

type SystemVdomRadiusServerMap map[string]SystemVdomRadiusServerInput

func (SystemVdomRadiusServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomRadiusServer)(nil)).Elem()
}

func (i SystemVdomRadiusServerMap) ToSystemVdomRadiusServerMapOutput() SystemVdomRadiusServerMapOutput {
	return i.ToSystemVdomRadiusServerMapOutputWithContext(context.Background())
}

func (i SystemVdomRadiusServerMap) ToSystemVdomRadiusServerMapOutputWithContext(ctx context.Context) SystemVdomRadiusServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomRadiusServerMapOutput)
}

type SystemVdomRadiusServerOutput struct{ *pulumi.OutputState }

func (SystemVdomRadiusServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomRadiusServer)(nil)).Elem()
}

func (o SystemVdomRadiusServerOutput) ToSystemVdomRadiusServerOutput() SystemVdomRadiusServerOutput {
	return o
}

func (o SystemVdomRadiusServerOutput) ToSystemVdomRadiusServerOutputWithContext(ctx context.Context) SystemVdomRadiusServerOutput {
	return o
}

type SystemVdomRadiusServerArrayOutput struct{ *pulumi.OutputState }

func (SystemVdomRadiusServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomRadiusServer)(nil)).Elem()
}

func (o SystemVdomRadiusServerArrayOutput) ToSystemVdomRadiusServerArrayOutput() SystemVdomRadiusServerArrayOutput {
	return o
}

func (o SystemVdomRadiusServerArrayOutput) ToSystemVdomRadiusServerArrayOutputWithContext(ctx context.Context) SystemVdomRadiusServerArrayOutput {
	return o
}

func (o SystemVdomRadiusServerArrayOutput) Index(i pulumi.IntInput) SystemVdomRadiusServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVdomRadiusServer {
		return vs[0].([]*SystemVdomRadiusServer)[vs[1].(int)]
	}).(SystemVdomRadiusServerOutput)
}

type SystemVdomRadiusServerMapOutput struct{ *pulumi.OutputState }

func (SystemVdomRadiusServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomRadiusServer)(nil)).Elem()
}

func (o SystemVdomRadiusServerMapOutput) ToSystemVdomRadiusServerMapOutput() SystemVdomRadiusServerMapOutput {
	return o
}

func (o SystemVdomRadiusServerMapOutput) ToSystemVdomRadiusServerMapOutputWithContext(ctx context.Context) SystemVdomRadiusServerMapOutput {
	return o
}

func (o SystemVdomRadiusServerMapOutput) MapIndex(k pulumi.StringInput) SystemVdomRadiusServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVdomRadiusServer {
		return vs[0].(map[string]*SystemVdomRadiusServer)[vs[1].(string)]
	}).(SystemVdomRadiusServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomRadiusServerInput)(nil)).Elem(), &SystemVdomRadiusServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomRadiusServerArrayInput)(nil)).Elem(), SystemVdomRadiusServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomRadiusServerMapInput)(nil)).Elem(), SystemVdomRadiusServerMap{})
	pulumi.RegisterOutputType(SystemVdomRadiusServerOutput{})
	pulumi.RegisterOutputType(SystemVdomRadiusServerArrayOutput{})
	pulumi.RegisterOutputType(SystemVdomRadiusServerMapOutput{})
}
