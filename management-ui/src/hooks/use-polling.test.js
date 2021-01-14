// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { render, fireEvent } from '@testing-library/react'
import usePolling, { INTERVAL } from './use-polling'

const pollFunctionA = jest.fn()
const pollFunctionB = jest.fn()

beforeEach(() => {
  jest.useFakeTimers()
})

afterEach(() => {
  pollFunctionA.mockReset()
  pollFunctionB.mockReset()
})

const CompA = () => {
  const [pausePolling, continuePolling] = usePolling(pollFunctionA)
  return (
    <>
      <button onClick={pausePolling}>pause A polling</button>
      <button onClick={continuePolling}>continue A polling</button>
    </>
  )
}

const CompB = () => {
  usePolling(pollFunctionB)
  return <p>component B</p>
}

test('sets and removes polling functions', () => {
  const { rerender } = render(
    <>
      <CompA />
      <CompB />
    </>,
  )

  jest.advanceTimersByTime(INTERVAL * 2)
  expect(pollFunctionA).toHaveBeenCalledTimes(2)
  expect(pollFunctionB).toHaveBeenCalledTimes(2)
  pollFunctionA.mockClear()
  pollFunctionB.mockClear()

  rerender(
    <>
      <CompA />
    </>,
  )

  jest.advanceTimersByTime(INTERVAL)
  expect(pollFunctionA).toHaveBeenCalledTimes(1)
  expect(pollFunctionB).toHaveBeenCalledTimes(0)
})

test('pause and continue polling functions', () => {
  const { getByText } = render(
    <>
      <CompA />
    </>,
  )

  fireEvent.click(getByText('pause A polling'))
  jest.advanceTimersByTime(INTERVAL * 2)
  expect(pollFunctionA).toHaveBeenCalledTimes(0)

  fireEvent.click(getByText('continue A polling'))
  jest.advanceTimersByTime(INTERVAL)
  expect(pollFunctionA).toHaveBeenCalledTimes(1)
})

test('function can only be added once', () => {
  const { getByText } = render(
    <>
      <CompA />
    </>,
  )

  fireEvent.click(getByText('continue A polling'))
  fireEvent.click(getByText('continue A polling'))
  jest.advanceTimersByTime(INTERVAL)
  expect(pollFunctionA).toHaveBeenCalledTimes(1)
})
