// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import { useTranslation } from 'react-i18next'
import Table from '../../../components/Table'
import TransactionLogModel, {
  DIRECTION_IN,
} from '../../../stores/models/TransactionLogModel'

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
        {transactionLog.order &&
        transactionLog.order.delegator &&
        transactionLog.order.delegator.serialNumber
          ? `${transactionLog.source.name} ${t('On behalf of')} ${
              transactionLog.order.delegator.name
            }`
          : transactionLog.source.name}
      </Table.Td>
      <Table.Td>{transactionLog.serviceName}</Table.Td>
    </Table.Tr>
  )
}

TransactionLogRow.propTypes = {
  transactionLog: instanceOf(TransactionLogModel),
}

export default TransactionLogRow
