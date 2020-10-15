// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { observable } from 'mobx'

export const mockInwaysStore = ({
  inways = [],
  isInitiallyFetched = true,
  error = '',
  fetchInways = jest.fn(),
  selectInway = jest.fn(),
}) =>
  observable({
    inwaysStore: {
      inways,
      isInitiallyFetched,
      error,
      fetchInways,
      selectInway,
    },
  })

export const mockInwayModel = (inway, fetch = jest.fn()) => ({
  ...inway,
  fetch,
})
