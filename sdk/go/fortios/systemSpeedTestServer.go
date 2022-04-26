// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure speed test server list. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// System SpeedTestServer can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemSpeedTestServer:SystemSpeedTestServer labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemSpeedTestServer:SystemSpeedTestServer labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSpeedTestServer struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Hosts of the server. The structure of `host` block is documented below.
	Hosts SystemSpeedTestServerHostArrayOutput `pulumi:"hosts"`
	// Speed test server name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Speed test server timestamp.
	Timestamp pulumi.IntOutput `pulumi:"timestamp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSpeedTestServer registers a new resource with the given unique name, arguments, and options.
func NewSystemSpeedTestServer(ctx *pulumi.Context,
	name string, args *SystemSpeedTestServerArgs, opts ...pulumi.ResourceOption) (*SystemSpeedTestServer, error) {
	if args == nil {
		args = &SystemSpeedTestServerArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSpeedTestServer
	err := ctx.RegisterResource("fortios:index/systemSpeedTestServer:SystemSpeedTestServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSpeedTestServer gets an existing SystemSpeedTestServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSpeedTestServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSpeedTestServerState, opts ...pulumi.ResourceOption) (*SystemSpeedTestServer, error) {
	var resource SystemSpeedTestServer
	err := ctx.ReadResource("fortios:index/systemSpeedTestServer:SystemSpeedTestServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSpeedTestServer resources.
type systemSpeedTestServerState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Hosts of the server. The structure of `host` block is documented below.
	Hosts []SystemSpeedTestServerHost `pulumi:"hosts"`
	// Speed test server name.
	Name *string `pulumi:"name"`
	// Speed test server timestamp.
	Timestamp *int `pulumi:"timestamp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSpeedTestServerState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Hosts of the server. The structure of `host` block is documented below.
	Hosts SystemSpeedTestServerHostArrayInput
	// Speed test server name.
	Name pulumi.StringPtrInput
	// Speed test server timestamp.
	Timestamp pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSpeedTestServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSpeedTestServerState)(nil)).Elem()
}

type systemSpeedTestServerArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Hosts of the server. The structure of `host` block is documented below.
	Hosts []SystemSpeedTestServerHost `pulumi:"hosts"`
	// Speed test server name.
	Name *string `pulumi:"name"`
	// Speed test server timestamp.
	Timestamp *int `pulumi:"timestamp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSpeedTestServer resource.
type SystemSpeedTestServerArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Hosts of the server. The structure of `host` block is documented below.
	Hosts SystemSpeedTestServerHostArrayInput
	// Speed test server name.
	Name pulumi.StringPtrInput
	// Speed test server timestamp.
	Timestamp pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSpeedTestServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSpeedTestServerArgs)(nil)).Elem()
}

type SystemSpeedTestServerInput interface {
	pulumi.Input

	ToSystemSpeedTestServerOutput() SystemSpeedTestServerOutput
	ToSystemSpeedTestServerOutputWithContext(ctx context.Context) SystemSpeedTestServerOutput
}

func (*SystemSpeedTestServer) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSpeedTestServer)(nil)).Elem()
}

func (i *SystemSpeedTestServer) ToSystemSpeedTestServerOutput() SystemSpeedTestServerOutput {
	return i.ToSystemSpeedTestServerOutputWithContext(context.Background())
}

func (i *SystemSpeedTestServer) ToSystemSpeedTestServerOutputWithContext(ctx context.Context) SystemSpeedTestServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedTestServerOutput)
}

// SystemSpeedTestServerArrayInput is an input type that accepts SystemSpeedTestServerArray and SystemSpeedTestServerArrayOutput values.
// You can construct a concrete instance of `SystemSpeedTestServerArrayInput` via:
//
//          SystemSpeedTestServerArray{ SystemSpeedTestServerArgs{...} }
type SystemSpeedTestServerArrayInput interface {
	pulumi.Input

	ToSystemSpeedTestServerArrayOutput() SystemSpeedTestServerArrayOutput
	ToSystemSpeedTestServerArrayOutputWithContext(context.Context) SystemSpeedTestServerArrayOutput
}

type SystemSpeedTestServerArray []SystemSpeedTestServerInput

func (SystemSpeedTestServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSpeedTestServer)(nil)).Elem()
}

func (i SystemSpeedTestServerArray) ToSystemSpeedTestServerArrayOutput() SystemSpeedTestServerArrayOutput {
	return i.ToSystemSpeedTestServerArrayOutputWithContext(context.Background())
}

func (i SystemSpeedTestServerArray) ToSystemSpeedTestServerArrayOutputWithContext(ctx context.Context) SystemSpeedTestServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedTestServerArrayOutput)
}

// SystemSpeedTestServerMapInput is an input type that accepts SystemSpeedTestServerMap and SystemSpeedTestServerMapOutput values.
// You can construct a concrete instance of `SystemSpeedTestServerMapInput` via:
//
//          SystemSpeedTestServerMap{ "key": SystemSpeedTestServerArgs{...} }
type SystemSpeedTestServerMapInput interface {
	pulumi.Input

	ToSystemSpeedTestServerMapOutput() SystemSpeedTestServerMapOutput
	ToSystemSpeedTestServerMapOutputWithContext(context.Context) SystemSpeedTestServerMapOutput
}

type SystemSpeedTestServerMap map[string]SystemSpeedTestServerInput

func (SystemSpeedTestServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSpeedTestServer)(nil)).Elem()
}

func (i SystemSpeedTestServerMap) ToSystemSpeedTestServerMapOutput() SystemSpeedTestServerMapOutput {
	return i.ToSystemSpeedTestServerMapOutputWithContext(context.Background())
}

func (i SystemSpeedTestServerMap) ToSystemSpeedTestServerMapOutputWithContext(ctx context.Context) SystemSpeedTestServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedTestServerMapOutput)
}

type SystemSpeedTestServerOutput struct{ *pulumi.OutputState }

func (SystemSpeedTestServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSpeedTestServer)(nil)).Elem()
}

func (o SystemSpeedTestServerOutput) ToSystemSpeedTestServerOutput() SystemSpeedTestServerOutput {
	return o
}

func (o SystemSpeedTestServerOutput) ToSystemSpeedTestServerOutputWithContext(ctx context.Context) SystemSpeedTestServerOutput {
	return o
}

type SystemSpeedTestServerArrayOutput struct{ *pulumi.OutputState }

func (SystemSpeedTestServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSpeedTestServer)(nil)).Elem()
}

func (o SystemSpeedTestServerArrayOutput) ToSystemSpeedTestServerArrayOutput() SystemSpeedTestServerArrayOutput {
	return o
}

func (o SystemSpeedTestServerArrayOutput) ToSystemSpeedTestServerArrayOutputWithContext(ctx context.Context) SystemSpeedTestServerArrayOutput {
	return o
}

func (o SystemSpeedTestServerArrayOutput) Index(i pulumi.IntInput) SystemSpeedTestServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSpeedTestServer {
		return vs[0].([]*SystemSpeedTestServer)[vs[1].(int)]
	}).(SystemSpeedTestServerOutput)
}

type SystemSpeedTestServerMapOutput struct{ *pulumi.OutputState }

func (SystemSpeedTestServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSpeedTestServer)(nil)).Elem()
}

func (o SystemSpeedTestServerMapOutput) ToSystemSpeedTestServerMapOutput() SystemSpeedTestServerMapOutput {
	return o
}

func (o SystemSpeedTestServerMapOutput) ToSystemSpeedTestServerMapOutputWithContext(ctx context.Context) SystemSpeedTestServerMapOutput {
	return o
}

func (o SystemSpeedTestServerMapOutput) MapIndex(k pulumi.StringInput) SystemSpeedTestServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSpeedTestServer {
		return vs[0].(map[string]*SystemSpeedTestServer)[vs[1].(string)]
	}).(SystemSpeedTestServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedTestServerInput)(nil)).Elem(), &SystemSpeedTestServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedTestServerArrayInput)(nil)).Elem(), SystemSpeedTestServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedTestServerMapInput)(nil)).Elem(), SystemSpeedTestServerMap{})
	pulumi.RegisterOutputType(SystemSpeedTestServerOutput{})
	pulumi.RegisterOutputType(SystemSpeedTestServerArrayOutput{})
	pulumi.RegisterOutputType(SystemSpeedTestServerMapOutput{})
}
