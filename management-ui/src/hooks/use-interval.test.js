// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { render } from '@testing-library/react'
import React from 'react'
import useInterval, { DEFAULT_INTERVAL } from './use-interval'

jest.useFakeTimers()

const pollFunctionA = jest.fn()
const pollFunctionB = jest.fn()

const CompA = () => {
  useInterval(pollFunctionA)
  return <p>component A</p>
}

const CompB = () => {
  useInterval(pollFunctionB)
  return <p>component B</p>
}

test('sets and removes polling functions', () => {
  const { rerender } = render(
    <>
      <CompA />
      <CompB />
    </>,
  )

  jest.advanceTimersByTime(DEFAULT_INTERVAL * 2)
  expect(pollFunctionA).toHaveBeenCalledTimes(2)
  expect(pollFunctionB).toHaveBeenCalledTimes(2)
  pollFunctionA.mockClear()
  pollFunctionB.mockClear()

  rerender(
    <>
      <CompA />
    </>,
  )

  jest.advanceTimersByTime(DEFAULT_INTERVAL)
  expect(pollFunctionA).toHaveBeenCalledTimes(1)
  expect(pollFunctionB).toHaveBeenCalledTimes(0)
})
