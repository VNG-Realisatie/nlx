// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, func } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Collapsible } from '@commonground/design-system'
import {
  DetailHeading,
  SectionGroup,
  StyledCollapsibleBody,
} from '../../../../components/DetailView'
import { IconCertificate } from '../../../../icons'
import {
  StyledActionsBar,
  StyledRemoveButton,
} from '../../../services/ServiceDetailPage/ServiceDetailView/index.styles'
import { useConfirmationModal } from '../../../../components/ConfirmationModal'
import {
  SubHeader,
  StyledIconOutway,
  StyledSpecList,
  StyledCode,
} from './index.styles'

const OutwayDetails = ({ outway, removeHandler }) => {
  const { t } = useTranslation()
  const { ipAddress, publicKeyPEM, version } = outway

  const [ConfirmRemoveModal, confirmRemove] = useConfirmationModal({
    okText: t('Remove'),
    children: <p>{t('Do you want to remove the outway?')}</p>,
  })

  const handleRemove = async () => {
    if (await confirmRemove()) {
      removeHandler()
    }
  }

  return (
    <>
      <SubHeader data-testid="gateway-type">
        <StyledIconOutway inline />
        outway
      </SubHeader>

      <StyledActionsBar>
        <StyledRemoveButton title={t('Remove outway')} onClick={handleRemove} />
      </StyledActionsBar>

      <StyledSpecList data-testid="outway-specs" alignValuesRight>
        <StyledSpecList.Item title={t('IP-address')} value={ipAddress} />
        <StyledSpecList.Item title={t('Version')} value={version} />
      </StyledSpecList>

      <SectionGroup>
        <Collapsible
          title={
            <DetailHeading>
              <IconCertificate />
              {t('Certificate')}
            </DetailHeading>
          }
          ariaLabel={t('Certificate')}
          buttonLabels={{
            open: t('Open'),
            close: t('Close'),
          }}
        >
          <StyledCollapsibleBody>
            <StyledCode>{publicKeyPEM}</StyledCode>
          </StyledCollapsibleBody>
        </Collapsible>
      </SectionGroup>

      <ConfirmRemoveModal />
    </>
  )
}

OutwayDetails.propTypes = {
  outway: shape({
    name: string.isRequired,
    ipAddress: string,
    publicKeyPEM: string,
  }),
  removeHandler: func,
}

OutwayDetails.defaultProps = {}

export default observer(OutwayDetails)
