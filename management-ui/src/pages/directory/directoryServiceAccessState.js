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
// Used until we can match accessProof with accessRequest
// We can't use date/time because timestamps are generated from different "clocks"
export const SHOW_ACCESS_REVOKED = 7

export default function getDirectoryServiceAccessUIState(
  outgoingAccessRequest,
  accessProof,
) {
  // It should not be possible to have accessProof and no access request
  if (!outgoingAccessRequest && !accessProof) {
    return SHOW_REQUEST_ACCESS
  }

  const hasNewerAccessRequest =
    outgoingAccessRequest &&
    accessProof &&
    accessProof.accessRequestId !== outgoingAccessRequest.id

  if (!hasNewerAccessRequest) {
    if (!accessProof.revokedAt) {
      return SHOW_HAS_ACCESS
    }

    if (accessProof.revokedAt) {
      return SHOW_ACCESS_REVOKED
    }
  } else {
    if (outgoingAccessRequest.state === CREATED) return SHOW_REQUEST_CREATED
    if (outgoingAccessRequest.state === FAILED) return SHOW_REQUEST_FAILED
    if (outgoingAccessRequest.state === RECEIVED) return SHOW_REQUEST_RECEIVED
    if (outgoingAccessRequest.state === APPROVED) return SHOW_REQUEST_RECEIVED
    if (outgoingAccessRequest.state === CANCELLED) return SHOW_REQUEST_CANCELLED
    if (outgoingAccessRequest.state === REJECTED) return SHOW_REQUEST_REJECTED
  }

  throw Error(
    'We have not forseen this combination of data, please report to a developer:',
    outgoingAccessRequest,
    accessProof,
  )
}
