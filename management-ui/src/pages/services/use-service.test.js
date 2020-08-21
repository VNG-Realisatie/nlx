// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { renderHook } from '@testing-library/react-hooks'
import { observable } from 'mobx'
import React from 'react'
import { StoreProvider } from '../../stores'
import useService from './use-service'

const storeWrapper = (store) => ({ children }) => (
  <StoreProvider store={store}>{children}</StoreProvider>
)

export const mockServicesStore = ({
  services = [],
  isReady = true,
  error = '',
  fetchServices = jest.fn(),
  selectService = jest.fn(),
  addService = jest.fn(),
  removeService = jest.fn(),
}) =>
  observable({
    servicesStore: {
      services,
      isReady,
      error,
      fetchServices,
      selectService,
      removeService,
      addService,
    },
  })

test('fetching and selecting a service', async () => {
  const selectedServiceFetch = jest.fn()
  const selectService = jest.fn().mockReturnValue({
    name: 'my-service',
    fetch: selectedServiceFetch,
  })
  const fetchServices = jest.fn().mockImplementation(() => {
    store.servicesStore.services = [
      {
        name: 'my-service',
      },
    ]
    store.servicesStore.isReady = true
  })
  const store = mockServicesStore({
    services: [],
    isReady: false,
    error: '',
    fetchServices,
    selectService,
  })

  const wrapper = storeWrapper(store)
  const { result, rerender } = renderHook(() => useService('my-service'), {
    wrapper,
  })
  expect(result.current).toEqual([null, '', false])
  expect(fetchServices).toHaveBeenCalledTimes(1)

  rerender()

  expect(selectService).toHaveBeenCalledTimes(1)
  expect(selectedServiceFetch).toHaveBeenCalledTimes(1)
  expect(result.current).toEqual([
    expect.objectContaining({ name: 'my-service' }),
    '',
    true,
  ])
})

test('an error is returned when the service is not found', async () => {
  const selectedServiceFetch = jest.fn()
  const selectService = jest.fn()
  const fetchServices = jest.fn().mockImplementation(() => {
    store.servicesStore.services = [
      {
        name: 'other-service',
      },
    ]
    store.servicesStore.isReady = true
  })
  const store = mockServicesStore({
    services: [],
    isReady: false,
    error: '',
    fetchServices,
    selectService,
  })

  const wrapper = storeWrapper(store)
  const { result, rerender } = renderHook(() => useService('my-service'), {
    wrapper,
  })
  expect(result.current).toEqual([null, '', false])
  expect(fetchServices).toHaveBeenCalledTimes(1)

  rerender()

  expect(selectService).toHaveBeenCalledTimes(1)
  expect(selectedServiceFetch).toHaveBeenCalledTimes(0)
  expect(result.current).toEqual([null, 'Service not found', true])
})
