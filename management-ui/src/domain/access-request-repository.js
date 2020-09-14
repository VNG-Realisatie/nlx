// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import env from '../env'
import { throwOnError } from './fetch-utils'

class AccessRequestRepository {
  static async requestAccess(payload) {
    const response = await fetch(
      `${env.managementApiBaseUrl}/v1/access-requests`,
      {
        method: 'POST',
        body: JSON.stringify(payload),
      },
    )

    throwOnError(response, {
      409: 'Request already sent, please refresh the page to see the latest state.',
    })

    return await response.json()
  }
}

export default AccessRequestRepository
