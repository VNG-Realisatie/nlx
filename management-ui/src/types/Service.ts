// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

interface Service {
  organization: {
    serialNumber: string
    name: string
  }
  service: string
  serviceName: string
  state: string
  apiSpecificationType?: string
  latestAccessRequest?: Record<string, unknown>
  latestAccessProof?: Record<string, unknown>

  requestAccess(): void
  retryRequestAccess(): void
}

export default Service
