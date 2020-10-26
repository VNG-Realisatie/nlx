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

const SHOW_REQUEST_ACCESS = 0
const SHOW_HAS_ACCESS = 1
const SHOW_REQUEST_CREATED = 2
const SHOW_REQUEST_FAILED = 3
const SHOW_REQUEST_RECEIVED = 4
const SHOW_REQUEST_CANCELLED = 5
const SHOW_REQUEST_REJECTED = 6

export default function getDirectoryServiceAccessUIState(
  outgoingAccessRequest,
  accessProof,
) {
  if (accessProof && !accessProof.revokedAt) {
    return SHOW_HAS_ACCESS
  }

  // It should not be possible to have accessProof and no access request
  if (!outgoingAccessRequest && !accessProof) {
    return SHOW_REQUEST_ACCESS
  }

  if (accessProof && accessProof.revokedAt > outgoingAccessRequest.createdAt) {
    return SHOW_REQUEST_ACCESS
  }

  if (!accessProof || accessProof.revokedAt < outgoingAccessRequest.createdAt) {
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

export {
  SHOW_REQUEST_ACCESS,
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_CREATED,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_CANCELLED,
  SHOW_REQUEST_REJECTED,
}
