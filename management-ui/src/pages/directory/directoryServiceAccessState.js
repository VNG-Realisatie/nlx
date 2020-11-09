// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ACCESS_REQUEST_STATES } from '../../models/OutgoingAccessRequestModel'

const {
  CREATED,
  FAILED,
  RECEIVED,
  CANCELLED,
  REJECTED,
  APPROVED,
} = ACCESS_REQUEST_STATES

export const SHOW_REQUEST_ACCESS = 0
export const SHOW_HAS_ACCESS = 1
export const SHOW_REQUEST_CREATED = 2
export const SHOW_REQUEST_FAILED = 3
export const SHOW_REQUEST_RECEIVED = 4
export const SHOW_REQUEST_CANCELLED = 5
export const SHOW_REQUEST_REJECTED = 6
export const SHOW_ACCESS_REVOKED = 7

export default function getDirectoryServiceAccessUIState(
  outgoingAccessRequest,
  accessProof,
) {
  if (!outgoingAccessRequest) {
    return SHOW_REQUEST_ACCESS
  }

  if (accessProof && accessProof.accessRequestId !== outgoingAccessRequest.id) {
    accessProof = null
  }

  if (accessProof && !accessProof.revokedAt) {
    return SHOW_HAS_ACCESS
  }

  if (accessProof && accessProof.revokedAt) {
    return SHOW_ACCESS_REVOKED
  }

  switch (outgoingAccessRequest.state) {
    case CREATED:
      return SHOW_REQUEST_CREATED

    case FAILED:
      return SHOW_REQUEST_FAILED

    case RECEIVED:
    case APPROVED:
      return SHOW_REQUEST_RECEIVED

    case CANCELLED:
      return SHOW_REQUEST_CANCELLED

    case REJECTED:
      return SHOW_REQUEST_REJECTED

    default:
      throw new Error(
        'unexpected combination of outgoingAccessRequest and accessProof',
      )
  }
}
