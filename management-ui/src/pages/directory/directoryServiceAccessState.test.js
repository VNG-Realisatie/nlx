// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../stores/models/OutgoingAccessRequestModel'
import getDirectoryServiceAccessState, {
  SHOW_ACCESS_REVOKED,
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_ACCESS,
  SHOW_REQUEST_CANCELLED,
  SHOW_REQUEST_CREATED,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_REJECTED,
} from './directoryServiceAccessState'

const { CREATED, FAILED, RECEIVED, CANCELLED, REJECTED, APPROVED } =
  ACCESS_REQUEST_STATES

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
  ).toBe(SHOW_HAS_ACCESS)
})

test('access request is cancelled', () => {
  expect(
    getDirectoryServiceAccessState(
      createOutgoingAccessRequestInstance({
        state: CANCELLED,
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

test('access request is approved, but revoked', () => {
  expect(
    getDirectoryServiceAccessState(
      new OutgoingAccessRequestModel({
        accessRequestData: {
          state: APPROVED,
        },
      }),
      new Date('2020-10-02'),
    ),
  ).toBe(SHOW_ACCESS_REVOKED)
})
