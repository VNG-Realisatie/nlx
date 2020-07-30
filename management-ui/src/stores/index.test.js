// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { render } from '../test-utils'
import useStores from '../hooks/use-stores'
import { StoreProvider } from './index'

test('store should be provided by useStore hook', () => {
  const Consumer = () => {
    const stores = useStores()

    return (
      <>
        {Object.keys(stores).map((store) => {
          const mockValue = stores[store]
          return <p key={mockValue}>{mockValue}</p>
        })}
      </>
    )
  }

  const store = {
    subStoreA: 'store a',
    subStoreB: 'store b',
  }

  const { getByText } = render(
    <StoreProvider store={store}>
      <Consumer />
    </StoreProvider>,
  )

  expect(getByText('store a')).toBeInTheDocument()
  expect(getByText('store b')).toBeInTheDocument()
})
