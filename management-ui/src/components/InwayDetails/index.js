// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, arrayOf } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Drawer } from '@commonground/design-system'

import Amount from '../Amount'
import Collapsible from '../Collapsible'
import {
  DetailHeading,
  StyledCollapsibleBody,
} from '../DetailViewStyles'
import SpecList from '../SpecList'
import { ReactComponent as IconServices } from './services.svg'
import { SubHeader, StyledIconInway, StyledSpecList } from './index.styles'

// Note: if inway- & outway details are interchangable, we can rename this to GatewayDetails
const InwayDetails = ({ inway }) => {
  const { t } = useTranslation()
  const { name, hostname, selfAddress, services } = inway

  return (
    <>
      <Drawer.Header
        title={name}
        closeButtonLabel={t('Close')}
        data-testid="service-name"
      />

      <Drawer.Content>
        <SubHeader data-testid="gateway-type">
          <StyledIconInway /> inway
        </SubHeader>

        <StyledSpecList data-testid="inway-specs" role="grid" alignValuesRight>
          <StyledSpecList.Item title={t('IP-address')} value="123.456.789.1" />
          <StyledSpecList.Item title={t('Hostname')} value={hostname} />
          <StyledSpecList.Item title={t('Self address')} value={selfAddress} />
        </StyledSpecList>

        <Collapsible
          title={
            <DetailHeading data-testid="inway-services">
              <IconServices />
              {t('Connected services')}
              <Amount value={services.length} />
            </DetailHeading>
          }
        >
          <StyledCollapsibleBody />
        </Collapsible>
      </Drawer.Content>
    </>
  )
}

InwayDetails.propTypes = {
  inway: shape({
    name: string.isRequired,
    hostname: string,
    selfAddress: string,
    version: string,
    services: arrayOf(string),
  }),
}

InwayDetails.defaultProps = {}

export default InwayDetails
