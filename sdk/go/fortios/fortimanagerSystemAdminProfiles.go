// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type FortimanagerSystemAdminProfiles struct {
	pulumi.CustomResourceState

	AdomPolicyPackages                  pulumi.StringPtrOutput `pulumi:"adomPolicyPackages"`
	AdomSwitch                          pulumi.StringPtrOutput `pulumi:"adomSwitch"`
	Assignment                          pulumi.StringPtrOutput `pulumi:"assignment"`
	ConfigRetrieve                      pulumi.StringPtrOutput `pulumi:"configRetrieve"`
	ConfigRevert                        pulumi.StringPtrOutput `pulumi:"configRevert"`
	ConsistencyCheck                    pulumi.StringPtrOutput `pulumi:"consistencyCheck"`
	DeployManagement                    pulumi.StringPtrOutput `pulumi:"deployManagement"`
	Description                         pulumi.StringPtrOutput `pulumi:"description"`
	DeviceAp                            pulumi.StringPtrOutput `pulumi:"deviceAp"`
	DeviceConfig                        pulumi.StringPtrOutput `pulumi:"deviceConfig"`
	DeviceForticlient                   pulumi.StringPtrOutput `pulumi:"deviceForticlient"`
	DeviceFortiswitch                   pulumi.StringPtrOutput `pulumi:"deviceFortiswitch"`
	DeviceManager                       pulumi.StringPtrOutput `pulumi:"deviceManager"`
	DeviceOperation                     pulumi.StringPtrOutput `pulumi:"deviceOperation"`
	DeviceProfile                       pulumi.StringPtrOutput `pulumi:"deviceProfile"`
	DeviceRevisionDeletion              pulumi.StringPtrOutput `pulumi:"deviceRevisionDeletion"`
	DeviceWanLinkLoadBalance            pulumi.StringPtrOutput `pulumi:"deviceWanLinkLoadBalance"`
	FortiguardCenter                    pulumi.StringPtrOutput `pulumi:"fortiguardCenter"`
	FortiguardCenterAdvanced            pulumi.StringPtrOutput `pulumi:"fortiguardCenterAdvanced"`
	FortiguardCenterFirmwareManagerment pulumi.StringPtrOutput `pulumi:"fortiguardCenterFirmwareManagerment"`
	FortiguardCenterLicensing           pulumi.StringPtrOutput `pulumi:"fortiguardCenterLicensing"`
	GlobalPolicyPackages                pulumi.StringPtrOutput `pulumi:"globalPolicyPackages"`
	ImportPolicyPackages                pulumi.StringPtrOutput `pulumi:"importPolicyPackages"`
	IntfMapping                         pulumi.StringPtrOutput `pulumi:"intfMapping"`
	LogViewer                           pulumi.StringPtrOutput `pulumi:"logViewer"`
	PolicyObjects                       pulumi.StringPtrOutput `pulumi:"policyObjects"`
	Profileid                           pulumi.StringOutput    `pulumi:"profileid"`
	SetInstallTargets                   pulumi.StringPtrOutput `pulumi:"setInstallTargets"`
	SystemSetting                       pulumi.StringPtrOutput `pulumi:"systemSetting"`
	TerminalAccess                      pulumi.StringPtrOutput `pulumi:"terminalAccess"`
	VpnManager                          pulumi.StringPtrOutput `pulumi:"vpnManager"`
}

