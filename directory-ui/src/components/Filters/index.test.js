// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Filters from './index'

describe('Filters', () => {
  describe('changing the text input value', () => {
    it('should call the onQueryChanged handler with the query', () => {
      const onQueryChangedSpy = jest.fn()

      const { getByPlaceholderText } = render(
        <Filters onQueryChanged={onQueryChangedSpy} />,
      )

      const input = getByPlaceholderText(
        'Search for an organization or service…',
      )

      userEvent.clear(input)
      userEvent.type(input, 'abc')

      expect(onQueryChangedSpy).toHaveBeenCalledWith('abc')
    })
  })

  describe('toggling the offline filter', () => {
    it('should call the onStatusFilterChanged handler with the checked state', () => {
      const onStatusFilterChangedSpy = jest.fn()

      const { getByLabelText } = render(
        <Filters onStatusFilterChanged={onStatusFilterChangedSpy} />,
      )

      const input = getByLabelText('Include offline')

      userEvent.click(input)

      expect(onStatusFilterChangedSpy).toHaveBeenCalledWith(false)
    })
  })
})
