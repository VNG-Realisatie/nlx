// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { array } from 'prop-types'
import { Collapsible, Table, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconCheckboxMultiple } from '../../../../../icons'
import AccessGrantRow from './AccessGrantRow'

const AccessGrantSection = ({ accessGrants }) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)

  const revokeHandler = async (accessGrant) => {
    await accessGrant.revoke()

    if (!accessGrant.error) {
      showToast({
        title: t('Access revoked'),
        variant: 'success',
      })
    } else {
      showToast({
        title: t('Failed to revoke access grant'),
        body: t('Please try again'),
        variant: 'error',
      })
      console.error(accessGrant.error)
    }
  }

  return (
    <Collapsible
      title={
        <DetailHeading data-testid="service-accessgrants">
          <IconCheckboxMultiple />
          {t('Organizations with access')}
          <Amount value={accessGrants.length} />
        </DetailHeading>
      }
      ariaLabel={t('Organizations with access')}
    >
      <StyledCollapsibleBody>
        <Table data-testid="service-accessgrant-list">
          <tbody>
            {accessGrants.length ? (
              accessGrants.map((accessGrant) => (
                <AccessGrantRow
                  key={accessGrant.id}
                  accessGrant={accessGrant}
                  revokeHandler={revokeHandler}
                />
              ))
            ) : (
              <Table.Tr data-testid="service-no-accessgrants">
                <Table.Td>
                  <StyledCollapsibleEmptyBody>
                    {t('There are no organizations with access')}
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
AccessGrantSection.defaultProps = {
  accessGrants: [],
}

export default AccessGrantSection
