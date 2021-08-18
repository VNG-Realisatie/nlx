// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import styled from 'styled-components'
import { mediaQueries } from '@commonground/design-system'
import Input from '@commonground/design-system/dist/components/Form/TextInput/Input'

export const StyledInput = styled(Input)`
  width: 100%;
`

export const StyledFilters = styled.div`
  display: flex;
  flex-direction: column;
  margin: 0 auto;

  > * + * {
    margin-top: ${(p) => p.theme.tokens.spacing06};
  }

  ${mediaQueries.smDown`
    > * {
        width: 100%;
    }
  `}
  ${mediaQueries.mdUp`
    flex-direction: row;
    justify-content: space-between;
    align-items: center;

    > * + * {
      margin: 0 0 0 ${(p) => p.theme.tokens.spacing05};
    }
  `}
`
