// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class WirelessControllerApStatus extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerApStatus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerApStatusState, opts?: pulumi.CustomResourceOptions): WirelessControllerApStatus {
        return new WirelessControllerApStatus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerApStatus:WirelessControllerApStatus';

    /**
     * Returns true if the given object is an instance of WirelessControllerApStatus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerApStatus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerApStatus.__pulumiType;
    }

    public readonly bssid!: pulumi.Output<string>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly ssid!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerApStatus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerApStatusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerApStatusArgs | WirelessControllerApStatusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerApStatusState | undefined;
            resourceInputs["bssid"] = state ? state.bssid : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["ssid"] = state ? state.ssid : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerApStatusArgs | undefined;
            resourceInputs["bssid"] = args ? args.bssid : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["ssid"] = args ? args.ssid : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerApStatus.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerApStatus resources.
 */
export interface WirelessControllerApStatusState {
    bssid?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    ssid?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerApStatus resource.
 */
export interface WirelessControllerApStatusArgs {
    bssid?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    ssid?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
