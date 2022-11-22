
import actions_pb2


# CodeContractOffer identifies a payload as a ContractOffer action message.
CodeContractOffer = b'C1'

# CodeContractFormation identifies a payload as a ContractFormation action message.
CodeContractFormation = b'C2'

# CodeContractAmendment identifies a payload as a ContractAmendment action message.
CodeContractAmendment = b'C3'

# CodeStaticContractFormation identifies a payload as a StaticContractFormation action message.
CodeStaticContractFormation = b'C4'

# CodeContractAddressChange identifies a payload as a ContractAddressChange action message.
CodeContractAddressChange = b'C5'

# CodeBodyOfAgreementOffer identifies a payload as a BodyOfAgreementOffer action message.
CodeBodyOfAgreementOffer = b'C6'

# CodeBodyOfAgreementFormation identifies a payload as a BodyOfAgreementFormation action message.
CodeBodyOfAgreementFormation = b'C7'

# CodeBodyOfAgreementAmendment identifies a payload as a BodyOfAgreementAmendment action message.
CodeBodyOfAgreementAmendment = b'C8'

# CodeInstrumentDefinition identifies a payload as a InstrumentDefinition action message.
CodeInstrumentDefinition = b'I1'

# CodeInstrumentCreation identifies a payload as a InstrumentCreation action message.
CodeInstrumentCreation = b'I2'

# CodeInstrumentModification identifies a payload as a InstrumentModification action message.
CodeInstrumentModification = b'I3'

# CodeTransfer identifies a payload as a Transfer action message.
CodeTransfer = b'T1'

# CodeSettlement identifies a payload as a Settlement action message.
CodeSettlement = b'T2'

# CodeRectificationSettlement identifies a payload as a RectificationSettlement action message.
CodeRectificationSettlement = b'T3'

# CodeProposal identifies a payload as a Proposal action message.
CodeProposal = b'G1'

# CodeVote identifies a payload as a Vote action message.
CodeVote = b'G2'

# CodeBallotCast identifies a payload as a BallotCast action message.
CodeBallotCast = b'G3'

# CodeBallotCounted identifies a payload as a BallotCounted action message.
CodeBallotCounted = b'G4'

# CodeResult identifies a payload as a Result action message.
CodeResult = b'G5'

# CodeOrder identifies a payload as a Order action message.
CodeOrder = b'E1'

# CodeFreeze identifies a payload as a Freeze action message.
CodeFreeze = b'E2'

# CodeThaw identifies a payload as a Thaw action message.
CodeThaw = b'E3'

# CodeConfiscation identifies a payload as a Confiscation action message.
CodeConfiscation = b'E4'

# CodeDeprecatedReconciliation identifies a payload as a DeprecatedReconciliation action message.
CodeDeprecatedReconciliation = b'E5'

# CodeEstablishment identifies a payload as a Establishment action message.
CodeEstablishment = b'R1'

# CodeAddition identifies a payload as a Addition action message.
CodeAddition = b'R2'

# CodeAlteration identifies a payload as a Alteration action message.
CodeAlteration = b'R3'

# CodeRemoval identifies a payload as a Removal action message.
CodeRemoval = b'R4'

# CodeMessage identifies a payload as a Message action message.
CodeMessage = b'M1'

# CodeRejection identifies a payload as a Rejection action message.
CodeRejection = b'M2'


def getActionObject(actionCode):
	if actionCode == CodeContractOffer:
		return actions_pb2.ContractOffer()
	if actionCode == CodeContractFormation:
		return actions_pb2.ContractFormation()
	if actionCode == CodeContractAmendment:
		return actions_pb2.ContractAmendment()
	if actionCode == CodeStaticContractFormation:
		return actions_pb2.StaticContractFormation()
	if actionCode == CodeContractAddressChange:
		return actions_pb2.ContractAddressChange()
	if actionCode == CodeBodyOfAgreementOffer:
		return actions_pb2.BodyOfAgreementOffer()
	if actionCode == CodeBodyOfAgreementFormation:
		return actions_pb2.BodyOfAgreementFormation()
	if actionCode == CodeBodyOfAgreementAmendment:
		return actions_pb2.BodyOfAgreementAmendment()
	if actionCode == CodeInstrumentDefinition:
		return actions_pb2.InstrumentDefinition()
	if actionCode == CodeInstrumentCreation:
		return actions_pb2.InstrumentCreation()
	if actionCode == CodeInstrumentModification:
		return actions_pb2.InstrumentModification()
	if actionCode == CodeTransfer:
		return actions_pb2.Transfer()
	if actionCode == CodeSettlement:
		return actions_pb2.Settlement()
	if actionCode == CodeRectificationSettlement:
		return actions_pb2.RectificationSettlement()
	if actionCode == CodeProposal:
		return actions_pb2.Proposal()
	if actionCode == CodeVote:
		return actions_pb2.Vote()
	if actionCode == CodeBallotCast:
		return actions_pb2.BallotCast()
	if actionCode == CodeBallotCounted:
		return actions_pb2.BallotCounted()
	if actionCode == CodeResult:
		return actions_pb2.Result()
	if actionCode == CodeOrder:
		return actions_pb2.Order()
	if actionCode == CodeFreeze:
		return actions_pb2.Freeze()
	if actionCode == CodeThaw:
		return actions_pb2.Thaw()
	if actionCode == CodeConfiscation:
		return actions_pb2.Confiscation()
	if actionCode == CodeDeprecatedReconciliation:
		return actions_pb2.DeprecatedReconciliation()
	if actionCode == CodeEstablishment:
		return actions_pb2.Establishment()
	if actionCode == CodeAddition:
		return actions_pb2.Addition()
	if actionCode == CodeAlteration:
		return actions_pb2.Alteration()
	if actionCode == CodeRemoval:
		return actions_pb2.Removal()
	if actionCode == CodeMessage:
		return actions_pb2.Message()
	if actionCode == CodeRejection:
		return actions_pb2.Rejection()

	raise Exception("Unsupported action code: ", str(actionCode))

def parseActionObject(actionCode, actionData):
	action = getActionObject(actionCode)
	action.ParseFromString(actionData)
	return action
