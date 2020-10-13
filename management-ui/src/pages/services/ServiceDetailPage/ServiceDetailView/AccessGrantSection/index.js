// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { array, func } from 'prop-types'
import { Collapsible, Table, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import ButtonWithIcon from '../../../../../components/ButtonWithIcon'
import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconCheckboxMultiple, IconRevoke } from '../../../../../icons'
import AccessGrantRepository from '../../../../../domain/access-grant-repository'
import { TdActions } from './index.styles'

const AccessGrantSection = ({ accessGrants, revokeAccessGrantHandler }) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)

  const handleRevokeGrantOnClick = (event, accessGrant) => {
    event.preventDefault()

    const confirmed = window.confirm(
      t(
        'Access will be revoked for the serviceName service from organizationName',
        {
          organizationName: accessGrant.organizationName,
          serviceName: accessGrant.serviceName,
        },
      ),
    )

    if (!confirmed) {
      return
    }

    revokeAccessGrant(accessGrant)
  }

  const revokeAccessGrant = async (accessGrant) => {
    await revokeAccessGrantHandler({
      organizationName: accessGrant.organizationName,
      serviceName: accessGrant.serviceName,
      accessGrantId: accessGrant.id,
    })

    showToast({
      title: t('Access revoked'),
      variant: 'success',
    })
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
                <Table.Tr
                  data-testid="service-accessgrants"
                  key={accessGrant.id}
                >
                  <Table.Td>{accessGrant.organizationName}</Table.Td>
                  <TdActions>
                    <ButtonWithIcon
                      size="small"
                      variant="link"
                      onClick={(e) => handleRevokeGrantOnClick(e, accessGrant)}
                    >
                      <IconRevoke />
                      {t('Revoke')}
                    </ButtonWithIcon>
                  </TdActions>
                </Table.Tr>
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
  revokeAccessGrantHandler: func,
}
AccessGrantSection.defaultProps = {
  accessGrants: [],
  revokeAccessGrantHandler: AccessGrantRepository.revokeAccessGrant,
}

export default AccessGrantSection
