// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure AntiSpam IP trust. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Import
 *
 * Spamfilter Iptrust can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/spamfilterIptrust:SpamfilterIptrust labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SpamfilterIptrust extends pulumi.CustomResource {
    /**
     * Get an existing SpamfilterIptrust resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpamfilterIptrustState, opts?: pulumi.CustomResourceOptions): SpamfilterIptrust {
        return new SpamfilterIptrust(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/spamfilterIptrust:SpamfilterIptrust';

    /**
     * Returns true if the given object is an instance of SpamfilterIptrust.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SpamfilterIptrust {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SpamfilterIptrust.__pulumiType;
    }

    /**
     * Optional comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Spam filter trusted IP addresses. The structure of `entries` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.SpamfilterIptrustEntry[] | undefined>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Name of table.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SpamfilterIptrust resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SpamfilterIptrustArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpamfilterIptrustArgs | SpamfilterIptrustState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SpamfilterIptrustState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["entries"] = state ? state.entries : undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SpamfilterIptrustArgs | undefined;
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            inputs["comment"] = args ? args.comment : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["entries"] = args ? args.entries : undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SpamfilterIptrust.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SpamfilterIptrust resources.
 */
export interface SpamfilterIptrustState {
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Spam filter trusted IP addresses. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.SpamfilterIptrustEntry>[]>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Name of table.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SpamfilterIptrust resource.
 */
export interface SpamfilterIptrustArgs {
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Spam filter trusted IP addresses. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.SpamfilterIptrustEntry>[]>;
    /**
     * ID.
     */
    fosid: pulumi.Input<number>;
    /**
     * Name of table.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
