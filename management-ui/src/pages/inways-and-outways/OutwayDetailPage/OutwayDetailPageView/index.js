// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'

import { SubHeader, StyledIconOutway, StyledSpecList } from './index.styles'

// Note: if outway- & outway details are interchangable, we can rename this to GatewayDetails
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
        <textarea disabled title={t('Public Key PEM')}>
          {publicKeyPEM}
        </textarea>
        <StyledSpecList.Item title={t('Version')} value={version} />
      </StyledSpecList>
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
