// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
class AccessRequestRepository {
  static async requestAccess(payload) {
    const result = await fetch(`/api/v1/access-requests`, {
      method: 'POST',
      body: JSON.stringify(payload),
    })

    if (result.status === 409) {
      alert(
        'Request already sent, please refresh the page to see the latest status.',
      )
      throw new Error(
        'Request already sent, please refresh the page to see the latest status.',
      )
    }

    if (result.status.toString().substring(0, 1) === '4') {
      alert('Generic error')
      throw new Error(`Error code: ${result.status}`)
    }

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    return await result.json()
  }
}

export default AccessRequestRepository
