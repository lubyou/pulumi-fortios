// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure replacement message groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemReplacemsgGroup("trname", {
 *     comment: "sgh",
 *     groupType: "utm",
 * });
 * ```
 *
 * ## Import
 *
 * System ReplacemsgGroup can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemReplacemsgGroup:SystemReplacemsgGroup labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemReplacemsgGroup extends pulumi.CustomResource {
    /**
     * Get an existing SystemReplacemsgGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemReplacemsgGroupState, opts?: pulumi.CustomResourceOptions): SystemReplacemsgGroup {
        return new SystemReplacemsgGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemReplacemsgGroup:SystemReplacemsgGroup';

    /**
     * Returns true if the given object is an instance of SystemReplacemsgGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemReplacemsgGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemReplacemsgGroup.__pulumiType;
    }

    /**
     * Replacement message table entries. The structure of `admin` block is documented below.
     */
    public readonly admins!: pulumi.Output<outputs.SystemReplacemsgGroupAdmin[] | undefined>;
    /**
     * Replacement message table entries. The structure of `alertmail` block is documented below.
     */
    public readonly alertmails!: pulumi.Output<outputs.SystemReplacemsgGroupAlertmail[] | undefined>;
    /**
     * Replacement message table entries. The structure of `auth` block is documented below.
     */
    public readonly auths!: pulumi.Output<outputs.SystemReplacemsgGroupAuth[] | undefined>;
    /**
     * Replacement message table entries. The structure of `automation` block is documented below.
     */
    public readonly automations!: pulumi.Output<outputs.SystemReplacemsgGroupAutomation[] | undefined>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Replacement message table entries. The structure of `customMessage` block is documented below.
     */
    public readonly customMessages!: pulumi.Output<outputs.SystemReplacemsgGroupCustomMessage[] | undefined>;
    /**
     * Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
     */
    public readonly deviceDetectionPortals!: pulumi.Output<outputs.SystemReplacemsgGroupDeviceDetectionPortal[] | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Replacement message table entries. The structure of `ec` block is documented below.
     */
    public readonly ecs!: pulumi.Output<outputs.SystemReplacemsgGroupEc[] | undefined>;
    /**
     * Replacement message table entries. The structure of `fortiguardWf` block is documented below.
     */
    public readonly fortiguardWfs!: pulumi.Output<outputs.SystemReplacemsgGroupFortiguardWf[] | undefined>;
    /**
     * Replacement message table entries. The structure of `ftp` block is documented below.
     */
    public readonly ftps!: pulumi.Output<outputs.SystemReplacemsgGroupFtp[] | undefined>;
    /**
     * Group type.
     */
    public readonly groupType!: pulumi.Output<string>;
    /**
     * Replacement message table entries. The structure of `http` block is documented below.
     */
    public readonly https!: pulumi.Output<outputs.SystemReplacemsgGroupHttp[] | undefined>;
    /**
     * Replacement message table entries. The structure of `icap` block is documented below.
     */
    public readonly icaps!: pulumi.Output<outputs.SystemReplacemsgGroupIcap[] | undefined>;
    /**
     * Replacement message table entries. The structure of `mail` block is documented below.
     */
    public readonly mails!: pulumi.Output<outputs.SystemReplacemsgGroupMail[] | undefined>;
    /**
     * Replacement message table entries. The structure of `nacQuar` block is documented below.
     */
    public readonly nacQuars!: pulumi.Output<outputs.SystemReplacemsgGroupNacQuar[] | undefined>;
    /**
     * Group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Replacement message table entries. The structure of `nntp` block is documented below.
     */
    public readonly nntps!: pulumi.Output<outputs.SystemReplacemsgGroupNntp[] | undefined>;
    /**
     * Replacement message table entries. The structure of `spam` block is documented below.
     */
    public readonly spams!: pulumi.Output<outputs.SystemReplacemsgGroupSpam[] | undefined>;
    /**
     * Replacement message table entries. The structure of `sslvpn` block is documented below.
     */
    public readonly sslvpns!: pulumi.Output<outputs.SystemReplacemsgGroupSslvpn[] | undefined>;
    /**
     * Replacement message table entries. The structure of `trafficQuota` block is documented below.
     */
    public readonly trafficQuotas!: pulumi.Output<outputs.SystemReplacemsgGroupTrafficQuota[] | undefined>;
    /**
     * Replacement message table entries. The structure of `utm` block is documented below.
     */
    public readonly utms!: pulumi.Output<outputs.SystemReplacemsgGroupUtm[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Replacement message table entries. The structure of `webproxy` block is documented below.
     */
    public readonly webproxies!: pulumi.Output<outputs.SystemReplacemsgGroupWebproxy[] | undefined>;

    /**
     * Create a SystemReplacemsgGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemReplacemsgGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemReplacemsgGroupArgs | SystemReplacemsgGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemReplacemsgGroupState | undefined;
            inputs["admins"] = state ? state.admins : undefined;
            inputs["alertmails"] = state ? state.alertmails : undefined;
            inputs["auths"] = state ? state.auths : undefined;
            inputs["automations"] = state ? state.automations : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["customMessages"] = state ? state.customMessages : undefined;
            inputs["deviceDetectionPortals"] = state ? state.deviceDetectionPortals : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["ecs"] = state ? state.ecs : undefined;
            inputs["fortiguardWfs"] = state ? state.fortiguardWfs : undefined;
            inputs["ftps"] = state ? state.ftps : undefined;
            inputs["groupType"] = state ? state.groupType : undefined;
            inputs["https"] = state ? state.https : undefined;
            inputs["icaps"] = state ? state.icaps : undefined;
            inputs["mails"] = state ? state.mails : undefined;
            inputs["nacQuars"] = state ? state.nacQuars : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nntps"] = state ? state.nntps : undefined;
            inputs["spams"] = state ? state.spams : undefined;
            inputs["sslvpns"] = state ? state.sslvpns : undefined;
            inputs["trafficQuotas"] = state ? state.trafficQuotas : undefined;
            inputs["utms"] = state ? state.utms : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["webproxies"] = state ? state.webproxies : undefined;
        } else {
            const args = argsOrState as SystemReplacemsgGroupArgs | undefined;
            if ((!args || args.groupType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupType'");
            }
            inputs["admins"] = args ? args.admins : undefined;
            inputs["alertmails"] = args ? args.alertmails : undefined;
            inputs["auths"] = args ? args.auths : undefined;
            inputs["automations"] = args ? args.automations : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["customMessages"] = args ? args.customMessages : undefined;
            inputs["deviceDetectionPortals"] = args ? args.deviceDetectionPortals : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["ecs"] = args ? args.ecs : undefined;
            inputs["fortiguardWfs"] = args ? args.fortiguardWfs : undefined;
            inputs["ftps"] = args ? args.ftps : undefined;
            inputs["groupType"] = args ? args.groupType : undefined;
            inputs["https"] = args ? args.https : undefined;
            inputs["icaps"] = args ? args.icaps : undefined;
            inputs["mails"] = args ? args.mails : undefined;
            inputs["nacQuars"] = args ? args.nacQuars : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nntps"] = args ? args.nntps : undefined;
            inputs["spams"] = args ? args.spams : undefined;
            inputs["sslvpns"] = args ? args.sslvpns : undefined;
            inputs["trafficQuotas"] = args ? args.trafficQuotas : undefined;
            inputs["utms"] = args ? args.utms : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["webproxies"] = args ? args.webproxies : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemReplacemsgGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemReplacemsgGroup resources.
 */
export interface SystemReplacemsgGroupState {
    /**
     * Replacement message table entries. The structure of `admin` block is documented below.
     */
    admins?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupAdmin>[]>;
    /**
     * Replacement message table entries. The structure of `alertmail` block is documented below.
     */
    alertmails?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupAlertmail>[]>;
    /**
     * Replacement message table entries. The structure of `auth` block is documented below.
     */
    auths?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupAuth>[]>;
    /**
     * Replacement message table entries. The structure of `automation` block is documented below.
     */
    automations?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupAutomation>[]>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `customMessage` block is documented below.
     */
    customMessages?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupCustomMessage>[]>;
    /**
     * Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
     */
    deviceDetectionPortals?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupDeviceDetectionPortal>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `ec` block is documented below.
     */
    ecs?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupEc>[]>;
    /**
     * Replacement message table entries. The structure of `fortiguardWf` block is documented below.
     */
    fortiguardWfs?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupFortiguardWf>[]>;
    /**
     * Replacement message table entries. The structure of `ftp` block is documented below.
     */
    ftps?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupFtp>[]>;
    /**
     * Group type.
     */
    groupType?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `http` block is documented below.
     */
    https?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupHttp>[]>;
    /**
     * Replacement message table entries. The structure of `icap` block is documented below.
     */
    icaps?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupIcap>[]>;
    /**
     * Replacement message table entries. The structure of `mail` block is documented below.
     */
    mails?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupMail>[]>;
    /**
     * Replacement message table entries. The structure of `nacQuar` block is documented below.
     */
    nacQuars?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupNacQuar>[]>;
    /**
     * Group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `nntp` block is documented below.
     */
    nntps?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupNntp>[]>;
    /**
     * Replacement message table entries. The structure of `spam` block is documented below.
     */
    spams?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupSpam>[]>;
    /**
     * Replacement message table entries. The structure of `sslvpn` block is documented below.
     */
    sslvpns?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupSslvpn>[]>;
    /**
     * Replacement message table entries. The structure of `trafficQuota` block is documented below.
     */
    trafficQuotas?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupTrafficQuota>[]>;
    /**
     * Replacement message table entries. The structure of `utm` block is documented below.
     */
    utms?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupUtm>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `webproxy` block is documented below.
     */
    webproxies?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupWebproxy>[]>;
}

/**
 * The set of arguments for constructing a SystemReplacemsgGroup resource.
 */
export interface SystemReplacemsgGroupArgs {
    /**
     * Replacement message table entries. The structure of `admin` block is documented below.
     */
    admins?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupAdmin>[]>;
    /**
     * Replacement message table entries. The structure of `alertmail` block is documented below.
     */
    alertmails?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupAlertmail>[]>;
    /**
     * Replacement message table entries. The structure of `auth` block is documented below.
     */
    auths?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupAuth>[]>;
    /**
     * Replacement message table entries. The structure of `automation` block is documented below.
     */
    automations?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupAutomation>[]>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `customMessage` block is documented below.
     */
    customMessages?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupCustomMessage>[]>;
    /**
     * Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
     */
    deviceDetectionPortals?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupDeviceDetectionPortal>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `ec` block is documented below.
     */
    ecs?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupEc>[]>;
    /**
     * Replacement message table entries. The structure of `fortiguardWf` block is documented below.
     */
    fortiguardWfs?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupFortiguardWf>[]>;
    /**
     * Replacement message table entries. The structure of `ftp` block is documented below.
     */
    ftps?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupFtp>[]>;
    /**
     * Group type.
     */
    groupType: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `http` block is documented below.
     */
    https?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupHttp>[]>;
    /**
     * Replacement message table entries. The structure of `icap` block is documented below.
     */
    icaps?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupIcap>[]>;
    /**
     * Replacement message table entries. The structure of `mail` block is documented below.
     */
    mails?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupMail>[]>;
    /**
     * Replacement message table entries. The structure of `nacQuar` block is documented below.
     */
    nacQuars?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupNacQuar>[]>;
    /**
     * Group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `nntp` block is documented below.
     */
    nntps?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupNntp>[]>;
    /**
     * Replacement message table entries. The structure of `spam` block is documented below.
     */
    spams?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupSpam>[]>;
    /**
     * Replacement message table entries. The structure of `sslvpn` block is documented below.
     */
    sslvpns?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupSslvpn>[]>;
    /**
     * Replacement message table entries. The structure of `trafficQuota` block is documented below.
     */
    trafficQuotas?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupTrafficQuota>[]>;
    /**
     * Replacement message table entries. The structure of `utm` block is documented below.
     */
    utms?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupUtm>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Replacement message table entries. The structure of `webproxy` block is documented below.
     */
    webproxies?: pulumi.Input<pulumi.Input<inputs.SystemReplacemsgGroupWebproxy>[]>;
}
