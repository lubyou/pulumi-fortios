// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WirelessControllerWtp struct {
	pulumi.CustomResourceState

	Admin                          pulumi.StringOutput                               `pulumi:"admin"`
	Allowaccess                    pulumi.StringOutput                               `pulumi:"allowaccess"`
	ApcfgProfile                   pulumi.StringOutput                               `pulumi:"apcfgProfile"`
	BonjourProfile                 pulumi.StringOutput                               `pulumi:"bonjourProfile"`
	CoordinateLatitude             pulumi.StringOutput                               `pulumi:"coordinateLatitude"`
	CoordinateLongitude            pulumi.StringOutput                               `pulumi:"coordinateLongitude"`
	DynamicSortSubtable            pulumi.StringPtrOutput                            `pulumi:"dynamicSortSubtable"`
	FirmwareProvision              pulumi.StringOutput                               `pulumi:"firmwareProvision"`
	FirmwareProvisionLatest        pulumi.StringOutput                               `pulumi:"firmwareProvisionLatest"`
	ImageDownload                  pulumi.StringOutput                               `pulumi:"imageDownload"`
	Index                          pulumi.IntOutput                                  `pulumi:"index"`
	IpFragmentPreventing           pulumi.StringOutput                               `pulumi:"ipFragmentPreventing"`
	Lan                            WirelessControllerWtpLanOutput                    `pulumi:"lan"`
	LedState                       pulumi.StringOutput                               `pulumi:"ledState"`
	Location                       pulumi.StringOutput                               `pulumi:"location"`
	LoginPasswd                    pulumi.StringPtrOutput                            `pulumi:"loginPasswd"`
	LoginPasswdChange              pulumi.StringOutput                               `pulumi:"loginPasswdChange"`
	MeshBridgeEnable               pulumi.StringOutput                               `pulumi:"meshBridgeEnable"`
	Name                           pulumi.StringOutput                               `pulumi:"name"`
	OverrideAllowaccess            pulumi.StringOutput                               `pulumi:"overrideAllowaccess"`
	OverrideIpFragment             pulumi.StringOutput                               `pulumi:"overrideIpFragment"`
	OverrideLan                    pulumi.StringOutput                               `pulumi:"overrideLan"`
	OverrideLedState               pulumi.StringOutput                               `pulumi:"overrideLedState"`
	OverrideLoginPasswdChange      pulumi.StringOutput                               `pulumi:"overrideLoginPasswdChange"`
	OverrideSplitTunnel            pulumi.StringOutput                               `pulumi:"overrideSplitTunnel"`
	OverrideWanPortMode            pulumi.StringOutput                               `pulumi:"overrideWanPortMode"`
	Radio1                         WirelessControllerWtpRadio1Output                 `pulumi:"radio1"`
	Radio2                         WirelessControllerWtpRadio2Output                 `pulumi:"radio2"`
	Radio3                         WirelessControllerWtpRadio3Output                 `pulumi:"radio3"`
	Radio4                         WirelessControllerWtpRadio4Output                 `pulumi:"radio4"`
	Region                         pulumi.StringOutput                               `pulumi:"region"`
	RegionX                        pulumi.StringOutput                               `pulumi:"regionX"`
	RegionY                        pulumi.StringOutput                               `pulumi:"regionY"`
	SplitTunnelingAclLocalApSubnet pulumi.StringOutput                               `pulumi:"splitTunnelingAclLocalApSubnet"`
	SplitTunnelingAclPath          pulumi.StringOutput                               `pulumi:"splitTunnelingAclPath"`
	SplitTunnelingAcls             WirelessControllerWtpSplitTunnelingAclArrayOutput `pulumi:"splitTunnelingAcls"`
	TunMtuDownlink                 pulumi.IntOutput                                  `pulumi:"tunMtuDownlink"`
	TunMtuUplink                   pulumi.IntOutput                                  `pulumi:"tunMtuUplink"`
	Uuid                           pulumi.StringOutput                               `pulumi:"uuid"`
	Vdomparam                      pulumi.StringPtrOutput                            `pulumi:"vdomparam"`
	WanPortMode                    pulumi.StringOutput                               `pulumi:"wanPortMode"`
	WtpId                          pulumi.StringOutput                               `pulumi:"wtpId"`
	WtpMode                        pulumi.StringOutput                               `pulumi:"wtpMode"`
	WtpProfile                     pulumi.StringOutput                               `pulumi:"wtpProfile"`
}

