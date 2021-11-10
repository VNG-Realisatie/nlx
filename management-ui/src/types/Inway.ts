// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

interface Inway {
  name: string
  ipAddress: string
  hostname: string
  selfAddress: string
  version: string
  services: {
    name: string
  }[]
}

export default Inway
