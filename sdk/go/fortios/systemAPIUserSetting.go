// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `SystemApiUser`, we recommend that you use the new resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSystemAPIUserSetting(ctx, "test2", &fortios.SystemAPIUserSettingArgs{
// 			Accprofile: pulumi.String("restAPIprofile"),
// 			Trusthosts: SystemAPIUserSettingTrusthostArray{
// 				&SystemAPIUserSettingTrusthostArgs{
// 					Ipv4Trusthost: pulumi.String("61.149.0.0 255.255.0.0"),
// 					Type:          pulumi.String("ipv4-trusthost"),
// 				},
// 				&SystemAPIUserSettingTrusthostArgs{
// 					Ipv4Trusthost: pulumi.String("22.22.0.0 255.255.0.0"),
// 					Type:          pulumi.String("ipv4-trusthost"),
// 				},
// 			},
// 			Vdoms: pulumi.StringArray{
// 				pulumi.String("root"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SystemAPIUserSetting struct {
	pulumi.CustomResourceState

	// Admin user access profile.
	Accprofile pulumi.StringOutput `pulumi:"accprofile"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// User name.
	Name       pulumi.StringOutput                      `pulumi:"name"`
	Trusthosts SystemAPIUserSettingTrusthostArrayOutput `pulumi:"trusthosts"`
	// Virtual domains.
	// * `trusthost-Type` - (Required) Trusthost type.
	// * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
	Vdoms pulumi.StringArrayOutput `pulumi:"vdoms"`
}

// NewSystemAPIUserSetting registers a new resource with the given unique name, arguments, and options.
func NewSystemAPIUserSetting(ctx *pulumi.Context,
	name string, args *SystemAPIUserSettingArgs, opts ...pulumi.ResourceOption) (*SystemAPIUserSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Accprofile == nil {
		return nil, errors.New("invalid value for required argument 'Accprofile'")
	}
	if args.Trusthosts == nil {
		return nil, errors.New("invalid value for required argument 'Trusthosts'")
	}
	if args.Vdoms == nil {
		return nil, errors.New("invalid value for required argument 'Vdoms'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAPIUserSetting
	err := ctx.RegisterResource("fortios:index/systemAPIUserSetting:SystemAPIUserSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAPIUserSetting gets an existing SystemAPIUserSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAPIUserSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAPIUserSettingState, opts ...pulumi.ResourceOption) (*SystemAPIUserSetting, error) {
	var resource SystemAPIUserSetting
	err := ctx.ReadResource("fortios:index/systemAPIUserSetting:SystemAPIUserSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAPIUserSetting resources.
type systemAPIUserSettingState struct {
	// Admin user access profile.
	Accprofile *string `pulumi:"accprofile"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// User name.
	Name       *string                         `pulumi:"name"`
	Trusthosts []SystemAPIUserSettingTrusthost `pulumi:"trusthosts"`
	// Virtual domains.
	// * `trusthost-Type` - (Required) Trusthost type.
	// * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
	Vdoms []string `pulumi:"vdoms"`
}

type SystemAPIUserSettingState struct {
	// Admin user access profile.
	Accprofile pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// User name.
	Name       pulumi.StringPtrInput
	Trusthosts SystemAPIUserSettingTrusthostArrayInput
	// Virtual domains.
	// * `trusthost-Type` - (Required) Trusthost type.
	// * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
	Vdoms pulumi.StringArrayInput
}

func (SystemAPIUserSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAPIUserSettingState)(nil)).Elem()
}

