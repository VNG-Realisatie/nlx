// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
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
  SubHeader,
  StyledIconOutway,
  StyledSpecList,
  StyledCode,
} from './index.styles'

const OutwayDetails = ({ outway }) => {
  const { t } = useTranslation()
  const { ipAddress, publicKeyPEM, version } = outway

  return (
    <>
      <SubHeader data-testid="gateway-type">
        <StyledIconOutway inline />
        outway
      </SubHeader>

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
    </>
  )
}

OutwayDetails.propTypes = {
  outway: shape({
    name: string.isRequired,
    ipAddress: string,
    publicKeyPEM: string,
  }),
}

OutwayDetails.defaultProps = {}

export default observer(OutwayDetails)