// NewWirelessControllerWtp registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerWtp(ctx *pulumi.Context,
	name string, args *WirelessControllerWtpArgs, opts ...pulumi.ResourceOption) (*WirelessControllerWtp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.WtpProfile == nil {
		return nil, errors.New("invalid value for required argument 'WtpProfile'")
	}
	if args.LoginPasswd != nil {
		args.LoginPasswd = pulumi.ToSecret(args.LoginPasswd).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"loginPasswd",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerWtp
	err := ctx.RegisterResource("fortios:index/wirelessControllerWtp:WirelessControllerWtp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerWtp gets an existing WirelessControllerWtp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerWtp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerWtpState, opts ...pulumi.ResourceOption) (*WirelessControllerWtp, error) {
	var resource WirelessControllerWtp
	err := ctx.ReadResource("fortios:index/wirelessControllerWtp:WirelessControllerWtp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerWtp resources.
type wirelessControllerWtpState struct {
	Admin                          *string                                  `pulumi:"admin"`
	Allowaccess                    *string                                  `pulumi:"allowaccess"`
	ApcfgProfile                   *string                                  `pulumi:"apcfgProfile"`
	BonjourProfile                 *string                                  `pulumi:"bonjourProfile"`
	CoordinateLatitude             *string                                  `pulumi:"coordinateLatitude"`
	CoordinateLongitude            *string                                  `pulumi:"coordinateLongitude"`
	DynamicSortSubtable            *string                                  `pulumi:"dynamicSortSubtable"`
	FirmwareProvision              *string                                  `pulumi:"firmwareProvision"`
	FirmwareProvisionLatest        *string                                  `pulumi:"firmwareProvisionLatest"`
	ImageDownload                  *string                                  `pulumi:"imageDownload"`
	Index                          *int                                     `pulumi:"index"`
	IpFragmentPreventing           *string                                  `pulumi:"ipFragmentPreventing"`
	Lan                            *WirelessControllerWtpLan                `pulumi:"lan"`
	LedState                       *string                                  `pulumi:"ledState"`
	Location                       *string                                  `pulumi:"location"`
	LoginPasswd                    *string                                  `pulumi:"loginPasswd"`
	LoginPasswdChange              *string                                  `pulumi:"loginPasswdChange"`
	MeshBridgeEnable               *string                                  `pulumi:"meshBridgeEnable"`
	Name                           *string                                  `pulumi:"name"`
	OverrideAllowaccess            *string                                  `pulumi:"overrideAllowaccess"`
	OverrideIpFragment             *string                                  `pulumi:"overrideIpFragment"`
	OverrideLan                    *string                                  `pulumi:"overrideLan"`
	OverrideLedState               *string                                  `pulumi:"overrideLedState"`
	OverrideLoginPasswdChange      *string                                  `pulumi:"overrideLoginPasswdChange"`
	OverrideSplitTunnel            *string                                  `pulumi:"overrideSplitTunnel"`
	OverrideWanPortMode            *string                                  `pulumi:"overrideWanPortMode"`
	Radio1                         *WirelessControllerWtpRadio1             `pulumi:"radio1"`
	Radio2                         *WirelessControllerWtpRadio2             `pulumi:"radio2"`
	Radio3                         *WirelessControllerWtpRadio3             `pulumi:"radio3"`
	Radio4                         *WirelessControllerWtpRadio4             `pulumi:"radio4"`
	Region                         *string                                  `pulumi:"region"`
	RegionX                        *string                                  `pulumi:"regionX"`
	RegionY                        *string                                  `pulumi:"regionY"`
	SplitTunnelingAclLocalApSubnet *string                                  `pulumi:"splitTunnelingAclLocalApSubnet"`
	SplitTunnelingAclPath          *string                                  `pulumi:"splitTunnelingAclPath"`
	SplitTunnelingAcls             []WirelessControllerWtpSplitTunnelingAcl `pulumi:"splitTunnelingAcls"`
	TunMtuDownlink                 *int                                     `pulumi:"tunMtuDownlink"`
	TunMtuUplink                   *int                                     `pulumi:"tunMtuUplink"`
	Uuid                           *string                                  `pulumi:"uuid"`
	Vdomparam                      *string                                  `pulumi:"vdomparam"`
	WanPortMode                    *string                                  `pulumi:"wanPortMode"`
	WtpId                          *string                                  `pulumi:"wtpId"`
	WtpMode                        *string                                  `pulumi:"wtpMode"`
	WtpProfile                     *string                                  `pulumi:"wtpProfile"`
}

type WirelessControllerWtpState struct {
	Admin                          pulumi.StringPtrInput
	Allowaccess                    pulumi.StringPtrInput
	ApcfgProfile                   pulumi.StringPtrInput
	BonjourProfile                 pulumi.StringPtrInput
	CoordinateLatitude             pulumi.StringPtrInput
	CoordinateLongitude            pulumi.StringPtrInput
	DynamicSortSubtable            pulumi.StringPtrInput
	FirmwareProvision              pulumi.StringPtrInput
	FirmwareProvisionLatest        pulumi.StringPtrInput
	ImageDownload                  pulumi.StringPtrInput
	Index                          pulumi.IntPtrInput
	IpFragmentPreventing           pulumi.StringPtrInput
	Lan                            WirelessControllerWtpLanPtrInput
	LedState                       pulumi.StringPtrInput
	Location                       pulumi.StringPtrInput
	LoginPasswd                    pulumi.StringPtrInput
	LoginPasswdChange              pulumi.StringPtrInput
	MeshBridgeEnable               pulumi.StringPtrInput
	Name                           pulumi.StringPtrInput
	OverrideAllowaccess            pulumi.StringPtrInput
	OverrideIpFragment             pulumi.StringPtrInput
	OverrideLan                    pulumi.StringPtrInput
	OverrideLedState               pulumi.StringPtrInput
	OverrideLoginPasswdChange      pulumi.StringPtrInput
	OverrideSplitTunnel            pulumi.StringPtrInput
	OverrideWanPortMode            pulumi.StringPtrInput
	Radio1                         WirelessControllerWtpRadio1PtrInput
	Radio2                         WirelessControllerWtpRadio2PtrInput
	Radio3                         WirelessControllerWtpRadio3PtrInput
	Radio4                         WirelessControllerWtpRadio4PtrInput
	Region                         pulumi.StringPtrInput
	RegionX                        pulumi.StringPtrInput
	RegionY                        pulumi.StringPtrInput
	SplitTunnelingAclLocalApSubnet pulumi.StringPtrInput
	SplitTunnelingAclPath          pulumi.StringPtrInput
	SplitTunnelingAcls             WirelessControllerWtpSplitTunnelingAclArrayInput
	TunMtuDownlink                 pulumi.IntPtrInput
	TunMtuUplink                   pulumi.IntPtrInput
	Uuid                           pulumi.StringPtrInput
	Vdomparam                      pulumi.StringPtrInput
	WanPortMode                    pulumi.StringPtrInput
	WtpId                          pulumi.StringPtrInput
	WtpMode                        pulumi.StringPtrInput
	WtpProfile                     pulumi.StringPtrInput
}

func (WirelessControllerWtpState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerWtpState)(nil)).Elem()
}

type wirelessControllerWtpArgs struct {
	Admin                          *string                                  `pulumi:"admin"`
	Allowaccess                    *string                                  `pulumi:"allowaccess"`
	ApcfgProfile                   *string                                  `pulumi:"apcfgProfile"`
	BonjourProfile                 *string                                  `pulumi:"bonjourProfile"`
	CoordinateLatitude             *string                                  `pulumi:"coordinateLatitude"`
	CoordinateLongitude            *string                                  `pulumi:"coordinateLongitude"`
	DynamicSortSubtable            *string                                  `pulumi:"dynamicSortSubtable"`
	FirmwareProvision              *string                                  `pulumi:"firmwareProvision"`
	FirmwareProvisionLatest        *string                                  `pulumi:"firmwareProvisionLatest"`
	ImageDownload                  *string                                  `pulumi:"imageDownload"`
	Index                          *int                                     `pulumi:"index"`
	IpFragmentPreventing           *string                                  `pulumi:"ipFragmentPreventing"`
	Lan                            *WirelessControllerWtpLan                `pulumi:"lan"`
	LedState                       *string                                  `pulumi:"ledState"`
	Location                       *string                                  `pulumi:"location"`
	LoginPasswd                    *string                                  `pulumi:"loginPasswd"`
	LoginPasswdChange              *string                                  `pulumi:"loginPasswdChange"`
	MeshBridgeEnable               *string                                  `pulumi:"meshBridgeEnable"`
	Name                           *string                                  `pulumi:"name"`
	OverrideAllowaccess            *string                                  `pulumi:"overrideAllowaccess"`
	OverrideIpFragment             *string                                  `pulumi:"overrideIpFragment"`
	OverrideLan                    *string                                  `pulumi:"overrideLan"`
	OverrideLedState               *string                                  `pulumi:"overrideLedState"`
	OverrideLoginPasswdChange      *string                                  `pulumi:"overrideLoginPasswdChange"`
	OverrideSplitTunnel            *string                                  `pulumi:"overrideSplitTunnel"`
	OverrideWanPortMode            *string                                  `pulumi:"overrideWanPortMode"`
	Radio1                         *WirelessControllerWtpRadio1             `pulumi:"radio1"`
	Radio2                         *WirelessControllerWtpRadio2             `pulumi:"radio2"`
	Radio3                         *WirelessControllerWtpRadio3             `pulumi:"radio3"`
	Radio4                         *WirelessControllerWtpRadio4             `pulumi:"radio4"`
	Region                         *string                                  `pulumi:"region"`
	RegionX                        *string                                  `pulumi:"regionX"`
	RegionY                        *string                                  `pulumi:"regionY"`
	SplitTunnelingAclLocalApSubnet *string                                  `pulumi:"splitTunnelingAclLocalApSubnet"`
	SplitTunnelingAclPath          *string                                  `pulumi:"splitTunnelingAclPath"`
	SplitTunnelingAcls             []WirelessControllerWtpSplitTunnelingAcl `pulumi:"splitTunnelingAcls"`
	TunMtuDownlink                 *int                                     `pulumi:"tunMtuDownlink"`
	TunMtuUplink                   *int                                     `pulumi:"tunMtuUplink"`
	Uuid                           *string                                  `pulumi:"uuid"`
	Vdomparam                      *string                                  `pulumi:"vdomparam"`
	WanPortMode                    *string                                  `pulumi:"wanPortMode"`
	WtpId                          *string                                  `pulumi:"wtpId"`
	WtpMode                        *string                                  `pulumi:"wtpMode"`
	WtpProfile                     string                                   `pulumi:"wtpProfile"`
}

// The set of arguments for constructing a WirelessControllerWtp resource.
type WirelessControllerWtpArgs struct {
	Admin                          pulumi.StringPtrInput
	Allowaccess                    pulumi.StringPtrInput
	ApcfgProfile                   pulumi.StringPtrInput
	BonjourProfile                 pulumi.StringPtrInput
	CoordinateLatitude             pulumi.StringPtrInput
	CoordinateLongitude            pulumi.StringPtrInput
	DynamicSortSubtable            pulumi.StringPtrInput
	FirmwareProvision              pulumi.StringPtrInput
	FirmwareProvisionLatest        pulumi.StringPtrInput
	ImageDownload                  pulumi.StringPtrInput
	Index                          pulumi.IntPtrInput
	IpFragmentPreventing           pulumi.StringPtrInput
	Lan                            WirelessControllerWtpLanPtrInput
	LedState                       pulumi.StringPtrInput
	Location                       pulumi.StringPtrInput
	LoginPasswd                    pulumi.StringPtrInput
	LoginPasswdChange              pulumi.StringPtrInput
	MeshBridgeEnable               pulumi.StringPtrInput
	Name                           pulumi.StringPtrInput
	OverrideAllowaccess            pulumi.StringPtrInput
	OverrideIpFragment             pulumi.StringPtrInput
	OverrideLan                    pulumi.StringPtrInput
	OverrideLedState               pulumi.StringPtrInput
	OverrideLoginPasswdChange      pulumi.StringPtrInput
	OverrideSplitTunnel            pulumi.StringPtrInput
	OverrideWanPortMode            pulumi.StringPtrInput
	Radio1                         WirelessControllerWtpRadio1PtrInput
	Radio2                         WirelessControllerWtpRadio2PtrInput
	Radio3                         WirelessControllerWtpRadio3PtrInput
	Radio4                         WirelessControllerWtpRadio4PtrInput
	Region                         pulumi.StringPtrInput
	RegionX                        pulumi.StringPtrInput
	RegionY                        pulumi.StringPtrInput
	SplitTunnelingAclLocalApSubnet pulumi.StringPtrInput
	SplitTunnelingAclPath          pulumi.StringPtrInput
	SplitTunnelingAcls             WirelessControllerWtpSplitTunnelingAclArrayInput
	TunMtuDownlink                 pulumi.IntPtrInput
	TunMtuUplink                   pulumi.IntPtrInput
	Uuid                           pulumi.StringPtrInput
	Vdomparam                      pulumi.StringPtrInput
	WanPortMode                    pulumi.StringPtrInput
	WtpId                          pulumi.StringPtrInput
	WtpMode                        pulumi.StringPtrInput
	WtpProfile                     pulumi.StringInput
}

func (WirelessControllerWtpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerWtpArgs)(nil)).Elem()
}

