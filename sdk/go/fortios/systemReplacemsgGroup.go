// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure replacement message groups.
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
// 		_, err := fortios.NewSystemReplacemsgGroup(ctx, "trname", &fortios.SystemReplacemsgGroupArgs{
// 			Comment:   pulumi.String("sgh"),
// 			GroupType: pulumi.String("utm"),
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
// System ReplacemsgGroup can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgGroup:SystemReplacemsgGroup labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemReplacemsgGroup struct {
	pulumi.CustomResourceState

	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins SystemReplacemsgGroupAdminArrayOutput `pulumi:"admins"`
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails SystemReplacemsgGroupAlertmailArrayOutput `pulumi:"alertmails"`
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths SystemReplacemsgGroupAuthArrayOutput `pulumi:"auths"`
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations SystemReplacemsgGroupAutomationArrayOutput `pulumi:"automations"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages SystemReplacemsgGroupCustomMessageArrayOutput `pulumi:"customMessages"`
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals SystemReplacemsgGroupDeviceDetectionPortalArrayOutput `pulumi:"deviceDetectionPortals"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs SystemReplacemsgGroupEcArrayOutput `pulumi:"ecs"`
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs SystemReplacemsgGroupFortiguardWfArrayOutput `pulumi:"fortiguardWfs"`
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps SystemReplacemsgGroupFtpArrayOutput `pulumi:"ftps"`
	// Group type.
	GroupType pulumi.StringOutput `pulumi:"groupType"`
	// Replacement message table entries. The structure of `http` block is documented below.
	Https SystemReplacemsgGroupHttpArrayOutput `pulumi:"https"`
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps SystemReplacemsgGroupIcapArrayOutput `pulumi:"icaps"`
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails SystemReplacemsgGroupMailArrayOutput `pulumi:"mails"`
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars SystemReplacemsgGroupNacQuarArrayOutput `pulumi:"nacQuars"`
	// Group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps SystemReplacemsgGroupNntpArrayOutput `pulumi:"nntps"`
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams SystemReplacemsgGroupSpamArrayOutput `pulumi:"spams"`
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns SystemReplacemsgGroupSslvpnArrayOutput `pulumi:"sslvpns"`
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas SystemReplacemsgGroupTrafficQuotaArrayOutput `pulumi:"trafficQuotas"`
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms SystemReplacemsgGroupUtmArrayOutput `pulumi:"utms"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies SystemReplacemsgGroupWebproxyArrayOutput `pulumi:"webproxies"`
}

// NewSystemReplacemsgGroup registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgGroup(ctx *pulumi.Context,
	name string, args *SystemReplacemsgGroupArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupType == nil {
		return nil, errors.New("invalid value for required argument 'GroupType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgGroup
	err := ctx.RegisterResource("fortios:index/systemReplacemsgGroup:SystemReplacemsgGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgGroup gets an existing SystemReplacemsgGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgGroupState, opts ...pulumi.ResourceOption) (*SystemReplacemsgGroup, error) {
	var resource SystemReplacemsgGroup
	err := ctx.ReadResource("fortios:index/systemReplacemsgGroup:SystemReplacemsgGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgGroup resources.
type systemReplacemsgGroupState struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins []SystemReplacemsgGroupAdmin `pulumi:"admins"`
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails []SystemReplacemsgGroupAlertmail `pulumi:"alertmails"`
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths []SystemReplacemsgGroupAuth `pulumi:"auths"`
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations []SystemReplacemsgGroupAutomation `pulumi:"automations"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages []SystemReplacemsgGroupCustomMessage `pulumi:"customMessages"`
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals []SystemReplacemsgGroupDeviceDetectionPortal `pulumi:"deviceDetectionPortals"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs []SystemReplacemsgGroupEc `pulumi:"ecs"`
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs []SystemReplacemsgGroupFortiguardWf `pulumi:"fortiguardWfs"`
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps []SystemReplacemsgGroupFtp `pulumi:"ftps"`
	// Group type.
	GroupType *string `pulumi:"groupType"`
	// Replacement message table entries. The structure of `http` block is documented below.
	Https []SystemReplacemsgGroupHttp `pulumi:"https"`
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps []SystemReplacemsgGroupIcap `pulumi:"icaps"`
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails []SystemReplacemsgGroupMail `pulumi:"mails"`
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars []SystemReplacemsgGroupNacQuar `pulumi:"nacQuars"`
	// Group name.
	Name *string `pulumi:"name"`
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps []SystemReplacemsgGroupNntp `pulumi:"nntps"`
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams []SystemReplacemsgGroupSpam `pulumi:"spams"`
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns []SystemReplacemsgGroupSslvpn `pulumi:"sslvpns"`
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas []SystemReplacemsgGroupTrafficQuota `pulumi:"trafficQuotas"`
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms []SystemReplacemsgGroupUtm `pulumi:"utms"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies []SystemReplacemsgGroupWebproxy `pulumi:"webproxies"`
}

