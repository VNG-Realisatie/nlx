// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { RootStore } from '../index'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'

let accessRequestData

beforeEach(() => {
  accessRequestData = {
    id: 'abcd',
    organization: {
      serialNumber: '00000000000000000001',
      name: 'Organization',
    },
    serviceName: 'Service',
    state: ACCESS_REQUEST_STATES.RECEIVED,
    createdAt: '2020-10-01',
    updatedAt: '2020-10-02',
    publicKeyFingerprint: 'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
    errorDetails: {
      cause: 'the cause of an error',
    },
  }
})

test('should properly construct object', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData,
    outgoingAccessRequestStore: {},
  })

  expect(accessRequest.id).toBe(accessRequestData.id)
  expect(accessRequest.organization.name).toBe(
    accessRequestData.organization.name,
  )
  expect(accessRequest.serviceName).toBe(accessRequestData.serviceName)
  expect(accessRequest.state).toBe(accessRequestData.state)
  expect(accessRequest.createdAt).toEqual(new Date(accessRequestData.createdAt))
  expect(accessRequest.updatedAt).toEqual(new Date(accessRequestData.updatedAt))
  expect(accessRequest.publicKeyFingerprint).toEqual(
    'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
  )
  expect(accessRequest.errorDetails.cause).toEqual('the cause of an error')
})

test('organization name is empty', () => {
  const model = new OutgoingAccessRequestModel({
    outgoingAccessRequestStore: {},
    accessRequestData: {
      organization: {
        name: '',
        serialNumber: '00000000000000000001',
      },
    },
  })

  expect(model.organization.name).toBe('00000000000000000001')
  expect(model.organization.serialNumber).toBe('00000000000000000001')
})

test('terminate access', async () => {
  configure({ safeDescriptors: false })

  const rootStore = new RootStore({})

  const model = new OutgoingAccessRequestModel({
    outgoingAccessRequestStore: rootStore.outgoingAccessRequestStore,
    accessRequestData: {
      organization: {
        name: '',
        serialNumber: '00000000000000000001',
      },
    },
  })

  jest
    .spyOn(rootStore.outgoingAccessRequestStore, 'terminate')
    .mockResolvedValue()

  await model.terminate()

  expect(rootStore.outgoingAccessRequestStore.terminate).toHaveBeenCalledWith(
    model,
  )
})

test('withdraw access', async () => {
  configure({ safeDescriptors: false })

  const rootStore = new RootStore({})

  const model = new OutgoingAccessRequestModel({
    outgoingAccessRequestStore: rootStore.outgoingAccessRequestStore,
    accessRequestData: {
      organization: {
        name: '',
        serialNumber: '00000000000000000001',
      },
    },
  })

  jest
    .spyOn(rootStore.outgoingAccessRequestStore, 'withdraw')
    .mockResolvedValue()

  await model.withdraw()

  expect(rootStore.outgoingAccessRequestStore.withdraw).toHaveBeenCalledWith(
    model,
  )
})
