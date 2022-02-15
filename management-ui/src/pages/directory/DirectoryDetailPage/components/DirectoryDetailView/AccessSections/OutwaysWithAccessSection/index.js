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
import State from '../components/State'
import { useOutwayStore } from '../../../../../../../hooks/use-stores'
import Table from '../../../../../../../components/Table'
import { OutwayName, Outways } from '../components/index.styles'
import Header from '../components/Header'

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
            {publicKeyFingerPrintsWithAccess.map((publicKeyFingerprint) => {
              const { accessRequest, accessProof } =
                service.getAccessStateFor(publicKeyFingerprint)

              return (
                <Table.Tr key={publicKeyFingerprint}>
                  <Table.Td>
                    <Outways>
                      {outwayStore
                        .getByPublicKeyFingerprint(publicKeyFingerprint)
                        .map((outway) => (
                          <OutwayName key={outway.name}>
                            {outway.name}
                          </OutwayName>
                        ))}
                    </Outways>

                    <State
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
