// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Web proxy address configuration.
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
// 		_, err := fortios.NewFirewallProxyAddress(ctx, "trname", &fortios.FirewallProxyAddressArgs{
// 			CaseSensitivity: pulumi.String("disable"),
// 			Color:           pulumi.Int(2),
// 			Referrer:        pulumi.String("enable"),
// 			Type:            pulumi.String("host-regex"),
// 			Visibility:      pulumi.String("enable"),
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
// Firewall ProxyAddress can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallProxyAddress:FirewallProxyAddress labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallProxyAddress:FirewallProxyAddress labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallProxyAddress struct {
	pulumi.CustomResourceState

	// Case sensitivity in pattern. Valid values: `disable`, `enable`.
	CaseSensitivity pulumi.StringOutput `pulumi:"caseSensitivity"`
	// Tag category.
	Categories FirewallProxyAddressCategoryArrayOutput `pulumi:"categories"`
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// HTTP header regular expression.
	Header pulumi.StringOutput `pulumi:"header"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups FirewallProxyAddressHeaderGroupArrayOutput `pulumi:"headerGroups"`
	// HTTP header.
	HeaderName pulumi.StringOutput `pulumi:"headerName"`
	// Address object for the host.
	Host pulumi.StringOutput `pulumi:"host"`
	// Host name as a regular expression.
	HostRegex pulumi.StringOutput `pulumi:"hostRegex"`
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method pulumi.StringOutput `pulumi:"method"`
	// Tag name.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL path as a regular expression.
	Path pulumi.StringOutput `pulumi:"path"`
	// Match the query part of the URL as a regular expression.
	Query pulumi.StringOutput `pulumi:"query"`
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer pulumi.StringOutput `pulumi:"referrer"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyAddressTaggingArrayOutput `pulumi:"taggings"`
	// Proxy address type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua pulumi.StringOutput `pulumi:"ua"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewFirewallProxyAddress registers a new resource with the given unique name, arguments, and options.
