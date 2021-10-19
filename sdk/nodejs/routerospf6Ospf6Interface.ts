// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * OSPF6 interface configuration.
 *
 * > The provider supports the definition of Ospf6-Interface in Router Ospf6 `fortios.RouterOspf6`, and also allows the definition of separate Ospf6-Interface resources `fortios.Routerospf6Ospf6Interface`, but do not use a `fortios.RouterOspf6` with in-line Ospf6-Interface in conjunction with any `fortios.Routerospf6Ospf6Interface` resources, otherwise conflicts and overwrite will occur.
 *
 * ## Import
 *
 * Routerospf6 Ospf6Interface can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerospf6Ospf6Interface:Routerospf6Ospf6Interface labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Routerospf6Ospf6Interface extends pulumi.CustomResource {
    /**
     * Get an existing Routerospf6Ospf6Interface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Routerospf6Ospf6InterfaceState, opts?: pulumi.CustomResourceOptions): Routerospf6Ospf6Interface {
        return new Routerospf6Ospf6Interface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerospf6Ospf6Interface:Routerospf6Ospf6Interface';

    /**
     * Returns true if the given object is an instance of Routerospf6Ospf6Interface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Routerospf6Ospf6Interface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Routerospf6Ospf6Interface.__pulumiType;
    }

    /**
     * A.B.C.D, in IPv4 address format.
     */
    public readonly areaId!: pulumi.Output<string>;
    /**
     * Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
     */
    public readonly authentication!: pulumi.Output<string>;
    /**
     * Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    public readonly bfd!: pulumi.Output<string>;
    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    public readonly cost!: pulumi.Output<number>;
    /**
     * Dead interval.
     */
    public readonly deadInterval!: pulumi.Output<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Hello interval.
     */
    public readonly helloInterval!: pulumi.Output<number>;
    /**
     * Configuration interface name.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
     */
    public readonly ipsecAuthAlg!: pulumi.Output<string>;
    /**
     * Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
     */
    public readonly ipsecEncAlg!: pulumi.Output<string>;
    /**
     * IPsec authentication and encryption keys. The structure of `ipsecKeys` block is documented below.
     */
    public readonly ipsecKeys!: pulumi.Output<outputs.Routerospf6Ospf6InterfaceIpsecKey[] | undefined>;
    /**
     * Key roll-over interval.
     */
    public readonly keyRolloverInterval!: pulumi.Output<number>;
    /**
     * MTU for OSPFv3 packets.
     */
    public readonly mtu!: pulumi.Output<number>;
    /**
     * Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
     */
    public readonly mtuIgnore!: pulumi.Output<string>;
    /**
     * Interface entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
     */
    public readonly neighbors!: pulumi.Output<outputs.Routerospf6Ospf6InterfaceNeighbor[] | undefined>;
    /**
     * Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * priority
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Retransmit interval.
     */
    public readonly retransmitInterval!: pulumi.Output<number>;
    /**
     * Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Transmit delay.
     */
    public readonly transmitDelay!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Routerospf6Ospf6Interface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Routerospf6Ospf6InterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Routerospf6Ospf6InterfaceArgs | Routerospf6Ospf6InterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Routerospf6Ospf6InterfaceState | undefined;
            inputs["areaId"] = state ? state.areaId : undefined;
            inputs["authentication"] = state ? state.authentication : undefined;
            inputs["bfd"] = state ? state.bfd : undefined;
            inputs["cost"] = state ? state.cost : undefined;
            inputs["deadInterval"] = state ? state.deadInterval : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["helloInterval"] = state ? state.helloInterval : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["ipsecAuthAlg"] = state ? state.ipsecAuthAlg : undefined;
            inputs["ipsecEncAlg"] = state ? state.ipsecEncAlg : undefined;
            inputs["ipsecKeys"] = state ? state.ipsecKeys : undefined;
            inputs["keyRolloverInterval"] = state ? state.keyRolloverInterval : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["mtuIgnore"] = state ? state.mtuIgnore : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["neighbors"] = state ? state.neighbors : undefined;
            inputs["networkType"] = state ? state.networkType : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["retransmitInterval"] = state ? state.retransmitInterval : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["transmitDelay"] = state ? state.transmitDelay : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Routerospf6Ospf6InterfaceArgs | undefined;
            inputs["areaId"] = args ? args.areaId : undefined;
            inputs["authentication"] = args ? args.authentication : undefined;
            inputs["bfd"] = args ? args.bfd : undefined;
            inputs["cost"] = args ? args.cost : undefined;
            inputs["deadInterval"] = args ? args.deadInterval : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["helloInterval"] = args ? args.helloInterval : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["ipsecAuthAlg"] = args ? args.ipsecAuthAlg : undefined;
            inputs["ipsecEncAlg"] = args ? args.ipsecEncAlg : undefined;
            inputs["ipsecKeys"] = args ? args.ipsecKeys : undefined;
            inputs["keyRolloverInterval"] = args ? args.keyRolloverInterval : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["mtuIgnore"] = args ? args.mtuIgnore : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["neighbors"] = args ? args.neighbors : undefined;
            inputs["networkType"] = args ? args.networkType : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["retransmitInterval"] = args ? args.retransmitInterval : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["transmitDelay"] = args ? args.transmitDelay : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Routerospf6Ospf6Interface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Routerospf6Ospf6Interface resources.
 */
export interface Routerospf6Ospf6InterfaceState {
    /**
     * A.B.C.D, in IPv4 address format.
     */
    areaId?: pulumi.Input<string>;
    /**
     * Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    bfd?: pulumi.Input<string>;
    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    cost?: pulumi.Input<number>;
    /**
     * Dead interval.
     */
    deadInterval?: pulumi.Input<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Hello interval.
     */
    helloInterval?: pulumi.Input<number>;
    /**
     * Configuration interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
     */
    ipsecAuthAlg?: pulumi.Input<string>;
    /**
     * Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
     */
    ipsecEncAlg?: pulumi.Input<string>;
    /**
     * IPsec authentication and encryption keys. The structure of `ipsecKeys` block is documented below.
     */
    ipsecKeys?: pulumi.Input<pulumi.Input<inputs.Routerospf6Ospf6InterfaceIpsecKey>[]>;
    /**
     * Key roll-over interval.
     */
    keyRolloverInterval?: pulumi.Input<number>;
    /**
     * MTU for OSPFv3 packets.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
     */
    mtuIgnore?: pulumi.Input<string>;
    /**
     * Interface entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.Routerospf6Ospf6InterfaceNeighbor>[]>;
    /**
     * Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * priority
     */
    priority?: pulumi.Input<number>;
    /**
     * Retransmit interval.
     */
    retransmitInterval?: pulumi.Input<number>;
    /**
     * Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Transmit delay.
     */
    transmitDelay?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Routerospf6Ospf6Interface resource.
 */
export interface Routerospf6Ospf6InterfaceArgs {
    /**
     * A.B.C.D, in IPv4 address format.
     */
    areaId?: pulumi.Input<string>;
    /**
     * Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    bfd?: pulumi.Input<string>;
    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    cost?: pulumi.Input<number>;
    /**
     * Dead interval.
     */
    deadInterval?: pulumi.Input<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Hello interval.
     */
    helloInterval?: pulumi.Input<number>;
    /**
     * Configuration interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
     */
    ipsecAuthAlg?: pulumi.Input<string>;
    /**
     * Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
     */
    ipsecEncAlg?: pulumi.Input<string>;
    /**
     * IPsec authentication and encryption keys. The structure of `ipsecKeys` block is documented below.
     */
    ipsecKeys?: pulumi.Input<pulumi.Input<inputs.Routerospf6Ospf6InterfaceIpsecKey>[]>;
    /**
     * Key roll-over interval.
     */
    keyRolloverInterval?: pulumi.Input<number>;
    /**
     * MTU for OSPFv3 packets.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
     */
    mtuIgnore?: pulumi.Input<string>;
    /**
     * Interface entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.Routerospf6Ospf6InterfaceNeighbor>[]>;
    /**
     * Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * priority
     */
    priority?: pulumi.Input<number>;
    /**
     * Retransmit interval.
     */
    retransmitInterval?: pulumi.Input<number>;
    /**
     * Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Transmit delay.
     */
    transmitDelay?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
