// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { object, func } from 'prop-types'
import { Table, Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useConfirmationModal } from '../../../../../../components/ConfirmationModal'
import { IconRevoke } from '../../../../../../icons'
import { TdActions } from './index.styles'

const AccessGrantRow = ({ accessGrant, revokeHandler }) => {
  const { t } = useTranslation()

  const [ConfirmRevokeModal, confirmRevoke] = useConfirmationModal({
    okText: t('Revoke'),
    children: (
      <p>
        {t(
          'Access will be revoked for the serviceName service from organizationName',
          {
            organizationName: accessGrant.organizationName,
            serviceName: accessGrant.serviceName,
          },
        )}
      </p>
    ),
  })

  const handleRevoke = async (evt) => {
    evt.stopPropagation()

    if (await confirmRevoke()) {
      revokeHandler(accessGrant)
    }
  }

  return (
    <Table.Tr data-testid="service-accessgrants" key={accessGrant.id}>
      <Table.Td>{accessGrant.organizationName}</Table.Td>
      <TdActions>
        <Button
          size="small"
          variant="link"
          onClick={handleRevoke}
          title={t('Revoke')}
        >
          <IconRevoke inline />
          {t('Revoke')}
        </Button>

        <ConfirmRevokeModal />
      </TdActions>
    </Table.Tr>
  )
}

AccessGrantRow.propTypes = {
  accessGrant: object,
  revokeHandler: func,
}

export default AccessGrantRow
