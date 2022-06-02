// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf, shape, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import Table from '../../../components/Table'
import { DIRECTION_IN } from '../../../stores/models/TransactionLogModel'

const TransactionLogRow = ({ transactionLog, ...props }) => {
  const { t } = useTranslation()
  return (
    <Table.Tr data-testid="transaction-log-record" {...props}>
      <Table.Td>
        {t('Transaction log created at', { date: transactionLog.createdAt })}
      </Table.Td>
      <Table.Td>{transactionLog.transactionID}</Table.Td>
      <Table.Td>
        {transactionLog.direction === DIRECTION_IN
          ? t('Incoming from')
          : t('Outgoing to')}
      </Table.Td>
      <Table.Td>
        {transactionLog.order && transactionLog.order.delegator
          ? `${transactionLog.source.serialNumber} ${t('On behalf of')} ${
              transactionLog.order.delegator
            }`
          : transactionLog.source.serialNumber}
      </Table.Td>
      <Table.Td>{transactionLog.serviceName}</Table.Td>
    </Table.Tr>
  )
}

TransactionLogRow.propTypes = {
  transactionLog: shape({
    direction: string,
    source: shape({
      serialNumber: string,
    }),
    createdAt: instanceOf(Date),
    order: shape({
      delegator: string,
      reference: string,
    }),
  }),
}

export default TransactionLogRow
