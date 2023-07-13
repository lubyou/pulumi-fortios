// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SwitchControllerVlan extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerVlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerVlanState, opts?: pulumi.CustomResourceOptions): SwitchControllerVlan {
        return new SwitchControllerVlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerVlan:SwitchControllerVlan';

    /**
     * Returns true if the given object is an instance of SwitchControllerVlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerVlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerVlan.__pulumiType;
    }

    public readonly auth!: pulumi.Output<string>;
    public readonly color!: pulumi.Output<number>;
    public readonly comments!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly portalMessageOverrideGroup!: pulumi.Output<string>;
    public readonly portalMessageOverrides!: pulumi.Output<outputs.SwitchControllerVlanPortalMessageOverrides>;
    public readonly radiusServer!: pulumi.Output<string>;
    public readonly security!: pulumi.Output<string>;
    public readonly selectedUsergroups!: pulumi.Output<outputs.SwitchControllerVlanSelectedUsergroup[] | undefined>;
    public readonly usergroup!: pulumi.Output<string>;
    public readonly vdom!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly vlanid!: pulumi.Output<number>;

    /**
     * Create a SwitchControllerVlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerVlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerVlanArgs | SwitchControllerVlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerVlanState | undefined;
            resourceInputs["auth"] = state ? state.auth : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["portalMessageOverrideGroup"] = state ? state.portalMessageOverrideGroup : undefined;
            resourceInputs["portalMessageOverrides"] = state ? state.portalMessageOverrides : undefined;
            resourceInputs["radiusServer"] = state ? state.radiusServer : undefined;
            resourceInputs["security"] = state ? state.security : undefined;
            resourceInputs["selectedUsergroups"] = state ? state.selectedUsergroups : undefined;
            resourceInputs["usergroup"] = state ? state.usergroup : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlanid"] = state ? state.vlanid : undefined;
        } else {
            const args = argsOrState as SwitchControllerVlanArgs | undefined;
            resourceInputs["auth"] = args ? args.auth : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["portalMessageOverrideGroup"] = args ? args.portalMessageOverrideGroup : undefined;
            resourceInputs["portalMessageOverrides"] = args ? args.portalMessageOverrides : undefined;
            resourceInputs["radiusServer"] = args ? args.radiusServer : undefined;
            resourceInputs["security"] = args ? args.security : undefined;
            resourceInputs["selectedUsergroups"] = args ? args.selectedUsergroups : undefined;
            resourceInputs["usergroup"] = args ? args.usergroup : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlanid"] = args ? args.vlanid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerVlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerVlan resources.
 */
export interface SwitchControllerVlanState {
    auth?: pulumi.Input<string>;
    color?: pulumi.Input<number>;
    comments?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    portalMessageOverrideGroup?: pulumi.Input<string>;
    portalMessageOverrides?: pulumi.Input<inputs.SwitchControllerVlanPortalMessageOverrides>;
    radiusServer?: pulumi.Input<string>;
    security?: pulumi.Input<string>;
    selectedUsergroups?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanSelectedUsergroup>[]>;
    usergroup?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vlanid?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SwitchControllerVlan resource.
 */
export interface SwitchControllerVlanArgs {
    auth?: pulumi.Input<string>;
    color?: pulumi.Input<number>;
    comments?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    portalMessageOverrideGroup?: pulumi.Input<string>;
    portalMessageOverrides?: pulumi.Input<inputs.SwitchControllerVlanPortalMessageOverrides>;
    radiusServer?: pulumi.Input<string>;
    security?: pulumi.Input<string>;
    selectedUsergroups?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanSelectedUsergroup>[]>;
    usergroup?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vlanid?: pulumi.Input<number>;
}