func NewFirewallProxyAddress(ctx *pulumi.Context,
	name string, args *FirewallProxyAddressArgs, opts ...pulumi.ResourceOption) (*FirewallProxyAddress, error) {
	if args == nil {
		args = &FirewallProxyAddressArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallProxyAddress
	err := ctx.RegisterResource("fortios:index/firewallProxyAddress:FirewallProxyAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallProxyAddress gets an existing FirewallProxyAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallProxyAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallProxyAddressState, opts ...pulumi.ResourceOption) (*FirewallProxyAddress, error) {
	var resource FirewallProxyAddress
	err := ctx.ReadResource("fortios:index/firewallProxyAddress:FirewallProxyAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallProxyAddress resources.
type firewallProxyAddressState struct {
	// Case sensitivity in pattern. Valid values: `disable`, `enable`.
	CaseSensitivity *string `pulumi:"caseSensitivity"`
	// Tag category.
	Categories []FirewallProxyAddressCategory `pulumi:"categories"`
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color *int `pulumi:"color"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// HTTP header regular expression.
	Header *string `pulumi:"header"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups []FirewallProxyAddressHeaderGroup `pulumi:"headerGroups"`
	// HTTP header.
	HeaderName *string `pulumi:"headerName"`
	// Address object for the host.
	Host *string `pulumi:"host"`
	// Host name as a regular expression.
	HostRegex *string `pulumi:"hostRegex"`
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method *string `pulumi:"method"`
	// Tag name.
	Name *string `pulumi:"name"`
	// URL path as a regular expression.
	Path *string `pulumi:"path"`
	// Match the query part of the URL as a regular expression.
	Query *string `pulumi:"query"`
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer *string `pulumi:"referrer"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallProxyAddressTagging `pulumi:"taggings"`
	// Proxy address type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`.
	Type *string `pulumi:"type"`
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua *string `pulumi:"ua"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type FirewallProxyAddressState struct {
	// Case sensitivity in pattern. Valid values: `disable`, `enable`.
	CaseSensitivity pulumi.StringPtrInput
	// Tag category.
	Categories FirewallProxyAddressCategoryArrayInput
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// HTTP header regular expression.
	Header pulumi.StringPtrInput
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups FirewallProxyAddressHeaderGroupArrayInput
	// HTTP header.
	HeaderName pulumi.StringPtrInput
	// Address object for the host.
	Host pulumi.StringPtrInput
	// Host name as a regular expression.
	HostRegex pulumi.StringPtrInput
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method pulumi.StringPtrInput
	// Tag name.
	Name pulumi.StringPtrInput
	// URL path as a regular expression.
	Path pulumi.StringPtrInput
	// Match the query part of the URL as a regular expression.
	Query pulumi.StringPtrInput
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyAddressTaggingArrayInput
	// Proxy address type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`.
	Type pulumi.StringPtrInput
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallProxyAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxyAddressState)(nil)).Elem()
}

type firewallProxyAddressArgs struct {
	// Case sensitivity in pattern. Valid values: `disable`, `enable`.
	CaseSensitivity *string `pulumi:"caseSensitivity"`
	// Tag category.
	Categories []FirewallProxyAddressCategory `pulumi:"categories"`
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color *int `pulumi:"color"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// HTTP header regular expression.
	Header *string `pulumi:"header"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups []FirewallProxyAddressHeaderGroup `pulumi:"headerGroups"`
	// HTTP header.
	HeaderName *string `pulumi:"headerName"`
	// Address object for the host.
	Host *string `pulumi:"host"`
	// Host name as a regular expression.
	HostRegex *string `pulumi:"hostRegex"`
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method *string `pulumi:"method"`
	// Tag name.
	Name *string `pulumi:"name"`
	// URL path as a regular expression.
	Path *string `pulumi:"path"`
	// Match the query part of the URL as a regular expression.
	Query *string `pulumi:"query"`
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer *string `pulumi:"referrer"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallProxyAddressTagging `pulumi:"taggings"`
	// Proxy address type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`.
	Type *string `pulumi:"type"`
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua *string `pulumi:"ua"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallProxyAddress resource.
type FirewallProxyAddressArgs struct {
	// Case sensitivity in pattern. Valid values: `disable`, `enable`.
	CaseSensitivity pulumi.StringPtrInput
	// Tag category.
	Categories FirewallProxyAddressCategoryArrayInput
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// HTTP header regular expression.
	Header pulumi.StringPtrInput
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups FirewallProxyAddressHeaderGroupArrayInput
	// HTTP header.
	HeaderName pulumi.StringPtrInput
	// Address object for the host.
	Host pulumi.StringPtrInput
	// Host name as a regular expression.
	HostRegex pulumi.StringPtrInput
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method pulumi.StringPtrInput
	// Tag name.
	Name pulumi.StringPtrInput
	// URL path as a regular expression.
	Path pulumi.StringPtrInput
	// Match the query part of the URL as a regular expression.
	Query pulumi.StringPtrInput
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyAddressTaggingArrayInput
	// Proxy address type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`.
	Type pulumi.StringPtrInput
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallProxyAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxyAddressArgs)(nil)).Elem()
}

type FirewallProxyAddressInput interface {
	pulumi.Input

	ToFirewallProxyAddressOutput() FirewallProxyAddressOutput
	ToFirewallProxyAddressOutputWithContext(ctx context.Context) FirewallProxyAddressOutput
}

func (*FirewallProxyAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxyAddress)(nil)).Elem()
}

func (i *FirewallProxyAddress) ToFirewallProxyAddressOutput() FirewallProxyAddressOutput {
	return i.ToFirewallProxyAddressOutputWithContext(context.Background())
}

func (i *FirewallProxyAddress) ToFirewallProxyAddressOutputWithContext(ctx context.Context) FirewallProxyAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyAddressOutput)
}

// FirewallProxyAddressArrayInput is an input type that accepts FirewallProxyAddressArray and FirewallProxyAddressArrayOutput values.
// You can construct a concrete instance of `FirewallProxyAddressArrayInput` via:
//
//          FirewallProxyAddressArray{ FirewallProxyAddressArgs{...} }
type FirewallProxyAddressArrayInput interface {
	pulumi.Input

	ToFirewallProxyAddressArrayOutput() FirewallProxyAddressArrayOutput
	ToFirewallProxyAddressArrayOutputWithContext(context.Context) FirewallProxyAddressArrayOutput
}

