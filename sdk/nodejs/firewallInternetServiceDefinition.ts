// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallInternetServiceDefinition extends pulumi.CustomResource {
    /**
     * Get an existing FirewallInternetServiceDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallInternetServiceDefinitionState, opts?: pulumi.CustomResourceOptions): FirewallInternetServiceDefinition {
        return new FirewallInternetServiceDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallInternetServiceDefinition:FirewallInternetServiceDefinition';

    /**
     * Returns true if the given object is an instance of FirewallInternetServiceDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallInternetServiceDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallInternetServiceDefinition.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly entries!: pulumi.Output<outputs.FirewallInternetServiceDefinitionEntry[] | undefined>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallInternetServiceDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallInternetServiceDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallInternetServiceDefinitionArgs | FirewallInternetServiceDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallInternetServiceDefinitionState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallInternetServiceDefinitionArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallInternetServiceDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallInternetServiceDefinition resources.
 */
export interface FirewallInternetServiceDefinitionState {
    dynamicSortSubtable?: pulumi.Input<string>;
    entries?: pulumi.Input<pulumi.Input<inputs.FirewallInternetServiceDefinitionEntry>[]>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallInternetServiceDefinition resource.
 */
export interface FirewallInternetServiceDefinitionArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    entries?: pulumi.Input<pulumi.Input<inputs.FirewallInternetServiceDefinitionEntry>[]>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