type WirelessControllerWtpInput interface {
	pulumi.Input

	ToWirelessControllerWtpOutput() WirelessControllerWtpOutput
	ToWirelessControllerWtpOutputWithContext(ctx context.Context) WirelessControllerWtpOutput
}

func (*WirelessControllerWtp) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerWtp)(nil)).Elem()
}

func (i *WirelessControllerWtp) ToWirelessControllerWtpOutput() WirelessControllerWtpOutput {
	return i.ToWirelessControllerWtpOutputWithContext(context.Background())
}

func (i *WirelessControllerWtp) ToWirelessControllerWtpOutputWithContext(ctx context.Context) WirelessControllerWtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerWtpOutput)
}

// WirelessControllerWtpArrayInput is an input type that accepts WirelessControllerWtpArray and WirelessControllerWtpArrayOutput values.
// You can construct a concrete instance of `WirelessControllerWtpArrayInput` via:
//
//	WirelessControllerWtpArray{ WirelessControllerWtpArgs{...} }
type WirelessControllerWtpArrayInput interface {
	pulumi.Input

	ToWirelessControllerWtpArrayOutput() WirelessControllerWtpArrayOutput
	ToWirelessControllerWtpArrayOutputWithContext(context.Context) WirelessControllerWtpArrayOutput
}

