// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, instanceOf, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import Table from '../../../../../../components/Table'
import EmptyContentMessage from '../../../../../../components/EmptyContentMessage'
import InwayModel from '../../../../../../stores/models/InwayModel'
import InwayRow from './InwayRow'

const Inways = ({ inways, selectedInwayName }) => {
  const { t } = useTranslation()

  return inways.length === 0 ? (
    <EmptyContentMessage>
      {t('There are no inways registered yet')}
    </EmptyContentMessage>
  ) : (
    <Table withLinks data-testid="inways-list" role="grid">
      <thead>
        <Table.TrHead>
          <Table.Th>{t('Type')}</Table.Th>
          <Table.Th>{t('Name')}</Table.Th>
          <Table.Th>{t('Hostname')}</Table.Th>
          <Table.Th>{t('Self address')}</Table.Th>
          <Table.Th>{t('Services')}</Table.Th>
          <Table.Th>{t('Version')}</Table.Th>
        </Table.TrHead>
      </thead>
      <tbody>
        {inways.map((inway, i) => (
          <InwayRow
            inway={inway}
            key={i}
            selected={inway.name === selectedInwayName}
          />
        ))}
      </tbody>
    </Table>
  )
}

Inways.propTypes = {
  inways: arrayOf(instanceOf(InwayModel)),
  selectedInwayName: string,
}

Inways.defaultProps = {
  inways: [],
}

export default Inways
