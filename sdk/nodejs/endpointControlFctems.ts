// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class EndpointControlFctems extends pulumi.CustomResource {
    /**
     * Get an existing EndpointControlFctems resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointControlFctemsState, opts?: pulumi.CustomResourceOptions): EndpointControlFctems {
        return new EndpointControlFctems(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/endpointControlFctems:EndpointControlFctems';

    /**
     * Returns true if the given object is an instance of EndpointControlFctems.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointControlFctems {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointControlFctems.__pulumiType;
    }

    public readonly adminPassword!: pulumi.Output<string | undefined>;
    public readonly adminUsername!: pulumi.Output<string>;
    public readonly callTimeout!: pulumi.Output<number>;
    public readonly capabilities!: pulumi.Output<string>;
    public readonly certificate!: pulumi.Output<string>;
    public readonly cloudServerType!: pulumi.Output<string>;
    public readonly dirtyReason!: pulumi.Output<string>;
    public readonly emsId!: pulumi.Output<number>;
    public readonly fortinetoneCloudAuthentication!: pulumi.Output<string>;
    public readonly httpsPort!: pulumi.Output<number>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly outOfSyncThreshold!: pulumi.Output<number>;
    public readonly preserveSslSession!: pulumi.Output<string>;
    public readonly pullAvatars!: pulumi.Output<string>;
    public readonly pullMalwareHash!: pulumi.Output<string>;
    public readonly pullSysinfo!: pulumi.Output<string>;
    public readonly pullTags!: pulumi.Output<string>;
    public readonly pullVulnerabilities!: pulumi.Output<string>;
    public readonly serialNumber!: pulumi.Output<string>;
    public readonly server!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly statusCheckInterval!: pulumi.Output<number>;
    public readonly tenantId!: pulumi.Output<string>;
    public readonly trustCaCn!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly websocketOverride!: pulumi.Output<string>;

    /**
     * Create a EndpointControlFctems resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointControlFctemsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointControlFctemsArgs | EndpointControlFctemsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointControlFctemsState | undefined;
            resourceInputs["adminPassword"] = state ? state.adminPassword : undefined;
            resourceInputs["adminUsername"] = state ? state.adminUsername : undefined;
            resourceInputs["callTimeout"] = state ? state.callTimeout : undefined;
            resourceInputs["capabilities"] = state ? state.capabilities : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["cloudServerType"] = state ? state.cloudServerType : undefined;
            resourceInputs["dirtyReason"] = state ? state.dirtyReason : undefined;
            resourceInputs["emsId"] = state ? state.emsId : undefined;
            resourceInputs["fortinetoneCloudAuthentication"] = state ? state.fortinetoneCloudAuthentication : undefined;
            resourceInputs["httpsPort"] = state ? state.httpsPort : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outOfSyncThreshold"] = state ? state.outOfSyncThreshold : undefined;
            resourceInputs["preserveSslSession"] = state ? state.preserveSslSession : undefined;
            resourceInputs["pullAvatars"] = state ? state.pullAvatars : undefined;
            resourceInputs["pullMalwareHash"] = state ? state.pullMalwareHash : undefined;
            resourceInputs["pullSysinfo"] = state ? state.pullSysinfo : undefined;
            resourceInputs["pullTags"] = state ? state.pullTags : undefined;
            resourceInputs["pullVulnerabilities"] = state ? state.pullVulnerabilities : undefined;
            resourceInputs["serialNumber"] = state ? state.serialNumber : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusCheckInterval"] = state ? state.statusCheckInterval : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["trustCaCn"] = state ? state.trustCaCn : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["websocketOverride"] = state ? state.websocketOverride : undefined;
        } else {
            const args = argsOrState as EndpointControlFctemsArgs | undefined;
            resourceInputs["adminPassword"] = args?.adminPassword ? pulumi.secret(args.adminPassword) : undefined;
            resourceInputs["adminUsername"] = args ? args.adminUsername : undefined;
            resourceInputs["callTimeout"] = args ? args.callTimeout : undefined;
            resourceInputs["capabilities"] = args ? args.capabilities : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["cloudServerType"] = args ? args.cloudServerType : undefined;
            resourceInputs["dirtyReason"] = args ? args.dirtyReason : undefined;
            resourceInputs["emsId"] = args ? args.emsId : undefined;
            resourceInputs["fortinetoneCloudAuthentication"] = args ? args.fortinetoneCloudAuthentication : undefined;
            resourceInputs["httpsPort"] = args ? args.httpsPort : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outOfSyncThreshold"] = args ? args.outOfSyncThreshold : undefined;
            resourceInputs["preserveSslSession"] = args ? args.preserveSslSession : undefined;
            resourceInputs["pullAvatars"] = args ? args.pullAvatars : undefined;
            resourceInputs["pullMalwareHash"] = args ? args.pullMalwareHash : undefined;
            resourceInputs["pullSysinfo"] = args ? args.pullSysinfo : undefined;
            resourceInputs["pullTags"] = args ? args.pullTags : undefined;
            resourceInputs["pullVulnerabilities"] = args ? args.pullVulnerabilities : undefined;
            resourceInputs["serialNumber"] = args ? args.serialNumber : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["statusCheckInterval"] = args ? args.statusCheckInterval : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["trustCaCn"] = args ? args.trustCaCn : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["websocketOverride"] = args ? args.websocketOverride : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["adminPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(EndpointControlFctems.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointControlFctems resources.
 */
export interface EndpointControlFctemsState {
    adminPassword?: pulumi.Input<string>;
    adminUsername?: pulumi.Input<string>;
    callTimeout?: pulumi.Input<number>;
    capabilities?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    cloudServerType?: pulumi.Input<string>;
    dirtyReason?: pulumi.Input<string>;
    emsId?: pulumi.Input<number>;
    fortinetoneCloudAuthentication?: pulumi.Input<string>;
    httpsPort?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    outOfSyncThreshold?: pulumi.Input<number>;
    preserveSslSession?: pulumi.Input<string>;
    pullAvatars?: pulumi.Input<string>;
    pullMalwareHash?: pulumi.Input<string>;
    pullSysinfo?: pulumi.Input<string>;
    pullTags?: pulumi.Input<string>;
    pullVulnerabilities?: pulumi.Input<string>;
    serialNumber?: pulumi.Input<string>;
    server?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    statusCheckInterval?: pulumi.Input<number>;
    tenantId?: pulumi.Input<string>;
    trustCaCn?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    websocketOverride?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointControlFctems resource.
 */
export interface EndpointControlFctemsArgs {
    adminPassword?: pulumi.Input<string>;
    adminUsername?: pulumi.Input<string>;
    callTimeout?: pulumi.Input<number>;
    capabilities?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    cloudServerType?: pulumi.Input<string>;
    dirtyReason?: pulumi.Input<string>;
    emsId?: pulumi.Input<number>;
    fortinetoneCloudAuthentication?: pulumi.Input<string>;
    httpsPort?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    outOfSyncThreshold?: pulumi.Input<number>;
    preserveSslSession?: pulumi.Input<string>;
    pullAvatars?: pulumi.Input<string>;
    pullMalwareHash?: pulumi.Input<string>;
    pullSysinfo?: pulumi.Input<string>;
    pullTags?: pulumi.Input<string>;
    pullVulnerabilities?: pulumi.Input<string>;
    serialNumber?: pulumi.Input<string>;
    server?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    statusCheckInterval?: pulumi.Input<number>;
    tenantId?: pulumi.Input<string>;
    trustCaCn?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    websocketOverride?: pulumi.Input<string>;
}
