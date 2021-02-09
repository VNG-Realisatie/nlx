// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { ManagementApi } from '../../api'
import { RootStore, StoreProvider } from '../../stores'
import { renderWithProviders, waitFor } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import { AUDIT_LOG_ACTION_LOGIN_SUCCESS } from '../../stores/models/AuditLogModel'
import AuditLogPage from './index'

test('fetching the audit logs', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListAuditLogs = jest.fn().mockResolvedValue({
    auditLogs: [
      {
        id: 42,
        action: AUDIT_LOG_ACTION_LOGIN_SUCCESS,
      },
    ],
  })

  const store = new RootStore({
    managementApiClient,
  })

  const { getByRole, findAllByTestId } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <AuditLogPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()

  const auditLogElements = await findAllByTestId('audit-log-record')
  expect(auditLogElements).toHaveLength(1)
})

test('failed to load audit logs', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListAuditLogs = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const store = new RootStore({
    managementApiClient,
  })

  const { queryByRole, getByTestId, findByText } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <AuditLogPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  await waitFor(() => {
    expect(queryByRole('progressbar')).not.toBeInTheDocument()
  })

  expect(() => getByTestId('audit-log-record')).toThrow()

  expect(await findByText(/^Failed to load audit logs$/)).toBeInTheDocument()
  expect(await findByText(/^arbitrary error$/)).toBeInTheDocument()
})
