// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Switch, Route, Redirect } from 'react-router-dom'

import LoginPage from '../pages/LoginPage/index'
import ServicesPage from '../pages/ServicesPage'
import InwaysPage from '../pages/InwaysPage'
import AddServicePage from '../pages/AddServicePage'
import DirectoryPage from '../pages/DirectoryPage'
import EditServicePage from '../pages/EditServicePage'
import NotFoundPage from '../pages/NotFoundPage'

import AuthenticatedRoute, { LoginRoutePath } from './authenticated-route'

const Routes = () => {
  return (
    <Switch>
      <Redirect exact path="/" to="/inways" />
      <Route path={LoginRoutePath} component={LoginPage} />

      <AuthenticatedRoute exact path="/inways" component={InwaysPage} />
      <AuthenticatedRoute
        path="/services/add-service"
        component={AddServicePage}
      />
      <AuthenticatedRoute
        path="/services/:name/edit-service"
        component={EditServicePage}
      />
      <AuthenticatedRoute path="/services" component={ServicesPage} />
      <AuthenticatedRoute path="/directory" component={DirectoryPage} />

      <Route path="*" component={NotFoundPage} />
    </Switch>
  )
}

export default Routes
