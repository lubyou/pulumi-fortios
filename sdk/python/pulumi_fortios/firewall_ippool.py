# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallIppoolArgs', 'FirewallIppool']

@pulumi.input_type
class FirewallIppoolArgs:
    def __init__(__self__, *,
                 endip: pulumi.Input[str],
                 startip: pulumi.Input[str],
                 add_nat64_route: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_interface: Optional[pulumi.Input[str]] = None,
                 block_size: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 endport: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat64: Optional[pulumi.Input[str]] = None,
                 num_blocks_per_user: Optional[pulumi.Input[int]] = None,
                 pba_timeout: Optional[pulumi.Input[int]] = None,
                 permit_any_host: Optional[pulumi.Input[str]] = None,
                 port_per_user: Optional[pulumi.Input[int]] = None,
                 source_endip: Optional[pulumi.Input[str]] = None,
                 source_startip: Optional[pulumi.Input[str]] = None,
                 startport: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallIppool resource.
        """
        pulumi.set(__self__, "endip", endip)
        pulumi.set(__self__, "startip", startip)
        if add_nat64_route is not None:
            pulumi.set(__self__, "add_nat64_route", add_nat64_route)
        if arp_intf is not None:
            pulumi.set(__self__, "arp_intf", arp_intf)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if associated_interface is not None:
            pulumi.set(__self__, "associated_interface", associated_interface)
        if block_size is not None:
            pulumi.set(__self__, "block_size", block_size)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if endport is not None:
            pulumi.set(__self__, "endport", endport)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nat64 is not None:
            pulumi.set(__self__, "nat64", nat64)
        if num_blocks_per_user is not None:
            pulumi.set(__self__, "num_blocks_per_user", num_blocks_per_user)
        if pba_timeout is not None:
            pulumi.set(__self__, "pba_timeout", pba_timeout)
        if permit_any_host is not None:
            pulumi.set(__self__, "permit_any_host", permit_any_host)
        if port_per_user is not None:
            pulumi.set(__self__, "port_per_user", port_per_user)
        if source_endip is not None:
            pulumi.set(__self__, "source_endip", source_endip)
        if source_startip is not None:
            pulumi.set(__self__, "source_startip", source_startip)
        if startport is not None:
            pulumi.set(__self__, "startport", startport)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def endip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endip")

    @endip.setter
    def endip(self, value: pulumi.Input[str]):
        pulumi.set(self, "endip", value)

    @property
    @pulumi.getter
    def startip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "startip")

    @startip.setter
    def startip(self, value: pulumi.Input[str]):
        pulumi.set(self, "startip", value)

    @property
    @pulumi.getter(name="addNat64Route")
    def add_nat64_route(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "add_nat64_route")

    @add_nat64_route.setter
    def add_nat64_route(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "add_nat64_route", value)

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_intf")

    @arp_intf.setter
    def arp_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_intf", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

    @property
    @pulumi.getter(name="associatedInterface")
    def associated_interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "associated_interface")

    @associated_interface.setter
    def associated_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associated_interface", value)

    @property
    @pulumi.getter(name="blockSize")
    def block_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "block_size")

    @block_size.setter
    def block_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "block_size", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def endport(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "endport")

    @endport.setter
    def endport(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "endport", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nat64(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nat64")

    @nat64.setter
    def nat64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat64", value)

    @property
    @pulumi.getter(name="numBlocksPerUser")
    def num_blocks_per_user(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "num_blocks_per_user")

    @num_blocks_per_user.setter
    def num_blocks_per_user(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "num_blocks_per_user", value)

    @property
    @pulumi.getter(name="pbaTimeout")
    def pba_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "pba_timeout")

    @pba_timeout.setter
    def pba_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pba_timeout", value)

    @property
    @pulumi.getter(name="permitAnyHost")
    def permit_any_host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "permit_any_host")

    @permit_any_host.setter
    def permit_any_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permit_any_host", value)

    @property
    @pulumi.getter(name="portPerUser")
    def port_per_user(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port_per_user")

    @port_per_user.setter
    def port_per_user(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port_per_user", value)

    @property
    @pulumi.getter(name="sourceEndip")
    def source_endip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_endip")

    @source_endip.setter
    def source_endip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_endip", value)

    @property
    @pulumi.getter(name="sourceStartip")
    def source_startip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_startip")

    @source_startip.setter
    def source_startip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_startip", value)

    @property
    @pulumi.getter
    def startport(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "startport")

    @startport.setter
    def startport(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "startport", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallIppoolState:
    def __init__(__self__, *,
                 add_nat64_route: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_interface: Optional[pulumi.Input[str]] = None,
                 block_size: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 endport: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat64: Optional[pulumi.Input[str]] = None,
                 num_blocks_per_user: Optional[pulumi.Input[int]] = None,
                 pba_timeout: Optional[pulumi.Input[int]] = None,
                 permit_any_host: Optional[pulumi.Input[str]] = None,
                 port_per_user: Optional[pulumi.Input[int]] = None,
                 source_endip: Optional[pulumi.Input[str]] = None,
                 source_startip: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 startport: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallIppool resources.
        """
        if add_nat64_route is not None:
            pulumi.set(__self__, "add_nat64_route", add_nat64_route)
        if arp_intf is not None:
            pulumi.set(__self__, "arp_intf", arp_intf)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if associated_interface is not None:
            pulumi.set(__self__, "associated_interface", associated_interface)
        if block_size is not None:
            pulumi.set(__self__, "block_size", block_size)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if endip is not None:
            pulumi.set(__self__, "endip", endip)
        if endport is not None:
            pulumi.set(__self__, "endport", endport)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nat64 is not None:
            pulumi.set(__self__, "nat64", nat64)
        if num_blocks_per_user is not None:
            pulumi.set(__self__, "num_blocks_per_user", num_blocks_per_user)
        if pba_timeout is not None:
            pulumi.set(__self__, "pba_timeout", pba_timeout)
        if permit_any_host is not None:
            pulumi.set(__self__, "permit_any_host", permit_any_host)
        if port_per_user is not None:
            pulumi.set(__self__, "port_per_user", port_per_user)
        if source_endip is not None:
            pulumi.set(__self__, "source_endip", source_endip)
        if source_startip is not None:
            pulumi.set(__self__, "source_startip", source_startip)
        if startip is not None:
            pulumi.set(__self__, "startip", startip)
        if startport is not None:
            pulumi.set(__self__, "startport", startport)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addNat64Route")
    def add_nat64_route(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "add_nat64_route")

    @add_nat64_route.setter
    def add_nat64_route(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "add_nat64_route", value)

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_intf")

    @arp_intf.setter
    def arp_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_intf", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

    @property
    @pulumi.getter(name="associatedInterface")
    def associated_interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "associated_interface")

    @associated_interface.setter
    def associated_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associated_interface", value)

    @property
    @pulumi.getter(name="blockSize")
    def block_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "block_size")

    @block_size.setter
    def block_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "block_size", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def endip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "endip")

    @endip.setter
    def endip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endip", value)

    @property
    @pulumi.getter
    def endport(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "endport")

    @endport.setter
    def endport(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "endport", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nat64(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nat64")

    @nat64.setter
    def nat64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat64", value)

    @property
    @pulumi.getter(name="numBlocksPerUser")
    def num_blocks_per_user(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "num_blocks_per_user")

    @num_blocks_per_user.setter
    def num_blocks_per_user(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "num_blocks_per_user", value)

    @property
    @pulumi.getter(name="pbaTimeout")
    def pba_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "pba_timeout")

    @pba_timeout.setter
    def pba_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pba_timeout", value)

    @property
    @pulumi.getter(name="permitAnyHost")
    def permit_any_host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "permit_any_host")

    @permit_any_host.setter
    def permit_any_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permit_any_host", value)

    @property
    @pulumi.getter(name="portPerUser")
    def port_per_user(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port_per_user")

    @port_per_user.setter
    def port_per_user(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port_per_user", value)

    @property
    @pulumi.getter(name="sourceEndip")
    def source_endip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_endip")

    @source_endip.setter
    def source_endip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_endip", value)

    @property
    @pulumi.getter(name="sourceStartip")
    def source_startip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_startip")

    @source_startip.setter
    def source_startip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_startip", value)

    @property
    @pulumi.getter
    def startip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "startip")

    @startip.setter
    def startip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "startip", value)

    @property
    @pulumi.getter
    def startport(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "startport")

    @startport.setter
    def startport(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "startport", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallIppool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_nat64_route: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_interface: Optional[pulumi.Input[str]] = None,
                 block_size: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 endport: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat64: Optional[pulumi.Input[str]] = None,
                 num_blocks_per_user: Optional[pulumi.Input[int]] = None,
                 pba_timeout: Optional[pulumi.Input[int]] = None,
                 permit_any_host: Optional[pulumi.Input[str]] = None,
                 port_per_user: Optional[pulumi.Input[int]] = None,
                 source_endip: Optional[pulumi.Input[str]] = None,
                 source_startip: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 startport: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallIppool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallIppoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallIppool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallIppoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallIppoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_nat64_route: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_interface: Optional[pulumi.Input[str]] = None,
                 block_size: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 endport: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat64: Optional[pulumi.Input[str]] = None,
                 num_blocks_per_user: Optional[pulumi.Input[int]] = None,
                 pba_timeout: Optional[pulumi.Input[int]] = None,
                 permit_any_host: Optional[pulumi.Input[str]] = None,
                 port_per_user: Optional[pulumi.Input[int]] = None,
                 source_endip: Optional[pulumi.Input[str]] = None,
                 source_startip: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 startport: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallIppoolArgs.__new__(FirewallIppoolArgs)

            __props__.__dict__["add_nat64_route"] = add_nat64_route
            __props__.__dict__["arp_intf"] = arp_intf
            __props__.__dict__["arp_reply"] = arp_reply
            __props__.__dict__["associated_interface"] = associated_interface
            __props__.__dict__["block_size"] = block_size
            __props__.__dict__["comments"] = comments
            if endip is None and not opts.urn:
                raise TypeError("Missing required property 'endip'")
            __props__.__dict__["endip"] = endip
            __props__.__dict__["endport"] = endport
            __props__.__dict__["name"] = name
            __props__.__dict__["nat64"] = nat64
            __props__.__dict__["num_blocks_per_user"] = num_blocks_per_user
            __props__.__dict__["pba_timeout"] = pba_timeout
            __props__.__dict__["permit_any_host"] = permit_any_host
            __props__.__dict__["port_per_user"] = port_per_user
            __props__.__dict__["source_endip"] = source_endip
            __props__.__dict__["source_startip"] = source_startip
            if startip is None and not opts.urn:
                raise TypeError("Missing required property 'startip'")
            __props__.__dict__["startip"] = startip
            __props__.__dict__["startport"] = startport
            __props__.__dict__["type"] = type
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallIppool, __self__).__init__(
            'fortios:index/firewallIppool:FirewallIppool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_nat64_route: Optional[pulumi.Input[str]] = None,
            arp_intf: Optional[pulumi.Input[str]] = None,
            arp_reply: Optional[pulumi.Input[str]] = None,
            associated_interface: Optional[pulumi.Input[str]] = None,
            block_size: Optional[pulumi.Input[int]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            endip: Optional[pulumi.Input[str]] = None,
            endport: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nat64: Optional[pulumi.Input[str]] = None,
            num_blocks_per_user: Optional[pulumi.Input[int]] = None,
            pba_timeout: Optional[pulumi.Input[int]] = None,
            permit_any_host: Optional[pulumi.Input[str]] = None,
            port_per_user: Optional[pulumi.Input[int]] = None,
            source_endip: Optional[pulumi.Input[str]] = None,
            source_startip: Optional[pulumi.Input[str]] = None,
            startip: Optional[pulumi.Input[str]] = None,
            startport: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallIppool':
        """
        Get an existing FirewallIppool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallIppoolState.__new__(_FirewallIppoolState)

        __props__.__dict__["add_nat64_route"] = add_nat64_route
        __props__.__dict__["arp_intf"] = arp_intf
        __props__.__dict__["arp_reply"] = arp_reply
        __props__.__dict__["associated_interface"] = associated_interface
        __props__.__dict__["block_size"] = block_size
        __props__.__dict__["comments"] = comments
        __props__.__dict__["endip"] = endip
        __props__.__dict__["endport"] = endport
        __props__.__dict__["name"] = name
        __props__.__dict__["nat64"] = nat64
        __props__.__dict__["num_blocks_per_user"] = num_blocks_per_user
        __props__.__dict__["pba_timeout"] = pba_timeout
        __props__.__dict__["permit_any_host"] = permit_any_host
        __props__.__dict__["port_per_user"] = port_per_user
        __props__.__dict__["source_endip"] = source_endip
        __props__.__dict__["source_startip"] = source_startip
        __props__.__dict__["startip"] = startip
        __props__.__dict__["startport"] = startport
        __props__.__dict__["type"] = type
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallIppool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addNat64Route")
    def add_nat64_route(self) -> pulumi.Output[str]:
        return pulumi.get(self, "add_nat64_route")

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arp_intf")

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arp_reply")

    @property
    @pulumi.getter(name="associatedInterface")
    def associated_interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "associated_interface")

    @property
    @pulumi.getter(name="blockSize")
    def block_size(self) -> pulumi.Output[int]:
        return pulumi.get(self, "block_size")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def endip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endip")

    @property
    @pulumi.getter
    def endport(self) -> pulumi.Output[int]:
        return pulumi.get(self, "endport")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nat64(self) -> pulumi.Output[str]:
        return pulumi.get(self, "nat64")

    @property
    @pulumi.getter(name="numBlocksPerUser")
    def num_blocks_per_user(self) -> pulumi.Output[int]:
        return pulumi.get(self, "num_blocks_per_user")

    @property
    @pulumi.getter(name="pbaTimeout")
    def pba_timeout(self) -> pulumi.Output[int]:
        return pulumi.get(self, "pba_timeout")

    @property
    @pulumi.getter(name="permitAnyHost")
    def permit_any_host(self) -> pulumi.Output[str]:
        return pulumi.get(self, "permit_any_host")

    @property
    @pulumi.getter(name="portPerUser")
    def port_per_user(self) -> pulumi.Output[int]:
        return pulumi.get(self, "port_per_user")

    @property
    @pulumi.getter(name="sourceEndip")
    def source_endip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_endip")

    @property
    @pulumi.getter(name="sourceStartip")
    def source_startip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_startip")

    @property
    @pulumi.getter
    def startip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "startip")

    @property
    @pulumi.getter
    def startport(self) -> pulumi.Output[int]:
        return pulumi.get(self, "startport")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

