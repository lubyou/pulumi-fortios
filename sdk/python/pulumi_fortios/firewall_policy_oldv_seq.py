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

__all__ = ['FirewallPolicyOldvSeqArgs', 'FirewallPolicyOldvSeq']

@pulumi.input_type
class FirewallPolicyOldvSeqArgs:
    def __init__(__self__, *,
                 alter_position: pulumi.Input[str],
                 policy_dst_id: pulumi.Input[int],
                 policy_src_id: pulumi.Input[int],
                 comment: Optional[pulumi.Input[str]] = None,
                 enable_state_checking: Optional[pulumi.Input[bool]] = None,
                 state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallPolicyOldvSeq resource.
        """
        pulumi.set(__self__, "alter_position", alter_position)
        pulumi.set(__self__, "policy_dst_id", policy_dst_id)
        pulumi.set(__self__, "policy_src_id", policy_src_id)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if enable_state_checking is not None:
            pulumi.set(__self__, "enable_state_checking", enable_state_checking)
        if state_policy_srcdst_pos is not None:
            pulumi.set(__self__, "state_policy_srcdst_pos", state_policy_srcdst_pos)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="alterPosition")
    def alter_position(self) -> pulumi.Input[str]:
        return pulumi.get(self, "alter_position")

    @alter_position.setter
    def alter_position(self, value: pulumi.Input[str]):
        pulumi.set(self, "alter_position", value)

    @property
    @pulumi.getter(name="policyDstId")
    def policy_dst_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "policy_dst_id")

    @policy_dst_id.setter
    def policy_dst_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "policy_dst_id", value)

    @property
    @pulumi.getter(name="policySrcId")
    def policy_src_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "policy_src_id")

    @policy_src_id.setter
    def policy_src_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "policy_src_id", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="enableStateChecking")
    def enable_state_checking(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_state_checking")

    @enable_state_checking.setter
    def enable_state_checking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_state_checking", value)

    @property
    @pulumi.getter(name="statePolicySrcdstPos")
    def state_policy_srcdst_pos(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state_policy_srcdst_pos")

    @state_policy_srcdst_pos.setter
    def state_policy_srcdst_pos(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_policy_srcdst_pos", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallPolicyOldvSeqState:
    def __init__(__self__, *,
                 alter_position: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 enable_state_checking: Optional[pulumi.Input[bool]] = None,
                 policy_dst_id: Optional[pulumi.Input[int]] = None,
                 policy_src_id: Optional[pulumi.Input[int]] = None,
                 state_policy_lists: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallPolicyOldvSeqStatePolicyListArgs']]]] = None,
                 state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallPolicyOldvSeq resources.
        """
        if alter_position is not None:
            pulumi.set(__self__, "alter_position", alter_position)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if enable_state_checking is not None:
            pulumi.set(__self__, "enable_state_checking", enable_state_checking)
        if policy_dst_id is not None:
            pulumi.set(__self__, "policy_dst_id", policy_dst_id)
        if policy_src_id is not None:
            pulumi.set(__self__, "policy_src_id", policy_src_id)
        if state_policy_lists is not None:
            pulumi.set(__self__, "state_policy_lists", state_policy_lists)
        if state_policy_srcdst_pos is not None:
            pulumi.set(__self__, "state_policy_srcdst_pos", state_policy_srcdst_pos)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="alterPosition")
    def alter_position(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "alter_position")

    @alter_position.setter
    def alter_position(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alter_position", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="enableStateChecking")
    def enable_state_checking(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_state_checking")

    @enable_state_checking.setter
    def enable_state_checking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_state_checking", value)

    @property
    @pulumi.getter(name="policyDstId")
    def policy_dst_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "policy_dst_id")

    @policy_dst_id.setter
    def policy_dst_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policy_dst_id", value)

    @property
    @pulumi.getter(name="policySrcId")
    def policy_src_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "policy_src_id")

    @policy_src_id.setter
    def policy_src_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policy_src_id", value)

    @property
    @pulumi.getter(name="statePolicyLists")
    def state_policy_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallPolicyOldvSeqStatePolicyListArgs']]]]:
        return pulumi.get(self, "state_policy_lists")

    @state_policy_lists.setter
    def state_policy_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallPolicyOldvSeqStatePolicyListArgs']]]]):
        pulumi.set(self, "state_policy_lists", value)

    @property
    @pulumi.getter(name="statePolicySrcdstPos")
    def state_policy_srcdst_pos(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state_policy_srcdst_pos")

    @state_policy_srcdst_pos.setter
    def state_policy_srcdst_pos(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_policy_srcdst_pos", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallPolicyOldvSeq(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alter_position: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 enable_state_checking: Optional[pulumi.Input[bool]] = None,
                 policy_dst_id: Optional[pulumi.Input[int]] = None,
                 policy_src_id: Optional[pulumi.Input[int]] = None,
                 state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallPolicyOldvSeq resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallPolicyOldvSeqArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallPolicyOldvSeq resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallPolicyOldvSeqArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallPolicyOldvSeqArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alter_position: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 enable_state_checking: Optional[pulumi.Input[bool]] = None,
                 policy_dst_id: Optional[pulumi.Input[int]] = None,
                 policy_src_id: Optional[pulumi.Input[int]] = None,
                 state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallPolicyOldvSeqArgs.__new__(FirewallPolicyOldvSeqArgs)

            if alter_position is None and not opts.urn:
                raise TypeError("Missing required property 'alter_position'")
            __props__.__dict__["alter_position"] = alter_position
            __props__.__dict__["comment"] = comment
            __props__.__dict__["enable_state_checking"] = enable_state_checking
            if policy_dst_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_dst_id'")
            __props__.__dict__["policy_dst_id"] = policy_dst_id
            if policy_src_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_src_id'")
            __props__.__dict__["policy_src_id"] = policy_src_id
            __props__.__dict__["state_policy_srcdst_pos"] = state_policy_srcdst_pos
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["state_policy_lists"] = None
        super(FirewallPolicyOldvSeq, __self__).__init__(
            'fortios:index/firewallPolicyOldvSeq:FirewallPolicyOldvSeq',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alter_position: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            enable_state_checking: Optional[pulumi.Input[bool]] = None,
            policy_dst_id: Optional[pulumi.Input[int]] = None,
            policy_src_id: Optional[pulumi.Input[int]] = None,
            state_policy_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallPolicyOldvSeqStatePolicyListArgs']]]]] = None,
            state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallPolicyOldvSeq':
        """
        Get an existing FirewallPolicyOldvSeq resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallPolicyOldvSeqState.__new__(_FirewallPolicyOldvSeqState)

        __props__.__dict__["alter_position"] = alter_position
        __props__.__dict__["comment"] = comment
        __props__.__dict__["enable_state_checking"] = enable_state_checking
        __props__.__dict__["policy_dst_id"] = policy_dst_id
        __props__.__dict__["policy_src_id"] = policy_src_id
        __props__.__dict__["state_policy_lists"] = state_policy_lists
        __props__.__dict__["state_policy_srcdst_pos"] = state_policy_srcdst_pos
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallPolicyOldvSeq(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alterPosition")
    def alter_position(self) -> pulumi.Output[str]:
        return pulumi.get(self, "alter_position")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="enableStateChecking")
    def enable_state_checking(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_state_checking")

    @property
    @pulumi.getter(name="policyDstId")
    def policy_dst_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "policy_dst_id")

    @property
    @pulumi.getter(name="policySrcId")
    def policy_src_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "policy_src_id")

    @property
    @pulumi.getter(name="statePolicyLists")
    def state_policy_lists(self) -> pulumi.Output[Sequence['outputs.FirewallPolicyOldvSeqStatePolicyList']]:
        return pulumi.get(self, "state_policy_lists")

    @property
    @pulumi.getter(name="statePolicySrcdstPos")
    def state_policy_srcdst_pos(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "state_policy_srcdst_pos")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")
