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

__all__ = ['WanoptCacheServiceArgs', 'WanoptCacheService']

@pulumi.input_type
class WanoptCacheServiceArgs:
    def __init__(__self__, *,
                 acceptable_connections: Optional[pulumi.Input[str]] = None,
                 collaboration: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 dst_peers: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceDstPeerArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 prefer_scenario: Optional[pulumi.Input[str]] = None,
                 src_peers: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceSrcPeerArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WanoptCacheService resource.
        :param pulumi.Input[str] acceptable_connections: Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        :param pulumi.Input[str] collaboration: Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] device_id: Device ID of this peer.
        :param pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceDstPeerArgs']]] dst_peers: Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] prefer_scenario: Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        :param pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceSrcPeerArgs']]] src_peers: Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if acceptable_connections is not None:
            pulumi.set(__self__, "acceptable_connections", acceptable_connections)
        if collaboration is not None:
            pulumi.set(__self__, "collaboration", collaboration)
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if dst_peers is not None:
            pulumi.set(__self__, "dst_peers", dst_peers)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if prefer_scenario is not None:
            pulumi.set(__self__, "prefer_scenario", prefer_scenario)
        if src_peers is not None:
            pulumi.set(__self__, "src_peers", src_peers)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="acceptableConnections")
    def acceptable_connections(self) -> Optional[pulumi.Input[str]]:
        """
        Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        """
        return pulumi.get(self, "acceptable_connections")

    @acceptable_connections.setter
    def acceptable_connections(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "acceptable_connections", value)

    @property
    @pulumi.getter
    def collaboration(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "collaboration")

    @collaboration.setter
    def collaboration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collaboration", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        Device ID of this peer.
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="dstPeers")
    def dst_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceDstPeerArgs']]]]:
        """
        Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        """
        return pulumi.get(self, "dst_peers")

    @dst_peers.setter
    def dst_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceDstPeerArgs']]]]):
        pulumi.set(self, "dst_peers", value)

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
    @pulumi.getter(name="preferScenario")
    def prefer_scenario(self) -> Optional[pulumi.Input[str]]:
        """
        Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        """
        return pulumi.get(self, "prefer_scenario")

    @prefer_scenario.setter
    def prefer_scenario(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefer_scenario", value)

    @property
    @pulumi.getter(name="srcPeers")
    def src_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceSrcPeerArgs']]]]:
        """
        Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        """
        return pulumi.get(self, "src_peers")

    @src_peers.setter
    def src_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceSrcPeerArgs']]]]):
        pulumi.set(self, "src_peers", value)

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
class _WanoptCacheServiceState:
    def __init__(__self__, *,
                 acceptable_connections: Optional[pulumi.Input[str]] = None,
                 collaboration: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 dst_peers: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceDstPeerArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 prefer_scenario: Optional[pulumi.Input[str]] = None,
                 src_peers: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceSrcPeerArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WanoptCacheService resources.
        :param pulumi.Input[str] acceptable_connections: Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        :param pulumi.Input[str] collaboration: Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] device_id: Device ID of this peer.
        :param pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceDstPeerArgs']]] dst_peers: Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] prefer_scenario: Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        :param pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceSrcPeerArgs']]] src_peers: Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if acceptable_connections is not None:
            pulumi.set(__self__, "acceptable_connections", acceptable_connections)
        if collaboration is not None:
            pulumi.set(__self__, "collaboration", collaboration)
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if dst_peers is not None:
            pulumi.set(__self__, "dst_peers", dst_peers)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if prefer_scenario is not None:
            pulumi.set(__self__, "prefer_scenario", prefer_scenario)
        if src_peers is not None:
            pulumi.set(__self__, "src_peers", src_peers)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="acceptableConnections")
    def acceptable_connections(self) -> Optional[pulumi.Input[str]]:
        """
        Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        """
        return pulumi.get(self, "acceptable_connections")

    @acceptable_connections.setter
    def acceptable_connections(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "acceptable_connections", value)

    @property
    @pulumi.getter
    def collaboration(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "collaboration")

    @collaboration.setter
    def collaboration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collaboration", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        Device ID of this peer.
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="dstPeers")
    def dst_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceDstPeerArgs']]]]:
        """
        Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        """
        return pulumi.get(self, "dst_peers")

    @dst_peers.setter
    def dst_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceDstPeerArgs']]]]):
        pulumi.set(self, "dst_peers", value)

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
    @pulumi.getter(name="preferScenario")
    def prefer_scenario(self) -> Optional[pulumi.Input[str]]:
        """
        Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        """
        return pulumi.get(self, "prefer_scenario")

    @prefer_scenario.setter
    def prefer_scenario(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefer_scenario", value)

    @property
    @pulumi.getter(name="srcPeers")
    def src_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceSrcPeerArgs']]]]:
        """
        Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        """
        return pulumi.get(self, "src_peers")

    @src_peers.setter
    def src_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptCacheServiceSrcPeerArgs']]]]):
        pulumi.set(self, "src_peers", value)

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


class WanoptCacheService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acceptable_connections: Optional[pulumi.Input[str]] = None,
                 collaboration: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 dst_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceDstPeerArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 prefer_scenario: Optional[pulumi.Input[str]] = None,
                 src_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceSrcPeerArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Designate cache-service for wan-optimization and webcache.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WanoptCacheService("trname",
            acceptable_connections="any",
            collaboration="disable",
            device_id="default_dev_id",
            prefer_scenario="balance")
        ```

        ## Import

        Wanopt CacheService can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wanoptCacheService:WanoptCacheService labelname WanoptCacheService
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wanoptCacheService:WanoptCacheService labelname WanoptCacheService
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acceptable_connections: Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        :param pulumi.Input[str] collaboration: Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] device_id: Device ID of this peer.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceDstPeerArgs']]]] dst_peers: Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] prefer_scenario: Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceSrcPeerArgs']]]] src_peers: Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WanoptCacheServiceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Designate cache-service for wan-optimization and webcache.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WanoptCacheService("trname",
            acceptable_connections="any",
            collaboration="disable",
            device_id="default_dev_id",
            prefer_scenario="balance")
        ```

        ## Import

        Wanopt CacheService can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wanoptCacheService:WanoptCacheService labelname WanoptCacheService
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wanoptCacheService:WanoptCacheService labelname WanoptCacheService
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WanoptCacheServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WanoptCacheServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acceptable_connections: Optional[pulumi.Input[str]] = None,
                 collaboration: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 dst_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceDstPeerArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 prefer_scenario: Optional[pulumi.Input[str]] = None,
                 src_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceSrcPeerArgs']]]]] = None,
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
            __props__ = WanoptCacheServiceArgs.__new__(WanoptCacheServiceArgs)

            __props__.__dict__["acceptable_connections"] = acceptable_connections
            __props__.__dict__["collaboration"] = collaboration
            __props__.__dict__["device_id"] = device_id
            __props__.__dict__["dst_peers"] = dst_peers
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["prefer_scenario"] = prefer_scenario
            __props__.__dict__["src_peers"] = src_peers
            __props__.__dict__["vdomparam"] = vdomparam
        super(WanoptCacheService, __self__).__init__(
            'fortios:index/wanoptCacheService:WanoptCacheService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acceptable_connections: Optional[pulumi.Input[str]] = None,
            collaboration: Optional[pulumi.Input[str]] = None,
            device_id: Optional[pulumi.Input[str]] = None,
            dst_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceDstPeerArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            prefer_scenario: Optional[pulumi.Input[str]] = None,
            src_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceSrcPeerArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WanoptCacheService':
        """
        Get an existing WanoptCacheService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acceptable_connections: Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        :param pulumi.Input[str] collaboration: Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] device_id: Device ID of this peer.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceDstPeerArgs']]]] dst_peers: Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] prefer_scenario: Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptCacheServiceSrcPeerArgs']]]] src_peers: Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WanoptCacheServiceState.__new__(_WanoptCacheServiceState)

        __props__.__dict__["acceptable_connections"] = acceptable_connections
        __props__.__dict__["collaboration"] = collaboration
        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["dst_peers"] = dst_peers
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["prefer_scenario"] = prefer_scenario
        __props__.__dict__["src_peers"] = src_peers
        __props__.__dict__["vdomparam"] = vdomparam
        return WanoptCacheService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptableConnections")
    def acceptable_connections(self) -> pulumi.Output[str]:
        """
        Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        """
        return pulumi.get(self, "acceptable_connections")

    @property
    @pulumi.getter
    def collaboration(self) -> pulumi.Output[str]:
        """
        Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "collaboration")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        Device ID of this peer.
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="dstPeers")
    def dst_peers(self) -> pulumi.Output[Optional[Sequence['outputs.WanoptCacheServiceDstPeer']]]:
        """
        Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        """
        return pulumi.get(self, "dst_peers")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="preferScenario")
    def prefer_scenario(self) -> pulumi.Output[str]:
        """
        Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        """
        return pulumi.get(self, "prefer_scenario")

    @property
    @pulumi.getter(name="srcPeers")
    def src_peers(self) -> pulumi.Output[Optional[Sequence['outputs.WanoptCacheServiceSrcPeer']]]:
        """
        Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        """
        return pulumi.get(self, "src_peers")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