type SystemReplacemsgGroupState struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins SystemReplacemsgGroupAdminArrayInput
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails SystemReplacemsgGroupAlertmailArrayInput
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths SystemReplacemsgGroupAuthArrayInput
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations SystemReplacemsgGroupAutomationArrayInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages SystemReplacemsgGroupCustomMessageArrayInput
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals SystemReplacemsgGroupDeviceDetectionPortalArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs SystemReplacemsgGroupEcArrayInput
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs SystemReplacemsgGroupFortiguardWfArrayInput
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps SystemReplacemsgGroupFtpArrayInput
	// Group type.
	GroupType pulumi.StringPtrInput
	// Replacement message table entries. The structure of `http` block is documented below.
	Https SystemReplacemsgGroupHttpArrayInput
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps SystemReplacemsgGroupIcapArrayInput
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails SystemReplacemsgGroupMailArrayInput
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars SystemReplacemsgGroupNacQuarArrayInput
	// Group name.
	Name pulumi.StringPtrInput
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps SystemReplacemsgGroupNntpArrayInput
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams SystemReplacemsgGroupSpamArrayInput
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns SystemReplacemsgGroupSslvpnArrayInput
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas SystemReplacemsgGroupTrafficQuotaArrayInput
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms SystemReplacemsgGroupUtmArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies SystemReplacemsgGroupWebproxyArrayInput
}

func (SystemReplacemsgGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgGroupState)(nil)).Elem()
}

type systemReplacemsgGroupArgs struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins []SystemReplacemsgGroupAdmin `pulumi:"admins"`
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails []SystemReplacemsgGroupAlertmail `pulumi:"alertmails"`
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths []SystemReplacemsgGroupAuth `pulumi:"auths"`
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations []SystemReplacemsgGroupAutomation `pulumi:"automations"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages []SystemReplacemsgGroupCustomMessage `pulumi:"customMessages"`
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals []SystemReplacemsgGroupDeviceDetectionPortal `pulumi:"deviceDetectionPortals"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs []SystemReplacemsgGroupEc `pulumi:"ecs"`
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs []SystemReplacemsgGroupFortiguardWf `pulumi:"fortiguardWfs"`
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps []SystemReplacemsgGroupFtp `pulumi:"ftps"`
	// Group type.
	GroupType string `pulumi:"groupType"`
	// Replacement message table entries. The structure of `http` block is documented below.
	Https []SystemReplacemsgGroupHttp `pulumi:"https"`
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps []SystemReplacemsgGroupIcap `pulumi:"icaps"`
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails []SystemReplacemsgGroupMail `pulumi:"mails"`
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars []SystemReplacemsgGroupNacQuar `pulumi:"nacQuars"`
	// Group name.
	Name *string `pulumi:"name"`
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps []SystemReplacemsgGroupNntp `pulumi:"nntps"`
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams []SystemReplacemsgGroupSpam `pulumi:"spams"`
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns []SystemReplacemsgGroupSslvpn `pulumi:"sslvpns"`
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas []SystemReplacemsgGroupTrafficQuota `pulumi:"trafficQuotas"`
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms []SystemReplacemsgGroupUtm `pulumi:"utms"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies []SystemReplacemsgGroupWebproxy `pulumi:"webproxies"`
}

// The set of arguments for constructing a SystemReplacemsgGroup resource.
type SystemReplacemsgGroupArgs struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins SystemReplacemsgGroupAdminArrayInput
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails SystemReplacemsgGroupAlertmailArrayInput
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths SystemReplacemsgGroupAuthArrayInput
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations SystemReplacemsgGroupAutomationArrayInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages SystemReplacemsgGroupCustomMessageArrayInput
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals SystemReplacemsgGroupDeviceDetectionPortalArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs SystemReplacemsgGroupEcArrayInput
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs SystemReplacemsgGroupFortiguardWfArrayInput
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps SystemReplacemsgGroupFtpArrayInput
	// Group type.
	GroupType pulumi.StringInput
	// Replacement message table entries. The structure of `http` block is documented below.
	Https SystemReplacemsgGroupHttpArrayInput
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps SystemReplacemsgGroupIcapArrayInput
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails SystemReplacemsgGroupMailArrayInput
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars SystemReplacemsgGroupNacQuarArrayInput
	// Group name.
	Name pulumi.StringPtrInput
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps SystemReplacemsgGroupNntpArrayInput
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams SystemReplacemsgGroupSpamArrayInput
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns SystemReplacemsgGroupSslvpnArrayInput
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas SystemReplacemsgGroupTrafficQuotaArrayInput
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms SystemReplacemsgGroupUtmArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies SystemReplacemsgGroupWebproxyArrayInput
}

func (SystemReplacemsgGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgGroupArgs)(nil)).Elem()
}

type SystemReplacemsgGroupInput interface {
	pulumi.Input

	ToSystemReplacemsgGroupOutput() SystemReplacemsgGroupOutput
	ToSystemReplacemsgGroupOutputWithContext(ctx context.Context) SystemReplacemsgGroupOutput
}

func (*SystemReplacemsgGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgGroup)(nil)).Elem()
}

func (i *SystemReplacemsgGroup) ToSystemReplacemsgGroupOutput() SystemReplacemsgGroupOutput {
	return i.ToSystemReplacemsgGroupOutputWithContext(context.Background())
}

func (i *SystemReplacemsgGroup) ToSystemReplacemsgGroupOutputWithContext(ctx context.Context) SystemReplacemsgGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgGroupOutput)
}

// SystemReplacemsgGroupArrayInput is an input type that accepts SystemReplacemsgGroupArray and SystemReplacemsgGroupArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgGroupArrayInput` via:
//
//          SystemReplacemsgGroupArray{ SystemReplacemsgGroupArgs{...} }
type SystemReplacemsgGroupArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgGroupArrayOutput() SystemReplacemsgGroupArrayOutput
	ToSystemReplacemsgGroupArrayOutputWithContext(context.Context) SystemReplacemsgGroupArrayOutput
}

