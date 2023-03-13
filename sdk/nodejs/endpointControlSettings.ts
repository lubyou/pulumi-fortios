// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class EndpointControlSettings extends pulumi.CustomResource {
    /**
     * Get an existing EndpointControlSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointControlSettingsState, opts?: pulumi.CustomResourceOptions): EndpointControlSettings {
        return new EndpointControlSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/endpointControlSettings:EndpointControlSettings';

    /**
     * Returns true if the given object is an instance of EndpointControlSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointControlSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointControlSettings.__pulumiType;
    }

    public readonly downloadCustomLink!: pulumi.Output<string>;
    public readonly downloadLocation!: pulumi.Output<string>;
    public readonly forticlientAvdbUpdateInterval!: pulumi.Output<number>;
    public readonly forticlientDeregUnsupportedClient!: pulumi.Output<string>;
    public readonly forticlientDisconnectUnsupportedClient!: pulumi.Output<string>;
    public readonly forticlientEmsRestApiCallTimeout!: pulumi.Output<number>;
    public readonly forticlientKeepaliveInterval!: pulumi.Output<number>;
    public readonly forticlientOfflineGrace!: pulumi.Output<string>;
    public readonly forticlientOfflineGraceInterval!: pulumi.Output<number>;
    public readonly forticlientRegKey!: pulumi.Output<string | undefined>;
    public readonly forticlientRegKeyEnforce!: pulumi.Output<string>;
    public readonly forticlientRegTimeout!: pulumi.Output<number>;
    public readonly forticlientSysUpdateInterval!: pulumi.Output<number>;
    public readonly forticlientUserAvatar!: pulumi.Output<string>;
    public readonly forticlientWarningInterval!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a EndpointControlSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointControlSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointControlSettingsArgs | EndpointControlSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointControlSettingsState | undefined;
            resourceInputs["downloadCustomLink"] = state ? state.downloadCustomLink : undefined;
            resourceInputs["downloadLocation"] = state ? state.downloadLocation : undefined;
            resourceInputs["forticlientAvdbUpdateInterval"] = state ? state.forticlientAvdbUpdateInterval : undefined;
            resourceInputs["forticlientDeregUnsupportedClient"] = state ? state.forticlientDeregUnsupportedClient : undefined;
            resourceInputs["forticlientDisconnectUnsupportedClient"] = state ? state.forticlientDisconnectUnsupportedClient : undefined;
            resourceInputs["forticlientEmsRestApiCallTimeout"] = state ? state.forticlientEmsRestApiCallTimeout : undefined;
            resourceInputs["forticlientKeepaliveInterval"] = state ? state.forticlientKeepaliveInterval : undefined;
            resourceInputs["forticlientOfflineGrace"] = state ? state.forticlientOfflineGrace : undefined;
            resourceInputs["forticlientOfflineGraceInterval"] = state ? state.forticlientOfflineGraceInterval : undefined;
            resourceInputs["forticlientRegKey"] = state ? state.forticlientRegKey : undefined;
            resourceInputs["forticlientRegKeyEnforce"] = state ? state.forticlientRegKeyEnforce : undefined;
            resourceInputs["forticlientRegTimeout"] = state ? state.forticlientRegTimeout : undefined;
            resourceInputs["forticlientSysUpdateInterval"] = state ? state.forticlientSysUpdateInterval : undefined;
            resourceInputs["forticlientUserAvatar"] = state ? state.forticlientUserAvatar : undefined;
            resourceInputs["forticlientWarningInterval"] = state ? state.forticlientWarningInterval : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as EndpointControlSettingsArgs | undefined;
            resourceInputs["downloadCustomLink"] = args ? args.downloadCustomLink : undefined;
            resourceInputs["downloadLocation"] = args ? args.downloadLocation : undefined;
            resourceInputs["forticlientAvdbUpdateInterval"] = args ? args.forticlientAvdbUpdateInterval : undefined;
            resourceInputs["forticlientDeregUnsupportedClient"] = args ? args.forticlientDeregUnsupportedClient : undefined;
            resourceInputs["forticlientDisconnectUnsupportedClient"] = args ? args.forticlientDisconnectUnsupportedClient : undefined;
            resourceInputs["forticlientEmsRestApiCallTimeout"] = args ? args.forticlientEmsRestApiCallTimeout : undefined;
            resourceInputs["forticlientKeepaliveInterval"] = args ? args.forticlientKeepaliveInterval : undefined;
            resourceInputs["forticlientOfflineGrace"] = args ? args.forticlientOfflineGrace : undefined;
            resourceInputs["forticlientOfflineGraceInterval"] = args ? args.forticlientOfflineGraceInterval : undefined;
            resourceInputs["forticlientRegKey"] = args?.forticlientRegKey ? pulumi.secret(args.forticlientRegKey) : undefined;
            resourceInputs["forticlientRegKeyEnforce"] = args ? args.forticlientRegKeyEnforce : undefined;
            resourceInputs["forticlientRegTimeout"] = args ? args.forticlientRegTimeout : undefined;
            resourceInputs["forticlientSysUpdateInterval"] = args ? args.forticlientSysUpdateInterval : undefined;
            resourceInputs["forticlientUserAvatar"] = args ? args.forticlientUserAvatar : undefined;
            resourceInputs["forticlientWarningInterval"] = args ? args.forticlientWarningInterval : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["forticlientRegKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(EndpointControlSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointControlSettings resources.
 */
export interface EndpointControlSettingsState {
    downloadCustomLink?: pulumi.Input<string>;
    downloadLocation?: pulumi.Input<string>;
    forticlientAvdbUpdateInterval?: pulumi.Input<number>;
    forticlientDeregUnsupportedClient?: pulumi.Input<string>;
    forticlientDisconnectUnsupportedClient?: pulumi.Input<string>;
    forticlientEmsRestApiCallTimeout?: pulumi.Input<number>;
    forticlientKeepaliveInterval?: pulumi.Input<number>;
    forticlientOfflineGrace?: pulumi.Input<string>;
    forticlientOfflineGraceInterval?: pulumi.Input<number>;
    forticlientRegKey?: pulumi.Input<string>;
    forticlientRegKeyEnforce?: pulumi.Input<string>;
    forticlientRegTimeout?: pulumi.Input<number>;
    forticlientSysUpdateInterval?: pulumi.Input<number>;
    forticlientUserAvatar?: pulumi.Input<string>;
    forticlientWarningInterval?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointControlSettings resource.
 */
export interface EndpointControlSettingsArgs {
    downloadCustomLink?: pulumi.Input<string>;
    downloadLocation?: pulumi.Input<string>;
    forticlientAvdbUpdateInterval?: pulumi.Input<number>;
    forticlientDeregUnsupportedClient?: pulumi.Input<string>;
    forticlientDisconnectUnsupportedClient?: pulumi.Input<string>;
    forticlientEmsRestApiCallTimeout?: pulumi.Input<number>;
    forticlientKeepaliveInterval?: pulumi.Input<number>;
    forticlientOfflineGrace?: pulumi.Input<string>;
    forticlientOfflineGraceInterval?: pulumi.Input<number>;
    forticlientRegKey?: pulumi.Input<string>;
    forticlientRegKeyEnforce?: pulumi.Input<string>;
    forticlientRegTimeout?: pulumi.Input<number>;
    forticlientSysUpdateInterval?: pulumi.Input<number>;
    forticlientUserAvatar?: pulumi.Input<string>;
    forticlientWarningInterval?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
