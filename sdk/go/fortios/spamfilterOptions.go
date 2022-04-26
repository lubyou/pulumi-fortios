// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam options. Applies to FortiOS Version `<= 6.2.0`.
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
// 		_, err := fortios.NewSpamfilterOptions(ctx, "trname", &fortios.SpamfilterOptionsArgs{
// 			DnsTimeout: pulumi.Int(7),
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
// Spamfilter Options can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/spamfilterOptions:SpamfilterOptions labelname SpamfilterOptions
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/spamfilterOptions:SpamfilterOptions labelname SpamfilterOptions
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SpamfilterOptions struct {
	pulumi.CustomResourceState

	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntOutput `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSpamfilterOptions registers a new resource with the given unique name, arguments, and options.
func NewSpamfilterOptions(ctx *pulumi.Context,
	name string, args *SpamfilterOptionsArgs, opts ...pulumi.ResourceOption) (*SpamfilterOptions, error) {
	if args == nil {
		args = &SpamfilterOptionsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SpamfilterOptions
	err := ctx.RegisterResource("fortios:index/spamfilterOptions:SpamfilterOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpamfilterOptions gets an existing SpamfilterOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpamfilterOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpamfilterOptionsState, opts ...pulumi.ResourceOption) (*SpamfilterOptions, error) {
	var resource SpamfilterOptions
	err := ctx.ReadResource("fortios:index/spamfilterOptions:SpamfilterOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpamfilterOptions resources.
type spamfilterOptionsState struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout *int `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SpamfilterOptionsState struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterOptionsState)(nil)).Elem()
}

type spamfilterOptionsArgs struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout *int `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SpamfilterOptions resource.
type SpamfilterOptionsArgs struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterOptionsArgs)(nil)).Elem()
}

type SpamfilterOptionsInput interface {
	pulumi.Input

	ToSpamfilterOptionsOutput() SpamfilterOptionsOutput
	ToSpamfilterOptionsOutputWithContext(ctx context.Context) SpamfilterOptionsOutput
}

func (*SpamfilterOptions) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterOptions)(nil)).Elem()
}

func (i *SpamfilterOptions) ToSpamfilterOptionsOutput() SpamfilterOptionsOutput {
	return i.ToSpamfilterOptionsOutputWithContext(context.Background())
}

func (i *SpamfilterOptions) ToSpamfilterOptionsOutputWithContext(ctx context.Context) SpamfilterOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterOptionsOutput)
}

// SpamfilterOptionsArrayInput is an input type that accepts SpamfilterOptionsArray and SpamfilterOptionsArrayOutput values.
// You can construct a concrete instance of `SpamfilterOptionsArrayInput` via:
//
//          SpamfilterOptionsArray{ SpamfilterOptionsArgs{...} }
type SpamfilterOptionsArrayInput interface {
	pulumi.Input

	ToSpamfilterOptionsArrayOutput() SpamfilterOptionsArrayOutput
	ToSpamfilterOptionsArrayOutputWithContext(context.Context) SpamfilterOptionsArrayOutput
}

type SpamfilterOptionsArray []SpamfilterOptionsInput

func (SpamfilterOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterOptions)(nil)).Elem()
}

func (i SpamfilterOptionsArray) ToSpamfilterOptionsArrayOutput() SpamfilterOptionsArrayOutput {
	return i.ToSpamfilterOptionsArrayOutputWithContext(context.Background())
}

func (i SpamfilterOptionsArray) ToSpamfilterOptionsArrayOutputWithContext(ctx context.Context) SpamfilterOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterOptionsArrayOutput)
}

// SpamfilterOptionsMapInput is an input type that accepts SpamfilterOptionsMap and SpamfilterOptionsMapOutput values.
// You can construct a concrete instance of `SpamfilterOptionsMapInput` via:
//
//          SpamfilterOptionsMap{ "key": SpamfilterOptionsArgs{...} }
type SpamfilterOptionsMapInput interface {
	pulumi.Input

	ToSpamfilterOptionsMapOutput() SpamfilterOptionsMapOutput
	ToSpamfilterOptionsMapOutputWithContext(context.Context) SpamfilterOptionsMapOutput
}

type SpamfilterOptionsMap map[string]SpamfilterOptionsInput

func (SpamfilterOptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterOptions)(nil)).Elem()
}

func (i SpamfilterOptionsMap) ToSpamfilterOptionsMapOutput() SpamfilterOptionsMapOutput {
	return i.ToSpamfilterOptionsMapOutputWithContext(context.Background())
}

func (i SpamfilterOptionsMap) ToSpamfilterOptionsMapOutputWithContext(ctx context.Context) SpamfilterOptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterOptionsMapOutput)
}

type SpamfilterOptionsOutput struct{ *pulumi.OutputState }

func (SpamfilterOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterOptions)(nil)).Elem()
}

func (o SpamfilterOptionsOutput) ToSpamfilterOptionsOutput() SpamfilterOptionsOutput {
	return o
}

func (o SpamfilterOptionsOutput) ToSpamfilterOptionsOutputWithContext(ctx context.Context) SpamfilterOptionsOutput {
	return o
}

type SpamfilterOptionsArrayOutput struct{ *pulumi.OutputState }

func (SpamfilterOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterOptions)(nil)).Elem()
}

func (o SpamfilterOptionsArrayOutput) ToSpamfilterOptionsArrayOutput() SpamfilterOptionsArrayOutput {
	return o
}

func (o SpamfilterOptionsArrayOutput) ToSpamfilterOptionsArrayOutputWithContext(ctx context.Context) SpamfilterOptionsArrayOutput {
	return o
}

func (o SpamfilterOptionsArrayOutput) Index(i pulumi.IntInput) SpamfilterOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SpamfilterOptions {
		return vs[0].([]*SpamfilterOptions)[vs[1].(int)]
	}).(SpamfilterOptionsOutput)
}

type SpamfilterOptionsMapOutput struct{ *pulumi.OutputState }

func (SpamfilterOptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterOptions)(nil)).Elem()
}

func (o SpamfilterOptionsMapOutput) ToSpamfilterOptionsMapOutput() SpamfilterOptionsMapOutput {
	return o
}

func (o SpamfilterOptionsMapOutput) ToSpamfilterOptionsMapOutputWithContext(ctx context.Context) SpamfilterOptionsMapOutput {
	return o
}

func (o SpamfilterOptionsMapOutput) MapIndex(k pulumi.StringInput) SpamfilterOptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SpamfilterOptions {
		return vs[0].(map[string]*SpamfilterOptions)[vs[1].(string)]
	}).(SpamfilterOptionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterOptionsInput)(nil)).Elem(), &SpamfilterOptions{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterOptionsArrayInput)(nil)).Elem(), SpamfilterOptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterOptionsMapInput)(nil)).Elem(), SpamfilterOptionsMap{})
	pulumi.RegisterOutputType(SpamfilterOptionsOutput{})
	pulumi.RegisterOutputType(SpamfilterOptionsArrayOutput{})
	pulumi.RegisterOutputType(SpamfilterOptionsMapOutput{})
}
