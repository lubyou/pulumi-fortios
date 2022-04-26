# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
        :param pulumi.Input[str] status: Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] eip: End IP.
        :param pulumi.Input[str] ip_mode: IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        :param pulumi.Input[str] local_ip: Local IP to be used for peer's remote IP.
        :param pulumi.Input[str] sip: Start IP.
        :param pulumi.Input[str] usrgrp: User group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
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
        """
        Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def eip(self) -> Optional[pulumi.Input[str]]:
        """
        End IP.
        """
        return pulumi.get(self, "eip")

    @eip.setter
    def eip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip", value)

    @property
    @pulumi.getter(name="ipMode")
    def ip_mode(self) -> Optional[pulumi.Input[str]]:
        """
        IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        """
        return pulumi.get(self, "ip_mode")

    @ip_mode.setter
    def ip_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_mode", value)

    @property
    @pulumi.getter(name="localIp")
    def local_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Local IP to be used for peer's remote IP.
        """
        return pulumi.get(self, "local_ip")

    @local_ip.setter
    def local_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_ip", value)

    @property
    @pulumi.getter
    def sip(self) -> Optional[pulumi.Input[str]]:
        """
        Start IP.
        """
        return pulumi.get(self, "sip")

    @sip.setter
    def sip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sip", value)

    @property
    @pulumi.getter
    def usrgrp(self) -> Optional[pulumi.Input[str]]:
        """
        User group.
        """
        return pulumi.get(self, "usrgrp")

    @usrgrp.setter
    def usrgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usrgrp", value)

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
        :param pulumi.Input[str] eip: End IP.
        :param pulumi.Input[str] ip_mode: IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        :param pulumi.Input[str] local_ip: Local IP to be used for peer's remote IP.
        :param pulumi.Input[str] sip: Start IP.
        :param pulumi.Input[str] status: Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usrgrp: User group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
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
        """
        End IP.
        """
        return pulumi.get(self, "eip")

    @eip.setter
    def eip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip", value)

    @property
    @pulumi.getter(name="ipMode")
    def ip_mode(self) -> Optional[pulumi.Input[str]]:
        """
        IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        """
        return pulumi.get(self, "ip_mode")

    @ip_mode.setter
    def ip_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_mode", value)

    @property
    @pulumi.getter(name="localIp")
    def local_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Local IP to be used for peer's remote IP.
        """
        return pulumi.get(self, "local_ip")

    @local_ip.setter
    def local_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_ip", value)

    @property
    @pulumi.getter
    def sip(self) -> Optional[pulumi.Input[str]]:
        """
        Start IP.
        """
        return pulumi.get(self, "sip")

    @sip.setter
    def sip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def usrgrp(self) -> Optional[pulumi.Input[str]]:
        """
        User group.
        """
        return pulumi.get(self, "usrgrp")

    @usrgrp.setter
    def usrgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usrgrp", value)

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
        Configure PPTP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.VpnPptp("trname",
            eip="1.1.1.22",
            ip_mode="range",
            local_ip="0.0.0.0",
            sip="1.1.1.1",
            status="enable",
            usrgrp="Guest-group")
        ```

        ## Import

        Vpn Pptp can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/vpnPptp:VpnPptp labelname VpnPptp
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/vpnPptp:VpnPptp labelname VpnPptp
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eip: End IP.
        :param pulumi.Input[str] ip_mode: IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        :param pulumi.Input[str] local_ip: Local IP to be used for peer's remote IP.
        :param pulumi.Input[str] sip: Start IP.
        :param pulumi.Input[str] status: Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usrgrp: User group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpnPptpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure PPTP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.VpnPptp("trname",
            eip="1.1.1.22",
            ip_mode="range",
            local_ip="0.0.0.0",
            sip="1.1.1.1",
            status="enable",
            usrgrp="Guest-group")
        ```

        ## Import

        Vpn Pptp can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/vpnPptp:VpnPptp labelname VpnPptp
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/vpnPptp:VpnPptp labelname VpnPptp
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

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
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
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
        :param pulumi.Input[str] eip: End IP.
        :param pulumi.Input[str] ip_mode: IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        :param pulumi.Input[str] local_ip: Local IP to be used for peer's remote IP.
        :param pulumi.Input[str] sip: Start IP.
        :param pulumi.Input[str] status: Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usrgrp: User group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
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
        """
        End IP.
        """
        return pulumi.get(self, "eip")

    @property
    @pulumi.getter(name="ipMode")
    def ip_mode(self) -> pulumi.Output[str]:
        """
        IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        """
        return pulumi.get(self, "ip_mode")

    @property
    @pulumi.getter(name="localIp")
    def local_ip(self) -> pulumi.Output[str]:
        """
        Local IP to be used for peer's remote IP.
        """
        return pulumi.get(self, "local_ip")

    @property
    @pulumi.getter
    def sip(self) -> pulumi.Output[str]:
        """
        Start IP.
        """
        return pulumi.get(self, "sip")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def usrgrp(self) -> pulumi.Output[str]:
        """
        User group.
        """
        return pulumi.get(self, "usrgrp")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

