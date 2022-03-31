// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { object, func } from 'prop-types'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import Table from '../../../../../../../components/Table'
import { useConfirmationModal } from '../../../../../../../components/ConfirmationModal'
import { IconRevoke } from '../../../../../../../icons'
import AccessDetails from '../../../AccessRequestSection/CollapsibleBody/IncomingAccessRequestRow/components/AccessDetails'
import { TdActions } from './index.styles'

const AccessGrantRow = ({ accessGrant, revokeHandler }) => {
  const { t } = useTranslation()

  const [ConfirmRevokeModal, confirmRevoke] = useConfirmationModal({
    title: t('Revoke access'),
    okText: t('Revoke'),
    children: (
      <AccessDetails
        subTitle={t(
          'Revoking this access grant will revoke access for the organization to the service. Are you sure?',
        )}
        organization={accessGrant.organization}
        serviceName={accessGrant.serviceName}
        publicKeyFingerprint={accessGrant.publicKeyFingerprint}
      />
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
      <Table.Td>
        {accessGrant.organization.name} <br />
        <small>
          {t('OIN serialNumber', {
            serialNumber: accessGrant.organization.serialNumber,
          })}
        </small>
        <br />
        <small>
          {t('Public Key Fingerprint publicKeyFingerprint', {
            publicKeyFingerprint: accessGrant.publicKeyFingerprint,
          })}
        </small>
      </Table.Td>
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