type systemAPIUserSettingArgs struct {
	// Admin user access profile.
	Accprofile string `pulumi:"accprofile"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// User name.
	Name       *string                         `pulumi:"name"`
	Trusthosts []SystemAPIUserSettingTrusthost `pulumi:"trusthosts"`
	// Virtual domains.
	// * `trusthost-Type` - (Required) Trusthost type.
	// * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
	Vdoms []string `pulumi:"vdoms"`
}

// The set of arguments for constructing a SystemAPIUserSetting resource.
type SystemAPIUserSettingArgs struct {
	// Admin user access profile.
	Accprofile pulumi.StringInput
	// Comment.
	Comments pulumi.StringPtrInput
	// User name.
	Name       pulumi.StringPtrInput
	Trusthosts SystemAPIUserSettingTrusthostArrayInput
	// Virtual domains.
	// * `trusthost-Type` - (Required) Trusthost type.
	// * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
	Vdoms pulumi.StringArrayInput
}

func (SystemAPIUserSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAPIUserSettingArgs)(nil)).Elem()
}

type SystemAPIUserSettingInput interface {
	pulumi.Input

	ToSystemAPIUserSettingOutput() SystemAPIUserSettingOutput
	ToSystemAPIUserSettingOutputWithContext(ctx context.Context) SystemAPIUserSettingOutput
}

func (*SystemAPIUserSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAPIUserSetting)(nil)).Elem()
}

func (i *SystemAPIUserSetting) ToSystemAPIUserSettingOutput() SystemAPIUserSettingOutput {
	return i.ToSystemAPIUserSettingOutputWithContext(context.Background())
}

func (i *SystemAPIUserSetting) ToSystemAPIUserSettingOutputWithContext(ctx context.Context) SystemAPIUserSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAPIUserSettingOutput)
}

// SystemAPIUserSettingArrayInput is an input type that accepts SystemAPIUserSettingArray and SystemAPIUserSettingArrayOutput values.
// You can construct a concrete instance of `SystemAPIUserSettingArrayInput` via:
//
//          SystemAPIUserSettingArray{ SystemAPIUserSettingArgs{...} }
type SystemAPIUserSettingArrayInput interface {
	pulumi.Input

	ToSystemAPIUserSettingArrayOutput() SystemAPIUserSettingArrayOutput
	ToSystemAPIUserSettingArrayOutputWithContext(context.Context) SystemAPIUserSettingArrayOutput
}

type SystemAPIUserSettingArray []SystemAPIUserSettingInput

func (SystemAPIUserSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAPIUserSetting)(nil)).Elem()
}

func (i SystemAPIUserSettingArray) ToSystemAPIUserSettingArrayOutput() SystemAPIUserSettingArrayOutput {
	return i.ToSystemAPIUserSettingArrayOutputWithContext(context.Background())
}

func (i SystemAPIUserSettingArray) ToSystemAPIUserSettingArrayOutputWithContext(ctx context.Context) SystemAPIUserSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAPIUserSettingArrayOutput)
}

// SystemAPIUserSettingMapInput is an input type that accepts SystemAPIUserSettingMap and SystemAPIUserSettingMapOutput values.
// You can construct a concrete instance of `SystemAPIUserSettingMapInput` via:
//
//          SystemAPIUserSettingMap{ "key": SystemAPIUserSettingArgs{...} }
type SystemAPIUserSettingMapInput interface {
	pulumi.Input

	ToSystemAPIUserSettingMapOutput() SystemAPIUserSettingMapOutput
	ToSystemAPIUserSettingMapOutputWithContext(context.Context) SystemAPIUserSettingMapOutput
}

type SystemAPIUserSettingMap map[string]SystemAPIUserSettingInput

func (SystemAPIUserSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAPIUserSetting)(nil)).Elem()
}

func (i SystemAPIUserSettingMap) ToSystemAPIUserSettingMapOutput() SystemAPIUserSettingMapOutput {
	return i.ToSystemAPIUserSettingMapOutputWithContext(context.Background())
}

func (i SystemAPIUserSettingMap) ToSystemAPIUserSettingMapOutputWithContext(ctx context.Context) SystemAPIUserSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAPIUserSettingMapOutput)
}

type SystemAPIUserSettingOutput struct{ *pulumi.OutputState }

func (SystemAPIUserSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAPIUserSetting)(nil)).Elem()
}

func (o SystemAPIUserSettingOutput) ToSystemAPIUserSettingOutput() SystemAPIUserSettingOutput {
	return o
}

func (o SystemAPIUserSettingOutput) ToSystemAPIUserSettingOutputWithContext(ctx context.Context) SystemAPIUserSettingOutput {
	return o
}

type SystemAPIUserSettingArrayOutput struct{ *pulumi.OutputState }

func (SystemAPIUserSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAPIUserSetting)(nil)).Elem()
}

func (o SystemAPIUserSettingArrayOutput) ToSystemAPIUserSettingArrayOutput() SystemAPIUserSettingArrayOutput {
	return o
}

func (o SystemAPIUserSettingArrayOutput) ToSystemAPIUserSettingArrayOutputWithContext(ctx context.Context) SystemAPIUserSettingArrayOutput {
	return o
}

func (o SystemAPIUserSettingArrayOutput) Index(i pulumi.IntInput) SystemAPIUserSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAPIUserSetting {
		return vs[0].([]*SystemAPIUserSetting)[vs[1].(int)]
	}).(SystemAPIUserSettingOutput)
}

type SystemAPIUserSettingMapOutput struct{ *pulumi.OutputState }

func (SystemAPIUserSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAPIUserSetting)(nil)).Elem()
}

func (o SystemAPIUserSettingMapOutput) ToSystemAPIUserSettingMapOutput() SystemAPIUserSettingMapOutput {
	return o
}

func (o SystemAPIUserSettingMapOutput) ToSystemAPIUserSettingMapOutputWithContext(ctx context.Context) SystemAPIUserSettingMapOutput {
	return o
}

func (o SystemAPIUserSettingMapOutput) MapIndex(k pulumi.StringInput) SystemAPIUserSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAPIUserSetting {
		return vs[0].(map[string]*SystemAPIUserSetting)[vs[1].(string)]
	}).(SystemAPIUserSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAPIUserSettingInput)(nil)).Elem(), &SystemAPIUserSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAPIUserSettingArrayInput)(nil)).Elem(), SystemAPIUserSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAPIUserSettingMapInput)(nil)).Elem(), SystemAPIUserSettingMap{})
	pulumi.RegisterOutputType(SystemAPIUserSettingOutput{})
	pulumi.RegisterOutputType(SystemAPIUserSettingArrayOutput{})
	pulumi.RegisterOutputType(SystemAPIUserSettingMapOutput{})
}
