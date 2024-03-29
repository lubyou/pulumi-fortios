// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WebProxyExplicit extends pulumi.CustomResource {
    /**
     * Get an existing WebProxyExplicit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebProxyExplicitState, opts?: pulumi.CustomResourceOptions): WebProxyExplicit {
        return new WebProxyExplicit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webProxyExplicit:WebProxyExplicit';

    /**
     * Returns true if the given object is an instance of WebProxyExplicit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebProxyExplicit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebProxyExplicit.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly ftpIncomingPort!: pulumi.Output<string>;
    public readonly ftpOverHttp!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly httpConnectionMode!: pulumi.Output<string>;
    public readonly httpIncomingPort!: pulumi.Output<string>;
    public readonly httpsIncomingPort!: pulumi.Output<string>;
    public readonly httpsReplacementMessage!: pulumi.Output<string>;
    public readonly incomingIp!: pulumi.Output<string>;
    public readonly incomingIp6!: pulumi.Output<string>;
    public readonly ipv6Status!: pulumi.Output<string>;
    public readonly messageUponServerError!: pulumi.Output<string>;
    public readonly outgoingIp!: pulumi.Output<string>;
    public readonly outgoingIp6!: pulumi.Output<string>;
    public readonly pacFileData!: pulumi.Output<string>;
    public readonly pacFileName!: pulumi.Output<string>;
    public readonly pacFileServerPort!: pulumi.Output<string>;
    public readonly pacFileServerStatus!: pulumi.Output<string>;
    public readonly pacFileThroughHttps!: pulumi.Output<string>;
    public readonly pacFileUrl!: pulumi.Output<string>;
    public readonly pacPolicies!: pulumi.Output<outputs.WebProxyExplicitPacPolicy[] | undefined>;
    public readonly prefDnsResult!: pulumi.Output<string>;
    public readonly realm!: pulumi.Output<string>;
    public readonly secDefaultAction!: pulumi.Output<string>;
    public readonly secureWebProxy!: pulumi.Output<string>;
    public readonly secureWebProxyCerts!: pulumi.Output<outputs.WebProxyExplicitSecureWebProxyCert[] | undefined>;
    public readonly socks!: pulumi.Output<string>;
    public readonly socksIncomingPort!: pulumi.Output<string>;
    public readonly sslAlgorithm!: pulumi.Output<string>;
    public readonly sslDhBits!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly strictGuest!: pulumi.Output<string>;
    public readonly traceAuthNoRsp!: pulumi.Output<string>;
    public readonly unknownHttpVersion!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebProxyExplicit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebProxyExplicitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebProxyExplicitArgs | WebProxyExplicitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebProxyExplicitState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["ftpIncomingPort"] = state ? state.ftpIncomingPort : undefined;
            resourceInputs["ftpOverHttp"] = state ? state.ftpOverHttp : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["httpConnectionMode"] = state ? state.httpConnectionMode : undefined;
            resourceInputs["httpIncomingPort"] = state ? state.httpIncomingPort : undefined;
            resourceInputs["httpsIncomingPort"] = state ? state.httpsIncomingPort : undefined;
            resourceInputs["httpsReplacementMessage"] = state ? state.httpsReplacementMessage : undefined;
            resourceInputs["incomingIp"] = state ? state.incomingIp : undefined;
            resourceInputs["incomingIp6"] = state ? state.incomingIp6 : undefined;
            resourceInputs["ipv6Status"] = state ? state.ipv6Status : undefined;
            resourceInputs["messageUponServerError"] = state ? state.messageUponServerError : undefined;
            resourceInputs["outgoingIp"] = state ? state.outgoingIp : undefined;
            resourceInputs["outgoingIp6"] = state ? state.outgoingIp6 : undefined;
            resourceInputs["pacFileData"] = state ? state.pacFileData : undefined;
            resourceInputs["pacFileName"] = state ? state.pacFileName : undefined;
            resourceInputs["pacFileServerPort"] = state ? state.pacFileServerPort : undefined;
            resourceInputs["pacFileServerStatus"] = state ? state.pacFileServerStatus : undefined;
            resourceInputs["pacFileThroughHttps"] = state ? state.pacFileThroughHttps : undefined;
            resourceInputs["pacFileUrl"] = state ? state.pacFileUrl : undefined;
            resourceInputs["pacPolicies"] = state ? state.pacPolicies : undefined;
            resourceInputs["prefDnsResult"] = state ? state.prefDnsResult : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["secDefaultAction"] = state ? state.secDefaultAction : undefined;
            resourceInputs["secureWebProxy"] = state ? state.secureWebProxy : undefined;
            resourceInputs["secureWebProxyCerts"] = state ? state.secureWebProxyCerts : undefined;
            resourceInputs["socks"] = state ? state.socks : undefined;
            resourceInputs["socksIncomingPort"] = state ? state.socksIncomingPort : undefined;
            resourceInputs["sslAlgorithm"] = state ? state.sslAlgorithm : undefined;
            resourceInputs["sslDhBits"] = state ? state.sslDhBits : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["strictGuest"] = state ? state.strictGuest : undefined;
            resourceInputs["traceAuthNoRsp"] = state ? state.traceAuthNoRsp : undefined;
            resourceInputs["unknownHttpVersion"] = state ? state.unknownHttpVersion : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebProxyExplicitArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["ftpIncomingPort"] = args ? args.ftpIncomingPort : undefined;
            resourceInputs["ftpOverHttp"] = args ? args.ftpOverHttp : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["httpConnectionMode"] = args ? args.httpConnectionMode : undefined;
            resourceInputs["httpIncomingPort"] = args ? args.httpIncomingPort : undefined;
            resourceInputs["httpsIncomingPort"] = args ? args.httpsIncomingPort : undefined;
            resourceInputs["httpsReplacementMessage"] = args ? args.httpsReplacementMessage : undefined;
            resourceInputs["incomingIp"] = args ? args.incomingIp : undefined;
            resourceInputs["incomingIp6"] = args ? args.incomingIp6 : undefined;
            resourceInputs["ipv6Status"] = args ? args.ipv6Status : undefined;
            resourceInputs["messageUponServerError"] = args ? args.messageUponServerError : undefined;
            resourceInputs["outgoingIp"] = args ? args.outgoingIp : undefined;
            resourceInputs["outgoingIp6"] = args ? args.outgoingIp6 : undefined;
            resourceInputs["pacFileData"] = args ? args.pacFileData : undefined;
            resourceInputs["pacFileName"] = args ? args.pacFileName : undefined;
            resourceInputs["pacFileServerPort"] = args ? args.pacFileServerPort : undefined;
            resourceInputs["pacFileServerStatus"] = args ? args.pacFileServerStatus : undefined;
            resourceInputs["pacFileThroughHttps"] = args ? args.pacFileThroughHttps : undefined;
            resourceInputs["pacFileUrl"] = args ? args.pacFileUrl : undefined;
            resourceInputs["pacPolicies"] = args ? args.pacPolicies : undefined;
            resourceInputs["prefDnsResult"] = args ? args.prefDnsResult : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["secDefaultAction"] = args ? args.secDefaultAction : undefined;
            resourceInputs["secureWebProxy"] = args ? args.secureWebProxy : undefined;
            resourceInputs["secureWebProxyCerts"] = args ? args.secureWebProxyCerts : undefined;
            resourceInputs["socks"] = args ? args.socks : undefined;
            resourceInputs["socksIncomingPort"] = args ? args.socksIncomingPort : undefined;
            resourceInputs["sslAlgorithm"] = args ? args.sslAlgorithm : undefined;
            resourceInputs["sslDhBits"] = args ? args.sslDhBits : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["strictGuest"] = args ? args.strictGuest : undefined;
            resourceInputs["traceAuthNoRsp"] = args ? args.traceAuthNoRsp : undefined;
            resourceInputs["unknownHttpVersion"] = args ? args.unknownHttpVersion : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebProxyExplicit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebProxyExplicit resources.
 */
