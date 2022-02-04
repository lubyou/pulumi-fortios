// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure global session TTL timers for this FortiGate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSystemSessionTtl(ctx, "trname", &fortios.SystemSessionTtlArgs{
// 			Default: pulumi.String("3600"),
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
// System SessionTtl can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemSessionTtl:SystemSessionTtl labelname SystemSessionTtl
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSessionTtl struct {
	pulumi.CustomResourceState

	// Default timeout.
	Default pulumi.StringOutput `pulumi:"default"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Session TTL port. The structure of `port` block is documented below.
	Ports SystemSessionTtlPortArrayOutput `pulumi:"ports"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSessionTtl registers a new resource with the given unique name, arguments, and options.
func NewSystemSessionTtl(ctx *pulumi.Context,
	name string, args *SystemSessionTtlArgs, opts ...pulumi.ResourceOption) (*SystemSessionTtl, error) {
	if args == nil {
		args = &SystemSessionTtlArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSessionTtl
	err := ctx.RegisterResource("fortios:index/systemSessionTtl:SystemSessionTtl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSessionTtl gets an existing SystemSessionTtl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSessionTtl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSessionTtlState, opts ...pulumi.ResourceOption) (*SystemSessionTtl, error) {
	var resource SystemSessionTtl
	err := ctx.ReadResource("fortios:index/systemSessionTtl:SystemSessionTtl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSessionTtl resources.
type systemSessionTtlState struct {
	// Default timeout.
	Default *string `pulumi:"default"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Session TTL port. The structure of `port` block is documented below.
	Ports []SystemSessionTtlPort `pulumi:"ports"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSessionTtlState struct {
	// Default timeout.
	Default pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Session TTL port. The structure of `port` block is documented below.
	Ports SystemSessionTtlPortArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSessionTtlState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSessionTtlState)(nil)).Elem()
}

type systemSessionTtlArgs struct {
	// Default timeout.
	Default *string `pulumi:"default"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Session TTL port. The structure of `port` block is documented below.
	Ports []SystemSessionTtlPort `pulumi:"ports"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSessionTtl resource.
type SystemSessionTtlArgs struct {
	// Default timeout.
	Default pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Session TTL port. The structure of `port` block is documented below.
	Ports SystemSessionTtlPortArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSessionTtlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSessionTtlArgs)(nil)).Elem()
}

type SystemSessionTtlInput interface {
	pulumi.Input

	ToSystemSessionTtlOutput() SystemSessionTtlOutput
	ToSystemSessionTtlOutputWithContext(ctx context.Context) SystemSessionTtlOutput
}

func (*SystemSessionTtl) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSessionTtl)(nil)).Elem()
}

func (i *SystemSessionTtl) ToSystemSessionTtlOutput() SystemSessionTtlOutput {
	return i.ToSystemSessionTtlOutputWithContext(context.Background())
}

func (i *SystemSessionTtl) ToSystemSessionTtlOutputWithContext(ctx context.Context) SystemSessionTtlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSessionTtlOutput)
}

// SystemSessionTtlArrayInput is an input type that accepts SystemSessionTtlArray and SystemSessionTtlArrayOutput values.
// You can construct a concrete instance of `SystemSessionTtlArrayInput` via:
//
//          SystemSessionTtlArray{ SystemSessionTtlArgs{...} }
type SystemSessionTtlArrayInput interface {
	pulumi.Input

	ToSystemSessionTtlArrayOutput() SystemSessionTtlArrayOutput
	ToSystemSessionTtlArrayOutputWithContext(context.Context) SystemSessionTtlArrayOutput
}

type SystemSessionTtlArray []SystemSessionTtlInput

func (SystemSessionTtlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSessionTtl)(nil)).Elem()
}

func (i SystemSessionTtlArray) ToSystemSessionTtlArrayOutput() SystemSessionTtlArrayOutput {
	return i.ToSystemSessionTtlArrayOutputWithContext(context.Background())
}

func (i SystemSessionTtlArray) ToSystemSessionTtlArrayOutputWithContext(ctx context.Context) SystemSessionTtlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSessionTtlArrayOutput)
}

// SystemSessionTtlMapInput is an input type that accepts SystemSessionTtlMap and SystemSessionTtlMapOutput values.
// You can construct a concrete instance of `SystemSessionTtlMapInput` via:
//
//          SystemSessionTtlMap{ "key": SystemSessionTtlArgs{...} }
type SystemSessionTtlMapInput interface {
	pulumi.Input

	ToSystemSessionTtlMapOutput() SystemSessionTtlMapOutput
	ToSystemSessionTtlMapOutputWithContext(context.Context) SystemSessionTtlMapOutput
}

type SystemSessionTtlMap map[string]SystemSessionTtlInput

func (SystemSessionTtlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSessionTtl)(nil)).Elem()
}

func (i SystemSessionTtlMap) ToSystemSessionTtlMapOutput() SystemSessionTtlMapOutput {
	return i.ToSystemSessionTtlMapOutputWithContext(context.Background())
}

func (i SystemSessionTtlMap) ToSystemSessionTtlMapOutputWithContext(ctx context.Context) SystemSessionTtlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSessionTtlMapOutput)
}

type SystemSessionTtlOutput struct{ *pulumi.OutputState }

func (SystemSessionTtlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSessionTtl)(nil)).Elem()
}

func (o SystemSessionTtlOutput) ToSystemSessionTtlOutput() SystemSessionTtlOutput {
	return o
}

func (o SystemSessionTtlOutput) ToSystemSessionTtlOutputWithContext(ctx context.Context) SystemSessionTtlOutput {
	return o
}

type SystemSessionTtlArrayOutput struct{ *pulumi.OutputState }

func (SystemSessionTtlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSessionTtl)(nil)).Elem()
}

func (o SystemSessionTtlArrayOutput) ToSystemSessionTtlArrayOutput() SystemSessionTtlArrayOutput {
	return o
}

func (o SystemSessionTtlArrayOutput) ToSystemSessionTtlArrayOutputWithContext(ctx context.Context) SystemSessionTtlArrayOutput {
	return o
}

func (o SystemSessionTtlArrayOutput) Index(i pulumi.IntInput) SystemSessionTtlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSessionTtl {
		return vs[0].([]*SystemSessionTtl)[vs[1].(int)]
	}).(SystemSessionTtlOutput)
}

type SystemSessionTtlMapOutput struct{ *pulumi.OutputState }

func (SystemSessionTtlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSessionTtl)(nil)).Elem()
}

func (o SystemSessionTtlMapOutput) ToSystemSessionTtlMapOutput() SystemSessionTtlMapOutput {
	return o
}

func (o SystemSessionTtlMapOutput) ToSystemSessionTtlMapOutputWithContext(ctx context.Context) SystemSessionTtlMapOutput {
	return o
}

func (o SystemSessionTtlMapOutput) MapIndex(k pulumi.StringInput) SystemSessionTtlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSessionTtl {
		return vs[0].(map[string]*SystemSessionTtl)[vs[1].(string)]
	}).(SystemSessionTtlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSessionTtlInput)(nil)).Elem(), &SystemSessionTtl{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSessionTtlArrayInput)(nil)).Elem(), SystemSessionTtlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSessionTtlMapInput)(nil)).Elem(), SystemSessionTtlMap{})
	pulumi.RegisterOutputType(SystemSessionTtlOutput{})
	pulumi.RegisterOutputType(SystemSessionTtlArrayOutput{})
	pulumi.RegisterOutputType(SystemSessionTtlMapOutput{})
}
