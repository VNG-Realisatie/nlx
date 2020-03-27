// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { bool } from 'prop-types'
import styled, { css } from 'styled-components'
import React from 'react'
import { ReactComponent as ChevronDown } from './chevron-down.svg'

const FlippingIconChevron = ({ flipHorizontal, ...props }) => (
  <ChevronDown {...props} />
)

FlippingIconChevron.propTypes = {
  flipHorizontal: bool,
}

const FlippingChevron = styled(FlippingIconChevron)`
  fill: ${(p) => p.theme.colorText};
  transition: 150ms ease-in-out;

  ${(p) =>
    p.flipHorizontal
      ? css`
          transform: rotate(180deg);
        `
      : ''}
`

export default FlippingChevron
