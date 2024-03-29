# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpnPptpArgs', 'VpnPptp']

@pulumi.input_type
class VpnPptpArgs:
    def __init__(__self__, *,
                 status: pulumi.Input[str],
                 eip: Optional[pulumi.Input[str]] = None,
                 ip_mode: Optional[pulumi.Input[str]] = None,
                 local_ip: Optional[pulumi.Input[str]] = None,
                 sip: Optional[pulumi.Input[str]] = None,
                 usrgrp: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpnPptp resource.
        """
        pulumi.set(__self__, "status", status)
        if eip is not None:
            pulumi.set(__self__, "eip", eip)
        if ip_mode is not None:
            pulumi.set(__self__, "ip_mode", ip_mode)
        if local_ip is not None:
            pulumi.set(__self__, "local_ip", local_ip)
        if sip is not None:
            pulumi.set(__self__, "sip", sip)
        if usrgrp is not None:
            pulumi.set(__self__, "usrgrp", usrgrp)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def eip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eip")

    @eip.setter
    def eip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip", value)

    @property
    @pulumi.getter(name="ipMode")
    def ip_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_mode")

    @ip_mode.setter
    def ip_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_mode", value)

    @property
    @pulumi.getter(name="localIp")
    def local_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "local_ip")

    @local_ip.setter
    def local_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_ip", value)

    @property
    @pulumi.getter
    def sip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sip")

    @sip.setter
    def sip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sip", value)

    @property
    @pulumi.getter
    def usrgrp(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "usrgrp")

    @usrgrp.setter
    def usrgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usrgrp", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _VpnPptpState:
    def __init__(__self__, *,
                 eip: Optional[pulumi.Input[str]] = None,
                 ip_mode: Optional[pulumi.Input[str]] = None,
                 local_ip: Optional[pulumi.Input[str]] = None,
                 sip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usrgrp: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpnPptp resources.
        """
        if eip is not None:
            pulumi.set(__self__, "eip", eip)
        if ip_mode is not None:
            pulumi.set(__self__, "ip_mode", ip_mode)
        if local_ip is not None:
            pulumi.set(__self__, "local_ip", local_ip)
        if sip is not None:
            pulumi.set(__self__, "sip", sip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if usrgrp is not None:
            pulumi.set(__self__, "usrgrp", usrgrp)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def eip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eip")

    @eip.setter
    def eip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip", value)

    @property
    @pulumi.getter(name="ipMode")
    def ip_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_mode")

    @ip_mode.setter
    def ip_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_mode", value)

    @property
    @pulumi.getter(name="localIp")
    def local_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "local_ip")

    @local_ip.setter
    def local_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_ip", value)

    @property
    @pulumi.getter
    def sip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sip")

    @sip.setter
    def sip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def usrgrp(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "usrgrp")

    @usrgrp.setter
    def usrgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usrgrp", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class VpnPptp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 ip_mode: Optional[pulumi.Input[str]] = None,
                 local_ip: Optional[pulumi.Input[str]] = None,
                 sip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usrgrp: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VpnPptp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpnPptpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpnPptp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpnPptpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpnPptpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 ip_mode: Optional[pulumi.Input[str]] = None,
                 local_ip: Optional[pulumi.Input[str]] = None,
                 sip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usrgrp: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpnPptpArgs.__new__(VpnPptpArgs)

            __props__.__dict__["eip"] = eip
            __props__.__dict__["ip_mode"] = ip_mode
            __props__.__dict__["local_ip"] = local_ip
            __props__.__dict__["sip"] = sip
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["usrgrp"] = usrgrp
            __props__.__dict__["vdomparam"] = vdomparam
        super(VpnPptp, __self__).__init__(
            'fortios:index/vpnPptp:VpnPptp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            eip: Optional[pulumi.Input[str]] = None,
            ip_mode: Optional[pulumi.Input[str]] = None,
            local_ip: Optional[pulumi.Input[str]] = None,
            sip: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            usrgrp: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'VpnPptp':
        """
        Get an existing VpnPptp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpnPptpState.__new__(_VpnPptpState)

        __props__.__dict__["eip"] = eip
        __props__.__dict__["ip_mode"] = ip_mode
        __props__.__dict__["local_ip"] = local_ip
        __props__.__dict__["sip"] = sip
        __props__.__dict__["status"] = status
        __props__.__dict__["usrgrp"] = usrgrp
        __props__.__dict__["vdomparam"] = vdomparam
        return VpnPptp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def eip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "eip")

    @property
    @pulumi.getter(name="ipMode")
    def ip_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ip_mode")

    @property
    @pulumi.getter(name="localIp")
    def local_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "local_ip")

    @property
    @pulumi.getter
    def sip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sip")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def usrgrp(self) -> pulumi.Output[str]:
        return pulumi.get(self, "usrgrp")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

