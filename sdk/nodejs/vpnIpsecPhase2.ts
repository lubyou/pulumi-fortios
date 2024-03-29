// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VpnIpsecPhase2 extends pulumi.CustomResource {
    /**
     * Get an existing VpnIpsecPhase2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnIpsecPhase2State, opts?: pulumi.CustomResourceOptions): VpnIpsecPhase2 {
        return new VpnIpsecPhase2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/vpnIpsecPhase2:VpnIpsecPhase2';

    /**
     * Returns true if the given object is an instance of VpnIpsecPhase2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnIpsecPhase2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnIpsecPhase2.__pulumiType;
    }

    public readonly addRoute!: pulumi.Output<string>;
    public readonly autoNegotiate!: pulumi.Output<string>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly dhcpIpsec!: pulumi.Output<string>;
    public readonly dhgrp!: pulumi.Output<string>;
    public readonly diffserv!: pulumi.Output<string>;
    public readonly diffservcode!: pulumi.Output<string>;
    public readonly dstAddrType!: pulumi.Output<string>;
    public readonly dstEndIp!: pulumi.Output<string>;
    public readonly dstEndIp6!: pulumi.Output<string>;
    public readonly dstName!: pulumi.Output<string>;
    public readonly dstName6!: pulumi.Output<string>;
    public readonly dstPort!: pulumi.Output<number>;
    public readonly dstStartIp!: pulumi.Output<string>;
    public readonly dstStartIp6!: pulumi.Output<string>;
    public readonly dstSubnet!: pulumi.Output<string>;
    public readonly dstSubnet6!: pulumi.Output<string>;
    public readonly encapsulation!: pulumi.Output<string>;
    public readonly inboundDscpCopy!: pulumi.Output<string>;
    public readonly initiatorTsNarrow!: pulumi.Output<string>;
    public readonly ipv4Df!: pulumi.Output<string>;
    public readonly keepalive!: pulumi.Output<string>;
    public readonly keylifeType!: pulumi.Output<string>;
    public readonly keylifekbs!: pulumi.Output<number>;
    public readonly keylifeseconds!: pulumi.Output<number>;
    public readonly l2tp!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly pfs!: pulumi.Output<string>;
    public readonly phase1name!: pulumi.Output<string>;
    public readonly proposal!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<number>;
    public readonly replay!: pulumi.Output<string>;
    public readonly routeOverlap!: pulumi.Output<string>;
    public readonly selectorMatch!: pulumi.Output<string>;
    public readonly singleSource!: pulumi.Output<string>;
    public readonly srcAddrType!: pulumi.Output<string>;
    public readonly srcEndIp!: pulumi.Output<string>;
    public readonly srcEndIp6!: pulumi.Output<string>;
    public readonly srcName!: pulumi.Output<string>;
    public readonly srcName6!: pulumi.Output<string>;
    public readonly srcPort!: pulumi.Output<number>;
    public readonly srcStartIp!: pulumi.Output<string>;
    public readonly srcStartIp6!: pulumi.Output<string>;
    public readonly srcSubnet!: pulumi.Output<string>;
    public readonly srcSubnet6!: pulumi.Output<string>;
    public readonly useNatip!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a VpnIpsecPhase2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnIpsecPhase2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnIpsecPhase2Args | VpnIpsecPhase2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnIpsecPhase2State | undefined;
            resourceInputs["addRoute"] = state ? state.addRoute : undefined;
            resourceInputs["autoNegotiate"] = state ? state.autoNegotiate : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dhcpIpsec"] = state ? state.dhcpIpsec : undefined;
            resourceInputs["dhgrp"] = state ? state.dhgrp : undefined;
            resourceInputs["diffserv"] = state ? state.diffserv : undefined;
            resourceInputs["diffservcode"] = state ? state.diffservcode : undefined;
            resourceInputs["dstAddrType"] = state ? state.dstAddrType : undefined;
            resourceInputs["dstEndIp"] = state ? state.dstEndIp : undefined;
            resourceInputs["dstEndIp6"] = state ? state.dstEndIp6 : undefined;
            resourceInputs["dstName"] = state ? state.dstName : undefined;
            resourceInputs["dstName6"] = state ? state.dstName6 : undefined;
            resourceInputs["dstPort"] = state ? state.dstPort : undefined;
            resourceInputs["dstStartIp"] = state ? state.dstStartIp : undefined;
            resourceInputs["dstStartIp6"] = state ? state.dstStartIp6 : undefined;
            resourceInputs["dstSubnet"] = state ? state.dstSubnet : undefined;
            resourceInputs["dstSubnet6"] = state ? state.dstSubnet6 : undefined;
            resourceInputs["encapsulation"] = state ? state.encapsulation : undefined;
            resourceInputs["inboundDscpCopy"] = state ? state.inboundDscpCopy : undefined;
            resourceInputs["initiatorTsNarrow"] = state ? state.initiatorTsNarrow : undefined;
            resourceInputs["ipv4Df"] = state ? state.ipv4Df : undefined;
            resourceInputs["keepalive"] = state ? state.keepalive : undefined;
            resourceInputs["keylifeType"] = state ? state.keylifeType : undefined;
            resourceInputs["keylifekbs"] = state ? state.keylifekbs : undefined;
            resourceInputs["keylifeseconds"] = state ? state.keylifeseconds : undefined;
            resourceInputs["l2tp"] = state ? state.l2tp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pfs"] = state ? state.pfs : undefined;
            resourceInputs["phase1name"] = state ? state.phase1name : undefined;
            resourceInputs["proposal"] = state ? state.proposal : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["replay"] = state ? state.replay : undefined;
            resourceInputs["routeOverlap"] = state ? state.routeOverlap : undefined;
            resourceInputs["selectorMatch"] = state ? state.selectorMatch : undefined;
            resourceInputs["singleSource"] = state ? state.singleSource : undefined;
            resourceInputs["srcAddrType"] = state ? state.srcAddrType : undefined;
            resourceInputs["srcEndIp"] = state ? state.srcEndIp : undefined;
            resourceInputs["srcEndIp6"] = state ? state.srcEndIp6 : undefined;
            resourceInputs["srcName"] = state ? state.srcName : undefined;
            resourceInputs["srcName6"] = state ? state.srcName6 : undefined;
            resourceInputs["srcPort"] = state ? state.srcPort : undefined;
            resourceInputs["srcStartIp"] = state ? state.srcStartIp : undefined;
            resourceInputs["srcStartIp6"] = state ? state.srcStartIp6 : undefined;
            resourceInputs["srcSubnet"] = state ? state.srcSubnet : undefined;
            resourceInputs["srcSubnet6"] = state ? state.srcSubnet6 : undefined;
            resourceInputs["useNatip"] = state ? state.useNatip : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as VpnIpsecPhase2Args | undefined;
            if ((!args || args.phase1name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'phase1name'");
            }
            if ((!args || args.proposal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proposal'");
            }
            resourceInputs["addRoute"] = args ? args.addRoute : undefined;
            resourceInputs["autoNegotiate"] = args ? args.autoNegotiate : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dhcpIpsec"] = args ? args.dhcpIpsec : undefined;
            resourceInputs["dhgrp"] = args ? args.dhgrp : undefined;
            resourceInputs["diffserv"] = args ? args.diffserv : undefined;
            resourceInputs["diffservcode"] = args ? args.diffservcode : undefined;
            resourceInputs["dstAddrType"] = args ? args.dstAddrType : undefined;
            resourceInputs["dstEndIp"] = args ? args.dstEndIp : undefined;
            resourceInputs["dstEndIp6"] = args ? args.dstEndIp6 : undefined;
            resourceInputs["dstName"] = args ? args.dstName : undefined;
            resourceInputs["dstName6"] = args ? args.dstName6 : undefined;
            resourceInputs["dstPort"] = args ? args.dstPort : undefined;
            resourceInputs["dstStartIp"] = args ? args.dstStartIp : undefined;
            resourceInputs["dstStartIp6"] = args ? args.dstStartIp6 : undefined;
            resourceInputs["dstSubnet"] = args ? args.dstSubnet : undefined;
            resourceInputs["dstSubnet6"] = args ? args.dstSubnet6 : undefined;
            resourceInputs["encapsulation"] = args ? args.encapsulation : undefined;
            resourceInputs["inboundDscpCopy"] = args ? args.inboundDscpCopy : undefined;
            resourceInputs["initiatorTsNarrow"] = args ? args.initiatorTsNarrow : undefined;
            resourceInputs["ipv4Df"] = args ? args.ipv4Df : undefined;
            resourceInputs["keepalive"] = args ? args.keepalive : undefined;
            resourceInputs["keylifeType"] = args ? args.keylifeType : undefined;
            resourceInputs["keylifekbs"] = args ? args.keylifekbs : undefined;
            resourceInputs["keylifeseconds"] = args ? args.keylifeseconds : undefined;
            resourceInputs["l2tp"] = args ? args.l2tp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pfs"] = args ? args.pfs : undefined;
            resourceInputs["phase1name"] = args ? args.phase1name : undefined;
            resourceInputs["proposal"] = args ? args.proposal : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["replay"] = args ? args.replay : undefined;
            resourceInputs["routeOverlap"] = args ? args.routeOverlap : undefined;
            resourceInputs["selectorMatch"] = args ? args.selectorMatch : undefined;
            resourceInputs["singleSource"] = args ? args.singleSource : undefined;
            resourceInputs["srcAddrType"] = args ? args.srcAddrType : undefined;
            resourceInputs["srcEndIp"] = args ? args.srcEndIp : undefined;
            resourceInputs["srcEndIp6"] = args ? args.srcEndIp6 : undefined;
            resourceInputs["srcName"] = args ? args.srcName : undefined;
            resourceInputs["srcName6"] = args ? args.srcName6 : undefined;
            resourceInputs["srcPort"] = args ? args.srcPort : undefined;
            resourceInputs["srcStartIp"] = args ? args.srcStartIp : undefined;
            resourceInputs["srcStartIp6"] = args ? args.srcStartIp6 : undefined;
            resourceInputs["srcSubnet"] = args ? args.srcSubnet : undefined;
            resourceInputs["srcSubnet6"] = args ? args.srcSubnet6 : undefined;
            resourceInputs["useNatip"] = args ? args.useNatip : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpnIpsecPhase2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnIpsecPhase2 resources.
 */
