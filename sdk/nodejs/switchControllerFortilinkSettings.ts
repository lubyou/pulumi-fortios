// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SwitchControllerFortilinkSettings extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerFortilinkSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerFortilinkSettingsState, opts?: pulumi.CustomResourceOptions): SwitchControllerFortilinkSettings {
        return new SwitchControllerFortilinkSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerFortilinkSettings:SwitchControllerFortilinkSettings';

    /**
     * Returns true if the given object is an instance of SwitchControllerFortilinkSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerFortilinkSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerFortilinkSettings.__pulumiType;
    }

    public readonly fortilink!: pulumi.Output<string>;
    public readonly inactiveTimer!: pulumi.Output<number>;
    public readonly linkDownFlush!: pulumi.Output<string>;
    public readonly nacPorts!: pulumi.Output<outputs.SwitchControllerFortilinkSettingsNacPorts>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerFortilinkSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerFortilinkSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerFortilinkSettingsArgs | SwitchControllerFortilinkSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerFortilinkSettingsState | undefined;
            resourceInputs["fortilink"] = state ? state.fortilink : undefined;
            resourceInputs["inactiveTimer"] = state ? state.inactiveTimer : undefined;
            resourceInputs["linkDownFlush"] = state ? state.linkDownFlush : undefined;
            resourceInputs["nacPorts"] = state ? state.nacPorts : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerFortilinkSettingsArgs | undefined;
            resourceInputs["fortilink"] = args ? args.fortilink : undefined;
            resourceInputs["inactiveTimer"] = args ? args.inactiveTimer : undefined;
            resourceInputs["linkDownFlush"] = args ? args.linkDownFlush : undefined;
            resourceInputs["nacPorts"] = args ? args.nacPorts : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerFortilinkSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerFortilinkSettings resources.
 */
export interface SwitchControllerFortilinkSettingsState {
    fortilink?: pulumi.Input<string>;
    inactiveTimer?: pulumi.Input<number>;
    linkDownFlush?: pulumi.Input<string>;
    nacPorts?: pulumi.Input<inputs.SwitchControllerFortilinkSettingsNacPorts>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerFortilinkSettings resource.
 */
export interface SwitchControllerFortilinkSettingsArgs {
    fortilink?: pulumi.Input<string>;
    inactiveTimer?: pulumi.Input<number>;
    linkDownFlush?: pulumi.Input<string>;
    nacPorts?: pulumi.Input<inputs.SwitchControllerFortilinkSettingsNacPorts>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
