// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Link } from 'react-router-dom'

export const StyledLink = styled(Link)`
  margin-left: ${(p) => p.theme.tokens.spacing06};
  text-decoration: none;
  font-weight: ${(p) => p.theme.tokens.fontWeightSemiBold};
  white-space: nowrap;
`
