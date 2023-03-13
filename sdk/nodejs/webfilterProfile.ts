// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WebfilterProfile extends pulumi.CustomResource {
    /**
     * Get an existing WebfilterProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebfilterProfileState, opts?: pulumi.CustomResourceOptions): WebfilterProfile {
        return new WebfilterProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webfilterProfile:WebfilterProfile';

    /**
     * Returns true if the given object is an instance of WebfilterProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebfilterProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebfilterProfile.__pulumiType;
    }

    public readonly antiphish!: pulumi.Output<outputs.WebfilterProfileAntiphish>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly extendedLog!: pulumi.Output<string>;
    public readonly featureSet!: pulumi.Output<string>;
    public readonly fileFilter!: pulumi.Output<outputs.WebfilterProfileFileFilter>;
    public readonly ftgdWf!: pulumi.Output<outputs.WebfilterProfileFtgdWf>;
    public readonly httpsReplacemsg!: pulumi.Output<string>;
    public readonly inspectionMode!: pulumi.Output<string>;
    public readonly logAllUrl!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly options!: pulumi.Output<string>;
    public readonly override!: pulumi.Output<outputs.WebfilterProfileOverride>;
    public readonly ovrdPerm!: pulumi.Output<string>;
    public readonly postAction!: pulumi.Output<string>;
    public readonly replacemsgGroup!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly web!: pulumi.Output<outputs.WebfilterProfileWeb>;
    public readonly webAntiphishingLog!: pulumi.Output<string>;
    public readonly webContentLog!: pulumi.Output<string>;
    public readonly webExtendedAllActionLog!: pulumi.Output<string>;
    public readonly webFilterActivexLog!: pulumi.Output<string>;
    public readonly webFilterAppletLog!: pulumi.Output<string>;
    public readonly webFilterCommandBlockLog!: pulumi.Output<string>;
    public readonly webFilterCookieLog!: pulumi.Output<string>;
    public readonly webFilterCookieRemovalLog!: pulumi.Output<string>;
    public readonly webFilterJsLog!: pulumi.Output<string>;
    public readonly webFilterJscriptLog!: pulumi.Output<string>;
    public readonly webFilterRefererLog!: pulumi.Output<string>;
    public readonly webFilterUnknownLog!: pulumi.Output<string>;
    public readonly webFilterVbsLog!: pulumi.Output<string>;
    public readonly webFtgdErrLog!: pulumi.Output<string>;
    public readonly webFtgdQuotaUsage!: pulumi.Output<string>;
    public readonly webInvalidDomainLog!: pulumi.Output<string>;
    public readonly webUrlLog!: pulumi.Output<string>;
    public readonly wisp!: pulumi.Output<string>;
    public readonly wispAlgorithm!: pulumi.Output<string>;
    public readonly wispServers!: pulumi.Output<outputs.WebfilterProfileWispServer[] | undefined>;
    public readonly youtubeChannelFilters!: pulumi.Output<outputs.WebfilterProfileYoutubeChannelFilter[] | undefined>;
    public readonly youtubeChannelStatus!: pulumi.Output<string>;

    /**
     * Create a WebfilterProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebfilterProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebfilterProfileArgs | WebfilterProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebfilterProfileState | undefined;
            resourceInputs["antiphish"] = state ? state.antiphish : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["fileFilter"] = state ? state.fileFilter : undefined;
            resourceInputs["ftgdWf"] = state ? state.ftgdWf : undefined;
            resourceInputs["httpsReplacemsg"] = state ? state.httpsReplacemsg : undefined;
            resourceInputs["inspectionMode"] = state ? state.inspectionMode : undefined;
            resourceInputs["logAllUrl"] = state ? state.logAllUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["override"] = state ? state.override : undefined;
            resourceInputs["ovrdPerm"] = state ? state.ovrdPerm : undefined;
            resourceInputs["postAction"] = state ? state.postAction : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["web"] = state ? state.web : undefined;
            resourceInputs["webAntiphishingLog"] = state ? state.webAntiphishingLog : undefined;
            resourceInputs["webContentLog"] = state ? state.webContentLog : undefined;
            resourceInputs["webExtendedAllActionLog"] = state ? state.webExtendedAllActionLog : undefined;
            resourceInputs["webFilterActivexLog"] = state ? state.webFilterActivexLog : undefined;
            resourceInputs["webFilterAppletLog"] = state ? state.webFilterAppletLog : undefined;
            resourceInputs["webFilterCommandBlockLog"] = state ? state.webFilterCommandBlockLog : undefined;
            resourceInputs["webFilterCookieLog"] = state ? state.webFilterCookieLog : undefined;
            resourceInputs["webFilterCookieRemovalLog"] = state ? state.webFilterCookieRemovalLog : undefined;
            resourceInputs["webFilterJsLog"] = state ? state.webFilterJsLog : undefined;
            resourceInputs["webFilterJscriptLog"] = state ? state.webFilterJscriptLog : undefined;
            resourceInputs["webFilterRefererLog"] = state ? state.webFilterRefererLog : undefined;
            resourceInputs["webFilterUnknownLog"] = state ? state.webFilterUnknownLog : undefined;
            resourceInputs["webFilterVbsLog"] = state ? state.webFilterVbsLog : undefined;
            resourceInputs["webFtgdErrLog"] = state ? state.webFtgdErrLog : undefined;
            resourceInputs["webFtgdQuotaUsage"] = state ? state.webFtgdQuotaUsage : undefined;
            resourceInputs["webInvalidDomainLog"] = state ? state.webInvalidDomainLog : undefined;
            resourceInputs["webUrlLog"] = state ? state.webUrlLog : undefined;
            resourceInputs["wisp"] = state ? state.wisp : undefined;
            resourceInputs["wispAlgorithm"] = state ? state.wispAlgorithm : undefined;
            resourceInputs["wispServers"] = state ? state.wispServers : undefined;
            resourceInputs["youtubeChannelFilters"] = state ? state.youtubeChannelFilters : undefined;
            resourceInputs["youtubeChannelStatus"] = state ? state.youtubeChannelStatus : undefined;
        } else {
            const args = argsOrState as WebfilterProfileArgs | undefined;
            resourceInputs["antiphish"] = args ? args.antiphish : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["fileFilter"] = args ? args.fileFilter : undefined;
            resourceInputs["ftgdWf"] = args ? args.ftgdWf : undefined;
            resourceInputs["httpsReplacemsg"] = args ? args.httpsReplacemsg : undefined;
            resourceInputs["inspectionMode"] = args ? args.inspectionMode : undefined;
            resourceInputs["logAllUrl"] = args ? args.logAllUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["override"] = args ? args.override : undefined;
            resourceInputs["ovrdPerm"] = args ? args.ovrdPerm : undefined;
            resourceInputs["postAction"] = args ? args.postAction : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["web"] = args ? args.web : undefined;
            resourceInputs["webAntiphishingLog"] = args ? args.webAntiphishingLog : undefined;
            resourceInputs["webContentLog"] = args ? args.webContentLog : undefined;
            resourceInputs["webExtendedAllActionLog"] = args ? args.webExtendedAllActionLog : undefined;
            resourceInputs["webFilterActivexLog"] = args ? args.webFilterActivexLog : undefined;
            resourceInputs["webFilterAppletLog"] = args ? args.webFilterAppletLog : undefined;
            resourceInputs["webFilterCommandBlockLog"] = args ? args.webFilterCommandBlockLog : undefined;
            resourceInputs["webFilterCookieLog"] = args ? args.webFilterCookieLog : undefined;
            resourceInputs["webFilterCookieRemovalLog"] = args ? args.webFilterCookieRemovalLog : undefined;
            resourceInputs["webFilterJsLog"] = args ? args.webFilterJsLog : undefined;
            resourceInputs["webFilterJscriptLog"] = args ? args.webFilterJscriptLog : undefined;
            resourceInputs["webFilterRefererLog"] = args ? args.webFilterRefererLog : undefined;
            resourceInputs["webFilterUnknownLog"] = args ? args.webFilterUnknownLog : undefined;
            resourceInputs["webFilterVbsLog"] = args ? args.webFilterVbsLog : undefined;
            resourceInputs["webFtgdErrLog"] = args ? args.webFtgdErrLog : undefined;
            resourceInputs["webFtgdQuotaUsage"] = args ? args.webFtgdQuotaUsage : undefined;
            resourceInputs["webInvalidDomainLog"] = args ? args.webInvalidDomainLog : undefined;
            resourceInputs["webUrlLog"] = args ? args.webUrlLog : undefined;
            resourceInputs["wisp"] = args ? args.wisp : undefined;
            resourceInputs["wispAlgorithm"] = args ? args.wispAlgorithm : undefined;
            resourceInputs["wispServers"] = args ? args.wispServers : undefined;
            resourceInputs["youtubeChannelFilters"] = args ? args.youtubeChannelFilters : undefined;
            resourceInputs["youtubeChannelStatus"] = args ? args.youtubeChannelStatus : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebfilterProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebfilterProfile resources.
 */
