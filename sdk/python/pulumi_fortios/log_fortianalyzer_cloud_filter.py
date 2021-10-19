# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LogFortianalyzerCloudFilterArgs', 'LogFortianalyzerCloudFilter']

@pulumi.input_type
class LogFortianalyzerCloudFilterArgs:
    def __init__(__self__, *,
                 anomaly: Optional[pulumi.Input[str]] = None,
                 dlp_archive: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 forward_traffic: Optional[pulumi.Input[str]] = None,
                 free_styles: Optional[pulumi.Input[Sequence[pulumi.Input['LogFortianalyzerCloudFilterFreeStyleArgs']]]] = None,
                 gtp: Optional[pulumi.Input[str]] = None,
                 local_traffic: Optional[pulumi.Input[str]] = None,
                 multicast_traffic: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sniffer_traffic: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogFortianalyzerCloudFilter resource.
        :param pulumi.Input[str] anomaly: Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dlp_archive: Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] filter: Free style filter string.
        :param pulumi.Input[str] filter_type: Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        :param pulumi.Input[str] forward_traffic: Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['LogFortianalyzerCloudFilterFreeStyleArgs']]] free_styles: Free Style Filters The structure of `free_style` block is documented below.
        :param pulumi.Input[str] gtp: Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] local_traffic: Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] multicast_traffic: Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] severity: Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        :param pulumi.Input[str] sniffer_traffic: Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] voip: Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        """
        if anomaly is not None:
            pulumi.set(__self__, "anomaly", anomaly)
        if dlp_archive is not None:
            pulumi.set(__self__, "dlp_archive", dlp_archive)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if filter_type is not None:
            pulumi.set(__self__, "filter_type", filter_type)
        if forward_traffic is not None:
            pulumi.set(__self__, "forward_traffic", forward_traffic)
        if free_styles is not None:
            pulumi.set(__self__, "free_styles", free_styles)
        if gtp is not None:
            pulumi.set(__self__, "gtp", gtp)
        if local_traffic is not None:
            pulumi.set(__self__, "local_traffic", local_traffic)
        if multicast_traffic is not None:
            pulumi.set(__self__, "multicast_traffic", multicast_traffic)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if sniffer_traffic is not None:
            pulumi.set(__self__, "sniffer_traffic", sniffer_traffic)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if voip is not None:
            pulumi.set(__self__, "voip", voip)

    @property
    @pulumi.getter
    def anomaly(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "anomaly")

    @anomaly.setter
    def anomaly(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "anomaly", value)

    @property
    @pulumi.getter(name="dlpArchive")
    def dlp_archive(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dlp_archive")

    @dlp_archive.setter
    def dlp_archive(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dlp_archive", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Free style filter string.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> Optional[pulumi.Input[str]]:
        """
        Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        """
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter(name="forwardTraffic")
    def forward_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "forward_traffic")

    @forward_traffic.setter
    def forward_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forward_traffic", value)

    @property
    @pulumi.getter(name="freeStyles")
    def free_styles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LogFortianalyzerCloudFilterFreeStyleArgs']]]]:
        """
        Free Style Filters The structure of `free_style` block is documented below.
        """
        return pulumi.get(self, "free_styles")

    @free_styles.setter
    def free_styles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LogFortianalyzerCloudFilterFreeStyleArgs']]]]):
        pulumi.set(self, "free_styles", value)

    @property
    @pulumi.getter
    def gtp(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "gtp")

    @gtp.setter
    def gtp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gtp", value)

    @property
    @pulumi.getter(name="localTraffic")
    def local_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "local_traffic")

    @local_traffic.setter
    def local_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_traffic", value)

    @property
    @pulumi.getter(name="multicastTraffic")
    def multicast_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_traffic")

    @multicast_traffic.setter
    def multicast_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_traffic", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter(name="snifferTraffic")
    def sniffer_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "sniffer_traffic")

    @sniffer_traffic.setter
    def sniffer_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sniffer_traffic", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def voip(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "voip")

    @voip.setter
    def voip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voip", value)


@pulumi.input_type
class _LogFortianalyzerCloudFilterState:
    def __init__(__self__, *,
                 anomaly: Optional[pulumi.Input[str]] = None,
                 dlp_archive: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 forward_traffic: Optional[pulumi.Input[str]] = None,
                 free_styles: Optional[pulumi.Input[Sequence[pulumi.Input['LogFortianalyzerCloudFilterFreeStyleArgs']]]] = None,
                 gtp: Optional[pulumi.Input[str]] = None,
                 local_traffic: Optional[pulumi.Input[str]] = None,
                 multicast_traffic: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sniffer_traffic: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogFortianalyzerCloudFilter resources.
        :param pulumi.Input[str] anomaly: Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dlp_archive: Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] filter: Free style filter string.
        :param pulumi.Input[str] filter_type: Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        :param pulumi.Input[str] forward_traffic: Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['LogFortianalyzerCloudFilterFreeStyleArgs']]] free_styles: Free Style Filters The structure of `free_style` block is documented below.
        :param pulumi.Input[str] gtp: Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] local_traffic: Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] multicast_traffic: Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] severity: Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        :param pulumi.Input[str] sniffer_traffic: Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] voip: Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        """
        if anomaly is not None:
            pulumi.set(__self__, "anomaly", anomaly)
        if dlp_archive is not None:
            pulumi.set(__self__, "dlp_archive", dlp_archive)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if filter_type is not None:
            pulumi.set(__self__, "filter_type", filter_type)
        if forward_traffic is not None:
            pulumi.set(__self__, "forward_traffic", forward_traffic)
        if free_styles is not None:
            pulumi.set(__self__, "free_styles", free_styles)
        if gtp is not None:
            pulumi.set(__self__, "gtp", gtp)
        if local_traffic is not None:
            pulumi.set(__self__, "local_traffic", local_traffic)
        if multicast_traffic is not None:
            pulumi.set(__self__, "multicast_traffic", multicast_traffic)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if sniffer_traffic is not None:
            pulumi.set(__self__, "sniffer_traffic", sniffer_traffic)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if voip is not None:
            pulumi.set(__self__, "voip", voip)

    @property
    @pulumi.getter
    def anomaly(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "anomaly")

    @anomaly.setter
    def anomaly(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "anomaly", value)

    @property
    @pulumi.getter(name="dlpArchive")
    def dlp_archive(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dlp_archive")

    @dlp_archive.setter
    def dlp_archive(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dlp_archive", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Free style filter string.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> Optional[pulumi.Input[str]]:
        """
        Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        """
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter(name="forwardTraffic")
    def forward_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "forward_traffic")

    @forward_traffic.setter
    def forward_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forward_traffic", value)

    @property
    @pulumi.getter(name="freeStyles")
    def free_styles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LogFortianalyzerCloudFilterFreeStyleArgs']]]]:
        """
        Free Style Filters The structure of `free_style` block is documented below.
        """
        return pulumi.get(self, "free_styles")

    @free_styles.setter
    def free_styles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LogFortianalyzerCloudFilterFreeStyleArgs']]]]):
        pulumi.set(self, "free_styles", value)

    @property
    @pulumi.getter
    def gtp(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "gtp")

    @gtp.setter
    def gtp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gtp", value)

    @property
    @pulumi.getter(name="localTraffic")
    def local_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "local_traffic")

    @local_traffic.setter
    def local_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_traffic", value)

    @property
    @pulumi.getter(name="multicastTraffic")
    def multicast_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_traffic")

    @multicast_traffic.setter
    def multicast_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_traffic", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter(name="snifferTraffic")
    def sniffer_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "sniffer_traffic")

    @sniffer_traffic.setter
    def sniffer_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sniffer_traffic", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def voip(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "voip")

    @voip.setter
    def voip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voip", value)


class LogFortianalyzerCloudFilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anomaly: Optional[pulumi.Input[str]] = None,
                 dlp_archive: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 forward_traffic: Optional[pulumi.Input[str]] = None,
                 free_styles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogFortianalyzerCloudFilterFreeStyleArgs']]]]] = None,
                 gtp: Optional[pulumi.Input[str]] = None,
                 local_traffic: Optional[pulumi.Input[str]] = None,
                 multicast_traffic: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sniffer_traffic: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Filters for FortiAnalyzer Cloud.

        ## Import

        LogFortianalyzerCloud Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/logFortianalyzerCloudFilter:LogFortianalyzerCloudFilter labelname LogFortianalyzerCloudFilter
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] anomaly: Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dlp_archive: Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] filter: Free style filter string.
        :param pulumi.Input[str] filter_type: Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        :param pulumi.Input[str] forward_traffic: Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogFortianalyzerCloudFilterFreeStyleArgs']]]] free_styles: Free Style Filters The structure of `free_style` block is documented below.
        :param pulumi.Input[str] gtp: Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] local_traffic: Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] multicast_traffic: Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] severity: Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        :param pulumi.Input[str] sniffer_traffic: Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] voip: Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LogFortianalyzerCloudFilterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Filters for FortiAnalyzer Cloud.

        ## Import

        LogFortianalyzerCloud Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/logFortianalyzerCloudFilter:LogFortianalyzerCloudFilter labelname LogFortianalyzerCloudFilter
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param LogFortianalyzerCloudFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogFortianalyzerCloudFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anomaly: Optional[pulumi.Input[str]] = None,
                 dlp_archive: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 forward_traffic: Optional[pulumi.Input[str]] = None,
                 free_styles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogFortianalyzerCloudFilterFreeStyleArgs']]]]] = None,
                 gtp: Optional[pulumi.Input[str]] = None,
                 local_traffic: Optional[pulumi.Input[str]] = None,
                 multicast_traffic: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sniffer_traffic: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogFortianalyzerCloudFilterArgs.__new__(LogFortianalyzerCloudFilterArgs)

            __props__.__dict__["anomaly"] = anomaly
            __props__.__dict__["dlp_archive"] = dlp_archive
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["filter"] = filter
            __props__.__dict__["filter_type"] = filter_type
            __props__.__dict__["forward_traffic"] = forward_traffic
            __props__.__dict__["free_styles"] = free_styles
            __props__.__dict__["gtp"] = gtp
            __props__.__dict__["local_traffic"] = local_traffic
            __props__.__dict__["multicast_traffic"] = multicast_traffic
            __props__.__dict__["severity"] = severity
            __props__.__dict__["sniffer_traffic"] = sniffer_traffic
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["voip"] = voip
        super(LogFortianalyzerCloudFilter, __self__).__init__(
            'fortios:index/logFortianalyzerCloudFilter:LogFortianalyzerCloudFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            anomaly: Optional[pulumi.Input[str]] = None,
            dlp_archive: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            filter_type: Optional[pulumi.Input[str]] = None,
            forward_traffic: Optional[pulumi.Input[str]] = None,
            free_styles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogFortianalyzerCloudFilterFreeStyleArgs']]]]] = None,
            gtp: Optional[pulumi.Input[str]] = None,
            local_traffic: Optional[pulumi.Input[str]] = None,
            multicast_traffic: Optional[pulumi.Input[str]] = None,
            severity: Optional[pulumi.Input[str]] = None,
            sniffer_traffic: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            voip: Optional[pulumi.Input[str]] = None) -> 'LogFortianalyzerCloudFilter':
        """
        Get an existing LogFortianalyzerCloudFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] anomaly: Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dlp_archive: Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] filter: Free style filter string.
        :param pulumi.Input[str] filter_type: Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        :param pulumi.Input[str] forward_traffic: Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogFortianalyzerCloudFilterFreeStyleArgs']]]] free_styles: Free Style Filters The structure of `free_style` block is documented below.
        :param pulumi.Input[str] gtp: Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] local_traffic: Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] multicast_traffic: Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] severity: Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        :param pulumi.Input[str] sniffer_traffic: Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] voip: Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogFortianalyzerCloudFilterState.__new__(_LogFortianalyzerCloudFilterState)

        __props__.__dict__["anomaly"] = anomaly
        __props__.__dict__["dlp_archive"] = dlp_archive
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["filter"] = filter
        __props__.__dict__["filter_type"] = filter_type
        __props__.__dict__["forward_traffic"] = forward_traffic
        __props__.__dict__["free_styles"] = free_styles
        __props__.__dict__["gtp"] = gtp
        __props__.__dict__["local_traffic"] = local_traffic
        __props__.__dict__["multicast_traffic"] = multicast_traffic
        __props__.__dict__["severity"] = severity
        __props__.__dict__["sniffer_traffic"] = sniffer_traffic
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["voip"] = voip
        return LogFortianalyzerCloudFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def anomaly(self) -> pulumi.Output[str]:
        """
        Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "anomaly")

    @property
    @pulumi.getter(name="dlpArchive")
    def dlp_archive(self) -> pulumi.Output[str]:
        """
        Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dlp_archive")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[str]:
        """
        Free style filter string.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> pulumi.Output[str]:
        """
        Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        """
        return pulumi.get(self, "filter_type")

    @property
    @pulumi.getter(name="forwardTraffic")
    def forward_traffic(self) -> pulumi.Output[str]:
        """
        Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "forward_traffic")

    @property
    @pulumi.getter(name="freeStyles")
    def free_styles(self) -> pulumi.Output[Optional[Sequence['outputs.LogFortianalyzerCloudFilterFreeStyle']]]:
        """
        Free Style Filters The structure of `free_style` block is documented below.
        """
        return pulumi.get(self, "free_styles")

    @property
    @pulumi.getter
    def gtp(self) -> pulumi.Output[str]:
        """
        Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "gtp")

    @property
    @pulumi.getter(name="localTraffic")
    def local_traffic(self) -> pulumi.Output[str]:
        """
        Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "local_traffic")

    @property
    @pulumi.getter(name="multicastTraffic")
    def multicast_traffic(self) -> pulumi.Output[str]:
        """
        Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_traffic")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        """
        Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter(name="snifferTraffic")
    def sniffer_traffic(self) -> pulumi.Output[str]:
        """
        Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "sniffer_traffic")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def voip(self) -> pulumi.Output[str]:
        """
        Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "voip")

