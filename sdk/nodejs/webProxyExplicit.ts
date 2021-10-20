// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure explicit Web proxy settings.
 *
 * ## Import
 *
 * WebProxy Explicit can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/webProxyExplicit:WebProxyExplicit labelname WebProxyExplicit
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
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

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
     */
    public readonly ftpIncomingPort!: pulumi.Output<string>;
    /**
     * Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
     */
    public readonly ftpOverHttp!: pulumi.Output<string>;
    /**
     * Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
     */
    public readonly httpIncomingPort!: pulumi.Output<string>;
    /**
     * Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
     */
    public readonly httpsIncomingPort!: pulumi.Output<string>;
    /**
     * Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
     */
    public readonly httpsReplacementMessage!: pulumi.Output<string>;
    /**
     * Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
     */
    public readonly incomingIp!: pulumi.Output<string>;
    /**
     * Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
     */
    public readonly incomingIp6!: pulumi.Output<string>;
    /**
     * Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
     */
    public readonly ipv6Status!: pulumi.Output<string>;
    /**
     * Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
     */
    public readonly messageUponServerError!: pulumi.Output<string>;
    /**
     * Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
     */
    public readonly outgoingIp!: pulumi.Output<string>;
    /**
     * Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
     */
    public readonly outgoingIp6!: pulumi.Output<string>;
    /**
     * PAC file contents enclosed in quotes (maximum of 256K bytes).
     */
    public readonly pacFileData!: pulumi.Output<string>;
    /**
     * Pac file name.
     */
    public readonly pacFileName!: pulumi.Output<string>;
    /**
     * Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
     */
    public readonly pacFileServerPort!: pulumi.Output<string>;
    /**
     * Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
     */
    public readonly pacFileServerStatus!: pulumi.Output<string>;
    /**
     * PAC file access URL.
     */
    public readonly pacFileUrl!: pulumi.Output<string>;
    /**
     * PAC policies. The structure of `pacPolicy` block is documented below.
     */
    public readonly pacPolicies!: pulumi.Output<outputs.WebProxyExplicitPacPolicy[] | undefined>;
    /**
     * Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
     */
    public readonly prefDnsResult!: pulumi.Output<string>;
    /**
     * Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
     */
    public readonly secDefaultAction!: pulumi.Output<string>;
    /**
     * Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
     */
    public readonly socks!: pulumi.Output<string>;
    /**
     * Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
     */
    public readonly socksIncomingPort!: pulumi.Output<string>;
    /**
     * Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
     */
    public readonly sslAlgorithm!: pulumi.Output<string>;
    /**
     * Enable/disable policy. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
     */
    public readonly strictGuest!: pulumi.Output<string>;
    /**
     * Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
     */
    public readonly traceAuthNoRsp!: pulumi.Output<string>;
    /**
     * Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
     */
    public readonly unknownHttpVersion!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebProxyExplicitState | undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["ftpIncomingPort"] = state ? state.ftpIncomingPort : undefined;
            inputs["ftpOverHttp"] = state ? state.ftpOverHttp : undefined;
            inputs["httpIncomingPort"] = state ? state.httpIncomingPort : undefined;
            inputs["httpsIncomingPort"] = state ? state.httpsIncomingPort : undefined;
            inputs["httpsReplacementMessage"] = state ? state.httpsReplacementMessage : undefined;
            inputs["incomingIp"] = state ? state.incomingIp : undefined;
            inputs["incomingIp6"] = state ? state.incomingIp6 : undefined;
            inputs["ipv6Status"] = state ? state.ipv6Status : undefined;
            inputs["messageUponServerError"] = state ? state.messageUponServerError : undefined;
            inputs["outgoingIp"] = state ? state.outgoingIp : undefined;
            inputs["outgoingIp6"] = state ? state.outgoingIp6 : undefined;
            inputs["pacFileData"] = state ? state.pacFileData : undefined;
            inputs["pacFileName"] = state ? state.pacFileName : undefined;
            inputs["pacFileServerPort"] = state ? state.pacFileServerPort : undefined;
            inputs["pacFileServerStatus"] = state ? state.pacFileServerStatus : undefined;
            inputs["pacFileUrl"] = state ? state.pacFileUrl : undefined;
            inputs["pacPolicies"] = state ? state.pacPolicies : undefined;
            inputs["prefDnsResult"] = state ? state.prefDnsResult : undefined;
            inputs["realm"] = state ? state.realm : undefined;
            inputs["secDefaultAction"] = state ? state.secDefaultAction : undefined;
            inputs["socks"] = state ? state.socks : undefined;
            inputs["socksIncomingPort"] = state ? state.socksIncomingPort : undefined;
            inputs["sslAlgorithm"] = state ? state.sslAlgorithm : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["strictGuest"] = state ? state.strictGuest : undefined;
            inputs["traceAuthNoRsp"] = state ? state.traceAuthNoRsp : undefined;
            inputs["unknownHttpVersion"] = state ? state.unknownHttpVersion : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebProxyExplicitArgs | undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["ftpIncomingPort"] = args ? args.ftpIncomingPort : undefined;
            inputs["ftpOverHttp"] = args ? args.ftpOverHttp : undefined;
            inputs["httpIncomingPort"] = args ? args.httpIncomingPort : undefined;
            inputs["httpsIncomingPort"] = args ? args.httpsIncomingPort : undefined;
            inputs["httpsReplacementMessage"] = args ? args.httpsReplacementMessage : undefined;
            inputs["incomingIp"] = args ? args.incomingIp : undefined;
            inputs["incomingIp6"] = args ? args.incomingIp6 : undefined;
            inputs["ipv6Status"] = args ? args.ipv6Status : undefined;
            inputs["messageUponServerError"] = args ? args.messageUponServerError : undefined;
            inputs["outgoingIp"] = args ? args.outgoingIp : undefined;
            inputs["outgoingIp6"] = args ? args.outgoingIp6 : undefined;
            inputs["pacFileData"] = args ? args.pacFileData : undefined;
            inputs["pacFileName"] = args ? args.pacFileName : undefined;
            inputs["pacFileServerPort"] = args ? args.pacFileServerPort : undefined;
            inputs["pacFileServerStatus"] = args ? args.pacFileServerStatus : undefined;
            inputs["pacFileUrl"] = args ? args.pacFileUrl : undefined;
            inputs["pacPolicies"] = args ? args.pacPolicies : undefined;
            inputs["prefDnsResult"] = args ? args.prefDnsResult : undefined;
            inputs["realm"] = args ? args.realm : undefined;
            inputs["secDefaultAction"] = args ? args.secDefaultAction : undefined;
            inputs["socks"] = args ? args.socks : undefined;
            inputs["socksIncomingPort"] = args ? args.socksIncomingPort : undefined;
            inputs["sslAlgorithm"] = args ? args.sslAlgorithm : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["strictGuest"] = args ? args.strictGuest : undefined;
            inputs["traceAuthNoRsp"] = args ? args.traceAuthNoRsp : undefined;
            inputs["unknownHttpVersion"] = args ? args.unknownHttpVersion : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebProxyExplicit.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebProxyExplicit resources.
 */
