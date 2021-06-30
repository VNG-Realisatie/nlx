// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

// via https://stackoverflow.com/a/41434763/363448
class LocalStorageMock {
  constructor() {
    this.store = {}
  }

  clear() {
    this.store = {}
  }

  getItem(key) {
    // eslint-disable-next-line security/detect-object-injection
    return this.store[key] || null
  }

  setItem(key, value) {
    // eslint-disable-next-line security/detect-object-injection
    this.store[key] = String(value)
  }

  removeItem(key) {
    // eslint-disable-next-line security/detect-object-injection
    delete this.store[key]
  }
}

export default LocalStorageMock
