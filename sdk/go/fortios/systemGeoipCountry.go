// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define geoip country name-ID table.
//
// ## Import
//
// System GeoipCountry can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemGeoipCountry:SystemGeoipCountry labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemGeoipCountry struct {
	pulumi.CustomResourceState

	// Country ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Country name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemGeoipCountry registers a new resource with the given unique name, arguments, and options.
func NewSystemGeoipCountry(ctx *pulumi.Context,
	name string, args *SystemGeoipCountryArgs, opts ...pulumi.ResourceOption) (*SystemGeoipCountry, error) {
	if args == nil {
		args = &SystemGeoipCountryArgs{}
	}

	var resource SystemGeoipCountry
	err := ctx.RegisterResource("fortios:index/systemGeoipCountry:SystemGeoipCountry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemGeoipCountry gets an existing SystemGeoipCountry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemGeoipCountry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemGeoipCountryState, opts ...pulumi.ResourceOption) (*SystemGeoipCountry, error) {
	var resource SystemGeoipCountry
	err := ctx.ReadResource("fortios:index/systemGeoipCountry:SystemGeoipCountry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemGeoipCountry resources.
type systemGeoipCountryState struct {
	// Country ID.
	Fosid *string `pulumi:"fosid"`
	// Country name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemGeoipCountryState struct {
	// Country ID.
	Fosid pulumi.StringPtrInput
	// Country name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemGeoipCountryState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGeoipCountryState)(nil)).Elem()
}

type systemGeoipCountryArgs struct {
	// Country ID.
	Fosid *string `pulumi:"fosid"`
	// Country name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemGeoipCountry resource.
type SystemGeoipCountryArgs struct {
	// Country ID.
	Fosid pulumi.StringPtrInput
	// Country name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemGeoipCountryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGeoipCountryArgs)(nil)).Elem()
}

type SystemGeoipCountryInput interface {
	pulumi.Input

	ToSystemGeoipCountryOutput() SystemGeoipCountryOutput
	ToSystemGeoipCountryOutputWithContext(ctx context.Context) SystemGeoipCountryOutput
}

func (*SystemGeoipCountry) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemGeoipCountry)(nil))
}

func (i *SystemGeoipCountry) ToSystemGeoipCountryOutput() SystemGeoipCountryOutput {
	return i.ToSystemGeoipCountryOutputWithContext(context.Background())
}

func (i *SystemGeoipCountry) ToSystemGeoipCountryOutputWithContext(ctx context.Context) SystemGeoipCountryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipCountryOutput)
}

func (i *SystemGeoipCountry) ToSystemGeoipCountryPtrOutput() SystemGeoipCountryPtrOutput {
	return i.ToSystemGeoipCountryPtrOutputWithContext(context.Background())
}

func (i *SystemGeoipCountry) ToSystemGeoipCountryPtrOutputWithContext(ctx context.Context) SystemGeoipCountryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipCountryPtrOutput)
}

type SystemGeoipCountryPtrInput interface {
	pulumi.Input

	ToSystemGeoipCountryPtrOutput() SystemGeoipCountryPtrOutput
	ToSystemGeoipCountryPtrOutputWithContext(ctx context.Context) SystemGeoipCountryPtrOutput
}

type systemGeoipCountryPtrType SystemGeoipCountryArgs

func (*systemGeoipCountryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGeoipCountry)(nil))
}

func (i *systemGeoipCountryPtrType) ToSystemGeoipCountryPtrOutput() SystemGeoipCountryPtrOutput {
	return i.ToSystemGeoipCountryPtrOutputWithContext(context.Background())
}

func (i *systemGeoipCountryPtrType) ToSystemGeoipCountryPtrOutputWithContext(ctx context.Context) SystemGeoipCountryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipCountryPtrOutput)
}

// SystemGeoipCountryArrayInput is an input type that accepts SystemGeoipCountryArray and SystemGeoipCountryArrayOutput values.
// You can construct a concrete instance of `SystemGeoipCountryArrayInput` via:
//
//          SystemGeoipCountryArray{ SystemGeoipCountryArgs{...} }
type SystemGeoipCountryArrayInput interface {
	pulumi.Input

	ToSystemGeoipCountryArrayOutput() SystemGeoipCountryArrayOutput
	ToSystemGeoipCountryArrayOutputWithContext(context.Context) SystemGeoipCountryArrayOutput
}

type SystemGeoipCountryArray []SystemGeoipCountryInput

func (SystemGeoipCountryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemGeoipCountry)(nil))
}

func (i SystemGeoipCountryArray) ToSystemGeoipCountryArrayOutput() SystemGeoipCountryArrayOutput {
	return i.ToSystemGeoipCountryArrayOutputWithContext(context.Background())
}

