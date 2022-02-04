// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system replacemsggroup
func LookupSystemReplacemsgGroup(ctx *pulumi.Context, args *LookupSystemReplacemsgGroupArgs, opts ...pulumi.InvokeOption) (*LookupSystemReplacemsgGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemReplacemsgGroupResult
	err := ctx.Invoke("fortios:index/getSystemReplacemsgGroup:GetSystemReplacemsgGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemReplacemsgGroup.
type LookupSystemReplacemsgGroupArgs struct {
	// Specify the name of the desired system replacemsggroup.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemReplacemsgGroup.
type LookupSystemReplacemsgGroupResult struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins []GetSystemReplacemsgGroupAdmin `pulumi:"admins"`
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails []GetSystemReplacemsgGroupAlertmail `pulumi:"alertmails"`
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths []GetSystemReplacemsgGroupAuth `pulumi:"auths"`
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations []GetSystemReplacemsgGroupAutomation `pulumi:"automations"`
	// Comment.
	Comment string `pulumi:"comment"`
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages []GetSystemReplacemsgGroupCustomMessage `pulumi:"customMessages"`
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals []GetSystemReplacemsgGroupDeviceDetectionPortal `pulumi:"deviceDetectionPortals"`
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs []GetSystemReplacemsgGroupEc `pulumi:"ecs"`
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs []GetSystemReplacemsgGroupFortiguardWf `pulumi:"fortiguardWfs"`
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps []GetSystemReplacemsgGroupFtp `pulumi:"ftps"`
	// Group type.
	GroupType string `pulumi:"groupType"`
	// Replacement message table entries. The structure of `http` block is documented below.
	Https []GetSystemReplacemsgGroupHttp `pulumi:"https"`
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps []GetSystemReplacemsgGroupIcap `pulumi:"icaps"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails []GetSystemReplacemsgGroupMail `pulumi:"mails"`
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars []GetSystemReplacemsgGroupNacQuar `pulumi:"nacQuars"`
	// Group name.
	Name string `pulumi:"name"`
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps []GetSystemReplacemsgGroupNntp `pulumi:"nntps"`
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams []GetSystemReplacemsgGroupSpam `pulumi:"spams"`
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns []GetSystemReplacemsgGroupSslvpn `pulumi:"sslvpns"`
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas []GetSystemReplacemsgGroupTrafficQuota `pulumi:"trafficQuotas"`
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms      []GetSystemReplacemsgGroupUtm `pulumi:"utms"`
	Vdomparam *string                       `pulumi:"vdomparam"`
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies []GetSystemReplacemsgGroupWebproxy `pulumi:"webproxies"`
}

func LookupSystemReplacemsgGroupOutput(ctx *pulumi.Context, args LookupSystemReplacemsgGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSystemReplacemsgGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemReplacemsgGroupResult, error) {
			args := v.(LookupSystemReplacemsgGroupArgs)
			r, err := LookupSystemReplacemsgGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemReplacemsgGroupResultOutput)
}

// A collection of arguments for invoking GetSystemReplacemsgGroup.
type LookupSystemReplacemsgGroupOutputArgs struct {
	// Specify the name of the desired system replacemsggroup.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemReplacemsgGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemReplacemsgGroupArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemReplacemsgGroup.
type LookupSystemReplacemsgGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSystemReplacemsgGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemReplacemsgGroupResult)(nil)).Elem()
}

func (o LookupSystemReplacemsgGroupResultOutput) ToLookupSystemReplacemsgGroupResultOutput() LookupSystemReplacemsgGroupResultOutput {
	return o
}

func (o LookupSystemReplacemsgGroupResultOutput) ToLookupSystemReplacemsgGroupResultOutputWithContext(ctx context.Context) LookupSystemReplacemsgGroupResultOutput {
	return o
}

// Replacement message table entries. The structure of `admin` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Admins() GetSystemReplacemsgGroupAdminArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupAdmin { return v.Admins }).(GetSystemReplacemsgGroupAdminArrayOutput)
}

// Replacement message table entries. The structure of `alertmail` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Alertmails() GetSystemReplacemsgGroupAlertmailArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupAlertmail { return v.Alertmails }).(GetSystemReplacemsgGroupAlertmailArrayOutput)
}

// Replacement message table entries. The structure of `auth` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Auths() GetSystemReplacemsgGroupAuthArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupAuth { return v.Auths }).(GetSystemReplacemsgGroupAuthArrayOutput)
}

// Replacement message table entries. The structure of `automation` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Automations() GetSystemReplacemsgGroupAutomationArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupAutomation { return v.Automations }).(GetSystemReplacemsgGroupAutomationArrayOutput)
}

// Comment.
func (o LookupSystemReplacemsgGroupResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `customMessage` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) CustomMessages() GetSystemReplacemsgGroupCustomMessageArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupCustomMessage {
		return v.CustomMessages
	}).(GetSystemReplacemsgGroupCustomMessageArrayOutput)
}

// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) DeviceDetectionPortals() GetSystemReplacemsgGroupDeviceDetectionPortalArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupDeviceDetectionPortal {
		return v.DeviceDetectionPortals
	}).(GetSystemReplacemsgGroupDeviceDetectionPortalArrayOutput)
}

// Replacement message table entries. The structure of `ec` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Ecs() GetSystemReplacemsgGroupEcArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupEc { return v.Ecs }).(GetSystemReplacemsgGroupEcArrayOutput)
}

// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) FortiguardWfs() GetSystemReplacemsgGroupFortiguardWfArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupFortiguardWf {
		return v.FortiguardWfs
	}).(GetSystemReplacemsgGroupFortiguardWfArrayOutput)
}

// Replacement message table entries. The structure of `ftp` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Ftps() GetSystemReplacemsgGroupFtpArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupFtp { return v.Ftps }).(GetSystemReplacemsgGroupFtpArrayOutput)
}

// Group type.
func (o LookupSystemReplacemsgGroupResultOutput) GroupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) string { return v.GroupType }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `http` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Https() GetSystemReplacemsgGroupHttpArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupHttp { return v.Https }).(GetSystemReplacemsgGroupHttpArrayOutput)
}

// Replacement message table entries. The structure of `icap` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Icaps() GetSystemReplacemsgGroupIcapArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupIcap { return v.Icaps }).(GetSystemReplacemsgGroupIcapArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemReplacemsgGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `mail` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Mails() GetSystemReplacemsgGroupMailArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupMail { return v.Mails }).(GetSystemReplacemsgGroupMailArrayOutput)
}

// Replacement message table entries. The structure of `nacQuar` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) NacQuars() GetSystemReplacemsgGroupNacQuarArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupNacQuar { return v.NacQuars }).(GetSystemReplacemsgGroupNacQuarArrayOutput)
}

// Group name.
func (o LookupSystemReplacemsgGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `nntp` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Nntps() GetSystemReplacemsgGroupNntpArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupNntp { return v.Nntps }).(GetSystemReplacemsgGroupNntpArrayOutput)
}

// Replacement message table entries. The structure of `spam` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Spams() GetSystemReplacemsgGroupSpamArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupSpam { return v.Spams }).(GetSystemReplacemsgGroupSpamArrayOutput)
}

// Replacement message table entries. The structure of `sslvpn` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Sslvpns() GetSystemReplacemsgGroupSslvpnArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupSslvpn { return v.Sslvpns }).(GetSystemReplacemsgGroupSslvpnArrayOutput)
}

// Replacement message table entries. The structure of `trafficQuota` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) TrafficQuotas() GetSystemReplacemsgGroupTrafficQuotaArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupTrafficQuota {
		return v.TrafficQuotas
	}).(GetSystemReplacemsgGroupTrafficQuotaArrayOutput)
}

// Replacement message table entries. The structure of `utm` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Utms() GetSystemReplacemsgGroupUtmArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupUtm { return v.Utms }).(GetSystemReplacemsgGroupUtmArrayOutput)
}

func (o LookupSystemReplacemsgGroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Replacement message table entries. The structure of `webproxy` block is documented below.
func (o LookupSystemReplacemsgGroupResultOutput) Webproxies() GetSystemReplacemsgGroupWebproxyArrayOutput {
	return o.ApplyT(func(v LookupSystemReplacemsgGroupResult) []GetSystemReplacemsgGroupWebproxy { return v.Webproxies }).(GetSystemReplacemsgGroupWebproxyArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemReplacemsgGroupResultOutput{})
}
