// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import {
  StyledHeader,
  StyledPageTitle,
  StyledDescription,
  StyledUserNavigation,
} from './index.styles'

const Header = ({ title, description }) => {
  return (
    <>
      <StyledHeader>
        {title && <StyledPageTitle>{title}</StyledPageTitle>}
        <StyledUserNavigation fullName="John Doe" />
      </StyledHeader>
      <StyledDescription>{description}</StyledDescription>
    </>
  )
}

Header.propTypes = {
  title: string,
  description: string,
}

Header.defaultProps = {
  description: '\u00A0',
}

export default Header
