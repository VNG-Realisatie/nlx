// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../components/PageTemplate'
import LoadingMessage from '../../components/LoadingMessage'
import useTransactionLogStore from '../../hooks/use-stores'
import Table from '../../components/Table'
import TransactionLogRow from './TransactionLogRow'

const TransactionLogPage = () => {
  const { t } = useTranslation()
  const { transactionLogStore } = useTransactionLogStore()
  const [error, setError] = useState()

  useEffect(() => {
    const fetchData = async () => {
      try {
        await transactionLogStore.fetchAll()
      } catch (err) {
        setError('something went wrong')
      }
    }

    fetchData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Transaction log')}
        description={t('Transaction logs of your inway and outways')}
      />

      {transactionLogStore.isLoading ? (
        <LoadingMessage />
      ) : error ? (
        <Alert
          variant="error"
          data-testid="error-message"
          title={t('Failed to load the transaction logs')}
        >
          {error}
        </Alert>
      ) : (
        <Table withLinks data-testid="transaction-log-list" role="grid">
          <thead>
            <Table.TrHead>
              <Table.Th>{t('Time')}</Table.Th>
              <Table.Th>{t('Direction')}</Table.Th>
              <Table.Th>{t('Organization')}</Table.Th>
              <Table.Th>{t('Service')}</Table.Th>
            </Table.TrHead>
          </thead>
          <tbody>
            {transactionLogStore.transactionLogs.map((transactionLog, i) => (
              <TransactionLogRow transactionLog={transactionLog} key={i} />
            ))}
          </tbody>
        </Table>
      )}
    </PageTemplate>
  )
}

export default observer(TransactionLogPage)
