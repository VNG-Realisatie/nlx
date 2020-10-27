// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../models/OutgoingAccessRequestModel'
import AccessProofModel from '../../models/AccessProofModel'

import getDirectoryServiceAccessState, {
  SHOW_REQUEST_ACCESS,
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_CREATED,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_CANCELLED,
  SHOW_REQUEST_REJECTED,
  SHOW_ACCESS_REVOKED,
} from './directoryServiceAccessState'

const {
  CREATED,
  FAILED,
  RECEIVED,
  CANCELLED,
  REJECTED,
  APPROVED,
} = ACCESS_REQUEST_STATES

const createOutgoingAccessRequestInstance = (accessRequestData) => {
  return new OutgoingAccessRequestModel({
    accessRequestData,
    outgoingAccessRequestStore: {},
  })
}

const createAccessProofInstance = (accessProofData) => {
  return new AccessProofModel({ accessProofData })
}

describe('Return SHOW_REQUEST_ACCESS when', () => {
  it('latestAccessRequest does not exist', () => {
    expect(getDirectoryServiceAccessState(null, null)).toBe(SHOW_REQUEST_ACCESS)
  })
})

describe('Return SHOW_ACCESS_REVOKED when', () => {
  it('latestAccessProof is revoked', () => {
    expect(
      getDirectoryServiceAccessState(
        createOutgoingAccessRequestInstance({
          state: APPROVED,
          createdAt: '2020-10-01',
        }),
        createAccessProofInstance({ revokedAt: '2020-10-02' }),
      ),
    ).toBe(SHOW_ACCESS_REVOKED)
  })
})

describe('Return SHOW_HAS_ACCESS when', () => {
  it('latestAccessProof is not revoked', () => {
    expect(
      getDirectoryServiceAccessState(
        createOutgoingAccessRequestInstance({
          state: APPROVED,
          createdAt: '2020-10-01',
        }),
        createAccessProofInstance({ revokedAt: null }),
      ),
    ).toBe(SHOW_HAS_ACCESS)
  })
})

describe('Return SHOW_REQUEST_CREATED when', () => {
  it('latestAccessRequest is created and not having latestAccessProof', () => {
    expect(
      getDirectoryServiceAccessState(
        createOutgoingAccessRequestInstance({
          state: CREATED,
          createdAt: '2020-10-01',
        }),
        null,
      ),
    ).toBe(SHOW_REQUEST_CREATED)
  })
})

describe('Return SHOW_REQUEST_FAILED when', () => {
  it('latestAccessRequest is failed while not having latestAccessProof', () => {
    expect(
      getDirectoryServiceAccessState(
        createOutgoingAccessRequestInstance({
          state: FAILED,
          createdAt: '2020-10-01',
        }),
        null,
      ),
    ).toBe(SHOW_REQUEST_FAILED)
  })
})

describe('Return SHOW_REQUEST_RECEIVED when', () => {
  it('latestAccessRequest is received while not having latestAccessProof', () => {
    expect(
      getDirectoryServiceAccessState(
        createOutgoingAccessRequestInstance({
          state: RECEIVED,
          createdAt: '2020-10-01',
        }),
        null,
      ),
    ).toBe(SHOW_REQUEST_RECEIVED)
  })

  it('latestAccessRequest is approved while not having latestAccessProof', () => {
    expect(
      getDirectoryServiceAccessState(
        createOutgoingAccessRequestInstance({
          state: APPROVED,
          createdAt: '2020-10-01',
        }),
        null,
      ),
    ).toBe(SHOW_REQUEST_RECEIVED)
  })
})

describe('Return SHOW_REQUEST_CANCELLED when', () => {
  it('latestAccessRequest is cancelled while not having latestAccessProof', () => {
    expect(
      getDirectoryServiceAccessState(
        createOutgoingAccessRequestInstance({
          state: CANCELLED,
          createdAt: '2020-10-01',
        }),
        null,
      ),
    ).toBe(SHOW_REQUEST_CANCELLED)
  })
})

describe('Return SHOW_REQUEST_REJECTED when', () => {
  it('latestAccessRequest is received while not having latestAccessProof', () => {
    expect(
      getDirectoryServiceAccessState(
        createOutgoingAccessRequestInstance({
          state: REJECTED,
          createdAt: '2020-10-01',
        }),
        null,
      ),
    ).toBe(SHOW_REQUEST_REJECTED)
  })
})
