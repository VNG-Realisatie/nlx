// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import { Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import { StyledCollapsibleBody } from '../../../../../../../components/DetailView'
import DirectoryServiceModel from '../../../../../../../stores/models/DirectoryServiceModel'
import { useOutwayStore } from '../../../../../../../hooks/use-stores'
import Table from '../../../../../../../components/Table'
import Header from '../components/Header'
import Row from './Row'

const CertificatesWithAccessSection = ({ service }) => {
  const { t } = useTranslation()
  const outwayStore = useOutwayStore()

  const publicKeyFingerPrintsWithAccess =
    outwayStore.publicKeyFingerprints.filter((publicKeyFingerprint) =>
      service.hasAccess(publicKeyFingerprint),
    )

  return publicKeyFingerPrintsWithAccess.length < 1 ? (
    <Header title={t('Outways with access')} label={t('None')} />
  ) : (
    <Collapsible
      title={<Header title={t('Outways with access')} />}
      ariaLabel={t('Outways with access')}
    >
      <StyledCollapsibleBody>
        <Table>
          <tbody>
            {publicKeyFingerPrintsWithAccess.map((publicKeyFingerprint) => (
              <Row
                key={publicKeyFingerprint}
                publicKeyFingerprint={publicKeyFingerprint}
                service={service}
                outways={outwayStore.getByPublicKeyFingerprint(
                  publicKeyFingerprint,
                )}
              />
            ))}
          </tbody>
        </Table>
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

CertificatesWithAccessSection.propTypes = {
  service: instanceOf(DirectoryServiceModel),
}

export default observer(CertificatesWithAccessSection)
