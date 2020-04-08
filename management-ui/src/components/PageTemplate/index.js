// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { node, string } from 'prop-types'
import PrimaryNavigation from '../PrimaryNavigation'
import {
  StyledMain,
  StyledContent,
  StyledPageDescription,
  StyledPageTitle,
  StyledUserNavigation,
  StyledPageHeader,
} from './index.styles'

const PageTemplate = ({ title, description, children }) => {
  return (
    <StyledMain>
      <PrimaryNavigation />
      <StyledContent>
        <StyledPageHeader>
          {title && <StyledPageTitle>{title}</StyledPageTitle>}
          <StyledUserNavigation fullName="John Doe" />
        </StyledPageHeader>
        <StyledPageDescription>{description}</StyledPageDescription>
        {children}
      </StyledContent>
    </StyledMain>
  )
}

PageTemplate.propTypes = {
  title: string,
  description: string,
  children: node,
}

PageTemplate.defaultProps = {
  description: '\u00A0',
}

export default PageTemplate