export interface WebProxyExplicitState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
     */
    ftpIncomingPort?: pulumi.Input<string>;
    /**
     * Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
     */
    ftpOverHttp?: pulumi.Input<string>;
    /**
     * Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
     */
    httpIncomingPort?: pulumi.Input<string>;
    /**
     * Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
     */
    httpsIncomingPort?: pulumi.Input<string>;
    /**
     * Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
     */
    httpsReplacementMessage?: pulumi.Input<string>;
    /**
     * Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
     */
    incomingIp?: pulumi.Input<string>;
    /**
     * Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
     */
    incomingIp6?: pulumi.Input<string>;
    /**
     * Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
     */
    ipv6Status?: pulumi.Input<string>;
    /**
     * Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
     */
    messageUponServerError?: pulumi.Input<string>;
    /**
     * Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
     */
    outgoingIp?: pulumi.Input<string>;
    /**
     * Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
     */
    outgoingIp6?: pulumi.Input<string>;
    /**
     * PAC file contents enclosed in quotes (maximum of 256K bytes).
     */
    pacFileData?: pulumi.Input<string>;
    /**
     * Pac file name.
     */
    pacFileName?: pulumi.Input<string>;
    /**
     * Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
     */
    pacFileServerPort?: pulumi.Input<string>;
    /**
     * Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
     */
    pacFileServerStatus?: pulumi.Input<string>;
    /**
     * PAC file access URL.
     */
    pacFileUrl?: pulumi.Input<string>;
    /**
     * PAC policies. The structure of `pacPolicy` block is documented below.
     */
    pacPolicies?: pulumi.Input<pulumi.Input<inputs.WebProxyExplicitPacPolicy>[]>;
    /**
     * Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
     */
    prefDnsResult?: pulumi.Input<string>;
    /**
     * Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
     */
    realm?: pulumi.Input<string>;
    /**
     * Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
     */
    secDefaultAction?: pulumi.Input<string>;
    /**
     * Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
     */
    socks?: pulumi.Input<string>;
    /**
     * Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
     */
    socksIncomingPort?: pulumi.Input<string>;
    /**
     * Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
     */
    sslAlgorithm?: pulumi.Input<string>;
    /**
     * Enable/disable policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
     */
    strictGuest?: pulumi.Input<string>;
    /**
     * Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
     */
    traceAuthNoRsp?: pulumi.Input<string>;
    /**
     * Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
     */
    unknownHttpVersion?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebProxyExplicit resource.
 */
