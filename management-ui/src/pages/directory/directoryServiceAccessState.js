// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ACCESS_REQUEST_STATES } from '../../stores/models/OutgoingAccessRequestModel'

const { CREATED, FAILED, RECEIVED, REJECTED, APPROVED } = ACCESS_REQUEST_STATES

export const SHOW_REQUEST_ACCESS = 0
export const SHOW_HAS_ACCESS = 1
export const SHOW_REQUEST_CREATED = 2
export const SHOW_REQUEST_FAILED = 3
export const SHOW_REQUEST_RECEIVED = 4
export const SHOW_REQUEST_REJECTED = 5
export const SHOW_ACCESS_REVOKED = 6

export default function getDirectoryServiceAccessUIState(
  outgoingAccessRequest,
  revokedAt,
) {
  if (!outgoingAccessRequest) {
    return SHOW_REQUEST_ACCESS
  }

  if (revokedAt) {
    return SHOW_ACCESS_REVOKED
  }

  switch (outgoingAccessRequest.state) {
    case CREATED:
      return SHOW_REQUEST_CREATED

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
