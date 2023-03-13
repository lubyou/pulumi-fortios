// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertemailSetting struct {
	pulumi.CustomResourceState

	AdminLoginLogs                    pulumi.StringOutput    `pulumi:"adminLoginLogs"`
	AlertInterval                     pulumi.IntOutput       `pulumi:"alertInterval"`
	AmcInterfaceBypassMode            pulumi.StringOutput    `pulumi:"amcInterfaceBypassMode"`
	AntivirusLogs                     pulumi.StringOutput    `pulumi:"antivirusLogs"`
	ConfigurationChangesLogs          pulumi.StringOutput    `pulumi:"configurationChangesLogs"`
	CriticalInterval                  pulumi.IntOutput       `pulumi:"criticalInterval"`
	DebugInterval                     pulumi.IntOutput       `pulumi:"debugInterval"`
	EmailInterval                     pulumi.IntOutput       `pulumi:"emailInterval"`
	EmergencyInterval                 pulumi.IntOutput       `pulumi:"emergencyInterval"`
	ErrorInterval                     pulumi.IntOutput       `pulumi:"errorInterval"`
	FdsLicenseExpiringDays            pulumi.IntOutput       `pulumi:"fdsLicenseExpiringDays"`
	FdsLicenseExpiringWarning         pulumi.StringOutput    `pulumi:"fdsLicenseExpiringWarning"`
	FdsUpdateLogs                     pulumi.StringOutput    `pulumi:"fdsUpdateLogs"`
	FilterMode                        pulumi.StringOutput    `pulumi:"filterMode"`
	FipsCcErrors                      pulumi.StringOutput    `pulumi:"fipsCcErrors"`
	FirewallAuthenticationFailureLogs pulumi.StringOutput    `pulumi:"firewallAuthenticationFailureLogs"`
	FortiguardLogQuotaWarning         pulumi.StringOutput    `pulumi:"fortiguardLogQuotaWarning"`
	FssoDisconnectLogs                pulumi.StringOutput    `pulumi:"fssoDisconnectLogs"`
	HaLogs                            pulumi.StringOutput    `pulumi:"haLogs"`
	InformationInterval               pulumi.IntOutput       `pulumi:"informationInterval"`
	IpsLogs                           pulumi.StringOutput    `pulumi:"ipsLogs"`
	IpsecErrorsLogs                   pulumi.StringOutput    `pulumi:"ipsecErrorsLogs"`
	LocalDiskUsage                    pulumi.IntOutput       `pulumi:"localDiskUsage"`
	LogDiskUsageWarning               pulumi.StringOutput    `pulumi:"logDiskUsageWarning"`
	Mailto1                           pulumi.StringOutput    `pulumi:"mailto1"`
	Mailto2                           pulumi.StringOutput    `pulumi:"mailto2"`
	Mailto3                           pulumi.StringOutput    `pulumi:"mailto3"`
	NotificationInterval              pulumi.IntOutput       `pulumi:"notificationInterval"`
	PppErrorsLogs                     pulumi.StringOutput    `pulumi:"pppErrorsLogs"`
	Severity                          pulumi.StringOutput    `pulumi:"severity"`
	SshLogs                           pulumi.StringOutput    `pulumi:"sshLogs"`
	SslvpnAuthenticationErrorsLogs    pulumi.StringOutput    `pulumi:"sslvpnAuthenticationErrorsLogs"`
	Username                          pulumi.StringOutput    `pulumi:"username"`
	Vdomparam                         pulumi.StringPtrOutput `pulumi:"vdomparam"`
	ViolationTrafficLogs              pulumi.StringOutput    `pulumi:"violationTrafficLogs"`
	WarningInterval                   pulumi.IntOutput       `pulumi:"warningInterval"`
	WebfilterLogs                     pulumi.StringOutput    `pulumi:"webfilterLogs"`
}

