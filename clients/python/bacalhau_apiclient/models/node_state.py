# coding: utf-8

"""
    Bacalhau API

    This page is the reference of the Bacalhau REST API. Project docs are available at https://docs.bacalhau.org/. Find more information about Bacalhau at https://github.com/bacalhau-project/bacalhau.  # noqa: E501

    OpenAPI spec version: ${VERSION}
    Contact: team@bacalhau.org
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""

import pprint
import re  # noqa: F401

import six

class NodeState(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """
    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'connection': 'NodeConnectionState',
        'info': 'NodeInfo',
        'membership': 'NodeMembershipState'
    }

    attribute_map = {
        'connection': 'Connection',
        'info': 'Info',
        'membership': 'Membership'
    }

    def __init__(self, connection=None, info=None, membership=None):  # noqa: E501
        """NodeState - a model defined in Swagger"""  # noqa: E501
        self._connection = None
        self._info = None
        self._membership = None
        self.discriminator = None
        if connection is not None:
            self.connection = connection
        if info is not None:
            self.info = info
        if membership is not None:
            self.membership = membership

    @property
    def connection(self):
        """Gets the connection of this NodeState.  # noqa: E501


        :return: The connection of this NodeState.  # noqa: E501
        :rtype: NodeConnectionState
        """
        return self._connection

    @connection.setter
    def connection(self, connection):
        """Sets the connection of this NodeState.


        :param connection: The connection of this NodeState.  # noqa: E501
        :type: NodeConnectionState
        """

        self._connection = connection

    @property
    def info(self):
        """Gets the info of this NodeState.  # noqa: E501


        :return: The info of this NodeState.  # noqa: E501
        :rtype: NodeInfo
        """
        return self._info

    @info.setter
    def info(self, info):
        """Sets the info of this NodeState.


        :param info: The info of this NodeState.  # noqa: E501
        :type: NodeInfo
        """

        self._info = info

    @property
    def membership(self):
        """Gets the membership of this NodeState.  # noqa: E501


        :return: The membership of this NodeState.  # noqa: E501
        :rtype: NodeMembershipState
        """
        return self._membership

    @membership.setter
    def membership(self, membership):
        """Sets the membership of this NodeState.


        :param membership: The membership of this NodeState.  # noqa: E501
        :type: NodeMembershipState
        """

        self._membership = membership

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(NodeState, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, NodeState):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other