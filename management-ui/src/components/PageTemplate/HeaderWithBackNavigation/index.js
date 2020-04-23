// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import BackButton from '../../BackButton'
import { StyledTitle } from './index.styles'

const HeaderWithBackNavigation = ({ title, backButtonTo }) => {
  return (
    <>
      <p>
        <BackButton to={backButtonTo} />
      </p>
      <StyledTitle>{title}</StyledTitle>
    </>
  )
}

HeaderWithBackNavigation.propTypes = {
  title: string,
  backButtonTo: string.isRequired,
}

export default HeaderWithBackNavigation