// NewAlertemailSetting registers a new resource with the given unique name, arguments, and options.
func NewAlertemailSetting(ctx *pulumi.Context,
	name string, args *AlertemailSettingArgs, opts ...pulumi.ResourceOption) (*AlertemailSetting, error) {
	if args == nil {
		args = &AlertemailSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AlertemailSetting
	err := ctx.RegisterResource("fortios:index/alertemailSetting:AlertemailSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertemailSetting gets an existing AlertemailSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertemailSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertemailSettingState, opts ...pulumi.ResourceOption) (*AlertemailSetting, error) {
	var resource AlertemailSetting
	err := ctx.ReadResource("fortios:index/alertemailSetting:AlertemailSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertemailSetting resources.
type alertemailSettingState struct {
	AdminLoginLogs                    *string `pulumi:"adminLoginLogs"`
	AlertInterval                     *int    `pulumi:"alertInterval"`
	AmcInterfaceBypassMode            *string `pulumi:"amcInterfaceBypassMode"`
	AntivirusLogs                     *string `pulumi:"antivirusLogs"`
	ConfigurationChangesLogs          *string `pulumi:"configurationChangesLogs"`
	CriticalInterval                  *int    `pulumi:"criticalInterval"`
	DebugInterval                     *int    `pulumi:"debugInterval"`
	EmailInterval                     *int    `pulumi:"emailInterval"`
	EmergencyInterval                 *int    `pulumi:"emergencyInterval"`
	ErrorInterval                     *int    `pulumi:"errorInterval"`
	FdsLicenseExpiringDays            *int    `pulumi:"fdsLicenseExpiringDays"`
	FdsLicenseExpiringWarning         *string `pulumi:"fdsLicenseExpiringWarning"`
	FdsUpdateLogs                     *string `pulumi:"fdsUpdateLogs"`
	FilterMode                        *string `pulumi:"filterMode"`
	FipsCcErrors                      *string `pulumi:"fipsCcErrors"`
	FirewallAuthenticationFailureLogs *string `pulumi:"firewallAuthenticationFailureLogs"`
	FortiguardLogQuotaWarning         *string `pulumi:"fortiguardLogQuotaWarning"`
	FssoDisconnectLogs                *string `pulumi:"fssoDisconnectLogs"`
	HaLogs                            *string `pulumi:"haLogs"`
	InformationInterval               *int    `pulumi:"informationInterval"`
	IpsLogs                           *string `pulumi:"ipsLogs"`
	IpsecErrorsLogs                   *string `pulumi:"ipsecErrorsLogs"`
	LocalDiskUsage                    *int    `pulumi:"localDiskUsage"`
	LogDiskUsageWarning               *string `pulumi:"logDiskUsageWarning"`
	Mailto1                           *string `pulumi:"mailto1"`
	Mailto2                           *string `pulumi:"mailto2"`
	Mailto3                           *string `pulumi:"mailto3"`
	NotificationInterval              *int    `pulumi:"notificationInterval"`
	PppErrorsLogs                     *string `pulumi:"pppErrorsLogs"`
	Severity                          *string `pulumi:"severity"`
	SshLogs                           *string `pulumi:"sshLogs"`
	SslvpnAuthenticationErrorsLogs    *string `pulumi:"sslvpnAuthenticationErrorsLogs"`
	Username                          *string `pulumi:"username"`
	Vdomparam                         *string `pulumi:"vdomparam"`
	ViolationTrafficLogs              *string `pulumi:"violationTrafficLogs"`
	WarningInterval                   *int    `pulumi:"warningInterval"`
	WebfilterLogs                     *string `pulumi:"webfilterLogs"`
}

type AlertemailSettingState struct {
	AdminLoginLogs                    pulumi.StringPtrInput
	AlertInterval                     pulumi.IntPtrInput
	AmcInterfaceBypassMode            pulumi.StringPtrInput
	AntivirusLogs                     pulumi.StringPtrInput
	ConfigurationChangesLogs          pulumi.StringPtrInput
	CriticalInterval                  pulumi.IntPtrInput
	DebugInterval                     pulumi.IntPtrInput
	EmailInterval                     pulumi.IntPtrInput
	EmergencyInterval                 pulumi.IntPtrInput
	ErrorInterval                     pulumi.IntPtrInput
	FdsLicenseExpiringDays            pulumi.IntPtrInput
	FdsLicenseExpiringWarning         pulumi.StringPtrInput
	FdsUpdateLogs                     pulumi.StringPtrInput
	FilterMode                        pulumi.StringPtrInput
	FipsCcErrors                      pulumi.StringPtrInput
	FirewallAuthenticationFailureLogs pulumi.StringPtrInput
	FortiguardLogQuotaWarning         pulumi.StringPtrInput
	FssoDisconnectLogs                pulumi.StringPtrInput
	HaLogs                            pulumi.StringPtrInput
	InformationInterval               pulumi.IntPtrInput
	IpsLogs                           pulumi.StringPtrInput
	IpsecErrorsLogs                   pulumi.StringPtrInput
	LocalDiskUsage                    pulumi.IntPtrInput
	LogDiskUsageWarning               pulumi.StringPtrInput
	Mailto1                           pulumi.StringPtrInput
	Mailto2                           pulumi.StringPtrInput
	Mailto3                           pulumi.StringPtrInput
	NotificationInterval              pulumi.IntPtrInput
	PppErrorsLogs                     pulumi.StringPtrInput
	Severity                          pulumi.StringPtrInput
	SshLogs                           pulumi.StringPtrInput
	SslvpnAuthenticationErrorsLogs    pulumi.StringPtrInput
	Username                          pulumi.StringPtrInput
	Vdomparam                         pulumi.StringPtrInput
	ViolationTrafficLogs              pulumi.StringPtrInput
	WarningInterval                   pulumi.IntPtrInput
	WebfilterLogs                     pulumi.StringPtrInput
}

func (AlertemailSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertemailSettingState)(nil)).Elem()
}

type alertemailSettingArgs struct {
	AdminLoginLogs                    *string `pulumi:"adminLoginLogs"`
	AlertInterval                     *int    `pulumi:"alertInterval"`
	AmcInterfaceBypassMode            *string `pulumi:"amcInterfaceBypassMode"`
	AntivirusLogs                     *string `pulumi:"antivirusLogs"`
	ConfigurationChangesLogs          *string `pulumi:"configurationChangesLogs"`
	CriticalInterval                  *int    `pulumi:"criticalInterval"`
	DebugInterval                     *int    `pulumi:"debugInterval"`
	EmailInterval                     *int    `pulumi:"emailInterval"`
	EmergencyInterval                 *int    `pulumi:"emergencyInterval"`
	ErrorInterval                     *int    `pulumi:"errorInterval"`
	FdsLicenseExpiringDays            *int    `pulumi:"fdsLicenseExpiringDays"`
	FdsLicenseExpiringWarning         *string `pulumi:"fdsLicenseExpiringWarning"`
	FdsUpdateLogs                     *string `pulumi:"fdsUpdateLogs"`
	FilterMode                        *string `pulumi:"filterMode"`
	FipsCcErrors                      *string `pulumi:"fipsCcErrors"`
	FirewallAuthenticationFailureLogs *string `pulumi:"firewallAuthenticationFailureLogs"`
	FortiguardLogQuotaWarning         *string `pulumi:"fortiguardLogQuotaWarning"`
	FssoDisconnectLogs                *string `pulumi:"fssoDisconnectLogs"`
	HaLogs                            *string `pulumi:"haLogs"`
	InformationInterval               *int    `pulumi:"informationInterval"`
	IpsLogs                           *string `pulumi:"ipsLogs"`
	IpsecErrorsLogs                   *string `pulumi:"ipsecErrorsLogs"`
	LocalDiskUsage                    *int    `pulumi:"localDiskUsage"`
	LogDiskUsageWarning               *string `pulumi:"logDiskUsageWarning"`
	Mailto1                           *string `pulumi:"mailto1"`
	Mailto2                           *string `pulumi:"mailto2"`
	Mailto3                           *string `pulumi:"mailto3"`
	NotificationInterval              *int    `pulumi:"notificationInterval"`
	PppErrorsLogs                     *string `pulumi:"pppErrorsLogs"`
	Severity                          *string `pulumi:"severity"`
	SshLogs                           *string `pulumi:"sshLogs"`
	SslvpnAuthenticationErrorsLogs    *string `pulumi:"sslvpnAuthenticationErrorsLogs"`
	Username                          *string `pulumi:"username"`
	Vdomparam                         *string `pulumi:"vdomparam"`
	ViolationTrafficLogs              *string `pulumi:"violationTrafficLogs"`
	WarningInterval                   *int    `pulumi:"warningInterval"`
	WebfilterLogs                     *string `pulumi:"webfilterLogs"`
}

// The set of arguments for constructing a AlertemailSetting resource.
type AlertemailSettingArgs struct {
	AdminLoginLogs                    pulumi.StringPtrInput
	AlertInterval                     pulumi.IntPtrInput
	AmcInterfaceBypassMode            pulumi.StringPtrInput
	AntivirusLogs                     pulumi.StringPtrInput
	ConfigurationChangesLogs          pulumi.StringPtrInput
	CriticalInterval                  pulumi.IntPtrInput
	DebugInterval                     pulumi.IntPtrInput
	EmailInterval                     pulumi.IntPtrInput
	EmergencyInterval                 pulumi.IntPtrInput
	ErrorInterval                     pulumi.IntPtrInput
	FdsLicenseExpiringDays            pulumi.IntPtrInput
	FdsLicenseExpiringWarning         pulumi.StringPtrInput
	FdsUpdateLogs                     pulumi.StringPtrInput
	FilterMode                        pulumi.StringPtrInput
	FipsCcErrors                      pulumi.StringPtrInput
	FirewallAuthenticationFailureLogs pulumi.StringPtrInput
	FortiguardLogQuotaWarning         pulumi.StringPtrInput
	FssoDisconnectLogs                pulumi.StringPtrInput
	HaLogs                            pulumi.StringPtrInput
	InformationInterval               pulumi.IntPtrInput
	IpsLogs                           pulumi.StringPtrInput
	IpsecErrorsLogs                   pulumi.StringPtrInput
	LocalDiskUsage                    pulumi.IntPtrInput
	LogDiskUsageWarning               pulumi.StringPtrInput
	Mailto1                           pulumi.StringPtrInput
	Mailto2                           pulumi.StringPtrInput
	Mailto3                           pulumi.StringPtrInput
	NotificationInterval              pulumi.IntPtrInput
	PppErrorsLogs                     pulumi.StringPtrInput
	Severity                          pulumi.StringPtrInput
	SshLogs                           pulumi.StringPtrInput
	SslvpnAuthenticationErrorsLogs    pulumi.StringPtrInput
	Username                          pulumi.StringPtrInput
	Vdomparam                         pulumi.StringPtrInput
	ViolationTrafficLogs              pulumi.StringPtrInput
	WarningInterval                   pulumi.IntPtrInput
	WebfilterLogs                     pulumi.StringPtrInput
}

func (AlertemailSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertemailSettingArgs)(nil)).Elem()
}

