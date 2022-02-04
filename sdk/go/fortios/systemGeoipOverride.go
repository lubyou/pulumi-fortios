// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
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
// 		_, err := fortios.NewSystemGeoipOverride(ctx, "trname", &fortios.SystemGeoipOverrideArgs{
// 			Description: pulumi.String("TEST for country"),
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
// System GeoipOverride can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemGeoipOverride:SystemGeoipOverride labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemGeoipOverride struct {
	pulumi.CustomResourceState

	// Two character Country ID code.
	CountryId pulumi.StringOutput `pulumi:"countryId"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges SystemGeoipOverrideIp6RangeArrayOutput `pulumi:"ip6Ranges"`
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges SystemGeoipOverrideIpRangeArrayOutput `pulumi:"ipRanges"`
	// Location name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemGeoipOverride registers a new resource with the given unique name, arguments, and options.
func NewSystemGeoipOverride(ctx *pulumi.Context,
	name string, args *SystemGeoipOverrideArgs, opts ...pulumi.ResourceOption) (*SystemGeoipOverride, error) {
	if args == nil {
		args = &SystemGeoipOverrideArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemGeoipOverride
	err := ctx.RegisterResource("fortios:index/systemGeoipOverride:SystemGeoipOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemGeoipOverride gets an existing SystemGeoipOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemGeoipOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemGeoipOverrideState, opts ...pulumi.ResourceOption) (*SystemGeoipOverride, error) {
	var resource SystemGeoipOverride
	err := ctx.ReadResource("fortios:index/systemGeoipOverride:SystemGeoipOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemGeoipOverride resources.
type systemGeoipOverrideState struct {
	// Two character Country ID code.
	CountryId *string `pulumi:"countryId"`
	// Description.
	Description *string `pulumi:"description"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges []SystemGeoipOverrideIp6Range `pulumi:"ip6Ranges"`
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges []SystemGeoipOverrideIpRange `pulumi:"ipRanges"`
	// Location name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemGeoipOverrideState struct {
	// Two character Country ID code.
	CountryId pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges SystemGeoipOverrideIp6RangeArrayInput
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges SystemGeoipOverrideIpRangeArrayInput
	// Location name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemGeoipOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGeoipOverrideState)(nil)).Elem()
}

