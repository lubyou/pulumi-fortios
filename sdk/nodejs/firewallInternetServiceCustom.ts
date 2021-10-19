// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure custom Internet Services.
 *
 * ## Import
 *
 * Firewall InternetServiceCustom can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallInternetServiceCustom:FirewallInternetServiceCustom labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallInternetServiceCustom extends pulumi.CustomResource {
    /**
     * Get an existing FirewallInternetServiceCustom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallInternetServiceCustomState, opts?: pulumi.CustomResourceOptions): FirewallInternetServiceCustom {
        return new FirewallInternetServiceCustom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallInternetServiceCustom:FirewallInternetServiceCustom';

    /**
     * Returns true if the given object is an instance of FirewallInternetServiceCustom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallInternetServiceCustom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallInternetServiceCustom.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.FirewallInternetServiceCustomEntry[] | undefined>;
    /**
     * Select the destination address or address group object from available options.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Reputation level of the custom Internet Service.
     */
    public readonly reputation!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallInternetServiceCustom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallInternetServiceCustomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallInternetServiceCustomArgs | FirewallInternetServiceCustomState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallInternetServiceCustomState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["entries"] = state ? state.entries : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["reputation"] = state ? state.reputation : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallInternetServiceCustomArgs | undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["entries"] = args ? args.entries : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["reputation"] = args ? args.reputation : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallInternetServiceCustom.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallInternetServiceCustom resources.
 */
export interface FirewallInternetServiceCustomState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.FirewallInternetServiceCustomEntry>[]>;
    /**
     * Select the destination address or address group object from available options.
     */
    name?: pulumi.Input<string>;
    /**
     * Reputation level of the custom Internet Service.
     */
    reputation?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallInternetServiceCustom resource.
 */
export interface FirewallInternetServiceCustomArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.FirewallInternetServiceCustomEntry>[]>;
    /**
     * Select the destination address or address group object from available options.
     */
    name?: pulumi.Input<string>;
    /**
     * Reputation level of the custom Internet Service.
     */
    reputation?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
