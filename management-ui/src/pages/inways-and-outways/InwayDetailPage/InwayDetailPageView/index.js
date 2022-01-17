// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Collapsible } from '@commonground/design-system'
import { func, instanceOf } from 'prop-types'
import Table from '../../../../components/Table'
import Amount from '../../../../components/Amount'
import {
  DetailHeading,
  StyledCollapsibleBody,
  SectionGroup,
} from '../../../../components/DetailView'
import { IconServices } from '../../../../icons'
import {
  StyledActionsBar,
  StyledRemoveButton,
} from '../../../services/ServiceDetailPage/ServiceDetailView/index.styles'
import { useConfirmationModal } from '../../../../components/ConfirmationModal'
import InwayModel from '../../../../stores/models/InwayModel'
import { SubHeader, StyledIconInway, StyledSpecList } from './index.styles'

const InwayDetailPageView = ({ inway, removeHandler }) => {
  const { t } = useTranslation()
  const { ipAddress, hostname, selfAddress, version, services } = inway

  const [ConfirmRemoveModal, confirmRemove] = useConfirmationModal({
    okText: t('Remove'),
    children: <p>{t('Do you want to remove the inway?')}</p>,
  })

  const handleRemove = async () => {
    if (await confirmRemove()) {
      removeHandler()
    }
  }

  return (
    <>
      <SubHeader data-testid="gateway-type">
        <StyledIconInway inline />
        inway
      </SubHeader>

      <StyledActionsBar>
        <StyledRemoveButton
          title={t('Remove service')}
          onClick={handleRemove}
        />
      </StyledActionsBar>

      <StyledSpecList data-testid="inway-specs" alignValuesRight>
        <StyledSpecList.Item title={t('IP-address')} value={ipAddress} />
        <StyledSpecList.Item title={t('Hostname')} value={hostname} />
        <StyledSpecList.Item title={t('Self address')} value={selfAddress} />
        <StyledSpecList.Item title={t('Version')} value={version} />
      </StyledSpecList>

      <SectionGroup>
        <Collapsible
          title={
            <DetailHeading data-testid="inway-services">
              <IconServices />
              {t('Connected services')}
              <Amount value={services.length} />
            </DetailHeading>
          }
          ariaLabel={t('Connected services')}
          buttonLabels={{
            open: t('Open'),
            close: t('Close'),
          }}
        >
          <StyledCollapsibleBody>
            {services.length ? (
              <Table data-testid="inway-services-list" role="grid" withLinks>
                <tbody>
                  {services.map(({ name }) => (
                    <Table.Tr name={name} key={name} to={`/services/${name}`}>
                      <Table.Td>{name}</Table.Td>
                    </Table.Tr>
                  ))}
                </tbody>
              </Table>
            ) : (
              <small>{t('No services have been connected')}</small>
            )}
          </StyledCollapsibleBody>
        </Collapsible>
      </SectionGroup>

      <ConfirmRemoveModal />
    </>
  )
}

InwayDetailPageView.propTypes = {
  inway: instanceOf(InwayModel),
  removeHandler: func,
}

export default observer(InwayDetailPageView)
