// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { func, instanceOf } from 'prop-types'
import { Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import { StyledCollapsibleBody } from '../../../../../../../components/DetailView'
import DirectoryServiceModel from '../../../../../../../stores/models/DirectoryServiceModel'
import { useOutwayStore } from '../../../../../../../hooks/use-stores'
import Table from '../../../../../../../components/Table'
import Header from '../components/Header'
import Row from './Row'

const OutwaysWithoutAccessSection = ({
  service,
  onShowConfirmRequestAccessModalHandler,
  onHideConfirmRequestAccessModalHandler,
}) => {
  const { t } = useTranslation()
  const outwayStore = useOutwayStore()

  const publicKeyFingerprintsWithoutAccess =
    outwayStore.publicKeyFingerprints.filter(
      (publicKeyFingerprint) => !service.hasAccess(publicKeyFingerprint),
    )

  return publicKeyFingerprintsWithoutAccess.length < 1 ? (
    <Header title={t('Outways without access')} label={t('None')} />
  ) : (
    <Collapsible
      title={<Header title={t('Outways without access')} />}
      ariaLabel={t('Outways without access')}
    >
      <StyledCollapsibleBody>
        <Table>
          <tbody>
            {publicKeyFingerprintsWithoutAccess.map((publicKeyFingerprint) => (
              <Row
                key={publicKeyFingerprint}
                publicKeyFingerprint={publicKeyFingerprint}
                service={service}
                outways={outwayStore.getByPublicKeyFingerprint(
                  publicKeyFingerprint,
                )}
                onShowConfirmRequestAccessModalHandler={
                  onShowConfirmRequestAccessModalHandler
                }
                onHideConfirmRequestAccessModalHandler={
                  onHideConfirmRequestAccessModalHandler
                }
              />
            ))}
          </tbody>
        </Table>
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

OutwaysWithoutAccessSection.propTypes = {
  service: instanceOf(DirectoryServiceModel),
  onShowConfirmRequestAccessModalHandler: func,
  onHideConfirmRequestAccessModalHandler: func,
}

export default observer(OutwaysWithoutAccessSection)
