// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure Internet Services Extension.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallInternetServiceExtension("trname", {
 *     comment: "EIWE",
 *     fosid: 65536,
 * });
 * ```
 *
 * ## Import
 *
 * Firewall InternetServiceExtension can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallInternetServiceExtension extends pulumi.CustomResource {
    /**
     * Get an existing FirewallInternetServiceExtension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallInternetServiceExtensionState, opts?: pulumi.CustomResourceOptions): FirewallInternetServiceExtension {
        return new FirewallInternetServiceExtension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension';

    /**
     * Returns true if the given object is an instance of FirewallInternetServiceExtension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallInternetServiceExtension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallInternetServiceExtension.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
     */
    public readonly disableEntries!: pulumi.Output<outputs.FirewallInternetServiceExtensionDisableEntry[] | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.FirewallInternetServiceExtensionEntry[] | undefined>;
    /**
     * Internet Service ID in the Internet Service database.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallInternetServiceExtension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallInternetServiceExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallInternetServiceExtensionArgs | FirewallInternetServiceExtensionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallInternetServiceExtensionState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["disableEntries"] = state ? state.disableEntries : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallInternetServiceExtensionArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["disableEntries"] = args ? args.disableEntries : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallInternetServiceExtension.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallInternetServiceExtension resources.
 */
export interface FirewallInternetServiceExtensionState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
     */
    disableEntries?: pulumi.Input<pulumi.Input<inputs.FirewallInternetServiceExtensionDisableEntry>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.FirewallInternetServiceExtensionEntry>[]>;
    /**
     * Internet Service ID in the Internet Service database.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallInternetServiceExtension resource.
 */
export interface FirewallInternetServiceExtensionArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
     */
    disableEntries?: pulumi.Input<pulumi.Input<inputs.FirewallInternetServiceExtensionDisableEntry>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.FirewallInternetServiceExtensionEntry>[]>;
    /**
     * Internet Service ID in the Internet Service database.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
