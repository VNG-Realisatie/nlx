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
  fetchAll = jest.fn(),
  getService = jest.fn(),
}) =>
  observable({
    directoryServicesStore: {
      services,
      isInitiallyFetched,
      error,
      fetchAll,
      getService,
    },
  })
