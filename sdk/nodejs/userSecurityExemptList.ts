// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class UserSecurityExemptList extends pulumi.CustomResource {
    /**
     * Get an existing UserSecurityExemptList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserSecurityExemptListState, opts?: pulumi.CustomResourceOptions): UserSecurityExemptList {
        return new UserSecurityExemptList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/userSecurityExemptList:UserSecurityExemptList';

    /**
     * Returns true if the given object is an instance of UserSecurityExemptList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserSecurityExemptList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserSecurityExemptList.__pulumiType;
    }

    public readonly description!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly rules!: pulumi.Output<outputs.UserSecurityExemptListRule[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a UserSecurityExemptList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserSecurityExemptListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserSecurityExemptListArgs | UserSecurityExemptListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserSecurityExemptListState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserSecurityExemptListArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserSecurityExemptList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserSecurityExemptList resources.
 */
export interface UserSecurityExemptListState {
    description?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.UserSecurityExemptListRule>[]>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserSecurityExemptList resource.
 */
export interface UserSecurityExemptListArgs {
    description?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.UserSecurityExemptListRule>[]>;
    vdomparam?: pulumi.Input<string>;
}