type AlertemailSettingInput interface {
	pulumi.Input

	ToAlertemailSettingOutput() AlertemailSettingOutput
	ToAlertemailSettingOutputWithContext(ctx context.Context) AlertemailSettingOutput
}

func (*AlertemailSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertemailSetting)(nil)).Elem()
}

func (i *AlertemailSetting) ToAlertemailSettingOutput() AlertemailSettingOutput {
	return i.ToAlertemailSettingOutputWithContext(context.Background())
}

func (i *AlertemailSetting) ToAlertemailSettingOutputWithContext(ctx context.Context) AlertemailSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertemailSettingOutput)
}

// AlertemailSettingArrayInput is an input type that accepts AlertemailSettingArray and AlertemailSettingArrayOutput values.
// You can construct a concrete instance of `AlertemailSettingArrayInput` via:
//
//	AlertemailSettingArray{ AlertemailSettingArgs{...} }
type AlertemailSettingArrayInput interface {
	pulumi.Input

	ToAlertemailSettingArrayOutput() AlertemailSettingArrayOutput
	ToAlertemailSettingArrayOutputWithContext(context.Context) AlertemailSettingArrayOutput
}

type AlertemailSettingArray []AlertemailSettingInput

func (AlertemailSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertemailSetting)(nil)).Elem()
}