type WirelessControllerWtpArray []WirelessControllerWtpInput

func (WirelessControllerWtpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerWtp)(nil)).Elem()
}

func (i WirelessControllerWtpArray) ToWirelessControllerWtpArrayOutput() WirelessControllerWtpArrayOutput {
	return i.ToWirelessControllerWtpArrayOutputWithContext(context.Background())
}

func (i WirelessControllerWtpArray) ToWirelessControllerWtpArrayOutputWithContext(ctx context.Context) WirelessControllerWtpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerWtpArrayOutput)
}

// WirelessControllerWtpMapInput is an input type that accepts WirelessControllerWtpMap and WirelessControllerWtpMapOutput values.
// You can construct a concrete instance of `WirelessControllerWtpMapInput` via:
//
//	WirelessControllerWtpMap{ "key": WirelessControllerWtpArgs{...} }
type WirelessControllerWtpMapInput interface {
	pulumi.Input

	ToWirelessControllerWtpMapOutput() WirelessControllerWtpMapOutput
	ToWirelessControllerWtpMapOutputWithContext(context.Context) WirelessControllerWtpMapOutput
}

type WirelessControllerWtpMap map[string]WirelessControllerWtpInput

func (WirelessControllerWtpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerWtp)(nil)).Elem()
}

func (i WirelessControllerWtpMap) ToWirelessControllerWtpMapOutput() WirelessControllerWtpMapOutput {
	return i.ToWirelessControllerWtpMapOutputWithContext(context.Background())
}

