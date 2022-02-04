// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch traffic policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SwitchControllerTrafficPolicy("trname", {
 *     guaranteedBandwidth: 0,
 *     guaranteedBurst: 0,
 *     maximumBurst: 0,
 *     policerStatus: "enable",
 *     type: "ingress",
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController TrafficPolicy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerTrafficPolicy:SwitchControllerTrafficPolicy labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerTrafficPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerTrafficPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerTrafficPolicyState, opts?: pulumi.CustomResourceOptions): SwitchControllerTrafficPolicy {
        return new SwitchControllerTrafficPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerTrafficPolicy:SwitchControllerTrafficPolicy';

    /**
     * Returns true if the given object is an instance of SwitchControllerTrafficPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerTrafficPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerTrafficPolicy.__pulumiType;
    }

    /**
     * COS queue(0 - 7), or unset to disable.
     */
    public readonly cos!: pulumi.Output<number>;
    /**
     * COS queue(0 - 7), or unset to disable.
     */
    public readonly cosQueue!: pulumi.Output<number>;
    /**
     * Description of the traffic policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * FSW Policer id
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Guaranteed bandwidth in kbps (max value = 524287000).
     */
    public readonly guaranteedBandwidth!: pulumi.Output<number>;
    /**
     * Guaranteed burst size in bytes (max value = 4294967295).
     */
    public readonly guaranteedBurst!: pulumi.Output<number>;
    /**
     * Maximum burst size in bytes (max value = 4294967295).
     */
    public readonly maximumBurst!: pulumi.Output<number>;
    /**
     * Traffic policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
     */
    public readonly policerStatus!: pulumi.Output<string>;
    /**
     * Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerTrafficPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerTrafficPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerTrafficPolicyArgs | SwitchControllerTrafficPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerTrafficPolicyState | undefined;
            resourceInputs["cos"] = state ? state.cos : undefined;
            resourceInputs["cosQueue"] = state ? state.cosQueue : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["guaranteedBandwidth"] = state ? state.guaranteedBandwidth : undefined;
            resourceInputs["guaranteedBurst"] = state ? state.guaranteedBurst : undefined;
            resourceInputs["maximumBurst"] = state ? state.maximumBurst : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policerStatus"] = state ? state.policerStatus : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerTrafficPolicyArgs | undefined;
            resourceInputs["cos"] = args ? args.cos : undefined;
            resourceInputs["cosQueue"] = args ? args.cosQueue : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["guaranteedBandwidth"] = args ? args.guaranteedBandwidth : undefined;
            resourceInputs["guaranteedBurst"] = args ? args.guaranteedBurst : undefined;
            resourceInputs["maximumBurst"] = args ? args.maximumBurst : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policerStatus"] = args ? args.policerStatus : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerTrafficPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerTrafficPolicy resources.
 */
export interface SwitchControllerTrafficPolicyState {
    /**
     * COS queue(0 - 7), or unset to disable.
     */
    cos?: pulumi.Input<number>;
    /**
     * COS queue(0 - 7), or unset to disable.
     */
    cosQueue?: pulumi.Input<number>;
    /**
     * Description of the traffic policy.
     */
    description?: pulumi.Input<string>;
    /**
     * FSW Policer id
     */
    fosid?: pulumi.Input<number>;
    /**
     * Guaranteed bandwidth in kbps (max value = 524287000).
     */
    guaranteedBandwidth?: pulumi.Input<number>;
    /**
     * Guaranteed burst size in bytes (max value = 4294967295).
     */
    guaranteedBurst?: pulumi.Input<number>;
    /**
     * Maximum burst size in bytes (max value = 4294967295).
     */
    maximumBurst?: pulumi.Input<number>;
    /**
     * Traffic policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
     */
    policerStatus?: pulumi.Input<string>;
    /**
     * Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerTrafficPolicy resource.
 */
export interface SwitchControllerTrafficPolicyArgs {
    /**
     * COS queue(0 - 7), or unset to disable.
     */
    cos?: pulumi.Input<number>;
    /**
     * COS queue(0 - 7), or unset to disable.
     */
    cosQueue?: pulumi.Input<number>;
    /**
     * Description of the traffic policy.
     */
    description?: pulumi.Input<string>;
    /**
     * FSW Policer id
     */
    fosid?: pulumi.Input<number>;
    /**
     * Guaranteed bandwidth in kbps (max value = 524287000).
     */
    guaranteedBandwidth?: pulumi.Input<number>;
    /**
     * Guaranteed burst size in bytes (max value = 4294967295).
     */
    guaranteedBurst?: pulumi.Input<number>;
    /**
     * Maximum burst size in bytes (max value = 4294967295).
     */
    maximumBurst?: pulumi.Input<number>;
    /**
     * Traffic policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
     */
    policerStatus?: pulumi.Input<string>;
    /**
     * Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
