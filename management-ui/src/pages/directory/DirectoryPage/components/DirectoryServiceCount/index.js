// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { array } from 'prop-types'

const DirectoryServiceCount = ({ services, ...props }) => (
  <span {...props}> ({services.length})</span>
)

DirectoryServiceCount.propTypes = {
  services: array.isRequired,
}

export default DirectoryServiceCount
