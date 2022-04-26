// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IP address type availability.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WirelessControllerHotspot20AnqpIpAddressType("trname", {
 *     ipv4AddressType: "public",
 *     ipv6AddressType: "not-available",
 * });
 * ```
 *
 * ## Import
 *
 * WirelessControllerHotspot20 AnqpIpAddressType can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerHotspot20AnqpIpAddressType:WirelessControllerHotspot20AnqpIpAddressType labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerHotspot20AnqpIpAddressType:WirelessControllerHotspot20AnqpIpAddressType labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerHotspot20AnqpIpAddressType extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerHotspot20AnqpIpAddressType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerHotspot20AnqpIpAddressTypeState, opts?: pulumi.CustomResourceOptions): WirelessControllerHotspot20AnqpIpAddressType {
        return new WirelessControllerHotspot20AnqpIpAddressType(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerHotspot20AnqpIpAddressType:WirelessControllerHotspot20AnqpIpAddressType';

    /**
     * Returns true if the given object is an instance of WirelessControllerHotspot20AnqpIpAddressType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerHotspot20AnqpIpAddressType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerHotspot20AnqpIpAddressType.__pulumiType;
    }

    /**
     * IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
     */
    public readonly ipv4AddressType!: pulumi.Output<string>;
    /**
     * IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
     */
    public readonly ipv6AddressType!: pulumi.Output<string>;
    /**
     * IP type name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerHotspot20AnqpIpAddressType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerHotspot20AnqpIpAddressTypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerHotspot20AnqpIpAddressTypeArgs | WirelessControllerHotspot20AnqpIpAddressTypeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerHotspot20AnqpIpAddressTypeState | undefined;
            resourceInputs["ipv4AddressType"] = state ? state.ipv4AddressType : undefined;
            resourceInputs["ipv6AddressType"] = state ? state.ipv6AddressType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerHotspot20AnqpIpAddressTypeArgs | undefined;
            resourceInputs["ipv4AddressType"] = args ? args.ipv4AddressType : undefined;
            resourceInputs["ipv6AddressType"] = args ? args.ipv6AddressType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerHotspot20AnqpIpAddressType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerHotspot20AnqpIpAddressType resources.
 */
export interface WirelessControllerHotspot20AnqpIpAddressTypeState {
    /**
     * IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
     */
    ipv4AddressType?: pulumi.Input<string>;
    /**
     * IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
     */
    ipv6AddressType?: pulumi.Input<string>;
    /**
     * IP type name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerHotspot20AnqpIpAddressType resource.
 */
export interface WirelessControllerHotspot20AnqpIpAddressTypeArgs {
    /**
     * IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
     */
    ipv4AddressType?: pulumi.Input<string>;
    /**
     * IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
     */
    ipv6AddressType?: pulumi.Input<string>;
    /**
     * IP type name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