export interface VpnIpsecPhase2State {
    addRoute?: pulumi.Input<string>;
    autoNegotiate?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    dhcpIpsec?: pulumi.Input<string>;
    dhgrp?: pulumi.Input<string>;
    diffserv?: pulumi.Input<string>;
    diffservcode?: pulumi.Input<string>;
    dstAddrType?: pulumi.Input<string>;
    dstEndIp?: pulumi.Input<string>;
    dstEndIp6?: pulumi.Input<string>;
    dstName?: pulumi.Input<string>;
    dstName6?: pulumi.Input<string>;
    dstPort?: pulumi.Input<number>;
    dstStartIp?: pulumi.Input<string>;
    dstStartIp6?: pulumi.Input<string>;
    dstSubnet?: pulumi.Input<string>;
    dstSubnet6?: pulumi.Input<string>;
    encapsulation?: pulumi.Input<string>;
    inboundDscpCopy?: pulumi.Input<string>;
    initiatorTsNarrow?: pulumi.Input<string>;
    ipv4Df?: pulumi.Input<string>;
    keepalive?: pulumi.Input<string>;
    keylifeType?: pulumi.Input<string>;
    keylifekbs?: pulumi.Input<number>;
    keylifeseconds?: pulumi.Input<number>;
    l2tp?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    pfs?: pulumi.Input<string>;
    phase1name?: pulumi.Input<string>;
    proposal?: pulumi.Input<string>;
    protocol?: pulumi.Input<number>;
    replay?: pulumi.Input<string>;
    routeOverlap?: pulumi.Input<string>;
    selectorMatch?: pulumi.Input<string>;
    singleSource?: pulumi.Input<string>;
    srcAddrType?: pulumi.Input<string>;
    srcEndIp?: pulumi.Input<string>;
    srcEndIp6?: pulumi.Input<string>;
    srcName?: pulumi.Input<string>;
    srcName6?: pulumi.Input<string>;
    srcPort?: pulumi.Input<number>;
    srcStartIp?: pulumi.Input<string>;
    srcStartIp6?: pulumi.Input<string>;
    srcSubnet?: pulumi.Input<string>;
    srcSubnet6?: pulumi.Input<string>;
    useNatip?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnIpsecPhase2 resource.
 */
export interface VpnIpsecPhase2Args {
    addRoute?: pulumi.Input<string>;
    autoNegotiate?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    dhcpIpsec?: pulumi.Input<string>;
    dhgrp?: pulumi.Input<string>;
    diffserv?: pulumi.Input<string>;
    diffservcode?: pulumi.Input<string>;
    dstAddrType?: pulumi.Input<string>;
    dstEndIp?: pulumi.Input<string>;
    dstEndIp6?: pulumi.Input<string>;
    dstName?: pulumi.Input<string>;
    dstName6?: pulumi.Input<string>;
    dstPort?: pulumi.Input<number>;
    dstStartIp?: pulumi.Input<string>;
    dstStartIp6?: pulumi.Input<string>;
    dstSubnet?: pulumi.Input<string>;
    dstSubnet6?: pulumi.Input<string>;
    encapsulation?: pulumi.Input<string>;
    inboundDscpCopy?: pulumi.Input<string>;
    initiatorTsNarrow?: pulumi.Input<string>;
    ipv4Df?: pulumi.Input<string>;
    keepalive?: pulumi.Input<string>;
    keylifeType?: pulumi.Input<string>;
    keylifekbs?: pulumi.Input<number>;
    keylifeseconds?: pulumi.Input<number>;
    l2tp?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    pfs?: pulumi.Input<string>;
    phase1name: pulumi.Input<string>;
    proposal: pulumi.Input<string>;
    protocol?: pulumi.Input<number>;
    replay?: pulumi.Input<string>;
    routeOverlap?: pulumi.Input<string>;
    selectorMatch?: pulumi.Input<string>;
    singleSource?: pulumi.Input<string>;
    srcAddrType?: pulumi.Input<string>;
    srcEndIp?: pulumi.Input<string>;
    srcEndIp6?: pulumi.Input<string>;
    srcName?: pulumi.Input<string>;
    srcName6?: pulumi.Input<string>;
    srcPort?: pulumi.Input<number>;
    srcStartIp?: pulumi.Input<string>;
    srcStartIp6?: pulumi.Input<string>;
    srcSubnet?: pulumi.Input<string>;
    srcSubnet6?: pulumi.Input<string>;
    useNatip?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
