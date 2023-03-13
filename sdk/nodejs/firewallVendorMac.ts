// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FirewallVendorMac extends pulumi.CustomResource {
    /**
     * Get an existing FirewallVendorMac resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallVendorMacState, opts?: pulumi.CustomResourceOptions): FirewallVendorMac {
        return new FirewallVendorMac(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallVendorMac:FirewallVendorMac';

    /**
     * Returns true if the given object is an instance of FirewallVendorMac.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallVendorMac {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallVendorMac.__pulumiType;
    }

    public readonly fosid!: pulumi.Output<number>;
    public readonly macNumber!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly obsolete!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallVendorMac resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallVendorMacArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallVendorMacArgs | FirewallVendorMacState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallVendorMacState | undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["macNumber"] = state ? state.macNumber : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["obsolete"] = state ? state.obsolete : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallVendorMacArgs | undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["macNumber"] = args ? args.macNumber : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["obsolete"] = args ? args.obsolete : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallVendorMac.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallVendorMac resources.
 */
export interface FirewallVendorMacState {
    fosid?: pulumi.Input<number>;
    macNumber?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    obsolete?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallVendorMac resource.
 */
export interface FirewallVendorMacArgs {
    fosid?: pulumi.Input<number>;
    macNumber?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    obsolete?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
