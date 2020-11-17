// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { Redirect, Switch, Route, useRouteMatch } from 'react-router-dom'

import PageTemplate from '../../components/PageTemplate'
import GeneralSettings from './GeneralSettings'
import InsightSettings from './InsightSettings'
import Navigation from './Navigation'
import { StyledContent, StyledMain, StyledSidebar } from './index.styles'

const SettingsPage = () => {
  const { t } = useTranslation()
  const { path } = useRouteMatch('/settings')

  return (
    <PageTemplate>
      <PageTemplate.Header title={t('Settings')} />

      <StyledMain>
        <StyledSidebar>
          <Navigation />
        </StyledSidebar>
        <StyledContent>
          <Switch>
            <Redirect exact path={path} to={`${path}/general`} />

            <Route path={`${path}/general`} component={GeneralSettings} />
            <Route path={`${path}/insight`} component={InsightSettings} />
          </Switch>
        </StyledContent>
      </StyledMain>
    </PageTemplate>
  )
}

export default SettingsPage