export interface WebfilterProfileState {
    antiphish?: pulumi.Input<inputs.WebfilterProfileAntiphish>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    extendedLog?: pulumi.Input<string>;
    featureSet?: pulumi.Input<string>;
    fileFilter?: pulumi.Input<inputs.WebfilterProfileFileFilter>;
    ftgdWf?: pulumi.Input<inputs.WebfilterProfileFtgdWf>;
    httpsReplacemsg?: pulumi.Input<string>;
    inspectionMode?: pulumi.Input<string>;
    logAllUrl?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    options?: pulumi.Input<string>;
    override?: pulumi.Input<inputs.WebfilterProfileOverride>;
    ovrdPerm?: pulumi.Input<string>;
    postAction?: pulumi.Input<string>;
    replacemsgGroup?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    web?: pulumi.Input<inputs.WebfilterProfileWeb>;
    webAntiphishingLog?: pulumi.Input<string>;
    webContentLog?: pulumi.Input<string>;
    webExtendedAllActionLog?: pulumi.Input<string>;
    webFilterActivexLog?: pulumi.Input<string>;
    webFilterAppletLog?: pulumi.Input<string>;
    webFilterCommandBlockLog?: pulumi.Input<string>;
    webFilterCookieLog?: pulumi.Input<string>;
    webFilterCookieRemovalLog?: pulumi.Input<string>;
    webFilterJsLog?: pulumi.Input<string>;
    webFilterJscriptLog?: pulumi.Input<string>;
    webFilterRefererLog?: pulumi.Input<string>;
    webFilterUnknownLog?: pulumi.Input<string>;
    webFilterVbsLog?: pulumi.Input<string>;
    webFtgdErrLog?: pulumi.Input<string>;
    webFtgdQuotaUsage?: pulumi.Input<string>;
    webInvalidDomainLog?: pulumi.Input<string>;
    webUrlLog?: pulumi.Input<string>;
    wisp?: pulumi.Input<string>;
    wispAlgorithm?: pulumi.Input<string>;
    wispServers?: pulumi.Input<pulumi.Input<inputs.WebfilterProfileWispServer>[]>;
    youtubeChannelFilters?: pulumi.Input<pulumi.Input<inputs.WebfilterProfileYoutubeChannelFilter>[]>;
    youtubeChannelStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebfilterProfile resource.
 */
export interface WebfilterProfileArgs {
    antiphish?: pulumi.Input<inputs.WebfilterProfileAntiphish>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    extendedLog?: pulumi.Input<string>;
    featureSet?: pulumi.Input<string>;
    fileFilter?: pulumi.Input<inputs.WebfilterProfileFileFilter>;
    ftgdWf?: pulumi.Input<inputs.WebfilterProfileFtgdWf>;
    httpsReplacemsg?: pulumi.Input<string>;
    inspectionMode?: pulumi.Input<string>;
    logAllUrl?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    options?: pulumi.Input<string>;
    override?: pulumi.Input<inputs.WebfilterProfileOverride>;
    ovrdPerm?: pulumi.Input<string>;
    postAction?: pulumi.Input<string>;
    replacemsgGroup?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    web?: pulumi.Input<inputs.WebfilterProfileWeb>;
    webAntiphishingLog?: pulumi.Input<string>;
    webContentLog?: pulumi.Input<string>;
    webExtendedAllActionLog?: pulumi.Input<string>;
    webFilterActivexLog?: pulumi.Input<string>;
    webFilterAppletLog?: pulumi.Input<string>;
    webFilterCommandBlockLog?: pulumi.Input<string>;
    webFilterCookieLog?: pulumi.Input<string>;
    webFilterCookieRemovalLog?: pulumi.Input<string>;
    webFilterJsLog?: pulumi.Input<string>;
    webFilterJscriptLog?: pulumi.Input<string>;
    webFilterRefererLog?: pulumi.Input<string>;
    webFilterUnknownLog?: pulumi.Input<string>;
    webFilterVbsLog?: pulumi.Input<string>;
    webFtgdErrLog?: pulumi.Input<string>;
    webFtgdQuotaUsage?: pulumi.Input<string>;
    webInvalidDomainLog?: pulumi.Input<string>;
    webUrlLog?: pulumi.Input<string>;
    wisp?: pulumi.Input<string>;
    wispAlgorithm?: pulumi.Input<string>;
    wispServers?: pulumi.Input<pulumi.Input<inputs.WebfilterProfileWispServer>[]>;
    youtubeChannelFilters?: pulumi.Input<pulumi.Input<inputs.WebfilterProfileYoutubeChannelFilter>[]>;
    youtubeChannelStatus?: pulumi.Input<string>;
}
