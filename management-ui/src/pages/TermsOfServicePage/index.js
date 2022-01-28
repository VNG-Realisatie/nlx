// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useTranslation } from 'react-i18next'
import { Button } from '@commonground/design-system'
import { useNavigate } from 'react-router-dom'
import ToSContext from '../../tos-context'
import {
  Content,
  StyledIconExternalLink,
  StyledNLXManagementLogo,
  StyledSidebar,
  Wrapper,
} from './index.styles'

const TermsOfServicePage = () => {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const tosContext = useContext(ToSContext)

  const tos = tosContext.tos

  const handleAcceptToS = async () => {
    await tosContext.accept()
    navigate('/')
  }

  return !tos ? null : (
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
          <a href={tos.url}>
            {t('Terms of Service')}
            <StyledIconExternalLink />
          </a>
        </p>
        <p>{t('Terms of Service - paragraph 4 of 4')}</p>
        <p>
          <Button type="button" onClick={handleAcceptToS}>
            {t('Confirm agreement')}
          </Button>
        </p>
      </Content>
    </Wrapper>
  )
}

export default TermsOfServicePage
