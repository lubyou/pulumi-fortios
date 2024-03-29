// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class WirelessControllerWagProfile extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerWagProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerWagProfileState, opts?: pulumi.CustomResourceOptions): WirelessControllerWagProfile {
        return new WirelessControllerWagProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerWagProfile:WirelessControllerWagProfile';

    /**
     * Returns true if the given object is an instance of WirelessControllerWagProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerWagProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerWagProfile.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dhcpIpAddr!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly pingInterval!: pulumi.Output<number>;
    public readonly pingNumber!: pulumi.Output<number>;
    public readonly returnPacketTimeout!: pulumi.Output<number>;
    public readonly tunnelType!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly wagIp!: pulumi.Output<string>;
    public readonly wagPort!: pulumi.Output<number>;

    /**
     * Create a WirelessControllerWagProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerWagProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerWagProfileArgs | WirelessControllerWagProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerWagProfileState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dhcpIpAddr"] = state ? state.dhcpIpAddr : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pingInterval"] = state ? state.pingInterval : undefined;
            resourceInputs["pingNumber"] = state ? state.pingNumber : undefined;
            resourceInputs["returnPacketTimeout"] = state ? state.returnPacketTimeout : undefined;
            resourceInputs["tunnelType"] = state ? state.tunnelType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wagIp"] = state ? state.wagIp : undefined;
            resourceInputs["wagPort"] = state ? state.wagPort : undefined;
        } else {
            const args = argsOrState as WirelessControllerWagProfileArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dhcpIpAddr"] = args ? args.dhcpIpAddr : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pingInterval"] = args ? args.pingInterval : undefined;
            resourceInputs["pingNumber"] = args ? args.pingNumber : undefined;
            resourceInputs["returnPacketTimeout"] = args ? args.returnPacketTimeout : undefined;
            resourceInputs["tunnelType"] = args ? args.tunnelType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wagIp"] = args ? args.wagIp : undefined;
            resourceInputs["wagPort"] = args ? args.wagPort : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerWagProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerWagProfile resources.
 */
export interface WirelessControllerWagProfileState {
    comment?: pulumi.Input<string>;
    dhcpIpAddr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    pingInterval?: pulumi.Input<number>;
    pingNumber?: pulumi.Input<number>;
    returnPacketTimeout?: pulumi.Input<number>;
    tunnelType?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    wagIp?: pulumi.Input<string>;
    wagPort?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a WirelessControllerWagProfile resource.
 */
export interface WirelessControllerWagProfileArgs {
    comment?: pulumi.Input<string>;
    dhcpIpAddr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    pingInterval?: pulumi.Input<number>;
    pingNumber?: pulumi.Input<number>;
    returnPacketTimeout?: pulumi.Input<number>;
    tunnelType?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    wagIp?: pulumi.Input<string>;
    wagPort?: pulumi.Input<number>;
}
