// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'
import { RootStore } from '../stores'
import InwayModel, { inwayModelPropTypes } from './InwayModel'

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const inwayModel = new InwayModel({ store: {}, inway: { name: 'service-a' } })

  checkPropTypes(inwayModelPropTypes, inwayModel, 'prop', 'InwayModel')

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('fetch should reload the model via the store', async () => {
  const rootStore = new RootStore({
    inwayRepository: {
      getByName: jest.fn().mockResolvedValue({ name: 'service-a' }),
    },
  })

  const inwayModel = new InwayModel({
    store: rootStore.inwaysStore,
    inway: {
      name: 'service-a',
    },
  })

  jest.spyOn(rootStore.inwaysStore, 'fetch')

  await inwayModel.fetch()

  expect(rootStore.inwaysStore.fetch).toHaveBeenCalledWith({
    name: 'service-a',
  })
})
