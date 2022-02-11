// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import { Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import { StyledCollapsibleBody } from '../../../../../../components/DetailView'
import DirectoryServiceModel from '../../../../../../stores/models/DirectoryServiceModel'
import AccessSection from '../AccessSection'
import { useOutwayStore } from '../../../../../../hooks/use-stores'
import Table from '../../../../../../components/Table'
import Header from './Header'
import { OutwayName, OutwayNames } from './index.styles'

const CertificatesWithAccessSection = ({ service }) => {
  const { t } = useTranslation()
  const outwayStore = useOutwayStore()

  const publicKeyFingerPrintsWithAccess =
    outwayStore.publicKeyFingerprints.filter((publicKeyFingerprint) =>
      service.hasAccess(publicKeyFingerprint),
    )

  return publicKeyFingerPrintsWithAccess.length < 1 ? (
    <Header label={t('None')} />
  ) : (
    <Collapsible title={<Header />} ariaLabel={t('Outways with access')}>
      <StyledCollapsibleBody>
        <Table>
          <tbody>
            {publicKeyFingerPrintsWithAccess.map((publicKeyFingerprint) => {
              const { accessRequest, accessProof } =
                service.getAccessStateFor(publicKeyFingerprint)

              return (
                <Table.Tr key={publicKeyFingerprint}>
                  <Table.Td>
                    <OutwayNames data-testid="outway-names">
                      {outwayStore
                        .getByPublicKeyFingerprint(publicKeyFingerprint)
                        .map((outway) => (
                          <OutwayName key={outway.name}>
                            {outway.name}
                          </OutwayName>
                        ))
                        .reduce(
                          (accu, elem) =>
                            accu === null ? [elem] : [...accu, ', ', elem],
                          null,
                        )}
                    </OutwayNames>

                    <AccessSection
                      accessRequest={accessRequest}
                      accessProof={accessProof}
                    />
                  </Table.Td>
                </Table.Tr>
              )
            })}
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
