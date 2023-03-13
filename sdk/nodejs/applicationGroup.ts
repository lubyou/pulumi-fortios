// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ApplicationGroup extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationGroupState, opts?: pulumi.CustomResourceOptions): ApplicationGroup {
        return new ApplicationGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/applicationGroup:ApplicationGroup';

    /**
     * Returns true if the given object is an instance of ApplicationGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationGroup.__pulumiType;
    }

    public readonly applications!: pulumi.Output<outputs.ApplicationGroupApplication[] | undefined>;
    public readonly behavior!: pulumi.Output<string>;
    public readonly categories!: pulumi.Output<outputs.ApplicationGroupCategory[] | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly popularity!: pulumi.Output<string>;
    public readonly protocols!: pulumi.Output<string>;
    public readonly risks!: pulumi.Output<outputs.ApplicationGroupRisk[] | undefined>;
    public readonly technology!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly vendor!: pulumi.Output<string>;

    /**
     * Create a ApplicationGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApplicationGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationGroupArgs | ApplicationGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationGroupState | undefined;
            resourceInputs["applications"] = state ? state.applications : undefined;
            resourceInputs["behavior"] = state ? state.behavior : undefined;
            resourceInputs["categories"] = state ? state.categories : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["popularity"] = state ? state.popularity : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["risks"] = state ? state.risks : undefined;
            resourceInputs["technology"] = state ? state.technology : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
        } else {
            const args = argsOrState as ApplicationGroupArgs | undefined;
            resourceInputs["applications"] = args ? args.applications : undefined;
            resourceInputs["behavior"] = args ? args.behavior : undefined;
            resourceInputs["categories"] = args ? args.categories : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["popularity"] = args ? args.popularity : undefined;
            resourceInputs["protocols"] = args ? args.protocols : undefined;
            resourceInputs["risks"] = args ? args.risks : undefined;
            resourceInputs["technology"] = args ? args.technology : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationGroup resources.
 */
export interface ApplicationGroupState {
    applications?: pulumi.Input<pulumi.Input<inputs.ApplicationGroupApplication>[]>;
    behavior?: pulumi.Input<string>;
    categories?: pulumi.Input<pulumi.Input<inputs.ApplicationGroupCategory>[]>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    popularity?: pulumi.Input<string>;
    protocols?: pulumi.Input<string>;
    risks?: pulumi.Input<pulumi.Input<inputs.ApplicationGroupRisk>[]>;
    technology?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vendor?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationGroup resource.
 */
export interface ApplicationGroupArgs {
    applications?: pulumi.Input<pulumi.Input<inputs.ApplicationGroupApplication>[]>;
    behavior?: pulumi.Input<string>;
    categories?: pulumi.Input<pulumi.Input<inputs.ApplicationGroupCategory>[]>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    popularity?: pulumi.Input<string>;
    protocols?: pulumi.Input<string>;
    risks?: pulumi.Input<pulumi.Input<inputs.ApplicationGroupRisk>[]>;
    technology?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vendor?: pulumi.Input<string>;
}
