// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure SNMP. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * WirelessController Snmp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerSnmp:WirelessControllerSnmp labelname WirelessControllerSnmp
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerSnmp extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerSnmp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerSnmpState, opts?: pulumi.CustomResourceOptions): WirelessControllerSnmp {
        return new WirelessControllerSnmp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerSnmp:WirelessControllerSnmp';

    /**
     * Returns true if the given object is an instance of WirelessControllerSnmp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerSnmp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerSnmp.__pulumiType;
    }

    /**
     * SNMP Community Configuration. The structure of `community` block is documented below.
     */
    public readonly communities!: pulumi.Output<outputs.WirelessControllerSnmpCommunity[] | undefined>;
    /**
     * Contact Information.
     */
    public readonly contactInfo!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * AC SNMP engineId string (maximum 24 characters).
     */
    public readonly engineId!: pulumi.Output<string>;
    /**
     * CPU usage when trap is sent.
     */
    public readonly trapHighCpuThreshold!: pulumi.Output<number>;
    /**
     * Memory usage when trap is sent.
     */
    public readonly trapHighMemThreshold!: pulumi.Output<number>;
    /**
     * SNMP User Configuration. The structure of `user` block is documented below.
     */
    public readonly users!: pulumi.Output<outputs.WirelessControllerSnmpUser[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerSnmp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerSnmpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerSnmpArgs | WirelessControllerSnmpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerSnmpState | undefined;
            resourceInputs["communities"] = state ? state.communities : undefined;
            resourceInputs["contactInfo"] = state ? state.contactInfo : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["engineId"] = state ? state.engineId : undefined;
            resourceInputs["trapHighCpuThreshold"] = state ? state.trapHighCpuThreshold : undefined;
            resourceInputs["trapHighMemThreshold"] = state ? state.trapHighMemThreshold : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerSnmpArgs | undefined;
            resourceInputs["communities"] = args ? args.communities : undefined;
            resourceInputs["contactInfo"] = args ? args.contactInfo : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["engineId"] = args ? args.engineId : undefined;
            resourceInputs["trapHighCpuThreshold"] = args ? args.trapHighCpuThreshold : undefined;
            resourceInputs["trapHighMemThreshold"] = args ? args.trapHighMemThreshold : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerSnmp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerSnmp resources.
 */
export interface WirelessControllerSnmpState {
    /**
     * SNMP Community Configuration. The structure of `community` block is documented below.
     */
    communities?: pulumi.Input<pulumi.Input<inputs.WirelessControllerSnmpCommunity>[]>;
    /**
     * Contact Information.
     */
    contactInfo?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * AC SNMP engineId string (maximum 24 characters).
     */
    engineId?: pulumi.Input<string>;
    /**
     * CPU usage when trap is sent.
     */
    trapHighCpuThreshold?: pulumi.Input<number>;
    /**
     * Memory usage when trap is sent.
     */
    trapHighMemThreshold?: pulumi.Input<number>;
    /**
     * SNMP User Configuration. The structure of `user` block is documented below.
     */
    users?: pulumi.Input<pulumi.Input<inputs.WirelessControllerSnmpUser>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerSnmp resource.
 */
export interface WirelessControllerSnmpArgs {
    /**
     * SNMP Community Configuration. The structure of `community` block is documented below.
     */
    communities?: pulumi.Input<pulumi.Input<inputs.WirelessControllerSnmpCommunity>[]>;
    /**
     * Contact Information.
     */
    contactInfo?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * AC SNMP engineId string (maximum 24 characters).
     */
    engineId?: pulumi.Input<string>;
    /**
     * CPU usage when trap is sent.
     */
    trapHighCpuThreshold?: pulumi.Input<number>;
    /**
     * Memory usage when trap is sent.
     */
    trapHighMemThreshold?: pulumi.Input<number>;
    /**
     * SNMP User Configuration. The structure of `user` block is documented below.
     */
    users?: pulumi.Input<pulumi.Input<inputs.WirelessControllerSnmpUser>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
