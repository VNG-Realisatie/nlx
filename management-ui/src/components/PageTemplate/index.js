import React from 'react'
import { element, string } from 'prop-types'
import PrimaryNavigation from '../PrimaryNavigation'
import UserNavigation from '../UserNavigation'
import {
  StyledMain,
  StyledContent,
  StyledPageDescription,
  StyledPageTitle,
  StyledPageHeader,
} from './index.styles'

const PageTemplate = ({ title, description, children }) => {
  return (
    <StyledMain>
      <PrimaryNavigation />
      <StyledContent>
        <StyledPageHeader>
          <StyledPageTitle>{title}</StyledPageTitle>
          <UserNavigation fullName="John Doe" />
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
  children: element,
}

PageTemplate.defaultProps = {
  description: '\u00A0',
}

export default PageTemplate
