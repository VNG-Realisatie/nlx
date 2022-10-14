// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../stores/models/OutgoingAccessRequestModel'
import { ManagementServiceApi } from '../api'
import OutgoingAccessRequestStore from './OutgoingAccessRequestStore'
import { RootStore } from './index'

test('initializing the store', () => {
  const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
    rootStore: new RootStore(),
  })

  expect(outgoingAccessRequestStore.outgoingAccessRequests.size).toEqual(0)
})

test('sending an outgoing access request', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceSendAccessRequest = jest
    .fn()
    .mockResolvedValue({
      outgoingAccessRequest: {
        id: '42',
        organization: {
          serialNumber: '00000000000000000001',
          name: 'organization-name',
        },
        serviceName: 'service-name',
        state: ACCESS_REQUEST_STATES.RECEIVED,
        createdAt: '2020-10-07T13:01:11.288349Z',
        updatedAt: null,
      },
    })

  const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
    managementApiClient,
  })

  const outgoingAccessRequest = await outgoingAccessRequestStore.send(
    '00000000000000000001',
    'service-name',
    'public-key-pem',
  )

  expect(
    managementApiClient.managementServiceSendAccessRequest,
  ).toHaveBeenCalledWith({
    body: {
      organizationSerialNumber: '00000000000000000001',
      serviceName: 'service-name',
      publicKeyPem: 'public-key-pem',
    },
  })

  expect(outgoingAccessRequest).toBeInstanceOf(OutgoingAccessRequestModel)
  expect(outgoingAccessRequest.id).toEqual('42')
  expect(outgoingAccessRequest.organization.serialNumber).toEqual(
    '00000000000000000001',
  )
  expect(outgoingAccessRequest.organization.name).toEqual('organization-name')
  expect(outgoingAccessRequest.serviceName).toEqual('service-name')
  expect(outgoingAccessRequest.state).toEqual(ACCESS_REQUEST_STATES.RECEIVED)
  expect(outgoingAccessRequest.createdAt).toEqual(
    new Date('2020-10-07T13:01:11.288349Z'),
  )
  expect(outgoingAccessRequest.updatedAt).toBeNull()
})

test('updating from server', async () => {
  const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
    rootStore: new RootStore(),
  })

  expect(outgoingAccessRequestStore.updateFromServer()).toBeNull()

  let outgoingAccessRequestModel =
    await outgoingAccessRequestStore.updateFromServer({
      id: '42',
      state: ACCESS_REQUEST_STATES.RECEIVED,
    })

  // new model should be created
  expect(outgoingAccessRequestStore.outgoingAccessRequests.size).toEqual(1)
  expect(outgoingAccessRequestModel).toBeInstanceOf(OutgoingAccessRequestModel)

  outgoingAccessRequestModel =
    await outgoingAccessRequestStore.updateFromServer({
      id: '42',
      state: ACCESS_REQUEST_STATES.APPROVED,
    })

  // existing model should be updated
  expect(outgoingAccessRequestStore.outgoingAccessRequests.size).toEqual(1)
  expect(outgoingAccessRequestModel).toBeInstanceOf(OutgoingAccessRequestModel)
  expect(outgoingAccessRequestModel.state).toEqual(
    ACCESS_REQUEST_STATES.APPROVED,
  )
})
