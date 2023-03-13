// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class RouterAccessList6 extends pulumi.CustomResource {
    /**
     * Get an existing RouterAccessList6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterAccessList6State, opts?: pulumi.CustomResourceOptions): RouterAccessList6 {
        return new RouterAccessList6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerAccessList6:RouterAccessList6';

    /**
     * Returns true if the given object is an instance of RouterAccessList6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterAccessList6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterAccessList6.__pulumiType;
    }

    public readonly comments!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly rules!: pulumi.Output<outputs.RouterAccessList6Rule[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterAccessList6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterAccessList6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterAccessList6Args | RouterAccessList6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterAccessList6State | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterAccessList6Args | undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterAccessList6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterAccessList6 resources.
 */
export interface RouterAccessList6State {
    comments?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.RouterAccessList6Rule>[]>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterAccessList6 resource.
 */
export interface RouterAccessList6Args {
    comments?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.RouterAccessList6Rule>[]>;
    vdomparam?: pulumi.Input<string>;
}
