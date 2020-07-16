// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'
const delay = (ms) => new Promise((resolve) => setTimeout(() => resolve(), ms))

class AccessRequestRepository {
  static async requestAccess(payload) {
    const response = await fetch(`/api/v1/access-requests`, {
      method: 'POST',
      body: JSON.stringify(payload),
    })

    throwOnError(response, {
      409: 'Request already sent, please refresh the page to see the latest status.',
    })

    await delay(2000)
    // throw new Error('yada')
    return await response.json()
  }
}

export default AccessRequestRepository