func (i WirelessControllerWtpMap) ToWirelessControllerWtpMapOutputWithContext(ctx context.Context) WirelessControllerWtpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerWtpMapOutput)
}

type WirelessControllerWtpOutput struct{ *pulumi.OutputState }

func (WirelessControllerWtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerWtp)(nil)).Elem()
}

func (o WirelessControllerWtpOutput) ToWirelessControllerWtpOutput() WirelessControllerWtpOutput {
	return o
}

func (o WirelessControllerWtpOutput) ToWirelessControllerWtpOutputWithContext(ctx context.Context) WirelessControllerWtpOutput {
	return o
}

func (o WirelessControllerWtpOutput) Admin() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.Admin }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) Allowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.Allowaccess }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) ApcfgProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.ApcfgProfile }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) BonjourProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.BonjourProfile }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) CoordinateLatitude() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.CoordinateLatitude }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) CoordinateLongitude() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.CoordinateLongitude }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerWtpOutput) FirmwareProvision() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.FirmwareProvision }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) FirmwareProvisionLatest() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.FirmwareProvisionLatest }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) ImageDownload() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.ImageDownload }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) Index() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.IntOutput { return v.Index }).(pulumi.IntOutput)
}

func (o WirelessControllerWtpOutput) IpFragmentPreventing() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.IpFragmentPreventing }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) Lan() WirelessControllerWtpLanOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) WirelessControllerWtpLanOutput { return v.Lan }).(WirelessControllerWtpLanOutput)
}

