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

__all__ = ['SystemApiUserArgs', 'SystemApiUser']

@pulumi.input_type
class SystemApiUserArgs:
    def __init__(__self__, *,
                 accprofile: pulumi.Input[str],
                 api_key: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 cors_allow_origin: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_auth: Optional[pulumi.Input[str]] = None,
                 peer_group: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserTrusthostArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserVdomArgs']]]] = None):
        """
        The set of arguments for constructing a SystemApiUser resource.
        :param pulumi.Input[str] accprofile: Admin user access profile.
        :param pulumi.Input[str] api_key: Admin user password.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] cors_allow_origin: Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Virtual domain name.
        :param pulumi.Input[str] peer_auth: Enable/disable peer authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] peer_group: Peer group name.
        :param pulumi.Input[str] schedule: Schedule name.
        :param pulumi.Input[Sequence[pulumi.Input['SystemApiUserTrusthostArgs']]] trusthosts: Trusthost. The structure of `trusthost` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['SystemApiUserVdomArgs']]] vdoms: Virtual domains. The structure of `vdom` block is documented below.
        """
        pulumi.set(__self__, "accprofile", accprofile)
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if cors_allow_origin is not None:
            pulumi.set(__self__, "cors_allow_origin", cors_allow_origin)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peer_auth is not None:
            pulumi.set(__self__, "peer_auth", peer_auth)
        if peer_group is not None:
            pulumi.set(__self__, "peer_group", peer_group)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if trusthosts is not None:
            pulumi.set(__self__, "trusthosts", trusthosts)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms is not None:
            pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter
    def accprofile(self) -> pulumi.Input[str]:
        """
        Admin user access profile.
        """
        return pulumi.get(self, "accprofile")

    @accprofile.setter
    def accprofile(self, value: pulumi.Input[str]):
        pulumi.set(self, "accprofile", value)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[str]]:
        """
        Admin user password.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="corsAllowOrigin")
    def cors_allow_origin(self) -> Optional[pulumi.Input[str]]:
        """
        Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        """
        return pulumi.get(self, "cors_allow_origin")

    @cors_allow_origin.setter
    def cors_allow_origin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cors_allow_origin", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual domain name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="peerAuth")
    def peer_auth(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable peer authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "peer_auth")

    @peer_auth.setter
    def peer_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_auth", value)

    @property
    @pulumi.getter(name="peerGroup")
    def peer_group(self) -> Optional[pulumi.Input[str]]:
        """
        Peer group name.
        """
        return pulumi.get(self, "peer_group")

    @peer_group.setter
    def peer_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_group", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule name.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def trusthosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserTrusthostArgs']]]]:
        """
        Trusthost. The structure of `trusthost` block is documented below.
        """
        return pulumi.get(self, "trusthosts")

    @trusthosts.setter
    def trusthosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserTrusthostArgs']]]]):
        pulumi.set(self, "trusthosts", value)

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
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserVdomArgs']]]]:
        """
        Virtual domains. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserVdomArgs']]]]):
        pulumi.set(self, "vdoms", value)


