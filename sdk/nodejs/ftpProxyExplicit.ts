// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FtpProxyExplicit extends pulumi.CustomResource {
    /**
     * Get an existing FtpProxyExplicit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FtpProxyExplicitState, opts?: pulumi.CustomResourceOptions): FtpProxyExplicit {
        return new FtpProxyExplicit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/ftpProxyExplicit:FtpProxyExplicit';

    /**
     * Returns true if the given object is an instance of FtpProxyExplicit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FtpProxyExplicit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FtpProxyExplicit.__pulumiType;
    }

    public readonly incomingIp!: pulumi.Output<string>;
    public readonly incomingPort!: pulumi.Output<string>;
    public readonly outgoingIp!: pulumi.Output<string>;
    public readonly secDefaultAction!: pulumi.Output<string>;
    public readonly serverDataMode!: pulumi.Output<string>;
    public readonly ssl!: pulumi.Output<string>;
    public readonly sslAlgorithm!: pulumi.Output<string>;
    public readonly sslCert!: pulumi.Output<string>;
    public readonly sslDhBits!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FtpProxyExplicit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FtpProxyExplicitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FtpProxyExplicitArgs | FtpProxyExplicitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FtpProxyExplicitState | undefined;
            resourceInputs["incomingIp"] = state ? state.incomingIp : undefined;
            resourceInputs["incomingPort"] = state ? state.incomingPort : undefined;
            resourceInputs["outgoingIp"] = state ? state.outgoingIp : undefined;
            resourceInputs["secDefaultAction"] = state ? state.secDefaultAction : undefined;
            resourceInputs["serverDataMode"] = state ? state.serverDataMode : undefined;
            resourceInputs["ssl"] = state ? state.ssl : undefined;
            resourceInputs["sslAlgorithm"] = state ? state.sslAlgorithm : undefined;
            resourceInputs["sslCert"] = state ? state.sslCert : undefined;
            resourceInputs["sslDhBits"] = state ? state.sslDhBits : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FtpProxyExplicitArgs | undefined;
            resourceInputs["incomingIp"] = args ? args.incomingIp : undefined;
            resourceInputs["incomingPort"] = args ? args.incomingPort : undefined;
            resourceInputs["outgoingIp"] = args ? args.outgoingIp : undefined;
            resourceInputs["secDefaultAction"] = args ? args.secDefaultAction : undefined;
            resourceInputs["serverDataMode"] = args ? args.serverDataMode : undefined;
            resourceInputs["ssl"] = args ? args.ssl : undefined;
            resourceInputs["sslAlgorithm"] = args ? args.sslAlgorithm : undefined;
            resourceInputs["sslCert"] = args ? args.sslCert : undefined;
            resourceInputs["sslDhBits"] = args ? args.sslDhBits : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FtpProxyExplicit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FtpProxyExplicit resources.
 */
export interface FtpProxyExplicitState {
    incomingIp?: pulumi.Input<string>;
    incomingPort?: pulumi.Input<string>;
    outgoingIp?: pulumi.Input<string>;
    secDefaultAction?: pulumi.Input<string>;
    serverDataMode?: pulumi.Input<string>;
    ssl?: pulumi.Input<string>;
    sslAlgorithm?: pulumi.Input<string>;
    sslCert?: pulumi.Input<string>;
    sslDhBits?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FtpProxyExplicit resource.
 */
export interface FtpProxyExplicitArgs {
    incomingIp?: pulumi.Input<string>;
    incomingPort?: pulumi.Input<string>;
    outgoingIp?: pulumi.Input<string>;
    secDefaultAction?: pulumi.Input<string>;
    serverDataMode?: pulumi.Input<string>;
    ssl?: pulumi.Input<string>;
    sslAlgorithm?: pulumi.Input<string>;
    sslCert?: pulumi.Input<string>;
    sslDhBits?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
