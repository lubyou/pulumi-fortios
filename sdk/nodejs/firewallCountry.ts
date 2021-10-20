// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Define country table. Applies to FortiOS Version `>= 6.4.0`.
 *
 * ## Import
 *
 * Firewall Country can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallCountry:FirewallCountry labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallCountry extends pulumi.CustomResource {
    /**
     * Get an existing FirewallCountry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallCountryState, opts?: pulumi.CustomResourceOptions): FirewallCountry {
        return new FirewallCountry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallCountry:FirewallCountry';

    /**
     * Returns true if the given object is an instance of FirewallCountry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallCountry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallCountry.__pulumiType;
    }

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Country ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Country name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Region ID list. The structure of `region` block is documented below.
     */
    public readonly regions!: pulumi.Output<outputs.FirewallCountryRegion[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallCountry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallCountryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallCountryArgs | FirewallCountryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallCountryState | undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["regions"] = state ? state.regions : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallCountryArgs | undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["regions"] = args ? args.regions : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallCountry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallCountry resources.
 */
export interface FirewallCountryState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Country ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Country name.
     */
    name?: pulumi.Input<string>;
    /**
     * Region ID list. The structure of `region` block is documented below.
     */
    regions?: pulumi.Input<pulumi.Input<inputs.FirewallCountryRegion>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallCountry resource.
 */
export interface FirewallCountryArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Country ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Country name.
     */
    name?: pulumi.Input<string>;
    /**
     * Region ID list. The structure of `region` block is documented below.
     */
    regions?: pulumi.Input<pulumi.Input<inputs.FirewallCountryRegion>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}