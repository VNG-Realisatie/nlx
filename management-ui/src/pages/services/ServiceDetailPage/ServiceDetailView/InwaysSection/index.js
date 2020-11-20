// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, string } from 'prop-types'
import { Collapsible, Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import {
  DetailHeading,
  StyledCollapsibleBody,
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
        {inways.length ? (
          <Table data-testid="service-inways-list" role="grid" withLinks>
            <tbody>
              {inways.map((inway, i) => (
                <Table.Tr
                  key={i}
                  data-testid={`service-inway-${i}`}
                  to={`/inways/${inway}`}
                >
                  <Table.Td>
                    <StyledInwayName>{inway}</StyledInwayName>
                  </Table.Td>
                </Table.Tr>
              ))}
            </tbody>
          </Table>
        ) : (
          <small>{t('No inways have been added')}</small>
        )}
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