// NewFortimanagerSystemAdminProfiles registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemAdminProfiles(ctx *pulumi.Context,
	name string, args *FortimanagerSystemAdminProfilesArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemAdminProfiles, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Profileid == nil {
		return nil, errors.New("invalid value for required argument 'Profileid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemAdminProfiles
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemAdminProfiles:FortimanagerSystemAdminProfiles", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemAdminProfiles gets an existing FortimanagerSystemAdminProfiles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemAdminProfiles(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemAdminProfilesState, opts ...pulumi.ResourceOption) (*FortimanagerSystemAdminProfiles, error) {
	var resource FortimanagerSystemAdminProfiles
	err := ctx.ReadResource("fortios:index/fortimanagerSystemAdminProfiles:FortimanagerSystemAdminProfiles", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemAdminProfiles resources.
type fortimanagerSystemAdminProfilesState struct {
	AdomPolicyPackages                  *string `pulumi:"adomPolicyPackages"`
	AdomSwitch                          *string `pulumi:"adomSwitch"`
	Assignment                          *string `pulumi:"assignment"`
	ConfigRetrieve                      *string `pulumi:"configRetrieve"`
	ConfigRevert                        *string `pulumi:"configRevert"`
	ConsistencyCheck                    *string `pulumi:"consistencyCheck"`
	DeployManagement                    *string `pulumi:"deployManagement"`
	Description                         *string `pulumi:"description"`
	DeviceAp                            *string `pulumi:"deviceAp"`
	DeviceConfig                        *string `pulumi:"deviceConfig"`
	DeviceForticlient                   *string `pulumi:"deviceForticlient"`
	DeviceFortiswitch                   *string `pulumi:"deviceFortiswitch"`
	DeviceManager                       *string `pulumi:"deviceManager"`
	DeviceOperation                     *string `pulumi:"deviceOperation"`
	DeviceProfile                       *string `pulumi:"deviceProfile"`
	DeviceRevisionDeletion              *string `pulumi:"deviceRevisionDeletion"`
	DeviceWanLinkLoadBalance            *string `pulumi:"deviceWanLinkLoadBalance"`
	FortiguardCenter                    *string `pulumi:"fortiguardCenter"`
	FortiguardCenterAdvanced            *string `pulumi:"fortiguardCenterAdvanced"`
	FortiguardCenterFirmwareManagerment *string `pulumi:"fortiguardCenterFirmwareManagerment"`
	FortiguardCenterLicensing           *string `pulumi:"fortiguardCenterLicensing"`
	GlobalPolicyPackages                *string `pulumi:"globalPolicyPackages"`
	ImportPolicyPackages                *string `pulumi:"importPolicyPackages"`
	IntfMapping                         *string `pulumi:"intfMapping"`
	LogViewer                           *string `pulumi:"logViewer"`
	PolicyObjects                       *string `pulumi:"policyObjects"`
	Profileid                           *string `pulumi:"profileid"`
	SetInstallTargets                   *string `pulumi:"setInstallTargets"`
	SystemSetting                       *string `pulumi:"systemSetting"`
	TerminalAccess                      *string `pulumi:"terminalAccess"`
	VpnManager                          *string `pulumi:"vpnManager"`
}

type FortimanagerSystemAdminProfilesState struct {
	AdomPolicyPackages                  pulumi.StringPtrInput
	AdomSwitch                          pulumi.StringPtrInput
	Assignment                          pulumi.StringPtrInput
	ConfigRetrieve                      pulumi.StringPtrInput
	ConfigRevert                        pulumi.StringPtrInput
	ConsistencyCheck                    pulumi.StringPtrInput
	DeployManagement                    pulumi.StringPtrInput
	Description                         pulumi.StringPtrInput
	DeviceAp                            pulumi.StringPtrInput
	DeviceConfig                        pulumi.StringPtrInput
	DeviceForticlient                   pulumi.StringPtrInput
	DeviceFortiswitch                   pulumi.StringPtrInput
	DeviceManager                       pulumi.StringPtrInput
	DeviceOperation                     pulumi.StringPtrInput
	DeviceProfile                       pulumi.StringPtrInput
	DeviceRevisionDeletion              pulumi.StringPtrInput
	DeviceWanLinkLoadBalance            pulumi.StringPtrInput
	FortiguardCenter                    pulumi.StringPtrInput
	FortiguardCenterAdvanced            pulumi.StringPtrInput
	FortiguardCenterFirmwareManagerment pulumi.StringPtrInput
	FortiguardCenterLicensing           pulumi.StringPtrInput
	GlobalPolicyPackages                pulumi.StringPtrInput
	ImportPolicyPackages                pulumi.StringPtrInput
	IntfMapping                         pulumi.StringPtrInput
	LogViewer                           pulumi.StringPtrInput
	PolicyObjects                       pulumi.StringPtrInput
	Profileid                           pulumi.StringPtrInput
	SetInstallTargets                   pulumi.StringPtrInput
	SystemSetting                       pulumi.StringPtrInput
	TerminalAccess                      pulumi.StringPtrInput
	VpnManager                          pulumi.StringPtrInput
}

func (FortimanagerSystemAdminProfilesState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemAdminProfilesState)(nil)).Elem()
}

type fortimanagerSystemAdminProfilesArgs struct {
	AdomPolicyPackages                  *string `pulumi:"adomPolicyPackages"`
	AdomSwitch                          *string `pulumi:"adomSwitch"`
	Assignment                          *string `pulumi:"assignment"`
	ConfigRetrieve                      *string `pulumi:"configRetrieve"`
	ConfigRevert                        *string `pulumi:"configRevert"`
	ConsistencyCheck                    *string `pulumi:"consistencyCheck"`
	DeployManagement                    *string `pulumi:"deployManagement"`
	Description                         *string `pulumi:"description"`
	DeviceAp                            *string `pulumi:"deviceAp"`
	DeviceConfig                        *string `pulumi:"deviceConfig"`
	DeviceForticlient                   *string `pulumi:"deviceForticlient"`
	DeviceFortiswitch                   *string `pulumi:"deviceFortiswitch"`
	DeviceManager                       *string `pulumi:"deviceManager"`
	DeviceOperation                     *string `pulumi:"deviceOperation"`
	DeviceProfile                       *string `pulumi:"deviceProfile"`
	DeviceRevisionDeletion              *string `pulumi:"deviceRevisionDeletion"`
	DeviceWanLinkLoadBalance            *string `pulumi:"deviceWanLinkLoadBalance"`
	FortiguardCenter                    *string `pulumi:"fortiguardCenter"`
	FortiguardCenterAdvanced            *string `pulumi:"fortiguardCenterAdvanced"`
	FortiguardCenterFirmwareManagerment *string `pulumi:"fortiguardCenterFirmwareManagerment"`
	FortiguardCenterLicensing           *string `pulumi:"fortiguardCenterLicensing"`
	GlobalPolicyPackages                *string `pulumi:"globalPolicyPackages"`
	ImportPolicyPackages                *string `pulumi:"importPolicyPackages"`
	IntfMapping                         *string `pulumi:"intfMapping"`
	LogViewer                           *string `pulumi:"logViewer"`
	PolicyObjects                       *string `pulumi:"policyObjects"`
	Profileid                           string  `pulumi:"profileid"`
	SetInstallTargets                   *string `pulumi:"setInstallTargets"`
	SystemSetting                       *string `pulumi:"systemSetting"`
	TerminalAccess                      *string `pulumi:"terminalAccess"`
	VpnManager                          *string `pulumi:"vpnManager"`
}

// The set of arguments for constructing a FortimanagerSystemAdminProfiles resource.
type FortimanagerSystemAdminProfilesArgs struct {
	AdomPolicyPackages                  pulumi.StringPtrInput
	AdomSwitch                          pulumi.StringPtrInput
	Assignment                          pulumi.StringPtrInput
	ConfigRetrieve                      pulumi.StringPtrInput
	ConfigRevert                        pulumi.StringPtrInput
	ConsistencyCheck                    pulumi.StringPtrInput
	DeployManagement                    pulumi.StringPtrInput
	Description                         pulumi.StringPtrInput
	DeviceAp                            pulumi.StringPtrInput
	DeviceConfig                        pulumi.StringPtrInput
	DeviceForticlient                   pulumi.StringPtrInput
	DeviceFortiswitch                   pulumi.StringPtrInput
	DeviceManager                       pulumi.StringPtrInput
	DeviceOperation                     pulumi.StringPtrInput
	DeviceProfile                       pulumi.StringPtrInput
	DeviceRevisionDeletion              pulumi.StringPtrInput
	DeviceWanLinkLoadBalance            pulumi.StringPtrInput
	FortiguardCenter                    pulumi.StringPtrInput
	FortiguardCenterAdvanced            pulumi.StringPtrInput
	FortiguardCenterFirmwareManagerment pulumi.StringPtrInput
	FortiguardCenterLicensing           pulumi.StringPtrInput
	GlobalPolicyPackages                pulumi.StringPtrInput
	ImportPolicyPackages                pulumi.StringPtrInput
	IntfMapping                         pulumi.StringPtrInput
	LogViewer                           pulumi.StringPtrInput
	PolicyObjects                       pulumi.StringPtrInput
	Profileid                           pulumi.StringInput
	SetInstallTargets                   pulumi.StringPtrInput
	SystemSetting                       pulumi.StringPtrInput
	TerminalAccess                      pulumi.StringPtrInput
	VpnManager                          pulumi.StringPtrInput
}

func (FortimanagerSystemAdminProfilesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemAdminProfilesArgs)(nil)).Elem()
}

type FortimanagerSystemAdminProfilesInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminProfilesOutput() FortimanagerSystemAdminProfilesOutput
	ToFortimanagerSystemAdminProfilesOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesOutput
}

func (*FortimanagerSystemAdminProfiles) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (i *FortimanagerSystemAdminProfiles) ToFortimanagerSystemAdminProfilesOutput() FortimanagerSystemAdminProfilesOutput {
	return i.ToFortimanagerSystemAdminProfilesOutputWithContext(context.Background())
}

func (i *FortimanagerSystemAdminProfiles) ToFortimanagerSystemAdminProfilesOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminProfilesOutput)
}

func (i *FortimanagerSystemAdminProfiles) ToOutput(ctx context.Context) pulumix.Output[*FortimanagerSystemAdminProfiles] {
	return pulumix.Output[*FortimanagerSystemAdminProfiles]{
		OutputState: i.ToFortimanagerSystemAdminProfilesOutputWithContext(ctx).OutputState,
	}
}

// FortimanagerSystemAdminProfilesArrayInput is an input type that accepts FortimanagerSystemAdminProfilesArray and FortimanagerSystemAdminProfilesArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemAdminProfilesArrayInput` via:
//
//	FortimanagerSystemAdminProfilesArray{ FortimanagerSystemAdminProfilesArgs{...} }
type FortimanagerSystemAdminProfilesArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminProfilesArrayOutput() FortimanagerSystemAdminProfilesArrayOutput
	ToFortimanagerSystemAdminProfilesArrayOutputWithContext(context.Context) FortimanagerSystemAdminProfilesArrayOutput
}