export interface WebProxyExplicitArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
     */
    ftpIncomingPort?: pulumi.Input<string>;
    /**
     * Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
     */
    ftpOverHttp?: pulumi.Input<string>;
    /**
     * Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
     */
    httpIncomingPort?: pulumi.Input<string>;
    /**
     * Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
     */
    httpsIncomingPort?: pulumi.Input<string>;
    /**
     * Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
     */
    httpsReplacementMessage?: pulumi.Input<string>;
    /**
     * Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
     */
    incomingIp?: pulumi.Input<string>;
    /**
     * Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
     */
    incomingIp6?: pulumi.Input<string>;
    /**
     * Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
     */
    ipv6Status?: pulumi.Input<string>;
    /**
     * Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
     */
    messageUponServerError?: pulumi.Input<string>;
    /**
     * Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
     */
    outgoingIp?: pulumi.Input<string>;
    /**
     * Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
     */
    outgoingIp6?: pulumi.Input<string>;
    /**
     * PAC file contents enclosed in quotes (maximum of 256K bytes).
     */
    pacFileData?: pulumi.Input<string>;
    /**
     * Pac file name.
     */
    pacFileName?: pulumi.Input<string>;
    /**
     * Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
     */
    pacFileServerPort?: pulumi.Input<string>;
    /**
     * Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
     */
    pacFileServerStatus?: pulumi.Input<string>;
    /**
     * PAC file access URL.
     */
    pacFileUrl?: pulumi.Input<string>;
    /**
     * PAC policies. The structure of `pacPolicy` block is documented below.
     */
    pacPolicies?: pulumi.Input<pulumi.Input<inputs.WebProxyExplicitPacPolicy>[]>;
    /**
     * Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
     */
    prefDnsResult?: pulumi.Input<string>;
    /**
     * Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
     */
    realm?: pulumi.Input<string>;
    /**
     * Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
     */
    secDefaultAction?: pulumi.Input<string>;
    /**
     * Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
     */
    socks?: pulumi.Input<string>;
    /**
     * Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
     */
    socksIncomingPort?: pulumi.Input<string>;
    /**
     * Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
     */
    sslAlgorithm?: pulumi.Input<string>;
    /**
     * Enable/disable policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
     */
    strictGuest?: pulumi.Input<string>;
    /**
     * Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
     */
    traceAuthNoRsp?: pulumi.Input<string>;
    /**
     * Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
     */
    unknownHttpVersion?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}