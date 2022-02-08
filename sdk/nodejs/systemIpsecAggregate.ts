// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure an aggregate of IPsec tunnels.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname1VpnIpsecPhase1Interface = new fortios.VpnIpsecPhase1Interface("trname1VpnIpsecPhase1Interface", {
 *     acctVerify: "disable",
 *     addGwRoute: "disable",
 *     addRoute: "enable",
 *     assignIp: "enable",
 *     assignIpFrom: "range",
 *     authmethod: "psk",
 *     autoDiscoveryForwarder: "disable",
 *     autoDiscoveryPsk: "disable",
 *     autoDiscoveryReceiver: "disable",
 *     autoDiscoverySender: "disable",
 *     autoNegotiate: "enable",
 *     certIdValidation: "enable",
 *     childlessIke: "disable",
 *     clientAutoNegotiate: "disable",
 *     clientKeepAlive: "disable",
 *     defaultGw: "0.0.0.0",
 *     defaultGwPriority: 0,
 *     dhgrp: "14 5",
 *     digitalSignatureAuth: "disable",
 *     distance: 15,
 *     dnsMode: "manual",
 *     dpd: "on-demand",
 *     dpdRetrycount: 3,
 *     dpdRetryinterval: "20",
 *     eap: "disable",
 *     eapIdentity: "use-id-payload",
 *     encapLocalGw4: "0.0.0.0",
 *     encapLocalGw6: "::",
 *     encapRemoteGw4: "0.0.0.0",
 *     encapRemoteGw6: "::",
 *     encapsulation: "none",
 *     encapsulationAddress: "ike",
 *     enforceUniqueId: "disable",
 *     exchangeInterfaceIp: "disable",
 *     exchangeIpAddr4: "0.0.0.0",
 *     exchangeIpAddr6: "::",
 *     forticlientEnforcement: "disable",
 *     fragmentation: "enable",
 *     fragmentationMtu: 1200,
 *     groupAuthentication: "disable",
 *     haSyncEspSeqno: "enable",
 *     idleTimeout: "disable",
 *     idleTimeoutinterval: 15,
 *     ikeVersion: "1",
 *     includeLocalLan: "disable",
 *     "interface": "port3",
 *     ipVersion: "4",
 *     ipv4DnsServer1: "0.0.0.0",
 *     ipv4DnsServer2: "0.0.0.0",
 *     ipv4DnsServer3: "0.0.0.0",
 *     ipv4EndIp: "0.0.0.0",
 *     ipv4Netmask: "255.255.255.255",
 *     ipv4StartIp: "0.0.0.0",
 *     ipv4WinsServer1: "0.0.0.0",
 *     ipv4WinsServer2: "0.0.0.0",
 *     ipv6DnsServer1: "::",
 *     ipv6DnsServer2: "::",
 *     ipv6DnsServer3: "::",
 *     ipv6EndIp: "::",
 *     ipv6Prefix: 128,
 *     ipv6StartIp: "::",
 *     keepalive: 10,
 *     keylife: 86400,
 *     localGw: "0.0.0.0",
 *     localGw6: "::",
 *     localidType: "auto",
 *     meshSelectorType: "disable",
 *     mode: "main",
 *     modeCfg: "disable",
 *     monitorHoldDownDelay: 0,
 *     monitorHoldDownTime: "00:00",
 *     monitorHoldDownType: "immediate",
 *     monitorHoldDownWeekday: "sunday",
 *     nattraversal: "enable",
 *     negotiateTimeout: 30,
 *     netDevice: "disable",
 *     passiveMode: "disable",
 *     peertype: "any",
 *     ppk: "disable",
 *     priority: 0,
 *     proposal: "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1",
 *     psksecret: "eweeeeeeeecee",
 *     reauth: "disable",
 *     rekey: "enable",
 *     remoteGw: "2.2.2.2",
 *     remoteGw6: "::",
 *     rsaSignatureFormat: "pkcs1",
 *     savePassword: "disable",
 *     sendCertChain: "enable",
 *     signatureHashAlg: "sha2-512 sha2-384 sha2-256 sha1",
 *     suiteB: "disable",
 *     tunnelSearch: "selectors",
 *     type: "static",
 *     unitySupport: "enable",
 *     wizardType: "custom",
 *     xauthtype: "disable",
 * });
 * const trname1VpnIpsecPhase2Interface = new fortios.VpnIpsecPhase2Interface("trname1VpnIpsecPhase2Interface", {
 *     addRoute: "phase1",
 *     autoDiscoveryForwarder: "phase1",
 *     autoDiscoverySender: "phase1",
 *     autoNegotiate: "disable",
 *     dhcpIpsec: "disable",
 *     dhgrp: "14 5",
 *     dstAddrType: "subnet",
 *     dstEndIp6: "::",
 *     dstPort: 0,
 *     dstSubnet: "0.0.0.0 0.0.0.0",
 *     encapsulation: "tunnel-mode",
 *     keepalive: "disable",
 *     keylifeType: "seconds",
 *     keylifekbs: 5120,
 *     keylifeseconds: 43200,
 *     l2tp: "disable",
 *     pfs: "enable",
 *     phase1name: trname1VpnIpsecPhase1Interface.name,
 *     proposal: "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305",
 *     protocol: 0,
 *     replay: "enable",
 *     routeOverlap: "use-new",
 *     singleSource: "disable",
 *     srcAddrType: "subnet",
 *     srcEndIp6: "::",
 *     srcPort: 0,
 *     srcSubnet: "0.0.0.0 0.0.0.0",
 * });
 * const trname = new fortios.SystemIpsecAggregate("trname", {
 *     algorithm: "round-robin",
 *     members: [{
 *         tunnelName: trname1VpnIpsecPhase1Interface.name,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * System IpsecAggregate can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemIpsecAggregate:SystemIpsecAggregate labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemIpsecAggregate extends pulumi.CustomResource {
    /**
     * Get an existing SystemIpsecAggregate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemIpsecAggregateState, opts?: pulumi.CustomResourceOptions): SystemIpsecAggregate {
        return new SystemIpsecAggregate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemIpsecAggregate:SystemIpsecAggregate';

    /**
     * Returns true if the given object is an instance of SystemIpsecAggregate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemIpsecAggregate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemIpsecAggregate.__pulumiType;
    }

    /**
     * Frame distribution algorithm.
     */
    public readonly algorithm!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Member tunnels of the aggregate. The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.SystemIpsecAggregateMember[]>;
    /**
     * IPsec aggregate name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemIpsecAggregate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemIpsecAggregateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemIpsecAggregateArgs | SystemIpsecAggregateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemIpsecAggregateState | undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemIpsecAggregateArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemIpsecAggregate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemIpsecAggregate resources.
 */
export interface SystemIpsecAggregateState {
    /**
     * Frame distribution algorithm.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Member tunnels of the aggregate. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.SystemIpsecAggregateMember>[]>;
    /**
     * IPsec aggregate name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemIpsecAggregate resource.
 */
export interface SystemIpsecAggregateArgs {
    /**
     * Frame distribution algorithm.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Member tunnels of the aggregate. The structure of `member` block is documented below.
     */
    members: pulumi.Input<pulumi.Input<inputs.SystemIpsecAggregateMember>[]>;
    /**
     * IPsec aggregate name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