func (o WirelessControllerWtpOutput) LedState() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.LedState }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) LoginPasswd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringPtrOutput { return v.LoginPasswd }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerWtpOutput) LoginPasswdChange() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.LoginPasswdChange }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) MeshBridgeEnable() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.MeshBridgeEnable }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) OverrideAllowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.OverrideAllowaccess }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) OverrideIpFragment() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.OverrideIpFragment }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) OverrideLan() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.OverrideLan }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) OverrideLedState() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.OverrideLedState }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) OverrideLoginPasswdChange() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.OverrideLoginPasswdChange }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) OverrideSplitTunnel() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.OverrideSplitTunnel }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) OverrideWanPortMode() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.OverrideWanPortMode }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) Radio1() WirelessControllerWtpRadio1Output {
	return o.ApplyT(func(v *WirelessControllerWtp) WirelessControllerWtpRadio1Output { return v.Radio1 }).(WirelessControllerWtpRadio1Output)
}

func (o WirelessControllerWtpOutput) Radio2() WirelessControllerWtpRadio2Output {
	return o.ApplyT(func(v *WirelessControllerWtp) WirelessControllerWtpRadio2Output { return v.Radio2 }).(WirelessControllerWtpRadio2Output)
}

func (o WirelessControllerWtpOutput) Radio3() WirelessControllerWtpRadio3Output {
	return o.ApplyT(func(v *WirelessControllerWtp) WirelessControllerWtpRadio3Output { return v.Radio3 }).(WirelessControllerWtpRadio3Output)
}