type FirewallProxyAddressArray []FirewallProxyAddressInput

func (FirewallProxyAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxyAddress)(nil)).Elem()
}

func (i FirewallProxyAddressArray) ToFirewallProxyAddressArrayOutput() FirewallProxyAddressArrayOutput {
	return i.ToFirewallProxyAddressArrayOutputWithContext(context.Background())
}

func (i FirewallProxyAddressArray) ToFirewallProxyAddressArrayOutputWithContext(ctx context.Context) FirewallProxyAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyAddressArrayOutput)
}

// FirewallProxyAddressMapInput is an input type that accepts FirewallProxyAddressMap and FirewallProxyAddressMapOutput values.
// You can construct a concrete instance of `FirewallProxyAddressMapInput` via:
//
//          FirewallProxyAddressMap{ "key": FirewallProxyAddressArgs{...} }
type FirewallProxyAddressMapInput interface {
	pulumi.Input

	ToFirewallProxyAddressMapOutput() FirewallProxyAddressMapOutput
	ToFirewallProxyAddressMapOutputWithContext(context.Context) FirewallProxyAddressMapOutput
}

type FirewallProxyAddressMap map[string]FirewallProxyAddressInput

func (FirewallProxyAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxyAddress)(nil)).Elem()
}

func (i FirewallProxyAddressMap) ToFirewallProxyAddressMapOutput() FirewallProxyAddressMapOutput {
	return i.ToFirewallProxyAddressMapOutputWithContext(context.Background())
}

func (i FirewallProxyAddressMap) ToFirewallProxyAddressMapOutputWithContext(ctx context.Context) FirewallProxyAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyAddressMapOutput)
}

type FirewallProxyAddressOutput struct{ *pulumi.OutputState }

func (FirewallProxyAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxyAddress)(nil)).Elem()
}

func (o FirewallProxyAddressOutput) ToFirewallProxyAddressOutput() FirewallProxyAddressOutput {
	return o
}

func (o FirewallProxyAddressOutput) ToFirewallProxyAddressOutputWithContext(ctx context.Context) FirewallProxyAddressOutput {
	return o
}

type FirewallProxyAddressArrayOutput struct{ *pulumi.OutputState }

func (FirewallProxyAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxyAddress)(nil)).Elem()
}

func (o FirewallProxyAddressArrayOutput) ToFirewallProxyAddressArrayOutput() FirewallProxyAddressArrayOutput {
	return o
}

func (o FirewallProxyAddressArrayOutput) ToFirewallProxyAddressArrayOutputWithContext(ctx context.Context) FirewallProxyAddressArrayOutput {
	return o
}

func (o FirewallProxyAddressArrayOutput) Index(i pulumi.IntInput) FirewallProxyAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallProxyAddress {
		return vs[0].([]*FirewallProxyAddress)[vs[1].(int)]
	}).(FirewallProxyAddressOutput)
}

type FirewallProxyAddressMapOutput struct{ *pulumi.OutputState }

func (FirewallProxyAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxyAddress)(nil)).Elem()
}

func (o FirewallProxyAddressMapOutput) ToFirewallProxyAddressMapOutput() FirewallProxyAddressMapOutput {
	return o
}

func (o FirewallProxyAddressMapOutput) ToFirewallProxyAddressMapOutputWithContext(ctx context.Context) FirewallProxyAddressMapOutput {
	return o
}

func (o FirewallProxyAddressMapOutput) MapIndex(k pulumi.StringInput) FirewallProxyAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallProxyAddress {
		return vs[0].(map[string]*FirewallProxyAddress)[vs[1].(string)]
	}).(FirewallProxyAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyAddressInput)(nil)).Elem(), &FirewallProxyAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyAddressArrayInput)(nil)).Elem(), FirewallProxyAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyAddressMapInput)(nil)).Elem(), FirewallProxyAddressMap{})
	pulumi.RegisterOutputType(FirewallProxyAddressOutput{})
	pulumi.RegisterOutputType(FirewallProxyAddressArrayOutput{})
	pulumi.RegisterOutputType(FirewallProxyAddressMapOutput{})
}
