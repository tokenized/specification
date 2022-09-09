
import binascii
import json
import struct
import sys
from google.protobuf.json_format import MessageToJson

import actions_generated
import actions_pb2

OP_FALSE  = 0x00
OP_RETURN = 0x6a
OP_MAX_SINGLE_BYTE_PUSH_DATA = 0x4b
OP_PUSH_DATA_1 = 0x4c
OP_PUSH_DATA_2 = 0x4d
OP_PUSH_DATA_4 = 0x4e

OP_1NEGATE = 0x4f
OP_0  = 0x00
OP_1  = 0x51
OP_2  = 0x52
OP_3  = 0x53
OP_4  = 0x54
OP_5  = 0x55
OP_6  = 0x56
OP_7  = 0x57
OP_8  = 0x58
OP_9  = 0x59
OP_10 = 0x5a
OP_11 = 0x5b
OP_12 = 0x5c
OP_13 = 0x5d
OP_14 = 0x5e
OP_15 = 0x5f
OP_16 = 0x60

PROTOCOL_ID = b'TKN'
TEST_PROTOCOL_ID = b'test.TKN'

def convertToJSON(action):
	return MessageToJson(action)


def deserialize(data, isTest):
	opCode = data[0]
	data = data[1:]
	if opCode != OP_RETURN:
		if opCode != OP_FALSE:
			raise Exception("Not an op return")

		opCode = data[0]
		data = data[1:]

	if opCode != OP_RETURN:
		raise Exception("Not an op return: " + str(opCode))

	data, envelopeProtocolID = parsePushData(data)
	if len(envelopeProtocolID) != 2:
		raise Exception("Not envelope protocol: " + str(envelopeProtocolID))

	if envelopeProtocolID[0] != 0xbd:
		raise Exception("Not envelope protocol: " + str(envelopeProtocolID))

	if envelopeProtocolID[1] != 0x01:
		raise Exception("Not envelope protocol version 1: " + str(envelopeProtocolID))

	data, protocolIDCount = parseInteger(data)
	if not isinstance(protocolIDCount, int):
		raise Exception("Protocol count is not a number: " + str(protocolIDCount))

	if protocolIDCount != 1:
		raise Exception("Only one protocol supported: " + str(protocolIDCount))

	protocolIDs = []
	for i in range(protocolIDCount):
		data, protocolID = parsePushData(data)
		protocolIDs.append(protocolID)

	if isTest:
		if protocolIDs[0] != TEST_PROTOCOL_ID:
			raise Exception("Wrong protocol ID: " + str(protocolIDs[0]))
	else:
		if protocolIDs[0] != PROTOCOL_ID:
			raise Exception("Wrong protocol ID: " + str(protocolIDs[0]))

	data, payloadCount = parseInteger(data)
	if not isinstance(payloadCount, int):
		raise Exception("Payload count is not an integer: " + str(payloadCount))

	if payloadCount != 2 and payloadCount != 3:
		raise Exception("Wrong payload count: " + str(payloadCount))

	data, version = parseInteger(data)
	if not isinstance(version, int):
		raise Exception("Version is not a number: " + str(version))

	if version != 0:
		raise Exception("Only Tokenized version 0 supported: " + str(version))

	# Parse the action code push data from the script
	data, actionCode = parsePushData(data)
	if len(actionCode) != 2:
		raise Exception("Wrong action code size: " + str(actionCode))

	if payloadCount > 2:
		# Parse push data from script that contains the protobuf for the action.
		data, actionData = parsePushData(data)

		# Parse protobuf data.
		return actions_generated.parseActionObject(actionCode, actionData)

	# Empty object
	return actions_generated.getActionObject(actionCode)


def parseInteger(data):
	opCode = data[0]

	if opCode == OP_FALSE:
		return data[1:], 0
	elif opCode >= OP_1 and opCode <= OP_16:
		return data[1:], opCode - OP_1 + 1
	elif opCode == OP_1NEGATE:
		return data[1:], -1
	else:
		data, pushData = parsePushData(data)
		value = convertPushDataToNumber(pushData)
		return data, value

def convertPushDataToNumber(data):
	if len(data) == 1:
		return struct.unpack("<B", data[:1])[0]
	elif len(data) == 2:
		return struct.unpack("<H", data[:2])[0]
	elif len(data) == 4:
		return struct.unpack("<I", data[:4])[0]
	else:
		data = bytes(8-len(data))+data
		return struct.unpack("<Q", data)[0]

def parsePushData(data):
	opCode = data[0]
	data = data[1:]

	if opCode <= OP_MAX_SINGLE_BYTE_PUSH_DATA:
		opData = data[:opCode]
		data = data[opCode:]
		return data, opData
	elif opCode == OP_PUSH_DATA_1:
		size = struct.unpack("<B", data[:1])[0]
		data = data[1:]
		opData = data[:size]
		data = data[size:]
		return data, opData
	elif opCode == OP_PUSH_DATA_2:
		size = struct.unpack("<H", data[:2])[0]
		data = data[2:]
		opData = data[:size]
		data = data[size:]
		return data, opData
	elif opCode == OP_PUSH_DATA_4:
		size = struct.unpack("<I", data[:4])[0]
		data = data[4:]
		opData = data[:size]
		data = data[size:]
		return data, opData
	else:
		raise Exception("Not push data: " + str(opCode))

if __name__ == "__main__":
	if len(sys.argv) != 2:
		print("Invalid arguments : protocol.py <hex locking script>")
		sys.exit(1)

	data = bytearray.fromhex(sys.argv[1])

	action = deserialize(data, True)
	print("Action : " + MessageToJson(action))
