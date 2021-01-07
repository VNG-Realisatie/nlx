// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../stores/models/OutgoingAccessRequestModel'
import AccessProofModel from '../../stores/models/AccessProofModel'
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

test('access request does not exist', () => {
  expect(getDirectoryServiceAccessState()).toBe(SHOW_REQUEST_ACCESS)
})

test('access request is created', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: CREATED,
      }),
    ),
  ).toBe(SHOW_REQUEST_CREATED)
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

test('access request is approved', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: APPROVED,
      }),
    ),
  ).toBe(SHOW_REQUEST_RECEIVED)
})

test('access request is cancelled', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: CANCELLED,
        createdAt: '2020-10-01',
      }),
    ),
  ).toBe(SHOW_REQUEST_CANCELLED)
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

test('access request is approved and has an access proof', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        id: '42',
        state: APPROVED,
      }),
      new AccessProofModel({
        accessProofData: {
          accessRequestId: '42',
        },
      }),
    ),
  ).toBe(SHOW_HAS_ACCESS)
})

test('access request is approved, but access proof has been revoked', () => {
  expect(
    getDirectoryServiceAccessState(
      new OutgoingAccessRequestModel({
        accessRequestData: {
          id: '42',
          state: APPROVED,
          createdAt: '2020-10-01',
        },
      }),
      new AccessProofModel({
        accessProofData: { revokedAt: '2020-10-02', accessRequestId: '42' },
      }),
    ),
  ).toBe(SHOW_ACCESS_REVOKED)
})

test('access request is received, and access proof from previous access request has been revoked', () => {
  expect(
    getDirectoryServiceAccessState(
      new OutgoingAccessRequestModel({
        accessRequestData: {
          id: '42',
          state: RECEIVED,
          createdAt: '2020-10-01',
        },
      }),
      new AccessProofModel({
        accessProofData: { revokedAt: '2020-10-02', id: '43' },
      }),
    ),
  ).toBe(SHOW_REQUEST_RECEIVED)
})
