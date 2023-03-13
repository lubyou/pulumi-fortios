# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FortimanagerFirewallObjectServiceArgs', 'FortimanagerFirewallObjectService']

@pulumi.input_type
class FortimanagerFirewallObjectServiceArgs:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 icmp_code: Optional[pulumi.Input[int]] = None,
                 icmp_type: Optional[pulumi.Input[int]] = None,
                 iprange: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_number: Optional[pulumi.Input[int]] = None,
                 proxy: Optional[pulumi.Input[str]] = None,
                 sctp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tcp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 udp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a FortimanagerFirewallObjectService resource.
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if icmp_code is not None:
            pulumi.set(__self__, "icmp_code", icmp_code)
        if icmp_type is not None:
            pulumi.set(__self__, "icmp_type", icmp_type)
        if iprange is not None:
            pulumi.set(__self__, "iprange", iprange)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if protocol_number is not None:
            pulumi.set(__self__, "protocol_number", protocol_number)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if sctp_portranges is not None:
            pulumi.set(__self__, "sctp_portranges", sctp_portranges)
        if tcp_portranges is not None:
            pulumi.set(__self__, "tcp_portranges", tcp_portranges)
        if udp_portranges is not None:
            pulumi.set(__self__, "udp_portranges", udp_portranges)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "icmp_code")

    @icmp_code.setter
    def icmp_code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_code", value)

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "icmp_type")

    @icmp_type.setter
    def icmp_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_type", value)

    @property
    @pulumi.getter
    def iprange(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iprange")

    @iprange.setter
    def iprange(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iprange", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="protocolNumber")
    def protocol_number(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "protocol_number")

    @protocol_number.setter
    def protocol_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "protocol_number", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter(name="sctpPortranges")
    def sctp_portranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "sctp_portranges")

    @sctp_portranges.setter
    def sctp_portranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "sctp_portranges", value)

    @property
    @pulumi.getter(name="tcpPortranges")
    def tcp_portranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tcp_portranges")

    @tcp_portranges.setter
    def tcp_portranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tcp_portranges", value)

    @property
    @pulumi.getter(name="udpPortranges")
    def udp_portranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "udp_portranges")

    @udp_portranges.setter
    def udp_portranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "udp_portranges", value)


@pulumi.input_type
class _FortimanagerFirewallObjectServiceState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 icmp_code: Optional[pulumi.Input[int]] = None,
                 icmp_type: Optional[pulumi.Input[int]] = None,
                 iprange: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_number: Optional[pulumi.Input[int]] = None,
                 proxy: Optional[pulumi.Input[str]] = None,
                 sctp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tcp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 udp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering FortimanagerFirewallObjectService resources.
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if icmp_code is not None:
            pulumi.set(__self__, "icmp_code", icmp_code)
        if icmp_type is not None:
            pulumi.set(__self__, "icmp_type", icmp_type)
        if iprange is not None:
            pulumi.set(__self__, "iprange", iprange)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if protocol_number is not None:
            pulumi.set(__self__, "protocol_number", protocol_number)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if sctp_portranges is not None:
            pulumi.set(__self__, "sctp_portranges", sctp_portranges)
        if tcp_portranges is not None:
            pulumi.set(__self__, "tcp_portranges", tcp_portranges)
        if udp_portranges is not None:
            pulumi.set(__self__, "udp_portranges", udp_portranges)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "icmp_code")

    @icmp_code.setter
    def icmp_code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_code", value)

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "icmp_type")

    @icmp_type.setter
    def icmp_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_type", value)

    @property
    @pulumi.getter
    def iprange(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iprange")

    @iprange.setter
    def iprange(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iprange", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="protocolNumber")
    def protocol_number(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "protocol_number")

    @protocol_number.setter
    def protocol_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "protocol_number", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter(name="sctpPortranges")
    def sctp_portranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "sctp_portranges")

    @sctp_portranges.setter
    def sctp_portranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "sctp_portranges", value)

    @property
    @pulumi.getter(name="tcpPortranges")
    def tcp_portranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tcp_portranges")

    @tcp_portranges.setter
    def tcp_portranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tcp_portranges", value)

    @property
    @pulumi.getter(name="udpPortranges")
    def udp_portranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "udp_portranges")

    @udp_portranges.setter
    def udp_portranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "udp_portranges", value)


