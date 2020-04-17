// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import {
  StyledSpinner,
  StyledBulletContainer,
  StyledBullet,
} from './index.styles'

const Spinner = () => (
  <StyledSpinner>
    <StyledBulletContainer>
      {Array.from({ length: 8 }).map((value, i) => (
        <StyledBullet data-test="bullet" key={i} />
      ))}
    </StyledBulletContainer>
  </StyledSpinner>
)

export default Spinner
