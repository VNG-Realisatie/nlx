// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { observable } from 'mobx'

export const mockServicesStore = ({
  services = [],
  isInitiallyFetched = true,
  error = '',
  fetchServices = jest.fn(),
  selectService = jest.fn(),
  addService = jest.fn(),
  removeService = jest.fn(),
}) =>
  observable({
    servicesStore: {
      services,
      isInitiallyFetched,
      error,
      fetchServices,
      selectService,
      removeService,
      addService,
    },
  })
export const mockServiceModel = (service, fetch = jest.fn()) => ({
  ...service,
  fetch,
})