func (i SystemGeoipCountryArray) ToSystemGeoipCountryArrayOutputWithContext(ctx context.Context) SystemGeoipCountryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipCountryArrayOutput)
}

// SystemGeoipCountryMapInput is an input type that accepts SystemGeoipCountryMap and SystemGeoipCountryMapOutput values.
// You can construct a concrete instance of `SystemGeoipCountryMapInput` via:
//
//          SystemGeoipCountryMap{ "key": SystemGeoipCountryArgs{...} }
type SystemGeoipCountryMapInput interface {
	pulumi.Input

	ToSystemGeoipCountryMapOutput() SystemGeoipCountryMapOutput
	ToSystemGeoipCountryMapOutputWithContext(context.Context) SystemGeoipCountryMapOutput
}

type SystemGeoipCountryMap map[string]SystemGeoipCountryInput

func (SystemGeoipCountryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemGeoipCountry)(nil))
}

func (i SystemGeoipCountryMap) ToSystemGeoipCountryMapOutput() SystemGeoipCountryMapOutput {
	return i.ToSystemGeoipCountryMapOutputWithContext(context.Background())
}

func (i SystemGeoipCountryMap) ToSystemGeoipCountryMapOutputWithContext(ctx context.Context) SystemGeoipCountryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipCountryMapOutput)
}

type SystemGeoipCountryOutput struct {
	*pulumi.OutputState
}

func (SystemGeoipCountryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemGeoipCountry)(nil))
}

func (o SystemGeoipCountryOutput) ToSystemGeoipCountryOutput() SystemGeoipCountryOutput {
	return o
}

func (o SystemGeoipCountryOutput) ToSystemGeoipCountryOutputWithContext(ctx context.Context) SystemGeoipCountryOutput {
	return o
}

func (o SystemGeoipCountryOutput) ToSystemGeoipCountryPtrOutput() SystemGeoipCountryPtrOutput {
	return o.ToSystemGeoipCountryPtrOutputWithContext(context.Background())
}

func (o SystemGeoipCountryOutput) ToSystemGeoipCountryPtrOutputWithContext(ctx context.Context) SystemGeoipCountryPtrOutput {
	return o.ApplyT(func(v SystemGeoipCountry) *SystemGeoipCountry {
		return &v
	}).(SystemGeoipCountryPtrOutput)
}

type SystemGeoipCountryPtrOutput struct {
	*pulumi.OutputState
}

func (SystemGeoipCountryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGeoipCountry)(nil))
}

func (o SystemGeoipCountryPtrOutput) ToSystemGeoipCountryPtrOutput() SystemGeoipCountryPtrOutput {
	return o
}

func (o SystemGeoipCountryPtrOutput) ToSystemGeoipCountryPtrOutputWithContext(ctx context.Context) SystemGeoipCountryPtrOutput {
	return o
}

type SystemGeoipCountryArrayOutput struct{ *pulumi.OutputState }

func (SystemGeoipCountryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemGeoipCountry)(nil))
}

func (o SystemGeoipCountryArrayOutput) ToSystemGeoipCountryArrayOutput() SystemGeoipCountryArrayOutput {
	return o
}

func (o SystemGeoipCountryArrayOutput) ToSystemGeoipCountryArrayOutputWithContext(ctx context.Context) SystemGeoipCountryArrayOutput {
	return o
}

func (o SystemGeoipCountryArrayOutput) Index(i pulumi.IntInput) SystemGeoipCountryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemGeoipCountry {
		return vs[0].([]SystemGeoipCountry)[vs[1].(int)]
	}).(SystemGeoipCountryOutput)
}

type SystemGeoipCountryMapOutput struct{ *pulumi.OutputState }

func (SystemGeoipCountryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemGeoipCountry)(nil))
}

func (o SystemGeoipCountryMapOutput) ToSystemGeoipCountryMapOutput() SystemGeoipCountryMapOutput {
	return o
}

func (o SystemGeoipCountryMapOutput) ToSystemGeoipCountryMapOutputWithContext(ctx context.Context) SystemGeoipCountryMapOutput {
	return o
}

func (o SystemGeoipCountryMapOutput) MapIndex(k pulumi.StringInput) SystemGeoipCountryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemGeoipCountry {
		return vs[0].(map[string]SystemGeoipCountry)[vs[1].(string)]
	}).(SystemGeoipCountryOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemGeoipCountryOutput{})
	pulumi.RegisterOutputType(SystemGeoipCountryPtrOutput{})
	pulumi.RegisterOutputType(SystemGeoipCountryArrayOutput{})
	pulumi.RegisterOutputType(SystemGeoipCountryMapOutput{})
}
