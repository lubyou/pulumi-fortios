// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure global session TTL timers for this FortiGate.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemSessionTtl("trname", {
 *     default: "3600",
 * });
 * ```
 *
 * ## Import
 *
 * System SessionTtl can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemSessionTtl:SystemSessionTtl labelname SystemSessionTtl
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemSessionTtl extends pulumi.CustomResource {
    /**
     * Get an existing SystemSessionTtl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSessionTtlState, opts?: pulumi.CustomResourceOptions): SystemSessionTtl {
        return new SystemSessionTtl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSessionTtl:SystemSessionTtl';

    /**
     * Returns true if the given object is an instance of SystemSessionTtl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSessionTtl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSessionTtl.__pulumiType;
    }

    /**
     * Default timeout.
     */
    public readonly default!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Session TTL port. The structure of `port` block is documented below.
     */
    public readonly ports!: pulumi.Output<outputs.SystemSessionTtlPort[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemSessionTtl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemSessionTtlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSessionTtlArgs | SystemSessionTtlState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSessionTtlState | undefined;
            inputs["default"] = state ? state.default : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["ports"] = state ? state.ports : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemSessionTtlArgs | undefined;
            inputs["default"] = args ? args.default : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["ports"] = args ? args.ports : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemSessionTtl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSessionTtl resources.
 */
export interface SystemSessionTtlState {
    /**
     * Default timeout.
     */
    default?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Session TTL port. The structure of `port` block is documented below.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.SystemSessionTtlPort>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemSessionTtl resource.
 */
export interface SystemSessionTtlArgs {
    /**
     * Default timeout.
     */
    default?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Session TTL port. The structure of `port` block is documented below.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.SystemSessionTtlPort>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}