// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch IGMP snooping global settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SwitchControllerIgmpSnooping("trname", {
 *     agingTime: 300,
 *     floodUnknownMulticast: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController IgmpSnooping can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerIgmpSnooping:SwitchControllerIgmpSnooping labelname SwitchControllerIgmpSnooping
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerIgmpSnooping extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerIgmpSnooping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerIgmpSnoopingState, opts?: pulumi.CustomResourceOptions): SwitchControllerIgmpSnooping {
        return new SwitchControllerIgmpSnooping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerIgmpSnooping:SwitchControllerIgmpSnooping';

    /**
     * Returns true if the given object is an instance of SwitchControllerIgmpSnooping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerIgmpSnooping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerIgmpSnooping.__pulumiType;
    }

    /**
     * Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
     */
    public readonly agingTime!: pulumi.Output<number>;
    /**
     * Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
     */
    public readonly floodUnknownMulticast!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerIgmpSnooping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerIgmpSnoopingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerIgmpSnoopingArgs | SwitchControllerIgmpSnoopingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerIgmpSnoopingState | undefined;
            resourceInputs["agingTime"] = state ? state.agingTime : undefined;
            resourceInputs["floodUnknownMulticast"] = state ? state.floodUnknownMulticast : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerIgmpSnoopingArgs | undefined;
            resourceInputs["agingTime"] = args ? args.agingTime : undefined;
            resourceInputs["floodUnknownMulticast"] = args ? args.floodUnknownMulticast : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerIgmpSnooping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerIgmpSnooping resources.
 */
export interface SwitchControllerIgmpSnoopingState {
    /**
     * Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
     */
    agingTime?: pulumi.Input<number>;
    /**
     * Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
     */
    floodUnknownMulticast?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerIgmpSnooping resource.
 */
export interface SwitchControllerIgmpSnoopingArgs {
    /**
     * Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
     */
    agingTime?: pulumi.Input<number>;
    /**
     * Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
     */
    floodUnknownMulticast?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
