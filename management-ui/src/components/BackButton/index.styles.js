// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Link } from 'react-router-dom'
import { IconChevron } from '../../icons'

export const StyledBackButton = styled(Link)`
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  line-height: 1rem;

  &:hover {
    svg {
      fill: ${(p) => p.theme.colorTextLinkHover};
    }
  }
`

export const StyledIconChevron = styled(IconChevron)`
  fill: ${(p) => p.theme.colorTextLink};
  width: ${(p) => p.theme.tokens.spacing05};
  height: ${(p) => p.theme.tokens.spacing05};
  margin: 0 ${(p) => p.theme.tokens.spacing04} 0 0;
  transform: rotate(-90deg);
`

export const StyledTitle = styled.h1`
  margin-bottom: ${(p) => p.theme.tokens.spacing09};
`
