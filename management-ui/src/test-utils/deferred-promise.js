// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
const deferredPromise = () => {
  let res, rej
  const p = new Promise((resolve, reject) => {
    res = resolve
    rej = reject
  })
  p.resolve = res
  p.reject = rej
  return p
}

export default deferredPromise
