// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

class AccessRequestRepository {
  static async requestAccess(payload) {
    const response = await fetch(`/api/v1/access-requests`, {
      method: 'POST',
      body: JSON.stringify(payload),
    })

    throwOnError(response, {
      409: 'Request already sent, please refresh the page to see the latest status.',
    })

    return await response.json()
  }
}

export default AccessRequestRepository
