// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { Navigate, Routes, Route } from 'react-router-dom'
import PageTemplate from '../../components/PageTemplate'
import GeneralSettings from './GeneralSettings'
import Navigation from './Navigation'
import { Wrapper, SettingsNav, Content } from './index.styles'

const SettingsPage = () => {
  const { t } = useTranslation()

  return (
    <PageTemplate>
      <PageTemplate.Header title={t('Settings')} id="settings-header" />

      <Wrapper>
        <SettingsNav aria-labelledby="settings-header">
          <Navigation />
        </SettingsNav>
        <Content>
          <Routes>
            <Route index element={<Navigate to="general" />} />
            <Route path="general" element={<GeneralSettings />} />
          </Routes>
        </Content>
      </Wrapper>
    </PageTemplate>
  )
}

export default SettingsPage
