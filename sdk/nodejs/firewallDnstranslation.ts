// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FirewallDnstranslation extends pulumi.CustomResource {
    /**
     * Get an existing FirewallDnstranslation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallDnstranslationState, opts?: pulumi.CustomResourceOptions): FirewallDnstranslation {
        return new FirewallDnstranslation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallDnstranslation:FirewallDnstranslation';

    /**
     * Returns true if the given object is an instance of FirewallDnstranslation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallDnstranslation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallDnstranslation.__pulumiType;
    }

    public readonly dst!: pulumi.Output<string>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly netmask!: pulumi.Output<string>;
    public readonly src!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallDnstranslation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallDnstranslationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallDnstranslationArgs | FirewallDnstranslationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallDnstranslationState | undefined;
            resourceInputs["dst"] = state ? state.dst : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["netmask"] = state ? state.netmask : undefined;
            resourceInputs["src"] = state ? state.src : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallDnstranslationArgs | undefined;
            resourceInputs["dst"] = args ? args.dst : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["netmask"] = args ? args.netmask : undefined;
            resourceInputs["src"] = args ? args.src : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallDnstranslation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallDnstranslation resources.
 */
export interface FirewallDnstranslationState {
    dst?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    netmask?: pulumi.Input<string>;
    src?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallDnstranslation resource.
 */
export interface FirewallDnstranslationArgs {
    dst?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    netmask?: pulumi.Input<string>;
    src?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
