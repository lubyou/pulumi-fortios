// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SystemFabricVpn extends pulumi.CustomResource {
    /**
     * Get an existing SystemFabricVpn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemFabricVpnState, opts?: pulumi.CustomResourceOptions): SystemFabricVpn {
        return new SystemFabricVpn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemFabricVpn:SystemFabricVpn';

    /**
     * Returns true if the given object is an instance of SystemFabricVpn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemFabricVpn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemFabricVpn.__pulumiType;
    }

    public readonly advertisedSubnets!: pulumi.Output<outputs.SystemFabricVpnAdvertisedSubnet[] | undefined>;
    public readonly bgpAs!: pulumi.Output<number>;
    public readonly branchName!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly healthChecks!: pulumi.Output<string>;
    public readonly loopbackAddressBlock!: pulumi.Output<string>;
    public readonly loopbackAdvertisedSubnet!: pulumi.Output<number>;
    public readonly loopbackInterface!: pulumi.Output<string>;
    public readonly overlays!: pulumi.Output<outputs.SystemFabricVpnOverlay[] | undefined>;
    public readonly policyRule!: pulumi.Output<string>;
    public readonly psksecret!: pulumi.Output<string | undefined>;
    public readonly sdwanZone!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly syncMode!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly vpnRole!: pulumi.Output<string>;

    /**
     * Create a SystemFabricVpn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemFabricVpnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemFabricVpnArgs | SystemFabricVpnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemFabricVpnState | undefined;
            resourceInputs["advertisedSubnets"] = state ? state.advertisedSubnets : undefined;
            resourceInputs["bgpAs"] = state ? state.bgpAs : undefined;
            resourceInputs["branchName"] = state ? state.branchName : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["healthChecks"] = state ? state.healthChecks : undefined;
            resourceInputs["loopbackAddressBlock"] = state ? state.loopbackAddressBlock : undefined;
            resourceInputs["loopbackAdvertisedSubnet"] = state ? state.loopbackAdvertisedSubnet : undefined;
            resourceInputs["loopbackInterface"] = state ? state.loopbackInterface : undefined;
            resourceInputs["overlays"] = state ? state.overlays : undefined;
            resourceInputs["policyRule"] = state ? state.policyRule : undefined;
            resourceInputs["psksecret"] = state ? state.psksecret : undefined;
            resourceInputs["sdwanZone"] = state ? state.sdwanZone : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["syncMode"] = state ? state.syncMode : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vpnRole"] = state ? state.vpnRole : undefined;
        } else {
            const args = argsOrState as SystemFabricVpnArgs | undefined;
            resourceInputs["advertisedSubnets"] = args ? args.advertisedSubnets : undefined;
            resourceInputs["bgpAs"] = args ? args.bgpAs : undefined;
            resourceInputs["branchName"] = args ? args.branchName : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["healthChecks"] = args ? args.healthChecks : undefined;
            resourceInputs["loopbackAddressBlock"] = args ? args.loopbackAddressBlock : undefined;
            resourceInputs["loopbackAdvertisedSubnet"] = args ? args.loopbackAdvertisedSubnet : undefined;
            resourceInputs["loopbackInterface"] = args ? args.loopbackInterface : undefined;
            resourceInputs["overlays"] = args ? args.overlays : undefined;
            resourceInputs["policyRule"] = args ? args.policyRule : undefined;
            resourceInputs["psksecret"] = args ? args.psksecret : undefined;
            resourceInputs["sdwanZone"] = args ? args.sdwanZone : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["syncMode"] = args ? args.syncMode : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vpnRole"] = args ? args.vpnRole : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemFabricVpn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemFabricVpn resources.
 */
export interface SystemFabricVpnState {
    advertisedSubnets?: pulumi.Input<pulumi.Input<inputs.SystemFabricVpnAdvertisedSubnet>[]>;
    bgpAs?: pulumi.Input<number>;
    branchName?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    healthChecks?: pulumi.Input<string>;
    loopbackAddressBlock?: pulumi.Input<string>;
    loopbackAdvertisedSubnet?: pulumi.Input<number>;
    loopbackInterface?: pulumi.Input<string>;
    overlays?: pulumi.Input<pulumi.Input<inputs.SystemFabricVpnOverlay>[]>;
    policyRule?: pulumi.Input<string>;
    psksecret?: pulumi.Input<string>;
    sdwanZone?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    syncMode?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vpnRole?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemFabricVpn resource.
 */
export interface SystemFabricVpnArgs {
    advertisedSubnets?: pulumi.Input<pulumi.Input<inputs.SystemFabricVpnAdvertisedSubnet>[]>;
    bgpAs?: pulumi.Input<number>;
    branchName?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    healthChecks?: pulumi.Input<string>;
    loopbackAddressBlock?: pulumi.Input<string>;
    loopbackAdvertisedSubnet?: pulumi.Input<number>;
    loopbackInterface?: pulumi.Input<string>;
    overlays?: pulumi.Input<pulumi.Input<inputs.SystemFabricVpnOverlay>[]>;
    policyRule?: pulumi.Input<string>;
    psksecret?: pulumi.Input<string>;
    sdwanZone?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    syncMode?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vpnRole?: pulumi.Input<string>;
}
