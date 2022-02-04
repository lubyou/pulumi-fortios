// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure Spanning Tree Protocol (STP). Applies to FortiOS Version `>= 7.0.4`.
 *
 * ## Import
 *
 * System Stp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemStp:SystemStp labelname SystemStp
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemStp extends pulumi.CustomResource {
    /**
     * Get an existing SystemStp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemStpState, opts?: pulumi.CustomResourceOptions): SystemStp {
        return new SystemStp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemStp:SystemStp';

    /**
     * Returns true if the given object is an instance of SystemStp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemStp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemStp.__pulumiType;
    }

    /**
     * Forward delay (4 - 30 sec, default = 15).
     */
    public readonly forwardDelay!: pulumi.Output<number>;
    /**
     * Hello time (1 - 10 sec, default = 2).
     */
    public readonly helloTime!: pulumi.Output<number>;
    /**
     * Maximum packet age (6 - 40 sec, default = 20).
     */
    public readonly maxAge!: pulumi.Output<number>;
    /**
     * Maximum number of hops (1 - 40, default = 20).
     */
    public readonly maxHops!: pulumi.Output<number>;
    /**
     * STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
     */
    public readonly switchPriority!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemStp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemStpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemStpArgs | SystemStpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemStpState | undefined;
            resourceInputs["forwardDelay"] = state ? state.forwardDelay : undefined;
            resourceInputs["helloTime"] = state ? state.helloTime : undefined;
            resourceInputs["maxAge"] = state ? state.maxAge : undefined;
            resourceInputs["maxHops"] = state ? state.maxHops : undefined;
            resourceInputs["switchPriority"] = state ? state.switchPriority : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemStpArgs | undefined;
            resourceInputs["forwardDelay"] = args ? args.forwardDelay : undefined;
            resourceInputs["helloTime"] = args ? args.helloTime : undefined;
            resourceInputs["maxAge"] = args ? args.maxAge : undefined;
            resourceInputs["maxHops"] = args ? args.maxHops : undefined;
            resourceInputs["switchPriority"] = args ? args.switchPriority : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemStp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemStp resources.
 */
export interface SystemStpState {
    /**
     * Forward delay (4 - 30 sec, default = 15).
     */
    forwardDelay?: pulumi.Input<number>;
    /**
     * Hello time (1 - 10 sec, default = 2).
     */
    helloTime?: pulumi.Input<number>;
    /**
     * Maximum packet age (6 - 40 sec, default = 20).
     */
    maxAge?: pulumi.Input<number>;
    /**
     * Maximum number of hops (1 - 40, default = 20).
     */
    maxHops?: pulumi.Input<number>;
    /**
     * STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
     */
    switchPriority?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemStp resource.
 */
export interface SystemStpArgs {
    /**
     * Forward delay (4 - 30 sec, default = 15).
     */
    forwardDelay?: pulumi.Input<number>;
    /**
     * Hello time (1 - 10 sec, default = 2).
     */
    helloTime?: pulumi.Input<number>;
    /**
     * Maximum packet age (6 - 40 sec, default = 20).
     */
    maxAge?: pulumi.Input<number>;
    /**
     * Maximum number of hops (1 - 40, default = 20).
     */
    maxHops?: pulumi.Input<number>;
    /**
     * STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
     */
    switchPriority?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
