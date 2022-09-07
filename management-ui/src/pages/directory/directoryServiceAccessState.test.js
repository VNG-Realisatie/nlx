// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../stores/models/OutgoingAccessRequestModel'
import AccessProofModel from '../../stores/models/AccessProofModel'
import getDirectoryServiceAccessState, {
  SHOW_ACCESS_REVOKED,
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_ACCESS,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_REJECTED,
} from './directoryServiceAccessState'

const { FAILED, RECEIVED, REJECTED, APPROVED } = ACCESS_REQUEST_STATES

const createOutgoingAccessRequestInstance = (accessRequestData) => {
  return new OutgoingAccessRequestModel({
    accessRequestData,
    outgoingAccessRequestStore: {},
  })
}

test('access request does not exist', () => {
  expect(getDirectoryServiceAccessState()).toBe(SHOW_REQUEST_ACCESS)
})

test('access request has failed', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: FAILED,
      }),
    ),
  ).toBe(SHOW_REQUEST_FAILED)
})

test('access request is received', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: RECEIVED,
      }),
    ),
  ).toBe(SHOW_REQUEST_RECEIVED)
})

test('access request is approved but has no access proof yet', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: APPROVED,
      }),
      null,
    ),
  ).toBe(SHOW_REQUEST_RECEIVED)
})

test('access request is approved', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: APPROVED,
      }),
      new AccessProofModel({
        accessProofData: {},
      }),
    ),
  ).toBe(SHOW_HAS_ACCESS)
})

test('access request is rejected', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: REJECTED,
      }),
    ),
  ).toBe(SHOW_REQUEST_REJECTED)
})

test('access request is approved, but revoked', () => {
  expect(
    getDirectoryServiceAccessState(
      new OutgoingAccessRequestModel({
        accessRequestData: {
          state: APPROVED,
        },
      }),
      new AccessProofModel({
        accessProofData: {
          revokedAt: new Date('2020-10-02'),
        },
      }),
    ),
  ).toBe(SHOW_ACCESS_REVOKED)
})
