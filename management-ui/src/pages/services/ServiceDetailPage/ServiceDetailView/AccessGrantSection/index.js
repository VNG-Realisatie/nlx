// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { array } from 'prop-types'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import Collapsible from '../../../../../components/Collapsible'
import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconCheckboxMultiple } from '../../../../../icons'

const AccessGrantSection = ({ accessGrants }) => {
  const { t } = useTranslation()

  return (
    <Collapsible
      title={
        <DetailHeading data-testid="service-accessgrants">
          <IconCheckboxMultiple />
          {t('Access grants')}
          <Amount value={accessGrants.length} />
        </DetailHeading>
      }
      ariaLabel={t('Access grants')}
    >
      <StyledCollapsibleBody>
        <Table data-testid="service-accessgrant-list">
          <tbody>
            {accessGrants.length ? (
              accessGrants.map(({ id, organizationName }) => (
                <Table.Tr data-testid="service-no-accessgrants" key={id}>
                  <Table.Td>{organizationName}</Table.Td>
                </Table.Tr>
              ))
            ) : (
              <Table.Tr data-testid="service-no-accessgrants">
                <Table.Td>
                  <StyledCollapsibleEmptyBody>
                    {t('There are no access grants')}
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

AccessGrantSection.propTypes = {
  accessGrants: array,
}
AccessGrantSection.defaultProps = {}

export default AccessGrantSection
