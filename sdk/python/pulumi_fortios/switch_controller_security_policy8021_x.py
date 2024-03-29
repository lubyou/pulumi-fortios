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

__all__ = ['SwitchControllerSecurityPolicy8021XArgs', 'SwitchControllerSecurityPolicy8021X']

@pulumi.input_type
class SwitchControllerSecurityPolicy8021XArgs:
    def __init__(__self__, *,
                 auth_fail_vlan: Optional[pulumi.Input[str]] = None,
                 auth_fail_vlan_id: Optional[pulumi.Input[str]] = None,
                 auth_fail_vlanid: Optional[pulumi.Input[int]] = None,
                 authserver_timeout_period: Optional[pulumi.Input[int]] = None,
                 authserver_timeout_vlan: Optional[pulumi.Input[str]] = None,
                 authserver_timeout_vlanid: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 eap_auto_untagged_vlans: Optional[pulumi.Input[str]] = None,
                 eap_passthru: Optional[pulumi.Input[str]] = None,
                 framevid_apply: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 guest_auth_delay: Optional[pulumi.Input[int]] = None,
                 guest_vlan: Optional[pulumi.Input[str]] = None,
                 guest_vlan_id: Optional[pulumi.Input[str]] = None,
                 guest_vlanid: Optional[pulumi.Input[int]] = None,
                 mac_auth_bypass: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_auth: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 radius_timeout_overwrite: Optional[pulumi.Input[str]] = None,
                 security_mode: Optional[pulumi.Input[str]] = None,
                 user_groups: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerSecurityPolicy8021X resource.
        """
        if auth_fail_vlan is not None:
            pulumi.set(__self__, "auth_fail_vlan", auth_fail_vlan)
        if auth_fail_vlan_id is not None:
            pulumi.set(__self__, "auth_fail_vlan_id", auth_fail_vlan_id)
        if auth_fail_vlanid is not None:
            pulumi.set(__self__, "auth_fail_vlanid", auth_fail_vlanid)
        if authserver_timeout_period is not None:
            pulumi.set(__self__, "authserver_timeout_period", authserver_timeout_period)
        if authserver_timeout_vlan is not None:
            pulumi.set(__self__, "authserver_timeout_vlan", authserver_timeout_vlan)
        if authserver_timeout_vlanid is not None:
            pulumi.set(__self__, "authserver_timeout_vlanid", authserver_timeout_vlanid)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if eap_auto_untagged_vlans is not None:
            pulumi.set(__self__, "eap_auto_untagged_vlans", eap_auto_untagged_vlans)
        if eap_passthru is not None:
            pulumi.set(__self__, "eap_passthru", eap_passthru)
        if framevid_apply is not None:
            pulumi.set(__self__, "framevid_apply", framevid_apply)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if guest_auth_delay is not None:
            pulumi.set(__self__, "guest_auth_delay", guest_auth_delay)
        if guest_vlan is not None:
            pulumi.set(__self__, "guest_vlan", guest_vlan)
        if guest_vlan_id is not None:
            pulumi.set(__self__, "guest_vlan_id", guest_vlan_id)
        if guest_vlanid is not None:
            pulumi.set(__self__, "guest_vlanid", guest_vlanid)
        if mac_auth_bypass is not None:
            pulumi.set(__self__, "mac_auth_bypass", mac_auth_bypass)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if open_auth is not None:
            pulumi.set(__self__, "open_auth", open_auth)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if radius_timeout_overwrite is not None:
            pulumi.set(__self__, "radius_timeout_overwrite", radius_timeout_overwrite)
        if security_mode is not None:
            pulumi.set(__self__, "security_mode", security_mode)
        if user_groups is not None:
            pulumi.set(__self__, "user_groups", user_groups)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authFailVlan")
    def auth_fail_vlan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_fail_vlan")

    @auth_fail_vlan.setter
    def auth_fail_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_fail_vlan", value)

    @property
    @pulumi.getter(name="authFailVlanId")
    def auth_fail_vlan_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_fail_vlan_id")

    @auth_fail_vlan_id.setter
    def auth_fail_vlan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_fail_vlan_id", value)

    @property
    @pulumi.getter(name="authFailVlanid")
    def auth_fail_vlanid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auth_fail_vlanid")

    @auth_fail_vlanid.setter
    def auth_fail_vlanid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auth_fail_vlanid", value)

    @property
    @pulumi.getter(name="authserverTimeoutPeriod")
    def authserver_timeout_period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "authserver_timeout_period")

    @authserver_timeout_period.setter
    def authserver_timeout_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "authserver_timeout_period", value)

    @property
    @pulumi.getter(name="authserverTimeoutVlan")
    def authserver_timeout_vlan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authserver_timeout_vlan")

    @authserver_timeout_vlan.setter
    def authserver_timeout_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authserver_timeout_vlan", value)

    @property
    @pulumi.getter(name="authserverTimeoutVlanid")
    def authserver_timeout_vlanid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authserver_timeout_vlanid")

    @authserver_timeout_vlanid.setter
    def authserver_timeout_vlanid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authserver_timeout_vlanid", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="eapAutoUntaggedVlans")
    def eap_auto_untagged_vlans(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eap_auto_untagged_vlans")

    @eap_auto_untagged_vlans.setter
    def eap_auto_untagged_vlans(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eap_auto_untagged_vlans", value)

    @property
    @pulumi.getter(name="eapPassthru")
    def eap_passthru(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eap_passthru")

    @eap_passthru.setter
    def eap_passthru(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eap_passthru", value)

    @property
    @pulumi.getter(name="framevidApply")
    def framevid_apply(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "framevid_apply")

    @framevid_apply.setter
    def framevid_apply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "framevid_apply", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="guestAuthDelay")
    def guest_auth_delay(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "guest_auth_delay")

    @guest_auth_delay.setter
    def guest_auth_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "guest_auth_delay", value)

    @property
    @pulumi.getter(name="guestVlan")
    def guest_vlan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "guest_vlan")

    @guest_vlan.setter
    def guest_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "guest_vlan", value)

    @property
    @pulumi.getter(name="guestVlanId")
    def guest_vlan_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "guest_vlan_id")

    @guest_vlan_id.setter
    def guest_vlan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "guest_vlan_id", value)

    @property
    @pulumi.getter(name="guestVlanid")
    def guest_vlanid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "guest_vlanid")

    @guest_vlanid.setter
    def guest_vlanid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "guest_vlanid", value)

    @property
    @pulumi.getter(name="macAuthBypass")
    def mac_auth_bypass(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mac_auth_bypass")

    @mac_auth_bypass.setter
    def mac_auth_bypass(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_auth_bypass", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openAuth")
    def open_auth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "open_auth")

    @open_auth.setter
    def open_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "open_auth", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter(name="radiusTimeoutOverwrite")
    def radius_timeout_overwrite(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "radius_timeout_overwrite")

    @radius_timeout_overwrite.setter
    def radius_timeout_overwrite(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "radius_timeout_overwrite", value)

    @property
    @pulumi.getter(name="securityMode")
    def security_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "security_mode")

    @security_mode.setter
    def security_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_mode", value)

    @property
    @pulumi.getter(name="userGroups")
    def user_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]]:
        return pulumi.get(self, "user_groups")

    @user_groups.setter
    def user_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]]):
        pulumi.set(self, "user_groups", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SwitchControllerSecurityPolicy8021XState:
    def __init__(__self__, *,
                 auth_fail_vlan: Optional[pulumi.Input[str]] = None,
                 auth_fail_vlan_id: Optional[pulumi.Input[str]] = None,
                 auth_fail_vlanid: Optional[pulumi.Input[int]] = None,
                 authserver_timeout_period: Optional[pulumi.Input[int]] = None,
                 authserver_timeout_vlan: Optional[pulumi.Input[str]] = None,
                 authserver_timeout_vlanid: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 eap_auto_untagged_vlans: Optional[pulumi.Input[str]] = None,
                 eap_passthru: Optional[pulumi.Input[str]] = None,
                 framevid_apply: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 guest_auth_delay: Optional[pulumi.Input[int]] = None,
                 guest_vlan: Optional[pulumi.Input[str]] = None,
                 guest_vlan_id: Optional[pulumi.Input[str]] = None,
                 guest_vlanid: Optional[pulumi.Input[int]] = None,
                 mac_auth_bypass: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_auth: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 radius_timeout_overwrite: Optional[pulumi.Input[str]] = None,
                 security_mode: Optional[pulumi.Input[str]] = None,
                 user_groups: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerSecurityPolicy8021X resources.
        """
        if auth_fail_vlan is not None:
            pulumi.set(__self__, "auth_fail_vlan", auth_fail_vlan)
        if auth_fail_vlan_id is not None:
            pulumi.set(__self__, "auth_fail_vlan_id", auth_fail_vlan_id)
        if auth_fail_vlanid is not None:
            pulumi.set(__self__, "auth_fail_vlanid", auth_fail_vlanid)
        if authserver_timeout_period is not None:
            pulumi.set(__self__, "authserver_timeout_period", authserver_timeout_period)
        if authserver_timeout_vlan is not None:
            pulumi.set(__self__, "authserver_timeout_vlan", authserver_timeout_vlan)
        if authserver_timeout_vlanid is not None:
            pulumi.set(__self__, "authserver_timeout_vlanid", authserver_timeout_vlanid)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if eap_auto_untagged_vlans is not None:
            pulumi.set(__self__, "eap_auto_untagged_vlans", eap_auto_untagged_vlans)
        if eap_passthru is not None:
            pulumi.set(__self__, "eap_passthru", eap_passthru)
        if framevid_apply is not None:
            pulumi.set(__self__, "framevid_apply", framevid_apply)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if guest_auth_delay is not None:
            pulumi.set(__self__, "guest_auth_delay", guest_auth_delay)
        if guest_vlan is not None:
            pulumi.set(__self__, "guest_vlan", guest_vlan)
        if guest_vlan_id is not None:
            pulumi.set(__self__, "guest_vlan_id", guest_vlan_id)
        if guest_vlanid is not None:
            pulumi.set(__self__, "guest_vlanid", guest_vlanid)
        if mac_auth_bypass is not None:
            pulumi.set(__self__, "mac_auth_bypass", mac_auth_bypass)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if open_auth is not None:
            pulumi.set(__self__, "open_auth", open_auth)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if radius_timeout_overwrite is not None:
            pulumi.set(__self__, "radius_timeout_overwrite", radius_timeout_overwrite)
        if security_mode is not None:
            pulumi.set(__self__, "security_mode", security_mode)
        if user_groups is not None:
            pulumi.set(__self__, "user_groups", user_groups)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authFailVlan")
    def auth_fail_vlan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_fail_vlan")

    @auth_fail_vlan.setter
    def auth_fail_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_fail_vlan", value)

    @property
    @pulumi.getter(name="authFailVlanId")
    def auth_fail_vlan_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_fail_vlan_id")

    @auth_fail_vlan_id.setter
    def auth_fail_vlan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_fail_vlan_id", value)

    @property
    @pulumi.getter(name="authFailVlanid")
    def auth_fail_vlanid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auth_fail_vlanid")

    @auth_fail_vlanid.setter
    def auth_fail_vlanid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auth_fail_vlanid", value)

    @property
    @pulumi.getter(name="authserverTimeoutPeriod")
    def authserver_timeout_period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "authserver_timeout_period")

    @authserver_timeout_period.setter
    def authserver_timeout_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "authserver_timeout_period", value)

    @property
    @pulumi.getter(name="authserverTimeoutVlan")
    def authserver_timeout_vlan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authserver_timeout_vlan")

    @authserver_timeout_vlan.setter
    def authserver_timeout_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authserver_timeout_vlan", value)

    @property
    @pulumi.getter(name="authserverTimeoutVlanid")
    def authserver_timeout_vlanid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authserver_timeout_vlanid")

    @authserver_timeout_vlanid.setter
    def authserver_timeout_vlanid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authserver_timeout_vlanid", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="eapAutoUntaggedVlans")
    def eap_auto_untagged_vlans(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eap_auto_untagged_vlans")

    @eap_auto_untagged_vlans.setter
    def eap_auto_untagged_vlans(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eap_auto_untagged_vlans", value)

    @property
    @pulumi.getter(name="eapPassthru")
    def eap_passthru(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eap_passthru")

    @eap_passthru.setter
    def eap_passthru(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eap_passthru", value)

    @property
    @pulumi.getter(name="framevidApply")
    def framevid_apply(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "framevid_apply")

    @framevid_apply.setter
    def framevid_apply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "framevid_apply", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="guestAuthDelay")
    def guest_auth_delay(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "guest_auth_delay")

    @guest_auth_delay.setter
    def guest_auth_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "guest_auth_delay", value)

    @property
    @pulumi.getter(name="guestVlan")
    def guest_vlan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "guest_vlan")

    @guest_vlan.setter
    def guest_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "guest_vlan", value)

    @property
    @pulumi.getter(name="guestVlanId")
    def guest_vlan_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "guest_vlan_id")

    @guest_vlan_id.setter
    def guest_vlan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "guest_vlan_id", value)

    @property
    @pulumi.getter(name="guestVlanid")
    def guest_vlanid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "guest_vlanid")

    @guest_vlanid.setter
    def guest_vlanid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "guest_vlanid", value)

    @property
    @pulumi.getter(name="macAuthBypass")
    def mac_auth_bypass(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mac_auth_bypass")

    @mac_auth_bypass.setter
    def mac_auth_bypass(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_auth_bypass", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openAuth")
    def open_auth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "open_auth")

    @open_auth.setter
    def open_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "open_auth", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter(name="radiusTimeoutOverwrite")
    def radius_timeout_overwrite(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "radius_timeout_overwrite")

    @radius_timeout_overwrite.setter
    def radius_timeout_overwrite(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "radius_timeout_overwrite", value)

    @property
    @pulumi.getter(name="securityMode")
    def security_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "security_mode")

    @security_mode.setter
    def security_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_mode", value)

    @property
    @pulumi.getter(name="userGroups")
    def user_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]]:
        return pulumi.get(self, "user_groups")

    @user_groups.setter
    def user_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]]):
        pulumi.set(self, "user_groups", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SwitchControllerSecurityPolicy8021X(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_fail_vlan: Optional[pulumi.Input[str]] = None,
                 auth_fail_vlan_id: Optional[pulumi.Input[str]] = None,
                 auth_fail_vlanid: Optional[pulumi.Input[int]] = None,
                 authserver_timeout_period: Optional[pulumi.Input[int]] = None,
                 authserver_timeout_vlan: Optional[pulumi.Input[str]] = None,
                 authserver_timeout_vlanid: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 eap_auto_untagged_vlans: Optional[pulumi.Input[str]] = None,
                 eap_passthru: Optional[pulumi.Input[str]] = None,
                 framevid_apply: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 guest_auth_delay: Optional[pulumi.Input[int]] = None,
                 guest_vlan: Optional[pulumi.Input[str]] = None,
                 guest_vlan_id: Optional[pulumi.Input[str]] = None,
                 guest_vlanid: Optional[pulumi.Input[int]] = None,
                 mac_auth_bypass: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_auth: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 radius_timeout_overwrite: Optional[pulumi.Input[str]] = None,
                 security_mode: Optional[pulumi.Input[str]] = None,
                 user_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SwitchControllerSecurityPolicy8021X resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerSecurityPolicy8021XArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SwitchControllerSecurityPolicy8021X resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SwitchControllerSecurityPolicy8021XArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerSecurityPolicy8021XArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_fail_vlan: Optional[pulumi.Input[str]] = None,
                 auth_fail_vlan_id: Optional[pulumi.Input[str]] = None,
                 auth_fail_vlanid: Optional[pulumi.Input[int]] = None,
                 authserver_timeout_period: Optional[pulumi.Input[int]] = None,
                 authserver_timeout_vlan: Optional[pulumi.Input[str]] = None,
                 authserver_timeout_vlanid: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 eap_auto_untagged_vlans: Optional[pulumi.Input[str]] = None,
                 eap_passthru: Optional[pulumi.Input[str]] = None,
                 framevid_apply: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 guest_auth_delay: Optional[pulumi.Input[int]] = None,
                 guest_vlan: Optional[pulumi.Input[str]] = None,
                 guest_vlan_id: Optional[pulumi.Input[str]] = None,
                 guest_vlanid: Optional[pulumi.Input[int]] = None,
                 mac_auth_bypass: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_auth: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 radius_timeout_overwrite: Optional[pulumi.Input[str]] = None,
                 security_mode: Optional[pulumi.Input[str]] = None,
                 user_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchControllerSecurityPolicy8021XArgs.__new__(SwitchControllerSecurityPolicy8021XArgs)

            __props__.__dict__["auth_fail_vlan"] = auth_fail_vlan
            __props__.__dict__["auth_fail_vlan_id"] = auth_fail_vlan_id
            __props__.__dict__["auth_fail_vlanid"] = auth_fail_vlanid
            __props__.__dict__["authserver_timeout_period"] = authserver_timeout_period
            __props__.__dict__["authserver_timeout_vlan"] = authserver_timeout_vlan
            __props__.__dict__["authserver_timeout_vlanid"] = authserver_timeout_vlanid
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["eap_auto_untagged_vlans"] = eap_auto_untagged_vlans
            __props__.__dict__["eap_passthru"] = eap_passthru
            __props__.__dict__["framevid_apply"] = framevid_apply
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["guest_auth_delay"] = guest_auth_delay
            __props__.__dict__["guest_vlan"] = guest_vlan
            __props__.__dict__["guest_vlan_id"] = guest_vlan_id
            __props__.__dict__["guest_vlanid"] = guest_vlanid
            __props__.__dict__["mac_auth_bypass"] = mac_auth_bypass
            __props__.__dict__["name"] = name
            __props__.__dict__["open_auth"] = open_auth
            __props__.__dict__["policy_type"] = policy_type
            __props__.__dict__["radius_timeout_overwrite"] = radius_timeout_overwrite
            __props__.__dict__["security_mode"] = security_mode
            __props__.__dict__["user_groups"] = user_groups
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerSecurityPolicy8021X, __self__).__init__(
            'fortios:index/switchControllerSecurityPolicy8021X:SwitchControllerSecurityPolicy8021X',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_fail_vlan: Optional[pulumi.Input[str]] = None,
            auth_fail_vlan_id: Optional[pulumi.Input[str]] = None,
            auth_fail_vlanid: Optional[pulumi.Input[int]] = None,
            authserver_timeout_period: Optional[pulumi.Input[int]] = None,
            authserver_timeout_vlan: Optional[pulumi.Input[str]] = None,
            authserver_timeout_vlanid: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            eap_auto_untagged_vlans: Optional[pulumi.Input[str]] = None,
            eap_passthru: Optional[pulumi.Input[str]] = None,
            framevid_apply: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            guest_auth_delay: Optional[pulumi.Input[int]] = None,
            guest_vlan: Optional[pulumi.Input[str]] = None,
            guest_vlan_id: Optional[pulumi.Input[str]] = None,
            guest_vlanid: Optional[pulumi.Input[int]] = None,
            mac_auth_bypass: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            open_auth: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None,
            radius_timeout_overwrite: Optional[pulumi.Input[str]] = None,
            security_mode: Optional[pulumi.Input[str]] = None,
            user_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerSecurityPolicy8021XUserGroupArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerSecurityPolicy8021X':
        """
        Get an existing SwitchControllerSecurityPolicy8021X resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerSecurityPolicy8021XState.__new__(_SwitchControllerSecurityPolicy8021XState)

        __props__.__dict__["auth_fail_vlan"] = auth_fail_vlan
        __props__.__dict__["auth_fail_vlan_id"] = auth_fail_vlan_id
        __props__.__dict__["auth_fail_vlanid"] = auth_fail_vlanid
        __props__.__dict__["authserver_timeout_period"] = authserver_timeout_period
        __props__.__dict__["authserver_timeout_vlan"] = authserver_timeout_vlan
        __props__.__dict__["authserver_timeout_vlanid"] = authserver_timeout_vlanid
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["eap_auto_untagged_vlans"] = eap_auto_untagged_vlans
        __props__.__dict__["eap_passthru"] = eap_passthru
        __props__.__dict__["framevid_apply"] = framevid_apply
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["guest_auth_delay"] = guest_auth_delay
        __props__.__dict__["guest_vlan"] = guest_vlan
        __props__.__dict__["guest_vlan_id"] = guest_vlan_id
        __props__.__dict__["guest_vlanid"] = guest_vlanid
        __props__.__dict__["mac_auth_bypass"] = mac_auth_bypass
        __props__.__dict__["name"] = name
        __props__.__dict__["open_auth"] = open_auth
        __props__.__dict__["policy_type"] = policy_type
        __props__.__dict__["radius_timeout_overwrite"] = radius_timeout_overwrite
        __props__.__dict__["security_mode"] = security_mode
        __props__.__dict__["user_groups"] = user_groups
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerSecurityPolicy8021X(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authFailVlan")
    def auth_fail_vlan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auth_fail_vlan")

    @property
    @pulumi.getter(name="authFailVlanId")
    def auth_fail_vlan_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auth_fail_vlan_id")

    @property
    @pulumi.getter(name="authFailVlanid")
    def auth_fail_vlanid(self) -> pulumi.Output[int]:
        return pulumi.get(self, "auth_fail_vlanid")

    @property
    @pulumi.getter(name="authserverTimeoutPeriod")
    def authserver_timeout_period(self) -> pulumi.Output[int]:
        return pulumi.get(self, "authserver_timeout_period")

    @property
    @pulumi.getter(name="authserverTimeoutVlan")
    def authserver_timeout_vlan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authserver_timeout_vlan")

    @property
    @pulumi.getter(name="authserverTimeoutVlanid")
    def authserver_timeout_vlanid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authserver_timeout_vlanid")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="eapAutoUntaggedVlans")
    def eap_auto_untagged_vlans(self) -> pulumi.Output[str]:
        return pulumi.get(self, "eap_auto_untagged_vlans")

    @property
    @pulumi.getter(name="eapPassthru")
    def eap_passthru(self) -> pulumi.Output[str]:
        return pulumi.get(self, "eap_passthru")

    @property
    @pulumi.getter(name="framevidApply")
    def framevid_apply(self) -> pulumi.Output[str]:
        return pulumi.get(self, "framevid_apply")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="guestAuthDelay")
    def guest_auth_delay(self) -> pulumi.Output[int]:
        return pulumi.get(self, "guest_auth_delay")

    @property
    @pulumi.getter(name="guestVlan")
    def guest_vlan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "guest_vlan")

    @property
    @pulumi.getter(name="guestVlanId")
    def guest_vlan_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "guest_vlan_id")

    @property
    @pulumi.getter(name="guestVlanid")
    def guest_vlanid(self) -> pulumi.Output[int]:
        return pulumi.get(self, "guest_vlanid")

    @property
    @pulumi.getter(name="macAuthBypass")
    def mac_auth_bypass(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mac_auth_bypass")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openAuth")
    def open_auth(self) -> pulumi.Output[str]:
        return pulumi.get(self, "open_auth")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="radiusTimeoutOverwrite")
    def radius_timeout_overwrite(self) -> pulumi.Output[str]:
        return pulumi.get(self, "radius_timeout_overwrite")

    @property
    @pulumi.getter(name="securityMode")
    def security_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "security_mode")

    @property
    @pulumi.getter(name="userGroups")
    def user_groups(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerSecurityPolicy8021XUserGroup']]]:
        return pulumi.get(self, "user_groups")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

