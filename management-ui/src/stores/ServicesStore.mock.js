// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { observable } from 'mobx'

export const mockServicesStore = ({
  services = [],
  isInitiallyFetched = true,
  error = '',
  fetch = jest.fn(),
  fetchAll = jest.fn(),
  getService = jest.fn(),
  create = jest.fn(),
  update = jest.fn(),
  removeService = jest.fn(),
}) =>
  observable({
    servicesStore: {
      services,
      isInitiallyFetched,
      error,
      fetch,
      fetchAll,
      getService,
      create,
      update,
      removeService,
    },
  })
