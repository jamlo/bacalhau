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

class ApiPutNodeRequest(object):
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
        'action': 'str',
        'credential': 'ApiHTTPCredential',
        'idempotency_token': 'str',
        'message': 'str',
        'namespace': 'str',
        'node_id': 'str'
    }

    attribute_map = {
        'action': 'action',
        'credential': 'credential',
        'idempotency_token': 'idempotencyToken',
        'message': 'message',
        'namespace': 'namespace',
        'node_id': 'nodeID'
    }

    def __init__(self, action=None, credential=None, idempotency_token=None, message=None, namespace=None, node_id=None):  # noqa: E501
        """ApiPutNodeRequest - a model defined in Swagger"""  # noqa: E501
        self._action = None
        self._credential = None
        self._idempotency_token = None
        self._message = None
        self._namespace = None
        self._node_id = None
        self.discriminator = None
        if action is not None:
            self.action = action
        if credential is not None:
            self.credential = credential
        if idempotency_token is not None:
            self.idempotency_token = idempotency_token
        if message is not None:
            self.message = message
        if namespace is not None:
            self.namespace = namespace
        if node_id is not None:
            self.node_id = node_id

    @property
    def action(self):
        """Gets the action of this ApiPutNodeRequest.  # noqa: E501


        :return: The action of this ApiPutNodeRequest.  # noqa: E501
        :rtype: str
        """
        return self._action

    @action.setter
    def action(self, action):
        """Sets the action of this ApiPutNodeRequest.


        :param action: The action of this ApiPutNodeRequest.  # noqa: E501
        :type: str
        """

        self._action = action

    @property
    def credential(self):
        """Gets the credential of this ApiPutNodeRequest.  # noqa: E501


        :return: The credential of this ApiPutNodeRequest.  # noqa: E501
        :rtype: ApiHTTPCredential
        """
        return self._credential

    @credential.setter
    def credential(self, credential):
        """Sets the credential of this ApiPutNodeRequest.


        :param credential: The credential of this ApiPutNodeRequest.  # noqa: E501
        :type: ApiHTTPCredential
        """

        self._credential = credential

    @property
    def idempotency_token(self):
        """Gets the idempotency_token of this ApiPutNodeRequest.  # noqa: E501


        :return: The idempotency_token of this ApiPutNodeRequest.  # noqa: E501
        :rtype: str
        """
        return self._idempotency_token

    @idempotency_token.setter
    def idempotency_token(self, idempotency_token):
        """Sets the idempotency_token of this ApiPutNodeRequest.


        :param idempotency_token: The idempotency_token of this ApiPutNodeRequest.  # noqa: E501
        :type: str
        """

        self._idempotency_token = idempotency_token

    @property
    def message(self):
        """Gets the message of this ApiPutNodeRequest.  # noqa: E501


        :return: The message of this ApiPutNodeRequest.  # noqa: E501
        :rtype: str
        """
        return self._message

    @message.setter
    def message(self, message):
        """Sets the message of this ApiPutNodeRequest.


        :param message: The message of this ApiPutNodeRequest.  # noqa: E501
        :type: str
        """

        self._message = message

    @property
    def namespace(self):
        """Gets the namespace of this ApiPutNodeRequest.  # noqa: E501


        :return: The namespace of this ApiPutNodeRequest.  # noqa: E501
        :rtype: str
        """
        return self._namespace

    @namespace.setter
    def namespace(self, namespace):
        """Sets the namespace of this ApiPutNodeRequest.


        :param namespace: The namespace of this ApiPutNodeRequest.  # noqa: E501
        :type: str
        """

        self._namespace = namespace

    @property
    def node_id(self):
        """Gets the node_id of this ApiPutNodeRequest.  # noqa: E501


        :return: The node_id of this ApiPutNodeRequest.  # noqa: E501
        :rtype: str
        """
        return self._node_id

    @node_id.setter
    def node_id(self, node_id):
        """Sets the node_id of this ApiPutNodeRequest.


        :param node_id: The node_id of this ApiPutNodeRequest.  # noqa: E501
        :type: str
        """

        self._node_id = node_id

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
        if issubclass(ApiPutNodeRequest, dict):
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
        if not isinstance(other, ApiPutNodeRequest):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other