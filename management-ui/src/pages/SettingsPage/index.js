// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { Redirect, Switch, Route, useRouteMatch } from 'react-router-dom'
import PageTemplate from '../../components/PageTemplate'
import GeneralSettings from './GeneralSettings'

const SettingsPage = () => {
  const { t } = useTranslation()
  const { path } = useRouteMatch('/settings')

  return (
    <PageTemplate>
      <PageTemplate.Header title={t('Settings')} />

      <Switch>
        <Redirect exact path={path} to={`${path}/general`} />

        <Route path={`${path}/general`} component={GeneralSettings} />
      </Switch>
    </PageTemplate>
  )
}

export default SettingsPage
