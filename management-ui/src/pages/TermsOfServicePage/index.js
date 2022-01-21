// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Button } from '@commonground/design-system'
import {
  Content,
  StyledIconExternalLink,
  StyledNLXManagementLogo,
  StyledSidebar,
  Wrapper,
} from './index.styles'

const TermsOfServicePage = () => {
  const { t } = useTranslation()
  const [termsOfServiceUrl, setTermsAndServiceUrl] = useState()

  useEffect(() => {
    const fetchUrl = async () => {
      const url = await Promise.resolve(
        'https://directory.nlx.io/Gebruiksvoorwaarden_NLX_maart_2020.pdf',
      )
      setTermsAndServiceUrl(url)
    }

    fetchUrl()
  }, [])

  return (
    <Wrapper>
      <StyledSidebar>
        <StyledNLXManagementLogo />
      </StyledSidebar>
      <Content>
        <h1>{t('Terms of Service')}</h1>
        <p>{t('Terms of Service - paragraph 1 of 4')}</p>
        <p>{t('Terms of Service - paragraph 2 of 4')}</p>
        <p>{t('Terms of Service - paragraph 3 of 4')}</p>
        <p>
          <a href={termsOfServiceUrl}>
            {termsOfServiceUrl}
            <StyledIconExternalLink />
          </a>
        </p>
        <p>{t('Terms of Service - paragraph 4 of 4')}</p>
        <p>
          <Button type="button">{t('Confirm agreement')}</Button>
        </p>
      </Content>
    </Wrapper>
  )
}

export default TermsOfServicePage
