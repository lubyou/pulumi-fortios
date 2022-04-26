# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WanoptPeerArgs', 'WanoptPeer']

@pulumi.input_type
class WanoptPeerArgs:
    def __init__(__self__, *,
                 ip: Optional[pulumi.Input[str]] = None,
                 peer_host_id: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WanoptPeer resource.
        :param pulumi.Input[str] ip: Peer IP address.
        :param pulumi.Input[str] peer_host_id: Peer host ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if peer_host_id is not None:
            pulumi.set(__self__, "peer_host_id", peer_host_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        Peer IP address.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="peerHostId")
    def peer_host_id(self) -> Optional[pulumi.Input[str]]:
        """
        Peer host ID.
        """
        return pulumi.get(self, "peer_host_id")

    @peer_host_id.setter
    def peer_host_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_host_id", value)

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
class _WanoptPeerState:
    def __init__(__self__, *,
                 ip: Optional[pulumi.Input[str]] = None,
                 peer_host_id: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WanoptPeer resources.
        :param pulumi.Input[str] ip: Peer IP address.
        :param pulumi.Input[str] peer_host_id: Peer host ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if peer_host_id is not None:
            pulumi.set(__self__, "peer_host_id", peer_host_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        Peer IP address.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="peerHostId")
    def peer_host_id(self) -> Optional[pulumi.Input[str]]:
        """
        Peer host ID.
        """
        return pulumi.get(self, "peer_host_id")

    @peer_host_id.setter
    def peer_host_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_host_id", value)

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


class WanoptPeer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 peer_host_id: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure WAN optimization peers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WanoptPeer("trname",
            ip="1.1.1.1",
            peer_host_id="1")
        ```

        ## Import

        Wanopt Peer can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wanoptPeer:WanoptPeer labelname {{peer_host_id}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wanoptPeer:WanoptPeer labelname {{peer_host_id}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip: Peer IP address.
        :param pulumi.Input[str] peer_host_id: Peer host ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WanoptPeerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure WAN optimization peers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WanoptPeer("trname",
            ip="1.1.1.1",
            peer_host_id="1")
        ```

        ## Import

        Wanopt Peer can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wanoptPeer:WanoptPeer labelname {{peer_host_id}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wanoptPeer:WanoptPeer labelname {{peer_host_id}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WanoptPeerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WanoptPeerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 peer_host_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = WanoptPeerArgs.__new__(WanoptPeerArgs)

            __props__.__dict__["ip"] = ip
            __props__.__dict__["peer_host_id"] = peer_host_id
            __props__.__dict__["vdomparam"] = vdomparam
        super(WanoptPeer, __self__).__init__(
            'fortios:index/wanoptPeer:WanoptPeer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip: Optional[pulumi.Input[str]] = None,
            peer_host_id: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WanoptPeer':
        """
        Get an existing WanoptPeer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip: Peer IP address.
        :param pulumi.Input[str] peer_host_id: Peer host ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WanoptPeerState.__new__(_WanoptPeerState)

        __props__.__dict__["ip"] = ip
        __props__.__dict__["peer_host_id"] = peer_host_id
        __props__.__dict__["vdomparam"] = vdomparam
        return WanoptPeer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        Peer IP address.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="peerHostId")
    def peer_host_id(self) -> pulumi.Output[str]:
        """
        Peer host ID.
        """
        return pulumi.get(self, "peer_host_id")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