type SystemReplacemsgGroupArray []SystemReplacemsgGroupInput

func (SystemReplacemsgGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgGroup)(nil)).Elem()
}

func (i SystemReplacemsgGroupArray) ToSystemReplacemsgGroupArrayOutput() SystemReplacemsgGroupArrayOutput {
	return i.ToSystemReplacemsgGroupArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgGroupArray) ToSystemReplacemsgGroupArrayOutputWithContext(ctx context.Context) SystemReplacemsgGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgGroupArrayOutput)
}

// SystemReplacemsgGroupMapInput is an input type that accepts SystemReplacemsgGroupMap and SystemReplacemsgGroupMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgGroupMapInput` via:
//
//          SystemReplacemsgGroupMap{ "key": SystemReplacemsgGroupArgs{...} }
type SystemReplacemsgGroupMapInput interface {
	pulumi.Input

	ToSystemReplacemsgGroupMapOutput() SystemReplacemsgGroupMapOutput
	ToSystemReplacemsgGroupMapOutputWithContext(context.Context) SystemReplacemsgGroupMapOutput
}

type SystemReplacemsgGroupMap map[string]SystemReplacemsgGroupInput

func (SystemReplacemsgGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgGroup)(nil)).Elem()
}

func (i SystemReplacemsgGroupMap) ToSystemReplacemsgGroupMapOutput() SystemReplacemsgGroupMapOutput {
	return i.ToSystemReplacemsgGroupMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgGroupMap) ToSystemReplacemsgGroupMapOutputWithContext(ctx context.Context) SystemReplacemsgGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgGroupMapOutput)
}

type SystemReplacemsgGroupOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgGroup)(nil)).Elem()
}

func (o SystemReplacemsgGroupOutput) ToSystemReplacemsgGroupOutput() SystemReplacemsgGroupOutput {
	return o
}

func (o SystemReplacemsgGroupOutput) ToSystemReplacemsgGroupOutputWithContext(ctx context.Context) SystemReplacemsgGroupOutput {
	return o
}

type SystemReplacemsgGroupArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgGroup)(nil)).Elem()
}

func (o SystemReplacemsgGroupArrayOutput) ToSystemReplacemsgGroupArrayOutput() SystemReplacemsgGroupArrayOutput {
	return o
}

func (o SystemReplacemsgGroupArrayOutput) ToSystemReplacemsgGroupArrayOutputWithContext(ctx context.Context) SystemReplacemsgGroupArrayOutput {
	return o
}

func (o SystemReplacemsgGroupArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgGroup {
		return vs[0].([]*SystemReplacemsgGroup)[vs[1].(int)]
	}).(SystemReplacemsgGroupOutput)
}

type SystemReplacemsgGroupMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgGroup)(nil)).Elem()
}

func (o SystemReplacemsgGroupMapOutput) ToSystemReplacemsgGroupMapOutput() SystemReplacemsgGroupMapOutput {
	return o
}

func (o SystemReplacemsgGroupMapOutput) ToSystemReplacemsgGroupMapOutputWithContext(ctx context.Context) SystemReplacemsgGroupMapOutput {
	return o
}

func (o SystemReplacemsgGroupMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgGroup {
		return vs[0].(map[string]*SystemReplacemsgGroup)[vs[1].(string)]
	}).(SystemReplacemsgGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgGroupInput)(nil)).Elem(), &SystemReplacemsgGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgGroupArrayInput)(nil)).Elem(), SystemReplacemsgGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgGroupMapInput)(nil)).Elem(), SystemReplacemsgGroupMap{})
	pulumi.RegisterOutputType(SystemReplacemsgGroupOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgGroupArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgGroupMapOutput{})
}
