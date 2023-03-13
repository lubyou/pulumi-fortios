// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class WafSignature extends pulumi.CustomResource {
    /**
     * Get an existing WafSignature resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WafSignatureState, opts?: pulumi.CustomResourceOptions): WafSignature {
        return new WafSignature(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wafSignature:WafSignature';

    /**
     * Returns true if the given object is an instance of WafSignature.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WafSignature {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WafSignature.__pulumiType;
    }

    public readonly desc!: pulumi.Output<string>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WafSignature resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WafSignatureArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WafSignatureArgs | WafSignatureState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WafSignatureState | undefined;
            resourceInputs["desc"] = state ? state.desc : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WafSignatureArgs | undefined;
            resourceInputs["desc"] = args ? args.desc : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WafSignature.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WafSignature resources.
 */
export interface WafSignatureState {
    desc?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WafSignature resource.
 */
export interface WafSignatureArgs {
    desc?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
