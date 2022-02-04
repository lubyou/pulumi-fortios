// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure virtual domain.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemVdom("trname", {
 *     shortName: "testvdom",
 *     temporary: 0,
 * });
 * ```
 *
 * ## Import
 *
 * System Vdom can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemVdom:SystemVdom labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemVdom extends pulumi.CustomResource {
    /**
     * Get an existing SystemVdom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemVdomState, opts?: pulumi.CustomResourceOptions): SystemVdom {
        return new SystemVdom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemVdom:SystemVdom';

    /**
     * Returns true if the given object is an instance of SystemVdom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemVdom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemVdom.__pulumiType;
    }

    /**
     * Flag.
     */
    public readonly flag!: pulumi.Output<number>;
    /**
     * VDOM name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * VDOM short name.
     */
    public readonly shortName!: pulumi.Output<string>;
    /**
     * Temporary.
     */
    public readonly temporary!: pulumi.Output<number>;
    /**
     * Virtual cluster ID (0 - 4294967295).
     */
    public readonly vclusterId!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemVdom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemVdomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemVdomArgs | SystemVdomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemVdomState | undefined;
            resourceInputs["flag"] = state ? state.flag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["shortName"] = state ? state.shortName : undefined;
            resourceInputs["temporary"] = state ? state.temporary : undefined;
            resourceInputs["vclusterId"] = state ? state.vclusterId : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemVdomArgs | undefined;
            resourceInputs["flag"] = args ? args.flag : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["shortName"] = args ? args.shortName : undefined;
            resourceInputs["temporary"] = args ? args.temporary : undefined;
            resourceInputs["vclusterId"] = args ? args.vclusterId : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemVdom.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemVdom resources.
 */
export interface SystemVdomState {
    /**
     * Flag.
     */
    flag?: pulumi.Input<number>;
    /**
     * VDOM name.
     */
    name?: pulumi.Input<string>;
    /**
     * VDOM short name.
     */
    shortName?: pulumi.Input<string>;
    /**
     * Temporary.
     */
    temporary?: pulumi.Input<number>;
    /**
     * Virtual cluster ID (0 - 4294967295).
     */
    vclusterId?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemVdom resource.
 */
export interface SystemVdomArgs {
    /**
     * Flag.
     */
    flag?: pulumi.Input<number>;
    /**
     * VDOM name.
     */
    name?: pulumi.Input<string>;
    /**
     * VDOM short name.
     */
    shortName?: pulumi.Input<string>;
    /**
     * Temporary.
     */
    temporary?: pulumi.Input<number>;
    /**
     * Virtual cluster ID (0 - 4294967295).
     */
    vclusterId?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
