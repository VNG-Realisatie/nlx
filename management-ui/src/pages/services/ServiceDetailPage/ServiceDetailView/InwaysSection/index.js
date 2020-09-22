// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, string } from 'prop-types'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import Collapsible from '../../../../../components/Collapsible'
import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconInway } from '../../../../../icons'
import { StyledInwayName } from '../index.styles'

const InwaysSection = ({ inways }) => {
  const { t } = useTranslation()

  return (
    <Collapsible
      title={
        <DetailHeading data-testid="service-inways">
          <IconInway />
          {t('Inways')}
          <Amount value={inways.length} />
        </DetailHeading>
      }
      ariaLabel={t('Inways')}
    >
      <StyledCollapsibleBody>
        <Table data-testid="service-inways-list" role="grid" withLinks>
          <tbody>
            {inways.length ? (
              inways.map((inway, i) => (
                <Table.Tr
                  key={i}
                  data-testid={`service-inway-${i}`}
                  to={`/inways/${inway}`}
                >
                  <Table.Td>
                    <StyledInwayName>{inway}</StyledInwayName>
                  </Table.Td>
                </Table.Tr>
              ))
            ) : (
              <Table.Tr data-testid="service-no-inways">
                <Table.Td>
                  <StyledCollapsibleEmptyBody>
                    {t('No inways have been added')}
                  </StyledCollapsibleEmptyBody>
                </Table.Td>
              </Table.Tr>
            )}
          </tbody>
        </Table>
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

InwaysSection.propTypes = {
  inways: arrayOf(string),
}

InwaysSection.defaultProps = {
  inways: [],
}

export default InwaysSection
