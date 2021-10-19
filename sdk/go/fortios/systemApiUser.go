// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// System ApiUser can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemApiUser:SystemApiUser labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemApiUser struct {
	pulumi.CustomResourceState

	// Admin user access profile.
	Accprofile pulumi.StringOutput `pulumi:"accprofile"`
	// Admin user password.
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin pulumi.StringOutput `pulumi:"corsAllowOrigin"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Virtual domain name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth pulumi.StringOutput `pulumi:"peerAuth"`
	// Peer group name.
	PeerGroup pulumi.StringOutput `pulumi:"peerGroup"`
	// Schedule name.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts SystemApiUserTrusthostArrayOutput `pulumi:"trusthosts"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms SystemApiUserVdomArrayOutput `pulumi:"vdoms"`
}

// NewSystemApiUser registers a new resource with the given unique name, arguments, and options.
func NewSystemApiUser(ctx *pulumi.Context,
	name string, args *SystemApiUserArgs, opts ...pulumi.ResourceOption) (*SystemApiUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Accprofile == nil {
		return nil, errors.New("invalid value for required argument 'Accprofile'")
	}
	var resource SystemApiUser
	err := ctx.RegisterResource("fortios:index/systemApiUser:SystemApiUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemApiUser gets an existing SystemApiUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemApiUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemApiUserState, opts ...pulumi.ResourceOption) (*SystemApiUser, error) {
	var resource SystemApiUser
	err := ctx.ReadResource("fortios:index/systemApiUser:SystemApiUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemApiUser resources.
type systemApiUserState struct {
	// Admin user access profile.
	Accprofile *string `pulumi:"accprofile"`
	// Admin user password.
	ApiKey *string `pulumi:"apiKey"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin *string `pulumi:"corsAllowOrigin"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Virtual domain name.
	Name *string `pulumi:"name"`
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth *string `pulumi:"peerAuth"`
	// Peer group name.
	PeerGroup *string `pulumi:"peerGroup"`
	// Schedule name.
	Schedule *string `pulumi:"schedule"`
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts []SystemApiUserTrusthost `pulumi:"trusthosts"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms []SystemApiUserVdom `pulumi:"vdoms"`
}

type SystemApiUserState struct {
	// Admin user access profile.
	Accprofile pulumi.StringPtrInput
	// Admin user password.
	ApiKey pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Virtual domain name.
	Name pulumi.StringPtrInput
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth pulumi.StringPtrInput
	// Peer group name.
	PeerGroup pulumi.StringPtrInput
	// Schedule name.
	Schedule pulumi.StringPtrInput
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts SystemApiUserTrusthostArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms SystemApiUserVdomArrayInput
}

func (SystemApiUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemApiUserState)(nil)).Elem()
}

type systemApiUserArgs struct {
	// Admin user access profile.
	Accprofile string `pulumi:"accprofile"`
	// Admin user password.
	ApiKey *string `pulumi:"apiKey"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin *string `pulumi:"corsAllowOrigin"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Virtual domain name.
	Name *string `pulumi:"name"`
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth *string `pulumi:"peerAuth"`
	// Peer group name.
	PeerGroup *string `pulumi:"peerGroup"`
	// Schedule name.
	Schedule *string `pulumi:"schedule"`
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts []SystemApiUserTrusthost `pulumi:"trusthosts"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms []SystemApiUserVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a SystemApiUser resource.
type SystemApiUserArgs struct {
	// Admin user access profile.
	Accprofile pulumi.StringInput
	// Admin user password.
	ApiKey pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Virtual domain name.
	Name pulumi.StringPtrInput
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth pulumi.StringPtrInput
	// Peer group name.
	PeerGroup pulumi.StringPtrInput
	// Schedule name.
	Schedule pulumi.StringPtrInput
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts SystemApiUserTrusthostArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms SystemApiUserVdomArrayInput
}

func (SystemApiUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemApiUserArgs)(nil)).Elem()
}

type SystemApiUserInput interface {
	pulumi.Input

	ToSystemApiUserOutput() SystemApiUserOutput
	ToSystemApiUserOutputWithContext(ctx context.Context) SystemApiUserOutput
}

func (*SystemApiUser) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemApiUser)(nil))
}

func (i *SystemApiUser) ToSystemApiUserOutput() SystemApiUserOutput {
	return i.ToSystemApiUserOutputWithContext(context.Background())
}

func (i *SystemApiUser) ToSystemApiUserOutputWithContext(ctx context.Context) SystemApiUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemApiUserOutput)
}

func (i *SystemApiUser) ToSystemApiUserPtrOutput() SystemApiUserPtrOutput {
	return i.ToSystemApiUserPtrOutputWithContext(context.Background())
}

func (i *SystemApiUser) ToSystemApiUserPtrOutputWithContext(ctx context.Context) SystemApiUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemApiUserPtrOutput)
}

type SystemApiUserPtrInput interface {
	pulumi.Input

	ToSystemApiUserPtrOutput() SystemApiUserPtrOutput
	ToSystemApiUserPtrOutputWithContext(ctx context.Context) SystemApiUserPtrOutput
}

type systemApiUserPtrType SystemApiUserArgs

