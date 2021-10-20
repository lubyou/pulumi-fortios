// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch SNMP trap threshold values globally.
 *
 * ## Import
 *
 * SwitchController SnmpTrapThreshold can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerSnmpTrapThreshold:SwitchControllerSnmpTrapThreshold labelname SwitchControllerSnmpTrapThreshold
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerSnmpTrapThreshold extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerSnmpTrapThreshold resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerSnmpTrapThresholdState, opts?: pulumi.CustomResourceOptions): SwitchControllerSnmpTrapThreshold {
        return new SwitchControllerSnmpTrapThreshold(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerSnmpTrapThreshold:SwitchControllerSnmpTrapThreshold';

    /**
     * Returns true if the given object is an instance of SwitchControllerSnmpTrapThreshold.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerSnmpTrapThreshold {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerSnmpTrapThreshold.__pulumiType;
    }

    /**
     * CPU usage when trap is sent.
     */
    public readonly trapHighCpuThreshold!: pulumi.Output<number>;
    /**
     * Log disk usage when trap is sent.
     */
    public readonly trapLogFullThreshold!: pulumi.Output<number>;
    /**
     * Memory usage when trap is sent.
     */
    public readonly trapLowMemoryThreshold!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerSnmpTrapThreshold resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerSnmpTrapThresholdArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerSnmpTrapThresholdArgs | SwitchControllerSnmpTrapThresholdState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerSnmpTrapThresholdState | undefined;
            inputs["trapHighCpuThreshold"] = state ? state.trapHighCpuThreshold : undefined;
            inputs["trapLogFullThreshold"] = state ? state.trapLogFullThreshold : undefined;
            inputs["trapLowMemoryThreshold"] = state ? state.trapLowMemoryThreshold : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerSnmpTrapThresholdArgs | undefined;
            inputs["trapHighCpuThreshold"] = args ? args.trapHighCpuThreshold : undefined;
            inputs["trapLogFullThreshold"] = args ? args.trapLogFullThreshold : undefined;
            inputs["trapLowMemoryThreshold"] = args ? args.trapLowMemoryThreshold : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerSnmpTrapThreshold.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerSnmpTrapThreshold resources.
 */
export interface SwitchControllerSnmpTrapThresholdState {
    /**
     * CPU usage when trap is sent.
     */
    trapHighCpuThreshold?: pulumi.Input<number>;
    /**
     * Log disk usage when trap is sent.
     */
    trapLogFullThreshold?: pulumi.Input<number>;
    /**
     * Memory usage when trap is sent.
     */
    trapLowMemoryThreshold?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerSnmpTrapThreshold resource.
 */
export interface SwitchControllerSnmpTrapThresholdArgs {
    /**
     * CPU usage when trap is sent.
     */
    trapHighCpuThreshold?: pulumi.Input<number>;
    /**
     * Log disk usage when trap is sent.
     */
    trapLogFullThreshold?: pulumi.Input<number>;
    /**
     * Memory usage when trap is sent.
     */
    trapLowMemoryThreshold?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}