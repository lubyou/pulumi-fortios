// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure application signatures.
 *
 * ## Import
 *
 * Application Name can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/applicationName:ApplicationName labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class ApplicationName extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationNameState, opts?: pulumi.CustomResourceOptions): ApplicationName {
        return new ApplicationName(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/applicationName:ApplicationName';

    /**
     * Returns true if the given object is an instance of ApplicationName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationName.__pulumiType;
    }

    /**
     * Application behavior.
     */
    public readonly behavior!: pulumi.Output<string>;
    /**
     * Application category ID.
     */
    public readonly category!: pulumi.Output<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Application ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Meta data. The structure of `metadata` block is documented below.
     */
    public readonly metadatas!: pulumi.Output<outputs.ApplicationNameMetadata[] | undefined>;
    /**
     * Parameter name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Application parameter name.
     */
    public readonly parameter!: pulumi.Output<string>;
    /**
     * Application parameters. The structure of `parameters` block is documented below.
     */
    public readonly parameters!: pulumi.Output<outputs.ApplicationNameParameter[] | undefined>;
    /**
     * Application popularity.
     */
    public readonly popularity!: pulumi.Output<number>;
    /**
     * Application protocol.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Application risk.
     */
    public readonly risk!: pulumi.Output<number>;
    /**
     * Application sub-category ID.
     */
    public readonly subCategory!: pulumi.Output<number>;
    /**
     * Application technology.
     */
    public readonly technology!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Application vendor.
     */
    public readonly vendor!: pulumi.Output<string>;
    /**
     * Application weight.
     */
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a ApplicationName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationNameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationNameArgs | ApplicationNameState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationNameState | undefined;
            resourceInputs["behavior"] = state ? state.behavior : undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["metadatas"] = state ? state.metadatas : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameter"] = state ? state.parameter : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["popularity"] = state ? state.popularity : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["risk"] = state ? state.risk : undefined;
            resourceInputs["subCategory"] = state ? state.subCategory : undefined;
            resourceInputs["technology"] = state ? state.technology : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as ApplicationNameArgs | undefined;
            if ((!args || args.category === undefined) && !opts.urn) {
                throw new Error("Missing required property 'category'");
            }
            resourceInputs["behavior"] = args ? args.behavior : undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["metadatas"] = args ? args.metadatas : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameter"] = args ? args.parameter : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["popularity"] = args ? args.popularity : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["risk"] = args ? args.risk : undefined;
            resourceInputs["subCategory"] = args ? args.subCategory : undefined;
            resourceInputs["technology"] = args ? args.technology : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationName.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationName resources.
 */
export interface ApplicationNameState {
    /**
     * Application behavior.
     */
    behavior?: pulumi.Input<string>;
    /**
     * Application category ID.
     */
    category?: pulumi.Input<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Application ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Meta data. The structure of `metadata` block is documented below.
     */
    metadatas?: pulumi.Input<pulumi.Input<inputs.ApplicationNameMetadata>[]>;
    /**
     * Parameter name.
     */
    name?: pulumi.Input<string>;
    /**
     * Application parameter name.
     */
    parameter?: pulumi.Input<string>;
    /**
     * Application parameters. The structure of `parameters` block is documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.ApplicationNameParameter>[]>;
    /**
     * Application popularity.
     */
    popularity?: pulumi.Input<number>;
    /**
     * Application protocol.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Application risk.
     */
    risk?: pulumi.Input<number>;
    /**
     * Application sub-category ID.
     */
    subCategory?: pulumi.Input<number>;
    /**
     * Application technology.
     */
    technology?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Application vendor.
     */
    vendor?: pulumi.Input<string>;
    /**
     * Application weight.
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ApplicationName resource.
 */
export interface ApplicationNameArgs {
    /**
     * Application behavior.
     */
    behavior?: pulumi.Input<string>;
    /**
     * Application category ID.
     */
    category: pulumi.Input<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Application ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Meta data. The structure of `metadata` block is documented below.
     */
    metadatas?: pulumi.Input<pulumi.Input<inputs.ApplicationNameMetadata>[]>;
    /**
     * Parameter name.
     */
    name?: pulumi.Input<string>;
    /**
     * Application parameter name.
     */
    parameter?: pulumi.Input<string>;
    /**
     * Application parameters. The structure of `parameters` block is documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.ApplicationNameParameter>[]>;
    /**
     * Application popularity.
     */
    popularity?: pulumi.Input<number>;
    /**
     * Application protocol.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Application risk.
     */
    risk?: pulumi.Input<number>;
    /**
     * Application sub-category ID.
     */
    subCategory?: pulumi.Input<number>;
    /**
     * Application technology.
     */
    technology?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Application vendor.
     */
    vendor?: pulumi.Input<string>;
    /**
     * Application weight.
     */
    weight?: pulumi.Input<number>;
}