func (*systemApiUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemApiUser)(nil))
}

func (i *systemApiUserPtrType) ToSystemApiUserPtrOutput() SystemApiUserPtrOutput {
	return i.ToSystemApiUserPtrOutputWithContext(context.Background())
}

func (i *systemApiUserPtrType) ToSystemApiUserPtrOutputWithContext(ctx context.Context) SystemApiUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemApiUserPtrOutput)
}

// SystemApiUserArrayInput is an input type that accepts SystemApiUserArray and SystemApiUserArrayOutput values.
// You can construct a concrete instance of `SystemApiUserArrayInput` via:
//
//          SystemApiUserArray{ SystemApiUserArgs{...} }
type SystemApiUserArrayInput interface {
	pulumi.Input

	ToSystemApiUserArrayOutput() SystemApiUserArrayOutput
	ToSystemApiUserArrayOutputWithContext(context.Context) SystemApiUserArrayOutput
}

type SystemApiUserArray []SystemApiUserInput

func (SystemApiUserArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemApiUser)(nil))
}

func (i SystemApiUserArray) ToSystemApiUserArrayOutput() SystemApiUserArrayOutput {
	return i.ToSystemApiUserArrayOutputWithContext(context.Background())
}

func (i SystemApiUserArray) ToSystemApiUserArrayOutputWithContext(ctx context.Context) SystemApiUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemApiUserArrayOutput)
}

// SystemApiUserMapInput is an input type that accepts SystemApiUserMap and SystemApiUserMapOutput values.
// You can construct a concrete instance of `SystemApiUserMapInput` via:
//
//          SystemApiUserMap{ "key": SystemApiUserArgs{...} }
type SystemApiUserMapInput interface {
	pulumi.Input

	ToSystemApiUserMapOutput() SystemApiUserMapOutput
	ToSystemApiUserMapOutputWithContext(context.Context) SystemApiUserMapOutput
}

type SystemApiUserMap map[string]SystemApiUserInput

func (SystemApiUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemApiUser)(nil))
}

func (i SystemApiUserMap) ToSystemApiUserMapOutput() SystemApiUserMapOutput {
	return i.ToSystemApiUserMapOutputWithContext(context.Background())
}

func (i SystemApiUserMap) ToSystemApiUserMapOutputWithContext(ctx context.Context) SystemApiUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemApiUserMapOutput)
}

type SystemApiUserOutput struct {
	*pulumi.OutputState
}

func (SystemApiUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemApiUser)(nil))
}

func (o SystemApiUserOutput) ToSystemApiUserOutput() SystemApiUserOutput {
	return o
}

func (o SystemApiUserOutput) ToSystemApiUserOutputWithContext(ctx context.Context) SystemApiUserOutput {
	return o
}

func (o SystemApiUserOutput) ToSystemApiUserPtrOutput() SystemApiUserPtrOutput {
	return o.ToSystemApiUserPtrOutputWithContext(context.Background())
}

func (o SystemApiUserOutput) ToSystemApiUserPtrOutputWithContext(ctx context.Context) SystemApiUserPtrOutput {
	return o.ApplyT(func(v SystemApiUser) *SystemApiUser {
		return &v
	}).(SystemApiUserPtrOutput)
}

type SystemApiUserPtrOutput struct {
	*pulumi.OutputState
}

func (SystemApiUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemApiUser)(nil))
}

func (o SystemApiUserPtrOutput) ToSystemApiUserPtrOutput() SystemApiUserPtrOutput {
	return o
}

func (o SystemApiUserPtrOutput) ToSystemApiUserPtrOutputWithContext(ctx context.Context) SystemApiUserPtrOutput {
	return o
}

type SystemApiUserArrayOutput struct{ *pulumi.OutputState }

func (SystemApiUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemApiUser)(nil))
}

func (o SystemApiUserArrayOutput) ToSystemApiUserArrayOutput() SystemApiUserArrayOutput {
	return o
}

func (o SystemApiUserArrayOutput) ToSystemApiUserArrayOutputWithContext(ctx context.Context) SystemApiUserArrayOutput {
	return o
}

func (o SystemApiUserArrayOutput) Index(i pulumi.IntInput) SystemApiUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemApiUser {
		return vs[0].([]SystemApiUser)[vs[1].(int)]
	}).(SystemApiUserOutput)
}

type SystemApiUserMapOutput struct{ *pulumi.OutputState }

func (SystemApiUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemApiUser)(nil))
}

func (o SystemApiUserMapOutput) ToSystemApiUserMapOutput() SystemApiUserMapOutput {
	return o
}

func (o SystemApiUserMapOutput) ToSystemApiUserMapOutputWithContext(ctx context.Context) SystemApiUserMapOutput {
	return o
}

func (o SystemApiUserMapOutput) MapIndex(k pulumi.StringInput) SystemApiUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemApiUser {
		return vs[0].(map[string]SystemApiUser)[vs[1].(string)]
	}).(SystemApiUserOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemApiUserOutput{})
	pulumi.RegisterOutputType(SystemApiUserPtrOutput{})
	pulumi.RegisterOutputType(SystemApiUserArrayOutput{})
	pulumi.RegisterOutputType(SystemApiUserMapOutput{})
}
