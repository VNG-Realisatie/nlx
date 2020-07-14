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

    throwOnError(response)

    return await response.json()
  }
}

export default AccessRequestRepository
