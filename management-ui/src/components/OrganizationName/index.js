// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { func, bool } from 'prop-types'
import usePromise from '../../hooks/use-promise'
import EnvironmentRepository from '../../domain/environment-repository'
import { StyledTextWithEllipsis } from './index.styles'

const OrganizationName = ({ getEnvironment, isHeader, ...props }) => {
  const { result } = usePromise(getEnvironment)

  return result ? (
    <StyledTextWithEllipsis
      {...(isHeader && { title: result.organizationName })}
      isHeader={isHeader}
      {...props}
    >
      {result.organizationName}
    </StyledTextWithEllipsis>
  ) : null
}

OrganizationName.propTypes = {
  getEnvironment: func,
  isHeader: bool,
}

OrganizationName.defaultProps = {
  getEnvironment: EnvironmentRepository.getCurrent,
  isHeader: false,
}

export default OrganizationName
