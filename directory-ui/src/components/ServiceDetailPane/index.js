// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { func, string } from 'prop-types'
import CloseIcon from '../CloseIcon'
import {
  StyledServiceDetailPane,
  StyledHeader,
  StyledTitle,
  StyledSecondTitle,
  StyledSubtitle,
  StyledDl,
  StyledLink,
  StyledCloseButton,
} from './index.styles'

const ServiceDetailPane = ({
  name,
  organizationName,
  contactEmailAddress,
  documentationUrl,
  closeHandler,
}) => (
  <StyledServiceDetailPane data-test="service-detail-pane">
    <StyledHeader>
      <StyledTitle>{name}</StyledTitle>
      <StyledCloseButton onClick={() => closeHandler()}>
        <CloseIcon />
      </StyledCloseButton>
    </StyledHeader>

    <StyledSecondTitle>{organizationName}</StyledSecondTitle>

    {contactEmailAddress ? (
      <>
        <StyledSubtitle>Support</StyledSubtitle>
        <StyledDl>
          <dt>Email address</dt>
          <dd>
            <StyledLink
              href={'mailto:' + contactEmailAddress}
              data-test="email-address-link"
            >
              {contactEmailAddress}
            </StyledLink>
          </dd>
        </StyledDl>
      </>
    ) : null}

    {documentationUrl ? (
      <>
        <StyledSubtitle>Documentation</StyledSubtitle>
        <StyledDl>
          <dt>URL</dt>
          <dd>
            <StyledLink
              href={documentationUrl}
              target="_blank"
              rel="noopener noreferrer"
              data-test="documentation-link"
            >
              {documentationUrl}
            </StyledLink>
          </dd>
        </StyledDl>
      </>
    ) : null}
  </StyledServiceDetailPane>
)

ServiceDetailPane.propTypes = {
  name: string,
  organizationName: string,
  contactEmailAddress: string,
  documentationUrl: string,
  closeHandler: func,
}

ServiceDetailPane.defaultProps = {
  closeHandler: () => {},
}

export default ServiceDetailPane
