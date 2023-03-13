// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemDedicatedMgmt extends pulumi.CustomResource {
    /**
     * Get an existing SystemDedicatedMgmt resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemDedicatedMgmtState, opts?: pulumi.CustomResourceOptions): SystemDedicatedMgmt {
        return new SystemDedicatedMgmt(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemDedicatedMgmt:SystemDedicatedMgmt';

    /**
     * Returns true if the given object is an instance of SystemDedicatedMgmt.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemDedicatedMgmt {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemDedicatedMgmt.__pulumiType;
    }

    public readonly defaultGateway!: pulumi.Output<string>;
    public readonly dhcpEndIp!: pulumi.Output<string>;
    public readonly dhcpNetmask!: pulumi.Output<string>;
    public readonly dhcpServer!: pulumi.Output<string>;
    public readonly dhcpStartIp!: pulumi.Output<string>;
    public readonly interface!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemDedicatedMgmt resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemDedicatedMgmtArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemDedicatedMgmtArgs | SystemDedicatedMgmtState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemDedicatedMgmtState | undefined;
            resourceInputs["defaultGateway"] = state ? state.defaultGateway : undefined;
            resourceInputs["dhcpEndIp"] = state ? state.dhcpEndIp : undefined;
            resourceInputs["dhcpNetmask"] = state ? state.dhcpNetmask : undefined;
            resourceInputs["dhcpServer"] = state ? state.dhcpServer : undefined;
            resourceInputs["dhcpStartIp"] = state ? state.dhcpStartIp : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemDedicatedMgmtArgs | undefined;
            resourceInputs["defaultGateway"] = args ? args.defaultGateway : undefined;
            resourceInputs["dhcpEndIp"] = args ? args.dhcpEndIp : undefined;
            resourceInputs["dhcpNetmask"] = args ? args.dhcpNetmask : undefined;
            resourceInputs["dhcpServer"] = args ? args.dhcpServer : undefined;
            resourceInputs["dhcpStartIp"] = args ? args.dhcpStartIp : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemDedicatedMgmt.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemDedicatedMgmt resources.
 */
export interface SystemDedicatedMgmtState {
    defaultGateway?: pulumi.Input<string>;
    dhcpEndIp?: pulumi.Input<string>;
    dhcpNetmask?: pulumi.Input<string>;
    dhcpServer?: pulumi.Input<string>;
    dhcpStartIp?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemDedicatedMgmt resource.
 */
export interface SystemDedicatedMgmtArgs {
    defaultGateway?: pulumi.Input<string>;
    dhcpEndIp?: pulumi.Input<string>;
    dhcpNetmask?: pulumi.Input<string>;
    dhcpServer?: pulumi.Input<string>;
    dhcpStartIp?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
