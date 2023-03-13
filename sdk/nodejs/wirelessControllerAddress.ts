// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class WirelessControllerAddress extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerAddressState, opts?: pulumi.CustomResourceOptions): WirelessControllerAddress {
        return new WirelessControllerAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerAddress:WirelessControllerAddress';

    /**
     * Returns true if the given object is an instance of WirelessControllerAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerAddress.__pulumiType;
    }

    public readonly fosid!: pulumi.Output<string>;
    public readonly mac!: pulumi.Output<string>;
    public readonly policy!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerAddressArgs | WirelessControllerAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerAddressState | undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerAddressArgs | undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["mac"] = args ? args.mac : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerAddress resources.
 */
export interface WirelessControllerAddressState {
    fosid?: pulumi.Input<string>;
    mac?: pulumi.Input<string>;
    policy?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerAddress resource.
 */
export interface WirelessControllerAddressArgs {
    fosid?: pulumi.Input<string>;
    mac?: pulumi.Input<string>;
    policy?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
