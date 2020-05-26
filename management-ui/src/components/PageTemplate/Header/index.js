// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { node, string } from 'prop-types'
import UserNavigation from '../../UserNavigation'
import OrganizationName from '../../OrganizationName'
import {
  StyledHeader,
  StyledPageTitle,
  StyledDescription,
  StyledHeaderItems,
} from './index.styles'

const Header = ({ title, description }) => (
  <>
    <StyledHeader>
      {title && <StyledPageTitle>{title}</StyledPageTitle>}
      <StyledHeaderItems>
        <OrganizationName isHeader />
        <UserNavigation />
      </StyledHeaderItems>
    </StyledHeader>
    <StyledDescription>{description}</StyledDescription>
  </>
)

Header.propTypes = {
  title: string,
  description: node,
}

Header.defaultProps = {
  description: '\u00A0',
}

export default Header
