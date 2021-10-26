// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, instanceOf, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Route } from 'react-router-dom'
import Table from '../../../../components/Table'
import EmptyContentMessage from '../../../../components/EmptyContentMessage'
import OutwayModel from '../../../../stores/models/OutwayModel'
import OutwayDetailPage from '../../OutwayDetailPage'
import { useOutwayStore } from '../../../../hooks/use-stores'
import OutwayRow from './OutwayRow'

const Outways = ({ outways, selectedOutwayName }) => {
  const { t } = useTranslation()
  const outwayStore = useOutwayStore()

  return outways.length === 0 ? (
    <EmptyContentMessage>
      {t('There are no outways registered yet')}
    </EmptyContentMessage>
  ) : (
    <>
      <Table withLinks data-testid="outways-list" role="grid">
        <thead>
          <Table.TrHead>
            <Table.Th>{t('Type')}</Table.Th>
            <Table.Th>{t('Name')}</Table.Th>
            <Table.Th>{t('Version')}</Table.Th>
          </Table.TrHead>
        </thead>
        <tbody>
          {outways.map((outway, i) => (
            <OutwayRow
              outway={outway}
              key={i}
              selected={outway.name === selectedOutwayName}
            />
          ))}
        </tbody>
      </Table>

      <Route
        path="/inways-and-outways/outways/:name"
        render={({ match }) => {
          const outway = outwayStore.getByName(match.params.name)

          return (
            <OutwayDetailPage
              parentUrl="/inways-and-outways/outways"
              outway={outway}
            />
          )
        }}
      />
    </>
  )
}

Outways.propTypes = {
  outways: arrayOf(instanceOf(OutwayModel)),
  selectedOutwayName: string,
}

Outways.defaultProps = {
  outways: [],
}

export default Outways