export interface WebProxyExplicitState {
    dynamicSortSubtable?: pulumi.Input<string>;
    ftpIncomingPort?: pulumi.Input<string>;
    ftpOverHttp?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    httpConnectionMode?: pulumi.Input<string>;
    httpIncomingPort?: pulumi.Input<string>;
    httpsIncomingPort?: pulumi.Input<string>;
    httpsReplacementMessage?: pulumi.Input<string>;
    incomingIp?: pulumi.Input<string>;
    incomingIp6?: pulumi.Input<string>;
    ipv6Status?: pulumi.Input<string>;
    messageUponServerError?: pulumi.Input<string>;
    outgoingIp?: pulumi.Input<string>;
    outgoingIp6?: pulumi.Input<string>;
    pacFileData?: pulumi.Input<string>;
    pacFileName?: pulumi.Input<string>;
    pacFileServerPort?: pulumi.Input<string>;
    pacFileServerStatus?: pulumi.Input<string>;
    pacFileThroughHttps?: pulumi.Input<string>;
    pacFileUrl?: pulumi.Input<string>;
    pacPolicies?: pulumi.Input<pulumi.Input<inputs.WebProxyExplicitPacPolicy>[]>;
    prefDnsResult?: pulumi.Input<string>;
    realm?: pulumi.Input<string>;
    secDefaultAction?: pulumi.Input<string>;
    secureWebProxy?: pulumi.Input<string>;
    secureWebProxyCerts?: pulumi.Input<pulumi.Input<inputs.WebProxyExplicitSecureWebProxyCert>[]>;
    socks?: pulumi.Input<string>;
    socksIncomingPort?: pulumi.Input<string>;
    sslAlgorithm?: pulumi.Input<string>;
    sslDhBits?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    strictGuest?: pulumi.Input<string>;
    traceAuthNoRsp?: pulumi.Input<string>;
    unknownHttpVersion?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebProxyExplicit resource.
 */
export interface WebProxyExplicitArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    ftpIncomingPort?: pulumi.Input<string>;
    ftpOverHttp?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    httpConnectionMode?: pulumi.Input<string>;
    httpIncomingPort?: pulumi.Input<string>;
    httpsIncomingPort?: pulumi.Input<string>;
    httpsReplacementMessage?: pulumi.Input<string>;
    incomingIp?: pulumi.Input<string>;
    incomingIp6?: pulumi.Input<string>;
    ipv6Status?: pulumi.Input<string>;
    messageUponServerError?: pulumi.Input<string>;
    outgoingIp?: pulumi.Input<string>;
    outgoingIp6?: pulumi.Input<string>;
    pacFileData?: pulumi.Input<string>;
    pacFileName?: pulumi.Input<string>;
    pacFileServerPort?: pulumi.Input<string>;
    pacFileServerStatus?: pulumi.Input<string>;
    pacFileThroughHttps?: pulumi.Input<string>;
    pacFileUrl?: pulumi.Input<string>;
    pacPolicies?: pulumi.Input<pulumi.Input<inputs.WebProxyExplicitPacPolicy>[]>;
    prefDnsResult?: pulumi.Input<string>;
    realm?: pulumi.Input<string>;
    secDefaultAction?: pulumi.Input<string>;
    secureWebProxy?: pulumi.Input<string>;
    secureWebProxyCerts?: pulumi.Input<pulumi.Input<inputs.WebProxyExplicitSecureWebProxyCert>[]>;
    socks?: pulumi.Input<string>;
    socksIncomingPort?: pulumi.Input<string>;
    sslAlgorithm?: pulumi.Input<string>;
    sslDhBits?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    strictGuest?: pulumi.Input<string>;
    traceAuthNoRsp?: pulumi.Input<string>;
    unknownHttpVersion?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
