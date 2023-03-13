// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemAdminProfiles extends pulumi.CustomResource {
    /**
     * Get an existing SystemAdminProfiles resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemAdminProfilesState, opts?: pulumi.CustomResourceOptions): SystemAdminProfiles {
        return new SystemAdminProfiles(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemAdminProfiles:SystemAdminProfiles';

    /**
     * Returns true if the given object is an instance of SystemAdminProfiles.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemAdminProfiles {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemAdminProfiles.__pulumiType;
    }

    public readonly admintimeoutOverride!: pulumi.Output<string>;
    public readonly authgrp!: pulumi.Output<string>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly ftviewgrp!: pulumi.Output<string>;
    public readonly fwgrp!: pulumi.Output<string>;
    public readonly loggrp!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly netgrp!: pulumi.Output<string>;
    public readonly scope!: pulumi.Output<string>;
    public readonly secfabgrp!: pulumi.Output<string>;
    public readonly sysgrp!: pulumi.Output<string>;
    public readonly utmgrp!: pulumi.Output<string>;
    public readonly vpngrp!: pulumi.Output<string>;
    public readonly wanoptgrp!: pulumi.Output<string>;
    public readonly wifi!: pulumi.Output<string>;

    /**
     * Create a SystemAdminProfiles resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemAdminProfilesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemAdminProfilesArgs | SystemAdminProfilesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemAdminProfilesState | undefined;
            resourceInputs["admintimeoutOverride"] = state ? state.admintimeoutOverride : undefined;
            resourceInputs["authgrp"] = state ? state.authgrp : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["ftviewgrp"] = state ? state.ftviewgrp : undefined;
            resourceInputs["fwgrp"] = state ? state.fwgrp : undefined;
            resourceInputs["loggrp"] = state ? state.loggrp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["netgrp"] = state ? state.netgrp : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["secfabgrp"] = state ? state.secfabgrp : undefined;
            resourceInputs["sysgrp"] = state ? state.sysgrp : undefined;
            resourceInputs["utmgrp"] = state ? state.utmgrp : undefined;
            resourceInputs["vpngrp"] = state ? state.vpngrp : undefined;
            resourceInputs["wanoptgrp"] = state ? state.wanoptgrp : undefined;
            resourceInputs["wifi"] = state ? state.wifi : undefined;
        } else {
            const args = argsOrState as SystemAdminProfilesArgs | undefined;
            resourceInputs["admintimeoutOverride"] = args ? args.admintimeoutOverride : undefined;
            resourceInputs["authgrp"] = args ? args.authgrp : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["ftviewgrp"] = args ? args.ftviewgrp : undefined;
            resourceInputs["fwgrp"] = args ? args.fwgrp : undefined;
            resourceInputs["loggrp"] = args ? args.loggrp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["netgrp"] = args ? args.netgrp : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["secfabgrp"] = args ? args.secfabgrp : undefined;
            resourceInputs["sysgrp"] = args ? args.sysgrp : undefined;
            resourceInputs["utmgrp"] = args ? args.utmgrp : undefined;
            resourceInputs["vpngrp"] = args ? args.vpngrp : undefined;
            resourceInputs["wanoptgrp"] = args ? args.wanoptgrp : undefined;
            resourceInputs["wifi"] = args ? args.wifi : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemAdminProfiles.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemAdminProfiles resources.
 */
export interface SystemAdminProfilesState {
    admintimeoutOverride?: pulumi.Input<string>;
    authgrp?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    ftviewgrp?: pulumi.Input<string>;
    fwgrp?: pulumi.Input<string>;
    loggrp?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    netgrp?: pulumi.Input<string>;
    scope?: pulumi.Input<string>;
    secfabgrp?: pulumi.Input<string>;
    sysgrp?: pulumi.Input<string>;
    utmgrp?: pulumi.Input<string>;
    vpngrp?: pulumi.Input<string>;
    wanoptgrp?: pulumi.Input<string>;
    wifi?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemAdminProfiles resource.
 */
export interface SystemAdminProfilesArgs {
    admintimeoutOverride?: pulumi.Input<string>;
    authgrp?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    ftviewgrp?: pulumi.Input<string>;
    fwgrp?: pulumi.Input<string>;
    loggrp?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    netgrp?: pulumi.Input<string>;
    scope?: pulumi.Input<string>;
    secfabgrp?: pulumi.Input<string>;
    sysgrp?: pulumi.Input<string>;
    utmgrp?: pulumi.Input<string>;
    vpngrp?: pulumi.Input<string>;
    wanoptgrp?: pulumi.Input<string>;
    wifi?: pulumi.Input<string>;
}
