// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure NAT64. Applies to FortiOS Version `<= 7.0.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemNat64("trname", {
 *     alwaysSynthesizeAaaaRecord: "enable",
 *     generateIpv6FragmentHeader: "disable",
 *     nat46ForceIpv4PacketForwarding: "disable",
 *     nat64Prefix: "2001:1:2:3::/96",
 *     secondaryPrefixStatus: "disable",
 *     status: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * System Nat64 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemNat64:SystemNat64 labelname SystemNat64
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemNat64 extends pulumi.CustomResource {
    /**
     * Get an existing SystemNat64 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemNat64State, opts?: pulumi.CustomResourceOptions): SystemNat64 {
        return new SystemNat64(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemNat64:SystemNat64';

    /**
     * Returns true if the given object is an instance of SystemNat64.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemNat64 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemNat64.__pulumiType;
    }

    /**
     * Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
     */
    public readonly alwaysSynthesizeAaaaRecord!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
     */
    public readonly generateIpv6FragmentHeader!: pulumi.Output<string>;
    /**
     * Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
     */
    public readonly nat46ForceIpv4PacketForwarding!: pulumi.Output<string>;
    /**
     * NAT64 prefix.
     */
    public readonly nat64Prefix!: pulumi.Output<string>;
    /**
     * Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
     */
    public readonly secondaryPrefixStatus!: pulumi.Output<string>;
    /**
     * Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
     */
    public readonly secondaryPrefixes!: pulumi.Output<outputs.SystemNat64SecondaryPrefix[] | undefined>;
    /**
     * Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemNat64 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemNat64Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemNat64Args | SystemNat64State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemNat64State | undefined;
            resourceInputs["alwaysSynthesizeAaaaRecord"] = state ? state.alwaysSynthesizeAaaaRecord : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["generateIpv6FragmentHeader"] = state ? state.generateIpv6FragmentHeader : undefined;
            resourceInputs["nat46ForceIpv4PacketForwarding"] = state ? state.nat46ForceIpv4PacketForwarding : undefined;
            resourceInputs["nat64Prefix"] = state ? state.nat64Prefix : undefined;
            resourceInputs["secondaryPrefixStatus"] = state ? state.secondaryPrefixStatus : undefined;
            resourceInputs["secondaryPrefixes"] = state ? state.secondaryPrefixes : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemNat64Args | undefined;
            if ((!args || args.nat64Prefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nat64Prefix'");
            }
            resourceInputs["alwaysSynthesizeAaaaRecord"] = args ? args.alwaysSynthesizeAaaaRecord : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["generateIpv6FragmentHeader"] = args ? args.generateIpv6FragmentHeader : undefined;
            resourceInputs["nat46ForceIpv4PacketForwarding"] = args ? args.nat46ForceIpv4PacketForwarding : undefined;
            resourceInputs["nat64Prefix"] = args ? args.nat64Prefix : undefined;
            resourceInputs["secondaryPrefixStatus"] = args ? args.secondaryPrefixStatus : undefined;
            resourceInputs["secondaryPrefixes"] = args ? args.secondaryPrefixes : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemNat64.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemNat64 resources.
 */
export interface SystemNat64State {
    /**
     * Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
     */
    alwaysSynthesizeAaaaRecord?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
     */
    generateIpv6FragmentHeader?: pulumi.Input<string>;
    /**
     * Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
     */
    nat46ForceIpv4PacketForwarding?: pulumi.Input<string>;
    /**
     * NAT64 prefix.
     */
    nat64Prefix?: pulumi.Input<string>;
    /**
     * Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
     */
    secondaryPrefixStatus?: pulumi.Input<string>;
    /**
     * Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
     */
    secondaryPrefixes?: pulumi.Input<pulumi.Input<inputs.SystemNat64SecondaryPrefix>[]>;
    /**
     * Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemNat64 resource.
 */
export interface SystemNat64Args {
    /**
     * Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
     */
    alwaysSynthesizeAaaaRecord?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
     */
    generateIpv6FragmentHeader?: pulumi.Input<string>;
    /**
     * Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
     */
    nat46ForceIpv4PacketForwarding?: pulumi.Input<string>;
    /**
     * NAT64 prefix.
     */
    nat64Prefix: pulumi.Input<string>;
    /**
     * Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
     */
    secondaryPrefixStatus?: pulumi.Input<string>;
    /**
     * Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
     */
    secondaryPrefixes?: pulumi.Input<pulumi.Input<inputs.SystemNat64SecondaryPrefix>[]>;
    /**
     * Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
