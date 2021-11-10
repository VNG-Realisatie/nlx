// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

// p.resolve and p.reject have been confirmed to work correctly
/* eslint-disable @typescript-eslint/ban-ts-comment */

const deferredPromise = <T>(): Promise<T> => {
  let res, rej
  const p = new Promise<T>((resolve, reject) => {
    res = resolve
    rej = reject
  })
  // @ts-ignore
  p.resolve = res
  // @ts-ignore
  p.reject = rej
  return p
}

export default deferredPromise
