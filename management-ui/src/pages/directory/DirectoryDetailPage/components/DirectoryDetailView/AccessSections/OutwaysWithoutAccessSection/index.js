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
import State from '../components/State'
import { useOutwayStore } from '../../../../../../../hooks/use-stores'
import Table from '../../../../../../../components/Table'
import { OutwayName, Outways } from '../components/index.styles'
import Header from '../components/Header'

const OutwaysWithoutAccessSection = ({
  service,
  requestAccessHandler,
  retryRequestAccessHandler,
}) => {
  const { t } = useTranslation()
  const outwayStore = useOutwayStore()

  const publicKeyFingerPrintsWithoutAccess =
    outwayStore.publicKeyFingerprints.filter(
      (publicKeyFingerprint) => !service.hasAccess(publicKeyFingerprint),
    )

  return publicKeyFingerPrintsWithoutAccess.length < 1 ? (
    <Header title={t('Outways without access')} label={t('None')} />
  ) : (
    <Collapsible
      title={<Header title={t('Outways without access')} />}
      ariaLabel={t('Outways without access')}
    >
      <StyledCollapsibleBody>
        <Table>
          <tbody>
            {publicKeyFingerPrintsWithoutAccess.map((publicKeyFingerprint) => {
              const onRequestAccess = () => {
                requestAccessHandler(publicKeyFingerprint)
              }

              const onRetryRequestAccess = () => {
                retryRequestAccessHandler(publicKeyFingerprint)
              }

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
                      onRequestAccess={onRequestAccess}
                      onRetryRequestAccess={onRetryRequestAccess}
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

OutwaysWithoutAccessSection.propTypes = {
  service: instanceOf(DirectoryServiceModel),
  requestAccessHandler: func,
  retryRequestAccessHandler: func,
}

export default observer(OutwaysWithoutAccessSection)
