// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

class InwayRepository {
  static async getAll() {
    const result = await fetch(`/api/v1/inways`)

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    const response = await result.json()
    return response.inways ? response.inways : []
  }

  static async getByName(name) {
    // Actually searches by key
    const result = await fetch(`/api/v1/inways/${name}`)

    if (result.status === 404) {
      throw new Error('not found')
    }

    if (result.status === 403) {
      throw new Error('forbidden')
    }

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    const inway = await result.json()

    if (inway.services) {
      inway.services = inway.services.map((service) => service.name)
    } else {
      inway.services = []
    }

    console.log(inway)
    return inway
  }
}

export default InwayRepository