type FortimanagerSystemAdminProfilesArray []FortimanagerSystemAdminProfilesInput

func (FortimanagerSystemAdminProfilesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (i FortimanagerSystemAdminProfilesArray) ToFortimanagerSystemAdminProfilesArrayOutput() FortimanagerSystemAdminProfilesArrayOutput {
	return i.ToFortimanagerSystemAdminProfilesArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemAdminProfilesArray) ToFortimanagerSystemAdminProfilesArrayOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminProfilesArrayOutput)
}

func (i FortimanagerSystemAdminProfilesArray) ToOutput(ctx context.Context) pulumix.Output[[]*FortimanagerSystemAdminProfiles] {
	return pulumix.Output[[]*FortimanagerSystemAdminProfiles]{
		OutputState: i.ToFortimanagerSystemAdminProfilesArrayOutputWithContext(ctx).OutputState,
	}
}

// FortimanagerSystemAdminProfilesMapInput is an input type that accepts FortimanagerSystemAdminProfilesMap and FortimanagerSystemAdminProfilesMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemAdminProfilesMapInput` via:
//
//	FortimanagerSystemAdminProfilesMap{ "key": FortimanagerSystemAdminProfilesArgs{...} }
type FortimanagerSystemAdminProfilesMapInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminProfilesMapOutput() FortimanagerSystemAdminProfilesMapOutput
	ToFortimanagerSystemAdminProfilesMapOutputWithContext(context.Context) FortimanagerSystemAdminProfilesMapOutput
}

type FortimanagerSystemAdminProfilesMap map[string]FortimanagerSystemAdminProfilesInput

func (FortimanagerSystemAdminProfilesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (i FortimanagerSystemAdminProfilesMap) ToFortimanagerSystemAdminProfilesMapOutput() FortimanagerSystemAdminProfilesMapOutput {
	return i.ToFortimanagerSystemAdminProfilesMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemAdminProfilesMap) ToFortimanagerSystemAdminProfilesMapOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminProfilesMapOutput)
}

func (i FortimanagerSystemAdminProfilesMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FortimanagerSystemAdminProfiles] {
	return pulumix.Output[map[string]*FortimanagerSystemAdminProfiles]{
		OutputState: i.ToFortimanagerSystemAdminProfilesMapOutputWithContext(ctx).OutputState,
	}
}

type FortimanagerSystemAdminProfilesOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdminProfilesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (o FortimanagerSystemAdminProfilesOutput) ToFortimanagerSystemAdminProfilesOutput() FortimanagerSystemAdminProfilesOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesOutput) ToFortimanagerSystemAdminProfilesOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesOutput) ToOutput(ctx context.Context) pulumix.Output[*FortimanagerSystemAdminProfiles] {
	return pulumix.Output[*FortimanagerSystemAdminProfiles]{
		OutputState: o.OutputState,
	}
}

func (o FortimanagerSystemAdminProfilesOutput) AdomPolicyPackages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.AdomPolicyPackages }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) AdomSwitch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.AdomSwitch }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) Assignment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.Assignment }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) ConfigRetrieve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.ConfigRetrieve }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) ConfigRevert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.ConfigRevert }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) ConsistencyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.ConsistencyCheck }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeployManagement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeployManagement }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceAp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceAp }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceConfig }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceForticlient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceForticlient }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceFortiswitch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceFortiswitch }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceManager() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceManager }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceOperation }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceProfile }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceRevisionDeletion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceRevisionDeletion }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) DeviceWanLinkLoadBalance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.DeviceWanLinkLoadBalance }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) FortiguardCenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.FortiguardCenter }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) FortiguardCenterAdvanced() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.FortiguardCenterAdvanced }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) FortiguardCenterFirmwareManagerment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput {
		return v.FortiguardCenterFirmwareManagerment
	}).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) FortiguardCenterLicensing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.FortiguardCenterLicensing }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) GlobalPolicyPackages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.GlobalPolicyPackages }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) ImportPolicyPackages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.ImportPolicyPackages }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) IntfMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.IntfMapping }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) LogViewer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.LogViewer }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) PolicyObjects() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.PolicyObjects }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) Profileid() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringOutput { return v.Profileid }).(pulumi.StringOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) SetInstallTargets() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.SetInstallTargets }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) SystemSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.SystemSetting }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) TerminalAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.TerminalAccess }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemAdminProfilesOutput) VpnManager() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemAdminProfiles) pulumi.StringPtrOutput { return v.VpnManager }).(pulumi.StringPtrOutput)
}

type FortimanagerSystemAdminProfilesArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdminProfilesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (o FortimanagerSystemAdminProfilesArrayOutput) ToFortimanagerSystemAdminProfilesArrayOutput() FortimanagerSystemAdminProfilesArrayOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesArrayOutput) ToFortimanagerSystemAdminProfilesArrayOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesArrayOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FortimanagerSystemAdminProfiles] {
	return pulumix.Output[[]*FortimanagerSystemAdminProfiles]{
		OutputState: o.OutputState,
	}
}

func (o FortimanagerSystemAdminProfilesArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemAdminProfilesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemAdminProfiles {
		return vs[0].([]*FortimanagerSystemAdminProfiles)[vs[1].(int)]
	}).(FortimanagerSystemAdminProfilesOutput)
}

type FortimanagerSystemAdminProfilesMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdminProfilesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (o FortimanagerSystemAdminProfilesMapOutput) ToFortimanagerSystemAdminProfilesMapOutput() FortimanagerSystemAdminProfilesMapOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesMapOutput) ToFortimanagerSystemAdminProfilesMapOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesMapOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FortimanagerSystemAdminProfiles] {
	return pulumix.Output[map[string]*FortimanagerSystemAdminProfiles]{
		OutputState: o.OutputState,
	}
}

func (o FortimanagerSystemAdminProfilesMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemAdminProfilesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemAdminProfiles {
		return vs[0].(map[string]*FortimanagerSystemAdminProfiles)[vs[1].(string)]
	}).(FortimanagerSystemAdminProfilesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdminProfilesInput)(nil)).Elem(), &FortimanagerSystemAdminProfiles{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdminProfilesArrayInput)(nil)).Elem(), FortimanagerSystemAdminProfilesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdminProfilesMapInput)(nil)).Elem(), FortimanagerSystemAdminProfilesMap{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminProfilesOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminProfilesArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminProfilesMapOutput{})
}
