// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import { mediaQueries } from '@commonground/design-system'

export default styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  text-align: left;

  ${mediaQueries.mdUp`
    flex-direction: row;

    a {
      margin-right: ${(p) => p.theme.tokens.spacing06};
    }
  `}
`
