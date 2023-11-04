// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type UserFortitoken struct {
	pulumi.CustomResourceState

	ActivationCode   pulumi.StringOutput    `pulumi:"activationCode"`
	ActivationExpire pulumi.IntOutput       `pulumi:"activationExpire"`
	Comments         pulumi.StringPtrOutput `pulumi:"comments"`
	License          pulumi.StringOutput    `pulumi:"license"`
	OsVer            pulumi.StringOutput    `pulumi:"osVer"`
	RegId            pulumi.StringOutput    `pulumi:"regId"`
	Seed             pulumi.StringOutput    `pulumi:"seed"`
	SerialNumber     pulumi.StringOutput    `pulumi:"serialNumber"`
	Status           pulumi.StringOutput    `pulumi:"status"`
	Vdomparam        pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserFortitoken registers a new resource with the given unique name, arguments, and options.
func NewUserFortitoken(ctx *pulumi.Context,
	name string, args *UserFortitokenArgs, opts ...pulumi.ResourceOption) (*UserFortitoken, error) {
	if args == nil {
		args = &UserFortitokenArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserFortitoken
	err := ctx.RegisterResource("fortios:index/userFortitoken:UserFortitoken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserFortitoken gets an existing UserFortitoken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserFortitoken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserFortitokenState, opts ...pulumi.ResourceOption) (*UserFortitoken, error) {
	var resource UserFortitoken
	err := ctx.ReadResource("fortios:index/userFortitoken:UserFortitoken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserFortitoken resources.
type userFortitokenState struct {
	ActivationCode   *string `pulumi:"activationCode"`
	ActivationExpire *int    `pulumi:"activationExpire"`
	Comments         *string `pulumi:"comments"`
	License          *string `pulumi:"license"`
	OsVer            *string `pulumi:"osVer"`
	RegId            *string `pulumi:"regId"`
	Seed             *string `pulumi:"seed"`
	SerialNumber     *string `pulumi:"serialNumber"`
	Status           *string `pulumi:"status"`
	Vdomparam        *string `pulumi:"vdomparam"`
}

type UserFortitokenState struct {
	ActivationCode   pulumi.StringPtrInput
	ActivationExpire pulumi.IntPtrInput
	Comments         pulumi.StringPtrInput
	License          pulumi.StringPtrInput
	OsVer            pulumi.StringPtrInput
	RegId            pulumi.StringPtrInput
	Seed             pulumi.StringPtrInput
	SerialNumber     pulumi.StringPtrInput
	Status           pulumi.StringPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (UserFortitokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*userFortitokenState)(nil)).Elem()
}

type userFortitokenArgs struct {
	ActivationCode   *string `pulumi:"activationCode"`
	ActivationExpire *int    `pulumi:"activationExpire"`
	Comments         *string `pulumi:"comments"`
	License          *string `pulumi:"license"`
	OsVer            *string `pulumi:"osVer"`
	RegId            *string `pulumi:"regId"`
	Seed             *string `pulumi:"seed"`
	SerialNumber     *string `pulumi:"serialNumber"`
	Status           *string `pulumi:"status"`
	Vdomparam        *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserFortitoken resource.
type UserFortitokenArgs struct {
	ActivationCode   pulumi.StringPtrInput
	ActivationExpire pulumi.IntPtrInput
	Comments         pulumi.StringPtrInput
	License          pulumi.StringPtrInput
	OsVer            pulumi.StringPtrInput
	RegId            pulumi.StringPtrInput
	Seed             pulumi.StringPtrInput
	SerialNumber     pulumi.StringPtrInput
	Status           pulumi.StringPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (UserFortitokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userFortitokenArgs)(nil)).Elem()
}

type UserFortitokenInput interface {
	pulumi.Input

	ToUserFortitokenOutput() UserFortitokenOutput
	ToUserFortitokenOutputWithContext(ctx context.Context) UserFortitokenOutput
}

func (*UserFortitoken) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFortitoken)(nil)).Elem()
}

func (i *UserFortitoken) ToUserFortitokenOutput() UserFortitokenOutput {
	return i.ToUserFortitokenOutputWithContext(context.Background())
}

func (i *UserFortitoken) ToUserFortitokenOutputWithContext(ctx context.Context) UserFortitokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFortitokenOutput)
}

func (i *UserFortitoken) ToOutput(ctx context.Context) pulumix.Output[*UserFortitoken] {
	return pulumix.Output[*UserFortitoken]{
		OutputState: i.ToUserFortitokenOutputWithContext(ctx).OutputState,
	}
}

// UserFortitokenArrayInput is an input type that accepts UserFortitokenArray and UserFortitokenArrayOutput values.
// You can construct a concrete instance of `UserFortitokenArrayInput` via:
//
//	UserFortitokenArray{ UserFortitokenArgs{...} }
type UserFortitokenArrayInput interface {
	pulumi.Input

	ToUserFortitokenArrayOutput() UserFortitokenArrayOutput
	ToUserFortitokenArrayOutputWithContext(context.Context) UserFortitokenArrayOutput
}

type UserFortitokenArray []UserFortitokenInput

func (UserFortitokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFortitoken)(nil)).Elem()
}

