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
        <p data-testid="store-a">{stores.subStoreA}</p>
        <p data-testid="store-b">{stores.subStoreB}</p>
      </>
    )
  }

  const store = {
    subStoreA: 'store a',
    subStoreB: 'store b',
  }

  const { getByTestId } = render(
    <StoreProvider store={store}>
      <Consumer />
    </StoreProvider>,
  )

  expect(getByTestId('store-a')).toHaveTextContent('store a')
  expect(getByTestId('store-b')).toHaveTextContent('store b')
})
