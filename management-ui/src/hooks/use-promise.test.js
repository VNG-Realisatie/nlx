// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { renderHook, act } from '@testing-library/react-hooks'
import usePromise from './use-promise'

test('with a resolving promise', async () => {
  let promiseResolve

  const promise = new Promise((resolve) => {
    promiseResolve = resolve
  })

  const handler = () => promise
  const { result } = renderHook(() => usePromise(handler))
  expect(result.current).toStrictEqual({
    error: null,
    loading: true,
    result: null,
  })

  await act(async () => {
    promiseResolve('arbitrary message')
  })

  expect(result.current).toStrictEqual({
    error: null,
    loading: false,
    result: 'arbitrary message',
  })
})

test('with a rejecting promise', async () => {
  let promiseReject

  const promise = new Promise((resolve, reject) => {
    promiseReject = reject
  })

  const handler = () => promise
  const { result } = renderHook(() => usePromise(handler))
  expect(result.current).toStrictEqual({
    error: null,
    loading: true,
    result: null,
  })

  await act(async () => {
    promiseReject(new Error('arbitrary message'))
  })

  expect(result.current).toStrictEqual({
    error: new Error('arbitrary message'),
    loading: false,
    result: null,
  })
})