func (i AlertemailSettingArray) ToAlertemailSettingArrayOutput() AlertemailSettingArrayOutput {
	return i.ToAlertemailSettingArrayOutputWithContext(context.Background())
}

func (i AlertemailSettingArray) ToAlertemailSettingArrayOutputWithContext(ctx context.Context) AlertemailSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertemailSettingArrayOutput)
}

// AlertemailSettingMapInput is an input type that accepts AlertemailSettingMap and AlertemailSettingMapOutput values.
// You can construct a concrete instance of `AlertemailSettingMapInput` via:
//
//	AlertemailSettingMap{ "key": AlertemailSettingArgs{...} }
type AlertemailSettingMapInput interface {
	pulumi.Input

	ToAlertemailSettingMapOutput() AlertemailSettingMapOutput
	ToAlertemailSettingMapOutputWithContext(context.Context) AlertemailSettingMapOutput
}

type AlertemailSettingMap map[string]AlertemailSettingInput

func (AlertemailSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertemailSetting)(nil)).Elem()
}

func (i AlertemailSettingMap) ToAlertemailSettingMapOutput() AlertemailSettingMapOutput {
	return i.ToAlertemailSettingMapOutputWithContext(context.Background())
}

func (i AlertemailSettingMap) ToAlertemailSettingMapOutputWithContext(ctx context.Context) AlertemailSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertemailSettingMapOutput)
}

