// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled, { css } from 'styled-components'

export const StyledTextWithEllipsis = styled.span`
  ${(p) =>
    p.title
      ? css`
          max-width: 20rem;
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
        `
      : null};
`