func (i UserFortitokenArray) ToUserFortitokenArrayOutput() UserFortitokenArrayOutput {
	return i.ToUserFortitokenArrayOutputWithContext(context.Background())
}

func (i UserFortitokenArray) ToUserFortitokenArrayOutputWithContext(ctx context.Context) UserFortitokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFortitokenArrayOutput)
}

func (i UserFortitokenArray) ToOutput(ctx context.Context) pulumix.Output[[]*UserFortitoken] {
	return pulumix.Output[[]*UserFortitoken]{
		OutputState: i.ToUserFortitokenArrayOutputWithContext(ctx).OutputState,
	}
}

// UserFortitokenMapInput is an input type that accepts UserFortitokenMap and UserFortitokenMapOutput values.
// You can construct a concrete instance of `UserFortitokenMapInput` via:
//
//	UserFortitokenMap{ "key": UserFortitokenArgs{...} }
type UserFortitokenMapInput interface {
	pulumi.Input

	ToUserFortitokenMapOutput() UserFortitokenMapOutput
	ToUserFortitokenMapOutputWithContext(context.Context) UserFortitokenMapOutput
}

type UserFortitokenMap map[string]UserFortitokenInput

func (UserFortitokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFortitoken)(nil)).Elem()
}

func (i UserFortitokenMap) ToUserFortitokenMapOutput() UserFortitokenMapOutput {
	return i.ToUserFortitokenMapOutputWithContext(context.Background())
}

func (i UserFortitokenMap) ToUserFortitokenMapOutputWithContext(ctx context.Context) UserFortitokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFortitokenMapOutput)
}

func (i UserFortitokenMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserFortitoken] {
	return pulumix.Output[map[string]*UserFortitoken]{
		OutputState: i.ToUserFortitokenMapOutputWithContext(ctx).OutputState,
	}
}

type UserFortitokenOutput struct{ *pulumi.OutputState }

func (UserFortitokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFortitoken)(nil)).Elem()
}

func (o UserFortitokenOutput) ToUserFortitokenOutput() UserFortitokenOutput {
	return o
}

func (o UserFortitokenOutput) ToUserFortitokenOutputWithContext(ctx context.Context) UserFortitokenOutput {
	return o
}

func (o UserFortitokenOutput) ToOutput(ctx context.Context) pulumix.Output[*UserFortitoken] {
	return pulumix.Output[*UserFortitoken]{
		OutputState: o.OutputState,
	}
}

func (o UserFortitokenOutput) ActivationCode() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.ActivationCode }).(pulumi.StringOutput)
}

func (o UserFortitokenOutput) ActivationExpire() pulumi.IntOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.IntOutput { return v.ActivationExpire }).(pulumi.IntOutput)
}

func (o UserFortitokenOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o UserFortitokenOutput) License() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.License }).(pulumi.StringOutput)
}

func (o UserFortitokenOutput) OsVer() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.OsVer }).(pulumi.StringOutput)
}

func (o UserFortitokenOutput) RegId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.RegId }).(pulumi.StringOutput)
}

func (o UserFortitokenOutput) Seed() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.Seed }).(pulumi.StringOutput)
}

func (o UserFortitokenOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o UserFortitokenOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o UserFortitokenOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserFortitokenArrayOutput struct{ *pulumi.OutputState }

func (UserFortitokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFortitoken)(nil)).Elem()
}

func (o UserFortitokenArrayOutput) ToUserFortitokenArrayOutput() UserFortitokenArrayOutput {
	return o
}

func (o UserFortitokenArrayOutput) ToUserFortitokenArrayOutputWithContext(ctx context.Context) UserFortitokenArrayOutput {
	return o
}

func (o UserFortitokenArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*UserFortitoken] {
	return pulumix.Output[[]*UserFortitoken]{
		OutputState: o.OutputState,
	}
}

func (o UserFortitokenArrayOutput) Index(i pulumi.IntInput) UserFortitokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserFortitoken {
		return vs[0].([]*UserFortitoken)[vs[1].(int)]
	}).(UserFortitokenOutput)
}

type UserFortitokenMapOutput struct{ *pulumi.OutputState }

func (UserFortitokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFortitoken)(nil)).Elem()
}

func (o UserFortitokenMapOutput) ToUserFortitokenMapOutput() UserFortitokenMapOutput {
	return o
}

func (o UserFortitokenMapOutput) ToUserFortitokenMapOutputWithContext(ctx context.Context) UserFortitokenMapOutput {
	return o
}

func (o UserFortitokenMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserFortitoken] {
	return pulumix.Output[map[string]*UserFortitoken]{
		OutputState: o.OutputState,
	}
}

func (o UserFortitokenMapOutput) MapIndex(k pulumi.StringInput) UserFortitokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserFortitoken {
		return vs[0].(map[string]*UserFortitoken)[vs[1].(string)]
	}).(UserFortitokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserFortitokenInput)(nil)).Elem(), &UserFortitoken{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFortitokenArrayInput)(nil)).Elem(), UserFortitokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFortitokenMapInput)(nil)).Elem(), UserFortitokenMap{})
	pulumi.RegisterOutputType(UserFortitokenOutput{})
	pulumi.RegisterOutputType(UserFortitokenArrayOutput{})
	pulumi.RegisterOutputType(UserFortitokenMapOutput{})
}
