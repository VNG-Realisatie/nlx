// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { bool, element, shape, string } from 'prop-types'
import { Route, Routes, Navigate } from 'react-router-dom'
import LoginOIDCPage from '../pages/LoginOIDCPage'
import ServicesPage from '../pages/services/ServicesPage'
import InwaysAndOutwaysPage from '../pages/inways-and-outways/InwaysAndOutwaysPage'
import AddServicePage from '../pages/services/AddServicePage'
import DirectoryPage from '../pages/directory/DirectoryPage'
import EditServicePage from '../pages/services/EditServicePage'
import FinancePage from '../pages/FinancePage'
import AuditLogPage from '../pages/AuditLogPage'
import TransactionLogPage from '../pages/TransactionLogPage'
import SettingsPage from '../pages/SettingsPage'
import NotFoundPage from '../pages/NotFoundPage'
import OrdersPage from '../pages/orders/OrdersPage'
import AddOrderPage from '../pages/orders/AddOrderPage'
import EditOrderPage from '../pages/orders/EditOrderPage'
import TermsOfService from '../pages/TermsOfServicePage'
import { ToSContextProvider } from '../tos-context'
import AuthenticatedRoute, { LoginRoutePath } from './authenticated-route'
import TermsOfServiceAcceptedRoute from './terms-of-service-accepted-route'

const AllRoutes = ({ authorizationPageElement, tos }) => (
  <Routes>
    <Route index element={<Navigate to="/inways-and-outways" />} />

    <Route path={LoginRoutePath} element={authorizationPageElement} />

    <Route
      path="*"
      element={
        <AuthenticatedRoute>
          <ToSContextProvider tos={tos}>
            <Routes>
              <Route path="/terms-of-service" element={<TermsOfService />} />

              <Route
                path="*"
                element={
                  <TermsOfServiceAcceptedRoute>
                    <Routes>
                      <Route
                        path="/inways-and-outways/*"
                        element={<InwaysAndOutwaysPage />}
                      />

                      <Route
                        path="/services/add-service"
                        element={<AddServicePage />}
                      />
                      <Route
                        path="/services/:name/edit-service"
                        element={<EditServicePage />}
                      />
                      <Route path="/services/*" element={<ServicesPage />} />

                      <Route path="/directory/*" element={<DirectoryPage />} />

                      <Route path="/finances" element={<FinancePage />} />

                      <Route path="/audit-log" element={<AuditLogPage />} />

                      <Route
                        path="/transaction-log"
                        element={<TransactionLogPage />}
                      />

                      <Route path="/settings/*" element={<SettingsPage />} />

                      <Route
                        path="/orders/add-order"
                        element={<AddOrderPage />}
                      />

                      <Route
                        path="/orders/outgoing/:delegateeSerialNumber/:reference/edit"
                        element={<EditOrderPage />}
                      />

                      <Route path="/orders/*" element={<OrdersPage />} />

                      <Route path="*" element={<NotFoundPage />} />
                    </Routes>
                  </TermsOfServiceAcceptedRoute>
                }
              />
            </Routes>
          </ToSContextProvider>
        </AuthenticatedRoute>
      }
    />
  </Routes>
)

AllRoutes.propTypes = {
  authorizationPageElement: element,
  tos: shape({
    enabled: bool,
    url: string,
    accepted: bool,
  }),
}

AllRoutes.defaultProps = {
  authorizationPage: <LoginOIDCPage />,
}

export default AllRoutes
