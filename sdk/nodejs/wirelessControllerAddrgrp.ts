// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure the MAC address group. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * WirelessController Addrgrp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerAddrgrp extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerAddrgrp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerAddrgrpState, opts?: pulumi.CustomResourceOptions): WirelessControllerAddrgrp {
        return new WirelessControllerAddrgrp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp';

    /**
     * Returns true if the given object is an instance of WirelessControllerAddrgrp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerAddrgrp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerAddrgrp.__pulumiType;
    }

    /**
     * Manually selected group of addresses. The structure of `addresses` block is documented below.
     */
    public readonly addresses!: pulumi.Output<outputs.WirelessControllerAddrgrpAddress[] | undefined>;
    /**
     * Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
     */
    public readonly defaultPolicy!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerAddrgrp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerAddrgrpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerAddrgrpArgs | WirelessControllerAddrgrpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerAddrgrpState | undefined;
            resourceInputs["addresses"] = state ? state.addresses : undefined;
            resourceInputs["defaultPolicy"] = state ? state.defaultPolicy : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerAddrgrpArgs | undefined;
            resourceInputs["addresses"] = args ? args.addresses : undefined;
            resourceInputs["defaultPolicy"] = args ? args.defaultPolicy : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerAddrgrp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerAddrgrp resources.
 */
export interface WirelessControllerAddrgrpState {
    /**
     * Manually selected group of addresses. The structure of `addresses` block is documented below.
     */
    addresses?: pulumi.Input<pulumi.Input<inputs.WirelessControllerAddrgrpAddress>[]>;
    /**
     * Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
     */
    defaultPolicy?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerAddrgrp resource.
 */
export interface WirelessControllerAddrgrpArgs {
    /**
     * Manually selected group of addresses. The structure of `addresses` block is documented below.
     */
    addresses?: pulumi.Input<pulumi.Input<inputs.WirelessControllerAddrgrpAddress>[]>;
    /**
     * Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
     */
    defaultPolicy?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
