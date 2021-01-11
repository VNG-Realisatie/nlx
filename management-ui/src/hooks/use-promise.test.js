// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { renderHook, act } from '@testing-library/react-hooks'
import deferredPromise from '../utils/deferred-promise'
import usePromise from './use-promise'

test('with a resolving promise', async () => {
  const promise = deferredPromise()

  const handler = () => promise
  const { result } = renderHook(() => usePromise(handler))
  expect(result.current).toEqual({
    error: null,
    isReady: false,
    result: null,
    reload: expect.any(Function),
  })

  await act(async () => {
    promise.resolve('arbitrary message')
  })

  expect(result.current).toEqual({
    error: null,
    isReady: true,
    result: 'arbitrary message',
    reload: expect.any(Function),
  })
})

test('with a rejecting promise', async () => {
  const promise = deferredPromise()

  const handler = () => promise
  const { result } = renderHook(() => usePromise(handler))

  expect(result.current).toEqual({
    error: null,
    isReady: false,
    result: null,
    reload: expect.any(Function),
  })

  await act(async () => {
    promise.reject(new Error('arbitrary message'))
  })

  expect(result.current).toEqual({
    error: new Error('arbitrary message'),
    isReady: true,
    result: null,
    reload: expect.any(Function),
  })
})

test('with an argument', async () => {
  const handler = async (argument) => argument
  const { result, waitForNextUpdate } = renderHook(() =>
    usePromise(handler, 'arbitrary argument'),
  )

  expect(result.current).toEqual({
    error: null,
    isReady: false,
    result: null,
    reload: expect.any(Function),
  })

  await waitForNextUpdate()

  expect(result.current).toEqual({
    error: null,
    isReady: true,
    result: 'arbitrary argument',
    reload: expect.any(Function),
  })
})

test('reloading a resource', async () => {
  const handler = jest
    .fn()
    .mockResolvedValueOnce('first-result')
    .mockResolvedValueOnce('second-result')
  const { result, waitForNextUpdate } = renderHook(() => usePromise(handler))

  expect(result.current).toEqual({
    error: null,
    isReady: false,
    result: null,
    reload: expect.any(Function),
  })

  await waitForNextUpdate()

  expect(result.current).toEqual({
    error: null,
    isReady: true,
    result: 'first-result',
    reload: expect.any(Function),
  })

  act(() => result.current.reload())
  await waitForNextUpdate()
  expect(result.current).toEqual({
    error: null,
    isReady: true,
    result: 'second-result',
    reload: expect.any(Function),
  })
  expect(handler).toBeCalledTimes(2)
})
