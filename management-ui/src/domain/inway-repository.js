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
}

export default InwayRepository