type AlertemailSettingOutput struct{ *pulumi.OutputState }

func (AlertemailSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertemailSetting)(nil)).Elem()
}

func (o AlertemailSettingOutput) ToAlertemailSettingOutput() AlertemailSettingOutput {
	return o
}

func (o AlertemailSettingOutput) ToAlertemailSettingOutputWithContext(ctx context.Context) AlertemailSettingOutput {
	return o
}

func (o AlertemailSettingOutput) AdminLoginLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.AdminLoginLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) AlertInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.AlertInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) AmcInterfaceBypassMode() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.AmcInterfaceBypassMode }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) AntivirusLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.AntivirusLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) ConfigurationChangesLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.ConfigurationChangesLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) CriticalInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.CriticalInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) DebugInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.DebugInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) EmailInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.EmailInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) EmergencyInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.EmergencyInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) ErrorInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.ErrorInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) FdsLicenseExpiringDays() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.FdsLicenseExpiringDays }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) FdsLicenseExpiringWarning() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.FdsLicenseExpiringWarning }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) FdsUpdateLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.FdsUpdateLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) FilterMode() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.FilterMode }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) FipsCcErrors() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.FipsCcErrors }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) FirewallAuthenticationFailureLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.FirewallAuthenticationFailureLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) FortiguardLogQuotaWarning() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.FortiguardLogQuotaWarning }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) FssoDisconnectLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.FssoDisconnectLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) HaLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.HaLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) InformationInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.InformationInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) IpsLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.IpsLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) IpsecErrorsLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.IpsecErrorsLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) LocalDiskUsage() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.LocalDiskUsage }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) LogDiskUsageWarning() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.LogDiskUsageWarning }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) Mailto1() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.Mailto1 }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) Mailto2() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.Mailto2 }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) Mailto3() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.Mailto3 }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) NotificationInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.NotificationInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) PppErrorsLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.PppErrorsLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) SshLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.SshLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) SslvpnAuthenticationErrorsLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.SslvpnAuthenticationErrorsLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o AlertemailSettingOutput) ViolationTrafficLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.ViolationTrafficLogs }).(pulumi.StringOutput)
}

func (o AlertemailSettingOutput) WarningInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.IntOutput { return v.WarningInterval }).(pulumi.IntOutput)
}

func (o AlertemailSettingOutput) WebfilterLogs() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertemailSetting) pulumi.StringOutput { return v.WebfilterLogs }).(pulumi.StringOutput)
}

type AlertemailSettingArrayOutput struct{ *pulumi.OutputState }

func (AlertemailSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertemailSetting)(nil)).Elem()
}

func (o AlertemailSettingArrayOutput) ToAlertemailSettingArrayOutput() AlertemailSettingArrayOutput {
	return o
}

func (o AlertemailSettingArrayOutput) ToAlertemailSettingArrayOutputWithContext(ctx context.Context) AlertemailSettingArrayOutput {
	return o
}

func (o AlertemailSettingArrayOutput) Index(i pulumi.IntInput) AlertemailSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlertemailSetting {
		return vs[0].([]*AlertemailSetting)[vs[1].(int)]
	}).(AlertemailSettingOutput)
}

type AlertemailSettingMapOutput struct{ *pulumi.OutputState }

func (AlertemailSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertemailSetting)(nil)).Elem()
}

func (o AlertemailSettingMapOutput) ToAlertemailSettingMapOutput() AlertemailSettingMapOutput {
	return o
}

func (o AlertemailSettingMapOutput) ToAlertemailSettingMapOutputWithContext(ctx context.Context) AlertemailSettingMapOutput {
	return o
}

func (o AlertemailSettingMapOutput) MapIndex(k pulumi.StringInput) AlertemailSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlertemailSetting {
		return vs[0].(map[string]*AlertemailSetting)[vs[1].(string)]
	}).(AlertemailSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertemailSettingInput)(nil)).Elem(), &AlertemailSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertemailSettingArrayInput)(nil)).Elem(), AlertemailSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertemailSettingMapInput)(nil)).Elem(), AlertemailSettingMap{})
	pulumi.RegisterOutputType(AlertemailSettingOutput{})
	pulumi.RegisterOutputType(AlertemailSettingArrayOutput{})
	pulumi.RegisterOutputType(AlertemailSettingMapOutput{})
}
