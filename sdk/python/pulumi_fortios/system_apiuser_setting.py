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

__all__ = ['SystemAPIUserSettingArgs', 'SystemAPIUserSetting']

@pulumi.input_type
class SystemAPIUserSettingArgs:
    def __init__(__self__, *,
                 accprofile: pulumi.Input[str],
                 trusthosts: pulumi.Input[Sequence[pulumi.Input['SystemAPIUserSettingTrusthostArgs']]],
                 vdoms: pulumi.Input[Sequence[pulumi.Input[str]]],
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemAPIUserSetting resource.
        :param pulumi.Input[str] accprofile: Admin user access profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vdoms: Virtual domains.
               * `trusthost-Type` - (Required) Trusthost type.
               * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: User name.
        """
        pulumi.set(__self__, "accprofile", accprofile)
        pulumi.set(__self__, "trusthosts", trusthosts)
        pulumi.set(__self__, "vdoms", vdoms)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
    @pulumi.getter
    def trusthosts(self) -> pulumi.Input[Sequence[pulumi.Input['SystemAPIUserSettingTrusthostArgs']]]:
        return pulumi.get(self, "trusthosts")

    @trusthosts.setter
    def trusthosts(self, value: pulumi.Input[Sequence[pulumi.Input['SystemAPIUserSettingTrusthostArgs']]]):
        pulumi.set(self, "trusthosts", value)

    @property
    @pulumi.getter
    def vdoms(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Virtual domains.
        * `trusthost-Type` - (Required) Trusthost type.
        * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "vdoms", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _SystemAPIUserSettingState:
    def __init__(__self__, *,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAPIUserSettingTrusthostArgs']]]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SystemAPIUserSetting resources.
        :param pulumi.Input[str] accprofile: Admin user access profile.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: User name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vdoms: Virtual domains.
               * `trusthost-Type` - (Required) Trusthost type.
               * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
        """
        if accprofile is not None:
            pulumi.set(__self__, "accprofile", accprofile)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if trusthosts is not None:
            pulumi.set(__self__, "trusthosts", trusthosts)
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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def trusthosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemAPIUserSettingTrusthostArgs']]]]:
        return pulumi.get(self, "trusthosts")

    @trusthosts.setter
    def trusthosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAPIUserSettingTrusthostArgs']]]]):
        pulumi.set(self, "trusthosts", value)

    @property
    @pulumi.getter
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Virtual domains.
        * `trusthost-Type` - (Required) Trusthost type.
        * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vdoms", value)


class SystemAPIUserSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAPIUserSettingTrusthostArgs']]]]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.

        !> **Warning:** The resource will be deprecated and replaced by new resource `SystemApiUser`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test2 = fortios.SystemAPIUserSetting("test2",
            accprofile="restAPIprofile",
            trusthosts=[
                fortios.SystemAPIUserSettingTrusthostArgs(
                    ipv4_trusthost="61.149.0.0 255.255.0.0",
                    type="ipv4-trusthost",
                ),
                fortios.SystemAPIUserSettingTrusthostArgs(
                    ipv4_trusthost="22.22.0.0 255.255.0.0",
                    type="ipv4-trusthost",
                ),
            ],
            vdoms=["root"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accprofile: Admin user access profile.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: User name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vdoms: Virtual domains.
               * `trusthost-Type` - (Required) Trusthost type.
               * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemAPIUserSettingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.

        !> **Warning:** The resource will be deprecated and replaced by new resource `SystemApiUser`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test2 = fortios.SystemAPIUserSetting("test2",
            accprofile="restAPIprofile",
            trusthosts=[
                fortios.SystemAPIUserSettingTrusthostArgs(
                    ipv4_trusthost="61.149.0.0 255.255.0.0",
                    type="ipv4-trusthost",
                ),
                fortios.SystemAPIUserSettingTrusthostArgs(
                    ipv4_trusthost="22.22.0.0 255.255.0.0",
                    type="ipv4-trusthost",
                ),
            ],
            vdoms=["root"])
        ```

        :param str resource_name: The name of the resource.
        :param SystemAPIUserSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemAPIUserSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAPIUserSettingTrusthostArgs']]]]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = SystemAPIUserSettingArgs.__new__(SystemAPIUserSettingArgs)

            if accprofile is None and not opts.urn:
                raise TypeError("Missing required property 'accprofile'")
            __props__.__dict__["accprofile"] = accprofile
            __props__.__dict__["comments"] = comments
            __props__.__dict__["name"] = name
            if trusthosts is None and not opts.urn:
                raise TypeError("Missing required property 'trusthosts'")
            __props__.__dict__["trusthosts"] = trusthosts
            if vdoms is None and not opts.urn:
                raise TypeError("Missing required property 'vdoms'")
            __props__.__dict__["vdoms"] = vdoms
        super(SystemAPIUserSetting, __self__).__init__(
            'fortios:index/systemAPIUserSetting:SystemAPIUserSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accprofile: Optional[pulumi.Input[str]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            trusthosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAPIUserSettingTrusthostArgs']]]]] = None,
            vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SystemAPIUserSetting':
        """
        Get an existing SystemAPIUserSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accprofile: Admin user access profile.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: User name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vdoms: Virtual domains.
               * `trusthost-Type` - (Required) Trusthost type.
               * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemAPIUserSettingState.__new__(_SystemAPIUserSettingState)

        __props__.__dict__["accprofile"] = accprofile
        __props__.__dict__["comments"] = comments
        __props__.__dict__["name"] = name
        __props__.__dict__["trusthosts"] = trusthosts
        __props__.__dict__["vdoms"] = vdoms
        return SystemAPIUserSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accprofile(self) -> pulumi.Output[str]:
        """
        Admin user access profile.
        """
        return pulumi.get(self, "accprofile")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def trusthosts(self) -> pulumi.Output[Sequence['outputs.SystemAPIUserSettingTrusthost']]:
        return pulumi.get(self, "trusthosts")

    @property
    @pulumi.getter
    def vdoms(self) -> pulumi.Output[Sequence[str]]:
        """
        Virtual domains.
        * `trusthost-Type` - (Required) Trusthost type.
        * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
        """
        return pulumi.get(self, "vdoms")

