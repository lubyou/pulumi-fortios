// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FortimanagerFirewallObjectIppool extends pulumi.CustomResource {
    /**
     * Get an existing FortimanagerFirewallObjectIppool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerFirewallObjectIppoolState, opts?: pulumi.CustomResourceOptions): FortimanagerFirewallObjectIppool {
        return new FortimanagerFirewallObjectIppool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fortimanagerFirewallObjectIppool:FortimanagerFirewallObjectIppool';

    /**
     * Returns true if the given object is an instance of FortimanagerFirewallObjectIppool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FortimanagerFirewallObjectIppool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FortimanagerFirewallObjectIppool.__pulumiType;
    }

    public readonly adom!: pulumi.Output<string | undefined>;
    public readonly arpIntf!: pulumi.Output<string | undefined>;
    public readonly arpReply!: pulumi.Output<string | undefined>;
    public readonly associatedIntf!: pulumi.Output<string | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly endip!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly startip!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a FortimanagerFirewallObjectIppool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FortimanagerFirewallObjectIppoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerFirewallObjectIppoolArgs | FortimanagerFirewallObjectIppoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerFirewallObjectIppoolState | undefined;
            resourceInputs["adom"] = state ? state.adom : undefined;
            resourceInputs["arpIntf"] = state ? state.arpIntf : undefined;
            resourceInputs["arpReply"] = state ? state.arpReply : undefined;
            resourceInputs["associatedIntf"] = state ? state.associatedIntf : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["endip"] = state ? state.endip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["startip"] = state ? state.startip : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FortimanagerFirewallObjectIppoolArgs | undefined;
            if ((!args || args.endip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endip'");
            }
            if ((!args || args.startip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startip'");
            }
            resourceInputs["adom"] = args ? args.adom : undefined;
            resourceInputs["arpIntf"] = args ? args.arpIntf : undefined;
            resourceInputs["arpReply"] = args ? args.arpReply : undefined;
            resourceInputs["associatedIntf"] = args ? args.associatedIntf : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["endip"] = args ? args.endip : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["startip"] = args ? args.startip : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FortimanagerFirewallObjectIppool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FortimanagerFirewallObjectIppool resources.
 */
export interface FortimanagerFirewallObjectIppoolState {
    adom?: pulumi.Input<string>;
    arpIntf?: pulumi.Input<string>;
    arpReply?: pulumi.Input<string>;
    associatedIntf?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    endip?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    startip?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FortimanagerFirewallObjectIppool resource.
 */
export interface FortimanagerFirewallObjectIppoolArgs {
    adom?: pulumi.Input<string>;
    arpIntf?: pulumi.Input<string>;
    arpReply?: pulumi.Input<string>;
    associatedIntf?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    endip: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    startip: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
