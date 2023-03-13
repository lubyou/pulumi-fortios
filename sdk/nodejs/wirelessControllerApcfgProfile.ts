// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WirelessControllerApcfgProfile extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerApcfgProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerApcfgProfileState, opts?: pulumi.CustomResourceOptions): WirelessControllerApcfgProfile {
        return new WirelessControllerApcfgProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerApcfgProfile:WirelessControllerApcfgProfile';

    /**
     * Returns true if the given object is an instance of WirelessControllerApcfgProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerApcfgProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerApcfgProfile.__pulumiType;
    }

    public readonly acIp!: pulumi.Output<string>;
    public readonly acPort!: pulumi.Output<number>;
    public readonly acTimer!: pulumi.Output<number>;
    public readonly acType!: pulumi.Output<string>;
    public readonly apFamily!: pulumi.Output<string>;
    public readonly commandLists!: pulumi.Output<outputs.WirelessControllerApcfgProfileCommandList[] | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerApcfgProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerApcfgProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerApcfgProfileArgs | WirelessControllerApcfgProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerApcfgProfileState | undefined;
            resourceInputs["acIp"] = state ? state.acIp : undefined;
            resourceInputs["acPort"] = state ? state.acPort : undefined;
            resourceInputs["acTimer"] = state ? state.acTimer : undefined;
            resourceInputs["acType"] = state ? state.acType : undefined;
            resourceInputs["apFamily"] = state ? state.apFamily : undefined;
            resourceInputs["commandLists"] = state ? state.commandLists : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerApcfgProfileArgs | undefined;
            resourceInputs["acIp"] = args ? args.acIp : undefined;
            resourceInputs["acPort"] = args ? args.acPort : undefined;
            resourceInputs["acTimer"] = args ? args.acTimer : undefined;
            resourceInputs["acType"] = args ? args.acType : undefined;
            resourceInputs["apFamily"] = args ? args.apFamily : undefined;
            resourceInputs["commandLists"] = args ? args.commandLists : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerApcfgProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerApcfgProfile resources.
 */
export interface WirelessControllerApcfgProfileState {
    acIp?: pulumi.Input<string>;
    acPort?: pulumi.Input<number>;
    acTimer?: pulumi.Input<number>;
    acType?: pulumi.Input<string>;
    apFamily?: pulumi.Input<string>;
    commandLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerApcfgProfileCommandList>[]>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerApcfgProfile resource.
 */
export interface WirelessControllerApcfgProfileArgs {
    acIp?: pulumi.Input<string>;
    acPort?: pulumi.Input<number>;
    acTimer?: pulumi.Input<number>;
    acType?: pulumi.Input<string>;
    apFamily?: pulumi.Input<string>;
    commandLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerApcfgProfileCommandList>[]>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
