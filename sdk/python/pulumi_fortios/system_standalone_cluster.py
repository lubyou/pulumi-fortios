# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemStandaloneClusterArgs', 'SystemStandaloneCluster']

@pulumi.input_type
class SystemStandaloneClusterArgs:
    def __init__(__self__, *,
                 encryption: Optional[pulumi.Input[str]] = None,
                 group_member_id: Optional[pulumi.Input[int]] = None,
                 layer2_connection: Optional[pulumi.Input[str]] = None,
                 psksecret: Optional[pulumi.Input[str]] = None,
                 session_sync_dev: Optional[pulumi.Input[str]] = None,
                 standalone_group_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemStandaloneCluster resource.
        :param pulumi.Input[str] encryption: Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] group_member_id: Cluster member ID (0 - 3).
        :param pulumi.Input[str] layer2_connection: Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        :param pulumi.Input[str] psksecret: Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] session_sync_dev: Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        :param pulumi.Input[int] standalone_group_id: Cluster group ID (0 - 255). Must be the same for all members.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if group_member_id is not None:
            pulumi.set(__self__, "group_member_id", group_member_id)
        if layer2_connection is not None:
            pulumi.set(__self__, "layer2_connection", layer2_connection)
        if psksecret is not None:
            pulumi.set(__self__, "psksecret", psksecret)
        if session_sync_dev is not None:
            pulumi.set(__self__, "session_sync_dev", session_sync_dev)
        if standalone_group_id is not None:
            pulumi.set(__self__, "standalone_group_id", standalone_group_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter(name="groupMemberId")
    def group_member_id(self) -> Optional[pulumi.Input[int]]:
        """
        Cluster member ID (0 - 3).
        """
        return pulumi.get(self, "group_member_id")

    @group_member_id.setter
    def group_member_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_member_id", value)

    @property
    @pulumi.getter(name="layer2Connection")
    def layer2_connection(self) -> Optional[pulumi.Input[str]]:
        """
        Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        """
        return pulumi.get(self, "layer2_connection")

    @layer2_connection.setter
    def layer2_connection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "layer2_connection", value)

    @property
    @pulumi.getter
    def psksecret(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psksecret")

    @psksecret.setter
    def psksecret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psksecret", value)

    @property
    @pulumi.getter(name="sessionSyncDev")
    def session_sync_dev(self) -> Optional[pulumi.Input[str]]:
        """
        Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        """
        return pulumi.get(self, "session_sync_dev")

    @session_sync_dev.setter
    def session_sync_dev(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_sync_dev", value)

    @property
    @pulumi.getter(name="standaloneGroupId")
    def standalone_group_id(self) -> Optional[pulumi.Input[int]]:
        """
        Cluster group ID (0 - 255). Must be the same for all members.
        """
        return pulumi.get(self, "standalone_group_id")

    @standalone_group_id.setter
    def standalone_group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "standalone_group_id", value)

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
class _SystemStandaloneClusterState:
    def __init__(__self__, *,
                 encryption: Optional[pulumi.Input[str]] = None,
                 group_member_id: Optional[pulumi.Input[int]] = None,
                 layer2_connection: Optional[pulumi.Input[str]] = None,
                 psksecret: Optional[pulumi.Input[str]] = None,
                 session_sync_dev: Optional[pulumi.Input[str]] = None,
                 standalone_group_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemStandaloneCluster resources.
        :param pulumi.Input[str] encryption: Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] group_member_id: Cluster member ID (0 - 3).
        :param pulumi.Input[str] layer2_connection: Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        :param pulumi.Input[str] psksecret: Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] session_sync_dev: Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        :param pulumi.Input[int] standalone_group_id: Cluster group ID (0 - 255). Must be the same for all members.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if group_member_id is not None:
            pulumi.set(__self__, "group_member_id", group_member_id)
        if layer2_connection is not None:
            pulumi.set(__self__, "layer2_connection", layer2_connection)
        if psksecret is not None:
            pulumi.set(__self__, "psksecret", psksecret)
        if session_sync_dev is not None:
            pulumi.set(__self__, "session_sync_dev", session_sync_dev)
        if standalone_group_id is not None:
            pulumi.set(__self__, "standalone_group_id", standalone_group_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter(name="groupMemberId")
    def group_member_id(self) -> Optional[pulumi.Input[int]]:
        """
        Cluster member ID (0 - 3).
        """
        return pulumi.get(self, "group_member_id")

    @group_member_id.setter
    def group_member_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_member_id", value)

    @property
    @pulumi.getter(name="layer2Connection")
    def layer2_connection(self) -> Optional[pulumi.Input[str]]:
        """
        Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        """
        return pulumi.get(self, "layer2_connection")

    @layer2_connection.setter
    def layer2_connection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "layer2_connection", value)

    @property
    @pulumi.getter
    def psksecret(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psksecret")

    @psksecret.setter
    def psksecret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psksecret", value)

    @property
    @pulumi.getter(name="sessionSyncDev")
    def session_sync_dev(self) -> Optional[pulumi.Input[str]]:
        """
        Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        """
        return pulumi.get(self, "session_sync_dev")

    @session_sync_dev.setter
    def session_sync_dev(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_sync_dev", value)

    @property
    @pulumi.getter(name="standaloneGroupId")
    def standalone_group_id(self) -> Optional[pulumi.Input[int]]:
        """
        Cluster group ID (0 - 255). Must be the same for all members.
        """
        return pulumi.get(self, "standalone_group_id")

    @standalone_group_id.setter
    def standalone_group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "standalone_group_id", value)

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


class SystemStandaloneCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 group_member_id: Optional[pulumi.Input[int]] = None,
                 layer2_connection: Optional[pulumi.Input[str]] = None,
                 psksecret: Optional[pulumi.Input[str]] = None,
                 session_sync_dev: Optional[pulumi.Input[str]] = None,
                 standalone_group_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        System StandaloneCluster can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemStandaloneCluster:SystemStandaloneCluster labelname SystemStandaloneCluster
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemStandaloneCluster:SystemStandaloneCluster labelname SystemStandaloneCluster
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] encryption: Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] group_member_id: Cluster member ID (0 - 3).
        :param pulumi.Input[str] layer2_connection: Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        :param pulumi.Input[str] psksecret: Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] session_sync_dev: Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        :param pulumi.Input[int] standalone_group_id: Cluster group ID (0 - 255). Must be the same for all members.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemStandaloneClusterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        System StandaloneCluster can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemStandaloneCluster:SystemStandaloneCluster labelname SystemStandaloneCluster
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemStandaloneCluster:SystemStandaloneCluster labelname SystemStandaloneCluster
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemStandaloneClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemStandaloneClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 group_member_id: Optional[pulumi.Input[int]] = None,
                 layer2_connection: Optional[pulumi.Input[str]] = None,
                 psksecret: Optional[pulumi.Input[str]] = None,
                 session_sync_dev: Optional[pulumi.Input[str]] = None,
                 standalone_group_id: Optional[pulumi.Input[int]] = None,
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
            __props__ = SystemStandaloneClusterArgs.__new__(SystemStandaloneClusterArgs)

            __props__.__dict__["encryption"] = encryption
            __props__.__dict__["group_member_id"] = group_member_id
            __props__.__dict__["layer2_connection"] = layer2_connection
            __props__.__dict__["psksecret"] = psksecret
            __props__.__dict__["session_sync_dev"] = session_sync_dev
            __props__.__dict__["standalone_group_id"] = standalone_group_id
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemStandaloneCluster, __self__).__init__(
            'fortios:index/systemStandaloneCluster:SystemStandaloneCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            encryption: Optional[pulumi.Input[str]] = None,
            group_member_id: Optional[pulumi.Input[int]] = None,
            layer2_connection: Optional[pulumi.Input[str]] = None,
            psksecret: Optional[pulumi.Input[str]] = None,
            session_sync_dev: Optional[pulumi.Input[str]] = None,
            standalone_group_id: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemStandaloneCluster':
        """
        Get an existing SystemStandaloneCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] encryption: Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] group_member_id: Cluster member ID (0 - 3).
        :param pulumi.Input[str] layer2_connection: Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        :param pulumi.Input[str] psksecret: Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] session_sync_dev: Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        :param pulumi.Input[int] standalone_group_id: Cluster group ID (0 - 255). Must be the same for all members.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemStandaloneClusterState.__new__(_SystemStandaloneClusterState)

        __props__.__dict__["encryption"] = encryption
        __props__.__dict__["group_member_id"] = group_member_id
        __props__.__dict__["layer2_connection"] = layer2_connection
        __props__.__dict__["psksecret"] = psksecret
        __props__.__dict__["session_sync_dev"] = session_sync_dev
        __props__.__dict__["standalone_group_id"] = standalone_group_id
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemStandaloneCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[str]:
        """
        Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="groupMemberId")
    def group_member_id(self) -> pulumi.Output[int]:
        """
        Cluster member ID (0 - 3).
        """
        return pulumi.get(self, "group_member_id")

    @property
    @pulumi.getter(name="layer2Connection")
    def layer2_connection(self) -> pulumi.Output[str]:
        """
        Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        """
        return pulumi.get(self, "layer2_connection")

    @property
    @pulumi.getter
    def psksecret(self) -> pulumi.Output[Optional[str]]:
        """
        Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psksecret")

    @property
    @pulumi.getter(name="sessionSyncDev")
    def session_sync_dev(self) -> pulumi.Output[str]:
        """
        Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        """
        return pulumi.get(self, "session_sync_dev")

    @property
    @pulumi.getter(name="standaloneGroupId")
    def standalone_group_id(self) -> pulumi.Output[int]:
        """
        Cluster group ID (0 - 255). Must be the same for all members.
        """
        return pulumi.get(self, "standalone_group_id")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

