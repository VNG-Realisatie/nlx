// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent } from '@testing-library/react'

import { renderWithProviders } from '../../test-utils'
import Collapsible from './index'

describe('Collapsible', () => {
  let utils
  beforeEach(() => {
    jest.useFakeTimers()

    utils = renderWithProviders(
      <Collapsible title={<span data-testid="title">title</span>}>
        <p data-testid="body">body</p>
      </Collapsible>,
    )
  })

  it('should render only title initially', () => {
    expect(utils.getByTestId('title').textContent).toBe('title')
    expect(utils.queryByTestId('body')).toBeNull()
  })

  it('should render the body after clicking the title', () => {
    fireEvent.click(utils.getByTestId('title'))
    expect(utils.getByTestId('body').textContent).toBe('body')
  })

  describe('when opened', () => {
    beforeEach(() => {
      // to open the Collapsible
      fireEvent.click(utils.getByTestId('title'))
    })
    it('should hide the body after clicking the title', () => {
      fireEvent.click(utils.getByTestId('title'))

      jest.runAllTimers()

      expect(utils.queryByTestId('body')).toBeNull()
    })

    it('should hide the body after clicking the body', () => {
      fireEvent.click(utils.getByTestId('body'))

      jest.runAllTimers()

      expect(utils.queryByTestId('body')).toBeNull()
    })
  })
})
