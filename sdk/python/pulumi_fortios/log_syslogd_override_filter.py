# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LogSyslogdOverrideFilterArgs', 'LogSyslogdOverrideFilter']

@pulumi.input_type
class LogSyslogdOverrideFilterArgs:
    def __init__(__self__, *,
                 anomaly: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 forward_traffic: Optional[pulumi.Input[str]] = None,
                 free_styles: Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdOverrideFilterFreeStyleArgs']]]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 gtp: Optional[pulumi.Input[str]] = None,
                 local_traffic: Optional[pulumi.Input[str]] = None,
                 multicast_traffic: Optional[pulumi.Input[str]] = None,
                 netscan_discovery: Optional[pulumi.Input[str]] = None,
                 netscan_vulnerability: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sniffer_traffic: Optional[pulumi.Input[str]] = None,
                 ssh: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip: Optional[pulumi.Input[str]] = None,
                 ztna_traffic: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogSyslogdOverrideFilter resource.
        """
        if anomaly is not None:
            pulumi.set(__self__, "anomaly", anomaly)
        if dns is not None:
            pulumi.set(__self__, "dns", dns)
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
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if gtp is not None:
            pulumi.set(__self__, "gtp", gtp)
        if local_traffic is not None:
            pulumi.set(__self__, "local_traffic", local_traffic)
        if multicast_traffic is not None:
            pulumi.set(__self__, "multicast_traffic", multicast_traffic)
        if netscan_discovery is not None:
            pulumi.set(__self__, "netscan_discovery", netscan_discovery)
        if netscan_vulnerability is not None:
            pulumi.set(__self__, "netscan_vulnerability", netscan_vulnerability)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if sniffer_traffic is not None:
            pulumi.set(__self__, "sniffer_traffic", sniffer_traffic)
        if ssh is not None:
            pulumi.set(__self__, "ssh", ssh)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if voip is not None:
            pulumi.set(__self__, "voip", voip)
        if ztna_traffic is not None:
            pulumi.set(__self__, "ztna_traffic", ztna_traffic)

    @property
    @pulumi.getter
    def anomaly(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "anomaly")

    @anomaly.setter
    def anomaly(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "anomaly", value)

    @property
    @pulumi.getter
    def dns(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dns")

    @dns.setter
    def dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter(name="forwardTraffic")
    def forward_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forward_traffic")

    @forward_traffic.setter
    def forward_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forward_traffic", value)

    @property
    @pulumi.getter(name="freeStyles")
    def free_styles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdOverrideFilterFreeStyleArgs']]]]:
        return pulumi.get(self, "free_styles")

    @free_styles.setter
    def free_styles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdOverrideFilterFreeStyleArgs']]]]):
        pulumi.set(self, "free_styles", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def gtp(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gtp")

    @gtp.setter
    def gtp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gtp", value)

    @property
    @pulumi.getter(name="localTraffic")
    def local_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "local_traffic")

    @local_traffic.setter
    def local_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_traffic", value)

    @property
    @pulumi.getter(name="multicastTraffic")
    def multicast_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "multicast_traffic")

    @multicast_traffic.setter
    def multicast_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_traffic", value)

    @property
    @pulumi.getter(name="netscanDiscovery")
    def netscan_discovery(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "netscan_discovery")

    @netscan_discovery.setter
    def netscan_discovery(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netscan_discovery", value)

    @property
    @pulumi.getter(name="netscanVulnerability")
    def netscan_vulnerability(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "netscan_vulnerability")

    @netscan_vulnerability.setter
    def netscan_vulnerability(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netscan_vulnerability", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter(name="snifferTraffic")
    def sniffer_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sniffer_traffic")

    @sniffer_traffic.setter
    def sniffer_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sniffer_traffic", value)

    @property
    @pulumi.getter
    def ssh(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssh")

    @ssh.setter
    def ssh(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def voip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "voip")

    @voip.setter
    def voip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voip", value)

    @property
    @pulumi.getter(name="ztnaTraffic")
    def ztna_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ztna_traffic")

    @ztna_traffic.setter
    def ztna_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ztna_traffic", value)


@pulumi.input_type
class _LogSyslogdOverrideFilterState:
    def __init__(__self__, *,
                 anomaly: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 forward_traffic: Optional[pulumi.Input[str]] = None,
                 free_styles: Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdOverrideFilterFreeStyleArgs']]]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 gtp: Optional[pulumi.Input[str]] = None,
                 local_traffic: Optional[pulumi.Input[str]] = None,
                 multicast_traffic: Optional[pulumi.Input[str]] = None,
                 netscan_discovery: Optional[pulumi.Input[str]] = None,
                 netscan_vulnerability: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sniffer_traffic: Optional[pulumi.Input[str]] = None,
                 ssh: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip: Optional[pulumi.Input[str]] = None,
                 ztna_traffic: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogSyslogdOverrideFilter resources.
        """
        if anomaly is not None:
            pulumi.set(__self__, "anomaly", anomaly)
        if dns is not None:
            pulumi.set(__self__, "dns", dns)
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
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if gtp is not None:
            pulumi.set(__self__, "gtp", gtp)
        if local_traffic is not None:
            pulumi.set(__self__, "local_traffic", local_traffic)
        if multicast_traffic is not None:
            pulumi.set(__self__, "multicast_traffic", multicast_traffic)
        if netscan_discovery is not None:
            pulumi.set(__self__, "netscan_discovery", netscan_discovery)
        if netscan_vulnerability is not None:
            pulumi.set(__self__, "netscan_vulnerability", netscan_vulnerability)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if sniffer_traffic is not None:
            pulumi.set(__self__, "sniffer_traffic", sniffer_traffic)
        if ssh is not None:
            pulumi.set(__self__, "ssh", ssh)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if voip is not None:
            pulumi.set(__self__, "voip", voip)
        if ztna_traffic is not None:
            pulumi.set(__self__, "ztna_traffic", ztna_traffic)

    @property
    @pulumi.getter
    def anomaly(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "anomaly")

    @anomaly.setter
    def anomaly(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "anomaly", value)

    @property
    @pulumi.getter
    def dns(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dns")

    @dns.setter
    def dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter(name="forwardTraffic")
    def forward_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forward_traffic")

    @forward_traffic.setter
    def forward_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forward_traffic", value)

    @property
    @pulumi.getter(name="freeStyles")
    def free_styles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdOverrideFilterFreeStyleArgs']]]]:
        return pulumi.get(self, "free_styles")

    @free_styles.setter
    def free_styles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdOverrideFilterFreeStyleArgs']]]]):
        pulumi.set(self, "free_styles", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def gtp(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gtp")

    @gtp.setter
    def gtp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gtp", value)

    @property
    @pulumi.getter(name="localTraffic")
    def local_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "local_traffic")

    @local_traffic.setter
    def local_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_traffic", value)

    @property
    @pulumi.getter(name="multicastTraffic")
    def multicast_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "multicast_traffic")

    @multicast_traffic.setter
    def multicast_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_traffic", value)

    @property
    @pulumi.getter(name="netscanDiscovery")
    def netscan_discovery(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "netscan_discovery")

    @netscan_discovery.setter
    def netscan_discovery(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netscan_discovery", value)

    @property
    @pulumi.getter(name="netscanVulnerability")
    def netscan_vulnerability(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "netscan_vulnerability")

    @netscan_vulnerability.setter
    def netscan_vulnerability(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netscan_vulnerability", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter(name="snifferTraffic")
    def sniffer_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sniffer_traffic")

    @sniffer_traffic.setter
    def sniffer_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sniffer_traffic", value)

    @property
    @pulumi.getter
    def ssh(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssh")

    @ssh.setter
    def ssh(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def voip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "voip")

    @voip.setter
    def voip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voip", value)

    @property
    @pulumi.getter(name="ztnaTraffic")
    def ztna_traffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ztna_traffic")

    @ztna_traffic.setter
    def ztna_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ztna_traffic", value)


class LogSyslogdOverrideFilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anomaly: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 forward_traffic: Optional[pulumi.Input[str]] = None,
                 free_styles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogSyslogdOverrideFilterFreeStyleArgs']]]]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 gtp: Optional[pulumi.Input[str]] = None,
                 local_traffic: Optional[pulumi.Input[str]] = None,
                 multicast_traffic: Optional[pulumi.Input[str]] = None,
                 netscan_discovery: Optional[pulumi.Input[str]] = None,
                 netscan_vulnerability: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sniffer_traffic: Optional[pulumi.Input[str]] = None,
                 ssh: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip: Optional[pulumi.Input[str]] = None,
                 ztna_traffic: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LogSyslogdOverrideFilter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LogSyslogdOverrideFilterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LogSyslogdOverrideFilter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LogSyslogdOverrideFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogSyslogdOverrideFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anomaly: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 forward_traffic: Optional[pulumi.Input[str]] = None,
                 free_styles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogSyslogdOverrideFilterFreeStyleArgs']]]]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 gtp: Optional[pulumi.Input[str]] = None,
                 local_traffic: Optional[pulumi.Input[str]] = None,
                 multicast_traffic: Optional[pulumi.Input[str]] = None,
                 netscan_discovery: Optional[pulumi.Input[str]] = None,
                 netscan_vulnerability: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sniffer_traffic: Optional[pulumi.Input[str]] = None,
                 ssh: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip: Optional[pulumi.Input[str]] = None,
                 ztna_traffic: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogSyslogdOverrideFilterArgs.__new__(LogSyslogdOverrideFilterArgs)

            __props__.__dict__["anomaly"] = anomaly
            __props__.__dict__["dns"] = dns
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["filter"] = filter
            __props__.__dict__["filter_type"] = filter_type
            __props__.__dict__["forward_traffic"] = forward_traffic
            __props__.__dict__["free_styles"] = free_styles
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["gtp"] = gtp
            __props__.__dict__["local_traffic"] = local_traffic
            __props__.__dict__["multicast_traffic"] = multicast_traffic
            __props__.__dict__["netscan_discovery"] = netscan_discovery
            __props__.__dict__["netscan_vulnerability"] = netscan_vulnerability
            __props__.__dict__["severity"] = severity
            __props__.__dict__["sniffer_traffic"] = sniffer_traffic
            __props__.__dict__["ssh"] = ssh
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["voip"] = voip
            __props__.__dict__["ztna_traffic"] = ztna_traffic
        super(LogSyslogdOverrideFilter, __self__).__init__(
            'fortios:index/logSyslogdOverrideFilter:LogSyslogdOverrideFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            anomaly: Optional[pulumi.Input[str]] = None,
            dns: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            filter_type: Optional[pulumi.Input[str]] = None,
            forward_traffic: Optional[pulumi.Input[str]] = None,
            free_styles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogSyslogdOverrideFilterFreeStyleArgs']]]]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            gtp: Optional[pulumi.Input[str]] = None,
            local_traffic: Optional[pulumi.Input[str]] = None,
            multicast_traffic: Optional[pulumi.Input[str]] = None,
            netscan_discovery: Optional[pulumi.Input[str]] = None,
            netscan_vulnerability: Optional[pulumi.Input[str]] = None,
            severity: Optional[pulumi.Input[str]] = None,
            sniffer_traffic: Optional[pulumi.Input[str]] = None,
            ssh: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            voip: Optional[pulumi.Input[str]] = None,
            ztna_traffic: Optional[pulumi.Input[str]] = None) -> 'LogSyslogdOverrideFilter':
        """
        Get an existing LogSyslogdOverrideFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogSyslogdOverrideFilterState.__new__(_LogSyslogdOverrideFilterState)

        __props__.__dict__["anomaly"] = anomaly
        __props__.__dict__["dns"] = dns
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["filter"] = filter
        __props__.__dict__["filter_type"] = filter_type
        __props__.__dict__["forward_traffic"] = forward_traffic
        __props__.__dict__["free_styles"] = free_styles
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["gtp"] = gtp
        __props__.__dict__["local_traffic"] = local_traffic
        __props__.__dict__["multicast_traffic"] = multicast_traffic
        __props__.__dict__["netscan_discovery"] = netscan_discovery
        __props__.__dict__["netscan_vulnerability"] = netscan_vulnerability
        __props__.__dict__["severity"] = severity
        __props__.__dict__["sniffer_traffic"] = sniffer_traffic
        __props__.__dict__["ssh"] = ssh
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["voip"] = voip
        __props__.__dict__["ztna_traffic"] = ztna_traffic
        return LogSyslogdOverrideFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def anomaly(self) -> pulumi.Output[str]:
        return pulumi.get(self, "anomaly")

    @property
    @pulumi.getter
    def dns(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dns")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[str]:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "filter_type")

    @property
    @pulumi.getter(name="forwardTraffic")
    def forward_traffic(self) -> pulumi.Output[str]:
        return pulumi.get(self, "forward_traffic")

    @property
    @pulumi.getter(name="freeStyles")
    def free_styles(self) -> pulumi.Output[Optional[Sequence['outputs.LogSyslogdOverrideFilterFreeStyle']]]:
        return pulumi.get(self, "free_styles")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def gtp(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gtp")

    @property
    @pulumi.getter(name="localTraffic")
    def local_traffic(self) -> pulumi.Output[str]:
        return pulumi.get(self, "local_traffic")

    @property
    @pulumi.getter(name="multicastTraffic")
    def multicast_traffic(self) -> pulumi.Output[str]:
        return pulumi.get(self, "multicast_traffic")

    @property
    @pulumi.getter(name="netscanDiscovery")
    def netscan_discovery(self) -> pulumi.Output[str]:
        return pulumi.get(self, "netscan_discovery")

    @property
    @pulumi.getter(name="netscanVulnerability")
    def netscan_vulnerability(self) -> pulumi.Output[str]:
        return pulumi.get(self, "netscan_vulnerability")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter(name="snifferTraffic")
    def sniffer_traffic(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sniffer_traffic")

    @property
    @pulumi.getter
    def ssh(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ssh")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def voip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "voip")

    @property
    @pulumi.getter(name="ztnaTraffic")
    def ztna_traffic(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ztna_traffic")

