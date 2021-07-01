// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { node } from 'prop-types'
import { useTranslation } from 'react-i18next'
import PrimaryNavigation from '../PrimaryNavigation'
import Header from './Header'
import HeaderWithBackNavigation from './HeaderWithBackNavigation'
import OrganizationInwayCheck from './OrganizationInwayCheck'
import { Page, SkipToContent, MainWrapper, Main } from './index.styles'

const contentId = 'content'

const PageTemplate = ({ children }) => {
  const { t } = useTranslation()
  return (
    <Page>
      <SkipToContent href={`#${contentId}`}>{t('Go to content')}</SkipToContent>
      <PrimaryNavigation />
      <MainWrapper id={contentId}>
        <OrganizationInwayCheck />
        <Main>{children}</Main>
      </MainWrapper>
    </Page>
  )
}

PageTemplate.propTypes = {
  children: node,
}

PageTemplate.Header = Header
PageTemplate.HeaderWithBackNavigation = HeaderWithBackNavigation

export default PageTemplate
