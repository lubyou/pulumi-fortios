// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FortimanagerFirewallObjectAddress extends pulumi.CustomResource {
    /**
     * Get an existing FortimanagerFirewallObjectAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerFirewallObjectAddressState, opts?: pulumi.CustomResourceOptions): FortimanagerFirewallObjectAddress {
        return new FortimanagerFirewallObjectAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fortimanagerFirewallObjectAddress:FortimanagerFirewallObjectAddress';

    /**
     * Returns true if the given object is an instance of FortimanagerFirewallObjectAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FortimanagerFirewallObjectAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FortimanagerFirewallObjectAddress.__pulumiType;
    }

    public readonly adom!: pulumi.Output<string | undefined>;
    public readonly allowRouting!: pulumi.Output<string | undefined>;
    public readonly associatedIntf!: pulumi.Output<string | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly endIp!: pulumi.Output<string | undefined>;
    public readonly fqdn!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly startIp!: pulumi.Output<string | undefined>;
    public readonly subnet!: pulumi.Output<string | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a FortimanagerFirewallObjectAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FortimanagerFirewallObjectAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerFirewallObjectAddressArgs | FortimanagerFirewallObjectAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerFirewallObjectAddressState | undefined;
            resourceInputs["adom"] = state ? state.adom : undefined;
            resourceInputs["allowRouting"] = state ? state.allowRouting : undefined;
            resourceInputs["associatedIntf"] = state ? state.associatedIntf : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["endIp"] = state ? state.endIp : undefined;
            resourceInputs["fqdn"] = state ? state.fqdn : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["startIp"] = state ? state.startIp : undefined;
            resourceInputs["subnet"] = state ? state.subnet : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FortimanagerFirewallObjectAddressArgs | undefined;
            resourceInputs["adom"] = args ? args.adom : undefined;
            resourceInputs["allowRouting"] = args ? args.allowRouting : undefined;
            resourceInputs["associatedIntf"] = args ? args.associatedIntf : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["endIp"] = args ? args.endIp : undefined;
            resourceInputs["fqdn"] = args ? args.fqdn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["startIp"] = args ? args.startIp : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FortimanagerFirewallObjectAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FortimanagerFirewallObjectAddress resources.
 */
export interface FortimanagerFirewallObjectAddressState {
    adom?: pulumi.Input<string>;
    allowRouting?: pulumi.Input<string>;
    associatedIntf?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    endIp?: pulumi.Input<string>;
    fqdn?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    startIp?: pulumi.Input<string>;
    subnet?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FortimanagerFirewallObjectAddress resource.
 */
export interface FortimanagerFirewallObjectAddressArgs {
    adom?: pulumi.Input<string>;
    allowRouting?: pulumi.Input<string>;
    associatedIntf?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    endIp?: pulumi.Input<string>;
    fqdn?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    startIp?: pulumi.Input<string>;
    subnet?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
