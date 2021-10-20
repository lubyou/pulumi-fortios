// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure AntiSpam options. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SpamfilterOptions("trname", {
 *     dnsTimeout: 7,
 * });
 * ```
 *
 * ## Import
 *
 * Spamfilter Options can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/spamfilterOptions:SpamfilterOptions labelname SpamfilterOptions
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SpamfilterOptions extends pulumi.CustomResource {
    /**
     * Get an existing SpamfilterOptions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpamfilterOptionsState, opts?: pulumi.CustomResourceOptions): SpamfilterOptions {
        return new SpamfilterOptions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/spamfilterOptions:SpamfilterOptions';

    /**
     * Returns true if the given object is an instance of SpamfilterOptions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SpamfilterOptions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SpamfilterOptions.__pulumiType;
    }

    /**
     * DNS query time out (1 - 30 sec).
     */
    public readonly dnsTimeout!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SpamfilterOptions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SpamfilterOptionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpamfilterOptionsArgs | SpamfilterOptionsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SpamfilterOptionsState | undefined;
            inputs["dnsTimeout"] = state ? state.dnsTimeout : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SpamfilterOptionsArgs | undefined;
            inputs["dnsTimeout"] = args ? args.dnsTimeout : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SpamfilterOptions.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SpamfilterOptions resources.
 */
export interface SpamfilterOptionsState {
    /**
     * DNS query time out (1 - 30 sec).
     */
    dnsTimeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SpamfilterOptions resource.
 */
export interface SpamfilterOptionsArgs {
    /**
     * DNS query time out (1 - 30 sec).
     */
    dnsTimeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}