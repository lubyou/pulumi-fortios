// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemVdom extends pulumi.CustomResource {
    /**
     * Get an existing SystemVdom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemVdomState, opts?: pulumi.CustomResourceOptions): SystemVdom {
        return new SystemVdom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemVdom:SystemVdom';

    /**
     * Returns true if the given object is an instance of SystemVdom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemVdom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemVdom.__pulumiType;
    }

    public readonly flag!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly shortName!: pulumi.Output<string>;
    public readonly temporary!: pulumi.Output<number>;
    public readonly vclusterId!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemVdom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemVdomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemVdomArgs | SystemVdomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemVdomState | undefined;
            resourceInputs["flag"] = state ? state.flag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["shortName"] = state ? state.shortName : undefined;
            resourceInputs["temporary"] = state ? state.temporary : undefined;
            resourceInputs["vclusterId"] = state ? state.vclusterId : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemVdomArgs | undefined;
            resourceInputs["flag"] = args ? args.flag : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["shortName"] = args ? args.shortName : undefined;
            resourceInputs["temporary"] = args ? args.temporary : undefined;
            resourceInputs["vclusterId"] = args ? args.vclusterId : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemVdom.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemVdom resources.
 */
export interface SystemVdomState {
    flag?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    shortName?: pulumi.Input<string>;
    temporary?: pulumi.Input<number>;
    vclusterId?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemVdom resource.
 */
export interface SystemVdomArgs {
    flag?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    shortName?: pulumi.Input<string>;
    temporary?: pulumi.Input<number>;
    vclusterId?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
