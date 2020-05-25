// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { fireEvent } from '@testing-library/react'
import { renderHook } from '@testing-library/react-hooks'
import useClickOutside from './use-click-outside'

test('a click inside component should do nothing', async () => {
  const ref = {
    current: {
      contains: () => true,
    },
  }
  const cb = jest.fn(() => {})

  renderHook(() => useClickOutside(ref, cb))
  fireEvent.mouseDown(document.body)

  expect(cb).not.toHaveBeenCalled()
})

test('a click outside component should fire callback', async () => {
  const ref = {
    current: {
      contains: () => false,
    },
  }
  const cb = jest.fn(() => {})

  renderHook(() => useClickOutside(ref, cb))
  fireEvent.mouseDown(document.body)

  expect(cb).toHaveBeenCalled()
})
