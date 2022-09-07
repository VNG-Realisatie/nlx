// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ACCESS_REQUEST_STATES } from '../../stores/models/OutgoingAccessRequestModel'

const { FAILED, RECEIVED, REJECTED, APPROVED } = ACCESS_REQUEST_STATES

export const SHOW_REQUEST_ACCESS = 0
export const SHOW_HAS_ACCESS = 1
export const SHOW_REQUEST_FAILED = 2
export const SHOW_REQUEST_RECEIVED = 3
export const SHOW_REQUEST_REJECTED = 4
export const SHOW_ACCESS_REVOKED = 5

export default function getDirectoryServiceAccessUIState(
  outgoingAccessRequest,
  accessProof,
) {
  if (!outgoingAccessRequest) {
    return SHOW_REQUEST_ACCESS
  }

  if (accessProof && accessProof.revokedAt) {
    return SHOW_ACCESS_REVOKED
  }

  if (!accessProof && outgoingAccessRequest.state === APPROVED) {
    return SHOW_REQUEST_RECEIVED
  }

  switch (outgoingAccessRequest.state) {
    case APPROVED:
      return SHOW_HAS_ACCESS

    case FAILED:
      return SHOW_REQUEST_FAILED

    case RECEIVED:
      return SHOW_REQUEST_RECEIVED

    case REJECTED:
      return SHOW_REQUEST_REJECTED

    default:
      throw new Error(
        'Unexpected combination of outgoingAccessRequest and accessProof',
      )
  }
}
