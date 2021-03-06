// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { checkPropTypes } from 'prop-types'
import { RootStore } from '../index'
import InwayModel, { inwayModelPropTypes } from './InwayModel'

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const inwayModel = new InwayModel({ store: {}, inway: { name: 'service-a' } })

  checkPropTypes(inwayModelPropTypes, inwayModel, 'prop', 'InwayModel')

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('fetch should reload the model via the store', async () => {
  configure({ safeDescriptors: false })
  const rootStore = new RootStore({})

  const inwayModel = new InwayModel({
    store: rootStore.inwayStore,
    inway: {
      name: 'service-a',
    },
  })

  jest.spyOn(rootStore.inwayStore, 'fetch').mockResolvedValue({
    name: 'service-a',
  })

  await inwayModel.fetch()

  expect(rootStore.inwayStore.fetch).toHaveBeenCalledWith({
    name: 'service-a',
  })
})
