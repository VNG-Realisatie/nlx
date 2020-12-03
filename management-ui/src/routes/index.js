// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Redirect, Route, Switch } from 'react-router-dom'

import LoginPage from '../pages/LoginPage/index'
import ServicesPage from '../pages/services/ServicesPage'
import InwaysPage from '../pages/inways/InwaysPage'
import AddServicePage from '../pages/services/AddServicePage'
import DirectoryPage from '../pages/directory/DirectoryPage'
import EditServicePage from '../pages/services/EditServicePage'
import NotFoundPage from '../pages/NotFoundPage'

import SettingsPage from '../pages/SettingsPage'
import AuthenticatedRoute, { LoginRoutePath } from './authenticated-route'

const Routes = () => {
  return (
    <Switch>
      <Redirect exact path="/" to="/inways" />
      <Route path={LoginRoutePath} component={LoginPage} />

      <AuthenticatedRoute path="/inways/:name?" component={InwaysPage} />

      <AuthenticatedRoute
        path="/services/add-service"
        component={AddServicePage}
      />
      <AuthenticatedRoute
        path="/services/:name/edit-service"
        component={EditServicePage}
      />
      <AuthenticatedRoute path="/services/:name?" component={ServicesPage} />
      <AuthenticatedRoute
        path="/directory/:organization?/:name?"
        component={DirectoryPage}
      />
      <AuthenticatedRoute path="/settings" component={SettingsPage} />

      <Route path="*" component={NotFoundPage} />
    </Switch>
  )
}

export default Routes
