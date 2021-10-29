// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { elementType } from 'prop-types'
import { Route, Switch, Redirect } from 'react-router-dom'
import LoginOIDCPage from '../pages/LoginOIDCPage'
import ServicesPage from '../pages/services/ServicesPage'
import InwaysOverview from '../pages/inways-and-outways/OverviewPage'
import AddServicePage from '../pages/services/AddServicePage'
import DirectoryPage from '../pages/directory/DirectoryPage'
import EditServicePage from '../pages/services/EditServicePage'
import FinancePage from '../pages/FinancePage'
import AuditLogPage from '../pages/AuditLogPage'
import SettingsPage from '../pages/SettingsPage'
import NotFoundPage from '../pages/NotFoundPage'
import OrdersPage from '../pages/orders/OrdersPage'
import AddOrderPage from '../pages/orders/AddOrderPage'
import AuthenticatedRoute, { LoginRoutePath } from './authenticated-route'

const Routes = ({ authorizationPage }) => (
  <Switch>
    <Redirect exact path="/" to="/inways-and-outways" />
    <Route path={LoginRoutePath} component={authorizationPage} />

    <AuthenticatedRoute
      path="/inways-and-outways/:name?"
      component={InwaysOverview}
    />
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
    <AuthenticatedRoute path="/finances" component={FinancePage} />
    <AuthenticatedRoute path="/audit-log" component={AuditLogPage} />
    <AuthenticatedRoute path="/orders/add-order" component={AddOrderPage} />

    <Redirect exact path="/orders" to="/orders/outgoing" />
    <AuthenticatedRoute
      path="/orders/:type(outgoing|incoming)"
      component={OrdersPage}
    />
    <AuthenticatedRoute
      path="/orders/outgoing/:delegatee/:reference"
      component={OrdersPage}
    />

    <AuthenticatedRoute path="/settings" component={SettingsPage} />

    <AuthenticatedRoute path="*" component={NotFoundPage} />
  </Switch>
)

Routes.propTypes = {
  authorizationPage: elementType,
}

Routes.defaultProps = {
  authorizationPage: LoginOIDCPage,
}

export default Routes
