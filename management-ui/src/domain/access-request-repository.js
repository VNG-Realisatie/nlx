// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

class AccessRequestRepository {
  static async createAccessRequest({ organizationName, serviceName }) {
    const response = await fetch(`/api/v1/access-requests`, {
      method: 'POST',
      body: JSON.stringify({
        organizationName: organizationName,
        serviceName: serviceName,
      }),
    })

    throwOnError(response, {
      409: 'Request already sent, please refresh the page to see the latest state.',
    })

    return await response.json()
  }

  static async listIncomingAccessRequests(serviceName) {
    const response = await fetch(
      `/api/v1/access-requests/incoming/services/${serviceName}`,
    )

    throwOnError(response, {
      500: 'Failed to retrieve incoming access requests.',
    })

    const result = await response.json()
    return result.accessRequests || []
  }

  static async approveIncomingAccessRequest({ serviceName, id }) {
    const response = await fetch(
      `/api/v1/access-requests/incoming/services/${serviceName}/${id}/approve`,
      { method: 'POST' },
    )

    throwOnError(response, {
      404: 'Request not found, please refresh the page to see the latest state.',
      409: 'Request already approved, please refresh the page to see the latest state.',
    })

    return null
  }

  static async rejectIncomingAccessRequest({ serviceName, id }) {
    const response = await fetch(
      `/api/v1/access-requests/incoming/services/${serviceName}/${id}/reject`,
      { method: 'POST' },
    )

    throwOnError(response, {
      404: 'Request not found, please refresh the page to see the latest state.',
    })

    return null
  }

  static async sendAccessRequest({ organizationName, serviceName, id }) {
    const response = await fetch(
      `/api/v1/access-requests/outgoing/organizations/${organizationName}/services/${serviceName}/${id}/send`,
      { method: 'POST' },
    )

    throwOnError(response)

    const result = await response.json()
    return result || {}
  }
}

export default AccessRequestRepository
