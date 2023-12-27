// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class RuleOtdt extends pulumi.CustomResource {
    /**
     * Get an existing RuleOtdt resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleOtdtState, opts?: pulumi.CustomResourceOptions): RuleOtdt {
        return new RuleOtdt(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/ruleOtdt:RuleOtdt';

    /**
     * Returns true if the given object is an instance of RuleOtdt.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleOtdt {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleOtdt.__pulumiType;
    }

    public readonly behavior!: pulumi.Output<string>;
    public readonly category!: pulumi.Output<number>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly metadatas!: pulumi.Output<outputs.RuleOtdtMetadata[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly parameters!: pulumi.Output<outputs.RuleOtdtParameter[] | undefined>;
    public readonly popularity!: pulumi.Output<number>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly risk!: pulumi.Output<number>;
    public readonly technology!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly vendor!: pulumi.Output<string>;
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a RuleOtdt resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleOtdtArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleOtdtArgs | RuleOtdtState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleOtdtState | undefined;
            resourceInputs["behavior"] = state ? state.behavior : undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["metadatas"] = state ? state.metadatas : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["popularity"] = state ? state.popularity : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["risk"] = state ? state.risk : undefined;
            resourceInputs["technology"] = state ? state.technology : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as RuleOtdtArgs | undefined;
            resourceInputs["behavior"] = args ? args.behavior : undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["metadatas"] = args ? args.metadatas : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["popularity"] = args ? args.popularity : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["risk"] = args ? args.risk : undefined;
            resourceInputs["technology"] = args ? args.technology : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleOtdt.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleOtdt resources.
 */
export interface RuleOtdtState {
    behavior?: pulumi.Input<string>;
    category?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.RuleOtdtMetadata>[]>;
    name?: pulumi.Input<string>;
    parameters?: pulumi.Input<pulumi.Input<inputs.RuleOtdtParameter>[]>;
    popularity?: pulumi.Input<number>;
    protocol?: pulumi.Input<string>;
    risk?: pulumi.Input<number>;
    technology?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vendor?: pulumi.Input<string>;
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RuleOtdt resource.
 */
export interface RuleOtdtArgs {
    behavior?: pulumi.Input<string>;
    category?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.RuleOtdtMetadata>[]>;
    name?: pulumi.Input<string>;
    parameters?: pulumi.Input<pulumi.Input<inputs.RuleOtdtParameter>[]>;
    popularity?: pulumi.Input<number>;
    protocol?: pulumi.Input<string>;
    risk?: pulumi.Input<number>;
    technology?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vendor?: pulumi.Input<string>;
    weight?: pulumi.Input<number>;
}