// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { number } from 'prop-types'
import { Collapsible } from '@commonground/design-system'
import Table from '../DirectoryTable/Table'
import { StyledCollapsibleBody } from '../DetailView'
import CollapsibleHeader from './CollapsibleHeader'
import { TdPrice } from './index.styles'
import NoCosts from './NoCosts'

const isServiceFreeToUse = (oneTimeCosts, monthlyCosts, requestCosts) =>
  (oneTimeCosts === 0 || oneTimeCosts === undefined) &&
  (monthlyCosts === 0 || monthlyCosts === undefined) &&
  (requestCosts === 0 || requestCosts === undefined)

const costFormatter = new Intl.NumberFormat('nl-NL', {
  style: 'currency',
  currency: 'EUR',
})

const CostsSection = ({ oneTimeCosts, monthlyCosts, requestCosts }) => {
  const costSummaryParts = []

  if (oneTimeCosts > 0) {
    costSummaryParts.push('one time')
  }

  if (monthlyCosts > 0) {
    costSummaryParts.push('monthly')
  }

  if (requestCosts > 0) {
    costSummaryParts.push('per request')
  }

  const costSummary =
    costSummaryParts.length <= 1
      ? costSummaryParts.toString()
      : costSummaryParts.length === 2
      ? costSummaryParts.join(` en `)
      : `${costSummaryParts.slice(0, -1).join(', ')} en ${
          costSummaryParts[costSummaryParts.length - 1]
        }`

  return isServiceFreeToUse(oneTimeCosts, monthlyCosts, requestCosts) ? (
    <NoCosts />
  ) : (
    <Collapsible
      title={<CollapsibleHeader label={costSummary} />}
      ariaLabel="kosten"
    >
      <StyledCollapsibleBody>
        <Table>
          <tbody>
            {oneTimeCosts > 0 ? (
              <Table.Tr>
                <Table.Td>Eenmalige kosten</Table.Td>
                <TdPrice>{costFormatter.format(oneTimeCosts)}</TdPrice>
              </Table.Tr>
            ) : null}

            {monthlyCosts > 0 ? (
              <Table.Tr>
                <Table.Td>Maandelijkse kosten</Table.Td>
                <TdPrice>{costFormatter.format(monthlyCosts)}</TdPrice>
              </Table.Tr>
            ) : null}

            {requestCosts > 0 ? (
              <Table.Tr>
                <Table.Td>Kosten per aanvraag</Table.Td>
                <TdPrice>{costFormatter.format(requestCosts)}</TdPrice>
              </Table.Tr>
            ) : null}
          </tbody>
        </Table>
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

CostsSection.propTypes = {
  oneTimeCosts: number,
  monthlyCosts: number,
  requestCosts: number,
}

export default CostsSection