func (o WirelessControllerWtpOutput) Radio4() WirelessControllerWtpRadio4Output {
	return o.ApplyT(func(v *WirelessControllerWtp) WirelessControllerWtpRadio4Output { return v.Radio4 }).(WirelessControllerWtpRadio4Output)
}

func (o WirelessControllerWtpOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) RegionX() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.RegionX }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) RegionY() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.RegionY }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) SplitTunnelingAclLocalApSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.SplitTunnelingAclLocalApSubnet }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) SplitTunnelingAclPath() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.SplitTunnelingAclPath }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) SplitTunnelingAcls() WirelessControllerWtpSplitTunnelingAclArrayOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) WirelessControllerWtpSplitTunnelingAclArrayOutput {
		return v.SplitTunnelingAcls
	}).(WirelessControllerWtpSplitTunnelingAclArrayOutput)
}

func (o WirelessControllerWtpOutput) TunMtuDownlink() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.IntOutput { return v.TunMtuDownlink }).(pulumi.IntOutput)
}

func (o WirelessControllerWtpOutput) TunMtuUplink() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.IntOutput { return v.TunMtuUplink }).(pulumi.IntOutput)
}

func (o WirelessControllerWtpOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerWtpOutput) WanPortMode() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.WanPortMode }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) WtpId() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.WtpId }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) WtpMode() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.WtpMode }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpOutput) WtpProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtp) pulumi.StringOutput { return v.WtpProfile }).(pulumi.StringOutput)
}

type WirelessControllerWtpArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerWtpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerWtp)(nil)).Elem()
}

func (o WirelessControllerWtpArrayOutput) ToWirelessControllerWtpArrayOutput() WirelessControllerWtpArrayOutput {
	return o
}

func (o WirelessControllerWtpArrayOutput) ToWirelessControllerWtpArrayOutputWithContext(ctx context.Context) WirelessControllerWtpArrayOutput {
	return o
}

func (o WirelessControllerWtpArrayOutput) Index(i pulumi.IntInput) WirelessControllerWtpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerWtp {
		return vs[0].([]*WirelessControllerWtp)[vs[1].(int)]
	}).(WirelessControllerWtpOutput)
}

type WirelessControllerWtpMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerWtpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerWtp)(nil)).Elem()
}

func (o WirelessControllerWtpMapOutput) ToWirelessControllerWtpMapOutput() WirelessControllerWtpMapOutput {
	return o
}

func (o WirelessControllerWtpMapOutput) ToWirelessControllerWtpMapOutputWithContext(ctx context.Context) WirelessControllerWtpMapOutput {
	return o
}

func (o WirelessControllerWtpMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerWtpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerWtp {
		return vs[0].(map[string]*WirelessControllerWtp)[vs[1].(string)]
	}).(WirelessControllerWtpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerWtpInput)(nil)).Elem(), &WirelessControllerWtp{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerWtpArrayInput)(nil)).Elem(), WirelessControllerWtpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerWtpMapInput)(nil)).Elem(), WirelessControllerWtpMap{})
	pulumi.RegisterOutputType(WirelessControllerWtpOutput{})
	pulumi.RegisterOutputType(WirelessControllerWtpArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerWtpMapOutput{})
}
