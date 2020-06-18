// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { array, shape } from 'prop-types'

const DirectoryServiceCount = ({ directoryServices, ...props }) => {
  const { services = [] } = directoryServices
  return <span {...props}> ({services.length})</span>
}

DirectoryServiceCount.propTypes = {
  directoryServices: shape({ services: array.isRequired }).isRequired,
}

export default DirectoryServiceCount
