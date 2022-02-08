// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure devices. Applies to FortiOS Version `<= 6.2.0`.
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
// 		_, err := fortios.NewUserDevice(ctx, "trname", &fortios.UserDeviceArgs{
// 			Alias:    pulumi.String("1"),
// 			Category: pulumi.String("amazon-device"),
// 			Mac:      pulumi.String("08:00:20:0a:8c:6d"),
// 			Type:     pulumi.String("unknown"),
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
// User Device can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/userDevice:UserDevice labelname {{alias}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type UserDevice struct {
	pulumi.CustomResourceState

	// Device alias.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Image file for avatar (maximum 4K base64 encoded).
	Avatar pulumi.StringPtrOutput `pulumi:"avatar"`
	// Tag category.
	Category pulumi.StringOutput `pulumi:"category"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Device MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Master device (optional).
	MasterDevice pulumi.StringOutput `pulumi:"masterDevice"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings UserDeviceTaggingArrayOutput `pulumi:"taggings"`
	// Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
	Type pulumi.StringOutput `pulumi:"type"`
	// User name.
	User pulumi.StringOutput `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserDevice registers a new resource with the given unique name, arguments, and options.
func NewUserDevice(ctx *pulumi.Context,
	name string, args *UserDeviceArgs, opts ...pulumi.ResourceOption) (*UserDevice, error) {
	if args == nil {
		args = &UserDeviceArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource UserDevice
	err := ctx.RegisterResource("fortios:index/userDevice:UserDevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserDevice gets an existing UserDevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserDeviceState, opts ...pulumi.ResourceOption) (*UserDevice, error) {
	var resource UserDevice
	err := ctx.ReadResource("fortios:index/userDevice:UserDevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserDevice resources.
type userDeviceState struct {
	// Device alias.
	Alias *string `pulumi:"alias"`
	// Image file for avatar (maximum 4K base64 encoded).
	Avatar *string `pulumi:"avatar"`
	// Tag category.
	Category *string `pulumi:"category"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Device MAC address.
	Mac *string `pulumi:"mac"`
	// Master device (optional).
	MasterDevice *string `pulumi:"masterDevice"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []UserDeviceTagging `pulumi:"taggings"`
	// Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
	Type *string `pulumi:"type"`
	// User name.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserDeviceState struct {
	// Device alias.
	Alias pulumi.StringPtrInput
	// Image file for avatar (maximum 4K base64 encoded).
	Avatar pulumi.StringPtrInput
	// Tag category.
	Category pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Device MAC address.
	Mac pulumi.StringPtrInput
	// Master device (optional).
	MasterDevice pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings UserDeviceTaggingArrayInput
	// Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
	Type pulumi.StringPtrInput
	// User name.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserDeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*userDeviceState)(nil)).Elem()
}

type userDeviceArgs struct {
	// Device alias.
	Alias *string `pulumi:"alias"`
	// Image file for avatar (maximum 4K base64 encoded).
	Avatar *string `pulumi:"avatar"`
	// Tag category.
	Category *string `pulumi:"category"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Device MAC address.
	Mac *string `pulumi:"mac"`
	// Master device (optional).
	MasterDevice *string `pulumi:"masterDevice"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []UserDeviceTagging `pulumi:"taggings"`
	// Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
	Type *string `pulumi:"type"`
	// User name.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserDevice resource.
type UserDeviceArgs struct {
	// Device alias.
	Alias pulumi.StringPtrInput
	// Image file for avatar (maximum 4K base64 encoded).
	Avatar pulumi.StringPtrInput
	// Tag category.
	Category pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Device MAC address.
	Mac pulumi.StringPtrInput
	// Master device (optional).
	MasterDevice pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings UserDeviceTaggingArrayInput
	// Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
	Type pulumi.StringPtrInput
	// User name.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userDeviceArgs)(nil)).Elem()
}

type UserDeviceInput interface {
	pulumi.Input

	ToUserDeviceOutput() UserDeviceOutput
	ToUserDeviceOutputWithContext(ctx context.Context) UserDeviceOutput
}

func (*UserDevice) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDevice)(nil)).Elem()
}

func (i *UserDevice) ToUserDeviceOutput() UserDeviceOutput {
	return i.ToUserDeviceOutputWithContext(context.Background())
}

func (i *UserDevice) ToUserDeviceOutputWithContext(ctx context.Context) UserDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDeviceOutput)
}

// UserDeviceArrayInput is an input type that accepts UserDeviceArray and UserDeviceArrayOutput values.
// You can construct a concrete instance of `UserDeviceArrayInput` via:
//
//          UserDeviceArray{ UserDeviceArgs{...} }
type UserDeviceArrayInput interface {
	pulumi.Input

	ToUserDeviceArrayOutput() UserDeviceArrayOutput
	ToUserDeviceArrayOutputWithContext(context.Context) UserDeviceArrayOutput
}

type UserDeviceArray []UserDeviceInput

func (UserDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDevice)(nil)).Elem()
}

func (i UserDeviceArray) ToUserDeviceArrayOutput() UserDeviceArrayOutput {
	return i.ToUserDeviceArrayOutputWithContext(context.Background())
}

func (i UserDeviceArray) ToUserDeviceArrayOutputWithContext(ctx context.Context) UserDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDeviceArrayOutput)
}

// UserDeviceMapInput is an input type that accepts UserDeviceMap and UserDeviceMapOutput values.
// You can construct a concrete instance of `UserDeviceMapInput` via:
//
//          UserDeviceMap{ "key": UserDeviceArgs{...} }
type UserDeviceMapInput interface {
	pulumi.Input

	ToUserDeviceMapOutput() UserDeviceMapOutput
	ToUserDeviceMapOutputWithContext(context.Context) UserDeviceMapOutput
}

type UserDeviceMap map[string]UserDeviceInput

func (UserDeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDevice)(nil)).Elem()
}

func (i UserDeviceMap) ToUserDeviceMapOutput() UserDeviceMapOutput {
	return i.ToUserDeviceMapOutputWithContext(context.Background())
}

func (i UserDeviceMap) ToUserDeviceMapOutputWithContext(ctx context.Context) UserDeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDeviceMapOutput)
}

type UserDeviceOutput struct{ *pulumi.OutputState }

func (UserDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDevice)(nil)).Elem()
}

func (o UserDeviceOutput) ToUserDeviceOutput() UserDeviceOutput {
	return o
}

func (o UserDeviceOutput) ToUserDeviceOutputWithContext(ctx context.Context) UserDeviceOutput {
	return o
}

type UserDeviceArrayOutput struct{ *pulumi.OutputState }

func (UserDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDevice)(nil)).Elem()
}

func (o UserDeviceArrayOutput) ToUserDeviceArrayOutput() UserDeviceArrayOutput {
	return o
}

func (o UserDeviceArrayOutput) ToUserDeviceArrayOutputWithContext(ctx context.Context) UserDeviceArrayOutput {
	return o
}

func (o UserDeviceArrayOutput) Index(i pulumi.IntInput) UserDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserDevice {
		return vs[0].([]*UserDevice)[vs[1].(int)]
	}).(UserDeviceOutput)
}

type UserDeviceMapOutput struct{ *pulumi.OutputState }

func (UserDeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDevice)(nil)).Elem()
}

func (o UserDeviceMapOutput) ToUserDeviceMapOutput() UserDeviceMapOutput {
	return o
}

func (o UserDeviceMapOutput) ToUserDeviceMapOutputWithContext(ctx context.Context) UserDeviceMapOutput {
	return o
}

func (o UserDeviceMapOutput) MapIndex(k pulumi.StringInput) UserDeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserDevice {
		return vs[0].(map[string]*UserDevice)[vs[1].(string)]
	}).(UserDeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserDeviceInput)(nil)).Elem(), &UserDevice{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDeviceArrayInput)(nil)).Elem(), UserDeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDeviceMapInput)(nil)).Elem(), UserDeviceMap{})
	pulumi.RegisterOutputType(UserDeviceOutput{})
	pulumi.RegisterOutputType(UserDeviceArrayOutput{})
	pulumi.RegisterOutputType(UserDeviceMapOutput{})
}
