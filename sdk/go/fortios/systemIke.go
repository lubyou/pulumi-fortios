// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IKE global attributes. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// System Ike can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemIke:SystemIke labelname SystemIke
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemIke:SystemIke labelname SystemIke
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemIke struct {
	pulumi.CustomResourceState

	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 SystemIkeDhGroup1PtrOutput `pulumi:"dhGroup1"`
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 SystemIkeDhGroup14PtrOutput `pulumi:"dhGroup14"`
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 SystemIkeDhGroup15PtrOutput `pulumi:"dhGroup15"`
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 SystemIkeDhGroup16PtrOutput `pulumi:"dhGroup16"`
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 SystemIkeDhGroup17PtrOutput `pulumi:"dhGroup17"`
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 SystemIkeDhGroup18PtrOutput `pulumi:"dhGroup18"`
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 SystemIkeDhGroup19PtrOutput `pulumi:"dhGroup19"`
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 SystemIkeDhGroup2PtrOutput `pulumi:"dhGroup2"`
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 SystemIkeDhGroup20PtrOutput `pulumi:"dhGroup20"`
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 SystemIkeDhGroup21PtrOutput `pulumi:"dhGroup21"`
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 SystemIkeDhGroup27PtrOutput `pulumi:"dhGroup27"`
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 SystemIkeDhGroup28PtrOutput `pulumi:"dhGroup28"`
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 SystemIkeDhGroup29PtrOutput `pulumi:"dhGroup29"`
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 SystemIkeDhGroup30PtrOutput `pulumi:"dhGroup30"`
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 SystemIkeDhGroup31PtrOutput `pulumi:"dhGroup31"`
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 SystemIkeDhGroup32PtrOutput `pulumi:"dhGroup32"`
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 SystemIkeDhGroup5PtrOutput `pulumi:"dhGroup5"`
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache pulumi.StringOutput `pulumi:"dhKeypairCache"`
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount pulumi.IntOutput `pulumi:"dhKeypairCount"`
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle pulumi.StringOutput `pulumi:"dhKeypairThrottle"`
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode pulumi.StringOutput `pulumi:"dhMode"`
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess pulumi.StringOutput `pulumi:"dhMultiprocess"`
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount pulumi.IntOutput `pulumi:"dhWorkerCount"`
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit pulumi.IntOutput `pulumi:"embryonicLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemIke registers a new resource with the given unique name, arguments, and options.
func NewSystemIke(ctx *pulumi.Context,
	name string, args *SystemIkeArgs, opts ...pulumi.ResourceOption) (*SystemIke, error) {
	if args == nil {
		args = &SystemIkeArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemIke
	err := ctx.RegisterResource("fortios:index/systemIke:SystemIke", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemIke gets an existing SystemIke resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemIke(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemIkeState, opts ...pulumi.ResourceOption) (*SystemIke, error) {
	var resource SystemIke
	err := ctx.ReadResource("fortios:index/systemIke:SystemIke", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemIke resources.
type systemIkeState struct {
	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 *SystemIkeDhGroup1 `pulumi:"dhGroup1"`
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 *SystemIkeDhGroup14 `pulumi:"dhGroup14"`
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 *SystemIkeDhGroup15 `pulumi:"dhGroup15"`
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 *SystemIkeDhGroup16 `pulumi:"dhGroup16"`
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 *SystemIkeDhGroup17 `pulumi:"dhGroup17"`
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 *SystemIkeDhGroup18 `pulumi:"dhGroup18"`
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 *SystemIkeDhGroup19 `pulumi:"dhGroup19"`
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 *SystemIkeDhGroup2 `pulumi:"dhGroup2"`
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 *SystemIkeDhGroup20 `pulumi:"dhGroup20"`
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 *SystemIkeDhGroup21 `pulumi:"dhGroup21"`
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 *SystemIkeDhGroup27 `pulumi:"dhGroup27"`
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 *SystemIkeDhGroup28 `pulumi:"dhGroup28"`
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 *SystemIkeDhGroup29 `pulumi:"dhGroup29"`
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 *SystemIkeDhGroup30 `pulumi:"dhGroup30"`
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 *SystemIkeDhGroup31 `pulumi:"dhGroup31"`
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 *SystemIkeDhGroup32 `pulumi:"dhGroup32"`
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 *SystemIkeDhGroup5 `pulumi:"dhGroup5"`
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache *string `pulumi:"dhKeypairCache"`
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount *int `pulumi:"dhKeypairCount"`
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle *string `pulumi:"dhKeypairThrottle"`
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode *string `pulumi:"dhMode"`
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess *string `pulumi:"dhMultiprocess"`
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount *int `pulumi:"dhWorkerCount"`
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit *int `pulumi:"embryonicLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemIkeState struct {
	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 SystemIkeDhGroup1PtrInput
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 SystemIkeDhGroup14PtrInput
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 SystemIkeDhGroup15PtrInput
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 SystemIkeDhGroup16PtrInput
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 SystemIkeDhGroup17PtrInput
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 SystemIkeDhGroup18PtrInput
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 SystemIkeDhGroup19PtrInput
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 SystemIkeDhGroup2PtrInput
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 SystemIkeDhGroup20PtrInput
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 SystemIkeDhGroup21PtrInput
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 SystemIkeDhGroup27PtrInput
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 SystemIkeDhGroup28PtrInput
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 SystemIkeDhGroup29PtrInput
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 SystemIkeDhGroup30PtrInput
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 SystemIkeDhGroup31PtrInput
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 SystemIkeDhGroup32PtrInput
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 SystemIkeDhGroup5PtrInput
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache pulumi.StringPtrInput
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount pulumi.IntPtrInput
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle pulumi.StringPtrInput
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode pulumi.StringPtrInput
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess pulumi.StringPtrInput
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount pulumi.IntPtrInput
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIkeState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIkeState)(nil)).Elem()
}