type systemGeoipOverrideArgs struct {
	// Two character Country ID code.
	CountryId *string `pulumi:"countryId"`
	// Description.
	Description *string `pulumi:"description"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges []SystemGeoipOverrideIp6Range `pulumi:"ip6Ranges"`
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges []SystemGeoipOverrideIpRange `pulumi:"ipRanges"`
	// Location name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemGeoipOverride resource.
type SystemGeoipOverrideArgs struct {
	// Two character Country ID code.
	CountryId pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges SystemGeoipOverrideIp6RangeArrayInput
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges SystemGeoipOverrideIpRangeArrayInput
	// Location name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemGeoipOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGeoipOverrideArgs)(nil)).Elem()
}

type SystemGeoipOverrideInput interface {
	pulumi.Input

	ToSystemGeoipOverrideOutput() SystemGeoipOverrideOutput
	ToSystemGeoipOverrideOutputWithContext(ctx context.Context) SystemGeoipOverrideOutput
}

func (*SystemGeoipOverride) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGeoipOverride)(nil)).Elem()
}

func (i *SystemGeoipOverride) ToSystemGeoipOverrideOutput() SystemGeoipOverrideOutput {
	return i.ToSystemGeoipOverrideOutputWithContext(context.Background())
}

func (i *SystemGeoipOverride) ToSystemGeoipOverrideOutputWithContext(ctx context.Context) SystemGeoipOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipOverrideOutput)
}

// SystemGeoipOverrideArrayInput is an input type that accepts SystemGeoipOverrideArray and SystemGeoipOverrideArrayOutput values.
// You can construct a concrete instance of `SystemGeoipOverrideArrayInput` via:
//
//          SystemGeoipOverrideArray{ SystemGeoipOverrideArgs{...} }
type SystemGeoipOverrideArrayInput interface {
	pulumi.Input

	ToSystemGeoipOverrideArrayOutput() SystemGeoipOverrideArrayOutput
	ToSystemGeoipOverrideArrayOutputWithContext(context.Context) SystemGeoipOverrideArrayOutput
}

type SystemGeoipOverrideArray []SystemGeoipOverrideInput

func (SystemGeoipOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGeoipOverride)(nil)).Elem()
}

func (i SystemGeoipOverrideArray) ToSystemGeoipOverrideArrayOutput() SystemGeoipOverrideArrayOutput {
	return i.ToSystemGeoipOverrideArrayOutputWithContext(context.Background())
}

func (i SystemGeoipOverrideArray) ToSystemGeoipOverrideArrayOutputWithContext(ctx context.Context) SystemGeoipOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipOverrideArrayOutput)
}

// SystemGeoipOverrideMapInput is an input type that accepts SystemGeoipOverrideMap and SystemGeoipOverrideMapOutput values.
// You can construct a concrete instance of `SystemGeoipOverrideMapInput` via:
//
//          SystemGeoipOverrideMap{ "key": SystemGeoipOverrideArgs{...} }
type SystemGeoipOverrideMapInput interface {
	pulumi.Input

	ToSystemGeoipOverrideMapOutput() SystemGeoipOverrideMapOutput
	ToSystemGeoipOverrideMapOutputWithContext(context.Context) SystemGeoipOverrideMapOutput
}

type SystemGeoipOverrideMap map[string]SystemGeoipOverrideInput

func (SystemGeoipOverrideMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGeoipOverride)(nil)).Elem()
}

func (i SystemGeoipOverrideMap) ToSystemGeoipOverrideMapOutput() SystemGeoipOverrideMapOutput {
	return i.ToSystemGeoipOverrideMapOutputWithContext(context.Background())
}

func (i SystemGeoipOverrideMap) ToSystemGeoipOverrideMapOutputWithContext(ctx context.Context) SystemGeoipOverrideMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipOverrideMapOutput)
}

type SystemGeoipOverrideOutput struct{ *pulumi.OutputState }

func (SystemGeoipOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGeoipOverride)(nil)).Elem()
}

func (o SystemGeoipOverrideOutput) ToSystemGeoipOverrideOutput() SystemGeoipOverrideOutput {
	return o
}

func (o SystemGeoipOverrideOutput) ToSystemGeoipOverrideOutputWithContext(ctx context.Context) SystemGeoipOverrideOutput {
	return o
}

type SystemGeoipOverrideArrayOutput struct{ *pulumi.OutputState }

func (SystemGeoipOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGeoipOverride)(nil)).Elem()
}

func (o SystemGeoipOverrideArrayOutput) ToSystemGeoipOverrideArrayOutput() SystemGeoipOverrideArrayOutput {
	return o
}

func (o SystemGeoipOverrideArrayOutput) ToSystemGeoipOverrideArrayOutputWithContext(ctx context.Context) SystemGeoipOverrideArrayOutput {
	return o
}

func (o SystemGeoipOverrideArrayOutput) Index(i pulumi.IntInput) SystemGeoipOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemGeoipOverride {
		return vs[0].([]*SystemGeoipOverride)[vs[1].(int)]
	}).(SystemGeoipOverrideOutput)
}

type SystemGeoipOverrideMapOutput struct{ *pulumi.OutputState }

func (SystemGeoipOverrideMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGeoipOverride)(nil)).Elem()
}

func (o SystemGeoipOverrideMapOutput) ToSystemGeoipOverrideMapOutput() SystemGeoipOverrideMapOutput {
	return o
}

func (o SystemGeoipOverrideMapOutput) ToSystemGeoipOverrideMapOutputWithContext(ctx context.Context) SystemGeoipOverrideMapOutput {
	return o
}

func (o SystemGeoipOverrideMapOutput) MapIndex(k pulumi.StringInput) SystemGeoipOverrideOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemGeoipOverride {
		return vs[0].(map[string]*SystemGeoipOverride)[vs[1].(string)]
	}).(SystemGeoipOverrideOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeoipOverrideInput)(nil)).Elem(), &SystemGeoipOverride{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeoipOverrideArrayInput)(nil)).Elem(), SystemGeoipOverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeoipOverrideMapInput)(nil)).Elem(), SystemGeoipOverrideMap{})
	pulumi.RegisterOutputType(SystemGeoipOverrideOutput{})
	pulumi.RegisterOutputType(SystemGeoipOverrideArrayOutput{})
	pulumi.RegisterOutputType(SystemGeoipOverrideMapOutput{})
}