class FortimanagerFirewallObjectService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 icmp_code: Optional[pulumi.Input[int]] = None,
                 icmp_type: Optional[pulumi.Input[int]] = None,
                 iprange: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_number: Optional[pulumi.Input[int]] = None,
                 proxy: Optional[pulumi.Input[str]] = None,
                 sctp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tcp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 udp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a FortimanagerFirewallObjectService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortimanagerFirewallObjectServiceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FortimanagerFirewallObjectService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FortimanagerFirewallObjectServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortimanagerFirewallObjectServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 icmp_code: Optional[pulumi.Input[int]] = None,
                 icmp_type: Optional[pulumi.Input[int]] = None,
                 iprange: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_number: Optional[pulumi.Input[int]] = None,
                 proxy: Optional[pulumi.Input[str]] = None,
                 sctp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tcp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 udp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortimanagerFirewallObjectServiceArgs.__new__(FortimanagerFirewallObjectServiceArgs)

            __props__.__dict__["adom"] = adom
            __props__.__dict__["category"] = category
            __props__.__dict__["comment"] = comment
            __props__.__dict__["fqdn"] = fqdn
            __props__.__dict__["icmp_code"] = icmp_code
            __props__.__dict__["icmp_type"] = icmp_type
            __props__.__dict__["iprange"] = iprange
            __props__.__dict__["name"] = name
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["protocol_number"] = protocol_number
            __props__.__dict__["proxy"] = proxy
            __props__.__dict__["sctp_portranges"] = sctp_portranges
            __props__.__dict__["tcp_portranges"] = tcp_portranges
            __props__.__dict__["udp_portranges"] = udp_portranges
        super(FortimanagerFirewallObjectService, __self__).__init__(
            'fortios:index/fortimanagerFirewallObjectService:FortimanagerFirewallObjectService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            category: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            icmp_code: Optional[pulumi.Input[int]] = None,
            icmp_type: Optional[pulumi.Input[int]] = None,
            iprange: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            protocol_number: Optional[pulumi.Input[int]] = None,
            proxy: Optional[pulumi.Input[str]] = None,
            sctp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tcp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            udp_portranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'FortimanagerFirewallObjectService':
        """
        Get an existing FortimanagerFirewallObjectService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortimanagerFirewallObjectServiceState.__new__(_FortimanagerFirewallObjectServiceState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["category"] = category
        __props__.__dict__["comment"] = comment
        __props__.__dict__["fqdn"] = fqdn
        __props__.__dict__["icmp_code"] = icmp_code
        __props__.__dict__["icmp_type"] = icmp_type
        __props__.__dict__["iprange"] = iprange
        __props__.__dict__["name"] = name
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["protocol_number"] = protocol_number
        __props__.__dict__["proxy"] = proxy
        __props__.__dict__["sctp_portranges"] = sctp_portranges
        __props__.__dict__["tcp_portranges"] = tcp_portranges
        __props__.__dict__["udp_portranges"] = udp_portranges
        return FortimanagerFirewallObjectService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "icmp_code")

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "icmp_type")

    @property
    @pulumi.getter
    def iprange(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "iprange")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolNumber")
    def protocol_number(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "protocol_number")

    @property
    @pulumi.getter
    def proxy(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "proxy")

    @property
    @pulumi.getter(name="sctpPortranges")
    def sctp_portranges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "sctp_portranges")

    @property
    @pulumi.getter(name="tcpPortranges")
    def tcp_portranges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tcp_portranges")

    @property
    @pulumi.getter(name="udpPortranges")
    def udp_portranges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "udp_portranges")