type systemIkeArgs struct {
	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 *SystemIkeDhGroup1 `pulumi:"dhGroup1"`
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 *SystemIkeDhGroup14 `pulumi:"dhGroup14"`
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 *SystemIkeDhGroup15 `pulumi:"dhGroup15"`
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 *SystemIkeDhGroup16 `pulumi:"dhGroup16"`
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 *SystemIkeDhGroup17 `pulumi:"dhGroup17"`
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 *SystemIkeDhGroup18 `pulumi:"dhGroup18"`
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 *SystemIkeDhGroup19 `pulumi:"dhGroup19"`
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 *SystemIkeDhGroup2 `pulumi:"dhGroup2"`
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 *SystemIkeDhGroup20 `pulumi:"dhGroup20"`
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 *SystemIkeDhGroup21 `pulumi:"dhGroup21"`
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 *SystemIkeDhGroup27 `pulumi:"dhGroup27"`
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 *SystemIkeDhGroup28 `pulumi:"dhGroup28"`
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 *SystemIkeDhGroup29 `pulumi:"dhGroup29"`
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 *SystemIkeDhGroup30 `pulumi:"dhGroup30"`
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 *SystemIkeDhGroup31 `pulumi:"dhGroup31"`
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 *SystemIkeDhGroup32 `pulumi:"dhGroup32"`
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 *SystemIkeDhGroup5 `pulumi:"dhGroup5"`
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache *string `pulumi:"dhKeypairCache"`
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount *int `pulumi:"dhKeypairCount"`
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle *string `pulumi:"dhKeypairThrottle"`
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode *string `pulumi:"dhMode"`
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess *string `pulumi:"dhMultiprocess"`
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount *int `pulumi:"dhWorkerCount"`
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit *int `pulumi:"embryonicLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemIke resource.
type SystemIkeArgs struct {
	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 SystemIkeDhGroup1PtrInput
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 SystemIkeDhGroup14PtrInput
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 SystemIkeDhGroup15PtrInput
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 SystemIkeDhGroup16PtrInput
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 SystemIkeDhGroup17PtrInput
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 SystemIkeDhGroup18PtrInput
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 SystemIkeDhGroup19PtrInput
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 SystemIkeDhGroup2PtrInput
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 SystemIkeDhGroup20PtrInput
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 SystemIkeDhGroup21PtrInput
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 SystemIkeDhGroup27PtrInput
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 SystemIkeDhGroup28PtrInput
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 SystemIkeDhGroup29PtrInput
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 SystemIkeDhGroup30PtrInput
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 SystemIkeDhGroup31PtrInput
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 SystemIkeDhGroup32PtrInput
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 SystemIkeDhGroup5PtrInput
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache pulumi.StringPtrInput
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount pulumi.IntPtrInput
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle pulumi.StringPtrInput
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode pulumi.StringPtrInput
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess pulumi.StringPtrInput
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount pulumi.IntPtrInput
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIkeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIkeArgs)(nil)).Elem()
}

type SystemIkeInput interface {
	pulumi.Input

	ToSystemIkeOutput() SystemIkeOutput
	ToSystemIkeOutputWithContext(ctx context.Context) SystemIkeOutput
}

func (*SystemIke) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIke)(nil)).Elem()
}

func (i *SystemIke) ToSystemIkeOutput() SystemIkeOutput {
	return i.ToSystemIkeOutputWithContext(context.Background())
}

func (i *SystemIke) ToSystemIkeOutputWithContext(ctx context.Context) SystemIkeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIkeOutput)
}

// SystemIkeArrayInput is an input type that accepts SystemIkeArray and SystemIkeArrayOutput values.
// You can construct a concrete instance of `SystemIkeArrayInput` via:
//
//          SystemIkeArray{ SystemIkeArgs{...} }
type SystemIkeArrayInput interface {
	pulumi.Input

	ToSystemIkeArrayOutput() SystemIkeArrayOutput
	ToSystemIkeArrayOutputWithContext(context.Context) SystemIkeArrayOutput
}

type SystemIkeArray []SystemIkeInput

func (SystemIkeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIke)(nil)).Elem()
}

func (i SystemIkeArray) ToSystemIkeArrayOutput() SystemIkeArrayOutput {
	return i.ToSystemIkeArrayOutputWithContext(context.Background())
}

func (i SystemIkeArray) ToSystemIkeArrayOutputWithContext(ctx context.Context) SystemIkeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIkeArrayOutput)
}

// SystemIkeMapInput is an input type that accepts SystemIkeMap and SystemIkeMapOutput values.
// You can construct a concrete instance of `SystemIkeMapInput` via:
//
//          SystemIkeMap{ "key": SystemIkeArgs{...} }
type SystemIkeMapInput interface {
	pulumi.Input

	ToSystemIkeMapOutput() SystemIkeMapOutput
	ToSystemIkeMapOutputWithContext(context.Context) SystemIkeMapOutput
}

type SystemIkeMap map[string]SystemIkeInput

func (SystemIkeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIke)(nil)).Elem()
}

func (i SystemIkeMap) ToSystemIkeMapOutput() SystemIkeMapOutput {
	return i.ToSystemIkeMapOutputWithContext(context.Background())
}

func (i SystemIkeMap) ToSystemIkeMapOutputWithContext(ctx context.Context) SystemIkeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIkeMapOutput)
}

type SystemIkeOutput struct{ *pulumi.OutputState }

func (SystemIkeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIke)(nil)).Elem()
}

func (o SystemIkeOutput) ToSystemIkeOutput() SystemIkeOutput {
	return o
}

func (o SystemIkeOutput) ToSystemIkeOutputWithContext(ctx context.Context) SystemIkeOutput {
	return o
}

type SystemIkeArrayOutput struct{ *pulumi.OutputState }

func (SystemIkeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIke)(nil)).Elem()
}

func (o SystemIkeArrayOutput) ToSystemIkeArrayOutput() SystemIkeArrayOutput {
	return o
}

func (o SystemIkeArrayOutput) ToSystemIkeArrayOutputWithContext(ctx context.Context) SystemIkeArrayOutput {
	return o
}

func (o SystemIkeArrayOutput) Index(i pulumi.IntInput) SystemIkeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemIke {
		return vs[0].([]*SystemIke)[vs[1].(int)]
	}).(SystemIkeOutput)
}

type SystemIkeMapOutput struct{ *pulumi.OutputState }

func (SystemIkeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIke)(nil)).Elem()
}

func (o SystemIkeMapOutput) ToSystemIkeMapOutput() SystemIkeMapOutput {
	return o
}

func (o SystemIkeMapOutput) ToSystemIkeMapOutputWithContext(ctx context.Context) SystemIkeMapOutput {
	return o
}

func (o SystemIkeMapOutput) MapIndex(k pulumi.StringInput) SystemIkeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemIke {
		return vs[0].(map[string]*SystemIke)[vs[1].(string)]
	}).(SystemIkeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIkeInput)(nil)).Elem(), &SystemIke{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIkeArrayInput)(nil)).Elem(), SystemIkeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIkeMapInput)(nil)).Elem(), SystemIkeMap{})
	pulumi.RegisterOutputType(SystemIkeOutput{})
	pulumi.RegisterOutputType(SystemIkeArrayOutput{})
	pulumi.RegisterOutputType(SystemIkeMapOutput{})
}
