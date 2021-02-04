// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../components/PageTemplate'
import LoadingMessage from '../../components/LoadingMessage'
import useStores from '../../hooks/use-stores'

const AuditLogPage = () => {
  const { t } = useTranslation()
  const { auditLogStore } = useStores()
  const [error, setError] = useState()

  useEffect(() => {
    const fetchData = async () => {
      try {
        await auditLogStore.fetchAll()
      } catch (err) {
        setError(err.message)
      }
    }

    fetchData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Audit log')}
        description={t('History of all mutations within your NLX setup.')}
      />

      {auditLogStore.isLoading ? (
        <LoadingMessage />
      ) : error ? (
        <Alert
          variant="error"
          data-testid="error-message"
          title={t('Failed to load audit logs')}
        >
          {error}
        </Alert>
      ) : (
        <ul>
          {auditLogStore.auditLogs.map((auditLog, i) => (
            <li key={i} data-testid="audit-log-record">
              {auditLog.action}
            </li>
          ))}
        </ul>
      )}
    </PageTemplate>
  )
}

export default observer(AuditLogPage)
