// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure MS Exchange server entries.
 *
 * ## Import
 *
 * User Exchange can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/userExchange:UserExchange labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class UserExchange extends pulumi.CustomResource {
    /**
     * Get an existing UserExchange resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserExchangeState, opts?: pulumi.CustomResourceOptions): UserExchange {
        return new UserExchange(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/userExchange:UserExchange';

    /**
     * Returns true if the given object is an instance of UserExchange.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserExchange {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserExchange.__pulumiType;
    }

    /**
     * Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
     */
    public readonly authLevel!: pulumi.Output<string>;
    /**
     * Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
     */
    public readonly authType!: pulumi.Output<string>;
    /**
     * Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
     */
    public readonly autoDiscoverKdc!: pulumi.Output<string>;
    /**
     * Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
     */
    public readonly connectProtocol!: pulumi.Output<string>;
    /**
     * MS Exchange server fully qualified domain name.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
     */
    public readonly httpAuthType!: pulumi.Output<string>;
    /**
     * Server IPv4 address.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
     */
    public readonly kdcIps!: pulumi.Output<outputs.UserExchangeKdcIp[] | undefined>;
    /**
     * MS Exchange server entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for the specified username.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * MS Exchange server hostname.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * User name used to sign in to the server. Must have proper permissions for service.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a UserExchange resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserExchangeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserExchangeArgs | UserExchangeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserExchangeState | undefined;
            inputs["authLevel"] = state ? state.authLevel : undefined;
            inputs["authType"] = state ? state.authType : undefined;
            inputs["autoDiscoverKdc"] = state ? state.autoDiscoverKdc : undefined;
            inputs["connectProtocol"] = state ? state.connectProtocol : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["httpAuthType"] = state ? state.httpAuthType : undefined;
            inputs["ip"] = state ? state.ip : undefined;
            inputs["kdcIps"] = state ? state.kdcIps : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["serverName"] = state ? state.serverName : undefined;
            inputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            inputs["username"] = state ? state.username : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserExchangeArgs | undefined;
            inputs["authLevel"] = args ? args.authLevel : undefined;
            inputs["authType"] = args ? args.authType : undefined;
            inputs["autoDiscoverKdc"] = args ? args.autoDiscoverKdc : undefined;
            inputs["connectProtocol"] = args ? args.connectProtocol : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["httpAuthType"] = args ? args.httpAuthType : undefined;
            inputs["ip"] = args ? args.ip : undefined;
            inputs["kdcIps"] = args ? args.kdcIps : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UserExchange.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserExchange resources.
 */
export interface UserExchangeState {
    /**
     * Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
     */
    authLevel?: pulumi.Input<string>;
    /**
     * Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
     */
    authType?: pulumi.Input<string>;
    /**
     * Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
     */
    autoDiscoverKdc?: pulumi.Input<string>;
    /**
     * Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
     */
    connectProtocol?: pulumi.Input<string>;
    /**
     * MS Exchange server fully qualified domain name.
     */
    domainName?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
     */
    httpAuthType?: pulumi.Input<string>;
    /**
     * Server IPv4 address.
     */
    ip?: pulumi.Input<string>;
    /**
     * KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
     */
    kdcIps?: pulumi.Input<pulumi.Input<inputs.UserExchangeKdcIp>[]>;
    /**
     * MS Exchange server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for the specified username.
     */
    password?: pulumi.Input<string>;
    /**
     * MS Exchange server hostname.
     */
    serverName?: pulumi.Input<string>;
    /**
     * Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * User name used to sign in to the server. Must have proper permissions for service.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserExchange resource.
 */
export interface UserExchangeArgs {
    /**
     * Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
     */
    authLevel?: pulumi.Input<string>;
    /**
     * Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
     */
    authType?: pulumi.Input<string>;
    /**
     * Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
     */
    autoDiscoverKdc?: pulumi.Input<string>;
    /**
     * Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
     */
    connectProtocol?: pulumi.Input<string>;
    /**
     * MS Exchange server fully qualified domain name.
     */
    domainName?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
     */
    httpAuthType?: pulumi.Input<string>;
    /**
     * Server IPv4 address.
     */
    ip?: pulumi.Input<string>;
    /**
     * KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
     */
    kdcIps?: pulumi.Input<pulumi.Input<inputs.UserExchangeKdcIp>[]>;
    /**
     * MS Exchange server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for the specified username.
     */
    password?: pulumi.Input<string>;
    /**
     * MS Exchange server hostname.
     */
    serverName?: pulumi.Input<string>;
    /**
     * Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * User name used to sign in to the server. Must have proper permissions for service.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
