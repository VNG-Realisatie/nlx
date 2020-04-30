// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { func } from 'prop-types'

const DirectoryServiceCount = ({ directoryServices, ...props }) => {
  const count = directoryServices().length
  return <span {...props}> ({count})</span>
}

DirectoryServiceCount.propTypes = {
  directoryServices: func.isRequired,
}

export default DirectoryServiceCount
