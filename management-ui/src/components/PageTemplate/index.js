// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { node } from 'prop-types'

import PrimaryNavigation from '../PrimaryNavigation'
import OrganizationInwayCheck from '../OrganizationInwayCheck'
import Header from './Header'
import HeaderWithBackNavigation from './HeaderWithBackNavigation'
import { Page, MainWrapper, Main } from './index.styles'

const PageTemplate = ({ children }) => (
  <Page>
    <PrimaryNavigation />
    <MainWrapper>
      <OrganizationInwayCheck />
      <Main>{children}</Main>
    </MainWrapper>
  </Page>
)

PageTemplate.propTypes = {
  children: node,
}

PageTemplate.Header = Header
PageTemplate.HeaderWithBackNavigation = HeaderWithBackNavigation

export default PageTemplate
