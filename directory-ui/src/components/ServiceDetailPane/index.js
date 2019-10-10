// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import React from 'react'
import { func, string } from 'prop-types'
import { StyledServiceDetailPane, StyledHeader, StyledTitle, StyledSecondTitle, StyledSubtitle, StyledDl, StyledCloseButton } from './index.styles'
import CloseIcon from '../CloseIcon'

const ServiceDetailPane = ({ name, organizationName, contactEmail, closeHandler }) =>
  <StyledServiceDetailPane>
    <StyledHeader>
      <StyledTitle>
        {name}
      </StyledTitle>
      <StyledCloseButton onClick={() => closeHandler()}>
        <CloseIcon />
      </StyledCloseButton>
    </StyledHeader>

    <StyledSecondTitle>
      {organizationName}
    </StyledSecondTitle>

    {
      contactEmail ?
        (
          <>
            <StyledSubtitle>Contact</StyledSubtitle>
            <StyledDl>
              <dt>Email</dt>
              <dd>{ contactEmail }</dd>
            </StyledDl>
          </>
      ) : null
    }
  </StyledServiceDetailPane>

ServiceDetailPane.propTypes = {
  name: string,
  organizationName: string,
  contactEmail: string,
  closeHandler: func,
}

ServiceDetailPane.defaultProps = {
  closeHandler: () => {}
}

export default ServiceDetailPane

