// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React, { useContext, useState } from 'react'
import { arrayOf, instanceOf, string } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Button, ToasterContext } from '@commonground/design-system'
import Table from '../../../../../../../../components/Table'
import { OutwayName, Outways } from '../../components/index.styles'
import AccessState from '../../components/AccessState'
import OutwayModel from '../../../../../../../../stores/models/OutwayModel'
import DirectoryServiceModel from '../../../../../../../../stores/models/DirectoryServiceModel'
import { IconClose } from '../../../../../../../../icons'
import { useConfirmationModal } from '../../../../../../../../components/ConfirmationModal'
import AccessDetails from './components/AccessDetails'
import { TdActions } from './index.styles'

const Row = ({ publicKeyFingerprint, outways, service }) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const [isLoading, setIsLoading] = useState(false)

  const { accessRequest, accessProof } =
    service.getAccessStateFor(publicKeyFingerprint)

  const [ConfirmTerminateModal, confirmTerminate] = useConfirmationModal({
    title: t('Terminate access'),
    okText: t('Terminate'),
    children: (
      <AccessDetails
        subTitle={t(
          'Terminating access will revoke access for the organization to the service. Are you sure?',
        )}
        organization={accessRequest.organization}
        serviceName={accessRequest.serviceName}
        publicKeyFingerprint={accessRequest.publicKeyFingerprint}
      />
    ),
  })

  const handleTerminate = async (evt) => {
    evt.stopPropagation()

    if (await confirmTerminate()) {
      try {
        setIsLoading(true)
        await service.terminateAccessRequest(accessRequest.publicKeyFingerprint)

        showToast({
          title: t('Access terminated'),
          variant: 'success',
        })
      } catch (err) {
        let message = err.message

        if (err.response && err.response.status === 403) {
          message = t(`You don't have the required permission.`)
        }

        showToast({
          title: t('Failed to terminate access'),
          body: message,
          variant: 'error',
        })
      }
    }
  }

  return (
    <Table.Tr>
      <Table.Td>
        <Outways>
          {outways.map((outway) => (
            <OutwayName key={outway.name}>{outway.name}</OutwayName>
          ))}
        </Outways>

        <AccessState
          isLoading={isLoading}
          accessRequest={accessRequest}
          accessProof={accessProof}
        />
      </Table.Td>

      <TdActions>
        <Button
          size="small"
          variant="link"
          onClick={handleTerminate}
          title={t(
            'Terminate access for Outways with public key fingerprint {{publicKeyFingerprint}}',
            {
              publicKeyFingerprint: accessRequest.publicKeyFingerprint,
            },
          )}
        >
          <IconClose inline />
          {t('Terminate')}
        </Button>

        <ConfirmTerminateModal />
      </TdActions>
    </Table.Tr>
  )
}

Row.propTypes = {
  publicKeyFingerprint: string,
  outways: arrayOf(instanceOf(OutwayModel)),
  service: instanceOf(DirectoryServiceModel),
}

export default observer(Row)
