// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure DNS domain filters.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.DnsfilterDomainFilter("trname", {
 *     entries: [{
 *         action: "block",
 *         domain: "bac.com",
 *         id: 1,
 *         status: "enable",
 *         type: "simple",
 *     }],
 *     fosid: 1,
 * });
 * ```
 *
 * ## Import
 *
 * Dnsfilter DomainFilter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/dnsfilterDomainFilter:DnsfilterDomainFilter labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class DnsfilterDomainFilter extends pulumi.CustomResource {
    /**
     * Get an existing DnsfilterDomainFilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsfilterDomainFilterState, opts?: pulumi.CustomResourceOptions): DnsfilterDomainFilter {
        return new DnsfilterDomainFilter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dnsfilterDomainFilter:DnsfilterDomainFilter';

    /**
     * Returns true if the given object is an instance of DnsfilterDomainFilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsfilterDomainFilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsfilterDomainFilter.__pulumiType;
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
     * DNS domain filter entries. The structure of `entries` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.DnsfilterDomainFilterEntry[] | undefined>;
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
     * Create a DnsfilterDomainFilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsfilterDomainFilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsfilterDomainFilterArgs | DnsfilterDomainFilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsfilterDomainFilterState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DnsfilterDomainFilterArgs | undefined;
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsfilterDomainFilter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsfilterDomainFilter resources.
 */
export interface DnsfilterDomainFilterState {
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * DNS domain filter entries. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.DnsfilterDomainFilterEntry>[]>;
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
 * The set of arguments for constructing a DnsfilterDomainFilter resource.
 */
export interface DnsfilterDomainFilterArgs {
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * DNS domain filter entries. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.DnsfilterDomainFilterEntry>[]>;
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