@pulumi.input_type
class _SystemApiUserState:
    def __init__(__self__, *,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 cors_allow_origin: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_auth: Optional[pulumi.Input[str]] = None,
                 peer_group: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserTrusthostArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserVdomArgs']]]] = None):
        """
        Input properties used for looking up and filtering SystemApiUser resources.
        :param pulumi.Input[str] accprofile: Admin user access profile.
        :param pulumi.Input[str] api_key: Admin user password.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] cors_allow_origin: Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Virtual domain name.
        :param pulumi.Input[str] peer_auth: Enable/disable peer authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] peer_group: Peer group name.
        :param pulumi.Input[str] schedule: Schedule name.
        :param pulumi.Input[Sequence[pulumi.Input['SystemApiUserTrusthostArgs']]] trusthosts: Trusthost. The structure of `trusthost` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['SystemApiUserVdomArgs']]] vdoms: Virtual domains. The structure of `vdom` block is documented below.
        """
        if accprofile is not None:
            pulumi.set(__self__, "accprofile", accprofile)
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if cors_allow_origin is not None:
            pulumi.set(__self__, "cors_allow_origin", cors_allow_origin)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peer_auth is not None:
            pulumi.set(__self__, "peer_auth", peer_auth)
        if peer_group is not None:
            pulumi.set(__self__, "peer_group", peer_group)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if trusthosts is not None:
            pulumi.set(__self__, "trusthosts", trusthosts)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms is not None:
            pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter
    def accprofile(self) -> Optional[pulumi.Input[str]]:
        """
        Admin user access profile.
        """
        return pulumi.get(self, "accprofile")

    @accprofile.setter
    def accprofile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accprofile", value)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[str]]:
        """
        Admin user password.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="corsAllowOrigin")
    def cors_allow_origin(self) -> Optional[pulumi.Input[str]]:
        """
        Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        """
        return pulumi.get(self, "cors_allow_origin")

    @cors_allow_origin.setter
    def cors_allow_origin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cors_allow_origin", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual domain name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="peerAuth")
    def peer_auth(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable peer authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "peer_auth")

    @peer_auth.setter
    def peer_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_auth", value)

    @property
    @pulumi.getter(name="peerGroup")
    def peer_group(self) -> Optional[pulumi.Input[str]]:
        """
        Peer group name.
        """
        return pulumi.get(self, "peer_group")

    @peer_group.setter
    def peer_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_group", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule name.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def trusthosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserTrusthostArgs']]]]:
        """
        Trusthost. The structure of `trusthost` block is documented below.
        """
        return pulumi.get(self, "trusthosts")

    @trusthosts.setter
    def trusthosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserTrusthostArgs']]]]):
        pulumi.set(self, "trusthosts", value)

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
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserVdomArgs']]]]:
        """
        Virtual domains. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemApiUserVdomArgs']]]]):
        pulumi.set(self, "vdoms", value)


class SystemApiUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 cors_allow_origin: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_auth: Optional[pulumi.Input[str]] = None,
                 peer_group: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserTrusthostArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserVdomArgs']]]]] = None,
                 __props__=None):
        """
        ## Import

        System ApiUser can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemApiUser:SystemApiUser labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accprofile: Admin user access profile.
        :param pulumi.Input[str] api_key: Admin user password.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] cors_allow_origin: Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Virtual domain name.
        :param pulumi.Input[str] peer_auth: Enable/disable peer authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] peer_group: Peer group name.
        :param pulumi.Input[str] schedule: Schedule name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserTrusthostArgs']]]] trusthosts: Trusthost. The structure of `trusthost` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserVdomArgs']]]] vdoms: Virtual domains. The structure of `vdom` block is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemApiUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        System ApiUser can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemApiUser:SystemApiUser labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemApiUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemApiUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 cors_allow_origin: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_auth: Optional[pulumi.Input[str]] = None,
                 peer_group: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserTrusthostArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserVdomArgs']]]]] = None,
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
            __props__ = SystemApiUserArgs.__new__(SystemApiUserArgs)

            if accprofile is None and not opts.urn:
                raise TypeError("Missing required property 'accprofile'")
            __props__.__dict__["accprofile"] = accprofile
            __props__.__dict__["api_key"] = api_key
            __props__.__dict__["comments"] = comments
            __props__.__dict__["cors_allow_origin"] = cors_allow_origin
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["peer_auth"] = peer_auth
            __props__.__dict__["peer_group"] = peer_group
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["trusthosts"] = trusthosts
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vdoms"] = vdoms
        super(SystemApiUser, __self__).__init__(
            'fortios:index/systemApiUser:SystemApiUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accprofile: Optional[pulumi.Input[str]] = None,
            api_key: Optional[pulumi.Input[str]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            cors_allow_origin: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            peer_auth: Optional[pulumi.Input[str]] = None,
            peer_group: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserTrusthostArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserVdomArgs']]]]] = None) -> 'SystemApiUser':
        """
        Get an existing SystemApiUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accprofile: Admin user access profile.
        :param pulumi.Input[str] api_key: Admin user password.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] cors_allow_origin: Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Virtual domain name.
        :param pulumi.Input[str] peer_auth: Enable/disable peer authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] peer_group: Peer group name.
        :param pulumi.Input[str] schedule: Schedule name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserTrusthostArgs']]]] trusthosts: Trusthost. The structure of `trusthost` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemApiUserVdomArgs']]]] vdoms: Virtual domains. The structure of `vdom` block is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemApiUserState.__new__(_SystemApiUserState)

        __props__.__dict__["accprofile"] = accprofile
        __props__.__dict__["api_key"] = api_key
        __props__.__dict__["comments"] = comments
        __props__.__dict__["cors_allow_origin"] = cors_allow_origin
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["peer_auth"] = peer_auth
        __props__.__dict__["peer_group"] = peer_group
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["trusthosts"] = trusthosts
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vdoms"] = vdoms
        return SystemApiUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accprofile(self) -> pulumi.Output[str]:
        """
        Admin user access profile.
        """
        return pulumi.get(self, "accprofile")

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[Optional[str]]:
        """
        Admin user password.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="corsAllowOrigin")
    def cors_allow_origin(self) -> pulumi.Output[str]:
        """
        Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        """
        return pulumi.get(self, "cors_allow_origin")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Virtual domain name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerAuth")
    def peer_auth(self) -> pulumi.Output[str]:
        """
        Enable/disable peer authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "peer_auth")

    @property
    @pulumi.getter(name="peerGroup")
    def peer_group(self) -> pulumi.Output[str]:
        """
        Peer group name.
        """
        return pulumi.get(self, "peer_group")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        Schedule name.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def trusthosts(self) -> pulumi.Output[Optional[Sequence['outputs.SystemApiUserTrusthost']]]:
        """
        Trusthost. The structure of `trusthost` block is documented below.
        """
        return pulumi.get(self, "trusthosts")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vdoms(self) -> pulumi.Output[Optional[Sequence['outputs.SystemApiUserVdom']]]:
        """
        Virtual domains. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

