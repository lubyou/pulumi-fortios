// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure WiFi SSID policies. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * WirelessController SsidPolicy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerSsidPolicy:WirelessControllerSsidPolicy labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerSsidPolicy extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerSsidPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerSsidPolicyState, opts?: pulumi.CustomResourceOptions): WirelessControllerSsidPolicy {
        return new WirelessControllerSsidPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerSsidPolicy:WirelessControllerSsidPolicy';

    /**
     * Returns true if the given object is an instance of WirelessControllerSsidPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerSsidPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerSsidPolicy.__pulumiType;
    }

    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * VLAN interface name.
     */
    public readonly vlan!: pulumi.Output<string>;

    /**
     * Create a WirelessControllerSsidPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerSsidPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerSsidPolicyArgs | WirelessControllerSsidPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerSsidPolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as WirelessControllerSsidPolicyArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerSsidPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerSsidPolicy resources.
 */
export interface WirelessControllerSsidPolicyState {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN interface name.
     */
    vlan?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerSsidPolicy resource.
 */
export interface WirelessControllerSsidPolicyArgs {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN interface name.
     */
    vlan?: pulumi.Input<string>;
}