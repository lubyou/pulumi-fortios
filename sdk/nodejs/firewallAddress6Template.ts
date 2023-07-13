// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallAddress6Template extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAddress6Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAddress6TemplateState, opts?: pulumi.CustomResourceOptions): FirewallAddress6Template {
        return new FirewallAddress6Template(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAddress6Template:FirewallAddress6Template';

    /**
     * Returns true if the given object is an instance of FirewallAddress6Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAddress6Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAddress6Template.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly fabricObject!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly ip6!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly subnetSegmentCount!: pulumi.Output<number>;
    public readonly subnetSegments!: pulumi.Output<outputs.FirewallAddress6TemplateSubnetSegment[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallAddress6Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallAddress6TemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAddress6TemplateArgs | FirewallAddress6TemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAddress6TemplateState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fabricObject"] = state ? state.fabricObject : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["ip6"] = state ? state.ip6 : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["subnetSegmentCount"] = state ? state.subnetSegmentCount : undefined;
            resourceInputs["subnetSegments"] = state ? state.subnetSegments : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallAddress6TemplateArgs | undefined;
            if ((!args || args.ip6 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip6'");
            }
            if ((!args || args.subnetSegmentCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetSegmentCount'");
            }
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fabricObject"] = args ? args.fabricObject : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["ip6"] = args ? args.ip6 : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["subnetSegmentCount"] = args ? args.subnetSegmentCount : undefined;
            resourceInputs["subnetSegments"] = args ? args.subnetSegments : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallAddress6Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAddress6Template resources.
 */
export interface FirewallAddress6TemplateState {
    dynamicSortSubtable?: pulumi.Input<string>;
    fabricObject?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    ip6?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    subnetSegmentCount?: pulumi.Input<number>;
    subnetSegments?: pulumi.Input<pulumi.Input<inputs.FirewallAddress6TemplateSubnetSegment>[]>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAddress6Template resource.
 */
export interface FirewallAddress6TemplateArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    fabricObject?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    ip6: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    subnetSegmentCount: pulumi.Input<number>;
    subnetSegments?: pulumi.Input<pulumi.Input<inputs.FirewallAddress6TemplateSubnetSegment>[]>;
    vdomparam?: pulumi.Input<string>;
}
