// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure the client with its MAC address.
 *
 * ## Import
 *
 * WirelessController Address can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerAddress:WirelessControllerAddress labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerAddress extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerAddressState, opts?: pulumi.CustomResourceOptions): WirelessControllerAddress {
        return new WirelessControllerAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerAddress:WirelessControllerAddress';

    /**
     * Returns true if the given object is an instance of WirelessControllerAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerAddress.__pulumiType;
    }

    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<string>;
    /**
     * MAC address.
     */
    public readonly mac!: pulumi.Output<string>;
    /**
     * Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerAddressArgs | WirelessControllerAddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerAddressState | undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["mac"] = state ? state.mac : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerAddressArgs | undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["mac"] = args ? args.mac : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WirelessControllerAddress.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerAddress resources.
 */
export interface WirelessControllerAddressState {
    /**
     * ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * MAC address.
     */
    mac?: pulumi.Input<string>;
    /**
     * Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
     */
    policy?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerAddress resource.
 */
export interface WirelessControllerAddressArgs {
    /**
     * ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * MAC address.
     */
    mac?: pulumi.Input<string>;
    /**
     * Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
     */
    policy?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}