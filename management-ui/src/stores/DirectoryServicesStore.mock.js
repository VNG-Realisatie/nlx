// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { observable } from 'mobx'

export const mockDirectoryServiceModel = ({
  fetch = jest.fn(),
  ...service
}) => ({
  ...service,
  fetch,
})

export const mockDirectoryServicesStore = ({
  services = [],
  isInitiallyFetched = true,
  error = '',
  fetchServices = jest.fn(),
  selectService = jest.fn(),
}) =>
  observable({
    directoryServicesStore: {
      services,
      isInitiallyFetched,
      error,
      fetchServices,
      selectService,
    },
  })
